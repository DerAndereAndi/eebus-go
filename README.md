# eebus-go

![Build Status](https://github.com/DerAndereAndi/eebus-go/actions/workflows/default.yml/badge.svg?branch=dev)

The goal is to provide a basic EEBUS implementation

## Introduction

This repository contains:

- adoptions of the SPINE and SHIP EEBUS model definitions, there are likely issues and some models are not 100% correct
- (De-)serialization for EEBUS specific JSON format requirements
- Certificate support incl. creating a compatible cert and key
- mDNS Support (announcement and connecting to an announced SKI)
- ... work in progress

## Usage

### HEMS

```sh
Usage: go run cmd/hems/main.go <serverport> <remoteski> <certfile> <keyfile>
```

### EVSE

```sh
Usage: go run cmd/evse/main.go <serverport> <remoteski> <certfile> <keyfile>
```

### Explanation

The remoteski is from the eebus service to connect to.
If no certfile or keyfile are provided, they are generated and printed in the console so they can be saved in a file and later used again. The local SKI is also printed.
