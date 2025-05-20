# 🖼️ Multi-Format Steganography & Encrypted Concealment Tool

> A powerful steganography utility written in Go (`crypha`) that encrypts messages and files (Argon2id + ChaCha20-Poly1305) and invisibly embeds them inside Images, Audio files, QR codes, Unicode Text, and PDFs.

[![Author](https://img.shields.io/badge/Made%20by-cyber--atharv-00ffcc?style=flat-square&logo=github)](https://github.com/cyber-atharv)
[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev)
[![Encryption](https://img.shields.io/badge/AEAD-ChaCha20--Poly1305-8B5CF6?style=flat-square)](https://datatracker.ietf.org/doc/html/rfc8439)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)

---

## 📌 What is Steganography?

While **Cryptography** scrambles a message so nobody can read it, **Steganography** hides the message so nobody even knows it exists! 

If you send an encrypted `.bin` file, network firewalls and monitors immediately flag it as suspicious. But if you hide that encrypted data inside the invisible least-significant pixel bits of a normal cat photo, an audio track, or invisible zero-width Unicode characters, the file looks and sounds 100% normal.

This tool, built by **cyber-atharv**, is a full-featured steganography toolkit supporting **5 distinct carrier mediums** with industrial-grade authenticated encryption.

---

## ✨ Supported Carrier Formats

| Medium | How Data is Concealed | Resulting File |
|---|---|---|
| **🖼️ Images (PNG, BMP)** | Embeds bits inside the Least Significant Bits (LSB) of Red, Green, and Blue color channels. | Looks visually identical to human eyes. |
| **🎵 Audio (WAV, FLAC)** | Modifies the lowest bits of 16-bit PCM audio samples. | Sounds completely identical when played. |
| **📱 QR Codes** | Exploits Reed-Solomon error correction slack space in QR codes. | Standard phone cameras read the normal URL; this tool extracts the secret payload! |
| **📝 Unicode Text** | Uses zero-width invisible Unicode characters (`U+200B`, `U+2060`). | Text looks 100% unchanged on screen and passes clipboard copy/paste. |
| **📄 PDF Documents** | Hides data inside embedded object streams or metadata streams. | Opens normally in Adobe Reader. |

---

## 🚀 Quick Start & Usage

### 1. Build the Binary
```bash
cd steganography-multi-tool
go build -o crypha ./cmd/crypha
```

### 2. Examples

#### 🔹 Hide an encrypted secret inside an image
```bash
./crypha hide -i cover.png -m "Top Secret Password: 123" -k "MySuperSecretKey" -o secret.png
```

#### 🔹 Reveal and decrypt the secret from the image
```bash
./crypha reveal -i secret.png -k "MySuperSecretKey"
# Output: Top Secret Password: 123
```

#### 🔹 Hide a secret file inside an invisible text string
```bash
./crypha hide -f text -i normal_notes.txt -m "Hidden Flag: CTF{s73g0_m4573r}" -o shared_notes.txt
```

#### 🔹 Interactive Terminal Wizard
Run `./crypha` without arguments to open a guided step-by-step interactive terminal wizard!

---

## 🧠 Why I Built This

Exploring steganography bridged the gap between pure mathematics, digital signal processing (image color spaces and audio PCM waveforms), and modern applied cryptography. It taught me how data can be concealed in plain sight and how forensic analysts use steganalysis (entropy checks and bit-plane analysis) to detect hidden payloads.

---

## 📜 Author & License

- **Author:** [cyber-atharv](https://github.com/cyber-atharv)
- **License:** Open source under the MIT / AGPL License.
