# Yoake - A Simple Tool for start up applications

Little CLI program for easy and universal application start on login for windows and linux.

Yoake (**夜明け**) is a noun meaning the dawn in Japanese. This project get started in the early days of the windows operating system.

## Install

```sh
go install github.com/mitsimi/yoake@latest
```

Or download the [latest release from GitHub Releases](https://github.com/mitsimi/yoake/releases/latest) and put it into `%USERPROFILE%/go/bin` for Windows or `$HOME/go/bin` for Linux. Make sure the directory is added to your PATH. Then you can use `yoake` command to use the application.

---

## DISCLAIMER

> The Linux version depends on the Desktop Environment autostarting applications inside the `~/.config/autostart` directory. 
> If your Desktop Environment does not use this directory, you may not be able to use the Linux version without any own work. 
>
> If you are using a Window Manager, you must call the program in the init script.

---

## Platforms Supported

- [x] Windows
- [x] Linux

## Features

- [x] Initialize yoake
- [x] Add an application to startup
- [x] Remove an application from startup
- [x] List all applications in startup
- [x] List information of a specific application
- [x] Enable or disable an application from startup
- [x] Enable or disable yoake startup
- [x] Show current set delay for startup applications
- [x] Change delay for startup applications

## Usage

### Available Commands

```sh
    yoake help

    yoake setup

    yoake add <app_name> <app_path>
    yoake remove <app_name>

    yoake show [app_name]

    yoake enable [app_name]
    yoake disable [app_name]

    yoake delay [delay_time]
```

<br/>

### Setup the autostart on login

`init` is an alias for `setup`

```sh
$ yoake setup
```

```sh
$ yoake init
```

### Add an application to startup

```sh
$ yoake add <app_name> <app_path>
```

### Remove an application from startup

`rm` is an alias for `remove`

```sh
$ yoake remove <app_name>
```

```sh
$ yoake rm <app_name>
```

### List all applications in startup

`show` is aliased as `list`, `ls`

```sh
$ yoake show
```
is equivalent to
```sh
$ yoake show all
```

```sh
$ yoake list
```

```sh
$ yoake ls
```



### List information about a specific application

```sh
$ yoake show <app_name>
```

```sh
$ yoake list <app_name>
```

```sh
$ yoake ls <app_name>
```

### Enable yoake startup

`en` is an alias for `enable`

```sh
$ yoake enable
```

```sh
$ yoake en
```

### Disable yoake startup

`dis` is an alias for `disable`

```sh
$ yoake disable
```

```sh
$ yoake dis
```

### Enable an application from startup

`en` is an alias for `enable`

```sh
$ yoake enable [app_name]
```

```sh
$ yoake en [app_name]
```

### Disable an application from startup

`dis` is an alias for `disable`

```sh
$ yoake disable [app_name]
```

```sh
$ yoake dis [app_name]
```


### Check current delay for startup

`d` is an alias for `delay`

```sh
$ yoake delay
```

### Set new delay for startup

delay is in seconds.

```sh
$ yoake delay <delay_time>
```

<br/>

## Defaults
---
- Delay is set to 10 seconds
- Config file `config.json` is located in the standard location for config files of your operating system.
- Log file `startup.log` is in the same directory as the config file.

Default location for config file and log file:
- Windows: %AppData%
- Linux: ~/.config

## Example Config File
---
```json
{
    "config": {
        "Enabled": true,
        "Delay": 10
    },
    "apps": {
        "app1": {
            "Enabled": true,
            "Path": "C:\\app1\\app1.exe"
        },
        "app2": {	
            "Enabled": false,
            "Path": "/usr/share/bin/app2"
        }
    }
}
```

## LICENSE
---
Yoake is released under the GPL-3.0 license. See [LICENSE](https://github.com/mitsimi/yoake/blob/main/LICENSE)

