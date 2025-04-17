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

type Hash struct {
	Hash float64 `json:"Value"`
}

func (h Hash) GetHash() string {
	return fmt.Sprint(h.Hash)
}

type IntValue struct {
	IntValue float64 `json:"IntValue"`
}

func (i IntValue) GetHash() string {
	return fmt.Sprint(i.IntValue)
}

type StringValue struct {
	StringValue string `json:"StringValue"`
}

func (s StringValue) GetHash() string {
	return s.StringValue
}

type Value struct {
	Value float64 `json:"Value"`
}

func (v Value) GetHash() string {
	return fmt.Sprint(v.Value)
}
