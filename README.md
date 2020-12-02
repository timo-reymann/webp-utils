webp-utils
===
[![GitHub Release](https://img.shields.io/github/v/release/timo-reymann/webp-utils.svg?label=version)](https://github.com/timo-reymann/webp-utils/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/timo-reymann/webp-utils)](https://goreportcard.com/report/github.com/timo-reymann/webp-utils)
[![CircleCI Build Status](https://circleci.com/gh/timo-reymann/webp-utils.svg?style=shield)](https://app.circleci.com/pipelines/github/timo-reymann/webp-utils)

> WIP: Wrapper around webp cli tools to allow easier batch processing. 
>
> For easier development a json schema is used as base for validation and paramter passing

# Usage

## Binary

### Get binary using go cli
If you already have go setup you can simply run a `go get`:
```bash
go get -u github.com/timo-reymann/webp-utils
```

### Get binary from GitHub release
tbd

### Docker
If you want to use docker:

```bash
docker run --user $(id -u):$(id -g) -v $PWD:/workspace --rm -it timoreymann/webp # args go here
```