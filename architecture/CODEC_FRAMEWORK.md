# iPod Classic 7G — MeCCA Codec Framework

## Overview

MeCCA (Media Codec/Container Architecture) is Apple's proprietary codec framework in the iPod Classic and Nano 3G firmware. It manages all audio and video decode/encode operations through a registry-based plugin architecture.

Analysis performed using Ghidra and Capstone disassembly.

---

## Pipeline Architecture

```
┌──────────┐    ┌───────────────┐    ┌───────────────┐    ┌───────────┐    ┌─────┐
│ iTunesDB │───>│ DiskReaderTask│───>│ StreamCache   │───>│ MeCCAIO   │───>│ DAC │
│ (track)  │    │  0x00153B48   │    │  0x00228C2C   │    │ 0x001F5354│    │(I2S)│
└──────────┘    └───────────────┘    └───────────────┘    └───────────┘    └─────┘
                                                                │
                                                        ┌───────┴────────┐
                                                        │ Codec Registry │
                                                        │  0x00150AB0    │
                                                        │  (6,028 bytes) │
                                                        └───────┬────────┘
                                                                │
                                          ┌─────────────────────┼───────────────────┐
                                          │                     │                   │
                                     ┌────┴───┐          ┌─────┴──┐          ┌─────┴──┐
                                     │  MP3   │          │  AAC   │          │  ALAC  │
                                     │Fraunhfr│          │FairPlay│          │ Apple  │
                                     └────────┘          └────────┘          └────────┘
```

### Data Flow

```
File on Disk → DiskReaderTask → StreamCacheReadTask → MeCCAIOTask → Decoder → DAC (I2S)
```

---

## Pipeline Tasks

| Task | Address | Size | Role |
|------|---------|------|------|
| DiskReaderTask | 0x00153B48 | 76 B | Reads raw file data from storage |
| StreamCacheReadTask | 0x00228C2C | 52 B | Buffers data in DRAM for continuous playback |
| StreamCacheTimeoutTask | 0x00228A70 | 116 B | Manages buffer timeout/refill thresholds |
| MeCCAIOTask | 0x001F5354 | 276 B | Routes data through codec, outputs PCM |

---

## Codec Registry (0x00150AB0)

The master codec registration function — the largest audio subsystem function at 6,028 bytes.

| Field | Value |
|-------|-------|
| Address | 0x00150AB0 |
| Size | 6,028 bytes |
| Called by | FUN_00152AA8 (192 bytes) |
| References | "AudioCodecs", "AppleLossless", "Audible", "MaximumSampleRate" |

### Registry Callees

| Function | Address | Size | Purpose |
|----------|---------|------|---------|
| FUN_00123EF8 | 0x00123EF8 | 68 B | Codec init helper |
| FUN_00124414 | 0x00124414 | 8 B | Codec registration stub |
| FUN_001508EC | 0x001508EC | 312 B | Codec capability query |
| FUN_00031744 | 0x00031744 | 32 B | String/name handler |
| FUN_001241D4 | 0x001241D4 | 52 B | Format registration |
| FUN_00106358 | 0x00106358 | 1,164 B | Major codec setup function |
| FUN_00124420 | 0x00124420 | 24 B | Codec finalize |
| FUN_0009E100 | 0x0009E100 | 112 B | Error handler |
| FUN_0012415C | 0x0012415C | 88 B | Codec teardown |

---

## Supported Audio Codecs

| Codec | String Evidence | DRM Variant | Notes |
|-------|----------------|-------------|-------|
| MP3 | Fraunhofer IIS license string | No | Licensed Fraunhofer decoder |
| AAC | `adrmmp4a`, `mp4aalac` atoms | Optional | Standard and protected variants |
| Protected AAC | `drmidrms` atoms | Yes (FairPlay) | FairPlay-encrypted AAC |
| ALAC | `AppleLossless`, `alac` | No | Apple Lossless Audio Codec |
| WAV/AIFF | `RIFFWAVEfmt data` | No | PCM container formats |
| Audible | `Audible` format flag | Yes | Audible audiobook format |

### Error Handling

- Codec Error at 0x0092B578: "ERROR: unknownCodec loaded !!!"
- Indicates runtime detection of unregistered codec format — the registry rejects unknown types

---

## Audio Output Path

### Hardware Chain

```
MeCCAIOTask → PCM Buffer → I2S Controller (0x38400000) → Cirrus Logic DAC → Headphone/Line Out
```

### Audio Hardware Addresses

| Component | Address | Purpose |
|-----------|---------|---------|
| I2S/Audio Controller | 0x38400000 | Digital audio interface |
| I2C Controller | 0x3C400000 | Codec chip configuration bus |
| DMA Controller | 0x38600000 | Audio buffer DMA transfers |

### Audio Data Flow

1. Codec decodes compressed audio → PCM samples in DRAM buffer
2. MeCCAIOTask double-buffers: one buffer plays via DMA while next fills
3. I2S controller clocks PCM data at configured sample rate
4. External Cirrus Logic DAC converts to analog
5. Analog signal to headphone jack and 30-pin dock line-out

---

## EQ Processing Chain

EQ processing occurs in the MeCCAIOTask pipeline between decode and I2S output:

| Property | Value |
|----------|-------|
| Processing location | Post-decode, pre-DMA |
| Settings storage | NOR NVRAM (0x2007C000) |
| Math | Fixed-point DSP (ARM9 MAC instructions) |
| Presets | Flat, Bass Booster, Classical, etc. |

Additional audio processing in this stage:
- Volume Limit enforcement
- Sound Check (volume normalization / ReplayGain equivalent)
- Audiobook speed adjustment (time-stretch DSP)

---

## Video Decode Pipeline

### Video Functions

| Function | Address | Size | Callers | Callees | Purpose |
|----------|---------|------|---------|---------|---------|
| MeCCA_VideoBufferMgr | 0x001BC134 | 656 B | 1 | 12 | Video decode buffer manager |
| MeCCAVideoDecode | 0x001BC528 | 144 B | 0 | 4 | Video decoder entry point |

### Video Buffer Manager Callees

| Function | Address | Size | Purpose |
|----------|---------|------|---------|
| FUN_0026FE58 | 0x0026FE58 | 68 B | Buffer allocation |
| FUN_001007DC | 0x001007DC | 8 B | Status check |
| FUN_001D6488 | 0x001D6488 | 304 B | Frame decode |
| FUN_001C1700 | 0x001C1700 | 88 B | Display sync |
| FUN_001BC484 | 0x001BC484 | 132 B | Buffer release |

---

## MeCCA MediaPlayer Controller

| Function | Address | Size | Purpose |
|----------|---------|------|---------|
| MeCCA_MediaPlayer | 0x001B1614 | 96 B | Main playback controller |
| MeCCA_RecordingBuffer | 0x001689A4 | 140 B | Recording input buffer |
| HandlePlayPause | 0x001CE770 | 3,308 B | Master play/pause (79 callees) |

### HandlePlayPause — The Master Audio Controller

At 3,308 bytes with 79 outgoing calls, this is the most complex audio control function. It handles:
- Play/pause state transitions
- Track advance (next/previous)
- Album/chapter navigation
- Audiobook chapter handling
- Crossfade transitions

Referenced strings:
- "HandleAudioPlayPause"
- "HandleAudioNext"
- "HandleAudioNextPressAndHold"
- "HandleAudioNextAlbum"
- "HandleAudioNextChapter"

---

## StreamCache Architecture

The stream cache provides continuous audio during disk spin-down (Classic) or NAND read latency (Nano):

| Component | Address | Size | Role |
|-----------|---------|------|------|
| StreamCacheReadTask | 0x00228C2C | 52 B | Fills buffer from storage |
| StreamCacheTimeoutTask | 0x00228A70 | 116 B | Monitors buffer level |
| DiskReaderTask | 0x00153B48 | 76 B | Low-level storage I/O |

The buffer allows the hard disk to spin down between reads — audio plays from RAM cache until buffer drops below threshold, triggering a disk read-ahead.

---

## GotoNowPlaying Entry Points

Multiple entry points for the Now Playing screen from different contexts:

| Function | Address | Size | Callers | Context |
|----------|---------|------|---------|---------|
| GotoNowPlaying | 0x0021960C | 556 B | 13 | Primary entry (most referenced) |
| GotoNowPlaying | 0x00219F08 | 572 B | 7 | Secondary entry |
| GotoNowPlaying | 0x0021A210 | 584 B | 4 | Tertiary entry |
| GotoNowPlaying | 0x0010CA60 | 408 B | 1 | From music menu |
| GotoNowPlaying | 0x00131728 | 856 B | 0 | From video player |
| GotoNowPlaying | 0x00226E2C | 520 B | 0 | Podcast context |
| GotoNowPlaying | 0x002270F8 | 568 B | 0 | Video context |
| GotoNowPlaying | 0x00230170 | 356 B | 1 | Audiobook context |
| GotoNowPlaying | 0x00235B7C | 436 B | 0 | Genius context |
| GotoNowPlaying | 0x002396A0 | 472 B | 0 | OTG playlist context |

---

## Sources

- Firmware binary analysis (Ghidra, ARM926EJ-S)
- String reference analysis for codec names
- Cross-reference counting for pipeline functions
- Task structure analysis via RTXC API thunks
