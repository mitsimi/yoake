# StariGo - A Simple Tool for start up applications

Little cli program for easy and universal application start on startup for windows and linux.

## Install

```sh
 go install github.com/mitsimi/starigo@latest
```

[Find latest Release Here](https://github.com/mitsimi/stariGo/releases/latest)

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

- [x] Initialize starigo
- [x] Add an application to startup
- [x] Remove an application from startup
- [ ] List all applications in startup
- [ ] List information of a specific application
- [x] Enable or disable an application from startup
- [x] Enable or disable starigo startup
- [x] Show current set delay for startup applications
- [x] Change delay for startup applications

## Usage

### Available Commands

```sh
    starigo help

    starigo add <app_name>
    starigo remove <app_name>

    starigo list [app_name]

    starigo enable [app_name]
    starigo disable [app_name]

    starigo delay [delay_time]
```

<br/>

### Add an application to startup

```sh
$ starigo add <app_name> <app_path>
```

### Remove an application from startup

`rm` is an alias for `remove`

```sh
$ starigo remove <app_name>
```

```sh
$ starigo rm <app_name>
```

### List all applications in startup

`ls` & `l` are aliases for `list`

```sh
$ starigo list
```

```sh
$ starigo ls
```

```sh
$ starigo l
```

### List information about a specific application

```sh
$ starigo list <app_name>
```

```sh
$ starigo ls <app_name>
```

```sh
$ starigo l <app_name>
```

### Enable starigo startup

`en` is an alias for `enable`

```sh
$ starigo enable
```

```sh
$ starigo en
```

### Disable starigo startup

`dis` is an alias for `disable`

```sh
$ starigo disable
```

```sh
$ starigo dis
```

### Enable an application from startup

`en` is an alias for `enable`

```sh
$ starigo enable [app_name]
```

```sh
$ starigo en [app_name]
```

### Disable an application from startup

`dis` is an alias for `disable`

```sh
$ starigo disable [app_name]
```

```sh
$ starigo dis [app_name]
```


### Check current delay for startup

`d` is an alias for `delay`

```sh
$ starigo delay
```

### Set new delay for startup

delay is in seconds.

```sh
$ starigo delay <delay_time>
```

## Defaults

- Delay is set to 10 seconds
- Config file `config.json` is located in the standard location for config files of your operating system.
- Log file `startup.log` is in the same directory as the config file.

Default location for config file and log file:
- Windows: %AppData%
- Linux: ~/.config

## LICENSE
Starigo is released under the GPL-3.0 license. See [LICENSE](https://github.com/mitsimi/starigo/blob/main/LICENSE)

