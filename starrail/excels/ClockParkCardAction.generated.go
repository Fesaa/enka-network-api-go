package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ClockParkCardAction struct {
	CardActionID      float64   `json:"CardActionID"`
	CardDesc          hash.Hash `json:"CardDesc"`
	DiceList          []float64 `json:"DiceList"`
	EffectList        []float64 `json:"EffectList"`
	ForeImgPath       string    `json:"ForeImgPath"`
	ImgPath           string    `json:"ImgPath"`
	ImgPath1          string    `json:"ImgPath1"`
	ImgPath2          string    `json:"ImgPath2"`
	ImgPath3          string    `json:"ImgPath3"`
	SuccessEffectList []float64 `json:"SuccessEffectList"`
}
type ClockParkCardActionAccessor struct {
	_data             []ClockParkCardAction
	_dataCardActionID map[float64]ClockParkCardAction
}

// LoadData retrieves the data. Must be called before ClockParkCardAction.GroupData
func (a *ClockParkCardActionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ClockParkCardAction.json")
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
func (a *ClockParkCardActionAccessor) Raw() ([]ClockParkCardAction, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ClockParkCardAction{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ClockParkCardActionAccessor.LoadData to preload everything
func (a *ClockParkCardActionAccessor) GroupData() {
	for _, d := range a._data {
		a._dataCardActionID[d.CardActionID] = d
	}
}

// ByCardActionID returns the ClockParkCardAction uniquely identified by CardActionID
//
// Error is only non-nil if the source errors out
func (a *ClockParkCardActionAccessor) ByCardActionID(identifier float64) (ClockParkCardAction, error) {
	if a._dataCardActionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ClockParkCardAction{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCardActionID[identifier], nil
}
