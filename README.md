# Calculator

A CLI application and library to add numbers.

## Problem Description

> As a fan of math,  
I want to model a calculator which accepts add, subtract, multiply, divide, cancel and exit as commands, so that I can do basic mathematics.

## Prerequisites

- Go (1.18.2)

## Test

    $ go test
_For detailed results with necessary information:_

    $ make test

## Build

    $ make build

## Library Usage

TODO

## CLI Application Usage

__Run__ the CLI Application:

    $ make run

_Note: The initial value is set as 0.00_
<br>

__Add__ an value to the balance:

    add amount

_Where __value__ is a float value_

Example:

    add 34.09


__Subtract__ an value from the balance:

    sub value

_Where __value__ is a float value_

Example:

    sub 34.09