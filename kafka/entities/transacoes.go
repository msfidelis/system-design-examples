package entities

import "github.com/uptrace/bun"

type Transacao struct {
	bun.BaseModel `bun:"table:transacoes,alias:t"`
	ID            string  `json:"id" faker:"uuid_hyphenated"`
	Valor         float64 `json:"valor" faker:"amount"`
	Tipo          string  `json:"tipo" faker:"oneof: c, d"`
	Descricao     string  `json:"descricao" faker:"sentence"`
	RealizadaEm   string  `json:"realizada_em" faker:"timestamp"`
}
