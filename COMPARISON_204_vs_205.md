# iPod Classic RetailOS — Firmware Comparison: 2.0.4 vs 2.0.5

## Summary

| Metric | 2.0.4 (FamilyID 35) | 2.0.5 (FamilyID 38) | Change |
|--------|---------------------|---------------------|--------|
| Binary size | 10,599,920 bytes | 10,634,528 bytes | +34,608 (+0.33%) |
| Strings (>=6 chars) | 55,243 | 55,712 | +469 |
| ARM functions | 17,721 | 17,762 | +41 |
| Thumb functions | 5,312 | 5,402 | +90 |
| Total functions | 23,033 | 23,164 | +131 |
| SoC identifier | `8702` | `8702` | Same |
| IPSW | iPod_35.2.0.4.ipsw | iPod_38.2.0.5.ipsw | — |
| Target hardware | Rev B (MC293/MC297) | Rev C (MD717/MD718) | Later revision |

## String Content Comparison

| Metric | Count |
|--------|-------|
| Strings common to both | 27,261 |
| New strings only in 2.0.5 | 789 |
| Removed strings (only in 2.0.4) | 493 |

### New Strings by Category

| Category | Count |
|----------|-------|
| Other | 635 |
| UI/Controllers | 31 |
| Database | 9 |
| Network/Web | 6 |
| DRM/Security | 4 |
| Settings | 3 |
| Audio/Codec | 2 |

### Removed Strings by Category

| Category | Count |
|----------|-------|
| Other | 434 |
| UI/Controllers | 3 |
| DRM/Security | 1 |

---

## Notable New Strings in 2.0.5

### Other (635 strings)

- `The author or authors of this code dedicate any and all copyright interest in this code to the publi...`
- `Permission is granted to anyone to use this software for any purpose, including commercial applicati...`
- `This software is provided 'as-is', without any express or implied warranty.  In no event will the au...`
- `   !!!!""""####$$$$%%%&&&&''''(((())))*****++++,,,,----.....////000001111222223333344444555556666677...`
- `2. Altered source versions must be plainly marked as such, and must not be misrepresented as being t...`
- `Hiermee beperkt u het maximale koptelefoonvolume tot het door de Europese Unie aanbevolen niveau.`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype\ttgload.c`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\afglobal.c`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftstream.c`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrgload.c`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfdrivr.c`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1load.c`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1cmap.c`
- `Beperk het maximale koptelefoonvolume tot het door de Europese Unie aanbevolen niveau.`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftobjs.c`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrcmap.c`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttcmap.c`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftutil.c`
- `c:\BWA\N25CFirmwareWin-247\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrobjs.c`
- `This will limit the maximum headphone volume to the European Union recommended level.`
- `Limita il volume massimo delle cuffie al livello consigliato dall'Unione europea.`
- `uvwxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxwvuyz{\|}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}\|{zy`
- `3. This notice may not be removed or altered from any source distribution.`
- `zlib.h -- interface of the 'zlib' general purpose compression library`
- `-Your legal department tells you that you have to purchase a license.`
- `Dette begrenser maksimalvolumet for hodetelefonene til niv`
- `ximo de los auriculares al nivel recomendado por la Uni`
- `Copyright (C) 1995-2005 Jean-loup Gailly and Mark Adler`
- ` limitato al livello consigliato dall'Unione europea.`
- `Ceci limite le volume du casque au maximum recommand`

*...and 605 more*

### UI/Controllers (31 strings)

- `27TSilverCntlrTransitionAddonI30TCSettings_EULimitConfirmationE`
- `TSilverSettingsMenuListCntlrTCSettings_EULimitConfirmation`
- `SettingsMenus_DialogNotice_EULimitConfirmation_Layout`
- `SettingsMenus_EUVolume_Confirmation_Screen_Default!`
- `SettingsMenus_EUVolume_Confirmation_Screen_Default`
- `controller.GotoScreen_EUVolumeLimitConfirmation1`
- `SettingsMenus_VolumeLimitEU_Screen_Default,`
- `SettingsMenus_EUVolume_Confirmation_Screen2`
- `SettingsMenus_EUVolume_Confirmation_Screen`
- `SettingsMenus_VolumeLimitEU_Screen_Default`
- `SettingsMenu_ListItem_VolumeLimitEU_Toggle`
- `SettingsMenu_VolumeLimit_Caption_String`
- `controller.GotoScreen_SettingsMenuEU1`
- `GotoScreen_EUVolumeLimitConfirmation`
- `GotoVolumeLimit_or_Lock_or_EU_Screen`
- `controller.GotoScreen_VolumeLimitEU1`
- `SettingsMenus_VolumeLimitEU_Screen*`
- `SettingsMenus_VolumeLimitEU_Screen!`
- `SettingsMenus_VolumeLimitEU_Screen"`
- `SettingsMenus_VolumeLimitEU_Screen,`
- `Settings_VolumeLimitControl_Layout#`
- `SettingsMenu_ListItem_VolumeLimitEU`
- `SettingsMenu_SetVolumeLimit_String`
- `SettingsMenus_VolumeLimitEU_Screen`
- `SettingsMenu_EUVolumeLimit_String`
- `SettingsMenu_EnableLimit_String`
- `SettingsMenu_HighVolume_String`
- `SettingsMenus_Main_Screen$`
- `GotoScreen_SettingsMenuEU`
- `GotoScreen_VolumeLimitEU`

*...and 1 more*

### Database (9 strings)

- `The previous paragraph applies to the deliverable code in SQLite - those parts of the SQLite library...`
- `All of the deliverable code in SQLite has been written from scratch. No code has been taken from oth...`
- `In order to keep SQLite completely free and unencumbered by copyright, all new contributors to the S...`
- `If you feel like you really have to purchase a license for SQLite, Hwaci, the company that employs t...`
- `-You are using SQLite in a jurisdiction that does not recognize the right of an author to dedicate t...`
- `-You want to hold a tangible legal document as evidence that you have the legal right to use and dis...`
- `Obtaining An Explicit License To Use SQLite`
- `Richard Hipp (SQLite)`
- `SQLite Copyright`

### Network/Web (6 strings)

- `                ! !!!!!!!#""######$$$$$%%%&&'((()))**"""""""""""!""""""""##################$#$$$$$$%...`
- `{{{\|\|\|xxyyyyyzzztttuuuuvvvvvjjjlllmmmmmwwwnnnnooooopppppqqqqqrqrrrr`
- `{{{\|\|\|xxxxyyyzzzztttuuuuvvvjjjllllmmmmwwwnnnnoooooppppqqqqq`
- `{{\|\|\|\|xxxyyyzzzztttuuuvvvjjjjlllmmmwwwwnnnnooooppp`
- `{{{\|\|\|xxxyyyzzztttuuuvvvjjlllmmmmwwwwnnnnn`
- `llmmmwwwwwn`

### DRM/Security (4 strings)

- `All of the deliverable code in SQLite has been dedicated to the public domain by the authors. All co...`
- `We are not able to accept patches or changes to SQLite that are not accompanied by a statement such ...`
- `tx3gdrmsp608aavdmp4aesdsD`
- `ltnCniMVxaMVlaVVtceRt`

### Settings (3 strings)

- `ToggleSetting_RecommendedVolumeLimit`
- `30TCSettings_EULimitConfirmation`
- `TCSettings_EULimitConfirmation`

### Audio/Codec (2 strings)

- `Even though SQLite is in the public domain and does not require a license, some users want to obtain...`
- `1. The origin of this software must not be misrepresented; you must not claim that you wrote the ori...`

---

## Notable Removed Strings (were in 2.0.4, gone in 2.0.5)

### Other (434 strings)

- `                ! !!!!!!!#""######$$$$$%%%&&'((()))**"""""""""""!""""""""##################$#$$$$$$%...`
- `   !!!!""""####$$$$%%%&&&&''''(((())))*****++++,,,,----.....////000001111222223333344444555556666677...`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\autofit\afglobal.c`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\truetype\ttgload.c`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftstream.c`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\type1\t1load.c`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\psaux\t1cmap.c`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrgload.c`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\bdf\bdfdrivr.c`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\sfnt\ttcmap.c`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftutil.c`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrobjs.c`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\pfr\pfrcmap.c`
- `c:\BWA\N25CFirmwareWin-75\srcroot\Firmware\Silver\3rdParty\freetype\src\base\ftobjs.c`
- `uvwxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxwvuyz{\|wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww\|{zy`
- `{{{\|\|\|\|}}}qqqsssssttttt~~~~uuuuuvvvvv`
- `list.pid.229442759.delayedselected1`
- `list.pid.229442283.delayedselected1`
- `list.pid.229443485.delayedselected1`
- `list.pid.229442771.delayedselected1`

*...and 414 more*

### UI/Controllers (3 strings)

- `SettingsMenus_Main_Screen!`
- `SettingsMenus_Main_Screen"`
- `SettingsMenus_Main_Screen,`

### DRM/Security (1 strings)

- `tx3gdrmsp608aavdmp4aesdsX{`

---

## Technical Observations

### 1. Same SoC, Same AES Key
Both firmware versions target the S5L8702 (`8702`). The AES encryption key is identical — a firmware extracted from FamilyID 35 IPSW can be decrypted on FamilyID 38 hardware and vice versa.

### 2. Code Layout
The two binaries share 27,261 string values, confirming they are the same codebase. The 789 new strings and 493 removed strings represent actual feature changes between versions.

### 3. Function Count
Firmware 2.0.5 has more functions, suggesting additional features or less aggressive inlining.

### 4. UpdaterFamilyID
- 2.0.4: FamilyID **35** — targets iPod Classic Initial/Rev A/Rev B (2007-2009)
- 2.0.5: FamilyID **38** — targets iPod Classic Rev C (Late 2012)

The Rev C hardware is identical SoC but has minor board changes. FamilyID 38 likely includes:
- Updated storage drivers for newer flash/HDD components
- Revised USB descriptors (DFU PID 0x1250 vs 0x1223)
- Any hardware-specific errata fixes

### 5. DFU Mode Difference
- Rev B (FamilyID 35): DFU presents as USB PID `0x1223`
- Rev C (FamilyID 38): DFU presents as USB PID `0x1250`

This is controlled by the SysCfg in NOR flash, not the firmware binary itself.
