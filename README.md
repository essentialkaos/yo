<p align="center"><a href="#readme"><img src=".github/images/card.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/w/yo/ci"><img src="https://kaos.sh/w/yo/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/yo/codeql"><img src="https://kaos.sh/w/yo/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="https://kaos.sh/r/yo"><img src="https://kaos.sh/r/yo.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/b/yo"><img src="https://kaos.sh/b/f9f024b1-a3b2-418f-b3a4-b4f1d0d4c73d.svg" alt="codebeat badge" /></a>
  <a href="#license"><img src=".github/images/license.svg"/></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#ci-status">CI Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

Yo is a command-line YAML processor.

### Installation

#### From source

To build the Yo from scratch, make sure you have a working Go 1.21+ workspace (_[instructions](https://go.dev/doc/install)_), then:

```
go install github.com/essentialkaos/yo@latest
```

#### From [ESSENTIAL KAOS Public Repository](https://kaos.sh/kaos-repo)

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

### CI Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://kaos.sh/w/yo/ci.svg?branch=master)](https://kaos.sh/w/yo/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/yo/ci.svg?branch=develop)](https://kaos.sh/w/yo/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
