package localization

import (
	"github.com/rs/zerolog"
	"io"
	"os"
	"path/filepath"
)

type Cache interface {
	Load(s string) ([]byte, error)
	Save(s string, data []byte) error
}

type diskCache struct {
	log zerolog.Logger
}

func (d *diskCache) Load(s string) ([]byte, error) {
	dir := os.Getenv("LOCALIZATION_DIR")
	if dir == "" {
		dir = os.TempDir()
	}

	cachePath := filepath.Join(dir, s)
	d.log.Debug().Str("cachePath", cachePath).Msg("Trying to load from disk")
	f, err := os.Open(cachePath)
	if err != nil {
		d.log.Debug().Err(err).Str("cachePath", cachePath).Msg("Failed to load from disk")
		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *diskCache) Save(s string, data []byte) error {
	dir := os.Getenv("LOCALIZATION_DIR")
	if dir == "" {
		dir = os.TempDir()
	}

	cachePath := filepath.Join(dir, s)
	d.log.Debug().Str("cachePath", cachePath).Msg("Trying to save to disk")
	f, err := os.Create(cachePath)
	if err != nil {
		d.log.Debug().Err(err).Str("cachePath", cachePath).Msg("Failed to save to disk")
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		d.log.Debug().Err(err).Str("cachePath", cachePath).Msg("Failed to save to disk")
		return err
	}

	return nil
}
