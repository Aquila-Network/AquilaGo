package src

type AquilaDbInterface interface {
	CreateDatabase(createDb *DataStructCreateDb, url string) (*CreateAquilaResponsStruct, error)
	SignDocument() // ???
	InsertDocument(docInsert *DatatDocInsertStruct, url string) (*DocInsertResponseStruct, error)
	DeleteDocument(docDelete *DeleteDataStruct, url string) (*DocDeleteResponseStruct, error)
	SearchKDocument(searchBody *DataSearchStruct, url string) (*DocSearchResponseStruct, error)
}

type AquilaHubInterface interface {
	CreateDatabase(createDb *DataStructCreateDb, url string) (*CreateAquilaHubResponsStruct, error)
	CompressDocument(a *AquilaHubRequestStruct, url string) (*AquilaHubResponseStruct, error)
}

type AquilaDb struct {
	AquilaDbInterface
	AquilaHubInterface
}

func NewAquila(wallet WalletStruct) *AquilaDb {
	return &AquilaDb{
		AquilaDbInterface:  NewAquilaDb(wallet),
		AquilaHubInterface: NewAquilaHub(wallet),
	}
}
