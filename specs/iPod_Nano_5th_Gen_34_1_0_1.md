# iPod Nano 5th Gen - RetailOS 1.0.1 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.0.1 |
| **IPSW** | iPod_1.0.1_34A10006.ipsw |
| **Device** | iPod Nano 5th Gen (2009, 8/16GB NAND, Click Wheel, Camera, Pedometer, FM Radio (First Release)) |
| **UpdaterFamilyID** | 34 |
| **Binary Size** | 7,276,688 bytes (6.94 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 7,274,640 bytes |
| **Total Strings (>=4)** | 90,635 |
| **Function Prologues** | 32,149 (ARM: 2,080, Thumb: 30,069) |
| **DRAM References** | 36,381 |
| **Peripheral Refs** | 8,463 |
| **Build** | N33FirmwareWin-206 |
| **SoC** | S5L8730 |
| **Architecture** | ARM Cortex-A8 (ARMv7) |
| **Codename** | N33 |
| **DFU PID** | 0x1231 |
| **SHA-256** | `ea4bad8dc8c8c57ea3144f9615b92d6865ab34bec267229b41e7b12cf6b491f4` |

---

## 1. Controllers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0004E154 | `TCSportTimer` | Known | Controller |
| 0x0004E16C | `TCSportTimerMenu` | Known | Controller |
| 0x0004E188 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x0004E1AC | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x0004EA78 | `TCFirewireUnsupported` | Known | Controller |
| 0x0004EBA0 | `TCNotesDispatcher` | Known | Controller |
| 0x0004EBBC | `TCNotesLoading` | Known | Controller |
| 0x0004EBD4 | `TCNotesList` | Known | Controller |
| 0x0004EBE8 | `TCNotesContents` | Known | Controller |
| 0x00050B2C | `TCDemoMode` | Known | Controller |
| 0x000540CC | `TCCamera` | Known | Controller |
| 0x000540E0 | `TCCameraInitial` | Known | Controller |
| 0x000540F8 | `TCCameraLocalMediaList` | Known | Controller |
| 0x00054118 | `TCCameraAllVideosList` | Known | Controller |
| 0x00054138 | `TCCameraDeleteAllDialog` | Known | Controller |
| 0x00054158 | `TCCameraDeleteDialog` | Known | Controller |
| 0x000549D0 | `TCPhotosDeleteAllDialog` | Known | Controller |
| 0x000549F0 | `TCPhotosDeleteDialog` | Known | Controller |
| 0x00055068 | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x00055094 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x000550C0 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x000550E8 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x00055114 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x0005513C | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x00055168 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x000556A0 | `TCRemoteUI` | Known | Controller |
| 0x000556B4 | `TCUnsupported` | Known | Controller |
| 0x00055AB0 | `TCSpeakers` | Known | Controller |
| 0x00055AC4 | `TCEQSetting` | Known | Controller |
| 0x00057048 | `TCVoiceMemos` | Known | Controller |
| 0x00057060 | `TCVoiceMemosIdle` | Known | Controller |
| 0x0005707C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x0005709C | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x000570BC | `TCVoiceMemosLabelSelectMenu` | Known | Controller |
| 0x000570E0 | `TCVoiceMemosLoading` | Known | Controller |
| 0x000570FC | `TCVoiceMemosTimedStatus` | Known | Controller |
| 0x00057694 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x000576BC | `TCSettings_MainMenu` | Known | Controller |
| 0x000576D8 | `TCSettings_MusicMenu` | Known | Controller |
| 0x000576F8 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x00057718 | `TCSettings_VolumeLimit_Dialogue` | Known | Controller |
| 0x00057740 | `TCSettings_Brightness` | Known | Controller |
| 0x00057760 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x00057784 | `TCSettings_EQ` | Known | Controller |
| 0x0005779C | `TCSettings_RadioRegions` | Known | Controller |
| 0x000577BC | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x000577E0 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x00057804 | `TCDateTimeScreen` | Known | Controller |
| 0x00057820 | `TCTimeZoneScreen` | Known | Controller |
| 0x0005783C | `TCAddressViewerLoadingScreenCntlr` | Known | Controller |
| 0x00057868 | `TCAddressViewerNoContactsCntlr` | Known | Controller |
| 0x000578B4 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x000578DC | `TCAboutCntlr` | Known | Controller |
| 0x000578F4 | `TCSettings_Language` | Known | Controller |
| 0x000588AC | `TCAddressViewerMainMenu` | Known | Controller |
| 0x000588CC | `TCAddressViewerDetails` | Known | Controller |
| 0x000588EC | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x00058910 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x00058938 | `TCAddressViewerContactGroups` | Known | Controller |
| 0x00059E74 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x00059E98 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x0005C148 | `TC_LockDialog` | Known | Controller |
| 0x0005C160 | `TC_LockScreen` | Known | Controller |
| 0x0005C178 | `TC_LockediPod` | Known | Controller |
| 0x0005C190 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x0005C1B4 | `TCLockChosenDispatcher` | Known | Controller |
| 0x0005C1D4 | `TCResetCombinationChosenDispatcher` | Known | Controller |
| 0x0005C200 | `TCLockAppMenu` | Known | Controller |
| 0x0005C670 | `TCClock` | Known | Controller |
| 0x0005C680 | `TCClockCityMenu` | Known | Controller |
| 0x0005C698 | `TCClockRegionMenu` | Known | Controller |
| 0x0005C6B4 | `TCAlarmMenu` | Known | Controller |
| 0x0005C6C8 | `TCSleepTimerMenu` | Known | Controller |
| 0x0005C6E4 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x0005C704 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x0005C72C | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x0005C750 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x0005C774 | `TCAlarmDatePicker` | Known | Controller |
| 0x0005C790 | `TCAlarmTriggered` | Known | Controller |
| 0x0006F4FC | `TSilverCntlr` | Known | Controller |
| 0x0006F514 | `TCExtrasMenu` | Known | Controller |
| 0x0006F52C | `TCGameScreen` | Known | Controller |
| 0x0006F544 | `TCGameControls` | Known | Controller |
| 0x0006F55C | `TCGamesMenu` | Known | Controller |
| 0x0006F570 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x0006F598 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x0006F5C0 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x0006F5EC | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x0006F610 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x0006F638 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x0006F660 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x0006F688 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x0006F6B0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0006F6D8 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0006F708 | `TSilverMediaListCntlr_iTunesU` | Known | Controller |
| 0x0006F730 | `TSilverMediaListCntlr_iTunesUEpisodes` | Known | Controller |
| 0x0006F760 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x0006F78C | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x0006F7BC | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x0006F7E4 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x0006F80C | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x0006F838 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x0006F860 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x0006F888 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0006F8B8 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x0006F8E8 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x0006FA14 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x0006FA44 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x0006FA6C | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x0006FA94 | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x0006FAC0 | `TCRentalNotification` | Known | Controller |
| 0x0006FAE0 | `TCRentalInfo` | Known | Controller |
| 0x0006FAF8 | `TCRentalConfirmDelete` | Known | Controller |
| 0x0006FB18 | `TCRentalDispatcher` | Known | Controller |
| 0x0006FB34 | `TSilverOverlayCntlr` | Known | Controller |
| 0x0006FBCC | `TSilverGlobalCntlr` | Known | Controller |
| 0x0006FBE8 | `TCGlobalCoverFlowEntry` | Known | Controller |
| 0x0008F20E | `TC3Fd` | Known | Controller |
| 0x000E99B0 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x00460418 | `TCClock` | Known | Controller |
| 0x00460420 | `TCAlarmTriggered` | Known | Controller |
| 0x00460444 | `TSilverCntlr` | Known | Controller |
| 0x00460454 | `TCClockRegionMenu` | Known | Controller |
| 0x00460468 | `TCClockCityMenu` | Known | Controller |
| 0x00460478 | `TCAlarmMenu` | Known | Controller |
| 0x00460484 | `TCSleepTimerMenu` | Known | Controller |
| 0x00460498 | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x004604B0 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x004604D0 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x004604EC | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x00460508 | `TCAlarmDatePicker` | Known | Controller |
| 0x00460630 | `TSilverCntlr` | Known | Controller |
| 0x00460640 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x00460660 | `TCSettings_Brightness` | Known | Controller |
| 0x00460678 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x00460694 | `TCSettings_RadioRegions` | Known | Controller |
| 0x004606AC | `TCSettings_EQ` | Known | Controller |
| 0x004606BC | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x004606D8 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x004606F8 | `TCAboutCntlr` | Known | Controller |
| 0x00460708 | `TCSettings_Language` | Known | Controller |
| 0x0046071C | `TCSettings_MainMenu` | Known | Controller |
| 0x00460730 | `TCSettings_MusicMenu` | Known | Controller |
| 0x00460748 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x00460760 | `TCSettings_VolumeLimit_Dialogue` | Known | Controller |
| 0x00460780 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x0046079C | `TCDateTimeScreen` | Known | Controller |
| 0x004607B0 | `TCTimeZoneScreen` | Known | Controller |
| 0x004607C4 | `TCAddressViewerLoadingScreenCntlr` | Known | Controller |
| 0x004607E8 | `TCAddressViewerNoContactsCntlr` | Known | Controller |
| 0x0046357C | `TSilverCntlr` | Known | Controller |
| 0x0046358C | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x004635B0 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x004635D4 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x004635F4 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x00463618 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x00463638 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x0046365C | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x004637FC | `TSilverCntlr` | Known | Controller |
| 0x0046380C | `TCCameraInitial` | Known | Controller |
| 0x0046381C | `TCCamera` | Known | Controller |
| 0x00463828 | `TCCameraLocalMediaList` | Known | Controller |
| 0x00463840 | `TCCameraAllVideosList` | Known | Controller |
| 0x00463858 | `TCCameraDeleteAllDialog` | Known | Controller |
| 0x00463870 | `TCCameraDeleteDialog` | Known | Controller |
| 0x00464690 | `TSilverCntlr` | Known | Controller |
| 0x004646A0 | `TC_LockDialog` | Known | Controller |
| 0x004646B0 | `TC_LockScreen` | Known | Controller |
| 0x004646C0 | `TC_LockediPod` | Known | Controller |
| 0x004646D0 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x004646EC | `TCLockChosenDispatcher` | Known | Controller |
| 0x00464704 | `TCResetCombinationChosenDispatcher` | Known | Controller |
| 0x00464728 | `TCLockAppMenu` | Known | Controller |
| 0x00464748 | `TSilverCntlr` | Known | Controller |
| 0x00464758 | `TCFirewireUnsupported` | Known | Controller |
| 0x004649C8 | `TCRemoteUI` | Known | Controller |
| 0x004649D4 | `TCUnsupported` | Known | Controller |
| 0x00464A68 | `TSilverCntlr` | Known | Controller |
| 0x00465038 | `TSilverCntlr` | Known | Controller |
| 0x00465264 | `TCDemoMode` | Known | Controller |
| 0x00465368 | `TSilverCntlr` | Known | Controller |
| 0x00465378 | `TCVoiceMemosIdle` | Known | Controller |
| 0x0046538C | `TCVoiceMemos` | Known | Controller |
| 0x0046539C | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x004653B4 | `TCVoiceMemosLabelSelectMenu` | Known | Controller |
| 0x004653D0 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x004653E8 | `TCVoiceMemosLoading` | Known | Controller |
| 0x004653FC | `TCVoiceMemosTimedStatus` | Known | Controller |
| 0x00466DD0 | `TSilverCntlr` | Known | Controller |
| 0x0047C00C | `TSilverCntlr` | Known | Controller |
| 0x0047C01C | `TCAddressViewerMainMenu` | Known | Controller |
| 0x0047C034 | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x0047C050 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x0047C070 | `TCAddressViewerDetails` | Known | Controller |
| 0x0047C088 | `TCAddressViewerContactGroups` | Known | Controller |
| 0x0047C5E8 | `TSilverCalendarCntlr_Base` | Known | Controller |
| 0x0047C8BC | `TCCameraInitial` | Known | Controller |
| 0x0047C8CC | `TCCamera` | Known | Controller |
| 0x0047C8D8 | `TCCameraMediaList_Base` | Known | Controller |
| 0x0047C900 | `TCCameraLocalMediaList` | Known | Controller |
| 0x0047C918 | `TCCameraAllVideosList` | Known | Controller |
| 0x0047C930 | `TCCameraDeleteAllDialog` | Known | Controller |
| 0x0047C948 | `TCCameraDeleteDialog` | Known | Controller |
| 0x0047CA70 | `TSilverMainMediaListCntlr_Base` | Known | Controller |
| 0x0047CAA0 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x0047CAC0 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x0047CAE0 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x0047CB14 | `TSilverCntlr` | Known | Controller |
| 0x0047CB24 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x0047CB40 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x0047CB60 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0047CB80 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0047CBA8 | `TSilverMediaListCntlr_iTunesU` | Known | Controller |
| 0x0047CBC8 | `TSilverMediaListCntlr_iTunesUEpisodes` | Known | Controller |
| 0x0047CBF0 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x0047CC14 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x0047CC3C | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x0047CC5C | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x0047CC7C | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x0047CC9C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x0047CCBC | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0047CCE4 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x0047CD0C | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x0047CD2C | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x0047CD60 | `TSilverCntlr` | Known | Controller |
| 0x0047CEDC | `TCNotesDispatcher` | Known | Controller |
| 0x0047CEF0 | `TCNotesLoading` | Known | Controller |
| 0x0047CF00 | `TCNotesList` | Known | Controller |
| 0x0047CF0C | `TCNotesContents` | Known | Controller |
| 0x0047CF1C | `TSilverCntlr` | Known | Controller |
| 0x0047D024 | `TCPhotosDeleteAllDialog` | Known | Controller |
| 0x0047D03C | `TCPhotosDeleteDialog` | Known | Controller |
| 0x0047D064 | `TSilverCntlr` | Known | Controller |
| 0x0047D138 | `TSilverCntlr` | Known | Controller |
| 0x0047D238 | `TSilverCntlrTestAppCntlr` | Known | Controller |
| 0x0047D254 | `TSilverCntlrTestCntlr` | Known | Controller |
| 0x0047D26C | `TCSpeakers` | Known | Controller |
| 0x0047D278 | `TCEQSetting` | Known | Controller |
| 0x0047D2B0 | `TSilverCntlr` | Known | Controller |
| 0x0047D2C0 | `TCSportTimer` | Known | Controller |
| 0x0047D2D0 | `TCSportTimerMenu` | Known | Controller |
| 0x0047D2E4 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x0047D300 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x0047D36C | `TSilverCntlr` | Known | Controller |
| 0x0047D37C | `TCVoiceMemosIdle` | Known | Controller |
| 0x0047D390 | `TCVoiceMemos` | Known | Controller |
| 0x0047D3A0 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x0047D3B8 | `TCVoiceMemosLabelSelectMenu` | Known | Controller |
| 0x0047D3D4 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x0047D3EC | `TCVoiceMemosLoading` | Known | Controller |
| 0x0047D400 | `TCVoiceMemosTimedStatus` | Known | Controller |
| 0x0047D458 | `TSilverCntlr` | Known | Controller |
| 0x0047D478 | `TCExtrasMenu` | Known | Controller |
| 0x0047D488 | `TCGamesMenu` | Known | Controller |
| 0x0047D494 | `TCGameControls` | Known | Controller |
| 0x0047D4A4 | `TCGameScreen` | Known | Controller |
| 0x0047D4B4 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x0047D4D0 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x0047D4F0 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x0047D510 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x0047D538 | `TSilverMediaListCntlr_iTunesU` | Known | Controller |
| 0x0047D558 | `TSilverMediaListCntlr_iTunesUEpisodes` | Known | Controller |
| 0x0047D580 | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x0047D5A4 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x0047D5CC | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x0047D5EC | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x0047D60C | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x0047D62C | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x0047D64C | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x0047D674 | `TSilverMediaListCntlr_PlaylistChooser` | Known | Controller |
| 0x0047D69C | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x0047D6BC | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x0047D6E0 | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x0047D700 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x0047D720 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x0047D744 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x0047D764 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x0047D784 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x0047D7A8 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x0047D7C8 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x0047D7F0 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x0047D81C | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x0047D83C | `TCRentalNotification` | Known | Controller |
| 0x0047D854 | `TCRentalInfo` | Known | Controller |
| 0x0047D864 | `TCRentalConfirmDelete` | Known | Controller |
| 0x0047D87C | `TCRentalDispatcher` | Known | Controller |
| 0x0047D890 | `TSilverGlobalCntlr` | Known | Controller |
| 0x0047D8A4 | `TCGlobalCoverFlowEntry` | Known | Controller |
| 0x0047D8BC | `TSilverOverlayCntlr` | Known | Controller |
| 0x0048BDE0 | `TCNotesDispatcher` | Known | Controller |
| 0x0048BDF4 | `TCNotesLoading` | Known | Controller |
| 0x0048BE04 | `TCNotesBase` | Known | Controller |
| 0x0048BE20 | `TCNotesList` | Known | Controller |
| 0x0048BE2C | `TCNotesContents` | Known | Controller |
| 0x0056DCA4 | `TCCameraInitial_InitialLayoutIsAppNotInitialized` | Known | Controller |
| 0x0056DD69 | `TCCameraInitial_InitialLayoutIsActive` | Known | Controller |
| 0x0056DDD5 | `TCCameraInitial_InitialLayoutIsDiskFull` | Known | Controller |
| 0x0056DED7 | `TCCameraMediaList_Base_DoDeleteAll` | Known | Controller |
| 0x0056DEFA | `TCCameraMediaList_Base_DoDeleteItem` | Known | Controller |
| 0x00571684 | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |
| 0x005AD90C | `TSilverGlobalCntlr` | Known | Controller |
| 0x005AD920 | `TSilverCntlr` | Known | Controller |
| 0x005AD930 | `TSilverCntlr` | Known | Controller |
| 0x005AD940 | `TSilverCntlr` | Known | Controller |
| 0x005AD950 | `TSilverCntlr` | Known | Controller |
| 0x005AD960 | `TCGlobalCoverFlowEntry` | Known | Controller |
| 0x005AD978 | `TSilverCntlr` | Known | Controller |
| 0x005AD988 | `TCGlobalCoverFlowEntry` | Known | Controller |
| 0x005AD9A0 | `TSilverCntlr` | Known | Controller |
| 0x005AD9B0 | `TSilverCntlr` | Known | Controller |
| 0x005AD9D0 | `TSilverCntlr` | Known | Controller |
| 0x005AD9E0 | `TSilverCntlr` | Known | Controller |
| 0x005AD9F0 | `TSilverCntlr` | Known | Controller |
| 0x005ADA00 | `TSilverCntlr` | Known | Controller |
| 0x005ADA10 | `TCGlobalCoverFlowEntry` | Known | Controller |
| 0x005ADA28 | `TSilverCntlr` | Known | Controller |
| 0x005BC1C4 | `TCFirewireUnsupported` | Known | Controller |
| 0x005BC1F4 | `TCExtrasMenu` | Known | Controller |
| 0x005BC22C | `TCAddressViewerNoContactsCntlr` | Known | Controller |
| 0x005BC24C | `TCAddressViewerContactGroups` | Known | Controller |
| 0x005BC26C | `TCAddressViewerMainMenu` | Known | Controller |
| 0x005BC284 | `TCAddressViewerDetails` | Known | Controller |
| 0x005BC29C | `TCAddressViewerPartialLoad` | Known | Controller |
| 0x005BC2B8 | `TCAddressViewerLoadingScreenCntlr` | Known | Controller |
| 0x005BC2DC | `TCAlarmMenu` | Known | Controller |
| 0x005BC2E8 | `TCSleepTimerMenu` | Known | Controller |
| 0x005BC2FC | `TCAlarmPropertiesMenu` | Known | Controller |
| 0x005BC314 | `TCAlarmPropertiesFrequencyMenu` | Known | Controller |
| 0x005BC334 | `TCAlarmPropertiesLabelMenu` | Known | Controller |
| 0x005BC350 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x005BC36C | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x005BC388 | `TCAlarmPropertiesSoundMenu` | Known | Controller |
| 0x005BC3A4 | `TCAlarmDatePicker` | Known | Controller |
| 0x005BC3B8 | `TCAlarmDatePicker` | Known | Controller |
| 0x005BC3CC | `TCAlarmTriggered` | Known | Controller |
| 0x005BC3E0 | `TSilverCalendarCntlr_Alarm` | Known | Controller |
| 0x005BC3FC | `TSilverCalendarCntlr_CalendarMenu` | Known | Controller |
| 0x005BC420 | `TSilverCalendarCntlr_MonthViewer` | Known | Controller |
| 0x005BC444 | `TSilverCalendarCntlr_DayViewer` | Known | Controller |
| 0x005BC464 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x005BC488 | `TSilverCalendarCntlr_ToDoList` | Known | Controller |
| 0x005BC4A8 | `TSilverCalendarCntlr_EventViewer` | Known | Controller |
| 0x005BC4CC | `TSilverCntlr` | Known | Controller |
| 0x005BC4F8 | `TCCamera` | Known | Controller |
| 0x005BC504 | `TCCameraLocalMediaList` | Known | Controller |
| 0x005BC51C | `TCCameraAllVideosList` | Known | Controller |
| 0x005BC54C | `TCCameraDeleteAllDialog` | Known | Controller |
| 0x005BC564 | `TCCameraDeleteDialog` | Known | Controller |
| 0x005BC57C | `TCCameraDeleteDialog` | Known | Controller |
| 0x005BC594 | `TCCameraDeleteDialog` | Known | Controller |
| 0x005BC5AC | `TCClockRegionMenu` | Known | Controller |
| 0x005BC5C0 | `TCClockCityMenu` | Known | Controller |
| 0x005BC5D0 | `TCClock` | Known | Controller |
| 0x005BC618 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x005BC638 | `TCDateTimeScreen` | Known | Controller |
| 0x005BC64C | `TCDateTimeScreen` | Known | Controller |
| 0x005BC660 | `TCTimeZoneScreen` | Known | Controller |
| 0x005BC674 | `TCDemoMode` | Known | Controller |
| 0x005BC698 | `TCGamesMenu` | Known | Controller |
| 0x005BC6A4 | `TCGameControls` | Known | Controller |
| 0x005BC6B4 | `TCGameScreen` | Known | Controller |
| 0x005BC710 | `TCLockAppMenu` | Known | Controller |
| 0x005BC720 | `TC_LockediPod` | Known | Controller |
| 0x005BC730 | `TC_LockScreen` | Known | Controller |
| 0x005BC740 | `TC_VolumeLimitLockScreen` | Known | Controller |
| 0x005BC75C | `TSilverMainMediaListCntlr_Main` | Known | Controller |
| 0x005BC77C | `TSilverCntlr` | Known | Controller |
| 0x005BC78C | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x005BC7AC | `TSilverMediaListCntlr_Audiobooks` | Known | Controller |
| 0x005BC7D0 | `TSilverMediaListCntlr_AudiobookChapters` | Known | Controller |
| 0x005BC7F8 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x005BC814 | `TSilverMediaListCntlr_Albums` | Known | Controller |
| 0x005BC834 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x005BC854 | `TSilverMediaListCntlr_Genres` | Known | Controller |
| 0x005BC874 | `TSilverMediaListCntlr_Genius` | Known | Controller |
| 0x005BC894 | `TSilverMediaListCntlr_GeniusMixes` | Known | Controller |
| 0x005BC8B8 | `TSilverMediaListCntlr_Composers` | Known | Controller |
| 0x005BC8D8 | `TSilverMainMediaListCntlr_Music` | Known | Controller |
| 0x005BC8F8 | `TSilverCntlr` | Known | Controller |
| 0x005BC908 | `TSilverMediaListCntlr_Podcasts` | Known | Controller |
| 0x005BC928 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | Controller |
| 0x005BC950 | `TSilverMediaListCntlr_iTunesU` | Known | Controller |
| 0x005BC970 | `TSilverMediaListCntlr_iTunesUEpisodes` | Known | Controller |
| 0x005BC998 | `TSilverMediaListCntlr_Playlists` | Known | Controller |
| 0x005BC9B8 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x005BC9E0 | `TSilverMediaListCntlr_NestedPlaylists` | Known | Controller |
| 0x005BCA20 | `TSilverCntlr` | Known | Controller |
| 0x005BCA30 | `TSilverMainMediaListCntlr_Videos` | Known | Controller |
| 0x005BCA54 | `TSilverMediaListCntlr_Movies` | Known | Controller |
| 0x005BCA74 | `TSilverMediaListCntlr_Rentals` | Known | Controller |
| 0x005BCA94 | `TSilverMediaListCntlr_TVShows` | Known | Controller |
| 0x005BCAB4 | `TSilverMediaListCntlr_TVSeasons` | Known | Controller |
| 0x005BCAD4 | `TSilverMediaListCntlr_TVEpisodes` | Known | Controller |
| 0x005BCAF8 | `TSilverMediaListCntlr_Artists` | Known | Controller |
| 0x005BCB18 | `TSilverMediaListCntlr_Songs` | Known | Controller |
| 0x005BCB34 | `TSilverMediaListCntlr_VideoPlaylists` | Known | Controller |
| 0x005BCB5C | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | Controller |
| 0x005BCB88 | `TSilverSettingsVideoCntlr` | Known | Controller |
| 0x005BCD04 | `TCRentalInfo` | Known | Controller |
| 0x005BCD14 | `TCRentalConfirmDelete` | Known | Controller |
| 0x005BCD2C | `TSilverCntlr` | Known | Controller |
| 0x005BCD3C | `TCRentalNotification` | Known | Controller |
| 0x005BCD54 | `TCRentalNotification` | Known | Controller |
| 0x005BCD6C | `TCRentalNotification` | Known | Controller |
| 0x005BCD84 | `TCNotesLoading` | Known | Controller |
| 0x005BCD94 | `TCNotesList` | Known | Controller |
| 0x005BCDA0 | `TCNotesList` | Known | Controller |
| 0x005BCDAC | `TCNotesContents` | Known | Controller |
| 0x005BCDBC | `TCNotesContents` | Known | Controller |
| 0x005BCDCC | `TCNotesContents` | Known | Controller |
| 0x005BCDDC | `TCNotesContents` | Known | Controller |
| 0x005BCE3C | `TSilverCntlr` | Known | Controller |
| 0x005BCF98 | `TCPhotosDeleteAllDialog` | Known | Controller |
| 0x005BCFB0 | `TCPhotosDeleteDialog` | Known | Controller |
| 0x005BCFC8 | `TCPhotosDeleteAllDialog` | Known | Controller |
| 0x005BCFE0 | `TCPhotosDeleteDialog` | Known | Controller |
| 0x005BD10C | `TCRemoteUI` | Known | Controller |
| 0x005BD118 | `TCUnsupported` | Known | Controller |
| 0x005BD144 | `TCAboutCntlr` | Known | Controller |
| 0x005BD154 | `TCAboutCntlr` | Known | Controller |
| 0x005BD164 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x005BD184 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x005BD1A4 | `TSilverSettingsMenuListCntlr` | Known | Controller |
| 0x005BD1C4 | `TCSettings_MainMenu` | Known | Controller |
| 0x005BD1D8 | `TCSettings_MusicMenu` | Known | Controller |
| 0x005BD1F0 | `TCShakeAdjust_Intensity` | Known | Controller |
| 0x005BD208 | `TCShakeAdjust_Duration` | Known | Controller |
| 0x005BD220 | `TCSettings_VolumeLimit` | Known | Controller |
| 0x005BD238 | `TCSettings_VolumeLimit_Dialogue` | Known | Controller |
| 0x005BD258 | `TCSettings_Brightness` | Known | Controller |
| 0x005BD270 | `TCSettings_BacklightTimer` | Known | Controller |
| 0x005BD28C | `TCSettings_RadioRegions` | Known | Controller |
| 0x005BD2A4 | `TCSettings_EQ` | Known | Controller |
| 0x005BD2B4 | `TCSettings_Language` | Known | Controller |
| 0x005BD2C8 | `TSilverCntlr` | Known | Controller |
| 0x005BD2D8 | `TCSettings_AdjustScrollingCntlr` | Known | Controller |
| 0x005BD2F8 | `TCSettings_ResetAllSettings` | Known | Controller |
| 0x005BD314 | `TCSettings_MainMenu` | Known | Controller |
| 0x005BD328 | `TCSettings_MusicMenu` | Known | Controller |
| 0x005BD340 | `TCSportTimer` | Known | Controller |
| 0x005BD350 | `TCSportTimerMenu` | Known | Controller |
| 0x005BD364 | `TCSportTimerSessionScreen` | Known | Controller |
| 0x005BD964 | `TSilverCntlr` | Known | Controller |
| 0x005BDAB0 | `TSilverCntlr` | Known | Controller |
| 0x005BDAC0 | `TSilverCntlr` | Known | Controller |
| 0x005BDB1C | `TSilverCntlr` | Known | Controller |
| 0x005BDB5C | `TSilverCntlr` | Known | Controller |
| 0x005BDB6C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BDB84 | `TCVoiceMemosIdle` | Known | Controller |
| 0x005BDB98 | `TCVoiceMemos` | Known | Controller |
| 0x005BDBC0 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BDBD8 | `TCVoiceMemosContextMenu` | Known | Controller |
| 0x005BDBF0 | `TCVoiceMemosLabelSelectMenu` | Known | Controller |
| 0x005BDC0C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BDC24 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BDC3C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BDC54 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BDC6C | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BDC84 | `TCVoiceMemosMainMenu` | Known | Controller |
| 0x005BDC9C | `TCVoiceMemosLoading` | Known | Controller |
| 0x005BDCB0 | `TCVoiceMemosTimedStatus` | Known | Controller |
| 0x005BDCD8 | `TCSpeakers` | Known | Controller |
| 0x005BDCE4 | `TCEQSetting` | Known | Controller |
| 0x00634910 | `TCAddressViewerMainDispatcher` | Known | Controller |
| 0x00634BEC | `TCNotesDispatcher` | Known | Controller |
| 0x00634CC4 | `TCSportTimerChosenDispatcher` | Known | Controller |
| 0x0063555C | `TSilverCalendarCntlr_ToDoDispatcher` | Known | Controller |
| 0x00635A40 | `TCCameraInitial` | Known | Controller |
| 0x00636C44 | `TCLockChosenDispatcher` | Known | Controller |
| 0x00636C78 | `TCResetCombinationChosenDispatcher` | Known | Controller |
| 0x0063A720 | `TCRentalDispatcher` | Known | Controller |
| 0x0063A87C | `TSilverMediaListCntlr_MixedVideoTracks_Screen_WithSubInfo` | Known | Controller |

---

## 2. Settings (Toggle/Show/TC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000ADB9C | `ToggleSetting_PreviewPanel` | Known | User setting |
| 0x000AF0E8 | `ToggleSetting_Repeat` | Known | User setting |
| 0x000AF104 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x000AF11C | `ToggleSetting_TVOut` | Known | User setting |
| 0x000AF130 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x000B5DF0 | `ToggleSetting_Audiobook` | Known | User setting |
| 0x000B5E0C | `ToggleSetting_Shuffle` | Known | User setting |
| 0x000B5E24 | `ToggleSetting_Repeat` | Known | User setting |
| 0x000B5E3C | `ToggleSetting_SortBy` | Known | User setting |
| 0x000B5E54 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x000B5E70 | `ToggleSetting_Clicker` | Known | User setting |
| 0x000B5E88 | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x000B5EA8 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x000B5EC4 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x000B5EE0 | `ToggleSetting_EnergySaver` | Known | User setting |
| 0x000B5EFC | `ToggleSetting_Crossfade` | Known | User setting |
| 0x000B5F14 | `ToggleSetting_FontSize` | Known | User setting |
| 0x000B5F2C | `ToggleSetting_Shake` | Known | User setting |
| 0x000B5F40 | `ToggleSetting_VoiceFeedback` | Known | User setting |
| 0x000B5F5C | `ToggleSetting_Rotate` | Known | User setting |
| 0x000B5F74 | `ShowSetting_About` | Known | User setting |
| 0x000B5F88 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x000B5FA0 | `ShowSetting_DateAndTime` | Known | User setting |
| 0x000B5FB8 | `ShowSetting_Legal` | Known | User setting |
| 0x0063559C | `ToggleSetting_Alarm` | Known | User setting |
| 0x00636784 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x006367BC | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x0063BC1C | `ToggleSetting_Repeat` | Known | User setting |
| 0x0063BC50 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x0063BCE4 | `ToggleSetting_TVOut` | Known | User setting |
| 0x0063BD14 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x0063D8C8 | `ShowSetting_About` | Known | User setting |
| 0x0063DBF8 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x0063DC6C | `ToggleSetting_Crossfade` | Known | User setting |
| 0x0063DCA0 | `ToggleSetting_Audiobook` | Known | User setting |
| 0x0063DCD4 | `ToggleSetting_Shake` | Known | User setting |
| 0x0063DD04 | `ToggleSetting_EnergySaver` | Known | User setting |
| 0x0063DDA4 | `ToggleSetting_Clicker` | Known | User setting |
| 0x0063DDD8 | `ToggleSetting_Rotate` | Known | User setting |
| 0x0063DE0C | `ToggleSetting_VoiceFeedback` | Known | User setting |
| 0x0063DEE4 | `ToggleSetting_FontSize` | Known | User setting |
| 0x0063DFE0 | `ToggleSetting_SortBy` | Known | User setting |
| 0x0063E014 | `ToggleSetting_PreviewPanel` | Known | User setting |

---

## 3. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0011AF2C | `Channel UnitTests` | Hidden | Developer Tool |
| 0x0051CF43 | `10TCDemoMode` | Hidden | Demo/Retail Mode |
| 0x0051D2E0 | `12TUnitTestApp` | Hidden | Developer Tool |
| 0x0051F37A | `27TSilverCntlrTransitionAddonI10TCDemoModeE` | Hidden | Demo/Retail Mode |
| 0x005BDB2C | `TUnitTestSuiteCntlr` | Hidden | Developer Tool |
| 0x005BDB40 | `TUnitTestSuiteTestsCntlr` | Hidden | Developer Tool |
| 0x0062A760 | `DemoMode` | Hidden | Demo/Retail Mode |

---

## 4. Audio System (MeCCA Framework)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00036460 | `AudioCodecs` | Known | Audio system |
| 0x00065F3C | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x0015D56C | `MeCCA_VdRecBufferMgr` | Known | Audio system |
| 0x0017DF08 | `MeCCA_GlobalBMHeap` | Known | Audio system |
| 0x001BFF17 | `"MeCCADecode` | Known | Audio system |
| 0x001C0684 | `MeCCAVideoDecode` | Known | Audio system |
| 0x00328F14 | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x004975A0 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Toolbox/MeCCA/MediaEngine/Video/Codec` | Known | Audio system |

---

## 5. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00071AA8 | `HandleNotesSelected` | Known | Event handler |
| 0x00071AC0 | `HandleNotesPop` | Known | Event handler |
| 0x00071AD0 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00072ACC | `HandleExit` | Known | Event handler |
| 0x00072ADC | `HandleLap` | Known | Event handler |
| 0x00072AE8 | `HandleResume` | Known | Event handler |
| 0x00072AF8 | `HandleStartStop` | Known | Event handler |
| 0x00073E9C | `HandleNotesPop` | Known | Event handler |
| 0x00073EB0 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0007470C | `HandleChosen` | Known | Event handler |
| 0x00076C44 | `HandleDelete` | Known | Event handler |
| 0x00076C58 | `HandleSelectLozinch` | Known | Event handler |
| 0x000771D8 | `HandleSelect` | Known | Event handler |
| 0x000778D8 | `HandleSelect` | Known | Event handler |
| 0x000778E8 | `HandleSelectRating` | Known | Event handler |
| 0x000778FC | `HandleSelectProgress` | Known | Event handler |
| 0x00077914 | `HandleWheelProgress` | Known | Event handler |
| 0x00077928 | `HandleSelectScrub` | Known | Event handler |
| 0x0007793C | `HandleWheelBrightness` | Known | Event handler |
| 0x00077954 | `HandleNextTrackGotoScrub` | Known | Event handler |
| 0x00077970 | `HandlePrevTrackGotoScrub` | Known | Event handler |
| 0x0007798C | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x000779AC | `HandleMikeyCenter` | Known | Event handler |
| 0x0007866C | `HandleSelect` | Known | Event handler |
| 0x00078680 | `HandleWheel` | Known | Event handler |
| 0x0007868C | `HandleWheelProgress` | Known | Event handler |
| 0x000786A0 | `HandleSelectProgress` | Known | Event handler |
| 0x000786B8 | `HandleSelectVolume` | Known | Event handler |
| 0x000786CC | `HandleSelectScrub` | Known | Event handler |
| 0x000786E0 | `HandleSelectGenius` | Known | Event handler |
| 0x000786F4 | `HandleSelectRating` | Known | Event handler |
| 0x00078708 | `HandleSelectExtraInfo` | Known | Event handler |
| 0x00078720 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x0007873C | `HandleWheelGenius` | Known | Event handler |
| 0x00078750 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0007876C | `HandleWheelBrightness` | Known | Event handler |
| 0x00078784 | `HandleAddToOTG` | Known | Event handler |
| 0x00078794 | `HandleViewArtist` | Known | Event handler |
| 0x000787A8 | `HandleViewAlbum` | Known | Event handler |
| 0x000787B8 | `HandleViewCompilation` | Known | Event handler |
| 0x000787D0 | `HandleStartGenius` | Known | Event handler |
| 0x000787E4 | `HandleAudiobookSlower` | Known | Event handler |
| 0x000787FC | `HandleAudiobookFaster` | Known | Event handler |
| 0x00078814 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0007882C | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00078848 | `HandlePushContextualMenu` | Known | Event handler |
| 0x000788B8 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x000788D8 | `HandleOrientationChange` | Known | Event handler |
| 0x000788F0 | `HandleRotationChange` | Known | Event handler |
| 0x00078908 | `HandlePlayPauseTV` | Known | Event handler |
| 0x0007891C | `HandleSwapToVideoScreen` | Known | Event handler |
| 0x00078934 | `HandleSwapToMusicScreen` | Known | Event handler |
| 0x0007894C | `HandleMikeyCenter` | Known | Event handler |
| 0x00078960 | `HandleRemoteMenu` | Known | Event handler |
| 0x00078D94 | `HandleAudiobookSlower` | Known | Event handler |
| 0x00078DAC | `HandleAudiobookNormal` | Known | Event handler |
| 0x00078DC4 | `HandleAudiobookFaster` | Known | Event handler |
| 0x00078DDC | `HandleStartGenius` | Known | Event handler |
| 0x00078DF0 | `HandleAddToOTG` | Known | Event handler |
| 0x00078E04 | `HandleViewCompilation` | Known | Event handler |
| 0x00078E1C | `HandleViewAlbum` | Known | Event handler |
| 0x00078E2C | `HandleViewArtist` | Known | Event handler |
| 0x00078E40 | `HandleCancel` | Known | Event handler |
| 0x00078E9C | `HandleSelect` | Known | Event handler |
| 0x00079078 | `HandleSelect` | Known | Event handler |
| 0x00087F38 | `HandleMenuSelection` | Known | Event handler |
| 0x000A1510 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000A1530 | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x000A276C | `HandleSelect` | Known | Event handler |
| 0x000A2780 | `HandleHilite` | Known | Event handler |
| 0x000A29E0 | `HandleEQSettingSelected` | Known | Event handler |
| 0x000A29FC | `HandleEQSettingPreview` | Known | Event handler |
| 0x000A315C | `HandleTunerContextMenu` | Known | Event handler |
| 0x000A31EC | `HandleVolumeChange` | Known | Event handler |
| 0x000A3200 | `HandleVolumeWheel` | Known | Event handler |
| 0x000A3214 | `HandleTunerWheel` | Known | Event handler |
| 0x000A3228 | `HandleBufferWheel` | Known | Event handler |
| 0x000A323C | `HandleBandWheel` | Known | Event handler |
| 0x000A324C | `HandlePreviousPress` | Known | Event handler |
| 0x000A3260 | `HandleNextPress` | Known | Event handler |
| 0x000A3270 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x000A328C | `HandleNextPressAndHold` | Known | Event handler |
| 0x000A32A4 | `HandlePreviousTuning` | Known | Event handler |
| 0x000A32BC | `HandleNextTuning` | Known | Event handler |
| 0x000A32D0 | `HandlePlayPause` | Known | Event handler |
| 0x000A32E0 | `HandleMikeyCenter` | Known | Event handler |
| 0x000A32F4 | `HandleMikeyPrevious` | Known | Event handler |
| 0x000A3308 | `HandleMikeyNext` | Known | Event handler |
| 0x000A3318 | `HandleMikeyVolume` | Known | Event handler |
| 0x000A3EB0 | `HandlePushToCapacity` | Known | Event handler |
| 0x000A3ECC | `HandlePopToCapacity` | Known | Event handler |
| 0x000A3EE0 | `HandlePushToCount` | Known | Event handler |
| 0x000A3EF4 | `HandlePopToCount` | Known | Event handler |
| 0x000A3F08 | `HandlePushToBasic` | Known | Event handler |
| 0x000A3F1C | `HandlePopToBasic` | Known | Event handler |
| 0x000A3F30 | `HandlePushToAccessoryCapacity` | Known | Event handler |
| 0x000A3F50 | `HandlePopToAccessoryCapacity` | Known | Event handler |
| 0x000A3F70 | `HandlePushToAccessoryCount` | Known | Event handler |
| 0x000A3F8C | `HandlePopToAccessoryCount` | Known | Event handler |
| 0x000A3FA8 | `HandlePushToAccessoryBasic` | Known | Event handler |
| 0x000A3FC4 | `HandlePopToAccessoryBasic` | Known | Event handler |
| 0x000A3FE0 | `HandlePushToAccessoryAccessory` | Known | Event handler |
| 0x000A4000 | `HandlePopToAccessoryAccessory` | Known | Event handler |
| 0x000A4AFC | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x000A7CE4 | `HandleSelectCity` | Known | Event handler |
| 0x000A802C | `HandleOrientationChange` | Known | Event handler |
| 0x000A8048 | `HandleVolumePopup` | Known | Event handler |
| 0x000A880C | `HandleLeaveAlarm` | Known | Event handler |
| 0x000A89A8 | `HandleSelect` | Known | Event handler |
| 0x000A8C58 | `HandleHiliteAlbum` | Known | Event handler |
| 0x000A8C70 | `HandleBrowseAlbum` | Known | Event handler |
| 0x000A8C84 | `HandlePlayPause` | Known | Event handler |
| 0x000A8C94 | `HandlePushEvents` | Known | Event handler |
| 0x000A8CA8 | `HandlePopEvents` | Known | Event handler |
| 0x000A8CB8 | `HandlePushFaces` | Known | Event handler |
| 0x000A8CC8 | `HandlePopFaces` | Known | Event handler |
| 0x000A8CD8 | `HandlePushPlaces` | Known | Event handler |
| 0x000A8CEC | `HandlePopPlaces` | Known | Event handler |
| 0x000AA814 | `HandleSelect` | Known | Event handler |
| 0x000AA9B0 | `HandleSelectRegion_Africa` | Known | Event handler |
| 0x000AA9D0 | `HandleSelectRegion_Asia` | Known | Event handler |
| 0x000AA9E8 | `HandleSelectRegion_Atlantic` | Known | Event handler |
| 0x000AAA04 | `HandleSelectRegion_Australia` | Known | Event handler |
| 0x000AAA24 | `HandleSelectRegion_Europe` | Known | Event handler |
| 0x000AAA40 | `HandleSelectRegion_NorthAmerica` | Known | Event handler |
| 0x000AAA60 | `HandleSelectRegion_Pacific` | Known | Event handler |
| 0x000AAA7C | `HandleSelectRegion_SouthAmerica` | Known | Event handler |
| 0x000AD954 | `HandleLanguage` | Known | Event handler |
| 0x000AD968 | `HandleLanguagePop` | Known | Event handler |
| 0x000ADB88 | `HandleMainMenu` | Known | Event handler |
| 0x000AE0B4 | `HandlePlayRadio` | Known | Event handler |
| 0x000AE0C8 | `HandleStopRadio` | Known | Event handler |
| 0x000AE0D8 | `HandleAutoTune` | Known | Event handler |
| 0x000AE0E8 | `HandleTogglePlayPause` | Known | Event handler |
| 0x000AE100 | `HandleToggleBufferSetting` | Known | Event handler |
| 0x000AE11C | `HandleScanLogging` | Known | Event handler |
| 0x000AEAE0 | `HandleSelect` | Known | Event handler |
| 0x000AEC28 | `HandleSelect` | Known | Event handler |
| 0x000AED98 | `HandleMusicMenu` | Known | Event handler |
| 0x000B050C | `HandleSelectPreset` | Known | Event handler |
| 0x000B0524 | `HandleTogglePlayPause` | Known | Event handler |
| 0x000B0730 | `HandlePrev` | Known | Event handler |
| 0x000B0740 | `HandleNext` | Known | Event handler |
| 0x000B074C | `HandlePlayPause` | Known | Event handler |
| 0x000B0988 | `HandleNextContact` | Known | Event handler |
| 0x000B09A0 | `HandlePreviousContact` | Known | Event handler |
| 0x000B0BE4 | `HandleSelectPressAndHold` | Known | Event handler |
| 0x000B0C04 | `HandleDeleteItem` | Known | Event handler |
| 0x000B0C18 | `HandleDeleteAllItems` | Known | Event handler |
| 0x000B15A4 | `HandleItemSelected` | Known | Event handler |
| 0x000B16B4 | `HandleSelect` | Known | Event handler |
| 0x000B17F8 | `HandleSelect` | Known | Event handler |
| 0x000B194C | `HandleRadioRegion` | Known | Event handler |
| 0x000B2A24 | `HandlePlayPause` | Known | Event handler |
| 0x000B3408 | `HandleCenterButtonSelected` | Known | Event handler |
| 0x000B3B0C | `HandleTVOutChanged` | Known | Event handler |
| 0x000B3B24 | `HandleTVSignalChanged` | Known | Event handler |
| 0x000B3B3C | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x000B3B5C | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x000B3B7C | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x000B3BA0 | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x000B3BC0 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x000B4254 | `HandleSelect` | Known | Event handler |
| 0x000B4B54 | `HandleLeaveAlarm` | Known | Event handler |
| 0x000B5A78 | `HandleItemSelected` | Known | Event handler |
| 0x000B5FEC | `HandleResetAllSettings` | Known | Event handler |
| 0x000B6004 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x000B67C4 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x000B6BC8 | `HandleSelect` | Known | Event handler |
| 0x000B7D90 | `HandleStartGenius` | Known | Event handler |
| 0x000B8174 | `HandleNextDay` | Known | Event handler |
| 0x000B8188 | `HandlePreviousDay` | Known | Event handler |
| 0x000B8448 | `HandleWheel` | Known | Event handler |
| 0x000B860C | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x000B8710 | `HandleNextDay` | Known | Event handler |
| 0x000B8724 | `HandlePreviousDay` | Known | Event handler |
| 0x000B8820 | `HandleSelect` | Known | Event handler |
| 0x000B94D8 | `HandleDeleteClock` | Known | Event handler |
| 0x000B94F0 | `HandleSelectClock` | Known | Event handler |
| 0x000B9504 | `HandleHilited` | Known | Event handler |
| 0x000B9514 | `HandleWheel` | Known | Event handler |
| 0x000B9520 | `HandleSelectLozinch` | Known | Event handler |
| 0x000B9534 | `HandleAddClock` | Known | Event handler |
| 0x000B9544 | `HandleEditClock` | Known | Event handler |
| 0x000C04B8 | `HandleWheel` | Known | Event handler |
| 0x000C04C4 | `HandlePlayPause` | Known | Event handler |
| 0x000C04D4 | `HandleSelectDown` | Known | Event handler |
| 0x000C04E8 | `HandleNext` | Known | Event handler |
| 0x000C04F4 | `HandlePrevious` | Known | Event handler |
| 0x000C0504 | `HandleNextPushAndHold` | Known | Event handler |
| 0x000C051C | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000C4394 | `HandlePortEvents` | Known | Event handler |
| 0x000C4B4C | `HandleWantPopFlow` | Known | Event handler |
| 0x000C4B64 | `HandleSwapToNowPlayingFromOrientation` | Known | Event handler |
| 0x000C4B8C | `HandleSwapToNowPlaying` | Known | Event handler |
| 0x000C4BA4 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x000C4BC0 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x000C4BDC | `HandleFlowNext` | Known | Event handler |
| 0x000C4BEC | `HandleFlowPrev` | Known | Event handler |
| 0x000C4BFC | `HandleFlowWheel` | Known | Event handler |
| 0x000C4C0C | `HandleAlbumSelected` | Known | Event handler |
| 0x000C4C20 | `HandlePlayPause` | Known | Event handler |
| 0x000C4C30 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x000C4C4C | `HandleScreenRotation` | Known | Event handler |
| 0x000C4C64 | `HandleNext` | Known | Event handler |
| 0x000C4C70 | `HandleNextPressAndHold` | Known | Event handler |
| 0x000C4C88 | `HandlePrevious` | Known | Event handler |
| 0x000C4C98 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x000E0444 | `HandleSelect` | Known | Event handler |
| 0x000E0458 | `HandleGameHilited` | Known | Event handler |
| 0x000E0EAC | `HandleLock` | Known | Event handler |
| 0x000E0EBC | `HandleGotoAddressBookScreen` | Known | Event handler |
| 0x000E0ED8 | `HandleGotoCalendarLoadingScreen` | Known | Event handler |
| 0x000E0EF8 | `HandleNikePlusSelected` | Known | Event handler |
| 0x000E0F10 | `HandleVoiceMemosSelected` | Known | Event handler |
| 0x000E0F2C | `HandleRadioSelected` | Known | Event handler |
| 0x000E0F40 | `HandleRadioPlayPause` | Known | Event handler |
| 0x000E12A0 | `HandleOrientationPortrait` | Known | Event handler |
| 0x000E12C0 | `HandleOrientationLandscape` | Known | Event handler |
| 0x000E12DC | `HandleScreenRotation` | Known | Event handler |
| 0x000E12F4 | `HandleGestureShake` | Known | Event handler |
| 0x000E1308 | `HandleGestureSteer` | Known | Event handler |
| 0x000E16A8 | `HandlePlayPause` | Known | Event handler |
| 0x000E16BC | `HandleAddChapterMark` | Known | Event handler |
| 0x000E2914 | `HandleExitUnsupported` | Known | Event handler |
| 0x000E356C | `HandleScreenRotation` | Known | Event handler |
| 0x000E3C30 | `HandleNext` | Known | Event handler |
| 0x000E3C40 | `HandlePrev` | Known | Event handler |
| 0x000E3C4C | `HandleNextPressAndHold` | Known | Event handler |
| 0x000E3C64 | `HandlePrevPressAndHold` | Known | Event handler |
| 0x000E3C7C | `HandleWheel` | Known | Event handler |
| 0x000E3C88 | `HandleOrientationAlt` | Known | Event handler |
| 0x000E3CA0 | `HandleOrientationDefault` | Known | Event handler |
| 0x000E3CBC | `HandleScreenRotation` | Known | Event handler |
| 0x000E3CD4 | `HandlePlayPause` | Known | Event handler |
| 0x000E3CE4 | `HandlePlay` | Known | Event handler |
| 0x000E3CF0 | `HandlePause` | Known | Event handler |
| 0x000E3CFC | `HandleMikeyPlayPause` | Known | Event handler |
| 0x000E3D14 | `HandleSelect` | Known | Event handler |
| 0x000E3D24 | `HandleMenuUp` | Known | Event handler |
| 0x000E5DAC | `HandleWheel` | Known | Event handler |
| 0x000E5DBC | `HandleArrowUp` | Known | Event handler |
| 0x000E5DCC | `HandleArrowDown` | Known | Event handler |
| 0x000E9610 | `HandleSteer` | Known | Event handler |
| 0x000E9904 | `HandleSelect` | Known | Event handler |
| 0x000E9B88 | `HandleShowRecordings` | Known | Event handler |
| 0x000E9BAC | `HandleDeleteAllSelect` | Known | Event handler |
| 0x000E9BC4 | `HandleDeleteSelect` | Known | Event handler |
| 0x000E9BD8 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x000E9BF8 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x000E9C1C | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x000E9C38 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x000E9C58 | `HandleMicrophoneRequired` | Known | Event handler |
| 0x000E9C74 | `HandleMicrophoneDisconnected` | Known | Event handler |
| 0x000EA09C | `HandleFrequencyChosen` | Known | Event handler |
| 0x000EA0B4 | `HandleDateChosen` | Known | Event handler |
| 0x000EA0C8 | `HandleTimeChosen` | Known | Event handler |
| 0x000EA0DC | `HandleSoundChosen` | Known | Event handler |
| 0x000EA0F0 | `HandleLabelChosen` | Known | Event handler |
| 0x000EA104 | `HandleDeleteChosen` | Known | Event handler |
| 0x000EA1F0 | `HandleSelect` | Known | Event handler |
| 0x000EA6A4 | `HandleOrientationAlt` | Known | Event handler |
| 0x000EAA04 | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x000EADE0 | `HandleSelectPressAndHold` | Known | Event handler |
| 0x000EAE00 | `HandleDeleteItem` | Known | Event handler |
| 0x000EAE14 | `HandleDeleteAllItems` | Known | Event handler |
| 0x000EBEC4 | `HandleSelectLabel` | Known | Event handler |
| 0x000ECD90 | `HandleStartGenius` | Known | Event handler |
| 0x000ECDA8 | `HandleViewArtist` | Known | Event handler |
| 0x000ECDBC | `HandleViewAlbum` | Known | Event handler |
| 0x000ECDCC | `HandleViewCompilation` | Known | Event handler |
| 0x000ECDE4 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x000ED544 | `HandleStartGenius` | Known | Event handler |
| 0x000ED558 | `HandleAddToOTG` | Known | Event handler |
| 0x000ED56C | `HandleViewCompilation` | Known | Event handler |
| 0x000ED584 | `HandleViewAlbum` | Known | Event handler |
| 0x000ED594 | `HandleViewArtist` | Known | Event handler |
| 0x000ED5A8 | `HandleCancel` | Known | Event handler |
| 0x000EDDCC | `HandleAddToOTG` | Known | Event handler |
| 0x000EDDDC | `HandleCancel` | Known | Event handler |
| 0x000EDEEC | `HandleStartGenius` | Known | Event handler |
| 0x000EDF04 | `HandleViewAlbum` | Known | Event handler |
| 0x000EDF14 | `HandleViewArtist` | Known | Event handler |
| 0x000EDF28 | `HandleViewCompilation` | Known | Event handler |
| 0x000EDF40 | `HandleSaveGeniusPlaylist` | Known | Event handler |
| 0x000EDF5C | `HandleRefreshPlaylist` | Known | Event handler |
| 0x000EDF74 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x000EEA90 | `HandleStartGenius` | Known | Event handler |
| 0x000EEAA4 | `HandleAddToOTG` | Known | Event handler |
| 0x000EEAB8 | `HandleViewCompilation` | Known | Event handler |
| 0x000EEAD0 | `HandleViewAlbum` | Known | Event handler |
| 0x000EEAE0 | `HandleViewArtist` | Known | Event handler |
| 0x000EEAF4 | `HandleCancel` | Known | Event handler |
| 0x000EED64 | `HandleAddToOTG` | Known | Event handler |
| 0x000EED74 | `HandleCancel` | Known | Event handler |
| 0x000EF640 | `HandleAddToOTG` | Known | Event handler |
| 0x000EF650 | `HandleCancel` | Known | Event handler |
| 0x000EFCF0 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x000F004C | `HandleAddToOTG` | Known | Event handler |
| 0x000F005C | `HandleCancel` | Known | Event handler |
| 0x000F02C0 | `HandleConfirmation` | Known | Event handler |
| 0x000F0758 | `HandleMusicHilited` | Known | Event handler |
| 0x000F0770 | `HandleVideosHilited` | Known | Event handler |
| 0x000F0784 | `HandlePodcastsHilited` | Known | Event handler |
| 0x000F079C | `HandleiTunesUHilited` | Known | Event handler |
| 0x000F07B4 | `HandleGenericHilited` | Known | Event handler |
| 0x000F07CC | `HandlePhotosHilited` | Known | Event handler |
| 0x000F07E0 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x000F07F8 | `HandleExtrasHilited` | Known | Event handler |
| 0x000F080C | `HandleSettingsHilited` | Known | Event handler |
| 0x000F0824 | `HandleCameraHilited` | Known | Event handler |
| 0x000F0838 | `HandleNoBookMusicHilited` | Known | Event handler |
| 0x000F0854 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x000F086C | `HandleArtistsHilited` | Known | Event handler |
| 0x000F0884 | `HandleGenresHilited` | Known | Event handler |
| 0x000F0898 | `HandleAlbumsHilited` | Known | Event handler |
| 0x000F08AC | `HandleCompilationsHilited` | Known | Event handler |
| 0x000F08C8 | `HandleComposersHilited` | Known | Event handler |
| 0x000F08E0 | `HandleSongsHilited` | Known | Event handler |
| 0x000F08F4 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x000F090C | `HandleGeniusMixesHilited` | Known | Event handler |
| 0x000F0928 | `HandleCoverflowHilited` | Known | Event handler |
| 0x000F0940 | `HandleTVShowsHilited` | Known | Event handler |
| 0x000F0958 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x000F0974 | `HandleMoviesHilited` | Known | Event handler |
| 0x000F0988 | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x000F09A4 | `HandleRentalsHilited` | Known | Event handler |
| 0x000F09BC | `HandleRadioHilited` | Known | Event handler |
| 0x000F09D0 | `HandleVoiceMemosHilited` | Known | Event handler |
| 0x000F09E8 | `HandleMusicSelected` | Known | Event handler |
| 0x000F09FC | `HandleVideosSelected` | Known | Event handler |
| 0x000F0A14 | `HandlePodcastsSelected` | Known | Event handler |
| 0x000F0A2C | `HandleiTunesUSelected` | Known | Event handler |
| 0x000F0A44 | `HandlePhotosSelected` | Known | Event handler |
| 0x000F0A5C | `HandleCoverFlowSelected` | Known | Event handler |
| 0x000F0A74 | `HandleSongsSelected` | Known | Event handler |
| 0x000F0A88 | `HandleAlbumsSelected` | Known | Event handler |
| 0x000F0AA0 | `HandleCompilationsSelected` | Known | Event handler |
| 0x000F0ABC | `HandleArtistsSelected` | Known | Event handler |
| 0x000F0AD4 | `HandleGenresSelected` | Known | Event handler |
| 0x000F0D48 | `HandleComposersSelected` | Known | Event handler |
| 0x000F0D60 | `HandleAudiobooksSelected` | Known | Event handler |
| 0x000F0D7C | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x000F0D98 | `HandlePlaylistsSelected` | Known | Event handler |
| 0x000F0DB0 | `HandleNowPlaying` | Known | Event handler |
| 0x000F0DC4 | `HandleCameraSelected` | Known | Event handler |
| 0x000F0DDC | `HandleGeniusMixesSelected` | Known | Event handler |
| 0x000F0DF8 | `HandleTVShowsSelected` | Known | Event handler |
| 0x000F0E10 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x000F0E2C | `HandleMoviesSelected` | Known | Event handler |
| 0x000F0E44 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x000F0E64 | `HandleRentalsSelected` | Known | Event handler |
| 0x000F0E7C | `HandleCameraVideosSelected` | Known | Event handler |
| 0x000F0E98 | `HandleAddressBookSelected` | Known | Event handler |
| 0x000F0EB4 | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x000F0ECC | `HandleSleepSelected` | Known | Event handler |
| 0x000F0EE0 | `HandleNikePlusSelected` | Known | Event handler |
| 0x000F0EF8 | `HandleNikePlusHilited` | Known | Event handler |
| 0x000F0F10 | `HandleRadioSelected` | Known | Event handler |
| 0x000F0F24 | `HandleRadioPreviewPlayPause` | Known | Event handler |
| 0x000F0F40 | `HandlePedometerSelected` | Known | Event handler |
| 0x000F0F58 | `HandlePedometerHilited` | Known | Event handler |
| 0x000F0F70 | `HandlePortraitToLandscape` | Known | Event handler |
| 0x000F14EC | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x000F1784 | `HandleAddToOTG` | Known | Event handler |
| 0x000F1794 | `HandleCancel` | Known | Event handler |
| 0x000F21AC | `HandleAddToOTG` | Known | Event handler |
| 0x000F21BC | `HandleCancel` | Known | Event handler |
| 0x000F2858 | `HandleAddToOTG` | Known | Event handler |
| 0x000F2868 | `HandleCancel` | Known | Event handler |
| 0x000F2A58 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x000F2ED8 | `HandleGeniusMixPlaylistReady` | Known | Event handler |
| 0x000F2EFC | `HandleSelectMix` | Known | Event handler |
| 0x000F2F0C | `HandlePrev` | Known | Event handler |
| 0x000F2F18 | `HandleNext` | Known | Event handler |
| 0x000F2F24 | `HandlePlayPause` | Known | Event handler |
| 0x000F380C | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x000F3828 | `HandleSaveOTGPlaylist` | Known | Event handler |
| 0x000F3840 | `HandleStartGenius` | Known | Event handler |
| 0x000F3854 | `HandleViewArtist` | Known | Event handler |
| 0x000F3868 | `HandleViewAlbum` | Known | Event handler |
| 0x000F3878 | `HandleViewCompilation` | Known | Event handler |
| 0x000F3890 | `HandleShowContextualMenu` | Known | Event handler |
| 0x000F38AC | `HandleRefreshPlaylist` | Known | Event handler |
| 0x000F38C4 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x000F4554 | `HandleStartGenius` | Known | Event handler |
| 0x000F4568 | `HandleAddToOTG` | Known | Event handler |
| 0x000F457C | `HandleViewCompilation` | Known | Event handler |
| 0x000F4594 | `HandleViewAlbum` | Known | Event handler |
| 0x000F45A4 | `HandleViewArtist` | Known | Event handler |
| 0x000F45B8 | `HandleCancel` | Known | Event handler |
| 0x000F460C | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x000F476C | `HandleAddToOTG` | Known | Event handler |
| 0x000F477C | `HandleCancel` | Known | Event handler |
| 0x000F5910 | `HandleSelect` | Known | Event handler |
| 0x000F5920 | `HandleCameraRemoteSelect` | Known | Event handler |
| 0x000F593C | `HandleMikeyCenter` | Known | Event handler |
| 0x000F5950 | `HandleMikeyCenterReleased` | Known | Event handler |
| 0x000F5988 | `HandleSelectPressAndHold` | Known | Event handler |
| 0x000F60F0 | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x000F6110 | `HandleMikeyCenterReleased` | Known | Event handler |
| 0x000F612C | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x000F614C | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x000F6170 | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x000F6190 | `HandleMikeyAllUp` | Known | Event handler |
| 0x000F61A4 | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x000F61B8 | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x000F61D0 | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x000F61E8 | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x000F6200 | `HandleSelect` | Known | Event handler |
| 0x000F6210 | `HandleCameraRemoteSelect` | Known | Event handler |
| 0x000F622C | `HandleMikeyCenter` | Known | Event handler |
| 0x000F6388 | `HandleSelect` | Known | Event handler |
| 0x000F639C | `HandleNext` | Known | Event handler |
| 0x000F63A8 | `HandlePrev` | Known | Event handler |
| 0x000F63B4 | `HandleWheel` | Known | Event handler |
| 0x000F6474 | `HandleAppInitialized` | Known | Event handler |
| 0x000FAD14 | `HandleSelect` | Known | Event handler |
| 0x001011A8 | `HandleRotationChange` | Known | Event handler |
| 0x001011C4 | `HandleSteerGesture` | Known | Event handler |
| 0x001011D8 | `HandlePlayPause` | Known | Event handler |
| 0x001011E8 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x00101204 | `HandleTouchAndHoldPlayPause` | Known | Event handler |
| 0x00101220 | `HandleNext` | Known | Event handler |
| 0x0010122C | `HandleNextPressAndHold` | Known | Event handler |
| 0x00101244 | `HandlePrevious` | Known | Event handler |
| 0x00101254 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x00101270 | `HandleRemotePlayPause` | Known | Event handler |
| 0x00101288 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x001012AC | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x001012C4 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x001012DC | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x001012F4 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x0010130C | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x00101324 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x00101340 | `HandleRemoteStop` | Known | Event handler |
| 0x00101354 | `HandleRemotePlay` | Known | Event handler |
| 0x00101368 | `HandleRemotePause` | Known | Event handler |
| 0x0010137C | `HandleRemoteMute` | Known | Event handler |
| 0x00101390 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x001013A8 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x001013C0 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x001013DC | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x001013F8 | `HandleRemoteShuffle` | Known | Event handler |
| 0x0010140C | `HandleRemoteRepeat` | Known | Event handler |
| 0x00101420 | `HandleRemoteOn` | Known | Event handler |
| 0x00101430 | `HandleRemoteOff` | Known | Event handler |
| 0x00101440 | `HandleRemoteBacklight` | Known | Event handler |
| 0x00101458 | `HandleRemoteFFDown` | Known | Event handler |
| 0x0010146C | `HandleRemoteFFUp` | Known | Event handler |
| 0x00101480 | `HandleRemoteRewDown` | Known | Event handler |
| 0x00101494 | `HandleRemoteRewUp` | Known | Event handler |
| 0x001014A8 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x001014C0 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x001014D4 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x001014EC | `HandleRemoteSelectUp` | Known | Event handler |
| 0x00101504 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x0010151C | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x00101534 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x00101910 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x00101928 | `HandleRemoteEvent` | Known | Event handler |
| 0x0010193C | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x00101958 | `HandleAudioPlayPause` | Known | Event handler |
| 0x00101970 | `HandleAudioNext` | Known | Event handler |
| 0x00101980 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x0010199C | `HandleAudioPrevious` | Known | Event handler |
| 0x001019B0 | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x001019D0 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x001019E8 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x00101A00 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x00101A18 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x00101A2C | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x00101A44 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x00101A5C | `HandleAudioStop` | Known | Event handler |
| 0x00101A6C | `HandleAudioPlay` | Known | Event handler |
| 0x00101A7C | `HandleAudioPause` | Known | Event handler |
| 0x00101A90 | `HandleAudioMute` | Known | Event handler |
| 0x00101AA0 | `HandleAudioNextChapter` | Known | Event handler |
| 0x00101AB8 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x00101AD0 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x00101AE8 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x00101B00 | `HandleAudioShuffle` | Known | Event handler |
| 0x00101B14 | `HandleAudioRepeat` | Known | Event handler |
| 0x00101B28 | `HandleAudioFFDown` | Known | Event handler |
| 0x00101B3C | `HandleAudioFFUp` | Known | Event handler |
| 0x00101B4C | `HandleAudioRewDown` | Known | Event handler |
| 0x00101B60 | `HandleAudioRewUp` | Known | Event handler |
| 0x00101B74 | `HandleVideoPlayPause` | Known | Event handler |
| 0x00101B8C | `HandleVideoNext` | Known | Event handler |
| 0x00101B9C | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x00101BB8 | `HandleVideoPrevious` | Known | Event handler |
| 0x00101BCC | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x00101BEC | `HandleVideoStop` | Known | Event handler |
| 0x00101BFC | `HandleVideoPlay` | Known | Event handler |
| 0x00101C0C | `HandleVideoPause` | Known | Event handler |
| 0x00101C20 | `HandleVideoFFDown` | Known | Event handler |
| 0x00101C34 | `HandleVideoFFUp` | Known | Event handler |
| 0x00101C44 | `HandleVideoRewDown` | Known | Event handler |
| 0x00101C58 | `HandleVideoRewUp` | Known | Event handler |
| 0x00101C6C | `HandleVideoNextChapter` | Known | Event handler |
| 0x00101C84 | `HandleVideoPrevChapter` | Known | Event handler |
| 0x00101F48 | `HandleVideoNextFrame` | Known | Event handler |
| 0x00101F60 | `HandleVideoPrevFrame` | Known | Event handler |
| 0x00101F78 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x00101F94 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00101FB0 | `HandleShakeShuffleSongsSelected` | Known | Event handler |
| 0x00101FD0 | `HandleGlobalVolume` | Known | Event handler |
| 0x00101FE4 | `HandleSimulateOrientationChange` | Known | Event handler |
| 0x00102004 | `HandleMikeyCenter` | Known | Event handler |
| 0x00102018 | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x00102038 | `HandleMikeyCenterReleased` | Known | Event handler |
| 0x00102054 | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x00102074 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x00102098 | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x001020B8 | `HandleMikeyCenterTripleClickAndHold` | Known | Event handler |
| 0x001020DC | `HandleMikeyAllUp` | Known | Event handler |
| 0x001020F0 | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x00102104 | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x0010211C | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x00102134 | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x0010214C | `HandleVoiceOverPlaylistSelected` | Known | Event handler |
| 0x0010216C | `HandleVoiceOverPodcastSelected` | Known | Event handler |
| 0x0010218C | `HandleVoiceOverAudiobookSelected` | Known | Event handler |
| 0x001021B0 | `HandleRadioTagUp` | Known | Event handler |
| 0x001021C4 | `HandleVoiceCommand` | Known | Event handler |
| 0x001021D8 | `HandleVoiceArtist` | Known | Event handler |
| 0x001021EC | `HandleVoiceAlbum` | Known | Event handler |
| 0x00102200 | `HandleNECIRMenuUp` | Known | Event handler |
| 0x00102214 | `HandleNECIRPlayPauseUp` | Known | Event handler |
| 0x0010222C | `HandleNECIRNextUp` | Known | Event handler |
| 0x00102240 | `HandleNECIRPrevUp` | Known | Event handler |
| 0x00102254 | `HandleNECIRVolumeDownUp` | Known | Event handler |
| 0x0010226C | `HandleNECIRVolumeUpUp` | Known | Event handler |
| 0x00102284 | `HandleCameraRemote` | Known | Event handler |
| 0x001062CC | `HandleLoadingCancelled` | Known | Event handler |
| 0x001B0E5C | `HandleWheelVolume` | Known | Event handler |
| 0x001B0E74 | `HandleMenuKey` | Known | Event handler |
| 0x001B0E84 | `HandlePauseKey` | Known | Event handler |
| 0x001B0E94 | `HandlePrevKey` | Known | Event handler |
| 0x001B0EA4 | `HandleNextKey` | Known | Event handler |
| 0x001B10F8 | `HandleSelect` | Known | Event handler |
| 0x001B14AC | `HandleChooseLink` | Known | Event handler |
| 0x001B14C4 | `HandleUnlink` | Known | Event handler |
| 0x001B15E4 | `HandleSelectedDayWorkout` | Known | Event handler |
| 0x001B1604 | `HandleMenuUp` | Known | Event handler |
| 0x001B1810 | `HandleSelect` | Known | Event handler |
| 0x001B1824 | `HandleMenu` | Known | Event handler |
| 0x001B1830 | `HandleLinkCancelOption` | Known | Event handler |
| 0x001B1848 | `HandleLinkNewRemote` | Known | Event handler |
| 0x001B185C | `HandleLinkNewHeartMonitor` | Known | Event handler |
| 0x001B1878 | `HandleCancelRemoteLinking` | Known | Event handler |
| 0x001B1ADC | `HandleChoosePlaylistsPlay` | Known | Event handler |
| 0x001B1AFC | `HandleChoosePodcastsPlay` | Known | Event handler |
| 0x001B1B18 | `HandleChooseAudiobooksPlay` | Known | Event handler |
| 0x001B1B34 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x001B1B50 | `HandleNoneSelected` | Known | Event handler |
| 0x001B1B64 | `HandleNowPlayingSelected` | Known | Event handler |
| 0x001B1B80 | `HandleMenuUp` | Known | Event handler |
| 0x001B1B90 | `HandleMenuKeyNop` | Known | Event handler |
| 0x001B2228 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x001B25BC | `HandleChoosePowerPlay` | Known | Event handler |
| 0x001B25D8 | `HandleChooseUnit` | Known | Event handler |
| 0x001B25EC | `HandleChooseVoicePrompts` | Known | Event handler |
| 0x001B2608 | `HandleChoosePedometer` | Known | Event handler |
| 0x001B3090 | `HandlePopupDismissed` | Known | Event handler |
| 0x001B3144 | `HandleSelect` | Known | Event handler |
| 0x001B32AC | `HandleListChoose` | Known | Event handler |
| 0x001B3630 | `HandleBasicSelected` | Known | Event handler |
| 0x001B3648 | `HandleTimedWorkoutSelected` | Known | Event handler |
| 0x001B3664 | `HandleDistanceWorkoutSelected` | Known | Event handler |
| 0x001B3684 | `HandleCaloriesWorkoutSelected` | Known | Event handler |
| 0x001B3C10 | `HandleWalkCalibrationSelection` | Known | Event handler |
| 0x001B3C34 | `HandleRunCalibrationSelection` | Known | Event handler |
| 0x001B3DE4 | `HandleNewWorkout` | Known | Event handler |
| 0x001B3DFC | `HandleNewBasicWorkout` | Known | Event handler |
| 0x001B3E14 | `HandleNewQuickstartWorkout` | Known | Event handler |
| 0x001B3E30 | `HandleResumeWorkout` | Known | Event handler |
| 0x001B4028 | `HandleChooseTimedWorkoutFromList` | Known | Event handler |
| 0x001B4050 | `HandleChooseDistanceWorkoutFromList` | Known | Event handler |
| 0x001B4074 | `HandleChooseCalorieWorkoutFromList` | Known | Event handler |
| 0x001B4170 | `HandleVerticalSelected` | Known | Event handler |
| 0x001B418C | `HandleRightSelected` | Known | Event handler |
| 0x001B41A0 | `HandleLeftSelected` | Known | Event handler |
| 0x001B483C | `HandleBegin` | Known | Event handler |
| 0x001B4E3C | `HandleItemSelected` | Known | Event handler |
| 0x001B4F84 | `HandleSelect` | Known | Event handler |
| 0x001B4F98 | `HandlePopBackToSongsScreen` | Known | Event handler |
| 0x001B5210 | `HandleUnlinkRemote` | Known | Event handler |
| 0x001B5320 | `HandleChooseRemoteSetting` | Known | Event handler |
| 0x001B5340 | `HandleChooseRemoteLink` | Known | Event handler |
| 0x001B5358 | `HandleChooseRemoteUnlink` | Known | Event handler |
| 0x001B5590 | `HandleDeleteWorkout` | Known | Event handler |
| 0x001B55A8 | `HandleDeleteAllWorkouts` | Known | Event handler |
| 0x001B55C0 | `HandleOrientationChange` | Known | Event handler |
| 0x001B55D8 | `HandleSelectNextWorkout` | Known | Event handler |
| 0x001B55F0 | `HandleSelectPrevWorkout` | Known | Event handler |
| 0x001B5AEC | `HandleDeleteAllWorkouts` | Known | Event handler |
| 0x001B5B08 | `HandleClearBests` | Known | Event handler |
| 0x001B5B1C | `HandleClearTotals` | Known | Event handler |
| 0x001B5B30 | `HandleResetWalkingCalibration` | Known | Event handler |
| 0x001B5B50 | `HandleResetRuningCalibration` | Known | Event handler |
| 0x001B5C8C | `HandlePopSelf` | Known | Event handler |
| 0x001B5CA0 | `HandlePressAndHold` | Known | Event handler |
| 0x001B6138 | `HandleHerculesKey` | Known | Event handler |
| 0x001B624C | `HandleUnlinkHeartMonitor` | Known | Event handler |
| 0x001B6350 | `HandleChooseHeartMonitorLink` | Known | Event handler |
| 0x001B6374 | `HandleChooseHeartMonitorUnlink` | Known | Event handler |
| 0x001B889C | `HandleAddToOTG` | Known | Event handler |
| 0x001B88AC | `HandleCancel` | Known | Event handler |
| 0x001B8BE8 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x001B8C88 | `HandleAddToOTG` | Known | Event handler |
| 0x001B8C98 | `HandleCancel` | Known | Event handler |
| 0x001B8CE0 | `HandleAudiobookChapter_ResumeSelected` | Known | Event handler |
| 0x001BB0C8 | `HandleSelectedNikeMainMenuItem` | Known | Event handler |
| 0x001BB264 | `HandleSelect` | Known | Event handler |
| 0x001BB278 | `HandleMenuKey` | Known | Event handler |
| 0x001BB288 | `HandlePauseWorkout` | Known | Event handler |
| 0x001BB29C | `HandleEndWorkout` | Known | Event handler |
| 0x001BB2B0 | `HandleResumeWorkout` | Known | Event handler |
| 0x001BB2C4 | `HandleChooseMusic` | Known | Event handler |
| 0x001C1C74 | `HandleOrientationChange` | Known | Event handler |
| 0x001C1C90 | `HandleNext` | Known | Event handler |
| 0x001C1C9C | `HandlePrev` | Known | Event handler |
| 0x001C1CA8 | `HandleWheel` | Known | Event handler |
| 0x001C1CB4 | `HandleSelect` | Known | Event handler |
| 0x001C290C | `HandleWeightWheel` | Known | Event handler |
| 0x001C2924 | `HandleWeightSelect` | Known | Event handler |
| 0x001C2938 | `HandleWeightSelectAltTrans` | Known | Event handler |
| 0x001C2954 | `HandleDistanceWheel` | Known | Event handler |
| 0x001C2968 | `HandleDistanceSelect` | Known | Event handler |
| 0x001C2980 | `HandleTimeWheel` | Known | Event handler |
| 0x001C2990 | `HandleTimeSelect` | Known | Event handler |
| 0x001C29A4 | `HandleCaloriesWheel` | Known | Event handler |
| 0x001C29B8 | `HandleCaloriesSelect` | Known | Event handler |
| 0x001C29D0 | `HandleStepGoalWheel` | Known | Event handler |
| 0x001C29E4 | `HandleStepGoalSelect` | Known | Event handler |
| 0x001C29FC | `HandleWeightSelectPedometer` | Known | Event handler |
| 0x001C2B90 | `HandleDistanceWheel` | Known | Event handler |
| 0x001C2BA8 | `HandleDistanceSelect` | Known | Event handler |
| 0x001CA814 | `HandlePauseKey` | Known | Event handler |
| 0x001CA828 | `HandlePauseHold` | Known | Event handler |
| 0x001CA838 | `HandleMenuKey` | Known | Event handler |
| 0x001CA848 | `HandleMenuKeyNop` | Known | Event handler |
| 0x001CA85C | `HandleWheel` | Known | Event handler |
| 0x001CA8AC | `HandleSelectKeyDown` | Known | Event handler |
| 0x001CA8C0 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x001CA8D8 | `HandlePowerPlay` | Known | Event handler |
| 0x001CA8E8 | `HandlePauseWorkout` | Known | Event handler |
| 0x001CA8FC | `HandleEndWorkout` | Known | Event handler |
| 0x001CA910 | `HandleResumeWorkout` | Known | Event handler |
| 0x001CA924 | `HandleChooseMusic` | Known | Event handler |
| 0x001CA938 | `HandleMikeyPressExtended` | Known | Event handler |
| 0x001CA954 | `Handle3BitModeFinished` | Known | Event handler |
| 0x00211B9C | `HandlePlayPause` | Known | Event handler |
| 0x00211BB0 | `HandleWheel` | Known | Event handler |
| 0x00211BBC | `HandleMTWheel` | Known | Event handler |
| 0x00211BCC | `HandleWheelRating` | Known | Event handler |
| 0x00211BE0 | `HandleWheelScrub` | Known | Event handler |
| 0x00211BF4 | `HandleWheelVolume` | Known | Event handler |
| 0x00211CB4 | `HandleMenuKey` | Known | Event handler |
| 0x00211CC4 | `HandleMenuLongpress` | Known | Event handler |
| 0x00211CD8 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x00211CF4 | `HandleSwapToCoverflow` | Known | Event handler |
| 0x00211D0C | `HandleDefaultOrientation` | Known | Event handler |
| 0x00215864 | `HandleSelect` | Known | Event handler |
| 0x00240104 | `HandleMikeyCenter` | Known | Event handler |
| 0x00240118 | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x00240138 | `HandleMikeyCenterReleased` | Known | Event handler |
| 0x00240154 | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x00240174 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x00240198 | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x002401B8 | `HandleMikeyAllUp` | Known | Event handler |
| 0x002401CC | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x002401E0 | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x002401F8 | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x00240210 | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x00240228 | `HandleNext` | Known | Event handler |
| 0x00240234 | `HandlePrev` | Known | Event handler |
| 0x00240240 | `HandleNextAndHold` | Known | Event handler |
| 0x00240254 | `HandlePrevAndHold` | Known | Event handler |
| 0x00240268 | `HandleScreenRotation` | Known | Event handler |
| 0x00240280 | `HandleAudioPlayPause` | Known | Event handler |
| 0x00240298 | `HandleCameraRemoteSelect` | Known | Event handler |
| 0x002955B8 | `HandleSelect` | Known | Event handler |
| 0x002955CC | `HandleHilite` | Known | Event handler |
| 0x002955DC | `HandlePlayPause` | Known | Event handler |
| 0x002955EC | `HandleAddToOTG` | Known | Event handler |
| 0x002955FC | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x0029561C | `HandleShowContextualMenu` | Known | Event handler |
| 0x00295638 | `HandleStartQuickNav` | Known | Event handler |
| 0x004650F8 | `HandlePauseKey` | Known | Event handler |
| 0x00465108 | `HandleMenuKey` | Known | Event handler |
| 0x0056DC75 | `HandleMikeyCenterReleased` | Known | Event handler |
| 0x0056DC8F | `HandleAppInitialized` | Known | Event handler |
| 0x0056DCD5 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x0056DCF9 | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x0056DD17 | `HandleSelectPressAndHold` | Known | Event handler |
| 0x0056DD30 | `HandleNextAndHold` | Known | Event handler |
| 0x0056DD42 | `HandlePrevAndHold` | Known | Event handler |
| 0x0056DD54 | `HandleAudioPlayPause` | Known | Event handler |
| 0x0056DD8F | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x0056DDAC | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x0056DDC9 | `HandleWheel` | Known | Event handler |
| 0x0056DDFD | `HandleScreenRotation` | Known | Event handler |
| 0x0056DE12 | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x0056DE28 | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x0056DE3C | `HandleMikeyAllUp` | Known | Event handler |
| 0x0056DE4D | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x0056DE65 | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x0056DE7B | `HandleMikeyCenter` | Known | Event handler |
| 0x0056DE8D | `HandleSelect` | Known | Event handler |
| 0x0056DE9A | `HandleCameraRemoteSelect` | Known | Event handler |
| 0x0056DEB3 | `HandleNext` | Known | Event handler |
| 0x0056DECC | `HandlePrev` | Known | Event handler |
| 0x0063325C | `HandleShakeShuffleSongsSelected` | Known | Event handler |
| 0x0063328C | `HandleAudioFFDown` | Known | Event handler |
| 0x006332AC | `HandleAudioFFUp` | Known | Event handler |
| 0x006332CC | `HandleAudioMute` | Known | Event handler |
| 0x006332F4 | `HandleAudioNextPressAndHold` | Known | Event handler |
| 0x00633320 | `HandleAudioNext` | Known | Event handler |
| 0x00633348 | `HandleAudioNextAlbum` | Known | Event handler |
| 0x00633378 | `HandleAudioNextChapter` | Known | Event handler |
| 0x006333A8 | `HandleAudioNextPlaylist` | Known | Event handler |
| 0x006333D4 | `HandleAudioPause` | Known | Event handler |
| 0x006333F8 | `HandleAudioPlay` | Known | Event handler |
| 0x0063341C | `HandleAudioPlayPause` | Known | Event handler |
| 0x0063344C | `HandleAudioPreviousPressAndHold` | Known | Event handler |
| 0x0063347C | `HandleAudioPrevious` | Known | Event handler |
| 0x006334A8 | `HandleAudioPrevAlbum` | Known | Event handler |
| 0x006334D8 | `HandleAudioPrevChapter` | Known | Event handler |
| 0x00633508 | `HandleAudioPrevPlaylist` | Known | Event handler |
| 0x00633534 | `HandleAudioRepeat` | Known | Event handler |
| 0x00633558 | `HandleAudioRewDown` | Known | Event handler |
| 0x0063357C | `HandleAudioRewUp` | Known | Event handler |
| 0x006335A4 | `HandleAudioShuffle` | Known | Event handler |
| 0x006335C8 | `HandleAudioStop` | Known | Event handler |
| 0x006335F0 | `HandleAudioVolumeDown` | Known | Event handler |
| 0x0063361C | `HandleAudioVolumeDownUp` | Known | Event handler |
| 0x00633648 | `HandleAudioVolumeUp` | Known | Event handler |
| 0x00633670 | `HandleAudioVolumeUpUp` | Known | Event handler |
| 0x006336EC | `HandleNextPressAndHold` | Known | Event handler |
| 0x00633704 | `HandleNext` | Known | Event handler |
| 0x00633734 | `HandlePlayPausePressAndHold` | Known | Event handler |
| 0x00633770 | `HandlePlayPause` | Known | Event handler |
| 0x00633780 | `HandlePreviousPressAndHold` | Known | Event handler |
| 0x0063379C | `HandlePrevious` | Known | Event handler |
| 0x006337C4 | `HandleCameraRemote` | Known | Event handler |
| 0x00633818 | `HandleRemoteDownArrowDown` | Known | Event handler |
| 0x00633840 | `HandleRemoteDownArrowUp` | Known | Event handler |
| 0x00633890 | `HandleRemoteMenuDown` | Known | Event handler |
| 0x006338B4 | `HandleRemoteMenuUp` | Known | Event handler |
| 0x00633904 | `HandleRemotePlayPause` | Known | Event handler |
| 0x00633928 | `HandleRemoteUpArrowDown` | Known | Event handler |
| 0x0063394C | `HandleRemoteUpArrowUp` | Known | Event handler |
| 0x006339C4 | `HandleMikeyAllUp` | Known | Event handler |
| 0x006339F0 | `HandleMikeyCenterReleased` | Known | Event handler |
| 0x00633A28 | `HandleMikeyCenterDoubleClick` | Known | Event handler |
| 0x00633A68 | `HandleMikeyCenterDoubleClickAndHold` | Known | Event handler |
| 0x00633AA8 | `HandleMikeyCenterPressAndHold` | Known | Event handler |
| 0x00633AE4 | `HandleMikeyCenterTripleClick` | Known | Event handler |
| 0x00633B24 | `HandleMikeyCenterTripleClickAndHold` | Known | Event handler |
| 0x00633B58 | `HandleMikeyCenter` | Known | Event handler |
| 0x00633B84 | `HandleMikeyVolumeDown` | Known | Event handler |
| 0x00633BB0 | `HandleMikeyVolumeDownUp` | Known | Event handler |
| 0x00633BDC | `HandleMikeyVolumeUp` | Known | Event handler |
| 0x00633C04 | `HandleMikeyVolumeUpUp` | Known | Event handler |
| 0x00633CB8 | `HandleRemoteVolumeDown` | Known | Event handler |
| 0x00633CE4 | `HandleRemoteVolumeDownUp` | Known | Event handler |
| 0x00633D14 | `HandleRemoteVolumeUp` | Known | Event handler |
| 0x00633D40 | `HandleRemoteVolumeUpUp` | Known | Event handler |
| 0x00633D70 | `HandleRadioTagUp` | Known | Event handler |
| 0x00633D98 | `HandleRemoteBacklight` | Known | Event handler |
| 0x00633DC8 | `HandleRemoteBacklightOff` | Known | Event handler |
| 0x00633E24 | `HandleRemoteEvent` | Known | Event handler |
| 0x00633E48 | `HandleRemoteFFDown` | Known | Event handler |
| 0x00633E6C | `HandleRemoteFFUp` | Known | Event handler |
| 0x00633EB8 | `HandleRemoteMute` | Known | Event handler |
| 0x00633EE4 | `HandleRemoteNextAlbum` | Known | Event handler |
| 0x00633F14 | `HandleRemoteNextChapter` | Known | Event handler |
| 0x00633F48 | `HandleRemoteNextPlaylist` | Known | Event handler |
| 0x00633F80 | `HandleRemoteOff` | Known | Event handler |
| 0x00633FB0 | `HandleRemoteOn` | Known | Event handler |
| 0x00633FD4 | `HandleRemotePause` | Known | Event handler |
| 0x00633FFC | `HandleRemotePlay` | Known | Event handler |
| 0x00634034 | `HandleRemotePlayPausePressAndHold` | Known | Event handler |
| 0x00634070 | `HandleRemotePrevAlbum` | Known | Event handler |
| 0x006340A0 | `HandleRemotePrevChapter` | Known | Event handler |
| 0x006340D4 | `HandleRemotePrevPlaylist` | Known | Event handler |
| 0x00634104 | `HandleRemoteRepeat` | Known | Event handler |
| 0x00634128 | `HandleRemoteRewDown` | Known | Event handler |
| 0x0063414C | `HandleRemoteRewUp` | Known | Event handler |
| 0x00634174 | `HandleRemoteSelectDown` | Known | Event handler |
| 0x006341A0 | `HandleRemoteSelectUp` | Known | Event handler |
| 0x006341CC | `HandleRemoteShuffle` | Known | Event handler |
| 0x006341F4 | `HandleRemoteStop` | Known | Event handler |
| 0x00634298 | `HandleRotationChange` | Known | Event handler |
| 0x006342C4 | `HandleVideoCaptionAdvance` | Known | Event handler |
| 0x006342F0 | `HandleVideoFFDown` | Known | Event handler |
| 0x00634310 | `HandleVideoFFUp` | Known | Event handler |
| 0x00634338 | `HandleVideoNextPressAndHold` | Known | Event handler |
| 0x00634364 | `HandleVideoNext` | Known | Event handler |
| 0x0063438C | `HandleVideoNextChapter` | Known | Event handler |
| 0x006343BC | `HandleVideoNextFrame` | Known | Event handler |
| 0x006343E8 | `HandleVideoPause` | Known | Event handler |
| 0x0063440C | `HandleVideoPlay` | Known | Event handler |
| 0x00634430 | `HandleVideoPlayPause` | Known | Event handler |
| 0x00634460 | `HandleVideoPreviousPressAndHold` | Known | Event handler |
| 0x00634490 | `HandleVideoPrevious` | Known | Event handler |
| 0x006344BC | `HandleVideoPrevChapter` | Known | Event handler |
| 0x006344EC | `HandleVideoPrevFrame` | Known | Event handler |
| 0x00634514 | `HandleVideoRewDown` | Known | Event handler |
| 0x00634538 | `HandleVideoRewUp` | Known | Event handler |
| 0x0063455C | `HandleVideoStop` | Known | Event handler |
| 0x00634590 | `HandleVoiceOverAudiobookSelected` | Known | Event handler |
| 0x006345D4 | `HandleVoiceOverPlaylistSelected` | Known | Event handler |
| 0x00634614 | `HandleVoiceOverPodcastSelected` | Known | Event handler |
| 0x00634664 | `HandleMenuSelection` | Known | Event handler |
| 0x0063478C | `HandleLoadingCancelled` | Known | Event handler |
| 0x006347FC | `HandleDialogDismiss` | Known | Event handler |
| 0x00634ACC | `HandleGotoCalendarLoadingScreen` | Known | Event handler |
| 0x00634B4C | `HandleGotoAddressBookScreen` | Known | Event handler |
| 0x00634C1C | `HandleRadioSelected` | Known | Event handler |
| 0x00634C50 | `HandleRadioPlayPause` | Known | Event handler |
| 0x00634D00 | `HandleVoiceMemosSelected` | Known | Event handler |
| 0x00634D88 | `HandleSteer` | Known | Event handler |
| 0x00634DA4 | `HandleItemSelected` | Known | Event handler |
| 0x00634E54 | `HandleNextContact` | Known | Event handler |
| 0x00634E68 | `HandlePreviousContact` | Known | Event handler |
| 0x00634E80 | `HandleSelect` | Known | Event handler |
| 0x00634EA0 | `HandleHilite` | Known | Event handler |
| 0x00635264 | `HandleDateChosen` | Known | Event handler |
| 0x00635294 | `HandleTimeChosen` | Known | Event handler |
| 0x006352C4 | `HandleFrequencyChosen` | Known | Event handler |
| 0x006352F8 | `HandleSoundChosen` | Known | Event handler |
| 0x00635328 | `HandleLabelChosen` | Known | Event handler |
| 0x00635358 | `HandleDeleteChosen` | Known | Event handler |
| 0x006354D8 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00635660 | `HandleNextDay` | Known | Event handler |
| 0x00635670 | `HandlePreviousDay` | Known | Event handler |
| 0x0063570C | `HandleNextAndHold` | Known | Event handler |
| 0x00635720 | `HandlePrevAndHold` | Known | Event handler |
| 0x00635734 | `HandlePrev` | Known | Event handler |
| 0x0063575C | `HandleSelectPressAndHold` | Known | Event handler |
| 0x00635778 | `HandleCameraRemoteSelect` | Known | Event handler |
| 0x006357B0 | `HandleAppInitialized` | Known | Event handler |
| 0x006358E4 | `HandleScreenRotation` | Known | Event handler |
| 0x00635958 | `HandleWheel` | Known | Event handler |
| 0x00635988 | `HandleDeleteAllItems` | Known | Event handler |
| 0x006359C0 | `HandleDeleteItem` | Known | Event handler |
| 0x00635DE8 | `HandleRemotePlayPauseForVideo` | Known | Event handler |
| 0x00635EB0 | `HandleSelectClock` | Known | Event handler |
| 0x00635EE0 | `HandleHilited` | Known | Event handler |
| 0x00635F28 | `HandleSelectRegion_Africa` | Known | Event handler |
| 0x00635F9C | `HandleSelectRegion_Asia` | Known | Event handler |
| 0x00635FD0 | `HandleSelectRegion_Atlantic` | Known | Event handler |
| 0x00636008 | `HandleSelectRegion_Australia` | Known | Event handler |
| 0x00636044 | `HandleSelectRegion_Europe` | Known | Event handler |
| 0x0063607C | `HandleSelectRegion_NorthAmerica` | Known | Event handler |
| 0x006360B8 | `HandleSelectRegion_Pacific` | Known | Event handler |
| 0x006360F0 | `HandleSelectRegion_SouthAmerica` | Known | Event handler |
| 0x00636110 | `HandleSelectCity` | Known | Event handler |
| 0x00636144 | `HandleAddClock` | Known | Event handler |
| 0x00636178 | `HandleDeleteClock` | Known | Event handler |
| 0x006361AC | `HandleEditClock` | Known | Event handler |
| 0x006363A0 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x006363F8 | `HandleSwapToNowPlaying` | Known | Event handler |
| 0x00636410 | `HandleFlowNext` | Known | Event handler |
| 0x00636420 | `HandleFlowPrev` | Known | Event handler |
| 0x00636430 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x00636474 | `HandleAlbumSelected` | Known | Event handler |
| 0x006365AC | `HandleSwapToNowPlayingFromOrientation` | Known | Event handler |
| 0x006365D4 | `HandleFlowWheel` | Known | Event handler |
| 0x00636600 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x006367F0 | `HandleArrowDown` | Known | Event handler |
| 0x00636814 | `HandleArrowUp` | Known | Event handler |
| 0x00636890 | `HandleGameHilited` | Known | Event handler |
| 0x00636AC4 | `HandleGestureShake` | Known | Event handler |
| 0x00636B20 | `HandleOrientationLandscape` | Known | Event handler |
| 0x00636B3C | `HandleOrientationPortrait` | Known | Event handler |
| 0x00636B58 | `HandleGestureSteer` | Known | Event handler |
| 0x00636BA4 | `HandleOrientionAlt` | Known | Event handler |
| 0x00636C00 | `HandlePauseKey` | Known | Event handler |
| 0x00637050 | `HandleOrientationAlt` | Known | Event handler |
| 0x00637084 | `HandleMusicSelected` | Known | Event handler |
| 0x006370B4 | `HandleMusicHilited` | Known | Event handler |
| 0x006370E4 | `HandleCoverFlowSelected` | Known | Event handler |
| 0x00637118 | `HandleCoverflowHilited` | Known | Event handler |
| 0x0063714C | `HandleGeniusMixesSelected` | Known | Event handler |
| 0x00637184 | `HandleGeniusMixesHilited` | Known | Event handler |
| 0x006371BC | `HandlePlaylistsSelected` | Known | Event handler |
| 0x006371F0 | `HandlePlaylistsHilited` | Known | Event handler |
| 0x00637224 | `HandleArtistsSelected` | Known | Event handler |
| 0x00637258 | `HandleArtistsHilited` | Known | Event handler |
| 0x0063728C | `HandleAlbumsSelected` | Known | Event handler |
| 0x006372C0 | `HandleAlbumsHilited` | Known | Event handler |
| 0x006372F0 | `HandleCompilationsSelected` | Known | Event handler |
| 0x00637328 | `HandleCompilationsHilited` | Known | Event handler |
| 0x00637360 | `HandleSongsSelected` | Known | Event handler |
| 0x00637390 | `HandleSongsHilited` | Known | Event handler |
| 0x006373C0 | `HandleGenresSelected` | Known | Event handler |
| 0x006373F4 | `HandleGenresHilited` | Known | Event handler |
| 0x00637424 | `HandleComposersSelected` | Known | Event handler |
| 0x00637458 | `HandleComposersHilited` | Known | Event handler |
| 0x0063748C | `HandleAudiobooksSelected` | Known | Event handler |
| 0x006374C4 | `HandleAudiobooksHilited` | Known | Event handler |
| 0x0063758C | `HandleVideosSelected` | Known | Event handler |
| 0x006375C0 | `HandleVideosHilited` | Known | Event handler |
| 0x006375F0 | `HandleVideoPlaylistsSelected` | Known | Event handler |
| 0x0063762C | `HandleVideoPlaylistsHilited` | Known | Event handler |
| 0x00637664 | `HandleMoviesSelected` | Known | Event handler |
| 0x00637698 | `HandleMoviesHilited` | Known | Event handler |
| 0x006376C8 | `HandleTVShowsSelected` | Known | Event handler |
| 0x006376FC | `HandleTVShowsHilited` | Known | Event handler |
| 0x00637730 | `HandleMusicVideosSelected` | Known | Event handler |
| 0x00637768 | `HandleMusicVideosHilited` | Known | Event handler |
| 0x006377A0 | `HandleRentalsSelected` | Known | Event handler |
| 0x006377D4 | `HandleRentalsHilited` | Known | Event handler |
| 0x00637808 | `HandlePhotosSelected` | Known | Event handler |
| 0x0063783C | `HandlePhotosHilited` | Known | Event handler |
| 0x0063786C | `HandlePodcastsSelected` | Known | Event handler |
| 0x006378A0 | `HandlePodcastsHilited` | Known | Event handler |
| 0x006378D4 | `HandleiTunesUSelected` | Known | Event handler |
| 0x00637908 | `HandleiTunesUHilited` | Known | Event handler |
| 0x00637958 | `HandleRadioHilited` | Known | Event handler |
| 0x006379A4 | `HandleCameraSelected` | Known | Event handler |
| 0x006379D8 | `HandleCameraHilited` | Known | Event handler |
| 0x00637A3C | `HandleExtrasHilited` | Known | Event handler |
| 0x00637B98 | `HandleAddressBookSelected` | Known | Event handler |
| 0x00637D48 | `HandleNikePlusSelected` | Known | Event handler |
| 0x00637D7C | `HandleNikePlusHilited` | Known | Event handler |
| 0x00637E80 | `HandleVoiceMemosHilited` | Known | Event handler |
| 0x00637F2C | `HandleGenericHilited` | Known | Event handler |
| 0x00637FBC | `HandleSettingsHilited` | Known | Event handler |
| 0x00638010 | `HandleSleepSelected` | Known | Event handler |
| 0x00638078 | `HandlePedometerSelected` | Known | Event handler |
| 0x006380AC | `HandlePedometerHilited` | Known | Event handler |
| 0x006380E0 | `HandleShuffleSongsSelected` | Known | Event handler |
| 0x00638154 | `HandleNowPlaying` | Known | Event handler |
| 0x00638184 | `HandleNowPlayingHilited` | Known | Event handler |
| 0x0063819C | `HandleMainMenuPlayPause` | Known | Event handler |
| 0x00639800 | `HandlePortraitToLandscape` | Known | Event handler |
| 0x0063981C | `HandleRadioPreviewPlayPause` | Known | Event handler |
| 0x006399D8 | `HandleShowContextualMenu` | Known | Event handler |
| 0x00639A14 | `HandleAddToOTG` | Known | Event handler |
| 0x00639AE4 | `HandleGeniusPlaylistReady` | Known | Event handler |
| 0x00639B40 | `HandleStartGenius` | Known | Event handler |
| 0x00639B74 | `HandleViewAlbum` | Known | Event handler |
| 0x00639BA4 | `HandleViewArtist` | Known | Event handler |
| 0x00639BE0 | `HandleViewCompilation` | Known | Event handler |
| 0x00639F50 | `HandleRefreshPlaylist` | Known | Event handler |
| 0x00639FA0 | `HandleSaveGeniusPlaylist` | Known | Event handler |
| 0x00639FBC | `HandleSelectMix` | Known | Event handler |
| 0x00639FF4 | `HandleGeniusMixPlaylistReady` | Known | Event handler |
| 0x0063A324 | `HandleClearOTGPlaylist` | Known | Event handler |
| 0x0063A4B4 | `HandleCameraVideosSelected` | Known | Event handler |
| 0x0063A8D4 | `HandleTVOutChanged` | Known | Event handler |
| 0x0063A904 | `HandleTVSignalChanged` | Known | Event handler |
| 0x0063A938 | `HandleVideoTVWideAspectRatioSettingsChanged` | Known | Event handler |
| 0x0063A980 | `HandleWideScreenSettingsChanged` | Known | Event handler |
| 0x0063A9BC | `HandleAlternateAudioSettingsChanged` | Known | Event handler |
| 0x0063A9FC | `HandleCaptionSettingsChanged` | Known | Event handler |
| 0x0063AA38 | `HandleSubtitleSettingsChanged` | Known | Event handler |
| 0x0063AA58 | `HandleMenuLongpress` | Known | Event handler |
| 0x0063AA6C | `HandleMenuKey` | Known | Event handler |
| 0x0063AAB4 | `HandlePushContextualMenu` | Known | Event handler |
| 0x0063AAF8 | `HandleAudiobookFaster` | Known | Event handler |
| 0x0063AB38 | `HandleAudiobookNormal` | Known | Event handler |
| 0x0063AB78 | `HandleAudiobookSlower` | Known | Event handler |
| 0x0063B0A4 | `HandleSwapToMusicScreen` | Known | Event handler |
| 0x0063B0DC | `HandleSwapToVideoScreen` | Known | Event handler |
| 0x0063B0FC | `HandleMTWheel` | Known | Event handler |
| 0x0063B10C | `HandleSwapToCoverflow` | Known | Event handler |
| 0x0063B124 | `HandleDefaultOrientation` | Known | Event handler |
| 0x0063B140 | `HandleSelectProgress` | Known | Event handler |
| 0x0063B158 | `HandleWheelProgress` | Known | Event handler |
| 0x0063B16C | `HandleSelectVolume` | Known | Event handler |
| 0x0063B180 | `HandleWheelVolume` | Known | Event handler |
| 0x0063B194 | `HandleSelectGenius` | Known | Event handler |
| 0x0063B1A8 | `HandleWheelGenius` | Known | Event handler |
| 0x0063B1BC | `HandleSelectRating` | Known | Event handler |
| 0x0063B1D0 | `HandleWheelRating` | Known | Event handler |
| 0x0063B1FC | `HandleSelectScrub` | Known | Event handler |
| 0x0063B210 | `HandleWheelScrub` | Known | Event handler |
| 0x0063B224 | `HandleWheelShuffleSlider` | Known | Event handler |
| 0x0063B240 | `HandleSelectShuffleSlider` | Known | Event handler |
| 0x0063B25C | `HandleSelectExtraInfo` | Known | Event handler |
| 0x0063B290 | `HandlePlayPauseWithTransition` | Known | Event handler |
| 0x0063B2C8 | `HandleOrientationChange` | Known | Event handler |
| 0x0063B2EC | `HandleWheelBrightness` | Known | Event handler |
| 0x0063B304 | `HandleRentalWarningChoice` | Known | Event handler |
| 0x0063B320 | `HandlePlayPauseTV` | Known | Event handler |
| 0x0063B448 | `HandleCenterButtonSelected` | Known | Event handler |
| 0x0063B570 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0063B58C | `HandleNotesPop` | Known | Event handler |
| 0x0063B814 | `HandleNotesSelected` | Known | Event handler |
| 0x0063B958 | `HandlePushEvents` | Known | Event handler |
| 0x0063B9BC | `HandlePushFaces` | Known | Event handler |
| 0x0063BA18 | `HandlePushPlaces` | Known | Event handler |
| 0x0063BAC8 | `HandleBrowseAlbum` | Known | Event handler |
| 0x0063BADC | `HandleHiliteAlbum` | Known | Event handler |
| 0x0063BAF0 | `HandlePopEvents` | Known | Event handler |
| 0x0063BB00 | `HandlePopFaces` | Known | Event handler |
| 0x0063BB10 | `HandlePopPlaces` | Known | Event handler |
| 0x0063BD2C | `HandleTransitionForSlideshowChosen` | Known | Event handler |
| 0x0063BD50 | `HandleDurationForSlideshowChosen` | Known | Event handler |
| 0x0063BDE0 | `HandleImageLast` | Known | Event handler |
| 0x0063BDF0 | `HandleScreenNext` | Known | Event handler |
| 0x0063BE04 | `HandleImageFirst` | Known | Event handler |
| 0x0063BE18 | `HandleScreenPrev` | Known | Event handler |
| 0x0063BE2C | `HandleBrowseLarge` | Known | Event handler |
| 0x0063BF28 | `HandleImageNext` | Known | Event handler |
| 0x0063BF48 | `HandleImagePrev` | Known | Event handler |
| 0x0063BF58 | `HandleMenuUp` | Known | Event handler |
| 0x0063BF68 | `HandlePrevPressAndHold` | Known | Event handler |
| 0x0063C478 | `HandlePause` | Known | Event handler |
| 0x0063C484 | `HandlePlay` | Known | Event handler |
| 0x0063C4A0 | `HandleMikeyPlayPause` | Known | Event handler |
| 0x0063C4B8 | `HandleOrientationDefault` | Known | Event handler |
| 0x0063C9E0 | `HandleReadyForLargeBrowse` | Known | Event handler |
| 0x0063CAD0 | `HandleNextTuning` | Known | Event handler |
| 0x0063CAF0 | `HandlePreviousTuning` | Known | Event handler |
| 0x0063CB08 | `HandleTunerContextMenu` | Known | Event handler |
| 0x0063CE44 | `HandleMikeyNext` | Known | Event handler |
| 0x0063CE54 | `HandleMikeyPrevious` | Known | Event handler |
| 0x0063CE68 | `HandleMikeyVolume` | Known | Event handler |
| 0x0063CE7C | `HandleVolumeChange` | Known | Event handler |
| 0x0063CE90 | `HandleVolumeWheel` | Known | Event handler |
| 0x0063CF2C | `HandleBufferWheel` | Known | Event handler |
| 0x0063CF54 | `HandleTunerWheel` | Known | Event handler |
| 0x0063CF98 | `HandleBandWheel` | Known | Event handler |
| 0x0063CFA8 | `HandleNextPress` | Known | Event handler |
| 0x0063CFB8 | `HandlePreviousPress` | Known | Event handler |
| 0x0063D080 | `HandleTogglePlayPause` | Known | Event handler |
| 0x0063D0EC | `HandlePlayRadio` | Known | Event handler |
| 0x0063D118 | `HandleStopRadio` | Known | Event handler |
| 0x0063D144 | `HandleAutoTune` | Known | Event handler |
| 0x0063D2D8 | `HandleToggleBufferSetting` | Known | Event handler |
| 0x0063D310 | `HandleScanLogging` | Known | Event handler |
| 0x0063D33C | `HandleSelectPreset` | Known | Event handler |
| 0x0063D370 | `HandleConfirmation` | Known | Event handler |
| 0x0063D4E4 | `HandleExitUnsupported` | Known | Event handler |
| 0x0063D5D0 | `HandlePushToCount` | Known | Event handler |
| 0x0063D5E4 | `HandlePopToBasic` | Known | Event handler |
| 0x0063D5F8 | `HandlePushToBasic` | Known | Event handler |
| 0x0063D60C | `HandlePopToCapacity` | Known | Event handler |
| 0x0063D620 | `HandlePushToCapacity` | Known | Event handler |
| 0x0063D638 | `HandlePopToCount` | Known | Event handler |
| 0x0063D64C | `HandlePushToAccessoryCount` | Known | Event handler |
| 0x0063D668 | `HandlePopToAccessoryAccessory` | Known | Event handler |
| 0x0063D688 | `HandlePushToAccessoryBasic` | Known | Event handler |
| 0x0063D6A4 | `HandlePopToAccessoryCapacity` | Known | Event handler |
| 0x0063D6C4 | `HandlePushToAccessoryAccessory` | Known | Event handler |
| 0x0063D6E4 | `HandlePopToAccessoryCount` | Known | Event handler |
| 0x0063D700 | `HandlePushToAccessoryCapacity` | Known | Event handler |
| 0x0063D720 | `HandlePopToAccessoryBasic` | Known | Event handler |
| 0x0063DAC0 | `HandleResetAllSettings` | Known | Event handler |
| 0x0063DE44 | `HandleSetting_BacklightTimer` | Known | Event handler |
| 0x0063E04C | `HandleMainMenu` | Known | Event handler |
| 0x0063E4C8 | `HandleMusicMenu` | Known | Event handler |
| 0x0063E87C | `HandleRadioRegion` | Known | Event handler |
| 0x0063E89C | `HandleLanguagePop` | Known | Event handler |
| 0x0063E8E4 | `HandleLanguage` | Known | Event handler |
| 0x0063EC48 | `HandleSelectKey` | Known | Event handler |
| 0x0063EC58 | `HandleExit` | Known | Event handler |
| 0x0063EC64 | `HandleStartStop` | Known | Event handler |
| 0x0063ECA8 | `HandleLap` | Known | Event handler |
| 0x0063ECB4 | `HandleChosen` | Known | Event handler |
| 0x0063ED30 | `HandleDelete` | Known | Event handler |
| 0x0063F240 | `HandleSelectedNikeMainMenuItem` | Known | Event handler |
| 0x0063F83C | `HandleBasicSelected` | Known | Event handler |
| 0x0063F86C | `HandleTimedWorkoutSelected` | Known | Event handler |
| 0x0063F8A4 | `HandleDistanceWorkoutSelected` | Known | Event handler |
| 0x0063F8E0 | `HandleCaloriesWorkoutSelected` | Known | Event handler |
| 0x0063F924 | `HandleChooseTimedWorkoutFromList` | Known | Event handler |
| 0x0063F948 | `HandleChooseDistanceWorkoutFromList` | Known | Event handler |
| 0x0063F96C | `HandleChooseCalorieWorkoutFromList` | Known | Event handler |
| 0x0063F990 | `HandleBegin` | Known | Event handler |
| 0x0063FC20 | `HandleLinkNewRemote` | Known | Event handler |
| 0x0063FCB0 | `HandleLinkNewHeartMonitor` | Known | Event handler |
| 0x0063FD3C | `HandleCancelRemoteLinking` | Known | Event handler |
| 0x0063FD7C | `HandleChooseMusic` | Known | Event handler |
| 0x0063FDB0 | `HandleEndWorkout` | Known | Event handler |
| 0x0063FDE8 | `HandlePauseWorkout` | Known | Event handler |
| 0x0063FE20 | `HandleResumeWorkout` | Known | Event handler |
| 0x006401D0 | `HandleNewWorkout` | Known | Event handler |
| 0x00640230 | `HandleNewBasicWorkout` | Known | Event handler |
| 0x00640248 | `HandleNewQuickstartWorkout` | Known | Event handler |
| 0x006402D8 | `HandleChoosePedometer` | Known | Event handler |
| 0x00640370 | `HandleChoosePowerPlay` | Known | Event handler |
| 0x006403A4 | `HandleChooseVoicePrompts` | Known | Event handler |
| 0x006403DC | `HandleChooseUnit` | Known | Event handler |
| 0x00640694 | `HandleChooseRemoteSetting` | Known | Event handler |
| 0x006406CC | `HandleChooseRemoteLink` | Known | Event handler |
| 0x00640700 | `HandleChooseRemoteUnlink` | Known | Event handler |
| 0x00640864 | `HandleChooseHeartMonitorLink` | Known | Event handler |
| 0x006408A0 | `HandleChooseHeartMonitorUnlink` | Known | Event handler |
| 0x00640958 | `HandleListChoose` | Known | Event handler |
| 0x0064096C | `HandlePopBackToSongsScreen` | Known | Event handler |
| 0x006409A4 | `HandleVerticalSelected` | Known | Event handler |
| 0x006409D8 | `HandleRightSelected` | Known | Event handler |
| 0x00640A08 | `HandleLeftSelected` | Known | Event handler |
| 0x00640AAC | `HandleChooseLink` | Known | Event handler |
| 0x00640B44 | `HandleMenuKeyNop` | Known | Event handler |
| 0x00640B7C | `HandlePauseHold` | Known | Event handler |
| 0x00640BA0 | `HandleSelectKeyDown` | Known | Event handler |
| 0x00640BB4 | `HandlePowerPlay` | Known | Event handler |
| 0x00640BC4 | `HandleOnDemandPrompt` | Known | Event handler |
| 0x006411A8 | `Handle3BitModeFinished` | Known | Event handler |
| 0x006411C0 | `HandleMikeyPressExtended` | Known | Event handler |
| 0x00641668 | `HandleNowPlayingSelected` | Known | Event handler |
| 0x006416A0 | `HandleChoosePlaylistsPlay` | Known | Event handler |
| 0x006416D8 | `HandleChoosePodcastsPlay` | Known | Event handler |
| 0x00641710 | `HandleChooseAudiobooksPlay` | Known | Event handler |
| 0x00641764 | `HandleNoneSelected` | Known | Event handler |
| 0x00641A8C | `HandleSelectedDayWorkout` | Known | Event handler |
| 0x00641AF8 | `HandleClearBests` | Known | Event handler |
| 0x00641B5C | `HandleClearTotals` | Known | Event handler |
| 0x00641B70 | `HandleHerculesKey` | Known | Event handler |
| 0x00641B84 | `HandlePopSelf` | Known | Event handler |
| 0x00641B94 | `HandlePressAndHold` | Known | Event handler |
| 0x00641C7C | `HandleSelectNextWorkout` | Known | Event handler |
| 0x00641C94 | `HandleSelectPrevWorkout` | Known | Event handler |
| 0x00641D24 | `HandleDeleteAllWorkouts` | Known | Event handler |
| 0x00641D60 | `HandleDeleteWorkout` | Known | Event handler |
| 0x00641F54 | `HandleNextKey` | Known | Event handler |
| 0x00641F64 | `HandlePrevKey` | Known | Event handler |
| 0x00642150 | `HandleWeightSelect` | Known | Event handler |
| 0x00642164 | `HandleWeightWheel` | Known | Event handler |
| 0x00642178 | `HandleWeightSelectAltTrans` | Known | Event handler |
| 0x00642194 | `HandleWeightSelectPedometer` | Known | Event handler |
| 0x006421B0 | `HandleDistanceSelect` | Known | Event handler |
| 0x006421C8 | `HandleDistanceWheel` | Known | Event handler |
| 0x006421DC | `HandleTimeSelect` | Known | Event handler |
| 0x006421F0 | `HandleTimeWheel` | Known | Event handler |
| 0x00642200 | `HandleCaloriesSelect` | Known | Event handler |
| 0x00642218 | `HandleCaloriesWheel` | Known | Event handler |
| 0x0064222C | `HandleStepGoalSelect` | Known | Event handler |
| 0x00642244 | `HandleStepGoalWheel` | Known | Event handler |
| 0x006422F0 | `HandleWalkCalibrationSelection` | Known | Event handler |
| 0x0064232C | `HandleRunCalibrationSelection` | Known | Event handler |
| 0x00642418 | `HandleResetRuningCalibration` | Known | Event handler |
| 0x00642438 | `HandleResetWalkingCalibration` | Known | Event handler |
| 0x00642734 | `HandleUnlinkRemote` | Known | Event handler |
| 0x00642784 | `HandleUnlinkHeartMonitor` | Known | Event handler |
| 0x00642BDC | `HandleShowRecordings` | Known | Event handler |
| 0x00642BF4 | `HandleAddChapterMark` | Known | Event handler |
| 0x00642E74 | `HandlePlaySelectIndexedRecording` | Known | Event handler |
| 0x00642F64 | `HandleSelectIndexedRecording` | Known | Event handler |
| 0x00642FBC | `HandlePlayCurrentRecording` | Known | Event handler |
| 0x00643070 | `HandleSelectLabel` | Known | Event handler |
| 0x00643084 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x0064309C | `HandleDeleteSelect` | Known | Event handler |
| 0x006430B0 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x006430CC | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x006430EC | `HandleMicrophoneRequired` | Known | Event handler |
| 0x00643108 | `HandleMicrophoneDisconnected` | Known | Event handler |
| 0x006431AC | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x0064323C | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x00643258 | `HandleEQSettingSelected` | Known | Event handler |
| 0x00643270 | `HandleEQSettingPreview` | Known | Event handler |

---

## 6. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00071CE8 | `GotoNowPlaying` | Known | Navigation |
| 0x00077070 | `GotoDefaultLayout` | Known | Navigation |
| 0x000770DC | `GotoVolumeLayout` | Known | Navigation |
| 0x00077764 | `GotoStatusBarLayout` | Known | Navigation |
| 0x00077778 | `GotoDefaultLayout` | Known | Navigation |
| 0x000778C8 | `GotoDefault` | Known | Navigation |
| 0x000779FC | `GotoProgressLayout` | Known | Navigation |
| 0x00077B4C | `GotoBrightnessLayout` | Known | Navigation |
| 0x00077BB0 | `GotoBrightnessLayout` | Known | Navigation |
| 0x00077C08 | `GotoVolumeLayout` | Known | Navigation |
| 0x00077C40 | `GotoScrubLayout` | Known | Navigation |
| 0x00077CD8 | `GotoStatusBarLayout` | Known | Navigation |
| 0x00077CEC | `GotoDefaultLayout` | Known | Navigation |
| 0x00077D48 | `GotoScrubLayout` | Known | Navigation |
| 0x00077D84 | `GotoScrubLayout` | Known | Navigation |
| 0x000781C8 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x000781E4 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00078200 | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x0007821C | `GotoDefaultLayout` | Known | Navigation |
| 0x00078294 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x000782AC | `GotoVolumeLayout` | Known | Navigation |
| 0x000789B8 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x00078B04 | `GotoStatusBarVideoLayout` | Known | Navigation |
| 0x00078B20 | `GotoDefaultVideoLayout` | Known | Navigation |
| 0x00078B3C | `GotoDefaultSubtitlesLayout` | Known | Navigation |
| 0x00078B58 | `GotoDefaultLayout` | Known | Navigation |
| 0x0008430C | `GotoMainMenu` | Known | Navigation |
| 0x0008907C | `GotoDialogueScreen` | Known | Navigation |
| 0x000A4F94 | `GotoNowPlaying` | Known | Navigation |
| 0x000A4FA8 | `GotoAlbums` | Known | Navigation |
| 0x000A4FB4 | `GotoSongs` | Known | Navigation |
| 0x000B0F90 | `GotoNowPlaying` | Known | Navigation |
| 0x000B11C8 | `GotoNowPlaying` | Known | Navigation |
| 0x000B26DC | `GotoScreen_PlaybackSettingsMenu` | Known | Navigation |
| 0x000B34D0 | `GotoMediumPedometerLayout` | Known | Navigation |
| 0x000B4680 | `GotoProgressLayout` | Known | Navigation |
| 0x000B472C | `GotoProgressLayout` | Known | Navigation |
| 0x000B4818 | `GotoProgressLayout` | Known | Navigation |
| 0x000B4AC8 | `GotoProgressLayout` | Known | Navigation |
| 0x000B5028 | `GotoMainCalendarPage` | Known | Navigation |
| 0x000B5FCC | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |
| 0x000B60D4 | `GotoFourCard_About` | Known | Navigation |
| 0x000B60E8 | `GotoThreeCard_About` | Known | Navigation |
| 0x000B6350 | `GotoScreen_ResetAllSettings` | Known | Navigation |
| 0x000B659C | `GotoScreen_BacklightTimer` | Known | Navigation |
| 0x000B660C | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation |
| 0x000B6630 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x000B6CC0 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x000B6CD8 | `GotoProgressLayout` | Known | Navigation |
| 0x000B6F78 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x000B6F90 | `GotoProgressLayout` | Known | Navigation |
| 0x000B7060 | `GotoProgressLayout` | Known | Navigation |
| 0x000B7180 | `GotoBrightnessVideoLayout` | Known | Navigation |
| 0x000B719C | `GotoGeniusLayout` | Known | Navigation |
| 0x000B71B0 | `GotoRatingLayout` | Known | Navigation |
| 0x000B731C | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x000B73D0 | `GotoRatingLayout` | Known | Navigation |
| 0x000B74B0 | `GotoShuffleLayout` | Known | Navigation |
| 0x000B74E8 | `GotoDefaultLayout` | Known | Navigation |
| 0x000B7798 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x000B77B0 | `GotoVolumeLayout` | Known | Navigation |
| 0x000B7838 | `GotoVolumeVideoLayout` | Known | Navigation |
| 0x000B7850 | `GotoVolumeLayout` | Known | Navigation |
| 0x000B794C | `GotoScrubVideoLayout` | Known | Navigation |
| 0x000B7964 | `GotoScrubLayout` | Known | Navigation |
| 0x000B79D8 | `GotoProgressVideoLayout` | Known | Navigation |
| 0x000B79F0 | `GotoProgressLayout` | Known | Navigation |
| 0x000B7F1C | `GotoProgressVideoLayout` | Known | Navigation |
| 0x000B7F34 | `GotoProgressLayout` | Known | Navigation |
| 0x000B7FC4 | `GotoExtraInfoLoadFailedLayout` | Known | Navigation |
| 0x000B7FE4 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x000B8B9C | `GotoScreen_AddressViewerLoaded` | Known | Navigation |
| 0x000C2E14 | `GotoScreen_LockediPod` | Known | Navigation |
| 0x000C2E2C | `GotoScreen_EnterPassKey` | Known | Navigation |
| 0x000C4CB4 | `GotoCoverFlowScreenBackside` | Known | Navigation |
| 0x000C9B4C | `GotoScreenMainMenu` | Known | Navigation |
| 0x000CB4B8 | `GotoScreen_Language` | Known | Navigation |
| 0x000CC07C | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation |
| 0x000E0D98 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation |
| 0x000E0DBC | `GotoScreen_LockDialog` | Known | Navigation |
| 0x000E10EC | `GotoScreen_AddressViewerLoading` | Known | Navigation |
| 0x000E110C | `GotoScreen_AddressViewerLoaded` | Known | Navigation |
| 0x000E1168 | `GotoScreen_CalendarViewerLoading` | Known | Navigation |
| 0x000E118C | `GotoScreen_CalendarView` | Known | Navigation |
| 0x000E2BA0 | `GotoScreen_LockediPod` | Known | Navigation |
| 0x000E2D94 | `GotoScreen_MainMenu` | Known | Navigation |
| 0x000EB710 | `GotoScreen_VolumeLimit` | Known | Navigation |
| 0x000ED20C | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x000EE2BC | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x000EE3CC | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x000EE530 | `GotoPlaylistScreen` | Known | Navigation |
| 0x000EE748 | `GotoNowPlaying` | Known | Navigation |
| 0x000EE950 | `GotoNowPlaying` | Known | Navigation |
| 0x000F13AC | `GotoFirstBoot` | Known | Navigation |
| 0x000F13BC | `GotoNotesApp` | Known | Navigation |
| 0x000F13D0 | `GotoLockApp` | Known | Navigation |
| 0x000F1EA0 | `GotoPlaylists` | Known | Navigation |
| 0x000F23BC | `GotoGenius` | Known | Navigation |
| 0x000F25E4 | `GotoGenius` | Known | Navigation |
| 0x000F25F0 | `GotoGeniusIntro` | Known | Navigation |
| 0x000F30C0 | `GotoNowPlaying` | Known | Navigation |
| 0x000F30F0 | `GotoGeniusMixLoadingScreen` | Known | Navigation |
| 0x000F3BFC | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x000F3D2C | `GotoGeniusLoadingScreen` | Known | Navigation |
| 0x000FAB20 | `GotoGameRunningLayout` | Known | Navigation |
| 0x000FACAC | `GotoGameRunningLayout` | Known | Navigation |
| 0x0010A8B0 | `GotoNowPlaying` | Known | Navigation |
| 0x0010ABE4 | `GotoNowPlaying` | Known | Navigation |
| 0x0010AC44 | `GotoScreen_AddressViewerLoading` | Known | Navigation |
| 0x0010AC64 | `GotoScreen_AddressViewerLoaded` | Known | Navigation |
| 0x0010AC84 | `GotoScreen_AddressViewerNoContacts` | Known | Navigation |
| 0x0010ADC8 | `GotoGeniusMixesIntro` | Known | Navigation |
| 0x0010B0F4 | `GotoNowPlaying` | Known | Navigation |
| 0x00129104 | `GotoErrorLayout` | Known | Navigation |
| 0x001B3CB8 | `GotoCalibrateRcvMissing` | Known | Navigation |
| 0x001B3D34 | `GotoCalibrateRcvMissing` | Known | Navigation |
| 0x001B4EF0 | `GotoSettings` | Known | Navigation |
| 0x001B4F00 | `GotoCustomStepGoal` | Known | Navigation |
| 0x001B8334 | `GotoRecording` | Known | Navigation |
| 0x0020FF10 | `GotoExtraInfoLayout` | Known | Navigation |
| 0x0020FF24 | `GotoExtraInfoLoadingLayout` | Known | Navigation |
| 0x00211C08 | `GotoDefault` | Known | Navigation |
| 0x00216010 | `GotoNowPlaying` | Known | Navigation |
| 0x00246EE4 | `GotoDefaultLayout` | Known | Navigation |
| 0x0024D054 | `GotoEnteringNowPlaying` | Known | Navigation |
| 0x00254EB0 | `GotoRentalWarningLayout` | Known | Navigation |
| 0x0028F860 | `GotoNowPlaying` | Known | Navigation |
| 0x00297F50 | `GotoNowPlaying` | Known | Navigation |
| 0x006364FC | `GotoCoverFlowScreenBackside` | Known | Navigation |
| 0x0063B2E0 | `GotoDefault` | Known | Navigation |
| 0x0063DC30 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation |

---

## 7. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0056E0BB | `VoiceMemos_Screen_InsufficientDiskSpace2` | Known | Screen layout |
| 0x0056E0E4 | `NikePlus_EndWorkout_Screen_Contextual_Default_L` | Known | Screen layout |
| 0x0056E114 | `MainMenus_Main_Screen_NoCamera` | Known | Screen layout |
| 0x0056E133 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0056E14B | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0056E169 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x0056E18D | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x0056E1AE | `NowPlaying_Screen_Music` | Known | Screen layout |
| 0x0056E1C6 | `NoContent_Screen_Music` | Known | Screen layout |
| 0x0056E1DD | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0056E1FB | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x0056E214 | `NikePlus_NowRunning_Screen_BasicNoMusic` | Known | Screen layout |
| 0x0056E23C | `NikePlus_NowRunning_Screen_DistanceNoMusic` | Known | Screen layout |
| 0x0056E267 | `NikePlus_NowRunning_Screen_TimeNoMusic` | Known | Screen layout |
| 0x0056E28E | `NikePlus_NowRunning_Screen_Basic_LandscapeNoMusic` | Known | Screen layout |
| 0x0056E2C0 | `NikePlus_NowRunning_Screen_Distance_LandscapeNoMusic` | Known | Screen layout |
| 0x0056E2F5 | `NikePlus_NowRunning_Screen_Time_LandscapeNoMusic` | Known | Screen layout |
| 0x0056E326 | `NikePlus_NowRunning_Screen_Calibrate_LandscapeNoMusic` | Known | Screen layout |
| 0x0056E35C | `NikePlus_NowRunning_Screen_Calories_LandscapeNoMusic` | Known | Screen layout |
| 0x0056E391 | `NikePlus_NowRunning_Screen_CalibrateNoMusic` | Known | Screen layout |
| 0x0056E3BD | `NoContent_Screen_MainNoMusic` | Known | Screen layout |
| 0x0056E3DA | `NikePlus_NowRunning_Screen_CaloriesNoMusic` | Known | Screen layout |
| 0x0056E405 | `NikePlus_Custom_Screen_Weight_ToWorkoutMusic` | Known | Screen layout |
| 0x0056E432 | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x0056E45D | `DeleteRental_Confirmation_Screen_Unwatched` | Known | Screen layout |
| 0x0056E488 | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x0056E4A6 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0056E5B8 | `NowPlaying_Screen_ExtraInfoLoadFailed` | Known | Screen layout |
| 0x0056E5DE | `Stopwatch_Screen_Stopped` | Known | Screen layout |
| 0x0056E5F7 | `VoiceMemos_Screen_MicrophoneRequired` | Known | Screen layout |
| 0x0056E61C | `PhotosGL_Camera_Screen_Paused` | Known | Screen layout |
| 0x0056E63A | `PhotosGL_Camera_Alt_Screen_Paused` | Known | Screen layout |
| 0x0056E65C | `PhotosGL_TvOut_Screen_Paused` | Known | Screen layout |
| 0x0056E679 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x0056E697 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x0056E6B7 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x0056E6DA | `VoiceMemos_Screen_MicrophoneDisconnected` | Known | Screen layout |
| 0x0056E703 | `VoiceMemos_Status_Screen_Inserted` | Known | Screen layout |
| 0x0056E725 | `VoiceMemos_Screen_Recording_ChapterInserted` | Known | Screen layout |
| 0x0056E751 | `Camera_Screen_Uninitialized` | Known | Screen layout |
| 0x0056E76D | `Search_Main_Screen_WithKeyboard` | Known | Screen layout |
| 0x0056E78D | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x0056E7AB | `VoiceMemos_Screen_InsufficientDiskSpace` | Known | Screen layout |
| 0x0056E7D3 | `NikePlus_NowRunning_Screen_Distance` | Known | Screen layout |
| 0x0056E7F7 | `NikePlus_Custom_Screen_Distance` | Known | Screen layout |
| 0x0056E817 | `NikePlus_Custom_Screen_Simple_CalibrationDistance` | Known | Screen layout |
| 0x0056E849 | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x0056E874 | `CoverFlow_Screen_Backside` | Known | Screen layout |
| 0x0056E88E | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x0056E8AB | `NowPlaying_Screen_TVOut_ConnectCable` | Known | Screen layout |
| 0x0056E8D0 | `PhotosGL_Camera_Screen_TvOut_ConnectCable` | Known | Screen layout |
| 0x0056E90A | `VoiceMemos_Screen_Idle` | Known | Screen layout |
| 0x0056E921 | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x0056E93B | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x0056E971 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0056E99D | `NikePlus_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x0056E9C7 | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x0056E9EF | `NikePlus_NowRunning_Screen_Time` | Known | Screen layout |
| 0x0056EA0F | `NikePlus_Custom_Screen_Time` | Known | Screen layout |
| 0x0056EA2B | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0056EA4D | `PhotosGL_Camera_Screen_Volume` | Known | Screen layout |
| 0x0056EA6B | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x0056EA84 | `PhotosGL_Camera_Alt_Screen_Volume` | Known | Screen layout |
| 0x0056EAA6 | `PhotosGL_TvOut_Screen_Volume` | Known | Screen layout |
| 0x0056EAC3 | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0056EAE2 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x0056EB07 | `NikePlus_StartWorkout_Screen_Resume` | Known | Screen layout |
| 0x0056EB2B | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x0056EB55 | `NikePlus_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x0056EB7D | `NikePlus_NowRunning_Screen_Basic_Landscape` | Known | Screen layout |
| 0x0056EBA8 | `NikePlus_NowRunning_Screen_Distance_Landscape` | Known | Screen layout |
| 0x0056EBD6 | `NikePlus_NowRunning_Screen_Time_Landscape` | Known | Screen layout |
| 0x0056EC00 | `NikePlus_NowRunning_Screen_Calibrate_Landscape` | Known | Screen layout |
| 0x0056EC2F | `NikePlus_EndWorkout_Screen_Contextual_Landscape` | Known | Screen layout |
| 0x0056EC5F | `NikePlus_EndWorkout_Screen_Calibration_Contextual_Landscape` | Known | Screen layout |
| 0x0056EC9B | `Alarms_Alarm_Clock_Triggered_Screen_Landscape` | Known | Screen layout |
| 0x0056ECC9 | `Alarms_Alarm_Triggered_Screen_Landscape` | Known | Screen layout |
| 0x0056ECF1 | `Nike_Volume_Screen_Landscape` | Known | Screen layout |
| 0x0056ED0E | `NikePlus_NowRunning_Screen_Landscape` | Known | Screen layout |
| 0x0056ED33 | `NikePlus_NowRunning_Screen_Calories_Landscape` | Known | Screen layout |
| 0x0056EE46 | `RemoteUI_Screen_DisplayImage_With_Unsupported_Firewire` | Known | Screen layout |
| 0x0056EE7D | `RemoteUI_Screen_Main_With_Unsupported_Firewire` | Known | Screen layout |
| 0x0056EEAC | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x0056EED1 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x0056EEFB | `Camera_Screen_Active` | Known | Screen layout |
| 0x0056EF10 | `Camera_Screen_ForcedOff` | Known | Screen layout |
| 0x0056EF28 | `NikePlus_CalibrationComplete_Screen_Pacing` | Known | Screen layout |
| 0x0056EF53 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x0056EF71 | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0056EF94 | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x0056EFB0 | `RemoteUI_Hercules_ScreenLayout_Recording` | Known | Screen layout |
| 0x0056EFD9 | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |
| 0x0056F059 | `NikePlus_Calibrate_ResetToDefault_Screen_Walking` | Known | Screen layout |
| 0x0056F08A | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x0056F0A3 | `NikePlus_Calibrate_ResetToDefault_Screen_Running` | Known | Screen layout |
| 0x0056F0D4 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0056F0FA | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x0056F120 | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x0056F139 | `CoverFlow_Screen_Exiting` | Known | Screen layout |
| 0x0056F173 | `VoiceMemos_Screen_Saving` | Known | Screen layout |
| 0x0056F18C | `PhotosGL_Camera_Screen_Playing` | Known | Screen layout |
| 0x0056F1AB | `PhotosGL_Camera_Alt_Screen_Playing` | Known | Screen layout |
| 0x0056F1CE | `PhotosGL_TvOut_Screen_Playing` | Known | Screen layout |
| 0x0056F1EC | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x0056F20D | `CoverFlow_Screen_EnteringNowPlaying` | Known | Screen layout |
| 0x0056F231 | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x0056F255 | `NikePlus_NowRunning_Screen_BasicPowerSong` | Known | Screen layout |
| 0x0056F27F | `NikePlus_NowRunning_Screen_DistancePowerSong` | Known | Screen layout |
| 0x0056F2AC | `NikePlus_NowRunning_Screen_TimePowerSong` | Known | Screen layout |
| 0x0056F2D5 | `NikePlus_NowRunning_Screen_Basic_LandscapePowerSong` | Known | Screen layout |
| 0x0056F309 | `NikePlus_NowRunning_Screen_Distance_LandscapePowerSong` | Known | Screen layout |
| 0x0056F340 | `NikePlus_NowRunning_Screen_Time_LandscapePowerSong` | Known | Screen layout |
| 0x0056F373 | `NikePlus_NowRunning_Screen_Calories_LandscapePowerSong` | Known | Screen layout |
| 0x0056F3AA | `NikePlus_NowRunning_Screen_CaloriesPowerSong` | Known | Screen layout |
| 0x0056F3D7 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x0056F3F1 | `MainMenu_Main_Screen_NoContentSearch` | Known | Screen layout |
| 0x0056F416 | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x0056F436 | `Genius_Error_Screen_NoGeniusInfoForTrack` | Known | Screen layout |
| 0x0056F45F | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x0056F47A | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x0056F4B5 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x0056F4D7 | `MediaLists_MusicVideos_Songs_Screen_WithAlbumAndArtwork` | Known | Screen layout |
| 0x0056F50F | `MediaLists_Songs_Screen_WithAlbumAndArtwork` | Known | Screen layout |
| 0x0056F53B | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x0056F568 | `PhotosGL_Camera_Screen_TvOut_Ask` | Known | Screen layout |
| 0x0056F589 | `VoiceMemos_Screen_DeleteAsk` | Known | Screen layout |
| 0x0056F5A5 | `VoiceMemos_Screen_DeleteAllAsk` | Known | Screen layout |
| 0x0056F5C4 | `NowPlaying_Screen_TVOutAsk` | Known | Screen layout |
| 0x0056F5DF | `NowPlaying_Screen_Initial` | Known | Screen layout |
| 0x0056F5F9 | `NikePlus_EndWorkout_Screen_Contextual` | Known | Screen layout |
| 0x0056F61F | `NikePlus_History_Screen_Contextual` | Known | Screen layout |
| 0x0056F677 | `Camera_Screen_DiskFull` | Known | Screen layout |
| 0x0056F68E | `MediaLists_Songs_Screen_WithAlbum` | Known | Screen layout |
| 0x0056F6B0 | `RemoteUI_Screen` | Known | Screen layout |
| 0x0056F6C0 | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x0056F6D8 | `MediaLists_iTunesU_Screen` | Known | Screen layout |
| 0x0056F6F2 | `MediaLists_Camera_Local_Media_Screen` | Known | Screen layout |
| 0x0056F717 | `Radio_InformationalOverlay_NoAntenna_Screen` | Known | Screen layout |
| 0x0056F743 | `PhotosGL_Camera_Screen` | Known | Screen layout |
| 0x0056F75A | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x0056F771 | `NikePlus_Workout_Music_Screen` | Known | Screen layout |
| 0x0056F78F | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x0056F7B3 | `AddressViewer_PartialLoad_Screen` | Known | Screen layout |
| 0x0056F7D4 | `NikePlus_IsLinked_Screen` | Known | Screen layout |
| 0x0056F7ED | `NikePlus_ActivityStopped_Screen` | Known | Screen layout |
| 0x0056F80D | `Alarms_Alarm_Clock_Triggered_Screen` | Known | Screen layout |
| 0x0056F831 | `Alarms_Alarm_Triggered_Screen` | Known | Screen layout |
| 0x0056F84F | `Pedometer_Trainer_Paused_Screen` | Known | Screen layout |
| 0x0056F86F | `Radio_InformationalOverlay_AccessoryConnected_Screen` | Known | Screen layout |
| 0x0056F8A4 | `Firewire_Charging_Unsupported_Screen` | Known | Screen layout |
| 0x0056F8C9 | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x0056F8E7 | `LockediPod_Screen` | Known | Screen layout |
| 0x0056F8F9 | `DiskMode_Screen` | Known | Screen layout |
| 0x0056F909 | `DemoMode_Screen` | Known | Screen layout |
| 0x0056F919 | `Notes_Image_Screen` | Known | Screen layout |
| 0x0056F92C | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x0056F94A | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x0056F960 | `NowPlaying_Idle_Screen` | Known | Screen layout |
| 0x0056F977 | `Game_Screen` | Known | Screen layout |
| 0x0056F983 | `NikePlus_Deleteme_Screen` | Known | Screen layout |
| 0x0056F99C | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x0056F9B9 | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x0056F9D2 | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x0056F9F3 | `Nike_Volume_Screen` | Known | Screen layout |
| 0x0056FA06 | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x0056FA2B | `NikePlus_NowRunning_Idle_Landscape_Screen` | Known | Screen layout |
| 0x0056FA55 | `Pedometer_Main_Landscape_Screen` | Known | Screen layout |
| 0x0056FA75 | `Pedometer_Ambient_Landscape_Screen` | Known | Screen layout |
| 0x0056FA98 | `NikePlus_Daily_landscape_Screen` | Known | Screen layout |
| 0x0056FAB8 | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x0056FAD5 | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x0056FAF6 | `PhotosRotate_Screen` | Known | Screen layout |
| 0x0056FB0A | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x0056FB2F | `ContextualMenu_ThreeItem_White_Screen` | Known | Screen layout |
| 0x0056FB55 | `ContextualMenu_FiveItem_White_Screen` | Known | Screen layout |
| 0x0056FB7A | `ContextualMenu_TwoItem_White_Screen` | Known | Screen layout |
| 0x0056FB9E | `ContextualMenu_FourItem_White_Screen` | Known | Screen layout |
| 0x0056FBC3 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x0056FBDA | `Calendar_Loading_Screen` | Known | Screen layout |
| 0x0056FBF2 | `AddressViewer_Loading_Screen` | Known | Screen layout |
| 0x0056FC0F | `Notes_Loading_Screen` | Known | Screen layout |
| 0x0056FC24 | `GeniusMixes_Loading_Screen` | Known | Screen layout |
| 0x0056FC3F | `Genius_Loading_Screen` | Known | Screen layout |
| 0x0056FC55 | `NikePlus_SensorSearching_Screen` | Known | Screen layout |
| 0x0056FC75 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x0056FC94 | `NikePlus_HeartMonitor_Linking_Screen` | Known | Screen layout |
| 0x0056FCB9 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x0056FCD1 | `NikePlus_Remote_Unlinking_Screen` | Known | Screen layout |
| 0x0056FCF2 | `NikePlus_HeartMonitor_Unlinking_Screen` | Known | Screen layout |
| 0x0056FD19 | `SettingsMenus_AdjustScrolling_Screen` | Known | Screen layout |
| 0x0056FD3E | `Game_Running_Screen` | Known | Screen layout |
| 0x0056FD52 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x0056FD6D | `Radio_NowPlaying_Screen` | Known | Screen layout |
| 0x0056FD85 | `NikePlus_SimpleCalibration_Walk_Dialog_Screen` | Known | Screen layout |
| 0x0056FDB3 | `NikePlus_DeleteAllWorkouts_Confirmation_Dialog_Screen` | Known | Screen layout |
| 0x0056FDE9 | `NikePlus_SimpleCalibration_Dialog_Screen` | Known | Screen layout |
| 0x0056FE12 | `NikePlus_SimpleCalibration_Run_Dialog_Screen` | Known | Screen layout |
| 0x0056FE3F | `Stopwatch_Screen` | Known | Screen layout |
| 0x0056FE50 | `NikePlus_History_WorkoutGraph_Screen` | Known | Screen layout |
| 0x0056FE75 | `SettingsMenus_Playback_Screen` | Known | Screen layout |
| 0x0056FE93 | `ContextualMenu_ThreeItem_Black_Screen` | Known | Screen layout |
| 0x0056FEB9 | `ContextualMenu_FiveItem_Black_Screen` | Known | Screen layout |
| 0x0056FEDE | `ContextualMenu_TwoItem_Black_Screen` | Known | Screen layout |
| 0x0056FF02 | `ContextualMenu_FourItem_Black_Screen` | Known | Screen layout |
| 0x0056FF27 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0056FF3E | `Clock_Screen` | Known | Screen layout |
| 0x0056FF4B | `Settings_Legal_Screen` | Known | Screen layout |
| 0x0056FF61 | `NikePlus_HeartMonitor_LinkingInitial_Screen` | Known | Screen layout |
| 0x0056FF8D | `Pedometer_Step_Goal_Screen` | Known | Screen layout |
| 0x0056FFA8 | `NikePlus_Custom_StepGoal_Screen` | Known | Screen layout |
| 0x0056FFC8 | `SettingsMenus_General_Screen` | Known | Screen layout |
| 0x0056FFE5 | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x00570003 | `Radio_InformationalOverlay_BufferFull_Screen` | Known | Screen layout |
| 0x00570030 | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x0057004C | `ToDo_Item_Screen` | Known | Screen layout |
| 0x0057005D | `PhotosZoom_Screen` | Known | Screen layout |
| 0x0057006F | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x00570086 | `MediaLists_OTG_ClearConfirm_Screen` | Known | Screen layout |
| 0x005700A9 | `Search_Main_Screen` | Known | Screen layout |
| 0x005700BC | `Location_Main_Screen` | Known | Screen layout |
| 0x005700D1 | `Pedometer_Main_Screen` | Known | Screen layout |
| 0x005700E7 | `AddressViewer_Main_Screen` | Known | Screen layout |
| 0x00570101 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x00570116 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x0057012C | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00570146 | `Clock_Region_Screen` | Known | Screen layout |
| 0x0057015A | `RentalDeleted_Notification_Screen` | Known | Screen layout |
| 0x0057017C | `SingleRentalExpiring_Notification_Screen` | Known | Screen layout |
| 0x005701A5 | `MultipleRentalsExpiring_Notification_Screen` | Known | Screen layout |
| 0x005701D1 | `RentalError_Notification_Screen` | Known | Screen layout |
| 0x005701F1 | `DeleteRental_Confirmation_Screen` | Known | Screen layout |
| 0x00570212 | `Stopwatch_DeleteConfirmation_Screen` | Known | Screen layout |
| 0x00570236 | `NikePlus_SimpleCalibration_Screen` | Known | Screen layout |
| 0x00570258 | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x00570286 | `NikePlus_StartCalibration_Screen` | Known | Screen layout |
| 0x005702A7 | `SettingsMenus_ShakeAdjust_Duration_Screen` | Known | Screen layout |
| 0x005702D1 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x005702EF | `Radio_ConfirmationOverlay_ChangeStation_Screen` | Known | Screen layout |
| 0x0057031E | `Hercules_Connection_Screen` | Known | Screen layout |
| 0x00570339 | `RentalInfo_Screen` | Known | Screen layout |
| 0x0057034B | `Genius_Intro_Screen` | Known | Screen layout |
| 0x0057035F | `NikePlus_Calendar_Screen` | Known | Screen layout |
| 0x00570378 | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x00570392 | `ChargingMode_LowPower_Screen` | Known | Screen layout |
| 0x005703AF | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x005703C9 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x005703E3 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x005703FD | `Genius_Error_Screen` | Known | Screen layout |
| 0x00570411 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x0057042A | `NikePlus_CalibrationCompleteError_Screen` | Known | Screen layout |
| 0x00570453 | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x0057046A | `NikePlus_HeartMonitor_Screen` | Known | Screen layout |
| 0x00570487 | `Extras_Screen` | Known | Screen layout |
| 0x00570495 | `PhotoBrowseThumbs_Screen` | Known | Screen layout |
| 0x005704AE | `Photos_Faces_Screen` | Known | Screen layout |
| 0x005704C2 | `Photos_Places_Screen` | Known | Screen layout |
| 0x005704D7 | `MediaLists_iTunesUEpisodes_Screen` | Known | Screen layout |
| 0x005704F9 | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x00570516 | `Nike_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x00570532 | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x00570554 | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x0057056D | `RemoteUI_Hercules_Screen` | Known | Screen layout |
| 0x00570586 | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x005705A4 | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x005705BD | `MediaLists_GeniusMixes_Screen` | Known | Screen layout |
| 0x005705DB | `Video_Settings_Screen` | Known | Screen layout |
| 0x005705F1 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x0057060A | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x00570630 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x00570646 | `MediaLists_MusicVideos_Songs_Screen` | Known | Screen layout |
| 0x0057066A | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x00570682 | `NikePlus_Songs_Screen` | Known | Screen layout |
| 0x00570698 | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x005706BB | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x005706D8 | `NikePlus_Audiobooks_Screen` | Known | Screen layout |
| 0x005706F3 | `MediaLists_Rentals_Screen` | Known | Screen layout |
| 0x0057070D | `NikePlus_History_Totals_Screen` | Known | Screen layout |
| 0x0057072C | `NikePlus_History_ClearTotals_Screen` | Known | Screen layout |
| 0x00570750 | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x00570774 | `Game_Controls_Screen` | Known | Screen layout |
| 0x00570789 | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x005707A2 | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x005707C4 | `ToDo_Instructions_Screen` | Known | Screen layout |
| 0x005707DD | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x005707F9 | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x00570813 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x00570834 | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x00570850 | `MediaLists_Camera_All_Videos_Screen` | Known | Screen layout |
| 0x00570874 | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x0057088C | `VoiceMemos_Screen` | Known | Screen layout |
| 0x0057089E | `No_Photos_Screen` | Known | Screen layout |
| 0x005708AF | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x005708C9 | `AddressViewer_ContactGroups_Screen` | Known | Screen layout |
| 0x005708EC | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x00570908 | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x0057092C | `NikePlus_AudiobookChapters_Screen` | Known | Screen layout |
| 0x0057094E | `NikePlus_CalibrationCompleteSuccess_Screen` | Known | Screen layout |
| 0x00570979 | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x00570999 | `ChargingMode_Progress_Screen` | Known | Screen layout |
| 0x005709B6 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x005709CC | `Photos_Events_Screen` | Known | Screen layout |
| 0x005709E1 | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x005709FC | `NikePlus_Podcasts_Screen` | Known | Screen layout |
| 0x00570A15 | `NikePlus_History_ClearBests_Screen` | Known | Screen layout |
| 0x00570A38 | `Location_Tests_Screen` | Known | Screen layout |
| 0x00570A4E | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x00570A6A | `NikePlus_Playlists_Screen` | Known | Screen layout |
| 0x00570A84 | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x00570AA6 | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x00570AC7 | `MediaLists_MusicVideos_Artists_Screen` | Known | Screen layout |
| 0x00570AED | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x00570B07 | `NikePlus_History_Day_Workouts_Screen` | Known | Screen layout |
| 0x00570B2C | `NikePlus_History_BestWorkouts_Screen` | Known | Screen layout |
| 0x00570B51 | `MediaLists_Genius_Screen` | Known | Screen layout |
| 0x00570B6A | `VoiceMemos_Status_Screen` | Known | Screen layout |
| 0x00570B83 | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x00570B9D | `VoiceMemos_Label_Select_Screen` | Known | Screen layout |
| 0x00570BBC | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x00570BDB | `NikePlus_NowRunning_Idle_Portrait_Screen` | Known | Screen layout |
| 0x00570C04 | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x00570C25 | `PhotosRotateAlt_Screen` | Known | Screen layout |
| 0x00570C3C | `PhotosZoomAlt_Screen` | Known | Screen layout |
| 0x00570C51 | `NikePlus_Calibrate_ResetToDefault_Screen` | Known | Screen layout |
| 0x00570C7A | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x00570C92 | `VoiceMemos_No_Content_Screen` | Known | Screen layout |
| 0x00570CAF | `AddressViewer_Intro_Content_Screen` | Known | Screen layout |
| 0x00570CD2 | `NoContent_Screen` | Known | Screen layout |
| 0x00570CE3 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00570CF9 | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x00570D0F | `NikePlus_EquipmentAlert_Screen` | Known | Screen layout |
| 0x00570D2E | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x00570D44 | `Notes_List_Screen` | Known | Screen layout |
| 0x00570D56 | `Radio_TagList_Screen` | Known | Screen layout |
| 0x00570D6B | `Radio_PresetList_Screen` | Known | Screen layout |
| 0x00570D83 | `Debug_TestList_Screen` | Known | Screen layout |
| 0x00570D99 | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x00570DBA | `NikePlus_PowerPlaylist_Screen` | Known | Screen layout |
| 0x00570DD8 | `MediaLists_GeniusPlaylist_Screen` | Known | Screen layout |
| 0x00570DF9 | `PhotosGL_TvOut_Screen` | Known | Screen layout |
| 0x00570E0F | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x00570E29 | `NikePlus_Dynamic_Workout_Screen` | Known | Screen layout |
| 0x00570E49 | `NikePlus_New_Workout_Screen` | Known | Screen layout |
| 0x00570E65 | `NikePlus_EndPausedWorkout_Screen` | Known | Screen layout |
| 0x00570E86 | `NikePlus_StartWorkout_Screen` | Known | Screen layout |
| 0x00570EA3 | `TODOS_Menu_Screen` | Known | Screen layout |
| 0x00570EB5 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x00570ECB | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x00570EE7 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x00570EFC | `Games_Menu_Screen` | Known | Screen layout |
| 0x00570F0E | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x00570F21 | `VoiceMemos_RecordingList_Menu_Screen` | Known | Screen layout |
| 0x00570F46 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x00570F65 | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x00570F84 | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x00570FA8 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x00570FBE | `Radio_MainMenu_Screen` | Known | Screen layout |
| 0x00570FD4 | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x00570FF2 | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x00571015 | `Radio_TunerTagContextMenu_Screen` | Known | Screen layout |
| 0x00571036 | `Radio_TunerContextMenu_Screen` | Known | Screen layout |
| 0x00571054 | `CoverFlow_Screen` | Known | Screen layout |
| 0x00571065 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x00571079 | `Volume_Overlay_Screen` | Known | Screen layout |
| 0x0057108F | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x005710B1 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x005710C9 | `Stopwatch_SessionSummary_Screen` | Known | Screen layout |
| 0x005710E9 | `NikePlus_End_WorkoutSummary_Screen` | Known | Screen layout |
| 0x0057110C | `NikePlus_EndHercules_WorkoutSummary_Screen` | Known | Screen layout |
| 0x00571137 | `NikePlus_History_WorkoutSummary_Screen` | Known | Screen layout |
| 0x0057115E | `SettingsMenu_AboutWithAccessory_Screen` | Known | Screen layout |
| 0x00571185 | `Location_NMEA_History_Screen` | Known | Screen layout |
| 0x005711A2 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x005711BA | `Radio_TrackHistory_Screen` | Known | Screen layout |
| 0x005711D4 | `Clock_City_Screen` | Known | Screen layout |
| 0x005711E6 | `SettingsMenus_ShakeAdjust_Intensity_Screen` | Known | Screen layout |
| 0x00571290 | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x005712A5 | `MediaLists_iTunesUEpisodes_Screen_Plain` | Known | Screen layout |
| 0x005712CD | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x005712F0 | `Nike_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x00571312 | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0057133A | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x00571358 | `NikePlus_History_WorkoutSummary_Screen_Pedometer_Session` | Known | Screen layout |
| 0x00571391 | `NikePlus_Custom_Screen_Weight_ToPedometerSession` | Known | Screen layout |
| 0x005714F1 | `RentalInfo_Screen_ExpiringSoon` | Known | Screen layout |
| 0x00571510 | `PhotosGL_TvOut_NTSC_Screen_Paused_Video` | Known | Screen layout |
| 0x00571538 | `PhotosGL_TvOut_PAL_Screen_Paused_Video` | Known | Screen layout |
| 0x0057155F | `PhotosGL_TvOut_NTSC_Screen_Volume_Video` | Known | Screen layout |
| 0x00571587 | `PhotosGL_TvOut_PAL_Screen_Volume_Video` | Known | Screen layout |
| 0x005715AE | `PhotosGL_TvOut_NTSC_Screen_Playing_Video` | Known | Screen layout |
| 0x005715D7 | `PhotosGL_TvOut_PAL_Screen_Playing_Video` | Known | Screen layout |
| 0x005715FF | `NowPlaying_Screen_Video` | Known | Screen layout |
| 0x00571617 | `PhotosGL_TvOut_NTSC_Screen_Default_Video` | Known | Screen layout |
| 0x00571640 | `PhotosGL_TvOut_PAL_Screen_Default_Video` | Known | Screen layout |
| 0x00571668 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x005716BE | `CoverFlow_Screen_DefaultNoAlbumInfo` | Known | Screen layout |
| 0x005716E2 | `MainMenus_Music_Screen_Radio` | Known | Screen layout |
| 0x005716FF | `Unsupported_Screen_Radio` | Known | Screen layout |
| 0x00571718 | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x00571734 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x00571750 | `MainMenus_Main_Screen_Filmstrip` | Known | Screen layout |
| 0x00571770 | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x005717FA | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x00571820 | `ChargingMode_ScreenLayout_LowPower` | Known | Screen layout |
| 0x00571843 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x0057185C | `GeniusMixes_Loading_Screen_Error` | Known | Screen layout |
| 0x0057187D | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x00571899 | `NikePlus_ActivityStopped_Screen_Contextual_FoundSensor` | Known | Screen layout |
| 0x005718D0 | `NikePlus_ActivityStopped_Screen_Contextual_NoSensor` | Known | Screen layout |
| 0x00571904 | `PhotosGL_Camera_Screen_Thumbs` | Known | Screen layout |
| 0x00571922 | `PhotosGL_Camera_Alt_Screen_Thumbs` | Known | Screen layout |
| 0x00571944 | `NikePlus_NowRunning_Screen_Calories` | Known | Screen layout |
| 0x00571968 | `NikePlus_Custom_Screen_Calories` | Known | Screen layout |
| 0x00571988 | `MainMenus_Main_Screen_NoMovies` | Known | Screen layout |
| 0x005719A7 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x005719C1 | `NoContent_Screen_MainNoMovies` | Known | Screen layout |
| 0x005719DF | `MainMenus_Main_Screen_Games` | Known | Screen layout |
| 0x005719FB | `MainMenu_Main_Screen_NoGenres` | Known | Screen layout |
| 0x00571A19 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x00571A55 | `MainMenu_Main_Screen_GeniusMixes` | Known | Screen layout |
| 0x00571A76 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x00571A95 | `MainMenu_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x00571AB7 | `MainMenus_Main_Screen_NoRentals` | Known | Screen layout |
| 0x00571AD7 | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00571AF2 | `NoContent_Screen_MainNoRentals` | Known | Screen layout |
| 0x00571B11 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x00571B2E | `MainMenu_Main_Screen_NoAlbums` | Known | Screen layout |
| 0x00571B4C | `MainMenu_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00571B70 | `Radio_TagList_Screen_Instructions` | Known | Screen layout |
| 0x00571B92 | `Radio_PresetList_Screen_Instructions` | Known | Screen layout |
| 0x00571BB7 | `Radio_TrackHistory_Screen_Instructions` | Known | Screen layout |
| 0x00571BDE | `MainMenus_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00571C02 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00571C21 | `NoContent_Screen_MainNoMusicVideos` | Known | Screen layout |
| 0x00571C44 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00571C63 | `MainMenus_Videos_Screen_NoVideos` | Known | Screen layout |
| 0x00571C84 | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x00571C9E | `NoContent_Screen_MainNoVideos` | Known | Screen layout |
| 0x00571CBC | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x00571CDD | `MainMenus_Main_Screen_Photos` | Known | Screen layout |
| 0x00571CFA | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00571D19 | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x00571D33 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x00571D52 | `MainMenu_Main_Screen_NoComposers` | Known | Screen layout |
| 0x00571D99 | `PhotosGL_Camera_Screen_Brightness` | Known | Screen layout |
| 0x00571DBB | `PhotosGL_Camera_Alt_Screen_Brightness` | Known | Screen layout |
| 0x00571DE1 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x00571E04 | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x00571E1F | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x00571E40 | `Camera_Screen_Effects` | Known | Screen layout |
| 0x00571E56 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x00571E73 | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x00571E94 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x00571EB0 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00571ED1 | `MainMenus_Main_Screen_NoVideoPlaylists` | Known | Screen layout |
| 0x00571EF8 | `MainMenu_Main_Screen_NoArtists` | Known | Screen layout |
| 0x00571F17 | `NowPlaying_Screen_Genius` | Known | Screen layout |
| 0x00571F30 | `Genius_Error_Screen_NoGenius` | Known | Screen layout |
| 0x00571F4D | `MainMenu_Main_Screen_NikePlus` | Known | Screen layout |
| 0x00571F6B | `MainMenus_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00571F8B | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x00571FA6 | `NoContent_Screen_MainNoTVShows` | Known | Screen layout |
| 0x00571FC5 | `NikePlus_Custom_Screen_Weight` | Known | Screen layout |
| 0x005720D8 | `Firewire_Charging_Unsupported_Screen_Alt` | Known | Screen layout |
| 0x00572101 | `NowPlaying_Idle_Screen_Alt` | Known | Screen layout |
| 0x0057211C | `Notes_Contents_Screen_Alt` | Known | Screen layout |
| 0x00572136 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x0057214C | `Volume_Overlay_Screen_Alt` | Known | Screen layout |
| 0x00572203 | `VoiceMemos_Screen_InsufficientDiskSpace2_Default` | Known | Screen layout |
| 0x00572234 | `VoiceMemos_Screen_MicrophoneRequired_Default` | Known | Screen layout |
| 0x00572261 | `VoiceMemos_Screen_MicrophoneDisconnected_Default` | Known | Screen layout |
| 0x00572292 | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x005722C5 | `VoiceMemos_Screen_InsufficientDiskSpace_Default` | Known | Screen layout |
| 0x005722F5 | `NowPlaying_Screen_TVOut_ConnectCable_Default` | Known | Screen layout |
| 0x00572322 | `VoiceMemos_Screen_Idle_Default` | Known | Screen layout |
| 0x0057235F | `NikePlus_EndWorkout_Screen_Calibration_Contextual_Landscape_Default` | Known | Screen layout |
| 0x005723A3 | `Nike_Volume_Screen_Landscape_Default` | Known | Screen layout |
| 0x0057241D | `VoiceMemos_Screen_DeleteAsk_Default` | Known | Screen layout |
| 0x00572441 | `VoiceMemos_Screen_DeleteAllAsk_Default` | Known | Screen layout |
| 0x00572468 | `NowPlaying_Screen_TVOutAsk_Default` | Known | Screen layout |
| 0x0057248B | `NikePlus_EndWorkout_Screen_Contextual_Default` | Known | Screen layout |
| 0x005724B9 | `NikePlus_History_Screen_Contextual_Default` | Known | Screen layout |
| 0x0057250F | `MediaLists_Camera_Local_Media_Screen_Default` | Known | Screen layout |
| 0x0057253C | `PhotosGL_Camera_Screen_Default` | Known | Screen layout |
| 0x0057255B | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x00572581 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0057259F | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x005725CB | `AddressViewer_PartialLoad_Screen_Default` | Known | Screen layout |
| 0x005725F4 | `NikePlus_IsLinked_Screen_Default` | Known | Screen layout |
| 0x00572615 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x0057263D | `Alarms_Alarm_Clock_Triggered_Screen_Default` | Known | Screen layout |
| 0x00572669 | `Alarms_Alarm_Triggered_Screen_Default` | Known | Screen layout |
| 0x0057268F | `Firewire_Charging_Unsupported_Screen_Default` | Known | Screen layout |
| 0x005726BC | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x005726E2 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x005726FA | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x00572715 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x00572732 | `Game_Screen_Default` | Known | Screen layout |
| 0x00572746 | `NikePlus_Deleteme_Screen_Default` | Known | Screen layout |
| 0x00572767 | `Settings_DateTime_Time_Screen_Default` | Known | Screen layout |
| 0x0057278D | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x005727AE | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x005727D7 | `Nike_Volume_Screen_Default` | Known | Screen layout |
| 0x005727F2 | `Settings_DateTime_TimeZone_Screen_Default` | Known | Screen layout |
| 0x0057281C | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x00572849 | `NikePlus_Daily_landscape_Screen_Default` | Known | Screen layout |
| 0x00572871 | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x0057289A | `PhotosRotate_Screen_Default` | Known | Screen layout |
| 0x005728B6 | `ContextualMenu_ThreeItem_White_Screen_Default` | Known | Screen layout |
| 0x005728E4 | `ContextualMenu_FiveItem_White_Screen_Default` | Known | Screen layout |
| 0x00572911 | `ContextualMenu_TwoItem_White_Screen_Default` | Known | Screen layout |
| 0x0057293D | `ContextualMenu_FourItem_White_Screen_Default` | Known | Screen layout |
| 0x0057296A | `Calendar_Loading_Screen_Default` | Known | Screen layout |
| 0x0057298A | `AddressViewer_Loading_Screen_Default` | Known | Screen layout |
| 0x005729AF | `Notes_Loading_Screen_Default` | Known | Screen layout |
| 0x005729CC | `GeniusMixes_Loading_Screen_Default` | Known | Screen layout |
| 0x005729EF | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x00572A0D | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x00572A35 | `NikePlus_Remote_Unlinking_Screen_Default` | Known | Screen layout |
| 0x00572A5E | `NikePlus_HeartMonitor_Unlinking_Screen_Default` | Known | Screen layout |
| 0x00572A8D | `NikePlus_History_WorkoutGraph_Screen_Default` | Known | Screen layout |
| 0x00572ABA | `ContextualMenu_ThreeItem_Black_Screen_Default` | Known | Screen layout |
| 0x00572AE8 | `ContextualMenu_FiveItem_Black_Screen_Default` | Known | Screen layout |
| 0x00572B15 | `ContextualMenu_TwoItem_Black_Screen_Default` | Known | Screen layout |
| 0x00572B41 | `ContextualMenu_FourItem_Black_Screen_Default` | Known | Screen layout |
| 0x00572B6E | `Clock_Screen_Default` | Known | Screen layout |
| 0x00572B83 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x00572BA1 | `NikePlus_HeartMonitor_LinkingInitial_Screen_Default` | Known | Screen layout |
| 0x00572BD5 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x00572BFB | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x00572C1F | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x00572C38 | `PhotosZoom_Screen_Default` | Known | Screen layout |
| 0x00572C52 | `Location_Main_Screen_Default` | Known | Screen layout |
| 0x00572C6F | `Pedometer_Main_Screen_Default` | Known | Screen layout |
| 0x00572C8D | `AddressViewer_Main_Screen_Default` | Known | Screen layout |
| 0x00572CAF | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x00572CCC | `MainMenus_Main_Screen_Default` | Known | Screen layout |
| 0x00572CEA | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x00572D07 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x00572D23 | `RentalDeleted_Notification_Screen_Default` | Known | Screen layout |
| 0x00572D4D | `SingleRentalExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x00572D7E | `MultipleRentalsExpiring_Notification_Screen_Default` | Known | Screen layout |
| 0x00572DB2 | `RentalError_Notification_Screen_Default` | Known | Screen layout |
| 0x00572DDA | `DeleteRental_Confirmation_Screen_Default` | Known | Screen layout |
| 0x00572E03 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x00572E2F | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x00572E58 | `RentalInfo_Screen_Default` | Known | Screen layout |
| 0x00572E72 | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00572E93 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x00572EAF | `NikePlus_Calendar_Screen_Default` | Known | Screen layout |
| 0x00572ED0 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x00572EF2 | `NikePlus_CalibrationCompleteError_Screen_Default` | Known | Screen layout |
| 0x00572F23 | `Extras_Screen_Default` | Known | Screen layout |
| 0x00572F39 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x00572F5F | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00572F80 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00572FA6 | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x00572FC4 | `NikePlus_Settings_Screen_Default` | Known | Screen layout |
| 0x00572FE5 | `NikePlus_Songs_Screen_Default` | Known | Screen layout |
| 0x00573003 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00573025 | `NikePlus_History_Totals_Screen_Default` | Known | Screen layout |
| 0x0057304C | `NikePlus_History_ClearTotals_Screen_Default` | Known | Screen layout |
| 0x00573078 | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x005730A4 | `ToDo_Instructions_Screen_Default` | Known | Screen layout |
| 0x005730C5 | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x005730E9 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x0057310B | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x0057312F | `MediaLists_Camera_All_Videos_Screen_Default` | Known | Screen layout |
| 0x0057315B | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x0057317A | `No_Photos_Screen_Default` | Known | Screen layout |
| 0x00573193 | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x005731B5 | `AddressViewer_ContactGroups_Screen_Default` | Known | Screen layout |
| 0x005731E0 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x00573204 | `NikePlus_CalibrationCompleteSuccess_Screen_Default` | Known | Screen layout |
| 0x00573237 | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x00573255 | `NikePlus_History_ClearBests_Screen_Default` | Known | Screen layout |
| 0x00573280 | `Location_Tests_Screen_Default` | Known | Screen layout |
| 0x0057329E | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x005732C2 | `NikePlus_Playlists_Screen_Default` | Known | Screen layout |
| 0x005732E4 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x0057330E | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00573337 | `MediaLists_MusicVideos_Artists_Screen_Default` | Known | Screen layout |
| 0x00573365 | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00573387 | `NikePlus_History_Day_Workouts_Screen_Default` | Known | Screen layout |
| 0x005733B4 | `NikePlus_History_BestWorkouts_Screen_Default` | Known | Screen layout |
| 0x005733E1 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00573402 | `VoiceMemos_Label_Select_Screen_Default` | Known | Screen layout |
| 0x00573429 | `PhotosGL_Camera_Alt_Screen_Default` | Known | Screen layout |
| 0x0057344C | `PhotosRotateAlt_Screen_Default` | Known | Screen layout |
| 0x0057346B | `PhotosZoomAlt_Screen_Default` | Known | Screen layout |
| 0x00573488 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x005734A8 | `Pedometer_Ambient_Screen_Default` | Known | Screen layout |
| 0x005734C9 | `AddressViewer_Intro_Content_Screen_Default` | Known | Screen layout |
| 0x005734F4 | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x00573512 | `Nike_Dummy_BarTest_Screen_Default` | Known | Screen layout |
| 0x00573534 | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x00573552 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x0057356C | `Radio_TagList_Screen_Default` | Known | Screen layout |
| 0x00573589 | `Radio_PresetList_Screen_Default` | Known | Screen layout |
| 0x005735A9 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x005735C7 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x005735F0 | `NikePlus_PowerPlaylist_Screen_Default` | Known | Screen layout |
| 0x00573616 | `MediaLists_GeniusPlaylist_Screen_Default` | Known | Screen layout |
| 0x0057363F | `PhotosGL_TvOut_Screen_Default` | Known | Screen layout |
| 0x0057365D | `NikePlus_New_Workout_Screen_Default` | Known | Screen layout |
| 0x00573681 | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x005736A6 | `TODOS_Menu_Screen_Default` | Known | Screen layout |
| 0x005736C0 | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x005736DE | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x005736FB | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x00573715 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x00573730 | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x0057374F | `VoiceMemos_RecordingList_Menu_Screen_Default` | Known | Screen layout |
| 0x0057377C | `Debug_MainMenu_Screen_Default` | Known | Screen layout |
| 0x0057379A | `Radio_TunerTagContextMenu_Screen_Default` | Known | Screen layout |
| 0x005737C3 | `Radio_TunerContextMenu_Screen_Default` | Known | Screen layout |
| 0x005737E9 | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x00573802 | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x0057381E | `Volume_Overlay_Screen_Default` | Known | Screen layout |
| 0x0057383C | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x00573866 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x00573886 | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x005738AE | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x005738D9 | `NikePlus_EndHercules_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x0057390C | `NikePlus_History_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x0057393B | `Location_NMEA_History_Screen_Default` | Known | Screen layout |
| 0x00573960 | `NikePlus_History_Screen_Default` | Known | Screen layout |
| 0x00573980 | `Radio_TrackHistory_Screen_Default` | Known | Screen layout |
| 0x005739A2 | `Clock_City_Screen_Default` | Known | Screen layout |
| 0x00573BE0 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x00573C00 | `NikePlus_ActivityStopped_Screen_Contextual_FoundSensor_Default` | Known | Screen layout |
| 0x00573C3F | `NikePlus_ActivityStopped_Screen_Contextual_NoSensor_Default` | Known | Screen layout |
| 0x00573C7B | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x00573CAB | `Firewire_Charging_Unsupported_Screen_Alt_Default` | Known | Screen layout |
| 0x00573CDC | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x00573CFE | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x00573D1C | `Volume_Overlay_Screen_Alt_Default` | Known | Screen layout |
| 0x00573DA7 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x00573E03 | `MediaLists_OTG_ClearConfirm_ScreenLayout_Default` | Known | Screen layout |
| 0x00573E34 | `RemoteUI_Hercules_ScreenLayout_Default` | Known | Screen layout |
| 0x00573E5B | `Radio_MainMenu_ScreenLayout_Default` | Known | Screen layout |
| 0x00573E7F | `Pedometer_Screen_Ambient` | Known | Screen layout |
| 0x00573E98 | `NikePlus_History_WorkoutSummary_Screen_Pedometer_Ambient` | Known | Screen layout |
| 0x00573ED1 | `MediaLists_iTunesU_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00573EFC | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00573F28 | `NikePlus_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00573F52 | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00573F7D | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x00573FA5 | `MainMenus_Main_Screen_NoiTunesUArt` | Known | Screen layout |
| 0x00573FC8 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00573FE9 | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x0057400A | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00574030 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x00574052 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00574076 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x0057409A | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x005740D1 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x005740F8 | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0057411B | `NowPlaying_Screen_Video_TVOut` | Known | Screen layout |
| 0x00574198 | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x005741C5 | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x005741F5 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x0057426A | `SettingsMenu_About_Screen_Basic_Layout` | Known | Screen layout |
| 0x00574291 | `SettingsMenu_AboutWithAccessory_Screen_Basic_Layout` | Known | Screen layout |
| 0x005744C8 | `NikePlus_EndPausedWorkout_Screen_BasicSave_Layout` | Known | Screen layout |
| 0x005744FA | `NikePlus_EndPausedWorkout_Screen_DistanceSave_Layout` | Known | Screen layout |
| 0x0057452F | `NikePlus_EndPausedWorkout_Screen_TimeSave_Layout` | Known | Screen layout |
| 0x00574560 | `NikePlus_EndPausedWorkout_Screen_CaloriesSave_Layout` | Known | Screen layout |
| 0x00574595 | `NikePlus_EndPausedWorkout_Screen_QuickstartSave_Layout` | Known | Screen layout |
| 0x005745CC | `MainMenu_Main_Screen_Pedometer_InActive_Layout` | Known | Screen layout |
| 0x005746DB | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x005746F6 | `Pedometer_Main_Landscape_Screen_Medium_Layout` | Known | Screen layout |
| 0x00574724 | `Pedometer_Trainer_Paused_Screen_Layout` | Known | Screen layout |
| 0x0057474B | `Pedometer_Main_Landscape_Screen_Layout` | Known | Screen layout |
| 0x00574772 | `Pedometer_Ambient_Landscape_Screen_Layout` | Known | Screen layout |
| 0x0057479C | `Hercules_Complete_Screen_Layout` | Known | Screen layout |
| 0x005747BC | `Pedometer_Ambient_Landscape_Medium_Screen_Layout` | Known | Screen layout |
| 0x005747ED | `Hercules_Connection_Screen_Layout` | Known | Screen layout |
| 0x00574874 | `MainMenu_Main_Screen_Pedometer_Layout` | Known | Screen layout |
| 0x0057489A | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x00574A0B | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00574B99 | `Lock_Screen_Default_Layout` | Known | Screen layout |
| 0x00574BB4 | `Pedometer_Step_Goal_Screen_Default_Layout` | Known | Screen layout |
| 0x00574C2B | `SettingsMenu_About_Screen_Count_Layout` | Known | Screen layout |
| 0x00574C52 | `SettingsMenu_AboutWithAccessory_Screen_Count_Layout` | Known | Screen layout |
| 0x00574D2B | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00574D42 | `SettingsMenu_AboutWithAccessory_Screen_Accessory_Layout` | Known | Screen layout |
| 0x00574D7A | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00574DA4 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x00574E1F | `Unsupported_Upgrade_ScreenLayout` | Known | Screen layout |
| 0x00574E40 | `Pedometer_Main_Screen_Medium_ScreenLayout` | Known | Screen layout |
| 0x00574E6A | `Pedometer_Ambient_Screen_Medium_ScreenLayout` | Known | Screen layout |
| 0x005751AF | `AddressViewer_Intro_Content_Screen_MainMenu` | Known | Screen layout |
| 0x00575229 | `CoverFlow_Screen_QuickNav` | Known | Screen layout |
| 0x00575243 | `MainMenus_Main_Screen_NoPreview` | Known | Screen layout |
| 0x0057527E | `PhotosGL_Camera_Screen_BatteryLow` | Known | Screen layout |
| 0x005752A0 | `PhotosGL_Camera_Alt_Screen_BatteryLow` | Known | Screen layout |
| 0x005752C6 | `PhotosGL_TvOut_Screen_BatteryLow` | Known | Screen layout |
| 0x005752E7 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x0057530A | `NowPlaying_Screen_Initial_From_CoverFlow` | Known | Screen layout |
| 0x00575333 | `NowPlaying_Screen_Exit_To_CoverFlow` | Known | Screen layout |
| 0x00575357 | `RentalInfo_Screen_ExpiresToday` | Known | Screen layout |
| 0x006347D8 | `GeniusMixes_Loading_Screen_Error` | Known | Screen layout |
| 0x00634810 | `Genius_Loading_Screen_Error` | Known | Screen layout |
| 0x006348AC | `MainMenus_Nike_Screen` | Known | Screen layout |
| 0x006348C4 | `MainMenu_Nike_Screen_Default` | Known | Screen layout |
| 0x0063495C | `AddressViewer_Loading_Screen` | Known | Screen layout |
| 0x0063497C | `AddressViewer_Loading_Screen_Default` | Known | Screen layout |
| 0x006349C8 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x006349E0 | `Calendar_Menu_Screen_Default` | Known | Screen layout |
| 0x00634A2C | `Calendar_Loading_Screen` | Known | Screen layout |
| 0x00634A44 | `Calendar_Loading_Screen_Default` | Known | Screen layout |
| 0x00634A80 | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x00634A94 | `Alarms_Menu_Screen_Default` | Known | Screen layout |
| 0x00634B08 | `Clock_Screen` | Known | Screen layout |
| 0x00634B18 | `Clock_Screen_Default` | Known | Screen layout |
| 0x00634BA0 | `Games_Menu_Screen` | Known | Screen layout |
| 0x00634BB4 | `Games_Menu_Screen_Default` | Known | Screen layout |
| 0x00634D38 | `Debug_MainMenu_Screen` | Known | Screen layout |
| 0x00634D50 | `Debug_MainMenu_Screen_Default` | Known | Screen layout |
| 0x00634DE8 | `Extras_Screen` | Known | Screen layout |
| 0x00634DF8 | `AddressViewer_ContactDetails_Screen` | Known | Screen layout |
| 0x00634E1C | `AddressViewer_ContactDetails_Screen_Default` | Known | Screen layout |
| 0x00634EC8 | `Alarms_Menu_Screen_Analog` | Known | Screen layout |
| 0x00634F00 | `Alarms_Timer_Props_Screen` | Known | Screen layout |
| 0x00634F1C | `Alarms_Timer_Props_Screen_Default` | Known | Screen layout |
| 0x00634F64 | `Calendar_Event_Screen` | Known | Screen layout |
| 0x00634F7C | `Calendar_Event_Screen_Default` | Known | Screen layout |
| 0x00634FD8 | `Alarms_Sleep_Timer_Screen` | Known | Screen layout |
| 0x00634FF4 | `Alarms_Sleep_Timer_Screen_Default` | Known | Screen layout |
| 0x0063504C | `Alarms_Set_Alarm_Date_Screen` | Known | Screen layout |
| 0x006350AC | `Alarms_Set_Alarm_Frequency_Screen` | Known | Screen layout |
| 0x006350D0 | `Alarms_Set_Alarm_Frequency_Screen_Default` | Known | Screen layout |
| 0x00635118 | `Alarms_Set_Alarm_Label_Screen` | Known | Screen layout |
| 0x00635138 | `Alarms_Set_Alarm_Label_Screen_Default` | Known | Screen layout |
| 0x0063517C | `Alarms_Set_Alarm_Sound_Screen` | Known | Screen layout |
| 0x0063519C | `Alarms_Set_Alarm_Sound_Screen_Default` | Known | Screen layout |
| 0x006351E0 | `Alarms_Set_Alarm_Time_Screen` | Known | Screen layout |
| 0x0063538C | `Alarms_Set_Alarm_Playlist_Screen` | Known | Screen layout |
| 0x006353B0 | `Alarms_Set_Alarm_Playlist_Screen_Default` | Known | Screen layout |
| 0x006353F8 | `Alarms_Set_Alarm_Tones_Screen` | Known | Screen layout |
| 0x00635418 | `Alarms_Set_Alarm_Tones_Screen_Default` | Known | Screen layout |
| 0x00635508 | `Calendar_Monthly_Screen` | Known | Screen layout |
| 0x00635520 | `Calendar_Monthly_Screen_Default` | Known | Screen layout |
| 0x006355E8 | `Calendar_Day_Screen` | Known | Screen layout |
| 0x006355FC | `Calendar_Day_Screen_Default` | Known | Screen layout |
| 0x00635640 | `Calendar_Day_Screen_NoEvents` | Known | Screen layout |
| 0x006356A0 | `ToDo_Item_Screen` | Known | Screen layout |
| 0x006356B4 | `ToDo_Item_Screen_Default` | Known | Screen layout |
| 0x00635828 | `Camera_Screen_Active` | Known | Screen layout |
| 0x00635858 | `Camera_Screen_DiskFull` | Known | Screen layout |
| 0x00635890 | `Camera_Screen_Effects` | Known | Screen layout |
| 0x006358C8 | `Camera_Screen_Uninitialized` | Known | Screen layout |
| 0x00635A70 | `PhotosGL_Camera_Screen` | Known | Screen layout |
| 0x00635A88 | `PhotosGL_Camera_Screen_Thumbs` | Known | Screen layout |
| 0x00635F44 | `Clock_City_Screen` | Known | Screen layout |
| 0x00635F58 | `Clock_City_Screen_Default` | Known | Screen layout |
| 0x00636290 | `Clock_Region_Screen` | Known | Screen layout |
| 0x006362A4 | `Clock_Region_Screen_Default` | Known | Screen layout |
| 0x006364A4 | `CoverFlow_Screen_QuickNav` | Known | Screen layout |
| 0x0063653C | `CoverFlow_Screen_EnteringNowPlaying` | Known | Screen layout |
| 0x006365E4 | `CoverFlow_Screen_Default` | Known | Screen layout |
| 0x00636638 | `Settings_DateTime_SetDate_Screen` | Known | Screen layout |
| 0x0063665C | `Settings_DateTime_SetDate_Screen_Default` | Known | Screen layout |
| 0x006366A4 | `Settings_DateTime_SetTime_Screen` | Known | Screen layout |
| 0x006366C8 | `Settings_DateTime_SetTime_Screen_Default` | Known | Screen layout |
| 0x00636710 | `Settings_DateTime_SetTimeZone_Screen` | Known | Screen layout |
| 0x00636738 | `Settings_DateTime_SetTimeZone_Screen_Default` | Known | Screen layout |
| 0x00636840 | `MainMenus_Main_Screen` | Known | Screen layout |
| 0x00636870 | `Game_Screen` | Known | Screen layout |
| 0x0063687C | `Game_Screen_Default` | Known | Screen layout |
| 0x00636950 | `Game_Controls_Screen` | Known | Screen layout |
| 0x006369D0 | `Game_Running_Screen` | Known | Screen layout |
| 0x00636A00 | `Game_Memory_Error_Screen` | Known | Screen layout |
| 0x00636A38 | `Game_Signing_Error_Screen` | Known | Screen layout |
| 0x00636A70 | `Game_Unknown_Error_Screen` | Known | Screen layout |
| 0x00636AA8 | `Game_Version_Error_Screen` | Known | Screen layout |
| 0x00636B6C | `Location_Tests_Screen` | Known | Screen layout |
| 0x00636B84 | `Location_Tests_Screen_Default` | Known | Screen layout |
| 0x00636BB8 | `Location_NMEA_History_Screen` | Known | Screen layout |
| 0x00636BD8 | `Location_NMEA_History_Screen_Default` | Known | Screen layout |
| 0x00636C10 | `Location_Main_Screen` | Known | Screen layout |
| 0x00636CE8 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00636E10 | `LockediPod_Screen` | Known | Screen layout |
| 0x00636E60 | `SettingsMenus_Main_Screen` | Known | Screen layout |
| 0x00636EA4 | `Lock_Screen_Confirm_Layout` | Known | Screen layout |
| 0x00636EE8 | `Lock_Screen_Enter_Layout` | Known | Screen layout |
| 0x00636F04 | `Lock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x00636F48 | `Lock_Screen_New_Layout` | Known | Screen layout |
| 0x00636F70 | `SettingsMenus_Playback_Screen` | Known | Screen layout |
| 0x00636FE0 | `SettingsMenus_VolumeLimit_Screen` | Known | Screen layout |
| 0x00637028 | `VolumeLimitLock_Screen_Incorrect_Layout` | Known | Screen layout |
| 0x006374F8 | `Search_Main_Screen` | Known | Screen layout |
| 0x0063750C | `Search_Main_Screen_NoKeyboard` | Known | Screen layout |
| 0x00637548 | `MainMenu_Main_Screen_NoContentSearch` | Known | Screen layout |
| 0x0063796C | `MainMenus_Main_Screen_Radio` | Known | Screen layout |
| 0x00637A08 | `Extras_Screen_Default` | Known | Screen layout |
| 0x00637A50 | `MainMenus_Main_Screen_Filmstrip` | Known | Screen layout |
| 0x00637AA8 | `MainMenus_Main_Screen_WorldClock` | Known | Screen layout |
| 0x00637B04 | `MainMenus_Main_Screen_Alarms` | Known | Screen layout |
| 0x00637B5C | `MainMenus_Main_Screen_Calendar` | Known | Screen layout |
| 0x00637BD0 | `MainMenus_Main_Screen_AddressBook` | Known | Screen layout |
| 0x00637C64 | `MainMenus_Main_Screen_Notes` | Known | Screen layout |
| 0x00637CB8 | `MainMenus_Main_Screen_Lock` | Known | Screen layout |
| 0x00637D0C | `MainMenus_Main_Screen_Stopwatch` | Known | Screen layout |
| 0x00637DB0 | `Location_Main_Screen_Default` | Known | Screen layout |
| 0x00637DEC | `MainMenus_Main_Screen_NoPreview` | Known | Screen layout |
| 0x00637E28 | `VoiceMemos_Main_Menu_Screen` | Known | Screen layout |
| 0x00637E44 | `VoiceMemos_Menu_Screen_Default` | Known | Screen layout |
| 0x00637E98 | `MainMenus_Main_Screen_VoiceMemos` | Known | Screen layout |
| 0x00637ED8 | `Speakers_Main_Screen` | Known | Screen layout |
| 0x00637EF0 | `Speakers_Main_Screen_Default` | Known | Screen layout |
| 0x00637F44 | `MainMenus_Main_Screen_Speakers` | Known | Screen layout |
| 0x00637FD4 | `MainMenus_Main_Screen_Settings` | Known | Screen layout |
| 0x00638040 | `MainMenus_Main_Screen_Sleep` | Known | Screen layout |
| 0x00638118 | `MainMenus_Main_Screen_Shuffle` | Known | Screen layout |
| 0x006381D0 | `MediaLists_Albums_Screen` | Known | Screen layout |
| 0x006381EC | `MediaLists_Albums_Screen_WithArtistName` | Known | Screen layout |
| 0x00638230 | `MediaLists_Artists_Screen` | Known | Screen layout |
| 0x0063824C | `MediaLists_Artists_Screen_Default` | Known | Screen layout |
| 0x00638290 | `MediaLists_Audiobooks_Screen` | Known | Screen layout |
| 0x006382B0 | `MediaLists_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x00638314 | `MediaLists_Camera_Local_Media_Screen` | Known | Screen layout |
| 0x0063833C | `MediaLists_Camera_Local_Media_Screen_Default` | Known | Screen layout |
| 0x006383AC | `MediaLists_Composers_Screen` | Known | Screen layout |
| 0x006383C8 | `MediaLists_Composers_Screen_Default` | Known | Screen layout |
| 0x0063840C | `CoverFlow_Screen` | Known | Screen layout |
| 0x00638440 | `MainMenu_Main_Screen_NoAlbums` | Known | Screen layout |
| 0x00638480 | `MainMenu_Main_Screen_NoArtists` | Known | Screen layout |
| 0x006384C4 | `MainMenu_Main_Screen_NoAudiobooks` | Known | Screen layout |
| 0x00638508 | `MainMenus_Main_Screen_NoCamera` | Known | Screen layout |
| 0x0063854C | `MainMenu_Main_Screen_NoCompilations` | Known | Screen layout |
| 0x00638594 | `MainMenu_Main_Screen_NoComposers` | Known | Screen layout |
| 0x006385DC | `NoContent_Screen_Music` | Known | Screen layout |
| 0x006385F4 | `NoContent_Screen_MainNoMusic` | Known | Screen layout |
| 0x00638638 | `MainMenu_Main_Screen_GeniusMixes` | Known | Screen layout |
| 0x0063867C | `MainMenu_Main_Screen_NoGenres` | Known | Screen layout |
| 0x006386BC | `MainMenus_Main_Screen_NoMovies` | Known | Screen layout |
| 0x006386FC | `NoContent_Screen` | Known | Screen layout |
| 0x00638710 | `NoContent_Screen_NoMovies` | Known | Screen layout |
| 0x0063874C | `MainMenus_Main_Screen_NoMusic` | Known | Screen layout |
| 0x0063878C | `NoContent_Screen_NoMusic` | Known | Screen layout |
| 0x006387CC | `MainMenus_Main_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00638814 | `NoContent_Screen_NoMusicVideos` | Known | Screen layout |
| 0x00638854 | `MainMenus_Main_Screen_NoPhotos` | Known | Screen layout |
| 0x00638894 | `NoContent_Screen_NoPhotos` | Known | Screen layout |
| 0x006388D0 | `MainMenu_Main_Screen_NoPlaylists` | Known | Screen layout |
| 0x00638938 | `MainMenus_Main_Screen_NoPodcasts` | Known | Screen layout |
| 0x00638980 | `NoContent_Screen_NoPodcasts` | Known | Screen layout |
| 0x006389BC | `MainMenus_Main_Screen_NoRentals` | Known | Screen layout |
| 0x006389FC | `NoContent_Screen_NoRentals` | Known | Screen layout |
| 0x00638A38 | `MainMenus_Main_Screen_NoTVShows` | Known | Screen layout |
| 0x00638A78 | `NoContent_Screen_NoTVShows` | Known | Screen layout |
| 0x00638ABC | `MainMenus_Main_Screen_NoVideoPlaylists` | Known | Screen layout |
| 0x00638B04 | `MainMenus_Main_Screen_NoVideos` | Known | Screen layout |
| 0x00638B44 | `NoContent_Screen_NoVideos` | Known | Screen layout |
| 0x00638B80 | `MediaLists_GeniusMixes_Screen` | Known | Screen layout |
| 0x00638BA0 | `MediaLists_GeniusMixes_Screen_Default` | Known | Screen layout |
| 0x00638BE4 | `MediaLists_Genres_Screen` | Known | Screen layout |
| 0x00638C00 | `MediaLists_Genres_Screen_Default` | Known | Screen layout |
| 0x00638C4C | `NikePlus_Custom_Screen` | Known | Screen layout |
| 0x00638C64 | `NikePlus_Custom_Screen_Weight_ToPedometerSession` | Known | Screen layout |
| 0x00638CC0 | `Pedometer_Ambient_Landscape_Screen` | Known | Screen layout |
| 0x00638CE4 | `Pedometer_Ambient_Landscape_Screen_Layout` | Known | Screen layout |
| 0x00638D38 | `Pedometer_Screen_Ambient` | Known | Screen layout |
| 0x00638D54 | `Pedometer_Ambient_Screen_Default` | Known | Screen layout |
| 0x00638DA4 | `Pedometer_Main_Landscape_Screen` | Known | Screen layout |
| 0x00638DC4 | `Pedometer_Main_Landscape_Screen_Layout` | Known | Screen layout |
| 0x00638E14 | `Pedometer_Main_Screen` | Known | Screen layout |
| 0x00638E2C | `Pedometer_Main_Screen_Default` | Known | Screen layout |
| 0x00638E68 | `SettingsMenus_Language_Screen` | Known | Screen layout |
| 0x00638EC8 | `Genius_Intro_Screen` | Known | Screen layout |
| 0x00638EDC | `GeniusMixes_Intro_Screen_Default` | Known | Screen layout |
| 0x00638F60 | `AddressViewer_Intro_Content_Screen` | Known | Screen layout |
| 0x00638F84 | `AddressViewer_Intro_Content_Screen_MainMenu` | Known | Screen layout |
| 0x00638FD0 | `MainMenu_Main_Screen_Default` | Known | Screen layout |
| 0x0063900C | `MediaLists_Movies_Screen` | Known | Screen layout |
| 0x00639028 | `MediaLists_Movies_Screen_WithTime` | Known | Screen layout |
| 0x0063908C | `MediaLists_Playlists_Screen` | Known | Screen layout |
| 0x006390A8 | `MediaLists_Playlists_Screen_Default` | Known | Screen layout |
| 0x006390E8 | `MainMenus_Music_Screen` | Known | Screen layout |
| 0x00639100 | `MainMenu_Music_Screen_Default` | Known | Screen layout |
| 0x0063915C | `MainMenu_Main_Screen_NikePlus` | Known | Screen layout |
| 0x0063919C | `MainMenus_Main_Screen_NoMusicArt` | Known | Screen layout |
| 0x006391E0 | `MainMenu_Music_Screen_NoMusicArt` | Known | Screen layout |
| 0x00639228 | `MainMenus_Main_Screen_NoNowPlayingArt` | Known | Screen layout |
| 0x00639270 | `MainMenu_Main_Screen_NoPlaylistsArt` | Known | Screen layout |
| 0x006392B4 | `MainMenus_Main_Screen_NoPodcastsArt` | Known | Screen layout |
| 0x006392F8 | `MainMenus_Main_Screen_NoVideosArt` | Known | Screen layout |
| 0x0063933C | `MainMenus_Videos_Screen` | Known | Screen layout |
| 0x00639354 | `MainMenus_Videos_Screen_NoVideosArt` | Known | Screen layout |
| 0x00639398 | `MainMenus_Main_Screen_NoiTunesUArt` | Known | Screen layout |
| 0x006393E0 | `MainMenus_Main_Screen_NowPlaying` | Known | Screen layout |
| 0x00639420 | `MainMenu_Main_Screen_Pedometer_Layout` | Known | Screen layout |
| 0x0063946C | `MainMenu_Main_Screen_Pedometer_InActive_Layout` | Known | Screen layout |
| 0x006394D4 | `Photos_Screen` | Known | Screen layout |
| 0x00639528 | `MediaLists_Podcasts_Screen` | Known | Screen layout |
| 0x00639544 | `MediaLists_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006395B8 | `MediaLists_Rentals_Screen` | Known | Screen layout |
| 0x006395D4 | `MediaLists_Rentals_Screen_Default` | Known | Screen layout |
| 0x00639614 | `MediaLists_Songs_Screen` | Known | Screen layout |
| 0x0063962C | `MediaLists_Songs_Screen_WithArtist` | Known | Screen layout |
| 0x0063966C | `MediaLists_TVShows_Screen` | Known | Screen layout |
| 0x00639688 | `MediaLists_TVShows_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x006396D8 | `MediaLists_VideoPlaylists_Screen` | Known | Screen layout |
| 0x006396FC | `MediaLists_VideoPlaylists_Screen_Default` | Known | Screen layout |
| 0x00639760 | `MainMenu_Videos_Screen_Default` | Known | Screen layout |
| 0x006397B8 | `MediaLists_iTunesU_Screen` | Known | Screen layout |
| 0x006397D4 | `MediaLists_iTunesU_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00639860 | `MainMenu_Music_Screen_NoMusic_Playlist` | Known | Screen layout |
| 0x00639A4C | `MediaLists_AudiobookChapters_Screen` | Known | Screen layout |
| 0x00639A70 | `MediaLists_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x00639C14 | `MediaLists_Genius_Screen` | Known | Screen layout |
| 0x00639C30 | `MediaLists_Genius_Screen_Default` | Known | Screen layout |
| 0x00639C78 | `Genius_Loading_Screen` | Known | Screen layout |
| 0x00639C90 | `Genius_Loading_Screen_Default` | Known | Screen layout |
| 0x00639CCC | `MediaLists_Songs_Screen_Plain` | Known | Screen layout |
| 0x00639D10 | `MediaLists_Songs_Screen_WithAlbum` | Known | Screen layout |
| 0x00639D5C | `MediaLists_Songs_Screen_WithAlbumAndArtwork` | Known | Screen layout |
| 0x00639D88 | `MediaLists_Songs_Screen_WithArtistAndArtwork` | Known | Screen layout |
| 0x00639E04 | `MediaLists_Albums_Screen_WithTrackCount` | Known | Screen layout |
| 0x0063A0A8 | `MediaLists_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x0063A0CC | `MediaLists_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0063A0F4 | `MediaLists_iTunesUEpisodes_Screen` | Known | Screen layout |
| 0x0063A118 | `MediaLists_iTunesUEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0063A164 | `MediaLists_GeniusPlaylist_Screen` | Known | Screen layout |
| 0x0063A188 | `MediaLists_GeniusPlaylist_Screen_Default` | Known | Screen layout |
| 0x0063A1E8 | `Genius_Intro_Screen_Default` | Known | Screen layout |
| 0x0063A220 | `MediaLists_NestedPlaylists_Screen` | Known | Screen layout |
| 0x0063A244 | `MediaLists_NestedPlaylists_Screen_Default` | Known | Screen layout |
| 0x0063A2A0 | `OnTheGo_Instructions_Screen` | Known | Screen layout |
| 0x0063A2BC | `OnTheGo_Instructions_Screen_Default` | Known | Screen layout |
| 0x0063A360 | `MediaLists_OTG_ClearConfirm_Screen` | Known | Screen layout |
| 0x0063A384 | `MediaLists_OTG_ClearConfirm_ScreenLayout_Default` | Known | Screen layout |
| 0x0063A604 | `Video_Settings_Screen` | Known | Screen layout |
| 0x0063A61C | `Video_Settings_Screen_Default` | Known | Screen layout |
| 0x0063A678 | `MediaLists_Camera_All_Videos_Screen` | Known | Screen layout |
| 0x0063A69C | `MediaLists_Camera_All_Videos_Screen_Default` | Known | Screen layout |
| 0x0063A6C8 | `MediaLists_MusicVideos_Artists_Screen` | Known | Screen layout |
| 0x0063A6F0 | `MediaLists_MusicVideos_Artists_Screen_Default` | Known | Screen layout |
| 0x0063A750 | `MediaLists_TVSeasons_Screen` | Known | Screen layout |
| 0x0063A76C | `MediaLists_TVSeasons_Screen_Default` | Known | Screen layout |
| 0x0063A7B8 | `MediaLists_TVEpisodes_Screen` | Known | Screen layout |
| 0x0063A7D8 | `MediaLists_TVEpisodes_Screen_Plain` | Known | Screen layout |
| 0x0063A7FC | `MediaLists_MusicVideos_Songs_Screen` | Known | Screen layout |
| 0x0063A820 | `MediaLists_MusicVideos_Songs_Screen_WithAlbumAndArtwork` | Known | Screen layout |
| 0x0063A858 | `MediaLists_MixedVideoTracks_Screen` | Known | Screen layout |
| 0x0063ABB8 | `NowPlaying_Screen_Video_Brightness` | Known | Screen layout |
| 0x0063AC14 | `NowPlaying_Screen_Progress` | Known | Screen layout |
| 0x0063AC54 | `NowPlaying_Screen_Video_Default` | Known | Screen layout |
| 0x0063AC94 | `NowPlaying_Screen_ExtraInfo` | Known | Screen layout |
| 0x0063ACDC | `NowPlaying_Screen_ExtraInfoLoadFailed` | Known | Screen layout |
| 0x0063AD2C | `NowPlaying_Screen_ExtraInfoLoading` | Known | Screen layout |
| 0x0063AD6C | `NowPlaying_Screen_Genius` | Known | Screen layout |
| 0x0063ADCC | `NowPlaying_Screen_Video_Progress` | Known | Screen layout |
| 0x0063AE0C | `NowPlaying_Screen_Rating` | Known | Screen layout |
| 0x0063AE4C | `NowPlaying_Screen_Video_Rating` | Known | Screen layout |
| 0x0063AE90 | `NowPlaying_Screen_Video_RentalWarning` | Known | Screen layout |
| 0x0063AED4 | `NowPlaying_Screen_Scrub` | Known | Screen layout |
| 0x0063AF0C | `NowPlaying_Screen_Video_Scrub` | Known | Screen layout |
| 0x0063AF4C | `NowPlaying_Screen_Shuffle` | Known | Screen layout |
| 0x0063AF8C | `NowPlaying_Screen_Video_StatusBar` | Known | Screen layout |
| 0x0063AFCC | `NowPlaying_Screen_Volume` | Known | Screen layout |
| 0x0063B00C | `NowPlaying_Screen_Video_Volume` | Known | Screen layout |
| 0x0063B334 | `NowPlaying_Screen_Video_TVOut_Default` | Known | Screen layout |
| 0x0063B384 | `NowPlaying_Screen_Video_TVOut_Subtitles_Default` | Known | Screen layout |
| 0x0063B3D4 | `NowPlaying_Screen_Video_TVOut_Scrub` | Known | Screen layout |
| 0x0063B3F8 | `NowPlaying_Screen_TVOut_RentalWarning` | Known | Screen layout |
| 0x0063B420 | `NowPlaying_Screen_Video_TVOut_Volume` | Known | Screen layout |
| 0x0063B4B0 | `Pedometer_Main_Screen_Medium_ScreenLayout` | Known | Screen layout |
| 0x0063B4DC | `Pedometer_Main_Landscape_Screen_Medium_Layout` | Known | Screen layout |
| 0x0063B50C | `Pedometer_Ambient_Screen_Medium_ScreenLayout` | Known | Screen layout |
| 0x0063B53C | `Pedometer_Ambient_Landscape_Medium_Screen_Layout` | Known | Screen layout |
| 0x0063B5D8 | `Notes_Contents_Screen_Alt` | Known | Screen layout |
| 0x0063B5F4 | `Notes_Contents_Screen_Alt_Default` | Known | Screen layout |
| 0x0063B638 | `Notes_List_Screen_Alt` | Known | Screen layout |
| 0x0063B650 | `Notes_List_Screen_Alt_Default` | Known | Screen layout |
| 0x0063B694 | `Notes_Contents_Screen` | Known | Screen layout |
| 0x0063B6AC | `Notes_Contents_Screen_Default` | Known | Screen layout |
| 0x0063B6EC | `Notes_Image_Screen` | Known | Screen layout |
| 0x0063B700 | `Notes_Image_Screen_Default` | Known | Screen layout |
| 0x0063B740 | `Notes_Image_Screen_Error` | Known | Screen layout |
| 0x0063B784 | `Notes_Instructions_Screen` | Known | Screen layout |
| 0x0063B7A0 | `Notes_Instructions_Screen_Default` | Known | Screen layout |
| 0x0063B7E4 | `Notes_List_Screen` | Known | Screen layout |
| 0x0063B7F8 | `Notes_List_Screen_Default` | Known | Screen layout |
| 0x0063B880 | `PhotosGL_Camera_Alt_Screen_Thumbs` | Known | Screen layout |
| 0x0063B8C8 | `PhotosGL_Screen` | Known | Screen layout |
| 0x0063B8D8 | `PhotosGL_Screen_Thumbs` | Known | Screen layout |
| 0x0063B918 | `PhotosGL_Alt_Screen_Thumbs` | Known | Screen layout |
| 0x0063B96C | `Photos_Events_Screen` | Known | Screen layout |
| 0x0063B9CC | `Photos_Faces_Screen` | Known | Screen layout |
| 0x0063BA2C | `Photos_Places_Screen` | Known | Screen layout |
| 0x0063BA78 | `PhotosSettings_Screen` | Known | Screen layout |
| 0x0063BB58 | `PhotosSettingsDuration_Screen` | Known | Screen layout |
| 0x0063BBB0 | `PhotosSettingsSlideshowMusic_Screen` | Known | Screen layout |
| 0x0063BBD4 | `PhotosSettingsSlideshowMusic_Screen_Default` | Known | Screen layout |
| 0x0063BC84 | `PhotosSettingsTransitions_Screen` | Known | Screen layout |
| 0x0063BDAC | `PhotosSettingsSlideshowMusic_Screen_Nested_Default` | Known | Screen layout |
| 0x0063BE60 | `PhotosZoomAlt_Screen` | Known | Screen layout |
| 0x0063BE78 | `PhotosZoomAlt_Screen_Default` | Known | Screen layout |
| 0x0063BEB0 | `PhotosGL_Screen_Default` | Known | Screen layout |
| 0x0063BEE8 | `PhotosZoom_Screen` | Known | Screen layout |
| 0x0063BEFC | `PhotosZoom_Screen_Default` | Known | Screen layout |
| 0x0063BF94 | `PhotosGL_Camera_Screen_TvOut_Ask` | Known | Screen layout |
| 0x0063BFD4 | `PhotosGL_Camera_Screen_Brightness` | Known | Screen layout |
| 0x0063C018 | `PhotosGL_Camera_Alt_Screen_Brightness` | Known | Screen layout |
| 0x0063C060 | `PhotosGL_Camera_Screen_TvOut_ConnectCable` | Known | Screen layout |
| 0x0063C0A4 | `PhotosGL_Camera_Screen_Default` | Known | Screen layout |
| 0x0063C0E0 | `PhotosGL_Camera_Alt_Screen_Default` | Known | Screen layout |
| 0x0063C11C | `PhotosGL_Camera_Screen_Paused` | Known | Screen layout |
| 0x0063C158 | `PhotosGL_Camera_Alt_Screen_Paused` | Known | Screen layout |
| 0x0063C194 | `PhotosGL_Camera_Screen_Playing` | Known | Screen layout |
| 0x0063C1D0 | `PhotosGL_Camera_Alt_Screen_Playing` | Known | Screen layout |
| 0x0063C264 | `PhotosGL_Camera_Screen_Volume` | Known | Screen layout |
| 0x0063C2A0 | `PhotosGL_Camera_Alt_Screen_Volume` | Known | Screen layout |
| 0x0063C6D8 | `PhotosGL_TvOut_Screen_Default` | Known | Screen layout |
| 0x0063C71C | `PhotosGL_TvOut_NTSC_Screen_Default_Video` | Known | Screen layout |
| 0x0063C76C | `PhotosGL_TvOut_PAL_Screen_Default_Video` | Known | Screen layout |
| 0x0063C794 | `PhotosGL_TvOut_Screen_Paused` | Known | Screen layout |
| 0x0063C7D8 | `PhotosGL_TvOut_NTSC_Screen_Paused_Video` | Known | Screen layout |
| 0x0063C824 | `PhotosGL_TvOut_PAL_Screen_Paused_Video` | Known | Screen layout |
| 0x0063C84C | `PhotosGL_TvOut_Screen_Playing` | Known | Screen layout |
| 0x0063C890 | `PhotosGL_TvOut_NTSC_Screen_Playing_Video` | Known | Screen layout |
| 0x0063C8E0 | `PhotosGL_TvOut_PAL_Screen_Playing_Video` | Known | Screen layout |
| 0x0063C908 | `PhotosGL_TvOut_Screen_Volume` | Known | Screen layout |
| 0x0063C94C | `PhotosGL_TvOut_NTSC_Screen_Volume_Video` | Known | Screen layout |
| 0x0063C998 | `PhotosGL_TvOut_PAL_Screen_Volume_Video` | Known | Screen layout |
| 0x0063CA18 | `SlideshowAlt_Screen` | Known | Screen layout |
| 0x0063CA48 | `Slideshow_Screen` | Known | Screen layout |
| 0x0063CA78 | `SlideshowAlt_Screen_Default` | Known | Screen layout |
| 0x0063CAAC | `Slideshow_Screen_Default` | Known | Screen layout |
| 0x0063CBB0 | `Radio_TunerContextMenu_Screen` | Known | Screen layout |
| 0x0063CBD0 | `Radio_TunerContextMenu_Screen_Default` | Known | Screen layout |
| 0x0063CC1C | `Radio_TunerTagContextMenu_Screen` | Known | Screen layout |
| 0x0063CC40 | `Radio_TunerTagContextMenu_Screen_Default` | Known | Screen layout |
| 0x0063D0B8 | `Radio_NowPlaying_Screen` | Known | Screen layout |
| 0x0063D170 | `Radio_PresetList_Screen` | Known | Screen layout |
| 0x0063D188 | `Radio_PresetList_Screen_Default` | Known | Screen layout |
| 0x0063D1C4 | `Radio_TagList_Screen` | Known | Screen layout |
| 0x0063D1DC | `Radio_TagList_Screen_Default` | Known | Screen layout |
| 0x0063D218 | `Radio_TrackHistory_Screen` | Known | Screen layout |
| 0x0063D234 | `Radio_TrackHistory_Screen_Default` | Known | Screen layout |
| 0x0063D274 | `SettingsMenus_RadioRegions_Screen` | Known | Screen layout |
| 0x0063D3B4 | `RemoteUI_Screen_Main_With_Unsupported_Firewire` | Known | Screen layout |
| 0x0063D408 | `RemoteUI_Screen_Main` | Known | Screen layout |
| 0x0063D448 | `RemoteUI_Screen_DisplayImage` | Known | Screen layout |
| 0x0063D4AC | `RemoteUI_Screen_DisplayImage_With_Unsupported_Firewire` | Known | Screen layout |
| 0x0063D534 | `RemoteUI_Hercules_ScreenLayout_Recording` | Known | Screen layout |
| 0x0063D578 | `NikePlus_History_WorkoutSummary_Screen` | Known | Screen layout |
| 0x0063D5A0 | `NikePlus_History_WorkoutSummary_Screen_Hercules` | Known | Screen layout |
| 0x0063D75C | `SettingsMenu_AboutWithAccessory_Screen` | Known | Screen layout |
| 0x0063D784 | `SettingsMenu_AboutWithAccessory_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0063D7E4 | `SettingsMenus_ResetAllSettings_Screen` | Known | Screen layout |
| 0x0063D864 | `SettingsMenu_About_Screen` | Known | Screen layout |
| 0x0063D880 | `SettingsMenu_About_Screen_Capacity_Layout` | Known | Screen layout |
| 0x0063D974 | `SettingsMenus_General_Screen` | Known | Screen layout |
| 0x0063D9D8 | `Settings_DateTime_Screen` | Known | Screen layout |
| 0x0063D9F4 | `Settings_DateTime_Screen_Default` | Known | Screen layout |
| 0x0063DA6C | `Settings_Legal_Screen` | Known | Screen layout |
| 0x0063DA84 | `Settings_Legal_Screen_Default` | Known | Screen layout |
| 0x0063DB08 | `VolumeLimitLock_Screen` | Known | Screen layout |
| 0x0063DB20 | `VolumeLimitLock_Screen_Locked` | Known | Screen layout |
| 0x0063DB70 | `VolumeLimitLock_Screen_Unlocked` | Known | Screen layout |
| 0x0063DBAC | `SettingsMenus_EQ_Screen` | Known | Screen layout |
| 0x0063DD48 | `SettingsMenus_Backlight_Screen` | Known | Screen layout |
| 0x0063DE80 | `SettingsMenus_Brightness_Screen` | Known | Screen layout |
| 0x0063DF18 | `SettingsMenus_MainMenu_Screen` | Known | Screen layout |
| 0x0063DF7C | `SettingsMenus_MusicMenu_Screen` | Known | Screen layout |
| 0x0063E44C | `SettingsMenus_ResetMainMenu_Screen` | Known | Screen layout |
| 0x0063E628 | `SettingsMenus_ResetMusicMenu_Screen` | Known | Screen layout |
| 0x0063EC74 | `Stopwatch_Screen_Running` | Known | Screen layout |
| 0x0063ECC4 | `Stopwatch_DeleteConfirmation_Screen` | Known | Screen layout |
| 0x0063ECE8 | `Stopwatch_DeleteConfirmation_Screen_Default` | Known | Screen layout |
| 0x0063ED94 | `Stopwatch_Menu_Screen` | Known | Screen layout |
| 0x0063EDAC | `Stopwatch_Menu_Screen_Default` | Known | Screen layout |
| 0x0063EDEC | `Stopwatch_SessionSummary_Screen_Default` | Known | Screen layout |
| 0x0063EE38 | `NikePlus_EquipmentAlert_Screen` | Known | Screen layout |
| 0x0063EECC | `NikePlus_Alert_Screen` | Known | Screen layout |
| 0x0063EF50 | `NikePlus_EndPausedWorkout_Screen` | Known | Screen layout |
| 0x0063EF74 | `NikePlus_EndPausedWorkout_Screen_QuickstartSave_Layout` | Known | Screen layout |
| 0x0063F000 | `NikePlus_New_Workout_Screen` | Known | Screen layout |
| 0x0063F01C | `NikePlus_New_Workout_Screen_Default` | Known | Screen layout |
| 0x0063F060 | `NikePlus_NowRunning_Screen` | Known | Screen layout |
| 0x0063F07C | `NikePlus_NowRunning_Screen_Basic` | Known | Screen layout |
| 0x0063F0CC | `NikePlus_NowRunning_Screen_Landscape` | Known | Screen layout |
| 0x0063F0F4 | `NikePlus_NowRunning_Screen_Basic_Landscape` | Known | Screen layout |
| 0x0063F140 | `NikePlus_SensorSearching_Screen` | Known | Screen layout |
| 0x0063F160 | `NikePlus_SensorSearching_Screen_Default` | Known | Screen layout |
| 0x0063F1A8 | `NikePlus_History_Screen` | Known | Screen layout |
| 0x0063F1C0 | `NikePlus_History_Screen_Default` | Known | Screen layout |
| 0x0063F200 | `NikePlus_Settings_Screen` | Known | Screen layout |
| 0x0063F21C | `NikePlus_Settings_Screen_Default` | Known | Screen layout |
| 0x0063F4CC | `NikePlus_Workout_Music_Screen` | Known | Screen layout |
| 0x0063F4EC | `NikePlus_Workout_Music_Screen_Default` | Known | Screen layout |
| 0x0063F514 | `NikePlus_Custom_Screen_Weight_ToWorkoutMusic` | Known | Screen layout |
| 0x0063F578 | `NikePlus_Dynamic_Workout_Screen` | Known | Screen layout |
| 0x0063F598 | `NikePlus_Dynamic_Workout_Screen_CalorieWorkout` | Known | Screen layout |
| 0x0063F5FC | `NikePlus_Dynamic_Workout_Screen_DistanceWorkout` | Known | Screen layout |
| 0x0063F660 | `NikePlus_Dynamic_Workout_Screen_TimedWorkout` | Known | Screen layout |
| 0x0063F6BC | `NikePlus_EndPausedWorkout_Screen_BasicSave_Layout` | Known | Screen layout |
| 0x0063F720 | `NikePlus_EndPausedWorkout_Screen_CaloriesSave_Layout` | Known | Screen layout |
| 0x0063F788 | `NikePlus_EndPausedWorkout_Screen_DistanceSave_Layout` | Known | Screen layout |
| 0x0063F7EC | `NikePlus_EndPausedWorkout_Screen_TimeSave_Layout` | Known | Screen layout |
| 0x0063FA04 | `NikePlus_StartWorkout_Screen` | Known | Screen layout |
| 0x0063FA24 | `NikePlus_StartWorkout_Screen_Resume` | Known | Screen layout |
| 0x0063FA74 | `NikePlus_StartWorkout_Screen_Default` | Known | Screen layout |
| 0x0063FAFC | `NikePlus_StartCalibration_Screen` | Known | Screen layout |
| 0x0063FB20 | `NikePlus_StartCalibration_Screen_Default` | Known | Screen layout |
| 0x0063FB4C | `NikePlus_Sensor_Screen` | Known | Screen layout |
| 0x0063FB64 | `NikePlus_Linking_Screen` | Known | Screen layout |
| 0x0063FBE0 | `NikePlus_EquipmentAlert_Screen_Default` | Known | Screen layout |
| 0x0063FC08 | `NikePlus_Remote_Screen` | Known | Screen layout |
| 0x0063FC54 | `NikePlus_Remote_Linking_Screen` | Known | Screen layout |
| 0x0063FC90 | `NikePlus_HeartMonitor_Screen` | Known | Screen layout |
| 0x0063FCF4 | `NikePlus_HeartMonitor_Linking_Screen` | Known | Screen layout |
| 0x0063FE5C | `NikePlus_ActivityStopped_Screen_Contextual_FoundSensor` | Known | Screen layout |
| 0x0063FE94 | `NikePlus_ActivityStopped_Screen_Contextual_FoundSensor_Default` | Known | Screen layout |
| 0x0063FF14 | `NikePlus_ActivityStopped_Screen_Contextual_NoSensor` | Known | Screen layout |
| 0x0063FF48 | `NikePlus_ActivityStopped_Screen_Contextual_NoSensor_Default` | Known | Screen layout |
| 0x0063FFA8 | `NikePlus_End_WorkoutSummary_Screen` | Known | Screen layout |
| 0x0063FFCC | `NikePlus_End_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x00640014 | `NikePlus_HeartMonitor_LinkingInitial_Screen` | Known | Screen layout |
| 0x00640040 | `NikePlus_HeartMonitor_LinkingInitial_Screen_Default` | Known | Screen layout |
| 0x00640074 | `NikePlus_Remote_Unlinking_Screen` | Known | Screen layout |
| 0x00640098 | `NikePlus_Remote_Unlinking_Screen_Default` | Known | Screen layout |
| 0x006400C4 | `NikePlus_HeartMonitor_Unlinking_Screen` | Known | Screen layout |
| 0x006400EC | `NikePlus_HeartMonitor_Unlinking_Screen_Default` | Known | Screen layout |
| 0x0064011C | `NikePlus_Custom_Screen_CalibrationDistance` | Known | Screen layout |
| 0x00640284 | `NikePlus_Songs_Screen` | Known | Screen layout |
| 0x0064029C | `NikePlus_Songs_Screen_Default` | Known | Screen layout |
| 0x0064030C | `Pedometer_Step_Goal_Screen` | Known | Screen layout |
| 0x00640328 | `Pedometer_Step_Goal_Screen_Default_Layout` | Known | Screen layout |
| 0x0064040C | `NikePlus_Custom_Screen_Weight` | Known | Screen layout |
| 0x006408E4 | `NikePlus_PowerPlaylist_Screen` | Known | Screen layout |
| 0x00640904 | `NikePlus_PowerPlaylist_Screen_Default` | Known | Screen layout |
| 0x00640ADC | `NikePlus_Calibration_ChooseCalibration_Screen` | Known | Screen layout |
| 0x00640D28 | `NikePlus_NowRunning_Idle_Portrait_Screen` | Known | Screen layout |
| 0x00640F20 | `NikePlus_CalibrationCompleteError_Screen` | Known | Screen layout |
| 0x00640F4C | `NikePlus_CalibrationCompleteError_Screen_Default` | Known | Screen layout |
| 0x00640FA8 | `NikePlus_CalibrationComplete_Screen_Pacing` | Known | Screen layout |
| 0x00640FFC | `NikePlus_CalibrationCompleteSuccess_Screen` | Known | Screen layout |
| 0x00641028 | `NikePlus_CalibrationCompleteSuccess_Screen_Default` | Known | Screen layout |
| 0x0064107C | `NikePlus_EndWorkout_Screen_Contextual` | Known | Screen layout |
| 0x006410A4 | `NikePlus_EndWorkout_Screen_Contextual_Default` | Known | Screen layout |
| 0x006410F4 | `NikePlus_ActivityStopped_Screen` | Known | Screen layout |
| 0x00641114 | `NikePlus_ActivityStopped_Screen_Default` | Known | Screen layout |
| 0x0064115C | `Nike_Volume_Screen` | Known | Screen layout |
| 0x00641170 | `Nike_Volume_Screen_Default` | Known | Screen layout |
| 0x00641204 | `NikePlus_NowRunning_Idle_Landscape_Screen` | Known | Screen layout |
| 0x006413AC | `NikePlus_EndWorkout_Screen_Calibration_Contextual_Landscape` | Known | Screen layout |
| 0x006413E8 | `NikePlus_EndWorkout_Screen_Calibration_Contextual_Landscape_Default` | Known | Screen layout |
| 0x00641454 | `NikePlus_EndWorkout_Screen_Contextual_Landscape` | Known | Screen layout |
| 0x00641484 | `NikePlus_EndWorkout_Screen_Contextual_Default_L` | Known | Screen layout |
| 0x006414B4 | `Nike_Volume_Screen_Landscape` | Known | Screen layout |
| 0x006414D4 | `Nike_Volume_Screen_Landscape_Default` | Known | Screen layout |
| 0x0064151C | `NikePlus_Audiobooks_Screen` | Known | Screen layout |
| 0x00641538 | `NikePlus_Audiobooks_Screen_WithAuthorName` | Known | Screen layout |
| 0x006415C4 | `NikePlus_Playlists_Screen` | Known | Screen layout |
| 0x006415E0 | `NikePlus_Playlists_Screen_Default` | Known | Screen layout |
| 0x00641604 | `NikePlus_Podcasts_Screen` | Known | Screen layout |
| 0x00641620 | `NikePlus_Podcasts_Screen_WithEpisodeCount` | Known | Screen layout |
| 0x00641778 | `Nike_PodcastEpisodes_Screen` | Known | Screen layout |
| 0x00641794 | `Nike_PodcastEpisodes_Screen_Plain` | Known | Screen layout |
| 0x006417B8 | `NikePlus_AudiobookChapters_Screen` | Known | Screen layout |
| 0x006417DC | `NikePlus_AudiobookChapters_Screen_Plane` | Known | Screen layout |
| 0x00641824 | `NikePlus_Calendar_Screen` | Known | Screen layout |
| 0x00641840 | `NikePlus_Calendar_Screen_Default` | Known | Screen layout |
| 0x0064189C | `NikePlus_History_Totals_Screen` | Known | Screen layout |
| 0x006418BC | `NikePlus_History_Totals_Screen_Default` | Known | Screen layout |
| 0x00641900 | `NikePlus_History_BestWorkouts_Screen` | Known | Screen layout |
| 0x00641928 | `NikePlus_History_BestWorkouts_Screen_Default` | Known | Screen layout |
| 0x0064198C | `NikePlus_History_WorkoutSummary_Screen_Pedometer_Session` | Known | Screen layout |
| 0x006419FC | `NikePlus_History_WorkoutSummary_Screen_Pedometer_Ambient` | Known | Screen layout |
| 0x00641A5C | `NikePlus_History_WorkoutSummary_Screen_Default` | Known | Screen layout |
| 0x00641AA8 | `NikePlus_History_ClearBests_Screen` | Known | Screen layout |
| 0x00641ACC | `NikePlus_History_ClearBests_Screen_Default` | Known | Screen layout |
| 0x00641B0C | `NikePlus_History_ClearTotals_Screen` | Known | Screen layout |
| 0x00641B30 | `NikePlus_History_ClearTotals_Screen_Default` | Known | Screen layout |
| 0x00641BCC | `NikePlus_SimpleCalibration_Dialog_Screen` | Known | Screen layout |
| 0x00641BF8 | `NikePlus_SimpleCalibration_Run_Dialog_Screen` | Known | Screen layout |
| 0x00641C4C | `NikePlus_SimpleCalibration_Walk_Dialog_Screen` | Known | Screen layout |
| 0x00641CAC | `NikePlus_History_Screen_Contextual` | Known | Screen layout |
| 0x00641CD0 | `NikePlus_History_Screen_Contextual_Default` | Known | Screen layout |
| 0x00641DBC | `NikePlus_DeleteAllWorkouts_Confirmation_Dialog_Screen` | Known | Screen layout |
| 0x00641E68 | `NikePlus_History_Day_Workouts_Screen` | Known | Screen layout |
| 0x00641EB8 | `NikePlus_Daily_landscape_Screen` | Known | Screen layout |
| 0x00641ED8 | `NikePlus_Daily_landscape_Screen_Default` | Known | Screen layout |
| 0x00642368 | `NikePlus_Calibrate_ResetToDefault_Screen` | Known | Screen layout |
| 0x00642394 | `NikePlus_Calibrate_ResetToDefault_Screen_Walking` | Known | Screen layout |
| 0x006423E4 | `NikePlus_Calibrate_ResetToDefault_Screen_Running` | Known | Screen layout |
| 0x00642458 | `NikePlus_NowRunning_Screen_Calibrate` | Known | Screen layout |
| 0x00642480 | `NikePlus_NowRunning_Screen_Calibrate_Landscape` | Known | Screen layout |
| 0x006424B0 | `NikePlus_SimpleCalibration_Screen` | Known | Screen layout |
| 0x006424D4 | `NikePlus_Custom_Screen_Simple_CalibrationDistance` | Known | Screen layout |
| 0x0064255C | `NikePlus_IsLinked_Screen` | Known | Screen layout |
| 0x00642578 | `NikePlus_IsLinked_Screen_Default` | Known | Screen layout |
| 0x00642910 | `NikePlus_Custom_StepGoal_Screen` | Known | Screen layout |
| 0x00642964 | `Debug_UnitTest_Screen` | Known | Screen layout |
| 0x0064297C | `Debug_UnitTest_Screen_Default` | Known | Screen layout |
| 0x006429B8 | `DemoMode_Screen` | Known | Screen layout |
| 0x006429C8 | `DemoMode_Screen_Default` | Known | Screen layout |
| 0x006429FC | `Debug_TestList_Screen` | Known | Screen layout |
| 0x00642A14 | `Debug_TestList_Screen_Default` | Known | Screen layout |
| 0x00642A50 | `Debug_TestResult_Screen` | Known | Screen layout |
| 0x00642A68 | `Debug_TestResult_Screen_Default` | Known | Screen layout |
| 0x00642AFC | `VoiceMemos_RecordingList_Menu_Screen` | Known | Screen layout |
| 0x00642B24 | `VoiceMemos_RecordingList_Menu_Screen_Default` | Known | Screen layout |
| 0x00642B78 | `VoiceMemos_No_Content_Screen` | Known | Screen layout |
| 0x00642CD4 | `VoiceMemos_Screen_Recording` | Known | Screen layout |
| 0x00642D10 | `VoiceMemos_Status_Screen` | Known | Screen layout |
| 0x00642D2C | `VoiceMemos_Status_Screen_Inserted` | Known | Screen layout |
| 0x00642D70 | `VoiceMemos_Context_Menu_Screen` | Known | Screen layout |
| 0x00642D90 | `VoiceMemos_Menu_Screen_Paused` | Known | Screen layout |
| 0x00642DD0 | `VoiceMemos_Screen_Saving` | Known | Screen layout |
| 0x00642EB0 | `VoiceMemos_Screen_DeleteAllAsk` | Known | Screen layout |
| 0x00642ED0 | `VoiceMemos_Screen_DeleteAllAsk_Default` | Known | Screen layout |
| 0x00642F10 | `VoiceMemos_Context_PlayDelete_Screen` | Known | Screen layout |
| 0x00642F38 | `VoiceMemos_Context_Menu_Screen_PlayDelete` | Known | Screen layout |
| 0x00642FF4 | `VoiceMemos_Label_Select_Screen` | Known | Screen layout |
| 0x00643014 | `VoiceMemos_Label_Select_Screen_Default` | Known | Screen layout |
| 0x00643188 | `Search_Main_Screen_WithKeyboardNow` | Known | Screen layout |
| 0x006431E0 | `Speakers_ToneControl_Screen` | Known | Screen layout |
| 0x006431FC | `Speakers_ToneControl_Screen_Default` | Known | Screen layout |
| 0x006432A8 | `DiskMode_ScreenLayout_Disconnected` | Known | Screen layout |
| 0x006432E8 | `DiskMode_ScreenLayout_Loading` | Known | Screen layout |
| 0x0064332C | `DiskMode_ScreenLayout_Synchronizing` | Known | Screen layout |
| 0x00643370 | `DiskMode_ScreenLayout_Connected` | Known | Screen layout |
| 0x006433AC | `ChargingMode_Progress_ScreenLayout_Charged` | Known | Screen layout |
| 0x006433F4 | `ChargingMode_Progress_ScreenLayout_Charging` | Known | Screen layout |

---

## 8. RTOS Architecture (RTXC)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00000628 | `Amici-1.0.1 1.0.1 34A10006 RTXC SCM Administrator@w02 2009/08/28 12:48:27 CL1522` | Known | RTOS |
| 0x00493BB0 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Known | RTOS |
| 0x0052370A | `N3ISL13TRFTuner_RTXCE` | Known | RTOS |
| 0x00523B45 | `N3ISL20TLocationDevice_RTXCE` | Known | RTOS |

---

## 9. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00005F78 | `BootTask` | Known | RTOS task thread |
| 0x0003554C | `USBDeviceTask` | Known | RTOS task thread |
| 0x000429AC | `FirewireTask` | Known | RTOS task thread |
| 0x000429C0 | `TouchwheelTask` | Known | RTOS task thread |
| 0x000429E8 | `DiskMgrTask` | Known | RTOS task thread |
| 0x000429F8 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x00042A0C | `MikeyTask` | Known | RTOS task thread |
| 0x00042A1C | `RadioTask` | Known | RTOS task thread |
| 0x00042A94 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x00042ABC | `AlarmTask` | Known | RTOS task thread |
| 0x00042ADB | `"USBAudioTask` | Known | RTOS task thread |
| 0x00042AF0 | `ChargeMgmtTask` | Known | RTOS task thread |
| 0x00048F18 | `Terminator Task` | Known | RTOS task thread |
| 0x0004C96C | `MainAppTask` | Known | RTOS task thread |
| 0x000C540C | `TLogPedDiskWritingTask` | Known | RTOS task thread |
| 0x000E846C | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x00145774 | `DiskReaderTask` | Known | RTOS task thread |
| 0x00152FB4 | `MeCCABufferedRDSUpdateTask` | Known | RTOS task thread |
| 0x001A1750 | `TPodMediaPlayerFileUpdate Task` | Known | RTOS task thread |
| 0x001A4BD0 | `TTrainerApp_LocaleChangedLoadingTask` | Known | RTOS task thread |
| 0x001B7C1C | `GeniusMixesTask` | Known | RTOS task thread |
| 0x0029B584 | `MeCCARecordingTask` | Known | RTOS task thread |
| 0x002BB09C | `MeCCAInputTask` | Known | RTOS task thread |
| 0x002BB0B0 | `MeCCAOutputTask` | Known | RTOS task thread |
| 0x002DF170 | `InputBufferLoadTask` | Known | RTOS task thread |
| 0x002ECB54 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x00306BE0 | `StreamCacheMassStorageManagerTimeOutTask` | Known | RTOS task thread |
| 0x00329AA0 | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x00350CCC | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x00350CE0 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0037CF3C | `Task` | Known | RTOS task thread |
| 0x00460CA4 | `HostOSTask` | Known | RTOS task thread |

---

## 10. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0011AD90 | `Channel Reserved` | Known | Logging channel |
| 0x0011ADA4 | `Channel AppBoot` | Known | Logging channel |
| 0x0011ADB4 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x0011ADD0 | `Channel PrefsWriting` | Known | Logging channel |
| 0x0011ADE8 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x0011AE08 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x0011AE20 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x0011AE3C | `Channel TestLogging` | Known | Logging channel |
| 0x0011AE50 | `Channel AppFileLoading` | Known | Logging channel |
| 0x0011AE68 | `Channel VCardReading` | Known | Logging channel |
| 0x0011AE80 | `Channel LongSongScanning` | Known | Logging channel |
| 0x0011AE9C | `Channel VoiceRecording` | Known | Logging channel |
| 0x0011AEB4 | `Channel PhotoImporting` | Known | Logging channel |
| 0x0011AECC | `Channel Notes` | Known | Logging channel |
| 0x0011AEDC | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x0011AEF8 | `Channel DiskMode` | Known | Logging channel |
| 0x0011AF0C | `Channel Firewire` | Known | Logging channel |
| 0x0011AF20 | `Channel USB` | Known | Logging channel |
| 0x0011AF40 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x0011AF58 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 11. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001EE240 | `gamedata_RW` | Known | Game system |
| 0x001EE25C | `gamedata_ShareRW` | Known | Game system |
| 0x001EE270 | `games_RO` | Known | Game system |
| 0x0051D05B | `11TCGamesMenu` | Known | Game system |
| 0x0051D169 | `12TCGameScreen` | Known | Game system |
| 0x0051D4C1 | `14TCGameControls` | Known | Game system |
| 0x0051F458 | `27TSilverCntlrTransitionAddonI11TCGamesMenuE` | Known | Game system |
| 0x0051F53B | `27TSilverCntlrTransitionAddonI12TCGameScreenE` | Known | Game system |
| 0x0051F769 | `27TSilverCntlrTransitionAddonI14TCGameControlsE` | Known | Game system |
| 0x0056B098 | `iPod_Control/games_RO/` | Known | Game system |
| 0x0056B0AF | `Resources/Games/games_RO/` | Known | Game system |

---

## 12. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00035B5C | `AppleDRMVersion` | Known | DRM system |
| 0x000364D8 | `AppleDRM` | Known | DRM system |
| 0x000368A8 | `AppleVideoDRM` | Known | DRM system |
| 0x000368E0 | `AppleDRM` | Known | DRM system |
| 0x00056384 | `FairPlayDeviceType` | Known | DRM system |
| 0x0018FC14 | `adrmmp4a` | Known | DRM system |
| 0x00190E50 | `drmttx3gp` | Known | DRM system |
| 0x00219AB8 | `tx3gdrmsdrmip608aavdmp4aesds` | Known | DRM system |
| 0x0056C64B | `DRMLevel` | Known | DRM system |
| 0x0069C0F4 | `$Apple FairPlay Certificate Authority0` | Known | DRM system |
| 0x0069C479 | `&Apple FairPlay Certification Authority0` | Known | DRM system |

---

## 13. Database System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0005634C | `SQLiteDB` | Known | SQLite database |
| 0x000BF50C | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x000DEC44 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x000DF024 | `iPod_Control\iTunes\Play Counts` | Known | iTunes database |
| 0x0019BF52 | `pGiPod_Control/iTunes/iTunes Library.itlp/Dynamic.itdb` | Known | iTunes database |
| 0x0019BF8E | `pGiPod_Control/iTunes/iTunes Library.itlp/Library.itdb` | Known | iTunes database |
| 0x001A2D54 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x001D3920 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x001D3938 | `iPod_Control\iTunes\firsttime` | Known | iTunes database |
| 0x001D7C48 | `iPod_Control\iTunes\iTunesDB` | Known | iTunes database |
| 0x001D7C6C | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x001EBA18 | `sqlite_temp_master` | Known | SQLite database |
| 0x001EBA2C | `sqlite_master` | Known | SQLite database |
| 0x001EF784 | `sqlite_stat1` | Known | SQLite database |
| 0x001EF794 | `CREATE TABLE %Q.sqlite_stat1(tbl,idx,stat)` | Known | SQLite database |
| 0x001EF7C0 | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x00203C50 | `sqlite_temp_master` | Known | SQLite database |
| 0x00203C64 | `sqlite_master` | Known | SQLite database |
| 0x00203FE4 | `iPod_Control/iTunes/` | Known | iTunes database |
| 0x002044DC | `sqlite_temp_master` | Known | SQLite database |
| 0x002044F0 | `sqlite_master` | Known | SQLite database |
| 0x00210E6C | `sqlite_temp_master` | Known | SQLite database |
| 0x00210E80 | `sqlite_master` | Known | SQLite database |
| 0x002204A8 | `sqlite_autoindex_` | Known | SQLite database |
| 0x002204BC | `sqlite_temp_master` | Known | SQLite database |
| 0x002204D0 | `sqlite_master` | Known | SQLite database |
| 0x002225BC | `sqlite3BtreeInitPage() returns error code %d` | Known | SQLite database |
| 0x00223420 | `sqlite_temp_master` | Known | SQLite database |
| 0x00223434 | `sqlite_master` | Known | SQLite database |
| 0x00223448 | `CREATE TABLE %Q.sqlite_sequence(name,seq)` | Known | SQLite database |
| 0x002249E4 | `sqlite_stat1` | Known | SQLite database |
| 0x002249F4 | `SELECT idx, stat FROM %Q.sqlite_stat1` | Known | SQLite database |
| 0x0022EBC8 | `sqlite_subquery_%p_` | Known | SQLite database |
| 0x00230330 | `sqlite_temp_master` | Known | SQLite database |
| 0x00230344 | `sqlite_master` | Known | SQLite database |
| 0x00230390 | `sqlite_sequence` | Known | SQLite database |
| 0x002315C0 | `sqlite_` | Known | SQLite database |
| 0x0023F3EC | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0028DFF8 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0028E010 | `iTunesDB` | Known | iTunes database |
| 0x002B4A1A | `pGiPod_Control/iTunes/iTunes Library.itlp/Extras.itdb` | Known | iTunes database |
| 0x002B4A52 | `pGiPod_Control/iTunes/iTunes Library.itlp/Genius.itdb` | Known | iTunes database |
| 0x002B4A8A | `pGiPod_Control/iTunes/iTunes Library.itlp/Locations.itdb` | Known | iTunes database |
| 0x002BE9EC | `sqlite3_extension_init` | Known | SQLite database |
| 0x002C83D4 | `sqlite_attach` | Known | SQLite database |
| 0x002C83E8 | `sqlite_detach` | Known | SQLite database |
| 0x002D74C0 | `%s/sqlite_` | Known | SQLite database |
| 0x002DA350 | `sqlite_attach` | Known | SQLite database |
| 0x002DA360 | `sqlite_detach` | Known | SQLite database |
| 0x002E1F9C | `sqlite_temp_master` | Known | SQLite database |
| 0x002E1FB0 | `sqlite_master` | Known | SQLite database |
| 0x002E21B4 | `sqlite_` | Known | SQLite database |
| 0x002E21F4 | `sqlite_temp_master` | Known | SQLite database |
| 0x002E2208 | `sqlite_master` | Known | SQLite database |
| 0x002E221C | `sqlite_sequence` | Known | SQLite database |
| 0x002E222C | `UPDATE "%w".sqlite_sequence set name = %Q WHERE name = %Q` | Known | SQLite database |
| 0x002E265C | `sqlite_` | Known | SQLite database |
| 0x002E26F4 | `sqlite_temp_master` | Known | SQLite database |
| 0x002E2708 | `sqlite_master` | Known | SQLite database |
| 0x002E29B4 | `sqlite_temp_master` | Known | SQLite database |
| 0x002E29C8 | `sqlite_master` | Known | SQLite database |
| 0x002E29F8 | `sqlite_stat1` | Known | SQLite database |
| 0x002E2A08 | `DELETE FROM %Q.sqlite_stat1 WHERE idx=%Q` | Known | SQLite database |
| 0x002E2CAC | `sqlite_temp_master` | Known | SQLite database |
| 0x002E2CC0 | `sqlite_master` | Known | SQLite database |
| 0x002E2D30 | `DELETE FROM %s.sqlite_sequence WHERE name=%Q` | Known | SQLite database |
| 0x002E2D98 | `sqlite_stat1` | Known | SQLite database |
| 0x002E2DA8 | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x002E32C0 | `sqlite_temp_master` | Known | SQLite database |
| 0x002E32D4 | `sqlite_master` | Known | SQLite database |
| 0x002E6EDC | `sqlite_temp_master` | Known | SQLite database |
| 0x002E6EF0 | `sqlite_master` | Known | SQLite database |
| 0x002E7944 | `sqlite_temp_master` | Known | SQLite database |
| 0x002E7958 | `sqlite_master` | Known | SQLite database |
| 0x00476840 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x00476880 | `iPod_Control\iTunes\` | Known | iTunes database |
| 0x0047723F | `SQLite format 3` | Known | SQLite database |
| 0x004798EC | `CREATE TABLE sqlite_master(` | Known | SQLite database |
| 0x00479954 | `CREATE TEMP TABLE sqlite_temp_master(` | Known | SQLite database |
| 0x0047A01C | `illegal return value (%d) from the authorization function - should be SQLITE_OK,` | Known | SQLite database |
| 0x0047A0D4 | `SELECT 'CREATE TABLE vacuum_db.' || substr(sql,14)   FROM sqlite_master WHERE ty` | Known | SQLite database |
| 0x0047A15C | `SELECT 'CREATE INDEX vacuum_db.' || substr(sql,14)  FROM sqlite_master WHERE sql` | Known | SQLite database |
| 0x0047A1C4 | `SELECT 'CREATE UNIQUE INDEX vacuum_db.' || substr(sql,21)   FROM sqlite_master W` | Known | SQLite database |
| 0x0047A23C | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x0047A2EC | `SELECT 'DELETE FROM vacuum_db.' || quote(name) || ';' FROM vacuum_db.sqlite_mast` | Known | SQLite database |
| 0x0047A360 | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x0047A3F8 | `INSERT INTO vacuum_db.sqlite_master   SELECT type, name, tbl_name, rootpage, sql` | Known | SQLite database |
| 0x0047A5B8 | `UPDATE %Q.%s SET sql = CASE WHEN type = 'trigger' THEN sqlite_rename_trigger(sql` | Known | SQLite database |
| 0x0047A72C | `UPDATE sqlite_temp_master SET sql = sqlite_rename_trigger(sql, %Q), tbl_name = %` | Known | SQLite database |
| 0x0047A940 | `sqlite3_get_table() called with two or more incompatible queries` | Known | SQLite database |
| 0x0056CA4E | `sqlite_rename_table` | Known | SQLite database |
| 0x0056CBCD | `sqlite_version` | Known | SQLite database |
| 0x0056CC67 | `sqlite_rename_trigger` | Known | SQLite database |
| 0x00575827 | `SQLite_iPod_VFS` | Known | SQLite database |
| 0x00579338 | `CREATE TABLE _SqliteDatabaseProperties (key TEXT, value TEXT, UNIQUE(key));` | Known | SQLite database |
| 0x0062E434 | `Richard Hipp (SQLite) SQLite Copyright` | Known | SQLite database |
| 0x0062E45C | `All of the deliverable code in SQLite has been dedicated to the public domain by` | Known | SQLite database |
| 0x0062E688 | `The previous paragraph applies to the deliverable code in SQLite - those parts o` | Known | SQLite database |
| 0x0062E858 | `All of the deliverable code in SQLite has been written from scratch. No code has` | Known | SQLite database |
| 0x0062E9C4 | `Obtaining An Explicit License To Use SQLite` | Known | SQLite database |
| 0x0062E9F0 | `Even though SQLite is in the public domain and does not require a license, some ` | Known | SQLite database |
| 0x0062EAE8 | `-You are using SQLite in a jurisdiction that does not recognize the right of an ` | Known | SQLite database |
| 0x0062EB6C | `-You want to hold a tangible legal document as evidence that you have the legal ` | Known | SQLite database |
| 0x0062EC28 | `If you feel like you really have to purchase a license for SQLite, Hwaci, the co` | Known | SQLite database |
| 0x0062ECE8 | `In order to keep SQLite completely free and unencumbered by copyright, all new c` | Known | SQLite database |
| 0x0062EFA4 | `We are not able to accept patches or changes to SQLite that are not accompanied ` | Known | SQLite database |

---

## 14. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00052C7C | `[FTL:MSG] Apple NAND Driver (AND) RW` | Known | Hardware |
| 0x00052CF4 | `[FTL:MSG] No NAND attached` | Known | Hardware |
| 0x000562A8 | `FireWireGUID` | Known | FireWire |
| 0x000562B8 | `FireWireVersion` | Known | FireWire |
| 0x00056524 | `FireWire` | Known | FireWire |
| 0x0007A470 | `[FIL:INF] could not find NAND config in the new NAND tables` | Known | Hardware |
| 0x0007EC6C | `NANDDRIVERSIGN` | Known | Hardware |
| 0x0007EDEC | `NANDDRIVERSIGN` | Known | Hardware |
| 0x0025853C | `NANDDRIVERSIGN` | Known | Hardware |
| 0x00350C34 | `NAND FLASH DRIVE` | Known | Hardware |
| 0x00493B74 | `[NAND] Panic! %s:%d` | Known | Hardware |
| 0x00493B8C | `[NAND] Failed WMR_ASSERT(%s) %s:%d` | Known | Hardware |
| 0x00493C20 | `[NAND] %s:%d IOCtl on buffer of size %d with %d bytes of src data!` | Known | Hardware |
| 0x00493E28 | `[WMR:ERR] NAND format invalid (mismatch, corrupt, read error or blank NAND devic` | Known | Hardware |
| 0x00493F18 | `AND: NAND initialisation failed due to format mismatch or uninitialised NAND.` | Known | Hardware |
| 0x004B07B4 | `[FTL:WRN] Recovering NAND Data Structures - this will take some time!` | Known | Hardware |
| 0x004B17F8 | `(bReadEdoClocks * dwMaxNSPerClock) < (_GetReadValidNanosecs() + SOC_RISE_TIME_NS` | Known | Hardware |
| 0x0051E4E2 | `21TCFirewireUnsupported` | Known | FireWire |
| 0x005200EA | `27TSilverCntlrTransitionAddonI21TCFirewireUnsupportedE` | Known | FireWire |

---

## 15. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0051E984 | `23TCSettings_RadioRegions` | Known | FM Radio |
| 0x0051ED1C | `24TSilverRadioTunerBarView` | Known | FM Radio |
| 0x0052049F | `27TSilverCntlrTransitionAddonI23TCSettings_RadioRegionsE` | Known | FM Radio |
| 0x0056AEED | `General.RadioRegion` | Known | FM Radio |
| 0x0057208C | `Radio_ConfirmationOverlay_ChangeStation_Layout_Portrait` | Known | FM Radio |
| 0x00574834 | `Radio_NowPlaying_TunerBar_Layout` | Known | FM Radio |
| 0x00574CEC | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |
| 0x0062C618 | `Please use the built in FM tuner to listen to the Radio.` | Known | FM Radio |
| 0x0062C88C | `Radio Regions` | Known | FM Radio |
| 0x0063CDE0 | `Radio_NowPlaying_TunerBar_Layout` | Known | FM Radio |
| 0x0063D298 | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 16. Camera

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00054994 | `TPhotosGLCntlrLcdCamera` | Known | Camera |
| 0x0008A104 | `CameraShutter` | Known | Camera |
| 0x000B0A78 | `EnterCameraScreen` | Known | Camera |
| 0x000B0F68 | `EnterCameraScreen` | Known | Camera |
| 0x000C7420 | `PushScreen_BrowseCameraPhotos` | Known | Camera |
| 0x000EA2DC | `Camera Videos` | Known | Camera |
| 0x000F560C | `FinishedPopCamera` | Known | Camera |
| 0x000F596C | `PopCamera` | Known | Camera |
| 0x000F5AF4 | `PopCamera` | Known | Camera |
| 0x000F5C74 | `PopCamera` | Known | Camera |
| 0x000F5FC8 | `PopCamera` | Known | Camera |
| 0x000F6240 | `PopCamera` | Known | Camera |
| 0x000F63D0 | `PopCamera` | Known | Camera |
| 0x000F6458 | `PopCamera` | Known | Camera |
| 0x000F8BA8 | `CameraApp` | Known | Camera |
| 0x00109908 | `EmptyCameraHilited` | Known | Camera |
| 0x0010991C | `CameraHilited` | Known | Camera |
| 0x00109D4C | `CameraSelected` | Known | Camera |
| 0x0010AF9C | `CameraVideosSelected` | Known | Camera |
| 0x0012AC98 | `CameraDeviceManager` | Known | Camera |
| 0x0014FB08 | `%d: GKCameraDriver - camera overflow occurred!` | Known | Camera |
| 0x0047CFF8 | `TPhotosGLCntlrLcdCamera` | Known | Camera |
| 0x0051CF6A | `10TCameraApp` | Known | Camera |
| 0x0051CFB8 | `10TPodCamera` | Known | Camera |
| 0x0051D49F | `14GKCameraDevice` | Known | Camera |
| 0x0051D4B0 | `14GKCameraDriver` | Known | Camera |
| 0x0051D6B8 | `15TCCameraInitial` | Known | Camera |
| 0x0051DC93 | `17TCameraMediaModel` | Known | Camera |
| 0x0051E285 | `20TCCameraDeleteDialog` | Known | Camera |
| 0x0051E4CA | `21TCCameraAllVideosList` | Known | Camera |
| 0x0051E70F | `22TCCameraLocalMediaList` | Known | Camera |
| 0x0051E728 | `22TCCameraMediaList_Base` | Known | Camera |
| 0x0051E950 | `23TCCameraDeleteAllDialog` | Known | Camera |
| 0x0051EA3A | `23TPhotosGLCntlrLcdCamera` | Known | Camera |
| 0x0051EAA2 | `23TRecentCameraMediaModel` | Known | Camera |
| 0x0051EC5F | `24TCameraMediaModel_Import` | Known | Camera |
| 0x0051ECE6 | `24TSilverCameraShutterView` | Known | Camera |
| 0x0051EE47 | `25TCameraApplication_Import` | Known | Camera |
| 0x0051F07C | `26TCameraMediaDatabaseEntity` | Known | Camera |
| 0x0051F2C6 | `27TCameraMediaDatabaseContext` | Known | Camera |
| 0x0051F7F9 | `27TSilverCntlrTransitionAddonI15TCCameraInitialE` | Known | Camera |
| 0x0051FE60 | `27TSilverCntlrTransitionAddonI20TCCameraDeleteDialogE` | Known | Camera |
| 0x005200B3 | `27TSilverCntlrTransitionAddonI21TCCameraAllVideosListE` | Known | Camera |
| 0x0052026C | `27TSilverCntlrTransitionAddonI22TCCameraLocalMediaListE` | Known | Camera |
| 0x0052042D | `27TSilverCntlrTransitionAddonI23TCCameraDeleteAllDialogE` | Known | Camera |
| 0x00520583 | `27TSilverCntlrTransitionAddonI23TPhotosGLCntlrLcdCameraE` | Known | Camera |
| 0x005220AF | `27TSilverCntlrTransitionAddonI8TCCameraE` | Known | Camera |
| 0x00522415 | `29TCameraMediaDatabaseInterface` | Known | Camera |
| 0x0052273B | `30TRecentCameraMediaModel_Import` | Known | Camera |
| 0x005231A4 | `8TCCamera` | Known | Camera |
| 0x005232AF | `N10TCameraApp19TCameraStateMachineE` | Known | Camera |
| 0x00523340 | `N17TCameraMediaModel15CameraItemPropsE` | Known | Camera |
| 0x00523367 | `N17TCameraMediaModel17LaunchCameraPropsE` | Known | Camera |
| 0x00523390 | `N17TCameraMediaModel18SyncedContentPropsE` | Known | Camera |
| 0x005233BA | `N17TCameraMediaModel22UnSyncedPhotoListPropsE` | Known | Camera |
| 0x005233E8 | `N17TCameraMediaModel22UnSyncedVideoFilePropsE` | Known | Camera |
| 0x00523416 | `N17TCameraMediaModel22UnSyncedVideoListPropsE` | Known | Camera |
| 0x00523444 | `N17TCameraMediaModel23UnSyncedAllContentPropsE` | Known | Camera |
| 0x0052349B | `N24TSilverCameraShutterView25TShutterAnimationStateMsgE` | Known | Camera |
| 0x005234D3 | `N27TCameraMediaDatabaseContext16ContextualEntityE` | Known | Camera |
| 0x0052351C | `N3ISL10IPodCameraE` | Known | Camera |
| 0x0056B18A | `cameraremote` | Known | Camera |
| 0x0056DC6B | `PopCamera` | Known | Camera |
| 0x005713C2 | `PhotosGL_Camera_Delete_All_Confirmation` | Known | Camera |
| 0x005713EA | `MediaLists_Camera_Delete_All_Confirmation` | Known | Camera |
| 0x00571414 | `PhotosGL_Camera_Delete_Item_Confirmation` | Known | Camera |
| 0x0057143D | `MediaLists_Camera_Delete_Video_Confirmation` | Known | Camera |
| 0x00571469 | `MediaLists_Camera_Delete_Video_Event_Confirmation` | Known | Camera |
| 0x0057149B | `MediaLists_Camera_Delete_Photo_Event_Confirmation` | Known | Camera |
| 0x00571792 | `PhotosGL_Camera_All_Media_Delete_Menu_Alt_NoStatusBar` | Known | Camera |
| 0x005717C8 | `PhotosGL_Camera_All_Media_Delete_Menu_NoStatusBar` | Known | Camera |
| 0x00572166 | `PhotosGL_Camera_Delete_All_Confirmation_Alt` | Known | Camera |
| 0x00572192 | `PhotosGL_Camera_Delete_Item_Confirmation_Alt` | Known | Camera |
| 0x005721BF | `PhotosGL_Camera_All_Media_Delete_Menu_Alt` | Known | Camera |
| 0x00573A79 | `PhotosGL_Camera_Delete_All_Confirmation_Default` | Known | Camera |
| 0x00573AA9 | `MediaLists_Camera_Delete_All_Confirmation_Default` | Known | Camera |
| 0x00573ADB | `PhotosGL_Camera_Delete_Item_Confirmation_Default` | Known | Camera |
| 0x00573B0C | `MediaLists_Camera_Delete_Video_Confirmation_Default` | Known | Camera |
| 0x00573B40 | `MediaLists_Camera_Delete_Video_Event_Confirmation_Default` | Known | Camera |
| 0x00573B7A | `MediaLists_Camera_Delete_Photo_Event_Confirmation_Default` | Known | Camera |
| 0x00573D3E | `PhotosGL_Camera_Delete_All_Confirmation_Alt_Default` | Known | Camera |
| 0x00573D72 | `PhotosGL_Camera_Delete_Item_Confirmation_Alt_Default` | Known | Camera |
| 0x0057501B | `PhotosGL_Camera_All_Media_Delete_Menu` | Known | Camera |
| 0x00575041 | `MediaLists_Camera_All_Media_Delete_Menu` | Known | Camera |
| 0x0057507E | `PhotosGL_Camera_All_Media_Contextual_Menu` | Known | Camera |
| 0x005750A8 | `MediaLists_Camera_All_Media_Contextual_Menu` | Known | Camera |
| 0x005750D4 | `PhotosGL_Camera_All_Media_Alt_Contextual_Menu` | Known | Camera |
| 0x00575102 | `MediaLists_Camera_All_Media_Delete_All_Menu` | Known | Camera |
| 0x00579E5D | `TPhotosGLCntlrLcdCamera_DoDeleteAll` | Known | Camera |
| 0x00579E81 | `TPhotosGLCntlrLcdCamera_DoDeleteItem` | Known | Camera |
| 0x005BCEEC | `TPhotosGLCntlrLcdCamera` | Known | Camera |
| 0x00628410 | `Video Camera` | Known | Camera |
| 0x00629AA0 | `Camera Roll` | Known | Camera |
| 0x00629AAC | `Camera Videos` | Known | Camera |
| 0x00629ABC | `Video Camera` | Known | Camera |
| 0x00629B44 | `Delete all camera videos from your iPod?` | Known | Camera |
| 0x00629BE4 | `Camera Initializing` | Known | Camera |
| 0x0062AD9C | `Video Camera` | Known | Camera |
| 0x0062AE20 | `Camera Videos` | Known | Camera |
| 0x0062BE48 | `Camera Roll` | Known | Camera |
| 0x0062C07C | `Delete all camera videos from your iPod?` | Known | Camera |
| 0x0062C0A8 | `Delete this camera video from your iPod?` | Known | Camera |
| 0x006337AC | `cameraremote.action.up` | Known | Camera |
| 0x006337D8 | `cameraremote.photo.up` | Known | Camera |
| 0x006337F0 | `cameraremote.video.up` | Known | Camera |
| 0x00635700 | `PopCamera` | Known | Camera |
| 0x006357C8 | `controller.FinishedPopCamera` | Known | Camera |
| 0x00635A20 | `controller.EnterCameraScreen` | Known | Camera |
| 0x00635AD4 | `MediaLists_Camera_Delete_All_Confirmation` | Known | Camera |
| 0x00635B00 | `MediaLists_Camera_Delete_All_Confirmation_Default` | Known | Camera |
| 0x00635B74 | `MediaLists_Camera_All_Media_Contextual_Menu` | Known | Camera |
| 0x00635BA0 | `MediaLists_Camera_All_Media_Delete_All_Menu` | Known | Camera |
| 0x00635BE8 | `MediaLists_Camera_All_Media_Delete_Menu` | Known | Camera |
| 0x00635C44 | `MediaLists_Camera_Delete_Photo_Event_Confirmation` | Known | Camera |
| 0x00635C78 | `MediaLists_Camera_Delete_Photo_Event_Confirmation_Default` | Known | Camera |
| 0x00635CE4 | `MediaLists_Camera_Delete_Video_Confirmation` | Known | Camera |
| 0x00635D10 | `MediaLists_Camera_Delete_Video_Confirmation_Default` | Known | Camera |
| 0x00635D78 | `MediaLists_Camera_Delete_Video_Event_Confirmation` | Known | Camera |
| 0x00635DAC | `MediaLists_Camera_Delete_Video_Event_Confirmation_Default` | Known | Camera |
| 0x006382DC | `controller.CameraHilited` | Known | Camera |
| 0x006382F8 | `controller.CameraSelected` | Known | Camera |
| 0x006384E8 | `controller.EmptyCameraHilited` | Known | Camera |
| 0x0063A658 | `controller.CameraVideosSelected` | Known | Camera |
| 0x0063B828 | `controller.PushScreen_BrowseCameraPhotos` | Known | Camera |
| 0x0063B854 | `controller.PushScreen_BrowseCameraPhotosAlt` | Known | Camera |
| 0x0063C2DC | `PhotosGL_Camera_All_Media_Contextual_Menu` | Known | Camera |
| 0x0063C308 | `PhotosGL_Camera_All_Media_Delete_Menu` | Known | Camera |
| 0x0063C330 | `PhotosGL_Camera_Delete_All_Confirmation` | Known | Camera |
| 0x0063C358 | `PhotosGL_Camera_Delete_All_Confirmation_Default` | Known | Camera |
| 0x0063C3B0 | `PhotosGL_Camera_Delete_Item_Confirmation` | Known | Camera |
| 0x0063C3DC | `PhotosGL_Camera_Delete_Item_Confirmation_Default` | Known | Camera |
| 0x0063C434 | `PhotosGL_Camera_All_Media_Delete_Menu_NoStatusBar` | Known | Camera |
| 0x0063C4D4 | `PhotosGL_Camera_All_Media_Alt_Contextual_Menu` | Known | Camera |
| 0x0063C504 | `PhotosGL_Camera_All_Media_Delete_Menu_Alt` | Known | Camera |
| 0x0063C530 | `PhotosGL_Camera_Delete_All_Confirmation_Alt` | Known | Camera |
| 0x0063C55C | `PhotosGL_Camera_Delete_All_Confirmation_Alt_Default` | Known | Camera |
| 0x0063C590 | `PhotosGL_Camera_Delete_Item_Confirmation_Alt` | Known | Camera |
| 0x0063C5C0 | `PhotosGL_Camera_Delete_Item_Confirmation_Alt_Default` | Known | Camera |
| 0x0063C5F8 | `PhotosGL_Camera_All_Media_Delete_Menu_Alt_NoStatusBar` | Known | Camera |

---

## 17. Pedometer

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00039530 | `Pedometer` | Known | Pedometer |
| 0x00055D38 | `TTrainer_Cntlr_Ambient_Pedometer` | Known | Pedometer |
| 0x000A37E4 | `] Step: ` | Known | Pedometer |
| 0x0010A4C4 | `PedometerHilited` | Known | Pedometer |
| 0x0010A4D8 | `PedometerInactiveHilited` | Known | Pedometer |
| 0x0011A0F0 | `/Pedometer/` | Known | Pedometer |
| 0x0015022C | `TPedometerHeartbeatThread` | Known | Pedometer |
| 0x00167EC4 | `/Pedometer/` | Known | Pedometer |
| 0x0019AAC0 | `TTrainer_Cntlr_Pedometer_Goal` | Known | Pedometer |
| 0x0019AAE8 | `TPedometer_Hourly_Cntlr` | Known | Pedometer |
| 0x001B11A8 | `GoToWorkoutPedometerAmbientSummaryScreen` | Known | Pedometer |
| 0x001B11D4 | `GoToWorkoutPedometerSummaryScreen` | Known | Pedometer |
| 0x001B16DC | `GoToWorkoutPedometerAmbientSummaryScreen` | Known | Pedometer |
| 0x001B1708 | `GoToPedometerSessionWorkoutSummaryScreen` | Known | Pedometer |
| 0x00231C90 | `pedometer` | Known | Pedometer |
| 0x0024373C | `pedometer` | Known | Pedometer |
| 0x0029FE28 | `TPedometerThread` | Known | Pedometer |
| 0x002CB058 | `GoToPedometerSession` | Known | Pedometer |
| 0x002CB070 | `GoToPedometerDaily` | Known | Pedometer |
| 0x002D6CB4 | `pedometer` | Known | Pedometer |
| 0x002F3A44 | `pedometer` | Known | Pedometer |
| 0x004605AF | `TotalSteps` | Known | Pedometer |
| 0x004645D8 | `PedometerModel - No steps for ambient workout . Discarding and deleting session!` | Known | Pedometer |
| 0x0046462C | `PedometerModel - No steps for session workout . Discarding and deleting session!` | Known | Pedometer |
| 0x00464DEC | `TPedometer_Hourly_Cntlr` | Known | Pedometer |
| 0x00464EB0 | `TTrainer_Cntlr_Pedometer_Goal` | Known | Pedometer |
| 0x0047C98C | `TTrainer_Cntlr_Ambient_Pedometer` | Known | Pedometer |
| 0x004808C8 | `Stepper` | Known | Pedometer |
| 0x00480A3B | `Steps` | Known | Pedometer |
| 0x00480ACF | `pedometer` | Known | Pedometer |
| 0x00480AD9 | `ambient_pedometer` | Known | Pedometer |
| 0x0048BD5F | `<TTrainer_Cntlr_Pedometer` | Known | Pedometer |
| 0x0048BD98 | `TTrainer_Cntlr_Ambient_Pedometer` | Known | Pedometer |
| 0x0051D43F | `13TPedometerApp` | Known | Pedometer |
| 0x0051D7C2 | `15TPedometerModel` | Known | Pedometer |
| 0x0051E7F0 | `22TPedometerModel_Import` | Known | Pedometer |
| 0x0051EA20 | `23TPedometer_Hourly_Cntlr` | Known | Pedometer |
| 0x0051EB0A | `23TSilverStepBarGraphView` | Known | Pedometer |
| 0x0051ED6D | `24TTrainer_Cntlr_Pedometer` | Known | Pedometer |
| 0x0052054A | `27TSilverCntlrTransitionAddonI23TPedometer_Hourly_CntlrE` | Known | Pedometer |
| 0x0052132E | `27TSilverCntlrTransitionAddonI29TTrainer_Cntlr_Pedometer_GoalE` | Known | Pedometer |
| 0x005219FF | `27TSilverCntlrTransitionAddonI32TTrainer_Cntlr_Ambient_PedometerE` | Known | Pedometer |
| 0x005221AA | `27TTrainer_PedometerGoalModel` | Known | Pedometer |
| 0x005225B5 | `29TTrainer_Cntlr_Pedometer_Goal` | Known | Pedometer |
| 0x00522BB7 | `32TTrainer_Cntlr_Ambient_Pedometer` | Known | Pedometer |
| 0x00522DFF | `34TTrainer_PedometerGoalModel_Import` | Known | Pedometer |
| 0x00523764 | `N3ISL14TStepPedometerE` | Known | Pedometer |
| 0x00523907 | `N3ISL17IPodStepPedometerE` | Known | Pedometer |
| 0x0056C2C8 | `Trainer.PedometerStepGoal` | Known | Pedometer |
| 0x0056C2F6 | `Trainer.Pedometer` | Known | Pedometer |
| 0x0056CC08 | `AggStep` | Known | Pedometer |
| 0x005BCCBC | `TTrainer_Cntlr_Ambient_Pedometer` | Known | Pedometer |
| 0x005BCCE0 | `TTrainer_Cntlr_Ambient_Pedometer` | Known | Pedometer |
| 0x005BDAE4 | `TPedometer_Hourly_Cntlr` | Known | Pedometer |
| 0x005BDAFC | `TTrainer_Cntlr_Pedometer_Goal` | Known | Pedometer |
| 0x00628528 | `Pedometer` | Known | Pedometer |
| 0x0062A774 | `Pedometer` | Known | Pedometer |
| 0x0062A794 | `Please quit your Nike+ workout to begin using the Pedometer.` | Known | Pedometer |
| 0x0062B7EC | `Pedometer` | Known | Pedometer |
| 0x0062B874 | `Step Workout` | Known | Pedometer |
| 0x006315CC | `Stepper Workout` | Known | Pedometer |
| 0x00632090 | `Pedometer` | Known | Pedometer |
| 0x006323D0 | `Steps` | Known | Pedometer |
| 0x006327BC | `Step away from all other sensors` | Known | Pedometer |
| 0x00632908 | `Step away from all other remotes` | Known | Pedometer |
| 0x00632BC4 | `Step away from all other monitors.` | Known | Pedometer |
| 0x00632BFC | `Daily Step View` | Known | Pedometer |
| 0x00632C18 | `Total Steps:` | Known | Pedometer |
| 0x00632C54 | `Steps` | Known | Pedometer |
| 0x00632C68 | `Step Goal` | Known | Pedometer |
| 0x00632CA4 | `Daily Step Goal` | Known | Pedometer |
| 0x00638C98 | `controller.GoToPedometerDailyLandscape` | Known | Pedometer |
| 0x00638D10 | `controller.GoToPedometerDailyPortrait` | Known | Pedometer |
| 0x00638D78 | `controller.GoToPedometerSessionLandscape` | Known | Pedometer |
| 0x00638DEC | `controller.GoToPedometerSessionPortrait` | Known | Pedometer |
| 0x00639404 | `controller.PedometerHilited` | Known | Pedometer |
| 0x00639448 | `controller.PedometerInactiveHilited` | Known | Pedometer |
| 0x0063B488 | `controller.GotoMediumPedometerLayout` | Known | Pedometer |
| 0x00641958 | `controller.GoToPedometerSessionWorkoutSummaryScreen` | Known | Pedometer |
| 0x006419C8 | `controller.GoToWorkoutPedometerAmbientSummaryScreen` | Known | Pedometer |
| 0x00641E90 | `controller.Goto_Pedometer_Daily_Graph` | Known | Pedometer |
| 0x006428C0 | `controller.GoToWorkoutPedometerSummaryScreen` | Known | Pedometer |
| 0x006428F0 | `controller.GotoCustomStepGoal` | Known | Pedometer |

---

## 18. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00053474 | `iPod_Control\Device\SysInfo` | Filesystem Path |  |
| 0x00065F50 | `iPod_Control/Device/Radio/RadioBuffer` | Filesystem Path |  |
| 0x00070324 | `iPod_Control/Device/Accessories/Tags/` | Filesystem Path |  |
| 0x0007C340 | `iPod_Control\Device` | Filesystem Path |  |
| 0x00081708 | `Resources/Sounds/camera.wav` | Filesystem Path |  |
| 0x00083778 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x000A37A0 | `iPod_Control/Device/Radio/Tuner_Scan.log` | Filesystem Path |  |
| 0x000DE208 | `iPod_Control` | Filesystem Path |  |
| 0x000DF564 | `iPod_Control/Logs/crash000.bin` | Filesystem Path |  |
| 0x000DF598 | `pytcgsmlrddamfducpafksthpsafpytegerfktsfglveiPod_Control/Logs` | Filesystem Path |  |
| 0x000DFDFC | `iPod_Control\Device\dst` | Filesystem Path |  |
| 0x001023D0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102554 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001025C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102640 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001027EC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010285C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001028CC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010293C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001029AC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102A1C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102A8C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102AFC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102B6C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102BE4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102C54 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102CCC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102D44 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102DB4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102E24 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102EA4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102F14 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102F7C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00102FF4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010306C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103154 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001031CC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103244 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001032C8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103344 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001033BC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010342C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001034A4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010351C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001035BC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103634 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001036F8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103770 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001037E8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103858 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001038D0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103950 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001039C8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103A40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103AC0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103B40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103BC0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103C38 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103CB0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103D30 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103E58 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103ED8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103F58 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00103FD0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104050 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104120 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001041AC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010422C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001042AC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010431C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010439C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010441C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010448C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010450C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010458C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010460C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104698 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104718 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001047A0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104828 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001048B0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104920 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001049A8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104A30 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104AA8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104B30 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104BA8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104C30 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104CB8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104D30 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104DA8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104E30 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104EA8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104F20 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00104FA0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105028 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001050A8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010513C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001051C4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010523C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001052BC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105334 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001053AC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010543C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001054D4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105564 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001055F4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010566C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001056E4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105780 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0010581C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001058B4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105960 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001059F8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105A9C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105B40 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105BB8 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105C5C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105CF4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105D6C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105E1C | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00105EBC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0011C1E4 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0011C234 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00152044 | `Resources/Sounds/volumebeep.wav` | Filesystem Path |  |
| 0x00170D64 | `iPod_Control/Device/Radio/RadioBuffer` | Filesystem Path |  |
| 0x001AF554 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x001B6E70 | `iPod_Control\Device` | Filesystem Path |  |
| 0x001B7224 | `iPod_Control/Device/Radio` | Filesystem Path |  |
| 0x001B7240 | `iPod_Control/Device/Radio/TunerSettings` | Filesystem Path |  |
| 0x001C0798 | `Resources/Games` | Filesystem Path |  |
| 0x001C07A8 | `iPod_Control` | Filesystem Path |  |
| 0x001C07C8 | `iPod_Control/%s%s%s` | Filesystem Path |  |
| 0x001D2A90 | `Resources/UI/` | Filesystem Path |  |
| 0x001D2AB0 | `Resources/UI/SilverDB.%s.LE.bin` | Filesystem Path |  |
| 0x001D390C | `iPod_Control` | Filesystem Path |  |
| 0x001D7C28 | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path |  |
| 0x001D9260 | `iPod_Control/Device/Users` | Filesystem Path |  |
| 0x001D93C4 | `iPod_Control/Device/PlayCounts` | Filesystem Path |  |
| 0x001DD12C | `/iPod_Control/Device/iPod_Contacts.db` | Filesystem Path |  |
| 0x001E3994 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x001EBF24 | `iPod_Control/Device/Radio/Tuner_Metadata.log` | Filesystem Path |  |
| 0x001EBFC8 | `iPod_Control/Device/Radio/Tuner_Readings.log` | Filesystem Path |  |
| 0x00220C18 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x002328DC | `iPod_Control\Device\` | Filesystem Path |  |
| 0x00232930 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0024B59C | `iPod_Control\Music\` | Filesystem Path |  |
| 0x0024BF80 | `Resources/TrainerTemplates` | Filesystem Path |  |
| 0x002606D0 | `iPod_Control\Device\` | Filesystem Path |  |
| 0x0028089C | `Resources/Sounds/shake.wav` | Filesystem Path |  |
| 0x00298FF4 | `iPod_Control/Device/alarms` | Filesystem Path |  |
| 0x0029CCA4 | `Resources/Sounds/marimba.wav` | Filesystem Path |  |
| 0x0029D170 | `iPod_Control\Device\Preferences` | Filesystem Path |  |
| 0x002F2768 | `iPod_Control/Device/Radio/TunerSettings` | Filesystem Path |  |
| 0x0031CC30 | `iPod_Control/Device/Accessories/Tags/` | Filesystem Path |  |
| 0x00321B04 | `iPod_Control:Device:voices:` | Filesystem Path |  |
| 0x0032FF18 | `Resources/Fonts` | Filesystem Path |  |
| 0x00339AD8 | `Resources/Fonts` | Filesystem Path |  |
| 0x0037D580 | `Resources/Sounds/clicker.wav` | Filesystem Path |  |
| 0x00460402 | `iPod_Control/Device` | Filesystem Path |  |
| 0x00463528 | `Resources/UI/SilverImagesDB.LE.bin` | Filesystem Path |  |
| 0x0047D28A | `iPod_Control/Device` | Filesystem Path |  |
| 0x0056AE19 | `Resources/Games/` | Filesystem Path |  |
| 0x0056C3F7 | `iPod_Control\Device\log` | Filesystem Path |  |
| 0x0056C417 | `/iPod_Control/Speakable` | Filesystem Path |  |
| 0x0056C42F | `/iPod_Control/Speakable/UISS.plist` | Filesystem Path |  |
| 0x0056C452 | `/iPod_Control/Speakable/CacheInfo.plist` | Filesystem Path |  |
| 0x0056C47A | `/iPod_Control/Speakable/ConfigInfo.plist` | Filesystem Path |  |
| 0x0056C4A3 | `/iPod_Control/Speakable/UISS_combined.plist.gz` | Filesystem Path |  |
| 0x0056C4D2 | `/Resources/Speakable/UISS_combined.plist.gz` | Filesystem Path |  |
| 0x0056C4FE | `/Resources/Speakable` | Filesystem Path |  |
| 0x0056C54B | `iPod_Control/Tones` | Filesystem Path |  |
| 0x0056C56B | `/iPod_Control/Device/` | Filesystem Path |  |
| 0x0056C581 | `iPod_Control/Device` | Filesystem Path |  |
| 0x0056C595 | `iPod_Control/Device/accessories.reg` | Filesystem Path |  |

---

## 19. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x004601A8 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Shared/Services/eAppHostLib/eAppHostL` | Build Path |  |
| 0x00460D98 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Shared/Services/eAppHostLib/eAppMotor` | Build Path |  |
| 0x00480E0C | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x00493CF0 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x00493DBC | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x004941A8 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Shared/Services/LIBXML/xpath.c` | Build Path |  |
| 0x004AF1B0 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Silver/3rdParty/freetype/src/base/ftu` | Build Path |  |
| 0x004AF208 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Silver/3rdParty/freetype/src/base/fts` | Build Path |  |
| 0x004AF260 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Silver/3rdParty/freetype/src/base/fto` | Build Path |  |
| 0x004AFE90 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x004B01C8 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x004B0DDC | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x004B14D8 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Shared/Services/hwapi/soc/samsung/nan` | Build Path |  |
| 0x004C2CE8 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Silver/3rdParty/freetype/src/autofit/` | Build Path |  |
| 0x004C3864 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Silver/3rdParty/freetype/src/bdf/bdfd` | Build Path |  |
| 0x004C4A60 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Silver/3rdParty/freetype/src/pfr/pfrg` | Build Path |  |
| 0x004C4AB8 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Silver/3rdParty/freetype/src/pfr/pfrc` | Build Path |  |
| 0x004C4B10 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Silver/3rdParty/freetype/src/pfr/pfro` | Build Path |  |
| 0x004C4E54 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Silver/3rdParty/freetype/src/psaux/t1` | Build Path |  |
| 0x004D41FC | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Silver/3rdParty/freetype/src/sfnt/ttc` | Build Path |  |
| 0x004D4478 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Silver/3rdParty/freetype/src/truetype` | Build Path |  |
| 0x004D49E4 | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Silver/3rdParty/freetype/src/type1/t1` | Build Path |  |
| 0x004D763C | `c:/bwa/N33FirmwareWin-206/srcroot/Firmware/Shared/Services/Image3/Image3.c` | Build Path |  |

---

## 20. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0056D99D | `Electronic` | EQ Preset |  |
| 0x0056D9A8 | `Acoustic` | EQ Preset |  |
| 0x0056D9C5 | `Dance` | EQ Preset |  |
| 0x0056D9CB | `Lounge` | EQ Preset |  |
| 0x0056D9D2 | `Rock` | EQ Preset |  |
| 0x0056D9D7 | `Classical` | EQ Preset |  |
| 0x0056D9E1 | `Latin` | EQ Preset |  |
| 0x0056D9E7 | `Piano` | EQ Preset |  |
| 0x0056DA48 | `Loudness` | EQ Preset |  |
| 0x0056DA56 | `Jazz` | EQ Preset |  |
| 0x0062CB24 | `Acoustic` | EQ Preset |  |
| 0x0062CB30 | `Bass Booster` | EQ Preset |  |
| 0x0062CB50 | `Classical` | EQ Preset |  |
| 0x0062CB5C | `Dance` | EQ Preset |  |
| 0x0062CB6C | `Electronic` | EQ Preset |  |
| 0x0062CB80 | `Hip Hop` | EQ Preset |  |
| 0x0062CB88 | `Jazz` | EQ Preset |  |
| 0x0062CB90 | `Latin` | EQ Preset |  |
| 0x0062CB98 | `Loudness` | EQ Preset |  |
| 0x0062CBA4 | `Lounge` | EQ Preset |  |
| 0x0062CBAC | `Piano` | EQ Preset |  |
| 0x0062CBBC | `Rock` | EQ Preset |  |
| 0x0062CBC4 | `Small Speakers` | EQ Preset |  |
| 0x0062CBD4 | `Spoken Word` | EQ Preset |  |
| 0x0062CBE0 | `Treble Booster` | EQ Preset |  |
| 0x0062CC00 | `Vocal Booster` | EQ Preset |  |

---
