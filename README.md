# Go collection

This package support you with paginate, pluck and convert your array

## Install

```bash
go get github.com/Limpid-LLC/go-helpers
```

## Update version

```bash
go get -u github.com/Limpid-LLC/go-helpers@v1.0.22
```

## Usage

```go
import "github.com/Limpid-LLC/go-helpers"
```

### Array

```go
func main () {

    Array = helpers.Array{
        Items:  [], // your array
        Skip:   1, // searching page
        Limit:  10, // limit items for select
        Pluck:  ["name", "age"], // fields for showing
    }

    Array.Get() // Return items from 0 to 10 with name and age fields
    
}
```