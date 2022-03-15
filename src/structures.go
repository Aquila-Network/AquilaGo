package src

// =====================================
// Doc insert struct
// =====================================

type MetadataStructDocInsert struct {
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
}

type PayloadStruct struct {
	Metadata MetadataStructDocInsert `json:"metadata" bson:"metadata"`
	Code     []float64               `json:"code" bson:"code"`
}

type DocsStruct struct {
	Payload PayloadStruct `json:"payload" bson:"payload"`
}

type DatatDocInsertStruct struct {
	Docs         []DocsStruct `json:"docs" bson:"docs"`
	DatabaseName string       `json:"database_name" bson:"database_name"`
}

type DocInsertRequestStruct struct {
	Data      DatatDocInsertStruct `json:"data"`
	Signature string               `json:"signature"`
}

// -----------------------
// Response

type DocInsertResponseStruct struct {
	Ids     []string `json:"ids"`
	Success bool     `json:"success"`
}

// =====================================
//  Aquila Hub
// =====================================

type AquilaDataRequestStruct struct {
	Text         []string `json:"text"`
	DatabaseName string   `json:"databaseName"`
}

type AquilaHubRequestStruct struct {
	Data AquilaDataRequestStruct `json:"data"`
}

// --------------------------------
// Response

type AquilaHubResponseStruct struct {
	Vectors [][]float64
	Success bool
}

// =====================================
// Db Search:
// =====================================

type DataSearchStruct struct {
	Matrix       [][]float64 `json:"matrix" bson:"matrix"`
	K            int         `json:"k" bson:"k"`
	R            int         `json:"r" bson:"r"`
	DatabaseName string      `json:"database_name" bson:"database_name"`
}

type SearchAquilaDbRequestStruct struct {
	Data      DataSearchStruct `json:"data" bson:"data"`
	Signature string           `json:"signature" bson:"signature"`
}

// --------------------------------
// Response

type MetadataSearchStruct struct {
	Age  int
	Name string
}

type DocSearchData struct {
	Cid      string
	Id       int
	Code     []float64
	Metadata MetadataSearchStruct
}

type DocSearchResponseStruct struct {
	Dist [][]float64
	Docs [][]DocSearchData
}

// =====================================
// Db Create:
// =====================================

type CreateAquilaHubResponsStruct struct {
	DatabaseName string `json:"databaseName"`
	Success      bool   `json:"success"`
}

type CreateAquilaResponsStruct struct {
	DatabaseName string `json:"database_name"`
	Success      bool   `json:"success"`
}

type MetadataStructCreateDb struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type SchemaStruct struct {
	Description string                 `json:"description"`
	Unique      string                 `json:"unique"`
	Encoder     string                 `json:"encoder"`
	Codelen     int                    `json:"codelen"`
	Metadata    MetadataStructCreateDb `json:"metadata"`
}

type DataStructCreateDb struct {
	Schema SchemaStruct `json:"schema"`
}

type CreateDbRequestStruct struct {
	Data      DataStructCreateDb `json:"data"`
	Signature string             `json:"signature"`
}

// =====================================
// Db Delete:
// =====================================

type DeleteDataStruct struct {
	Ids          []string `json:"ids" bson:"ids"`
	DatabaseName string   `json:"database_name" bson:"database_name"`
}

type DocDeleteRequestStruct struct {
	Data      DeleteDataStruct `json:"data" bson:"data"`
	Signature string           `json:"signature" bson:"signature"`
}

// ----------------------------
// Response

type DocDeleteResponseStruct struct {
	Ids     []string `json:"ids"`
	Success bool     `json:"success"`
}
