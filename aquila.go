package aquiladb

import "github.com/Aquila-Network/go-aquila/src"

func AquilaModule(wallet src.WalletStruct) *src.AquilaDb {
	// var wallet src.WalletStruct
	return src.NewAquila(wallet)
}
