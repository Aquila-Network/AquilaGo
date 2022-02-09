package aquiladb

import "github.com/Aquila-Network/go-aquila/src"

func InitAquilaDb() *src.AquilaDb {
	wallet := src.Wallet()
	return src.NewAquilaDb(&wallet)
}
