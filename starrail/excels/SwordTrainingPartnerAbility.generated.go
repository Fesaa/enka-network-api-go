package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type SwordTrainingPartnerAbility struct {
	AbilityDesc      hash.Hash `json:"AbilityDesc"`
	AbilityIcon      string    `json:"AbilityIcon"`
	AbilityName      hash.Hash `json:"AbilityName"`
	DescParamList    []float64 `json:"DescParamList"`
	EffectIDList     []float64 `json:"EffectIDList"`
	PartnerAbilityID float64   `json:"PartnerAbilityID"`
	Rare             float64   `json:"Rare"`
}
type SwordTrainingPartnerAbilityAccessor struct {
	_data                 []SwordTrainingPartnerAbility
	_dataPartnerAbilityID map[float64]SwordTrainingPartnerAbility
}

// LoadData retrieves the data. Must be called before SwordTrainingPartnerAbility.GroupData
func (a *SwordTrainingPartnerAbilityAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SwordTrainingPartnerAbility.json")
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
func (a *SwordTrainingPartnerAbilityAccessor) Raw() ([]SwordTrainingPartnerAbility, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SwordTrainingPartnerAbility{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SwordTrainingPartnerAbilityAccessor.LoadData to preload everything
func (a *SwordTrainingPartnerAbilityAccessor) GroupData() {
	for _, d := range a._data {
		a._dataPartnerAbilityID[d.PartnerAbilityID] = d
	}
}

// ByPartnerAbilityID returns the SwordTrainingPartnerAbility uniquely identified by PartnerAbilityID
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingPartnerAbilityAccessor) ByPartnerAbilityID(identifier float64) (SwordTrainingPartnerAbility, error) {
	if a._dataPartnerAbilityID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SwordTrainingPartnerAbility{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPartnerAbilityID[identifier], nil
}
