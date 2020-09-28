# Contains

Contains - A command line tool to check if words exists in a text.

# Synopsis

```
contains [ OPTIONS ] < WORDS >
```

# Description

Contains is a command line tool to check if given words exists in a text and return a exit code depending on the result. It is  created  in order to simplify bash scripts.

# Installation

**On linux**

Clone the repository and cd into it.

```
git clone git@github.com:eric-von-knorring/go-contains.git
cd go-contains
```

In order to install the binary and manual pages run.

```
sudo make install
```

# Options

```
  --any  Exit with status code 0 if any of the words matched.

  -f, --file
        Reads from file instead of standard input.

  -h, --help
        Prints help text for contains.
```


# Exit Status

The following are all the status codes that the program will terminate with.
```
  0.     The given words where found.

  1.     The given words where not found.

  127.   The file could not be opened.
```

# Usage

This program can be used for scripts in order to check for specific values in output form other programs. The resulting status code from contains can then be used in order to run different scripts that are depending on certain conditions.

```
xrandr  |  grep ' connected' | cut -d' ' -f1 | contains DP1 HDM2 && $HOME/.screenlayout/setup.sh
```

