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

1. Install dependencies:
``` bash
make install-arch-deps
```

2. Install package:
``` bash
makepkg -si
```

#### Other Distros:

``` bash
make install 
```

#### MacOS:

 Make sure you have [Homebrew](https://brew.sh) installed

1. Add my [brew tap](https://github.com/darwincereska/homebrew-software):

``` bash
brew tap darwincereska/software
```

2. Install app:

``` bash
brew install xswitcher
```

## Uninstall

#### Linux:

``` bash
make uninstall
```

#### MacOS:

``` bash
brew uninstall xswitcher
```

## Usage

1. Initialize the config file `(~/.config/xswitcher.yml)`:

``` bash
xswitcher --init
```

2. Change the layouts that fit the **setxkbmap** usage:

``` yaml
# xswitcher.yml
layouts:
  main: "us" # Layout you use
  secondary: "us colemak" # Layout the app switches to
```

3. Run the program:

``` bash
xswitcher <command> [args...]
```

## Compatibility

- Works with Steams `%command%` launch option:
``` bash
xswitcher %command%
```
- Supports Environment Variables
- Works with any app that constantly runs. It will not work with apps such as **VSCode** due to that it makes a new process every time

## The Future

I plan on adding support for macos and windows using their given api.
