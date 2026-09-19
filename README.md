<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# paiza

A collection of independent Go programs, one per directory, each solving a single paiza.jp coding-practice problem by reading its judge-formatted standard input and printing the expected answer.

**English** · [简体中文](README.zh-CN.md)

[![CI](https://github.com/anyingiit/paiza/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/paiza/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/paiza)](LICENSE)

[Report a bug](https://github.com/anyingiit/paiza/issues/new?template=bug_report.yml) · [Request a feature](https://github.com/anyingiit/paiza/issues/new?template=feature_request.yml)

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About The Project

Each directory in this repository is named after a paiza.jp practice-problem ID
(`D007`, `D170`, and so on) and holds one self-contained Go program: a
`package main` whose `main()` reads that problem's input straight from stdin,
in the exact format paiza's judge sends it, and prints the single computed
answer. A few of the solutions (`D156`, `D189`, `D200`) read a two-line,
N-by-M grid of numbers through a shared `GetInputWithPaizaStanderedDatas`
helper; `lib/standerInputV3.31` and the numbered `playground/standerInputV3`
directories keep the successive drafts of that same input parser from before
it was copied into the solved problems.

This repository is archived: it is a record of problems paiza's own judge has
already accepted, not a tool under active development. Some directories (for
example `D007` and `D170`) also carry a `docment.md` or `document.md` file
reproducing the paiza.jp problem statement in full, including its own sample
input/output, alongside a note on the score the submitted solution received;
that text originates from paiza.jp, not from this repository's author.

## Getting Started

### Prerequisites

- Go 1.18 or newer, the floor declared in `go.mod`

### Installation

```sh
git clone https://github.com/anyingiit/paiza.git
cd paiza
go build ./...
```

`go build ./...` compiles every directory's solution as its own binary; it
does not produce a single combined program, because there isn't one.

## Usage

Each directory is a separate program. Run one with `go run` and give it the
problem's input on stdin, in the format its `document.md`/`docment.md` (where
present) describes. For example, `D007` prints `N` asterisks for an input `N`:

```sh
go run ./D007 <<< "4"
# ****
```

`D170` reads two lines -- a lap distance and a lap count -- and prints their
product:

```sh
printf '40\n15\n' | go run ./D170
# 600
```

## Contributing

Contributions are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md) for how to open an issue or a pull request, and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for the standards expected of everyone taking part.

Please do not report security issues in public issues or pull requests. [SECURITY.md](SECURITY.md) explains how to report them privately.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.

## Contact

Project link: [https://github.com/anyingiit/paiza](https://github.com/anyingiit/paiza)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
