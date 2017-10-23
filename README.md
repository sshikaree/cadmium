# Cadmium

**Cadmium** is a multiprotocol messenger written in Go and Qt5.

Supported protocols (current and planned in near future):

* IRC (planned)
* XMPP (planned)
* Matrix (planned)

# The idea

This project was started because there was no good native multiprotocol
messenger available for macOS. Adium is good, but looks like it stopped
being developed, as their Trac at the moment of writing this README was
broken and latest source on Github was about 1.5.10.2 version.

Due to nature of Go and Qt5, this messenger is crossplatform and should
be able to run on Windows, Linux and macOS.

It may be possible to run this messenger also on *BSD, if you'll get
Qt 5.9 running on it. Last time version 5.7.1 was available.

# Dependencies

Cadmium requires:

* Go 1.9+ (lower versions might work and might not. 1.9 is a recommended
minumim).
* Qt 5.9.

# Installation

## Binary releases

Just download binary from Releases section and run it.

## Source code

Refer to [installation documentation](/doc/install_from_source.md).

# Developers chat

We're using Matrix as our primary support source. You can access
chat room as guest with [this link](https://riot.im/app/#/room/cadmium:matrix.feder8.ru).