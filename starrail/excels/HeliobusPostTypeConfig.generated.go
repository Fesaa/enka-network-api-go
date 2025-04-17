package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type HeliobusPostTypeConfig struct {
	PostType                   string    `json:"PostType"`
	PostTypeIconPath           string    `json:"PostTypeIconPath"`
	PostTypeIconPathUnselected string    `json:"PostTypeIconPathUnselected"`
	PostTypeName               hash.Hash `json:"PostTypeName"`
}
type HeliobusPostTypeConfigAccessor struct {
	_data         []HeliobusPostTypeConfig
	_dataPostType map[string]HeliobusPostTypeConfig
}

// LoadData retrieves the data. Must be called before HeliobusPostTypeConfig.GroupData
func (a *HeliobusPostTypeConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HeliobusPostTypeConfig.json")
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
func (a *HeliobusPostTypeConfigAccessor) Raw() ([]HeliobusPostTypeConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HeliobusPostTypeConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with HeliobusPostTypeConfigAccessor.LoadData to preload everything
func (a *HeliobusPostTypeConfigAccessor) GroupData() {
	a._dataPostType = map[string]HeliobusPostTypeConfig{}
	for _, d := range a._data {
		a._dataPostType[d.PostType] = d
	}
}

// ByPostType returns the HeliobusPostTypeConfig uniquely identified by PostType
//
// Error is only non-nil if the source errors out
func (a *HeliobusPostTypeConfigAccessor) ByPostType(identifier string) (HeliobusPostTypeConfig, error) {
	if a._dataPostType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusPostTypeConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPostType[identifier], nil
}
