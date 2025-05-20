# Changelog

All notable changes to steganography-multi-tool are documented here.

### [2025-12-14]
- fix: prevent duplicate event emission during rapid retry bursts

### [2025-12-18]
- feat: add support for custom timeout configuration via CLI flags

### [2025-12-29]
- fix: resolve race condition during concurrent worker initialization

### [2026-01-07]
- perf: optimize memory allocation in buffer pool

### [2026-01-25]
- fix: patch edge-case buffer truncation in stream reader

### [2026-02-08]
- feat: add support for custom timeout configuration via CLI flags

### [2026-02-10]
- security: enforce strict bounds checking on dynamic byte slices

### [2026-02-14]
- test: add fuzzing harness for packet decoding routine

### [2026-02-17]
- feat: add support for custom timeout configuration via CLI flags

### [2026-03-20]
- fix: ensure file descriptors are properly closed on error exits

### [2026-03-23]
- fix: resolve race condition during concurrent worker initialization

### [2026-04-04]
- style: clean up trailing whitespace and fix alignment

### [2026-04-13]
- security: sanitize input strings to mitigate format string risks

### [2026-05-09]
- refactor: use enum types for status codes instead of magic numbers

### [2026-05-12]
- fix: handle malformed HTTP header parsing without crashing

### [2026-06-10]
- fix: patch edge-case buffer truncation in stream reader

### [2026-06-27]
- docs: add architecture diagram and sequence flow explanation

### [2026-07-02]
- test: add fuzzing harness for packet decoding routine

### [2026-07-28]
- test: verify backward compatibility with legacy message format

### [2026-08-26]
- fix: resolve memory leak in idle connection reaper

### [2026-09-03]
- fix: handle malformed HTTP header parsing without crashing

