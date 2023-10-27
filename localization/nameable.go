package localization

import "fmt"

type Nameable interface {
	GetHash() string
}

type HashNameAble struct {
	Hash int64 `json:"Hash"`
}

func (h *HashNameAble) GetHash() string {
	return fmt.Sprint(h.Hash)
}
