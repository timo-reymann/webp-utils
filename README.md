webp-utils
===
[![GitHub Release](https://img.shields.io/github/v/release/timo-reymann/webp-utils.svg?label=version)](https://github.com/timo-reymann/webp-utils/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/timo-reymann/webp-utils)](https://goreportcard.com/report/github.com/timo-reymann/webp-utils)
[![CircleCI Build Status](https://circleci.com/gh/timo-reymann/webp-utils.svg?style=shield)](https://app.circleci.com/pipelines/github/timo-reymann/webp-utils)

Wrapper around [webp cli tools](https://developers.google.com/speed/webp/docs/using) to allow easier batch processing. 


# Currently supported webp cli tools:
- cwebp


# Technical notes
For easier development a json schema is used as base for validation and paramter passing.


# Install

## Binary

### IMPORTANT: webcp tools must be installed!
To use the tools with a binary you will also need to install the webp cli tools yourself.

Just make sure they are placed inside your PATH or add them accordingly.

More information about installation can be found on [developers.google.com](https://developers.google.com/speed/webp/docs/using). 


### Get binary using go cli
If you already have go setup you can simply run a `go get`:
```bash
go get -u github.com/timo-reymann/webp-utils
```
### Get binary from GitHub release
Navigate to the [latest release](https://github.com/timo-reymann/webp-utils/releases/latest) and download the binary for your platform.

## Docker
If you want to use docker:

```bash
docker run --user $(id -u):$(id -g) -v $PWD:/workspace --rm -it timoreymann/webp-utils cwebp --config /etc/webp-utils/default.json --file-glob *.png
```


# Usage
tbd
