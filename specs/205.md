# iPod Classic 7G (Rev C) - RetailOS 2.0.5 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 2.0.5 |
| **IPSW** | iPod_38.2.0.5.ipsw |
| **Device** | iPod Classic 7G (Rev C) (2012, Click Wheel, Cover Flow, Genius, EU Volume) |
| **UpdaterFamilyID** | 38 |
| **Binary Size** | 10,634,528 bytes (10.14 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,632,480 bytes |
| **Total Strings (>=6)** | 55,712 |
| **Function Prologues** | 23,164 (ARM: 17,762, Thumb: 5,402) |
| **SoC** | Samsung S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Encrypted** | Yes (HW AES) |
| **SHA-256** | `b467cb5151ef9b45fc830466aca106322fd19e795c404e5bfa14ade632933a7a` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001449A4 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x0015B6D4 | `MockupMode/MockupMode.xml` | Hidden | Developer Tool |
| 0x0015BBD4 | `MockupMode/` | Hidden | Developer Tool |
| 0x00189C74 | `TCDemoMode` | Hidden | Demo/Retail Mode |
| 0x001E13BC | `TSilverCntlrTestAppCntlr` | Hidden | Developer Tool |
| 0x00266BD8 | `Channel UnitTests` | Hidden | Developer Tool |
| 0x002C13DD | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x002C1420 | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x002C1435 | `RTXCbug> ` | Hidden | Developer Tool |
| 0x002C1E11 | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x002DB6F0 | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x003960E5 | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x003961AD | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x003F4E0D | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x003F8860 | `TCDemoMode` | Hidden | Demo/Retail Mode |
| 0x00403E0C | `TSilverCntlrTestAppCntlr` | Hidden | Developer Tool |
| 0x0073B49E | `Debug_MainMenu_Screen` | Hidden | Debug/Diagnostic |
| 0x0073B4B7 | `Debug_MainMenu_Screen_Default"` | Hidden | Debug/Diagnostic |
| 0x0073B525 | `Extras_Screen_Debug` | Hidden | Debug/Diagnostic |
| 0x0075B0C8 | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDe` | Hidden | Demo/Retail Mode |
| 0x0075BA0C | `TSilverCntlrTUnitTestSuiteCntlr` | Hidden | Developer Tool |
| 0x0075BA2C | `TUnitTestSuiteTestsCntlrTSilverCntlrTCVoiceMemosMainMenuTCVoiceMemosMenuTCVoiceM` | Hidden | Developer Tool |
| 0x007E81B6 | `Debug_UnitTest_Screen` | Hidden | Developer Tool |
| 0x007E81CF | `Debug_UnitTest_Screen_Default` | Hidden | Developer Tool |
| 0x007E8232 | `DemoMode_Screen` | Hidden | Demo/Retail Mode |
| 0x007E8245 | `DemoMode_Screen_Default` | Hidden | Demo/Retail Mode |
| 0x007E82B2 | `Debug_TestList_Screen` | Hidden | Debug/Diagnostic |
| 0x007E82CB | `Debug_TestList_Screen_Default` | Hidden | Debug/Diagnostic |
| 0x007E833E | `Debug_TestResult_Screen` | Hidden | Debug/Diagnostic |
| 0x007E8359 | `Debug_TestResult_Screen_Default` | Hidden | Debug/Diagnostic |
| 0x00805CD0 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00844E40 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00858050 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00870660 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008834C0 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0088D730 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00897684 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008AD38C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008B7568 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008DF12C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008FE89C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00908150 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00911EBC | `WaveFileDebugTask` | Hidden | Debug/Diagnostic |
| 0x00913040 | `TCMockupModeNavScreen` | Hidden | Developer Tool |
| 0x009900A1 | `10TCDemoMode` | Hidden | Demo/Retail Mode |
| 0x00990A18 | `21TCMockupModeNavScreen` | Hidden | Developer Tool |
| 0x00990C7E | `24TSilverCntlrTestAppCntlr` | Hidden | Developer Tool |
| 0x00990ED8 | `27TSilverCntlrTransitionAddonI10TCDemoModeE` | Hidden | Demo/Retail Mode |
| 0x00991B9F | `27TSilverCntlrTransitionAddonI24TSilverCntlrTestAppCntlrE` | Hidden | Developer Tool |
| 0x009BC360 | `Returning from RTXCBug` | Hidden | Developer Tool |
| 0x009BFCCE | `Debug_ListItem_DemoMode` | Hidden | Debug/Diagnostic |
| 0x009BFCE6 | `Debug_MenuItem_DemoMode` | Hidden | Debug/Diagnostic |
| 0x009C03EB | `DemoMode_Background_Image` | Hidden | Demo/Retail Mode |
| 0x009C1077 | `DemoMode_Icon_Image` | Hidden | Demo/Retail Mode |
| 0x009C2C85 | `Debug_Menu_Title` | Hidden | Debug/Diagnostic |
| 0x009C2CAA | `UnitTest_Menu_Title` | Hidden | Developer Tool |
| 0x009CADD2 | `Extras_Screen_Debug` | Hidden | Debug/Diagnostic |
| 0x009CADE6 | `MainMenu_List_Debug` | Hidden | Debug/Diagnostic |
| 0x009CADFA | `ExtrasMenu_Debug` | Hidden | Debug/Diagnostic |
| 0x009CBB69 | `UnitTestModel` | Hidden | Developer Tool |
| 0x009CC548 | `UnitTest_ListItem` | Hidden | Developer Tool |
| 0x009CC88B | `DemoMode_Screen` | Hidden | Demo/Retail Mode |
| 0x009CCAF4 | `DemoMode_Main_Screen` | Hidden | Demo/Retail Mode |
| 0x009CD265 | `Debug_TestResult_Screen` | Hidden | Debug/Diagnostic |
| 0x009CD2B5 | `Debug_UnitTest_Screen` | Hidden | Developer Tool |
| 0x009CD2DD | `Debug_TestList_Screen` | Hidden | Debug/Diagnostic |
| 0x009CD445 | `Debug_MainMenu_Screen` | Hidden | Debug/Diagnostic |
| 0x009CD71F | `DemoMode_View_Main` | Hidden | Demo/Retail Mode |
| 0x009CD91C | `DemoMode_Icon` | Hidden | Demo/Retail Mode |
| 0x009CE70F | `SilverTestApp` | Hidden | Developer Tool |
| 0x009CE71D | `UnitTestApp` | Hidden | Developer Tool |
| 0x009CECCF | `Debug_ListItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009CECEA | `Debug_MenuItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009CF446 | `DemoMode_Text_Color` | Hidden | Demo/Retail Mode |
| 0x009CF85B | `Debug_ListItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009CF872 | `Debug_MenuItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009D25C3 | `DemoMode_Screen_Default` | Hidden | Demo/Retail Mode |
| 0x009D2D1F | `Debug_TestResult_Screen_Default` | Hidden | Debug/Diagnostic |
| 0x009D2D76 | `Debug_UnitTest_Screen_Default` | Hidden | Developer Tool |
| 0x009D2DAE | `Debug_TestList_Screen_Default` | Hidden | Debug/Diagnostic |
| 0x009D2EE5 | `Debug_MainMenu_Screen_Default` | Hidden | Debug/Diagnostic |
| 0x009D3A9A | `Debug_ListItem_UnitTest` | Hidden | Developer Tool |
| 0x009D3AB2 | `Debug_MenuItem_UnitTest` | Hidden | Developer Tool |
| 0x009D834E | `Debug_ListItem_Memory` | Hidden | Debug/Diagnostic |
| 0x009D8364 | `Debug_MenuItem_Memory` | Hidden | Debug/Diagnostic |
| 0x00A23AF0 | `DebugUtil` | Hidden | Debug/Diagnostic |

---

## 2. Controllers (TSilver/TC Classes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000AC1D8 | `TSilverCntlr` | Known | UI controller |
| 0x000AC1F0 | `TCExtrasMenu` | Known | UI controller |
| 0x000AC208 | `TCGameScreen` | Known | UI controller |
| 0x000AC220 | `TCGamesMenu` | Known | UI controller |
| 0x000AC234 | `TSilverMainMediaListCntlr_Main` | Known | UI controller |
| 0x000AC25C | `TSilverMainMediaListCntlr_Music` | Known | UI controller |
| 0x000AC284 | `TSilverMainMediaListCntlr_Videos` | Known | UI controller |
| 0x000AC2B0 | `TSilverMediaListCntlr_Songs` | Known | UI controller |
| 0x000AC2D4 | `TSilverMediaListCntlr_Albums` | Known | UI controller |
| 0x000AC2FC | `TSilverMediaListCntlr_Artists` | Known | UI controller |
| 0x000AC324 | `TSilverMediaListCntlr_Genres` | Known | UI controller |
| 0x000AC34C | `TSilverMediaListCntlr_Composers` | Known | UI controller |
| 0x000AC374 | `TSilverMediaListCntlr_Podcasts` | Known | UI controller |
| 0x000AC39C | `TSilverMediaListCntlr_PodcastEpisodes` | Known | UI controller |
| 0x000AC3CC | `TSilverMediaListCntlr_iTunesU` | Known | UI controller |
| 0x000AC3F4 | `TSilverMediaListCntlr_iTunesUEpisodes` | Known | UI controller |
| 0x000AC424 | `TSilverMediaListCntlr_Audiobooks` | Known | UI controller |
| 0x000AC450 | `TSilverMediaListCntlr_AudiobookChapters` | Known | UI controller |
| 0x000AC480 | `TSilverMediaListCntlr_TVShows` | Known | UI controller |
| 0x000AC4A8 | `TSilverMediaListCntlr_TVSeasons` | Known | UI controller |
| 0x000AC4D0 | `TSilverMediaListCntlr_TVEpisodes` | Known | UI controller |
| 0x000AC4FC | `TSilverMediaListCntlr_Movies` | Known | UI controller |
| 0x000AC524 | `TSilverMediaListCntlr_Playlists` | Known | UI controller |
| 0x000AC54C | `TSilverMediaListCntlr_NestedPlaylists` | Known | UI controller |
| 0x000AC57C | `TSilverMediaListCntlr_VideoPlaylists` | Known | UI controller |
| 0x000AC718 | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | UI controller |
| 0x000AC74C | `TSilverMediaListCntlr_PlaylistChooser` | Known | UI controller |
| 0x000AC77C | `TSilverMediaListCntlr_Rentals` | Known | UI controller |
| 0x000AC7A4 | `TSilverMediaListCntlr_Genius` | Known | UI controller |
| 0x000AC7CC | `TSilverMediaListCntlr_GeniusMixes` | Known | UI controller |
| 0x000AC7F8 | `TCRentalNotification` | Known | UI controller |
| 0x000AC818 | `TCRentalInfo` | Known | UI controller |
| 0x000AC830 | `TCRentalConfirmDelete` | Known | UI controller |
| 0x000AC850 | `TCRentalDispatcher` | Known | UI controller |
| 0x000AC8A8 | `TSilverGlobalCntlr` | Known | UI controller |
| 0x000AC8C4 | `TSilverTrainerCntlr` | Known | UI controller |
| 0x001041A4 | `TCSlideshowLCD` | Known | UI controller |
| 0x001041BC | `TCSlideshowTVOut` | Known | UI controller |
| 0x001041D8 | `TCSlideshow_TVOutAsk` | Known | UI controller |
| 0x001041F8 | `TCSlideshow_TVOutCableConnect` | Known | UI controller |
| 0x001280B8 | `TSilverCalendarCntlr_CalendarMenu` | Known | UI controller |
| 0x001280E4 | `TSilverCalendarCntlr_MonthViewer` | Known | UI controller |
| 0x00128110 | `TSilverCalendarCntlr_DayViewer` | Known | UI controller |
| 0x00128138 | `TSilverCalendarCntlr_EventViewer` | Known | UI controller |
| 0x00128164 | `TSilverCalendarCntlr_ToDoList` | Known | UI controller |
| 0x0012818C | `TSilverCalendarCntlr_ToDoDispatcher` | Known | UI controller |
| 0x001281B8 | `TSilverCalendarCntlr_Alarm` | Known | UI controller |
| 0x0012FB38 | `TCRemoteUI` | Known | UI controller |
| 0x0012FB4C | `TCUnsupported` | Known | UI controller |
| 0x00136350 | `TCSpeakers` | Known | UI controller |
| 0x00136364 | `TCEQSetting` | Known | UI controller |
| 0x0015F85C | `TCSportTimer` | Known | UI controller |
| 0x0015F874 | `TCSportTimerMenu` | Known | UI controller |
| 0x0015F890 | `TCSportTimerSessionScreen` | Known | UI controller |
| 0x0015F8B4 | `TCSportTimerChosenDispatcher` | Known | UI controller |
| 0x00160C64 | `TCVoiceMemos` | Known | UI controller |
| 0x00160C7C | `TCVoiceMemosMenu` | Known | UI controller |
| 0x00160C98 | `TCVoiceMemosMainMenu` | Known | UI controller |
| 0x00160CB8 | `TCVoiceMemosPlayback` | Known | UI controller |
| 0x00160CD8 | `TCVoiceMemosContextMenu` | Known | UI controller |
| 0x00160CF8 | `TCVoiceMemosAlert` | Known | UI controller |
| 0x00172E70 | `TSilverSettingsMenuListCntlr` | Known | UI controller |
| 0x00172E98 | `TCSettings_MainMenu` | Known | UI controller |
| 0x00172EB4 | `TCSettings_MusicMenu` | Known | UI controller |
| 0x00172ED4 | `TCSettings_VolumeLimit` | Known | UI controller |
| 0x00172EF4 | `TCSettings_Brightness` | Known | UI controller |
| 0x00172F14 | `TCSettings_BacklightTimer` | Known | UI controller |
| 0x00172F38 | `TCSettings_EQ` | Known | UI controller |
| 0x00172F50 | `TCSettings_AudiobookSettings` | Known | UI controller |
| 0x00172F78 | `TCSettings_RadioRegions` | Known | UI controller |
| 0x00172F98 | `TCSettings_ResetAllSettings` | Known | UI controller |
| 0x00172FBC | `TCSettings_EULimitConfirmation` | Known | UI controller |
| 0x00172FE4 | `TSilverSettingsVideoCntlr` | Known | UI controller |
| 0x00173008 | `TCDateTimeScreen` | Known | UI controller |
| 0x00173024 | `TCTimeZoneScreen` | Known | UI controller |
| 0x00173040 | `TCSettings_AdjustScrollingCntlr` | Known | UI controller |
| 0x00173068 | `TCFirstBoot` | Known | UI controller |
| 0x001B2E80 | `TCAddressViewerMainMenu` | Known | UI controller |
| 0x001B2EA0 | `TCAddressViewerDetails` | Known | UI controller |
| 0x001B2EC0 | `TCAddressViewerPartialLoad` | Known | UI controller |
| 0x001B2EE4 | `TCAddressViewerMainDispatcher` | Known | UI controller |
| 0x001E13E0 | `TSilverCntlrTestCntlr` | Known | UI controller |
| 0x001E8C24 | `TSilverMainMediaListCntlr_Videos` | Known | UI controller |
| 0x00281590 | `TC_LockDialog` | Known | UI controller |
| 0x002815A8 | `TC_LockScreen` | Known | UI controller |
| 0x002815C0 | `TC_LockediPod` | Known | UI controller |
| 0x002815D8 | `TC_VolumeLimitLockScreen` | Known | UI controller |
| 0x002815FC | `TCLockChosenDispatcher` | Known | UI controller |
| 0x002871BC | `TCClock` | Known | UI controller |
| 0x002871CC | `TCClockCityMenu` | Known | UI controller |
| 0x002871E4 | `TCClockRegionMenu` | Known | UI controller |
| 0x00287200 | `TCAlarmMenu` | Known | UI controller |
| 0x00287214 | `TCSleepTimerMenu` | Known | UI controller |
| 0x00287230 | `TCAlarmPropertiesMenu` | Known | UI controller |
| 0x00287250 | `TCAlarmPropertiesFrequencyMenu` | Known | UI controller |
| 0x00287278 | `TCAlarmPropertiesLabelMenu` | Known | UI controller |
| 0x0028729C | `TCAlarmPropertiesSoundMenu` | Known | UI controller |
| 0x002872C0 | `TCAlarmDatePicker` | Known | UI controller |
| 0x002872DC | `TCAlarmTriggered` | Known | UI controller |
| 0x0028E284 | `TCNotesDispatcher` | Known | UI controller |
| | *...and 405 more* | | |

---

## 3. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000FF928 | `HandleWheel` | Known | Event handler |
| 0x000FF934 | `HandlePlayPause` | Known | Event handler |
| 0x000FF944 | `HandleSelectDown` | Known | Event handler |
| 0x000FF958 | `HandleNext` | Known | Event handler |
| 0x000FF964 | `HandlePrevious` | Known | Event handler |
| 0x000FF974 | `HandleNextPushAndHold` | Known | Event handler |
| 0x000FF98C | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000FFC24 | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000FFC44 | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x0010C06C | `HandleSelect` | Known | Event handler |
| 0x0010C080 | `HandleHilite` | Known | Event handler |
| 0x0010C418 | `HandleEQSettingSelected` | Known | Event handler |
| 0x0010C848 | `HandleSelect` | Known | Event handler |
| 0x0010C85C | `HandleGameHilited` | Known | Event handler |
| 0x0010CB0C | `HandleNotesSelected` | Known | Event handler |
| 0x0010CB24 | `HandleNotesPop` | Known | Event handler |
| 0x0010CB34 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0011B270 | `HandleVolumeWheel` | Known | Event handler |
| 0x0011B284 | `HandleVolumeChange` | Known | Event handler |
| 0x0011B298 | `HandleTimerDone` | Known | Event handler |
| 0x0011B2A8 | `HandleFrequencyChange` | Known | Event handler |
| 0x0011B320 | `HandleTuning` | Known | Event handler |
| 0x0011B330 | `HandleTuningSelect` | Known | Event handler |
| 0x00125E74 | `HandleLock` | Known | Event handler |
| 0x00125E84 | `HandleAddressBook` | Known | Event handler |
| 0x0012656C | `HandleSelect` | Known | Event handler |
| 0x00126AA4 | `HandleExit` | Known | Event handler |
| 0x00126AB4 | `HandleLap` | Known | Event handler |
| 0x00126AC0 | `HandleResume` | Known | Event handler |
| 0x00126AD0 | `HandleStartStop` | Known | Event handler |
| 0x00126D84 | `HandleWheel` | Known | Event handler |
| 0x00126D94 | `HandlePlayPause` | Known | Event handler |
| 0x00126DA4 | `HandleSelectDown` | Known | Event handler |
| 0x00126DB8 | `HandleHilite` | Known | Event handler |
| 0x00126DDC | `HandleFinishRecording` | Known | Event handler |
| 0x0013135C | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x0014001C | `HandleExitUnsupported` | Known | Event handler |
| 0x00156F20 | `HandleNotesPop` | Known | Event handler |
| 0x00156F34 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00157E40 | `HandleSelect` | Known | Event handler |
| 0x00157E54 | `HandleWheel` | Known | Event handler |
| 0x00157E60 | `HandleImageNext` | Known | Event handler |
| 0x00157E70 | `HandleImagePrev` | Known | Event handler |
| 0x00157E80 | `HandleImageLast` | Known | Event handler |
| 0x00157E90 | `HandleImageFirst` | Known | Event handler |
| 0x00157EA4 | `HandlePlayPause` | Known | Event handler |
| 0x00157EB4 | `HandlePlay` | Known | Event handler |
| 0x00157EC0 | `HandlePause` | Known | Event handler |
| 0x00157ECC | `HandleMikeyCenter` | Known | Event handler |
| 0x0016D160 | `HandleSelectCity` | Known | Event handler |
| 0x0016D178 | `HandleHighlightCity` | Known | Event handler |
| 0x0016E264 | `HandleWantPopFlow` | Known | Event handler |
| 0x0016E27C | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x0016E298 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0016E2B4 | `HandleFlowNext` | Known | Event handler |
| 0x0016E2C4 | `HandleFlowPrev` | Known | Event handler |
| 0x0016E2D4 | `HandleFlowWheel` | Known | Event handler |
| 0x0016E2E4 | `HandleAlbumSelected` | Known | Event handler |
| 0x0016E2F8 | `HandlePlayPause` | Known | Event handler |
| 0x0016E308 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x0019A3FC | `HandleLeaveAlarm` | Known | Event handler |
| 0x0019A7EC | `HandleSelect` | Known | Event handler |
| 0x0019B6D4 | `HandleSelect` | Known | Event handler |
| 0x0019B6E8 | `HandleWheel` | Known | Event handler |
| 0x0019B6F4 | `HandleImageNext` | Known | Event handler |
| 0x0019B704 | `HandleImagePrev` | Known | Event handler |
| 0x0019B714 | `HandleImageLast` | Known | Event handler |
| 0x0019B724 | `HandleImageFirst` | Known | Event handler |
| 0x0019B738 | `HandlePlayPause` | Known | Event handler |
| 0x0019B748 | `HandlePlay` | Known | Event handler |
| 0x0019B754 | `HandlePause` | Known | Event handler |
| 0x0019B760 | `HandleMikeyCenter` | Known | Event handler |
| 0x0019BC08 | `HandleNew` | Known | Event handler |
| 0x0019BC18 | `HandleClear` | Known | Event handler |
| 0x0019BC24 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x0019BC40 | `HandleSelectIndexedSession` | Known | Event handler |
| 0x0019BF50 | `HandleWheel` | Known | Event handler |
| 0x0019BF60 | `HandleArrowUp` | Known | Event handler |
| 0x0019BF70 | `HandleArrowDown` | Known | Event handler |
| 0x0019EBFC | `HandleHiliteAlbum` | Known | Event handler |
| 0x0019EC14 | `HandleBrowseAlbum` | Known | Event handler |
| 0x0019EC28 | `HandlePlayPause` | Known | Event handler |
| 0x001B56DC | `HandleSelect` | Known | Event handler |
| 0x001B586C | `HandleSelectRegion` | Known | Event handler |
| 0x001B5BE4 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x001B5C00 | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x001B5C1C | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x001CC9C4 | `HandleImageWheel` | Known | Event handler |
| 0x001CC9DC | `HandlePlayPause` | Known | Event handler |
| 0x001CC9EC | `HandleBrowseLarge` | Known | Event handler |
| 0x001CCA00 | `HandleBrowseSmall` | Known | Event handler |
| 0x001CCA14 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x001CCA2C | `HandleImageNext` | Known | Event handler |
| 0x001CCA3C | `HandleImagePrev` | Known | Event handler |
| 0x001CCA4C | `HandleHilite` | Known | Event handler |
| 0x001CCA5C | `HandleImageLast` | Known | Event handler |
| 0x001CCA6C | `HandleImageFirst` | Known | Event handler |
| 0x001CCA80 | `HandleScreenNext` | Known | Event handler |
| 0x001CCA94 | `HandleScreenPrev` | Known | Event handler |
| 0x001CF35C | `HandlePlayPause` | Known | Event handler |
| | *...and 1846 more* | | |

---

## 4. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00125DB4 | `GotoScreen_LockDialog` | Known | Navigation handler |
| 0x00125DCC | `GotoScreen_SetCombinationFirstTime` | Known | Navigation handler |
| 0x00125F44 | `GotoScreen_AddressBook` | Known | Navigation handler |
| 0x001403D4 | `GotoScreen_EnterPassKey` | Known | Navigation handler |
| 0x001403EC | `GotoScreen_LockediPod` | Known | Navigation handler |
| 0x00140DF0 | `GotoScreen_MainMenu` | Known | Navigation handler |
| 0x001DDA00 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation handler |
| 0x001E8EC8 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation handler |
| 0x0020483C | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation handler |
| 0x00210260 | `GotoScreen_VolumeLimit` | Known | Navigation handler |
| 0x00210364 | `GotoScreen_SettingsMenu` | Known | Navigation handler |
| 0x0021037C | `GotoScreen_SettingsMenuEU` | Known | Navigation handler |
| 0x0021E034 | `GotoScreen_Language` | Known | Navigation handler |
| 0x0021F748 | `GotoGeniusLoadingScreen` | Known | Navigation handler |
| 0x0022498C | `GotoGeniusLoadingScreen` | Known | Navigation handler |
| 0x00227E60 | `GotoGeniusLoadingScreen` | Known | Navigation handler |
| 0x00228044 | `GotoGeniusLoadingScreen` | Known | Navigation handler |
| 0x002299B4 | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation handler |
| 0x002299D4 | `GotoVolumeLimit_or_Lock_or_EU_Screen` | Known | Navigation handler |
| 0x00229D00 | `GotoScreen_BacklightTimer` | Known | Navigation handler |
| 0x00229D90 | `GotoScreen_VolumeLimit` | Known | Navigation handler |
| 0x00229DA8 | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation handler |
| 0x00229E54 | `GotoScreen_VolumeLimitEU` | Known | Navigation handler |
| 0x00229E70 | `GotoScreen_VolumeLimit` | Known | Navigation handler |
| 0x00229E88 | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation handler |
| 0x00229F18 | `GotoScreen_EUVolumeLimitConfirmation` | Known | Navigation handler |
| 0x0022F214 | `GotoScreen_LockDialog` | Known | Navigation handler |
| 0x0022F22C | `GotoScreen_SetCombinationFirstTime` | Known | Navigation handler |
| 0x00237680 | `GotoGeniusMixLoadingScreen` | Known | Navigation handler |
| 0x0023A7FC | `GotoGeniusLoadingScreen` | Known | Navigation handler |
| 0x0023A9D4 | `GotoGeniusLoadingScreen` | Known | Navigation handler |
| 0x00746F97 | `GotoVolumeLimit_or_Lock_or_EU_Screen` | Known | Navigation handler |
| 0x0074818E | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation handler |

---

## 5. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016ED34 | `CoverFlow_Screen` | Known | Screen layout |
| 0x0073AE46 | `Clock_Screen` | Known | Screen layout |
| 0x0073AE56 | `Clock_Screen_Default"` | Known | Screen layout |
| 0x0073AEBB | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x0073AF19 | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x0073AF31 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x0073AF9E | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x0073B03C | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x0073B09B | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x0073B0B1 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x0073B11C | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x0073B176 | `Games_Menu_Screen` | Known | Screen layout |
| 0x0073B18B | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x0073B1F5 | `Extras_Screen_Games` | Known | Screen layout |
| 0x0073B2B4 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x0073B378 | `Extras_Screen_Lock` | Known | Screen layout |
| 0x0073B441 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x0073B664 | `Clock_Africa_City_Screen ` | Known | Screen layout |
| 0x0073B680 | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x0073B704 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x0073B71E | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x0073B7A0 | `Clock_Atlantic_City_Screen"` | Known | Screen layout |
| 0x0073B7BE | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x0073B844 | `Clock_Australia_City_Screen#` | Known | Screen layout |
| 0x0073B863 | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x0073B8EA | `Clock_Europe_City_Screen ` | Known | Screen layout |
| 0x0073B906 | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x0073B98A | `Clock_NorthAmerica_City_Screen&` | Known | Screen layout |
| 0x0073B9AC | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0073BA36 | `Clock_Pacific_City_Screen!` | Known | Screen layout |
| 0x0073BA53 | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x0073BAD8 | `Clock_SouthAmerica_City_Screen&` | Known | Screen layout |
| 0x0073BAFA | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x0073BB87 | `Clock_Screen"` | Known | Screen layout |
| 0x0073BC2C | `Clock_Screen"` | Known | Screen layout |
| 0x0073BCD1 | `Clock_Screen"` | Known | Screen layout |
| 0x0073BD76 | `Clock_Screen"` | Known | Screen layout |
| 0x0073BE1B | `Clock_Screen"` | Known | Screen layout |
| 0x0073BEC0 | `Clock_Screen"` | Known | Screen layout |
| 0x0073BF65 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C00A | `Clock_Screen"` | Known | Screen layout |
| 0x0073C0AF | `Clock_Screen"` | Known | Screen layout |
| 0x0073C154 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C1F9 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C29E | `Clock_Screen"` | Known | Screen layout |
| 0x0073C343 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C3E8 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C48D | `Clock_Screen"` | Known | Screen layout |
| 0x0073C532 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C5D7 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C67C | `Clock_Screen"` | Known | Screen layout |
| 0x0073C721 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C7C6 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C86B | `Clock_Screen"` | Known | Screen layout |
| 0x0073C910 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C9B5 | `Clock_Screen"` | Known | Screen layout |
| 0x0073CA5A | `Clock_Screen"` | Known | Screen layout |
| 0x0073CAFF | `Clock_Screen"` | Known | Screen layout |
| 0x0073CBA4 | `Clock_Screen"` | Known | Screen layout |
| 0x0073CC49 | `Clock_Screen"` | Known | Screen layout |
| 0x0073CCEE | `Clock_Screen"` | Known | Screen layout |
| 0x0073CD93 | `Clock_Screen"` | Known | Screen layout |
| 0x0073CE38 | `Clock_Screen"` | Known | Screen layout |
| 0x0073CEDD | `Clock_Screen"` | Known | Screen layout |
| 0x0073CF87 | `Clock_Screen"` | Known | Screen layout |
| 0x0073D02C | `Clock_Screen"` | Known | Screen layout |
| 0x0073D0D1 | `Clock_Screen"` | Known | Screen layout |
| 0x0073D176 | `Clock_Screen"` | Known | Screen layout |
| 0x0073D21B | `Clock_Screen"` | Known | Screen layout |
| 0x0073D2C0 | `Clock_Screen"` | Known | Screen layout |
| 0x0073D365 | `Clock_Screen"` | Known | Screen layout |
| 0x0073D40A | `Clock_Screen"` | Known | Screen layout |
| 0x0073D4AF | `Clock_Screen"` | Known | Screen layout |
| 0x0073D554 | `Clock_Screen"` | Known | Screen layout |
| 0x0073D5F9 | `Clock_Screen"` | Known | Screen layout |
| 0x0073D69E | `Clock_Screen"` | Known | Screen layout |
| 0x0073D743 | `Clock_Screen"` | Known | Screen layout |
| 0x0073D7E8 | `Clock_Screen"` | Known | Screen layout |
| 0x0073D88D | `Clock_Screen"` | Known | Screen layout |
| 0x0073D932 | `Clock_Screen"` | Known | Screen layout |
| 0x0073D9D7 | `Clock_Screen"` | Known | Screen layout |
| 0x0073DA7C | `Clock_Screen"` | Known | Screen layout |
| 0x0073DB21 | `Clock_Screen"` | Known | Screen layout |
| 0x0073DBC6 | `Clock_Screen"` | Known | Screen layout |
| 0x0073DC6B | `Clock_Screen"` | Known | Screen layout |
| 0x0073DD10 | `Clock_Screen"` | Known | Screen layout |
| 0x0073DDB5 | `Clock_Screen"` | Known | Screen layout |
| 0x0073DE5A | `Clock_Screen"` | Known | Screen layout |
| 0x0073DEFF | `Clock_Screen"` | Known | Screen layout |
| 0x0073DFA4 | `Clock_Screen"` | Known | Screen layout |
| 0x0073E049 | `Clock_Screen"` | Known | Screen layout |
| 0x0073E0EE | `Clock_Screen"` | Known | Screen layout |
| 0x0073E193 | `Clock_Screen"` | Known | Screen layout |
| 0x0073E238 | `Clock_Screen"` | Known | Screen layout |
| 0x0073E2DD | `Clock_Screen"` | Known | Screen layout |
| 0x0073E382 | `Clock_Screen"` | Known | Screen layout |
| 0x0073E427 | `Clock_Screen"` | Known | Screen layout |
| 0x0073E4CC | `Clock_Screen"` | Known | Screen layout |
| 0x0073E571 | `Clock_Screen"` | Known | Screen layout |
| 0x0073E616 | `Clock_Screen"` | Known | Screen layout |
| | *...and 6691 more* | | |

---

## 6. Settings (Toggle/Show)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0013FDE8 | `ShowSetting_EQ` | Known | User setting |
| 0x001EABD4 | `ToggleSetting_Repeat` | Known | User setting |
| 0x001EABF0 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001EAC08 | `ToggleSetting_TVOut` | Known | User setting |
| 0x001EAC1C | `ToggleSetting_TVSignal` | Known | User setting |
| 0x002140EC | `ShowSetting_Backlight` | Known | User setting |
| 0x00229414 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x00229430 | `ToggleSetting_Repeat` | Known | User setting |
| 0x00229448 | `ToggleSetting_SortBy` | Known | User setting |
| 0x00229460 | `ToggleSetting_ClassicUI` | Known | User setting |
| 0x00229478 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x00229494 | `ToggleSetting_Clicker` | Known | User setting |
| 0x002294AC | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x002294CC | `ToggleSetting_24HourClock` | Known | User setting |
| 0x002294E8 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x00229504 | `ShowSetting_Shuffle` | Known | User setting |
| 0x002296B0 | `ShowSetting_Repeat` | Known | User setting |
| 0x002296C4 | `ShowSetting_About` | Known | User setting |
| 0x002296D8 | `ShowSetting_MainMenu` | Known | User setting |
| 0x002296F0 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x00229708 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x00229720 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x0022973C | `ShowSetting_Brightness` | Known | User setting |
| 0x00229754 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x0022976C | `ShowSetting_RadioRegions` | Known | User setting |
| 0x00229788 | `ShowSetting_EQ` | Known | User setting |
| 0x00229798 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x00229918 | `ShowSetting_Clicker` | Known | User setting |
| 0x0022992C | `ShowSetting_DateAndTime` | Known | User setting |
| 0x00229944 | `ShowSetting_SortBy` | Known | User setting |
| 0x00229958 | `ShowSetting_ClassicUI` | Known | User setting |
| 0x00229970 | `ShowSetting_Language` | Known | User setting |
| 0x00229988 | `ShowSetting_Legal` | Known | User setting |
| 0x0022999C | `ShowSetting_ResetAll` | Known | User setting |
| 0x002299FC | `ToggleSetting_RecommendedVolumeLimit` | Known | User setting |
| 0x007442C9 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x00744379 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x00746B62 | `ShowSetting_About` | Known | User setting |
| 0x00746C04 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x00746C48 | `ShowSetting_Shuffle` | Known | User setting |
| 0x00746CBF | `ToggleSetting_Repeat` | Known | User setting |
| 0x00746D02 | `ShowSetting_Repeat` | Known | User setting |
| 0x00746E0C | `ShowSetting_MainMenu` | Known | User setting |
| 0x00746F1C | `ShowSetting_MusicMenu` | Known | User setting |
| 0x00746FEA | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x007470B4 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x007471CC | `ShowSetting_Brightness` | Known | User setting |
| 0x00747302 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x00747413 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x00747514 | `ShowSetting_EQ` | Known | User setting |
| 0x00747581 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x007475C8 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x00747645 | `ToggleSetting_Clicker` | Known | User setting |
| 0x00747689 | `ShowSetting_Clicker` | Known | User setting |
| 0x007477F0 | `ToggleSetting_SortBy` | Known | User setting |
| 0x00747833 | `ShowSetting_SortBy` | Known | User setting |
| 0x00747934 | `ShowSetting_Language` | Known | User setting |
| 0x00747A44 | `ShowSetting_Legal` | Known | User setting |
| 0x00747B75 | `ShowSetting_ResetAll` | Known | User setting |
| 0x00747CE8 | `ShowSetting_Backlight` | Known | User setting |
| 0x00747D98 | `ShowSetting_Backlight` | Known | User setting |
| 0x00747E48 | `ShowSetting_Backlight` | Known | User setting |
| 0x00747EF9 | `ShowSetting_Backlight` | Known | User setting |
| 0x00747FAA | `ShowSetting_Backlight` | Known | User setting |
| 0x0074805B | `ShowSetting_Backlight` | Known | User setting |
| 0x0074810F | `ShowSetting_Backlight` | Known | User setting |
| 0x007481D2 | `ToggleSetting_RecommendedVolumeLimit` | Known | User setting |
| 0x00748256 | `ShowSetting_EQ` | Known | User setting |
| 0x007482CB | `ShowSetting_Language` | Known | User setting |
| 0x007DC080 | `ToggleSetting_Repeat` | Known | User setting |
| 0x007DC0BA | `ToggleSetting_Shuffle` | Known | User setting |
| 0x007DC17C | `ToggleSetting_TVOut` | Known | User setting |
| 0x007DC1B5 | `ToggleSetting_TVSignal` | Known | User setting |

---

## 7. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000091AB | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | Known | RTOS task thread |
| 0x000E9134 | `HostOSTask` | Known | RTOS task thread |
| 0x00149E6C | `USBDeviceTask` | Known | RTOS task thread |
| 0x001541BC | `DiskReaderTask` | Known | RTOS task thread |
| 0x001645E8 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x001645FC | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0019D400 | `GeniusMixesTask` | Known | RTOS task thread |
| 0x001BA43C | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x001F6358 | `MeCCAIOTask` | Known | RTOS task thread |
| 0x0022A3A0 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x002AF250 | `FirewireTask` | Known | RTOS task thread |
| 0x002AF264 | `TouchwheelTask` | Known | RTOS task thread |
| 0x002AF278 | `AudioOutStateTask` | Known | RTOS task thread |
| 0x002AF2A4 | `DiskMgrTask` | Known | RTOS task thread |
| 0x002AF2B4 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x002AF2C8 | `MikeyTask` | Known | RTOS task thread |
| 0x002AF2D8 | `TopPlugTask` | Known | RTOS task thread |
| 0x002AF2E8 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x002AF360 | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x002AF388 | `AlarmTask` | Known | RTOS task thread |
| 0x002AF3A7 | `"USBAudioTask` | Known | RTOS task thread |
| 0x002C0F3D | `** Clock Snapshot **` | Known | RTOS task thread |
| 0x002C13F4 | `  K - RTXC` | Known | RTOS task thread |
| 0x002C1559 | `** Mailbox Snapshot **` | Known | RTOS task thread |
| 0x002C179D | `** Queue Snapshot **` | Known | RTOS task thread |
| 0x002C1A00 | `** Task Register Snapshot **` | Known | RTOS task thread |
| 0x002C1A94 | `Undefined Task` | Known | RTOS task thread |
| 0x002C1BD5 | `** Resource Snapshot **` | Known | RTOS task thread |
| 0x002C202D | `** Semaphore Snapshot **` | Known | RTOS task thread |
| 0x002C2365 | `** Stack Snapshot **` | Known | RTOS task thread |
| 0x002C23FC | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | Known | RTOS task thread |
| 0x002C2679 | `** Task Snapshot **` | Known | RTOS task thread |
| 0x003F6674 | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x003F9E58 | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x00402564 | `MeCCARecordingTask` | Known | RTOS task thread |

---

## 8. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002669E4 | `Channel Reserved` | Known | Logging channel |
| 0x002669F8 | `Channel AppBoot` | Known | Logging channel |
| 0x00266A08 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x00266A24 | `Channel PrefsWriting` | Known | Logging channel |
| 0x00266A3C | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x00266A5C | `Channel PlayFromDisk` | Known | Logging channel |
| 0x00266A74 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x00266A90 | `Channel TestLogging` | Known | Logging channel |
| 0x00266AA4 | `Channel AppFileLoading` | Known | Logging channel |
| 0x00266ABC | `Channel VCardReading` | Known | Logging channel |
| 0x00266AD4 | `Channel LongSongScanning` | Known | Logging channel |
| 0x00266B48 | `Channel VoiceRecording` | Known | Logging channel |
| 0x00266B60 | `Channel PhotoImporting` | Known | Logging channel |
| 0x00266B78 | `Channel Notes` | Known | Logging channel |
| 0x00266B88 | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x00266BA4 | `Channel DiskMode` | Known | Logging channel |
| 0x00266BB8 | `Channel Firewire` | Known | Logging channel |
| 0x00266BCC | `Channel USB` | Known | Logging channel |
| 0x00266BEC | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x00266C04 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 9. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0000680F | `"MeCCADecode` | Known | Audio system |
| 0x0013FDFC | `HandleEQ` | Known | Audio system |
| 0x001514FC | `AudioCodecs` | Known | Audio system |
| 0x00152B1C | `VideoCodecs` | Known | Audio system |
| 0x001692B8 | `MeCCA_RecordingBuffer` | Known | Audio system |
| 0x00198560 | `MeCCA_PCM_Output.wav` | Known | Audio system |
| 0x001B20BC | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x001BCE0C | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x001BD014 | `MeCCAVideoDecode` | Known | Audio system |
| 0x007474D3 | `Settings_EQMenu_Layout"` | Known | Audio system |
| 0x00747548 | `SettingsMenus_EQ_Layout` | Known | Audio system |
| 0x00748226 | `HandleEQ` | Known | Audio system |
| 0x007FFEA8 | `ACELP is either registered trademark or trademark of VoiceAge Corporation in the` | Known | Audio system |
| 0x0091EF14 | `MeCCA_StreamCache` | Known | Audio system |
| 0x00933224 | `ERROR: unknownCodec loaded !!!` | Known | Audio system |
| 0x00990127 | `11TCEQSetting` | Known | Audio system |
| 0x009902EC | `13TCSettings_EQ` | Known | Audio system |
| 0x00990F89 | `27TSilverCntlrTransitionAddonI11TCEQSettingE` | Known | Audio system |
| 0x009911AC | `27TSilverCntlrTransitionAddonI13TCSettings_EQE` | Known | Audio system |
| 0x009BE7F4 | `SettingsMenu_ListItem_EQ` | Known | Audio system |
| 0x009C00D6 | `Settings_EQ_RandB_Image` | Known | Audio system |
| 0x009C0198 | `Settings_EQ_Electronic_Image` | Known | Audio system |
| 0x009C01CF | `Settings_EQ_Acoustic_Image` | Known | Audio system |
| 0x009C0540 | `Settings_EQ_SpokenWord_Image` | Known | Audio system |
| 0x009C05BD | `Settings_EQ_Dance_Image` | Known | Audio system |
| 0x009C065B | `Settings_EQ_Lounge_Image` | Known | Audio system |
| 0x009C0C35 | `Settings_EQ_Rock_Image` | Known | Audio system |
| 0x009C0CBA | `Settings_EQ_Classical_Image` | Known | Audio system |
| 0x009C102F | `Settings_EQ_Latin_Image` | Known | Audio system |
| 0x009C1227 | `Settings_EQ_Piano_Image` | Known | Audio system |
| 0x009C1664 | `Settings_EQ_Deep_Image` | Known | Audio system |
| 0x009C16C1 | `Settings_EQ_HipHop_Image` | Known | Audio system |
| 0x009C16DA | `Settings_EQ_Pop_Image` | Known | Audio system |
| 0x009C1877 | `Settings_EQ_TrebleReducer_Image` | Known | Audio system |
| 0x009C1897 | `Settings_EQ_BassReducer_Image` | Known | Audio system |
| 0x009C1B33 | `Settings_EQ_TrebleBooster_Image` | Known | Audio system |
| 0x009C1B53 | `Settings_EQ_VocalBooster_Image` | Known | Audio system |
| 0x009C1B72 | `Settings_EQ_BassBooster_Image` | Known | Audio system |
| 0x009C1C99 | `Settings_EQ_SmallSpeakers_Image` | Known | Audio system |
| 0x009C1CDD | `Settings_EQ_Loudness_Image` | Known | Audio system |
| 0x009C1DA2 | `Settings_EQ_Flat_Image` | Known | Audio system |
| 0x009C27C4 | `Settings_EQ_Jazz_Image` | Known | Audio system |
| 0x009C349C | `SettingsEQ_Template` | Known | Audio system |
| 0x009C50CD | `SettingsMenu_EQ_String` | Known | Audio system |
| 0x009D3E2E | `SettingsMenus_EQ_Layout` | Known | Audio system |
| 0x009D40DD | `SettingsEQ_Template_Layout` | Known | Audio system |
| 0x009D4C2F | `Settings_EQMenu_Layout` | Known | Audio system |
| 0x009DAF9A | `msCodeCom` | Known | Audio system |

---

## 10. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001515D8 | `Audible` | Known | Audible audiobook format |
| 0x007FFFEC | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x00800041 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x0080605C | `, %d Audibles` | Known | Audible audiobook format |
| 0x0080606C | `, 1 Audible` | Known | Audible audiobook format |
| 0x008451C0 | `, %d Audible` | Known | Audible audiobook format |
| 0x008451D0 | `, 1 Audible` | Known | Audible audiobook format |
| 0x0084EAC4 | `, %d Audibles` | Known | Audible audiobook format |
| 0x0084EAD4 | `, 1 Audible` | Known | Audible audiobook format |
| 0x00858464 | `, %d Audibles` | Known | Audible audiobook format |
| 0x00858474 | `, 1 Audible` | Known | Audible audiobook format |
| 0x00870A38 | `, %d Audibles` | Known | Audible audiobook format |
| 0x00870A48 | `, 1 Audible` | Known | Audible audiobook format |
| 0x0087A318 | `, %d Audiblea` | Known | Audible audiobook format |
| 0x0087A328 | `, 1 Audible` | Known | Audible audiobook format |
| 0x008A2488 | ` Audible` | Known | Audible audiobook format |
| 0x008A249F | ` Audible` | Known | Audible audiobook format |
| 0x008AD7F2 | ` Audible` | Known | Audible audiobook format |
| 0x008AD805 | ` Audible` | Known | Audible audiobook format |
| 0x008C1160 | `, %d Audible` | Known | Audible audiobook format |
| 0x008C1170 | `, 1 Audible` | Known | Audible audiobook format |
| 0x008CA754 | `, %d Audible` | Known | Audible audiobook format |
| 0x008CA764 | `, 1 Audible` | Known | Audible audiobook format |
| 0x008D4460 | `, %d Audibles` | Known | Audible audiobook format |
| 0x008D4470 | `, 1 Audible` | Known | Audible audiobook format |
| 0x008DF73C | `, %d Audibles` | Known | Audible audiobook format |
| 0x008DF74C | `, 1 Audible` | Known | Audible audiobook format |
| 0x008EB900 | `, %d Audible` | Known | Audible audiobook format |
| 0x008EB910 | `, 1 Audible` | Known | Audible audiobook format |
| 0x008F508C | `, %d Audible` | Known | Audible audiobook format |
| 0x008F509C | `, 1 Audible` | Known | Audible audiobook format |
| 0x008FEC4D | ` Audible` | Known | Audible audiobook format |
| 0x008FEC60 | ` Audible` | Known | Audible audiobook format |
| 0x009084FD | ` Audible` | Known | Audible audiobook format |
| 0x00908510 | ` Audible` | Known | Audible audiobook format |

---

## 11. Audio/Codec - Apple Lossless

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000A6C43 | `"alac: bit depth = %d, pb = 0x%X, mb = 0x%X, kb = 0x%X ` | Known | Apple Lossless codec |
| 0x001515AC | `AppleLossless` | Known | Apple Lossless codec |
| 0x00155F4C | `alacmp4v@KL` | Known | Apple Lossless codec |
| 0x001C4A14 | `elsttkhdmdhdstsdsttsstszstscstcomp4aalac` | Known | Apple Lossless codec |
| 0x008F84AD | ` geri alacakt` | Known | Apple Lossless codec |
| 0x008F8522 | ` geri alacakt` | Known | Apple Lossless codec |

---

## 12. Audio/Codec - AAC

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0049F118 | `!"#$%%#!&%"'##&%()))))))(()())))*+,*+++++**--.+*///0//00/000/00/1221113111411551` | Known | AAC codec |
| 0x004E110C | `#$%$$%$$$$$%$$$%!&$!$$%$$$$%$!%$%$%%$$%$#'())((()(())((((((((()(((()((())((()(()` | Known | AAC codec |
| 0x005C0C97 | `AAAAAAAAAAC` | Known | AAC codec |
| 0x006C14CF | `B+A22AAAC` | Known | AAC codec |

---

## 13. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x007FFE4C | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |

---

## 14. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000A9129 | `;=1sinf` | Known | DRM system |
| 0x001514D0 | `AppleDRMVersion` | Known | DRM system |
| 0x00151570 | `AppleDRM` | Known | DRM system |
| 0x00152B30 | `AppleVideoDRM` | Known | DRM system |
| 0x00155F30 | `tx3gdrmsp608aavdmp4aesdsD` | Known | DRM system |
| 0x001C30E4 | `tkhdedtselstmdiamdhdminfstblstsdstcoco64stscstszsttsstssdrmidrms` | Known | DRM system |
| 0x008C2044 | `Ingen enhetsinfo tilgjengelig.` | Known | DRM system |
| 0x008EC877 | `rsinformation finns tillg` | Known | DRM system |
| 0x008EE3CD | `ningsinformationens riktighet.` | Known | DRM system |
| 0x009BC8AA | `DRMLevel` | Known | DRM system |
| 0x009E2E40 | `$Apple FairPlay Certificate Authority0` | Known | DRM system |
| 0x009E31C5 | `&Apple FairPlay Certification Authority0` | Known | DRM system |
| 0x00A17D6D | `&Apple FairPlay Certification Authority0` | Known | DRM system |
| 0x00A17DE3 | `Apple FairPlay1402` | Known | DRM system |

---

## 15. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000987A8 | `gamedata_RW` | Known | Game system |
| 0x000987C4 | `gamedata_ShareRW` | Known | Game system |
| 0x000987D8 | `games_RO` | Known | Game system |
| 0x001261DC | `StartGame` | Known | Game system |
| 0x0075EBEA | `controller.StartGame1` | Known | Game system |
| 0x0075EDCE | `controller.StartGame1` | Known | Game system |
| 0x0075EFB2 | `controller.StartGame1` | Known | Game system |
| 0x0075F196 | `controller.StartGame1` | Known | Game system |
| 0x0075F37A | `controller.StartGame1` | Known | Game system |
| 0x0075F744 | `controller.StartGame1` | Known | Game system |
| 0x00990143 | `11TCGamesMenu` | Known | Game system |
| 0x00990217 | `12TCGameScreen` | Known | Game system |
| 0x00990FE3 | `27TSilverCntlrTransitionAddonI11TCGamesMenuE` | Known | Game system |
| 0x00991098 | `27TSilverCntlrTransitionAddonI12TCGameScreenE` | Known | Game system |
| 0x009C833A | `GamesMenu_StatusBar_String` | Known | Game system |
| 0x009C8AC2 | `AboutScreen_Games_String` | Known | Game system |
| 0x009CFBDC | `MainMenu_List_Games` | Known | Game system |
| 0x009CFBF0 | `ExtrasMenu_Games` | Known | Game system |
| 0x009D7D8B | `MainMenu_List_Games_x` | Known | Game system |

---

## 16. Photo System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00069230 | ` rtSPhotos\Photo Database` | Known | Photo system |
| 0x00104220 | `TPhotosBrowseCntlr` | Known | Photo system |
| 0x0010423C | `TPhotosBrowseTransitionCntlr` | Known | Photo system |
| 0x00104264 | `TPhotosMenuCntlr` | Known | Photo system |
| 0x00104280 | `TPhotosSettingsCntlr` | Known | Photo system |
| 0x001042A0 | `TPhotosSettingsCntlr_Transitions` | Known | Photo system |
| 0x001042CC | `TPhotosSettingsCntlr_Duration` | Known | Photo system |
| 0x001042F4 | `TPhotosSettingsSlideshowPlaylistCntlr` | Known | Photo system |
| 0x001CE8F0 | `PhotoBrowse/Slideshow` | Known | Photo system |
| 0x0026895C | `Photo Database Size` | Known | Photo system |
| 0x004039EC | `TPhotosBrowseCntlr` | Known | Photo system |
| 0x00403A00 | `TPhotosBrowseTransitionCntlr` | Known | Photo system |
| 0x00403A20 | `TPhotosMenuCntlr` | Known | Photo system |
| 0x00403A34 | `TPhotosSettingsCntlr` | Known | Photo system |
| 0x00403A4C | `TPhotosSettingsCntlr_Transitions` | Known | Photo system |
| 0x00403A70 | `TPhotosSettingsCntlr_Duration` | Known | Photo system |
| 0x00403CBC | `TPhotosSettingsSlideshowPlaylistCntlr` | Known | Photo system |
| 0x00746AEF | `Photos_SettingsMenu` | Known | Photo system |
| 0x0075B6D8 | `TPhotosMenuCntlrTSilverCntlrTPhotosSettingsCntlrTPhotosSettingsCntlr_Transitions` | Known | Photo system |
| 0x0075B748 | `TPhotosSettingsSlideshowPlaylistCntlr` | Known | Photo system |
| 0x0075B770 | `TPhotosBrowseCntlr` | Known | Photo system |
| 0x00763D4F | `Photos_Menu` | Known | Photo system |
| 0x00765E1B | `Photos_Menu` | Known | Photo system |
| 0x00767EE7 | `Photos_Menu` | Known | Photo system |
| 0x00769FB3 | `Photos_Menu` | Known | Photo system |
| 0x0076C07F | `Photos_Menu` | Known | Photo system |
| 0x0076E14B | `Photos_Menu` | Known | Photo system |
| 0x00770217 | `Photos_Menu` | Known | Photo system |
| 0x007722E3 | `Photos_Menu` | Known | Photo system |
| 0x007743AF | `Photos_Menu` | Known | Photo system |
| 0x0077647B | `Photos_Menu` | Known | Photo system |
| 0x00778547 | `Photos_Menu` | Known | Photo system |
| 0x0077A613 | `Photos_Menu` | Known | Photo system |
| 0x0077C6DF | `Photos_Menu` | Known | Photo system |
| 0x0077E7AB | `Photos_Menu` | Known | Photo system |
| 0x00780877 | `Photos_Menu` | Known | Photo system |
| 0x00782943 | `Photos_Menu` | Known | Photo system |
| 0x00784A0F | `Photos_Menu` | Known | Photo system |
| 0x00786ADB | `Photos_Menu` | Known | Photo system |
| 0x00788BA7 | `Photos_Menu` | Known | Photo system |
| 0x0078AC73 | `Photos_Menu` | Known | Photo system |
| 0x0078CD3F | `Photos_Menu` | Known | Photo system |
| 0x0078EE0B | `Photos_Menu` | Known | Photo system |
| 0x00790ED7 | `Photos_Menu` | Known | Photo system |
| 0x00792FA3 | `Photos_Menu` | Known | Photo system |
| 0x0079506F | `Photos_Menu` | Known | Photo system |
| 0x0079713B | `Photos_Menu` | Known | Photo system |
| 0x00799207 | `Photos_Menu` | Known | Photo system |
| 0x0079B2D3 | `Photos_Menu` | Known | Photo system |
| 0x0079D39F | `Photos_Menu` | Known | Photo system |
| 0x0079F46B | `Photos_Menu` | Known | Photo system |
| 0x007A1537 | `Photos_Menu` | Known | Photo system |
| 0x007A3603 | `Photos_Menu` | Known | Photo system |
| 0x007A56CF | `Photos_Menu` | Known | Photo system |
| 0x007A779B | `Photos_Menu` | Known | Photo system |
| 0x007A9867 | `Photos_Menu` | Known | Photo system |
| 0x007AB933 | `Photos_Menu` | Known | Photo system |
| 0x007AD9FF | `Photos_Menu` | Known | Photo system |
| 0x007AFACB | `Photos_Menu` | Known | Photo system |
| 0x007DBDBF | `PhotoBrowse_Small` | Known | Photo system |
| 0x007DBE68 | `Photos_SettingsMenu` | Known | Photo system |
| 0x007DBEE8 | `PhotoBrowse_Small"` | Known | Photo system |
| 0x007DBFAB | `Photos_SettingsDurationMenu` | Known | Photo system |
| 0x007DC13D | `Photos_SettingsTransitionMenu` | Known | Photo system |
| 0x00807F2C | `PhotoBrowse_Large` | Known | Photo system |
| 0x0091331C | `TPhotosSettingsCntlr` | Known | Photo system |
| 0x00913334 | `TPhotosSettingsCntlr_Transitions` | Known | Photo system |
| 0x00913358 | `TPhotosSettingsCntlr_Duration` | Known | Photo system |
| 0x009905C8 | `16TPhotosMenuCntlr` | Known | Photo system |
| 0x00990813 | `18TPhotosBrowseCntlr` | Known | Photo system |
| 0x0099098C | `20TPhotosSettingsCntlr` | Known | Photo system |
| 0x00991518 | `27TSilverCntlrTransitionAddonI16TPhotosMenuCntlrE` | Known | Photo system |
| 0x0099167E | `27TSilverCntlrTransitionAddonI18TPhotosBrowseCntlrE` | Known | Photo system |
| 0x009918C9 | `27TSilverCntlrTransitionAddonI20TPhotosSettingsCntlrE` | Known | Photo system |
| 0x00991F9E | `27TSilverCntlrTransitionAddonI28TPhotosBrowseTransitionCntlrE` | Known | Photo system |
| 0x00992190 | `27TSilverCntlrTransitionAddonI29TPhotosSettingsCntlr_DurationE` | Known | Photo system |
| 0x009925CF | `27TSilverCntlrTransitionAddonI32TPhotosSettingsCntlr_TransitionsE` | Known | Photo system |
| 0x0099286C | `27TSilverCntlrTransitionAddonI37TPhotosSettingsSlideshowPlaylistCntlrE` | Known | Photo system |
| 0x00992B31 | `28TPhotosBrowseTransitionCntlr` | Known | Photo system |
| 0x00992C2B | `29TPhotosSettingsCntlr_Duration` | Known | Photo system |
| 0x00992E7C | `32TPhotosSettingsCntlr_Transitions` | Known | Photo system |
| 0x00993056 | `37TPhotosSettingsSlideshowPlaylistCntlr` | Known | Photo system |
| 0x009BF73E | `PhotoBrowse_Grid` | Known | Photo system |
| 0x009BF87E | `Settings_Capacity_PhotosLegend` | Known | Photo system |
| 0x009C03A3 | `Settings_About_Capacity_PhotosLegend_Image` | Known | Photo system |
| 0x009C2917 | `PhotoBrowse_Large` | Known | Photo system |
| 0x009C38E6 | `PhotoBrowse_Template` | Known | Photo system |
| 0x009C5530 | `Photos_Settings_Music_String` | Known | Photo system |
| 0x009C5CD9 | `Photos_Import_Browse_Choice_String` | Known | Photo system |
| 0x009C6026 | `Photos_Settings_Time_Per_Slide_String` | Known | Photo system |
| 0x009C6619 | `Photos_All_Photos_Browse_String` | Known | Photo system |
| 0x009C69B9 | `Photos_Settings_Music_Off_String` | Known | Photo system |
| 0x009C6F20 | `Photos_Settings_Music_NowPlaying_String` | Known | Photo system |
| 0x009C7272 | `Photos_Settings_TV_Signal_String` | Known | Photo system |
| 0x009C81A7 | `Photos_Browse_1_Photo_String` | Known | Photo system |
| 0x009C81F6 | `Photos_Settings_Music_FromiPhoto_String` | Known | Photo system |
| 0x009C8BF9 | `Photos_Settings_String` | Known | Photo system |
| 0x009C90C7 | `Photos_Settings_Transitions_String` | Known | Photo system |
| 0x009C91C4 | `Photos_Settings_Shuffle_Photos_String` | Known | Photo system |
| 0x009C921D | `MainMenu_Photos_String` | Known | Photo system |
| | *...and 33 more* | | |

---

## 17. Video System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x009CDF92 | `NowPlaying_RentalWarning_Dialog_Template_Video` | Known | Video system |
| 0x009CDFE5 | `NowPlaying_RentalWarning_Overlay_Template_Video` | Known | Video system |
| 0x009D011B | `MainMenu_Video_List_Rentals` | Known | Video system |
| 0x009D7DB7 | `MainMenu_Video_List_Rentals_x` | Known | Video system |

---

## 18. Genius

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000AC88C | `TGeniusLoadingCntlr` | Known | Genius system |
| 0x000B608C | `GeniusPlaylist` | Known | Genius system |
| 0x00113E38 | `GeniusPlaylist_` | Known | Genius system |
| 0x00152A50 | `SupportsGenius` | Known | Genius system |
| 0x00152A60 | `GeniusConfigMinVersion` | Known | Genius system |
| 0x00152A78 | `GeniusMetadataMinVersion` | Known | Genius system |
| 0x00152A94 | `GeniusSimilaritiesMinVersion` | Known | Genius system |
| 0x00152AB4 | `GeniusConfigMaxVersion` | Known | Genius system |
| 0x00152ACC | `GeniusMetadataMaxVersion` | Known | Genius system |
| 0x00152AE8 | `GeniusSimilaritiesMaxVersion` | Known | Genius system |
| 0x00152B08 | `SupportsGeniusMixes` | Known | Genius system |
| 0x001B6878 | `GeniusMixArtwork` | Known | Genius system |
| 0x001DFBCC | `RefreshingGenius` | Known | Genius system |
| 0x001DFBE4 | `CreatingGeniusMix` | Known | Genius system |
| 0x001DFED0 | `GeniusPlaylistReady` | Known | Genius system |
| 0x001DFEE4 | `GeniusMixPlaylistReady` | Known | Genius system |
| 0x0021F5A4 | `GotoGeniusLayout` | Known | Genius system |
| 0x0021F72C | `GotoGeniusError_NoGenius` | Known | Genius system |
| 0x0021F760 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Genius system |
| 0x0021F788 | `StartGenius` | Known | Genius system |
| 0x0022077C | `GotoGeniusError_NoGenius` | Known | Genius system |
| 0x00220798 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Genius system |
| 0x00220C50 | `StartGenius` | Known | Genius system |
| 0x00224978 | `StartGenius` | Known | Genius system |
| 0x00224BA0 | `StartGenius` | Known | Genius system |
| 0x00227E44 | `GotoGeniusError_NoGenius` | Known | Genius system |
| 0x00227E78 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Genius system |
| 0x00227EA0 | `StartGenius` | Known | Genius system |
| 0x00228020 | `StartGenius` | Known | Genius system |
| 0x00228030 | `RefreshingGenius` | Known | Genius system |
| 0x002282E8 | `StartGenius` | Known | Genius system |
| 0x002312D8 | `GotoGeniusMixesIntro` | Known | Genius system |
| 0x002312F4 | `GotoGeniusMixes` | Known | Genius system |
| 0x00231304 | `GotoSingleGeniusMix` | Known | Genius system |
| 0x00232B74 | `StartGenius` | Known | Genius system |
| 0x00233E18 | `StartGenius` | Known | Genius system |
| 0x00234624 | `StartGenius` | Known | Genius system |
| 0x0023463C | `GotoGenius` | Known | Genius system |
| 0x00234654 | `SavedGeniusPlaylist` | Known | Genius system |
| 0x002348C4 | `SavedGeniusPlaylist` | Known | Genius system |
| 0x00234AC0 | `GotoGeniusIntro` | Known | Genius system |
| 0x00234AD4 | `GotoGenius` | Known | Genius system |
| 0x00234B14 | `GeniusPlaylistSelected` | Known | Genius system |
| 0x00237668 | `CreatingGeniusMix` | Known | Genius system |
| 0x0023A7E0 | `GotoGeniusError_NoGenius` | Known | Genius system |
| 0x0023A814 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Genius system |
| 0x0023A83C | `StartGenius` | Known | Genius system |
| 0x0023A9BC | `RefreshingGenius` | Known | Genius system |
| 0x0023AE14 | `StartGenius` | Known | Genius system |
| 0x0023B1A4 | `GeniusPlaylistSelected` | Known | Genius system |
| 0x004043BC | `TGeniusLoadingCntlr` | Known | Genius system |
| 0x0075ADA8 | `TContextualMenuCntlrTCExtrasMenuTSilverCntlrTGeniusLoadingCntlr` | Known | Genius system |
| 0x0076256D | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x00763062 | `controller.GotoGenius1` | Known | Genius system |
| 0x007630E0 | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x00763159 | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x007631E6 | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x00763565 | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x00764639 | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x0076512E | `controller.GotoGenius1` | Known | Genius system |
| 0x007651AC | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x00765225 | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x007652B2 | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x00765631 | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x00766705 | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x007671FA | `controller.GotoGenius1` | Known | Genius system |
| 0x00767278 | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x007672F1 | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x0076737E | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x007676FD | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x007687D1 | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x007692C6 | `controller.GotoGenius1` | Known | Genius system |
| 0x00769344 | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x007693BD | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x0076944A | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x007697C9 | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x0076A89D | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x0076B392 | `controller.GotoGenius1` | Known | Genius system |
| 0x0076B410 | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x0076B489 | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x0076B516 | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x0076B895 | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x0076C969 | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x0076D45E | `controller.GotoGenius1` | Known | Genius system |
| 0x0076D4DC | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x0076D555 | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x0076D5E2 | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x0076D961 | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x0076EA35 | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x0076F52A | `controller.GotoGenius1` | Known | Genius system |
| 0x0076F5A8 | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x0076F621 | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x0076F6AE | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x0076FA2D | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x00770B01 | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x007715F6 | `controller.GotoGenius1` | Known | Genius system |
| 0x00771674 | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x007716ED | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x0077177A | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x00771AF9 | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| | *...and 821 more* | | |

---

## 19. Database (SQLite)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00238964 | `%s/sqlite_` | Known | SQLite database |
| 0x002C4EE8 | `sqlite3BtreeInitPage() returns error code %d` | Known | SQLite database |
| 0x002C8200 | `sqlite_master` | Known | SQLite database |
| 0x002C8210 | `sqlite_temp_master` | Known | SQLite database |
| 0x002DE3D0 | `sqlite_stat1` | Known | SQLite database |
| 0x002DE3E0 | `CREATE TABLE %Q.sqlite_stat1(tbl,idx,stat)` | Known | SQLite database |
| 0x002DE40C | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x002E9148 | `sqlite_subquery_%p_` | Known | SQLite database |
| 0x0037104C | `sqlite_master` | Known | SQLite database |
| 0x0037105C | `sqlite_temp_master` | Known | SQLite database |
| 0x00371380 | `sqlite_` | Known | SQLite database |
| 0x003713C0 | `sqlite_master` | Known | SQLite database |
| 0x003713D0 | `sqlite_temp_master` | Known | SQLite database |
| 0x003713E8 | `sqlite_sequence` | Known | SQLite database |
| 0x003713F8 | `UPDATE "%w".sqlite_sequence set name = %Q WHERE name = %Q` | Known | SQLite database |
| 0x003714DC | `sqlite_stat1` | Known | SQLite database |
| 0x003714EC | `SELECT idx, stat FROM %Q.sqlite_stat1` | Known | SQLite database |
| 0x003721C8 | `sqlite_` | Known | SQLite database |
| 0x003723C4 | `sqlite_master` | Known | SQLite database |
| 0x003723D4 | `sqlite_temp_master` | Known | SQLite database |
| 0x003750F0 | `sqlite_` | Known | SQLite database |
| 0x003763DC | `sqlite_autoindex_` | Known | SQLite database |
| 0x003763F0 | `sqlite_master` | Known | SQLite database |
| 0x00376400 | `sqlite_temp_master` | Known | SQLite database |
| 0x00377858 | `sqlite_master` | Known | SQLite database |
| 0x00377868 | `sqlite_temp_master` | Known | SQLite database |
| 0x0037789C | `sqlite_stat1` | Known | SQLite database |
| 0x003778AC | `DELETE FROM %Q.sqlite_stat1 WHERE idx=%Q` | Known | SQLite database |
| 0x00377B94 | `sqlite_master` | Known | SQLite database |
| 0x00377BA4 | `sqlite_temp_master` | Known | SQLite database |
| 0x00377C18 | `DELETE FROM %s.sqlite_sequence WHERE name=%Q` | Known | SQLite database |
| 0x00377C80 | `sqlite_stat1` | Known | SQLite database |
| 0x00377C90 | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x00378008 | `sqlite_master` | Known | SQLite database |
| 0x00378018 | `sqlite_temp_master` | Known | SQLite database |
| 0x00378430 | `sqlite_master` | Known | SQLite database |
| 0x00378440 | `sqlite_temp_master` | Known | SQLite database |
| 0x00378458 | `CREATE TABLE %Q.sqlite_sequence(name,seq)` | Known | SQLite database |
| 0x0037B6E0 | `sqlite_master` | Known | SQLite database |
| 0x0037B6F0 | `sqlite_temp_master` | Known | SQLite database |
| 0x0037DAD8 | `sqlite_temp_master` | Known | SQLite database |
| 0x0037DAF0 | `sqlite_master` | Known | SQLite database |
| 0x0037F2CC | `sqlite3_extension_init` | Known | SQLite database |
| 0x0037FAC0 | `sqlite_master` | Known | SQLite database |
| 0x0037FAD0 | `sqlite_temp_master` | Known | SQLite database |
| 0x00383EB0 | `sqlite_attach` | Known | SQLite database |
| 0x00383EC4 | `sqlite_detach` | Known | SQLite database |
| 0x00386BF8 | `sqlite_master` | Known | SQLite database |
| 0x00386C08 | `sqlite_temp_master` | Known | SQLite database |
| 0x00386C58 | `sqlite_sequence` | Known | SQLite database |
| 0x0038C4E4 | `sqlite_master` | Known | SQLite database |
| 0x0038C4F4 | `sqlite_temp_master` | Known | SQLite database |
| 0x0038F888 | `sqlite_master` | Known | SQLite database |
| 0x0038F898 | `sqlite_temp_master` | Known | SQLite database |
| 0x0039D880 | `sqlite_attach` | Known | SQLite database |
| 0x0039D890 | `sqlite_detach` | Known | SQLite database |
| 0x0080319C | `Richard Hipp (SQLite)` | Known | SQLite database |
| 0x008031B4 | `SQLite Copyright` | Known | SQLite database |
| 0x008031C8 | `All of the deliverable code in SQLite has been dedicated to the public domain by` | Known | SQLite database |
| 0x008033F4 | `The previous paragraph applies to the deliverable code in SQLite - those parts o` | Known | SQLite database |
| 0x008035C4 | `All of the deliverable code in SQLite has been written from scratch. No code has` | Known | SQLite database |
| 0x00803730 | `Obtaining An Explicit License To Use SQLite` | Known | SQLite database |
| 0x0080375C | `Even though SQLite is in the public domain and does not require a license, some ` | Known | SQLite database |
| 0x00803854 | `-You are using SQLite in a jurisdiction that does not recognize the right of an ` | Known | SQLite database |
| 0x008038D8 | `-You want to hold a tangible legal document as evidence that you have the legal ` | Known | SQLite database |
| 0x00803994 | `If you feel like you really have to purchase a license for SQLite, Hwaci, the co` | Known | SQLite database |
| 0x00803A54 | `In order to keep SQLite completely free and unencumbered by copyright, all new c` | Known | SQLite database |
| 0x00803D10 | `We are not able to accept patches or changes to SQLite that are not accompanied ` | Known | SQLite database |
| 0x0090DF0B | `SQLite format 3` | Known | SQLite database |
| 0x009105B8 | `CREATE TABLE sqlite_master(` | Known | SQLite database |
| 0x00910620 | `CREATE TEMP TABLE sqlite_temp_master(` | Known | SQLite database |
| 0x00910CE8 | `illegal return value (%d) from the authorization function - should be SQLITE_OK,` | Known | SQLite database |
| 0x00910DA0 | `SELECT 'CREATE TABLE vacuum_db.' || substr(sql,14)   FROM sqlite_master WHERE ty` | Known | SQLite database |
| 0x00910E28 | `SELECT 'CREATE INDEX vacuum_db.' || substr(sql,14)  FROM sqlite_master WHERE sql` | Known | SQLite database |
| 0x00910E90 | `SELECT 'CREATE UNIQUE INDEX vacuum_db.' || substr(sql,21)   FROM sqlite_master W` | Known | SQLite database |
| 0x00910F08 | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x00910FB8 | `SELECT 'DELETE FROM vacuum_db.' || quote(name) || ';' FROM vacuum_db.sqlite_mast` | Known | SQLite database |
| 0x0091102C | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x009110C4 | `INSERT INTO vacuum_db.sqlite_master   SELECT type, name, tbl_name, rootpage, sql` | Known | SQLite database |
| 0x00911284 | `UPDATE %Q.%s SET sql = CASE WHEN type = 'trigger' THEN sqlite_rename_trigger(sql` | Known | SQLite database |
| 0x009113F8 | `UPDATE sqlite_temp_master SET sql = sqlite_rename_trigger(sql, %Q), tbl_name = %` | Known | SQLite database |
| 0x00911634 | `sqlite3_get_table() called with two or more incompatible queries` | Known | SQLite database |
| 0x009D8890 | `sqlite_rename_table` | Known | SQLite database |
| 0x009D8A13 | `sqlite_version` | Known | SQLite database |
| 0x009D8AAD | `sqlite_rename_trigger` | Known | SQLite database |
| 0x009D8DD1 | `SQLite_iPod_VFS` | Known | SQLite database |

---

## 20. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0011F7A0 | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x0014B534 | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x00151F70 | `iTunesUSupported` | Known | iTunes database |
| 0x0020A2BC | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x0020D1F0 | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x002124D8 | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x002255C0 | `iTunes Image DB` | Known | iTunes database |
| 0x00231604 | `iTunesUSelected` | Known | iTunes database |
| 0x00231614 | `EmptyiTunesUSelected` | Known | iTunes database |
| 0x003EF904 | `iTunesDB` | Known | iTunes database |
| 0x00763BB8 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00764127 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00765C84 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007661F3 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00767D50 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007682BF | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00769E1C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0076A38B | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0076BEE8 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0076C457 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0076DFB4 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0076E523 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00770080 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007705EF | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0077214C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007726BB | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00774218 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00774787 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007762E4 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00776853 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007783B0 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0077891F | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0077A47C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0077A9EB | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0077C548 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0077CAB7 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0077E614 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0077EB83 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007806E0 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00780C4F | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007827AC | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00782D1B | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00784878 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00784DE7 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00786944 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00786EB3 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00788A10 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00788F7F | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0078AADC | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0078B04B | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0078CBA8 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0078D117 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0078EC74 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0078F1E3 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00790D40 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007912AF | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00792E0C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0079337B | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00794ED8 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00795447 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00796FA4 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00797513 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00799070 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007995DF | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0079B13C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0079B6AB | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0079D208 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0079D777 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0079F2D4 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0079F843 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007A13A0 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007A190F | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007A346C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007A39DB | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007A5538 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007A5AA7 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007A7604 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007A7B73 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007A96D0 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007A9C3F | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007AB79C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007ABD0B | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007AD868 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007ADDD7 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007AF934 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007AFEA3 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00803F5C | `iTunes` | Known | iTunes database |
| 0x008041B8 | `You can download music from iTunes.` | Known | iTunes database |
| 0x008041DC | `You can download videos from iTunes.` | Known | iTunes database |
| 0x00804204 | `You can download podcasts from iTunes.` | Known | iTunes database |
| 0x0080422C | `You can download audiobooks from iTunes.` | Known | iTunes database |
| 0x00804258 | `You can download TV shows from iTunes.` | Known | iTunes database |
| 0x00804280 | `You can download movies from iTunes.` | Known | iTunes database |
| 0x008042A8 | `You can download music videos from iTunes.` | Known | iTunes database |
| 0x008042D4 | `You can sync Photos via iTunes.` | Known | iTunes database |
| 0x008042F4 | `You can create playlists and sync via iTunes.` | Known | iTunes database |
| 0x00804324 | `You can download rentals from iTunes.` | Known | iTunes database |
| 0x008045AC | `If you forget the combination, connect to iTunes to unlock your iPod.` | Known | iTunes database |
| 0x00804AD0 | `You can download contacts from iTunes.` | Known | iTunes database |
| 0x00805208 | `To view your To Do items here, enable syncing from iTunes under the Calendar sec` | Known | iTunes database |
| | *...and 430 more* | | |

---

## 21. Nike+ iPod

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00762EFD | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00762F40 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00762F57 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00762F76 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00764FC9 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0076500C | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00765023 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00765042 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00767095 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x007670D8 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x007670EF | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0076710E | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00769161 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x007691A4 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x007691BB | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x007691DA | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0076B22D | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0076B270 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0076B287 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0076B2A6 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0076D2F9 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0076D33C | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0076D353 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0076D372 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0076F3C5 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0076F408 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0076F41F | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0076F43E | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00771491 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x007714D4 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x007714EB | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0077150A | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0077355D | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x007735A0 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x007735B7 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x007735D6 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00775629 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0077566C | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00775683 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x007756A2 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x007776F5 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00777738 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0077774F | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0077776E | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x007797C1 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00779804 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0077981B | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0077983A | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0077B88D | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0077B8D0 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0077B8E7 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0077B906 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0077D959 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0077D99C | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0077D9B3 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0077D9D2 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0077FA25 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0077FA68 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0077FA7F | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0077FA9E | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00781AF1 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00781B34 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00781B4B | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00781B6A | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00783BBD | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00783C00 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00783C17 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00783C36 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00785C89 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00785CCC | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00785CE3 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00785D02 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00787D55 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00787D98 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00787DAF | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00787DCE | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00789E21 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00789E64 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00789E7B | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00789E9A | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0078BEED | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0078BF30 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0078BF47 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0078BF66 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0078DFB9 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0078DFFC | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0078E013 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0078E032 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00790085 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x007900C8 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x007900DF | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x007900FE | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00792151 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00792194 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x007921AB | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x007921CA | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0079421D | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00794260 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00794277 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00794296 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| | *...and 320 more* | | |

---

## 22. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x007473C8 | `Settings_RadioRegionsMenu_Layout"` | Known | FM Radio |
| 0x00747451 | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x00806FBC | `Radio Regions` | Known | FM Radio |
| 0x008596E8 | `Radio-Regionen` | Known | FM Radio |
| 0x00990BE0 | `23TCSettings_RadioRegions` | Known | FM Radio |
| 0x00991AF3 | `27TSilverCntlrTransitionAddonI23TCSettings_RadioRegionsE` | Known | FM Radio |
| 0x009C11D8 | `Settings_Radio_Image` | Known | FM Radio |
| 0x009C530D | `Settings_RadioRegions_Australia_String` | Known | FM Radio |
| 0x009C5334 | `Settings_RadioRegions_Asia_String` | Known | FM Radio |
| 0x009C65B8 | `Settings_RadioRegions_Europe_String` | Known | FM Radio |
| 0x009C7BC7 | `Settings_RadioRegions_Japan_String` | Known | FM Radio |
| 0x009C88DF | `Settings_RadioRegions_Americas_String` | Known | FM Radio |
| 0x009C8FFB | `SettingsMenu_RadioRegions_String` | Known | FM Radio |
| 0x009CC55A | `RadioRegionsMenuItem` | Known | FM Radio |
| 0x009D032F | `SettingsMenu_ListItem_RadioRegions` | Known | FM Radio |
| 0x009D46C2 | `SettingsInfo_Template_RadioRegions_Layout` | Known | FM Radio |
| 0x009D46EC | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x009D4D4E | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 23. Clock/Alarms

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00805C48 | `24 Hour Clock` | Known | Clock system |
| 0x00990448 | `15TCClockCityMenu` | Known | Clock system |
| 0x009906F4 | `17TCClockRegionMenu` | Known | Clock system |
| 0x00991327 | `27TSilverCntlrTransitionAddonI15TCClockCityMenuE` | Known | Clock system |
| 0x0099157D | `27TSilverCntlrTransitionAddonI17TCClockRegionMenuE` | Known | Clock system |
| 0x00992AAD | `27TSilverCntlrTransitionAddonI7TCClockE` | Known | Clock system |
| 0x009931EF | `7TCClock` | Known | Clock system |
| 0x009BDD11 | `Clock_Hours_Image_24` | Known | Clock system |
| 0x009C715B | `Settings_DateTime_24HrClock_String` | Known | Clock system |
| 0x009CB5AB | `DateTime_List_24HrClock` | Known | Clock system |

---

## 24. Storage (CE-ATA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00075748 | `cI: Set drive to MMC high speed failed` | Known | CE-ATA/storage interface |
| 0x00075804 | `cI: could not read CE-ATA task file` | Known | CE-ATA/storage interface |
| 0x0007582C | `cI: CE-ATA signature missing (%x,%x)` | Known | CE-ATA/storage interface |
| 0x00075884 | `cI: CE-ATA interrupt enable failed` | Known | CE-ATA/storage interface |
| 0x000EEE50 | `mI: card not in MMC TRAN state as expected` | Known | CE-ATA/storage interface |
| 0x0036C374 | `MMC init failed` | Known | CE-ATA/storage interface |
| 0x0036C388 | `CE-ATA init failed` | Known | CE-ATA/storage interface |
| 0x0036C848 | `ISDIE: CE-ATA interrupt enable failed` | Known | CE-ATA/storage interface |
| 0x005C524D | `KMMKKKMMMC` | Known | CE-ATA/storage interface |

---

## 25. Storage (NAND Flash)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0013C4B8 | `NAND FLASH DRIVE` | Known | NAND flash interface |

---

## 26. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00151A38 | `FireWireGUID` | Known | FireWire interface |
| 0x00151A48 | `FireWireVersion` | Known | FireWire interface |
| 0x00152424 | `FireWire` | Known | FireWire interface |

---

## 27. Hardware (GPIO)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x003F050D | `GPIO_REG_WRITE` | Known | GPIO hardware |
| 0x003F051E | `GPIO_INT_INIT` | Known | GPIO hardware |

---

## 28. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00129268 | `MonoHope-LCD` | Known | Hardware interface |
| 0x00129880 | `MonoHope-LCD` | Known | Hardware interface |
| 0x00144A1C | `TDiskModeCntlr` | Known | Hardware interface |
| 0x00152530 | `ForcedDiskMode` | Known | Hardware interface |
| 0x0015A1DC | ` rtSltnCpaMBDiskModeImage_SyncArrow1` | Known | Hardware interface |
| 0x001CF134 | `MonoHope-LCD` | Known | Hardware interface |
| 0x002688BC | `Enter Disk Mode` | Known | Hardware interface |
| 0x002688CC | `Exit Disk Mode` | Known | Hardware interface |
| 0x003EFB14 | `TDiskModeCntlr` | Known | Hardware interface |
| 0x003F0551 | `I2C_MASTER` | Known | Hardware interface |
| 0x003F4E62 | `S_I2C_DONE` | Known | Hardware interface |
| 0x003F8A50 | `TDiskModeCntlr` | Known | Hardware interface |
| 0x003F8AB0 | `TDiskModeCntlr` | Known | Hardware interface |
| 0x0075BAF8 | `TDiskModeCntlr` | Known | Hardware interface |
| 0x009903BD | `14TCSlideshowLCD` | Known | Hardware interface |
| 0x009903DF | `14TDiskModeCntlr` | Known | Hardware interface |
| 0x009912C7 | `27TSilverCntlrTransitionAddonI14TCSlideshowLCDE` | Known | Hardware interface |
| 0x009912F7 | `27TSilverCntlrTransitionAddonI14TDiskModeCntlrE` | Known | Hardware interface |
| 0x009BD2E6 | `DiskModeImage_SyncArrow11` | Known | Hardware interface |
| 0x009BD385 | `DiskModeImage_SyncArrow21` | Known | Hardware interface |
| 0x009BD39F | `DiskModeImage_SyncArrow31` | Known | Hardware interface |
| 0x009BD543 | `DiskMode_Text1` | Known | Hardware interface |
| 0x009BD591 | `DiskModeImage_SyncArrow1` | Known | Hardware interface |
| 0x009BD893 | `DiskMode_Text2` | Known | Hardware interface |
| 0x009BD9CB | `DiskModeImage_SyncArrow13` | Known | Hardware interface |
| 0x009BDA3A | `DiskModeImage_SyncArrow23` | Known | Hardware interface |
| 0x009BDA54 | `DiskModeImage_SyncArrow33` | Known | Hardware interface |
| 0x009BDC1B | `DiskModeImage_SyncArrow3` | Known | Hardware interface |
| 0x009BDF05 | `DiskModeImage_SyncArrow15` | Known | Hardware interface |
| 0x009BDF64 | `DiskModeImage_SyncArrow25` | Known | Hardware interface |
| 0x009BDF7E | `DiskModeImage_SyncArrow35` | Known | Hardware interface |
| 0x009BE079 | `DiskModeImage_SyncArrow5` | Known | Hardware interface |
| 0x009BE2A1 | `DiskModeImage_SyncArrow17` | Known | Hardware interface |
| 0x009BE300 | `DiskModeImage_SyncArrow27` | Known | Hardware interface |
| 0x009BE3DA | `DiskModeImage_SyncArrow7` | Known | Hardware interface |
| 0x009BE602 | `DiskModeImage_SyncArrow19` | Known | Hardware interface |
| 0x009BE661 | `DiskModeImage_SyncArrow29` | Known | Hardware interface |
| 0x009BE73B | `DiskModeImage_SyncArrow9` | Known | Hardware interface |
| 0x009BF696 | `DiskMode_View_Connected` | Known | Hardware interface |
| 0x009BF6D1 | `DiskMode_View_Disconnected` | Known | Hardware interface |
| 0x009BFAD7 | `DiskMode_iPod` | Known | Hardware interface |
| 0x009C10C9 | `DiskMode_SyncIcon_Image` | Known | Hardware interface |
| 0x009C10E1 | `DiskMode_ConnectedIcon_Image` | Known | Hardware interface |
| 0x009C111A | `DiskMode_DisconnectIcon_Image` | Known | Hardware interface |
| 0x009C1D50 | `DiskMode_SyncArrows_Image` | Known | Hardware interface |
| 0x009C46B2 | `DiskMode_View_Loading` | Known | Hardware interface |
| 0x009C52FC | `DiskMode__String` | Known | Hardware interface |
| 0x009C5B13 | `DiskMode_Connected_String` | Known | Hardware interface |
| 0x009C6AF5 | `DiskMode_Syncing_String` | Known | Hardware interface |
| 0x009C6B0D | `DiskMode_Loading_String` | Known | Hardware interface |
| 0x009C6F63 | `DiskMode_Synchronizing_String` | Known | Hardware interface |
| 0x009C99D0 | `DiskMode_UseiTunesToEject_String` | Known | Hardware interface |
| 0x009C9A12 | `DiskMode_OKToDisconnect_String` | Known | Hardware interface |
| 0x009C9A31 | `DiskMode_OkayToDisconnect_String` | Known | Hardware interface |
| 0x009C9A52 | `DiskMode_EjectingYouMayDisconnect_String` | Known | Hardware interface |
| 0x009C9C4B | `DiskMode_PleaseWait_String` | Known | Hardware interface |
| 0x009C9C66 | `DiskMode_EjectingPleaseWait_String` | Known | Hardware interface |
| 0x009CAC6E | `DiskMode_View_Synchronizing` | Known | Hardware interface |
| 0x009CBB49 | `DiskModeModel` | Known | Hardware interface |
| 0x009CBD14 | `DiskModeImage_Progress_Full_Fill` | Known | Hardware interface |
| 0x009CBD5C | `DiskModeImage_Progress_Empty_Fill` | Known | Hardware interface |
| 0x009CDA23 | `DiskModeImage_SyncIcon` | Known | Hardware interface |
| 0x009CDA3A | `DiskModeImage_ConnectedIcon` | Known | Hardware interface |
| 0x009CDA56 | `DiskModeImage_DisconnectIcon` | Known | Hardware interface |
| 0x009CE51F | `DiskModeImage_Progress_Full_LeftCap` | Known | Hardware interface |
| 0x009CE543 | `DiskModeImage_Progress_Empty_LeftCap` | Known | Hardware interface |
| 0x009CE568 | `DiskModeImage_Progress_Full_RightCap` | Known | Hardware interface |
| 0x009CE58D | `DiskModeImage_Progress_Empty_RightCap` | Known | Hardware interface |
| 0x009CF38C | `DiskMode_Arrows_Color` | Known | Hardware interface |
| 0x009CF432 | `DiskMode_Text_Color` | Known | Hardware interface |
| 0x009D37BC | `DiskModeLargeFont` | Known | Hardware interface |
| 0x009D37E8 | `DiskModeSmallFont` | Known | Hardware interface |
| 0x009D6296 | `DiskMode_View` | Known | Hardware interface |
| 0x009D6F42 | `DiskMode_Progress_View` | Known | Hardware interface |
| 0x009DD8F4 | `OCSP_RESPID` | Known | Hardware interface |

---

## 29. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00003554 | `iPodPowerProfile.txt` | Known | Power management |
| 0x00144A34 | `TChargingModeCntlr` | Known | Power management |
| 0x00144A50 | `TChargingModeLowPowerCntlr` | Known | Power management |
| 0x00152400 | `PowerInformation` | Known | Power management |
| 0x001529C0 | `BatteryPollInterval` | Known | Power management |
| 0x00268928 | `Begin Charging` | Known | Power management |
| 0x00268938 | `Stop Charging` | Known | Power management |
| 0x002AF290 | `USBPowerSense` | Known | Power management |
| 0x002AF350 | `PCFPowerMgr` | Known | Power management |
| 0x002AF398 | `PowerMgmt` | Known | Power management |
| 0x003EFAF8 | `TChargingModeLowPowerCntlr` | Known | Power management |
| 0x003EFB24 | `TChargingModeCntlr` | Known | Power management |
| 0x003F8A60 | `TChargingModeCntlr` | Known | Power management |
| 0x003F8A74 | `TChargingModeLowPowerCntlr` | Known | Power management |
| 0x003F8B1C | `TChargingModeCntlr` | Known | Power management |
| 0x003F8B30 | `SwitchToCharging` | Known | Power management |
| 0x003F9E44 | `TChargingModeCntlr` | Known | Power management |
| 0x00403EAC | `TChargingModeCntlr` | Known | Power management |
| 0x0075BB08 | `TChargingModeCntlr` | Known | Power management |
| 0x0075BB1C | `TChargingModeLowPowerCntlr` | Known | Power management |
| 0x007DC8DB | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DCE17 | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DD353 | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DD88F | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DDDCB | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DE307 | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DE843 | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DED07 | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DF1CB | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DF68F | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DFB53 | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007E95A4 | `controller.SwitchToCharging1` | Known | Power management |
| 0x00804790 | `Low Battery` | Known | Power management |
| 0x0080479C | `Connect to Power` | Known | Power management |
| 0x00807910 | `Charging` | Known | Power management |
| 0x00807924 | `Low Battery` | Known | Power management |
| 0x00807930 | `Connect to Power` | Known | Power management |
| 0x008470A8 | `PowerSong` | Known | Power management |
| 0x00850738 | `PowerSong` | Known | Power management |
| 0x0085A438 | `PowerSong` | Known | Power management |
| 0x00867608 | `PowerSong` | Known | Power management |
| 0x00872A90 | `PowerSong` | Known | Power management |
| 0x0087BF64 | `PowerSong` | Known | Power management |
| 0x00885A04 | `PowerSong` | Known | Power management |
| 0x0088FCE4 | `PowerSong` | Known | Power management |
| 0x00899868 | `PowerSong` | Known | Power management |
| 0x008A4B90 | `PowerSong` | Known | Power management |
| 0x008AF968 | `PowerSong` | Known | Power management |
| 0x008B972C | `PowerSong` | Known | Power management |
| 0x008C2D10 | `PowerSong` | Known | Power management |
| 0x008CC748 | `PowerSong` | Known | Power management |
| 0x008D6370 | `PowerSong` | Known | Power management |
| 0x008E2E7C | `PowerSong` | Known | Power management |
| 0x008ED518 | `PowerSong` | Known | Power management |
| 0x008F6EE0 | `PowerSong` | Known | Power management |
| 0x00900890 | `PowerSong` | Known | Power management |
| 0x0090A1C8 | `PowerSong` | Known | Power management |
| 0x009905DB | `16TPowerStatusView` | Known | Power management |
| 0x009907E9 | `18TChargingModeCntlr` | Known | Power management |
| 0x00990DD0 | `26TChargingModeLowPowerCntlr` | Known | Power management |
| 0x0099164A | `27TSilverCntlrTransitionAddonI18TChargingModeCntlrE` | Known | Power management |
| 0x00991D3E | `27TSilverCntlrTransitionAddonI26TChargingModeLowPowerCntlrE` | Known | Power management |
| 0x00993398 | `N3ISL17IPodPowerListenerE` | Known | Power management |
| 0x009BD0A5 | `StatusBarWhite_Battery_Image_10` | Known | Power management |
| 0x009BD0C5 | `StatusBarBlack_Battery_Image_10` | Known | Power management |
| 0x009BD12A | `StatusBarWhite_Battery_Image_20` | Known | Power management |
| 0x009BD14A | `StatusBarBlack_Battery_Image_20` | Known | Power management |
| 0x009BD1D2 | `StatusBarWhite_Battery_Image_0` | Known | Power management |
| 0x009BD1F1 | `StatusBarBlack_Battery_Image_0` | Known | Power management |
| 0x009BD2A6 | `StatusBarWhite_Battery_Image_11` | Known | Power management |
| 0x009BD2C6 | `StatusBarBlack_Battery_Image_11` | Known | Power management |
| 0x009BD345 | `StatusBarWhite_Battery_Image_21` | Known | Power management |
| 0x009BD365 | `StatusBarBlack_Battery_Image_21` | Known | Power management |
| 0x009BD442 | `StatusBarWhite_Battery_Image_1` | Known | Power management |
| 0x009BD461 | `StatusBarBlack_Battery_Image_1` | Known | Power management |
| 0x009BD617 | `StatusBarWhite_Battery_Image_12` | Known | Power management |
| 0x009BD637 | `StatusBarBlack_Battery_Image_12` | Known | Power management |
| 0x009BD69C | `StatusBarWhite_Battery_Image_22` | Known | Power management |
| 0x009BD6BC | `StatusBarBlack_Battery_Image_22` | Known | Power management |
| 0x009BD765 | `StatusBarWhite_Battery_Image_2` | Known | Power management |
| 0x009BD784 | `StatusBarBlack_Battery_Image_2` | Known | Power management |
| 0x009BD98B | `StatusBarWhite_Battery_Image_13` | Known | Power management |
| 0x009BD9AB | `StatusBarBlack_Battery_Image_13` | Known | Power management |
| 0x009BDAF7 | `StatusBarWhite_Battery_Image_3` | Known | Power management |
| 0x009BDB16 | `StatusBarBlack_Battery_Image_3` | Known | Power management |
| 0x009BDCA1 | `StatusBarWhite_Battery_Image_14` | Known | Power management |
| 0x009BDCC1 | `StatusBarBlack_Battery_Image_14` | Known | Power management |
| 0x009BDDAF | `StatusBarWhite_Battery_Image_4` | Known | Power management |
| 0x009BDDCE | `StatusBarBlack_Battery_Image_4` | Known | Power management |
| 0x009BDEC5 | `StatusBarWhite_Battery_Image_15` | Known | Power management |
| 0x009BDEE5 | `StatusBarBlack_Battery_Image_15` | Known | Power management |
| 0x009BE012 | `StatusBarWhite_Battery_Image_5` | Known | Power management |
| 0x009BE031 | `StatusBarBlack_Battery_Image_5` | Known | Power management |
| 0x009BE0D7 | `StatusBarWhite_Battery_Image_16` | Known | Power management |
| 0x009BE0F7 | `StatusBarBlack_Battery_Image_16` | Known | Power management |
| 0x009BE1C4 | `StatusBarWhite_Battery_Image_6` | Known | Power management |
| 0x009BE1E3 | `StatusBarBlack_Battery_Image_6` | Known | Power management |
| 0x009BE261 | `StatusBarWhite_Battery_Image_17` | Known | Power management |
| 0x009BE281 | `StatusBarBlack_Battery_Image_17` | Known | Power management |
| 0x009BE382 | `StatusBarWhite_Battery_Image_7` | Known | Power management |
| | *...and 55 more* | | |

---

## 30. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000AB7E8 | `Calendars/` | Known | UI element |
| 0x000BDFF4 | `Calendars` | Known | UI element |
| 0x0010C120 | `AlarmHilited` | Known | UI element |
| 0x0010C234 | `NewAlarmSelected` | Known | UI element |
| 0x0010C248 | `AlarmSelected` | Known | UI element |
| 0x0010C258 | `CalendarEventSelected` | Known | UI element |
| 0x0010CE28 | `GotoNowPlaying` | Known | UI element |
| 0x0010CEA0 | `GotoMainMenu` | Known | UI element |
| 0x00131D40 | `GotoNowPlaying` | Known | UI element |
| 0x00131D54 | `GotoAlbums` | Known | UI element |
| 0x00131D60 | `GotoSongs` | Known | UI element |
| 0x00157164 | `GotoMainMenu` | Known | UI element |
| 0x001E958C | `GotoPlayDeleteMenu` | Known | UI element |
| 0x001E9610 | `GotoNowPlaying` | Known | UI element |
| 0x001F6CE4 | `ToggleAlarm` | Known | UI element |
| 0x00217C20 | `AlarmTonesChosen` | Known | UI element |
| 0x00217C34 | `AlarmToneAt` | Known | UI element |
| 0x00217FDC | `AlarmToneAt` | Known | UI element |
| 0x0021830C | `GotoDefaultLayout` | Known | UI element |
| 0x00218390 | `GotoVolumeLayout` | Known | UI element |
| 0x002184C8 | `GotoProgressLayout` | Known | UI element |
| 0x002187E4 | `GotoDefault` | Known | UI element |
| 0x00218B18 | `GotoProgressLayout` | Known | UI element |
| 0x00218D8C | `GotoRentalWarningLayout` | Known | UI element |
| 0x00218E10 | `GotoProgressLayout` | Known | UI element |
| 0x00219120 | `GotoProgressLayout` | Known | UI element |
| 0x0021AD20 | `GotoNowPlaying` | Known | UI element |
| 0x0021B630 | `GotoNowPlaying` | Known | UI element |
| 0x0021B93C | `GotoNowPlaying` | Known | UI element |
| 0x0021E394 | `GotoStatusBarVideoLayout` | Known | UI element |
| 0x0021E3B0 | `GotoDefaultVideoLayout` | Known | UI element |
| 0x0021E3C8 | `GotoDefaultLayout` | Known | UI element |
| 0x0021E3DC | `GotoDefaultSubtitlesLayout` | Known | UI element |
| 0x0021E474 | `GotoVolumeLayout` | Known | UI element |
| 0x0021E488 | `GotoVolumeVideoLayout` | Known | UI element |
| 0x0021E528 | `GotoProgressLayout` | Known | UI element |
| 0x0021E53C | `GotoProgressVideoLayout` | Known | UI element |
| 0x0021ECF0 | `GotoProgressVideoLayout` | Known | UI element |
| 0x0021F158 | `GotoCaptionVideoLayout` | Known | UI element |
| 0x0021F3C4 | `GotoProgressLayout` | Known | UI element |
| 0x0021F3D8 | `GotoProgressVideoLayout` | Known | UI element |
| 0x0021F580 | `GotoBrightnessVideoLayout` | Known | UI element |
| 0x0021F5B8 | `GotoRatingLayout` | Known | UI element |
| 0x0021FA80 | `GotoChapterArtLayout` | Known | UI element |
| 0x0021FA98 | `GotoShuffleLayout` | Known | UI element |
| 0x0021FE28 | `GotoExtraInfoLayout` | Known | UI element |
| 0x0021FE3C | `GotoExtraInfoLoadingLayout` | Known | UI element |
| 0x0021FF0C | `GotoVolumeLayout` | Known | UI element |
| 0x0021FF24 | `GotoVolumeVideoLayout` | Known | UI element |
| 0x0021FFB8 | `GotoVolumeLayout` | Known | UI element |
| 0x0021FFCC | `GotoVolumeVideoLayout` | Known | UI element |
| 0x002201DC | `GotoScrubLayout` | Known | UI element |
| 0x002201EC | `GotoScrubVideoLayout` | Known | UI element |
| 0x0022027C | `GotoProgressLayout` | Known | UI element |
| 0x00220290 | `GotoProgressVideoLayout` | Known | UI element |
| 0x002204E8 | `GotoStatusBarVideoLayout` | Known | UI element |
| 0x00220504 | `GotoDefaultVideoLayout` | Known | UI element |
| 0x0022051C | `GotoDefaultSubtitlesLayout` | Known | UI element |
| 0x00220538 | `GotoDefaultLayout` | Known | UI element |
| 0x00220D64 | `GotoChapterArtLayout` | Known | UI element |
| 0x00220E5C | `GotoProgressLayout` | Known | UI element |
| 0x00220EE8 | `GotoProgressLayout` | Known | UI element |
| 0x00220EFC | `GotoProgressVideoLayout` | Known | UI element |
| 0x00220FD8 | `GotoExtraInfoLoadFailedLayout` | Known | UI element |
| 0x00220FF8 | `GotoExtraInfoLayout` | Known | UI element |
| 0x00221434 | `GotoStatusBarLayout` | Known | UI element |
| 0x00221448 | `GotoDefaultLayout` | Known | UI element |
| 0x00221620 | `GotoDefault` | Known | UI element |
| 0x00221754 | `GotoProgressLayout` | Known | UI element |
| 0x00221914 | `GotoCaptionVideoLayout` | Known | UI element |
| 0x00221A64 | `GotoBrightnessLayout` | Known | UI element |
| 0x00221AE8 | `GotoBrightnessLayout` | Known | UI element |
| 0x00221B68 | `GotoVolumeLayout` | Known | UI element |
| 0x00221BB4 | `GotoScrubLayout` | Known | UI element |
| 0x00221C7C | `GotoStatusBarLayout` | Known | UI element |
| 0x00221C90 | `GotoDefaultLayout` | Known | UI element |
| 0x00221D68 | `GotoScrubLayout` | Known | UI element |
| 0x00221DB8 | `GotoScrubLayout` | Known | UI element |
| 0x00228558 | `GotoNowPlaying` | Known | UI element |
| 0x00228854 | `GotoNowPlaying` | Known | UI element |
| 0x00229C00 | `GotoFourCard_About` | Known | UI element |
| 0x00229C14 | `GotoThreeCard_About` | Known | UI element |
| 0x0022EF44 | `looBCalendarEventlessDaySelected` | Known | UI element |
| 0x0022EF68 | `CalendarEventfullDaySelected` | Known | UI element |
| 0x0022F084 | `looBCalendarEventlessDaySelected` | Known | UI element |
| 0x0022F0A8 | `CalendarEventfullDaySelected` | Known | UI element |
| 0x0022F174 | `CalendarEventSelected` | Known | UI element |
| 0x00231BFC | `GotoNowPlaying` | Known | UI element |
| 0x00232310 | `GotoNowPlaying` | Known | UI element |
| 0x00232B10 | `GotoFirstBoot` | Known | UI element |
| 0x00232B20 | `GotoNotesApp` | Known | UI element |
| 0x00232B34 | `GotoLockApp` | Known | UI element |
| 0x00235D88 | `looBCalendarEventlessDaySelected` | Known | UI element |
| 0x00235DAC | `CalendarEventfullDaySelected` | Known | UI element |
| 0x00236888 | `CalendarSelected` | Known | UI element |
| 0x00237654 | `GotoNowPlaying` | Known | UI element |
| 0x0023B1D4 | `GotoNowPlaying` | Known | UI element |
| 0x003F6F8C | `AlarmLabelChosen` | Known | UI element |
| 0x003F6FA0 | `AlarmSoundChosen` | Known | UI element |
| 0x003F6FB4 | `AlarmTimeChosen` | Known | UI element |
| | *...and 309 more* | | |

---

## 31. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00803F50 | `Podcasts` | Known | Menu item |
| 0x00803F68 | `Photos` | Known | Menu item |
| 0x00803F70 | `Videos` | Known | Menu item |
| 0x00803F78 | `Extras` | Known | Menu item |
| 0x00803F94 | `Settings` | Known | Menu item |
| 0x00803FA0 | `Shuffle Songs` | Known | Menu item |
| 0x00803FB0 | `Now Playing` | Known | Menu item |
| 0x00803FE4 | `Playlists` | Known | Menu item |
| 0x00803FF0 | `Artists` | Known | Menu item |
| 0x00803FF8 | `Albums` | Known | Menu item |
| 0x00804018 | `Genres` | Known | Menu item |
| 0x00804020 | `Composers` | Known | Menu item |
| 0x0080402C | `Audiobooks` | Known | Menu item |
| 0x00804C8C | `Playlists` | Known | Menu item |
| 0x00805F50 | `Artists` | Known | Menu item |
| 0x00805F58 | `Albums` | Known | Menu item |
| 0x00805F68 | `Genres` | Known | Menu item |
| 0x00805F70 | `Composers` | Known | Menu item |
| 0x00805FA0 | `Photos` | Known | Menu item |
| 0x00805FA8 | `Playlists` | Known | Menu item |
| 0x00805FC4 | `Audiobooks` | Known | Menu item |
| 0x00805FD0 | `Podcasts` | Known | Menu item |
| 0x00806350 | `Settings` | Known | Menu item |
| 0x0080652C | `Albums` | Known | Menu item |
| 0x00806540 | `Now Playing` | Known | Menu item |
| 0x00806B74 | `Photos` | Known | Menu item |
| 0x00806B88 | `Albums` | Known | Menu item |
| 0x00806B90 | `Settings` | Known | Menu item |
| 0x00806BDC | `Settings` | Known | Menu item |
| 0x00806C54 | `Now Playing` | Known | Menu item |
| 0x00806DCC | `Podcasts` | Known | Menu item |
| 0x00806E00 | `Videos` | Known | Menu item |
| 0x00806E08 | `Photos` | Known | Menu item |
| 0x00806F44 | `Main Menu` | Known | Menu item |
| 0x00806FB0 | `Audiobooks` | Known | Menu item |
| 0x00807058 | `Albums` | Known | Menu item |
| 0x0084315C | `Podcasts` | Known | Menu item |
| 0x00845114 | `Podcasts` | Known | Menu item |
| 0x00846110 | `Podcasts` | Known | Menu item |
| 0x0084CA2C | `Podcasts` | Known | Menu item |
| 0x0084EA30 | `Podcasts` | Known | Menu item |
| 0x0085605C | `Podcasts` | Known | Menu item |
| 0x00856070 | `Videos` | Known | Menu item |
| 0x00856078 | `Extras` | Known | Menu item |
| 0x00856130 | `Genres` | Known | Menu item |
| 0x00858358 | `Genres` | Known | Menu item |
| 0x008583D4 | `Podcasts` | Known | Menu item |
| 0x008594CC | `Podcasts` | Known | Menu item |
| 0x008594F4 | `Videos` | Known | Menu item |
| 0x00860114 | `Podcasts` | Known | Menu item |
| 0x00863FA0 | `Podcasts` | Known | Menu item |
| 0x00865BB0 | `Podcasts` | Known | Menu item |
| 0x0086E74C | `Podcasts` | Known | Menu item |
| 0x0086E768 | `Extras` | Known | Menu item |
| 0x008709A8 | `Podcasts` | Known | Menu item |
| 0x00871AC0 | `Podcasts` | Known | Menu item |
| 0x008813B8 | `Podcasts` | Known | Menu item |
| 0x008813C4 | `Photos` | Known | Menu item |
| 0x008813D4 | `Extras` | Known | Menu item |
| 0x00881468 | `Albums` | Known | Menu item |
| 0x0088148C | `Genres` | Known | Menu item |
| 0x008837BC | `Albums` | Known | Menu item |
| 0x008837D0 | `Genres` | Known | Menu item |
| 0x00883814 | `Photos` | Known | Menu item |
| 0x0088385C | `Podcasts` | Known | Menu item |
| 0x00883EF0 | `Albums` | Known | Menu item |
| 0x00884634 | `Photos` | Known | Menu item |
| 0x00884644 | `Albums` | Known | Menu item |
| 0x00884908 | `Podcasts` | Known | Menu item |
| 0x0088493C | `Photos` | Known | Menu item |
| 0x00884C30 | `Albums` | Known | Menu item |
| 0x008B5724 | `Podcasts` | Known | Menu item |
| 0x008B57D4 | `Albums` | Known | Menu item |
| 0x008B57F0 | `Genres` | Known | Menu item |
| 0x008B7820 | `Albums` | Known | Menu item |
| 0x008B7830 | `Genres` | Known | Menu item |
| 0x008B78A0 | `Podcasts` | Known | Menu item |
| 0x008B7ED0 | `Albums` | Known | Menu item |
| 0x008B85A4 | `Albums` | Known | Menu item |
| 0x008B884C | `Podcasts` | Known | Menu item |
| 0x008B8AE8 | `Albums` | Known | Menu item |
| 0x008D20F0 | `Podcasts` | Known | Menu item |
| 0x008D210C | `Extras` | Known | Menu item |
| 0x008D43A8 | `Podcasts` | Known | Menu item |
| 0x008D54C0 | `Podcasts` | Known | Menu item |
| 0x00A23BE4 | `Settings` | Known | Menu item |

---

## 32. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00047DEC | `iPod_Control` | Filesystem Path | |
| 0x00047E00 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x00047E18 | `iPod_Control\iTunes\firsttime` | Filesystem Path | |
| 0x00047E58 | `iPod_Control\Device` | Filesystem Path | |
| 0x00056B70 | `iPod_Control\Device` | Filesystem Path | |
| 0x00058C38 | `iPod_Control` | Filesystem Path | |
| 0x000592A8 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x0006924C | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path | |
| 0x0006926C | `iPod_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x00069294 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x0006BE94 | `iPod_Control\Music\` | Filesystem Path | |
| 0x0006EED4 | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x0006F050 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000946E0 | `iPod_Control/iTunes/` | Filesystem Path | |
| 0x00098738 | `iPod_Control/iTunes/iTunesDB.p7b` | Filesystem Path | |
| 0x000A27D0 | `iPod_Control` | Filesystem Path | |
| 0x000A27E0 | `Resources/Games` | Filesystem Path | |
| 0x000A27F0 | `iPod_Control/%s%s%s` | Filesystem Path | |
| 0x000ABE40 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000B5550 | `iPod_Control/iTunes/` | Filesystem Path | |
| 0x000B5738 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000B6064 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000BE20C | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000BF7B0 | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000BF8B0 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x00100354 | `iPod_Control\Device\dst` | Filesystem Path | |
| 0x0010B48C | `iPod_Control/Device/alarms` | Filesystem Path | |
| 0x0011BD40 | `iPod_Control/Device/radio` | Filesystem Path | |
| 0x0011D280 | `iPod_Control/Device` | Filesystem Path | |
| 0x0011D294 | `iPod_Control/Device/radio` | Filesystem Path | |
| 0x00135060 | `/Resources/Icons/` | Filesystem Path | |
| 0x00137644 | `iPod_Control/Device/Users` | Filesystem Path | |
| 0x0013CC54 | `iPod_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x00165A8C | `/iPod_Control/Device/1da` | Filesystem Path | |
| 0x00165CE8 | `/iPod_Control/Device/1da` | Filesystem Path | |
| 0x001728F4 | `Resources/UI/active.bin` | Filesystem Path | |
| 0x0017290C | `Resources/UI/` | Filesystem Path | |
| 0x00196B94 | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x00197AC0 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path | |
| 0x00197AE8 | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001BDA84 | `iPod_Control/Device/PlayCounts` | Filesystem Path | |
| 0x001D3E58 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D3F08 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D4084 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D421C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D42C4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D4484 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D4528 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D45CC | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D4670 | `iPod_Control\Device\` | Filesystem Path | |

---

## 33. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00914610 | `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftutil.c` | Build Path | |
| 0x00914668 | `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftstream.c` | Build Path | |
| 0x009146C4 | `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftobjs.c` | Build Path | |
| 0x0091F064 | `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\afglobal.c` | Build Path | |
| 0x0091FBE0 | `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfdrivr.c` | Build Path | |
| 0x00920DDC | `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrgload.c` | Build Path | |
| 0x00920E34 | `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrcmap.c` | Build Path | |
| 0x00920E8C | `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrobjs.c` | Build Path | |
| 0x009211D0 | `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1cmap.c` | Build Path | |
| 0x00930578 | `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttcmap.c` | Build Path | |
| 0x009307F4 | `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype\ttgload.c` | Build Path | |
| 0x00930D60 | `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1load.c` | Build Path | |

---

## 34. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000A0418 | `Acoustic` | EQ Preset | |
| 0x000A0424 | `Bass Booster` | EQ Preset | |
| 0x000A0444 | `Classical` | EQ Preset | |
| 0x000A0460 | `Electronic` | EQ Preset | |
| 0x000A0474 | `Hip Hop` | EQ Preset | |
| 0x000A048C | `Loudness` | EQ Preset | |
| 0x000A0498 | `Lounge` | EQ Preset | |
| 0x000A04BC | `Small Speakers` | EQ Preset | |
| 0x000A04CC | `Spoken Word` | EQ Preset | |
| 0x000A04D8 | `Treble Booster` | EQ Preset | |
| 0x000A0524 | `Vocal Booster` | EQ Preset | |
| 0x0080732C | `Acoustic` | EQ Preset | |
| 0x00807338 | `Bass Booster` | EQ Preset | |
| 0x00807358 | `Classical` | EQ Preset | |
| 0x00807374 | `Electronic` | EQ Preset | |
| 0x00807388 | `Hip Hop` | EQ Preset | |
| 0x008073A0 | `Loudness` | EQ Preset | |
| 0x008073AC | `Lounge` | EQ Preset | |
| 0x008073CC | `Small Speakers` | EQ Preset | |
| 0x008073DC | `Spoken Word` | EQ Preset | |
| 0x008073E8 | `Treble Booster` | EQ Preset | |
| 0x00807408 | `Vocal Booster` | EQ Preset | |
| 0x008467BC | `Acoustic` | EQ Preset | |
| 0x008467C8 | `Bass Booster` | EQ Preset | |
| 0x008467E8 | `Classical` | EQ Preset | |
| 0x00846804 | `Electronic` | EQ Preset | |
| 0x00846818 | `Hip Hop` | EQ Preset | |
| 0x00846830 | `Loudness` | EQ Preset | |
| 0x0084683C | `Lounge` | EQ Preset | |
| 0x0084685C | `Small Speakers` | EQ Preset | |
| 0x0084686C | `Spoken Word` | EQ Preset | |
| 0x00846878 | `Treble Booster` | EQ Preset | |
| 0x00846898 | `Vocal Booster` | EQ Preset | |
| 0x0084FEBC | `Acoustic` | EQ Preset | |
| 0x0084FEC8 | `Bass Booster` | EQ Preset | |
| 0x0084FEE8 | `Classical` | EQ Preset | |
| 0x0084FF04 | `Electronic` | EQ Preset | |
| 0x0084FF18 | `Hip Hop` | EQ Preset | |
| 0x0084FF30 | `Loudness` | EQ Preset | |
| 0x0084FF3C | `Lounge` | EQ Preset | |
| 0x0084FF5C | `Small Speakers` | EQ Preset | |
| 0x0084FF6C | `Spoken Word` | EQ Preset | |
| 0x0084FF78 | `Treble Booster` | EQ Preset | |
| 0x0084FF98 | `Vocal Booster` | EQ Preset | |
| 0x00859B28 | `Acoustic` | EQ Preset | |
| 0x00859B68 | `Electronic` | EQ Preset | |
| 0x00859B94 | `Loudness` | EQ Preset | |
| 0x008721C0 | `Hip Hop` | EQ Preset | |
| 0x008721D0 | `Latina` | EQ Preset | |
| 0x008721D8 | `Loudness` | EQ Preset | |
| 0x008721E4 | `Lounge` | EQ Preset | |
| 0x0087B958 | `Lounge` | EQ Preset | |
| 0x00885034 | `Hip Hop` | EQ Preset | |
| 0x00885044 | `Latino` | EQ Preset | |
| 0x00885058 | `Lounge` | EQ Preset | |
| 0x00898FEC | `Hip Hop` | EQ Preset | |
| 0x00898FFC | `Latina` | EQ Preset | |
| 0x00899004 | `Loudness` | EQ Preset | |
| 0x00899010 | `Lounge` | EQ Preset | |
| 0x008A40B4 | `Acoustic` | EQ Preset | |
| 0x008A40C0 | `Bass Booster` | EQ Preset | |
| 0x008A40E0 | `Classical` | EQ Preset | |
| 0x008A40FC | `Electronic` | EQ Preset | |
| 0x008A4110 | `Hip Hop` | EQ Preset | |
| 0x008A4128 | `Loudness` | EQ Preset | |
| 0x008A4134 | `Lounge` | EQ Preset | |
| 0x008A4154 | `Small Speakers` | EQ Preset | |
| 0x008A4164 | `Spoken Word` | EQ Preset | |
| 0x008A4170 | `Treble Booster` | EQ Preset | |
| 0x008A4190 | `Vocal Booster` | EQ Preset | |
| 0x008AEFF0 | `Acoustic` | EQ Preset | |
| 0x008AEFFC | `Bass Booster` | EQ Preset | |
| 0x008AF01C | `Classical` | EQ Preset | |
| 0x008AF038 | `Electronic` | EQ Preset | |
| 0x008AF04C | `Hip Hop` | EQ Preset | |
| 0x008AF064 | `Loudness` | EQ Preset | |
| 0x008AF070 | `Lounge` | EQ Preset | |
| 0x008AF090 | `Small Speakers` | EQ Preset | |
| 0x008AF0A0 | `Spoken Word` | EQ Preset | |
| 0x008AF0AC | `Treble Booster` | EQ Preset | |
| 0x008AF0CC | `Vocal Booster` | EQ Preset | |
| 0x008B8EC8 | `Loudness` | EQ Preset | |
| 0x008B8ED4 | `Lounge` | EQ Preset | |
| 0x008C2534 | `Latino` | EQ Preset | |
| 0x008C253C | `Loudness` | EQ Preset | |
| 0x008C2548 | `Lounge` | EQ Preset | |
| 0x008CBE5C | `Hip Hop` | EQ Preset | |
| 0x008CBE88 | `Lounge` | EQ Preset | |
| 0x008D5BB8 | `Hip Hop` | EQ Preset | |
| 0x008D5BC8 | `Latina` | EQ Preset | |
| 0x008D5BDC | `Lounge` | EQ Preset | |
| 0x008ECCF8 | `Acoustic` | EQ Preset | |
| 0x008ECD04 | `Bass Booster` | EQ Preset | |
| 0x008ECD24 | `Classical` | EQ Preset | |
| 0x008ECD40 | `Electronic` | EQ Preset | |
| 0x008ECD54 | `Hip Hop` | EQ Preset | |
| 0x008ECD6C | `Loudness` | EQ Preset | |
| 0x008ECD78 | `Lounge` | EQ Preset | |
| 0x008ECD98 | `Small Speakers` | EQ Preset | |
| 0x008ECDA8 | `Spoken Word` | EQ Preset | |
| 0x008ECDB4 | `Treble Booster` | EQ Preset | |
| 0x008ECDD4 | `Vocal Booster` | EQ Preset | |
| 0x008F6624 | `Hip Hop` | EQ Preset | |
| 0x00900054 | `Acoustic` | EQ Preset | |
| 0x00900060 | `Bass Booster` | EQ Preset | |
| 0x00900080 | `Classical` | EQ Preset | |
| 0x0090009C | `Electronic` | EQ Preset | |
| 0x009000B0 | `Hip Hop` | EQ Preset | |
| 0x009000C8 | `Loudness` | EQ Preset | |
| 0x009000D4 | `Lounge` | EQ Preset | |
| 0x009000F4 | `Small Speakers` | EQ Preset | |
| 0x00900104 | `Spoken Word` | EQ Preset | |
| 0x00900110 | `Treble Booster` | EQ Preset | |
| 0x00900130 | `Vocal Booster` | EQ Preset | |
| 0x00909968 | `Acoustic` | EQ Preset | |
| 0x00909974 | `Bass Booster` | EQ Preset | |
| 0x00909994 | `Classical` | EQ Preset | |
| 0x009099B0 | `Electronic` | EQ Preset | |
| 0x009099C4 | `Hip Hop` | EQ Preset | |
| 0x009099DC | `Loudness` | EQ Preset | |
| 0x009099E8 | `Lounge` | EQ Preset | |
| 0x00909A08 | `Small Speakers` | EQ Preset | |
| 0x00909A18 | `Spoken Word` | EQ Preset | |
| 0x00909A24 | `Treble Booster` | EQ Preset | |
| 0x00909A44 | `Vocal Booster` | EQ Preset | |

---

## 35. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000CE960 | `cIC12: ATA Status Error! Could not get error code.` | Diagnostic | |
| 0x000CE994 | `cIC12: ATA Status Error! Error code (0x%2x)` | Diagnostic | |
| 0x000DC384 | `Error[%d] has occurred in rule %d` | Diagnostic | |
| 0x0010CE10 | `SwitchToNotesImageError` | Diagnostic | |
| 0x00118BC8 | `%s Error in file %s.` | Diagnostic | |
| 0x001DFEB8 | `GotoErrorLayout` | Diagnostic | |
| 0x0075BF41 | `controller.GotoErrorLayout1` | Diagnostic | |
| 0x0075C013 | `controller.GotoErrorLayout1` | Diagnostic | |
| 0x0075EA60 | `controller.ShowMemoryError1` | Diagnostic | |
| 0x0075EAC1 | `controller.ShowSigningError1` | Diagnostic | |
| 0x0075EB24 | `controller.ShowUnknownError1` | Diagnostic | |
| 0x0075EB87 | `controller.ShowVersionError1` | Diagnostic | |
| 0x0075EC44 | `controller.ShowMemoryError1` | Diagnostic | |
| 0x0075ECA5 | `controller.ShowSigningError1` | Diagnostic | |
| 0x0075ED08 | `controller.ShowUnknownError1` | Diagnostic | |
| 0x0075ED6B | `controller.ShowVersionError1` | Diagnostic | |
| 0x0075EE28 | `controller.ShowMemoryError1` | Diagnostic | |
| 0x0075EE89 | `controller.ShowSigningError1` | Diagnostic | |
| 0x0075EEEC | `controller.ShowUnknownError1` | Diagnostic | |
| 0x0075EF4F | `controller.ShowVersionError1` | Diagnostic | |
| 0x0075F00C | `controller.ShowMemoryError1` | Diagnostic | |
| 0x0075F06D | `controller.ShowSigningError1` | Diagnostic | |
| 0x0075F0D0 | `controller.ShowUnknownError1` | Diagnostic | |
| 0x0075F133 | `controller.ShowVersionError1` | Diagnostic | |
| 0x0075F1F0 | `controller.ShowMemoryError1` | Diagnostic | |
| 0x0075F251 | `controller.ShowSigningError1` | Diagnostic | |
| 0x0075F2B4 | `controller.ShowUnknownError1` | Diagnostic | |
| 0x0075F317 | `controller.ShowVersionError1` | Diagnostic | |
| 0x0075F5BA | `controller.ShowMemoryError1` | Diagnostic | |
| 0x0075F61B | `controller.ShowSigningError1` | Diagnostic | |

---

## 36. Assertions

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0004F79C | `assertion failed on line %d of file %s` | Assertion | |
| 0x0004F888 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0004F954 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0004F9E8 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0004FA64 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0004FADC | `assertion failed on line %d of file %s` | Assertion | |
| 0x0004FE20 | `assertion failed on line %d of file %s` | Assertion | |
| 0x00050150 | `assertion failed on line %d of file %s` | Assertion | |
| 0x00050290 | `assertion failed on line %d of file %s` | Assertion | |
| 0x000503D0 | `assertion failed on line %d of file %s` | Assertion | |
| 0x00050500 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0006ADC4 | `assertion failed on line %d of file %s` | Assertion | |
| 0x00086664 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0009C4C0 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0009D494 | `assertion failed on line %d of file %s` | Assertion | |
| 0x000A3B78 | `assertion failed on line %d of file %s` | Assertion | |
| 0x000A99A8 | `assertion failed on line %d of file %s` | Assertion | |
| 0x000B64F4 | `assertion failed on line %d of file %s` | Assertion | |
| 0x000B7108 | `assertion failed on line %d of file %s` | Assertion | |
| 0x000B7244 | `assertion failed on line %d of file %s` | Assertion | |

---
