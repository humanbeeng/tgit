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

<div>
    <a href="https://www.loom.com/share/79374698750b4ad88560e569ae2fe2b3">
      <p>tgit demonstration - Click to watch video</p>
    </a>
    <a href="https://www.loom.com/share/79374698750b4ad88560e569ae2fe2b3">
      <img style="max-width:300px;" src="https://cdn.loom.com/sessions/thumbnails/79374698750b4ad88560e569ae2fe2b3-with-play.gif">
    </a>
  </div>

#### Features
- Commits are made against the checked out branch.
- Uncreated branch cannot be checked out to.
- Cannot reinit a tgit repository.
- Cannot add duplicate files to staging if they are unmodified. Only those files that are either modified or haven't been staged before can be added and the old one will be overwritten in staged area.
- Invalid command checks.
- Checkout <branch-name> will yield all committed files that we made upto branch-name HEAD.
- Cannot commit an empty staged area.

#### References
[Git internals](https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain)
