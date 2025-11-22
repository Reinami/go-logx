# go-logx

Just a simple logging utility that i use in some of my proejcts. You are more than welcome to use it
or if you want to change something would be happy to review a PR if you want something added.

## Installation

```bash
go get github.com/Reinami/go-logx@latest
```

## Usage

```go
func main() {
	logger := slog.New(&logx.ColorLogger{})
	slog.SetDefault(logger)

    slog.Error("some error occured")
}
```
