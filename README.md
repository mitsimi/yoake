# StariGo - A Simple Tool for start up applications

Little cli program for easy and universal application start on startup for windows and linux.

## Install

```sh
 go install github.com/mitsimi/starigo@latest
```

[Find latest Release Here](https://github.com/mitsimi/stariGo/releases/latest)

Or just download the binary from releases page and put it in your PATH.

---

## DISCLAIMER

The Linux version depends on the Desktop Environment autostarting applications inside the `~/.config/autostart` directory. 
If your Desktop Environment does not use this directory, you may not be able to use the Linux version without any own work. 
It is only needed that on start this application will be run.

---

## Platforms Supported

- [ ] Windows
- [ ] Linux

## Features

- [ ] Initialize starigo
- [ ] Add an application to startup
- [ ] Remove an application from startup
- [ ] List all applications in startup
- [ ] List a specific added application
- [ ] Enable or disable an application from startup
- [ ] Enable or disable starigo startup
- [ ] Show current set delay for startup applications
- [ ] Change delay for startup applications

## Usage

### Available Commands

```sh
    starigo help

    starigo add <app_name>
    starigo rm <app_name>

    starigo list
    starigo list <app_name>

    starigo enable <app_name>
    starigo disable <app_name>

    starigo start
    starigo stop

    starigo delay
    starigo delay <delay_time>
```

### Initialize a starigo

```sh
    starigo init
```
### Add an application to startup

```sh
$ starigo add <app_name> <app_path>
```

### Remove an application from startup

`remove` is also aliased as `rm`

```sh
$ starigo remove <app_name>
```

```sh
$ starigo rm <app_name>
```

### List all applications in startup

`list` is also aliased as `ls` & `l`


```sh
$ starigo list
```

```sh
$ starigo ls
```

```sh
$ starigo l
```

### List a specific added application

```sh
$ starigo list <app_name>
```

```sh
$ starigo ls <app_name>
```

```sh
$ starigo l <app_name>
```

### Enable an application from startup

`enable` is also aliased as `en`

```sh
$ starigo enable <app_name>
```

```sh
$ starigo en <app_name>
```

### Disable an application from startup

`disable` is also aliased as `dis`

```sh
$ starigo disable <app_name>
```

```sh
$ starigo dis <app_name>
```

### Run starigo at startup

```sh
$ starigo start
```

### Stop starigo from running at startup

```sh
$ starigo stop
```

### Check current delay for startup

`delay` is also aliased as `d`

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
- Config file `config.json` is located in the standard location for config files for your operating system.
- LOG file `startup.log` is in the same directory as the config file.

## LICENSE
Starigo is released under the GPL-3.0 license. See [LICENSE](https://github.com/mitsimi/starigo/blob/main/LICENSE)

