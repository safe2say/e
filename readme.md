[![](https://godoc.org/github.com/nii236/e?status.svg)](http://godoc.org/github.com/nii236/e)

# E

*This package is under review*

Structured errors package, inspired mostly by [Ben Johnson](https://middlemost.com/failure-is-your-domain/) and [Upspin](https://github.com/upspin/upspin) with some modifications (extra meta data, JSON marshalling, type coercions, HTTP codes).

```bash
$ cd example
$ go run *.go
serving on :8080
visit:
http://localhost:8080?username=nii236
http://localhost:8080?username=nii237
```

## Usage

Return your structured error:

```go
return &e.Error{
        Op:      op,
        Code:    e.ECONFLICT,
        Message: "Username is already in use. Please choose a different username.",
        Meta:    map[string]string{"username": username, "filename": "storage.go"},
    }
```

Handle your structured _root_ error:

```go
if err != nil {
    fmt.Println(err)
    http.Error(w, e.ErrorMessage(err), e.HTTPCode(err))
}
```

Server result:

```
main.handlerExample: main.createUser: <conflict> [filename=storage.go username=nii236] Username is already in use. Please choose a different username.
```

API response:

```
Username is already in use. Please choose a different username.
```
