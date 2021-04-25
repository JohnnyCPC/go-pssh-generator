# go-pssh-generator

[![Build Status](https://travis-ci.com/JohnnyCPC/go-pssh-generator.svg?branch=main)](https://travis-ci.com/JohnnyCPC/go-pssh-generator)
[![Go Report Card](https://goreportcard.com/badge/github.com/JohnnyCPC/go-pssh-generator)](https://goreportcard.com/report/github.com/JohnnyCPC/go-pssh-generator)
[![codecov](https://codecov.io/gh/JohnnyCPC/go-pssh-generator/branch/main/graph/badge.svg?token=0NT3MZGAEP)](https://codecov.io/gh/JohnnyCPC/go-pssh-generator)
[![Go Reference](https://pkg.go.dev/badge/github.com/JohnnyCPC/go-pssh-generator.svg)](https://pkg.go.dev/github.com/JohnnyCPC/go-pssh-generator)

go-pssh-generator is a Protection System Specific Header ('pssh') boxes generator written in Go (Golang).

## Usage

Download and install it:

```bash
$ go get github.com/JohnnyCPC/go-pssh-generator
```

### PSSH data/box generation
An example to generate a Widevine PSSH Data/Box with content-id 'test_id' and provider 'widevinve_test'
```bash
$ go run main.go -contentid test_id -provider widevinve_test
```
Will get the result: (both hex and base64)
```bash
PSSH Data: 08011a0d7769646576696e655f746573742207746573745f6964
PSSH Data(base64): CAEaDXdpZGV2aW5lX3Rlc3QiB3Rlc3RfaWQ=
PSSH BOX: 0000003A7073736800000000EDEF8BA979D64ACEA3C827DCD51D21ED0000001A08011a0d7769646576696e655f746573742207746573745f6964 
PSSH BOX(base64): AAAAOnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABoIARoNd2lkZXZpbmVfdGVzdCIHdGVzdF9pZA==+LqXnWSs6jyCfc1R0h7QAAABsIARoOd2lkZXZpbnZlX3Rlc3QiB3Rlc3RfaWQ=
```

## Useful Links
- ["cenc" Initialization Data Format by W3C](https://www.w3.org/TR/eme-initdata-cenc/)
  (This page contains the defination of PSSH box format.)
