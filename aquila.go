package aquiladb

import "github.com/Aquila-Network/go-aquila/src"

func AquilaModule() *src.AquilaDb {
	wallet := src.Wallet()
	return src.NewAquila(&wallet)
}
