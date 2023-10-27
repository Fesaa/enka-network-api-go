package localization

import "fmt"

func (l *Localization) loadHonkaiLocalization() {
	if _, ok := l.honkaiLocalizationCache[l.key]; ok {
		return
	}

	l.log.Infof("(Honkai) Loading new localization %s...", l.key)
	localizationMap, err := l.fetchJson(fmt.Sprintf(HONKAI_BASE_URL, string(l.key)))
	if err != nil {
		l.log.Errorf("(Honkai) Couldn't load localization for %s, falling back to %s. \n Error: %s", l.key, l.defaultKey, err.Error())
		l.key = l.defaultKey
		l.honkaiLocalizationCache = make(map[LocalizationKey]LocalizationMap)
		return
	}

	l.honkaiLocalizationCache[l.key] = *localizationMap
	l.log.Infof("(Honkai) Loaded localization %s!", l.key)

}

// GetHonkaiLocale tries to retrieve the string for a hash
// in Honkai: Star Rail
//
// Takes a nameable and returns a nilable string
//
// Methods on structs will use GetHonkaiLocaleOrHash, call this directly
// for full control over it's behavior
func GetHonkaiLocale(nameable Nameable) *string {
	if m, ok := localization.honkaiLocalizationCache[localization.key]; ok {
		if str, ok := m[nameable.GetHash()]; ok {
			return &str
		}
	}
	return nil
}

// GetHonkaiLocaleOrHash tries to retrieve the string for a hash
// in Honkai: Star Rail
//
// Returns the hash if string is not found
func GetHonkaiLocaleOrHash(nameable Nameable) string {
	name := GetHonkaiLocale(nameable)
	if name != nil {
		return *name
	}

	return nameable.GetHash()
}
