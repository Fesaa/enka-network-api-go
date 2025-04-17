package main

import (
	"errors"
	"github.com/Fesaa/yoitsu"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	yoitsu.SliceNameFormatter = func(s string) string {
		return s
	}

	jsons, err := os.ReadDir("../jsons/")
	if err != nil {
		panic(err)
	}

	for _, json := range jsons {
		ext := filepath.Ext(json.Name())
		if ext != ".json" {
			continue
		}

		base := strings.TrimSuffix(json.Name(), ext)

		err = generateJson(base)
		if err != nil && !errors.Is(err, yoitsu.ErrNoData) {
			panic(err)
		}
	}
}

func generateJson(name string) error {
	src := &LocalFileExtUrl{name: name}

	y := yoitsu.New(src,
		yoitsu.WithPackageName("excels"),
		yoitsu.WithUniverse(getUniverse()),
		yoitsu.WithGenerateAccessors(func(accessors *yoitsu.Accessors) {
			accessors.Generate = true
			accessors.ById = true
		}))

	if err := y.GenerateFile(); err != nil {
		return err
	}

	return y.WriteToDisk("../")
}

func getUniverse() yoitsu.Universe {
	universe := yoitsu.NewUniverse()
	universe.AddType(&yoitsu.StructType{
		Name:   "hash.Hash",
		Import: "github.com/Fesaa/enka-network-api-go/hash",
		Fields: map[string]*yoitsu.StructField{
			"Hash": {
				Type: yoitsu.Float64Type,
				Tag:  "Hash",
			},
		},
	})
	universe.AddType(&yoitsu.StructType{
		Name:   "hash.IntValue",
		Import: "github.com/Fesaa/enka-network-api-go/hash",
		Fields: map[string]*yoitsu.StructField{
			"IntValue": {
				Type: yoitsu.Float64Type,
				Tag:  "IntValue",
			},
		},
	})
	universe.AddType(&yoitsu.StructType{
		Name:   "hash.Value",
		Import: "github.com/Fesaa/enka-network-api-go/hash",
		Fields: map[string]*yoitsu.StructField{
			"Value": {
				Type: yoitsu.Float64Type,
				Tag:  "Value",
			},
		},
	})
	universe.AddType(&yoitsu.StructType{
		Name:   "hash.StringValue",
		Import: "github.com/Fesaa/enka-network-api-go/hash",
		Fields: map[string]*yoitsu.StructField{
			"StringValue": {
				Type: yoitsu.StringType,
				Tag:  "StringValue",
			},
		},
	})

	return universe
}
