# iPod Classic Firmware Comparison: 6G vs 6.5G vs 7G (2.0.4) vs 7G (2.0.5)

## Summary

All four firmware versions run on the **Samsung S5L8702** SoC (ARM926EJ-S) with the
**RTXC v3.2b** RTOS and share the **N25C** build codename and Silver UI framework.

| Metric | 6G (1.1.2) | 6.5G (2.0.1) | 7G (2.0.4) | 7G (2.0.5) |
|--------|-----------|-------------|-----------|-----------|
| Binary size | 9,926,528 | 10,514,000 | 10,599,920 | 10,634,528 |
| Strings (>=6 chars) | 25,696 | 27,560 | 27,754 | 28,050 |
| ARM functions | 16,096 | 17,413 | 17,721 | 17,762 |
| Thumb functions | 5,107 | 5,397 | 5,315 | 5,402 |
| Total functions | 21,203 | 22,810 | 23,036 | 23,164 |
| UpdaterFamilyID | 24 | 33 | 35 | 38 |
| Release year | 2007 | 2008 | 2009 | 2012 |
| Build | N25FirmwareWin-435 | N25BFirmwareWin-93 | N25CFirmwareWin-75 | N25CFirmwareWin-247 |
| SoC | S5L8702 | S5L8702 | S5L8702 | S5L8702 |
| Storage | 80/160GB HDD | 120GB HDD | 160GB HDD | 160GB HDD |
| Target HW | Initial | Rev A | Rev B | Rev C |

---

## Feature Comparison

| Feature | 6G (1.1.2) | 6.5G (2.0.1) | 7G (2.0.4) | 7G (2.0.5) |
|---------|-----------|-------------|-----------|-----------|
| Cover Flow | ✅ | ✅ | ✅ | ✅ |
| Demo Mode | ✅ | ✅ | ✅ | ✅ |
| Disk Mode | ✅ | ✅ | ✅ | ✅ |
| EU Volume Limit | — | — | — | ✅ |
| FM Radio | ✅ | ✅ | ✅ | ✅ |
| FairPlay DRM | ✅ | ✅ | ✅ | ✅ |
| FreeType2 Fonts | ✅ | ✅ | ✅ | ✅ |
| Games | ✅ | ✅ | ✅ | ✅ |
| Genius | — | ✅ | ✅ | ✅ |
| Genius Mixes | — | — | ✅ | ✅ |
| MeCCA (Codec Framework) | ✅ | ✅ | ✅ | ✅ |
| Nike+ iPod | ✅ | ✅ | ✅ | ✅ |
| Photos | ✅ | ✅ | ✅ | ✅ |
| RTXC (RTOS) | ✅ | ✅ | ✅ | ✅ |
| SQLite Database | — | ✅ | ✅ | ✅ |
| Silver UI Framework | ✅ | ✅ | ✅ | ✅ |
| USB Audio | ✅ | ✅ | ✅ | ✅ |
| Video Playback | ✅ | ✅ | ✅ | ✅ |
| Voice Memos | ✅ | ✅ | ✅ | ✅ |

---

## String Content Comparison

| Comparison | Common | Only in A | Only in B |
|-----------|--------|-----------|-----------|
| 6G vs 6.5G | 25,104 | 592 | 2,456 |
| 6.5G vs 7G (2.0.4) | 26,857 | 703 | 897 |
| 7G (2.0.4) vs 7G (2.0.5) | 27,261 | 493 | 789 |
| Common to ALL four | 24,526 | — | — |

---

## New Strings in 6.5G (Not in 6G)

Total new strings: **2,456**

| Category | Count |
|----------|-------|
| Other | 2114 |
| Genius | 213 |
| Database | 37 |
| Handlers | 33 |
| Settings | 32 |
| Build Paths | 12 |
| UI/Controllers | 6 |
| FM Radio | 5 |
| Screens | 3 |
| DRM/Security | 1 |

### Genius (213 strings)

- ` Genius`
- ` Genius `
- ` Genius %lu`
- ` Genius, `
- ` Genius.`
- ` iTunes pour activer Genius.`
- ` keskipainiketta painettuna, kunnes Genius-vaihtoehto tulee n`
- ` muligheden Genius vises.`
- ` se seznam Genius`
- ` visas alternativet Genius.`
- ` vise Genius-funksjonen.`
- `19TGeniusLoadingCntlr`
- `27TSilverCntlrTransitionAddonI19TGeniusLoadingCntlrE`
- `27TSilverCntlrTransitionAddonI28TSilverMediaListCntlr_GeniusE`
- `28TSilverMediaListCntlr_Genius`
- `A Genius aktiv`
- `A Genius nem el`
- `A Genius olyan dalokat j`
- `A criar lista Genius`
- `Aloita Genius`
- *...and 193 more*

### UI/Controllers (6 strings)

- `TCVoiceMemosAlert`
- `TCVoiceMemosTCVoiceMemosAlert`
- `TContextualMenuCntlr`
- `TContextualMenuCntlrTCExtrasMenuTSilverCntlrTGeniusLoadingCntlr`
- `TSilverMediaListCntlr_Genius`
- `TSilverMediaListCntlr_GeniusTSilverMediaListCntlr_NestedPlaylists`

### Handlers (33 strings)

- `11TMsgHandler`
- `32iMAImageCacheClientEventHandlers`
- `HandleAudiobookFaster`
- `HandleAudiobookNormal`
- `HandleAudiobookSlower`
- `HandleCancel`
- `HandleFinishRecording`
- `HandleLoadingCancelled`
- `HandleMenuSelection`
- `HandleMikeyAllUp`
- `HandleMikeyCenter`
- `HandleMikeyCenterDoubleClick`
- `HandleMikeyCenterDoubleClickAndHold`
- `HandleMikeyCenterPressAndHold`
- `HandleMikeyCenterTripleClick`
- `HandleMikeyVolumeDown`
- `HandleMikeyVolumeDownUp`
- `HandleMikeyVolumeUp`
- `HandleMikeyVolumeUpUp`
- `HandlePushContextualMenu`
- *...and 13 more*

### Settings (32 strings)

- `11TCEQSetting`
- `13TCSettings_EQ`
- `19TCSettings_MainMenu`
- `20TCSettings_MusicMenu`
- `20TPhotosSettingsCntlr`
- `21TCSettings_Brightness`
- `22TCSettings_VolumeLimit`
- `25TCSettings_BacklightTimer`
- `25TSilverSettingsVideoCntlr`
- `27TCSettings_ResetAllSettings`
- `27TSilverCntlrTransitionAddonI11TCEQSettingE`
- `27TSilverCntlrTransitionAddonI13TCSettings_EQE`
- `27TSilverCntlrTransitionAddonI19TCSettings_MainMenuE`
- `27TSilverCntlrTransitionAddonI20TCSettings_MusicMenuE`
- `27TSilverCntlrTransitionAddonI20TPhotosSettingsCntlrE`
- `27TSilverCntlrTransitionAddonI21TCSettings_BrightnessE`
- `27TSilverCntlrTransitionAddonI22TCSettings_VolumeLimitE`
- `27TSilverCntlrTransitionAddonI25TCSettings_BacklightTimerE`
- `27TSilverCntlrTransitionAddonI25TSilverSettingsVideoCntlrE`
- `27TSilverCntlrTransitionAddonI27TCSettings_ResetAllSettingsE`
- *...and 12 more*

### FM Radio (5 strings)

- `11TRadioCntlr`
- `23TCSettings_RadioRegions`
- `27TSilverCntlrTransitionAddonI11TRadioCntlrE`
- `27TSilverCntlrTransitionAddonI23TCSettings_RadioRegionsE`
- `PFMd2(-`

---

## New Strings in 7G 2.0.4 (Not in 6.5G)

Total new strings: **897**

| Category | Count |
|----------|-------|
| Other | 763 |
| Genius | 104 |
| Build Paths | 12 |
| Screens | 8 |
| DRM/Security | 4 |
| Handlers | 3 |
| UI/Controllers | 3 |

### UI/Controllers (3 strings)

- `TSilverMediaListCntlr_GeniusMixes`
- `TSilverMediaListCntlr_iTunesU`
- `TSilverMediaListCntlr_iTunesUEpisodes`

### Handlers (3 strings)

- `HandleSelectMix`
- `HandleiTunesUHilited`
- `HandleiTunesUSelected`

### Build Paths (12 strings)

- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fto`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\base\fts`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftu`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfd`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrc`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrg`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfro`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttc`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1`

---

## New Strings in 7G 2.0.5 (Not in 7G 2.0.4)

Total new strings: **789**

| Category | Count |
|----------|-------|
| Other | 726 |
| EU Volume Limit | 27 |
| Database | 12 |
| Build Paths | 12 |
| Settings | 8 |
| UI/Controllers | 2 |
| Screens | 1 |
| DRM/Security | 1 |

### EU Volume Limit (27 strings)

- `27TSilverCntlrTransitionAddonI30TCSettings_EULimitConfirmationE`
- `30TCSettings_EULimitConfirmation`
- `Do you want to enable EU Volume Limit?`
- `EU Volume Limit`
- `GotoScreen_EUVolumeLimitConfirmation`
- `GotoScreen_VolumeLimitEU`
- `GotoVolumeLimit_or_Lock_or_EU_Screen`
- `Limit hlasitosti (EU)`
- `SettingsMenu_EUVolumeLimit_String`
- `SettingsMenu_ListItem_VolumeLimitEU`
- `SettingsMenu_ListItem_VolumeLimitEU_Toggle`
- `SettingsMenus_DialogNotice_EULimitConfirmation_Layout`
- `SettingsMenus_EUVolume_Confirmation_Screen`
- `SettingsMenus_EUVolume_Confirmation_Screen2`
- `SettingsMenus_EUVolume_Confirmation_Screen_Default`
- *...and 12 more*

### UI/Controllers (2 strings)

- `TCSettings_EULimitConfirmation`
- `TSilverSettingsMenuListCntlrTCSettings_EULimitConfirmation`

### Settings (8 strings)

- `GotoScreen_SettingsMenuEU`
- `SettingsMenu_EnableLimit_String`
- `SettingsMenu_HighVolume_String`
- `SettingsMenu_SetVolumeLimit_String`
- `SettingsMenu_VolumeLimit_Caption_String`
- `Settings_VolumeLimitControl_Layout#`
- `ToggleSetting_RecommendedVolumeLimit`
- `controller.GotoScreen_SettingsMenuEU1`

### Build Paths (12 strings)

- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ft`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ft`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ft`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdf`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfr`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfr`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfr`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\tt`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\truetyp`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t`

---

## Technical Observations

### 1. Same SoC, Same AES Key
All four firmware versions target the S5L8702. The AES GID key is identical —
any firmware decrypted from one revision can be analyzed to understand all revisions.

### 2. Firmware Size Growth

| Transition | Size Delta | Percentage | Primary Cause |
|-----------|-----------|-----------|---------------|
| 6G → 6.5G | +587,472 bytes | +5% | Genius system + SQLite |
| 6.5G → 7G (2.0.4) | +85,920 bytes | +0% | Bug fixes, minor additions |
| 7G (2.0.4) → 7G (2.0.5) | +34,608 bytes | +0% | EU Volume Limit, FreeType2 update |

### 3. Function Count Growth

| Transition | New Functions | Notes |
|-----------|--------------|-------|
| 6G → 6.5G | +1,607 | Genius playlist generation, Genius Mixes UI |
| 6.5G → 7G (2.0.4) | +226 | Minor additions |
| 7G (2.0.4) → 7G (2.0.5) | +128 | EU Volume Limit UI + updated FreeType2 |

### 4. Hardware Revision Differences

| Aspect | 6G Initial | 6.5G Rev A | 7G Rev B | 7G Rev C |
|--------|-----------|-----------|---------|---------|
| HDD | 80/160GB dual-platter | 120GB single-platter | 160GB single-platter | 160GB single-platter |
| DFU PID | 0x1223 | 0x1223 | 0x1223 | 0x1250 |
| WTF PID | 0x1241 | 0x1245 | 0x1247 | 0x1250 |
| Model Numbers | MB029/MB147 | MB562/MB565 | MC293/MC297 | MD717/MD718 |

### 5. Genius Timeline

- **6G (1.1.2):** No Genius — released before iTunes 8 (Sept 2008)
- **6.5G (2.0.1):** Genius Playlists + Genius Mixes added (first firmware with iTunes 8 integration)
- **7G (2.0.4/2.0.5):** Genius retained, no further changes

### 6. EU Volume Limit Timeline

- **6G, 6.5G, 7G 2.0.4:** No EU Volume Limit
- **7G 2.0.5:** EU Volume Limit added (compliance with EU regulation 2006/95/EC)
