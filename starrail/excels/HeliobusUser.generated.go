package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type HeliobusUser struct {
	HeliobusUserID   float64   `json:"HeliobusUserID"`
	HeliobusUserName hash.Hash `json:"HeliobusUserName"`
	UserIconPath     string    `json:"UserIconPath"`
}
type HeliobusUserAccessor struct {
	_data               []HeliobusUser
	_dataHeliobusUserID map[float64]HeliobusUser
	_dataUserIconPath   map[string]HeliobusUser
}

// LoadData retrieves the data. Must be called before HeliobusUser.GroupData
func (a *HeliobusUserAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HeliobusUser.json")
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
func (a *HeliobusUserAccessor) Raw() ([]HeliobusUser, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HeliobusUser{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with HeliobusUserAccessor.LoadData to preload everything
func (a *HeliobusUserAccessor) GroupData() {
	a._dataHeliobusUserID = map[float64]HeliobusUser{}
	a._dataUserIconPath = map[string]HeliobusUser{}
	for _, d := range a._data {
		a._dataHeliobusUserID[d.HeliobusUserID] = d
		a._dataUserIconPath[d.UserIconPath] = d
	}
}

// ByHeliobusUserID returns the HeliobusUser uniquely identified by HeliobusUserID
//
// Error is only non-nil if the source errors out
func (a *HeliobusUserAccessor) ByHeliobusUserID(identifier float64) (HeliobusUser, error) {
	if a._dataHeliobusUserID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusUser{}, err
			}
		}
		a.GroupData()
	}
	return a._dataHeliobusUserID[identifier], nil
}

// ByUserIconPath returns the HeliobusUser uniquely identified by UserIconPath
//
// Error is only non-nil if the source errors out
func (a *HeliobusUserAccessor) ByUserIconPath(identifier string) (HeliobusUser, error) {
	if a._dataUserIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusUser{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUserIconPath[identifier], nil
}
