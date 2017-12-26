# hexip

![GitHub release](http://img.shields.io/github/release/bells17/hexip.svg?style=flat-square)
[![CircleCI](https://img.shields.io/circleci/project/github/bells17/hexip.svg)](https://circleci.com/gh/bells17/hexip)
![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)

## Installation

Get binary from [here](https://github.com/bells17/hexip/releases/latest) or build yourself.

```
go get github.com/bells17/hexip
cd $GOPATH/src/github.com/bells17/hexip
make && make install
```

## Usage

```
hexip 0a000001
10.0.0.1
```

## Help

```
hexip --help
Usage:
  hexip [OPTIONS]

Application Options:
      --version  print version

Help Options:
  -h, --help     Show this help message
```

## Development

### Requirements

- [ghr](https://github.com/tcnksm/ghr)
  - And need set your Github Token(export GITHUB_TOKEN="...")
- [zopfli](https://github.com/google/zopfli)


### Use Docker

```
docker-compose run builder bash
```

### Initialize

```
make init
make bundle
```

### Run Testing

```
make test
```

### Build

```
make build
```

### Release

```
make tag
make release
```
