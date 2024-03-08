# cerberusd-go

![Build status](https://github.com/Cerberus-Wallet/cerberusd-go/actions/workflows/check-go-validation.yml/badge.svg) ![Installer build status](https://github.com/Cerberus-Wallet/cerberusd-go/actions/workflows/build-unsigned-installers.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/Cerberus-Wallet/cerberusd-go)](https://goreportcard.com/report/Cerberus-Wallet/cerberusd-go)

Cerberus Communication Daemon aka Cerberus Bridge.

**Only compatible with Chrome (version 53 or later) and Firefox (version 55 or later).**

We officially don't support Windows 7 and older; it could run, but we don't guarantee it.

## What does cerberusd do and why it is needed?

Cerberusd is a tiny http server, that allows webpages (like Cerberus Suite in web mode) to communicate with Cerberus directly.

Our new devices now support WebUSB, which should eliminate the need for Cerberus Bridge; however, there are some reasons, why bridge is still needed.

1. Firefox does not allow WebUSB ([see discussion here](https://github.com/mozilla/standards-positions/issues/100)).
2. Devices with old firmware (2018 and older) support only HID and not WebUSB.
3. WebUSB does not allow synchronization of USB access between domains.

## Install and run from source

cerberusd-go requires go >= 1.18.

```
git clone --recursive https://github.com/Cerberus-Wallet/cerberusd-go.git
cd cerberusd-go
go build .
./cerberusd-go -h
```

On Linux don't forget to install the [udev rules](https://github.com/Cerberus-Wallet/cerberus-common/blob/master/udev/51-cerberus.rules) if you are running from source and not using pre-built packages.

#### Debug mode

When built with `-tags debug` a debug mode is enabled. This disables CORS which is helpful for local development and when run inside a docker image.

## Build release packages

Prerequisites:

* install `docker`
* make sure `docker` is in `$PATH`
* `make build-release`; the installers are in `release/installers`, binaries in `release/binaries`

The base docker images are all built for both ARM and Intel 64, so they should work on both x64 architectures and ARM.

The base images are quite big and can take a while to download (mainly the musl cross-compiler, about 1 GB) and build (mainly the Rust-based apple-codesign). However, it should be cached correctly and run fast next time.

## Signing release packages

By default, the binaries and installers are unsigned and unnotarized. The build does not require any certificates or private keys, but produces unsigned binaries and packages.

The notarization and signing is all done in Docker, so it can run everywhere. (No need to run the mac notarization on macOS, etc.)

If you want to sign the packages, you need the following:

* For Linux, you need to put GPG private key into `release/linux/privkey.asc`.
* For Windows, you need to put GPG private key into `release/windows/privkey.asc` and an authenticode to `release/windows/authenticode.key` and `release/windows/authenticode.crt`.
* For macOS:
  1. You need to put GPG private key into `release/macos/privkey.asc`.
  2. Then you need to generate and put a lot of things for notarization and signing into `release/macos/certs`; see the details in top comment of `release/macos/release.sh`.

All those files are ignored by `.gitignore` so they are not accidentally put into git.

## Emulator support

Cerberusd supports emulators for all Cerberus versions. However, you need to enable it manually; it is disabled by default. After enabling, services that work with emulator can work with all services that support cerberusd.

To enable emulator, run cerberusd with a parameter `-e` followed by port, for every emulator with an enabled port:

`./cerberusd-go -e 21324`

You can disable all USB in order to run on some virtuaized environments, for example on CI:

`./cerberusd-go -e 21324 -u=false`

## API documentation

`cerberusd-go` starts a HTTP server on `http://localhost:21325`. AJAX calls are only enabled from cerberus.uraanai.com subdomains.

Server supports following API calls:

| url <br> method | parameters | result type | description |
|-------------|------------|-------------|-------------|
| `/` <br> POST | | {`version`:&nbsp;string} | Returns current version of bridge |
| `/enumerate` <br> POST | | Array&lt;{`path`:&nbsp;string, <br>`session`:&nbsp;string&nbsp;&#124;&nbsp;null}&gt; | Lists devices.<br>`path` uniquely defines device between more connected devices. Two different devices (or device connected and disconnected) will return different paths.<br>If `session` is null, nobody else is using the device; if it's string, it identifies who is using it. |
| `/listen` <br> POST | request body: previous, as JSON | like `enumerate` | Listen to changes and returns either on change or after 30 second timeout. Compares change from `previous` that is sent as a parameter. "Change" is both connecting/disconnecting and session change. |
| `/acquire/PATH/PREVIOUS` <br> POST | `PATH`: path of device<br>`PREVIOUS`: previous session (or string "null") | {`session`:&nbsp;string} | Acquires the device at `PATH`. By "acquiring" the device, you are claiming the device for yourself.<br>Before acquiring, checks that the current session is `PREVIOUS`.<br>If two applications call `acquire` on a newly connected device at the same time, only one of them succeed. |
| `/release/SESSION`<br>POST | `SESSION`: session to release | {} | Releases the device with the given session.<br>By "releasing" the device, you claim that you don't want to use the device anymore. |
| `/call/SESSION`<br>POST | `SESSION`: session to call<br><br>request body: hexadecimal string | hexadecimal string | Both input and output are hexadecimal, encoded in following way:<br>first 2 bytes (4 characters in the hexadecimal) is the message type<br>next 4 bytes (8 in hex) is length of the data<br>the rest is the actual encoded protobuf data.<br>Protobuf messages are defined in [this protobuf file](https://github.com/Cerberus-Wallet/cerberus-common/blob/master/protob/messages.proto) and the app, calling cerberusd, should encode/decode it itself. |
| `/post/SESSION`<br>POST | `SESSION`: session to call<br><br>request body: hexadecimal string | 0 | Similar to `call`, just doesn't read response back. Also forces the message to be sent even if another call is in progress. Usable mainly for debug link and workflow cancelling on Cerberus.  |
| `/read/SESSION`<br>POST | `SESSION`: session to call | 0 | Similar to `call`, just doesn't post, only reads. Usable mainly for debug link. |

## Debug link support

Cerberusd has support for debug link.

To support an emulator with debug link, run

`./cerberusd-go -ed 21324:21325 -u=false`

this will detect emulator debug link on port 21325, with regular device on 21324.

To support WebUSB devices with debug link, no option is needed, just run cerberusd-go.

In the `enumerate` and `listen` results, there are now two new fields: `debug` and `debugSession`. `debug` signals that device can receive debug link messages.

Session management is separate for debug link and normal interface, so you can have two applications - one controlling cerberus and one "normal".

There are new calls:

* `/debug/acquire/PATH`, which has the same path as normal `acquire`, and returns a `SESSION`
* `/debug/release/SESSION` releases session
* `/debug/call/SESSION`, `/debug/post/SESSION`, `/debug/read/SESSION` work as with normal interface

The session IDs for debug link start with the string "debug".

## Copyright

* (C) 2018 Karel Bilek, Jan Pochyla
* CORS Copyright (c) 2013 The Gorilla Handlers Authors, [BSD license](https://github.com/gorilla/handlers/blob/master/LICENSE)
* (c) 2017 Jason T. Harris (also see https://github.com/deadsy/libusb for comprehensive list)
* (C) 2017 Péter Szilágyi (also see https://github.com/karalabe/hid for comprehensive list)
* (C) 2010-2016 Pete Batard <pete@akeo.ie> (also see https://github.com/pbatard/libwdi/ for comprehensive list)
* Licensed under LGPLv3
