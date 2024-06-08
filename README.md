<p align="center"><a href="#readme"><img src="https://gh.kaos.st/yo.svg"/></a></p>

<p align="center">
  <a href="https://github.com/essentialkaos/yo/actions"><img src="https://github.com/essentialkaos/yo/workflows/CI/badge.svg" alt="GitHub Actions Status" /></a>
  <a href="https://github.com/essentialkaos/yo/actions?query=workflow%3ACodeQL"><img src="https://github.com/essentialkaos/yo/workflows/CodeQL/badge.svg" /></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/yo"><img src="https://goreportcard.com/badge/github.com/essentialkaos/yo"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-yo-master"><img alt="codebeat badge" src="https://codebeat.co/badges/f9f024b1-a3b2-418f-b3a4-b4f1d0d4c73d" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

Yo is a command-line YAML processor.

### Installation

#### From source

To build the Yo from scratch, make sure you have a working Go 1.19+ workspace (_[instructions](https://go.dev/doc/install)_), then:

```
go install github.com/essentialkaos/yo@latest
```

#### From [ESSENTIAL KAOS Public Repository](https://pkgs.kaos.st)

```bash
sudo yum install -y https://pkgs.kaos.st/kaos-repo-latest.el$(grep 'CPE_NAME' /etc/os-release | tr -d '"' | cut -d':' -f5).noarch.rpm
sudo yum install yo
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and macOS from [EK Apps Repository](https://apps.kaos.st/yo/latest):

```bash
bash <(curl -fsSL https://apps.kaos.st/get) yo
```

### Usage

<img src=".github/images/usage.svg" />

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://github.com/essentialkaos/yo/workflows/CI/badge.svg?branch=master)](https://github.com/essentialkaos/yo/actions) |
| `develop` | [![CI](https://github.com/essentialkaos/yo/workflows/CI/badge.svg?branch=develop)](https://github.com/essentialkaos/yo/actions) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
