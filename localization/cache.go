package localization

import (
	"io"
	"os"
	"path/filepath"
)

type Cache interface {
	Load(s string) ([]byte, error)
	Save(s string, data []byte) (error, bool)
}

type diskCache struct{}

func (d *diskCache) Load(s string) ([]byte, error) {
	return d.tryFromDisk(s)
}

func (d *diskCache) Save(s string, data []byte) (error, bool) {
	return d.saveToDisk(s, data)
}

func (d *diskCache) tryFromDisk(s string) ([]byte, error) {
	dir := os.Getenv("LOCALIZATION_DIR")
	if dir == "" {
		dir = os.TempDir()
	}

	cachePath := filepath.Join(dir, s)
	f, err := os.Open(cachePath)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *diskCache) saveToDisk(s string, data []byte) (error, bool) {
	dir := os.Getenv("LOCALIZATION_DIR")
	if dir == "" {
		dir = os.TempDir()
	}

	cachePath := filepath.Join(dir, s)
	f, err := os.Create(cachePath)
	if err != nil {
		return err, false
	}

	_, err = f.Write(data)
	if err != nil {
		return err, false
	}

	return nil, true
}
