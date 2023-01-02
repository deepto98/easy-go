## Learning
1. `go mod vendor` - save local copies of dependencies in vendor folder, instead of fetching everytime
2. Go now has native support for multi module workspaces (https://github.com/golang/tools/blob/master/gopls/doc/workspace.md). <br>
 ```
    go work init
    go work use ./app1 ./app2 (Add all the separate modules here)
```
