# gosh

A URL shortener in Go, made as an experiment in learning Go. Currently operates in-memory only.

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/gosh
```

Add your long-running agent logic to `command/agent/command.go`, and any status or action commands you need to `commands.go`.

### Testing

``make test``

## License

BSD licenced.

## Contributing

See `CONTRIBUTING.md` for more details.
