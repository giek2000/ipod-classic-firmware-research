# iPod Classic 6.5G Rev A (120GB) - RetailOS 2.0 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 2.0 |
| **IPSW** | iPod_33.2.0.ipsw |
| **Device** | iPod Classic 6.5G Rev A (120GB) (2008, 120GB, Click Wheel, Cover Flow, Genius, CE-ATA HDD (First Release)) |
| **UpdaterFamilyID** | 33 |
| **Binary Size** | 10,509,968 bytes (10.02 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,507,920 bytes |
| **Total Strings (>=4)** | 71,427 |
| **Function Prologues** | 24,448 (ARM: 17,394, Thumb: 7,054) |
| **DRAM References** | 106,485 |
| **Peripheral Refs** | 7,230 |
| **Build** | N25BFirmwareWin-69 |
| **SoC** | S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Codename** | N25C |
| **DFU PID** | 0x1223 |
| **SHA-256** | `2cf1c158129f28751d742f34e8ed259d05c73f1e4bc03d75b89548fa78b2c9b9` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00095294 | `TSilverCntlr` | Known | Controller |
| 0x000952AC | `TCExtrasMenu` | Known | Controller |
| 0x000952C4 | `TCGameScreen` | Known | Controller |
| 0x000952DC | `TCGamesMenu` | Known | Controller |
| 0x000952F0 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x00095318 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00095340 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x0009536C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x00095390 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x000953B8 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x000953E0 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x00095408 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x00095430 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x00095458 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x00095488 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x000954B4 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x000954E4 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x0009550C | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x00095534 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x00095560 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x00095588 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x000955B0 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x000955E0 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x00095610 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x0009576C | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x0009579C | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x000957C4 | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x000957EC | `TCRentalNotification` | Known | Controller |
| 0x0009580C | `TCRentalInfo` | Known | Controller |
| 0x00095824 | `TCRentalConfirmDelete` | Known | Controller |
| 0x00095844 | `TCRentalDispatcher` | Known | Controller |
| 0x0009589C | `TSilverGlobalCntlr` | Known | Controller |
| 0x000958B8 | `TSilverTrainerCntlr` | Known | Controller |
| 0x000EC694 | `TCSlideshowLCD` | Known | Controller |
| 0x000EC6AC | `TCSlideshowTVOut` | Known | Controller |
| 0x000EC6C8 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x000EC6E8 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x001100EC | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x00110118 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x00110144 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x0011016C | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x00110198 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x001101C0 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x001101EC | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x00117D38 | `TCRemoteUI` | Known | Controller |
| 0x00117D4C | `TCUnsupported` | Known | Controller |
| 0x0011E1A4 | `TCSpeakers` | Known | Controller |
| 0x0011E1B8 | `TCEQSetting` | Known | Controller |
| 0x00147320 | `TCSportTimer` | Known | Controller |
| 0x00147338 | `TCSportTimerMenu` | Known | Controller |
| 0x00147354 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x00147378 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x00148724 | `TCVoiceMemos` | Known | Controller |
| 0x0014873C | `TCVoiceMemosMenu` | Known | Controller |
| 0x00148758 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x00148778 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x00148798 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x001487B8 | `TCVoiceMemosAlert` | Known | Controller |
| 0x00159EB4 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x00159EDC | `TCSettings_MainMenu` | Known | Controller |
| 0x00159EF8 | `TCSettings_MusicMenu` | Known | Controller |
| 0x00159F18 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x00159F38 | `TCSettings_Brightness` | Known | Controller |
| 0x00159F58 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x00159F7C | `TCSettings_EQ` | Known | Controller |
| 0x00159F94 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x00159FBC | `TCSettings_RadioRegions` | Known | Controller |
| 0x00159FDC | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0015A000 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x0015A024 | `TCDateTimeScreen` | Known | Controller |
| 0x0015A040 | `TCTimeZoneScreen` | Known | Controller |
| 0x0015A05C | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0015A084 | `TCFirstBoot` | Known | Controller |
| 0x00170428 | `TCDemoMode` | Known | Controller |
| 0x00198BE0 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x00198C00 | `TCAddressViewerDetails` | Known | Controller |
| 0x00198C20 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x00198C44 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x001C547C | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x001C54A0 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x001CCCC8 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x00261FA0 | `TC_LockDialog` | Known | Controller |
| 0x00261FB8 | `TC_LockScreen` | Known | Controller |
| 0x00261FD0 | `TC_LockediPod` | Known | Controller |
| 0x00261FE8 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x0026200C | `TCLockChosenDispatcher` | Known | Controller |
| 0x00267B90 | `TCClock` | Known | Controller |
| 0x00267BA0 | `TCClockCityMenu` | Known | Controller |
| 0x00267BB8 | `TCClockRegionMenu` | Known | Controller |
| 0x00267BD4 | `TCAlarmMenu` | Known | Controller |
| 0x00267BE8 | `TCSleepTimerMenu` | Known | Controller |
| 0x00267C04 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x00267C24 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x00267C4C | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x00267C70 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00267C94 | `TCAlarmDatePicker` | Known | Controller |
| 0x00267CB0 | `TCAlarmTriggered` | Known | Controller |
| 0x0026EBD4 | `TCNotesDispatcher` | Known | Controller |
| 0x0026EBF0 | `TCNotesLoading` | Known | Controller |
| 0x0026EC08 | `TCNotesList` | Known | Controller |
| 0x0026EC1C | `TCNotesContents` | Known | Controller |
| 0x003DCDD8 | `TCAlarmTriggered` | Known | Controller |
| 0x003DCDEC | `TSilverCntlr` | Known | Controller |
| 0x003DCE0C | `TCClock` | Known | Controller |
| 0x003DCE14 | `TCClockRegionMenu` | Known | Controller |
| 0x003DCE28 | `TCClockCityMenu` | Known | Controller |
| 0x003DCE38 | `TCAlarmMenu` | Known | Controller |
| 0x003DCE44 | `TCSleepTimerMenu` | Known | Controller |
| 0x003DCE58 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003DCE70 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003DCE90 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003DCEAC | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003DCEC8 | `TCAlarmDatePicker` | Known | Controller |
| 0x003DCF00 | `TSilverCntlr` | Known | Controller |
| 0x003DCF20 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003DD0B0 | `TSilverCntlr` | Known | Controller |
| 0x003DD0D0 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x003DD0F0 | `TCSettings_Brightness` | Known | Controller |
| 0x003DD108 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x003DD124 | `TCSettings_AudiobookSettings` | Known | Controller |
| 0x003DD144 | `TCSettings_RadioRegions` | Known | Controller |
| 0x003DD15C | `TCSettings_EQ` | Known | Controller |
| 0x003DD16C | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x003DD188 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x003DD1A8 | `TCFirstBoot` | Known | Controller |
| 0x003DD1B4 | `TCSettings_MainMenu` | Known | Controller |
| 0x003DD1C8 | `TCSettings_MusicMenu` | Known | Controller |
| 0x003DD1E0 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x003DD1F8 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x003DD214 | `TCDateTimeScreen` | Known | Controller |
| 0x003DD228 | `TCTimeZoneScreen` | Known | Controller |
| 0x003E42F4 | `TSilverCntlr` | Known | Controller |
| 0x003E4314 | `TCClock` | Known | Controller |
| 0x003E431C | `TCClockRegionMenu` | Known | Controller |
| 0x003E4330 | `TCClockCityMenu` | Known | Controller |
| 0x003E4340 | `TCAlarmMenu` | Known | Controller |
| 0x003E434C | `TCSleepTimerMenu` | Known | Controller |
| 0x003E4360 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003E43D8 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003E43F8 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003E4414 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003E4448 | `TCAlarmDatePicker` | Known | Controller |
| 0x003E445C | `TCAlarmTriggered` | Known | Controller |
| 0x003E5EDC | `TSilverCntlr` | Known | Controller |
| 0x003E5EFC | `TC_LockDialog` | Known | Controller |
| 0x003E5F0C | `TC_LockScreen` | Known | Controller |
| 0x003E5F1C | `TC_LockediPod` | Known | Controller |
| 0x003E5F2C | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003E5F48 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003E5F60 | `TSilverCntlr` | Known | Controller |
| 0x003E616C | `TSilverCntlr` | Known | Controller |
| 0x003E6188 | `TSilverCntlr` | Known | Controller |
| 0x003E61EC | `TSilverCntlr` | Known | Controller |
| 0x003E620C | `TCNotesDispatcher` | Known | Controller |
| 0x003E6220 | `TCNotesLoading` | Known | Controller |
| 0x003E6230 | `TCNotesBase` | Known | Controller |
| 0x003E623C | `TCNotesList` | Known | Controller |
| 0x003E6248 | `TCNotesContents` | Known | Controller |
| 0x003E6258 | `TSilverCntlr` | Known | Controller |
| 0x003E6278 | `TCRemoteUI` | Known | Controller |
| 0x003E6284 | `TCUnsupported` | Known | Controller |
| 0x003E6294 | `TSilverCntlr` | Known | Controller |
| 0x003E62F8 | `TSilverCntlr` | Known | Controller |
| 0x003E6318 | `TCSportTimer` | Known | Controller |
| 0x003E6328 | `TCSportTimerMenu` | Known | Controller |
| 0x003E633C | `TCSportTimerSessionScreen` | Known | Controller |
| 0x003E6358 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x003E6388 | `TSilverCntlr` | Known | Controller |
| 0x003E64B0 | `TSilverCntlr` | Known | Controller |
| 0x003E64D0 | `TCDemoMode` | Known | Controller |
| 0x003E64DC | `TCClock` | Known | Controller |
| 0x003E64E4 | `TCClockRegionMenu` | Known | Controller |
| 0x003E64F8 | `TCClockCityMenu` | Known | Controller |
| 0x003E6508 | `TCAlarmMenu` | Known | Controller |
| 0x003E6514 | `TCSleepTimerMenu` | Known | Controller |
| 0x003E6528 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x003E6540 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x003E6560 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x003E657C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x003E6598 | `TCAlarmDatePicker` | Known | Controller |
| 0x003E65AC | `TCAlarmTriggered` | Known | Controller |
| 0x003E65CC | `TSilverCntlr` | Known | Controller |
| 0x003E65E8 | `TSilverCntlr` | Known | Controller |
| 0x003E65F8 | `TSilverCntlr` | Known | Controller |
| 0x003E6618 | `TCVoiceMemos` | Known | Controller |
| 0x003E6628 | `TCVoiceMemosMenu` | Known | Controller |
| 0x003E663C | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x003E6654 | `TCVoiceMemosAlert` | Known | Controller |
| 0x003E6668 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x003E6680 | `TCVoiceMemosPlayback` | Known | Controller |
| 0x003E66A0 | `TSilverCntlr` | Known | Controller |
| 0x003E6700 | `TSilverCntlr` | Known | Controller |
| 0x003E676C | `TSilverCntlr` | Known | Controller |
| 0x003E7AB4 | `TSilverCntlr` | Known | Controller |
| 0x003E7BC0 | `TSilverCntlr` | Known | Controller |
| 0x003F043C | `TSilverCntlr` | Known | Controller |
| 0x003F045C | `TCAddressViewerMainMenu` | Known | Controller |
| 0x003F0474 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x003F0490 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x003F04B0 | `TCAddressViewerDetails` | Known | Controller |
| 0x003F04C8 | `TSilverCntlr` | Known | Controller |
| 0x003F04E8 | `TSilverCalendarCntlr_Base` | Known | Controller |
| 0x003F0504 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x003F0528 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x003F054C | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x003F056C | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x003F0590 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x003F05B0 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x003F05D4 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x003F07AC | `TSilverCntlr` | Known | Controller |
| 0x003F07CC | `TC_LockDialog` | Known | Controller |
| 0x003F07DC | `TC_LockScreen` | Known | Controller |
| 0x003F07EC | `TC_LockediPod` | Known | Controller |
| 0x003F07FC | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x003F0820 | `TCLockChosenDispatcher` | Known | Controller |
| 0x003F0940 | `TSilverCntlr` | Known | Controller |
| 0x003F0A74 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003F0A90 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003F0AB0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003F0AD0 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003F0AF8 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003F0B1C | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003F0B44 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003F0B64 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003F0B84 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003F0BA4 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003F0BC4 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003F0BEC | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003F0C14 | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x003F0C34 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003F0C54 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003F0C74 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003F0C98 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003F0CB8 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003F0CDC | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003F0D04 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003F0D30 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003F0D50 | `TCRentalNotification` | Known | Controller |
| 0x003F0D68 | `TCRentalInfo` | Known | Controller |
| 0x003F0D78 | `TCRentalConfirmDelete` | Known | Controller |
| 0x003F0D90 | `TCRentalDispatcher` | Known | Controller |
| 0x003F1680 | `TSilverCntlr` | Known | Controller |
| 0x003F1744 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003F1760 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003F1780 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003F17A0 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003F17C8 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003F17EC | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003F1814 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003F1834 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003F1854 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003F1874 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003F1894 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003F18BC | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003F18E4 | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x003F192C | `TCSlideshowTVOut` | Known | Controller |
| 0x003F1940 | `TCSlideshowLCD` | Known | Controller |
| 0x003F1950 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x003F1968 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x003F1988 | `TSilverCntlr` | Known | Controller |
| 0x003F19B4 | `TSilverCntlr` | Known | Controller |
| 0x003F19D4 | `TCUnsupported` | Known | Controller |
| 0x003F19F4 | `TSilverCntlr` | Known | Controller |
| 0x003F1A34 | `TSilverCntlr` | Known | Controller |
| 0x003F1A54 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x003F1A70 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x003F1A88 | `TSilverCntlr` | Known | Controller |
| 0x003F1AA8 | `TCSpeakers` | Known | Controller |
| 0x003F1AB4 | `TCEQSetting` | Known | Controller |
| 0x003F1AD4 | `TSilverCntlr` | Known | Controller |
| 0x003F1B3C | `TSilverCntlr` | Known | Controller |
| 0x003F1B5C | `TCExtrasMenu` | Known | Controller |
| 0x003F1B6C | `TCGamesMenu` | Known | Controller |
| 0x003F1B78 | `TCGameScreen` | Known | Controller |
| 0x003F1B88 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x003F1BA8 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x003F1BC8 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x003F1BE8 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x003F1C0C | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x003F1C28 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x003F1C48 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x003F1C68 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x003F1C90 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x003F1CB4 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x003F1CDC | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x003F1CFC | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x003F1D1C | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x003F1D3C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x003F1D5C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x003F1D84 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x003F1DAC | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x003F1DCC | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x003F1DEC | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x003F1E0C | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x003F1E30 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x003F1E50 | `TSilverMediaListCntlr_MusicVideos` | Known | Controller |
| 0x003F1E74 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x003F1E9C | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x003F1EC8 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x003F1EE8 | `TCRentalNotification` | Known | Controller |
| 0x003F1F00 | `TCRentalInfo` | Known | Controller |
| 0x003F1F10 | `TCRentalConfirmDelete` | Known | Controller |
| 0x003F1F28 | `TCRentalDispatcher` | Known | Controller |
| 0x003F1F3C | `TSilverGlobalCntlr` | Known | Controller |
| 0x003F1F50 | `TSilverTrainerCntlr` | Known | Controller |
| 0x0047733C | `TCCCCCCCCCCCCCCCCCCCCCCT` | Known | Controller |
| 0x0071A462 | `TCNotesDispatcher"` | Known | Controller |
| 0x0071A521 | `TCLockChosenDispatcher"` | Known | Controller |
| 0x0071A5E4 | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x007246C9 | `TCNotesDispatcher"` | Known | Controller |
| 0x0072482B | `TCSportTimerChosenDispatcher"` | Known | Controller |
| 0x00739B34 | `TCAddressViewerMainMenu` | Known | Controller |
| 0x00739B4C | `TCAddressViewerDetails` | Known | Controller |
| 0x00739B64 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x00739B80 | `TCAlarmMenu` | Known | Controller |
| 0x00739B8C | `TCSleepTimerMenuTCAlarmPropertiesMenu` | Known | Controller |
| 0x00739BB4 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x00739BD4 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x00739BF0 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00739C0C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00739C28 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00739C44 | `TCAlarmDatePicker` | Known | Controller |
| 0x00739C58 | `TCAlarmDatePicker` | Known | Controller |
| 0x00739C6C | `TCAlarmTriggeredTSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x00739C98 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x00739CBC | `TSilverCalendarCntlr_MonthViewerTSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x00739CFC | `TSilverCalendarCntlr_EventViewerTSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x00739D3C | `TSilverCalendarCntlr_EventViewerTSilverCntlrTCClockRegionMenu` | Known | Controller |
| 0x00739D7C | `TCClockCityMenu` | Known | Controller |
| 0x00739D8C | `TCClockCityMenu` | Known | Controller |
| 0x00739D9C | `TCClockCityMenu` | Known | Controller |
| 0x00739DAC | `TCClockCityMenu` | Known | Controller |
| 0x00739DBC | `TCClockCityMenu` | Known | Controller |
| 0x00739DCC | `TCClockCityMenu` | Known | Controller |
| 0x00739DDC | `TCClockCityMenu` | Known | Controller |
| 0x00739DEC | `TCClockCityMenu` | Known | Controller |
| 0x00739DFC | `TCClock` | Known | Controller |
| 0x00739E14 | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDe` | Known | Controller |
| 0x00739E6C | `TCGamesMenu` | Known | Controller |
| 0x00739E78 | `TCGameScreenTC_LockediPod` | Known | Controller |
| 0x00739E94 | `TC_LockDialog` | Known | Controller |
| 0x00739EA4 | `TC_LockScreen` | Known | Controller |
| 0x00739EB4 | `TC_VolumeLimitLockScreenTSilverCntlrTSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x00739EF8 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x00739F18 | `TSilverMediaListCntlr_AudiobooksTSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x00739F60 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x00739F7C | `TSilverMediaListCntlr_AlbumsTSilverMediaListCntlr_Artists` | Known | Controller |
| 0x00739FB8 | `TSilverMediaListCntlr_GenresTSilverMediaListCntlr_Composers` | Known | Controller |
| 0x00739FF4 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0073A014 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0073A03C | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x0073A05C | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x0073A07C | `TSilverMediaListCntlr_TVEpisodesTSilverMediaListCntlr_MoviesTSilverMediaListCntl` | Known | Controller |
| 0x0073A0D8 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0073A100 | `TSilverMediaListCntlr_VideoPlaylistsTSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x0073A144 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x0073A170 | `TSilverCntlrTSilverMainMediaListCntlr_VideosTSilverSettingsVideoCntlr` | Known | Controller |
| 0x0073A1B8 | `TCFirstBoot` | Known | Controller |
| 0x0073A1C4 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x0073A1E4 | `TSilverMediaListCntlr_GeniusTSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0073A2C4 | `TCRentalInfoTCRentalConfirmDelete` | Known | Controller |
| 0x0073A2E8 | `TSilverCntlrTCRentalNotificationTCRentalNotificationTCRentalNotificationTCNotesL` | Known | Controller |
| 0x0073A340 | `TCNotesList` | Known | Controller |
| 0x0073A34C | `TCNotesList` | Known | Controller |
| 0x0073A358 | `TCNotesContents` | Known | Controller |
| 0x0073A368 | `TCNotesContents` | Known | Controller |
| 0x0073A378 | `TCNotesContents` | Known | Controller |
| 0x0073A388 | `TCNotesContents` | Known | Controller |
| 0x0073A444 | `TCSlideshowLCD` | Known | Controller |
| 0x0073A454 | `TCSlideshowTVOutTCSlideshow_TVOutAskTCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x0073A4A4 | `TCRemoteUI` | Known | Controller |
| 0x0073A4B0 | `TCUnsupported` | Known | Controller |
| 0x0073A4C0 | `TSilverSettingsMenuListCntlrTSilverSettingsMenuListCntlrTSilverSettingsMenuListC` | Known | Controller |
| 0x0073A528 | `TCSettings_MusicMenuTCSettings_VolumeLimit` | Known | Controller |
| 0x0073A554 | `TCSettings_Brightness` | Known | Controller |
| 0x0073A56C | `TCSettings_BacklightTimer` | Known | Controller |
| 0x0073A588 | `TCSettings_AudiobookSettingsTCSettings_RadioRegions` | Known | Controller |
| 0x0073A5BC | `TCSettings_EQ` | Known | Controller |
| 0x0073A5CC | `TSilverSettingsMenuListCntlrTSilverCntlrTCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x0073A614 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x0073A630 | `TCSettings_MainMenu` | Known | Controller |
| 0x0073A644 | `TCSettings_MusicMenuTCSportTimerTCSportTimerMenuTCSportTimerSessionScreen` | Known | Controller |
| 0x0073A690 | `TSilverCntlrTUnitTestSuiteCntlr` | Known | Controller |
| 0x0073A710 | `TCVoiceMemosTCVoiceMemosAlert` | Known | Controller |
| 0x0073A730 | `TCVoiceMemosAlert` | Known | Controller |
| 0x0073A744 | `TCVoiceMemosAlert` | Known | Controller |
| 0x0073A770 | `TCEQSetting` | Known | Controller |
| 0x0073A8DE | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x0073BE3D | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x00741BA4 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00741C02 | `TCNotesDispatcher` | Known | Controller |
| 0x00743940 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074399E | `TCNotesDispatcher` | Known | Controller |
| 0x007456DC | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074573A | `TCNotesDispatcher` | Known | Controller |
| 0x00747478 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007474D6 | `TCNotesDispatcher` | Known | Controller |
| 0x00749214 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00749272 | `TCNotesDispatcher` | Known | Controller |
| 0x0074AFB0 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074B00E | `TCNotesDispatcher` | Known | Controller |
| 0x0074CD4C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074CDAA | `TCNotesDispatcher` | Known | Controller |
| 0x0074EAE8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0074EB46 | `TCNotesDispatcher` | Known | Controller |
| 0x00750884 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007508E2 | `TCNotesDispatcher` | Known | Controller |
| 0x00752620 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075267E | `TCNotesDispatcher` | Known | Controller |
| 0x007543BC | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075441A | `TCNotesDispatcher` | Known | Controller |
| 0x00756158 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007561B6 | `TCNotesDispatcher` | Known | Controller |
| 0x00757EF4 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00757F52 | `TCNotesDispatcher` | Known | Controller |
| 0x00759C90 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00759CEE | `TCNotesDispatcher` | Known | Controller |
| 0x0075BA2C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075BA8A | `TCNotesDispatcher` | Known | Controller |
| 0x0075D7C8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075D826 | `TCNotesDispatcher` | Known | Controller |
| 0x0075F564 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0075F5C2 | `TCNotesDispatcher` | Known | Controller |
| 0x00761300 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076135E | `TCNotesDispatcher` | Known | Controller |
| 0x0076309C | `TCLockChosenDispatcher` | Known | Controller |
| 0x007630FA | `TCNotesDispatcher` | Known | Controller |
| 0x00764E38 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00764E96 | `TCNotesDispatcher` | Known | Controller |
| 0x00766BD4 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00766C32 | `TCNotesDispatcher` | Known | Controller |
| 0x00768970 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007689CE | `TCNotesDispatcher` | Known | Controller |
| 0x0076A70C | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076A76A | `TCNotesDispatcher` | Known | Controller |
| 0x0076C4A8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076C506 | `TCNotesDispatcher` | Known | Controller |
| 0x0076E244 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0076E2A2 | `TCNotesDispatcher` | Known | Controller |
| 0x0076FFE0 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077003E | `TCNotesDispatcher` | Known | Controller |
| 0x00771D7C | `TCLockChosenDispatcher` | Known | Controller |
| 0x00771DDA | `TCNotesDispatcher` | Known | Controller |
| 0x00773B18 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00773B76 | `TCNotesDispatcher` | Known | Controller |
| 0x007758B4 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00775912 | `TCNotesDispatcher` | Known | Controller |
| 0x00777650 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007776AE | `TCNotesDispatcher` | Known | Controller |
| 0x007793EC | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077944A | `TCNotesDispatcher` | Known | Controller |
| 0x0077B188 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077B1E6 | `TCNotesDispatcher` | Known | Controller |
| 0x0077CF24 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077CF82 | `TCNotesDispatcher` | Known | Controller |
| 0x0077ECC0 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0077ED1E | `TCNotesDispatcher` | Known | Controller |
| 0x00780A5C | `TCLockChosenDispatcher` | Known | Controller |
| 0x00780ABA | `TCNotesDispatcher` | Known | Controller |
| 0x007827F8 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00782856 | `TCNotesDispatcher` | Known | Controller |
| 0x00784594 | `TCLockChosenDispatcher` | Known | Controller |
| 0x007845F2 | `TCNotesDispatcher` | Known | Controller |
| 0x00792118 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x007923DA | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x00792C10 | `TCRentalDispatcher` | Known | Controller |
| 0x007934C8 | `TCRentalDispatcher` | Known | Controller |
| 0x00793D80 | `TCRentalDispatcher` | Known | Controller |
| 0x00794638 | `TCRentalDispatcher` | Known | Controller |
| 0x00794EF0 | `TCRentalDispatcher` | Known | Controller |
| 0x007957A8 | `TCRentalDispatcher` | Known | Controller |
| 0x00796060 | `TCRentalDispatcher` | Known | Controller |
| 0x00796918 | `TCRentalDispatcher` | Known | Controller |
| 0x008DD844 | `TCMockupModeNavScreen` | Known | Controller |
| 0x008DD85C | `TSilverCntlr` | Known | Controller |
| 0x008DD87C | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x008DD8CC | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x008DD8EC | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x008DD90C | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x008DD930 | `TCExtrasMenu` | Known | Controller |
| 0x008DDA40 | `TSilverCntlr` | Known | Controller |
| 0x008DDA60 | `TCSlideshowTVOut` | Known | Controller |
| 0x008DDA74 | `TCSlideshowLCD` | Known | Controller |
| 0x008DDA84 | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x008DDA9C | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x008DDABC | `TSilverGlobalCntlr` | Known | Controller |
| 0x008DDAEC | `TSilverCntlr` | Known | Controller |
| 0x008DDB68 | `TCSlideshowTVOut` | Known | Controller |
| 0x008DDB7C | `TCSlideshowLCD` | Known | Controller |
| 0x008DDB8C | `TCSlideshow_TVOutAsk` | Known | Controller |
| 0x008DDBA4 | `TCSlideshow_TVOutCableConnect` | Known | Controller |
| 0x008DDBC4 | `TSilverCntlr` | Known | Controller |
| 0x008DDC0C | `TSilverCntlr` | Known | Controller |
| 0x008DDC2C | `TCGamesMenu` | Known | Controller |
| 0x008DDC38 | `TCGameScreen` | Known | Controller |
| 0x0099B664 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |

---

## 2. Settings (Toggle/Show/TC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00127868 | `ShowSetting_EQ` | Known | User setting |
| 0x001CEC78 | `ToggleSetting_Repeat` | Known | User setting |
| 0x001CEC94 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001CECAC | `ToggleSetting_TVOut` | Known | User setting |
| 0x001CECC0 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x001F79D4 | `ShowSetting_Backlight` | Known | User setting |
| 0x0020C998 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x0020C9B4 | `ToggleSetting_Repeat` | Known | User setting |
| 0x0020C9CC | `ToggleSetting_SortBy` | Known | User setting |
| 0x0020C9E4 | `ToggleSetting_ClassicUI` | Known | User setting |
| 0x0020C9FC | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x0020CA18 | `ToggleSetting_Clicker` | Known | User setting |
| 0x0020CA30 | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x0020CA50 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x0020CA6C | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x0020CA88 | `ShowSetting_Shuffle` | Known | User setting |
| 0x0020CC34 | `ShowSetting_Repeat` | Known | User setting |
| 0x0020CC48 | `ShowSetting_About` | Known | User setting |
| 0x0020CC5C | `ShowSetting_MainMenu` | Known | User setting |
| 0x0020CC74 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x0020CC8C | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x0020CCA4 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x0020CCC0 | `ShowSetting_Brightness` | Known | User setting |
| 0x0020CCD8 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x0020CCF0 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x0020CD0C | `ShowSetting_EQ` | Known | User setting |
| 0x0020CD1C | `ShowSetting_SoundCheck` | Known | User setting |
| 0x0020CEB8 | `ShowSetting_Clicker` | Known | User setting |
| 0x0020CECC | `ShowSetting_DateAndTime` | Known | User setting |
| 0x0020CEE4 | `ShowSetting_SortBy` | Known | User setting |
| 0x0020CEF8 | `ShowSetting_ClassicUI` | Known | User setting |
| 0x0020CF10 | `ShowSetting_Language` | Known | User setting |
| 0x0020CF28 | `ShowSetting_Legal` | Known | User setting |
| 0x0020CF3C | `ShowSetting_ResetAll` | Known | User setting |
| 0x007234D9 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x00723589 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x00725CD6 | `ShowSetting_About` | Known | User setting |
| 0x00725D78 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x00725DBC | `ShowSetting_Shuffle` | Known | User setting |
| 0x00725E33 | `ToggleSetting_Repeat` | Known | User setting |
| 0x00725E76 | `ShowSetting_Repeat` | Known | User setting |
| 0x00725F80 | `ShowSetting_MainMenu` | Known | User setting |
| 0x00726090 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x00726158 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x00726222 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x0072633A | `ShowSetting_Brightness` | Known | User setting |
| 0x00726470 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x00726581 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x00726682 | `ShowSetting_EQ` | Known | User setting |
| 0x007266EF | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x00726736 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x007267B3 | `ToggleSetting_Clicker` | Known | User setting |
| 0x007267F7 | `ShowSetting_Clicker` | Known | User setting |
| 0x0072695E | `ToggleSetting_SortBy` | Known | User setting |
| 0x007269A1 | `ShowSetting_SortBy` | Known | User setting |
| 0x00726AA2 | `ShowSetting_Language` | Known | User setting |
| 0x00726BB2 | `ShowSetting_Legal` | Known | User setting |
| 0x00726CE3 | `ShowSetting_ResetAll` | Known | User setting |
| 0x00726E54 | `ShowSetting_Backlight` | Known | User setting |
| 0x00726F04 | `ShowSetting_Backlight` | Known | User setting |
| 0x00726FB4 | `ShowSetting_Backlight` | Known | User setting |
| 0x00727065 | `ShowSetting_Backlight` | Known | User setting |
| 0x00727116 | `ShowSetting_Backlight` | Known | User setting |
| 0x007271C7 | `ShowSetting_Backlight` | Known | User setting |
| 0x0072727B | `ShowSetting_Backlight` | Known | User setting |
| 0x0072732A | `ShowSetting_EQ` | Known | User setting |
| 0x0072739F | `ShowSetting_Language` | Known | User setting |
| 0x007AEF68 | `ToggleSetting_Repeat` | Known | User setting |
| 0x007AEFA2 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x007AF064 | `ToggleSetting_TVOut` | Known | User setting |
| 0x007AF09D | `ToggleSetting_TVSignal` | Known | User setting |

---

## 3. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00143198 | `MockupMode/MockupMode.xml` | Hidden | Developer Tool |
| 0x00143698 | `MockupMode/` | Hidden | Developer Tool |
| 0x00247E64 | `Channel UnitTests` | Hidden | Developer Tool |
| 0x002A141D | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x002A1460 | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x002A1475 | `RTXCbug> ` | Hidden | Developer Tool |
| 0x002A1E51 | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x002BB7EC | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x00383CDD | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x00383DA5 | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x003E221D | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x0073A6B0 | `TUnitTestSuiteTestsCntlrTSilverCntlrTCVoiceMemosMainMenuTCVoiceMemosMenuTCVoiceM` | Hidden | Developer Tool |
| 0x007D5D58 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00812C54 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0082593C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0083D820 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00850170 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0085A13C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00863DD8 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008795AC | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008834C0 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008AA678 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008C9470 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008D2AF0 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0095EE2D | `10TCDemoMode` | Hidden | Demo/Retail Mode |
| 0x0095F7A4 | `21TCMockupModeNavScreen` | Hidden | Developer Tool |
| 0x0095FC64 | `27TSilverCntlrTransitionAddonI10TCDemoModeE` | Hidden | Demo/Retail Mode |
| 0x0098D8A7 | `Debug_ListItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x0098D8BF | `Debug_MenuItem_DemoMode` | Hidden | Demo/Retail Mode |
| 0x0098DFC4 | `DemoMode_Background_Image` | Hidden | Demo/Retail Mode |
| 0x0098EB5E | `DemoMode_Icon_Image` | Hidden | Demo/Retail Mode |
| 0x00990720 | `Debug_Menu_Title` | Hidden | Debug/Diagnostic |
| 0x00990745 | `UnitTest_Menu_Title` | Hidden | Developer Tool |
| 0x0099921C | `UnitTestModel` | Hidden | Developer Tool |
| 0x00999BFB | `UnitTest_ListItem` | Hidden | Developer Tool |
| 0x0099AD11 | `DemoMode_View_Main` | Hidden | Demo/Retail Mode |
| 0x0099AEE6 | `DemoMode_Icon` | Hidden | Demo/Retail Mode |
| 0x0099BCCC | `UnitTestApp` | Hidden | Developer Tool |
| 0x0099C27E | `Debug_ListItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x0099C299 | `Debug_MenuItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x0099C9AF | `DemoMode_Text_Color` | Hidden | Demo/Retail Mode |
| 0x0099CDC4 | `Debug_ListItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x0099CDDB | `Debug_MenuItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009A0E5A | `Debug_ListItem_UnitTest` | Hidden | Developer Tool |
| 0x009A0E72 | `Debug_MenuItem_UnitTest` | Hidden | Developer Tool |
| 0x009A526C | `Debug_ListItem_Memory` | Hidden | Debug/Diagnostic |
| 0x009A5282 | `Debug_MenuItem_Memory` | Hidden | Debug/Diagnostic |

---

## 4. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000067B3 | `"MeCCADecode` | Known | Audio system |
| 0x001391DC | `AudioCodecs` | Known | Audio system |
| 0x0017EBDC | `MeCCA_PCM_Output.wav` | Known | Audio system |
| 0x00197E1C | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x001A1734 | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x001A193C | `MeCCAVideoDecode` | Known | Audio system |
| 0x008E99B0 | `MeCCA_StreamCache` | Known | Audio system |

---

## 5. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000E8308 | `HandleWheel` | Known | Event handler |
| 0x000E8314 | `HandlePlayPause` | Known | Event handler |
| 0x000E8324 | `HandleSelectDown` | Known | Event handler |
| 0x000E8338 | `HandleNext` | Known | Event handler |
| 0x000E8344 | `HandlePrevious` | Known | Event handler |
| 0x000E8354 | `HandleNextPushAndHold` | Known | Event handler |
| 0x000E836C | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000E8604 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000E8624 | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x000F444C | `HandleSelect` | Known | Event handler |
| 0x000F4460 | `HandleHilite` | Known | Event handler |
| 0x000F47F8 | `HandleEQSettingSelected` | Known | Event handler |
| 0x000F4C28 | `HandleSelect` | Known | Event handler |
| 0x000F4C3C | `HandleGameHilited` | Known | Event handler |
| 0x000F4EEC | `HandleNotesSelected` | Known | Event handler |
| 0x000F4F04 | `HandleNotesPop` | Known | Event handler |
| 0x000F4F14 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x001032A8 | `HandleVolumeWheel` | Known | Event handler |
| 0x001032BC | `HandleVolumeChange` | Known | Event handler |
| 0x001032D0 | `HandleTimerDone` | Known | Event handler |
| 0x001032E0 | `HandleFrequencyChange` | Known | Event handler |
| 0x00103358 | `HandleTuning` | Known | Event handler |
| 0x00103368 | `HandleTuningSelect` | Known | Event handler |
| 0x0010DEAC | `HandleLock` | Known | Event handler |
| 0x0010DEBC | `HandleAddressBook` | Known | Event handler |
| 0x0010E5A4 | `HandleSelect` | Known | Event handler |
| 0x0010EADC | `HandleExit` | Known | Event handler |
| 0x0010EAEC | `HandleLap` | Known | Event handler |
| 0x0010EAF8 | `HandleResume` | Known | Event handler |
| 0x0010EB08 | `HandleStartStop` | Known | Event handler |
| 0x0010EDB8 | `HandleWheel` | Known | Event handler |
| 0x0010EDC8 | `HandlePlayPause` | Known | Event handler |
| 0x0010EDD8 | `HandleSelectDown` | Known | Event handler |
| 0x0010EDEC | `HandleHilite` | Known | Event handler |
| 0x0010EE10 | `HandleFinishRecording` | Known | Event handler |
| 0x00119520 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x00127A9C | `HandleExitUnsupported` | Known | Event handler |
| 0x0013EA18 | `HandleNotesPop` | Known | Event handler |
| 0x0013EA2C | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0013F938 | `HandleSelect` | Known | Event handler |
| 0x0013F94C | `HandleWheel` | Known | Event handler |
| 0x0013F958 | `HandleImageNext` | Known | Event handler |
| 0x0013F968 | `HandleImagePrev` | Known | Event handler |
| 0x0013F978 | `HandleImageLast` | Known | Event handler |
| 0x0013F988 | `HandleImageFirst` | Known | Event handler |
| 0x0013F99C | `HandlePlayPause` | Known | Event handler |
| 0x0013F9AC | `HandlePlay` | Known | Event handler |
| 0x0013F9B8 | `HandlePause` | Known | Event handler |
| 0x0013F9C4 | `HandleMikeyCenter` | Known | Event handler |
| 0x001543C8 | `HandleSelectCity` | Known | Event handler |
| 0x001543E0 | `HandleHighlightCity` | Known | Event handler |
| 0x001554CC | `HandleWantPopFlow` | Known | Event handler |
| 0x001554E4 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x00155500 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0015551C | `HandleFlowNext` | Known | Event handler |
| 0x0015552C | `HandleFlowPrev` | Known | Event handler |
| 0x0015553C | `HandleFlowWheel` | Known | Event handler |
| 0x0015554C | `HandleAlbumSelected` | Known | Event handler |
| 0x00155560 | `HandlePlayPause` | Known | Event handler |
| 0x00155570 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00180A78 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00180E68 | `HandleSelect` | Known | Event handler |
| 0x00181D50 | `HandleSelect` | Known | Event handler |
| 0x00181D64 | `HandleWheel` | Known | Event handler |
| 0x00181D70 | `HandleImageNext` | Known | Event handler |
| 0x00181D80 | `HandleImagePrev` | Known | Event handler |
| 0x00181D90 | `HandleImageLast` | Known | Event handler |
| 0x00181DA0 | `HandleImageFirst` | Known | Event handler |
| 0x00181DB4 | `HandlePlayPause` | Known | Event handler |
| 0x00181DC4 | `HandlePlay` | Known | Event handler |
| 0x00181DD0 | `HandlePause` | Known | Event handler |
| 0x00181DDC | `HandleMikeyCenter` | Known | Event handler |
| 0x00182284 | `HandleNew` | Known | Event handler |
| 0x00182294 | `HandleClear` | Known | Event handler |
| 0x001822A0 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x001822BC | `HandleSelectIndexedSession` | Known | Event handler |
| 0x001825CC | `HandleWheel` | Known | Event handler |
| 0x001825DC | `HandleArrowUp` | Known | Event handler |
| 0x001825EC | `HandleArrowDown` | Known | Event handler |
| 0x00184B0C | `HandleHiliteAlbum` | Known | Event handler |
| 0x00184B24 | `HandleBrowseAlbum` | Known | Event handler |
| 0x00184B38 | `HandlePlayPause` | Known | Event handler |
| 0x0019B43C | `HandleSelect` | Known | Event handler |
| 0x0019B5CC | `HandleSelectRegion` | Known | Event handler |
| 0x0019B944 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x0019B960 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x0019B97C | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x001B1010 | `HandleImageWheel` | Known | Event handler |
| 0x001B1028 | `HandlePlayPause` | Known | Event handler |
| 0x001B1038 | `HandleBrowseLarge` | Known | Event handler |
| 0x001B104C | `HandleBrowseSmall` | Known | Event handler |
| 0x001B1060 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x001B1078 | `HandleImageNext` | Known | Event handler |
| 0x001B1088 | `HandleImagePrev` | Known | Event handler |
| 0x001B1098 | `HandleHilite` | Known | Event handler |
| 0x001B10A8 | `HandleImageLast` | Known | Event handler |
| 0x001B10B8 | `HandleImageFirst` | Known | Event handler |
| 0x001B10CC | `HandleScreenNext` | Known | Event handler |
| 0x001B10E0 | `HandleScreenPrev` | Known | Event handler |
| 0x001B39A8 | `HandlePlayPause` | Known | Event handler |
| 0x001B39BC | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x001B39D8 | `HandleNext` | Known | Event handler |
| 0x001B39E4 | `HandleNextPressAndHold` | Known | Event handler |
| 0x001B39FC | `HandlePrevious` | Known | Event handler |
| 0x001B3A0C | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x001B3A28 | `HandleRemotePlayPause` | Known | Event handler |
| 0x001B3A40 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x001B3A64 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x001B3A7C | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x001B3A94 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x001B3C38 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x001B3C50 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x001B3C68 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x001B3C84 | `HandleRemoteStop` | Known | Event handler |
| 0x001B3C98 | `HandleRemotePlay` | Known | Event handler |
| 0x001B3CAC | `HandleRemotePause` | Known | Event handler |
| 0x001B3CC0 | `HandleRemoteMute` | Known | Event handler |
| 0x001B3CD4 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x001B3CEC | `HandleRemotePrevChapter` | Known | Event handler |
| 0x001B3D04 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x001B3D20 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x001B3F28 | `HandleRemoteShuffle` | Known | Event handler |
| 0x001B3F3C | `HandleRemoteRepeat` | Known | Event handler |
| 0x001B3F50 | `HandleRemoteOn` | Known | Event handler |
| 0x001B3F64 | `HandleRemoteOff` | Known | Event handler |
| 0x001B3F74 | `HandleRemoteBacklight` | Known | Event handler |
| 0x001B3F8C | `HandleRemoteFFDown` | Known | Event handler |
| 0x001B3FA0 | `HandleRemoteFFUp` | Known | Event handler |
| 0x001B3FB4 | `HandleRemoteRewDown` | Known | Event handler |
| 0x001B3FC8 | `HandleRemoteRewUp` | Known | Event handler |
| 0x001B3FDC | `HandleRemoteMenuDown` | Known | Event handler |
| 0x001B3FF4 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x001B4008 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x001B4020 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x001B41D0 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x001B41E8 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x001B4200 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x001B421C | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x001B4234 | `HandleRemoteEvent` | Known | Event handler |
| 0x001B4248 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x001B4264 | `HandleAudioPlayPause` | Known | Event handler |
| 0x001B427C | `HandleAudioNext` | Known | Event handler |
| 0x001B428C | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x001B42A8 | `HandleAudioPrevious` | Known | Event handler |
| 0x001B42BC | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x001B444C | `HandleAudioNextAlbum` | Known | Event handler |
| 0x001B4464 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x001B447C | `HandleAudioVolumeDown` | Known | Event handler |
| 0x001B4494 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x001B44A8 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x001B44C0 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x001B44D8 | `HandleAudioStop` | Known | Event handler |
| 0x001B44E8 | `HandleAudioPlay` | Known | Event handler |
| 0x001B44F8 | `HandleAudioPause` | Known | Event handler |
| 0x001B450C | `HandleAudioMute` | Known | Event handler |
| 0x001B451C | `HandleAudioNextChapter` | Known | Event handler |
| 0x001B4534 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x001B4720 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x001B4738 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x001B4750 | `HandleAudioShuffle` | Known | Event handler |
| 0x001B4764 | `HandleAudioRepeat` | Known | Event handler |
| 0x001B4778 | `HandleAudioFFDown` | Known | Event handler |
| 0x001B478C | `HandleAudioFFUp` | Known | Event handler |
| 0x001B479C | `HandleAudioRewDown` | Known | Event handler |
| 0x001B47B0 | `HandleAudioRewUp` | Known | Event handler |
| 0x001B47C4 | `HandleVideoPlayPause` | Known | Event handler |
| 0x001B47DC | `HandleVideoNext` | Known | Event handler |
| 0x001B47EC | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x001B4808 | `HandleVideoPrevious` | Known | Event handler |
| 0x001B481C | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x001B49E0 | `HandleVideoStop` | Known | Event handler |
| 0x001B49F0 | `HandleVideoPlay` | Known | Event handler |
| 0x001B4A00 | `HandleVideoPause` | Known | Event handler |
| 0x001B4A14 | `HandleVideoFFDown` | Known | Event handler |
| 0x001B4A28 | `HandleVideoFFUp` | Known | Event handler |
| 0x001B4A38 | `HandleVideoRewDown` | Known | Event handler |
| 0x001B4A4C | `HandleVideoRewUp` | Known | Event handler |
| 0x001B4A60 | `HandleVideoNextChapter` | Known | Event handler |
| 0x001B4A78 | `HandleVideoPrevChapter` | Known | Event handler |
| 0x001B4A90 | `HandleVideoNextFrame` | Known | Event handler |
| 0x001B4AA8 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x001B4AC0 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x001B4ADC | `HandleMikeyCenter` | Known | Event handler |
| 0x001B4C3C | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x001B4C5C | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x001B4C7C | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x001B4CA0 | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x001B4CC0 | `HandleMikeyAllUp` | Known | Event handler |
| 0x001B4CD4 | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x001B4CE8 | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x001B4D00 | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x001B4D18 | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x001C1AE0 | `HandleMainMenu` | Known | Event handler |
| 0x001C3C2C | `HandleLoadingCancelled` | Known | Event handler |
| 0x001C6660 | `HandlePowerSongSelected` | Known | Event handler |
| 0x001C667C | `HandlePowerSongChosen` | Known | Event handler |
| 0x001C6694 | `HandleSelectTimedWorkout` | Known | Event handler |
| 0x001CCBE0 | `HandleSelect` | Known | Event handler |
| 0x001CCE88 | `HandleMusicMenu` | Known | Event handler |
| 0x001CD148 | `HandleSelect` | Known | Event handler |
| 0x001CD44C | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x001CD46C | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x001CD928 | `HandleWheel` | Known | Event handler |
| 0x001CD938 | `HandlePlayPause` | Known | Event handler |
| 0x001CD948 | `HandleSelectDown` | Known | Event handler |
| 0x001CD95C | `HandleNext` | Known | Event handler |
| 0x001CD968 | `HandlePrevious` | Known | Event handler |
| 0x001CD978 | `HandleNextPushAndHold` | Known | Event handler |
| 0x001CD990 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x001CE084 | `HandleMenuSelection` | Known | Event handler |
| 0x001CE098 | `HandleViewAlbum` | Known | Event handler |
| 0x001CE0A8 | `HandleViewArtist` | Known | Event handler |
| 0x001CE0BC | `HandleViewCompilation` | Known | Event handler |
| 0x001CE0D4 | `HandleStartGenius` | Known | Event handler |
| 0x001DACCC | `HandleFrequencyChosen` | Known | Event handler |
| 0x001DACE4 | `HandleDateChosen` | Known | Event handler |
| 0x001DACF8 | `HandleTimeChosen` | Known | Event handler |
| 0x001DAD0C | `HandleSoundChosen` | Known | Event handler |
| 0x001DAD20 | `HandleLabelChosen` | Known | Event handler |
| 0x001DAD34 | `HandleDeleteChosen` | Known | Event handler |
| 0x001DBE14 | `HandleSelect` | Known | Event handler |
| 0x001E0730 | `HandlePrev` | Known | Event handler |
| 0x001E0740 | `HandleNext` | Known | Event handler |
| 0x001E074C | `HandlePlayPause` | Known | Event handler |
| 0x001E8108 | `HandleNextContact` | Known | Event handler |
| 0x001E8120 | `HandlePreviousContact` | Known | Event handler |
| 0x001EFCC4 | `HandleItemSelected` | Known | Event handler |
| 0x001EFEBC | `HandleRadioRegion` | Known | Event handler |
| 0x001F00A4 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x001F4318 | `HandlePlayPause` | Known | Event handler |
| 0x001F7CB0 | `HandleDelete` | Known | Event handler |
| 0x001F7CC4 | `HandleSelectLozinch` | Known | Event handler |
| 0x001F7F6C | `HandleSelect` | Known | Event handler |
| 0x001F8238 | `HandleTVOutChanged` | Known | Event handler |
| 0x001F8250 | `HandleTVSignalChanged` | Known | Event handler |
| 0x001F8268 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x001F8288 | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x001F82A8 | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x001F82CC | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x001F82EC | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x001FB14C | `HandleSelectKey` | Known | Event handler |
| 0x001FB2F4 | `HandleSelect` | Known | Event handler |
| 0x001FC070 | `HandlePlayPause` | Known | Event handler |
| 0x001FC084 | `HandleWheel` | Known | Event handler |
| 0x001FC090 | `HandleWheelRating` | Known | Event handler |
| 0x001FC0A4 | `HandleWheelScrub` | Known | Event handler |
| 0x001FC0B8 | `HandleWheelVolume` | Known | Event handler |
| 0x001FC178 | `HandleMenuKey` | Known | Event handler |
| 0x001FC1E4 | `HandleMenuLongpress` | Known | Event handler |
| 0x001FC1F8 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x001FCE00 | `HandleSelect` | Known | Event handler |
| 0x001FD6F8 | `HandleLeaveAlarm` | Known | Event handler |
| 0x001FE610 | `HandleSelect` | Known | Event handler |
| 0x001FE624 | `HandleHilite` | Known | Event handler |
| 0x001FE634 | `HandlePlayPause` | Known | Event handler |
| 0x001FE644 | `HandleAddToOTG` | Known | Event handler |
| 0x001FE654 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x001FE674 | `HandleShowContextualMenu` | Known | Event handler |
| 0x002016FC | `HandleLanguageAfterReset` | Known | Event handler |
| 0x00201F0C | `HandleSelect` | Known | Event handler |
| 0x00201F20 | `HandleWheel` | Known | Event handler |
| 0x00201F2C | `HandleWheelProgress` | Known | Event handler |
| 0x00201F40 | `HandleSelectProgress` | Known | Event handler |
| 0x00201F58 | `HandleSelectVolume` | Known | Event handler |
| 0x00201F6C | `HandleSelectScrub` | Known | Event handler |
| 0x00201F80 | `HandleSelectGenius` | Known | Event handler |
| 0x00201F94 | `HandleSelectRating` | Known | Event handler |
| 0x00201FA8 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x00201FC0 | `HandleSelectChapterArt` | Known | Event handler |
| 0x00201FD8 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x00201FF4 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x002021F0 | `HandleWheelGenius` | Known | Event handler |
| 0x00202204 | `HandleWheelBrightness` | Known | Event handler |
| 0x00202270 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x00202290 | `HandlePushContextualMenu` | Known | Event handler |
| 0x002022AC | `HandleAddToOTG` | Known | Event handler |
| 0x002022BC | `HandleViewArtist` | Known | Event handler |
| 0x002022D0 | `HandleViewAlbum` | Known | Event handler |
| 0x002022E0 | `HandleViewCompilation` | Known | Event handler |
| 0x002023D8 | `HandleStartGenius` | Known | Event handler |
| 0x002023EC | `HandleAudiobookSlower` | Known | Event handler |
| 0x00202404 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0020241C | `HandleAudiobookNormal` | Known | Event handler |
| 0x00202434 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00203ED0 | `HandleStartGenius` | Known | Event handler |
| 0x00204224 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0020423C | `HandleAudiobookNormal` | Known | Event handler |
| 0x00204254 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0020426C | `HandleStartGenius` | Known | Event handler |
| 0x00204280 | `HandleAddToOTG` | Known | Event handler |
| 0x00204290 | `HandleViewCompilation` | Known | Event handler |
| 0x002042A8 | `HandleViewAlbum` | Known | Event handler |
| 0x002042B8 | `HandleViewArtist` | Known | Event handler |
| 0x002042CC | `HandleCancel` | Known | Event handler |
| 0x00204D68 | `HandleSelect` | Known | Event handler |
| 0x00204D78 | `HandleSelectRating` | Known | Event handler |
| 0x00204D8C | `HandleSelectProgress` | Known | Event handler |
| 0x00204DA4 | `HandleWheelProgress` | Known | Event handler |
| 0x00204DB8 | `HandleSelectScrub` | Known | Event handler |
| 0x00204DCC | `HandleWheelBrightness` | Known | Event handler |
| 0x00204DE4 | `HandleNextTrackGotoScrub` | Known | Event handler |
| 0x00204E00 | `HandlePrevTrackGotoScrub` | Known | Event handler |
| 0x00204E1C | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x00207A04 | `HandleStartGenius` | Known | Event handler |
| 0x00207A1C | `HandleViewArtist` | Known | Event handler |
| 0x00207A30 | `HandleViewAlbum` | Known | Event handler |
| 0x00207A40 | `HandleViewCompilation` | Known | Event handler |
| 0x00207A58 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x002083DC | `HandleStartGenius` | Known | Event handler |
| 0x002083F0 | `HandleAddToOTG` | Known | Event handler |
| 0x00208400 | `HandleViewCompilation` | Known | Event handler |
| 0x00208418 | `HandleViewAlbum` | Known | Event handler |
| 0x00208428 | `HandleViewArtist` | Known | Event handler |
| 0x0020843C | `HandleCancel` | Known | Event handler |
| 0x0020ADCC | `HandleAddToOTG` | Known | Event handler |
| 0x0020ADDC | `HandleCancel` | Known | Event handler |
| 0x0020AFD0 | `HandleStartGenius` | Known | Event handler |
| 0x0020AFE8 | `HandleViewAlbum` | Known | Event handler |
| 0x0020AFF8 | `HandleViewArtist` | Known | Event handler |
| 0x0020B00C | `HandleViewCompilation` | Known | Event handler |
| 0x0020B024 | `HandleSaveGeniusPlaylist` | Known | Event handler |
| 0x0020B040 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x0020B058 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0020BFC0 | `HandleStartGenius` | Known | Event handler |
| 0x0020BFD4 | `HandleAddToOTG` | Known | Event handler |
| 0x0020BFE4 | `HandleViewCompilation` | Known | Event handler |
| 0x0020BFFC | `HandleViewAlbum` | Known | Event handler |
| 0x0020C00C | `HandleViewArtist` | Known | Event handler |
| 0x0020C020 | `HandleCancel` | Known | Event handler |
| 0x0020C4CC | `HandleAddToOTG` | Known | Event handler |
| 0x0020C4DC | `HandleCancel` | Known | Event handler |
| 0x0020CF74 | `HandleLanguage` | Known | Event handler |
| 0x0020CF84 | `HandleResetAllSettings` | Known | Event handler |
| 0x0020CF9C | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x0020D908 | `HandleSelect` | Known | Event handler |
| 0x0020DB38 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x0020E964 | `HandleAddToOTG` | Known | Event handler |
| 0x0020E974 | `HandleCancel` | Known | Event handler |
| 0x0021145C | `HandleSelect` | Known | Event handler |
| 0x002115F8 | `HandleSelect` | Known | Event handler |
| 0x00211898 | `HandleNextDay` | Known | Event handler |
| 0x002118AC | `HandlePreviousDay` | Known | Event handler |
| 0x002120B0 | `HandleMusicHilited` | Known | Event handler |
| 0x002120C8 | `HandleVideosHilited` | Known | Event handler |
| 0x002120DC | `HandlePodcastsHilited` | Known | Event handler |
| 0x002120F4 | `HandleGenericHilited` | Known | Event handler |
| 0x0021210C | `HandlePhotosHilited` | Known | Event handler |
| 0x00212120 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x00212138 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x00212154 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x0021216C | `HandleArtistsHilited` | Known | Event handler |
| 0x00212184 | `HandleGenresHilited` | Known | Event handler |
| 0x00212198 | `HandleAlbumsHilited` | Known | Event handler |
| 0x002121AC | `HandleCompilationsHilited` | Known | Event handler |
| 0x0021237C | `HandleComposersHilited` | Known | Event handler |
| 0x00212394 | `HandleSongsHilited` | Known | Event handler |
| 0x002123A8 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x002123C0 | `HandleGeniusHilited` | Known | Event handler |
| 0x002123D4 | `HandleTVShowsHilited` | Known | Event handler |
| 0x002123EC | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00212408 | `HandleMoviesHilited` | Known | Event handler |
| 0x0021241C | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00212438 | `HandleRentalsHilited` | Known | Event handler |
| 0x00212450 | `HandleMusicSelected` | Known | Event handler |
| 0x00212464 | `HandleVideosSelected` | Known | Event handler |
| 0x00212634 | `HandlePodcastsSelected` | Known | Event handler |
| 0x0021264C | `HandlePhotosSelected` | Known | Event handler |
| 0x00212664 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x0021267C | `HandleSongsSelected` | Known | Event handler |
| 0x00212690 | `HandleAlbumsSelected` | Known | Event handler |
| 0x002126A8 | `HandleCompilationsSelected` | Known | Event handler |
| 0x002126C4 | `HandleArtistsSelected` | Known | Event handler |
| 0x002126DC | `HandleGenresSelected` | Known | Event handler |
| 0x002126F4 | `HandleComposersSelected` | Known | Event handler |
| 0x0021270C | `HandleAudiobooksSelected` | Known | Event handler |
| 0x00212728 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x002128FC | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00212914 | `HandleNowPlaying` | Known | Event handler |
| 0x00212928 | `HandleGotoGenius` | Known | Event handler |
| 0x0021293C | `HandleTVShowsSelected` | Known | Event handler |
| 0x00212954 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00212970 | `HandleMoviesSelected` | Known | Event handler |
| 0x00212988 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x002129A8 | `HandleRentalsSelected` | Known | Event handler |
| 0x002129C0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x002129D8 | `HandleLock` | Known | Event handler |
| 0x002129E4 | `HandleBacklightSelected` | Known | Event handler |
| 0x00212A44 | `HandleSleepSelected` | Known | Event handler |
| 0x00212A58 | `HandleNikePlusSelected` | Known | Event handler |
| 0x0021547C | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00215A18 | `HandleAddToOTG` | Known | Event handler |
| 0x00215A28 | `HandleCancel` | Known | Event handler |
| 0x00215BF8 | `HandleWheel` | Known | Event handler |
| 0x00216A74 | `HandleAddToOTG` | Known | Event handler |
| 0x00216A84 | `HandleCancel` | Known | Event handler |
| 0x00217114 | `HandleAddToOTG` | Known | Event handler |
| 0x00217124 | `HandleCancel` | Known | Event handler |
| 0x00217AD4 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x00217D2C | `HandleNextDay` | Known | Event handler |
| 0x00217D40 | `HandlePreviousDay` | Known | Event handler |
| 0x00217F88 | `HandleSelect` | Known | Event handler |
| 0x00218224 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x002186F4 | `HandleAddToOTG` | Known | Event handler |
| 0x00218704 | `HandleCancel` | Known | Event handler |
| 0x0021BDCC | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x0021BDE8 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x0021BE00 | `HandleStartGenius` | Known | Event handler |
| 0x0021BE14 | `HandleViewArtist` | Known | Event handler |
| 0x0021BE28 | `HandleViewAlbum` | Known | Event handler |
| 0x0021BE38 | `HandleViewCompilation` | Known | Event handler |
| 0x0021BE50 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0021BE6C | `HandleRefreshPlaylist` | Known | Event handler |
| 0x0021BE84 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0021D13C | `HandleStartGenius` | Known | Event handler |
| 0x0021D150 | `HandleAddToOTG` | Known | Event handler |
| 0x0021D160 | `HandleViewCompilation` | Known | Event handler |
| 0x0021D178 | `HandleViewAlbum` | Known | Event handler |
| 0x0021D188 | `HandleViewArtist` | Known | Event handler |
| 0x0021D19C | `HandleCancel` | Known | Event handler |
| 0x0021D910 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0021DB74 | `HandleAddToOTG` | Known | Event handler |
| 0x0021DB84 | `HandleCancel` | Known | Event handler |
| 0x0021E078 | `HandleSelect` | Known | Event handler |
| 0x0021E744 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x00257720 | `HandleDeleteClock` | Known | Event handler |
| 0x00257738 | `HandleSelectClock` | Known | Event handler |
| 0x0025774C | `HandleHilited` | Known | Event handler |
| 0x0025775C | `HandleWheel` | Known | Event handler |
| 0x00257768 | `HandleSelectLozinch` | Known | Event handler |
| 0x00412392 | `HandleAudioFFDown` | Known | Event handler |
| 0x004123BB | `HandleAudioFFUp` | Known | Event handler |
| 0x004123E6 | `HandleAudioMute` | Known | Event handler |
| 0x00412419 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x0041244E | `HandleAudioNext` | Known | Event handler |
| 0x0041247E | `HandleAudioNextAlbum` | Known | Event handler |
| 0x004124B5 | `HandleAudioNextChapter` | Known | Event handler |
| 0x004124EF | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x00412523 | `HandleAudioPause` | Known | Event handler |
| 0x0041254F | `HandleAudioPlay` | Known | Event handler |
| 0x0041257D | `HandleAudioPlayPause` | Known | Event handler |
| 0x004125B5 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x004125EE | `HandleAudioPrevious` | Known | Event handler |
| 0x00412622 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x00412659 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x00412693 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x004126C8 | `HandleAudioRepeat` | Known | Event handler |
| 0x004126F4 | `HandleAudioRewDown` | Known | Event handler |
| 0x0041271F | `HandleAudioRewUp` | Known | Event handler |
| 0x0041274E | `HandleAudioShuffle` | Known | Event handler |
| 0x0041277C | `HandleAudioStop` | Known | Event handler |
| 0x004127AD | `HandleAudioVolumeDown` | Known | Event handler |
| 0x004127E2 | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x00412819 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x0041284A | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x00412903 | `HandleNextPressAndHold` | Known | Event handler |
| 0x00412934 | `HandleNext` | Known | Event handler |
| 0x0041296C | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x004129A7 | `HandlePlayPause` | Known | Event handler |
| 0x004129DB | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x00412A10 | `HandlePrevious` | Known | Event handler |
| 0x00412AA2 | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x00412AEA | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x00412B33 | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x00412B75 | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x00412BAD | `HandleMikeyCenter` | Known | Event handler |
| 0x00412BE0 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x00412C16 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x00412C4E | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x00412C80 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x00412CB6 | `HandleRemoteBacklight` | Known | Event handler |
| 0x00412CEE | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x00412D28 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x00412D61 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x00412D96 | `HandleRemoteEvent` | Known | Event handler |
| 0x00412DC2 | `HandleRemoteFFDown` | Known | Event handler |
| 0x00412DED | `HandleRemoteFFUp` | Known | Event handler |
| 0x00412E1A | `HandleRemoteMenuDown` | Known | Event handler |
| 0x00412E49 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x00412E78 | `HandleRemoteMute` | Known | Event handler |
| 0x00412EAA | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x00412EE3 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x00412F1F | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x00412F5F | `HandleRemoteOff` | Known | Event handler |
| 0x00412F88 | `HandleRemoteOff` | Known | Event handler |
| 0x00412FB2 | `HandleRemoteOn` | Known | Event handler |
| 0x00412FDE | `HandleRemotePause` | Known | Event handler |
| 0x0041300C | `HandleRemotePlay` | Known | Event handler |
| 0x0041304A | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x0041308B | `HandleRemotePlayPause` | Known | Event handler |
| 0x004130C2 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x004130FB | `HandleRemotePrevChapter` | Known | Event handler |
| 0x00413137 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x0041316E | `HandleRemoteRepeat` | Known | Event handler |
| 0x0041319C | `HandleRemoteRewDown` | Known | Event handler |
| 0x004131C9 | `HandleRemoteRewUp` | Known | Event handler |
| 0x004131F9 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x0041322C | `HandleRemoteSelectUp` | Known | Event handler |
| 0x00413260 | `HandleRemoteShuffle` | Known | Event handler |
| 0x00413290 | `HandleRemoteStop` | Known | Event handler |
| 0x004132C0 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x004132F5 | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x0041332D | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x00413364 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x0041339D | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x004133D0 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x00413405 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x00413438 | `HandleVideoFFDown` | Known | Event handler |
| 0x00413461 | `HandleVideoFFUp` | Known | Event handler |
| 0x00413494 | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x004134C9 | `HandleVideoNext` | Known | Event handler |
| 0x004134FB | `HandleVideoNextChapter` | Known | Event handler |
| 0x00413532 | `HandleVideoNextFrame` | Known | Event handler |
| 0x00413563 | `HandleVideoPause` | Known | Event handler |
| 0x0041358F | `HandleVideoPlay` | Known | Event handler |
| 0x004135BD | `HandleVideoPlayPause` | Known | Event handler |
| 0x004135F5 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x0041362E | `HandleVideoPrevious` | Known | Event handler |
| 0x00413664 | `HandleVideoPrevChapter` | Known | Event handler |
| 0x0041369B | `HandleVideoPrevFrame` | Known | Event handler |
| 0x004136CA | `HandleVideoRewDown` | Known | Event handler |
| 0x004136F5 | `HandleVideoRewUp` | Known | Event handler |
| 0x00413721 | `HandleVideoStop` | Known | Event handler |
| 0x0071A1E6 | `HandleAddressBook` | Known | Event handler |
| 0x0071A782 | `HandleSelect` | Known | Event handler |
| 0x0071A7BD | `HandleHilite` | Known | Event handler |
| 0x0071A83E | `HandleSelectRegion` | Known | Event handler |
| 0x0071A8DE | `HandleSelectRegion` | Known | Event handler |
| 0x0071A97A | `HandleSelectRegion` | Known | Event handler |
| 0x0071AA1E | `HandleSelectRegion` | Known | Event handler |
| 0x0071AAC4 | `HandleSelectRegion` | Known | Event handler |
| 0x0071AB64 | `HandleSelectRegion` | Known | Event handler |
| 0x0071AC10 | `HandleSelectRegion` | Known | Event handler |
| 0x0071ACB2 | `HandleSelectRegion` | Known | Event handler |
| 0x0071AD62 | `HandleSelectCity` | Known | Event handler |
| 0x0071ADCE | `HandleHighlightCity` | Known | Event handler |
| 0x0071AE07 | `HandleSelectCity` | Known | Event handler |
| 0x0071AE73 | `HandleHighlightCity` | Known | Event handler |
| 0x0071AEAC | `HandleSelectCity` | Known | Event handler |
| 0x0071AF18 | `HandleHighlightCity` | Known | Event handler |
| 0x0071AF51 | `HandleSelectCity` | Known | Event handler |
| 0x0071AFBD | `HandleHighlightCity` | Known | Event handler |
| 0x0071AFF6 | `HandleSelectCity` | Known | Event handler |
| 0x0071B062 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B09B | `HandleSelectCity` | Known | Event handler |
| 0x0071B107 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B140 | `HandleSelectCity` | Known | Event handler |
| 0x0071B1AC | `HandleHighlightCity` | Known | Event handler |
| 0x0071B1E5 | `HandleSelectCity` | Known | Event handler |
| 0x0071B251 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B28A | `HandleSelectCity` | Known | Event handler |
| 0x0071B2F6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B32F | `HandleSelectCity` | Known | Event handler |
| 0x0071B39B | `HandleHighlightCity` | Known | Event handler |
| 0x0071B3D4 | `HandleSelectCity` | Known | Event handler |
| 0x0071B440 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B479 | `HandleSelectCity` | Known | Event handler |
| 0x0071B4E5 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B51E | `HandleSelectCity` | Known | Event handler |
| 0x0071B58A | `HandleHighlightCity` | Known | Event handler |
| 0x0071B5C3 | `HandleSelectCity` | Known | Event handler |
| 0x0071B62F | `HandleHighlightCity` | Known | Event handler |
| 0x0071B668 | `HandleSelectCity` | Known | Event handler |
| 0x0071B6D4 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B70D | `HandleSelectCity` | Known | Event handler |
| 0x0071B779 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B7B2 | `HandleSelectCity` | Known | Event handler |
| 0x0071B81E | `HandleHighlightCity` | Known | Event handler |
| 0x0071B857 | `HandleSelectCity` | Known | Event handler |
| 0x0071B8C3 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B8FC | `HandleSelectCity` | Known | Event handler |
| 0x0071B968 | `HandleHighlightCity` | Known | Event handler |
| 0x0071B9A1 | `HandleSelectCity` | Known | Event handler |
| 0x0071BA0D | `HandleHighlightCity` | Known | Event handler |
| 0x0071BA46 | `HandleSelectCity` | Known | Event handler |
| 0x0071BAB2 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BAEB | `HandleSelectCity` | Known | Event handler |
| 0x0071BB57 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BB90 | `HandleSelectCity` | Known | Event handler |
| 0x0071BBFC | `HandleHighlightCity` | Known | Event handler |
| 0x0071BC35 | `HandleSelectCity` | Known | Event handler |
| 0x0071BCA1 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BCDA | `HandleSelectCity` | Known | Event handler |
| 0x0071BD46 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BD7F | `HandleSelectCity` | Known | Event handler |
| 0x0071BDEB | `HandleHighlightCity` | Known | Event handler |
| 0x0071BE24 | `HandleSelectCity` | Known | Event handler |
| 0x0071BE90 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BEC9 | `HandleSelectCity` | Known | Event handler |
| 0x0071BF35 | `HandleHighlightCity` | Known | Event handler |
| 0x0071BF6E | `HandleSelectCity` | Known | Event handler |
| 0x0071BFDA | `HandleHighlightCity` | Known | Event handler |
| 0x0071C013 | `HandleSelectCity` | Known | Event handler |
| 0x0071C07F | `HandleHighlightCity` | Known | Event handler |
| 0x0071C0B8 | `HandleSelectCity` | Known | Event handler |
| 0x0071C124 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C162 | `HandleSelectCity` | Known | Event handler |
| 0x0071C1CE | `HandleHighlightCity` | Known | Event handler |
| 0x0071C207 | `HandleSelectCity` | Known | Event handler |
| 0x0071C273 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C2AC | `HandleSelectCity` | Known | Event handler |
| 0x0071C318 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C351 | `HandleSelectCity` | Known | Event handler |
| 0x0071C3BD | `HandleHighlightCity` | Known | Event handler |
| 0x0071C3F6 | `HandleSelectCity` | Known | Event handler |
| 0x0071C462 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C49B | `HandleSelectCity` | Known | Event handler |
| 0x0071C507 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C540 | `HandleSelectCity` | Known | Event handler |
| 0x0071C5AC | `HandleHighlightCity` | Known | Event handler |
| 0x0071C5E5 | `HandleSelectCity` | Known | Event handler |
| 0x0071C651 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C68A | `HandleSelectCity` | Known | Event handler |
| 0x0071C6F6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C72F | `HandleSelectCity` | Known | Event handler |
| 0x0071C79B | `HandleHighlightCity` | Known | Event handler |
| 0x0071C7D4 | `HandleSelectCity` | Known | Event handler |
| 0x0071C840 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C879 | `HandleSelectCity` | Known | Event handler |
| 0x0071C8E5 | `HandleHighlightCity` | Known | Event handler |
| 0x0071C91E | `HandleSelectCity` | Known | Event handler |
| 0x0071C98A | `HandleHighlightCity` | Known | Event handler |
| 0x0071C9C3 | `HandleSelectCity` | Known | Event handler |
| 0x0071CA2F | `HandleHighlightCity` | Known | Event handler |
| 0x0071CA68 | `HandleSelectCity` | Known | Event handler |
| 0x0071CAD4 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CB0D | `HandleSelectCity` | Known | Event handler |
| 0x0071CB79 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CBB2 | `HandleSelectCity` | Known | Event handler |
| 0x0071CC1E | `HandleHighlightCity` | Known | Event handler |
| 0x0071CC57 | `HandleSelectCity` | Known | Event handler |
| 0x0071CCC3 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CCFC | `HandleSelectCity` | Known | Event handler |
| 0x0071CD68 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CDA1 | `HandleSelectCity` | Known | Event handler |
| 0x0071CE0D | `HandleHighlightCity` | Known | Event handler |
| 0x0071CE46 | `HandleSelectCity` | Known | Event handler |
| 0x0071CEB2 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CEEB | `HandleSelectCity` | Known | Event handler |
| 0x0071CF57 | `HandleHighlightCity` | Known | Event handler |
| 0x0071CF90 | `HandleSelectCity` | Known | Event handler |
| 0x0071CFFC | `HandleHighlightCity` | Known | Event handler |
| 0x0071D035 | `HandleSelectCity` | Known | Event handler |
| 0x0071D0A1 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D0DA | `HandleSelectCity` | Known | Event handler |
| 0x0071D146 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D17F | `HandleSelectCity` | Known | Event handler |
| 0x0071D1EB | `HandleHighlightCity` | Known | Event handler |
| 0x0071D224 | `HandleSelectCity` | Known | Event handler |
| 0x0071D290 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D2C9 | `HandleSelectCity` | Known | Event handler |
| 0x0071D335 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D36E | `HandleSelectCity` | Known | Event handler |
| 0x0071D3DA | `HandleHighlightCity` | Known | Event handler |
| 0x0071D413 | `HandleSelectCity` | Known | Event handler |
| 0x0071D47F | `HandleHighlightCity` | Known | Event handler |
| 0x0071D4B8 | `HandleSelectCity` | Known | Event handler |
| 0x0071D524 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D55D | `HandleSelectCity` | Known | Event handler |
| 0x0071D5C9 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D602 | `HandleSelectCity` | Known | Event handler |
| 0x0071D66E | `HandleHighlightCity` | Known | Event handler |
| 0x0071D6A7 | `HandleSelectCity` | Known | Event handler |
| 0x0071D713 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D74C | `HandleSelectCity` | Known | Event handler |
| 0x0071D7B8 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D7F1 | `HandleSelectCity` | Known | Event handler |
| 0x0071D85D | `HandleHighlightCity` | Known | Event handler |
| 0x0071D896 | `HandleSelectCity` | Known | Event handler |
| 0x0071D902 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D93B | `HandleSelectCity` | Known | Event handler |
| 0x0071D9A7 | `HandleHighlightCity` | Known | Event handler |
| 0x0071D9E0 | `HandleSelectCity` | Known | Event handler |
| 0x0071DA4C | `HandleHighlightCity` | Known | Event handler |
| 0x0071DA85 | `HandleSelectCity` | Known | Event handler |
| 0x0071DAF1 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DB2A | `HandleSelectCity` | Known | Event handler |
| 0x0071DB96 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DBCF | `HandleSelectCity` | Known | Event handler |
| 0x0071DC3B | `HandleHighlightCity` | Known | Event handler |
| 0x0071DC74 | `HandleSelectCity` | Known | Event handler |
| 0x0071DCE0 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DD19 | `HandleSelectCity` | Known | Event handler |
| 0x0071DD85 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DDBE | `HandleSelectCity` | Known | Event handler |
| 0x0071DE2A | `HandleHighlightCity` | Known | Event handler |
| 0x0071DE63 | `HandleSelectCity` | Known | Event handler |
| 0x0071DECF | `HandleHighlightCity` | Known | Event handler |
| 0x0071DF08 | `HandleSelectCity` | Known | Event handler |
| 0x0071DF74 | `HandleHighlightCity` | Known | Event handler |
| 0x0071DFAD | `HandleSelectCity` | Known | Event handler |
| 0x0071E019 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E052 | `HandleSelectCity` | Known | Event handler |
| 0x0071E0BE | `HandleHighlightCity` | Known | Event handler |
| 0x0071E0F7 | `HandleSelectCity` | Known | Event handler |
| 0x0071E163 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E19C | `HandleSelectCity` | Known | Event handler |
| 0x0071E208 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E241 | `HandleSelectCity` | Known | Event handler |
| 0x0071E2AD | `HandleHighlightCity` | Known | Event handler |
| 0x0071E2E6 | `HandleSelectCity` | Known | Event handler |
| 0x0071E352 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E38B | `HandleSelectCity` | Known | Event handler |
| 0x0071E3F7 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E430 | `HandleSelectCity` | Known | Event handler |
| 0x0071E49C | `HandleHighlightCity` | Known | Event handler |
| 0x0071E4D5 | `HandleSelectCity` | Known | Event handler |
| 0x0071E541 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E57A | `HandleSelectCity` | Known | Event handler |
| 0x0071E5E6 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E626 | `HandleSelectCity` | Known | Event handler |
| 0x0071E692 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E6CB | `HandleSelectCity` | Known | Event handler |
| 0x0071E737 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E770 | `HandleSelectCity` | Known | Event handler |
| 0x0071E7DC | `HandleHighlightCity` | Known | Event handler |
| 0x0071E81A | `HandleSelectCity` | Known | Event handler |
| 0x0071E886 | `HandleHighlightCity` | Known | Event handler |
| 0x0071E8BF | `HandleSelectCity` | Known | Event handler |
| 0x0071E92B | `HandleHighlightCity` | Known | Event handler |
| 0x0071E964 | `HandleSelectCity` | Known | Event handler |
| 0x0071E9D0 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EA09 | `HandleSelectCity` | Known | Event handler |
| 0x0071EA75 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EAAE | `HandleSelectCity` | Known | Event handler |
| 0x0071EB1A | `HandleHighlightCity` | Known | Event handler |
| 0x0071EB53 | `HandleSelectCity` | Known | Event handler |
| 0x0071EBBF | `HandleHighlightCity` | Known | Event handler |
| 0x0071EBF8 | `HandleSelectCity` | Known | Event handler |
| 0x0071EC64 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EC9D | `HandleSelectCity` | Known | Event handler |
| 0x0071ED09 | `HandleHighlightCity` | Known | Event handler |
| 0x0071ED46 | `HandleSelectCity` | Known | Event handler |
| 0x0071EDB2 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EDEB | `HandleSelectCity` | Known | Event handler |
| 0x0071EE57 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EE90 | `HandleSelectCity` | Known | Event handler |
| 0x0071EEFC | `HandleHighlightCity` | Known | Event handler |
| 0x0071EF35 | `HandleSelectCity` | Known | Event handler |
| 0x0071EFA1 | `HandleHighlightCity` | Known | Event handler |
| 0x0071EFDA | `HandleSelectCity` | Known | Event handler |
| 0x0071F046 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F07F | `HandleSelectCity` | Known | Event handler |
| 0x0071F0EB | `HandleHighlightCity` | Known | Event handler |
| 0x0071F124 | `HandleSelectCity` | Known | Event handler |
| 0x0071F190 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F1C9 | `HandleSelectCity` | Known | Event handler |
| 0x0071F235 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F26E | `HandleSelectCity` | Known | Event handler |
| 0x0071F2DA | `HandleHighlightCity` | Known | Event handler |
| 0x0071F313 | `HandleSelectCity` | Known | Event handler |
| 0x0071F37F | `HandleHighlightCity` | Known | Event handler |
| 0x0071F3B8 | `HandleSelectCity` | Known | Event handler |
| 0x0071F424 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F45D | `HandleSelectCity` | Known | Event handler |
| 0x0071F4C9 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F502 | `HandleSelectCity` | Known | Event handler |
| 0x0071F56E | `HandleHighlightCity` | Known | Event handler |
| 0x0071F5A7 | `HandleSelectCity` | Known | Event handler |
| 0x0071F613 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F64C | `HandleSelectCity` | Known | Event handler |
| 0x0071F6B8 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F6F1 | `HandleSelectCity` | Known | Event handler |
| 0x0071F75D | `HandleHighlightCity` | Known | Event handler |
| 0x0071F796 | `HandleSelectCity` | Known | Event handler |
| 0x0071F802 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F83B | `HandleSelectCity` | Known | Event handler |
| 0x0071F8A7 | `HandleHighlightCity` | Known | Event handler |
| 0x0071F8E0 | `HandleSelectCity` | Known | Event handler |
| 0x0071F94C | `HandleHighlightCity` | Known | Event handler |
| 0x0071F985 | `HandleSelectCity` | Known | Event handler |
| 0x0071F9F1 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FA2A | `HandleSelectCity` | Known | Event handler |
| 0x0071FA96 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FACF | `HandleSelectCity` | Known | Event handler |
| 0x0071FB3B | `HandleHighlightCity` | Known | Event handler |
| 0x0071FB74 | `HandleSelectCity` | Known | Event handler |
| 0x0071FBE0 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FC19 | `HandleSelectCity` | Known | Event handler |
| 0x0071FC85 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FCBE | `HandleSelectCity` | Known | Event handler |
| 0x0071FD2A | `HandleHighlightCity` | Known | Event handler |
| 0x0071FD63 | `HandleSelectCity` | Known | Event handler |
| 0x0071FDCF | `HandleHighlightCity` | Known | Event handler |
| 0x0071FE08 | `HandleSelectCity` | Known | Event handler |
| 0x0071FE74 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FEAD | `HandleSelectCity` | Known | Event handler |
| 0x0071FF19 | `HandleHighlightCity` | Known | Event handler |
| 0x0071FF52 | `HandleSelectCity` | Known | Event handler |
| 0x0071FFBE | `HandleHighlightCity` | Known | Event handler |
| 0x0071FFF7 | `HandleSelectCity` | Known | Event handler |
| 0x00720063 | `HandleHighlightCity` | Known | Event handler |
| 0x0072009C | `HandleSelectCity` | Known | Event handler |
| 0x00720108 | `HandleHighlightCity` | Known | Event handler |
| 0x00720141 | `HandleSelectCity` | Known | Event handler |
| 0x007201AD | `HandleHighlightCity` | Known | Event handler |
| 0x007201E6 | `HandleSelectCity` | Known | Event handler |
| 0x00720252 | `HandleHighlightCity` | Known | Event handler |
| 0x0072028B | `HandleSelectCity` | Known | Event handler |
| 0x007202F7 | `HandleHighlightCity` | Known | Event handler |
| 0x00720336 | `HandleSelectCity` | Known | Event handler |
| 0x007203A2 | `HandleHighlightCity` | Known | Event handler |
| 0x007203DB | `HandleSelectCity` | Known | Event handler |
| 0x00720447 | `HandleHighlightCity` | Known | Event handler |
| 0x00720480 | `HandleSelectCity` | Known | Event handler |
| 0x007204EC | `HandleHighlightCity` | Known | Event handler |
| 0x00720525 | `HandleSelectCity` | Known | Event handler |
| 0x00720591 | `HandleHighlightCity` | Known | Event handler |
| 0x007205CA | `HandleSelectCity` | Known | Event handler |
| 0x00720636 | `HandleHighlightCity` | Known | Event handler |
| 0x0072066F | `HandleSelectCity` | Known | Event handler |
| 0x007206DB | `HandleHighlightCity` | Known | Event handler |
| 0x00720714 | `HandleSelectCity` | Known | Event handler |
| 0x00720780 | `HandleHighlightCity` | Known | Event handler |
| 0x007207B9 | `HandleSelectCity` | Known | Event handler |
| 0x00720825 | `HandleHighlightCity` | Known | Event handler |
| 0x0072085E | `HandleSelectCity` | Known | Event handler |
| 0x007208CA | `HandleHighlightCity` | Known | Event handler |
| 0x00720903 | `HandleSelectCity` | Known | Event handler |
| 0x0072096F | `HandleHighlightCity` | Known | Event handler |
| 0x007209A8 | `HandleSelectCity` | Known | Event handler |
| 0x00720A14 | `HandleHighlightCity` | Known | Event handler |
| 0x00720A4D | `HandleSelectCity` | Known | Event handler |
| 0x00720AB9 | `HandleHighlightCity` | Known | Event handler |
| 0x00720AF2 | `HandleSelectCity` | Known | Event handler |
| 0x00720B5E | `HandleHighlightCity` | Known | Event handler |
| 0x00720B97 | `HandleSelectCity` | Known | Event handler |
| 0x00720C03 | `HandleHighlightCity` | Known | Event handler |
| 0x00720C3C | `HandleSelectCity` | Known | Event handler |
| 0x00720CA8 | `HandleHighlightCity` | Known | Event handler |
| 0x00720CE1 | `HandleSelectCity` | Known | Event handler |
| 0x00720D4D | `HandleHighlightCity` | Known | Event handler |
| 0x00720D86 | `HandleSelectCity` | Known | Event handler |
| 0x00720DF2 | `HandleHighlightCity` | Known | Event handler |
| 0x00720E2B | `HandleSelectCity` | Known | Event handler |
| 0x00720E97 | `HandleHighlightCity` | Known | Event handler |
| 0x00720ED0 | `HandleSelectCity` | Known | Event handler |
| 0x00720F3C | `HandleHighlightCity` | Known | Event handler |
| 0x00720F75 | `HandleSelectCity` | Known | Event handler |
| 0x00720FE1 | `HandleHighlightCity` | Known | Event handler |
| 0x0072101A | `HandleSelectCity` | Known | Event handler |
| 0x00721086 | `HandleHighlightCity` | Known | Event handler |
| 0x007210BF | `HandleSelectCity` | Known | Event handler |
| 0x0072112B | `HandleHighlightCity` | Known | Event handler |
| 0x00721164 | `HandleSelectCity` | Known | Event handler |
| 0x007211D0 | `HandleHighlightCity` | Known | Event handler |
| 0x00721209 | `HandleSelectCity` | Known | Event handler |
| 0x00721275 | `HandleHighlightCity` | Known | Event handler |
| 0x007212AE | `HandleSelectCity` | Known | Event handler |
| 0x0072131A | `HandleHighlightCity` | Known | Event handler |
| 0x00721353 | `HandleSelectCity` | Known | Event handler |
| 0x007213BF | `HandleHighlightCity` | Known | Event handler |
| 0x007213F8 | `HandleSelectCity` | Known | Event handler |
| 0x00721464 | `HandleHighlightCity` | Known | Event handler |
| 0x0072149D | `HandleSelectCity` | Known | Event handler |
| 0x00721509 | `HandleHighlightCity` | Known | Event handler |
| 0x00721542 | `HandleSelectCity` | Known | Event handler |
| 0x007215AE | `HandleHighlightCity` | Known | Event handler |
| 0x007215E7 | `HandleSelectCity` | Known | Event handler |
| 0x00721653 | `HandleHighlightCity` | Known | Event handler |
| 0x0072168C | `HandleSelectCity` | Known | Event handler |
| 0x007216F8 | `HandleHighlightCity` | Known | Event handler |
| 0x00721731 | `HandleSelectCity` | Known | Event handler |
| 0x0072179D | `HandleHighlightCity` | Known | Event handler |
| 0x007217D6 | `HandleSelectCity` | Known | Event handler |
| 0x00721842 | `HandleHighlightCity` | Known | Event handler |
| 0x0072187B | `HandleSelectCity` | Known | Event handler |
| 0x007218E7 | `HandleHighlightCity` | Known | Event handler |
| 0x00721920 | `HandleSelectCity` | Known | Event handler |
| 0x0072198C | `HandleHighlightCity` | Known | Event handler |
| 0x007219C5 | `HandleSelectCity` | Known | Event handler |
| 0x00721A31 | `HandleHighlightCity` | Known | Event handler |
| 0x00721A6A | `HandleSelectCity` | Known | Event handler |
| 0x00721AD6 | `HandleHighlightCity` | Known | Event handler |
| 0x00721B0F | `HandleSelectCity` | Known | Event handler |
| 0x00721B7B | `HandleHighlightCity` | Known | Event handler |
| 0x00721BB4 | `HandleSelectCity` | Known | Event handler |
| 0x00721C20 | `HandleHighlightCity` | Known | Event handler |
| 0x00721C59 | `HandleSelectCity` | Known | Event handler |
| 0x00721CC5 | `HandleHighlightCity` | Known | Event handler |
| 0x00721CFE | `HandleSelectCity` | Known | Event handler |
| 0x00721D6A | `HandleHighlightCity` | Known | Event handler |
| 0x00721DA3 | `HandleSelectCity` | Known | Event handler |
| 0x00721E0F | `HandleHighlightCity` | Known | Event handler |
| 0x00721E48 | `HandleSelectCity` | Known | Event handler |
| 0x00721EB4 | `HandleHighlightCity` | Known | Event handler |
| 0x00721EED | `HandleSelectCity` | Known | Event handler |
| 0x00721F59 | `HandleHighlightCity` | Known | Event handler |
| 0x00721F92 | `HandleSelectCity` | Known | Event handler |
| 0x00721FFE | `HandleHighlightCity` | Known | Event handler |
| 0x00722037 | `HandleSelectCity` | Known | Event handler |
| 0x007220A3 | `HandleHighlightCity` | Known | Event handler |
| 0x007220DC | `HandleSelectCity` | Known | Event handler |
| 0x00722148 | `HandleHighlightCity` | Known | Event handler |
| 0x00722181 | `HandleSelectCity` | Known | Event handler |
| 0x007221ED | `HandleHighlightCity` | Known | Event handler |
| 0x00722226 | `HandleSelectCity` | Known | Event handler |
| 0x00722292 | `HandleHighlightCity` | Known | Event handler |
| 0x007222CB | `HandleSelectCity` | Known | Event handler |
| 0x00722337 | `HandleHighlightCity` | Known | Event handler |
| 0x00722376 | `HandleSelectCity` | Known | Event handler |
| 0x007223E2 | `HandleHighlightCity` | Known | Event handler |
| 0x0072241B | `HandleSelectCity` | Known | Event handler |
| 0x00722487 | `HandleHighlightCity` | Known | Event handler |
| 0x007224C0 | `HandleSelectCity` | Known | Event handler |
| 0x0072252C | `HandleHighlightCity` | Known | Event handler |
| 0x00722565 | `HandleSelectCity` | Known | Event handler |
| 0x007225D1 | `HandleHighlightCity` | Known | Event handler |
| 0x0072260A | `HandleSelectCity` | Known | Event handler |
| 0x00722676 | `HandleHighlightCity` | Known | Event handler |
| 0x007226B6 | `HandleSelectCity` | Known | Event handler |
| 0x00722722 | `HandleHighlightCity` | Known | Event handler |
| 0x0072275B | `HandleSelectCity` | Known | Event handler |
| 0x007227C7 | `HandleHighlightCity` | Known | Event handler |
| 0x00722800 | `HandleSelectCity` | Known | Event handler |
| 0x0072286C | `HandleHighlightCity` | Known | Event handler |
| 0x007228A5 | `HandleSelectCity` | Known | Event handler |
| 0x00722911 | `HandleHighlightCity` | Known | Event handler |
| 0x0072294A | `HandleSelectCity` | Known | Event handler |
| 0x007229B6 | `HandleHighlightCity` | Known | Event handler |
| 0x007229EF | `HandleSelectCity` | Known | Event handler |
| 0x00722A5B | `HandleHighlightCity` | Known | Event handler |
| 0x00722A94 | `HandleSelectCity` | Known | Event handler |
| 0x00722B00 | `HandleHighlightCity` | Known | Event handler |
| 0x00722B39 | `HandleSelectCity` | Known | Event handler |
| 0x00722BA5 | `HandleHighlightCity` | Known | Event handler |
| 0x00722BDE | `HandleSelectCity` | Known | Event handler |
| 0x00722C4A | `HandleHighlightCity` | Known | Event handler |
| 0x00722C83 | `HandleSelectCity` | Known | Event handler |
| 0x00722CEF | `HandleHighlightCity` | Known | Event handler |
| 0x00722D28 | `HandleSelectCity` | Known | Event handler |
| 0x00722D94 | `HandleHighlightCity` | Known | Event handler |
| 0x00722DCD | `HandleSelectCity` | Known | Event handler |
| 0x00722E39 | `HandleHighlightCity` | Known | Event handler |
| 0x00722E72 | `HandleSelectCity` | Known | Event handler |
| 0x00722EDE | `HandleHighlightCity` | Known | Event handler |
| 0x00722F17 | `HandleSelectCity` | Known | Event handler |
| 0x00722F83 | `HandleHighlightCity` | Known | Event handler |
| 0x00722FBC | `HandleSelectCity` | Known | Event handler |
| 0x00723028 | `HandleHighlightCity` | Known | Event handler |
| 0x00723061 | `HandleSelectCity` | Known | Event handler |
| 0x007230CD | `HandleHighlightCity` | Known | Event handler |
| 0x00723106 | `HandleSelectCity` | Known | Event handler |
| 0x00723172 | `HandleHighlightCity` | Known | Event handler |
| 0x0072366A | `HandleMusicSelected` | Known | Event handler |
| 0x007236AC | `HandleMusicHilited` | Known | Event handler |
| 0x007236E4 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x0072372A | `HandleMusicHilited` | Known | Event handler |
| 0x00723762 | `HandleGotoGenius` | Known | Event handler |
| 0x007237A1 | `HandleGeniusHilited` | Known | Event handler |
| 0x007237DA | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00723820 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x0072385C | `HandleArtistsSelected` | Known | Event handler |
| 0x007238A0 | `HandleArtistsHilited` | Known | Event handler |
| 0x007238DA | `HandleAlbumsSelected` | Known | Event handler |
| 0x0072391D | `HandleAlbumsHilited` | Known | Event handler |
| 0x00723956 | `HandleCompilationsSelected` | Known | Event handler |
| 0x0072399F | `HandleCompilationsHilited` | Known | Event handler |
| 0x007239DE | `HandleSongsSelected` | Known | Event handler |
| 0x00723A20 | `HandleSongsHilited` | Known | Event handler |
| 0x00723A58 | `HandleGenresSelected` | Known | Event handler |
| 0x00723A9B | `HandleGenresHilited` | Known | Event handler |
| 0x00723AD4 | `HandleComposersSelected` | Known | Event handler |
| 0x00723B1A | `HandleComposersHilited` | Known | Event handler |
| 0x00723B56 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x00723B9D | `HandleAudiobooksHilited` | Known | Event handler |
| 0x00723C5C | `HandleMusicHilited` | Known | Event handler |
| 0x00723C94 | `HandleVideosSelected` | Known | Event handler |
| 0x00723CD7 | `HandleVideosHilited` | Known | Event handler |
| 0x00723D10 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00723D5B | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00723D9C | `HandleMoviesSelected` | Known | Event handler |
| 0x00723DDF | `HandleMoviesHilited` | Known | Event handler |
| 0x00723E18 | `HandleTVShowsSelected` | Known | Event handler |
| 0x00723E5C | `HandleTVShowsHilited` | Known | Event handler |
| 0x00723E96 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00723EDE | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00723F1C | `HandleRentalsSelected` | Known | Event handler |
| 0x00723F60 | `HandleRentalsHilited` | Known | Event handler |
| 0x00723F9A | `HandlePhotosSelected` | Known | Event handler |
| 0x00723FDD | `HandlePhotosHilited` | Known | Event handler |
| 0x00724016 | `HandlePhotosSelected` | Known | Event handler |
| 0x00724059 | `HandlePhotosHilited` | Known | Event handler |
| 0x00724092 | `HandlePodcastsSelected` | Known | Event handler |
| 0x007240D7 | `HandlePodcastsHilited` | Known | Event handler |
| 0x0072418A | `HandleGenericHilited` | Known | Event handler |
| 0x00724283 | `HandleGenericHilited` | Known | Event handler |
| 0x00724768 | `HandleLock` | Known | Event handler |
| 0x007248D9 | `HandleNikePlusSelected` | Known | Event handler |
| 0x0072491E | `HandleGenericHilited` | Known | Event handler |
| 0x00724A24 | `HandleGenericHilited` | Known | Event handler |
| 0x00724B23 | `HandleGenericHilited` | Known | Event handler |
| 0x00724C10 | `HandleGenericHilited` | Known | Event handler |
| 0x00724D0D | `HandleGenericHilited` | Known | Event handler |
| 0x00724D87 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00724DD0 | `HandleGenericHilited` | Known | Event handler |
| 0x00724E49 | `HandleBacklightSelected` | Known | Event handler |
| 0x00724E8F | `HandleGenericHilited` | Known | Event handler |
| 0x00724F0A | `HandleSleepSelected` | Known | Event handler |
| 0x00724F4C | `HandleGenericHilited` | Known | Event handler |
| 0x00724FC3 | `HandleNowPlaying` | Known | Event handler |
| 0x0072503B | `HandleNowPlayingHilited` | Known | Event handler |
| 0x0072507E | `HandleCoverFlowSelected` | Known | Event handler |
| 0x007250C4 | `HandleMusicHilited` | Known | Event handler |
| 0x007250FC | `HandleGotoGenius` | Known | Event handler |
| 0x00725132 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x00725178 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x007251B6 | `HandleArtistsSelected` | Known | Event handler |
| 0x007251FA | `HandleArtistsHilited` | Known | Event handler |
| 0x00725234 | `HandleAlbumsSelected` | Known | Event handler |
| 0x00725277 | `HandleAlbumsHilited` | Known | Event handler |
| 0x007252B0 | `HandleCompilationsSelected` | Known | Event handler |
| 0x007252F9 | `HandleCompilationsHilited` | Known | Event handler |
| 0x00725338 | `HandleSongsSelected` | Known | Event handler |
| 0x0072537A | `HandleSongsHilited` | Known | Event handler |
| 0x00725425 | `HandleGenericHilited` | Known | Event handler |
| 0x0072549D | `HandleGenresSelected` | Known | Event handler |
| 0x007254E0 | `HandleGenresHilited` | Known | Event handler |
| 0x00725519 | `HandleComposersSelected` | Known | Event handler |
| 0x0072555F | `HandleComposersHilited` | Known | Event handler |
| 0x0072559B | `HandleAudiobooksSelected` | Known | Event handler |
| 0x007255E2 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x007256A1 | `HandleMusicHilited` | Known | Event handler |
| 0x00725715 | `HandlePlayPause` | Known | Event handler |
| 0x0072574A | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x00725834 | `HandleSelect` | Known | Event handler |
| 0x0072587A | `HandleMoviesSelected` | Known | Event handler |
| 0x007258BD | `HandleMoviesHilited` | Known | Event handler |
| 0x007258F6 | `HandleRentalsSelected` | Known | Event handler |
| 0x0072593A | `HandleRentalsHilited` | Known | Event handler |
| 0x00725974 | `HandleTVShowsSelected` | Known | Event handler |
| 0x007259B8 | `HandleTVShowsHilited` | Known | Event handler |
| 0x007259F2 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00725A3A | `HandleMusicVideosHilited` | Known | Event handler |
| 0x00725A78 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x00725AC3 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00725B89 | `HandleVideosHilited` | Known | Event handler |
| 0x007261D7 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x00726D5E | `HandleMainMenu` | Known | Event handler |
| 0x00726D96 | `HandleMusicMenu` | Known | Event handler |
| 0x007272BE | `HandleRadioRegion` | Known | Event handler |
| 0x00727362 | `HandleLanguage` | Known | Event handler |
| 0x00727468 | `HandleNew` | Known | Event handler |
| 0x007274E3 | `HandleClear` | Known | Event handler |
| 0x00727514 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x007275D0 | `HandleSelectIndexedSession` | Known | Event handler |
| 0x00727739 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x0072778C | `HandleSelect` | Known | Event handler |
| 0x007278B6 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x007278F0 | `HandleEQSettingSelected` | Known | Event handler |
| 0x00727928 | `HandleEQSettingSelected` | Known | Event handler |
| 0x0073A862 | `HandleMenuSelection` | Known | Event handler |
| 0x0073ABA7 | `HandleLoadingCancelled` | Known | Event handler |
| 0x0073AC43 | `HandleLoadingCancelled` | Known | Event handler |
| 0x0073AD10 | `HandleItemSelected` | Known | Event handler |
| 0x0073AE5B | `HandleNextContact` | Known | Event handler |
| 0x0073AE87 | `HandlePreviousContact` | Known | Event handler |
| 0x0073AEBD | `HandleSelectKey` | Known | Event handler |
| 0x0073B4CE | `HandleSelect` | Known | Event handler |
| 0x0073B7F5 | `HandleDateChosen` | Known | Event handler |
| 0x0073B82B | `HandleTimeChosen` | Known | Event handler |
| 0x0073B861 | `HandleFrequencyChosen` | Known | Event handler |
| 0x0073B89C | `HandleSoundChosen` | Known | Event handler |
| 0x0073B8D3 | `HandleLabelChosen` | Known | Event handler |
| 0x0073B90A | `HandleDeleteChosen` | Known | Event handler |
| 0x0073B946 | `HandleSelect` | Known | Event handler |
| 0x0073B97E | `HandleSelect` | Known | Event handler |
| 0x0073BCBF | `HandleLeaveAlarm` | Known | Event handler |
| 0x0073BCEC | `HandleLeaveAlarm` | Known | Event handler |
| 0x0073BD1B | `HandleLeaveAlarm` | Known | Event handler |
| 0x0073BD48 | `HandleLeaveAlarm` | Known | Event handler |
| 0x0073BE82 | `HandleSelect` | Known | Event handler |
| 0x0073BEB0 | `HandleSelect` | Known | Event handler |
| 0x0073C00F | `HandleNextDay` | Known | Event handler |
| 0x0073C037 | `HandlePreviousDay` | Known | Event handler |
| 0x0073C1E6 | `HandleSelect` | Known | Event handler |
| 0x0073C213 | `HandleNextDay` | Known | Event handler |
| 0x0073C23B | `HandlePreviousDay` | Known | Event handler |
| 0x0073C3E3 | `HandleNextDay` | Known | Event handler |
| 0x0073C40B | `HandlePreviousDay` | Known | Event handler |
| 0x0073C4CC | `HandleSelect` | Known | Event handler |
| 0x0073C4F7 | `HandleNextDay` | Known | Event handler |
| 0x0073C51F | `HandlePreviousDay` | Known | Event handler |
| 0x0073C696 | `HandleSelectLozinch` | Known | Event handler |
| 0x0073C80E | `HandleSelectLozinch` | Known | Event handler |
| 0x0073C92D | `HandleFlowNext` | Known | Event handler |
| 0x0073C95B | `HandlePlayPause` | Known | Event handler |
| 0x0073C9AA | `HandleFlowPrev` | Known | Event handler |
| 0x0073C9D5 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x0073CAC9 | `HandleAlbumSelected` | Known | Event handler |
| 0x0073CC64 | `HandleFlowNext` | Known | Event handler |
| 0x0073CCB2 | `HandleFlowNext` | Known | Event handler |
| 0x0073CCE0 | `HandlePlayPause` | Known | Event handler |
| 0x0073CD2F | `HandleFlowPrev` | Known | Event handler |
| 0x0073CD5B | `HandleFlowPrev` | Known | Event handler |
| 0x0073CD7B | `HandleFlowWheel` | Known | Event handler |
| 0x0073D10B | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0073D536 | `HandleArrowDown` | Known | Event handler |
| 0x0073D5A0 | `HandleArrowUp` | Known | Event handler |
| 0x0073D5BF | `HandleWheel` | Known | Event handler |
| 0x0073D648 | `HandleSelect` | Known | Event handler |
| 0x0073D6C5 | `HandleGameHilited` | Known | Event handler |
| 0x00740B2B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007428C7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00744663 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007463FF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074819B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00749F37 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074BCD3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074DA6F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0074F80B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007515A7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00753343 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007550DF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00756E7B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00758C17 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075A9B3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075C74F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0075E4EB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00760287 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00762023 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00763DBF | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00765B5B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007678F7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00769693 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076B42F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076D1CB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0076EF67 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00770D03 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00772A9F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077483B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x007765D7 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00778373 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077A10F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077BEAB | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077DC47 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0077F9E3 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078177F | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078351B | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078529C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00785F18 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00786B94 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00787810 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078848C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00789108 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00789D84 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078AA00 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078B67C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078C2F8 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078CF74 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078DBF0 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x0078E86C | `HandlePlayPause` | Known | Event handler |
| 0x0078E8A2 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078E8E4 | `HandleAddToOTG` | Known | Event handler |
| 0x0078EA81 | `HandlePlayPause` | Known | Event handler |
| 0x0078EAA8 | `HandleSelect` | Known | Event handler |
| 0x0078EAD5 | `HandleHilite` | Known | Event handler |
| 0x0078EB08 | `HandlePlayPause` | Known | Event handler |
| 0x0078EB9B | `HandlePlayPause` | Known | Event handler |
| 0x0078EBC2 | `HandleSelect` | Known | Event handler |
| 0x0078EC28 | `HandleHilite` | Known | Event handler |
| 0x0078EC5A | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x0078ECA4 | `HandlePlayPause` | Known | Event handler |
| 0x0078ECDA | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078ED21 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0078ED64 | `HandleAddToOTG` | Known | Event handler |
| 0x0078EDC7 | `HandleStartGenius` | Known | Event handler |
| 0x0078EE03 | `HandleViewAlbum` | Known | Event handler |
| 0x0078EE3E | `HandleViewArtist` | Known | Event handler |
| 0x0078EE7F | `HandleViewCompilation` | Known | Event handler |
| 0x0078F01F | `HandlePlayPause` | Known | Event handler |
| 0x0078F046 | `HandleSelect` | Known | Event handler |
| 0x0078F0B0 | `HandlePlayPause` | Known | Event handler |
| 0x0078F0E6 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078F12D | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0078F170 | `HandleAddToOTG` | Known | Event handler |
| 0x0078F1D3 | `HandleStartGenius` | Known | Event handler |
| 0x0078F20F | `HandleViewAlbum` | Known | Event handler |
| 0x0078F24A | `HandleViewArtist` | Known | Event handler |
| 0x0078F28B | `HandleViewCompilation` | Known | Event handler |
| 0x0078F42B | `HandlePlayPause` | Known | Event handler |
| 0x0078F452 | `HandleSelect` | Known | Event handler |
| 0x0078F4BC | `HandlePlayPause` | Known | Event handler |
| 0x0078F4FA | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0078F53D | `HandleAddToOTG` | Known | Event handler |
| 0x0078F5A0 | `HandleStartGenius` | Known | Event handler |
| 0x0078F5DC | `HandleViewAlbum` | Known | Event handler |
| 0x0078F617 | `HandleViewArtist` | Known | Event handler |
| 0x0078F658 | `HandleViewCompilation` | Known | Event handler |
| 0x0078F7EB | `HandleSelect` | Known | Event handler |
| 0x0078F850 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0078F894 | `HandlePlayPause` | Known | Event handler |
| 0x0078F8CA | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078F90C | `HandleAddToOTG` | Known | Event handler |
| 0x0078FB66 | `HandlePlayPause` | Known | Event handler |
| 0x0078FB8D | `HandleSelect` | Known | Event handler |
| 0x0078FBBA | `HandleHilite` | Known | Event handler |
| 0x0078FBEC | `HandlePlayPause` | Known | Event handler |
| 0x0078FC22 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078FC64 | `HandleAddToOTG` | Known | Event handler |
| 0x0078FEBE | `HandlePlayPause` | Known | Event handler |
| 0x0078FEE5 | `HandleSelect` | Known | Event handler |
| 0x0078FF12 | `HandleHilite` | Known | Event handler |
| 0x0078FF44 | `HandlePlayPause` | Known | Event handler |
| 0x0078FF7A | `HandleShowContextualMenu` | Known | Event handler |
| 0x0078FFBC | `HandleAddToOTG` | Known | Event handler |
| 0x007902CF | `HandlePlayPause` | Known | Event handler |
| 0x007902F6 | `HandleSelect` | Known | Event handler |
| 0x00790328 | `HandlePlayPause` | Known | Event handler |
| 0x0079035E | `HandleShowContextualMenu` | Known | Event handler |
| 0x007903A0 | `HandleAddToOTG` | Known | Event handler |
| 0x0079045A | `HandlePlayPause` | Known | Event handler |
| 0x00790481 | `HandleSelect` | Known | Event handler |
| 0x00790510 | `HandlePlayPause` | Known | Event handler |
| 0x00790546 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00790588 | `HandleAddToOTG` | Known | Event handler |
| 0x00790769 | `HandlePlayPause` | Known | Event handler |
| 0x00790790 | `HandleSelect` | Known | Event handler |
| 0x007907C0 | `HandlePlayPause` | Known | Event handler |
| 0x007907F6 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00790838 | `HandleAddToOTG` | Known | Event handler |
| 0x007908E5 | `HandleSelect` | Known | Event handler |
| 0x0079097E | `HandleHilite` | Known | Event handler |
| 0x007909AA | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007909EC | `HandlePlayPause` | Known | Event handler |
| 0x00790A22 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00790A64 | `HandleAddToOTG` | Known | Event handler |
| 0x00790B11 | `HandleSelect` | Known | Event handler |
| 0x00790B76 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00790BB8 | `HandlePlayPause` | Known | Event handler |
| 0x00790D5C | `HandleSelect` | Known | Event handler |
| 0x00790D89 | `HandleHilite` | Known | Event handler |
| 0x00790DB5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00790DF8 | `HandlePlayPause` | Known | Event handler |
| 0x00790E7E | `HandleSelect` | Known | Event handler |
| 0x00790F0C | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00790F50 | `HandlePlayPause` | Known | Event handler |
| 0x00790FD6 | `HandleSelect` | Known | Event handler |
| 0x0079103B | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079107C | `HandlePlayPause` | Known | Event handler |
| 0x00791102 | `HandleSelect` | Known | Event handler |
| 0x00791168 | `HandleHilite` | Known | Event handler |
| 0x00791194 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007911D8 | `HandlePlayPause` | Known | Event handler |
| 0x0079120E | `HandleShowContextualMenu` | Known | Event handler |
| 0x00791250 | `HandleAddToOTG` | Known | Event handler |
| 0x007914D5 | `HandlePlayPause` | Known | Event handler |
| 0x007914FC | `HandleSelect` | Known | Event handler |
| 0x0079152C | `HandlePlayPause` | Known | Event handler |
| 0x00791562 | `HandleShowContextualMenu` | Known | Event handler |
| 0x007915A9 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007915EC | `HandleAddToOTG` | Known | Event handler |
| 0x0079164F | `HandleStartGenius` | Known | Event handler |
| 0x0079168B | `HandleViewAlbum` | Known | Event handler |
| 0x007916C6 | `HandleViewArtist` | Known | Event handler |
| 0x00791707 | `HandleViewCompilation` | Known | Event handler |
| 0x00791B57 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x00791B9C | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00791BDF | `HandleAddToOTG` | Known | Event handler |
| 0x00791C42 | `HandleStartGenius` | Known | Event handler |
| 0x00791C7E | `HandleViewAlbum` | Known | Event handler |
| 0x00791CB9 | `HandleViewArtist` | Known | Event handler |
| 0x00791CFA | `HandleViewCompilation` | Known | Event handler |
| 0x00792038 | `HandlePlayPause` | Known | Event handler |
| 0x00792165 | `HandleSelect` | Known | Event handler |
| 0x00792191 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007921D4 | `HandlePlayPause` | Known | Event handler |
| 0x0079225A | `HandleSelect` | Known | Event handler |
| 0x00792287 | `HandleHilite` | Known | Event handler |
| 0x007922B3 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007922F4 | `HandlePlayPause` | Known | Event handler |
| 0x00792427 | `HandleSelect` | Known | Event handler |
| 0x00792453 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00792D65 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079361D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00793ED5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0079478D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00795045 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007958FD | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x007961B5 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00796A6D | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00796AB6 | `HandleTVOutChanged` | Known | Event handler |
| 0x00796AEE | `HandleTVSignalChanged` | Known | Event handler |
| 0x00796B29 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x00796B7A | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x00796BBF | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x00796C08 | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x00796C4A | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x00796C94 | `HandlePlayPause` | Known | Event handler |
| 0x00796CCA | `HandleShowContextualMenu` | Known | Event handler |
| 0x00796D11 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00796D54 | `HandleAddToOTG` | Known | Event handler |
| 0x00796DB7 | `HandleStartGenius` | Known | Event handler |
| 0x00796DF3 | `HandleViewAlbum` | Known | Event handler |
| 0x00796E2E | `HandleViewArtist` | Known | Event handler |
| 0x00796E6F | `HandleViewCompilation` | Known | Event handler |
| 0x007970AB | `HandlePlayPause` | Known | Event handler |
| 0x007970D2 | `HandleSelect` | Known | Event handler |
| 0x00797104 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x0079713F | `HandleSaveGeniusPlaylist` | Known | Event handler |
| 0x007971E0 | `HandlePlayPause` | Known | Event handler |
| 0x00797216 | `HandleShowContextualMenu` | Known | Event handler |
| 0x0079725D | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007972A0 | `HandleAddToOTG` | Known | Event handler |
| 0x00797303 | `HandleStartGenius` | Known | Event handler |
| 0x0079733F | `HandleViewAlbum` | Known | Event handler |
| 0x0079737A | `HandleViewArtist` | Known | Event handler |
| 0x007973BB | `HandleViewCompilation` | Known | Event handler |
| 0x00797429 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x00797851 | `HandlePlayPause` | Known | Event handler |
| 0x00797878 | `HandleSelect` | Known | Event handler |
| 0x007978AA | `HandleRefreshPlaylist` | Known | Event handler |
| 0x007978E1 | `HandleSelect` | Known | Event handler |
| 0x00797911 | `HandleSelect` | Known | Event handler |
| 0x00797949 | `HandleMenuLongpress` | Known | Event handler |
| 0x00797977 | `HandleMenuKey` | Known | Event handler |
| 0x007979FD | `HandlePlayPause` | Known | Event handler |
| 0x00797A87 | `HandlePushContextualMenu` | Known | Event handler |
| 0x00797ABC | `HandleSelect` | Known | Event handler |
| 0x00797AF7 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00797B3A | `HandleAddToOTG` | Known | Event handler |
| 0x00797B79 | `HandleAudiobookFaster` | Known | Event handler |
| 0x00797BBF | `HandleAudiobookNormal` | Known | Event handler |
| 0x00797C05 | `HandleAudiobookSlower` | Known | Event handler |
| 0x00797C6F | `HandleStartGenius` | Known | Event handler |
| 0x00797CAB | `HandleViewAlbum` | Known | Event handler |
| 0x00797CE6 | `HandleViewArtist` | Known | Event handler |
| 0x00797D27 | `HandleViewCompilation` | Known | Event handler |
| 0x00798769 | `HandleStartGenius` | Known | Event handler |
| 0x0079887C | `HandlePlayPause` | Known | Event handler |
| 0x007988F1 | `HandleWheelProgress` | Known | Event handler |
| 0x0079892D | `HandleMenuLongpress` | Known | Event handler |
| 0x0079895B | `HandleMenuKey` | Known | Event handler |
| 0x007989E1 | `HandlePlayPause` | Known | Event handler |
| 0x00798A6B | `HandlePushContextualMenu` | Known | Event handler |
| 0x00798AA0 | `HandleSelectProgress` | Known | Event handler |
| 0x00798AE3 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00798B26 | `HandleAddToOTG` | Known | Event handler |
| 0x00798B65 | `HandleAudiobookFaster` | Known | Event handler |
| 0x00798BAB | `HandleAudiobookNormal` | Known | Event handler |
| 0x00798BF1 | `HandleAudiobookSlower` | Known | Event handler |
| 0x00798C5B | `HandleStartGenius` | Known | Event handler |
| 0x00798C97 | `HandleViewAlbum` | Known | Event handler |
| 0x00798CD2 | `HandleViewArtist` | Known | Event handler |
| 0x00798D13 | `HandleViewCompilation` | Known | Event handler |
| 0x00799755 | `HandleStartGenius` | Known | Event handler |
| 0x00799868 | `HandlePlayPause` | Known | Event handler |
| 0x007998DD | `HandleWheelProgress` | Known | Event handler |
| 0x00799919 | `HandleMenuLongpress` | Known | Event handler |
| 0x00799947 | `HandleMenuKey` | Known | Event handler |
| 0x007999CD | `HandlePlayPause` | Known | Event handler |
| 0x00799A57 | `HandlePushContextualMenu` | Known | Event handler |
| 0x00799A8C | `HandleSelectVolume` | Known | Event handler |
| 0x00799ACD | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00799B10 | `HandleAddToOTG` | Known | Event handler |
| 0x00799B4F | `HandleAudiobookFaster` | Known | Event handler |
| 0x00799B95 | `HandleAudiobookNormal` | Known | Event handler |
| 0x00799BDB | `HandleAudiobookSlower` | Known | Event handler |
| 0x00799C45 | `HandleStartGenius` | Known | Event handler |
| 0x00799C81 | `HandleViewAlbum` | Known | Event handler |
| 0x00799CBC | `HandleViewArtist` | Known | Event handler |
| 0x00799CFD | `HandleViewCompilation` | Known | Event handler |
| 0x0079A73F | `HandleStartGenius` | Known | Event handler |
| 0x0079A852 | `HandlePlayPause` | Known | Event handler |
| 0x0079A8C7 | `HandleWheelVolume` | Known | Event handler |
| 0x0079A901 | `HandleMenuLongpress` | Known | Event handler |
| 0x0079A92F | `HandleMenuKey` | Known | Event handler |
| 0x0079A9B5 | `HandlePlayPause` | Known | Event handler |
| 0x0079AA3F | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079AA74 | `HandleSelectRating` | Known | Event handler |
| 0x0079AAB5 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079AAF8 | `HandleAddToOTG` | Known | Event handler |
| 0x0079AB37 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079AB7D | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079ABC3 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079AC2D | `HandleStartGenius` | Known | Event handler |
| 0x0079AC69 | `HandleViewAlbum` | Known | Event handler |
| 0x0079ACA4 | `HandleViewArtist` | Known | Event handler |
| 0x0079ACE5 | `HandleViewCompilation` | Known | Event handler |
| 0x0079B727 | `HandleStartGenius` | Known | Event handler |
| 0x0079B83A | `HandlePlayPause` | Known | Event handler |
| 0x0079B8AF | `HandleWheelRating` | Known | Event handler |
| 0x0079B8E9 | `HandleMenuLongpress` | Known | Event handler |
| 0x0079B917 | `HandleMenuKey` | Known | Event handler |
| 0x0079B98F | `HandlePlayPause` | Known | Event handler |
| 0x0079BA10 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079BA45 | `HandleSelectScrub` | Known | Event handler |
| 0x0079BA85 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079BAC8 | `HandleAddToOTG` | Known | Event handler |
| 0x0079BB07 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079BB4D | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079BB93 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079BBFD | `HandleStartGenius` | Known | Event handler |
| 0x0079BC39 | `HandleViewAlbum` | Known | Event handler |
| 0x0079BC74 | `HandleViewArtist` | Known | Event handler |
| 0x0079BCB5 | `HandleViewCompilation` | Known | Event handler |
| 0x0079C6F7 | `HandleStartGenius` | Known | Event handler |
| 0x0079C7FC | `HandlePlayPause` | Known | Event handler |
| 0x0079C868 | `HandleWheelScrub` | Known | Event handler |
| 0x0079C8A1 | `HandleMenuLongpress` | Known | Event handler |
| 0x0079C8CF | `HandleMenuKey` | Known | Event handler |
| 0x0079C955 | `HandlePlayPause` | Known | Event handler |
| 0x0079C9DF | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079CA14 | `HandleSelectGenius` | Known | Event handler |
| 0x0079CA55 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079CA98 | `HandleAddToOTG` | Known | Event handler |
| 0x0079CAD7 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079CB1D | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079CB63 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079CBCD | `HandleStartGenius` | Known | Event handler |
| 0x0079CC09 | `HandleViewAlbum` | Known | Event handler |
| 0x0079CC44 | `HandleViewArtist` | Known | Event handler |
| 0x0079CC85 | `HandleViewCompilation` | Known | Event handler |
| 0x0079D6C7 | `HandleStartGenius` | Known | Event handler |
| 0x0079D7DA | `HandlePlayPause` | Known | Event handler |
| 0x0079D84F | `HandleWheelGenius` | Known | Event handler |
| 0x0079D889 | `HandleMenuLongpress` | Known | Event handler |
| 0x0079D8B7 | `HandleMenuKey` | Known | Event handler |
| 0x0079D914 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0079D94C | `HandlePlayPause` | Known | Event handler |
| 0x0079D9A6 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0079D9E5 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079DA1A | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x0079DA62 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079DAA5 | `HandleAddToOTG` | Known | Event handler |
| 0x0079DAE4 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079DB2A | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079DB70 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079DBDA | `HandleStartGenius` | Known | Event handler |
| 0x0079DC16 | `HandleViewAlbum` | Known | Event handler |
| 0x0079DC51 | `HandleViewArtist` | Known | Event handler |
| 0x0079DC92 | `HandleViewCompilation` | Known | Event handler |
| 0x0079E6D4 | `HandleStartGenius` | Known | Event handler |
| 0x0079E7E7 | `HandlePlayPause` | Known | Event handler |
| 0x0079E85C | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0079E89D | `HandleMenuLongpress` | Known | Event handler |
| 0x0079E8CB | `HandleMenuKey` | Known | Event handler |
| 0x0079E951 | `HandlePlayPause` | Known | Event handler |
| 0x0079E9DB | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079EA10 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x0079EA54 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079EA97 | `HandleAddToOTG` | Known | Event handler |
| 0x0079EAD6 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079EB1C | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079EB62 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079EBCC | `HandleStartGenius` | Known | Event handler |
| 0x0079EC08 | `HandleViewAlbum` | Known | Event handler |
| 0x0079EC43 | `HandleViewArtist` | Known | Event handler |
| 0x0079EC84 | `HandleViewCompilation` | Known | Event handler |
| 0x0079F6C6 | `HandleStartGenius` | Known | Event handler |
| 0x0079F7D9 | `HandlePlayPause` | Known | Event handler |
| 0x0079F879 | `HandleMenuLongpress` | Known | Event handler |
| 0x0079F8A7 | `HandleMenuKey` | Known | Event handler |
| 0x0079F92D | `HandlePlayPause` | Known | Event handler |
| 0x0079F9B7 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0079F9EC | `HandleSelectExtraInfo` | Known | Event handler |
| 0x0079FA30 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x0079FA73 | `HandleAddToOTG` | Known | Event handler |
| 0x0079FAB2 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0079FAF8 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0079FB3E | `HandleAudiobookSlower` | Known | Event handler |
| 0x0079FBA8 | `HandleStartGenius` | Known | Event handler |
| 0x0079FBE4 | `HandleViewAlbum` | Known | Event handler |
| 0x0079FC1F | `HandleViewArtist` | Known | Event handler |
| 0x0079FC60 | `HandleViewCompilation` | Known | Event handler |
| 0x007A06A2 | `HandleStartGenius` | Known | Event handler |
| 0x007A07B5 | `HandlePlayPause` | Known | Event handler |
| 0x007A0855 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A0883 | `HandleMenuKey` | Known | Event handler |
| 0x007A0909 | `HandlePlayPause` | Known | Event handler |
| 0x007A0993 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A09C8 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x007A0A0C | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A0A4F | `HandleAddToOTG` | Known | Event handler |
| 0x007A0A8E | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A0AD4 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A0B1A | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A0B84 | `HandleStartGenius` | Known | Event handler |
| 0x007A0BC0 | `HandleViewAlbum` | Known | Event handler |
| 0x007A0BFB | `HandleViewArtist` | Known | Event handler |
| 0x007A0C3C | `HandleViewCompilation` | Known | Event handler |
| 0x007A167E | `HandleStartGenius` | Known | Event handler |
| 0x007A1791 | `HandlePlayPause` | Known | Event handler |
| 0x007A1831 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A185F | `HandleMenuKey` | Known | Event handler |
| 0x007A18E5 | `HandlePlayPause` | Known | Event handler |
| 0x007A196F | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A19A4 | `HandleSelectChapterArt` | Known | Event handler |
| 0x007A19E9 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A1A2C | `HandleAddToOTG` | Known | Event handler |
| 0x007A1A6B | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A1AB1 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A1AF7 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A1B61 | `HandleStartGenius` | Known | Event handler |
| 0x007A1B9D | `HandleViewAlbum` | Known | Event handler |
| 0x007A1BD8 | `HandleViewArtist` | Known | Event handler |
| 0x007A1C19 | `HandleViewCompilation` | Known | Event handler |
| 0x007A265B | `HandleStartGenius` | Known | Event handler |
| 0x007A276E | `HandlePlayPause` | Known | Event handler |
| 0x007A27E3 | `HandleWheelVolume` | Known | Event handler |
| 0x007A281D | `HandleMenuLongpress` | Known | Event handler |
| 0x007A284B | `HandleMenuKey` | Known | Event handler |
| 0x007A28DA | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007A297B | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A29B0 | `HandleSelect` | Known | Event handler |
| 0x007A29EB | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A2A2E | `HandleAddToOTG` | Known | Event handler |
| 0x007A2A6D | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A2AB3 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A2AF9 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A2B63 | `HandleStartGenius` | Known | Event handler |
| 0x007A2B9F | `HandleViewAlbum` | Known | Event handler |
| 0x007A2BDA | `HandleViewArtist` | Known | Event handler |
| 0x007A2C1B | `HandleViewCompilation` | Known | Event handler |
| 0x007A365D | `HandleStartGenius` | Known | Event handler |
| 0x007A3779 | `HandlePlayPause` | Known | Event handler |
| 0x007A37F7 | `HandleWheel` | Known | Event handler |
| 0x007A382D | `HandleMenuLongpress` | Known | Event handler |
| 0x007A385B | `HandleMenuKey` | Known | Event handler |
| 0x007A38EA | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x007A398B | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A39C0 | `HandleSelect` | Known | Event handler |
| 0x007A39FB | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A3A3E | `HandleAddToOTG` | Known | Event handler |
| 0x007A3A7D | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A3AC3 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A3B09 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A3B73 | `HandleStartGenius` | Known | Event handler |
| 0x007A3BAF | `HandleViewAlbum` | Known | Event handler |
| 0x007A3BEA | `HandleViewArtist` | Known | Event handler |
| 0x007A3C2B | `HandleViewCompilation` | Known | Event handler |
| 0x007A466D | `HandleStartGenius` | Known | Event handler |
| 0x007A4789 | `HandlePlayPause` | Known | Event handler |
| 0x007A4807 | `HandleWheel` | Known | Event handler |
| 0x007A483D | `HandleMenuLongpress` | Known | Event handler |
| 0x007A486B | `HandleMenuKey` | Known | Event handler |
| 0x007A48F1 | `HandlePlayPause` | Known | Event handler |
| 0x007A497B | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A49B0 | `HandleSelect` | Known | Event handler |
| 0x007A49EB | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A4A2E | `HandleAddToOTG` | Known | Event handler |
| 0x007A4A6D | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A4AB3 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A4AF9 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A4B63 | `HandleStartGenius` | Known | Event handler |
| 0x007A4B9F | `HandleViewAlbum` | Known | Event handler |
| 0x007A4BDA | `HandleViewArtist` | Known | Event handler |
| 0x007A4C1B | `HandleViewCompilation` | Known | Event handler |
| 0x007A565D | `HandleStartGenius` | Known | Event handler |
| 0x007A5770 | `HandlePlayPause` | Known | Event handler |
| 0x007A57E5 | `HandleWheel` | Known | Event handler |
| 0x007A5819 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A5847 | `HandleMenuKey` | Known | Event handler |
| 0x007A58CD | `HandlePlayPause` | Known | Event handler |
| 0x007A5957 | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A598C | `HandleSelectProgress` | Known | Event handler |
| 0x007A59CF | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A5A12 | `HandleAddToOTG` | Known | Event handler |
| 0x007A5A51 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A5A97 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A5ADD | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A5B47 | `HandleStartGenius` | Known | Event handler |
| 0x007A5B83 | `HandleViewAlbum` | Known | Event handler |
| 0x007A5BBE | `HandleViewArtist` | Known | Event handler |
| 0x007A5BFF | `HandleViewCompilation` | Known | Event handler |
| 0x007A6641 | `HandleStartGenius` | Known | Event handler |
| 0x007A6754 | `HandlePlayPause` | Known | Event handler |
| 0x007A67C9 | `HandleWheelProgress` | Known | Event handler |
| 0x007A6805 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A6833 | `HandleMenuKey` | Known | Event handler |
| 0x007A68AB | `HandlePlayPause` | Known | Event handler |
| 0x007A692C | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A6961 | `HandleSelectScrub` | Known | Event handler |
| 0x007A69A1 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A69E4 | `HandleAddToOTG` | Known | Event handler |
| 0x007A6A23 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A6A69 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A6AAF | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A6B19 | `HandleStartGenius` | Known | Event handler |
| 0x007A6B55 | `HandleViewAlbum` | Known | Event handler |
| 0x007A6B90 | `HandleViewArtist` | Known | Event handler |
| 0x007A6BD1 | `HandleViewCompilation` | Known | Event handler |
| 0x007A7613 | `HandleStartGenius` | Known | Event handler |
| 0x007A7718 | `HandlePlayPause` | Known | Event handler |
| 0x007A7784 | `HandleWheelScrub` | Known | Event handler |
| 0x007A77BD | `HandleMenuLongpress` | Known | Event handler |
| 0x007A77EB | `HandleMenuKey` | Known | Event handler |
| 0x007A7871 | `HandlePlayPause` | Known | Event handler |
| 0x007A78FB | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A796A | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A79AD | `HandleAddToOTG` | Known | Event handler |
| 0x007A79EC | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A7A32 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A7A78 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A7AE2 | `HandleStartGenius` | Known | Event handler |
| 0x007A7B1E | `HandleViewAlbum` | Known | Event handler |
| 0x007A7B59 | `HandleViewArtist` | Known | Event handler |
| 0x007A7B9A | `HandleViewCompilation` | Known | Event handler |
| 0x007A85DC | `HandleStartGenius` | Known | Event handler |
| 0x007A86EF | `HandlePlayPause` | Known | Event handler |
| 0x007A8764 | `HandleWheelVolume` | Known | Event handler |
| 0x007A87A1 | `HandleMenuLongpress` | Known | Event handler |
| 0x007A87CF | `HandleMenuKey` | Known | Event handler |
| 0x007A8855 | `HandlePlayPause` | Known | Event handler |
| 0x007A88DF | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A894E | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A8991 | `HandleAddToOTG` | Known | Event handler |
| 0x007A89D0 | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A8A16 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A8A5C | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A8AC6 | `HandleStartGenius` | Known | Event handler |
| 0x007A8B02 | `HandleViewAlbum` | Known | Event handler |
| 0x007A8B3D | `HandleViewArtist` | Known | Event handler |
| 0x007A8B7E | `HandleViewCompilation` | Known | Event handler |
| 0x007A95C0 | `HandleStartGenius` | Known | Event handler |
| 0x007A96D3 | `HandlePlayPause` | Known | Event handler |
| 0x007A9748 | `HandleWheelBrightness` | Known | Event handler |
| 0x007A986B | `HandlePushContextualMenu` | Known | Event handler |
| 0x007A98A0 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007A98E8 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x007A992B | `HandleAddToOTG` | Known | Event handler |
| 0x007A996A | `HandleAudiobookFaster` | Known | Event handler |
| 0x007A99B0 | `HandleAudiobookNormal` | Known | Event handler |
| 0x007A99F6 | `HandleAudiobookSlower` | Known | Event handler |
| 0x007A9A60 | `HandleStartGenius` | Known | Event handler |
| 0x007A9A9C | `HandleViewAlbum` | Known | Event handler |
| 0x007A9AD7 | `HandleViewArtist` | Known | Event handler |
| 0x007A9B18 | `HandleViewCompilation` | Known | Event handler |
| 0x007AA55A | `HandleStartGenius` | Known | Event handler |
| 0x007AA6A6 | `HandleWheel` | Known | Event handler |
| 0x007AA6DD | `HandleMenuLongpress` | Known | Event handler |
| 0x007AA70B | `HandleMenuKey` | Known | Event handler |
| 0x007AA791 | `HandlePlayPause` | Known | Event handler |
| 0x007AA811 | `HandleSelect` | Known | Event handler |
| 0x007AACB3 | `HandlePlayPause` | Known | Event handler |
| 0x007AAD41 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AAD6F | `HandleMenuKey` | Known | Event handler |
| 0x007AADF5 | `HandlePlayPause` | Known | Event handler |
| 0x007AAE75 | `HandleSelectProgress` | Known | Event handler |
| 0x007AB31F | `HandlePlayPause` | Known | Event handler |
| 0x007AB394 | `HandleWheelProgress` | Known | Event handler |
| 0x007AB3D1 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AB3FF | `HandleMenuKey` | Known | Event handler |
| 0x007AB485 | `HandlePlayPause` | Known | Event handler |
| 0x007AB505 | `HandleSelectProgress` | Known | Event handler |
| 0x007AB9AF | `HandlePlayPause` | Known | Event handler |
| 0x007ABA24 | `HandleWheelProgress` | Known | Event handler |
| 0x007ABA61 | `HandleMenuLongpress` | Known | Event handler |
| 0x007ABA8F | `HandleMenuKey` | Known | Event handler |
| 0x007ABB15 | `HandlePlayPause` | Known | Event handler |
| 0x007ABB95 | `HandleSelectProgress` | Known | Event handler |
| 0x007ABFCB | `HandlePlayPause` | Known | Event handler |
| 0x007AC040 | `HandleWheelProgress` | Known | Event handler |
| 0x007AC07D | `HandleMenuLongpress` | Known | Event handler |
| 0x007AC0AB | `HandleMenuKey` | Known | Event handler |
| 0x007AC118 | `HandlePlayPause` | Known | Event handler |
| 0x007AC184 | `HandleSelectScrub` | Known | Event handler |
| 0x007AC59E | `HandlePlayPause` | Known | Event handler |
| 0x007AC5FF | `HandleWheelScrub` | Known | Event handler |
| 0x007AC639 | `HandleMenuLongpress` | Known | Event handler |
| 0x007AC667 | `HandleMenuKey` | Known | Event handler |
| 0x007AC6ED | `HandlePlayPause` | Known | Event handler |
| 0x007AC76D | `HandleSelectVolume` | Known | Event handler |
| 0x007ACBA1 | `HandlePlayPause` | Known | Event handler |
| 0x007ACC16 | `HandleWheelVolume` | Known | Event handler |
| 0x007ACD29 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x007AD1C8 | `HandleSelect` | Known | Event handler |
| 0x007AD1F5 | `HandleSelect` | Known | Event handler |
| 0x007AD225 | `HandleSelect` | Known | Event handler |
| 0x007AD255 | `HandleSelect` | Known | Event handler |
| 0x007AD285 | `HandleSelect` | Known | Event handler |
| 0x007AD2B5 | `HandleSelect` | Known | Event handler |
| 0x007AD2E5 | `HandleSelect` | Known | Event handler |
| 0x007AD315 | `HandleSelect` | Known | Event handler |
| 0x007AD345 | `HandleSelect` | Known | Event handler |
| 0x007AD3B5 | `HandleSelect` | Known | Event handler |
| 0x007AD3E5 | `HandleSelect` | Known | Event handler |
| 0x007AD45D | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AD490 | `HandleNotesPop` | Known | Event handler |
| 0x007AD50D | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AD540 | `HandleNotesPop` | Known | Event handler |
| 0x007AD9FC | `HandleNotesSelected` | Known | Event handler |
| 0x007ADA39 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007ADA6C | `HandleNotesPop` | Known | Event handler |
| 0x007ADF28 | `HandleNotesSelected` | Known | Event handler |
| 0x007ADF65 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007ADF98 | `HandleNotesPop` | Known | Event handler |
| 0x007ADFC3 | `HandleNotesSelected` | Known | Event handler |
| 0x007AE495 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AE4C8 | `HandleNotesPop` | Known | Event handler |
| 0x007AE4F3 | `HandleNotesSelected` | Known | Event handler |
| 0x007AE9C5 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AE9F8 | `HandleNotesPop` | Known | Event handler |
| 0x007AEA75 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AEAA8 | `HandleNotesPop` | Known | Event handler |
| 0x007AEB25 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x007AEB58 | `HandleNotesPop` | Known | Event handler |
| 0x007AEBD0 | `HandlePlayPause` | Known | Event handler |
| 0x007AEBF9 | `HandlePlayPause` | Known | Event handler |
| 0x007AEC27 | `HandlePlayPause` | Known | Event handler |
| 0x007AEC5C | `HandleBrowseAlbum` | Known | Event handler |
| 0x007AECDC | `HandleHiliteAlbum` | Known | Event handler |
| 0x007AED85 | `HandleBrowseAlbum` | Known | Event handler |
| 0x007AEE0C | `HandleHiliteAlbum` | Known | Event handler |
| 0x007AF0D0 | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x007AF12C | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x007AF2E3 | `HandleSelect` | Known | Event handler |
| 0x007AF467 | `HandleSelect` | Known | Event handler |
| 0x007AF4A1 | `HandleImageLast` | Known | Event handler |
| 0x007AF4CB | `HandleImageNext` | Known | Event handler |
| 0x007AF4FA | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AF534 | `HandleImageFirst` | Known | Event handler |
| 0x007AF55F | `HandleImagePrev` | Known | Event handler |
| 0x007AF58B | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AF5BA | `HandleImageNext` | Known | Event handler |
| 0x007AF5E3 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AF617 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007AF646 | `HandleImagePrev` | Known | Event handler |
| 0x007AF667 | `HandleImageWheel` | Known | Event handler |
| 0x007AF705 | `HandleImageNext` | Known | Event handler |
| 0x007AF734 | `HandlePlayPause` | Known | Event handler |
| 0x007AF783 | `HandleImagePrev` | Known | Event handler |
| 0x007AF7AF | `HandleSelect` | Known | Event handler |
| 0x007AFA7F | `HandleImageNext` | Known | Event handler |
| 0x007AFAA9 | `HandlePause` | Known | Event handler |
| 0x007AFACE | `HandlePlay` | Known | Event handler |
| 0x007AFAF7 | `HandlePlayPause` | Known | Event handler |
| 0x007AFB20 | `HandleImagePrev` | Known | Event handler |
| 0x007AFB83 | `HandleMikeyCenter` | Known | Event handler |
| 0x007AFBA6 | `HandleWheel` | Known | Event handler |
| 0x007AFC41 | `HandleImageNext` | Known | Event handler |
| 0x007AFC70 | `HandlePlayPause` | Known | Event handler |
| 0x007AFCBF | `HandleImagePrev` | Known | Event handler |
| 0x007AFCEB | `HandleSelect` | Known | Event handler |
| 0x007AFFBB | `HandleImageNext` | Known | Event handler |
| 0x007AFFE5 | `HandlePause` | Known | Event handler |
| 0x007B000A | `HandlePlay` | Known | Event handler |
| 0x007B0033 | `HandlePlayPause` | Known | Event handler |
| 0x007B005C | `HandleImagePrev` | Known | Event handler |
| 0x007B00BF | `HandleMikeyCenter` | Known | Event handler |
| 0x007B00E2 | `HandleWheel` | Known | Event handler |
| 0x007B017D | `HandleImageNext` | Known | Event handler |
| 0x007B01AC | `HandlePlayPause` | Known | Event handler |
| 0x007B01FB | `HandleImagePrev` | Known | Event handler |
| 0x007B0227 | `HandleSelect` | Known | Event handler |
| 0x007B04F7 | `HandleImageNext` | Known | Event handler |
| 0x007B0521 | `HandlePause` | Known | Event handler |
| 0x007B0546 | `HandlePlay` | Known | Event handler |
| 0x007B056F | `HandlePlayPause` | Known | Event handler |
| 0x007B0598 | `HandleImagePrev` | Known | Event handler |
| 0x007B05FB | `HandleMikeyCenter` | Known | Event handler |
| 0x007B061E | `HandleWheel` | Known | Event handler |
| 0x007B06B9 | `HandleImageNext` | Known | Event handler |
| 0x007B06E8 | `HandlePlayPause` | Known | Event handler |
| 0x007B0737 | `HandleImagePrev` | Known | Event handler |
| 0x007B0763 | `HandleSelect` | Known | Event handler |
| 0x007B0A33 | `HandleImageNext` | Known | Event handler |
| 0x007B0A5D | `HandlePause` | Known | Event handler |
| 0x007B0A82 | `HandlePlay` | Known | Event handler |
| 0x007B0AAB | `HandlePlayPause` | Known | Event handler |
| 0x007B0AD4 | `HandleImagePrev` | Known | Event handler |
| 0x007B0B37 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B0B5A | `HandleWheel` | Known | Event handler |
| 0x007B0BF5 | `HandleImageNext` | Known | Event handler |
| 0x007B0C24 | `HandlePlayPause` | Known | Event handler |
| 0x007B0C73 | `HandleImagePrev` | Known | Event handler |
| 0x007B0C9F | `HandleSelect` | Known | Event handler |
| 0x007B0F6F | `HandleImageNext` | Known | Event handler |
| 0x007B0F99 | `HandlePause` | Known | Event handler |
| 0x007B0FBE | `HandlePlay` | Known | Event handler |
| 0x007B0FE7 | `HandlePlayPause` | Known | Event handler |
| 0x007B1010 | `HandleImagePrev` | Known | Event handler |
| 0x007B1073 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B1096 | `HandleWheel` | Known | Event handler |
| 0x007B1131 | `HandleImageNext` | Known | Event handler |
| 0x007B1160 | `HandlePlayPause` | Known | Event handler |
| 0x007B11AF | `HandleImagePrev` | Known | Event handler |
| 0x007B11DB | `HandleSelect` | Known | Event handler |
| 0x007B14AB | `HandleImageNext` | Known | Event handler |
| 0x007B14D5 | `HandlePause` | Known | Event handler |
| 0x007B14FA | `HandlePlay` | Known | Event handler |
| 0x007B1523 | `HandlePlayPause` | Known | Event handler |
| 0x007B154C | `HandleImagePrev` | Known | Event handler |
| 0x007B15AF | `HandleMikeyCenter` | Known | Event handler |
| 0x007B15D2 | `HandleWheel` | Known | Event handler |
| 0x007B166D | `HandleImageNext` | Known | Event handler |
| 0x007B169C | `HandlePlayPause` | Known | Event handler |
| 0x007B16EB | `HandleImagePrev` | Known | Event handler |
| 0x007B1717 | `HandleSelect` | Known | Event handler |
| 0x007B1962 | `HandleImageNext` | Known | Event handler |
| 0x007B198C | `HandlePause` | Known | Event handler |
| 0x007B19B1 | `HandlePlay` | Known | Event handler |
| 0x007B19DA | `HandlePlayPause` | Known | Event handler |
| 0x007B1A03 | `HandleImagePrev` | Known | Event handler |
| 0x007B1A76 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B1A99 | `HandleWheel` | Known | Event handler |
| 0x007B1B31 | `HandleImageNext` | Known | Event handler |
| 0x007B1B60 | `HandlePlayPause` | Known | Event handler |
| 0x007B1BAF | `HandleImagePrev` | Known | Event handler |
| 0x007B1BDB | `HandleSelect` | Known | Event handler |
| 0x007B1E26 | `HandleImageNext` | Known | Event handler |
| 0x007B1E50 | `HandlePause` | Known | Event handler |
| 0x007B1E75 | `HandlePlay` | Known | Event handler |
| 0x007B1E9E | `HandlePlayPause` | Known | Event handler |
| 0x007B1EC7 | `HandleImagePrev` | Known | Event handler |
| 0x007B1F3A | `HandleMikeyCenter` | Known | Event handler |
| 0x007B1F5D | `HandleWheel` | Known | Event handler |
| 0x007B1FF5 | `HandleImageNext` | Known | Event handler |
| 0x007B2024 | `HandlePlayPause` | Known | Event handler |
| 0x007B2073 | `HandleImagePrev` | Known | Event handler |
| 0x007B209F | `HandleSelect` | Known | Event handler |
| 0x007B22EA | `HandleImageNext` | Known | Event handler |
| 0x007B2314 | `HandlePause` | Known | Event handler |
| 0x007B2339 | `HandlePlay` | Known | Event handler |
| 0x007B2362 | `HandlePlayPause` | Known | Event handler |
| 0x007B238B | `HandleImagePrev` | Known | Event handler |
| 0x007B23FE | `HandleMikeyCenter` | Known | Event handler |
| 0x007B2421 | `HandleWheel` | Known | Event handler |
| 0x007B24B9 | `HandleImageNext` | Known | Event handler |
| 0x007B24E8 | `HandlePlayPause` | Known | Event handler |
| 0x007B2537 | `HandleImagePrev` | Known | Event handler |
| 0x007B2563 | `HandleSelect` | Known | Event handler |
| 0x007B27AE | `HandleImageNext` | Known | Event handler |
| 0x007B27D8 | `HandlePause` | Known | Event handler |
| 0x007B27FD | `HandlePlay` | Known | Event handler |
| 0x007B2826 | `HandlePlayPause` | Known | Event handler |
| 0x007B284F | `HandleImagePrev` | Known | Event handler |
| 0x007B28C2 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B28E5 | `HandleWheel` | Known | Event handler |
| 0x007B297D | `HandleImageNext` | Known | Event handler |
| 0x007B29AC | `HandlePlayPause` | Known | Event handler |
| 0x007B29FB | `HandleImagePrev` | Known | Event handler |
| 0x007B2A27 | `HandleSelect` | Known | Event handler |
| 0x007B2C72 | `HandleImageNext` | Known | Event handler |
| 0x007B2C9C | `HandlePause` | Known | Event handler |
| 0x007B2CC1 | `HandlePlay` | Known | Event handler |
| 0x007B2CEA | `HandlePlayPause` | Known | Event handler |
| 0x007B2D13 | `HandleImagePrev` | Known | Event handler |
| 0x007B2D86 | `HandleMikeyCenter` | Known | Event handler |
| 0x007B2DA9 | `HandleWheel` | Known | Event handler |
| 0x007B2DD5 | `HandleSelect` | Known | Event handler |
| 0x007B2E05 | `HandleSelect` | Known | Event handler |
| 0x007B2F28 | `HandleTuning` | Known | Event handler |
| 0x007B30E4 | `HandleVolumeChange` | Known | Event handler |
| 0x007B3230 | `HandleVolumeWheel` | Known | Event handler |
| 0x007B338B | `HandleTuningSelect` | Known | Event handler |
| 0x007B366A | `HandleFrequencyChange` | Known | Event handler |
| 0x007B37C7 | `HandleTuningSelect` | Known | Event handler |
| 0x007B3AA6 | `HandleFrequencyChange` | Known | Event handler |
| 0x007B3BD0 | `HandleTimerDone` | Known | Event handler |
| 0x007B3DC5 | `HandleVolumeChange` | Known | Event handler |
| 0x007B3EDC | `HandleVolumeWheel` | Known | Event handler |
| 0x007B472F | `HandleExitUnsupported` | Known | Event handler |
| 0x007B4761 | `HandleExitUnsupported` | Known | Event handler |
| 0x007B9795 | `HandleSelectKey` | Known | Event handler |
| 0x007B97CA | `HandleWheel` | Known | Event handler |
| 0x007B9918 | `HandleLanguageAfterReset` | Known | Event handler |
| 0x007B996B | `HandleSelectKey` | Known | Event handler |
| 0x007B9993 | `HandleSelectKey` | Known | Event handler |
| 0x007B99C3 | `HandleExit` | Known | Event handler |
| 0x007B99ED | `HandleStartStop` | Known | Event handler |
| 0x007B9A53 | `HandleStartStop` | Known | Event handler |
| 0x007B9B6B | `HandleExit` | Known | Event handler |
| 0x007B9B95 | `HandleStartStop` | Known | Event handler |
| 0x007B9BC1 | `HandleLap` | Known | Event handler |
| 0x007B9CC5 | `HandleSelectLozinch` | Known | Event handler |
| 0x007B9EE2 | `HandleSelect` | Known | Event handler |
| 0x007B9F6E | `HandleSelect` | Known | Event handler |
| 0x007B9FFC | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x007BA2FA | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x007BA3E5 | `HandleFinishRecording` | Known | Event handler |
| 0x007BA436 | `HandlePlayPause` | Known | Event handler |
| 0x007BA4C4 | `HandlePlayPause` | Known | Event handler |
| 0x007BA555 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x007BA58D | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x007BA5C9 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x007BA60C | `HandlePlayPause` | Known | Event handler |
| 0x007BA642 | `HandleAddToOTG` | Known | Event handler |
| 0x007BA897 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007BAAF3 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x007D79B6 | `HandleSelectClock` | Known | Event handler |
| 0x007D79EF | `HandleHilited` | Known | Event handler |
| 0x007D7A21 | `HandleWheel` | Known | Event handler |
| 0x007D7A68 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x007D7AED | `HandleBacksideSongSelected` | Known | Event handler |
| 0x007D7CF9 | `HandleImageLast` | Known | Event handler |
| 0x007D7D23 | `HandleScreenNext` | Known | Event handler |
| 0x007D7D53 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D7D8D | `HandleImageFirst` | Known | Event handler |
| 0x007D7DB8 | `HandleScreenPrev` | Known | Event handler |
| 0x007D7DE5 | `HandleBrowseLarge` | Known | Event handler |
| 0x007D7E65 | `HandleImageNext` | Known | Event handler |
| 0x007D7E8E | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D7EC2 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x007D7EF1 | `HandleImagePrev` | Known | Event handler |
| 0x007D7F1F | `HandleBrowseSmall` | Known | Event handler |

---

## 6. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000F5208 | `GotoNowPlaying` | Known | Navigation |
| 0x000F5280 | `GotoMainMenu` | Known | Navigation |
| 0x0010DDEC | `GotoScreen_LockDialog` | Known | Navigation |
| 0x0010DE04 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x0010DF7C | `GotoScreen_AddressBook` | Known | Navigation |
| 0x00119F04 | `GotoNowPlaying` | Known | Navigation |
| 0x00119F18 | `GotoAlbums` | Known | Navigation |
| 0x00119F24 | `GotoSongs` | Known | Navigation |
| 0x00127E54 | `GotoScreen_EnterPassKey` | Known | Navigation |
| 0x00127E6C | `GotoScreen_LockediPod` | Known | Navigation |
| 0x00128870 | `GotoScreen_MainMenu` | Known | Navigation |
| 0x0013EC5C | `GotoMainMenu` | Known | Navigation |
| 0x001C1BC4 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001C3F94 | `GotoErrorLayout` | Known | Navigation |
| 0x001CCF6C | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation |
| 0x001CD630 | `GotoPlayDeleteMenu` | Known | Navigation |
| 0x001CD6B4 | `GotoNowPlaying` | Known | Navigation |
| 0x001E85E4 | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation |
| 0x001F3FD4 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x001F40CC | `GotoScreen_SettingsMenu` | Known | Navigation |
| 0x001FBBF4 | `GotoDefaultLayout` | Known | Navigation |
| 0x001FBC78 | `GotoVolumeLayout` | Known | Navigation |
| 0x001FBDB0 | `GotoProgressLayout` | Known | Navigation |
| 0x001FC0CC | `GotoDefault` | Known | Navigation |
| 0x001FC400 | `GotoProgressLayout` | Known | Navigation |
| 0x001FC5C0 | `GotoRentalWarningLayout` | Known | Navigation |
| 0x001FC644 | `GotoProgressLayout` | Known | Navigation |
| 0x001FC954 | `GotoProgressLayout` | Known | Navigation |
| 0x001FE4E0 | `GotoNowPlaying` | Known | Navigation |
| 0x001FEDF0 | `GotoNowPlaying` | Known | Navigation |
| 0x001FF0F4 | `GotoNowPlaying` | Known | Navigation |
| 0x002017EC | `GotoScreen_Language` | Known | Navigation |
| 0x00201B4C | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x00201B68 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00201B80 | `GotoDefaultLayout` | Known | Navigation |
| 0x00201B94 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00201C2C | `GotoVolumeLayout` | Known | Navigation |
| 0x00201C40 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x00201CE0 | `GotoProgressLayout` | Known | Navigation |
| 0x00201CF4 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x002024A8 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00202910 | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x00202B7C | `GotoProgressLayout` | Known | Navigation |
| 0x00202B90 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00202D28 | `GotoBrightnessVideoLayout` | Known | Navigation |
| 0x00202D4C | `GotoGeniusLayout` | Known | Navigation |
| 0x00202D60 | `GotoRatingLayout` | Known | Navigation |
| 0x00202ED4 | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x00202EF0 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x00202F08 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x00203208 | `GotoChapterArtLayout` | Known | Navigation |
| 0x00203220 | `GotoShuffleLayout` | Known | Navigation |
| 0x002035B0 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x002035C4 | `GotoExtraInfoLoadingLayout` | Known | Navigation |
| 0x00203694 | `GotoVolumeLayout` | Known | Navigation |
| 0x002036AC | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x00203738 | `GotoVolumeLayout` | Known | Navigation |
| 0x0020374C | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x0020395C | `GotoScrubLayout` | Known | Navigation |
| 0x0020396C | `GotoScrubVideoLayout` | Known | Navigation |
| 0x002039FC | `GotoProgressLayout` | Known | Navigation |
| 0x00203A10 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00203C68 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x00203C84 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00203C9C | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00203CB8 | `GotoDefaultLayout` | Known | Navigation |
| 0x00203EE4 | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x00203F00 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x0020449C | `GotoChapterArtLayout` | Known | Navigation |
| 0x00204594 | `GotoProgressLayout` | Known | Navigation |
| 0x00204620 | `GotoProgressLayout` | Known | Navigation |
| 0x00204634 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00204710 | `GotoExtraInfoLoadFailedLayout` | Known | Navigation |
| 0x00204730 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x00204B6C | `GotoStatusBarLayout` | Known | Navigation |
| 0x00204B80 | `GotoDefaultLayout` | Known | Navigation |
| 0x00204D58 | `GotoDefault` | Known | Navigation |
| 0x00204E8C | `GotoProgressLayout` | Known | Navigation |
| 0x0020504C | `GotoCaptionVideoLayout` | Known | Navigation |
| 0x0020519C | `GotoBrightnessLayout` | Known | Navigation |
| 0x00205220 | `GotoBrightnessLayout` | Known | Navigation |
| 0x002052A0 | `GotoVolumeLayout` | Known | Navigation |
| 0x002052EC | `GotoScrubLayout` | Known | Navigation |
| 0x002053B4 | `GotoStatusBarLayout` | Known | Navigation |
| 0x002053C8 | `GotoDefaultLayout` | Known | Navigation |
| 0x002054A0 | `GotoScrubLayout` | Known | Navigation |
| 0x002054F0 | `GotoScrubLayout` | Known | Navigation |
| 0x00207FC4 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020B438 | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x0020B454 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020B46C | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x0020B61C | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0020BB10 | `GotoNowPlaying` | Known | Navigation |
| 0x0020BDF8 | `GotoNowPlaying` | Known | Navigation |
| 0x0020CF54 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x0020D0E4 | `GotoFourCard_About` | Known | Navigation |
| 0x0020D0F8 | `GotoThreeCard_About` | Known | Navigation |
| 0x0020D1E4 | `GotoScreen_BacklightTimer` | Known | Navigation |
| 0x0020D274 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x0020D28C | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation |
| 0x00211C30 | `GotoScreen_LockDialog` | Known | Navigation |
| 0x00211C48 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x00212BD8 | `GotoGeniusIntro` | Known | Navigation |
| 0x00212BEC | `GotoGenius` | Known | Navigation |
| 0x00214294 | `GotoNowPlaying` | Known | Navigation |
| 0x002149A4 | `GotoNowPlaying` | Known | Navigation |
| 0x00215188 | `GotoFirstBoot` | Known | Navigation |
| 0x00215198 | `GotoNotesApp` | Known | Navigation |
| 0x002151AC | `GotoLockApp` | Known | Navigation |
| 0x002164E4 | `GotoGenius` | Known | Navigation |
| 0x0021C248 | `GotoGeniusError_NoGenius` | Known | Navigation |
| 0x0021C264 | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0021C27C | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Navigation |
| 0x0021C42C | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x0021CB98 | `GotoNowPlaying` | Known | Navigation |
| 0x003F0A14 | `GotoRatingLayout` | Known | Navigation |
| 0x003F0A28 | `GotoProgressLayout` | Known | Navigation |
| 0x0072610B | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x007A7930 | `GotoDefault` | Known | Navigation |
| 0x007A8914 | `GotoDefault` | Known | Navigation |
| 0x00899624 | `Gotowe?` | Known | Navigation |

---

## 7. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00155F9C | `CoverFlow_Screen` | Known | Screen layout |
| 0x0071A056 | `Clock_Screen` | Known | Screen layout |
| 0x0071A066 | `Clock_Screen_Default"` | Known | Screen layout |
| 0x0071A0CB | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x0071A129 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0071A141 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x0071A1AE | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x0071A24C | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x0071A2AB | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0071A2C1 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x0071A32C | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x0071A386 | `Games_Menu_Screen` | Known | Screen layout |
| 0x0071A39B | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x0071A405 | `Extras_Screen_Games` | Known | Screen layout |
| 0x0071A4C4 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x0071A588 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0071A651 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x0071A6AE | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x0071A6C7 | `Debug_MainMenu_Screen_Default"` | Known | Screen layout |
| 0x0071A735 | `Extras_Screen_Debug` | Known | Screen layout |
| 0x0071A874 | `Clock_Africa_City_Screen ` | Known | Screen layout |
| 0x0071A890 | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x0071A914 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x0071A92E | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x0071A9B0 | `Clock_Atlantic_City_Screen"` | Known | Screen layout |
| 0x0071A9CE | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x0071AA54 | `Clock_Australia_City_Screen#` | Known | Screen layout |
| 0x0071AA73 | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x0071AAFA | `Clock_Europe_City_Screen ` | Known | Screen layout |
| 0x0071AB16 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x0071AB9A | `Clock_NorthAmerica_City_Screen&` | Known | Screen layout |
| 0x0071ABBC | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0071AC46 | `Clock_Pacific_City_Screen!` | Known | Screen layout |
| 0x0071AC63 | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x0071ACE8 | `Clock_SouthAmerica_City_Screen&` | Known | Screen layout |
| 0x0071AD0A | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0071AD97 | `Clock_Screen"` | Known | Screen layout |
| 0x0071AE3C | `Clock_Screen"` | Known | Screen layout |
| 0x0071AEE1 | `Clock_Screen"` | Known | Screen layout |
| 0x0071AF86 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B02B | `Clock_Screen"` | Known | Screen layout |
| 0x0071B0D0 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B175 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B21A | `Clock_Screen"` | Known | Screen layout |
| 0x0071B2BF | `Clock_Screen"` | Known | Screen layout |
| 0x0071B364 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B409 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B4AE | `Clock_Screen"` | Known | Screen layout |
| 0x0071B553 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B5F8 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B69D | `Clock_Screen"` | Known | Screen layout |
| 0x0071B742 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B7E7 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B88C | `Clock_Screen"` | Known | Screen layout |
| 0x0071B931 | `Clock_Screen"` | Known | Screen layout |
| 0x0071B9D6 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BA7B | `Clock_Screen"` | Known | Screen layout |
| 0x0071BB20 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BBC5 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BC6A | `Clock_Screen"` | Known | Screen layout |
| 0x0071BD0F | `Clock_Screen"` | Known | Screen layout |
| 0x0071BDB4 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BE59 | `Clock_Screen"` | Known | Screen layout |
| 0x0071BEFE | `Clock_Screen"` | Known | Screen layout |
| 0x0071BFA3 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C048 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C0ED | `Clock_Screen"` | Known | Screen layout |
| 0x0071C197 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C23C | `Clock_Screen"` | Known | Screen layout |
| 0x0071C2E1 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C386 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C42B | `Clock_Screen"` | Known | Screen layout |
| 0x0071C4D0 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C575 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C61A | `Clock_Screen"` | Known | Screen layout |
| 0x0071C6BF | `Clock_Screen"` | Known | Screen layout |
| 0x0071C764 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C809 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C8AE | `Clock_Screen"` | Known | Screen layout |
| 0x0071C953 | `Clock_Screen"` | Known | Screen layout |
| 0x0071C9F8 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CA9D | `Clock_Screen"` | Known | Screen layout |
| 0x0071CB42 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CBE7 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CC8C | `Clock_Screen"` | Known | Screen layout |
| 0x0071CD31 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CDD6 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CE7B | `Clock_Screen"` | Known | Screen layout |
| 0x0071CF20 | `Clock_Screen"` | Known | Screen layout |
| 0x0071CFC5 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D06A | `Clock_Screen"` | Known | Screen layout |
| 0x0071D10F | `Clock_Screen"` | Known | Screen layout |
| 0x0071D1B4 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D259 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D2FE | `Clock_Screen"` | Known | Screen layout |
| 0x0071D3A3 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D448 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D4ED | `Clock_Screen"` | Known | Screen layout |
| 0x0071D592 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D637 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D6DC | `Clock_Screen"` | Known | Screen layout |
| 0x0071D781 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D826 | `Clock_Screen"` | Known | Screen layout |
| 0x0071D8CB | `Clock_Screen"` | Known | Screen layout |
| 0x0071D970 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DA15 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DABA | `Clock_Screen"` | Known | Screen layout |
| 0x0071DB5F | `Clock_Screen"` | Known | Screen layout |
| 0x0071DC04 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DCA9 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DD4E | `Clock_Screen"` | Known | Screen layout |
| 0x0071DDF3 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DE98 | `Clock_Screen"` | Known | Screen layout |
| 0x0071DF3D | `Clock_Screen"` | Known | Screen layout |
| 0x0071DFE2 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E087 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E12C | `Clock_Screen"` | Known | Screen layout |
| 0x0071E1D1 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E276 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E31B | `Clock_Screen"` | Known | Screen layout |
| 0x0071E3C0 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E465 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E50A | `Clock_Screen"` | Known | Screen layout |
| 0x0071E5AF | `Clock_Screen"` | Known | Screen layout |
| 0x0071E65B | `Clock_Screen"` | Known | Screen layout |
| 0x0071E700 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E7A5 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E84F | `Clock_Screen"` | Known | Screen layout |
| 0x0071E8F4 | `Clock_Screen"` | Known | Screen layout |
| 0x0071E999 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EA3E | `Clock_Screen"` | Known | Screen layout |
| 0x0071EAE3 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EB88 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EC2D | `Clock_Screen"` | Known | Screen layout |
| 0x0071ECD2 | `Clock_Screen"` | Known | Screen layout |
| 0x0071ED7B | `Clock_Screen"` | Known | Screen layout |
| 0x0071EE20 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EEC5 | `Clock_Screen"` | Known | Screen layout |
| 0x0071EF6A | `Clock_Screen"` | Known | Screen layout |
| 0x0071F00F | `Clock_Screen"` | Known | Screen layout |
| 0x0071F0B4 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F159 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F1FE | `Clock_Screen"` | Known | Screen layout |
| 0x0071F2A3 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F348 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F3ED | `Clock_Screen"` | Known | Screen layout |
| 0x0071F492 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F537 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F5DC | `Clock_Screen"` | Known | Screen layout |
| 0x0071F681 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F726 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F7CB | `Clock_Screen"` | Known | Screen layout |
| 0x0071F870 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F915 | `Clock_Screen"` | Known | Screen layout |
| 0x0071F9BA | `Clock_Screen"` | Known | Screen layout |
| 0x0071FA5F | `Clock_Screen"` | Known | Screen layout |
| 0x0071FB04 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FBA9 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FC4E | `Clock_Screen"` | Known | Screen layout |
| 0x0071FCF3 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FD98 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FE3D | `Clock_Screen"` | Known | Screen layout |
| 0x0071FEE2 | `Clock_Screen"` | Known | Screen layout |
| 0x0071FF87 | `Clock_Screen"` | Known | Screen layout |
| 0x0072002C | `Clock_Screen"` | Known | Screen layout |
| 0x007200D1 | `Clock_Screen"` | Known | Screen layout |
| 0x00720176 | `Clock_Screen"` | Known | Screen layout |
| 0x0072021B | `Clock_Screen"` | Known | Screen layout |
| 0x007202C0 | `Clock_Screen"` | Known | Screen layout |
| 0x0072036B | `Clock_Screen"` | Known | Screen layout |
| 0x00720410 | `Clock_Screen"` | Known | Screen layout |
| 0x007204B5 | `Clock_Screen"` | Known | Screen layout |
| 0x0072055A | `Clock_Screen"` | Known | Screen layout |
| 0x007205FF | `Clock_Screen"` | Known | Screen layout |
| 0x007206A4 | `Clock_Screen"` | Known | Screen layout |
| 0x00720749 | `Clock_Screen"` | Known | Screen layout |
| 0x007207EE | `Clock_Screen"` | Known | Screen layout |
| 0x00720893 | `Clock_Screen"` | Known | Screen layout |
| 0x00720938 | `Clock_Screen"` | Known | Screen layout |
| 0x007209DD | `Clock_Screen"` | Known | Screen layout |
| 0x00720A82 | `Clock_Screen"` | Known | Screen layout |
| 0x00720B27 | `Clock_Screen"` | Known | Screen layout |
| 0x00720BCC | `Clock_Screen"` | Known | Screen layout |
| 0x00720C71 | `Clock_Screen"` | Known | Screen layout |
| 0x00720D16 | `Clock_Screen"` | Known | Screen layout |
| 0x00720DBB | `Clock_Screen"` | Known | Screen layout |
| 0x00720E60 | `Clock_Screen"` | Known | Screen layout |
| 0x00720F05 | `Clock_Screen"` | Known | Screen layout |
| 0x00720FAA | `Clock_Screen"` | Known | Screen layout |
| 0x0072104F | `Clock_Screen"` | Known | Screen layout |
| 0x007210F4 | `Clock_Screen"` | Known | Screen layout |
| 0x00721199 | `Clock_Screen"` | Known | Screen layout |
| 0x0072123E | `Clock_Screen"` | Known | Screen layout |
| 0x007212E3 | `Clock_Screen"` | Known | Screen layout |
| 0x00721388 | `Clock_Screen"` | Known | Screen layout |
| 0x0072142D | `Clock_Screen"` | Known | Screen layout |
| 0x007214D2 | `Clock_Screen"` | Known | Screen layout |
| 0x00721577 | `Clock_Screen"` | Known | Screen layout |
| 0x0072161C | `Clock_Screen"` | Known | Screen layout |
| 0x007216C1 | `Clock_Screen"` | Known | Screen layout |
| 0x00721766 | `Clock_Screen"` | Known | Screen layout |
| 0x0072180B | `Clock_Screen"` | Known | Screen layout |
| 0x007218B0 | `Clock_Screen"` | Known | Screen layout |
| 0x00721955 | `Clock_Screen"` | Known | Screen layout |
| 0x007219FA | `Clock_Screen"` | Known | Screen layout |
| 0x00721A9F | `Clock_Screen"` | Known | Screen layout |
| 0x00721B44 | `Clock_Screen"` | Known | Screen layout |
| 0x00721BE9 | `Clock_Screen"` | Known | Screen layout |
| 0x00721C8E | `Clock_Screen"` | Known | Screen layout |
| 0x00721D33 | `Clock_Screen"` | Known | Screen layout |
| 0x00721DD8 | `Clock_Screen"` | Known | Screen layout |
| 0x00721E7D | `Clock_Screen"` | Known | Screen layout |
| 0x00721F22 | `Clock_Screen"` | Known | Screen layout |
| 0x00721FC7 | `Clock_Screen"` | Known | Screen layout |
| 0x0072206C | `Clock_Screen"` | Known | Screen layout |
| 0x00722111 | `Clock_Screen"` | Known | Screen layout |
| 0x007221B6 | `Clock_Screen"` | Known | Screen layout |
| 0x0072225B | `Clock_Screen"` | Known | Screen layout |
| 0x00722300 | `Clock_Screen"` | Known | Screen layout |
| 0x007223AB | `Clock_Screen"` | Known | Screen layout |
| 0x00722450 | `Clock_Screen"` | Known | Screen layout |
| 0x007224F5 | `Clock_Screen"` | Known | Screen layout |
| 0x0072259A | `Clock_Screen"` | Known | Screen layout |
| 0x0072263F | `Clock_Screen"` | Known | Screen layout |
| 0x007226EB | `Clock_Screen"` | Known | Screen layout |
| 0x00722790 | `Clock_Screen"` | Known | Screen layout |
| 0x00722835 | `Clock_Screen"` | Known | Screen layout |
| 0x007228DA | `Clock_Screen"` | Known | Screen layout |
| 0x0072297F | `Clock_Screen"` | Known | Screen layout |
| 0x00722A24 | `Clock_Screen"` | Known | Screen layout |
| 0x00722AC9 | `Clock_Screen"` | Known | Screen layout |
| 0x00722B6E | `Clock_Screen"` | Known | Screen layout |
| 0x00722C13 | `Clock_Screen"` | Known | Screen layout |
| 0x00722CB8 | `Clock_Screen"` | Known | Screen layout |
| 0x00722D5D | `Clock_Screen"` | Known | Screen layout |
| 0x00722E02 | `Clock_Screen"` | Known | Screen layout |
| 0x00722EA7 | `Clock_Screen"` | Known | Screen layout |
| 0x00722F4C | `Clock_Screen"` | Known | Screen layout |
| 0x00722FF1 | `Clock_Screen"` | Known | Screen layout |
| 0x00723096 | `Clock_Screen"` | Known | Screen layout |
| 0x0072313B | `Clock_Screen"` | Known | Screen layout |
| 0x007231DE | `Settings_DateTime_SetDate_Screen(` | Known | Screen layout |
| 0x00723202 | `Settings_DateTime_SetDate_Screen_Default"` | Known | Screen layout |
| 0x0072327B | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x007232E1 | `Settings_DateTime_SetTime_Screen(` | Known | Screen layout |
| 0x00723305 | `Settings_DateTime_SetTime_Screen_Default"` | Known | Screen layout |
| 0x0072337E | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x007233E9 | `Settings_DateTime_SetTimeZone_Screen,` | Known | Screen layout |
| 0x00723411 | `Settings_DateTime_SetTimeZone_Screen_Default"` | Known | Screen layout |
| 0x0072348E | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x00723547 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x007235F7 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x00723BFE | `Search_Main_Screen` | Known | Screen layout |
| 0x00723C14 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x00724136 | `Extras_Screen` | Known | Screen layout |
| 0x00724147 | `Extras_Screen_WorldClock"` | Known | Screen layout |
| 0x007241C4 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x00724226 | `Clock_Screen` | Known | Screen layout |
| 0x00724236 | `Clock_Screen_Default` | Known | Screen layout |
| 0x007242BD | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x00724323 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x00724339 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x007243A4 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x00724406 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0072441E | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x0072448B | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x007244EF | `AddressViewer_Main_Screen!` | Known | Screen layout |
| 0x0072450C | `AddressViewer_Main_Screen_Default"` | Known | Screen layout |
| 0x0072457E | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x007245E5 | `Games_Menu_Screen` | Known | Screen layout |
| 0x007245FA | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x00724664 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x0072472B | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x007247C7 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x00724898 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x00724958 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x007249BC | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007249DB | `VoiceMemos_Menu_Screen_Default"` | Known | Screen layout |
| 0x00724A5E | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x00724AC4 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x00724ADC | `Speakers_Main_Screen_Default"` | Known | Screen layout |
| 0x00724B5D | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x00724BC1 | `Radio_Screen` | Known | Screen layout |
| 0x00724BD1 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x00724C4A | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x00724CAB | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00724D47 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x00724E0A | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x00724EC9 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x00724F86 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x007253D6 | `Radio_Screen` | Known | Screen layout |
| 0x007253E6 | `Radio_Screen_Default"` | Known | Screen layout |
| 0x0072545F | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x00725643 | `Search_Main_Screen` | Known | Screen layout |
| 0x00725659 | `Search_Main_Screen_NoKeyboard"` | Known | Screen layout |
| 0x00725784 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x007257E7 | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x00725B28 | `Video_Settings_Screen` | Known | Screen layout |
| 0x00725B41 | `Video_Settings_Screen_Default"` | Known | Screen layout |
| 0x00725C4A | `PhotosSettings_Screen` | Known | Screen layout |
| 0x00725F0F | `SettingsMenus_MainMenu_Screen%` | Known | Screen layout |
| 0x0072601D | `SettingsMenus_MusicMenu_Screen&` | Known | Screen layout |
| 0x007262C6 | `SettingsMenus_Brightness_Screen&` | Known | Screen layout |
| 0x007263DB | `SettingsMenus_AudiobookSettings_Screen.` | Known | Screen layout |
| 0x00726511 | `SettingsMenus_RadioRegions_Screen ` | Known | Screen layout |
| 0x00726626 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x00726892 | `Settings_DateTime_Screen ` | Known | Screen layout |
| 0x007268AE | `Settings_DateTime_Screen_Default"` | Known | Screen layout |
| 0x00726A3A | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x00726B3F | `Settings_Legal_Screen` | Known | Screen layout |
| 0x00726B58 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x00726C49 | `SettingsMenus_ResetAllSettings_Screen4` | Known | Screen layout |
| 0x0072741A | `Stopwatch_Screen` | Known | Screen layout |
| 0x0072742E | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x00727495 | `Stopwatch_Screen` | Known | Screen layout |
| 0x007274A9 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x00727552 | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x00727575 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x0072760E | `Stopwatch_SessionSummary_Screen'` | Known | Screen layout |
| 0x00727631 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007277E4 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x00727852 | `Speakers_ToneControl_Screen#` | Known | Screen layout |
| 0x00727871 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x0073A955 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073A9D8 | `LockediPod_Screen` | Known | Screen layout |
| 0x0073AA60 | `Lock_Screen` | Known | Screen layout |
| 0x0073AA6F | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073AC0A | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x0073ACDC | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x0073AD46 | `AddressViewer_ContactDetails_Screen+` | Known | Screen layout |
| 0x0073AD6D | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x0073ADE8 | `Extras_Screen` | Known | Screen layout |
| 0x0073AE33 | `Extras_Screen` | Known | Screen layout |
| 0x0073AF1A | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x0073AF78 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0073AF95 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x0073B003 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0073B01C | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0073B093 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0073B0B0 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x0073B11B | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x0073B138 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0073B19F | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x0073B206 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x0073B264 | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0073B281 | `Alarms_Timer_Props_Screen_Default ` | Known | Screen layout |
| 0x0073B2EF | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0073B308 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0073B37F | `Alarms_Timer_Props_Screen!` | Known | Screen layout |
| 0x0073B39C | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x0073B407 | `Alarms_Sleep_Timer_Screen!` | Known | Screen layout |
| 0x0073B424 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0073B48B | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x0073B52B | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x0073B5B4 | `Alarms_Set_Alarm_Frequency_Screen)` | Known | Screen layout |
| 0x0073B5D9 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x0073B64A | `Alarms_Set_Alarm_Label_Screen%` | Known | Screen layout |
| 0x0073B66B | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x0073B6D8 | `Alarms_Set_Alarm_Sound_Screen%` | Known | Screen layout |
| 0x0073B6F9 | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x0073B765 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x0073B9E0 | `Alarms_Set_Alarm_Playlist_Screen(` | Known | Screen layout |
| 0x0073BA04 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x0073BA74 | `Alarms_Set_Alarm_Tones_Screen%` | Known | Screen layout |
| 0x0073BA95 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x0073BDA8 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x0073BDC3 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x0073BF14 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0073BF2B | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x0073BFAC | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0073BFC3 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0073C099 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0073C0B2 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0073C137 | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x0073C1A8 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0073C29D | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0073C2B6 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x0073C33B | `Calendar_Day_Screen_Default'` | Known | Screen layout |
| 0x0073C3AC | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0073C46C | `ToDo_Item_Screen` | Known | Screen layout |
| 0x0073C480 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x0073C5AF | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x0073C612 | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x0073C669 | `Clock_Screen_Default` | Known | Screen layout |
| 0x0073C6FA | `Clock_Region_Screen` | Known | Screen layout |
| 0x0073C711 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x0073C78A | `Clock_Screen_Default` | Known | Screen layout |
| 0x0073C7E1 | `Clock_Screen_Default` | Known | Screen layout |
| 0x0073C872 | `Clock_Region_Screen` | Known | Screen layout |
| 0x0073C889 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x0073CA14 | `CoverFlow_Screen_DefaultNoAlbumInfo&` | Known | Screen layout |
| 0x0073CB02 | `CoverFlow_Screen_DefaultNoAlbumInfo#` | Known | Screen layout |
| 0x0073CB77 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073CE6D | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073D01D | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073D14B | `CoverFlow_Screen_Default&` | Known | Screen layout |
| 0x0073D221 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073D3B6 | `CoverFlow_Screen_Backside"` | Known | Screen layout |
| 0x0073D61B | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x0073D678 | `Game_Screen` | Known | Screen layout |
| 0x0073D687 | `Game_Screen_Default` | Known | Screen layout |
| 0x0073D729 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073D78B | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073D7EE | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073D851 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073D8AD | `Game_Running_Screen` | Known | Screen layout |
| 0x0073D90D | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073D96F | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073D9D2 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073DA35 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073DA91 | `Game_Running_Screen` | Known | Screen layout |
| 0x0073DAF1 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073DB53 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073DBB6 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073DC19 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073DC75 | `Game_Running_Screen` | Known | Screen layout |
| 0x0073DCD5 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073DD37 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073DD9A | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073DDFD | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073DE59 | `Game_Running_Screen` | Known | Screen layout |
| 0x0073DEB9 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073DF1B | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073DF7E | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073DFE1 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073E03D | `Game_Running_Screen` | Known | Screen layout |
| 0x0073E283 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0073E2E5 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0073E348 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0073E3AB | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0073E407 | `Game_Running_Screen` | Known | Screen layout |
| 0x0073E4BE | `Extras_Screen` | Known | Screen layout |
| 0x0073E4CF | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073E52D | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0073E6CA | `Extras_Screen` | Known | Screen layout |
| 0x0073E6DB | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073E739 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0073E8D6 | `Extras_Screen` | Known | Screen layout |
| 0x0073E8E7 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073E945 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0073EAE2 | `Extras_Screen` | Known | Screen layout |
| 0x0073EAF3 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073EB51 | `MainMenus_Main_Screen&` | Known | Screen layout |
| 0x0073ECF3 | `Lock_Screen` | Known | Screen layout |
| 0x0073ED02 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0073ED64 | `Extras_Screen` | Known | Screen layout |
| 0x0073ED75 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0073EDD4 | `LockediPod_Screen` | Known | Screen layout |
| 0x0073EE4E | `MainMenus_Main_Screen)` | Known | Screen layout |
| 0x0073F01F | `Lock_Screen` | Known | Screen layout |
| 0x0073F02E | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0073F090 | `Extras_Screen` | Known | Screen layout |
| 0x0073F0A1 | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0073F100 | `LockediPod_Screen` | Known | Screen layout |
| 0x0073F17A | `MainMenus_Main_Screen%` | Known | Screen layout |
| 0x0073F1E1 | `LockediPod_Screen` | Known | Screen layout |
| 0x0073F1F6 | `LockediPod_Screen_Default,` | Known | Screen layout |
| 0x0073F345 | `Lock_Screen` | Known | Screen layout |
| 0x0073F354 | `Lock_Screen_Default_Layout"` | Known | Screen layout |
| 0x0073F3BD | `Lock_Screen` | Known | Screen layout |
| 0x0073F3CC | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x0073F42E | `Extras_Screen` | Known | Screen layout |
| 0x0073F43F | `Extras_Screen_Lock ` | Known | Screen layout |
| 0x0073F49E | `LockediPod_Screen` | Known | Screen layout |
| 0x0073F518 | `MainMenus_Main_Screen,` | Known | Screen layout |
| 0x0073F674 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0073F6DA | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0073F73E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073F7CD | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0073F83A | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0073F8A7 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0073F914 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073F97C | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0073F9E2 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0073FA46 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073FAD5 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0073FB42 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0073FBAF | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0073FC1C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073FC84 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0073FCEA | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0073FD4E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0073FDDD | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0073FE4A | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x0073FEB7 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0073FF24 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0073FF8C | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0073FFF2 | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x00740056 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007400E5 | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x00740152 | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x007401BF | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x0074022C | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00740294 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x007402FA | `LockConfirmation_Screen ` | Known | Screen layout |
| 0x0074035E | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007403ED | `Lock_Screen_Confirm_Layout$` | Known | Screen layout |
| 0x0074045A | `Lock_Screen_Enter_Layout&` | Known | Screen layout |
| 0x007404C7 | `Lock_Screen_Incorrect_Layout"` | Known | Screen layout |
| 0x00740534 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074058D | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x007405F6 | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x0074065D | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007406F8 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00740761 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x007407CA | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00740831 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007408CC | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00740935 | `SettingsMenus_Main_Screen"` | Known | Screen layout |
| 0x0074099E | `SettingsMenus_Main_Screen!` | Known | Screen layout |
| 0x00740A05 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x00740AA0 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00740B8C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00740BA8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00740C16 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00740C33 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00740C9E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00740CBE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00740D35 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00740D51 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00740DC1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00740DE0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00740E4C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00740E60 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00740ED9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00740F4D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00740FBD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00741024 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074108C | `NoContent_Screen` | Known | Screen layout |
| 0x007410A0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00741104 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074116B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00741185 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007411F3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00741265 | `NoContent_Screen` | Known | Screen layout |
| 0x00741279 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007412E3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074134C | `No_Photos_Screen` | Known | Screen layout |
| 0x00741360 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007413C6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00741434 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007414A1 | `NoContent_Screen` | Known | Screen layout |
| 0x007414B5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074151D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00741587 | `NoContent_Screen` | Known | Screen layout |
| 0x0074159B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00741602 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074166C | `NoContent_Screen` | Known | Screen layout |
| 0x00741680 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007416ED | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074175F | `NoContent_Screen` | Known | Screen layout |
| 0x00741773 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007417DB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00741844 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074185F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007418C5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007418E1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007419C0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007419D9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00741A3A | `FirstBoot_Screen` | Known | Screen layout |
| 0x00741A4E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00741AA8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00741AC4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00741B2B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00741B42 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00741CB3 | `Radio_Screen` | Known | Screen layout |
| 0x00741CC3 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00741D24 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00741DA7 | `LockediPod_Screen` | Known | Screen layout |
| 0x00741E2F | `Lock_Screen` | Known | Screen layout |
| 0x00741E3E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00741EA1 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00741F03 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00741F1F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00741F91 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00741FB0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00742018 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00742032 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074209A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007420B7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00742123 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074218D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007421A7 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00742217 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074228A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007422FB | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074236A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007423D6 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007423F1 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00742466 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007424CD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074252F | `Photos_Screen` | Known | Screen layout |
| 0x00742593 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007425B1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00742623 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00742640 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007426A6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007426C1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074272A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00742747 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007427BE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007427E2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00742850 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074286B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00742928 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00742944 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007429B2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007429CF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00742A3A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00742A5A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00742AD1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00742AED | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00742B5D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00742B7C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00742BE8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00742BFC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00742C75 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00742CE9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00742D59 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00742DC0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00742E28 | `NoContent_Screen` | Known | Screen layout |
| 0x00742E3C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00742EA0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00742F07 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00742F21 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00742F8F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00743001 | `NoContent_Screen` | Known | Screen layout |
| 0x00743015 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074307F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007430E8 | `No_Photos_Screen` | Known | Screen layout |
| 0x007430FC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00743162 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007431D0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074323D | `NoContent_Screen` | Known | Screen layout |
| 0x00743251 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007432B9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00743323 | `NoContent_Screen` | Known | Screen layout |
| 0x00743337 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074339E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00743408 | `NoContent_Screen` | Known | Screen layout |
| 0x0074341C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00743489 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007434FB | `NoContent_Screen` | Known | Screen layout |
| 0x0074350F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00743577 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007435E0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007435FB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00743661 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074367D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074375C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00743775 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007437D6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007437EA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00743844 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00743860 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007438C7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007438DE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00743A4F | `Radio_Screen` | Known | Screen layout |
| 0x00743A5F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00743AC0 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00743B43 | `LockediPod_Screen` | Known | Screen layout |
| 0x00743BCB | `Lock_Screen` | Known | Screen layout |
| 0x00743BDA | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00743C3D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00743C9F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00743CBB | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00743D2D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00743D4C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00743DB4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00743DCE | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00743E36 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00743E53 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00743EBF | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00743F29 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00743F43 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00743FB3 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00744026 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00744097 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00744106 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00744172 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074418D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00744202 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00744269 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007442CB | `Photos_Screen` | Known | Screen layout |
| 0x0074432F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074434D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007443BF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007443DC | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00744442 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074445D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007444C6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007444E3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074455A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074457E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007445EC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00744607 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007446C4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007446E0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074474E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074476B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007447D6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007447F6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074486D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00744889 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007448F9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00744918 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00744984 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00744998 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00744A11 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00744A85 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00744AF5 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00744B5C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00744BC4 | `NoContent_Screen` | Known | Screen layout |
| 0x00744BD8 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00744C3C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00744CA3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00744CBD | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00744D2B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00744D9D | `NoContent_Screen` | Known | Screen layout |
| 0x00744DB1 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00744E1B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00744E84 | `No_Photos_Screen` | Known | Screen layout |
| 0x00744E98 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00744EFE | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00744F6C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00744FD9 | `NoContent_Screen` | Known | Screen layout |
| 0x00744FED | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00745055 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007450BF | `NoContent_Screen` | Known | Screen layout |
| 0x007450D3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074513A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007451A4 | `NoContent_Screen` | Known | Screen layout |
| 0x007451B8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00745225 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00745297 | `NoContent_Screen` | Known | Screen layout |
| 0x007452AB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00745313 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074537C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00745397 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007453FD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00745419 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007454F8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00745511 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00745572 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00745586 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007455E0 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007455FC | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00745663 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074567A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007457EB | `Radio_Screen` | Known | Screen layout |
| 0x007457FB | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074585C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007458DF | `LockediPod_Screen` | Known | Screen layout |
| 0x00745967 | `Lock_Screen` | Known | Screen layout |
| 0x00745976 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007459D9 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00745A3B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00745A57 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00745AC9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00745AE8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00745B50 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00745B6A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00745BD2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00745BEF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00745C5B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00745CC5 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00745CDF | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00745D4F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00745DC2 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00745E33 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00745EA2 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00745F0E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00745F29 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00745F9E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00746005 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00746067 | `Photos_Screen` | Known | Screen layout |
| 0x007460CB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007460E9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074615B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00746178 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007461DE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007461F9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00746262 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074627F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007462F6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074631A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00746388 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007463A3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00746460 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074647C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007464EA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00746507 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00746572 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00746592 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00746609 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00746625 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00746695 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007466B4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00746720 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00746734 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007467AD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00746821 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00746891 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007468F8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00746960 | `NoContent_Screen` | Known | Screen layout |
| 0x00746974 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007469D8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00746A3F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00746A59 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00746AC7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00746B39 | `NoContent_Screen` | Known | Screen layout |
| 0x00746B4D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00746BB7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00746C20 | `No_Photos_Screen` | Known | Screen layout |
| 0x00746C34 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00746C9A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00746D08 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00746D75 | `NoContent_Screen` | Known | Screen layout |
| 0x00746D89 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00746DF1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00746E5B | `NoContent_Screen` | Known | Screen layout |
| 0x00746E6F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00746ED6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00746F40 | `NoContent_Screen` | Known | Screen layout |
| 0x00746F54 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00746FC1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00747033 | `NoContent_Screen` | Known | Screen layout |
| 0x00747047 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007470AF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00747118 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00747133 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00747199 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007471B5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00747294 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007472AD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074730E | `FirstBoot_Screen` | Known | Screen layout |
| 0x00747322 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074737C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00747398 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007473FF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00747416 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00747587 | `Radio_Screen` | Known | Screen layout |
| 0x00747597 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007475F8 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074767B | `LockediPod_Screen` | Known | Screen layout |
| 0x00747703 | `Lock_Screen` | Known | Screen layout |
| 0x00747712 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00747775 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007477D7 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007477F3 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00747865 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00747884 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007478EC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00747906 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074796E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074798B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007479F7 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00747A61 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00747A7B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00747AEB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00747B5E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00747BCF | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00747C3E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00747CAA | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00747CC5 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00747D3A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00747DA1 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00747E03 | `Photos_Screen` | Known | Screen layout |
| 0x00747E67 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00747E85 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00747EF7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00747F14 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00747F7A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00747F95 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00747FFE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074801B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00748092 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007480B6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00748124 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074813F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007481FC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00748218 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00748286 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007482A3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074830E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074832E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007483A5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007483C1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00748431 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00748450 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007484BC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007484D0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00748549 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007485BD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074862D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00748694 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007486FC | `NoContent_Screen` | Known | Screen layout |
| 0x00748710 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00748774 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007487DB | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007487F5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00748863 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007488D5 | `NoContent_Screen` | Known | Screen layout |
| 0x007488E9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00748953 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007489BC | `No_Photos_Screen` | Known | Screen layout |
| 0x007489D0 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00748A36 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00748AA4 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00748B11 | `NoContent_Screen` | Known | Screen layout |
| 0x00748B25 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00748B8D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00748BF7 | `NoContent_Screen` | Known | Screen layout |
| 0x00748C0B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00748C72 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00748CDC | `NoContent_Screen` | Known | Screen layout |
| 0x00748CF0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00748D5D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00748DCF | `NoContent_Screen` | Known | Screen layout |
| 0x00748DE3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00748E4B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00748EB4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00748ECF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00748F35 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00748F51 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00749030 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00749049 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007490AA | `FirstBoot_Screen` | Known | Screen layout |
| 0x007490BE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00749118 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00749134 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074919B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007491B2 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00749323 | `Radio_Screen` | Known | Screen layout |
| 0x00749333 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00749394 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00749417 | `LockediPod_Screen` | Known | Screen layout |
| 0x0074949F | `Lock_Screen` | Known | Screen layout |
| 0x007494AE | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00749511 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00749573 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074958F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00749601 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00749620 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00749688 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007496A2 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074970A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00749727 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00749793 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007497FD | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00749817 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00749887 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007498FA | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074996B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007499DA | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00749A46 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00749A61 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00749AD6 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00749B3D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00749B9F | `Photos_Screen` | Known | Screen layout |
| 0x00749C03 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00749C21 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00749C93 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00749CB0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00749D16 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00749D31 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00749D9A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00749DB7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00749E2E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00749E52 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00749EC0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00749EDB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00749F98 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00749FB4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074A022 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074A03F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074A0AA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074A0CA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074A141 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074A15D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074A1CD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074A1EC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074A258 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074A26C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074A2E5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074A359 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074A3C9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074A430 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074A498 | `NoContent_Screen` | Known | Screen layout |
| 0x0074A4AC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074A510 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074A577 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074A591 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074A5FF | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074A671 | `NoContent_Screen` | Known | Screen layout |
| 0x0074A685 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074A6EF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074A758 | `No_Photos_Screen` | Known | Screen layout |
| 0x0074A76C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074A7D2 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074A840 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074A8AD | `NoContent_Screen` | Known | Screen layout |
| 0x0074A8C1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074A929 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074A993 | `NoContent_Screen` | Known | Screen layout |
| 0x0074A9A7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074AA0E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074AA78 | `NoContent_Screen` | Known | Screen layout |
| 0x0074AA8C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074AAF9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074AB6B | `NoContent_Screen` | Known | Screen layout |
| 0x0074AB7F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074ABE7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074AC50 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074AC6B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074ACD1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074ACED | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074ADCC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074ADE5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074AE46 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074AE5A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074AEB4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074AED0 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074AF37 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074AF4E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074B0BF | `Radio_Screen` | Known | Screen layout |
| 0x0074B0CF | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074B130 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074B1B3 | `LockediPod_Screen` | Known | Screen layout |
| 0x0074B23B | `Lock_Screen` | Known | Screen layout |
| 0x0074B24A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074B2AD | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074B30F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074B32B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0074B39D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074B3BC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074B424 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074B43E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074B4A6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074B4C3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074B52F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074B599 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074B5B3 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074B623 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074B696 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074B707 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074B776 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074B7E2 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074B7FD | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074B872 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074B8D9 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074B93B | `Photos_Screen` | Known | Screen layout |
| 0x0074B99F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074B9BD | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074BA2F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074BA4C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074BAB2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074BACD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074BB36 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074BB53 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074BBCA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074BBEE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074BC5C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074BC77 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074BD34 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074BD50 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074BDBE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074BDDB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074BE46 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074BE66 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074BEDD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074BEF9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074BF69 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074BF88 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074BFF4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074C008 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074C081 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074C0F5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074C165 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074C1CC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074C234 | `NoContent_Screen` | Known | Screen layout |
| 0x0074C248 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074C2AC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074C313 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074C32D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074C39B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074C40D | `NoContent_Screen` | Known | Screen layout |
| 0x0074C421 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074C48B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074C4F4 | `No_Photos_Screen` | Known | Screen layout |
| 0x0074C508 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074C56E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074C5DC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074C649 | `NoContent_Screen` | Known | Screen layout |
| 0x0074C65D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074C6C5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074C72F | `NoContent_Screen` | Known | Screen layout |
| 0x0074C743 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074C7AA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074C814 | `NoContent_Screen` | Known | Screen layout |
| 0x0074C828 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074C895 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074C907 | `NoContent_Screen` | Known | Screen layout |
| 0x0074C91B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074C983 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074C9EC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074CA07 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074CA6D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074CA89 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074CB68 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074CB81 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074CBE2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074CBF6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074CC50 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074CC6C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074CCD3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074CCEA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074CE5B | `Radio_Screen` | Known | Screen layout |
| 0x0074CE6B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074CECC | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074CF4F | `LockediPod_Screen` | Known | Screen layout |
| 0x0074CFD7 | `Lock_Screen` | Known | Screen layout |
| 0x0074CFE6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074D049 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074D0AB | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074D0C7 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0074D139 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074D158 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074D1C0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074D1DA | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074D242 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074D25F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074D2CB | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074D335 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074D34F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074D3BF | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074D432 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074D4A3 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074D512 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074D57E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074D599 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074D60E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074D675 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074D6D7 | `Photos_Screen` | Known | Screen layout |
| 0x0074D73B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074D759 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074D7CB | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074D7E8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074D84E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074D869 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074D8D2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074D8EF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074D966 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074D98A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074D9F8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074DA13 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074DAD0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074DAEC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074DB5A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074DB77 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074DBE2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074DC02 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074DC79 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074DC95 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074DD05 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074DD24 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074DD90 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074DDA4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074DE1D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074DE91 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074DF01 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074DF68 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074DFD0 | `NoContent_Screen` | Known | Screen layout |
| 0x0074DFE4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074E048 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074E0AF | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074E0C9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074E137 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074E1A9 | `NoContent_Screen` | Known | Screen layout |
| 0x0074E1BD | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074E227 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0074E290 | `No_Photos_Screen` | Known | Screen layout |
| 0x0074E2A4 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0074E30A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074E378 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0074E3E5 | `NoContent_Screen` | Known | Screen layout |
| 0x0074E3F9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0074E461 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0074E4CB | `NoContent_Screen` | Known | Screen layout |
| 0x0074E4DF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0074E546 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0074E5B0 | `NoContent_Screen` | Known | Screen layout |
| 0x0074E5C4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0074E631 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0074E6A3 | `NoContent_Screen` | Known | Screen layout |
| 0x0074E6B7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0074E71F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0074E788 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0074E7A3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0074E809 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0074E825 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0074E904 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0074E91D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0074E97E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0074E992 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0074E9EC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0074EA08 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0074EA6F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0074EA86 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0074EBF7 | `Radio_Screen` | Known | Screen layout |
| 0x0074EC07 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0074EC68 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0074ECEB | `LockediPod_Screen` | Known | Screen layout |
| 0x0074ED73 | `Lock_Screen` | Known | Screen layout |
| 0x0074ED82 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0074EDE5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0074EE47 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0074EE63 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0074EED5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0074EEF4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0074EF5C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074EF76 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0074EFDE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074EFFB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074F067 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0074F0D1 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0074F0EB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0074F15B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0074F1CE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0074F23F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0074F2AE | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0074F31A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0074F335 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0074F3AA | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0074F411 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0074F473 | `Photos_Screen` | Known | Screen layout |
| 0x0074F4D7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0074F4F5 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0074F567 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0074F584 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0074F5EA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0074F605 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0074F66E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0074F68B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0074F702 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0074F726 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0074F794 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0074F7AF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0074F86C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074F888 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074F8F6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0074F913 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0074F97E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0074F99E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0074FA15 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0074FA31 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0074FAA1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0074FAC0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0074FB2C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0074FB40 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0074FBB9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0074FC2D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0074FC9D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0074FD04 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0074FD6C | `NoContent_Screen` | Known | Screen layout |
| 0x0074FD80 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0074FDE4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0074FE4B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0074FE65 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0074FED3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0074FF45 | `NoContent_Screen` | Known | Screen layout |
| 0x0074FF59 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0074FFC3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075002C | `No_Photos_Screen` | Known | Screen layout |
| 0x00750040 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007500A6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00750114 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00750181 | `NoContent_Screen` | Known | Screen layout |
| 0x00750195 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007501FD | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00750267 | `NoContent_Screen` | Known | Screen layout |
| 0x0075027B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007502E2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075034C | `NoContent_Screen` | Known | Screen layout |
| 0x00750360 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007503CD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075043F | `NoContent_Screen` | Known | Screen layout |
| 0x00750453 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007504BB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00750524 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075053F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007505A5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007505C1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007506A0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007506B9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075071A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075072E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00750788 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007507A4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075080B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00750822 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00750993 | `Radio_Screen` | Known | Screen layout |
| 0x007509A3 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00750A04 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00750A87 | `LockediPod_Screen` | Known | Screen layout |
| 0x00750B0F | `Lock_Screen` | Known | Screen layout |
| 0x00750B1E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00750B81 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00750BE3 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00750BFF | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00750C71 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00750C90 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00750CF8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00750D12 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00750D7A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00750D97 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00750E03 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00750E6D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00750E87 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00750EF7 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00750F6A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00750FDB | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075104A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007510B6 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007510D1 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00751146 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007511AD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075120F | `Photos_Screen` | Known | Screen layout |
| 0x00751273 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00751291 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00751303 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00751320 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00751386 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007513A1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075140A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00751427 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075149E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007514C2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00751530 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075154B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00751608 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00751624 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00751692 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007516AF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075171A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075173A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007517B1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007517CD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075183D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075185C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007518C8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007518DC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00751955 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007519C9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00751A39 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00751AA0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00751B08 | `NoContent_Screen` | Known | Screen layout |
| 0x00751B1C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00751B80 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00751BE7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00751C01 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00751C6F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00751CE1 | `NoContent_Screen` | Known | Screen layout |
| 0x00751CF5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00751D5F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00751DC8 | `No_Photos_Screen` | Known | Screen layout |
| 0x00751DDC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00751E42 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00751EB0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00751F1D | `NoContent_Screen` | Known | Screen layout |
| 0x00751F31 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00751F99 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00752003 | `NoContent_Screen` | Known | Screen layout |
| 0x00752017 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075207E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007520E8 | `NoContent_Screen` | Known | Screen layout |
| 0x007520FC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00752169 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007521DB | `NoContent_Screen` | Known | Screen layout |
| 0x007521EF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00752257 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007522C0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007522DB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00752341 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075235D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075243C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00752455 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007524B6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007524CA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00752524 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00752540 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007525A7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007525BE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075272F | `Radio_Screen` | Known | Screen layout |
| 0x0075273F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007527A0 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00752823 | `LockediPod_Screen` | Known | Screen layout |
| 0x007528AB | `Lock_Screen` | Known | Screen layout |
| 0x007528BA | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075291D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075297F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075299B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00752A0D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00752A2C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00752A94 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00752AAE | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00752B16 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00752B33 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00752B9F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00752C09 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00752C23 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00752C93 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00752D06 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00752D77 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00752DE6 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00752E52 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00752E6D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00752EE2 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00752F49 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00752FAB | `Photos_Screen` | Known | Screen layout |
| 0x0075300F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075302D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075309F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007530BC | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00753122 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075313D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007531A6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007531C3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075323A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075325E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007532CC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007532E7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007533A4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007533C0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075342E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075344B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007534B6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007534D6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075354D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00753569 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007535D9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007535F8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00753664 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00753678 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007536F1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00753765 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007537D5 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075383C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007538A4 | `NoContent_Screen` | Known | Screen layout |
| 0x007538B8 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075391C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00753983 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075399D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00753A0B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00753A7D | `NoContent_Screen` | Known | Screen layout |
| 0x00753A91 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00753AFB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00753B64 | `No_Photos_Screen` | Known | Screen layout |
| 0x00753B78 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00753BDE | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00753C4C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00753CB9 | `NoContent_Screen` | Known | Screen layout |
| 0x00753CCD | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00753D35 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00753D9F | `NoContent_Screen` | Known | Screen layout |
| 0x00753DB3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00753E1A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00753E84 | `NoContent_Screen` | Known | Screen layout |
| 0x00753E98 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00753F05 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00753F77 | `NoContent_Screen` | Known | Screen layout |
| 0x00753F8B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00753FF3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075405C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00754077 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007540DD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007540F9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007541D8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007541F1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00754252 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00754266 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007542C0 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007542DC | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00754343 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075435A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007544CB | `Radio_Screen` | Known | Screen layout |
| 0x007544DB | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075453C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007545BF | `LockediPod_Screen` | Known | Screen layout |
| 0x00754647 | `Lock_Screen` | Known | Screen layout |
| 0x00754656 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007546B9 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075471B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00754737 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007547A9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007547C8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00754830 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075484A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007548B2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007548CF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075493B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007549A5 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007549BF | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00754A2F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00754AA2 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00754B13 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00754B82 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00754BEE | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00754C09 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00754C7E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00754CE5 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00754D47 | `Photos_Screen` | Known | Screen layout |
| 0x00754DAB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00754DC9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00754E3B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00754E58 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00754EBE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00754ED9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00754F42 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00754F5F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00754FD6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00754FFA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00755068 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00755083 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00755140 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075515C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007551CA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007551E7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00755252 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00755272 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007552E9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00755305 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00755375 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00755394 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00755400 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00755414 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075548D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00755501 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00755571 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007555D8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00755640 | `NoContent_Screen` | Known | Screen layout |
| 0x00755654 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007556B8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075571F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00755739 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007557A7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00755819 | `NoContent_Screen` | Known | Screen layout |
| 0x0075582D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00755897 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00755900 | `No_Photos_Screen` | Known | Screen layout |
| 0x00755914 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075597A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007559E8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00755A55 | `NoContent_Screen` | Known | Screen layout |
| 0x00755A69 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00755AD1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00755B3B | `NoContent_Screen` | Known | Screen layout |
| 0x00755B4F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00755BB6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00755C20 | `NoContent_Screen` | Known | Screen layout |
| 0x00755C34 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00755CA1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00755D13 | `NoContent_Screen` | Known | Screen layout |
| 0x00755D27 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00755D8F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00755DF8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00755E13 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00755E79 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00755E95 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00755F74 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00755F8D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00755FEE | `FirstBoot_Screen` | Known | Screen layout |
| 0x00756002 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075605C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00756078 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007560DF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007560F6 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00756267 | `Radio_Screen` | Known | Screen layout |
| 0x00756277 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007562D8 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075635B | `LockediPod_Screen` | Known | Screen layout |
| 0x007563E3 | `Lock_Screen` | Known | Screen layout |
| 0x007563F2 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00756455 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007564B7 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007564D3 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00756545 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00756564 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007565CC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007565E6 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075664E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075666B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007566D7 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00756741 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075675B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007567CB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075683E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007568AF | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075691E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075698A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007569A5 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00756A1A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00756A81 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00756AE3 | `Photos_Screen` | Known | Screen layout |
| 0x00756B47 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00756B65 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00756BD7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00756BF4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00756C5A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00756C75 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00756CDE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00756CFB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00756D72 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00756D96 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00756E04 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00756E1F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00756EDC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00756EF8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00756F66 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00756F83 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00756FEE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075700E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00757085 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007570A1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00757111 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00757130 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075719C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007571B0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00757229 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075729D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075730D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00757374 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007573DC | `NoContent_Screen` | Known | Screen layout |
| 0x007573F0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00757454 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007574BB | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007574D5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00757543 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007575B5 | `NoContent_Screen` | Known | Screen layout |
| 0x007575C9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00757633 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075769C | `No_Photos_Screen` | Known | Screen layout |
| 0x007576B0 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00757716 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00757784 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007577F1 | `NoContent_Screen` | Known | Screen layout |
| 0x00757805 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075786D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007578D7 | `NoContent_Screen` | Known | Screen layout |
| 0x007578EB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00757952 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007579BC | `NoContent_Screen` | Known | Screen layout |
| 0x007579D0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00757A3D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00757AAF | `NoContent_Screen` | Known | Screen layout |
| 0x00757AC3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00757B2B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00757B94 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00757BAF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00757C15 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00757C31 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00757D10 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00757D29 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00757D8A | `FirstBoot_Screen` | Known | Screen layout |
| 0x00757D9E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00757DF8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00757E14 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00757E7B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00757E92 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00758003 | `Radio_Screen` | Known | Screen layout |
| 0x00758013 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00758074 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007580F7 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075817F | `Lock_Screen` | Known | Screen layout |
| 0x0075818E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007581F1 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00758253 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075826F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007582E1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00758300 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00758368 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00758382 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007583EA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00758407 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00758473 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007584DD | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007584F7 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00758567 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007585DA | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075864B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007586BA | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00758726 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00758741 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007587B6 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075881D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075887F | `Photos_Screen` | Known | Screen layout |
| 0x007588E3 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00758901 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00758973 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00758990 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007589F6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00758A11 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00758A7A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00758A97 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00758B0E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00758B32 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00758BA0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00758BBB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00758C78 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00758C94 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00758D02 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00758D1F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00758D8A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00758DAA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00758E21 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00758E3D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00758EAD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00758ECC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00758F38 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00758F4C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00758FC5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00759039 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007590A9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00759110 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00759178 | `NoContent_Screen` | Known | Screen layout |
| 0x0075918C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007591F0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00759257 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00759271 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007592DF | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00759351 | `NoContent_Screen` | Known | Screen layout |
| 0x00759365 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007593CF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00759438 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075944C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007594B2 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00759520 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075958D | `NoContent_Screen` | Known | Screen layout |
| 0x007595A1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00759609 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00759673 | `NoContent_Screen` | Known | Screen layout |
| 0x00759687 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007596EE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00759758 | `NoContent_Screen` | Known | Screen layout |
| 0x0075976C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007597D9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075984B | `NoContent_Screen` | Known | Screen layout |
| 0x0075985F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007598C7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00759930 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075994B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007599B1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007599CD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00759AAC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00759AC5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00759B26 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00759B3A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00759B94 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00759BB0 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00759C17 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00759C2E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00759D9F | `Radio_Screen` | Known | Screen layout |
| 0x00759DAF | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00759E10 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00759E93 | `LockediPod_Screen` | Known | Screen layout |
| 0x00759F1B | `Lock_Screen` | Known | Screen layout |
| 0x00759F2A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00759F8D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00759FEF | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075A00B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075A07D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075A09C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075A104 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075A11E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075A186 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075A1A3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075A20F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075A279 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075A293 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075A303 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075A376 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075A3E7 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075A456 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075A4C2 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075A4DD | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075A552 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075A5B9 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075A61B | `Photos_Screen` | Known | Screen layout |
| 0x0075A67F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075A69D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075A70F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075A72C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075A792 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075A7AD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075A816 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075A833 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075A8AA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075A8CE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075A93C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075A957 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075AA14 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075AA30 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075AA9E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075AABB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075AB26 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075AB46 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075ABBD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075ABD9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075AC49 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075AC68 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075ACD4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075ACE8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075AD61 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075ADD5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075AE45 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075AEAC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075AF14 | `NoContent_Screen` | Known | Screen layout |
| 0x0075AF28 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075AF8C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075AFF3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075B00D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075B07B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075B0ED | `NoContent_Screen` | Known | Screen layout |
| 0x0075B101 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075B16B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075B1D4 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075B1E8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075B24E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075B2BC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075B329 | `NoContent_Screen` | Known | Screen layout |
| 0x0075B33D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075B3A5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075B40F | `NoContent_Screen` | Known | Screen layout |
| 0x0075B423 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075B48A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075B4F4 | `NoContent_Screen` | Known | Screen layout |
| 0x0075B508 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075B575 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075B5E7 | `NoContent_Screen` | Known | Screen layout |
| 0x0075B5FB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075B663 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075B6CC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075B6E7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075B74D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075B769 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075B848 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075B861 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075B8C2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075B8D6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075B930 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075B94C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075B9B3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075B9CA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075BB3B | `Radio_Screen` | Known | Screen layout |
| 0x0075BB4B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075BBAC | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075BC2F | `LockediPod_Screen` | Known | Screen layout |
| 0x0075BCB7 | `Lock_Screen` | Known | Screen layout |
| 0x0075BCC6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075BD29 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075BD8B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075BDA7 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075BE19 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075BE38 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075BEA0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075BEBA | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075BF22 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075BF3F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075BFAB | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075C015 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075C02F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075C09F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075C112 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075C183 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075C1F2 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075C25E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075C279 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075C2EE | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075C355 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075C3B7 | `Photos_Screen` | Known | Screen layout |
| 0x0075C41B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075C439 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075C4AB | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075C4C8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075C52E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075C549 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075C5B2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075C5CF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075C646 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075C66A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075C6D8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075C6F3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075C7B0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075C7CC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075C83A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075C857 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075C8C2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075C8E2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075C959 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075C975 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075C9E5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075CA04 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075CA70 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075CA84 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075CAFD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075CB71 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075CBE1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075CC48 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075CCB0 | `NoContent_Screen` | Known | Screen layout |
| 0x0075CCC4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075CD28 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075CD8F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075CDA9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075CE17 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075CE89 | `NoContent_Screen` | Known | Screen layout |
| 0x0075CE9D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075CF07 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075CF70 | `No_Photos_Screen` | Known | Screen layout |
| 0x0075CF84 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075CFEA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075D058 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075D0C5 | `NoContent_Screen` | Known | Screen layout |
| 0x0075D0D9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075D141 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075D1AB | `NoContent_Screen` | Known | Screen layout |
| 0x0075D1BF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075D226 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075D290 | `NoContent_Screen` | Known | Screen layout |
| 0x0075D2A4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075D311 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075D383 | `NoContent_Screen` | Known | Screen layout |
| 0x0075D397 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075D3FF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075D468 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075D483 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075D4E9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075D505 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075D5E4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075D5FD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075D65E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075D672 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075D6CC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075D6E8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075D74F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075D766 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075D8D7 | `Radio_Screen` | Known | Screen layout |
| 0x0075D8E7 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075D948 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075D9CB | `LockediPod_Screen` | Known | Screen layout |
| 0x0075DA53 | `Lock_Screen` | Known | Screen layout |
| 0x0075DA62 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075DAC5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075DB27 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075DB43 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075DBB5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075DBD4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075DC3C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075DC56 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075DCBE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075DCDB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075DD47 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075DDB1 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075DDCB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075DE3B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075DEAE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075DF1F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075DF8E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075DFFA | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075E015 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075E08A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075E0F1 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075E153 | `Photos_Screen` | Known | Screen layout |
| 0x0075E1B7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075E1D5 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075E247 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0075E264 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0075E2CA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0075E2E5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0075E34E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0075E36B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0075E3E2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0075E406 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0075E474 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0075E48F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0075E54C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075E568 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075E5D6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075E5F3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075E65E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0075E67E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0075E6F5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0075E711 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0075E781 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0075E7A0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0075E80C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0075E820 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0075E899 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0075E90D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0075E97D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0075E9E4 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0075EA4C | `NoContent_Screen` | Known | Screen layout |
| 0x0075EA60 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0075EAC4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0075EB2B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075EB45 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0075EBB3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0075EC25 | `NoContent_Screen` | Known | Screen layout |
| 0x0075EC39 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0075ECA3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0075ED0C | `No_Photos_Screen` | Known | Screen layout |
| 0x0075ED20 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0075ED86 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075EDF4 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0075EE61 | `NoContent_Screen` | Known | Screen layout |
| 0x0075EE75 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0075EEDD | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0075EF47 | `NoContent_Screen` | Known | Screen layout |
| 0x0075EF5B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0075EFC2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0075F02C | `NoContent_Screen` | Known | Screen layout |
| 0x0075F040 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0075F0AD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0075F11F | `NoContent_Screen` | Known | Screen layout |
| 0x0075F133 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0075F19B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0075F204 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0075F21F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0075F285 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0075F2A1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0075F380 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0075F399 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0075F3FA | `FirstBoot_Screen` | Known | Screen layout |
| 0x0075F40E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0075F468 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0075F484 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0075F4EB | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0075F502 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0075F673 | `Radio_Screen` | Known | Screen layout |
| 0x0075F683 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0075F6E4 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0075F767 | `LockediPod_Screen` | Known | Screen layout |
| 0x0075F7EF | `Lock_Screen` | Known | Screen layout |
| 0x0075F7FE | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0075F861 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0075F8C3 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0075F8DF | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0075F951 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0075F970 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0075F9D8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0075F9F2 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0075FA5A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0075FA77 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0075FAE3 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0075FB4D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0075FB67 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0075FBD7 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0075FC4A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0075FCBB | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0075FD2A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0075FD96 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0075FDB1 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0075FE26 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0075FE8D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0075FEEF | `Photos_Screen` | Known | Screen layout |
| 0x0075FF53 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0075FF71 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0075FFE3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00760000 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00760066 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00760081 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007600EA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00760107 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076017E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007601A2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00760210 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076022B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007602E8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00760304 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00760372 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076038F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007603FA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076041A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00760491 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007604AD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076051D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076053C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007605A8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007605BC | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00760635 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007606A9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00760719 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00760780 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007607E8 | `NoContent_Screen` | Known | Screen layout |
| 0x007607FC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00760860 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007608C7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007608E1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076094F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007609C1 | `NoContent_Screen` | Known | Screen layout |
| 0x007609D5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00760A3F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00760AA8 | `No_Photos_Screen` | Known | Screen layout |
| 0x00760ABC | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00760B22 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00760B90 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00760BFD | `NoContent_Screen` | Known | Screen layout |
| 0x00760C11 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00760C79 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00760CE3 | `NoContent_Screen` | Known | Screen layout |
| 0x00760CF7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00760D5E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00760DC8 | `NoContent_Screen` | Known | Screen layout |
| 0x00760DDC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00760E49 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00760EBB | `NoContent_Screen` | Known | Screen layout |
| 0x00760ECF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00760F37 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00760FA0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00760FBB | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00761021 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076103D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076111C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00761135 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00761196 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007611AA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00761204 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00761220 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00761287 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076129E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076140F | `Radio_Screen` | Known | Screen layout |
| 0x0076141F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00761480 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00761503 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076158B | `Lock_Screen` | Known | Screen layout |
| 0x0076159A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007615FD | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076165F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076167B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007616ED | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076170C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00761774 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076178E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007617F6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00761813 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076187F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007618E9 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00761903 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00761973 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007619E6 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00761A57 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00761AC6 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00761B32 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00761B4D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00761BC2 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00761C29 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00761C8B | `Photos_Screen` | Known | Screen layout |
| 0x00761CEF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00761D0D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00761D7F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00761D9C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00761E02 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00761E1D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00761E86 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00761EA3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00761F1A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00761F3E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00761FAC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00761FC7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00762084 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007620A0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076210E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076212B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00762196 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007621B6 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076222D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00762249 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007622B9 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007622D8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00762344 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00762358 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007623D1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00762445 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007624B5 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076251C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00762584 | `NoContent_Screen` | Known | Screen layout |
| 0x00762598 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007625FC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00762663 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076267D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007626EB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076275D | `NoContent_Screen` | Known | Screen layout |
| 0x00762771 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007627DB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00762844 | `No_Photos_Screen` | Known | Screen layout |
| 0x00762858 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007628BE | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076292C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00762999 | `NoContent_Screen` | Known | Screen layout |
| 0x007629AD | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00762A15 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00762A7F | `NoContent_Screen` | Known | Screen layout |
| 0x00762A93 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00762AFA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00762B64 | `NoContent_Screen` | Known | Screen layout |
| 0x00762B78 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00762BE5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00762C57 | `NoContent_Screen` | Known | Screen layout |
| 0x00762C6B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00762CD3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00762D3C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00762D57 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00762DBD | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00762DD9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00762EB8 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00762ED1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00762F32 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00762F46 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00762FA0 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00762FBC | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00763023 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076303A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007631AB | `Radio_Screen` | Known | Screen layout |
| 0x007631BB | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076321C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076329F | `LockediPod_Screen` | Known | Screen layout |
| 0x00763327 | `Lock_Screen` | Known | Screen layout |
| 0x00763336 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00763399 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007633FB | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00763417 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00763489 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007634A8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00763510 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076352A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00763592 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007635AF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076361B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00763685 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076369F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076370F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00763782 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007637F3 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00763862 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007638CE | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007638E9 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076395E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007639C5 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00763A27 | `Photos_Screen` | Known | Screen layout |
| 0x00763A8B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00763AA9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00763B1B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00763B38 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00763B9E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00763BB9 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00763C22 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00763C3F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00763CB6 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00763CDA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00763D48 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00763D63 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00763E20 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00763E3C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00763EAA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00763EC7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00763F32 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00763F52 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00763FC9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00763FE5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00764055 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00764074 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007640E0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007640F4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076416D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007641E1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00764251 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007642B8 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00764320 | `NoContent_Screen` | Known | Screen layout |
| 0x00764334 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00764398 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007643FF | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00764419 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00764487 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007644F9 | `NoContent_Screen` | Known | Screen layout |
| 0x0076450D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00764577 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007645E0 | `No_Photos_Screen` | Known | Screen layout |
| 0x007645F4 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076465A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007646C8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00764735 | `NoContent_Screen` | Known | Screen layout |
| 0x00764749 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007647B1 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076481B | `NoContent_Screen` | Known | Screen layout |
| 0x0076482F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00764896 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00764900 | `NoContent_Screen` | Known | Screen layout |
| 0x00764914 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00764981 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007649F3 | `NoContent_Screen` | Known | Screen layout |
| 0x00764A07 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00764A6F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00764AD8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00764AF3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00764B59 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00764B75 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00764C54 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00764C6D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00764CCE | `FirstBoot_Screen` | Known | Screen layout |
| 0x00764CE2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00764D3C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00764D58 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00764DBF | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00764DD6 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00764F47 | `Radio_Screen` | Known | Screen layout |
| 0x00764F57 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00764FB8 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076503B | `LockediPod_Screen` | Known | Screen layout |
| 0x007650C3 | `Lock_Screen` | Known | Screen layout |
| 0x007650D2 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00765135 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00765197 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007651B3 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00765225 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00765244 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007652AC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007652C6 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076532E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076534B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007653B7 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00765421 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076543B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007654AB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076551E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076558F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007655FE | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076566A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00765685 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007656FA | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00765761 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007657C3 | `Photos_Screen` | Known | Screen layout |
| 0x00765827 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00765845 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007658B7 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007658D4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076593A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00765955 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007659BE | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007659DB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00765A52 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00765A76 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00765AE4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00765AFF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00765BBC | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00765BD8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00765C46 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00765C63 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00765CCE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00765CEE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00765D65 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00765D81 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00765DF1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00765E10 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00765E7C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00765E90 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00765F09 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00765F7D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00765FED | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00766054 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007660BC | `NoContent_Screen` | Known | Screen layout |
| 0x007660D0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00766134 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076619B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007661B5 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00766223 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00766295 | `NoContent_Screen` | Known | Screen layout |
| 0x007662A9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00766313 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076637C | `No_Photos_Screen` | Known | Screen layout |
| 0x00766390 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007663F6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00766464 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007664D1 | `NoContent_Screen` | Known | Screen layout |
| 0x007664E5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076654D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007665B7 | `NoContent_Screen` | Known | Screen layout |
| 0x007665CB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00766632 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076669C | `NoContent_Screen` | Known | Screen layout |
| 0x007666B0 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076671D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076678F | `NoContent_Screen` | Known | Screen layout |
| 0x007667A3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076680B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00766874 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076688F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007668F5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00766911 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007669F0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00766A09 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00766A6A | `FirstBoot_Screen` | Known | Screen layout |
| 0x00766A7E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00766AD8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00766AF4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00766B5B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00766B72 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00766CE3 | `Radio_Screen` | Known | Screen layout |
| 0x00766CF3 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00766D54 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00766DD7 | `LockediPod_Screen` | Known | Screen layout |
| 0x00766E5F | `Lock_Screen` | Known | Screen layout |
| 0x00766E6E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00766ED1 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00766F33 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00766F4F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00766FC1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00766FE0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00767048 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00767062 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007670CA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007670E7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00767153 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007671BD | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007671D7 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00767247 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007672BA | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076732B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076739A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00767406 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00767421 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00767496 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007674FD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076755F | `Photos_Screen` | Known | Screen layout |
| 0x007675C3 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007675E1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00767653 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00767670 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007676D6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007676F1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076775A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00767777 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007677EE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00767812 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00767880 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076789B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00767958 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00767974 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007679E2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007679FF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00767A6A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00767A8A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00767B01 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00767B1D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00767B8D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00767BAC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00767C18 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00767C2C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00767CA5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00767D19 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00767D89 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00767DF0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00767E58 | `NoContent_Screen` | Known | Screen layout |
| 0x00767E6C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00767ED0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00767F37 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00767F51 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00767FBF | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00768031 | `NoContent_Screen` | Known | Screen layout |
| 0x00768045 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007680AF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00768118 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076812C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00768192 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00768200 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076826D | `NoContent_Screen` | Known | Screen layout |
| 0x00768281 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007682E9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00768353 | `NoContent_Screen` | Known | Screen layout |
| 0x00768367 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007683CE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00768438 | `NoContent_Screen` | Known | Screen layout |
| 0x0076844C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007684B9 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076852B | `NoContent_Screen` | Known | Screen layout |
| 0x0076853F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007685A7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00768610 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076862B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00768691 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007686AD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076878C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007687A5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00768806 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076881A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00768874 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00768890 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007688F7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076890E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00768A7F | `Radio_Screen` | Known | Screen layout |
| 0x00768A8F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00768AF0 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00768B73 | `LockediPod_Screen` | Known | Screen layout |
| 0x00768BFB | `Lock_Screen` | Known | Screen layout |
| 0x00768C0A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00768C6D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00768CCF | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00768CEB | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00768D5D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00768D7C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00768DE4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00768DFE | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00768E66 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00768E83 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00768EEF | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00768F59 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00768F73 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00768FE3 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00769056 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007690C7 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00769136 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007691A2 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007691BD | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00769232 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00769299 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007692FB | `Photos_Screen` | Known | Screen layout |
| 0x0076935F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076937D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007693EF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076940C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00769472 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076948D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007694F6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00769513 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076958A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007695AE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076961C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00769637 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007696F4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00769710 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076977E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076979B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00769806 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00769826 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076989D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007698B9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00769929 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00769948 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007699B4 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007699C8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00769A41 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00769AB5 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00769B25 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00769B8C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00769BF4 | `NoContent_Screen` | Known | Screen layout |
| 0x00769C08 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00769C6C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00769CD3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00769CED | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00769D5B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00769DCD | `NoContent_Screen` | Known | Screen layout |
| 0x00769DE1 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00769E4B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00769EB4 | `No_Photos_Screen` | Known | Screen layout |
| 0x00769EC8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00769F2E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00769F9C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076A009 | `NoContent_Screen` | Known | Screen layout |
| 0x0076A01D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076A085 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076A0EF | `NoContent_Screen` | Known | Screen layout |
| 0x0076A103 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076A16A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076A1D4 | `NoContent_Screen` | Known | Screen layout |
| 0x0076A1E8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076A255 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076A2C7 | `NoContent_Screen` | Known | Screen layout |
| 0x0076A2DB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076A343 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076A3AC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076A3C7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076A42D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076A449 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076A528 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076A541 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076A5A2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076A5B6 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076A610 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076A62C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076A693 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076A6AA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076A81B | `Radio_Screen` | Known | Screen layout |
| 0x0076A82B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076A88C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076A90F | `LockediPod_Screen` | Known | Screen layout |
| 0x0076A997 | `Lock_Screen` | Known | Screen layout |
| 0x0076A9A6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076AA09 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076AA6B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076AA87 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076AAF9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076AB18 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076AB80 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076AB9A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076AC02 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076AC1F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076AC8B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076ACF5 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076AD0F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076AD7F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076ADF2 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076AE63 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076AED2 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076AF3E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076AF59 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076AFCE | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076B035 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076B097 | `Photos_Screen` | Known | Screen layout |
| 0x0076B0FB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076B119 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076B18B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076B1A8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076B20E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076B229 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076B292 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076B2AF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076B326 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076B34A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076B3B8 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076B3D3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076B490 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076B4AC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076B51A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076B537 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076B5A2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076B5C2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076B639 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076B655 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076B6C5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076B6E4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076B750 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076B764 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076B7DD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076B851 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076B8C1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076B928 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076B990 | `NoContent_Screen` | Known | Screen layout |
| 0x0076B9A4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076BA08 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076BA6F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076BA89 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076BAF7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076BB69 | `NoContent_Screen` | Known | Screen layout |
| 0x0076BB7D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076BBE7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076BC50 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076BC64 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076BCCA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076BD38 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076BDA5 | `NoContent_Screen` | Known | Screen layout |
| 0x0076BDB9 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076BE21 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076BE8B | `NoContent_Screen` | Known | Screen layout |
| 0x0076BE9F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076BF06 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076BF70 | `NoContent_Screen` | Known | Screen layout |
| 0x0076BF84 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076BFF1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076C063 | `NoContent_Screen` | Known | Screen layout |
| 0x0076C077 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076C0DF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076C148 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076C163 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076C1C9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076C1E5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076C2C4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076C2DD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076C33E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076C352 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076C3AC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076C3C8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076C42F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076C446 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076C5B7 | `Radio_Screen` | Known | Screen layout |
| 0x0076C5C7 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076C628 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076C6AB | `LockediPod_Screen` | Known | Screen layout |
| 0x0076C733 | `Lock_Screen` | Known | Screen layout |
| 0x0076C742 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076C7A5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076C807 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076C823 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076C895 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076C8B4 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076C91C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076C936 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076C99E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076C9BB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076CA27 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076CA91 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076CAAB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076CB1B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076CB8E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076CBFF | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076CC6E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076CCDA | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076CCF5 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076CD6A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076CDD1 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076CE33 | `Photos_Screen` | Known | Screen layout |
| 0x0076CE97 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076CEB5 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076CF27 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076CF44 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076CFAA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076CFC5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076D02E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076D04B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076D0C2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076D0E6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076D154 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076D16F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076D22C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076D248 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076D2B6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076D2D3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076D33E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076D35E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076D3D5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076D3F1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076D461 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076D480 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076D4EC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076D500 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076D579 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076D5ED | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076D65D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076D6C4 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076D72C | `NoContent_Screen` | Known | Screen layout |
| 0x0076D740 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076D7A4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076D80B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076D825 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076D893 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076D905 | `NoContent_Screen` | Known | Screen layout |
| 0x0076D919 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076D983 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076D9EC | `No_Photos_Screen` | Known | Screen layout |
| 0x0076DA00 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076DA66 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076DAD4 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076DB41 | `NoContent_Screen` | Known | Screen layout |
| 0x0076DB55 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076DBBD | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076DC27 | `NoContent_Screen` | Known | Screen layout |
| 0x0076DC3B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076DCA2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076DD0C | `NoContent_Screen` | Known | Screen layout |
| 0x0076DD20 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076DD8D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076DDFF | `NoContent_Screen` | Known | Screen layout |
| 0x0076DE13 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076DE7B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076DEE4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076DEFF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076DF65 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076DF81 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076E060 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076E079 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076E0DA | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076E0EE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076E148 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076E164 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076E1CB | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076E1E2 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0076E353 | `Radio_Screen` | Known | Screen layout |
| 0x0076E363 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0076E3C4 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0076E447 | `LockediPod_Screen` | Known | Screen layout |
| 0x0076E4CF | `Lock_Screen` | Known | Screen layout |
| 0x0076E4DE | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0076E541 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0076E5A3 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0076E5BF | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0076E631 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0076E650 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0076E6B8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076E6D2 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0076E73A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076E757 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076E7C3 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0076E82D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0076E847 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0076E8B7 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0076E92A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0076E99B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0076EA0A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0076EA76 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0076EA91 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0076EB06 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0076EB6D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0076EBCF | `Photos_Screen` | Known | Screen layout |
| 0x0076EC33 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0076EC51 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0076ECC3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0076ECE0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0076ED46 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0076ED61 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0076EDCA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0076EDE7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0076EE5E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0076EE82 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0076EEF0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0076EF0B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0076EFC8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076EFE4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076F052 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0076F06F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0076F0DA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0076F0FA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0076F171 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0076F18D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0076F1FD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0076F21C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0076F288 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0076F29C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0076F315 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0076F389 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0076F3F9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0076F460 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0076F4C8 | `NoContent_Screen` | Known | Screen layout |
| 0x0076F4DC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0076F540 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0076F5A7 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0076F5C1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0076F62F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0076F6A1 | `NoContent_Screen` | Known | Screen layout |
| 0x0076F6B5 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0076F71F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0076F788 | `No_Photos_Screen` | Known | Screen layout |
| 0x0076F79C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0076F802 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076F870 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0076F8DD | `NoContent_Screen` | Known | Screen layout |
| 0x0076F8F1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0076F959 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0076F9C3 | `NoContent_Screen` | Known | Screen layout |
| 0x0076F9D7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0076FA3E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0076FAA8 | `NoContent_Screen` | Known | Screen layout |
| 0x0076FABC | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0076FB29 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0076FB9B | `NoContent_Screen` | Known | Screen layout |
| 0x0076FBAF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0076FC17 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0076FC80 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0076FC9B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0076FD01 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0076FD1D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0076FDFC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0076FE15 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0076FE76 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0076FE8A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0076FEE4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0076FF00 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0076FF67 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0076FF7E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007700EF | `Radio_Screen` | Known | Screen layout |
| 0x007700FF | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00770160 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007701E3 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077026B | `Lock_Screen` | Known | Screen layout |
| 0x0077027A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007702DD | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077033F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077035B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007703CD | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007703EC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00770454 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077046E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007704D6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007704F3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077055F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007705C9 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007705E3 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00770653 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007706C6 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00770737 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007707A6 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00770812 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077082D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007708A2 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00770909 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077096B | `Photos_Screen` | Known | Screen layout |
| 0x007709CF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007709ED | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00770A5F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00770A7C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00770AE2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00770AFD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00770B66 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00770B83 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00770BFA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00770C1E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00770C8C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00770CA7 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00770D64 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00770D80 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00770DEE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00770E0B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00770E76 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00770E96 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00770F0D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00770F29 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00770F99 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00770FB8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00771024 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00771038 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007710B1 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00771125 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00771195 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x007711FC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00771264 | `NoContent_Screen` | Known | Screen layout |
| 0x00771278 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x007712DC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00771343 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077135D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007713CB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077143D | `NoContent_Screen` | Known | Screen layout |
| 0x00771451 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007714BB | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00771524 | `No_Photos_Screen` | Known | Screen layout |
| 0x00771538 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077159E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077160C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00771679 | `NoContent_Screen` | Known | Screen layout |
| 0x0077168D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007716F5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077175F | `NoContent_Screen` | Known | Screen layout |
| 0x00771773 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007717DA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00771844 | `NoContent_Screen` | Known | Screen layout |
| 0x00771858 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007718C5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00771937 | `NoContent_Screen` | Known | Screen layout |
| 0x0077194B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007719B3 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00771A1C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00771A37 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00771A9D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00771AB9 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00771B98 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00771BB1 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00771C12 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00771C26 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00771C80 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00771C9C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00771D03 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00771D1A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00771E8B | `Radio_Screen` | Known | Screen layout |
| 0x00771E9B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00771EFC | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00771F7F | `LockediPod_Screen` | Known | Screen layout |
| 0x00772007 | `Lock_Screen` | Known | Screen layout |
| 0x00772016 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00772079 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007720DB | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007720F7 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00772169 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00772188 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007721F0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077220A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00772272 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077228F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007722FB | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00772365 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077237F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007723EF | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00772462 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007724D3 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00772542 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007725AE | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007725C9 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077263E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007726A5 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00772707 | `Photos_Screen` | Known | Screen layout |
| 0x0077276B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00772789 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007727FB | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00772818 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077287E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00772899 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00772902 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077291F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00772996 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007729BA | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00772A28 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00772A43 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00772B00 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00772B1C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00772B8A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00772BA7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00772C12 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00772C32 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00772CA9 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00772CC5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00772D35 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00772D54 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00772DC0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00772DD4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00772E4D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00772EC1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00772F31 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00772F98 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00773000 | `NoContent_Screen` | Known | Screen layout |
| 0x00773014 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00773078 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007730DF | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007730F9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00773167 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007731D9 | `NoContent_Screen` | Known | Screen layout |
| 0x007731ED | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00773257 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x007732C0 | `No_Photos_Screen` | Known | Screen layout |
| 0x007732D4 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077333A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007733A8 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00773415 | `NoContent_Screen` | Known | Screen layout |
| 0x00773429 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00773491 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007734FB | `NoContent_Screen` | Known | Screen layout |
| 0x0077350F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00773576 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007735E0 | `NoContent_Screen` | Known | Screen layout |
| 0x007735F4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00773661 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007736D3 | `NoContent_Screen` | Known | Screen layout |
| 0x007736E7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077374F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007737B8 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007737D3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00773839 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00773855 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00773934 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077394D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007739AE | `FirstBoot_Screen` | Known | Screen layout |
| 0x007739C2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00773A1C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00773A38 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00773A9F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00773AB6 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00773C27 | `Radio_Screen` | Known | Screen layout |
| 0x00773C37 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00773C98 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00773D1B | `LockediPod_Screen` | Known | Screen layout |
| 0x00773DA3 | `Lock_Screen` | Known | Screen layout |
| 0x00773DB2 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00773E15 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00773E77 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00773E93 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00773F05 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00773F24 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00773F8C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00773FA6 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077400E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077402B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00774097 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00774101 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077411B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077418B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x007741FE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077426F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x007742DE | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077434A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00774365 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007743DA | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00774441 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007744A3 | `Photos_Screen` | Known | Screen layout |
| 0x00774507 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00774525 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00774597 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007745B4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077461A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00774635 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077469E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007746BB | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00774732 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00774756 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007747C4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007747DF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077489C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007748B8 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00774926 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00774943 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007749AE | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007749CE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00774A45 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00774A61 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00774AD1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00774AF0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00774B5C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00774B70 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00774BE9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00774C5D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00774CCD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00774D34 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00774D9C | `NoContent_Screen` | Known | Screen layout |
| 0x00774DB0 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00774E14 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00774E7B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00774E95 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00774F03 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00774F75 | `NoContent_Screen` | Known | Screen layout |
| 0x00774F89 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00774FF3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077505C | `No_Photos_Screen` | Known | Screen layout |
| 0x00775070 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x007750D6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00775144 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007751B1 | `NoContent_Screen` | Known | Screen layout |
| 0x007751C5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077522D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00775297 | `NoContent_Screen` | Known | Screen layout |
| 0x007752AB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00775312 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077537C | `NoContent_Screen` | Known | Screen layout |
| 0x00775390 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007753FD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077546F | `NoContent_Screen` | Known | Screen layout |
| 0x00775483 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007754EB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00775554 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077556F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007755D5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007755F1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007756D0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007756E9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077574A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077575E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007757B8 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007757D4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077583B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00775852 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007759C3 | `Radio_Screen` | Known | Screen layout |
| 0x007759D3 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00775A34 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00775AB7 | `LockediPod_Screen` | Known | Screen layout |
| 0x00775B3F | `Lock_Screen` | Known | Screen layout |
| 0x00775B4E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00775BB1 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00775C13 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00775C2F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00775CA1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00775CC0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00775D28 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00775D42 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00775DAA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00775DC7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00775E33 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00775E9D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00775EB7 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00775F27 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00775F9A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077600B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077607A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x007760E6 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00776101 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00776176 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x007761DD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077623F | `Photos_Screen` | Known | Screen layout |
| 0x007762A3 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x007762C1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00776333 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00776350 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007763B6 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007763D1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077643A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00776457 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007764CE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007764F2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00776560 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077657B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00776638 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00776654 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007766C2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007766DF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077674A | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077676A | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007767E1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007767FD | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077686D | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077688C | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007768F8 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077690C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00776985 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x007769F9 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00776A69 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00776AD0 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00776B38 | `NoContent_Screen` | Known | Screen layout |
| 0x00776B4C | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00776BB0 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00776C17 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00776C31 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00776C9F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00776D11 | `NoContent_Screen` | Known | Screen layout |
| 0x00776D25 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00776D8F | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00776DF8 | `No_Photos_Screen` | Known | Screen layout |
| 0x00776E0C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00776E72 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00776EE0 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00776F4D | `NoContent_Screen` | Known | Screen layout |
| 0x00776F61 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00776FC9 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00777033 | `NoContent_Screen` | Known | Screen layout |
| 0x00777047 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007770AE | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00777118 | `NoContent_Screen` | Known | Screen layout |
| 0x0077712C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00777199 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077720B | `NoContent_Screen` | Known | Screen layout |
| 0x0077721F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00777287 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007772F0 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077730B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00777371 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077738D | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077746C | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00777485 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007774E6 | `FirstBoot_Screen` | Known | Screen layout |
| 0x007774FA | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00777554 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00777570 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007775D7 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007775EE | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077775F | `Radio_Screen` | Known | Screen layout |
| 0x0077776F | `Radio_Screen_Default ` | Known | Screen layout |
| 0x007777D0 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00777853 | `LockediPod_Screen` | Known | Screen layout |
| 0x007778DB | `Lock_Screen` | Known | Screen layout |
| 0x007778EA | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077794D | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007779AF | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007779CB | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00777A3D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00777A5C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00777AC4 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00777ADE | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00777B46 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00777B63 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00777BCF | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00777C39 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00777C53 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00777CC3 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00777D36 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00777DA7 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00777E16 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00777E82 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00777E9D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00777F12 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00777F79 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00777FDB | `Photos_Screen` | Known | Screen layout |
| 0x0077803F | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077805D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007780CF | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007780EC | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00778152 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077816D | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007781D6 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007781F3 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077826A | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077828E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007782FC | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00778317 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007783D4 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007783F0 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077845E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077847B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007784E6 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00778506 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077857D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00778599 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00778609 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00778628 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00778694 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007786A8 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00778721 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00778795 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00778805 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077886C | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x007788D4 | `NoContent_Screen` | Known | Screen layout |
| 0x007788E8 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077894C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x007789B3 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x007789CD | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00778A3B | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00778AAD | `NoContent_Screen` | Known | Screen layout |
| 0x00778AC1 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00778B2B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00778B94 | `No_Photos_Screen` | Known | Screen layout |
| 0x00778BA8 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00778C0E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00778C7C | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00778CE9 | `NoContent_Screen` | Known | Screen layout |
| 0x00778CFD | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00778D65 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00778DCF | `NoContent_Screen` | Known | Screen layout |
| 0x00778DE3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00778E4A | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00778EB4 | `NoContent_Screen` | Known | Screen layout |
| 0x00778EC8 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00778F35 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00778FA7 | `NoContent_Screen` | Known | Screen layout |
| 0x00778FBB | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00779023 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077908C | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007790A7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077910D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00779129 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00779208 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00779221 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00779282 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00779296 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007792F0 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077930C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00779373 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077938A | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007794FB | `Radio_Screen` | Known | Screen layout |
| 0x0077950B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077956C | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007795EF | `LockediPod_Screen` | Known | Screen layout |
| 0x00779677 | `Lock_Screen` | Known | Screen layout |
| 0x00779686 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x007796E9 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077974B | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00779767 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x007797D9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007797F8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00779860 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077987A | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x007798E2 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007798FF | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077996B | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x007799D5 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x007799EF | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00779A5F | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00779AD2 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00779B43 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00779BB2 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00779C1E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00779C39 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00779CAE | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00779D15 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00779D77 | `Photos_Screen` | Known | Screen layout |
| 0x00779DDB | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00779DF9 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00779E6B | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00779E88 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00779EEE | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00779F09 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00779F72 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00779F8F | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077A006 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077A02A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077A098 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077A0B3 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077A170 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077A18C | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077A1FA | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077A217 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077A282 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077A2A2 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077A319 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077A335 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077A3A5 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077A3C4 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077A430 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077A444 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077A4BD | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077A531 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077A5A1 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077A608 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077A670 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A684 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077A6E8 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077A74F | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077A769 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077A7D7 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077A849 | `NoContent_Screen` | Known | Screen layout |
| 0x0077A85D | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077A8C7 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077A930 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077A944 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077A9AA | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077AA18 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077AA85 | `NoContent_Screen` | Known | Screen layout |
| 0x0077AA99 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077AB01 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077AB6B | `NoContent_Screen` | Known | Screen layout |
| 0x0077AB7F | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077ABE6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077AC50 | `NoContent_Screen` | Known | Screen layout |
| 0x0077AC64 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077ACD1 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077AD43 | `NoContent_Screen` | Known | Screen layout |
| 0x0077AD57 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077ADBF | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077AE28 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077AE43 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077AEA9 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077AEC5 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077AFA4 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077AFBD | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077B01E | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077B032 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077B08C | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077B0A8 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077B10F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077B126 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077B297 | `Radio_Screen` | Known | Screen layout |
| 0x0077B2A7 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077B308 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077B38B | `LockediPod_Screen` | Known | Screen layout |
| 0x0077B413 | `Lock_Screen` | Known | Screen layout |
| 0x0077B422 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077B485 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077B4E7 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077B503 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077B575 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077B594 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077B5FC | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077B616 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077B67E | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077B69B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077B707 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077B771 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077B78B | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077B7FB | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077B86E | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077B8DF | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077B94E | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077B9BA | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077B9D5 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077BA4A | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077BAB1 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077BB13 | `Photos_Screen` | Known | Screen layout |
| 0x0077BB77 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077BB95 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077BC07 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077BC24 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077BC8A | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077BCA5 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077BD0E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077BD2B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077BDA2 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077BDC6 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077BE34 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077BE4F | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077BF0C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077BF28 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077BF96 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077BFB3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077C01E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077C03E | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077C0B5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077C0D1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077C141 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077C160 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077C1CC | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077C1E0 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077C259 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077C2CD | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077C33D | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077C3A4 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077C40C | `NoContent_Screen` | Known | Screen layout |
| 0x0077C420 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077C484 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077C4EB | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077C505 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077C573 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077C5E5 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C5F9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077C663 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077C6CC | `No_Photos_Screen` | Known | Screen layout |
| 0x0077C6E0 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077C746 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077C7B4 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077C821 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C835 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077C89D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077C907 | `NoContent_Screen` | Known | Screen layout |
| 0x0077C91B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077C982 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077C9EC | `NoContent_Screen` | Known | Screen layout |
| 0x0077CA00 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077CA6D | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077CADF | `NoContent_Screen` | Known | Screen layout |
| 0x0077CAF3 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077CB5B | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077CBC4 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077CBDF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077CC45 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077CC61 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077CD40 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077CD59 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077CDBA | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077CDCE | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077CE28 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077CE44 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077CEAB | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077CEC2 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077D033 | `Radio_Screen` | Known | Screen layout |
| 0x0077D043 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077D0A4 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077D127 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077D1AF | `Lock_Screen` | Known | Screen layout |
| 0x0077D1BE | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077D221 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077D283 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077D29F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077D311 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077D330 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077D398 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077D3B2 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077D41A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077D437 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077D4A3 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077D50D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077D527 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077D597 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077D60A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077D67B | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077D6EA | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077D756 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077D771 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077D7E6 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077D84D | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077D8AF | `Photos_Screen` | Known | Screen layout |
| 0x0077D913 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077D931 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077D9A3 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077D9C0 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077DA26 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077DA41 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077DAAA | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077DAC7 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077DB3E | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077DB62 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077DBD0 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077DBEB | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077DCA8 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077DCC4 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077DD32 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077DD4F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077DDBA | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077DDDA | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077DE51 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077DE6D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077DEDD | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077DEFC | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077DF68 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077DF7C | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077DFF5 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077E069 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077E0D9 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077E140 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077E1A8 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E1BC | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077E220 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0077E287 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077E2A1 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x0077E30F | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0077E381 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E395 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0077E3FF | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0077E468 | `No_Photos_Screen` | Known | Screen layout |
| 0x0077E47C | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0077E4E2 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077E550 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x0077E5BD | `NoContent_Screen` | Known | Screen layout |
| 0x0077E5D1 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0077E639 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0077E6A3 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E6B7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0077E71E | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0077E788 | `NoContent_Screen` | Known | Screen layout |
| 0x0077E79C | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x0077E809 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0077E87B | `NoContent_Screen` | Known | Screen layout |
| 0x0077E88F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0077E8F7 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0077E960 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0077E97B | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0077E9E1 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0077E9FD | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0077EADC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0077EAF5 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0077EB56 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0077EB6A | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x0077EBC4 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0077EBE0 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0077EC47 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0077EC5E | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0077EDCF | `Radio_Screen` | Known | Screen layout |
| 0x0077EDDF | `Radio_Screen_Default ` | Known | Screen layout |
| 0x0077EE40 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0077EEC3 | `LockediPod_Screen` | Known | Screen layout |
| 0x0077EF4B | `Lock_Screen` | Known | Screen layout |
| 0x0077EF5A | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x0077EFBD | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0077F01F | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0077F03B | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x0077F0AD | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0077F0CC | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0077F134 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0077F14E | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0077F1B6 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077F1D3 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077F23F | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0077F2A9 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0077F2C3 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x0077F333 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x0077F3A6 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x0077F417 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0077F486 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0077F4F2 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x0077F50D | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0077F582 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0077F5E9 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0077F64B | `Photos_Screen` | Known | Screen layout |
| 0x0077F6AF | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x0077F6CD | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x0077F73F | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x0077F75C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0077F7C2 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0077F7DD | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0077F846 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0077F863 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0077F8DA | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0077F8FE | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x0077F96C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0077F987 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0077FA44 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077FA60 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077FACE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0077FAEB | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0077FB56 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0077FB76 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0077FBED | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0077FC09 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0077FC79 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0077FC98 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0077FD04 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0077FD18 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x0077FD91 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0077FE05 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0077FE75 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x0077FEDC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0077FF44 | `NoContent_Screen` | Known | Screen layout |
| 0x0077FF58 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0077FFBC | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00780023 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0078003D | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x007800AB | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0078011D | `NoContent_Screen` | Known | Screen layout |
| 0x00780131 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0078019B | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00780204 | `No_Photos_Screen` | Known | Screen layout |
| 0x00780218 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078027E | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x007802EC | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00780359 | `NoContent_Screen` | Known | Screen layout |
| 0x0078036D | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x007803D5 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0078043F | `NoContent_Screen` | Known | Screen layout |
| 0x00780453 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007804BA | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00780524 | `NoContent_Screen` | Known | Screen layout |
| 0x00780538 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007805A5 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00780617 | `NoContent_Screen` | Known | Screen layout |
| 0x0078062B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00780693 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x007806FC | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x00780717 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0078077D | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00780799 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00780878 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x00780891 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x007808F2 | `FirstBoot_Screen` | Known | Screen layout |
| 0x00780906 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00780960 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078097C | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x007809E3 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007809FA | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00780B6B | `Radio_Screen` | Known | Screen layout |
| 0x00780B7B | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00780BDC | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00780C5F | `LockediPod_Screen` | Known | Screen layout |
| 0x00780CE7 | `Lock_Screen` | Known | Screen layout |
| 0x00780CF6 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00780D59 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00780DBB | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00780DD7 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00780E49 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00780E68 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00780ED0 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00780EEA | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00780F52 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00780F6F | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00780FDB | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00781045 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x0078105F | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x007810CF | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00781142 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x007811B3 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00781222 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078128E | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x007812A9 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x0078131E | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00781385 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x007813E7 | `Photos_Screen` | Known | Screen layout |
| 0x0078144B | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00781469 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x007814DB | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x007814F8 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0078155E | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00781579 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007815E2 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007815FF | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00781676 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x0078169A | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00781708 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00781723 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007817E0 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007817FC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078186A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00781887 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007818F2 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00781912 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00781989 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007819A5 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00781A15 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00781A34 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00781AA0 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00781AB4 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x00781B2D | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00781BA1 | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00781C11 | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00781C78 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00781CE0 | `NoContent_Screen` | Known | Screen layout |
| 0x00781CF4 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00781D58 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00781DBF | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00781DD9 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00781E47 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00781EB9 | `NoContent_Screen` | Known | Screen layout |
| 0x00781ECD | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00781F37 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00781FA0 | `No_Photos_Screen` | Known | Screen layout |
| 0x00781FB4 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x0078201A | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00782088 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x007820F5 | `NoContent_Screen` | Known | Screen layout |
| 0x00782109 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00782171 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x007821DB | `NoContent_Screen` | Known | Screen layout |
| 0x007821EF | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00782256 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x007822C0 | `NoContent_Screen` | Known | Screen layout |
| 0x007822D4 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00782341 | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007823B3 | `NoContent_Screen` | Known | Screen layout |
| 0x007823C7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0078242F | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00782498 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x007824B3 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00782519 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00782535 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00782614 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0078262D | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078268E | `FirstBoot_Screen` | Known | Screen layout |
| 0x007826A2 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x007826FC | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00782718 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078277F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00782796 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00782907 | `Radio_Screen` | Known | Screen layout |
| 0x00782917 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00782978 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x007829FB | `LockediPod_Screen` | Known | Screen layout |
| 0x00782A83 | `Lock_Screen` | Known | Screen layout |
| 0x00782A92 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00782AF5 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00782B57 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00782B73 | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00782BE5 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00782C04 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00782C6C | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00782C86 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00782CEE | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00782D0B | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00782D77 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00782DE1 | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00782DFB | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00782E6B | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00782EDE | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00782F4F | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00782FBE | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0078302A | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00783045 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x007830BA | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00783121 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00783183 | `Photos_Screen` | Known | Screen layout |
| 0x007831E7 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00783205 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00783277 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00783294 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x007832FA | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00783315 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078337E | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0078339B | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00783412 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00783436 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007834A4 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x007834BF | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0078357C | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00783598 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00783606 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00783623 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078368E | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007836AE | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00783725 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00783741 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007837B1 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007837D0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078383C | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00783850 | `CoverFlow_Screen_DefaultNoTextNoStatusBar!` | Known | Screen layout |
| 0x007838C9 | `MainMenus_Main_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078393D | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x007839AD | `MainMenu_Main_Screen_Genius` | Known | Screen layout |
| 0x00783A14 | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x00783A7C | `NoContent_Screen` | Known | Screen layout |
| 0x00783A90 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x00783AF4 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x00783B5B | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00783B75 | `MainMenu_Music_Screen_NoMusic"` | Known | Screen layout |
| 0x00783BE3 | `MainMenu_Main_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00783C55 | `NoContent_Screen` | Known | Screen layout |
| 0x00783C69 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00783CD3 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00783D3C | `No_Photos_Screen` | Known | Screen layout |
| 0x00783D50 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00783DB6 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00783E24 | `MainMenus_Main_Screen_NoPodcasts ` | Known | Screen layout |
| 0x00783E91 | `NoContent_Screen` | Known | Screen layout |
| 0x00783EA5 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00783F0D | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00783F77 | `NoContent_Screen` | Known | Screen layout |
| 0x00783F8B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00783FF2 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0078405C | `NoContent_Screen` | Known | Screen layout |
| 0x00784070 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007840DD | `MainMenu_Main_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0078414F | `NoContent_Screen` | Known | Screen layout |
| 0x00784163 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007841CB | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00784234 | `MainMenus_Videos_Screen ` | Known | Screen layout |
| 0x0078424F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007842B5 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007842D1 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x007843B0 | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x007843C9 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0078442A | `FirstBoot_Screen` | Known | Screen layout |
| 0x0078443E | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x00784498 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007844B4 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078451B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00784532 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007846A3 | `Radio_Screen` | Known | Screen layout |
| 0x007846B3 | `Radio_Screen_Default ` | Known | Screen layout |
| 0x00784714 | `LockConfirmation_Screen` | Known | Screen layout |
| 0x00784797 | `LockediPod_Screen` | Known | Screen layout |
| 0x0078481F | `Lock_Screen` | Known | Screen layout |
| 0x0078482E | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00784891 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x007848F3 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x0078490F | `MediaLists_Movies_Screen_WithTime$` | Known | Screen layout |
| 0x00784981 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007849A0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00784A08 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00784A22 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x00784A8A | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00784AA7 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00784B13 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x00784B7D | `MainMenus_Music_Screen ` | Known | Screen layout |
| 0x00784B97 | `MainMenu_Music_Screen_NoMusicArt!` | Known | Screen layout |
| 0x00784C07 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00784C7A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x00784CEB | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x00784D5A | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00784DC6 | `MainMenus_Videos_Screen#` | Known | Screen layout |
| 0x00784DE1 | `MainMenus_Videos_Screen_NoVideosArt#` | Known | Screen layout |
| 0x00784E56 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00784EBD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00784F1F | `Photos_Screen` | Known | Screen layout |
| 0x00784F83 | `MediaLists_Podcasts_Screen+` | Known | Screen layout |
| 0x00784FA1 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00785013 | `MediaLists_Rentals_Screen!` | Known | Screen layout |
| 0x00785030 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00785096 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007850B1 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078511A | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00785137 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007851AE | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007851D2 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00785240 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0078525B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007852FD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00785319 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00785387 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007853A4 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078540F | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078542F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x007854A6 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007854C2 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00785532 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00785551 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007855BD | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007855D1 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00785646 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007856B1 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00785720 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00785791 | `NoContent_Screen` | Known | Screen layout |
| 0x007857A5 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00785814 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00785887 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007858F4 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078595D | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007859CD | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00785A3D | `NoContent_Screen` | Known | Screen layout |
| 0x00785A51 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00785AB4 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00785B17 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00785B33 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00785B95 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00785BB1 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00785C18 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00785C2F | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00785CEA | `Radio_Screen` | Known | Screen layout |
| 0x00785CFA | `Radio_Screen_Default` | Known | Screen layout |
| 0x00785D5B | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00785DC9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00785DE8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00785E56 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00785EBB | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00785ED6 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00785F79 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00785F95 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00786003 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00786020 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078608B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007860AB | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00786122 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078613E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007861AE | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007861CD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00786239 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078624D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007862C2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078632D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078639C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078640D | `NoContent_Screen` | Known | Screen layout |
| 0x00786421 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00786490 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00786503 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00786570 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007865D9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00786649 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007866B9 | `NoContent_Screen` | Known | Screen layout |
| 0x007866CD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00786730 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00786793 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007867AF | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00786811 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078682D | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00786894 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007868AB | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00786966 | `Radio_Screen` | Known | Screen layout |
| 0x00786976 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007869D7 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00786A45 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00786A64 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00786AD2 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00786B37 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00786B52 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00786BF5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00786C11 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00786C7F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00786C9C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00786D07 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00786D27 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00786D9E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00786DBA | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00786E2A | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00786E49 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00786EB5 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00786EC9 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00786F3E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00786FA9 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00787018 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00787089 | `NoContent_Screen` | Known | Screen layout |
| 0x0078709D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078710C | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078717F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x007871EC | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00787255 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x007872C5 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00787335 | `NoContent_Screen` | Known | Screen layout |
| 0x00787349 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x007873AC | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078740F | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078742B | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078748D | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007874A9 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00787510 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00787527 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x007875E2 | `Radio_Screen` | Known | Screen layout |
| 0x007875F2 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00787653 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x007876C1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007876E0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078774E | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007877B3 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007877CE | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00787871 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078788D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007878FB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00787918 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00787983 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x007879A3 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00787A1A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00787A36 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00787AA6 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00787AC5 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00787B31 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x00787B45 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00787BBA | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x00787C25 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00787C94 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00787D05 | `NoContent_Screen` | Known | Screen layout |
| 0x00787D19 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00787D88 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00787DFB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00787E68 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00787ED1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00787F41 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00787FB1 | `NoContent_Screen` | Known | Screen layout |
| 0x00787FC5 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00788028 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078808B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x007880A7 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00788109 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00788125 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078818C | `Genius_Intro_Screen` | Known | Screen layout |
| 0x007881A3 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078825E | `Radio_Screen` | Known | Screen layout |
| 0x0078826E | `Radio_Screen_Default` | Known | Screen layout |
| 0x007882CF | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078833D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078835C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x007883CA | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078842F | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078844A | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x007884ED | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00788509 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00788577 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00788594 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007885FF | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078861F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00788696 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007886B2 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00788722 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x00788741 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x007887AD | `CoverFlow_Screen)` | Known | Screen layout |
| 0x007887C1 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x00788836 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x007888A1 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x00788910 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x00788981 | `NoContent_Screen` | Known | Screen layout |
| 0x00788995 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00788A04 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x00788A77 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00788AE4 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x00788B4D | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00788BBD | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00788C2D | `NoContent_Screen` | Known | Screen layout |
| 0x00788C41 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00788CA4 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00788D07 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x00788D23 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00788D85 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00788DA1 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00788E08 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00788E1F | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00788EDA | `Radio_Screen` | Known | Screen layout |
| 0x00788EEA | `Radio_Screen_Default` | Known | Screen layout |
| 0x00788F4B | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00788FB9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00788FD8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00789046 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x007890AB | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007890C6 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00789169 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00789185 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x007891F3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00789210 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078927B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078929B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00789312 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078932E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078939E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x007893BD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00789429 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078943D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x007894B2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078951D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078958C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x007895FD | `NoContent_Screen` | Known | Screen layout |
| 0x00789611 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x00789680 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x007896F3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x00789760 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x007897C9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x00789839 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x007898A9 | `NoContent_Screen` | Known | Screen layout |
| 0x007898BD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x00789920 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x00789983 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078999F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00789A01 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00789A1D | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00789A84 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00789A9B | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00789B56 | `Radio_Screen` | Known | Screen layout |
| 0x00789B66 | `Radio_Screen_Default` | Known | Screen layout |
| 0x00789BC7 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x00789C35 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00789C54 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00789CC2 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00789D27 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x00789D42 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00789DE5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00789E01 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00789E6F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00789E8C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00789EF7 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x00789F17 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00789F8E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00789FAA | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078A01A | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078A039 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078A0A5 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078A0B9 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078A12E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078A199 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078A208 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078A279 | `NoContent_Screen` | Known | Screen layout |
| 0x0078A28D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078A2FC | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078A36F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078A3DC | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078A445 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078A4B5 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078A525 | `NoContent_Screen` | Known | Screen layout |
| 0x0078A539 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078A59C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078A5FF | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078A61B | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078A67D | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078A699 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078A700 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078A717 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078A7D2 | `Radio_Screen` | Known | Screen layout |
| 0x0078A7E2 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078A843 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078A8B1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078A8D0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078A93E | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078A9A3 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078A9BE | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078AA61 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078AA7D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078AAEB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078AB08 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078AB73 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078AB93 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078AC0A | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078AC26 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078AC96 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078ACB5 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078AD21 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078AD35 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078ADAA | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078AE15 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078AE84 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078AEF5 | `NoContent_Screen` | Known | Screen layout |
| 0x0078AF09 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078AF78 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078AFEB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078B058 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078B0C1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078B131 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078B1A1 | `NoContent_Screen` | Known | Screen layout |
| 0x0078B1B5 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078B218 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078B27B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078B297 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078B2F9 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078B315 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078B37C | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078B393 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078B44E | `Radio_Screen` | Known | Screen layout |
| 0x0078B45E | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078B4BF | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078B52D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078B54C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078B5BA | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078B61F | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078B63A | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078B6DD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078B6F9 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078B767 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078B784 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078B7EF | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078B80F | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078B886 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078B8A2 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078B912 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078B931 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078B99D | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078B9B1 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078BA26 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078BA91 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078BB00 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078BB71 | `NoContent_Screen` | Known | Screen layout |
| 0x0078BB85 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078BBF4 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078BC67 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078BCD4 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078BD3D | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078BDAD | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078BE1D | `NoContent_Screen` | Known | Screen layout |
| 0x0078BE31 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078BE94 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078BEF7 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078BF13 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078BF75 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078BF91 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078BFF8 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078C00F | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078C0CA | `Radio_Screen` | Known | Screen layout |
| 0x0078C0DA | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078C13B | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078C1A9 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078C1C8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078C236 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078C29B | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078C2B6 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078C359 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078C375 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078C3E3 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078C400 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078C46B | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078C48B | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078C502 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078C51E | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078C58E | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078C5AD | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078C619 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078C62D | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078C6A2 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078C70D | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078C77C | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078C7ED | `NoContent_Screen` | Known | Screen layout |
| 0x0078C801 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078C870 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078C8E3 | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078C950 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078C9B9 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078CA29 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078CA99 | `NoContent_Screen` | Known | Screen layout |
| 0x0078CAAD | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078CB10 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078CB73 | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078CB8F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078CBF1 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078CC0D | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078CC74 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078CC8B | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078CD46 | `Radio_Screen` | Known | Screen layout |
| 0x0078CD56 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078CDB7 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078CE25 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078CE44 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078CEB2 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078CF17 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078CF32 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078CFD5 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078CFF1 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078D05F | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078D07C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078D0E7 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078D107 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078D17E | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078D19A | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078D20A | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078D229 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078D295 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078D2A9 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078D31E | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078D389 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078D3F8 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078D469 | `NoContent_Screen` | Known | Screen layout |
| 0x0078D47D | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078D4EC | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078D55F | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078D5CC | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078D635 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078D6A5 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078D715 | `NoContent_Screen` | Known | Screen layout |
| 0x0078D729 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078D78C | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078D7EF | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078D80B | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078D86D | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078D889 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078D8F0 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078D907 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078D9C2 | `Radio_Screen` | Known | Screen layout |
| 0x0078D9D2 | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078DA33 | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078DAA1 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078DAC0 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078DB2E | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078DB93 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078DBAE | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078DC51 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078DC6D | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078DCDB | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0078DCF8 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0078DD63 | `MediaLists_Audiobooks_Screen+` | Known | Screen layout |
| 0x0078DD83 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0078DDFA | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x0078DE16 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0078DE86 | `MediaLists_Composers_Screen#` | Known | Screen layout |
| 0x0078DEA5 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0078DF11 | `CoverFlow_Screen)` | Known | Screen layout |
| 0x0078DF25 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0078DF9A | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0078E005 | `MainMenu_Music_Screen_NoArtists!` | Known | Screen layout |
| 0x0078E074 | `MainMenu_Music_Screen_NoAudiobooks"` | Known | Screen layout |
| 0x0078E0E5 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E0F9 | `NoContent_Screen_NoAudiobooks#` | Known | Screen layout |
| 0x0078E168 | `MainMenu_Music_Screen_NoCompilations ` | Known | Screen layout |
| 0x0078E1DB | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0078E248 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0078E2B1 | `MainMenu_Music_Screen_NoMusic$` | Known | Screen layout |
| 0x0078E321 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x0078E391 | `NoContent_Screen` | Known | Screen layout |
| 0x0078E3A5 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0078E408 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0078E46B | `MediaLists_Genres_Screen ` | Known | Screen layout |
| 0x0078E487 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0078E4E9 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078E505 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x0078E56C | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0078E583 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0078E63E | `Radio_Screen` | Known | Screen layout |
| 0x0078E64E | `Radio_Screen_Default` | Known | Screen layout |
| 0x0078E6AF | `MainMenu_Music_Screen_Default$` | Known | Screen layout |
| 0x0078E71D | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0078E73C | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x0078E7AA | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x0078E80F | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078E82A | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0078E94A | `MediaLists_AudiobookChapters_Screen)` | Known | Screen layout |
| 0x0078E971 | `MediaLists_AudiobookChapters_Screen_Plane!` | Known | Screen layout |
| 0x0078EEDE | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078EEFA | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0078EF69 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0078EF82 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0078F2EA | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078F306 | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0078F375 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0078F38E | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0078F6B7 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x0078F6D3 | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0078F742 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0078F75B | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0078F98B | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0078F9A6 | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x0078FA11 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078FA2C | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x0078FA9F | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0078FABA | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0078FCE3 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0078FCFE | `MediaLists_Songs_Screen_Plain!` | Known | Screen layout |
| 0x0078FD69 | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x0078FD84 | `MediaLists_Songs_Screen_WithArtist$` | Known | Screen layout |
| 0x0078FDF7 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x0078FE12 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x00790046 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00790062 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x007900DD | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007900F9 | `MediaLists_Albums_Screen_WithTrackCount%` | Known | Screen layout |
| 0x00790172 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0079018D | `MediaLists_Songs_Screen_Plain1` | Known | Screen layout |
| 0x00790208 | `MediaLists_Songs_Screen,` | Known | Screen layout |
| 0x00790223 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x007904B1 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007904CE | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00790615 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x00790631 | `MediaLists_Albums_Screen_WithArtistName'` | Known | Screen layout |
| 0x007906AC | `MediaLists_Songs_Screen"` | Known | Screen layout |
| 0x007906C7 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x00790915 | `MediaLists_PodcastEpisodes_Screen'` | Known | Screen layout |
| 0x0079093A | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00790C72 | `MediaLists_TVSeasons_Screen#` | Known | Screen layout |
| 0x00790C91 | `MediaLists_TVSeasons_Screen_Default%` | Known | Screen layout |
| 0x00790D06 | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x00790D26 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00790EAE | `MediaLists_TVEpisodes_Screen"` | Known | Screen layout |
| 0x00790ECE | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x007912D8 | `MediaLists_GeniusPlaylist_Screen(` | Known | Screen layout |
| 0x007912FC | `MediaLists_GeniusPlaylist_Screen_Default` | Known | Screen layout |
| 0x007913CB | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x007913F0 | `MediaLists_NestedPlaylists_Screen_Default,` | Known | Screen layout |
| 0x00791472 | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x00791491 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00791770 | `Genius_Error_Screen` | Known | Screen layout |
| 0x00791787 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007917FF | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00791816 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x00791884 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007918A0 | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0079190F | `Genius_Loading_Screen` | Known | Screen layout |
| 0x00791928 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007919F2 | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x00791A17 | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x00791A8F | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x00791AAE | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00791B15 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x00791D64 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x00791DD6 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x00791E41 | `Genius_Error_Screen` | Known | Screen layout |
| 0x00791E58 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x00791ED0 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00791EE7 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x00791F55 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x00791F71 | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x00791FE0 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x00791FF9 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007920F2 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x007923B4 | `MediaLists_MixedVideoTracks_Screen9` | Known | Screen layout |
| 0x007924B4 | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x00792520 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079258A | `NoContent_Screen` | Known | Screen layout |
| 0x0079259E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00792608 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079267C | `NoContent_Screen` | Known | Screen layout |
| 0x00792690 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007926FB | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00792767 | `NoContent_Screen` | Known | Screen layout |
| 0x0079277B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007927E2 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0079284E | `NoContent_Screen` | Known | Screen layout |
| 0x00792862 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007928CF | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00792943 | `NoContent_Screen` | Known | Screen layout |
| 0x00792957 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007929BF | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00792A2C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00792A90 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00792AAC | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00792B18 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00792B35 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00792BA2 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00792C69 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00792C86 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00792CFD | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00792D21 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00792DD8 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00792E42 | `NoContent_Screen` | Known | Screen layout |
| 0x00792E56 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00792EC0 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00792F34 | `NoContent_Screen` | Known | Screen layout |
| 0x00792F48 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00792FB3 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0079301F | `NoContent_Screen` | Known | Screen layout |
| 0x00793033 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079309A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00793106 | `NoContent_Screen` | Known | Screen layout |
| 0x0079311A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00793187 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007931FB | `NoContent_Screen` | Known | Screen layout |
| 0x0079320F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00793277 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007932E4 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00793348 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00793364 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007933D0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007933ED | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079345A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00793521 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079353E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x007935B5 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007935D9 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00793690 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007936FA | `NoContent_Screen` | Known | Screen layout |
| 0x0079370E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00793778 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007937EC | `NoContent_Screen` | Known | Screen layout |
| 0x00793800 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0079386B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007938D7 | `NoContent_Screen` | Known | Screen layout |
| 0x007938EB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00793952 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007939BE | `NoContent_Screen` | Known | Screen layout |
| 0x007939D2 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00793A3F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00793AB3 | `NoContent_Screen` | Known | Screen layout |
| 0x00793AC7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00793B2F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00793B9C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00793C00 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00793C1C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00793C88 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00793CA5 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00793D12 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00793DD9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00793DF6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00793E6D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00793E91 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00793F48 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00793FB2 | `NoContent_Screen` | Known | Screen layout |
| 0x00793FC6 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00794030 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x007940A4 | `NoContent_Screen` | Known | Screen layout |
| 0x007940B8 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00794123 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0079418F | `NoContent_Screen` | Known | Screen layout |
| 0x007941A3 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079420A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00794276 | `NoContent_Screen` | Known | Screen layout |
| 0x0079428A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007942F7 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079436B | `NoContent_Screen` | Known | Screen layout |
| 0x0079437F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007943E7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00794454 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x007944B8 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007944D4 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00794540 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079455D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007945CA | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00794691 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007946AE | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00794725 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00794749 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00794800 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0079486A | `NoContent_Screen` | Known | Screen layout |
| 0x0079487E | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007948E8 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x0079495C | `NoContent_Screen` | Known | Screen layout |
| 0x00794970 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x007949DB | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00794A47 | `NoContent_Screen` | Known | Screen layout |
| 0x00794A5B | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00794AC2 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00794B2E | `NoContent_Screen` | Known | Screen layout |
| 0x00794B42 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00794BAF | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00794C23 | `NoContent_Screen` | Known | Screen layout |
| 0x00794C37 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00794C9F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00794D0C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00794D70 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00794D8C | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00794DF8 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00794E15 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00794E82 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00794F49 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x00794F66 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00794FDD | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00795001 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x007950B8 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00795122 | `NoContent_Screen` | Known | Screen layout |
| 0x00795136 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x007951A0 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00795214 | `NoContent_Screen` | Known | Screen layout |
| 0x00795228 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00795293 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x007952FF | `NoContent_Screen` | Known | Screen layout |
| 0x00795313 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0079537A | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x007953E6 | `NoContent_Screen` | Known | Screen layout |
| 0x007953FA | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00795467 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x007954DB | `NoContent_Screen` | Known | Screen layout |
| 0x007954EF | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00795557 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x007955C4 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00795628 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00795644 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x007956B0 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x007956CD | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x0079573A | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00795801 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079581E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00795895 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x007958B9 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00795970 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x007959DA | `NoContent_Screen` | Known | Screen layout |
| 0x007959EE | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00795A58 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00795ACC | `NoContent_Screen` | Known | Screen layout |
| 0x00795AE0 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00795B4B | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x00795BB7 | `NoContent_Screen` | Known | Screen layout |
| 0x00795BCB | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00795C32 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00795C9E | `NoContent_Screen` | Known | Screen layout |
| 0x00795CB2 | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x00795D1F | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x00795D93 | `NoContent_Screen` | Known | Screen layout |
| 0x00795DA7 | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x00795E0F | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00795E7C | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00795EE0 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x00795EFC | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00795F68 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x00795F85 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00795FF2 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x007960B9 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x007960D6 | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x0079614D | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00796171 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00796228 | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x00796292 | `NoContent_Screen` | Known | Screen layout |
| 0x007962A6 | `NoContent_Screen_NoMovies"` | Known | Screen layout |
| 0x00796310 | `MainMenu_Videos_Screen_NoMusicVideos#` | Known | Screen layout |
| 0x00796384 | `NoContent_Screen` | Known | Screen layout |
| 0x00796398 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00796403 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0079646F | `NoContent_Screen` | Known | Screen layout |
| 0x00796483 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x007964EA | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x00796556 | `NoContent_Screen` | Known | Screen layout |
| 0x0079656A | `NoContent_Screen_NoTVShows$` | Known | Screen layout |
| 0x007965D7 | `MainMenu_Videos_Screen_NoPlaylists%` | Known | Screen layout |
| 0x0079664B | `NoContent_Screen` | Known | Screen layout |
| 0x0079665F | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x007966C7 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00796734 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x00796798 | `MediaLists_Movies_Screen!` | Known | Screen layout |
| 0x007967B4 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00796820 | `MediaLists_Artists_Screen!` | Known | Screen layout |
| 0x0079683D | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x007968AA | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00796971 | `MediaLists_TVShows_Screen*` | Known | Screen layout |
| 0x0079698E | `MediaLists_TVShows_Screen_WithEpisodeCount ` | Known | Screen layout |
| 0x00796A05 | `MediaLists_VideoPlaylists_Screen(` | Known | Screen layout |
| 0x00796A29 | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00796ED8 | `Genius_Error_Screen` | Known | Screen layout |
| 0x00796EEF | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x00796F67 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00796F7E | `Genius_Error_Screen_NoGeniusInfoForTrack"` | Known | Screen layout |
| 0x00796FF5 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079700E | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079717B | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x0079719A | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x00797464 | `MediaLists_Playlists_Screen#` | Known | Screen layout |
| 0x007974CF | `Genius_Error_Screen` | Known | Screen layout |
| 0x007974E6 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079755E | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00797575 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007975E3 | `MediaLists_Genius_Screen ` | Known | Screen layout |
| 0x007975FF | `MediaLists_Genius_Screen_Default"` | Known | Screen layout |
| 0x0079766E | `Genius_Loading_Screen` | Known | Screen layout |
| 0x00797687 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x00797751 | `MediaLists_NestedPlaylists_Screen)` | Known | Screen layout |
| 0x00797776 | `MediaLists_NestedPlaylists_Screen_Default"` | Known | Screen layout |
| 0x007977EE | `OnTheGo_Instructions_Screen#` | Known | Screen layout |
| 0x0079780D | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x00797D93 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00797E05 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00797E70 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00797ED5 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00797F3F | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00797FA9 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00798019 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x00798090 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x00798102 | `Genius_Error_Screen` | Known | Screen layout |
| 0x00798119 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x00798191 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007981A8 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079821A | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x00798281 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079829A | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x00798303 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079836E | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007983D8 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079843F | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007984AE | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079851C | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x00798581 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007985E9 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x00798654 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007986BF | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x00798726 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x00798D7F | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00798DF1 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00798E5C | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00798EC1 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00798F2B | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00798F95 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00799005 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079907C | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007990EE | `Genius_Error_Screen` | Known | Screen layout |
| 0x00799105 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079917D | `Genius_Error_Screen(` | Known | Screen layout |
| 0x00799194 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x00799206 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079926D | `Genius_Loading_Screen` | Known | Screen layout |
| 0x00799286 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007992EF | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079935A | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007993C4 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079942B | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079949A | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x00799508 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079956D | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007995D5 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x00799640 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007996AB | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x00799712 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x00799D69 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x00799DDB | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x00799E46 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x00799EAB | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x00799F15 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00799F7F | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x00799FEF | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079A066 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079A0D8 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079A0EF | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079A167 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079A17E | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079A1F0 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079A257 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079A270 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079A2D9 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079A344 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079A3AE | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079A415 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079A484 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079A4F2 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079A557 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079A5BF | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079A62A | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079A695 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079A6FC | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079AD51 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079ADC3 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079AE2E | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079AE93 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079AEFD | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079AF67 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079AFD7 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079B04E | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079B0C0 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079B0D7 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079B14F | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079B166 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079B1D8 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079B23F | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079B258 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079B2C1 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079B32C | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079B396 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079B3FD | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079B46C | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079B4DA | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079B53F | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079B5A7 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079B612 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079B67D | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079B6E4 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079BD21 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079BD93 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079BDFE | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079BE63 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079BECD | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079BF37 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079BFA7 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079C01E | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079C090 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079C0A7 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079C11F | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079C136 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079C1A8 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079C20F | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079C228 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079C291 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079C2FC | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079C366 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079C3CD | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079C43C | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079C4AA | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079C50F | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079C577 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079C5E2 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079C64D | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079C6B4 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079CCF1 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079CD63 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079CDCE | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079CE33 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079CE9D | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079CF07 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079CF77 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079CFEE | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079D060 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079D077 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079D0EF | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079D106 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079D178 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079D1DF | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079D1F8 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079D261 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079D2CC | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079D336 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079D39D | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079D40C | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079D47A | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079D4DF | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079D547 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079D5B2 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079D61D | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079D684 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079DCFE | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079DD70 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079DDDB | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079DE40 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079DEAA | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079DF14 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079DF84 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079DFFB | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079E06D | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079E084 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079E0FC | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079E113 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079E185 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079E1EC | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079E205 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079E26E | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079E2D9 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079E343 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079E3AA | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079E419 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079E487 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079E4EC | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079E554 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079E5BF | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079E62A | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079E691 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079ECF0 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079ED62 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079EDCD | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079EE32 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079EE9C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079EF06 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079EF76 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079EFED | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x0079F05F | `Genius_Error_Screen` | Known | Screen layout |
| 0x0079F076 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x0079F0EE | `Genius_Error_Screen(` | Known | Screen layout |
| 0x0079F105 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0079F177 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x0079F1DE | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0079F1F7 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0079F260 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x0079F2CB | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0079F335 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x0079F39C | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x0079F40B | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0079F479 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0079F4DE | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0079F546 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x0079F5B1 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0079F61C | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x0079F683 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0079FCCC | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x0079FD3E | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0079FDA9 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x0079FE0E | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x0079FE78 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0079FEE2 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x0079FF52 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x0079FFC9 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A003B | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A0052 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A00CA | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A00E1 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A0153 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A01BA | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A01D3 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A023C | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A02A7 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A0311 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A0378 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A03E7 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A0455 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A04BA | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A0522 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A058D | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A05F8 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A065F | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A0CA8 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A0D1A | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A0D85 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A0DEA | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A0E54 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A0EBE | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A0F2E | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A0FA5 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A1017 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A102E | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A10A6 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A10BD | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A112F | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A1196 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A11AF | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A1218 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A1283 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A12ED | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A1354 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A13C3 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A1431 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A1496 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A14FE | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A1569 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A15D4 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A163B | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A1C85 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A1CF7 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A1D62 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A1DC7 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A1E31 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A1E9B | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A1F0B | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A1F82 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A1FF4 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A200B | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A2083 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A209A | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A210C | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A2173 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A218C | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A21F5 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A2260 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A22CA | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A2331 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A23A0 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A240E | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A2473 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A24DB | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A2546 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A25B1 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A2618 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A2C87 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A2CF9 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A2D64 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A2DC9 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A2E33 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A2E9D | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A2F0D | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A2F84 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A2FF6 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A300D | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A3085 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A309C | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A310E | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A3175 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A318E | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A31F7 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A3262 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A32CC | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A3333 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A33A2 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A3410 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A3475 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A34DD | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A3548 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A35B3 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A361A | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A3C97 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A3D09 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A3D74 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A3DD9 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A3E43 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A3EAD | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A3F1D | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A3F94 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A4006 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A401D | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A4095 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A40AC | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A411E | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A4185 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A419E | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A4207 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A4272 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A42DC | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A4343 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A43B2 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A4420 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A4485 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A44ED | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A4558 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A45C3 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A462A | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A4C87 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A4CF9 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A4D64 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A4DC9 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A4E33 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A4E9D | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A4F0D | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A4F84 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A4FF6 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A500D | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A5085 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A509C | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A510E | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A5175 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A518E | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A51F7 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A5262 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A52CC | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A5333 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A53A2 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A5410 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A5475 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A54DD | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A5548 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A55B3 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A561A | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A5C6B | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A5CDD | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A5D48 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A5DAD | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A5E17 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A5E81 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A5EF1 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A5F68 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A5FDA | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A5FF1 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A6069 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A6080 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A60F2 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A6159 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A6172 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A61DB | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A6246 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A62B0 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A6317 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A6386 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A63F4 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A6459 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A64C1 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A652C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A6597 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A65FE | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A6C3D | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A6CAF | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A6D1A | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A6D7F | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A6DE9 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A6E53 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A6EC3 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A6F3A | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A6FAC | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A6FC3 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A703B | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A7052 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A70C4 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A712B | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A7144 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A71AD | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A7218 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A7282 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A72E9 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A7358 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A73C6 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A742B | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A7493 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A74FE | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A7569 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A75D0 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A7C06 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A7C78 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A7CE3 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A7D48 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A7DB2 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A7E1C | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A7E8C | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A7F03 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A7F75 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A7F8C | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A8004 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A801B | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A808D | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A80F4 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A810D | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A8176 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A81E1 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A824B | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A82B2 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A8321 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A838F | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A83F4 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A845C | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A84C7 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A8532 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A8599 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A8BEA | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A8C5C | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A8CC7 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A8D2C | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A8D96 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A8E00 | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A8E70 | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A8EE7 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A8F59 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A8F70 | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A8FE8 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A8FFF | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007A9071 | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007A90D8 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007A90F1 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007A915A | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007A91C5 | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007A922F | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007A9296 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007A9305 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007A9373 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007A93D8 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007A9440 | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007A94AB | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007A9516 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007A957D | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007A9B84 | `NowPlaying_Screen_Video_Brightness!` | Known | Screen layout |
| 0x007A9BF6 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x007A9C61 | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x007A9CC6 | `NowPlaying_Screen_Progress!` | Known | Screen layout |
| 0x007A9D30 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x007A9D9A | `NowPlaying_Screen_ExtraInfo(` | Known | Screen layout |
| 0x007A9E0A | `NowPlaying_Screen_ExtraInfoLoadFailed%` | Known | Screen layout |
| 0x007A9E81 | `NowPlaying_Screen_ExtraInfoLoading#` | Known | Screen layout |
| 0x007A9EF3 | `Genius_Error_Screen` | Known | Screen layout |
| 0x007A9F0A | `Genius_Error_Screen_NoGenius/` | Known | Screen layout |
| 0x007A9F82 | `Genius_Error_Screen(` | Known | Screen layout |
| 0x007A9F99 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x007AA00B | `NowPlaying_Screen_Genius"` | Known | Screen layout |
| 0x007AA072 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x007AA08B | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x007AA0F4 | `NowPlaying_Screen_Progress"` | Known | Screen layout |
| 0x007AA15F | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x007AA1C9 | `NowPlaying_Screen_Rating ` | Known | Screen layout |
| 0x007AA230 | `NowPlaying_Screen_Video_Rating"` | Known | Screen layout |
| 0x007AA29F | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x007AA30D | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x007AA372 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x007AA3DA | `NowPlaying_Screen_Shuffle#` | Known | Screen layout |
| 0x007AA445 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x007AA4B0 | `NowPlaying_Screen_Volume ` | Known | Screen layout |
| 0x007AA517 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x007AA86A | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AA8E1 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AA95E | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AA9D0 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AAA40 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AAAB6 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AAB24 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AAB91 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AAED6 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AAF4D | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AAFCA | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AB03C | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AB0AC | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AB122 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AB190 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AB1FD | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AB566 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AB5DD | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AB65A | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AB6CC | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AB73C | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AB7B2 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AB820 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AB88D | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007ABBF6 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007ABC6D | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007ABCE8 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007ABD58 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007ABDCE | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007ABE3C | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007ABEA9 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AC1E2 | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AC259 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AC2D4 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AC344 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AC3BA | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007AC428 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AC495 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AC7CC | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007AC843 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007AC8BE | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007AC92E | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007AC9A4 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007ACA12 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007ACA7F | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007ACD8F | `NowPlaying_Screen_Video_TVOut_Default%` | Known | Screen layout |
| 0x007ACE06 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x007ACE81 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x007ACEF1 | `NowPlaying_Screen_Video_TVOut_Default"` | Known | Screen layout |
| 0x007ACF67 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x007ACFD5 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x007AD042 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x007AD646 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007AD663 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007AD6DE | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007AD6F7 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007AD76F | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007AD788 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007AD7FD | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AD813 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007AD88A | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AD8A0 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007AD917 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007AD934 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007AD9AC | `Notes_List_Screen` | Known | Screen layout |
| 0x007AD9C1 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007ADB72 | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007ADB8F | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007ADC0A | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007ADC23 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007ADC9B | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007ADCB4 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007ADD29 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007ADD3F | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007ADDB6 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007ADDCC | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007ADE43 | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007ADE60 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007ADED8 | `Notes_List_Screen` | Known | Screen layout |
| 0x007ADEED | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007AE0CE | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007AE0EB | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007AE166 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007AE17F | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007AE1F7 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007AE210 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007AE285 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AE29B | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007AE312 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AE328 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007AE39F | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007AE3BC | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007AE434 | `Notes_List_Screen` | Known | Screen layout |
| 0x007AE449 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007AE5FE | `Notes_Contents_Screen_Alt!` | Known | Screen layout |
| 0x007AE61B | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x007AE696 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x007AE6AF | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x007AE727 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x007AE740 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x007AE7B5 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AE7CB | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x007AE842 | `Notes_Image_Screen` | Known | Screen layout |
| 0x007AE858 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x007AE8CF | `Notes_Instructions_Screen!` | Known | Screen layout |
| 0x007AE8EC | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x007AE964 | `Notes_List_Screen` | Known | Screen layout |
| 0x007AE979 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x007AEC91 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007AED37 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AEDBA | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x007AEE72 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x007AEEF4 | `PhotosSettingsSlideshowMusic_Screen+` | Known | Screen layout |
| 0x007AEF1B | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x007AF001 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x007AF1B9 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AF219 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AF276 | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007AF29D | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007AF33D | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AF39D | `PhotosSettings_Screen` | Known | Screen layout |
| 0x007AF3FA | `PhotosSettingsSlideshowMusic_Screen2` | Known | Screen layout |
| 0x007AF421 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x007AF6BC | `Photos_Screen` | Known | Screen layout |
| 0x007AF808 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AF86C | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007AF8CD | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007AF92A | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007AF987 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007AF9F5 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007AFA52 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007AFBF8 | `Photos_Screen` | Known | Screen layout |
| 0x007AFD44 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007AFDA8 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007AFE09 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007AFE66 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007AFEC3 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007AFF31 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007AFF8E | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B0134 | `Photos_Screen` | Known | Screen layout |
| 0x007B0280 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B02E4 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007B0345 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007B03A2 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007B03FF | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B046D | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B04CA | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B0670 | `Photos_Screen` | Known | Screen layout |
| 0x007B07BC | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B0820 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007B0881 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007B08DE | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007B093B | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B09A9 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B0A06 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B0BAC | `Photos_Screen` | Known | Screen layout |
| 0x007B0CF8 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B0D5C | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007B0DBD | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007B0E1A | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007B0E77 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B0EE5 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B0F42 | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B10E8 | `Photos_Screen` | Known | Screen layout |
| 0x007B1234 | `Slideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B1298 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x007B12F9 | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x007B1356 | `Slideshow_Screen_Paused` | Known | Screen layout |
| 0x007B13B3 | `Slideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B1421 | `Slideshow_Screen_Playing` | Known | Screen layout |
| 0x007B147E | `Slideshow_Screen_Volume` | Known | Screen layout |
| 0x007B1624 | `Photos_Screen` | Known | Screen layout |
| 0x007B1770 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B17D6 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B1838 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B189A | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B1930 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B1A51 | `Photos_Screen` | Known | Screen layout |
| 0x007B1AE8 | `Photos_Screen` | Known | Screen layout |
| 0x007B1C34 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B1C9A | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B1CFC | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B1D5E | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B1DF4 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B1F15 | `Photos_Screen` | Known | Screen layout |
| 0x007B1FAC | `Photos_Screen` | Known | Screen layout |
| 0x007B20F8 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B215E | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B21C0 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B2222 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B22B8 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B23D9 | `Photos_Screen` | Known | Screen layout |
| 0x007B2470 | `Photos_Screen` | Known | Screen layout |
| 0x007B25BC | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B2622 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B2684 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B26E6 | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B277C | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B289D | `Photos_Screen` | Known | Screen layout |
| 0x007B2934 | `Photos_Screen` | Known | Screen layout |
| 0x007B2A80 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x007B2AE6 | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x007B2B48 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x007B2BAA | `TVOutSlideshow_Screen_Playing'` | Known | Screen layout |
| 0x007B2C40 | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x007B2D61 | `Photos_Screen` | Known | Screen layout |
| 0x007B2F81 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B2FE3 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B3051 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B30B7 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B311C | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B33EA | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B344C | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B34BA | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B3520 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3826 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B3888 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B38F6 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B395C | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B3C05 | `Radio_Screen_Default` | Known | Screen layout |
| 0x007B3C62 | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x007B3CC4 | `Radio_Screen_Tuning_Default$` | Known | Screen layout |
| 0x007B3D32 | `Radio_Screen_Default#` | Known | Screen layout |
| 0x007B3D98 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x007B4092 | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007B40FC | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007B44A2 | `RemoteUI_Screen_Main'` | Known | Screen layout |
| 0x007B450C | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x007B4801 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4864 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B48C9 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4931 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B4994 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B49FC | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B4A65 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B4ACB | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4B30 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B4B9D | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4C0D | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007B4C83 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4CF9 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B4D69 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B4DDE | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B4E55 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007B4EC9 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x007B4F3B | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x007B4FB5 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B5028 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x007B509A | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B511E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B5148 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B51CF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B525C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B52FB | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5315 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B538D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B53A7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B5411 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B542E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B54A6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B54D0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B5557 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B55E4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B5683 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B569D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B5715 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B572F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B5799 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B57B6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B582E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B5858 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B58DF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B596C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B5A0B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5A25 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B5A9D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5AB7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B5B21 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B5B3E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B5BB6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B5BE0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B5C67 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B5CF4 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B5D93 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5DAD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B5E25 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B5E3F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B5EA9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B5EC6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B5F3E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B5F68 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B5FEF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B607C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B611B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6135 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B61AD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B61C7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B6231 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B624E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B62C6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B62F0 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B6377 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B6404 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B64A3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B64BD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B6535 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B654F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B65B9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B65D6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B664E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B6678 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B66FF | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B678C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B682B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6845 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B68BD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B68D7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B6941 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B695E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B69D6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B6A00 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B6A87 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B6B14 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B6BB3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6BCD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B6C45 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6C5F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B6CC9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B6CE6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B6D5E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B6D88 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B6E0F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B6E9C | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B6F3B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6F55 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B6FCD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B6FE7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B7051 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B706E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B70E6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B7110 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B7197 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B7224 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B72C3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B72DD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B7355 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B736F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B73D9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B73F6 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B746E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B7498 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B751F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B75AC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B764B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7665 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B76DD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B76F7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B7761 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B777E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B77F6 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B7820 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B78A7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B7934 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B79D3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B79ED | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B7A65 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7A7F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B7AE9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B7B06 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B7B7E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B7BA8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B7C2F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B7CBC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B7D5B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7D75 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B7DED | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B7E07 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B7E71 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B7E8E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B7F06 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B7F30 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B7FB7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B8044 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B80E3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B80FD | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B8175 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B818F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B81F9 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B8216 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B828E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B82B8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B833F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B83CC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B846B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8485 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B84FD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8517 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B8581 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B859E | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B8616 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B8640 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B86C7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B8754 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B87F3 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B880D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B8885 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B889F | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B8909 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B8926 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B899E | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B89C8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B8A4F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B8ADC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B8B7B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8B95 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B8C0D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8C27 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B8C91 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B8CAE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B8D26 | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B8D50 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B8DD7 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B8E64 | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B8F03 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8F1D | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B8F95 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B8FAF | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B9019 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B9036 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B90AE | `SettingsMenu_AboutWithAccessory_Screen6` | Known | Screen layout |
| 0x007B90D8 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout$` | Known | Screen layout |
| 0x007B915F | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x007B91EC | `SettingsMenus_VolumeLimit_Screen"` | Known | Screen layout |
| 0x007B928B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B92A5 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B931D | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B9337 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B93A1 | `SettingsMenu_About_Screen)` | Known | Screen layout |
| 0x007B93BE | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x007B9445 | `SettingsMenus_ResetMainMenu_Screen9` | Known | Screen layout |
| 0x007B9515 | `SettingsMenus_ResetMusicMenu_Screen:` | Known | Screen layout |
| 0x007B95C9 | `SettingsMenus_Main_Screen,` | Known | Screen layout |
| 0x007B963B | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B9655 | `VolumeLimitLock_Screen_Locked.` | Known | Screen layout |
| 0x007B96CD | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x007B96E7 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x007B9A22 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007B9A88 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x007B9AE5 | `Extras_Screen` | Known | Screen layout |
| 0x007B9B39 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x007B9C17 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x007B9C85 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007B9D23 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x007B9D3C | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x007B9DA4 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x007B9E16 | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x007B9E2F | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x007B9E92 | `DemoMode_Screen` | Known | Screen layout |
| 0x007B9EA5 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x007B9F12 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x007B9F2B | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x007B9F9E | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x007B9FB9 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x007BA0C9 | `VoiceMemos_Context_PlayDelete_Screen)` | Known | Screen layout |
| 0x007BA0F1 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x007BA24A | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007BA2B9 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007BA3A5 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x007BA469 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007BA48B | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007BA4F7 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x007BA519 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x007BA696 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BA6B2 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007BA779 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007BA794 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007BA7F7 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007BA85A | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007BA8F1 | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BA90D | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007BA9D4 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007BA9EF | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007BAA52 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007BAAB5 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007BAB4D | `MediaLists_Albums_Screen'` | Known | Screen layout |
| 0x007BAB69 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x007BAC30 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x007BAC4B | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x007BACAE | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x007BAD11 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x007BAD8E | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007BADF9 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007BAE65 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007BAED7 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007BAF44 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x007BAFAF | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x007BB01B | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007BB083 | `DiskMode_ScreenLayout_Loading ` | Known | Screen layout |
| 0x007BB0EF | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x007BB163 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x007BB1D1 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x007BB24A | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x007D7AA8 | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x007D7B2D | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x007D7E1A | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x0098B407 | `VoiceMemos_Screen_InsufficientDiskSpace2` | Known | Screen layout |
| 0x0098CC8B | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0098CCA3 | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0098CCC1 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x0098CDCD | `NowPlaying_Screen_Music` | Known | Screen layout |
| 0x0098CDF9 | `MainMenu_Music_Screen_NoMusic` | Known | Screen layout |
| 0x0098CE17 | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0098CE35 | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0098CF36 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x0098CFEA | `DeleteRental_Confirmation_Screen_Unwatched` | Known | Screen layout |
| 0x0098D040 | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x0098D08C | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0098D18E | `NowPlaying_Screen_ExtraInfoLoadFailed` | Known | Screen layout |
| 0x0098D1E9 | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x0098D202 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x0098D220 | `TVOutSlideshow_Screen_Paused` | Known | Screen layout |
| 0x0098D24F | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x0098D287 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x0098D6BE | `Video_Settings_TV_Screen_Standard` | Known | Screen layout |
| 0x0098D6F0 | `Search_Main_Screen_WithKeyboard` | Known | Screen layout |
| 0x0098D710 | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x0098D755 | `VoiceMemos_Screen_InsufficientDiskSpace` | Known | Screen layout |
| 0x0098D819 | `Video_Settings_TV_Screen_Wide` | Known | Screen layout |
| 0x0098D861 | `CoverFlow_Screen_Backside` | Known | Screen layout |
| 0x009902C7 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x009904CC | `NowPlaying_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x009904F1 | `Slideshow_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x009905C1 | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x009905DB | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x0099066E | `RentalDeleted_Screen_Title` | Known | Screen layout |
| 0x00990689 | `SingleRentalExpiring_Screen_Title` | Known | Screen layout |
| 0x009906AB | `MultipleRentalsExpiring_Screen_Title` | Known | Screen layout |
| 0x009906D0 | `DeleteRental_Screen_Title` | Known | Screen layout |
| 0x00990773 | `VoiceMemos_DeleteAll_ScreenTitle` | Known | Screen layout |
| 0x00990810 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00990853 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00990A44 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x00990B2D | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x00990B46 | `Radio_Screen_Volume` | Known | Screen layout |
| 0x00990B5A | `TVOutSlideshow_Screen_Volume` | Known | Screen layout |
| 0x00990B77 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x00990B96 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x00990C62 | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x00990DB8 | `CoverFlow_Screen_BacksideForCapture` | Known | Screen layout |
| 0x00991DB7 | `Clock_Screen_AddEditDelete` | Known | Screen layout |
| 0x00991DD2 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x009920C9 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x009920FD | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0099213A | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x0099224C | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x0099239C | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x009923D4 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x009923FA | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x0099827A | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x009982A5 | `TVOutSlideshow_Screen_Playing` | Known | Screen layout |
| 0x009982C3 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x009982FD | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x0099839A | `MediaLists_NestedPlaylists_Screen_OTGDialog` | Known | Screen layout |
| 0x00998405 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x00998485 | `Extras_Screen_Debug` | Known | Screen layout |
| 0x0099858F | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x009985AF | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x00998ADD | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x00998B3B | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x00998B56 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00998B69 | `MainMenu_List_ScreenLock` | Known | Screen layout |
| 0x00998B82 | `ExtrasMenu_ScreenLock` | Known | Screen layout |
| 0x00998BF5 | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x00998C16 | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x00998CE9 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x00998D0B | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x00998E12 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x00998E52 | `VoiceMemos_Screen_DeletAllAsk` | Known | Screen layout |
| 0x00998E70 | `NowPlaying_Screen_TVOutAsk` | Known | Screen layout |
| 0x00998FCC | `NowPlaying_Screen_Initial` | Known | Screen layout |
| 0x00998FE6 | `NowPlaying_Screen_Video_TVOut_Initial` | Known | Screen layout |
| 0x00999D4E | `Radio_Screen_Tuning_Japan` | Known | Screen layout |
| 0x00999DCF | `RemoteUI_Screen` | Known | Screen layout |
| 0x00999DDF | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x00999DF7 | `Video_Settings_TV_Screen` | Known | Screen layout |
| 0x00999E10 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00999E27 | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x00999E4B | `AddressViewer_PartialLoad_Screen` | Known | Screen layout |
| 0x00999E6C | `Alarms_Alarm_Clock_Triggered_Screen` | Known | Screen layout |
| 0x00999E90 | `Alarms_Alarm_Triggered_Screen` | Known | Screen layout |
| 0x00999EAE | `Unsupported_Screen` | Known | Screen layout |
| 0x00999EC1 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x00999EDF | `LockediPod_Screen` | Known | Screen layout |
| 0x00999EF1 | `DiskMode_Screen` | Known | Screen layout |
| 0x00999F01 | `DemoMode_Screen` | Known | Screen layout |
| 0x00999F11 | `Notes_Image_Screen` | Known | Screen layout |
| 0x00999F24 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x00999F42 | `NowPlaying_Idle_Screen` | Known | Screen layout |
| 0x00999F59 | `Game_Screen` | Known | Screen layout |
| 0x00999F65 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x00999F82 | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x00999F9B | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x00999FBC | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x00999FE1 | `PhotoBrowse_Screen` | Known | Screen layout |
| 0x00999FF4 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x0099A011 | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x0099A032 | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x0099A057 | `Notes_Loading_Screen` | Known | Screen layout |
| 0x0099A06C | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0099A082 | `SettingsMenus_AdjustScrolling_Screen` | Known | Screen layout |
| 0x0099A0A7 | `Game_Running_Screen` | Known | Screen layout |
| 0x0099A0BB | `Stopwatch_Screen` | Known | Screen layout |
| 0x0099A0CC | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0099A0E3 | `Clock_Screen` | Known | Screen layout |
| 0x0099A0F0 | `TVOutSlideshowAsk_Screen` | Known | Screen layout |
| 0x0099A109 | `Settings_Legal_Screen` | Known | Screen layout |
| 0x0099A11F | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x0099A13D | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x0099A159 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x0099A16A | `DemoMode_Main_Screen` | Known | Screen layout |
| 0x0099A17F | `Search_Main_Screen` | Known | Screen layout |
| 0x0099A192 | `AddressViewer_Main_Screen` | Known | Screen layout |
| 0x0099A1AC | `Speakers_Main_Screen` | Known | Screen layout |
| 0x0099A1C1 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x0099A1D7 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x0099A1F1 | `Clock_Region_Screen` | Known | Screen layout |
| 0x0099A205 | `RentalDeleted_Notification_Screen` | Known | Screen layout |
| 0x0099A227 | `SingleRentalExpiring_Notification_Screen` | Known | Screen layout |
| 0x0099A250 | `MultipleRentalsExpiring_Notification_Screen` | Known | Screen layout |
| 0x0099A27C | `RentalError_Notification_Screen` | Known | Screen layout |
| 0x0099A29C | `DeleteRental_Confirmation_Screen` | Known | Screen layout |
| 0x0099A2BD | `LockConfirmation_Screen` | Known | Screen layout |
| 0x0099A2D5 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x0099A2F3 | `Video_Settings_Fit_To_Screen` | Known | Screen layout |
| 0x0099A310 | `RentalInfo_Screen` | Known | Screen layout |
| 0x0099A322 | `Radio_Screen` | Known | Screen layout |
| 0x0099A32F | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0099A343 | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x0099A35D | `ChargingMode_LowPower_Screen` | Known | Screen layout |
| 0x0099A37A | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x0099A394 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x0099A3AE | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x0099A3C8 | `Genius_Error_Screen` | Known | Screen layout |
| 0x0099A3DC | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0099A3F5 | `Extras_Screen` | Known | Screen layout |
| 0x0099A403 | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x0099A420 | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x0099A442 | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x0099A45B | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x0099A479 | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x0099A492 | `Video_Settings_Screen` | Known | Screen layout |
| 0x0099A4A8 | `SettingsMenus_AudiobookSettings_Screen` | Known | Screen layout |
| 0x0099A4CF | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x0099A4F5 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0099A50B | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0099A523 | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x0099A546 | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x0099A563 | `MediaLists_Rentals_Screen` | Known | Screen layout |
| 0x0099A57D | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x0099A5A1 | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x0099A5BA | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x0099A5DC | `ToDo_Instructions_Screen` | Known | Screen layout |
| 0x0099A5F5 | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x0099A611 | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x0099A62B | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x0099A64C | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x0099A668 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0099A680 | `VoiceMemos_Screen` | Known | Screen layout |
| 0x0099A692 | `No_Photos_Screen` | Known | Screen layout |
| 0x0099A6A3 | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x0099A6BD | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x0099A6D9 | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x0099A6FD | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x0099A71D | `ChargingMode_Progress_Screen` | Known | Screen layout |
| 0x0099A73A | `Notes_Contents_Screen` | Known | Screen layout |
| 0x0099A750 | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x0099A76B | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x0099A787 | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x0099A7A9 | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x0099A7CA | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x0099A7E4 | `MediaLists_Genius_Screen` | Known | Screen layout |
| 0x0099A7FD | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x0099A817 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0099A836 | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x0099A857 | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x0099A86F | `NoContent_Screen` | Known | Screen layout |
| 0x0099A880 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x0099A896 | `FirstBoot_Screen` | Known | Screen layout |
| 0x0099A8A7 | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x0099A8BD | `Notes_List_Screen` | Known | Screen layout |
| 0x0099A8CF | `Debug_TestList_Screen` | Known | Screen layout |
| 0x0099A8E5 | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x0099A906 | `MediaLists_GeniusPlaylist_Screen` | Known | Screen layout |
| 0x0099A927 | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x0099A941 | `TODOS_Menu_Screen` | Known | Screen layout |
| 0x0099A953 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x0099A969 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x0099A985 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0099A99A | `Games_Menu_Screen` | Known | Screen layout |
| 0x0099A9AC | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0099A9BF | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x0099A9DE | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x0099A9FD | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x0099AA21 | `ContextualMenu_Screen` | Known | Screen layout |
| 0x0099AA37 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x0099AA4D | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x0099AA6B | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x0099AA8E | `TVOutSlideshow_Screen` | Known | Screen layout |
| 0x0099AAA4 | `CoverFlow_Screen` | Known | Screen layout |
| 0x0099AAB5 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x0099AAC9 | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x0099AAEB | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x0099AB03 | `Stopwatch_SessionSummary_Screen` | Known | Screen layout |
| 0x0099AB23 | `SettingsMenu_AboutWithAccessory_Screen` | Known | Screen layout |
| 0x0099AB4A | `Clock_NorthAmerica_City_Screen` | Known | Screen layout |
| 0x0099AB69 | `Clock_SouthAmerica_City_Screen` | Known | Screen layout |
| 0x0099AB88 | `Clock_Africa_City_Screen` | Known | Screen layout |
| 0x0099ABA1 | `Clock_Australia_City_Screen` | Known | Screen layout |
| 0x0099ABBD | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x0099ABD4 | `Clock_Pacific_City_Screen` | Known | Screen layout |
| 0x0099ABEE | `Clock_Atlantic_City_Screen` | Known | Screen layout |
| 0x0099AC09 | `Clock_Europe_City_Screen` | Known | Screen layout |
| 0x0099ACE9 | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x0099AD3A | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0099AD5D | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0099AD85 | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x0099B120 | `NowPlaying_Screen_Video_Caption` | Known | Screen layout |
| 0x0099B223 | `RentalInfo_Screen_ExpiringSoon` | Known | Screen layout |
| 0x0099B279 | `RentalInfo_Screen_NoAlbumArt_ExpiringSoon` | Known | Screen layout |
| 0x0099B648 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x0099B69E | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x0099B7EF | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x0099B80C | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x0099BBE0 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x0099BD02 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0099BD24 | `CoverFlow_Screen_DefaultNoTextNoStatusBar` | Known | Screen layout |
| 0x0099BD91 | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x0099BDB0 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x0099C3F2 | `ChargingMode_ScreenLayout_LowPower` | Known | Screen layout |
| 0x0099CD8F | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x0099CDA8 | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x0099CEF0 | `MainMenus_Main_Screen_Extras` | Known | Screen layout |
| 0x0099CFCC | `MainMenu_Main_Screen_NoMovies` | Known | Screen layout |
| 0x0099CFEA | `MainMenu_Videos_Screen_NoMovies` | Known | Screen layout |
| 0x0099D00A | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0099D115 | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x0099D131 | `Extras_Screen_Games` | Known | Screen layout |
| 0x0099D237 | `MainMenu_Music_Screen_NoGenres` | Known | Screen layout |
| 0x0099D256 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x0099D272 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x0099D33D | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x0099D418 | `MainMenu_Music_Screen_NoSongs` | Known | Screen layout |
| 0x0099D5E6 | `MainMenu_Music_Screen_NoAudiobooks` | Known | Screen layout |
| 0x0099D609 | `MainMenus_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x0099D62C | `NoContent_Screen_NoAudiobooks` | Known | Screen layout |
| 0x0099D666 | `MainMenu_Main_Screen_NoRentals` | Known | Screen layout |
| 0x0099D685 | `MainMenu_Videos_Screen_NoRentals` | Known | Screen layout |
| 0x0099D6A6 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x0099D755 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x0099D772 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x0099D7F1 | `MainMenu_Music_Screen_NoAlbums` | Known | Screen layout |
| 0x0099D8D5 | `MainMenu_Music_Screen_NoCompilations` | Known | Screen layout |
| 0x0099D8FA | `MainMenus_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x0099DA81 | `MainMenu_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0099DAA4 | `MainMenu_Videos_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0099DAC9 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x0099DAE8 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x0099DB07 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x0099DB28 | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x0099DB66 | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x0099DB87 | `Extras_Screen_VoiceMemos` | Known | Screen layout |
| 0x0099DBF2 | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x0099DC24 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x0099DC43 | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x0099DCF0 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x0099DD5C | `MainMenu_Music_Screen_NoComposers` | Known | Screen layout |
| 0x0099DE55 | `Slideshow_Screen_Brightness` | Known | Screen layout |
| 0x0099DE71 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x0099DEF4 | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x0099DF0F | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0099DF30 | `NowPlaying_Screen_Video_TVOut_Progress` | Known | Screen layout |
| 0x0099DFDF | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x0099E013 | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x0099E034 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x0099E0D7 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x0099E0F8 | `MainMenu_Videos_Screen_NoPlaylists` | Known | Screen layout |
| 0x0099E11B | `NoContent_Screen_NoPlaylists` | Known | Screen layout |
| 0x0099E16A | `MainMenu_Music_Screen_NoArtists` | Known | Screen layout |
| 0x0099E211 | `NowPlaying_Screen_Genius` | Known | Screen layout |
| 0x0099E25A | `Genius_Error_Screen_NoGenius` | Known | Screen layout |
| 0x0099E277 | `MainMenus_Main_Screen_NikePlus` | Known | Screen layout |
| 0x0099E296 | `Extras_Screen_NikePlus` | Known | Screen layout |
| 0x0099E3E6 | `MainMenu_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x0099E405 | `MainMenu_Videos_Screen_NoTVShows` | Known | Screen layout |
| 0x0099E426 | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x0099E891 | `MainMenus_Main_Screen_Backlight` | Known | Screen layout |
| 0x0099E944 | `Clock_Screen_AddEdit` | Known | Screen layout |
| 0x0099E9BE | `Notes_Contents_Screen_Alt` | Known | Screen layout |
| 0x0099E9D8 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x0099EA84 | `VoiceMemos_Screen_InsufficientDiskSpace2_Default` | Known | Screen layout |
| 0x0099EB36 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x0099EBDB | `VoiceMemos_Screen_InsufficientDiskSpace_Default` | Known | Screen layout |
| 0x0099EC0B | `NowPlaying_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x0099EC38 | `Slideshow_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x0099F8C9 | `Radio_Screen_Tuning_Default` | Known | Screen layout |
| 0x0099F92A | `VoiceMemos_Screen_DeletAllAsk_Default` | Known | Screen layout |
| 0x0099F950 | `NowPlaying_Screen_TVOutAsk_Default` | Known | Screen layout |
| 0x0099F973 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0099F991 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x0099F9BD | `AddressViewer_PartialLoad_Screen_Default` | Known | Screen layout |
| 0x0099F9E6 | `Alarms_Alarm_Clock_Triggered_Screen_Default` | Known | Screen layout |
| 0x0099FA12 | `Alarms_Alarm_Triggered_Screen_Default` | Known | Screen layout |
| 0x0099FA38 | `Unsupported_Screen_Default` | Known | Screen layout |
| 0x0099FA53 | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x0099FA79 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x0099FA91 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x0099FAAC | `Game_Screen_Default` | Known | Screen layout |
| 0x0099FAC0 | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x0099FAE6 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0099FB07 | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x0099FB30 | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x0099FB5A | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x0099FB87 | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x0099FBB0 | `Notes_Loading_Screen_Default` | Known | Screen layout |
| 0x0099FBCD | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x0099FBEB | `Clock_Screen_Default` | Known | Screen layout |
| 0x0099FC00 | `TVOutSlideshowAsk_Screen_Default` | Known | Screen layout |
| 0x0099FC21 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x0099FC3F | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x0099FC65 | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x0099FC89 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x0099FCA2 | `AddressViewer_Main_Screen_Default` | Known | Screen layout |
| 0x0099FCC4 | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x0099FCE1 | `MainMenus_Main_Screen_Default` | Known | Screen layout |
| 0x0099FCFF | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0099FD1C | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x0099FD38 | `RentalDeleted_Notification_Screen_Default` | Known | Screen layout |
| 0x0099FD62 | `SingleRentalExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x0099FD93 | `MultipleRentalsExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x0099FDC7 | `RentalError_Notification_Screen_Default` | Known | Screen layout |
| 0x0099FDEF | `DeleteRental_Confirmation_Screen_Default` | Known | Screen layout |
| 0x0099FE18 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x0099FE44 | `RentalInfo_Screen_Default` | Known | Screen layout |
| 0x0099FE5E | `Radio_Screen_Default` | Known | Screen layout |
| 0x0099FE73 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0099FE8F | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0099FEB1 | `Extras_Screen_Default` | Known | Screen layout |
| 0x0099FEC7 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x0099FEED | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x0099FF0E | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x0099FF2C | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x0099FF4E | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x0099FF7A | `ToDo_Instructions_Screen_Default` | Known | Screen layout |
| 0x0099FF9B | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x0099FFBF | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x0099FFE1 | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x009A0005 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x009A0024 | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x009A003D | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x009A005F | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x009A0083 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x009A00A1 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x009A00C5 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x009A00EF | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x009A0118 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x009A013A | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x009A015B | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x009A017B | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x009A0199 | `FirstBoot_Screen_Default` | Known | Screen layout |
| 0x009A01B2 | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x009A01D0 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x009A01EA | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x009A0208 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x009A0231 | `MediaLists_GeniusPlaylist_Screen_Default` | Known | Screen layout |
| 0x009A025A | `TODOS_Menu_Screen_Default` | Known | Screen layout |
| 0x009A0274 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x009A0292 | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x009A02AF | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x009A02C9 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x009A02E4 | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x009A0303 | `ContextualMenu_Screen_Default` | Known | Screen layout |
| 0x009A0321 | `Debug_MainMenu_Screen_Default` | Known | Screen layout |
| 0x009A033F | `TVOutSlideshow_Screen_Default` | Known | Screen layout |
| 0x009A035D | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x009A0376 | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x009A0392 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x009A03BC | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x009A03DC | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x009A0404 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009A042B | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x009A0452 | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x009A0473 | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x009A0497 | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x009A04B6 | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x009A04D8 | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x009A04FB | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x009A051C | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x009A05AA | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x009A05DA | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x009A05FC | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x009A066D | `RentalInfo_Screen_NoAlbumArt_Default` | Known | Screen layout |
| 0x009A0692 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x009A0C6D | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009A0C99 | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x009A0CDE | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x009A0D06 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x009A0D27 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x009A0D48 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x009A0D6E | `NowPlaying_Screen_ChapterArt` | Known | Screen layout |
| 0x009A0D8B | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x009A0DAD | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x009A0DD1 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x009A0DF5 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x009A0FC5 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x009A10A0 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x009A10F1 | `NowPlaying_Screen_Video_TVOut` | Known | Screen layout |
| 0x009A1263 | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x009A128A | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x009A17C3 | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x009A1980 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x009A1B72 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x009A1E3E | `Lock_Screen_Default_Layout` | Known | Screen layout |
| 0x009A1ED4 | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x009A1EFB | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x009A2117 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x009A21F1 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x009A2258 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009A2282 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x009A4B86 | `TVOutSlideshow_Screen_BatteryLow` | Known | Screen layout |
| 0x009A4BD2 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x009A4CB0 | `MainMenu_List_ScreenLock_x` | Known | Screen layout |
| 0x009A4F7E | `RentalInfo_Screen_ExpiresToday` | Known | Screen layout |
| 0x009A4FD4 | `RentalInfo_Screen_NoAlbumArt_ExpiresToday` | Known | Screen layout |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000909B | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | Known | RTOS |
| 0x002A1434 | `  K - RTXC` | Known | RTOS |
| 0x002A243C | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | Known | RTOS |
| 0x00989FF8 | `Returning from RTXCBug` | Known | RTOS |

---

## 9. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000D1F28 | `HostOSTask` | Known | RTOS task thread |
| 0x0012C414 | `MP3ExampleTask` | Known | RTOS task thread |
| 0x001318D4 | `USBDeviceTask` | Known | RTOS task thread |
| 0x0013BCE0 | `DiskReaderTask` | Known | RTOS task thread |
| 0x0014BE4C | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x0014BE60 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0019ED64 | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x001DA330 | `MeCCAIOTask` | Known | RTOS task thread |
| 0x0020D594 | `StreamCacheMassStorageManagerTimeOutTask` | Known | RTOS task thread |
| 0x0020D710 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x0028F5D4 | `FirewireTask` | Known | RTOS task thread |
| 0x0028F5E8 | `TouchwheelTask` | Known | RTOS task thread |
| 0x0028F5FC | `AudioOutStateTask` | Known | RTOS task thread |
| 0x0028F628 | `DiskMgrTask` | Known | RTOS task thread |
| 0x0028F638 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x0028F64C | `MikeyTask` | Known | RTOS task thread |
| 0x0028F65C | `TopPlugTask` | Known | RTOS task thread |
| 0x0028F66C | `HPhoneDetTask` | Known | RTOS task thread |
| 0x0028F6E4 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x0028F70C | `AlarmTask` | Known | RTOS task thread |
| 0x0028F72B | `"USBAudioTask` | Known | RTOS task thread |
| 0x002A1AD4 | `Undefined Task` | Known | RTOS task thread |
| 0x003E3A60 | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x003E7AE8 | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x003F01F8 | `MeCCARecordingTask` | Known | RTOS task thread |
| 0x008DC6C0 | `WaveFileDebugTask` | Known | RTOS task thread |

---

## 10. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00247C70 | `Channel Reserved` | Known | Logging channel |
| 0x00247C84 | `Channel AppBoot` | Known | Logging channel |
| 0x00247C94 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x00247CB0 | `Channel PrefsWriting` | Known | Logging channel |
| 0x00247CC8 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x00247CE8 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x00247D00 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x00247D1C | `Channel TestLogging` | Known | Logging channel |
| 0x00247D30 | `Channel AppFileLoading` | Known | Logging channel |
| 0x00247D48 | `Channel VCardReading` | Known | Logging channel |
| 0x00247D60 | `Channel LongSongScanning` | Known | Logging channel |
| 0x00247DD4 | `Channel VoiceRecording` | Known | Logging channel |
| 0x00247DEC | `Channel PhotoImporting` | Known | Logging channel |
| 0x00247E04 | `Channel Notes` | Known | Logging channel |
| 0x00247E14 | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x00247E30 | `Channel DiskMode` | Known | Logging channel |
| 0x00247E44 | `Channel Firewire` | Known | Logging channel |
| 0x00247E58 | `Channel USB` | Known | Logging channel |
| 0x00247E78 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x00247E90 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00081B70 | `gamedata_RW` | Known | Game system |
| 0x00081B8C | `gamedata_ShareRW` | Known | Game system |
| 0x00081BA0 | `games_RO` | Known | Game system |
| 0x0095EECF | `11TCGamesMenu` | Known | Game system |
| 0x0095EFA3 | `12TCGameScreen` | Known | Game system |
| 0x0095FD6F | `27TSilverCntlrTransitionAddonI11TCGamesMenuE` | Known | Game system |
| 0x0095FE24 | `27TSilverCntlrTransitionAddonI12TCGameScreenE` | Known | Game system |
| 0x0098A052 | `iPod_Control/games_RO/` | Known | Game system |
| 0x0098A069 | `Resources/Games/games_RO/` | Known | Game system |
| 0x00995A8C | `GamesMenu_StatusBar_String` | Known | Game system |
| 0x00996214 | `AboutScreen_Games_String` | Known | Game system |
| 0x0099D145 | `MainMenu_List_Games` | Known | Game system |
| 0x0099D159 | `ExtrasMenu_Games` | Known | Game system |
| 0x009A4D1F | `MainMenu_List_Games_x` | Known | Game system |

---

## 12. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00092064 | `adrmmp4a` | Known | DRM system |
| 0x001391B0 | `AppleDRMVersion` | Known | DRM system |
| 0x00139250 | `AppleDRM` | Known | DRM system |
| 0x0013A4C0 | `AppleVideoDRM` | Known | DRM system |
| 0x0013DA50 | `tx3gdrmsp608aavdmp4aesds ` | Known | DRM system |
| 0x001E7CD8 | `drmttx3g` | Known | DRM system |
| 0x0098A4DB | `DRMLevel` | Known | DRM system |

---

## 13. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00031360 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00031378 | `iPod_Control\iTunes\firsttime` | Known | iTunes database |
| 0x00052750 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x00052778 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00058A3C | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0007DB88 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x00081B00 | `iPod_Control/iTunes/iTunesDB.p7b` | Known | iTunes database |
| 0x00094F18 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0009E53C | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x0009E724 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x0009EFEC | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000A7164 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000A8624 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000A8724 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00124A94 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x0021A3E4 | `%s/sqlite_` | Known | SQLite database |
| 0x002806F0 | `iPod_Control/iTunes/primary.db` | Known | iTunes database |
| 0x002812B8 | `iPod_Control/iTunes/Extras.itdb` | Known | iTunes database |
| 0x002A4F28 | `sqlite3BtreeInitPage() returns error code %d` | Known | SQLite database |
| 0x002A8204 | `sqlite_master` | Known | SQLite database |
| 0x002A8214 | `sqlite_temp_master` | Known | SQLite database |
| 0x002BF124 | `sqlite_stat1` | Known | SQLite database |
| 0x002BF134 | `CREATE TABLE %Q.sqlite_stat1(tbl,idx,stat)` | Known | SQLite database |
| 0x002BF160 | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x002C9B08 | `sqlite_subquery_%p_` | Known | SQLite database |
| 0x0035ED2C | `sqlite_master` | Known | SQLite database |
| 0x0035ED3C | `sqlite_temp_master` | Known | SQLite database |
| 0x0035F060 | `sqlite_` | Known | SQLite database |
| 0x0035F0A0 | `sqlite_master` | Known | SQLite database |
| 0x0035F0B0 | `sqlite_temp_master` | Known | SQLite database |
| 0x0035F0C8 | `sqlite_sequence` | Known | SQLite database |
| 0x0035F0D8 | `UPDATE "%w".sqlite_sequence set name = %Q WHERE name = %Q` | Known | SQLite database |
| 0x0035F1BC | `sqlite_stat1` | Known | SQLite database |
| 0x0035F1CC | `SELECT idx, stat FROM %Q.sqlite_stat1` | Known | SQLite database |
| 0x0035FEA8 | `sqlite_` | Known | SQLite database |
| 0x003600A4 | `sqlite_master` | Known | SQLite database |
| 0x003600B4 | `sqlite_temp_master` | Known | SQLite database |
| 0x00362DD0 | `sqlite_` | Known | SQLite database |
| 0x003640BC | `sqlite_autoindex_` | Known | SQLite database |
| 0x003640D0 | `sqlite_master` | Known | SQLite database |
| 0x003640E0 | `sqlite_temp_master` | Known | SQLite database |
| 0x00365538 | `sqlite_master` | Known | SQLite database |
| 0x00365548 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036557C | `sqlite_stat1` | Known | SQLite database |
| 0x0036558C | `DELETE FROM %Q.sqlite_stat1 WHERE idx=%Q` | Known | SQLite database |
| 0x00365874 | `sqlite_master` | Known | SQLite database |
| 0x00365884 | `sqlite_temp_master` | Known | SQLite database |
| 0x003658F8 | `DELETE FROM %s.sqlite_sequence WHERE name=%Q` | Known | SQLite database |
| 0x00365960 | `sqlite_stat1` | Known | SQLite database |
| 0x00365970 | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x00365CE8 | `sqlite_master` | Known | SQLite database |
| 0x00365CF8 | `sqlite_temp_master` | Known | SQLite database |
| 0x00366110 | `sqlite_master` | Known | SQLite database |
| 0x00366120 | `sqlite_temp_master` | Known | SQLite database |
| 0x00366138 | `CREATE TABLE %Q.sqlite_sequence(name,seq)` | Known | SQLite database |
| 0x003693C0 | `sqlite_master` | Known | SQLite database |
| 0x003693D0 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036B7B8 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036B7D0 | `sqlite_master` | Known | SQLite database |
| 0x0036CFAC | `sqlite3_extension_init` | Known | SQLite database |
| 0x0036D7A0 | `sqlite_master` | Known | SQLite database |
| 0x0036D7B0 | `sqlite_temp_master` | Known | SQLite database |
| 0x00371B90 | `sqlite_attach` | Known | SQLite database |
| 0x00371BA4 | `sqlite_detach` | Known | SQLite database |
| 0x003748D8 | `sqlite_master` | Known | SQLite database |
| 0x003748E8 | `sqlite_temp_master` | Known | SQLite database |
| 0x00374938 | `sqlite_sequence` | Known | SQLite database |
| 0x0037A1C4 | `sqlite_master` | Known | SQLite database |
| 0x0037A1D4 | `sqlite_temp_master` | Known | SQLite database |
| 0x0037D568 | `sqlite_master` | Known | SQLite database |
| 0x0037D578 | `sqlite_temp_master` | Known | SQLite database |
| 0x0038B62C | `sqlite_attach` | Known | SQLite database |
| 0x0038B63C | `sqlite_detach` | Known | SQLite database |
| 0x003DCD44 | `iTunesDB` | Known | iTunes database |
| 0x003DCD50 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x008D870F | `SQLite format 3` | Known | SQLite database |
| 0x008DADBC | `CREATE TABLE sqlite_master(` | Known | SQLite database |
| 0x008DAE24 | `CREATE TEMP TABLE sqlite_temp_master(` | Known | SQLite database |
| 0x008DB4EC | `illegal return value (%d) from the authorization function - should be SQLITE_OK,` | Known | SQLite database |
| 0x008DB5A4 | `SELECT 'CREATE TABLE vacuum_db.' || substr(sql,14)   FROM sqlite_master WHERE ty` | Known | SQLite database |
| 0x008DB62C | `SELECT 'CREATE INDEX vacuum_db.' || substr(sql,14)  FROM sqlite_master WHERE sql` | Known | SQLite database |
| 0x008DB694 | `SELECT 'CREATE UNIQUE INDEX vacuum_db.' || substr(sql,21)   FROM sqlite_master W` | Known | SQLite database |
| 0x008DB70C | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x008DB7BC | `SELECT 'DELETE FROM vacuum_db.' || quote(name) || ';' FROM vacuum_db.sqlite_mast` | Known | SQLite database |
| 0x008DB830 | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x008DB8C8 | `INSERT INTO vacuum_db.sqlite_master   SELECT type, name, tbl_name, rootpage, sql` | Known | SQLite database |
| 0x008DBA88 | `UPDATE %Q.%s SET sql = CASE WHEN type = 'trigger' THEN sqlite_rename_trigger(sql` | Known | SQLite database |
| 0x008DBBFC | `UPDATE sqlite_temp_master SET sql = sqlite_rename_trigger(sql, %Q), tbl_name = %` | Known | SQLite database |
| 0x008DBE38 | `sqlite3_get_table() called with two or more incompatible queries` | Known | SQLite database |
| 0x009A57AE | `sqlite_rename_table` | Known | SQLite database |
| 0x009A5931 | `sqlite_version` | Known | SQLite database |
| 0x009A59CB | `sqlite_rename_trigger` | Known | SQLite database |
| 0x009A5CEF | `SQLite_iPod_VFS` | Known | SQLite database |

---

## 14. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0005EE34 | `cI: could not read CE-ATA task file` | Known | Hardware |
| 0x0005EE5C | `cI: CE-ATA signature missing (%x,%x)` | Known | Hardware |
| 0x0005EEB4 | `cI: CE-ATA interrupt enable failed` | Known | Hardware |
| 0x001242F8 | `NAND FLASH DRIVE` | Known | Hardware |
| 0x001396F8 | `FireWireGUID` | Known | FireWire |
| 0x00139708 | `FireWireVersion` | Known | FireWire |
| 0x00139DE4 | `FireWire` | Known | FireWire |
| 0x0035A09C | `CE-ATA init failed` | Known | Hardware |
| 0x0035A55C | `ISDIE: CE-ATA interrupt enable failed` | Known | Hardware |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00726536 | `Settings_RadioRegionsMenu_Layout"` | Known | FM Radio |
| 0x007265BF | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x007D6F50 | `Radio Regions` | Known | FM Radio |
| 0x00826EDC | `Radio-Regionen` | Known | FM Radio |
| 0x0095F96C | `23TCSettings_RadioRegions` | Known | FM Radio |
| 0x0096087F | `27TSilverCntlrTransitionAddonI23TCSettings_RadioRegionsE` | Known | FM Radio |
| 0x00992AED | `Settings_RadioRegions_Australia_String` | Known | FM Radio |
| 0x00992B14 | `Settings_RadioRegions_Asia_String` | Known | FM Radio |
| 0x00993D79 | `Settings_RadioRegions_Europe_String` | Known | FM Radio |
| 0x00995388 | `Settings_RadioRegions_Japan_String` | Known | FM Radio |
| 0x00996031 | `Settings_RadioRegions_Americas_String` | Known | FM Radio |
| 0x00996713 | `SettingsMenu_RadioRegions_String` | Known | FM Radio |
| 0x00999C0D | `RadioRegionsMenuItem` | Known | FM Radio |
| 0x0099D85E | `SettingsMenu_ListItem_RadioRegions` | Known | FM Radio |
| 0x009A1A4C | `SettingsInfo_Template_RadioRegions_Layout` | Known | FM Radio |
| 0x009A1A76 | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x009A20D8 | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 16. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00865CC4 | `Fotocamera` | Known | Camera |
| 0x00866228 | `Collega fotocamera o inserisci scheda` | Known | Camera |
| 0x008662A0 | `Fotocamera non supportata` | Known | Camera |
| 0x0088536C | `Camera` | Known | Camera |
| 0x008858EC | `Sluit camera of kaart aan` | Known | Camera |
| 0x00885958 | `Camera niet ondersteund` | Known | Camera |
| 0x00992B36 | `Photos_DevName_Camera_String` | Known | Camera |

---

## 17. Pedometer

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x009A5087 | `NikePlus_Remote_Step_Away` | Known | Pedometer |
| 0x009A50A1 | `NikePlus_Step_Away` | Known | Pedometer |
| 0x009A596C | `AggStep` | Known | Pedometer |

---

## 18. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0003134C | `iPod_Control` | Filesystem Path |  |
| 0x000313B8 | `iPod_Control\Device` | Filesystem Path |  |
| 0x0003FC14 | `iPod_Control\Device` | Filesystem Path |  |
| 0x00041CA0 | `iPod_Control` | Filesystem Path |  |
| 0x0004230C | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x00052730 | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path |  |
| 0x000552D8 | `iPod_Control\Music\` | Filesystem Path |  |
| 0x000588BC | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x0008BA48 | `iPod_Control` | Filesystem Path |  |
| 0x0008BA58 | `Resources/Games` | Filesystem Path |  |
| 0x0008BA68 | `iPod_Control/%s%s%s` | Filesystem Path |  |
| 0x000F386C | `iPod_Control/Device/alarms` | Filesystem Path |  |
| 0x00103D78 | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x001052B8 | `iPod_Control/Device` | Filesystem Path |  |
| 0x001052CC | `iPod_Control/Device/radio` | Filesystem Path |  |
| 0x0011F498 | `iPod_Control/Device/Users` | Filesystem Path |  |
| 0x0014D2F0 | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x0014D54C | `/iPod_Control/Device/1da` | Filesystem Path |  |
| 0x00159954 | `Resources/UI/active.bin` | Filesystem Path |  |
| 0x0015996C | `Resources/UI/` | Filesystem Path |  |
| 0x0017D210 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x0017E13C | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path |  |
| 0x0017E164 | `/iPod_Control/Device/Accessories` | Filesystem Path |  |
| 0x001A23AC | `iPod_Control/Device/PlayCounts` | Filesystem Path |  |
| 0x001B837C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B842C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B85A8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B8740 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B87E8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B8998 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B8A3C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B8AE0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B8B84 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B8C28 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B8CD8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B8D7C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B8E20 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B8ED0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B8F80 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9030 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B919C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B924C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B92FC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B93A0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9450 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9544 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B95E8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B969C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9758 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9808 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B992C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B99E8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9A98 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9C54 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9D18 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9DC8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9E84 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001B9FC0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA08C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA148 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA1EC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA290 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA34C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA408 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA4D0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA574 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA63C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA704 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA7B4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA87C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA944 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BA9F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAAA4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAB68 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAC18 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BACC8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAD78 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAE4C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BAF20 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB020 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB100 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB208 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001BB2F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003DCDC2 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003E3304 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x003E5E90 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x003E62E2 | `iPod_Control/Device` | Filesystem Path |  |
| 0x003E7C54 | `Resources/Fonts` | Filesystem Path |  |
| 0x003F01C4 | `iPod_Control:Device:voices:` | Filesystem Path |  |
| 0x00989F2D | `Resources/Games/` | Filesystem Path |  |
| 0x0098A34B | `iPod_Control/Device` | Filesystem Path |  |
| 0x0098A35F | `iPod_Control/Device/accessories.reg` | Filesystem Path |  |
| 0x0098A452 | `iPod_Control/Tones` | Filesystem Path |  |

---

## 19. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x008DEE00 | `c:\bwa\N25BFirmwareWin-69\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftu` | Build Path |  |
| 0x008DEE58 | `c:\bwa\N25BFirmwareWin-69\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fts` | Build Path |  |
| 0x008DEEB0 | `c:\bwa\N25BFirmwareWin-69\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fto` | Build Path |  |
| 0x008E9B00 | `c:\bwa\N25BFirmwareWin-69\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\` | Build Path |  |
| 0x008EA67C | `c:\bwa\N25BFirmwareWin-69\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfd` | Build Path |  |
| 0x008EB878 | `c:\bwa\N25BFirmwareWin-69\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrg` | Build Path |  |
| 0x008EB8D0 | `c:\bwa\N25BFirmwareWin-69\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrc` | Build Path |  |
| 0x008EB928 | `c:\bwa\N25BFirmwareWin-69\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfro` | Build Path |  |
| 0x008EBC6C | `c:\bwa\N25BFirmwareWin-69\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1` | Build Path |  |
| 0x008FB014 | `c:\bwa\N25BFirmwareWin-69\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttc` | Build Path |  |
| 0x008FB290 | `c:\bwa\N25BFirmwareWin-69\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype` | Build Path |  |
| 0x008FB7FC | `c:\bwa\N25BFirmwareWin-69\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1` | Build Path |  |

---

## 20. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00089690 | `Acoustic` | EQ Preset |  |
| 0x0008969C | `Bass Booster` | EQ Preset |  |
| 0x000896BC | `Classical` | EQ Preset |  |
| 0x000896C8 | `Dance` | EQ Preset |  |
| 0x000896D8 | `Electronic` | EQ Preset |  |
| 0x000896EC | `Hip Hop` | EQ Preset |  |
| 0x000896F4 | `Jazz` | EQ Preset |  |
| 0x000896FC | `Latin` | EQ Preset |  |
| 0x00089704 | `Loudness` | EQ Preset |  |
| 0x00089710 | `Lounge` | EQ Preset |  |
| 0x00089718 | `Piano` | EQ Preset |  |
| 0x0008972C | `Rock` | EQ Preset |  |
| 0x00089734 | `Small Speakers` | EQ Preset |  |
| 0x00089744 | `Spoken Word` | EQ Preset |  |
| 0x00089750 | `Treble Booster` | EQ Preset |  |
| 0x0008979C | `Vocal Booster` | EQ Preset |  |
| 0x007D7240 | `Acoustic` | EQ Preset |  |
| 0x007D724C | `Bass Booster` | EQ Preset |  |
| 0x007D726C | `Classical` | EQ Preset |  |
| 0x007D7278 | `Dance` | EQ Preset |  |
| 0x007D7288 | `Electronic` | EQ Preset |  |
| 0x007D729C | `Hip Hop` | EQ Preset |  |
| 0x007D72A4 | `Jazz` | EQ Preset |  |
| 0x007D72AC | `Latin` | EQ Preset |  |
| 0x007D72B4 | `Loudness` | EQ Preset |  |
| 0x007D72C0 | `Lounge` | EQ Preset |  |
| 0x007D72C8 | `Piano` | EQ Preset |  |
| 0x007D72D8 | `Rock` | EQ Preset |  |
| 0x007D72E0 | `Small Speakers` | EQ Preset |  |
| 0x007D72F0 | `Spoken Word` | EQ Preset |  |
| 0x007D72FC | `Treble Booster` | EQ Preset |  |
| 0x007D731C | `Vocal Booster` | EQ Preset |  |
| 0x00814444 | `Acoustic` | EQ Preset |  |
| 0x00814450 | `Bass Booster` | EQ Preset |  |
| 0x00814470 | `Classical` | EQ Preset |  |
| 0x0081447C | `Dance` | EQ Preset |  |
| 0x0081448C | `Electronic` | EQ Preset |  |
| 0x008144A0 | `Hip Hop` | EQ Preset |  |
| 0x008144A8 | `Jazz` | EQ Preset |  |
| 0x008144B0 | `Latin` | EQ Preset |  |
| 0x008144B8 | `Loudness` | EQ Preset |  |
| 0x008144C4 | `Lounge` | EQ Preset |  |
| 0x008144CC | `Piano` | EQ Preset |  |
| 0x008144DC | `Rock` | EQ Preset |  |
| 0x008144E4 | `Small Speakers` | EQ Preset |  |
| 0x008144F4 | `Spoken Word` | EQ Preset |  |
| 0x00814500 | `Treble Booster` | EQ Preset |  |
| 0x00814520 | `Vocal Booster` | EQ Preset |  |
| 0x0081D8B8 | `Acoustic` | EQ Preset |  |
| 0x0081D8C4 | `Bass Booster` | EQ Preset |  |
| 0x0081D8E4 | `Classical` | EQ Preset |  |
| 0x0081D8F0 | `Dance` | EQ Preset |  |
| 0x0081D900 | `Electronic` | EQ Preset |  |
| 0x0081D914 | `Hip Hop` | EQ Preset |  |
| 0x0081D91C | `Jazz` | EQ Preset |  |
| 0x0081D924 | `Latin` | EQ Preset |  |
| 0x0081D92C | `Loudness` | EQ Preset |  |
| 0x0081D938 | `Lounge` | EQ Preset |  |
| 0x0081D940 | `Piano` | EQ Preset |  |
| 0x0081D950 | `Rock` | EQ Preset |  |
| 0x0081D958 | `Small Speakers` | EQ Preset |  |
| 0x0081D968 | `Spoken Word` | EQ Preset |  |
| 0x0081D974 | `Treble Booster` | EQ Preset |  |
| 0x0081D994 | `Vocal Booster` | EQ Preset |  |
| 0x00827284 | `Acoustic` | EQ Preset |  |
| 0x008272B4 | `Dance` | EQ Preset |  |
| 0x008272C4 | `Electronic` | EQ Preset |  |
| 0x008272E0 | `Jazz` | EQ Preset |  |
| 0x008272E8 | `Latin` | EQ Preset |  |
| 0x008272F0 | `Loudness` | EQ Preset |  |
| 0x00827304 | `Piano` | EQ Preset |  |
| 0x00827314 | `Rock` | EQ Preset |  |
| 0x0083F1C0 | `Dance` | EQ Preset |  |
| 0x0083F1E8 | `Hip Hop` | EQ Preset |  |
| 0x0083F1F0 | `Jazz` | EQ Preset |  |
| 0x0083F200 | `Loudness` | EQ Preset |  |
| 0x0083F20C | `Lounge` | EQ Preset |  |
| 0x0083F214 | `Piano` | EQ Preset |  |
| 0x0083F224 | `Rock` | EQ Preset |  |
| 0x008486E4 | `Jazz` | EQ Preset |  |
| 0x008486EC | `Latin` | EQ Preset |  |
| 0x00848700 | `Lounge` | EQ Preset |  |
| 0x00848708 | `Piano` | EQ Preset |  |
| 0x00848718 | `Rock` | EQ Preset |  |
| 0x00851B38 | `Hip Hop` | EQ Preset |  |
| 0x00851B40 | `Jazz` | EQ Preset |  |
| 0x00851B5C | `Lounge` | EQ Preset |  |
| 0x00851B64 | `Piano` | EQ Preset |  |
| 0x00851B7C | `Rock` | EQ Preset |  |
| 0x0085BBF0 | `Latin` | EQ Preset |  |
| 0x0085BC1C | `Rock` | EQ Preset |  |
| 0x008655B0 | `Dance` | EQ Preset |  |
| 0x008655D4 | `Hip Hop` | EQ Preset |  |
| 0x008655DC | `Jazz` | EQ Preset |  |
| 0x008655EC | `Loudness` | EQ Preset |  |
| 0x008655F8 | `Lounge` | EQ Preset |  |
| 0x00865600 | `Piano` | EQ Preset |  |
| 0x00865610 | `Rock` | EQ Preset |  |
| 0x008703E4 | `Acoustic` | EQ Preset |  |
| 0x008703F0 | `Bass Booster` | EQ Preset |  |
| 0x00870410 | `Classical` | EQ Preset |  |
| 0x0087041C | `Dance` | EQ Preset |  |
| 0x0087042C | `Electronic` | EQ Preset |  |
| 0x00870440 | `Hip Hop` | EQ Preset |  |
| 0x00870448 | `Jazz` | EQ Preset |  |
| 0x00870450 | `Latin` | EQ Preset |  |
| 0x00870458 | `Loudness` | EQ Preset |  |
| 0x00870464 | `Lounge` | EQ Preset |  |
| 0x0087046C | `Piano` | EQ Preset |  |
| 0x0087047C | `Rock` | EQ Preset |  |
| 0x00870484 | `Small Speakers` | EQ Preset |  |
| 0x00870494 | `Spoken Word` | EQ Preset |  |
| 0x008704A0 | `Treble Booster` | EQ Preset |  |
| 0x008704C0 | `Vocal Booster` | EQ Preset |  |
| 0x0087B050 | `Acoustic` | EQ Preset |  |
| 0x0087B05C | `Bass Booster` | EQ Preset |  |
| 0x0087B07C | `Classical` | EQ Preset |  |
| 0x0087B088 | `Dance` | EQ Preset |  |
| 0x0087B098 | `Electronic` | EQ Preset |  |
| 0x0087B0AC | `Hip Hop` | EQ Preset |  |
| 0x0087B0B4 | `Jazz` | EQ Preset |  |
| 0x0087B0BC | `Latin` | EQ Preset |  |
| 0x0087B0C4 | `Loudness` | EQ Preset |  |
| 0x0087B0D0 | `Lounge` | EQ Preset |  |
| 0x0087B0D8 | `Piano` | EQ Preset |  |
| 0x0087B0E8 | `Rock` | EQ Preset |  |
| 0x0087B0F0 | `Small Speakers` | EQ Preset |  |
| 0x0087B100 | `Spoken Word` | EQ Preset |  |
| 0x0087B10C | `Treble Booster` | EQ Preset |  |
| 0x0087B12C | `Vocal Booster` | EQ Preset |  |
| 0x00884C50 | `Dance` | EQ Preset |  |
| 0x00884C84 | `Jazz` | EQ Preset |  |
| 0x00884C8C | `Latin` | EQ Preset |  |
| 0x00884C94 | `Loudness` | EQ Preset |  |
| 0x00884CA0 | `Lounge` | EQ Preset |  |
| 0x00884CA8 | `Piano` | EQ Preset |  |
| 0x00884CB8 | `Rock` | EQ Preset |  |
| 0x0088E06C | `Dance` | EQ Preset |  |
| 0x0088E098 | `Jazz` | EQ Preset |  |
| 0x0088E0A8 | `Loudness` | EQ Preset |  |
| 0x0088E0B4 | `Lounge` | EQ Preset |  |
| 0x0088E0BC | `Piano` | EQ Preset |  |
| 0x0088E0CC | `Rock` | EQ Preset |  |
| 0x00897734 | `Hip Hop` | EQ Preset |  |
| 0x0089773C | `Jazz` | EQ Preset |  |
| 0x00897760 | `Lounge` | EQ Preset |  |
| 0x00897778 | `Rock` | EQ Preset |  |
| 0x008A1204 | `Hip Hop` | EQ Preset |  |
| 0x008A120C | `Jazz` | EQ Preset |  |
| 0x008A1228 | `Lounge` | EQ Preset |  |
| 0x008A1230 | `Piano` | EQ Preset |  |
| 0x008A1240 | `Rock` | EQ Preset |  |
| 0x008B7C8C | `Acoustic` | EQ Preset |  |
| 0x008B7C98 | `Bass Booster` | EQ Preset |  |
| 0x008B7CB8 | `Classical` | EQ Preset |  |
| 0x008B7CC4 | `Dance` | EQ Preset |  |
| 0x008B7CD4 | `Electronic` | EQ Preset |  |
| 0x008B7CE8 | `Hip Hop` | EQ Preset |  |
| 0x008B7CF0 | `Jazz` | EQ Preset |  |
| 0x008B7CF8 | `Latin` | EQ Preset |  |
| 0x008B7D00 | `Loudness` | EQ Preset |  |
| 0x008B7D0C | `Lounge` | EQ Preset |  |
| 0x008B7D14 | `Piano` | EQ Preset |  |
| 0x008B7D24 | `Rock` | EQ Preset |  |
| 0x008B7D2C | `Small Speakers` | EQ Preset |  |
| 0x008B7D3C | `Spoken Word` | EQ Preset |  |
| 0x008B7D48 | `Treble Booster` | EQ Preset |  |
| 0x008B7D68 | `Vocal Booster` | EQ Preset |  |
| 0x008C130C | `Hip Hop` | EQ Preset |  |
| 0x008C1318 | `Latin` | EQ Preset |  |
| 0x008C1350 | `Rock` | EQ Preset |  |
| 0x008CAACC | `Acoustic` | EQ Preset |  |
| 0x008CAAD8 | `Bass Booster` | EQ Preset |  |
| 0x008CAAF8 | `Classical` | EQ Preset |  |
| 0x008CAB04 | `Dance` | EQ Preset |  |
| 0x008CAB14 | `Electronic` | EQ Preset |  |
| 0x008CAB28 | `Hip Hop` | EQ Preset |  |
| 0x008CAB30 | `Jazz` | EQ Preset |  |
| 0x008CAB38 | `Latin` | EQ Preset |  |
| 0x008CAB40 | `Loudness` | EQ Preset |  |
| 0x008CAB4C | `Lounge` | EQ Preset |  |
| 0x008CAB54 | `Piano` | EQ Preset |  |
| 0x008CAB64 | `Rock` | EQ Preset |  |
| 0x008CAB6C | `Small Speakers` | EQ Preset |  |
| 0x008CAB7C | `Spoken Word` | EQ Preset |  |
| 0x008CAB88 | `Treble Booster` | EQ Preset |  |
| 0x008CABA8 | `Vocal Booster` | EQ Preset |  |
| 0x008D41BC | `Acoustic` | EQ Preset |  |
| 0x008D41C8 | `Bass Booster` | EQ Preset |  |
| 0x008D41E8 | `Classical` | EQ Preset |  |
| 0x008D41F4 | `Dance` | EQ Preset |  |
| 0x008D4204 | `Electronic` | EQ Preset |  |
| 0x008D4218 | `Hip Hop` | EQ Preset |  |
| 0x008D4220 | `Jazz` | EQ Preset |  |
| 0x008D4228 | `Latin` | EQ Preset |  |
| 0x008D4230 | `Loudness` | EQ Preset |  |
| 0x008D423C | `Lounge` | EQ Preset |  |
| 0x008D4244 | `Piano` | EQ Preset |  |
| 0x008D4254 | `Rock` | EQ Preset |  |
| 0x008D425C | `Small Speakers` | EQ Preset |  |
| 0x008D426C | `Spoken Word` | EQ Preset |  |
| 0x008D4278 | `Treble Booster` | EQ Preset |  |
| 0x008D4298 | `Vocal Booster` | EQ Preset |  |

---
