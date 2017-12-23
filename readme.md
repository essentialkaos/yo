# Yo [![Build Status](https://travis-ci.org/essentialkaos/yo.svg?branch=master)](https://travis-ci.org/essentialkaos/yo) [![Go Report Card](https://goreportcard.com/badge/github.com/essentialkaos/yo)](https://goreportcard.com/report/github.com/essentialkaos/yo) [![codebeat badge](https://codebeat.co/badges/f9f024b1-a3b2-418f-b3a4-b4f1d0d4c73d)](https://codebeat.co/projects/github-com-essentialkaos-yo-master) [![License](https://gh.kaos.io/ekol.svg)](https://essentialkaos.com/ekol)

`yo` is command-line YAML processor.

* [Installation](#installation)
* [Usage](#usage)
* [Build Status](#build-status)
* [Contributing](#contributing)
* [License](#license)

### Installation

#### From source

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)):

```
git config --global http.https://pkg.re.followRedirects true
```

To build the Yo from scratch, make sure you have a working Go 1.5+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get github.com/essentialkaos/yo
```

If you want update Yo to latest stable release, do:

```
go get -u github.com/essentialkaos/yo
```

#### From ESSENTIAL KAOS Public repo for RHEL6/CentOS6

```bash
[sudo] yum install -y https://yum.kaos.io/6/release/x86_64/kaos-repo-8.0-0.el6.noarch.rpm
[sudo] yum install yo
```

#### From ESSENTIAL KAOS Public repo for RHEL7/CentOS7

```bash
[sudo] yum install -y https://yum.kaos.io/7/release/x86_64/kaos-repo-8.0-0.el7.noarch.rpm
[sudo] yum install yo
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and OS X from [EK Apps Repository](https://apps.kaos.io/yo/latest).

### Usage

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

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![Build Status](https://travis-ci.org/essentialkaos/yo.svg?branch=master)](https://travis-ci.org/essentialkaos/yo) |
| `develop` | [![Build Status](https://travis-ci.org/essentialkaos/yo.svg?branch=develop)](https://travis-ci.org/essentialkaos/yo) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[EKOL](https://essentialkaos.com/ekol)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.io/ekgh.svg"/></a></p>
