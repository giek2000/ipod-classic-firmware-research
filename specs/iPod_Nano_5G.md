# iPod Nano 5th Generation (8/16GB) — RetailOS Firmware Feature Specification

## Document Summary

| Field | Value |
|-------|-------|
| **Firmware** | RetailOS 1.0.2 |
| **IPSW** | iPod_1.0.2_34A20020.ipsw |
| **Device** | iPod Nano 5th Generation (2009, Click Wheel, Camera, Pedometer, FM Radio) |
| **UpdaterFamilyID** | 34 |
| **Build Codename** | N33 |
| **Build** | N33FirmwareWin-261 |
| **Binary Size** | 7,286,720 bytes (6.95 MB) |
| **ARM Code Start** | 0x800 |
| **Total Strings (>=6)** | 48,164 |
| **Function Prologues** | 32,358 (ARM: 1,649, Thumb: 30,709) |
| **SoC** | Samsung S5L8730 |
| **Architecture** | ARM Cortex-A8 (ARMv7) |
| **Encrypted** | Yes (HW AES-128-CBC, GID key) |
| **SHA-256** | `3269d9eda2e7e7c406d7bfd6895bd83f27603f03a2c9f52c6cf415924e41af81` |
| **DFU USB PID** | 0x1231 |

---

## Hardware Overview

| Component | Specification |
|-----------|---------------|
| SoC | Samsung S5L8730 (ARM Cortex-A8, ARMv7) |
| RAM | 64 MB |
| Storage | 8GB / 16GB NAND Flash |
| Display | 240×376 RGB565, 2.2" TFT |
| Camera | 640×480 VGA video recording |
| Sensors | Accelerometer, Pedometer |
| Radio | FM Tuner with RDS, Live Pause |
| Audio | Cirrus Logic codec (I2S) |
| Connectivity | USB 2.0, 30-pin dock |
| Form Factor | Nano slim, click wheel + video camera |

---

## Significance

The iPod Nano 5G is architecturally significant as a major departure from the ARM926EJ-S lineage:

1. **First ARM Cortex-A8 iPod** — Transitions from the ARMv5TEJ (ARM926EJ-S) architecture used in the Nano 3G/4G and Classic 6G/7G to the ARMv7 Cortex-A8 core. This represents a full generational leap in instruction set architecture.

2. **First Thumb-dominant firmware** — With 30,709 Thumb-mode functions versus only 1,649 ARM-mode functions (94.9% Thumb), this is the first iPod firmware compiled predominantly in Thumb-2 mode. The Nano 3G, by comparison, was 76.5% ARM-mode. This indicates a compiler toolchain modernization targeting the Thumb-2 instruction set introduced with ARMv7.

3. **Smaller binary, greater capability** — Despite being 3.5 MB smaller than the Nano 3G binary (7.29 MB vs. 10.79 MB), the Nano 5G contains significantly more event handlers (484 vs. 289), more screens (721 vs. 613), and substantially more user-facing features. Thumb-2 code density enables this reduction.

4. **First iPod with integrated camera** — A dedicated camera subsystem with video recording, media management, and playback.

5. **First iPod with hardware pedometer** — A built-in step counter with persistent logging, independent of the Nike+ accessory system.

6. **First FM Radio with Live Pause** — Buffered FM reception with RDS metadata processing enables pause-and-resume of live radio.

7. **First iPod with VoiceOver accessibility** — A text-to-speech accessibility layer, representing Apple's first implementation of screen reader technology on a dedicated media player.

---

## 1. Hidden/Disabled Features

| Offset | Symbol | Classification | Description |
|--------|--------|----------------|-------------|
| — | `RTXCbug` | Hidden | Interactive RTXC RTOS kernel debugger with command prompt and object inspector |
| — | `Debug_MainMenu_Screen` | Hidden | Debug menu main screen, accessible via hidden key sequence |
| — | `TCDemoMode` | Hidden | Apple Store retail demonstration/kiosk mode |
| — | `WaveFileDebugTask` | Hidden | Audio subsystem debug recording task for PCM capture |
| — | `MockupMode/MockupMode.xml` | Hidden | Developer UI prototyping tool with navigation screen |
| — | `Channel UnitTests` | Hidden | Built-in unit test execution logging channel |
| — | `Channel TestLogging` | Hidden | Development test logging infrastructure |

---

## 2. Camera System

The Nano 5G introduced a 640×480 video camera — the first camera on any iPod. Analysis reveals a full camera controller hierarchy:

| Symbol | Classification | Description |
|--------|----------------|-------------|
| `TCCamera` | Known | Primary camera controller |
| `TCCameraInitial` | Known | Camera initialization state machine |
| `TCCameraInitial_InitialLayoutIsActive` | Known | Camera active layout state |
| `TCCameraInitial_InitialLayoutIsAppNotInitialized` | Known | Camera pre-initialization state |
| `TCCameraInitial_InitialLayoutIsDiskFull` | Known | Disk full error state for recording |
| `TCCameraAllVideosList` | Known | All recorded videos browser |
| `TCCameraLocalMediaList` | Known | Local camera media list controller |
| `TCCameraMediaList_Base` | Known | Base media list for camera content |
| `TCCameraMediaList_Base_DoDeleteAll` | Known | Delete all camera media action |
| `TCCameraMediaList_Base_DoDeleteItem` | Known | Delete individual camera item |
| `TCCameraDeleteDialog` | Known | Single item deletion confirmation |
| `TCCameraDeleteAllDialog` | Known | Bulk deletion confirmation |

The camera subsystem includes state handling for active recording, disk space management, and a dedicated media browser separate from the Photos application.

---

## 3. Pedometer System

The Nano 5G was the first iPod with a built-in hardware pedometer (accelerometer-based step counter), independent of the Nike+ external sensor:

| Symbol | Classification | Description |
|--------|----------------|-------------|
| `TLogPedDiskWritingTask` | Known | RTOS task for persistent step data logging to NAND |
| `TTrainerApp_LocaleChangedLoadingTask` | Known | Trainer app locale reloading task |
| `Nike+ iPod` (feature flag) | Known | Nike+ integration layer (sensor + built-in pedometer) |

The pedometer system operates through the TLogPedDiskWritingTask, which continuously logs step data to persistent storage. This is distinct from the Nike+ wireless sensor ecosystem — the built-in accelerometer provides always-on step counting without an external accessory.

---

## 4. Shake Gesture System

The Nano 5G introduced accelerometer-based shake gestures for shuffle control:

| Symbol | Classification | Description |
|--------|----------------|-------------|
| `ToggleSetting_Shake` | Known | Enable/disable shake-to-shuffle gesture |
| `HandleGestureShake` (inferred) | Known | Shake gesture event handler |

The ToggleSetting_Shake setting allows users to enable or disable the shake gesture. When enabled, a physical shake of the device triggers shuffle. The accelerometer sensitivity and duration thresholds are configurable internally.

---

## 5. VoiceOver Accessibility

The Nano 5G was the first iPod with VoiceOver — Apple's text-to-speech accessibility system:

| Symbol | Classification | Description |
|--------|----------------|-------------|
| `ToggleSetting_VoiceFeedback` | Known | Enable/disable VoiceOver spoken feedback |
| `HandleVoiceOver*` (handler family) | Known | VoiceOver event handler set |

VoiceOver on the Nano 5G provides spoken menu names, song titles, artist names, and playlist information. The ToggleSetting_VoiceFeedback controls the feature globally. This implementation predates VoiceOver on the iPod Shuffle 3G (which used pre-rendered speech) by using a real-time synthesis approach.

---

## 6. FM Radio System

The Nano 5G introduced FM radio with Live Pause — the first buffered radio implementation on an iPod:

| Symbol | Classification | Description |
|--------|----------------|-------------|
| `RadioTask` | Known | Primary FM radio RTOS task |
| `MeCCABufferedRDSUpdateTask` | Known | Buffered RDS (Radio Data System) metadata processing task |
| `MeCCARecordingTask` | Known | Audio recording task (shared with Voice Memos) |

The FM radio subsystem operates through two dedicated RTOS tasks. RadioTask manages tuner control, frequency selection, and audio routing. MeCCABufferedRDSUpdateTask processes RDS data (station names, program information) through a buffer, enabling the Live Pause feature where radio audio is continuously recorded to a circular buffer and can be paused/resumed without losing content.

---

## 7. Settings (24)

| # | Setting | New vs. Nano 3G |
|---|---------|-----------------|
| 1 | ShowSetting_About | Existing |
| 2 | ShowSetting_DateAndTime | Existing |
| 3 | ShowSetting_Legal | Existing |
| 4 | ShowSetting_VolumeLimit | Existing |
| 5 | ToggleSetting_24HourClock | Existing |
| 6 | ToggleSetting_Alarm | Existing |
| 7 | ToggleSetting_Audiobook | Existing |
| 8 | ToggleSetting_Clicker | Existing |
| 9 | ToggleSetting_Crossfade | **New** |
| 10 | ToggleSetting_DaylightSavings | Existing |
| 11 | ToggleSetting_EnergySaver | **New** |
| 12 | ToggleSetting_FontSize | **New** |
| 13 | ToggleSetting_MonoAudio | **New** |
| 14 | ToggleSetting_PreviewPanel | **New** |
| 15 | ToggleSetting_Repeat | Existing |
| 16 | ToggleSetting_Rotate | **New** |
| 17 | ToggleSetting_Shake | **New** |
| 18 | ToggleSetting_Shuffle | Existing |
| 19 | ToggleSetting_SortBy | Existing |
| 20 | ToggleSetting_SoundCheck | Existing |
| 21 | ToggleSetting_TVOut | Existing |
| 22 | ToggleSetting_TVSignal | Existing |
| 23 | ToggleSetting_TimeInTitle | Existing |
| 24 | ToggleSetting_VoiceFeedback | **New** |

**New settings in Nano 5G (not present in Nano 3G):**
- **Shake** — Accelerometer shake-to-shuffle gesture toggle
- **Rotate** — Screen rotation based on accelerometer orientation
- **EnergySaver** — Display power management / auto-dimming
- **FontSize** — Accessibility font size adjustment
- **MonoAudio** — Accessibility mono audio downmix (hearing impairment support)
- **Crossfade** — Audio crossfade between tracks
- **VoiceFeedback** — VoiceOver text-to-speech accessibility
- **PreviewPanel** — Cover Flow / album art preview panel toggle

---

## 8. RTOS Tasks (30)

| # | Task Name | New vs. Classic 7G |
|---|-----------|-------------------|
| 1 | USBAudioTask | Existing |
| 2 | ATAWorkLoopIRQTask | Existing |
| 3 | ATAWorkLoopTask | Existing |
| 4 | **AlarmTask** | **New** |
| 5 | BootTask | Existing |
| 6 | ChargeMgmtTask | Existing |
| 7 | DiskMgrTask | Existing |
| 8 | DiskReaderTask | Existing |
| 9 | FirewireTask | Existing |
| 10 | **GeniusMixesTask** | **New** |
| 11 | HoldSwitchTask | Existing |
| 12 | HostOSTask | Existing |
| 13 | InputBufferLoadTask | Existing |
| 14 | LowBattDebounceTask | Existing |
| 15 | MainAppTask | Existing |
| 16 | **MeCCABufferedRDSUpdateTask** | **New** |
| 17 | MeCCAInputTask | Existing |
| 18 | MeCCAOutputTask | Existing |
| 19 | **MeCCARecordingTask** | **New** |
| 20 | MikeyTask | Existing |
| 21 | **RadioTask** | **New** |
| 22 | StreamCacheReadTask | Existing |
| 23 | **TLogPedDiskWritingTask** | **New** |
| 24 | TMusicLoadingTask | Existing |
| 25 | TPodMediaPlayer Task | Existing |
| 26 | TPodMediaPlayerFileUpdate Task | Existing |
| 27 | **TTrainerApp_LocaleChangedLoadingTask** | **New** |
| 28 | Terminator Task | Existing |
| 29 | TouchwheelTask | Existing |
| 30 | USBDeviceTask | Existing |

**New RTOS tasks (not present in Classic 7G):**
- **AlarmTask** — Dedicated alarm clock management (Classic uses MainAppTask)
- **GeniusMixesTask** — Background Genius Mixes generation
- **MeCCABufferedRDSUpdateTask** — FM Radio RDS metadata buffering for Live Pause
- **MeCCARecordingTask** — Audio recording pipeline (camera video, voice memos)
- **RadioTask** — FM tuner control and audio routing
- **TLogPedDiskWritingTask** — Pedometer step data persistent logging
- **TTrainerApp_LocaleChangedLoadingTask** — Nike+ trainer locale reload

---

## 9. Firmware Statistics Comparison

| Metric | Nano 5G (1.0.2) | Classic 7G (2.0.4) | Nano 3G (1.1.3) |
|--------|-----------------|-------------------|-----------------|
| Binary Size | 7,286,720 B | 10,599,920 B | 10,792,304 B |
| SoC | S5L8730 | S5L8702 | S5L8702 |
| Architecture | Cortex-A8 (ARMv7) | ARM926EJ-S (ARMv5) | ARM926EJ-S (ARMv5) |
| ARM Functions | 1,649 | 17,721 | 17,473 |
| Thumb Functions | 30,709 | — | 5,347 |
| Total Functions | 32,358 | 17,721 | 22,820 |
| Total Strings | 48,164 | 55,243 | 53,545 |
| Controllers | 129 | 321 | 420+ |
| Handlers | 484 | 289 | 289 |
| Screens | 721 | 613 | 613 |
| Settings | 24 | — | 10 |
| Tasks | 30 | — | — |
| Channels | 21 | — | — |
| Features Detected | 27 | 19 | 17 |

**Observations:**
- The Nano 5G is 32% smaller in binary size than the Nano 3G despite having 60% more features
- Thumb-2 dominance (94.9%) versus Nano 3G ARM dominance (76.5%) accounts for the size reduction
- Handler count increased 67% (484 vs. 289), reflecting the richer input model (accelerometer, camera)
- Screen count increased 18% (721 vs. 613), reflecting camera UI, pedometer UI, and radio UI additions

---

## 10. IPSW Versions

| IPSW | UpdaterFamilyID | Notes |
|------|-----------------|-------|
| iPod_1.0.1_34A10006.ipsw | 34 | Initial release firmware |
| **iPod_1.0.2_34A20020.ipsw** | **34** | **Final version (analyzed)** |

---

## 11. Decryption and Build Information

| Parameter | Value |
|-----------|-------|
| Build | N33FirmwareWin-261 |
| SoC | S5L8730 |
| DFU USB PID | 0x1231 |
| Architecture | ARMv7 (Cortex-A8) |
| Code Start | 0x800 (2048 bytes) |
| Header | 8730 |
| OS Type | RTXC |

Build path references observed in the binary confirm Windows-hosted cross-compilation:
- `c:/bwa/N33FirmwareWin-261/srcroot/Firmware/Shared/Services/`

The N33 codename and S5L8730 SoC header identify this as the Nano 5th Generation hardware platform.
