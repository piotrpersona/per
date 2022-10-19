# per

Simple error wrapper library

## Install

```console
go get github.com/piotrpersona/per
```

## Usage

Wrap error:
```go
err := ReadFile()
err = per.Error(err, "cannot read file")

fmt.Println(err)
// cannot read file, err: OH NO
```

With format:
```go
err = per.Errorf(ReadFile(), "cannot read file: %s", "file.txt")

fmt.Println(err)
// cannot read file: file.txt, err: OH NO
```

Overwrite default format:
```go
per.Format = "Custom MSG: %s, ERR: %s"
err = per.Errorf(ReadFile(), "cannot read file: %s", "file.txt")

fmt.Println(err)
// Custom MSG: cannot read file: file.txt, ERR: OH NO
```

## Examples

[Live example at: https://goplay.space/#HciC635WRLQ](https://goplay.space/#HciC635WRLQ)
[examples/main.go](./examples/main.go)