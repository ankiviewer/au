[![Build Status](https://travis-ci.org/shouston3/au.svg?branch=master)](https://travis-ci.org/shouston3/au)
[![codecov](https://codecov.io/gh/shouston3/au/branch/master/graph/badge.svg)](https://codecov.io/gh/shouston3/au)

# AV

av is a companion program for the `av_umbrella` project

See project here: [ankiviewer/av_umbrella](https://github.com/ankiviewer/av_umbrella)

# Quickstart

Requires Go version >=1.8

If you have your go workspace set up, run:

```bash
go get github.com/shouston3/au
```

Otherwise run:

```bash
git clone https://github.com/shouston3/au
cd au
go build
```

Place this `au` executable somewhere in your `$PATH`

For useage run: `au` or `au help` after installation

# Tests

Files with: `//+build !test` at the top will not be covered

Run the coverage locally and view the output with:

```bash
go test -coverprofile=c.out -tags=test && go tool cover -html=c.out
```

Just run the coverage output in the terminal with:

```bash
go test -cover -tags=test
```

add the `-v` flag for a verbose output
