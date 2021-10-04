# Currency

## Dependencies
- Docker
- Docker Compose
- Go v1.16

## Build

```sh
make
```

## Execute
### Set environment variables
```sh
set -a
. ./env.env
set +a
```

### Run
On distinct terminals run:
```sh
./bin/server
```


```sh
./bin/gateway
```

## Clean (remove the binary and the container)
```sh
make clean
```
## Remove database data
```sh
make destroy
```
