package aquiladb

import "github.com/Aquila-Network/go-aquila/src"

func AquilaModule() *src.AquilaDb {
	return src.NewAquila()
}
