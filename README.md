# godotenv

## Testing

[![Coverage Status](https://coveralls.io/repos/github/StevenDStanton/dotenvgo/badge.svg?branch=master)](https://coveralls.io/github/StevenDStanton/dotenvgo?branch=master)

### Normal Tests

`go test`

### Fuzzing Tests

`go test -fuzz=Fuzz -fuzztime=30s`

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
