package cache

import (
	"log/slog"

	"github.com/Fesaa/enka-network-api-go/utils"
)

func Default() utils.ErrorSupplier[EnkaCache] {
	return func() (EnkaCache, error) {
		return NewMemoryCache(slog.Default())
	}
}
