# visual-git
This is a simple tool to visualize local Git contributions.
Original idea from [Visualize your local Git contributions with Go ](https://flaviocopes.com/go-git-contributions/)
Special Thanks to [Flavio Copes](https://flaviocopes.com/) for the great article.

This repo aims to implement the same idea with `go-git/v5`, increase the readibility and add some features.

## Usage
```
$ make
$ ./main -add <path-to-search-git-repo>
```

## Environment
- [Go 1.20](https://golang.org/dl/)

## Dependencies
- [github.com/go-git/go-git/v5](https://pkg.go.dev/github.com/go-git/go-git/v5)

