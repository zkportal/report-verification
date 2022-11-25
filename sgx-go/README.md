# SGX report verification example in Go

## Prerequisites

1. Install Go 1.18
2. Install EGo by following [EGo's README](https://github.com/edgelesssys/ego)

## Running

If you installed EGo from Snap:

```bash
EGOPATH=/snap/ego-dev/current/opt/ego CGO_CFLAGS=-I$EGOPATH/include CGO_LDFLAGS=-L$EGOPATH/lib go run main.go
```

If you installed EGo from a DEB package:
```bash
CGO_CFLAGS=-I/opt/ego/include CGO_LDFLAGS=-L/opt/ego/lib go run main.go
```
