package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MuseumAreaConfig struct {
	AreaID             float64   `json:"AreaID"`
	AreaItemNoTextID   hash.Hash `json:"AreaItemNoTextID"`
	FirstWorldText     string    `json:"FirstWorldText"`
	MuseumAreaHintIcon string    `json:"MuseumAreaHintIcon"`
	MuseumAreaName     hash.Hash `json:"MuseumAreaName"`
	MuseumAreaTabIcon  string    `json:"MuseumAreaTabIcon"`
}
type MuseumAreaConfigAccessor struct {
	_data                   []MuseumAreaConfig
	_dataAreaID             map[float64]MuseumAreaConfig
	_dataMuseumAreaHintIcon map[string]MuseumAreaConfig
	_dataMuseumAreaTabIcon  map[string]MuseumAreaConfig
}

// LoadData retrieves the data. Must be called before MuseumAreaConfig.GroupData
func (a *MuseumAreaConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MuseumAreaConfig.json")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &a._data)
}

// Raw returns the raw data.
func (a *MuseumAreaConfigAccessor) Raw() ([]MuseumAreaConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MuseumAreaConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MuseumAreaConfigAccessor.LoadData to preload everything
func (a *MuseumAreaConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataAreaID[d.AreaID] = d
		a._dataMuseumAreaHintIcon[d.MuseumAreaHintIcon] = d
		a._dataMuseumAreaTabIcon[d.MuseumAreaTabIcon] = d
	}
}

// ByAreaID returns the MuseumAreaConfig uniquely identified by AreaID
//
// Error is only non-nil if the source errors out
func (a *MuseumAreaConfigAccessor) ByAreaID(identifier float64) (MuseumAreaConfig, error) {
	if a._dataAreaID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MuseumAreaConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAreaID[identifier], nil
}

// ByMuseumAreaHintIcon returns the MuseumAreaConfig uniquely identified by MuseumAreaHintIcon
//
// Error is only non-nil if the source errors out
func (a *MuseumAreaConfigAccessor) ByMuseumAreaHintIcon(identifier string) (MuseumAreaConfig, error) {
	if a._dataMuseumAreaHintIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MuseumAreaConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMuseumAreaHintIcon[identifier], nil
}

// ByMuseumAreaTabIcon returns the MuseumAreaConfig uniquely identified by MuseumAreaTabIcon
//
// Error is only non-nil if the source errors out
func (a *MuseumAreaConfigAccessor) ByMuseumAreaTabIcon(identifier string) (MuseumAreaConfig, error) {
	if a._dataMuseumAreaTabIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MuseumAreaConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMuseumAreaTabIcon[identifier], nil
}
