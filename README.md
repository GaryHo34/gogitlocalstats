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
```
[scan]: found 8 new git folders
    /Users/yho/Documents/Code/Javascript/garyho34.github.io/.git
    /Users/yho/Documents/Code/Python/SeattleBot/.git
    /Users/yho/Documents/Code/Rust/simple_database_in_rust/.git
    /Users/yho/Documents/Code/design-patterns-notes/.git
    /Users/yho/Documents/Code/distributed-systems/.git
    /Users/yho/Documents/Code/http-c/.git
    /Users/yho/Documents/Code/incubator-devlake/.git
    /Users/yho/Documents/Code/visual-git-go/.git
[repo]: loading repo: garyho34.github.io
[repo]: loading repo: SeattleBot
[repo]: loading repo: simple_database_in_rust
[repo]: loading repo: design-patterns-notes
[repo]: loading repo: distributed-systems
[repo]: loading repo: http-c
[repo]: loading repo: incubator-devlake
[repo]: loading repo: visual-git-go
```

![Screenshot 2023-05-01 at 11 20 48 AM](https://user-images.githubusercontent.com/24312717/235505376-a3138ace-6c3e-48c1-b59b-b6d3b7c44cc2.png)



## Environment
- [Go 1.20](https://golang.org/dl/)

## Dependencies
- [github.com/go-git/go-git/v5](https://pkg.go.dev/github.com/go-git/go-git/v5)

