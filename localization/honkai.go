package localization

import (
	"fmt"
	"github.com/Fesaa/enka-network-api-go/hash"
)

func (l *Localization) loadHsrLocalization() {
	if _, ok := l.hsrCache[l.key]; ok {
		return
	}

	l.log.Info().Str("game", "hsr").Any("key", l.key).Msg("loading localization")
	localizationMap, err := l.fetchJson("honkai_", fmt.Sprintf(HONKAI_BASE_URL, string(l.key)))
	if err != nil {
		if l.key != l.defaultKey {
			l.log.Error().Err(err).Str("game", "hsr").
				Any("key", l.key).
				Any("defaultKey", l.defaultKey).
				Msg("falling back after failure")
			l.key = l.defaultKey
			l.loadHsrLocalization()
			return
		}
		l.key = l.defaultKey
		l.hsrCache = make(map[LocalizationKey]map[string]string)
		l.log.Error().Err(err).Str("game", "hsr").Any("key", l.key).
			Msg("failed to load, localization won't be avaible")
		return
	}

	l.hsrCache[l.key] = *localizationMap
	l.log.Info().Str("game", "hsr").Any("key", l.key).Msg("loaded localization")
}

// GetHsrLocale tries to retrieve the string for a hash
// in Honkai: Star Rail. Takes a nameable and returns a nilable string
//
// Methods on structs will use GetHsrLocaleOrHash, call this directly
// for full control over its behavior
func GetHsrLocale(nameable hash.Nameable) *string {
	if m, ok := localization.hsrCache[localization.key]; ok {
		if str, ok := m[nameable.GetHash()]; ok {
			return &str
		}
	}
	return nil
}

// GetHsrLocaleOrHash tries to retrieve the string for a hash
// in Honkai: Star Rail
//
// Returns the hash if string is not found
func GetHsrLocaleOrHash(nameable hash.Nameable) string {
	name := GetHsrLocale(nameable)
	if name != nil {
		return *name
	}

	return nameable.GetHash()
}
