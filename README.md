# Golang Bitbucket API

**NOTE**: not ready for production

## Install

```
go get github.com/compotlab/bitterbucket
```

## Godoc

- <http://godoc.org/github.com/compotlab/bitterbucket>
 
## Example
 
    1. Repository resource

```
c := bitterbucket.Client{Auth: bitterbucket.Auth{Token: "TOKEN"}}

req := bitterbucket.RepositoryRequest{
    Owner: "owner",
    Slug:  "repo_slug":,
}

c.Repository().GetRepository(req)
```

    2. Commits resource

```
c := bitterbucket.Client{Auth: bitterbucket.Auth{Token: "TOKEN"}}

req := bitterbucket.CommitsRequest{
    Owner:       rd.Owner,
    Slug:        rd.Slug,
    BranchOrTag: "master",
}

c.Commits().GetCommits(req)
```

License
===
Licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/compotlab/bitterbucket/blob/master/LICENSE) for the full license text.