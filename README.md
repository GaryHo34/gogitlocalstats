# visual-git
This is a simple tool to visualize local Git contributions.
Original idea from [Visualize your local Git contributions with Go ](https://flaviocopes.com/go-git-contributions/)
Special Thanks to [Flavio Copes](https://flaviocopes.com/) for the great article.

This repo aims to implement the same idea with `go-git/v5`, increase the readibility and add some features.

## Usage
```
$ make
$ ./visualgit -add <path-to-search-git-repo> -e <email> 
```

output:
![Screenshot 2023-05-01 at 11 16 40 AM](https://user-images.githubusercontent.com/24312717/235505088-9fb6718a-1751-405a-b24d-843ede6e8d8f.png)


## Environment
- [Go 1.20](https://golang.org/dl/)

## Dependencies
- [github.com/go-git/go-git/v5](https://pkg.go.dev/github.com/go-git/go-git/v5)

