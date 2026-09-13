
# Quad

A team project from the 1st week of Go Piscine at Zone01 Athens.

## Description

The goal of this project is to create five Go functions that print rectangles with a given width (`x`) and height (`y`).

The functions are:

- `QuadA`
- `QuadB`
- `QuadC`
- `QuadD`
- `QuadE`

Each function prints a rectangle using a different combination of characters.

If `x` or `y` is less than or equal to zero, the function prints nothing.

## Project Structure

```text
quad/
├── README.md
├── quadA.go
├── quadB.go
├── quadC.go
├── quadD.go
├── quadE.go
├── go.mod
├── go.sum
├── QuadA/
│   ├
│   ├── test1/
│   │   └── main.go
│   ├── test2/
│   │   └── main.go
│   ├── test3/
│   │   └── main.go
│   └── test4/
│       └── main.go
├── QuadB/
│   ├
│   ├── test1/
│   │   └── main.go
│   └── ...
├── QuadC/
│   ├
│   ├── test1/
│   │   └── main.go
│   └── ...
├── QuadD/
│   ├
│   ├── test1/
│   │   └── main.go
│   └── ...
└── QuadE/
    ├
    ├── ...
    │   └── ...
    └── ...
```


The file names above are examples of the expected organization.

## Functions

### QuadA

Prints a rectangle using:

- `o` for the corners
- `-` for the top and bottom borders
- `|` for the left and right borders

Example:

```text
o---o
|   |
o---o
```

### QuadB

Prints a rectangle using:

- `/` for the top-left corner
- `\` for the top-right and bottom-left corners
- `*` for the borders and inside edges

Example:

```text
/***\
*   *
\***/
```

### QuadC

Prints a rectangle using:

- `A` for the top-left and top-right corners
- `B` for the borders and inside edges
- `C` for the bottom-left and bottom-right corners

Example:

```text
ABBBA
B   B
CBBBC
```

### QuadD

Prints a rectangle using:

- `A` for the top-left and bottom-left corners
- `C` for the top-right and bottom-right corners
- `B` for the borders and inside edges

Example:

```text
ABBBC
B   B
ABBBC
```

### QuadE

Prints a rectangle using:

- `A` for the top-left corner
- `C` for the bottom-right corner
- `B` for the borders and inside edges

Example:

```text
ABBBC
B   B
CBBBA
```

## How to Open

Clone the repository:

```bash
git clone https://platform.zone01.gr/git/magora/quad
```

Move into the project directory:

```bash
cd quad
```

## How to Test

To test a function, create a Go program that imports the `piscine` package and calls the function.

Example:

```go
package main

import "piscine"

func main() {
    piscine.QuadA(5, 3)
}
```

Expected output:

```text
o---o
|   |
o---o
```

The same approach can be used to test `QuadB`, `QuadC`, `QuadD`, and `QuadE`.

## Team

- Marios Agora
- Nikolaos Isankoulidis
- Eirini Fourou

## Technologies

- Go
- Gitea

## Learning Goals

Through this project, we practiced:

- Writing functions in Go.
- Using loops.
- Working with integers and conditions.
- Understanding how to work on a shared Git repository.
- Collaborating as a team.

## Authors

This project was created as part of the Go Piscine at Zone01 Athens.

Team repository:

https://platform.zone01.gr/git/magora/quad