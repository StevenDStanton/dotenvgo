# godotenv

## Testing

[![Coverage Status](https://coveralls.io/repos/github/StevenDStanton/dotenvgo/badge.svg?branch=master)](https://coveralls.io/github/StevenDStanton/dotenvgo?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/StevenDStanton/dotenvgo)](https://goreportcard.com/report/github.com/StevenDStanton/dotenvgo)
![MIT LICENSE](https://camo.githubusercontent.com/2bb6ac78e5a9f4f688a6a066cc71b62012101802fcdb478e6e4c6b6ec75dc694/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6c6963656e73652d4d49542d626c75652e737667)

### Normal Tests

`go test`

### Fuzzing Tests

`go test -fuzz=Fuzz -fuzztime=30s`

## Usage

Default `.env` current directory returning a map and setting the enviromental variables

```go
variables, err := dotenvgo.Load(dotenvgo.Both)
```

`.env` in `test/` sub directory and setting only the enviromental variables

```go
_, err := dotenvgo.Load(dotenvgo.Enviroment, "test/")
```

## Why

There are already several godotenv packages out there, why write my own?

I hate external dependancies. Especially for something as trivial as this. It is nearly as bad
as left-pad over in the npm ecosystem. Part of why I love go is that the standard library is
powerful enough I can write nearly anything I want on my own without relying on external
resources.

## How

Well I have used .env files for a long time so I think I have the gist of them. Now I may miss
some nuances. I have asked ChatGPT for a run down of requirements which I will list below.

If you feel I need some more features or missed something feel free to reach out to me. I do
not accept outside code in most of my repos but I am happy to address any issues.

## Who

This package is written for me to use. If other people want to use it I have no problem with that.
However, I would encourage you to go write your own, its really easy!

## Requirements

- [x] Accept optional parameter for path to .env
- [x] Search current directory or provided directory for .env
- [x] Read File into memory
- [x] Ignore anything after the # character
- [x] Ignore blank lines
- [x] Parse key-value pairs deliminated by `=`
- [x] Trim any whitespace around the `=`
- [x] Values can be quoted or unquoted
- [x] Unquoted lines end at line break
- [ ] Special characters must be surrounded by quotes
- [ ] Multiline values must use double quotes
- [x] Ignore any line without a `=`
- [x] Return Value as a Map
- [x] Save to os.Getenv
- [ ] Set up overload bool to prevent overwriting existing variables
- [ ] Set up autoload
