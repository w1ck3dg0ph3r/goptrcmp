# goptrcmp

A golang linter that reports comparison between pointer values.

## Usage

### Standalone

```sh
go run github.com/w1ck3dg0ph3r/goptrcmp/cmd/goptrcmp@latest ./...
```

### Using `go vet`

Install the goptrcmp binary:
```sh
go install github.com/w1ck3dg0ph3r/goptrcmp/cmd/goptrcmp@latest
```

Run go vet:
```sh
go vet -vettool=goptrcmp ./...
```
