# Go collection

This package support you with paginate, pluck and convert your array

## Install

```bash
go get github.com/wellsivi/go-collection
```

## Usage

```go
import "github.com/wellsivi/go-collection"

func main () {

    Collection = collection.Array{
        Items:  [], // your array
        Skip:   1, // searching page
        Limit:  10, // limit items for select
        Pluck:  ["name", "age"], // fields for showing
    }
    
    Collection.Get() // Return items from 0 to 10 with name and age fields
    
}
```