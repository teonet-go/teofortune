# Teonet fortune microservice

This is simple Teonet micriservice application which return linux fortune messages. The [fortune](https://linux.die.net/man/6/fortune) application should be installed first.

## Run the Teonet fortune microservice

### From sources

```bash
git clone https://github.com/teonet-go/teofortune
cd teofortune
go run .
```

### Install binary from github

```bash
git install https://github.com/teonet-go/teofortune .
teofortune
```

### Using docker
```bash
docker run ...
```

## How to use

### Use `teofortune` in Teonet cli application

Install Teonet cli application:
```bash
go install github.com/teonet-go/teonet ./cmd/teonet
```

Run the Teonet cli application:
```bash
teonet
```

In the Teonet cli application print commands:
```
connectto 8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF
api 8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF
api 8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF forta
```


# Licence

[BSD](LICENSE)
