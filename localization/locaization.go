package localization

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"sync"
)

const (
	HONKAI_BASE_URL  = "https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/TextMap/TextMap%s.json"
	GENSHIN_BASE_URL = "https://gitlab.com/Dimbreath/AnimeGameData/-/raw/master/TextMap/TextMap%s.json"
)

var (
	LoadGenshin  = true
	LoadStarRail = true
)

var localization = &Localization{
	defaultKey:   ENGLISH,
	key:          ENGLISH,
	cache:        nil,
	hsrCache:     map[LocalizationKey]map[string]string{},
	genshinCache: map[LocalizationKey]map[string]string{},
	log:          zerolog.Nop(),
}

func Get() *Localization {
	return localization
}

type Localization struct {
	defaultKey LocalizationKey
	key        LocalizationKey

	cache Cache

	hsrCache     map[LocalizationKey]map[string]string
	genshinCache map[LocalizationKey]map[string]string

	log zerolog.Logger
}

func Init(logger zerolog.Logger, caches ...Cache) {
	var cache Cache
	if len(caches) > 0 {
		cache = caches[0]
	} else {
		cache = &diskCache{
			log: logger.With().Str("handler", "diskcache").Logger(),
		}
	}

	l := Localization{
		defaultKey: ENGLISH,
		cache:      cache,

		hsrCache:     map[LocalizationKey]map[string]string{},
		genshinCache: map[LocalizationKey]map[string]string{},

		log: logger.With().Str("handler", "localization").Logger(),
	}

	l.key = l.defaultKey
	l.loadLocalizations()

	localization = &l
}

func (l *Localization) loadLocalizations() {
	wg := sync.WaitGroup{}

	if LoadStarRail {
		wg.Add(1)
		go func() {
			l.loadHsrLocalization()
			wg.Add(-1)
		}()
	}

	if LoadGenshin {
		wg.Add(1)
		go func() {
			l.loadGenshinLocalization()
			wg.Add(-1)
		}()
	}

	wg.Wait()
}

// SetLocalization updates the localization for the given language. Removes current locale if flushOld is true
func SetLocalization(locale LocalizationKey, flushOld ...bool) {
	flush := func() bool {
		if len(flushOld) > 0 {
			return flushOld[0]
		}
		return false
	}()

	localization.key = locale
	localization.loadLocalizations()

	if flush {
		localization.log.Info().Any("locale", locale).Msg("Deleting cached localization")
		delete(localization.hsrCache, locale)
		delete(localization.genshinCache, locale)
	}
}

func (l *Localization) fetchJson(s string, url string) (*map[string]string, error) {
	cacheKey := s + string(l.key)

	data, err := l.cache.Load(cacheKey)
	if data != nil {
		l.log.Debug().Str("cacheKey", cacheKey).Str("url", url).Msg("Loaded from cache")
		return unmarshal(data)
	}

	if err != nil {
		l.log.Debug().Str("cacheKey", cacheKey).Str("url", url).Msg("Loaded from cache errored")
	} else {
		l.log.Debug().Str("cacheKey", cacheKey).Str("url", url).Msg("no data in cache")
	}

	l.log.Info().Str("url", url).Msg("Fetching localization from network")
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()
	if req.StatusCode != 200 {
		l.log.Warn().Str("url", url).Int("status", req.StatusCode).Msg("Failed to fetch localization from network")
		return nil, errors.New("enka-network-api-go: Non 200 status code returned: " + req.Status)
	}

	data, err = io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer func() {
		l.log.Debug().Str("url", url).Msg("saving to cache")
		if err = l.cache.Save(cacheKey, data); err != nil {
			l.log.Warn().Str("cacheKey", cacheKey).Str("url", url).Msg("Failed to save to cache")
		}
	}()

	return unmarshal(data)
}

func unmarshal(data []byte) (*map[string]string, error) {
	var localizationJson map[string]string
	err := json.Unmarshal(data, &localizationJson)
	if err != nil {
		return nil, err
	}

	return &localizationJson, nil
}
