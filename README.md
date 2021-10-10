webp-utils
===
[![GitHub Release](https://img.shields.io/github/v/release/timo-reymann/webp-utils.svg?label=version)](https://github.com/timo-reymann/webp-utils/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/timo-reymann/webp-utils)](https://goreportcard.com/report/github.com/timo-reymann/webp-utils)
[![CircleCI Build Status](https://circleci.com/gh/timo-reymann/webp-utils.svg?style=shield)](https://app.circleci.com/pipelines/github/timo-reymann/webp-utils)
[![codecov](https://codecov.io/gh/timo-reymann/webp-utils/branch/main/graph/badge.svg?token=OV4RC1ZQ7D)](https://codecov.io/gh/timo-reymann/webp-utils)
[![Dependabot](https://badgen.net/badge/Dependabot/enabled/green?icon=dependabot)](https://dependabot.com/)

Wrapper around [webp cli tools](https://developers.google.com/speed/webp/docs/using) to allow easier batch processing. 

# Currently supported webp cli tools:
- cwebp
- gif2webp

# Technical notes
For easier development a json schema is used as base for validation and paramter passing.

# Install

## Binary

### IMPORTANT: webcp tools must be installed!
To use the tools with a binary you will also need to install the webp cli tools yourself.

Just make sure they are placed inside your PATH or add them accordingly.

More information about installation can be found on [developers.google.com](https://developers.google.com/speed/webp/docs/precompiled). 

### Get binary using go cli
If you already have go setup you can simply run a `go get`:
```bash
go get -u github.com/timo-reymann/webp-utils
```
### Get binary from GitHub release
Navigate to the [latest release](https://github.com/timo-reymann/webp-utils/releases/latest) and download the binary for your platform.

## Docker
If you want to use docker:

e. g. for cwebp
```bash
docker run --user $(id -u):$(id -g) -v $PWD:/workspace --rm -it timoreymann/webp-utils \
    cwebp --config /etc/webp-utils/default.json --file-glob '*.png'
```

# Usage
1. **OPTIONAL:** Setup schema for your IDE, for available schemas look into [pkg/schema](./pkg/schema)
2. Create a file named `config.json`, containing all parameters for the command you want to execute, 
    you can also create a file named differently and use the `--config`-Flag 
3. Run the command and specify a file glob for the files to be processed, like this:
    `webp-utils cwebp --file-glob '*.png'`

## Configuration file
Each json key value pair represents an argument and its value, that can be templated using go templates.
You can also use the template functions from [sprig](https://github.com/Masterminds/sprig) as documented 
[here](http://masterminds.github.io/sprig/).

You can use any other key or the following builtins:

| Name              | Content                               |
| :---------------- | :------------------------------------ |
| source_file       | Source file path with extension       |
| source_file_name  | Source file name without extension    |

You can find all schemas inside [pkg/schema](./pkg/schema), you can also set up your ide/editor to give you
autocomplete based on the json schema.

## Examples
Feel free to browse around at [examples](./examples)
