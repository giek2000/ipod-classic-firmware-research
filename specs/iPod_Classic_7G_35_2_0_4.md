# iPod Classic 6G/7G (Rev A/B) - RetailOS 2.0.4 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 2.0.4 |
| **IPSW** | iPod_35.2.0.4.ipsw |
| **Device** | iPod Classic 6G/7G (Rev A/B) (2007, Click Wheel, Cover Flow, Genius) |
| **UpdaterFamilyID** | 35 |
| **Binary Size** | 10,599,920 bytes (10.11 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,597,872 bytes |
| **Total Strings (>=6)** | 55,243 |
| **Function Prologues** | 23,036 (ARM: 17,721, Thumb: 5,315) |
| **SoC** | Samsung S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Encrypted** | Yes (HW AES) |
| **SHA-256** | `f4368251a58b2fdc7b46acf3178dae1d24bc1e029736240741015851256c65c4` |

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001443A0 | `MP3ExampleTask` | Hidden | Hidden Test |
| 0x0015B0B4 | `MockupMode/MockupMode.xml` | Hidden | Developer Tool |
| 0x0015B5B4 | `MockupMode/` | Hidden | Developer Tool |
| 0x001892F4 | `TCDemoMode` | Hidden | Demo/Retail Mode |
| 0x001E04F4 | `TSilverCntlrTestAppCntlr` | Hidden | Developer Tool |
| 0x002652B4 | `Channel UnitTests` | Hidden | Developer Tool |
| 0x002BFA1D | `** RTXCbug - ` | Hidden | Developer Tool |
| 0x002BFA60 | `  X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x002BFA75 | `RTXCbug> ` | Hidden | Developer Tool |
| 0x002C0451 | `RTXCbug - RTXC Objects> ` | Hidden | Developer Tool |
| 0x002D9D24 | `X - Exit RTXCbug` | Hidden | Developer Tool |
| 0x00394719 | `$RTXCbug> ` | Hidden | Developer Tool |
| 0x003947E1 | `Re-entering RTXCbug mode` | Hidden | Developer Tool |
| 0x003F3175 | `S_RTXCBUG` | Hidden | Developer Tool |
| 0x003F6AB0 | `TCDemoMode` | Hidden | Demo/Retail Mode |
| 0x004020E8 | `TSilverCntlrTestAppCntlr` | Hidden | Developer Tool |
| 0x007395C2 | `Debug_MainMenu_Screen` | Hidden | Debug/Diagnostic |
| 0x007395DB | `Debug_MainMenu_Screen_Default"` | Hidden | Debug/Diagnostic |
| 0x00739649 | `Extras_Screen_Debug` | Hidden | Debug/Diagnostic |
| 0x00759004 | `TSilverSettingsMenuListCntlrTCDateTimeScreenTCDateTimeScreenTCTimeZoneScreenTCDe` | Hidden | Demo/Retail Mode |
| 0x0075990C | `TSilverCntlrTUnitTestSuiteCntlr` | Hidden | Developer Tool |
| 0x0075992C | `TUnitTestSuiteTestsCntlrTSilverCntlrTCVoiceMemosMainMenuTCVoiceMemosMenuTCVoiceM` | Hidden | Developer Tool |
| 0x007E4EE6 | `Debug_UnitTest_Screen` | Hidden | Developer Tool |
| 0x007E4EFF | `Debug_UnitTest_Screen_Default` | Hidden | Developer Tool |
| 0x007E4F62 | `DemoMode_Screen` | Hidden | Demo/Retail Mode |
| 0x007E4F75 | `DemoMode_Screen_Default` | Hidden | Demo/Retail Mode |
| 0x007E4FE2 | `Debug_TestList_Screen` | Hidden | Debug/Diagnostic |
| 0x007E4FFB | `Debug_TestList_Screen_Default` | Hidden | Debug/Diagnostic |
| 0x007E506E | `Debug_TestResult_Screen` | Hidden | Debug/Diagnostic |
| 0x007E5089 | `Debug_TestResult_Screen_Default` | Hidden | Debug/Diagnostic |
| 0x008013B4 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0083F1EC | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008520B0 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0086A278 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0087CDCC | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00886EAC | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00890C6C | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008A6654 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008B0694 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008D7C38 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x008F6E04 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x00900588 | `DemoMode` | Hidden | Demo/Retail Mode |
| 0x0090A214 | `WaveFileDebugTask` | Hidden | Debug/Diagnostic |
| 0x0090B398 | `TCMockupModeNavScreen` | Hidden | Developer Tool |
| 0x009883D9 | `10TCDemoMode` | Hidden | Demo/Retail Mode |
| 0x00988D50 | `21TCMockupModeNavScreen` | Hidden | Developer Tool |
| 0x00988FB6 | `24TSilverCntlrTestAppCntlr` | Hidden | Developer Tool |
| 0x00989210 | `27TSilverCntlrTransitionAddonI10TCDemoModeE` | Hidden | Demo/Retail Mode |
| 0x00989ED7 | `27TSilverCntlrTransitionAddonI24TSilverCntlrTestAppCntlrE` | Hidden | Developer Tool |
| 0x009B439C | `Returning from RTXCBug` | Hidden | Developer Tool |
| 0x009B7C7F | `Debug_ListItem_DemoMode` | Hidden | Debug/Diagnostic |
| 0x009B7C97 | `Debug_MenuItem_DemoMode` | Hidden | Debug/Diagnostic |
| 0x009B839C | `DemoMode_Background_Image` | Hidden | Demo/Retail Mode |
| 0x009B8FF8 | `DemoMode_Icon_Image` | Hidden | Demo/Retail Mode |
| 0x009BABDB | `Debug_Menu_Title` | Hidden | Debug/Diagnostic |
| 0x009BAC00 | `UnitTest_Menu_Title` | Hidden | Developer Tool |
| 0x009C2A2D | `Extras_Screen_Debug` | Hidden | Debug/Diagnostic |
| 0x009C2A41 | `MainMenu_List_Debug` | Hidden | Debug/Diagnostic |
| 0x009C2A55 | `ExtrasMenu_Debug` | Hidden | Debug/Diagnostic |
| 0x009C37C4 | `UnitTestModel` | Hidden | Developer Tool |
| 0x009C41A3 | `UnitTest_ListItem` | Hidden | Developer Tool |
| 0x009C44C3 | `DemoMode_Screen` | Hidden | Demo/Retail Mode |
| 0x009C472C | `DemoMode_Main_Screen` | Hidden | Demo/Retail Mode |
| 0x009C4E72 | `Debug_TestResult_Screen` | Hidden | Debug/Diagnostic |
| 0x009C4EC2 | `Debug_UnitTest_Screen` | Hidden | Developer Tool |
| 0x009C4EEA | `Debug_TestList_Screen` | Hidden | Debug/Diagnostic |
| 0x009C5052 | `Debug_MainMenu_Screen` | Hidden | Debug/Diagnostic |
| 0x009C532C | `DemoMode_View_Main` | Hidden | Demo/Retail Mode |
| 0x009C5529 | `DemoMode_Icon` | Hidden | Demo/Retail Mode |
| 0x009C631C | `SilverTestApp` | Hidden | Developer Tool |
| 0x009C632A | `UnitTestApp` | Hidden | Developer Tool |
| 0x009C68DC | `Debug_ListItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009C68F7 | `Debug_MenuItem_DiskBrowser` | Hidden | Debug/Diagnostic |
| 0x009C7053 | `DemoMode_Text_Color` | Hidden | Demo/Retail Mode |
| 0x009C7468 | `Debug_ListItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009C747F | `Debug_MenuItem_Monitor` | Hidden | Debug/Diagnostic |
| 0x009CA1A5 | `DemoMode_Screen_Default` | Hidden | Demo/Retail Mode |
| 0x009CA8CE | `Debug_TestResult_Screen_Default` | Hidden | Debug/Diagnostic |
| 0x009CA925 | `Debug_UnitTest_Screen_Default` | Hidden | Developer Tool |
| 0x009CA95D | `Debug_TestList_Screen_Default` | Hidden | Debug/Diagnostic |
| 0x009CAA94 | `Debug_MainMenu_Screen_Default` | Hidden | Debug/Diagnostic |
| 0x009CB649 | `Debug_ListItem_UnitTest` | Hidden | Developer Tool |
| 0x009CB661 | `Debug_MenuItem_UnitTest` | Hidden | Developer Tool |
| 0x009CFC52 | `Debug_ListItem_Memory` | Hidden | Debug/Diagnostic |
| 0x009CFC68 | `Debug_MenuItem_Memory` | Hidden | Debug/Diagnostic |
| 0x00A1B3D8 | `DebugUtil` | Hidden | Debug/Diagnostic |

---

## 2. Controllers (TSilver/TC Classes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000AC06C | `TSilverCntlr` | Known | UI controller |
| 0x000AC084 | `TCExtrasMenu` | Known | UI controller |
| 0x000AC09C | `TCGameScreen` | Known | UI controller |
| 0x000AC0B4 | `TCGamesMenu` | Known | UI controller |
| 0x000AC0C8 | `TSilverMainMediaListCntlr_Main` | Known | UI controller |
| 0x000AC0F0 | `TSilverMainMediaListCntlr_Music` | Known | UI controller |
| 0x000AC118 | `TSilverMainMediaListCntlr_Videos` | Known | UI controller |
| 0x000AC144 | `TSilverMediaListCntlr_Songs` | Known | UI controller |
| 0x000AC168 | `TSilverMediaListCntlr_Albums` | Known | UI controller |
| 0x000AC190 | `TSilverMediaListCntlr_Artists` | Known | UI controller |
| 0x000AC1B8 | `TSilverMediaListCntlr_Genres` | Known | UI controller |
| 0x000AC1E0 | `TSilverMediaListCntlr_Composers` | Known | UI controller |
| 0x000AC208 | `TSilverMediaListCntlr_Podcasts` | Known | UI controller |
| 0x000AC230 | `TSilverMediaListCntlr_PodcastEpisodes` | Known | UI controller |
| 0x000AC260 | `TSilverMediaListCntlr_iTunesU` | Known | UI controller |
| 0x000AC288 | `TSilverMediaListCntlr_iTunesUEpisodes` | Known | UI controller |
| 0x000AC2B8 | `TSilverMediaListCntlr_Audiobooks` | Known | UI controller |
| 0x000AC2E4 | `TSilverMediaListCntlr_AudiobookChapters` | Known | UI controller |
| 0x000AC314 | `TSilverMediaListCntlr_TVShows` | Known | UI controller |
| 0x000AC33C | `TSilverMediaListCntlr_TVSeasons` | Known | UI controller |
| 0x000AC364 | `TSilverMediaListCntlr_TVEpisodes` | Known | UI controller |
| 0x000AC390 | `TSilverMediaListCntlr_Movies` | Known | UI controller |
| 0x000AC3B8 | `TSilverMediaListCntlr_Playlists` | Known | UI controller |
| 0x000AC3E0 | `TSilverMediaListCntlr_NestedPlaylists` | Known | UI controller |
| 0x000AC410 | `TSilverMediaListCntlr_VideoPlaylists` | Known | UI controller |
| 0x000AC5AC | `TSilverMediaListCntlr_NestedVideoPlaylists` | Known | UI controller |
| 0x000AC5E0 | `TSilverMediaListCntlr_PlaylistChooser` | Known | UI controller |
| 0x000AC610 | `TSilverMediaListCntlr_Rentals` | Known | UI controller |
| 0x000AC638 | `TSilverMediaListCntlr_Genius` | Known | UI controller |
| 0x000AC660 | `TSilverMediaListCntlr_GeniusMixes` | Known | UI controller |
| 0x000AC68C | `TCRentalNotification` | Known | UI controller |
| 0x000AC6AC | `TCRentalInfo` | Known | UI controller |
| 0x000AC6C4 | `TCRentalConfirmDelete` | Known | UI controller |
| 0x000AC6E4 | `TCRentalDispatcher` | Known | UI controller |
| 0x000AC73C | `TSilverGlobalCntlr` | Known | UI controller |
| 0x000AC758 | `TSilverTrainerCntlr` | Known | UI controller |
| 0x0010402C | `TCSlideshowLCD` | Known | UI controller |
| 0x00104044 | `TCSlideshowTVOut` | Known | UI controller |
| 0x00104060 | `TCSlideshow_TVOutAsk` | Known | UI controller |
| 0x00104080 | `TCSlideshow_TVOutCableConnect` | Known | UI controller |
| 0x00127E04 | `TSilverCalendarCntlr_CalendarMenu` | Known | UI controller |
| 0x00127E30 | `TSilverCalendarCntlr_MonthViewer` | Known | UI controller |
| 0x00127E5C | `TSilverCalendarCntlr_DayViewer` | Known | UI controller |
| 0x00127E84 | `TSilverCalendarCntlr_EventViewer` | Known | UI controller |
| 0x00127EB0 | `TSilverCalendarCntlr_ToDoList` | Known | UI controller |
| 0x00127ED8 | `TSilverCalendarCntlr_ToDoDispatcher` | Known | UI controller |
| 0x00127F04 | `TSilverCalendarCntlr_Alarm` | Known | UI controller |
| 0x0012F880 | `TCRemoteUI` | Known | UI controller |
| 0x0012F894 | `TCUnsupported` | Known | UI controller |
| 0x00135D60 | `TCSpeakers` | Known | UI controller |
| 0x00135D74 | `TCEQSetting` | Known | UI controller |
| 0x0015F23C | `TCSportTimer` | Known | UI controller |
| 0x0015F254 | `TCSportTimerMenu` | Known | UI controller |
| 0x0015F270 | `TCSportTimerSessionScreen` | Known | UI controller |
| 0x0015F294 | `TCSportTimerChosenDispatcher` | Known | UI controller |
| 0x00160644 | `TCVoiceMemos` | Known | UI controller |
| 0x0016065C | `TCVoiceMemosMenu` | Known | UI controller |
| 0x00160678 | `TCVoiceMemosMainMenu` | Known | UI controller |
| 0x00160698 | `TCVoiceMemosPlayback` | Known | UI controller |
| 0x001606B8 | `TCVoiceMemosContextMenu` | Known | UI controller |
| 0x001606D8 | `TCVoiceMemosAlert` | Known | UI controller |
| 0x001725D0 | `TSilverSettingsMenuListCntlr` | Known | UI controller |
| 0x001725F8 | `TCSettings_MainMenu` | Known | UI controller |
| 0x00172614 | `TCSettings_MusicMenu` | Known | UI controller |
| 0x00172634 | `TCSettings_VolumeLimit` | Known | UI controller |
| 0x00172654 | `TCSettings_Brightness` | Known | UI controller |
| 0x00172674 | `TCSettings_BacklightTimer` | Known | UI controller |
| 0x00172698 | `TCSettings_EQ` | Known | UI controller |
| 0x001726B0 | `TCSettings_AudiobookSettings` | Known | UI controller |
| 0x001726D8 | `TCSettings_RadioRegions` | Known | UI controller |
| 0x001726F8 | `TCSettings_ResetAllSettings` | Known | UI controller |
| 0x0017271C | `TSilverSettingsVideoCntlr` | Known | UI controller |
| 0x00172740 | `TCDateTimeScreen` | Known | UI controller |
| 0x0017275C | `TCTimeZoneScreen` | Known | UI controller |
| 0x00172778 | `TCSettings_AdjustScrollingCntlr` | Known | UI controller |
| 0x001727A0 | `TCFirstBoot` | Known | UI controller |
| 0x001B243C | `TCAddressViewerMainMenu` | Known | UI controller |
| 0x001B245C | `TCAddressViewerDetails` | Known | UI controller |
| 0x001B247C | `TCAddressViewerPartialLoad` | Known | UI controller |
| 0x001B24A0 | `TCAddressViewerMainDispatcher` | Known | UI controller |
| 0x001E0518 | `TSilverCntlrTestCntlr` | Known | UI controller |
| 0x001E7D5C | `TSilverMainMediaListCntlr_Videos` | Known | UI controller |
| 0x0027FC00 | `TC_LockDialog` | Known | UI controller |
| 0x0027FC18 | `TC_LockScreen` | Known | UI controller |
| 0x0027FC30 | `TC_LockediPod` | Known | UI controller |
| 0x0027FC48 | `TC_VolumeLimitLockScreen` | Known | UI controller |
| 0x0027FC6C | `TCLockChosenDispatcher` | Known | UI controller |
| 0x0028582C | `TCClock` | Known | UI controller |
| 0x0028583C | `TCClockCityMenu` | Known | UI controller |
| 0x00285854 | `TCClockRegionMenu` | Known | UI controller |
| 0x00285870 | `TCAlarmMenu` | Known | UI controller |
| 0x00285884 | `TCSleepTimerMenu` | Known | UI controller |
| 0x002858A0 | `TCAlarmPropertiesMenu` | Known | UI controller |
| 0x002858C0 | `TCAlarmPropertiesFrequencyMenu` | Known | UI controller |
| 0x002858E8 | `TCAlarmPropertiesLabelMenu` | Known | UI controller |
| 0x0028590C | `TCAlarmPropertiesSoundMenu` | Known | UI controller |
| 0x00285930 | `TCAlarmDatePicker` | Known | UI controller |
| 0x0028594C | `TCAlarmTriggered` | Known | UI controller |
| 0x0028C8F4 | `TCNotesDispatcher` | Known | UI controller |
| 0x0028C910 | `TCNotesLoading` | Known | UI controller |
| | *...and 407 more* | | |

---

## 3. Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000FF7B0 | `HandleWheel` | Known | Event handler |
| 0x000FF7BC | `HandlePlayPause` | Known | Event handler |
| 0x000FF7CC | `HandleSelectDown` | Known | Event handler |
| 0x000FF7E0 | `HandleNext` | Known | Event handler |
| 0x000FF7EC | `HandlePrevious` | Known | Event handler |
| 0x000FF7FC | `HandleNextPushAndHold` | Known | Event handler |
| 0x000FF814 | `HandlePreviousPushAndHold` | Known | Event handler |
| 0x000FFAAC | `HandleCycleBacklightSetting` | Known | Event handler |
| 0x000FFACC | `HandleCycleLargeAlbumSetting` | Known | Event handler |
| 0x0010BEF4 | `HandleSelect` | Known | Event handler |
| 0x0010BF08 | `HandleHilite` | Known | Event handler |
| 0x0010C2A0 | `HandleEQSettingSelected` | Known | Event handler |
| 0x0010C6D0 | `HandleSelect` | Known | Event handler |
| 0x0010C6E4 | `HandleGameHilited` | Known | Event handler |
| 0x0010C994 | `HandleNotesSelected` | Known | Event handler |
| 0x0010C9AC | `HandleNotesPop` | Known | Event handler |
| 0x0010C9BC | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x0011AFBC | `HandleVolumeWheel` | Known | Event handler |
| 0x0011AFD0 | `HandleVolumeChange` | Known | Event handler |
| 0x0011AFE4 | `HandleTimerDone` | Known | Event handler |
| 0x0011AFF4 | `HandleFrequencyChange` | Known | Event handler |
| 0x0011B06C | `HandleTuning` | Known | Event handler |
| 0x0011B07C | `HandleTuningSelect` | Known | Event handler |
| 0x00125BC0 | `HandleLock` | Known | Event handler |
| 0x00125BD0 | `HandleAddressBook` | Known | Event handler |
| 0x001262B8 | `HandleSelect` | Known | Event handler |
| 0x001267F0 | `HandleExit` | Known | Event handler |
| 0x00126800 | `HandleLap` | Known | Event handler |
| 0x0012680C | `HandleResume` | Known | Event handler |
| 0x0012681C | `HandleStartStop` | Known | Event handler |
| 0x00126AD0 | `HandleWheel` | Known | Event handler |
| 0x00126AE0 | `HandlePlayPause` | Known | Event handler |
| 0x00126AF0 | `HandleSelectDown` | Known | Event handler |
| 0x00126B04 | `HandleHilite` | Known | Event handler |
| 0x00126B28 | `HandleFinishRecording` | Known | Event handler |
| 0x001310A4 | `HandleMenuOnKeyboard` | Known | Event handler |
| 0x0013FA18 | `HandleExitUnsupported` | Known | Event handler |
| 0x00156900 | `HandleNotesPop` | Known | Event handler |
| 0x00156914 | `HandleNotesPopToMainMenu` | Known | Event handler |
| 0x00157820 | `HandleSelect` | Known | Event handler |
| 0x00157834 | `HandleWheel` | Known | Event handler |
| 0x00157840 | `HandleImageNext` | Known | Event handler |
| 0x00157850 | `HandleImagePrev` | Known | Event handler |
| 0x00157860 | `HandleImageLast` | Known | Event handler |
| 0x00157870 | `HandleImageFirst` | Known | Event handler |
| 0x00157884 | `HandlePlayPause` | Known | Event handler |
| 0x00157894 | `HandlePlay` | Known | Event handler |
| 0x001578A0 | `HandlePause` | Known | Event handler |
| 0x001578AC | `HandleMikeyCenter` | Known | Event handler |
| 0x0016C8DC | `HandleSelectCity` | Known | Event handler |
| 0x0016C8F4 | `HandleHighlightCity` | Known | Event handler |
| 0x0016D9E0 | `HandleWantPopFlow` | Known | Event handler |
| 0x0016D9F8 | `HandleFlipToAlbumBackside` | Known | Event handler |
| 0x0016DA14 | `HandleFlipToAlbumFrontside` | Known | Event handler |
| 0x0016DA30 | `HandleFlowNext` | Known | Event handler |
| 0x0016DA40 | `HandleFlowPrev` | Known | Event handler |
| 0x0016DA50 | `HandleFlowWheel` | Known | Event handler |
| 0x0016DA60 | `HandleAlbumSelected` | Known | Event handler |
| 0x0016DA74 | `HandlePlayPause` | Known | Event handler |
| 0x0016DA84 | `HandleBacksideSongSelected` | Known | Event handler |
| 0x00199A78 | `HandleLeaveAlarm` | Known | Event handler |
| 0x00199E68 | `HandleSelect` | Known | Event handler |
| 0x0019AD50 | `HandleSelect` | Known | Event handler |
| 0x0019AD64 | `HandleWheel` | Known | Event handler |
| 0x0019AD70 | `HandleImageNext` | Known | Event handler |
| 0x0019AD80 | `HandleImagePrev` | Known | Event handler |
| 0x0019AD90 | `HandleImageLast` | Known | Event handler |
| 0x0019ADA0 | `HandleImageFirst` | Known | Event handler |
| 0x0019ADB4 | `HandlePlayPause` | Known | Event handler |
| 0x0019ADC4 | `HandlePlay` | Known | Event handler |
| 0x0019ADD0 | `HandlePause` | Known | Event handler |
| 0x0019ADDC | `HandleMikeyCenter` | Known | Event handler |
| 0x0019B284 | `HandleNew` | Known | Event handler |
| 0x0019B294 | `HandleClear` | Known | Event handler |
| 0x0019B2A0 | `HandleSelectCurrentSession` | Known | Event handler |
| 0x0019B2BC | `HandleSelectIndexedSession` | Known | Event handler |
| 0x0019B5CC | `HandleWheel` | Known | Event handler |
| 0x0019B5DC | `HandleArrowUp` | Known | Event handler |
| 0x0019B5EC | `HandleArrowDown` | Known | Event handler |
| 0x0019E278 | `HandleHiliteAlbum` | Known | Event handler |
| 0x0019E290 | `HandleBrowseAlbum` | Known | Event handler |
| 0x0019E2A4 | `HandlePlayPause` | Known | Event handler |
| 0x001B4C98 | `HandleSelect` | Known | Event handler |
| 0x001B4E28 | `HandleSelectRegion` | Known | Event handler |
| 0x001B51A0 | `HandleDeleteAllSelect` | Known | Event handler |
| 0x001B51BC | `HandleInsufficientDiskSpace` | Known | Event handler |
| 0x001B51D8 | `HandleInsufficientDiskSpace2` | Known | Event handler |
| 0x001CBF80 | `HandleImageWheel` | Known | Event handler |
| 0x001CBF98 | `HandlePlayPause` | Known | Event handler |
| 0x001CBFA8 | `HandleBrowseLarge` | Known | Event handler |
| 0x001CBFBC | `HandleBrowseSmall` | Known | Event handler |
| 0x001CBFD0 | `HandleBrowseSlideshow` | Known | Event handler |
| 0x001CBFE8 | `HandleImageNext` | Known | Event handler |
| 0x001CBFF8 | `HandleImagePrev` | Known | Event handler |
| 0x001CC008 | `HandleHilite` | Known | Event handler |
| 0x001CC018 | `HandleImageLast` | Known | Event handler |
| 0x001CC028 | `HandleImageFirst` | Known | Event handler |
| 0x001CC03C | `HandleScreenNext` | Known | Event handler |
| 0x001CC050 | `HandleScreenPrev` | Known | Event handler |
| 0x001CE918 | `HandlePlayPause` | Known | Event handler |
| | *...and 1846 more* | | |

---

## 4. Navigation Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00125B00 | `GotoScreen_LockDialog` | Known | Navigation handler |
| 0x00125B18 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation handler |
| 0x00125C90 | `GotoScreen_AddressBook` | Known | Navigation handler |
| 0x0013FDD0 | `GotoScreen_EnterPassKey` | Known | Navigation handler |
| 0x0013FDE8 | `GotoScreen_LockediPod` | Known | Navigation handler |
| 0x001407EC | `GotoScreen_MainMenu` | Known | Navigation handler |
| 0x001DCB38 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation handler |
| 0x001E8000 | `GotoScreen_ConfirmCancelResetMenu` | Known | Navigation handler |
| 0x00203798 | `GotoScreen_VolumeLimitLock_Unlocked` | Known | Navigation handler |
| 0x0020F194 | `GotoScreen_VolumeLimit` | Known | Navigation handler |
| 0x0020F28C | `GotoScreen_SettingsMenu` | Known | Navigation handler |
| 0x0021CB50 | `GotoScreen_Language` | Known | Navigation handler |
| 0x0021E264 | `GotoGeniusLoadingScreen` | Known | Navigation handler |
| 0x00223468 | `GotoGeniusLoadingScreen` | Known | Navigation handler |
| 0x0022693C | `GotoGeniusLoadingScreen` | Known | Navigation handler |
| 0x00226B20 | `GotoGeniusLoadingScreen` | Known | Navigation handler |
| 0x002284AC | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation handler |
| 0x0022873C | `GotoScreen_BacklightTimer` | Known | Navigation handler |
| 0x002287CC | `GotoScreen_VolumeLimit` | Known | Navigation handler |
| 0x002287E4 | `GotoScreen_VolumeLimitLock_Locked` | Known | Navigation handler |
| 0x0022D8F0 | `GotoScreen_LockDialog` | Known | Navigation handler |
| 0x0022D908 | `GotoScreen_SetCombinationFirstTime` | Known | Navigation handler |
| 0x00235D5C | `GotoGeniusMixLoadingScreen` | Known | Navigation handler |
| 0x00238ED8 | `GotoGeniusLoadingScreen` | Known | Navigation handler |
| 0x002390B0 | `GotoGeniusLoadingScreen` | Known | Navigation handler |
| 0x007450BB | `GotoVolumeLimit_or_Lock_Screen` | Known | Navigation handler |

---

## 5. Screen Layouts

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016E4B0 | `CoverFlow_Screen` | Known | Screen layout |
| 0x00738F6A | `Clock_Screen` | Known | Screen layout |
| 0x00738F7A | `Clock_Screen_Default"` | Known | Screen layout |
| 0x00738FDF | `Extras_Screen_WorldClock` | Known | Screen layout |
| 0x0073903D | `Calendar_Menu_Screen` | Known | Screen layout |
| 0x00739055 | `Calendar_Menu_Screen_Default"` | Known | Screen layout |
| 0x007390C2 | `Extras_Screen_Calendar` | Known | Screen layout |
| 0x00739160 | `Extras_Screen_AddressBook` | Known | Screen layout |
| 0x007391BF | `Alarms_Menu_Screen` | Known | Screen layout |
| 0x007391D5 | `Alarms_Menu_Screen_Default"` | Known | Screen layout |
| 0x00739240 | `Extras_Screen_Alarms` | Known | Screen layout |
| 0x0073929A | `Games_Menu_Screen` | Known | Screen layout |
| 0x007392AF | `Games_Menu_Screen_Default"` | Known | Screen layout |
| 0x00739319 | `Extras_Screen_Games` | Known | Screen layout |
| 0x007393D8 | `Extras_Screen_Notes` | Known | Screen layout |
| 0x0073949C | `Extras_Screen_Lock` | Known | Screen layout |
| 0x00739565 | `Extras_Screen_Stopwatch` | Known | Screen layout |
| 0x00739788 | `Clock_Africa_City_Screen ` | Known | Screen layout |
| 0x007397A4 | `Clock_Africa_City_Screen_Default` | Known | Screen layout |
| 0x00739828 | `Clock_Asia_City_Screen` | Known | Screen layout |
| 0x00739842 | `Clock_Asia_City_Screen_Default` | Known | Screen layout |
| 0x007398C4 | `Clock_Atlantic_City_Screen"` | Known | Screen layout |
| 0x007398E2 | `Clock_Atlantic_City_Screen_Default` | Known | Screen layout |
| 0x00739968 | `Clock_Australia_City_Screen#` | Known | Screen layout |
| 0x00739987 | `Clock_Australia_City_Screen_Default` | Known | Screen layout |
| 0x00739A0E | `Clock_Europe_City_Screen ` | Known | Screen layout |
| 0x00739A2A | `Clock_Europe_City_Screen_Default` | Known | Screen layout |
| 0x00739AAE | `Clock_NorthAmerica_City_Screen&` | Known | Screen layout |
| 0x00739AD0 | `Clock_NorthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x00739B5A | `Clock_Pacific_City_Screen!` | Known | Screen layout |
| 0x00739B77 | `Clock_Pacific_City_Screen_Default` | Known | Screen layout |
| 0x00739BFC | `Clock_SouthAmerica_City_Screen&` | Known | Screen layout |
| 0x00739C1E | `Clock_SouthAmerica_City_Screen_Default` | Known | Screen layout |
| 0x00739CAB | `Clock_Screen"` | Known | Screen layout |
| 0x00739D50 | `Clock_Screen"` | Known | Screen layout |
| 0x00739DF5 | `Clock_Screen"` | Known | Screen layout |
| 0x00739E9A | `Clock_Screen"` | Known | Screen layout |
| 0x00739F3F | `Clock_Screen"` | Known | Screen layout |
| 0x00739FE4 | `Clock_Screen"` | Known | Screen layout |
| 0x0073A089 | `Clock_Screen"` | Known | Screen layout |
| 0x0073A12E | `Clock_Screen"` | Known | Screen layout |
| 0x0073A1D3 | `Clock_Screen"` | Known | Screen layout |
| 0x0073A278 | `Clock_Screen"` | Known | Screen layout |
| 0x0073A31D | `Clock_Screen"` | Known | Screen layout |
| 0x0073A3C2 | `Clock_Screen"` | Known | Screen layout |
| 0x0073A467 | `Clock_Screen"` | Known | Screen layout |
| 0x0073A50C | `Clock_Screen"` | Known | Screen layout |
| 0x0073A5B1 | `Clock_Screen"` | Known | Screen layout |
| 0x0073A656 | `Clock_Screen"` | Known | Screen layout |
| 0x0073A6FB | `Clock_Screen"` | Known | Screen layout |
| 0x0073A7A0 | `Clock_Screen"` | Known | Screen layout |
| 0x0073A845 | `Clock_Screen"` | Known | Screen layout |
| 0x0073A8EA | `Clock_Screen"` | Known | Screen layout |
| 0x0073A98F | `Clock_Screen"` | Known | Screen layout |
| 0x0073AA34 | `Clock_Screen"` | Known | Screen layout |
| 0x0073AAD9 | `Clock_Screen"` | Known | Screen layout |
| 0x0073AB7E | `Clock_Screen"` | Known | Screen layout |
| 0x0073AC23 | `Clock_Screen"` | Known | Screen layout |
| 0x0073ACC8 | `Clock_Screen"` | Known | Screen layout |
| 0x0073AD6D | `Clock_Screen"` | Known | Screen layout |
| 0x0073AE12 | `Clock_Screen"` | Known | Screen layout |
| 0x0073AEB7 | `Clock_Screen"` | Known | Screen layout |
| 0x0073AF5C | `Clock_Screen"` | Known | Screen layout |
| 0x0073B001 | `Clock_Screen"` | Known | Screen layout |
| 0x0073B0AB | `Clock_Screen"` | Known | Screen layout |
| 0x0073B150 | `Clock_Screen"` | Known | Screen layout |
| 0x0073B1F5 | `Clock_Screen"` | Known | Screen layout |
| 0x0073B29A | `Clock_Screen"` | Known | Screen layout |
| 0x0073B33F | `Clock_Screen"` | Known | Screen layout |
| 0x0073B3E4 | `Clock_Screen"` | Known | Screen layout |
| 0x0073B489 | `Clock_Screen"` | Known | Screen layout |
| 0x0073B52E | `Clock_Screen"` | Known | Screen layout |
| 0x0073B5D3 | `Clock_Screen"` | Known | Screen layout |
| 0x0073B678 | `Clock_Screen"` | Known | Screen layout |
| 0x0073B71D | `Clock_Screen"` | Known | Screen layout |
| 0x0073B7C2 | `Clock_Screen"` | Known | Screen layout |
| 0x0073B867 | `Clock_Screen"` | Known | Screen layout |
| 0x0073B90C | `Clock_Screen"` | Known | Screen layout |
| 0x0073B9B1 | `Clock_Screen"` | Known | Screen layout |
| 0x0073BA56 | `Clock_Screen"` | Known | Screen layout |
| 0x0073BAFB | `Clock_Screen"` | Known | Screen layout |
| 0x0073BBA0 | `Clock_Screen"` | Known | Screen layout |
| 0x0073BC45 | `Clock_Screen"` | Known | Screen layout |
| 0x0073BCEA | `Clock_Screen"` | Known | Screen layout |
| 0x0073BD8F | `Clock_Screen"` | Known | Screen layout |
| 0x0073BE34 | `Clock_Screen"` | Known | Screen layout |
| 0x0073BED9 | `Clock_Screen"` | Known | Screen layout |
| 0x0073BF7E | `Clock_Screen"` | Known | Screen layout |
| 0x0073C023 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C0C8 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C16D | `Clock_Screen"` | Known | Screen layout |
| 0x0073C212 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C2B7 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C35C | `Clock_Screen"` | Known | Screen layout |
| 0x0073C401 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C4A6 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C54B | `Clock_Screen"` | Known | Screen layout |
| 0x0073C5F0 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C695 | `Clock_Screen"` | Known | Screen layout |
| 0x0073C73A | `Clock_Screen"` | Known | Screen layout |
| | *...and 6634 more* | | |

---

## 6. Settings (Toggle/Show)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0013F7E4 | `ShowSetting_EQ` | Known | User setting |
| 0x001E9D0C | `ToggleSetting_Repeat` | Known | User setting |
| 0x001E9D28 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x001E9D40 | `ToggleSetting_TVOut` | Known | User setting |
| 0x001E9D54 | `ToggleSetting_TVSignal` | Known | User setting |
| 0x00212D30 | `ShowSetting_Backlight` | Known | User setting |
| 0x00227EF0 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x00227F0C | `ToggleSetting_Repeat` | Known | User setting |
| 0x00227F24 | `ToggleSetting_SortBy` | Known | User setting |
| 0x00227F3C | `ToggleSetting_ClassicUI` | Known | User setting |
| 0x00227F54 | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x00227F70 | `ToggleSetting_Clicker` | Known | User setting |
| 0x00227F88 | `ToggleSetting_DaylightSavings` | Known | User setting |
| 0x00227FA8 | `ToggleSetting_24HourClock` | Known | User setting |
| 0x00227FC4 | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x00227FE0 | `ShowSetting_Shuffle` | Known | User setting |
| 0x0022818C | `ShowSetting_Repeat` | Known | User setting |
| 0x002281A0 | `ShowSetting_About` | Known | User setting |
| 0x002281B4 | `ShowSetting_MainMenu` | Known | User setting |
| 0x002281CC | `ShowSetting_MusicMenu` | Known | User setting |
| 0x002281E4 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x002281FC | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x00228218 | `ShowSetting_Brightness` | Known | User setting |
| 0x00228230 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x00228248 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x00228264 | `ShowSetting_EQ` | Known | User setting |
| 0x00228274 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x00228410 | `ShowSetting_Clicker` | Known | User setting |
| 0x00228424 | `ShowSetting_DateAndTime` | Known | User setting |
| 0x0022843C | `ShowSetting_SortBy` | Known | User setting |
| 0x00228450 | `ShowSetting_ClassicUI` | Known | User setting |
| 0x00228468 | `ShowSetting_Language` | Known | User setting |
| 0x00228480 | `ShowSetting_Legal` | Known | User setting |
| 0x00228494 | `ShowSetting_ResetAll` | Known | User setting |
| 0x007423ED | `ToggleSetting_24HourClock` | Known | User setting |
| 0x0074249D | `ToggleSetting_TimeInTitle` | Known | User setting |
| 0x00744C86 | `ShowSetting_About` | Known | User setting |
| 0x00744D28 | `ToggleSetting_Shuffle` | Known | User setting |
| 0x00744D6C | `ShowSetting_Shuffle` | Known | User setting |
| 0x00744DE3 | `ToggleSetting_Repeat` | Known | User setting |
| 0x00744E26 | `ShowSetting_Repeat` | Known | User setting |
| 0x00744F30 | `ShowSetting_MainMenu` | Known | User setting |
| 0x00745040 | `ShowSetting_MusicMenu` | Known | User setting |
| 0x00745108 | `ShowSetting_VolumeLimit` | Known | User setting |
| 0x007451D2 | `ShowSetting_BacklightTimer` | Known | User setting |
| 0x007452EA | `ShowSetting_Brightness` | Known | User setting |
| 0x00745420 | `ShowSetting_Audiobooks` | Known | User setting |
| 0x00745531 | `ShowSetting_RadioRegions` | Known | User setting |
| 0x00745632 | `ShowSetting_EQ` | Known | User setting |
| 0x0074569F | `ToggleSetting_SoundCheck` | Known | User setting |
| 0x007456E6 | `ShowSetting_SoundCheck` | Known | User setting |
| 0x00745763 | `ToggleSetting_Clicker` | Known | User setting |
| 0x007457A7 | `ShowSetting_Clicker` | Known | User setting |
| 0x0074590E | `ToggleSetting_SortBy` | Known | User setting |
| 0x00745951 | `ShowSetting_SortBy` | Known | User setting |
| 0x00745A52 | `ShowSetting_Language` | Known | User setting |
| 0x00745B62 | `ShowSetting_Legal` | Known | User setting |
| 0x00745C93 | `ShowSetting_ResetAll` | Known | User setting |
| 0x00745E04 | `ShowSetting_Backlight` | Known | User setting |
| 0x00745EB4 | `ShowSetting_Backlight` | Known | User setting |
| 0x00745F64 | `ShowSetting_Backlight` | Known | User setting |
| 0x00746015 | `ShowSetting_Backlight` | Known | User setting |
| 0x007460C6 | `ShowSetting_Backlight` | Known | User setting |
| 0x00746177 | `ShowSetting_Backlight` | Known | User setting |
| 0x0074622B | `ShowSetting_Backlight` | Known | User setting |
| 0x007462DA | `ShowSetting_EQ` | Known | User setting |
| 0x0074634F | `ShowSetting_Language` | Known | User setting |
| 0x007D9D04 | `ToggleSetting_Repeat` | Known | User setting |
| 0x007D9D3E | `ToggleSetting_Shuffle` | Known | User setting |
| 0x007D9E00 | `ToggleSetting_TVOut` | Known | User setting |
| 0x007D9E39 | `ToggleSetting_TVSignal` | Known | User setting |

---

## 7. RTOS Tasks

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00009107 | `"RTXC v3.2b for ARM and Thumb - ARM ADS 1.0.1 Nov-17-00 Key: 23971` | Known | RTOS task thread |
| 0x000E8FBC | `HostOSTask` | Known | RTOS task thread |
| 0x00149868 | `USBDeviceTask` | Known | RTOS task thread |
| 0x00153B9C | `DiskReaderTask` | Known | RTOS task thread |
| 0x00163D64 | `ATAWorkLoopTask` | Known | RTOS task thread |
| 0x00163D78 | `ATAWorkLoopIRQTask` | Known | RTOS task thread |
| 0x0019CA7C | `GeniusMixesTask` | Known | RTOS task thread |
| 0x001B99F8 | `TMusicLoadingTask` | Known | RTOS task thread |
| 0x001F5470 | `MeCCAIOTask` | Known | RTOS task thread |
| 0x00228C68 | `StreamCacheReadTask` | Known | RTOS task thread |
| 0x002AD8AC | `FirewireTask` | Known | RTOS task thread |
| 0x002AD8C0 | `TouchwheelTask` | Known | RTOS task thread |
| 0x002AD8D4 | `AudioOutStateTask` | Known | RTOS task thread |
| 0x002AD900 | `DiskMgrTask` | Known | RTOS task thread |
| 0x002AD910 | `HoldSwitchTask` | Known | RTOS task thread |
| 0x002AD924 | `MikeyTask` | Known | RTOS task thread |
| 0x002AD934 | `TopPlugTask` | Known | RTOS task thread |
| 0x002AD944 | `HPhoneDetTask` | Known | RTOS task thread |
| 0x002AD9BC | `LowBattDebounceTask` | Known | RTOS task thread |
| 0x002AD9E4 | `AlarmTask` | Known | RTOS task thread |
| 0x002ADA03 | `"USBAudioTask` | Known | RTOS task thread |
| 0x002BF57D | `** Clock Snapshot **` | Known | RTOS task thread |
| 0x002BFA34 | `  K - RTXC` | Known | RTOS task thread |
| 0x002BFB99 | `** Mailbox Snapshot **` | Known | RTOS task thread |
| 0x002BFDDD | `** Queue Snapshot **` | Known | RTOS task thread |
| 0x002C0040 | `** Task Register Snapshot **` | Known | RTOS task thread |
| 0x002C00D4 | `Undefined Task` | Known | RTOS task thread |
| 0x002C0215 | `** Resource Snapshot **` | Known | RTOS task thread |
| 0x002C066D | `** Semaphore Snapshot **` | Known | RTOS task thread |
| 0x002C09A5 | `** Stack Snapshot **` | Known | RTOS task thread |
| 0x002C0A3C | `(/H'RTXC Kernel 'X2,R16,Z-,UL8,Z+,R10,3(UI6),//N)` | Known | RTOS task thread |
| 0x002C0CB9 | `** Task Snapshot **` | Known | RTOS task thread |
| 0x003F49DC | `ArtworkLoadTask` | Known | RTOS task thread |
| 0x003F80A8 | `TPodMediaPlayer Task` | Known | RTOS task thread |
| 0x004007B4 | `MeCCARecordingTask` | Known | RTOS task thread |

---

## 8. Logging Channels

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002650C0 | `Channel Reserved` | Known | Logging channel |
| 0x002650D4 | `Channel AppBoot` | Known | Logging channel |
| 0x002650E4 | `Channel BufferedSongReading` | Known | Logging channel |
| 0x00265100 | `Channel PrefsWriting` | Known | Logging channel |
| 0x00265118 | `Channel GeneralUserExperience` | Known | Logging channel |
| 0x00265138 | `Channel PlayFromDisk` | Known | Logging channel |
| 0x00265150 | `Channel CacheSpinupDrive` | Known | Logging channel |
| 0x0026516C | `Channel TestLogging` | Known | Logging channel |
| 0x00265180 | `Channel AppFileLoading` | Known | Logging channel |
| 0x00265198 | `Channel VCardReading` | Known | Logging channel |
| 0x002651B0 | `Channel LongSongScanning` | Known | Logging channel |
| 0x00265224 | `Channel VoiceRecording` | Known | Logging channel |
| 0x0026523C | `Channel PhotoImporting` | Known | Logging channel |
| 0x00265254 | `Channel Notes` | Known | Logging channel |
| 0x00265264 | `Channel PhotoFileManagement` | Known | Logging channel |
| 0x00265280 | `Channel DiskMode` | Known | Logging channel |
| 0x00265294 | `Channel Firewire` | Known | Logging channel |
| 0x002652A8 | `Channel USB` | Known | Logging channel |
| 0x002652C8 | `Channel FreeSpaceCache` | Known | Logging channel |
| 0x002652E0 | `Channel OnTheGoFileMgmt` | Known | Logging channel |

---

## 9. Audio System (MeCCA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000067EF | `"MeCCADecode` | Known | Audio system |
| 0x0013F7F8 | `HandleEQ` | Known | Audio system |
| 0x00150EDC | `AudioCodecs` | Known | Audio system |
| 0x001524FC | `VideoCodecs` | Known | Audio system |
| 0x00168A34 | `MeCCA_RecordingBuffer` | Known | Audio system |
| 0x00197BE0 | `MeCCA_PCM_Output.wav` | Known | Audio system |
| 0x001B1678 | `MeCCA_MediaPlayer` | Known | Audio system |
| 0x001BC3C8 | `MeCCA_VideoBufferMgr` | Known | Audio system |
| 0x001BC5D0 | `MeCCAVideoDecode` | Known | Audio system |
| 0x007455F1 | `Settings_EQMenu_Layout"` | Known | Audio system |
| 0x00745666 | `SettingsMenus_EQ_Layout` | Known | Audio system |
| 0x007462AA | `HandleEQ` | Known | Audio system |
| 0x007FC75C | `ACELP is either registered trademark or trademark of VoiceAge Corporation in the` | Known | Audio system |
| 0x00917268 | `MeCCA_StreamCache` | Known | Audio system |
| 0x0092B578 | `ERROR: unknownCodec loaded !!!` | Known | Audio system |
| 0x0098845F | `11TCEQSetting` | Known | Audio system |
| 0x00988624 | `13TCSettings_EQ` | Known | Audio system |
| 0x009892C1 | `27TSilverCntlrTransitionAddonI11TCEQSettingE` | Known | Audio system |
| 0x009894E4 | `27TSilverCntlrTransitionAddonI13TCSettings_EQE` | Known | Audio system |
| 0x009B67C9 | `SettingsMenu_ListItem_EQ` | Known | Audio system |
| 0x009B8087 | `Settings_EQ_RandB_Image` | Known | Audio system |
| 0x009B8149 | `Settings_EQ_Electronic_Image` | Known | Audio system |
| 0x009B8180 | `Settings_EQ_Acoustic_Image` | Known | Audio system |
| 0x009B84F1 | `Settings_EQ_SpokenWord_Image` | Known | Audio system |
| 0x009B856E | `Settings_EQ_Dance_Image` | Known | Audio system |
| 0x009B85DC | `Settings_EQ_Lounge_Image` | Known | Audio system |
| 0x009B8BB6 | `Settings_EQ_Rock_Image` | Known | Audio system |
| 0x009B8C3B | `Settings_EQ_Classical_Image` | Known | Audio system |
| 0x009B8FB0 | `Settings_EQ_Latin_Image` | Known | Audio system |
| 0x009B91A8 | `Settings_EQ_Piano_Image` | Known | Audio system |
| 0x009B95E5 | `Settings_EQ_Deep_Image` | Known | Audio system |
| 0x009B9642 | `Settings_EQ_HipHop_Image` | Known | Audio system |
| 0x009B965B | `Settings_EQ_Pop_Image` | Known | Audio system |
| 0x009B97F8 | `Settings_EQ_TrebleReducer_Image` | Known | Audio system |
| 0x009B9818 | `Settings_EQ_BassReducer_Image` | Known | Audio system |
| 0x009B9AB4 | `Settings_EQ_TrebleBooster_Image` | Known | Audio system |
| 0x009B9AD4 | `Settings_EQ_VocalBooster_Image` | Known | Audio system |
| 0x009B9AF3 | `Settings_EQ_BassBooster_Image` | Known | Audio system |
| 0x009B9C1A | `Settings_EQ_SmallSpeakers_Image` | Known | Audio system |
| 0x009B9C5E | `Settings_EQ_Loudness_Image` | Known | Audio system |
| 0x009B9D23 | `Settings_EQ_Flat_Image` | Known | Audio system |
| 0x009BA745 | `Settings_EQ_Jazz_Image` | Known | Audio system |
| 0x009BB3F2 | `SettingsEQ_Template` | Known | Audio system |
| 0x009BCE1B | `SettingsMenu_EQ_String` | Known | Audio system |
| 0x009CB9DD | `SettingsMenus_EQ_Layout` | Known | Audio system |
| 0x009CBC8C | `SettingsEQ_Template_Layout` | Known | Audio system |
| 0x009CC7A8 | `Settings_EQMenu_Layout` | Known | Audio system |
| 0x009D289E | `msCodeCom` | Known | Audio system |

---

## 10. Audio/Codec - Audible

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00150FB8 | `Audible` | Known | Audible audiobook format |
| 0x007FC8A0 | `The Audible software in this product is used under license from Audible. Copyrig` | Known | Audible audiobook format |
| 0x007FC8F5 | ` 2002 by Audible, Inc. All rights reserved.` | Known | Audible audiobook format |
| 0x00801740 | `, %d Audibles` | Known | Audible audiobook format |
| 0x00801750 | `, 1 Audible` | Known | Audible audiobook format |
| 0x0083F56C | `, %d Audible` | Known | Audible audiobook format |
| 0x0083F57C | `, 1 Audible` | Known | Audible audiobook format |
| 0x00848CD8 | `, %d Audibles` | Known | Audible audiobook format |
| 0x00848CE8 | `, 1 Audible` | Known | Audible audiobook format |
| 0x008524C4 | `, %d Audibles` | Known | Audible audiobook format |
| 0x008524D4 | `, 1 Audible` | Known | Audible audiobook format |
| 0x0086A650 | `, %d Audibles` | Known | Audible audiobook format |
| 0x0086A660 | `, 1 Audible` | Known | Audible audiobook format |
| 0x00873D94 | `, %d Audiblea` | Known | Audible audiobook format |
| 0x00873DA4 | `, 1 Audible` | Known | Audible audiobook format |
| 0x0089B8F0 | ` Audible` | Known | Audible audiobook format |
| 0x0089B907 | ` Audible` | Known | Audible audiobook format |
| 0x008A6ABA | ` Audible` | Known | Audible audiobook format |
| 0x008A6ACD | ` Audible` | Known | Audible audiobook format |
| 0x008BA100 | `, %d Audible` | Known | Audible audiobook format |
| 0x008BA110 | `, 1 Audible` | Known | Audible audiobook format |
| 0x008C358C | `, %d Audible` | Known | Audible audiobook format |
| 0x008C359C | `, 1 Audible` | Known | Audible audiobook format |
| 0x008CD0EC | `, %d Audibles` | Known | Audible audiobook format |
| 0x008CD0FC | `, 1 Audible` | Known | Audible audiobook format |
| 0x008D8248 | `, %d Audibles` | Known | Audible audiobook format |
| 0x008D8258 | `, 1 Audible` | Known | Audible audiobook format |
| 0x008E4174 | `, %d Audible` | Known | Audible audiobook format |
| 0x008E4184 | `, 1 Audible` | Known | Audible audiobook format |
| 0x008ED7A4 | `, %d Audible` | Known | Audible audiobook format |
| 0x008ED7B4 | `, 1 Audible` | Known | Audible audiobook format |
| 0x008F71B5 | ` Audible` | Known | Audible audiobook format |
| 0x008F71C8 | ` Audible` | Known | Audible audiobook format |
| 0x00900935 | ` Audible` | Known | Audible audiobook format |
| 0x00900948 | ` Audible` | Known | Audible audiobook format |

---

## 11. Audio/Codec - Apple Lossless

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000A6ACF | `"alac: bit depth = %d, pb = 0x%X, mb = 0x%X, kb = 0x%X ` | Known | Apple Lossless codec |
| 0x00150F8C | `AppleLossless` | Known | Apple Lossless codec |
| 0x0015592C | `alacmp4v@KL` | Known | Apple Lossless codec |
| 0x001C3FD0 | `elsttkhdmdhdstsdsttsstszstscstcomp4aalac` | Known | Apple Lossless codec |
| 0x008F0A69 | ` geri alacakt` | Known | Apple Lossless codec |
| 0x008F0ADE | ` geri alacakt` | Known | Apple Lossless codec |

---

## 12. Audio/Codec - AAC

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0049CEC0 | `!"#$%%#!&%"'##&%()))))))(()())))*+,*+++++**--.+*///0//00/000/00/1221113111411551` | Known | AAC codec |
| 0x004DEEB8 | `#$%$$%$$$$$%$$$%!&$!$$%$$$$%$!%$%$%%$$%$#'())((()(())((((((((()(((()((())((()(()` | Known | AAC codec |
| 0x005BEB7F | `AAAAAAAAAAC` | Known | AAC codec |
| 0x006BF413 | `B+A22AAAC` | Known | AAC codec |

---

## 13. Audio/Codec - MP3

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x007FC700 | `MPEG Layer-3 audio coding technology licensed from Fraunhofer IIS and THOMSON mu` | Known | MP3 codec |

---

## 14. DRM/Security

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000A8FBD | `;=1sinf` | Known | DRM system |
| 0x00150EB0 | `AppleDRMVersion` | Known | DRM system |
| 0x00150F50 | `AppleDRM` | Known | DRM system |
| 0x00152510 | `AppleVideoDRM` | Known | DRM system |
| 0x00155910 | `tx3gdrmsp608aavdmp4aesdsX{` | Known | DRM system |
| 0x001C26A0 | `tkhdedtselstmdiamdhdminfstblstsdstcoco64stscstszsttsstssdrmidrms` | Known | DRM system |
| 0x008BAFD8 | `Ingen enhetsinfo tilgjengelig.` | Known | DRM system |
| 0x008E50DF | `rsinformation finns tillg` | Known | DRM system |
| 0x008E6B39 | `ningsinformationens riktighet.` | Known | DRM system |
| 0x009B487F | `DRMLevel` | Known | DRM system |
| 0x009DA728 | `$Apple FairPlay Certificate Authority0` | Known | DRM system |
| 0x009DAAAD | `&Apple FairPlay Certification Authority0` | Known | DRM system |
| 0x00A0F655 | `&Apple FairPlay Certification Authority0` | Known | DRM system |
| 0x00A0F6CB | `Apple FairPlay1402` | Known | DRM system |

---

## 15. Game System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00098648 | `gamedata_RW` | Known | Game system |
| 0x00098664 | `gamedata_ShareRW` | Known | Game system |
| 0x00098678 | `games_RO` | Known | Game system |
| 0x00125F28 | `StartGame` | Known | Game system |
| 0x0075CAEA | `controller.StartGame1` | Known | Game system |
| 0x0075CCCE | `controller.StartGame1` | Known | Game system |
| 0x0075CEB2 | `controller.StartGame1` | Known | Game system |
| 0x0075D096 | `controller.StartGame1` | Known | Game system |
| 0x0075D27A | `controller.StartGame1` | Known | Game system |
| 0x0075D644 | `controller.StartGame1` | Known | Game system |
| 0x0098847B | `11TCGamesMenu` | Known | Game system |
| 0x0098854F | `12TCGameScreen` | Known | Game system |
| 0x0098931B | `27TSilverCntlrTransitionAddonI11TCGamesMenuE` | Known | Game system |
| 0x009893D0 | `27TSilverCntlrTransitionAddonI12TCGameScreenE` | Known | Game system |
| 0x009BFFFA | `GamesMenu_StatusBar_String` | Known | Game system |
| 0x009C0782 | `AboutScreen_Games_String` | Known | Game system |
| 0x009C77E9 | `MainMenu_List_Games` | Known | Game system |
| 0x009C77FD | `ExtrasMenu_Games` | Known | Game system |
| 0x009CF68F | `MainMenu_List_Games_x` | Known | Game system |

---

## 16. Photo System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000690CC | ` rtSPhotos\Photo Database` | Known | Photo system |
| 0x001040A8 | `TPhotosBrowseCntlr` | Known | Photo system |
| 0x001040C4 | `TPhotosBrowseTransitionCntlr` | Known | Photo system |
| 0x001040EC | `TPhotosMenuCntlr` | Known | Photo system |
| 0x00104108 | `TPhotosSettingsCntlr` | Known | Photo system |
| 0x00104128 | `TPhotosSettingsCntlr_Transitions` | Known | Photo system |
| 0x00104154 | `TPhotosSettingsCntlr_Duration` | Known | Photo system |
| 0x0010417C | `TPhotosSettingsSlideshowPlaylistCntlr` | Known | Photo system |
| 0x001CDEAC | `PhotoBrowse/Slideshow` | Known | Photo system |
| 0x00267038 | `Photo Database Size` | Known | Photo system |
| 0x00401CC8 | `TPhotosBrowseCntlr` | Known | Photo system |
| 0x00401CDC | `TPhotosBrowseTransitionCntlr` | Known | Photo system |
| 0x00401CFC | `TPhotosMenuCntlr` | Known | Photo system |
| 0x00401D10 | `TPhotosSettingsCntlr` | Known | Photo system |
| 0x00401D28 | `TPhotosSettingsCntlr_Transitions` | Known | Photo system |
| 0x00401D4C | `TPhotosSettingsCntlr_Duration` | Known | Photo system |
| 0x00401F98 | `TPhotosSettingsSlideshowPlaylistCntlr` | Known | Photo system |
| 0x00744C13 | `Photos_SettingsMenu` | Known | Photo system |
| 0x00759614 | `TPhotosMenuCntlrTSilverCntlrTPhotosSettingsCntlrTPhotosSettingsCntlr_Transitions` | Known | Photo system |
| 0x00759684 | `TPhotosSettingsSlideshowPlaylistCntlr` | Known | Photo system |
| 0x007596AC | `TPhotosBrowseCntlr` | Known | Photo system |
| 0x007619D3 | `Photos_Menu` | Known | Photo system |
| 0x00763A9F | `Photos_Menu` | Known | Photo system |
| 0x00765B6B | `Photos_Menu` | Known | Photo system |
| 0x00767C37 | `Photos_Menu` | Known | Photo system |
| 0x00769D03 | `Photos_Menu` | Known | Photo system |
| 0x0076BDCF | `Photos_Menu` | Known | Photo system |
| 0x0076DE9B | `Photos_Menu` | Known | Photo system |
| 0x0076FF67 | `Photos_Menu` | Known | Photo system |
| 0x00772033 | `Photos_Menu` | Known | Photo system |
| 0x007740FF | `Photos_Menu` | Known | Photo system |
| 0x007761CB | `Photos_Menu` | Known | Photo system |
| 0x00778297 | `Photos_Menu` | Known | Photo system |
| 0x0077A363 | `Photos_Menu` | Known | Photo system |
| 0x0077C42F | `Photos_Menu` | Known | Photo system |
| 0x0077E4FB | `Photos_Menu` | Known | Photo system |
| 0x007805C7 | `Photos_Menu` | Known | Photo system |
| 0x00782693 | `Photos_Menu` | Known | Photo system |
| 0x0078475F | `Photos_Menu` | Known | Photo system |
| 0x0078682B | `Photos_Menu` | Known | Photo system |
| 0x007888F7 | `Photos_Menu` | Known | Photo system |
| 0x0078A9C3 | `Photos_Menu` | Known | Photo system |
| 0x0078CA8F | `Photos_Menu` | Known | Photo system |
| 0x0078EB5B | `Photos_Menu` | Known | Photo system |
| 0x00790C27 | `Photos_Menu` | Known | Photo system |
| 0x00792CF3 | `Photos_Menu` | Known | Photo system |
| 0x00794DBF | `Photos_Menu` | Known | Photo system |
| 0x00796E8B | `Photos_Menu` | Known | Photo system |
| 0x00798F57 | `Photos_Menu` | Known | Photo system |
| 0x0079B023 | `Photos_Menu` | Known | Photo system |
| 0x0079D0EF | `Photos_Menu` | Known | Photo system |
| 0x0079F1BB | `Photos_Menu` | Known | Photo system |
| 0x007A1287 | `Photos_Menu` | Known | Photo system |
| 0x007A3353 | `Photos_Menu` | Known | Photo system |
| 0x007A541F | `Photos_Menu` | Known | Photo system |
| 0x007A74EB | `Photos_Menu` | Known | Photo system |
| 0x007A95B7 | `Photos_Menu` | Known | Photo system |
| 0x007AB683 | `Photos_Menu` | Known | Photo system |
| 0x007AD74F | `Photos_Menu` | Known | Photo system |
| 0x007D9A43 | `PhotoBrowse_Small` | Known | Photo system |
| 0x007D9AEC | `Photos_SettingsMenu` | Known | Photo system |
| 0x007D9B6C | `PhotoBrowse_Small"` | Known | Photo system |
| 0x007D9C2F | `Photos_SettingsDurationMenu` | Known | Photo system |
| 0x007D9DC1 | `Photos_SettingsTransitionMenu` | Known | Photo system |
| 0x00803560 | `PhotoBrowse_Large` | Known | Photo system |
| 0x0090B674 | `TPhotosSettingsCntlr` | Known | Photo system |
| 0x0090B68C | `TPhotosSettingsCntlr_Transitions` | Known | Photo system |
| 0x0090B6B0 | `TPhotosSettingsCntlr_Duration` | Known | Photo system |
| 0x00988900 | `16TPhotosMenuCntlr` | Known | Photo system |
| 0x00988B4B | `18TPhotosBrowseCntlr` | Known | Photo system |
| 0x00988CC4 | `20TPhotosSettingsCntlr` | Known | Photo system |
| 0x00989850 | `27TSilverCntlrTransitionAddonI16TPhotosMenuCntlrE` | Known | Photo system |
| 0x009899B6 | `27TSilverCntlrTransitionAddonI18TPhotosBrowseCntlrE` | Known | Photo system |
| 0x00989C01 | `27TSilverCntlrTransitionAddonI20TPhotosSettingsCntlrE` | Known | Photo system |
| 0x0098A2D6 | `27TSilverCntlrTransitionAddonI28TPhotosBrowseTransitionCntlrE` | Known | Photo system |
| 0x0098A4C8 | `27TSilverCntlrTransitionAddonI29TPhotosSettingsCntlr_DurationE` | Known | Photo system |
| 0x0098A8C7 | `27TSilverCntlrTransitionAddonI32TPhotosSettingsCntlr_TransitionsE` | Known | Photo system |
| 0x0098AB64 | `27TSilverCntlrTransitionAddonI37TPhotosSettingsSlideshowPlaylistCntlrE` | Known | Photo system |
| 0x0098AE29 | `28TPhotosBrowseTransitionCntlr` | Known | Photo system |
| 0x0098AF23 | `29TPhotosSettingsCntlr_Duration` | Known | Photo system |
| 0x0098B153 | `32TPhotosSettingsCntlr_Transitions` | Known | Photo system |
| 0x0098B32D | `37TPhotosSettingsSlideshowPlaylistCntlr` | Known | Photo system |
| 0x009B76EF | `PhotoBrowse_Grid` | Known | Photo system |
| 0x009B782F | `Settings_Capacity_PhotosLegend` | Known | Photo system |
| 0x009B8354 | `Settings_About_Capacity_PhotosLegend_Image` | Known | Photo system |
| 0x009BA898 | `PhotoBrowse_Large` | Known | Photo system |
| 0x009BB83C | `PhotoBrowse_Template` | Known | Photo system |
| 0x009BD27E | `Photos_Settings_Music_String` | Known | Photo system |
| 0x009BDA27 | `Photos_Import_Browse_Choice_String` | Known | Photo system |
| 0x009BDD74 | `Photos_Settings_Time_Per_Slide_String` | Known | Photo system |
| 0x009BE348 | `Photos_All_Photos_Browse_String` | Known | Photo system |
| 0x009BE6E8 | `Photos_Settings_Music_Off_String` | Known | Photo system |
| 0x009BEC4F | `Photos_Settings_Music_NowPlaying_String` | Known | Photo system |
| 0x009BEFA1 | `Photos_Settings_TV_Signal_String` | Known | Photo system |
| 0x009BFE67 | `Photos_Browse_1_Photo_String` | Known | Photo system |
| 0x009BFEB6 | `Photos_Settings_Music_FromiPhoto_String` | Known | Photo system |
| 0x009C08B9 | `Photos_Settings_String` | Known | Photo system |
| 0x009C0D87 | `Photos_Settings_Transitions_String` | Known | Photo system |
| 0x009C0E84 | `Photos_Settings_Shuffle_Photos_String` | Known | Photo system |
| 0x009C0EDD | `MainMenu_Photos_String` | Known | Photo system |
| | *...and 33 more* | | |

---

## 17. Video System

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x009C5B9F | `NowPlaying_RentalWarning_Dialog_Template_Video` | Known | Video system |
| 0x009C5BF2 | `NowPlaying_RentalWarning_Overlay_Template_Video` | Known | Video system |
| 0x009C7D28 | `MainMenu_Video_List_Rentals` | Known | Video system |
| 0x009CF6BB | `MainMenu_Video_List_Rentals_x` | Known | Video system |

---

## 18. Genius

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000AC720 | `TGeniusLoadingCntlr` | Known | Genius system |
| 0x000B5F20 | `GeniusPlaylist` | Known | Genius system |
| 0x00113BD0 | `GeniusPlaylist_` | Known | Genius system |
| 0x00152430 | `SupportsGenius` | Known | Genius system |
| 0x00152440 | `GeniusConfigMinVersion` | Known | Genius system |
| 0x00152458 | `GeniusMetadataMinVersion` | Known | Genius system |
| 0x00152474 | `GeniusSimilaritiesMinVersion` | Known | Genius system |
| 0x00152494 | `GeniusConfigMaxVersion` | Known | Genius system |
| 0x001524AC | `GeniusMetadataMaxVersion` | Known | Genius system |
| 0x001524C8 | `GeniusSimilaritiesMaxVersion` | Known | Genius system |
| 0x001524E8 | `SupportsGeniusMixes` | Known | Genius system |
| 0x001B5E34 | `GeniusMixArtwork` | Known | Genius system |
| 0x001DED04 | `RefreshingGenius` | Known | Genius system |
| 0x001DED1C | `CreatingGeniusMix` | Known | Genius system |
| 0x001DF008 | `GeniusPlaylistReady` | Known | Genius system |
| 0x001DF01C | `GeniusMixPlaylistReady` | Known | Genius system |
| 0x0021E0C0 | `GotoGeniusLayout` | Known | Genius system |
| 0x0021E248 | `GotoGeniusError_NoGenius` | Known | Genius system |
| 0x0021E27C | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Genius system |
| 0x0021E2A4 | `StartGenius` | Known | Genius system |
| 0x0021F290 | `GotoGeniusError_NoGenius` | Known | Genius system |
| 0x0021F2AC | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Genius system |
| 0x0021F764 | `StartGenius` | Known | Genius system |
| 0x00223454 | `StartGenius` | Known | Genius system |
| 0x0022367C | `StartGenius` | Known | Genius system |
| 0x00226920 | `GotoGeniusError_NoGenius` | Known | Genius system |
| 0x00226954 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Genius system |
| 0x0022697C | `StartGenius` | Known | Genius system |
| 0x00226AFC | `StartGenius` | Known | Genius system |
| 0x00226B0C | `RefreshingGenius` | Known | Genius system |
| 0x00226DC4 | `StartGenius` | Known | Genius system |
| 0x0022F9B4 | `GotoGeniusMixesIntro` | Known | Genius system |
| 0x0022F9D0 | `GotoGeniusMixes` | Known | Genius system |
| 0x0022F9E0 | `GotoSingleGeniusMix` | Known | Genius system |
| 0x00231250 | `StartGenius` | Known | Genius system |
| 0x002324F4 | `StartGenius` | Known | Genius system |
| 0x00232D00 | `StartGenius` | Known | Genius system |
| 0x00232D18 | `GotoGenius` | Known | Genius system |
| 0x00232D30 | `SavedGeniusPlaylist` | Known | Genius system |
| 0x00232FA0 | `SavedGeniusPlaylist` | Known | Genius system |
| 0x0023319C | `GotoGeniusIntro` | Known | Genius system |
| 0x002331B0 | `GotoGenius` | Known | Genius system |
| 0x002331F0 | `GeniusPlaylistSelected` | Known | Genius system |
| 0x00235D44 | `CreatingGeniusMix` | Known | Genius system |
| 0x00238EBC | `GotoGeniusError_NoGenius` | Known | Genius system |
| 0x00238EF0 | `GotoGeniusError_NoGeniusInfoForTrack` | Known | Genius system |
| 0x00238F18 | `StartGenius` | Known | Genius system |
| 0x00239098 | `RefreshingGenius` | Known | Genius system |
| 0x002394F0 | `StartGenius` | Known | Genius system |
| 0x00239880 | `GeniusPlaylistSelected` | Known | Genius system |
| 0x00402698 | `TGeniusLoadingCntlr` | Known | Genius system |
| 0x00758CE4 | `TContextualMenuCntlrTCExtrasMenuTSilverCntlrTGeniusLoadingCntlr` | Known | Genius system |
| 0x007601F1 | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x00760CE6 | `controller.GotoGenius1` | Known | Genius system |
| 0x00760D64 | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x00760DDD | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x00760E6A | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x007611E9 | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x007622BD | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x00762DB2 | `controller.GotoGenius1` | Known | Genius system |
| 0x00762E30 | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x00762EA9 | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x00762F36 | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x007632B5 | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x00764389 | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x00764E7E | `controller.GotoGenius1` | Known | Genius system |
| 0x00764EFC | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x00764F75 | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x00765002 | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x00765381 | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x00766455 | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x00766F4A | `controller.GotoGenius1` | Known | Genius system |
| 0x00766FC8 | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x00767041 | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x007670CE | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x0076744D | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x00768521 | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x00769016 | `controller.GotoGenius1` | Known | Genius system |
| 0x00769094 | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x0076910D | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x0076919A | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x00769519 | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x0076A5ED | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x0076B0E2 | `controller.GotoGenius1` | Known | Genius system |
| 0x0076B160 | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x0076B1D9 | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x0076B266 | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x0076B5E5 | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x0076C6B9 | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x0076D1AE | `controller.GotoGenius1` | Known | Genius system |
| 0x0076D22C | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x0076D2A5 | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x0076D332 | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x0076D6B1 | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| 0x0076E785 | `controller.EmptyGeniusHilited1` | Known | Genius system |
| 0x0076F27A | `controller.GotoGenius1` | Known | Genius system |
| 0x0076F2F8 | `controller.GotoGeniusIntro1` | Known | Genius system |
| 0x0076F371 | `controller.GotoGeniusMixes1` | Known | Genius system |
| 0x0076F3FE | `controller.GotoGeniusMixesIntro1` | Known | Genius system |
| 0x0076F77D | `controller.GotoSingleGeniusMix1` | Known | Genius system |
| | *...and 821 more* | | |

---

## 19. Database (SQLite)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00237040 | `%s/sqlite_` | Known | SQLite database |
| 0x002C3528 | `sqlite3BtreeInitPage() returns error code %d` | Known | SQLite database |
| 0x002C6840 | `sqlite_master` | Known | SQLite database |
| 0x002C6850 | `sqlite_temp_master` | Known | SQLite database |
| 0x002DCA04 | `sqlite_stat1` | Known | SQLite database |
| 0x002DCA14 | `CREATE TABLE %Q.sqlite_stat1(tbl,idx,stat)` | Known | SQLite database |
| 0x002DCA40 | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x002E777C | `sqlite_subquery_%p_` | Known | SQLite database |
| 0x0036F680 | `sqlite_master` | Known | SQLite database |
| 0x0036F690 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036F9B4 | `sqlite_` | Known | SQLite database |
| 0x0036F9F4 | `sqlite_master` | Known | SQLite database |
| 0x0036FA04 | `sqlite_temp_master` | Known | SQLite database |
| 0x0036FA1C | `sqlite_sequence` | Known | SQLite database |
| 0x0036FA2C | `UPDATE "%w".sqlite_sequence set name = %Q WHERE name = %Q` | Known | SQLite database |
| 0x0036FB10 | `sqlite_stat1` | Known | SQLite database |
| 0x0036FB20 | `SELECT idx, stat FROM %Q.sqlite_stat1` | Known | SQLite database |
| 0x003707FC | `sqlite_` | Known | SQLite database |
| 0x003709F8 | `sqlite_master` | Known | SQLite database |
| 0x00370A08 | `sqlite_temp_master` | Known | SQLite database |
| 0x00373724 | `sqlite_` | Known | SQLite database |
| 0x00374A10 | `sqlite_autoindex_` | Known | SQLite database |
| 0x00374A24 | `sqlite_master` | Known | SQLite database |
| 0x00374A34 | `sqlite_temp_master` | Known | SQLite database |
| 0x00375E8C | `sqlite_master` | Known | SQLite database |
| 0x00375E9C | `sqlite_temp_master` | Known | SQLite database |
| 0x00375ED0 | `sqlite_stat1` | Known | SQLite database |
| 0x00375EE0 | `DELETE FROM %Q.sqlite_stat1 WHERE idx=%Q` | Known | SQLite database |
| 0x003761C8 | `sqlite_master` | Known | SQLite database |
| 0x003761D8 | `sqlite_temp_master` | Known | SQLite database |
| 0x0037624C | `DELETE FROM %s.sqlite_sequence WHERE name=%Q` | Known | SQLite database |
| 0x003762B4 | `sqlite_stat1` | Known | SQLite database |
| 0x003762C4 | `DELETE FROM %Q.sqlite_stat1 WHERE tbl=%Q` | Known | SQLite database |
| 0x0037663C | `sqlite_master` | Known | SQLite database |
| 0x0037664C | `sqlite_temp_master` | Known | SQLite database |
| 0x00376A64 | `sqlite_master` | Known | SQLite database |
| 0x00376A74 | `sqlite_temp_master` | Known | SQLite database |
| 0x00376A8C | `CREATE TABLE %Q.sqlite_sequence(name,seq)` | Known | SQLite database |
| 0x00379D14 | `sqlite_master` | Known | SQLite database |
| 0x00379D24 | `sqlite_temp_master` | Known | SQLite database |
| 0x0037C10C | `sqlite_temp_master` | Known | SQLite database |
| 0x0037C124 | `sqlite_master` | Known | SQLite database |
| 0x0037D900 | `sqlite3_extension_init` | Known | SQLite database |
| 0x0037E0F4 | `sqlite_master` | Known | SQLite database |
| 0x0037E104 | `sqlite_temp_master` | Known | SQLite database |
| 0x003824E4 | `sqlite_attach` | Known | SQLite database |
| 0x003824F8 | `sqlite_detach` | Known | SQLite database |
| 0x0038522C | `sqlite_master` | Known | SQLite database |
| 0x0038523C | `sqlite_temp_master` | Known | SQLite database |
| 0x0038528C | `sqlite_sequence` | Known | SQLite database |
| 0x0038AB18 | `sqlite_master` | Known | SQLite database |
| 0x0038AB28 | `sqlite_temp_master` | Known | SQLite database |
| 0x0038DEBC | `sqlite_master` | Known | SQLite database |
| 0x0038DECC | `sqlite_temp_master` | Known | SQLite database |
| 0x0039BEB4 | `sqlite_attach` | Known | SQLite database |
| 0x0039BEC4 | `sqlite_detach` | Known | SQLite database |
| 0x00906263 | `SQLite format 3` | Known | SQLite database |
| 0x00908910 | `CREATE TABLE sqlite_master(` | Known | SQLite database |
| 0x00908978 | `CREATE TEMP TABLE sqlite_temp_master(` | Known | SQLite database |
| 0x00909040 | `illegal return value (%d) from the authorization function - should be SQLITE_OK,` | Known | SQLite database |
| 0x009090F8 | `SELECT 'CREATE TABLE vacuum_db.' || substr(sql,14)   FROM sqlite_master WHERE ty` | Known | SQLite database |
| 0x00909180 | `SELECT 'CREATE INDEX vacuum_db.' || substr(sql,14)  FROM sqlite_master WHERE sql` | Known | SQLite database |
| 0x009091E8 | `SELECT 'CREATE UNIQUE INDEX vacuum_db.' || substr(sql,21)   FROM sqlite_master W` | Known | SQLite database |
| 0x00909260 | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x00909310 | `SELECT 'DELETE FROM vacuum_db.' || quote(name) || ';' FROM vacuum_db.sqlite_mast` | Known | SQLite database |
| 0x00909384 | `SELECT 'INSERT INTO vacuum_db.' || quote(name) || ' SELECT * FROM ' || quote(nam` | Known | SQLite database |
| 0x0090941C | `INSERT INTO vacuum_db.sqlite_master   SELECT type, name, tbl_name, rootpage, sql` | Known | SQLite database |
| 0x009095DC | `UPDATE %Q.%s SET sql = CASE WHEN type = 'trigger' THEN sqlite_rename_trigger(sql` | Known | SQLite database |
| 0x00909750 | `UPDATE sqlite_temp_master SET sql = sqlite_rename_trigger(sql, %Q), tbl_name = %` | Known | SQLite database |
| 0x0090998C | `sqlite3_get_table() called with two or more incompatible queries` | Known | SQLite database |
| 0x009D0194 | `sqlite_rename_table` | Known | SQLite database |
| 0x009D0317 | `sqlite_version` | Known | SQLite database |
| 0x009D03B1 | `sqlite_rename_trigger` | Known | SQLite database |
| 0x009D06D5 | `SQLite_iPod_VFS` | Known | SQLite database |

---

## 20. Database (iTunes)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0011F4EC | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x0014AF30 | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x00151950 | `iTunesUSupported` | Known | iTunes database |
| 0x00209208 | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x0020C13C | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x0021111C | `iTunes Image DB.itdb` | Known | iTunes database |
| 0x0022409C | `iTunes Image DB` | Known | iTunes database |
| 0x0022FCE0 | `iTunesUSelected` | Known | iTunes database |
| 0x0022FCF0 | `EmptyiTunesUSelected` | Known | iTunes database |
| 0x003EDC9C | `iTunesDB` | Known | iTunes database |
| 0x0076183C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00761DAB | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00763908 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00763E77 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007659D4 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00765F43 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00767AA0 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0076800F | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00769B6C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0076A0DB | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0076BC38 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0076C1A7 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0076DD04 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0076E273 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0076FDD0 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0077033F | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00771E9C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0077240B | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00773F68 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007744D7 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00776034 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007765A3 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00778100 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0077866F | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0077A1CC | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0077A73B | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0077C298 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0077C807 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0077E364 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0077E8D3 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00780430 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0078099F | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007824FC | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00782A6B | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007845C8 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00784B37 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00786694 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00786C03 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00788760 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00788CCF | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0078A82C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0078AD9B | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0078C8F8 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0078CE67 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0078E9C4 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0078EF33 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00790A90 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00790FFF | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00792B5C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007930CB | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00794C28 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00795197 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00796CF4 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x00797263 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x00798DC0 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0079932F | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0079AE8C | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0079B3FB | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0079CF58 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0079D4C7 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x0079F024 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x0079F593 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007A10F0 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007A165F | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007A31BC | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007A372B | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007A5288 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007A57F7 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007A7354 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007A78C3 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007A9420 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007A998F | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007AB4EC | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007ABA5B | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007AD5B8 | `controller.NoArtiTunesUHilited1` | Known | iTunes database |
| 0x007ADB27 | `controller.iTunesUSelected1` | Known | iTunes database |
| 0x007FF640 | `iTunes` | Known | iTunes database |
| 0x007FF89C | `You can download music from iTunes.` | Known | iTunes database |
| 0x007FF8C0 | `You can download videos from iTunes.` | Known | iTunes database |
| 0x007FF8E8 | `You can download podcasts from iTunes.` | Known | iTunes database |
| 0x007FF910 | `You can download audiobooks from iTunes.` | Known | iTunes database |
| 0x007FF93C | `You can download TV shows from iTunes.` | Known | iTunes database |
| 0x007FF964 | `You can download movies from iTunes.` | Known | iTunes database |
| 0x007FF98C | `You can download music videos from iTunes.` | Known | iTunes database |
| 0x007FF9B8 | `You can sync Photos via iTunes.` | Known | iTunes database |
| 0x007FF9D8 | `You can create playlists and sync via iTunes.` | Known | iTunes database |
| 0x007FFA08 | `You can download rentals from iTunes.` | Known | iTunes database |
| 0x007FFC90 | `If you forget the combination, connect to iTunes to unlock your iPod.` | Known | iTunes database |
| 0x008001B4 | `You can download contacts from iTunes.` | Known | iTunes database |
| 0x008008EC | `To view your To Do items here, enable syncing from iTunes under the Calendar sec` | Known | iTunes database |
| | *...and 430 more* | | |

---

## 21. Nike+ iPod

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00760B81 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00760BC4 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00760BDB | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00760BFA | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00762C4D | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00762C90 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00762CA7 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00762CC6 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00764D19 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00764D5C | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00764D73 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00764D92 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00766DE5 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00766E28 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00766E3F | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00766E5E | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00768EB1 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00768EF4 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00768F0B | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00768F2A | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0076AF7D | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0076AFC0 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0076AFD7 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0076AFF6 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0076D049 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0076D08C | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0076D0A3 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0076D0C2 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0076F115 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0076F158 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0076F16F | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0076F18E | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x007711E1 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00771224 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0077123B | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0077125A | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x007732AD | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x007732F0 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00773307 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00773326 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00775379 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x007753BC | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x007753D3 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x007753F2 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00777445 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00777488 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0077749F | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x007774BE | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00779511 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00779554 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0077956B | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0077958A | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0077B5DD | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0077B620 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0077B637 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0077B656 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0077D6A9 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0077D6EC | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0077D703 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0077D722 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0077F775 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0077F7B8 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0077F7CF | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0077F7EE | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00781841 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00781884 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0078189B | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x007818BA | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0078390D | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00783950 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00783967 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00783986 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x007859D9 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00785A1C | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00785A33 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00785A52 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00787AA5 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00787AE8 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00787AFF | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00787B1E | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00789B71 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00789BB4 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00789BCB | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00789BEA | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0078BC3D | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0078BC80 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0078BC97 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0078BCB6 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0078DD09 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0078DD4C | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0078DD63 | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0078DD82 | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x0078FDD5 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x0078FE18 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x0078FE2F | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x0078FE4E | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| 0x00791EA1 | `controller.GoToLoadTrainer1` | Known | Nike+ fitness |
| 0x00791EE4 | `NikePlus_LoadScreen` | Known | Nike+ fitness |
| 0x00791EFB | `NikePlus_LoadScreen_Default` | Known | Nike+ fitness |
| 0x00791F1A | `controller.GoToTrainerApp1` | Known | Nike+ fitness |
| | *...and 320 more* | | |

---

## 22. FM Radio

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x007454E6 | `Settings_RadioRegionsMenu_Layout"` | Known | FM Radio |
| 0x0074556F | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x00802670 | `Radio Regions` | Known | FM Radio |
| 0x008536F8 | `Radio-Regionen` | Known | FM Radio |
| 0x00988F18 | `23TCSettings_RadioRegions` | Known | FM Radio |
| 0x00989E2B | `27TSilverCntlrTransitionAddonI23TCSettings_RadioRegionsE` | Known | FM Radio |
| 0x009B9159 | `Settings_Radio_Image` | Known | FM Radio |
| 0x009BD05B | `Settings_RadioRegions_Australia_String` | Known | FM Radio |
| 0x009BD082 | `Settings_RadioRegions_Asia_String` | Known | FM Radio |
| 0x009BE2E7 | `Settings_RadioRegions_Europe_String` | Known | FM Radio |
| 0x009BF8F6 | `Settings_RadioRegions_Japan_String` | Known | FM Radio |
| 0x009C059F | `Settings_RadioRegions_Americas_String` | Known | FM Radio |
| 0x009C0CBB | `SettingsMenu_RadioRegions_String` | Known | FM Radio |
| 0x009C41B5 | `RadioRegionsMenuItem` | Known | FM Radio |
| 0x009C7F3C | `SettingsMenu_ListItem_RadioRegions` | Known | FM Radio |
| 0x009CC23B | `SettingsInfo_Template_RadioRegions_Layout` | Known | FM Radio |
| 0x009CC265 | `SettingsMenus_RadioRegions_Layout` | Known | FM Radio |
| 0x009CC8C7 | `Settings_RadioRegionsMenu_Layout` | Known | FM Radio |

---

## 23. Clock/Alarms

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0080132C | `24 Hour Clock` | Known | Clock system |
| 0x00988780 | `15TCClockCityMenu` | Known | Clock system |
| 0x00988A2C | `17TCClockRegionMenu` | Known | Clock system |
| 0x0098965F | `27TSilverCntlrTransitionAddonI15TCClockCityMenuE` | Known | Clock system |
| 0x009898B5 | `27TSilverCntlrTransitionAddonI17TCClockRegionMenuE` | Known | Clock system |
| 0x0098ADA5 | `27TSilverCntlrTransitionAddonI7TCClockE` | Known | Clock system |
| 0x0098B4C6 | `7TCClock` | Known | Clock system |
| 0x009B5CE6 | `Clock_Hours_Image_24` | Known | Clock system |
| 0x009BEE8A | `Settings_DateTime_24HrClock_String` | Known | Clock system |
| 0x009C3206 | `DateTime_List_24HrClock` | Known | Clock system |

---

## 24. Storage (CE-ATA)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000755EC | `cI: Set drive to MMC high speed failed` | Known | CE-ATA/storage interface |
| 0x000756A8 | `cI: could not read CE-ATA task file` | Known | CE-ATA/storage interface |
| 0x000756D0 | `cI: CE-ATA signature missing (%x,%x)` | Known | CE-ATA/storage interface |
| 0x00075728 | `cI: CE-ATA interrupt enable failed` | Known | CE-ATA/storage interface |
| 0x000EECD8 | `mI: card not in MMC TRAN state as expected` | Known | CE-ATA/storage interface |
| 0x0036A9A8 | `MMC init failed` | Known | CE-ATA/storage interface |
| 0x0036A9BC | `CE-ATA init failed` | Known | CE-ATA/storage interface |
| 0x0036AE7C | `ISDIE: CE-ATA interrupt enable failed` | Known | CE-ATA/storage interface |
| 0x005C3135 | `KMMKKKMMMC` | Known | CE-ATA/storage interface |

---

## 25. Storage (NAND Flash)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0013BEB4 | `NAND FLASH DRIVE` | Known | NAND flash interface |

---

## 26. FireWire

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00151418 | `FireWireGUID` | Known | FireWire interface |
| 0x00151428 | `FireWireVersion` | Known | FireWire interface |
| 0x00151E04 | `FireWire` | Known | FireWire interface |

---

## 27. Hardware (GPIO)

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x003EE875 | `GPIO_REG_WRITE` | Known | GPIO hardware |
| 0x003EE886 | `GPIO_INT_INIT` | Known | GPIO hardware |

---

## 28. Hardware Interfaces

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00128FB4 | `MonoHope-LCD` | Known | Hardware interface |
| 0x001295CC | `MonoHope-LCD` | Known | Hardware interface |
| 0x00144418 | `TDiskModeCntlr` | Known | Hardware interface |
| 0x00151F10 | `ForcedDiskMode` | Known | Hardware interface |
| 0x00159BBC | ` rtSltnCpaMBDiskModeImage_SyncArrow1` | Known | Hardware interface |
| 0x001CE6F0 | `MonoHope-LCD` | Known | Hardware interface |
| 0x00266F98 | `Enter Disk Mode` | Known | Hardware interface |
| 0x00266FA8 | `Exit Disk Mode` | Known | Hardware interface |
| 0x003EDEAC | `TDiskModeCntlr` | Known | Hardware interface |
| 0x003EE8B9 | `I2C_MASTER` | Known | Hardware interface |
| 0x003F31CA | `S_I2C_DONE` | Known | Hardware interface |
| 0x003F6CA0 | `TDiskModeCntlr` | Known | Hardware interface |
| 0x003F6D00 | `TDiskModeCntlr` | Known | Hardware interface |
| 0x007599F8 | `TDiskModeCntlr` | Known | Hardware interface |
| 0x009886F5 | `14TCSlideshowLCD` | Known | Hardware interface |
| 0x00988717 | `14TDiskModeCntlr` | Known | Hardware interface |
| 0x009895FF | `27TSilverCntlrTransitionAddonI14TCSlideshowLCDE` | Known | Hardware interface |
| 0x0098962F | `27TSilverCntlrTransitionAddonI14TDiskModeCntlrE` | Known | Hardware interface |
| 0x009B52BB | `DiskModeImage_SyncArrow11` | Known | Hardware interface |
| 0x009B535A | `DiskModeImage_SyncArrow21` | Known | Hardware interface |
| 0x009B5374 | `DiskModeImage_SyncArrow31` | Known | Hardware interface |
| 0x009B5518 | `DiskMode_Text1` | Known | Hardware interface |
| 0x009B5566 | `DiskModeImage_SyncArrow1` | Known | Hardware interface |
| 0x009B5868 | `DiskMode_Text2` | Known | Hardware interface |
| 0x009B59A0 | `DiskModeImage_SyncArrow13` | Known | Hardware interface |
| 0x009B5A0F | `DiskModeImage_SyncArrow23` | Known | Hardware interface |
| 0x009B5A29 | `DiskModeImage_SyncArrow33` | Known | Hardware interface |
| 0x009B5BF0 | `DiskModeImage_SyncArrow3` | Known | Hardware interface |
| 0x009B5EDA | `DiskModeImage_SyncArrow15` | Known | Hardware interface |
| 0x009B5F39 | `DiskModeImage_SyncArrow25` | Known | Hardware interface |
| 0x009B5F53 | `DiskModeImage_SyncArrow35` | Known | Hardware interface |
| 0x009B604E | `DiskModeImage_SyncArrow5` | Known | Hardware interface |
| 0x009B6276 | `DiskModeImage_SyncArrow17` | Known | Hardware interface |
| 0x009B62D5 | `DiskModeImage_SyncArrow27` | Known | Hardware interface |
| 0x009B63AF | `DiskModeImage_SyncArrow7` | Known | Hardware interface |
| 0x009B65D7 | `DiskModeImage_SyncArrow19` | Known | Hardware interface |
| 0x009B6636 | `DiskModeImage_SyncArrow29` | Known | Hardware interface |
| 0x009B6710 | `DiskModeImage_SyncArrow9` | Known | Hardware interface |
| 0x009B7647 | `DiskMode_View_Connected` | Known | Hardware interface |
| 0x009B7682 | `DiskMode_View_Disconnected` | Known | Hardware interface |
| 0x009B7A88 | `DiskMode_iPod` | Known | Hardware interface |
| 0x009B904A | `DiskMode_SyncIcon_Image` | Known | Hardware interface |
| 0x009B9062 | `DiskMode_ConnectedIcon_Image` | Known | Hardware interface |
| 0x009B909B | `DiskMode_DisconnectIcon_Image` | Known | Hardware interface |
| 0x009B9CD1 | `DiskMode_SyncArrows_Image` | Known | Hardware interface |
| 0x009BC608 | `DiskMode_View_Loading` | Known | Hardware interface |
| 0x009BD04A | `DiskMode__String` | Known | Hardware interface |
| 0x009BD861 | `DiskMode_Connected_String` | Known | Hardware interface |
| 0x009BE824 | `DiskMode_Syncing_String` | Known | Hardware interface |
| 0x009BE83C | `DiskMode_Loading_String` | Known | Hardware interface |
| 0x009BEC92 | `DiskMode_Synchronizing_String` | Known | Hardware interface |
| 0x009C1690 | `DiskMode_UseiTunesToEject_String` | Known | Hardware interface |
| 0x009C16D2 | `DiskMode_OKToDisconnect_String` | Known | Hardware interface |
| 0x009C16F1 | `DiskMode_OkayToDisconnect_String` | Known | Hardware interface |
| 0x009C1712 | `DiskMode_EjectingYouMayDisconnect_String` | Known | Hardware interface |
| 0x009C190B | `DiskMode_PleaseWait_String` | Known | Hardware interface |
| 0x009C1926 | `DiskMode_EjectingPleaseWait_String` | Known | Hardware interface |
| 0x009C28C9 | `DiskMode_View_Synchronizing` | Known | Hardware interface |
| 0x009C37A4 | `DiskModeModel` | Known | Hardware interface |
| 0x009C396F | `DiskModeImage_Progress_Full_Fill` | Known | Hardware interface |
| 0x009C39B7 | `DiskModeImage_Progress_Empty_Fill` | Known | Hardware interface |
| 0x009C5630 | `DiskModeImage_SyncIcon` | Known | Hardware interface |
| 0x009C5647 | `DiskModeImage_ConnectedIcon` | Known | Hardware interface |
| 0x009C5663 | `DiskModeImage_DisconnectIcon` | Known | Hardware interface |
| 0x009C612C | `DiskModeImage_Progress_Full_LeftCap` | Known | Hardware interface |
| 0x009C6150 | `DiskModeImage_Progress_Empty_LeftCap` | Known | Hardware interface |
| 0x009C6175 | `DiskModeImage_Progress_Full_RightCap` | Known | Hardware interface |
| 0x009C619A | `DiskModeImage_Progress_Empty_RightCap` | Known | Hardware interface |
| 0x009C6F99 | `DiskMode_Arrows_Color` | Known | Hardware interface |
| 0x009C703F | `DiskMode_Text_Color` | Known | Hardware interface |
| 0x009CB36B | `DiskModeLargeFont` | Known | Hardware interface |
| 0x009CB397 | `DiskModeSmallFont` | Known | Hardware interface |
| 0x009CDBB9 | `DiskMode_View` | Known | Hardware interface |
| 0x009CE865 | `DiskMode_Progress_View` | Known | Hardware interface |
| 0x009D51F8 | `OCSP_RESPID` | Known | Hardware interface |

---

## 29. Power Management

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00003554 | `iPodPowerProfile.txt` | Known | Power management |
| 0x00144430 | `TChargingModeCntlr` | Known | Power management |
| 0x0014444C | `TChargingModeLowPowerCntlr` | Known | Power management |
| 0x00151DE0 | `PowerInformation` | Known | Power management |
| 0x001523A0 | `BatteryPollInterval` | Known | Power management |
| 0x00267004 | `Begin Charging` | Known | Power management |
| 0x00267014 | `Stop Charging` | Known | Power management |
| 0x002AD8EC | `USBPowerSense` | Known | Power management |
| 0x002AD9AC | `PCFPowerMgr` | Known | Power management |
| 0x002AD9F4 | `PowerMgmt` | Known | Power management |
| 0x003EDE90 | `TChargingModeLowPowerCntlr` | Known | Power management |
| 0x003EDEBC | `TChargingModeCntlr` | Known | Power management |
| 0x003F6CB0 | `TChargingModeCntlr` | Known | Power management |
| 0x003F6CC4 | `TChargingModeLowPowerCntlr` | Known | Power management |
| 0x003F6D6C | `TChargingModeCntlr` | Known | Power management |
| 0x003F6D80 | `SwitchToCharging` | Known | Power management |
| 0x003F8094 | `TChargingModeCntlr` | Known | Power management |
| 0x00402188 | `TChargingModeCntlr` | Known | Power management |
| 0x00759A08 | `TChargingModeCntlr` | Known | Power management |
| 0x00759A1C | `TChargingModeLowPowerCntlr` | Known | Power management |
| 0x007DA55F | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DAA9B | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DAFD7 | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DB513 | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DBA4F | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DBF8B | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DC4C7 | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DC98B | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DCE4F | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DD313 | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007DD7D7 | `controller.EnterBatteryLow1` | Known | Power management |
| 0x007E62D4 | `controller.SwitchToCharging1` | Known | Power management |
| 0x007FFE74 | `Low Battery` | Known | Power management |
| 0x007FFE80 | `Connect to Power` | Known | Power management |
| 0x00802F44 | `Charging` | Known | Power management |
| 0x00802F58 | `Low Battery` | Known | Power management |
| 0x00802F64 | `Connect to Power` | Known | Power management |
| 0x00841310 | `PowerSong` | Known | Power management |
| 0x0084A7EC | `PowerSong` | Known | Power management |
| 0x00854354 | `PowerSong` | Known | Power management |
| 0x00861274 | `PowerSong` | Known | Power management |
| 0x0086C560 | `PowerSong` | Known | Power management |
| 0x008758C4 | `PowerSong` | Known | Power management |
| 0x0087F1D4 | `PowerSong` | Known | Power management |
| 0x00889320 | `PowerSong` | Known | Power management |
| 0x00892D24 | `PowerSong` | Known | Power management |
| 0x0089DEAC | `PowerSong` | Known | Power management |
| 0x008A8AE8 | `PowerSong` | Known | Power management |
| 0x008B2720 | `PowerSong` | Known | Power management |
| 0x008BBB9C | `PowerSong` | Known | Power management |
| 0x008C5428 | `PowerSong` | Known | Power management |
| 0x008CEED0 | `PowerSong` | Known | Power management |
| 0x008DB744 | `PowerSong` | Known | Power management |
| 0x008E5C84 | `PowerSong` | Known | Power management |
| 0x008EF49C | `PowerSong` | Known | Power management |
| 0x008F8D1C | `PowerSong` | Known | Power management |
| 0x00902520 | `PowerSong` | Known | Power management |
| 0x00988913 | `16TPowerStatusView` | Known | Power management |
| 0x00988B21 | `18TChargingModeCntlr` | Known | Power management |
| 0x00989108 | `26TChargingModeLowPowerCntlr` | Known | Power management |
| 0x00989982 | `27TSilverCntlrTransitionAddonI18TChargingModeCntlrE` | Known | Power management |
| 0x0098A076 | `27TSilverCntlrTransitionAddonI26TChargingModeLowPowerCntlrE` | Known | Power management |
| 0x0098B66F | `N3ISL17IPodPowerListenerE` | Known | Power management |
| 0x009B507A | `StatusBarWhite_Battery_Image_10` | Known | Power management |
| 0x009B509A | `StatusBarBlack_Battery_Image_10` | Known | Power management |
| 0x009B50FF | `StatusBarWhite_Battery_Image_20` | Known | Power management |
| 0x009B511F | `StatusBarBlack_Battery_Image_20` | Known | Power management |
| 0x009B51A7 | `StatusBarWhite_Battery_Image_0` | Known | Power management |
| 0x009B51C6 | `StatusBarBlack_Battery_Image_0` | Known | Power management |
| 0x009B527B | `StatusBarWhite_Battery_Image_11` | Known | Power management |
| 0x009B529B | `StatusBarBlack_Battery_Image_11` | Known | Power management |
| 0x009B531A | `StatusBarWhite_Battery_Image_21` | Known | Power management |
| 0x009B533A | `StatusBarBlack_Battery_Image_21` | Known | Power management |
| 0x009B5417 | `StatusBarWhite_Battery_Image_1` | Known | Power management |
| 0x009B5436 | `StatusBarBlack_Battery_Image_1` | Known | Power management |
| 0x009B55EC | `StatusBarWhite_Battery_Image_12` | Known | Power management |
| 0x009B560C | `StatusBarBlack_Battery_Image_12` | Known | Power management |
| 0x009B5671 | `StatusBarWhite_Battery_Image_22` | Known | Power management |
| 0x009B5691 | `StatusBarBlack_Battery_Image_22` | Known | Power management |
| 0x009B573A | `StatusBarWhite_Battery_Image_2` | Known | Power management |
| 0x009B5759 | `StatusBarBlack_Battery_Image_2` | Known | Power management |
| 0x009B5960 | `StatusBarWhite_Battery_Image_13` | Known | Power management |
| 0x009B5980 | `StatusBarBlack_Battery_Image_13` | Known | Power management |
| 0x009B5ACC | `StatusBarWhite_Battery_Image_3` | Known | Power management |
| 0x009B5AEB | `StatusBarBlack_Battery_Image_3` | Known | Power management |
| 0x009B5C76 | `StatusBarWhite_Battery_Image_14` | Known | Power management |
| 0x009B5C96 | `StatusBarBlack_Battery_Image_14` | Known | Power management |
| 0x009B5D84 | `StatusBarWhite_Battery_Image_4` | Known | Power management |
| 0x009B5DA3 | `StatusBarBlack_Battery_Image_4` | Known | Power management |
| 0x009B5E9A | `StatusBarWhite_Battery_Image_15` | Known | Power management |
| 0x009B5EBA | `StatusBarBlack_Battery_Image_15` | Known | Power management |
| 0x009B5FE7 | `StatusBarWhite_Battery_Image_5` | Known | Power management |
| 0x009B6006 | `StatusBarBlack_Battery_Image_5` | Known | Power management |
| 0x009B60AC | `StatusBarWhite_Battery_Image_16` | Known | Power management |
| 0x009B60CC | `StatusBarBlack_Battery_Image_16` | Known | Power management |
| 0x009B6199 | `StatusBarWhite_Battery_Image_6` | Known | Power management |
| 0x009B61B8 | `StatusBarBlack_Battery_Image_6` | Known | Power management |
| 0x009B6236 | `StatusBarWhite_Battery_Image_17` | Known | Power management |
| 0x009B6256 | `StatusBarBlack_Battery_Image_17` | Known | Power management |
| 0x009B6357 | `StatusBarWhite_Battery_Image_7` | Known | Power management |
| | *...and 55 more* | | |

---

## 30. UI Elements

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000AB67C | `Calendars/` | Known | UI element |
| 0x000BDE88 | `Calendars` | Known | UI element |
| 0x0010BFA8 | `AlarmHilited` | Known | UI element |
| 0x0010C0BC | `NewAlarmSelected` | Known | UI element |
| 0x0010C0D0 | `AlarmSelected` | Known | UI element |
| 0x0010C0E0 | `CalendarEventSelected` | Known | UI element |
| 0x0010CCB0 | `GotoNowPlaying` | Known | UI element |
| 0x0010CD28 | `GotoMainMenu` | Known | UI element |
| 0x00131A88 | `GotoNowPlaying` | Known | UI element |
| 0x00131A9C | `GotoAlbums` | Known | UI element |
| 0x00131AA8 | `GotoSongs` | Known | UI element |
| 0x00156B44 | `GotoMainMenu` | Known | UI element |
| 0x001E86C4 | `GotoPlayDeleteMenu` | Known | UI element |
| 0x001E8748 | `GotoNowPlaying` | Known | UI element |
| 0x001F5DFC | `ToggleAlarm` | Known | UI element |
| 0x00216864 | `AlarmTonesChosen` | Known | UI element |
| 0x00216878 | `AlarmToneAt` | Known | UI element |
| 0x00216C20 | `AlarmToneAt` | Known | UI element |
| 0x00216F50 | `GotoDefaultLayout` | Known | UI element |
| 0x00216FD4 | `GotoVolumeLayout` | Known | UI element |
| 0x0021710C | `GotoProgressLayout` | Known | UI element |
| 0x00217428 | `GotoDefault` | Known | UI element |
| 0x0021775C | `GotoProgressLayout` | Known | UI element |
| 0x0021791C | `GotoRentalWarningLayout` | Known | UI element |
| 0x002179A0 | `GotoProgressLayout` | Known | UI element |
| 0x00217CB0 | `GotoProgressLayout` | Known | UI element |
| 0x0021983C | `GotoNowPlaying` | Known | UI element |
| 0x0021A14C | `GotoNowPlaying` | Known | UI element |
| 0x0021A458 | `GotoNowPlaying` | Known | UI element |
| 0x0021CEB0 | `GotoStatusBarVideoLayout` | Known | UI element |
| 0x0021CECC | `GotoDefaultVideoLayout` | Known | UI element |
| 0x0021CEE4 | `GotoDefaultLayout` | Known | UI element |
| 0x0021CEF8 | `GotoDefaultSubtitlesLayout` | Known | UI element |
| 0x0021CF90 | `GotoVolumeLayout` | Known | UI element |
| 0x0021CFA4 | `GotoVolumeVideoLayout` | Known | UI element |
| 0x0021D044 | `GotoProgressLayout` | Known | UI element |
| 0x0021D058 | `GotoProgressVideoLayout` | Known | UI element |
| 0x0021D80C | `GotoProgressVideoLayout` | Known | UI element |
| 0x0021DC74 | `GotoCaptionVideoLayout` | Known | UI element |
| 0x0021DEE0 | `GotoProgressLayout` | Known | UI element |
| 0x0021DEF4 | `GotoProgressVideoLayout` | Known | UI element |
| 0x0021E09C | `GotoBrightnessVideoLayout` | Known | UI element |
| 0x0021E0D4 | `GotoRatingLayout` | Known | UI element |
| 0x0021E59C | `GotoChapterArtLayout` | Known | UI element |
| 0x0021E5B4 | `GotoShuffleLayout` | Known | UI element |
| 0x0021E944 | `GotoExtraInfoLayout` | Known | UI element |
| 0x0021E958 | `GotoExtraInfoLoadingLayout` | Known | UI element |
| 0x0021EA28 | `GotoVolumeLayout` | Known | UI element |
| 0x0021EA40 | `GotoVolumeVideoLayout` | Known | UI element |
| 0x0021EACC | `GotoVolumeLayout` | Known | UI element |
| 0x0021EAE0 | `GotoVolumeVideoLayout` | Known | UI element |
| 0x0021ECF0 | `GotoScrubLayout` | Known | UI element |
| 0x0021ED00 | `GotoScrubVideoLayout` | Known | UI element |
| 0x0021ED90 | `GotoProgressLayout` | Known | UI element |
| 0x0021EDA4 | `GotoProgressVideoLayout` | Known | UI element |
| 0x0021EFFC | `GotoStatusBarVideoLayout` | Known | UI element |
| 0x0021F018 | `GotoDefaultVideoLayout` | Known | UI element |
| 0x0021F030 | `GotoDefaultSubtitlesLayout` | Known | UI element |
| 0x0021F04C | `GotoDefaultLayout` | Known | UI element |
| 0x0021F878 | `GotoChapterArtLayout` | Known | UI element |
| 0x0021F970 | `GotoProgressLayout` | Known | UI element |
| 0x0021F9FC | `GotoProgressLayout` | Known | UI element |
| 0x0021FA10 | `GotoProgressVideoLayout` | Known | UI element |
| 0x0021FAEC | `GotoExtraInfoLoadFailedLayout` | Known | UI element |
| 0x0021FB0C | `GotoExtraInfoLayout` | Known | UI element |
| 0x0021FF48 | `GotoStatusBarLayout` | Known | UI element |
| 0x0021FF5C | `GotoDefaultLayout` | Known | UI element |
| 0x00220134 | `GotoDefault` | Known | UI element |
| 0x00220268 | `GotoProgressLayout` | Known | UI element |
| 0x00220428 | `GotoCaptionVideoLayout` | Known | UI element |
| 0x00220578 | `GotoBrightnessLayout` | Known | UI element |
| 0x002205FC | `GotoBrightnessLayout` | Known | UI element |
| 0x0022067C | `GotoVolumeLayout` | Known | UI element |
| 0x002206C8 | `GotoScrubLayout` | Known | UI element |
| 0x00220790 | `GotoStatusBarLayout` | Known | UI element |
| 0x002207A4 | `GotoDefaultLayout` | Known | UI element |
| 0x0022087C | `GotoScrubLayout` | Known | UI element |
| 0x002208CC | `GotoScrubLayout` | Known | UI element |
| 0x00227034 | `GotoNowPlaying` | Known | UI element |
| 0x00227330 | `GotoNowPlaying` | Known | UI element |
| 0x0022863C | `GotoFourCard_About` | Known | UI element |
| 0x00228650 | `GotoThreeCard_About` | Known | UI element |
| 0x0022D620 | `looBCalendarEventlessDaySelected` | Known | UI element |
| 0x0022D644 | `CalendarEventfullDaySelected` | Known | UI element |
| 0x0022D760 | `looBCalendarEventlessDaySelected` | Known | UI element |
| 0x0022D784 | `CalendarEventfullDaySelected` | Known | UI element |
| 0x0022D850 | `CalendarEventSelected` | Known | UI element |
| 0x002302D8 | `GotoNowPlaying` | Known | UI element |
| 0x002309EC | `GotoNowPlaying` | Known | UI element |
| 0x002311EC | `GotoFirstBoot` | Known | UI element |
| 0x002311FC | `GotoNotesApp` | Known | UI element |
| 0x00231210 | `GotoLockApp` | Known | UI element |
| 0x00234464 | `looBCalendarEventlessDaySelected` | Known | UI element |
| 0x00234488 | `CalendarEventfullDaySelected` | Known | UI element |
| 0x00234F64 | `CalendarSelected` | Known | UI element |
| 0x00235D30 | `GotoNowPlaying` | Known | UI element |
| 0x002398B0 | `GotoNowPlaying` | Known | UI element |
| 0x003F52F4 | `AlarmLabelChosen` | Known | UI element |
| 0x003F5308 | `AlarmSoundChosen` | Known | UI element |
| 0x003F531C | `AlarmTimeChosen` | Known | UI element |
| | *...and 309 more* | | |

---

## 31. Menu Items

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x007FF634 | `Podcasts` | Known | Menu item |
| 0x007FF64C | `Photos` | Known | Menu item |
| 0x007FF654 | `Videos` | Known | Menu item |
| 0x007FF65C | `Extras` | Known | Menu item |
| 0x007FF678 | `Settings` | Known | Menu item |
| 0x007FF684 | `Shuffle Songs` | Known | Menu item |
| 0x007FF694 | `Now Playing` | Known | Menu item |
| 0x007FF6C8 | `Playlists` | Known | Menu item |
| 0x007FF6D4 | `Artists` | Known | Menu item |
| 0x007FF6DC | `Albums` | Known | Menu item |
| 0x007FF6FC | `Genres` | Known | Menu item |
| 0x007FF704 | `Composers` | Known | Menu item |
| 0x007FF710 | `Audiobooks` | Known | Menu item |
| 0x00800370 | `Playlists` | Known | Menu item |
| 0x00801634 | `Artists` | Known | Menu item |
| 0x0080163C | `Albums` | Known | Menu item |
| 0x0080164C | `Genres` | Known | Menu item |
| 0x00801654 | `Composers` | Known | Menu item |
| 0x00801684 | `Photos` | Known | Menu item |
| 0x0080168C | `Playlists` | Known | Menu item |
| 0x008016A8 | `Audiobooks` | Known | Menu item |
| 0x008016B4 | `Podcasts` | Known | Menu item |
| 0x00801A34 | `Settings` | Known | Menu item |
| 0x00801C10 | `Albums` | Known | Menu item |
| 0x00801C24 | `Now Playing` | Known | Menu item |
| 0x0080224C | `Photos` | Known | Menu item |
| 0x00802260 | `Albums` | Known | Menu item |
| 0x00802268 | `Settings` | Known | Menu item |
| 0x008022B4 | `Settings` | Known | Menu item |
| 0x0080232C | `Now Playing` | Known | Menu item |
| 0x008024A4 | `Podcasts` | Known | Menu item |
| 0x008024D8 | `Videos` | Known | Menu item |
| 0x008024E0 | `Photos` | Known | Menu item |
| 0x0080261C | `Main Menu` | Known | Menu item |
| 0x00802664 | `Audiobooks` | Known | Menu item |
| 0x0080270C | `Albums` | Known | Menu item |
| 0x0083D508 | `Podcasts` | Known | Menu item |
| 0x0083F4C0 | `Podcasts` | Known | Menu item |
| 0x008404A8 | `Podcasts` | Known | Menu item |
| 0x00846C40 | `Podcasts` | Known | Menu item |
| 0x00848C44 | `Podcasts` | Known | Menu item |
| 0x008500BC | `Podcasts` | Known | Menu item |
| 0x008500D0 | `Videos` | Known | Menu item |
| 0x008500D8 | `Extras` | Known | Menu item |
| 0x00850190 | `Genres` | Known | Menu item |
| 0x008523B8 | `Genres` | Known | Menu item |
| 0x00852434 | `Podcasts` | Known | Menu item |
| 0x00853518 | `Podcasts` | Known | Menu item |
| 0x00853540 | `Videos` | Known | Menu item |
| 0x00859FDC | `Podcasts` | Known | Menu item |
| 0x0085DE68 | `Podcasts` | Known | Menu item |
| 0x0085FA54 | `Podcasts` | Known | Menu item |
| 0x00868364 | `Podcasts` | Known | Menu item |
| 0x00868380 | `Extras` | Known | Menu item |
| 0x0086A5C0 | `Podcasts` | Known | Menu item |
| 0x0086B6C8 | `Podcasts` | Known | Menu item |
| 0x0087ACC4 | `Podcasts` | Known | Menu item |
| 0x0087ACD0 | `Photos` | Known | Menu item |
| 0x0087ACE0 | `Extras` | Known | Menu item |
| 0x0087AD74 | `Albums` | Known | Menu item |
| 0x0087AD98 | `Genres` | Known | Menu item |
| 0x0087D0C8 | `Albums` | Known | Menu item |
| 0x0087D0DC | `Genres` | Known | Menu item |
| 0x0087D120 | `Photos` | Known | Menu item |
| 0x0087D168 | `Podcasts` | Known | Menu item |
| 0x0087D7FC | `Albums` | Known | Menu item |
| 0x0087DF30 | `Photos` | Known | Menu item |
| 0x0087DF40 | `Albums` | Known | Menu item |
| 0x0087E204 | `Podcasts` | Known | Menu item |
| 0x0087E238 | `Photos` | Known | Menu item |
| 0x0087E4F0 | `Albums` | Known | Menu item |
| 0x008AE850 | `Podcasts` | Known | Menu item |
| 0x008AE900 | `Albums` | Known | Menu item |
| 0x008AE91C | `Genres` | Known | Menu item |
| 0x008B094C | `Albums` | Known | Menu item |
| 0x008B095C | `Genres` | Known | Menu item |
| 0x008B09CC | `Podcasts` | Known | Menu item |
| 0x008B0FFC | `Albums` | Known | Menu item |
| 0x008B16C4 | `Albums` | Known | Menu item |
| 0x008B196C | `Podcasts` | Known | Menu item |
| 0x008B1BD8 | `Albums` | Known | Menu item |
| 0x008CAD7C | `Podcasts` | Known | Menu item |
| 0x008CAD98 | `Extras` | Known | Menu item |
| 0x008CD034 | `Podcasts` | Known | Menu item |
| 0x008CE140 | `Podcasts` | Known | Menu item |
| 0x00A1B4CC | `Settings` | Known | Menu item |

---

## 32. Filesystem Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x00047D2C | `iPod_Control` | Filesystem Path | |
| 0x00047D40 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x00047D58 | `iPod_Control\iTunes\firsttime` | Filesystem Path | |
| 0x00047D98 | `iPod_Control\Device` | Filesystem Path | |
| 0x00056A84 | `iPod_Control\Device` | Filesystem Path | |
| 0x00058B4C | `iPod_Control` | Filesystem Path | |
| 0x000591C0 | `iPod_Control\Device\SysInfo` | Filesystem Path | |
| 0x000690E8 | `iPod_Control\Artwork\ArtworkDB` | Filesystem Path | |
| 0x00069108 | `iPod_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x00069130 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x0006BD30 | `iPod_Control\Music\` | Filesystem Path | |
| 0x0006ED7C | `iPod_Control\Device\Preferences` | Filesystem Path | |
| 0x0006EEFC | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x00094584 | `iPod_Control/iTunes/` | Filesystem Path | |
| 0x000985D8 | `iPod_Control/iTunes/iTunesDB.p7b` | Filesystem Path | |
| 0x000A265C | `iPod_Control` | Filesystem Path | |
| 0x000A266C | `Resources/Games` | Filesystem Path | |
| 0x000A267C | `iPod_Control/%s%s%s` | Filesystem Path | |
| 0x000ABCD4 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000B53E4 | `iPod_Control/iTunes/` | Filesystem Path | |
| 0x000B55CC | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000B5EF8 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000BE0A0 | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x000BF64C | `iPod_Control\iTunes\Play Counts` | Filesystem Path | |
| 0x000BF74C | `iPod_Control\iTunes\` | Filesystem Path | |
| 0x001001DC | `iPod_Control\Device\dst` | Filesystem Path | |
| 0x0010B314 | `iPod_Control/Device/alarms` | Filesystem Path | |
| 0x0011BA8C | `iPod_Control/Device/radio` | Filesystem Path | |
| 0x0011CFCC | `iPod_Control/Device` | Filesystem Path | |
| 0x0011CFE0 | `iPod_Control/Device/radio` | Filesystem Path | |
| 0x00137054 | `iPod_Control/Device/Users` | Filesystem Path | |
| 0x0013C650 | `iPod_Control\iTunes\iTunesDB` | Filesystem Path | |
| 0x00165208 | `/iPod_Control/Device/1da` | Filesystem Path | |
| 0x00165464 | `/iPod_Control/Device/1da` | Filesystem Path | |
| 0x00172070 | `Resources/UI/active.bin` | Filesystem Path | |
| 0x00172088 | `Resources/UI/` | Filesystem Path | |
| 0x00196214 | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x00197140 | `/iPod_Control/Device/Accessories/Tags` | Filesystem Path | |
| 0x00197168 | `/iPod_Control/Device/Accessories` | Filesystem Path | |
| 0x001BD040 | `iPod_Control/Device/PlayCounts` | Filesystem Path | |
| 0x001D32F0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D33A0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D351C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D36B4 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D375C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D390C | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D39B0 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D3A54 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D3AF8 | `iPod_Control\Device\` | Filesystem Path | |
| 0x001D3B9C | `iPod_Control\Device\` | Filesystem Path | |

---

## 33. Build Paths

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0090C968 | `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftutil.c` | Build Path | |
| 0x0090C9C0 | `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftstream.c` | Build Path | |
| 0x0090CA18 | `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftobjs.c` | Build Path | |
| 0x009173B8 | `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\afglobal.c` | Build Path | |
| 0x00917F34 | `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfdrivr.c` | Build Path | |
| 0x00919130 | `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrgload.c` | Build Path | |
| 0x00919188 | `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrcmap.c` | Build Path | |
| 0x009191E0 | `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrobjs.c` | Build Path | |
| 0x00919524 | `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1cmap.c` | Build Path | |
| 0x009288CC | `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttcmap.c` | Build Path | |
| 0x00928B48 | `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype\ttgload.c` | Build Path | |
| 0x009290B4 | `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1load.c` | Build Path | |

---

## 34. EQ Presets

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000A02A4 | `Acoustic` | EQ Preset | |
| 0x000A02B0 | `Bass Booster` | EQ Preset | |
| 0x000A02D0 | `Classical` | EQ Preset | |
| 0x000A02EC | `Electronic` | EQ Preset | |
| 0x000A0300 | `Hip Hop` | EQ Preset | |
| 0x000A0318 | `Loudness` | EQ Preset | |
| 0x000A0324 | `Lounge` | EQ Preset | |
| 0x000A0348 | `Small Speakers` | EQ Preset | |
| 0x000A0358 | `Spoken Word` | EQ Preset | |
| 0x000A0364 | `Treble Booster` | EQ Preset | |
| 0x000A03B0 | `Vocal Booster` | EQ Preset | |
| 0x00802960 | `Acoustic` | EQ Preset | |
| 0x0080296C | `Bass Booster` | EQ Preset | |
| 0x0080298C | `Classical` | EQ Preset | |
| 0x008029A8 | `Electronic` | EQ Preset | |
| 0x008029BC | `Hip Hop` | EQ Preset | |
| 0x008029D4 | `Loudness` | EQ Preset | |
| 0x008029E0 | `Lounge` | EQ Preset | |
| 0x00802A00 | `Small Speakers` | EQ Preset | |
| 0x00802A10 | `Spoken Word` | EQ Preset | |
| 0x00802A1C | `Treble Booster` | EQ Preset | |
| 0x00802A3C | `Vocal Booster` | EQ Preset | |
| 0x00840A94 | `Acoustic` | EQ Preset | |
| 0x00840AA0 | `Bass Booster` | EQ Preset | |
| 0x00840AC0 | `Classical` | EQ Preset | |
| 0x00840ADC | `Electronic` | EQ Preset | |
| 0x00840AF0 | `Hip Hop` | EQ Preset | |
| 0x00840B08 | `Loudness` | EQ Preset | |
| 0x00840B14 | `Lounge` | EQ Preset | |
| 0x00840B34 | `Small Speakers` | EQ Preset | |
| 0x00840B44 | `Spoken Word` | EQ Preset | |
| 0x00840B50 | `Treble Booster` | EQ Preset | |
| 0x00840B70 | `Vocal Booster` | EQ Preset | |
| 0x00849FF0 | `Acoustic` | EQ Preset | |
| 0x00849FFC | `Bass Booster` | EQ Preset | |
| 0x0084A01C | `Classical` | EQ Preset | |
| 0x0084A038 | `Electronic` | EQ Preset | |
| 0x0084A04C | `Hip Hop` | EQ Preset | |
| 0x0084A064 | `Loudness` | EQ Preset | |
| 0x0084A070 | `Lounge` | EQ Preset | |
| 0x0084A090 | `Small Speakers` | EQ Preset | |
| 0x0084A0A0 | `Spoken Word` | EQ Preset | |
| 0x0084A0AC | `Treble Booster` | EQ Preset | |
| 0x0084A0CC | `Vocal Booster` | EQ Preset | |
| 0x00853AA0 | `Acoustic` | EQ Preset | |
| 0x00853AE0 | `Electronic` | EQ Preset | |
| 0x00853B0C | `Loudness` | EQ Preset | |
| 0x0086BCFC | `Hip Hop` | EQ Preset | |
| 0x0086BD0C | `Latina` | EQ Preset | |
| 0x0086BD14 | `Loudness` | EQ Preset | |
| 0x0086BD20 | `Lounge` | EQ Preset | |
| 0x00875320 | `Lounge` | EQ Preset | |
| 0x0087E868 | `Hip Hop` | EQ Preset | |
| 0x0087E878 | `Latino` | EQ Preset | |
| 0x0087E88C | `Lounge` | EQ Preset | |
| 0x0089250C | `Hip Hop` | EQ Preset | |
| 0x0089251C | `Latina` | EQ Preset | |
| 0x00892524 | `Loudness` | EQ Preset | |
| 0x00892530 | `Lounge` | EQ Preset | |
| 0x0089D44C | `Acoustic` | EQ Preset | |
| 0x0089D458 | `Bass Booster` | EQ Preset | |
| 0x0089D478 | `Classical` | EQ Preset | |
| 0x0089D494 | `Electronic` | EQ Preset | |
| 0x0089D4A8 | `Hip Hop` | EQ Preset | |
| 0x0089D4C0 | `Loudness` | EQ Preset | |
| 0x0089D4CC | `Lounge` | EQ Preset | |
| 0x0089D4EC | `Small Speakers` | EQ Preset | |
| 0x0089D4FC | `Spoken Word` | EQ Preset | |
| 0x0089D508 | `Treble Booster` | EQ Preset | |
| 0x0089D528 | `Vocal Booster` | EQ Preset | |
| 0x008A81E4 | `Acoustic` | EQ Preset | |
| 0x008A81F0 | `Bass Booster` | EQ Preset | |
| 0x008A8210 | `Classical` | EQ Preset | |
| 0x008A822C | `Electronic` | EQ Preset | |
| 0x008A8240 | `Hip Hop` | EQ Preset | |
| 0x008A8258 | `Loudness` | EQ Preset | |
| 0x008A8264 | `Lounge` | EQ Preset | |
| 0x008A8284 | `Small Speakers` | EQ Preset | |
| 0x008A8294 | `Spoken Word` | EQ Preset | |
| 0x008A82A0 | `Treble Booster` | EQ Preset | |
| 0x008A82C0 | `Vocal Booster` | EQ Preset | |
| 0x008B1F28 | `Loudness` | EQ Preset | |
| 0x008B1F34 | `Lounge` | EQ Preset | |
| 0x008BB420 | `Latino` | EQ Preset | |
| 0x008BB428 | `Loudness` | EQ Preset | |
| 0x008BB434 | `Lounge` | EQ Preset | |
| 0x008C4BAC | `Hip Hop` | EQ Preset | |
| 0x008C4BD8 | `Lounge` | EQ Preset | |
| 0x008CE780 | `Hip Hop` | EQ Preset | |
| 0x008CE790 | `Latina` | EQ Preset | |
| 0x008CE7A4 | `Lounge` | EQ Preset | |
| 0x008E54C0 | `Acoustic` | EQ Preset | |
| 0x008E54CC | `Bass Booster` | EQ Preset | |
| 0x008E54EC | `Classical` | EQ Preset | |
| 0x008E5508 | `Electronic` | EQ Preset | |
| 0x008E551C | `Hip Hop` | EQ Preset | |
| 0x008E5534 | `Loudness` | EQ Preset | |
| 0x008E5540 | `Lounge` | EQ Preset | |
| 0x008E5560 | `Small Speakers` | EQ Preset | |
| 0x008E5570 | `Spoken Word` | EQ Preset | |
| 0x008E557C | `Treble Booster` | EQ Preset | |
| 0x008E559C | `Vocal Booster` | EQ Preset | |
| 0x008EEC5C | `Hip Hop` | EQ Preset | |
| 0x008F8524 | `Acoustic` | EQ Preset | |
| 0x008F8530 | `Bass Booster` | EQ Preset | |
| 0x008F8550 | `Classical` | EQ Preset | |
| 0x008F856C | `Electronic` | EQ Preset | |
| 0x008F8580 | `Hip Hop` | EQ Preset | |
| 0x008F8598 | `Loudness` | EQ Preset | |
| 0x008F85A4 | `Lounge` | EQ Preset | |
| 0x008F85C4 | `Small Speakers` | EQ Preset | |
| 0x008F85D4 | `Spoken Word` | EQ Preset | |
| 0x008F85E0 | `Treble Booster` | EQ Preset | |
| 0x008F8600 | `Vocal Booster` | EQ Preset | |
| 0x00901D08 | `Acoustic` | EQ Preset | |
| 0x00901D14 | `Bass Booster` | EQ Preset | |
| 0x00901D34 | `Classical` | EQ Preset | |
| 0x00901D50 | `Electronic` | EQ Preset | |
| 0x00901D64 | `Hip Hop` | EQ Preset | |
| 0x00901D7C | `Loudness` | EQ Preset | |
| 0x00901D88 | `Lounge` | EQ Preset | |
| 0x00901DA8 | `Small Speakers` | EQ Preset | |
| 0x00901DB8 | `Spoken Word` | EQ Preset | |
| 0x00901DC4 | `Treble Booster` | EQ Preset | |
| 0x00901DE4 | `Vocal Booster` | EQ Preset | |

---

## 35. Diagnostics

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000CE7FC | `cIC12: ATA Status Error! Could not get error code.` | Diagnostic | |
| 0x000CE830 | `cIC12: ATA Status Error! Error code (0x%2x)` | Diagnostic | |
| 0x000DC218 | `Error[%d] has occurred in rule %d` | Diagnostic | |
| 0x0010CC98 | `SwitchToNotesImageError` | Diagnostic | |
| 0x00118914 | `%s Error in file %s.` | Diagnostic | |
| 0x001DEFF0 | `GotoErrorLayout` | Diagnostic | |
| 0x00759E41 | `controller.GotoErrorLayout1` | Diagnostic | |
| 0x00759F13 | `controller.GotoErrorLayout1` | Diagnostic | |
| 0x0075C960 | `controller.ShowMemoryError1` | Diagnostic | |
| 0x0075C9C1 | `controller.ShowSigningError1` | Diagnostic | |
| 0x0075CA24 | `controller.ShowUnknownError1` | Diagnostic | |
| 0x0075CA87 | `controller.ShowVersionError1` | Diagnostic | |
| 0x0075CB44 | `controller.ShowMemoryError1` | Diagnostic | |
| 0x0075CBA5 | `controller.ShowSigningError1` | Diagnostic | |
| 0x0075CC08 | `controller.ShowUnknownError1` | Diagnostic | |
| 0x0075CC6B | `controller.ShowVersionError1` | Diagnostic | |
| 0x0075CD28 | `controller.ShowMemoryError1` | Diagnostic | |
| 0x0075CD89 | `controller.ShowSigningError1` | Diagnostic | |
| 0x0075CDEC | `controller.ShowUnknownError1` | Diagnostic | |
| 0x0075CE4F | `controller.ShowVersionError1` | Diagnostic | |
| 0x0075CF0C | `controller.ShowMemoryError1` | Diagnostic | |
| 0x0075CF6D | `controller.ShowSigningError1` | Diagnostic | |
| 0x0075CFD0 | `controller.ShowUnknownError1` | Diagnostic | |
| 0x0075D033 | `controller.ShowVersionError1` | Diagnostic | |
| 0x0075D0F0 | `controller.ShowMemoryError1` | Diagnostic | |
| 0x0075D151 | `controller.ShowSigningError1` | Diagnostic | |
| 0x0075D1B4 | `controller.ShowUnknownError1` | Diagnostic | |
| 0x0075D217 | `controller.ShowVersionError1` | Diagnostic | |
| 0x0075D4BA | `controller.ShowMemoryError1` | Diagnostic | |
| 0x0075D51B | `controller.ShowSigningError1` | Diagnostic | |

---

## 36. Assertions

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0004F6E8 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0004F7D4 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0004F8A0 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0004F934 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0004F9B0 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0004FA28 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0004FD6C | `assertion failed on line %d of file %s` | Assertion | |
| 0x0005009C | `assertion failed on line %d of file %s` | Assertion | |
| 0x000501DC | `assertion failed on line %d of file %s` | Assertion | |
| 0x0005031C | `assertion failed on line %d of file %s` | Assertion | |
| 0x0005044C | `assertion failed on line %d of file %s` | Assertion | |
| 0x0006AC60 | `assertion failed on line %d of file %s` | Assertion | |
| 0x00086508 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0009C360 | `assertion failed on line %d of file %s` | Assertion | |
| 0x0009D334 | `assertion failed on line %d of file %s` | Assertion | |
| 0x000A3A04 | `assertion failed on line %d of file %s` | Assertion | |
| 0x000A983C | `assertion failed on line %d of file %s` | Assertion | |
| 0x000B6388 | `assertion failed on line %d of file %s` | Assertion | |
| 0x000B6F9C | `assertion failed on line %d of file %s` | Assertion | |
| 0x000B70D8 | `assertion failed on line %d of file %s` | Assertion | |

---
