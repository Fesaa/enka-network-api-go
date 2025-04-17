package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueDLCBossDecay struct {
	BossDecayComeFrom hash.Hash    `json:"BossDecayComeFrom"`
	BossDecayDesc     hash.Hash    `json:"BossDecayDesc"`
	BossDecayID       float64      `json:"BossDecayID"`
	BossDecayName     hash.Hash    `json:"BossDecayName"`
	BossEffectIcon    string       `json:"BossEffectIcon"`
	DecayIcon         string       `json:"DecayIcon"`
	DescParam         []hash.Value `json:"DescParam"`
	EffectParamList   []float64    `json:"EffectParamList"`
	EffectType        string       `json:"EffectType"`
	ExtraDesc         []float64    `json:"ExtraDesc"`
}
type RogueDLCBossDecayAccessor struct {
	_data            []RogueDLCBossDecay
	_dataBossDecayID map[float64]RogueDLCBossDecay
}

// LoadData retrieves the data. Must be called before RogueDLCBossDecay.GroupData
func (a *RogueDLCBossDecayAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCBossDecay.json")
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
func (a *RogueDLCBossDecayAccessor) Raw() ([]RogueDLCBossDecay, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCBossDecay{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCBossDecayAccessor.LoadData to preload everything
func (a *RogueDLCBossDecayAccessor) GroupData() {
	a._dataBossDecayID = map[float64]RogueDLCBossDecay{}
	for _, d := range a._data {
		a._dataBossDecayID[d.BossDecayID] = d
	}
}

// ByBossDecayID returns the RogueDLCBossDecay uniquely identified by BossDecayID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCBossDecayAccessor) ByBossDecayID(identifier float64) (RogueDLCBossDecay, error) {
	if a._dataBossDecayID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCBossDecay{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBossDecayID[identifier], nil
}
