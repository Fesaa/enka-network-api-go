package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueDLCAeonDiceSurface struct {
	AeonDiceID        float64         `json:"AeonDiceID"`
	AeonSurfaceDiceID float64         `json:"AeonSurfaceDiceID"`
	DescParam         []hash.IntValue `json:"DescParam"`
	Dice3DSurfaceList []float64       `json:"Dice3DSurfaceList"`
	DiceActiveStage   float64         `json:"DiceActiveStage"`
	DiceEffectParam   []float64       `json:"DiceEffectParam"`
	DiceEffectType    string          `json:"DiceEffectType"`
	DiceSurfaceDesc   hash.Hash       `json:"DiceSurfaceDesc"`
	DiceSurfaceIcon   string          `json:"DiceSurfaceIcon"`
	DiceSurfaceName   hash.Hash       `json:"DiceSurfaceName"`
	ExtraEffect       []float64       `json:"ExtraEffect"`
	Rarity            float64         `json:"Rarity"`
	Sort              float64         `json:"Sort"`
}
type RogueDLCAeonDiceSurfaceAccessor struct {
	_data                  []RogueDLCAeonDiceSurface
	_dataAeonSurfaceDiceID map[float64]RogueDLCAeonDiceSurface
	_dataDiceSurfaceIcon   map[string]RogueDLCAeonDiceSurface
}

// LoadData retrieves the data. Must be called before RogueDLCAeonDiceSurface.GroupData
func (a *RogueDLCAeonDiceSurfaceAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCAeonDiceSurface.json")
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
func (a *RogueDLCAeonDiceSurfaceAccessor) Raw() ([]RogueDLCAeonDiceSurface, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCAeonDiceSurface{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCAeonDiceSurfaceAccessor.LoadData to preload everything
func (a *RogueDLCAeonDiceSurfaceAccessor) GroupData() {
	for _, d := range a._data {
		a._dataAeonSurfaceDiceID[d.AeonSurfaceDiceID] = d
		a._dataDiceSurfaceIcon[d.DiceSurfaceIcon] = d
	}
}

// ByAeonSurfaceDiceID returns the RogueDLCAeonDiceSurface uniquely identified by AeonSurfaceDiceID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonDiceSurfaceAccessor) ByAeonSurfaceDiceID(identifier float64) (RogueDLCAeonDiceSurface, error) {
	if a._dataAeonSurfaceDiceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeonDiceSurface{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAeonSurfaceDiceID[identifier], nil
}

// ByDiceSurfaceIcon returns the RogueDLCAeonDiceSurface uniquely identified by DiceSurfaceIcon
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonDiceSurfaceAccessor) ByDiceSurfaceIcon(identifier string) (RogueDLCAeonDiceSurface, error) {
	if a._dataDiceSurfaceIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeonDiceSurface{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDiceSurfaceIcon[identifier], nil
}
