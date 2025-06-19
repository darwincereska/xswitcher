# XSwitcher

**Auto-switch keyboard layouts based on the active application.**

**XSwitcher** is a lightweight CLI tool that launches programs using different keyboard layouts depending on what you're working on. Great for bilingual users, coders, or anyone who needs to switch layouts on the fly.

## ✨ Features

- 🔁 Automatically switches between **main** and **secondary** layouts
- ⚙️ Simple config file for layout preferences
- 🧠 Stays out of your way—just run and focus

## 📦 Installation

1. Clone the repository:

``` bash
git clone https://darwincereska/xswitcher.git
```

2. Enter the folder:

``` bash
cd xswitcher
```

#### Arch Linux:

``` bash
make install-arch-deps

makepkg -si
```

#### Other Distros:

``` bash
make install 
```

## Uninstall

- To uninstall run:

``` bash
make uninstall
```

## Usage

1. Init the config file `(~/.config/xswitcher.yml)`:

``` bash
xswitcher --init
```

2. Change the layouts that fit the **setxkbmap** usage:

``` yaml
layouts:
  main: "us" # Layout you use
  secondary: "us colemak" # Layout the app switches to
```

3. Run the program:

``` bash
xswitcher <command> [args...]
```

## Compatibility

- Works with Steams `%command%` launch option
- Works with most apps

## The Future

I plan on adding support for macos and windows using their given api.