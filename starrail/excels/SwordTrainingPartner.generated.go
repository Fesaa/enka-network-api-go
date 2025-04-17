package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type SwordTrainingPartner struct {
	AvatarID             float64   `json:"AvatarID"`
	PartnerAbilityIDList []float64 `json:"PartnerAbilityIDList"`
	PartnerID            float64   `json:"PartnerID"`
	PartnerImage         string    `json:"PartnerImage"`
	PartnerName          hash.Hash `json:"PartnerName"`
}
type SwordTrainingPartnerAccessor struct {
	_data             []SwordTrainingPartner
	_dataPartnerID    map[float64]SwordTrainingPartner
	_dataPartnerImage map[string]SwordTrainingPartner
}

// LoadData retrieves the data. Must be called before SwordTrainingPartner.GroupData
func (a *SwordTrainingPartnerAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SwordTrainingPartner.json")
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
func (a *SwordTrainingPartnerAccessor) Raw() ([]SwordTrainingPartner, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SwordTrainingPartner{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SwordTrainingPartnerAccessor.LoadData to preload everything
func (a *SwordTrainingPartnerAccessor) GroupData() {
	for _, d := range a._data {
		a._dataPartnerID[d.PartnerID] = d
		a._dataPartnerImage[d.PartnerImage] = d
	}
}

// ByPartnerID returns the SwordTrainingPartner uniquely identified by PartnerID
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingPartnerAccessor) ByPartnerID(identifier float64) (SwordTrainingPartner, error) {
	if a._dataPartnerID == nil {
		err := a.LoadData()
		if err != nil {
			return SwordTrainingPartner{}, err
		}
		a.GroupData()
	}
	return a._dataPartnerID[identifier], nil
}

// ByPartnerImage returns the SwordTrainingPartner uniquely identified by PartnerImage
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingPartnerAccessor) ByPartnerImage(identifier string) (SwordTrainingPartner, error) {
	if a._dataPartnerImage == nil {
		err := a.LoadData()
		if err != nil {
			return SwordTrainingPartner{}, err
		}
		a.GroupData()
	}
	return a._dataPartnerImage[identifier], nil
}
