package hash

import "fmt"

type Nameable interface {
	GetHash() string
}

type String struct {
	Hash string
}

func (h String) GetHash() string {
	return h.Hash
}

type Int struct {
	Hash int
}

func (h Int) GetHash() string {
	return fmt.Sprint(h.Hash)
}

type Int64 struct {
	Hash int64
}

func (h Int64) GetHash() string {
	return fmt.Sprint(h.Hash)
}

type UInt64 struct {
	Hash uint64
}

func (h UInt64) GetHash() string {
	return fmt.Sprint(h.Hash)
}
