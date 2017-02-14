## Yo [![License](https://gh.kaos.io/ekol.svg)](https://essentialkaos.com/ekol)

`yo` is command-line YAML processor.

* [Installation](#installation)
* [Usage](#usage)
* [Build Status](#build-status)
* [Contributing](#contributing)
* [License](#license)

#### Installation

To build the Yo from scratch, make sure you have a working Go 1.5+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get github.com/essentialkaos/yo
```

If you want update Yo to latest stable release, do:

```
go get -u github.com/essentialkaos/yo
```

#### Usage

```
Usage: yo {options} query

Options

  --from-file, -f filename    Read data from file
  --no-color, -nc             Disable colors in output
  --help, -h                  Show this help message
  --version, -v               Show version

Examples

  yo '.foo'
  Return value for key foo

  yo '.foo | length'
  Print value length

  yo '.foo[]'
  Return all items from array

  yo '.bar[2:]'
  Return subarray started from item with index 2

  yo '.bar[1,2,5]'
  Return items with index 1, 2 and 5 from array

  yo '.bar[] | length'
  Print array size

  yo '.xyz | keys'
  Print hash map keys

  yo '.xyz | keys | length'
  Print number of hash map keys

```

#### Build Status

| Repository | Status |
|------------|--------|
| Stable | [![Build Status](https://travis-ci.org/essentialkaos/yo.svg?branch=master)](https://travis-ci.org/essentialkaos/yo) |
| Unstable | [![Build Status](https://travis-ci.org/essentialkaos/yo.svg?branch=develop)](https://travis-ci.org/essentialkaos/yo) |

#### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

#### License

[EKOL](https://essentialkaos.com/ekol)