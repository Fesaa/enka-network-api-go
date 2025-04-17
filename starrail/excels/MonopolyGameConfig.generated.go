package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MonopolyGameConfig struct {
	BaseRaiseMaxValue  float64   `json:"BaseRaiseMaxValue"`
	GameID             float64   `json:"GameID"`
	GameIcon           string    `json:"GameIcon"`
	GameResourceIDList []float64 `json:"GameResourceIDList"`
	GameType           string    `json:"GameType"`
	IntroDesc          hash.Hash `json:"IntroDesc"`
	Name               hash.Hash `json:"Name"`
	ParamStr1          string    `json:"ParamStr1"`
	ParamStr2          string    `json:"ParamStr2"`
	RaiseCurveID       float64   `json:"RaiseCurveID"`
}
type MonopolyGameConfigAccessor struct {
	_data          []MonopolyGameConfig
	_dataGameType  map[string]MonopolyGameConfig
	_dataGameID    map[float64]MonopolyGameConfig
	_dataParamStr1 map[string]MonopolyGameConfig
}

// LoadData retrieves the data. Must be called before MonopolyGameConfig.GroupData
func (a *MonopolyGameConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonopolyGameConfig.json")
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
func (a *MonopolyGameConfigAccessor) Raw() ([]MonopolyGameConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonopolyGameConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonopolyGameConfigAccessor.LoadData to preload everything
func (a *MonopolyGameConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataGameType[d.GameType] = d
		a._dataGameID[d.GameID] = d
		a._dataParamStr1[d.ParamStr1] = d
	}
}

// ByGameType returns the MonopolyGameConfig uniquely identified by GameType
//
// Error is only non-nil if the source errors out
func (a *MonopolyGameConfigAccessor) ByGameType(identifier string) (MonopolyGameConfig, error) {
	if a._dataGameType == nil {
		err := a.LoadData()
		if err != nil {
			return MonopolyGameConfig{}, err
		}
		a.GroupData()
	}
	return a._dataGameType[identifier], nil
}

// ByGameID returns the MonopolyGameConfig uniquely identified by GameID
//
// Error is only non-nil if the source errors out
func (a *MonopolyGameConfigAccessor) ByGameID(identifier float64) (MonopolyGameConfig, error) {
	if a._dataGameID == nil {
		err := a.LoadData()
		if err != nil {
			return MonopolyGameConfig{}, err
		}
		a.GroupData()
	}
	return a._dataGameID[identifier], nil
}

// ByParamStr1 returns the MonopolyGameConfig uniquely identified by ParamStr1
//
// Error is only non-nil if the source errors out
func (a *MonopolyGameConfigAccessor) ByParamStr1(identifier string) (MonopolyGameConfig, error) {
	if a._dataParamStr1 == nil {
		err := a.LoadData()
		if err != nil {
			return MonopolyGameConfig{}, err
		}
		a.GroupData()
	}
	return a._dataParamStr1[identifier], nil
}
