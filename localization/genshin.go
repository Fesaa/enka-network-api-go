package localization

import (
	"fmt"
	"github.com/Fesaa/enka-network-api-go/hash"
)

func (l *Localization) loadGenshinLocalization() {
	if _, ok := l.genshinCache[l.key]; ok {
		return
	}

	l.log.Info().Str("game", "genshin").Any("key", l.key).Msg("loading localization")
	localizationMap, err := l.fetchJson("genshin_", fmt.Sprintf(GENSHIN_BASE_URL, string(l.key)))
	if err != nil {
		if l.key != l.defaultKey {
			l.log.Error().Err(err).Str("game", "genshin").
				Any("key", l.key).
				Any("defaultKey", l.defaultKey).
				Msg("falling back after failure")
			l.key = l.defaultKey
			l.loadHsrLocalization()
			return
		}
		l.key = l.defaultKey
		l.genshinCache = make(map[LocalizationKey]map[string]string)
		l.log.Error().Err(err).Str("game", "genshin").Any("key", l.key).
			Msg("failed to load, localization won't be avaible")
		return
	}

	l.genshinCache[l.key] = *localizationMap
	l.log.Info().Str("game", "genshin").Any("key", l.key).Msg("loaded localization")
}

// GetGenshinLocale tries to retrieve the string for a hash
// in Genish impact. Takes a nameable and returns a nilable string
//
// Methods on structs will use GetGensinLocaleOrHash, call this directly
// for full control over its behavior
func GetGenshinLocale(nameable hash.Nameable) *string {
	if m, ok := localization.genshinCache[localization.key]; ok {
		if name, ok := m[nameable.GetHash()]; ok {
			return &name
		}
	}
	return nil
}

// GetGenshinLocaleOrHash tries to retrieve the string for a hash
// in Genish impact
//
// Returns the hash if the string is not found
func GetGenshinLocaleOrHash(nameable hash.Nameable) string {
	name := GetGenshinLocale(nameable)
	if name != nil {
		return *name
	}
	return nameable.GetHash()
}
