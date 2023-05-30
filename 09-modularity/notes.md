# Modules & Packages

## Modules
- Any code that is versioned and distributed together
- A folder with a go.mod file
    - module name
    - go version
    - info about dependency modules

### Creating a module
```
    go mod init <module_name>
```
module_name = repo path

### Run a module
```
    go run .
```

### Build a module
```
    go build .
    go build -o <binary_name> .
```

### Package
- logical grouping of application code
- typically organized as folders


# Using 3rd party modules

## Search portal for available modules
```
    https://pkg.go.dev
```

## Installing 3rd party modules
```
    go get <module_name>
    go get -u <module_name>

    [ code downloaded to GOPATH/pkg ]
```

## To update the go.mod file 
```
    go mod tidy
```

## To download the dependency modules mentioned in the go.mod file
```
    go mod download
```

## To localize the dependencies 
```
    go mod vendor
```

## To consider the module cache (GOPATH/pkg)
```
    go run -mod=mod .
```

## Other useful go mod commands
```
    go mod graph
    go mod why <module_name>
```

## Reference:
```
    https://go.dev/ref/mod
```

## Curated list of 3rd party modules
```
    https://github.com/avelino/awesome-go
```