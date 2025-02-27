package localization

import "fmt"

type Nameable interface {
	GetHash() string
}

type HashNameAble struct {
	Hash string `json:"Hash"`
}

func (h *HashNameAble) GetHash() string {
	return fmt.Sprint(h.Hash)
}

type HashString struct {
	Hash string
}

func (h *HashString) GetHash() string {
	return h.Hash
}

type HashInt struct {
	Hash int
}

func (h *HashInt) GetHash() string {
	return fmt.Sprint(h.Hash)
}
