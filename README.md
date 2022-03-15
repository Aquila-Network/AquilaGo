# AquilaGo
Golang client library to access Aquila Network Neural Search Engine

## Import:
```
$ go get github.com/Aquila-Network/go-aquila
```
## Usage
Import 
```
import (
    moduleDb "github.com/Aquila-Network/go-aquila"
	moduleDbSrc "github.com/Aquila-Network/go-aquila/src"
)
```

## Create aquiladb:
```
createAquilaDb := &moduleDbSrc.DataStructCreateDb{
    Schema: moduleDbSrc.SchemaStruct{
        Description: fmt.Sprintf("Database of %v %v", customer.FirstName, customer.LastName),
        Unique:      customer.SecretKey,
        Encoder:     "strn:msmarco-distilbert-base-tas-b",
        Codelen:     768,
        Metadata: moduleDbSrc.MetadataStructCreateDb{
            Name: "string",
            Age:  "number",
        },
    },
}

response, err := moduleDb.AquilaModule().AquilaDbInterface.CreateDatabase(createAquilaDb, url)
```

## Doc insert
Example struct:
```
&moduleDbSrc.DatatDocInsertStruct{
    Docs: []moduleDbSrc.DocsStruct{
        {
            Payload: moduleDbSrc.PayloadStruct{
                Metadata: moduleDbSrc.MetadataStructDocInsert{
                    Name: "name1",
                    Age:  20,
                },
                Code: []float64{0.1, 0.2, 0.3},
            },
        },
        {
            Payload: moduleDbSrc.PayloadStruct{
                Metadata: moduleDbSrc.MetadataStructDocInsert{
                    Name: "name1",
                    Age:  20,
                },
                Code: []float64{0.1, 0.2, 0.3},
            },
        },
    },
    DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
}
```

Insert:
```
responseInsert, err := moduleDb.AquilaModule().AquilaDbInterface.InsertDocument(docInsert, url)
```

## Doc sarch
Example struct:
```
matrix := make([][]float64, 1)
matrix[0] = make([]float64, 1)
matrix[0] = []float64{
    -0.01806008443236351, -0.17380790412425995, 0.03992759436368942, 0.43514639139175415,
}
searchBody := &moduleDbSrc.SearchAquilaDbRequestStruct{
    Data: moduleDbSrc.DataSearchStruct{
        Matrix:       matrix,
        K:            10,
        R:            0,
        DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
    },
}
```

Search:
```
response, err := moduleDb.AquilaModule().AquilaDbInterface.SearchKDocument(searchBody, url)
```

## Doc delete
Example struct:
```
docDelete := &moduleDbSrc.DeleteDataStruct{
    Ids: []string{
        "3gwTnetiYJfHTBcqGwoxETLsmmdGYVsd5MRBohuTG22C",
        "BXsbHy9B3tU9zaHwU41jATzDBisNEFa67XKvYZhB2fzQ",
    },
    DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
}
```

Delete:
```
responseDelete, err := moduleDb.AquilaModule().AquilaDbInterface.DeleteDocument(docDelete, url)
```

Run test:
```
$ go test ./src/
or
$ go test ./src/ -v
```
Clear test cache:
```
$ go clean -testcache
```