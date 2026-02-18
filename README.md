<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go">
  <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge">
  <img src="https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-blue?style=for-the-badge">
</p>

<h1 align="center">👻 PHANTOMTLS</h1>

<p align="center">
  <b>High-Performance TLS Testing & Local Forwarding Framework</b><br>
  Made with ♥ by Yashvir Gaming
</p>

---

## 🚀 Overview

PHANTOMTLS is a production-grade CLI framework written in Go for:

- 🔐 TLS fingerprint testing  
- 🔁 Proxy rotation  
- ⚡ Multithreaded traffic simulation  
- 🌐 Local request forwarding  
- 📊 Real-time colored logging  

Designed for controlled testing of infrastructure you own or have permission to test.

---

## ✨ Features

- ✅ Local Listener Mode (localhost:2024)
- ✅ Colored status code logging
- ✅ Multithread request engine
- ✅ Proxy list rotation
- ✅ TLS fingerprint rotation
- ✅ Live metrics display
- ✅ Interactive CLI interface
- ✅ Single compiled binary (no runtime required)

---

# 📦 Installation

## 1️⃣ Install Go

Download Go from:

https://go.dev/dl/

Verify installation:

```bash
go version
```

You should see something like:

```bash
go version go1.22.x windows/amd64
```

---

## 2️⃣ Clone Repository

```bash
git clone https://github.com/YOURUSERNAME/PhantomTLS.git
cd PhantomTLS
```

---

## 3️⃣ Install Dependencies

```bash
go mod tidy
```

---

## 4️⃣ Build

### Windows

```bash
go build -o PhantomTLS.exe
```

### Linux / macOS

```bash
go build -o PhantomTLS
```

---

# 🖥 Usage

Run:

```bash
PhantomTLS.exe
```

Select:

```
[1] Start Local Listener (localhost:2024)
```

Send requests to:

```
http://localhost:2024
```

---

## 📌 Required Headers

```
x-url: https://example.com
x-proxy: http://user:pass@ip:port
x-identifier: chrome | firefox | safari_ios | safari | okhttp
x-session-id: any-random-guid
```

---

## 📌 Optional Headers

```
host: example.com
content-encoding: gzip
```

---

# 🔐 TLS Fingerprint Profiles

Supported profiles:

- chrome
- firefox
- safari
- safari_ios
- okhttp
- randomized

---

# 📊 Logging Example

```
[REGISTERED] GET    200 - from 127.0.0.1     - https://example.com
```

Status colors:

- 🟢 200–299 → Success  
- 🟡 400–499 → Client Error  
- 🔴 500–599 → Server Error  

---

# 🤖 GitHub Actions

This repository includes automatic builds for:

- Windows
- Linux
- macOS

Triggered on push to `main`.

---

# 📁 Project Structure

```
PhantomTLS/
│
├── main.go
├── ui.go
├── server.go
├── worker.go
├── tlsclient.go
├── go.mod
├── go.sum
├── LICENSE
├── README.md
├── .gitignore
└── .github/
    └── workflows/
        └── build.yml
```

---

# ⚠ Disclaimer

PHANTOMTLS is intended for:

- Educational purposes
- Testing infrastructure you own
- TLS research and development

Do not use against systems you do not own or have explicit authorization to test.

---

# 📜 License

MIT License

Copyright (c) 2026 Yashvir Gaming

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction.
