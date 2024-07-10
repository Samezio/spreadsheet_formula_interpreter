# spreadsheet_formula_interpreter
Welcome to the **Spreadsheet Formula Interpreter** project! This project involves creating a Lexer, Parser, and Interpreter for solving spreadsheet-like formulae. As of now, the project supports basic arithmetic operations including addition, subtraction, multiplication, division, and handling of brackets.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Running](#running)
- [License](#license)

## Introduction

Spreadsheet Formula Interpreter is a tool designed to parse and evaluate arithmetic expressions similar to those used in spreadsheet applications. The interpreter takes a formula as input and returns the computed result.

## Features

- **Lexer**: Tokenizes the input string into meaningful components.
- **Parser**: Builds a syntax tree from the tokens.
- **Interpreter**: Evaluates the syntax tree to compute the result.
- **Supports Basic Arithmetic**:
  - Addition (`+`)
  - Subtraction (`-`)
  - Multiplication (`*`)
  - Division (`/`)
  - Brackets for operation precedence (`()`)
  - String concatenation  (`&`)
  - Equality (`==` and `!=`)
  - Comparision (`>`, `<`, `>=`, `<=`)
  - Cell reference (`column``row` eg `A1`, `XA123`): You'll need to provide `RetriveCelldata`([Refer](interpreter/interpreter.go)) to the interpreter to fetch the value. This test version gives 1 for all cells

## Running

To use the Spreadsheet Formula Interpreter, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/Samezio/spreadsheet_formula_interpreter.git
    ```

2. Change to the project directory:
    ```bash
    cd spreadsheet_formula_interpreter
    ```

3. Build the project:
    ```
    go run .
    ```
4. Write your expression

## Exmaple

## License
[License](LICENSE)