[English](README.md) · **简体中文**

> 英文版是规范版本。本页与 [README.md](README.md) 不一致时，以英文版为准。

<!-- translation-of: README.md sha256:931bb10bb17ddb1c -->

<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# paiza

一组彼此独立的 Go 程序，每个目录对应一个 paiza.jp 编程练习题，按题目判题系统规定的格式读取标准输入，并输出该题期望的答案。

[![CI](https://github.com/anyingiit/paiza/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/paiza/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/paiza)](LICENSE)

[报告问题](https://github.com/anyingiit/paiza/issues/new?template=bug_report.yml) · [提出需求](https://github.com/anyingiit/paiza/issues/new?template=feature_request.yml)

<details>
  <summary>目录</summary>
  <ol>
    <li><a href="#about-the-project">关于本项目</a></li>
    <li><a href="#getting-started">开始使用</a></li>
    <li><a href="#usage">用法</a></li>
    <li><a href="#contributing">参与贡献</a></li>
    <li><a href="#license">许可证</a></li>
    <li><a href="#contact">联系方式</a></li>
  </ol>
</details>

## 关于本项目

仓库中的每个目录都以一个 paiza.jp 练习题的编号命名（例如 `D007`、`D170`），
并且各自包含一个独立完整的 Go 程序：一个 `package main`，其 `main()` 会按
paiza 判题系统发送的确切格式从标准输入读取该题的输入，并输出唯一计算出的
答案。其中一部分解答（`D156`、`D189`、`D200`）通过共用的
`GetInputWithPaizaStanderedDatas` 辅助函数读取两行、N 行 M 列的数字网格；
`lib/standerInputV3.31` 以及编号的 `playground/standerInputV3` 系列目录，
保存了这个输入解析器在被复制进已解出的题目之前的历次演进版本。

本仓库已归档：它记录的是 paiza 判题系统已经判为正确的题目，而不是一个正在
积极开发的工具。部分目录（例如 `D007` 和 `D170`）还带有 `docment.md` 或
`document.md` 文件，其中完整复制了 paiza.jp 的题目描述，包括其自带的输入
输出示例，以及提交解答获得的分数；这些文字来自 paiza.jp 本身，并非本仓库
作者所写。

## 开始使用

### 环境要求

- Go 1.18 或更新版本，这是 `go.mod` 声明的最低要求

### 安装

```sh
git clone https://github.com/anyingiit/paiza.git
cd paiza
go build ./...
```

`go build ./...` 会把每个目录下的解答各自编译成一个独立的可执行文件；
它不会产出单一的合并程序，因为本来就没有这样一个程序。

## 用法

每个目录都是一个独立的程序。用 `go run` 运行其中一个，并把题目要求的输入
通过标准输入传给它，格式以该目录下的 `document.md`/`docment.md`（如果有）
为准。例如，`D007` 会针对输入 `N` 输出 N 个星号：

```sh
go run ./D007 <<< "4"
# ****
```

`D170` 读取两行输入——一圈的距离和圈数——并输出两者的乘积：

```sh
printf '40\n15\n' | go run ./D170
# 600
```

## 参与贡献

欢迎参与。[CONTRIBUTING.md](CONTRIBUTING.md) 说明如何提交 issue 或 pull request，[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 说明对所有参与者的行为要求。

请不要在公开的 issue 或 pull request 中报告安全问题。[SECURITY.md](SECURITY.md) 说明了私下报告的方式。

## 许可证

以 MIT 许可证分发。详见 [LICENSE](LICENSE)。

## 联系方式

项目地址：[https://github.com/anyingiit/paiza](https://github.com/anyingiit/paiza)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
