# Teonet fortune microservice

This is simple [Teonet](https://github.com/teonet-go/teonet) micriservice application which return linux fortune messages. The [fortune](https://linux.die.net/man/6/fortune) application should be installed first.

[![GoDoc](https://godoc.org/github.com/teonet-go/teofortune?status.svg)](https://godoc.org/github.com/teonet-go/teofortune/)
[![Go Report Card](https://goreportcard.com/badge/github.com/teonet-go/teofortune)](https://goreportcard.com/report/github.com/teonet-go/teofortune)

## Run the Teonet fortune microservice

There are various ways to start this Teonet microservice application:

### 1. From sources

```bash
git clone https://github.com/teonet-go/teofortune
cd teofortune
go run .
```

### 2. Install binary from github

```bash
go install github.com/teonet-go/teofortune .
teofortune
```

### 3. Using docker

```bash
docker run -d -it --network=host --restart=always --name teofortune -v \
$HOME/.config/teonet/teofortune:/root/.config/teonet/teofortune ghcr.io/teonet-go/\
teofortune:latest teofortune -loglevel=debug
```

## How to use

### Use `teofortune` in Teonet cli application

First of all you can use and check `teofortune` in Teonet cli application which is placed in Teonet package.

Install Teonet cli application:

```bash
go install github.com/teonet-go/teonet ./cmd/teonet
```

Run the Teonet cli application:

```bash
teonet
```

In the Teonet cli application print commands:

_Change address to your application address which you start before. Or you can use this address, but than you will connect to `teofortune` microservice application running in Teonet cloud. The address prints after you start teonet application in string:_  
`Teonet address: 8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF`

```
connectto 8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF
api 8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF
api 8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF forta
```

## Licence

[BSD](LICENSE)
