package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ContentUnlockDescConfig struct {
	ContentID    float64   `json:"ContentID"`
	UnlockDesc01 hash.Hash `json:"UnlockDesc01"`
	UnlockDesc02 hash.Hash `json:"UnlockDesc02"`
	UnlockDesc03 hash.Hash `json:"UnlockDesc03"`
}
type ContentUnlockDescConfigAccessor struct {
	_data          []ContentUnlockDescConfig
	_dataContentID map[float64]ContentUnlockDescConfig
}

// LoadData retrieves the data. Must be called before ContentUnlockDescConfig.GroupData
func (a *ContentUnlockDescConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ContentUnlockDescConfig.json")
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
func (a *ContentUnlockDescConfigAccessor) Raw() ([]ContentUnlockDescConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ContentUnlockDescConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ContentUnlockDescConfigAccessor.LoadData to preload everything
func (a *ContentUnlockDescConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataContentID[d.ContentID] = d
	}
}

// ByContentID returns the ContentUnlockDescConfig uniquely identified by ContentID
//
// Error is only non-nil if the source errors out
func (a *ContentUnlockDescConfigAccessor) ByContentID(identifier float64) (ContentUnlockDescConfig, error) {
	if a._dataContentID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ContentUnlockDescConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataContentID[identifier], nil
}
