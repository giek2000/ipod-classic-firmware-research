# iPod Nano 3rd Generation (4/8GB) — RetailOS 1.1.3 Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.1.3 |
| **IPSW** | iPod_26.1.1.3.ipsw |
| **Device** | iPod Nano 3rd Generation (2007, Click Wheel, Cover Flow, Video) |
| **UpdaterFamilyID** | 26 |
| **FamilyID** | 12 |
| **Build Codename** | N46 |
| **Binary Size** | 10,792,304 bytes (10.29 MB) |
| **ARM Code Start** | 0x800 |
| **ARM Code Size** | 10,790,256 bytes |
| **Total Strings (>=6)** | 53,545 |
| **Function Prologues** | 22,820 (ARM: 17,473, Thumb: 5,347) |
| **SoC** | Samsung S5L8702 |
| **Architecture** | ARM926EJ-S (ARMv5TEJ) |
| **Encrypted** | Yes (HW AES-128-CBC, GID key) |
| **SHA-256** | `41f3782d9ae5ab8437e30e9f4b26789e959fea2dae2eea7b60fc14272c669877` |
| **DFU USB PID** | 0x1229 |
| **WTF USB PID** | 0x1231 |

---

## Hardware Overview

| Component | Specification |
|-----------|---------------|
| SoC | Samsung S5L8702 (ARM926EJ-S, ARMv5TEJ) |
| Clock | ~200 MHz |
| RAM | 64 MB SDRAM |
| Storage | 4GB / 8GB NAND Flash |
| Display | 320×240 RGB565, 2.0" TFT |
| Audio | Cirrus Logic codec (I2S) |
| Connectivity | USB 2.0, 30-pin dock |
| Battery | ~580 mAh Li-ion |
| Form Factor | 69.8 × 52.3 × 6.5 mm |

---

## Significance

The iPod Nano 3G is architecturally significant because:
1. **First S5L8702 device** — shares the exact same SoC as the iPod Classic 6G/7G
2. **First Nano with video playback** — 320×240 display with Cover Flow
3. **Same RTOS and framework** — RTXC kernel + Silver UI framework identical to Classic
4. **Same encryption** — AES-128-CBC with GID key, decryptable via wInd3x
5. **Native wInd3x support** — No device descriptor patch needed (unlike Classic)

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0013ED04 | `MP3ExampleTask` | Hidden | MP3 decode test task (development leftover) |
| 0x00158808 | `MockupMode/MockupMode.xml` | Hidden | Developer UI prototyping tool |
| 0x00184F68 | `TCDemoMode` | Hidden | Apple Store demo/retail kiosk mode |
| 0x001E22E0 | `TSilverCntlrTestAppCntlr` | Hidden | UI framework test application |
| 0x0026CC5C | `Channel UnitTests` | Hidden | Built-in unit test logging channel |
| 0x002C3141 | `** RTXCbug - ` | Hidden | Interactive RTXC RTOS debugger |
| 0x002C3184 | `  X - Exit RTXCbug` | Hidden | RTXCbug exit command |
| 0x002C3199 | `RTXCbug> ` | Hidden | RTXCbug command prompt |
| 0x002C3B75 | `RTXCbug - RTXC Objects> ` | Hidden | RTXCbug object inspector |
| 0x003CC709 | `S_RTXCBUG` | Hidden | RTXCbug state identifier |
| 0x003D0EB0 | `TCDemoMode` | Hidden | Demo mode controller (duplicate ref) |
| 0x0091F7F0 | `WaveFileDebugTask` | Hidden | Audio debug recording task |
| 0x00920974 | `TCMockupModeNavScreen` | Hidden | Mockup mode navigation screen |
| 0x009CBD10 | `Returning from RTXCBug` | Hidden | RTXCbug exit confirmation |
| 0x00754412 | `Debug_MainMenu_Screen` | Hidden | Debug menu main screen layout |
| 0x007F4526 | `Debug_UnitTest_Screen` | Hidden | Unit test execution screen |
| 0x007F45A2 | `DemoMode_Screen` | Hidden | Demo mode display screen |
| 0x007F4622 | `Debug_TestList_Screen` | Hidden | Test list browser screen |
| 0x007F46AE | `Debug_TestResult_Screen` | Hidden | Test result display screen |
| 0x009DF986 | `SilverTestApp` | Hidden | Silver framework test application |
| 0x009DF994 | `UnitTestApp` | Hidden | Unit test application entry |
| 0x00A4A360 | `DebugUtil` | Hidden | Debug utility library reference |

---

## 2. UI Controllers (TSilver/TC Classes)

The Nano 3G uses the same Silver UI framework as the Classic, with controllers for all media types:

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x000A0B48 | `TSilverCntlr` | Known | Base UI controller class |
| 0x000A0B60 | `TCExtrasMenu` | Known | Extras menu controller |
| 0x000A0B78 | `TCGameScreen` | Known | Game execution screen |
| 0x000A0B90 | `TCGamesMenu` | Known | Games list menu |
| 0x000A0BA4 | `TSilverMainMediaListCntlr_Main` | Known | Main menu media list |
| 0x000A0BCC | `TSilverMainMediaListCntlr_Music` | Known | Music top-level controller |
| 0x000A0BF4 | `TSilverMainMediaListCntlr_Videos` | Known | Videos top-level controller |
| 0x000A0C20 | `TSilverMediaListCntlr_Songs` | Known | Songs list controller |
| 0x000A0C44 | `TSilverMediaListCntlr_Albums` | Known | Albums list controller |
| 0x000A0C6C | `TSilverMediaListCntlr_Artists` | Known | Artists list controller |
| 0x000A0C94 | `TSilverMediaListCntlr_Genres` | Known | Genres list controller |
| 0x000A0CBC | `TSilverMediaListCntlr_Composers` | Known | Composers list controller |
| 0x000A0CE4 | `TSilverMediaListCntlr_Podcasts` | Known | Podcasts list controller |
| 0x000A0D3C | `TSilverMediaListCntlr_Audiobooks` | Known | Audiobooks list controller |
| 0x000A0D98 | `TSilverMediaListCntlr_TVShows` | Known | TV Shows list controller |
| 0x000A0DE8 | `TSilverMediaListCntlr_TVEpisodes` | Known | TV Episodes list controller |
| 0x000A0E14 | `TSilverMediaListCntlr_Movies` | Known | Movies list controller |
| 0x000A0E3C | `TSilverMediaListCntlr_Playlists` | Known | Playlists list controller |
| 0x000A0FFC | `TSilverMediaListCntlr_Rentals` | Known | iTunes rentals controller |
| 0x000A1098 | `TSilverGlobalCntlr` | Known | Global state controller |
| 0x000A10B4 | `TSilverTrainerCntlr` | Known | Nike+ training controller |
| 0x000FA23C | `TCSlideshowLCD` | Known | Photo slideshow (LCD output) |
| 0x000FA254 | `TCSlideshowTVOut` | Known | Photo slideshow (TV output) |
| 0x0015C990 | `TCSportTimer` | Known | Sport timer controller |
| 0x0015DD68 | `TCVoiceMemos` | Known | Voice memos controller |
| 0x0016F4EC | `TSilverSettingsMenuListCntlr` | Known | Settings menu controller |
| 0x0016F6BC | `TCFirstBoot` | Known | First-boot setup wizard |
| 0x0016B9A8 | `CoverFlow_Screen` | Known | Cover Flow rendering screen |
| 0x00287D0C | `TC_LockDialog` | Known | Screen lock dialog |
| 0x0028DC24 | `TCClock` | Known | Clock/world clock controller |
| 0x0028DC68 | `TCAlarmMenu` | Known | Alarm settings controller |
| 0x00294C64 | `TCNotesDispatcher` | Known | Notes viewer dispatcher |

**Total identified controllers:** 420+

---

## 3. Cover Flow Event Handlers

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x0016AEEC | `HandleWantPopFlow` | Known | Exit Cover Flow view |
| 0x0016AF04 | `HandleFlipToAlbumBackside` | Known | Flip album to show track list |
| 0x0016AF20 | `HandleFlipToAlbumFrontside` | Known | Flip back to album art |
| 0x0016AF3C | `HandleFlowNext` | Known | Navigate to next album |
| 0x0016AF4C | `HandleFlowPrev` | Known | Navigate to previous album |
| 0x0016AF5C | `HandleFlowWheel` | Known | Scroll wheel in Cover Flow |
| 0x0016AF6C | `HandleAlbumSelected` | Known | Album selected (enter track list) |
| 0x0016AF80 | `HandlePlayPause` | Known | Play/pause from Cover Flow |
| 0x0016AF90 | `HandleBacksideSongSelected` | Known | Song selected from backside |

---

## 4. Nike+ Training Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x001459B4 | `HandleTimedWorkoutSelected` | Known | Timed workout selected |
| 0x001459D0 | `HandleDistanceWorkoutSelected` | Known | Distance-based workout |
| 0x001459F0 | `HandleCaloriesWorkoutSelected` | Known | Calorie-based workout |
| 0x00145A10 | `HandleSelectWorkout` | Known | Generic workout selection |
| 0x001BA2F8 | `HandleChooseLink` | Known | Nike+ sensor link setup |
| 0x001BA310 | `HandleChooseCalibrate` | Known | Nike+ sensor calibration |
| 0x001BA328 | `HandleUnlink` | Known | Nike+ sensor unlink |

---

## 5. Settings

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| 0x002306CC | `ToggleSetting_ClassicUI` | Known | Toggle classic/modern UI style |
| 0x002306E4 | `ToggleSetting_SoundCheck` | Known | Volume normalization toggle |
| 0x00230700 | `ToggleSetting_Clicker` | Known | Click wheel audio feedback |
| 0x00230718 | `ToggleSetting_DaylightSavings` | Known | DST toggle |
| 0x00230738 | `ToggleSetting_24HourClock` | Known | 24-hour clock format |
| 0x00230754 | `ToggleSetting_TimeInTitle` | Known | Show time in title bar |
| 0x001EBA8C | `ToggleSetting_Repeat` | Known | Repeat mode cycle |
| 0x001EBAA8 | `ToggleSetting_Shuffle` | Known | Shuffle toggle |
| 0x001EBAC0 | `ToggleSetting_TVOut` | Known | TV output enable |
| 0x001EBAD4 | `ToggleSetting_TVSignal` | Known | TV signal format (PAL/NTSC) |

---

## 6. Firmware Statistics Comparison (Nano 3G vs Classic 7G)

| Metric | Nano 3G (26.1.1.3) | Classic 7G (2.0.4) | Delta |
|--------|--------------------|--------------------|-------|
| Binary Size | 10,792,304 B | 10,599,920 B | +192 KB |
| ARM Functions | 17,473 | 17,721 | −248 |
| Thumb Functions | 5,347 | 5,315 | +32 |
| Total Functions | 22,820 | 23,036 | −216 |
| Total Strings | 53,545 | 55,243 | −1,698 |
| Hidden Features | 78 | 82 | −4 |
| Controllers | 420+ | 407+ | +13 |
| SoC | S5L8702 | S5L8702 | Same |

---

## 7. Features NOT Present (vs. Classic 7G)

| Feature | Nano 3G | Classic 7G | Notes |
|---------|---------|------------|-------|
| Genius Mixes | ❌ | ✅ | Added in later firmware/hardware |
| FM Radio | ❌ | ✅ | Classic has FM tuner hardware |
| Disk Mode (USB MSC) | ❌ | ✅ | Nano 3G uses NAND, no user disk mode |
| Hard Disk | ❌ | ✅ | Nano uses NAND flash storage |
| CE-ATA | ❌ | ✅ | No hard disk controller needed |
| 160GB Storage | ❌ | ✅ | Max 8GB NAND |
| iTunesU | ❌ | ✅ | Added in later Classic builds |

---

## 8. IPSW Versions Available

| IPSW | Firmware File | Size | Notes |
|------|--------------|------|-------|
| iPod_26.1.0.1.ipsw | Firmware-26.9.0.1 | 87.67 MB | Initial release |
| iPod_26.1.0.2.ipsw | Firmware-26.9.0.2 | 87.83 MB | Bug fix |
| iPod_26.1.0.3.ipsw | Firmware-26.9.0.3 | 89.09 MB | Bug fix |
| iPod_26.1.1.ipsw | Firmware-26.9.1 | 89.45 MB | Feature update |
| iPod_26.1.1.2.ipsw | Firmware-26.9.1.2 | 89.51 MB | Bug fix |
| **iPod_26.1.1.3.ipsw** | **Firmware-26.9.1.3** | **89.51 MB** | **Final version (analyzed)** |

---

## 9. Decryption Notes

The Nano 3G has **native wInd3x support** — no device descriptor patch is required:

| Parameter | Value |
|-----------|-------|
| DFU PID | 0x1229 |
| WTF PID | 0x1231 |
| SoC | S5L8702 |
| Exploit | DFUProtoVersion1 (same as Classic) |
| Decryption Time | ~30–60 minutes |
| Recovery Flag | Supported (`-r`) |

The Nano 3G was the first device to receive native wInd3x support, as it uses the same S5L8702 BootROM as the Classic.
