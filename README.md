# Go collection

This package support you with paginate, pluck and convert your array

## Install

```bash
go get github.com/Limpid-LLC/go-helpers
```

## Usage

```go
import "github.com/Limpid-LLC/go-helpers"
```

### Collection

```go
func main () {

    Collection = helpers.Collection{
        Items:  [], // your array
        Skip:   1, // searching page
        Limit:  10, // limit items for select
        Pluck:  ["name", "age"], // fields for showing
    }
    
    Collection.Get() // Return items from 0 to 10 with name and age fields
    
}
```