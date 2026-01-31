# Shadow Stream

Shadow Stream is a Go-based systems project that streams **raw RGB video frames** from FFmpeg directly into Go using stdout pipes.

The goal is to work at the **byte level**, not behind high-level video libraries.

---

## Why this project exists

Most video projects hide complexity behind OpenCV or media frameworks.  
Shadow Stream does the opposite.

It focuses on:
- Understanding how video decoding *actually* works
- Streaming raw pixel data across process boundaries
- Working with continuous byte streams safely in Go

This project is a foundation for:
- Video steganography
- Frame-level analysis
- Custom video pipelines
- Systems / backend engineering practice

---

## High-level architecture
