# go-pssh-generator

[![Build Status](https://travis-ci.com/JohnnyCPC/go-pssh-generator.svg?branch=main)](https://travis-ci.com/JohnnyCPC/go-pssh-generator)
[![Go Report Card](https://goreportcard.com/badge/github.com/JohnnyCPC/go-pssh-generator)](https://goreportcard.com/report/github.com/JohnnyCPC/go-pssh-generator)
[![codecov](https://codecov.io/gh/JohnnyCPC/go-pssh-generator/branch/main/graph/badge.svg?token=0NT3MZGAEP)](https://codecov.io/gh/JohnnyCPC/go-pssh-generator)

go-pssh-generator is a Protection System Specific Header ('pssh') boxes generator written in Go (Golang).

## Usage
### PSSH data/box generation
An example to generate a Widevine PSSH Data/Box with content-id 'test_id' and provider 'widevinve_test'
```bash
$ go run main.go -contentid test_id -provider widevinve_test
```
Will get the result: (both hex and base64)
```bash
PSSH Data: 08011a0e7769646576696e76655f746573742207746573745f6964
PSSH Data(base64): CAEaDndpZGV2aW52ZV90ZXN0Igd0ZXN0X2lk
PSSH BOX: 0000003B7073736800000000EDEF8BA979D64ACEA3C827DCD51D21ED0000001B08011a0e7769646576696e76655f746573742207746573745f6964 
PSSH BOX(base64): AAAAO3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABsIARoOd2lkZXZpbnZlX3Rlc3QiB3Rlc3RfaWQ=
```

## Useful Links
- ["cenc" Initialization Data Format by W3C](https://www.w3.org/TR/eme-initdata-cenc/)
  (This page contains the defination of PSSH box format.)
