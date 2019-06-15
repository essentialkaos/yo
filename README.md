<p align="center"><a href="#readme"><img src="https://gh.kaos.st/yo.svg"/></a></p>

<p align="center">
  <a href="https://travis-ci.org/essentialkaos/yo"><img src="https://travis-ci.org/essentialkaos/yo.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/yo"><img src="https://goreportcard.com/badge/github.com/essentialkaos/yo"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-yo-master"><img alt="codebeat badge" src="https://codebeat.co/badges/f9f024b1-a3b2-418f-b3a4-b4f1d0d4c73d" /></a>
  <a href="https://essentialkaos.com/ekol"><img src="https://gh.kaos.st/ekol.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

Yo is a command-line YAML processor.

### Installation

#### From source

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

To build the Yo from scratch, make sure you have a working Go 1.10+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get github.com/essentialkaos/yo
```

If you want update Yo to latest stable release, do:

```
go get -u github.com/essentialkaos/yo
```

#### From ESSENTIAL KAOS Public repo for RHEL6/CentOS6

```bash
[sudo] yum install -y https://yum.kaos.st/kaos-repo-latest.el6.noarch.rpm
[sudo] yum install yo
```

#### From ESSENTIAL KAOS Public repo for RHEL7/CentOS7

```bash
[sudo] yum install -y https://yum.kaos.st/kaos-repo-latest.el7.noarch.rpm
[sudo] yum install yo
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and OS X from [EK Apps Repository](https://apps.kaos.st/yo/latest).

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

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
