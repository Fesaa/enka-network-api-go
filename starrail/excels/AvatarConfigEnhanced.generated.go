package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AvatarConfigEnhanced struct {
	AvatarID   float64    `json:"AvatarID"`
	EnhancedID float64    `json:"EnhancedID"`
	JsonPath   string     `json:"JsonPath"`
	RankIDList []float64  `json:"RankIDList"`
	SPNeed     hash.Value `json:"SPNeed"`
	SkillList  []float64  `json:"SkillList"`
}
type AvatarConfigEnhancedAccessor struct {
	_data         []AvatarConfigEnhanced
	_dataAvatarID map[float64]AvatarConfigEnhanced
	_dataJsonPath map[string]AvatarConfigEnhanced
}

// LoadData retrieves the data. Must be called before AvatarConfigEnhanced.GroupData
func (a *AvatarConfigEnhancedAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarConfigEnhanced.json")
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
func (a *AvatarConfigEnhancedAccessor) Raw() ([]AvatarConfigEnhanced, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarConfigEnhanced{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AvatarConfigEnhancedAccessor.LoadData to preload everything
func (a *AvatarConfigEnhancedAccessor) GroupData() {
	a._dataAvatarID = map[float64]AvatarConfigEnhanced{}
	a._dataJsonPath = map[string]AvatarConfigEnhanced{}
	for _, d := range a._data {
		a._dataAvatarID[d.AvatarID] = d
		a._dataJsonPath[d.JsonPath] = d
	}
}

// ByAvatarID returns the AvatarConfigEnhanced uniquely identified by AvatarID
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigEnhancedAccessor) ByAvatarID(identifier float64) (AvatarConfigEnhanced, error) {
	if a._dataAvatarID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigEnhanced{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarID[identifier], nil
}

// ByJsonPath returns the AvatarConfigEnhanced uniquely identified by JsonPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigEnhancedAccessor) ByJsonPath(identifier string) (AvatarConfigEnhanced, error) {
	if a._dataJsonPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigEnhanced{}, err
			}
		}
		a.GroupData()
	}
	return a._dataJsonPath[identifier], nil
}
