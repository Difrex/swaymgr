# swaymgr

The [Swaywm](https://github.com/swaywm/sway) autotiling manager.

This project adds a autotiling feature to the SwayWM based on the [gosway](https://github.com/Difrex/gosway) IPC library.

## Install

### Build

You need a Go installed in your system.

```sh
git clone https://github.com/Difrex/swaymgr.git
cd swaymgr/swaymgr
go get -v
go build -o ~/.local/bin/swaymgr .
```

### From AUR

**swaymgr** package is available in the Arch Linux AUR. Install it with the favorite tool.

## Configure

* Autostart swaymgr

  Add this to the config:

  ```
  exec --no-startup-id swaymgr
  ```

* Set keybindings for changing layouts setup

```
bindsym --to-code $mod+Alt+s exec swaymgr -s 'spiral'
bindsym --to-code $mod+Alt+l exec swaymgr -s 'left'
bindsym --to-code $mod+Alt+m exec swaymgr -s 'manual'
```

## Commands

Commands can be sended to the control socket by the `-s` option.

* *get layout* -- returns information about current focused workspace in the JSON format
  ```
  swaymgr -s 'get layout' | jq
  {
      "name": "2:",
      "layout": "spiral",
      "managed": true
  }
  ```

* *set spiral* -- mark workspace as managed and set it to the spiral windows placement
  ```
  swaymgr -s 'set spiral'
  ```

* *set left* -- mark workspace as managed and set it to the left windows placement
  ```
  swaymgr -s 'set left'
  ```

* *set manual* -- mark workspace as unmanaged
  ```
  swaymgr -s 'set manual'
  ```

## Known issues

* Only spiral layout is working fine.

* Left layout is buggy.
