<!-- markdownlint-configure-file {
  "MD013": {
    "code_blocks": false,
    "tables": false
  },
  "MD033": false,
  "MD041": false
} -->

<div align="center" markdown="1">

# Bridgeftp

[![Tests](https://github.com/abhijeetmohanan/bridgeftp/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/abhijeetmohanan/bridgeftp/actions/workflows/test.yml/badge.svg?query=branch%3Amain)
![GitHub Release](https://img.shields.io/github/v/release/abhijeetmohanan/bridgeftp)

  <br/>

</div>

## About

A tool that can copy files from Remote to Remote FTP Servers

* It streams data from source to destination [ no need to save files locally just to transfer ]

* Supported Protocols
  * ftp
  * sftp

## Installation

**1: Binary Installation**

download one of [releases](https://github.com/abhijeetmohanan/bridgeftp/releases)

## Getting Started

<b> FTP </b>

```bash
bridgeftp --src ftp://username:password@server:port/filepath \
  --dest ftp://username@password@server:port/filepath \
  --bs 512
```

<b> SFTP </b>

```bash
bridgeftp --src sftp://username:password@server:port/filepath \
  --dest sftp://username@password@server:port/filepath \
  --bs 512
```

##
