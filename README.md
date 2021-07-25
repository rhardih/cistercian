# cistercian

Commandline tool to generate [Cistercian numerals].

> The medieval Cistercian numerals, or "ciphers" in nineteenth-century parlance,
> were developed by the Cistercian monastic order in the early thirteenth
> century at about the time that Arabic numerals were introduced to northwestern
> Europe. They are more compact than Arabic or Roman numerals, with a single
> glyph able to indicate any integer from 1 to 9,999. Digits are based on a
> horizontal or vertical stave, with the position of the digit on the stave
> indicating its place value (units, tens, hundreds or thousands). These digits
> are compounded on a single stave to indicate more complex numbers.

## Installation

```bash
go get github.com/rhardih/cistercian
```

## Example usage

### Text

```bash
$ cistercian 7323
    ╷
    │╲
    │ ╲
    │  ╲
────┤
    │
    │
    │
╷   │
│   │  ╱
│   │ ╱
│   │╱
└───┘
```

### SVG

```bash
$ cistercian -svg 5221 > 5221.svg
```

Result:

![5221](https://raw.githubusercontent.com/rhardih/cistercian/main/5221.svg)

[Cistercian numerals]: https://en.wikipedia.org/wiki/Cistercian_numerals
