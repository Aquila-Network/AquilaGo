package src

type AquilaDbInterface interface {
	CreateDatabase(createDb *CreateDbRequestStruct, url string) (*CreateAquilaResponsStruct, error)
	SignDocument() // ???
	InsertDocument(docInsert *DocInsertRequestStruct, url string) (*DocInsertResponseStruct, error)
	DeleteDocument(docDelete *DocDeleteRequestStruct, url string) (*DocDeleteResponseStruct, error)
	SearchKDocument(searchBody *SearchAquilaDbRequestStruct, url string) (*DocSearchResponseStruct, error)
}

type AquilaHubInterface interface {
	CreateHubDatabase(createDb *CreateDbRequestStruct, url string) (*CreateAquilaResponsStruct, error)
	CompressDocument(a *AquilaHubRequestStruct, url string) (*AquilaHubResponseStruct, error)
}

type AquilaDb struct {
	AquilaDbInterface
	AquilaHubInterface
}

func NewAquila(wallet *WalletStruct) *AquilaDb {
	return &AquilaDb{
		AquilaDbInterface:  NewAquilaDb(),
		AquilaHubInterface: NewAquilaHub(),
	}
}
