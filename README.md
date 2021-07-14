# cistercian

Commandline tool to generate [Cistercian numerals].

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
