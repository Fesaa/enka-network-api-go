package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type EquipmentConfig struct {
	AvatarBaseType       string    `json:"AvatarBaseType"`
	AvatarDetailOffset   []float64 `json:"AvatarDetailOffset"`
	BattleDialogOffset   []float64 `json:"BattleDialogOffset"`
	CoinCost             float64   `json:"CoinCost"`
	EquipmentID          float64   `json:"EquipmentID"`
	EquipmentName        hash.Hash `json:"EquipmentName"`
	ExpProvide           float64   `json:"ExpProvide"`
	ExpType              float64   `json:"ExpType"`
	GachaResultOffset    []float64 `json:"GachaResultOffset"`
	ImagePath            string    `json:"ImagePath"`
	ItemRightPanelOffset []float64 `json:"ItemRightPanelOffset"`
	MaxPromotion         float64   `json:"MaxPromotion"`
	MaxRank              float64   `json:"MaxRank"`
	RankUpCostList       []float64 `json:"RankUpCostList"`
	Rarity               string    `json:"Rarity"`
	Release              bool      `json:"Release"`
	SkillID              float64   `json:"SkillID"`
	ThumbnailPath        string    `json:"ThumbnailPath"`
}
type EquipmentConfigAccessor struct {
	_data              []EquipmentConfig
	_dataEquipmentID   map[float64]EquipmentConfig
	_dataImagePath     map[string]EquipmentConfig
	_dataSkillID       map[float64]EquipmentConfig
	_dataThumbnailPath map[string]EquipmentConfig
}

// LoadData retrieves the data. Must be called before EquipmentConfig.GroupData
func (a *EquipmentConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EquipmentConfig.json")
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
func (a *EquipmentConfigAccessor) Raw() ([]EquipmentConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EquipmentConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EquipmentConfigAccessor.LoadData to preload everything
func (a *EquipmentConfigAccessor) GroupData() {
	a._dataEquipmentID = map[float64]EquipmentConfig{}
	a._dataImagePath = map[string]EquipmentConfig{}
	a._dataSkillID = map[float64]EquipmentConfig{}
	a._dataThumbnailPath = map[string]EquipmentConfig{}
	for _, d := range a._data {
		a._dataEquipmentID[d.EquipmentID] = d
		a._dataImagePath[d.ImagePath] = d
		a._dataSkillID[d.SkillID] = d
		a._dataThumbnailPath[d.ThumbnailPath] = d
	}
}

// ByEquipmentID returns the EquipmentConfig uniquely identified by EquipmentID
//
// Error is only non-nil if the source errors out
func (a *EquipmentConfigAccessor) ByEquipmentID(identifier float64) (EquipmentConfig, error) {
	if a._dataEquipmentID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EquipmentConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEquipmentID[identifier], nil
}

// ByImagePath returns the EquipmentConfig uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *EquipmentConfigAccessor) ByImagePath(identifier string) (EquipmentConfig, error) {
	if a._dataImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EquipmentConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}

// BySkillID returns the EquipmentConfig uniquely identified by SkillID
//
// Error is only non-nil if the source errors out
func (a *EquipmentConfigAccessor) BySkillID(identifier float64) (EquipmentConfig, error) {
	if a._dataSkillID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EquipmentConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSkillID[identifier], nil
}

// ByThumbnailPath returns the EquipmentConfig uniquely identified by ThumbnailPath
//
// Error is only non-nil if the source errors out
func (a *EquipmentConfigAccessor) ByThumbnailPath(identifier string) (EquipmentConfig, error) {
	if a._dataThumbnailPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EquipmentConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataThumbnailPath[identifier], nil
}
