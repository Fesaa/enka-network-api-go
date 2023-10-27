package enkanetworkapigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/withmandala/go-log"
)

type LocalizationKey string

const (
	TRADITIONAL_CHINESE LocalizationKey = "CHT"
	SIMPLIFIED_CHINESE  LocalizationKey = "CHS"
	GERMAN              LocalizationKey = "DE"
	ENGLISH             LocalizationKey = "EN"
	ESPANOL             LocalizationKey = "ES"
	FRANCAIS            LocalizationKey = "FR"
	INDONESIAN          LocalizationKey = "ID"
	// Do not use for Honkai: Star Rail
	ITALIAN    LocalizationKey = "IT"
	JAPANESE   LocalizationKey = "JP"
	KOREAN     LocalizationKey = "KR"
	PORTUGUESE LocalizationKey = "PT"
	RUSSIAN    LocalizationKey = "RU"
	THAI       LocalizationKey = "TH"
	// Do not use for Honkai: Star Rail
	TURKISH    LocalizationKey = "TR"
	VIETNAMESE LocalizationKey = "VI"
)

const (
	HONKAI_BASE_URL = "https://raw.githubusercontent.com/Dimbreath/StarRailData/master/TextMap/TextMap%s.json"
)

type Nameable interface {
	GetHash() string
}

type HashNameAble struct {
	Hash int64 `json:"Hash"`
}

func (h *HashNameAble) GetHash() string {
	return fmt.Sprint(h.Hash)
}

type LocalizationMap map[string]string

type Localization struct {
	defaultKey LocalizationKey
	key        LocalizationKey

	honkaiLocalizationCache map[LocalizationKey]LocalizationMap

	log *log.Logger
}

func newLocalization() Localization {
	l := Localization{
		defaultKey: ENGLISH,

		honkaiLocalizationCache: map[LocalizationKey]LocalizationMap{},

		log: log.New(os.Stdout).WithColor(),
	}

	l.key = l.defaultKey
	l.loadLocalizations()

	return l
}

func (l *Localization) loadLocalizations() {
	l.loadHonkaiLocalization()
}

func (l *Localization) SetLocalization(locale LocalizationKey) {
	l.key = locale
	l.loadLocalizations()
}

func (l *Localization) SetDebug(debug bool) {
	if debug {
		l.log.WithDebug()
	} else {
		l.log.WithoutDebug()
	}
}

func (l *Localization) loadHonkaiLocalization() {
	if _, ok := l.honkaiLocalizationCache[l.key]; ok {
		return
	}

	l.log.Infof("(Honkai) Loading new localization %s...", l.key)
	localizationMap, err := l.fetchJson(fmt.Sprintf(HONKAI_BASE_URL, string(l.key)))
	if err != nil {
		l.log.Errorf("(Honkai) Couldn't load localization for %s, falling back to %s. \n Error: %s", l.key, l.defaultKey, err.Error())
		l.key = l.defaultKey
		return
	}

	l.honkaiLocalizationCache[l.key] = *localizationMap
	l.log.Infof("(Honkai) Loaded localization %s!", l.key)

}

func (l *Localization) fetchJson(url string) (*LocalizationMap, error) {

	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if req.StatusCode != 200 {
		l.log.Debugf("Returned a non 200 status code. Got %d", req.StatusCode)
		return nil, errors.New("enka-network-api-go: Non 200 status code returned: " + req.Status)
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var localizationJson LocalizationMap
	err = json.Unmarshal(data, &localizationJson)
	if err != nil {
		return nil, err
	}

	return &localizationJson, nil
}

// GetHonkaiLocale tries to retrieve the string for a hash
// in Honkai: Star Rail
// Takes a nameable and returns a nilable string
func (l *Localization) GetHonkaiLocale(nameable Nameable) *string {
	if m, ok := l.honkaiLocalizationCache[l.key]; ok {
		if str, ok := m[nameable.GetHash()]; ok {
			return &str
		}
	}
	return nil
}
