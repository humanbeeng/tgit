# tgit
Tiny Git - Lightweight Version Control System


![Go build](https://github.com/humanbeeng/tgit/actions/workflows/go.yml/badge.svg)
[![Code style: gofumpt](https://img.shields.io/badge/code%20style-gofumpt-29BEB0.svg)](https://github.com/mvdan/gofumpt)


#### Installation
```
go install github.com/humanbeeng/tgit@latest
```

#### Commands
- `init`: Initialises an empty tgit repository.
- `branch`: Creates a new branch by having current branch as head.
- `add`: Stage file(s)
- `commit`: Commit staged files to branch.
- `checkout`: Checkouts to branch with file changes upto checked out branch head.
- `help`: Displays a help message


#### References
[Git internals](https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain)
