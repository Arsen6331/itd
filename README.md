# ITD
## InfiniTime Daemon

`itd` is a daemon that uses my infinitime [library](https://go.arsenm.dev/infinitime) to interact with the [PineTime](https://www.pine64.org/pinetime/) running [InfiniTime](https://infinitime.io).

[![Build status](https://ci.appveyor.com/api/projects/status/xgj5sobw76ndqaod?svg=true)](https://ci.appveyor.com/project/moussaelianarsen/itd)
[![Binary downloads](https://img.shields.io/badge/download-binary-orange)](https://minio.arsenm.dev/minio/itd/)
[![AUR package](https://img.shields.io/aur/version/itd-git?label=itd-git&logo=archlinux)](https://aur.archlinux.org/packages/itd-git/)

---

### Features

- Notification relay
- Notification transliteration
- Call Notifications (ModemManager)
- Music control
- Get info from watch (HRM, Battery level, Firmware version, Motion)
- Set current time
- Control socket
- Firmware upgrades
- Weather
- BLE Filesystem

---

### Socket

This daemon creates a UNIX socket at `/tmp/itd/socket`. It allows you to directly control the daemon and, by extension, the connected watch.

The socket uses my [lrpc](https://gitea.arsenm.dev/Arsen6331/lrpc) library for requests. This library accepts requests in msgpack, with the following format:

```json
{"Receiver": "ITD", "Method": "Notify", "Arg": {"title": "title1", "body": "body1"}, "ID": "some-id-here"}
```

It will return a msgpack response, the format of which can be found [here](https://gitea.arsenm.dev/Arsen6331/lrpc/src/branch/master/internal/types/types.go#L30). The response will have the same ID as was sent in the request in order to allow the client to keep track of which request the response belongs to.

---

### Transliteration

Since the PineTime does not have enough space to store all unicode glyphs, it only stores the ASCII space and Cyrillic. Therefore, this daemon can transliterate unsupported characters into supported ones. Since some languages have different transliterations, the transliterators to be used must be specified in the config. Here are the available transliterators:

- eASCII
- Scandinavian
- German
- Hebrew
- Greek
- Russian
- Ukranian
- Arabic
- Farsi
- Polish
- Lithuanian
- Estonian
- Icelandic
- Czech
- French
- Armenian
- Korean
- Chinese
- Romanian
- Emoji

Place the desired map names in an array as `notifs.translit.use`. They will be evaluated in order. You can also put custom transliterations in `notifs.translit.custom`. These take priority over any other maps. The `notifs.translit` config section should look like this:

```toml
[notifs.translit]
    use = ["eASCII", "Russian", "Emoji"]
    custom = [
        "test", "replaced"
    ]
```

---

### `itctl`

This daemon comes with a binary called `itctl` which uses the socket to control the daemon from the command line. As such, it can be scripted using bash.

This is the `itctl` usage screen:
```
Control the itd daemon for InfiniTime smartwatches

Usage:
  itctl [flags]
  itctl [command]

Available Commands:
  firmware    Manage InfiniTime firmware
  get         Get information from InfiniTime
  help        Help about any command
  notify      Send notification to InfiniTime
  set         Set information on InfiniTime

Flags:
  -h, --help                 help for itctl
  -s, --socket-path string   Path to itd socket

Use "itctl [command] --help" for more information about a command.
```

---

### `itgui`

In `cmd/itgui`, there is a gui frontend to the socket of `itd`. It uses the [fyne library](https://fyne.io/) for Go. It can be compiled by running:

```shell
go build ./cmd/itgui
```

#### Screenshots

![Info tab](cmd/itgui/screenshots/info.png)

![Motion tab](cmd/itgui/screenshots/motion.png)

![Notify tab](cmd/itgui/screenshots/notify.png)

![FS tab](cmd/itgui/screenshots/fs.png)

![FS mkdir](cmd/itgui/screenshots/mkdir.png)

![Time tab](cmd/itgui/screenshots/time.png)

![Firmware tab](cmd/itgui/screenshots/firmware.png)

![Upgrade in progress](cmd/itgui/screenshots/progress.png)

![Metrics tab](cmd/itgui/screenshots/metrics.png)

---

### Installation

To install, install the go compiler and make. Usually, go is provided by a package either named `go` or `golang`, and make is usually provided by `make`. The go compiler must be version 1.17 or newer for various new `reflect` features.

To install, run
```shell
make && sudo make install
```

---

### Starting

To start the daemon, run the following **without root**:

```shell
systemctl --user start itd
```

To autostart on login, run:
```shell
systemctl --user enable itd
```

---

### Cross compiling

To cross compile, simply set the go environment variables. For example, for PinePhone, use:

```shell
make GOOS=linux GOARCH=arm64
```

This will compile `itd` and `itctl` for Linux aarch64 which is what runs on the PinePhone. This daemon only runs on Linux due to the library's dependencies (`dbus`, and `bluez` specifically).

---

### Configuration

This daemon places a config file at `/etc/itd.toml`. This is the global config. `itd` will also look for a config at `~/.config/itd.toml`.

Most of the time, the daemon does not need to be restarted for config changes to take effect.

---

### Attribution

Location data from OpenStreetMap Nominatim, &copy; [OpenStreetMap](https://www.openstreetmap.org/copyright) contributors

Weather data from the [Norwegian Meteorological Institute](https://www.met.no/en)