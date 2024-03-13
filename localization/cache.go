package localization

import (
	"io"
	"os"
)

type LocalizationCache interface {
	Load(s string) ([]byte, error)
	Save(s string, data []byte) (error, bool)
}

type diskCache struct{}

func (d *diskCache) Load(s string) ([]byte, error) {
	return tryFromDisk(s)
}

func (d *diskCache) Save(s string, data []byte) (error, bool) {
	return saveToDisk(s, data)
}

func tryFromDisk(s string) ([]byte, error) {
	dir := os.Getenv("LOCALIZATION_DIR")
	if dir == "" {
		return nil, nil
	}

	f, err := os.Open(dir + s)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func saveToDisk(s string, data []byte) (error, bool) {
	dir := os.Getenv("LOCALIZATION_DIR")
	if dir == "" {
		return nil, false
	}

	f, err := os.Create(dir + s)
	if err != nil {
		return err, false
	}

	_, err = f.Write(data)
	if err != nil {
		return err, false
	}

	return nil, true
}
