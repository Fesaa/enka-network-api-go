package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type FightFestAvatarInfo struct {
	AvatarID       float64   `json:"AvatarID"`
	AvatarName     hash.Hash `json:"AvatarName"`
	FigureOffset   []float64 `json:"FigureOffset"`
	FullFigurePath string    `json:"FullFigurePath"`
	HalfFigurePath string    `json:"HalfFigurePath"`
	IconPath       string    `json:"IconPath"`
	VSImgPath      string    `json:"VSImgPath"`
}
type FightFestAvatarInfoAccessor struct {
	_data          []FightFestAvatarInfo
	_dataAvatarID  map[float64]FightFestAvatarInfo
	_dataVSImgPath map[string]FightFestAvatarInfo
}

// LoadData retrieves the data. Must be called before FightFestAvatarInfo.GroupData
func (a *FightFestAvatarInfoAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FightFestAvatarInfo.json")
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
func (a *FightFestAvatarInfoAccessor) Raw() ([]FightFestAvatarInfo, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FightFestAvatarInfo{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FightFestAvatarInfoAccessor.LoadData to preload everything
func (a *FightFestAvatarInfoAccessor) GroupData() {
	a._dataAvatarID = map[float64]FightFestAvatarInfo{}
	a._dataVSImgPath = map[string]FightFestAvatarInfo{}
	for _, d := range a._data {
		a._dataAvatarID[d.AvatarID] = d
		a._dataVSImgPath[d.VSImgPath] = d
	}
}

// ByAvatarID returns the FightFestAvatarInfo uniquely identified by AvatarID
//
// Error is only non-nil if the source errors out
func (a *FightFestAvatarInfoAccessor) ByAvatarID(identifier float64) (FightFestAvatarInfo, error) {
	if a._dataAvatarID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestAvatarInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarID[identifier], nil
}

// ByVSImgPath returns the FightFestAvatarInfo uniquely identified by VSImgPath
//
// Error is only non-nil if the source errors out
func (a *FightFestAvatarInfoAccessor) ByVSImgPath(identifier string) (FightFestAvatarInfo, error) {
	if a._dataVSImgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestAvatarInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataVSImgPath[identifier], nil
}
