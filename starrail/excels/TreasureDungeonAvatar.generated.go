package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TreasureDungeonAvatar struct {
	AvatarPickID   float64   `json:"AvatarPickID"`
	Dialogue1      hash.Hash `json:"Dialogue1"`
	FigureDiff     []float64 `json:"FigureDiff"`
	FigureScale    float64   `json:"FigureScale"`
	SpecialAvataID float64   `json:"SpecialAvataID"`
}
type TreasureDungeonAvatarAccessor struct {
	_data             []TreasureDungeonAvatar
	_dataAvatarPickID map[float64]TreasureDungeonAvatar
}

// LoadData retrieves the data. Must be called before TreasureDungeonAvatar.GroupData
func (a *TreasureDungeonAvatarAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TreasureDungeonAvatar.json")
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
func (a *TreasureDungeonAvatarAccessor) Raw() ([]TreasureDungeonAvatar, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TreasureDungeonAvatar{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TreasureDungeonAvatarAccessor.LoadData to preload everything
func (a *TreasureDungeonAvatarAccessor) GroupData() {
	for _, d := range a._data {
		a._dataAvatarPickID[d.AvatarPickID] = d
	}
}

// ByAvatarPickID returns the TreasureDungeonAvatar uniquely identified by AvatarPickID
//
// Error is only non-nil if the source errors out
func (a *TreasureDungeonAvatarAccessor) ByAvatarPickID(identifier float64) (TreasureDungeonAvatar, error) {
	if a._dataAvatarPickID == nil {
		err := a.LoadData()
		if err != nil {
			return TreasureDungeonAvatar{}, err
		}
		a.GroupData()
	}
	return a._dataAvatarPickID[identifier], nil
}
