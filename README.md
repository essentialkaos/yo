<p align="center"><a href="#readme"><img src=".github/images/card.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/w/yo/ci"><img src="https://kaos.sh/w/yo/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/yo/codeql"><img src="https://kaos.sh/w/yo/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="https://kaos.sh/r/yo"><img src="https://kaos.sh/r/yo.svg" alt="GoReportCard" /></a>
  <a href="#license"><img src=".github/images/license.svg"/></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#ci-status">CI Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

Yo is a command-line YAML processor.

### Installation

#### From source

To build the Yo from scratch, make sure you have a working Go [1.24+](https://github.com/essentialkaos/.github/blob/master/GO-VERSION-SUPPORT.md) workspace (_[instructions](https://go.dev/doc/install)_), then:

```
go install github.com/essentialkaos/yo@latest
```

#### From [ESSENTIAL KAOS Public Repository](https://kaos.sh/kaos-repo)

```bash
sudo dnf install -y https://pkgs.kaos.st/kaos-repo-latest.el$(grep 'CPE_NAME' /etc/os-release | tr -d '"' | cut -d':' -f5).noarch.rpm
sudo dnf install yo
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and macOS from [EK Apps Repository](https://apps.kaos.st/yo/latest):

```bash
bash <(curl -fsSL https://apps.kaos.st/get) yo
```

### Upgrading

Since version `1.1.0` you can update `yo` to the latest release using [self-update feature](https://github.com/essentialkaos/.github/blob/master/APPS-UPDATE.md):

```bash
yo --update
```

This command will runs a self-update in interactive mode. If you want to run a quiet update (_no output_), use the following command:

```bash
yo --update=quiet
```

> [!NOTE]
> Please note that the self-update feature only works with binaries that are downloaded from the [EK Apps Repository](https://apps.kaos.st/yo/latest). Binaries from packages do not have a self-update feature and must be upgraded via the package manager.

### Usage

<img src=".github/images/usage.svg" />

### CI Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://kaos.sh/w/yo/ci.svg?branch=master)](https://kaos.sh/w/yo/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/yo/ci.svg?branch=develop)](https://kaos.sh/w/yo/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/.github/blob/master/CONTRIBUTING.md).

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://kaos.dev"><img src="https://raw.githubusercontent.com/essentialkaos/.github/refs/heads/master/images/ekgh.svg"/></a></p>
