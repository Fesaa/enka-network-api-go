package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type AetherDivideBadge struct {
	ItemID         float64 `json:"ItemID"`
	MaxSpiritLevel float64 `json:"MaxSpiritLevel"`
	Number         float64 `json:"Number"`
}
type AetherDivideBadgeAccessor struct {
	_data               []AetherDivideBadge
	_dataMaxSpiritLevel map[float64]AetherDivideBadge
}

// LoadData retrieves the data. Must be called before AetherDivideBadge.GroupData
func (a *AetherDivideBadgeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AetherDivideBadge.json")
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
func (a *AetherDivideBadgeAccessor) Raw() ([]AetherDivideBadge, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AetherDivideBadge{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AetherDivideBadgeAccessor.LoadData to preload everything
func (a *AetherDivideBadgeAccessor) GroupData() {
	for _, d := range a._data {
		a._dataMaxSpiritLevel[d.MaxSpiritLevel] = d
	}
}

// ByMaxSpiritLevel returns the AetherDivideBadge uniquely identified by MaxSpiritLevel
//
// Error is only non-nil if the source errors out
func (a *AetherDivideBadgeAccessor) ByMaxSpiritLevel(identifier float64) (AetherDivideBadge, error) {
	if a._dataMaxSpiritLevel == nil {
		err := a.LoadData()
		if err != nil {
			return AetherDivideBadge{}, err
		}
		a.GroupData()
	}
	return a._dataMaxSpiritLevel[identifier], nil
}
