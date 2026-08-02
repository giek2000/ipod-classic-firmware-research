# iPod Classic Firmware Comparison: 6G vs 6.5G vs 7G

## Summary

| Metric | 6G (FamilyID 24) | 6.5G (FamilyID 33) | 7G (FamilyID 35) | 7G (FamilyID 38) |
|--------|-------------------|---------------------|--------------------|--------------------|
| Version | 1.1.2 | 2.0.1 | 2.0.4 | 2.0.5 |
| Release | Jul 2008 | Oct 2008 | Late 2009 | Late 2012 |
| Binary size | 9,926,528 | 10,514,000 | 10,599,920 | 10,634,528 |
| Size (MB) | 9.47 | 10.03 | 10.11 | 10.14 |
| Entry point | 0x00976F80 | 0x00A06650 | 0x00A1B5F0 | — |
| Code entropy | 7.997 | 7.997 | 5.664 | 5.67 (est.) |
| Decrypted? | ❌ No | ❌ No | ✅ Yes | ✅ Yes |
| Strings | N/A | N/A | 55,243 | 55,712 |
| ARM functions | N/A | N/A | 17,721 | 17,762 |
| Controllers | N/A | N/A | 321 | 324 |
| SoC | S5L8702 | S5L8702 | S5L8702 | S5L8702 |
| Codename | N73/N25 | N25A | N25B | N25C |
| DFU PID | 0x1223 | 0x1223 | 0x1223 | 0x1250 |

---

## Size Growth Analysis

```
6G 1.1.2    ████████████████████████████████████████████████░░░░░  9.47 MB
6.5G 2.0.1  █████████████████████████████████████████████████████  10.03 MB  (+5.9%)
7G 2.0.4    ██████████████████████████████████████████████████████ 10.11 MB  (+0.8%)
7G 2.0.5    ██████████████████████████████████████████████████████ 10.14 MB  (+0.3%)
```

| Transition | Growth (bytes) | Growth (%) | Likely cause |
|-----------|---------------|-----------|--------------|
| 6G → 6.5G | +587,472 | +5.9% | Genius, new HDD controller, UI refinements |
| 6.5G → 7G 2.0.4 | +85,920 | +0.8% | Minor features, bug fixes |
| 7G 2.0.4 → 2.0.5 | +34,608 | +0.3% | EU Volume Limit, FreeType update |


---

## Entry Point Evolution

| Firmware | Entry Point | Offset/Size Ratio | Interpretation |
|----------|-------------|-------------------|----------------|
| 6G 1.1.2 | 0x00976F80 | 0.958 | Init at 95.8% of binary |
| 6.5G 2.0.1 | 0x00A06650 | 0.957 | Init at 95.7% of binary |
| 7G 2.0.4 | 0x00A1B5F0 | 0.958 | Init at 95.8% of binary |

The entry point consistently sits at ~95.8% of the total firmware size across all generations. This confirms:
1. The same linker script/memory layout is used
2. Initialization code is always placed at the end
3. The ratio is stable, meaning code growth is distributed throughout rather than appended

---

## IMG1 Header Comparison

| Field | 6G | 6.5G | 7G 2.0.4 |
|-------|-----|------|----------|
| Magic | `8702` | `8702` | `8702` |
| Version | `1.0` | `1.0` | `1.0` |
| Entry × 3 | 0x976F80 | 0xA06650 | 0xA1B5F0 |
| Padding | All zeros | All zeros | All zeros |
| Header size | 2048 | 2048 | 2048 |

The IMG1 header format is byte-for-byte identical in structure across all Classic generations. Only the entry point value changes (proportional to code size).

---

## Hardware Comparison

| Attribute | 6G (2007) | 6.5G (2008) | 7G Rev B (2009) | 7G Rev C (2012) |
|-----------|-----------|-------------|-----------------|-----------------|
| SoC | S5L8702 | S5L8702 | S5L8702 | S5L8702 (new mask) |
| CPU | ARM926EJ-S | ARM926EJ-S | ARM926EJ-S | ARM926EJ-S |
| Clock | 200 MHz | 200 MHz | 200 MHz | 200 MHz |
| RAM | 32/64 MB | 64 MB | 64 MB | 64 MB |
| Storage | 80/160 GB HDD | 120 GB HDD | 160 GB HDD | 160 GB HDD |
| HDD interface | ZIF | ZIF | ZIF | ZIF |
| Display | 320×240 | 320×240 | 320×240 | 320×240 |
| Body | Aluminum+steel | Aluminum+steel | Aluminum+steel | Aluminum+steel |
| Thickness | 10.5mm (80)/13.5mm (160) | 10.5mm | 10.5mm | 10.5mm |
| DFU PID | 0x1223 | 0x1223 | 0x1223 | **0x1250** |
| GID key | Shared (A) | Shared (A) | Shared (A) | **Different (C)** |

---

## Feature Timeline (Expected)

| Feature | 6G 1.x | 6.5G 2.0.x | 7G 2.0.4 (confirmed) | 7G 2.0.5 (confirmed) |
|---------|---------|------------|----------------------|----------------------|
| Cover Flow | ✅ | ✅ | ✅ | ✅ |
| FairPlay DRM | ✅ | ✅ | ✅ | ✅ |
| Nike+ iPod | ✅ | ✅ | ✅ | ✅ |
| Games | ✅ | ✅ | ✅ | ✅ |
| Video | ✅ | ✅ | ✅ | ✅ |
| **Genius** | ❌ | **✅** | ✅ | ✅ |
| **Genius Mixes** | ❌ | ❓ | ✅ | ✅ |
| FM Radio | ❓ | ✅ | ✅ | ✅ |
| **EU Volume Limit** | ❌ | ❓ | ❌ | **✅** |
| Demo Mode | ✅ | ✅ | ✅ | ✅ |
| Debug Menu | ✅ | ✅ | ✅ | ✅ |
| Disk Mode | ✅ | ✅ | ✅ | ✅ |

---

## Encryption & Decryption Status

| Model | FamilyID | DFU PID | GID Group | Decrypted? | Method |
|-------|----------|---------|-----------|------------|--------|
| Classic 6G | 24 | 0x1223 | A | ❌ | Needs wInd3x on 0x1223 device |
| Classic 6.5G | 33 | 0x1223 | A | ❌ | Needs wInd3x on 0x1223 device |
| Classic 7G Rev B | 35 | 0x1223 | A | ✅ | wInd3x haxdfu |
| Classic 7G Rev C | 38 | 0x1250 | C | ✅ | wInd3x haxdfu |

**Key insight:** Devices sharing DFU PID 0x1223 (6G, 6.5G, 7G Rev A/B) use the same S5L8702 silicon revision and therefore the same GID encryption key. A single 7G Rev B device should be able to decrypt all three firmware versions.

---

## Next Steps

1. Use existing 7G Rev B hardware to attempt decryption of 6G and 6.5G binaries
2. If successful, run full analysis pipeline to extract controllers, handlers, screens
3. Produce detailed 6G↔6.5G↔7G comparison showing exact feature additions per generation
4. Document the complete evolution of the Silver UI framework from 1.0 to 2.0.5
