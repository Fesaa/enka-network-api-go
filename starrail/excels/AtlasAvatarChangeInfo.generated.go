package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type AtlasAvatarChangeInfo struct {
	CCOFCKBMMMI float64 `json:"CCOFCKBMMMI"`
	DJPCAIKIONP float64 `json:"DJPCAIKIONP"`
	EBBHGBKEPAA float64 `json:"EBBHGBKEPAA"`
	OELNFIJLCOL float64 `json:"OELNFIJLCOL"`
}
type AtlasAvatarChangeInfoAccessor struct {
	_data            []AtlasAvatarChangeInfo
	_dataOELNFIJLCOL map[float64]AtlasAvatarChangeInfo
	_dataCCOFCKBMMMI map[float64]AtlasAvatarChangeInfo
	_dataDJPCAIKIONP map[float64]AtlasAvatarChangeInfo
	_dataEBBHGBKEPAA map[float64]AtlasAvatarChangeInfo
}

// LoadData retrieves the data. Must be called before AtlasAvatarChangeInfo.GroupData
func (a *AtlasAvatarChangeInfoAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AtlasAvatarChangeInfo.json")
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
func (a *AtlasAvatarChangeInfoAccessor) Raw() ([]AtlasAvatarChangeInfo, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AtlasAvatarChangeInfo{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AtlasAvatarChangeInfoAccessor.LoadData to preload everything
func (a *AtlasAvatarChangeInfoAccessor) GroupData() {
	for _, d := range a._data {
		a._dataOELNFIJLCOL[d.OELNFIJLCOL] = d
		a._dataCCOFCKBMMMI[d.CCOFCKBMMMI] = d
		a._dataDJPCAIKIONP[d.DJPCAIKIONP] = d
		a._dataEBBHGBKEPAA[d.EBBHGBKEPAA] = d
	}
}

// ByOELNFIJLCOL returns the AtlasAvatarChangeInfo uniquely identified by OELNFIJLCOL
//
// Error is only non-nil if the source errors out
func (a *AtlasAvatarChangeInfoAccessor) ByOELNFIJLCOL(identifier float64) (AtlasAvatarChangeInfo, error) {
	if a._dataOELNFIJLCOL == nil {
		err := a.LoadData()
		if err != nil {
			return AtlasAvatarChangeInfo{}, err
		}
		a.GroupData()
	}
	return a._dataOELNFIJLCOL[identifier], nil
}

// ByCCOFCKBMMMI returns the AtlasAvatarChangeInfo uniquely identified by CCOFCKBMMMI
//
// Error is only non-nil if the source errors out
func (a *AtlasAvatarChangeInfoAccessor) ByCCOFCKBMMMI(identifier float64) (AtlasAvatarChangeInfo, error) {
	if a._dataCCOFCKBMMMI == nil {
		err := a.LoadData()
		if err != nil {
			return AtlasAvatarChangeInfo{}, err
		}
		a.GroupData()
	}
	return a._dataCCOFCKBMMMI[identifier], nil
}

// ByDJPCAIKIONP returns the AtlasAvatarChangeInfo uniquely identified by DJPCAIKIONP
//
// Error is only non-nil if the source errors out
func (a *AtlasAvatarChangeInfoAccessor) ByDJPCAIKIONP(identifier float64) (AtlasAvatarChangeInfo, error) {
	if a._dataDJPCAIKIONP == nil {
		err := a.LoadData()
		if err != nil {
			return AtlasAvatarChangeInfo{}, err
		}
		a.GroupData()
	}
	return a._dataDJPCAIKIONP[identifier], nil
}

// ByEBBHGBKEPAA returns the AtlasAvatarChangeInfo uniquely identified by EBBHGBKEPAA
//
// Error is only non-nil if the source errors out
func (a *AtlasAvatarChangeInfoAccessor) ByEBBHGBKEPAA(identifier float64) (AtlasAvatarChangeInfo, error) {
	if a._dataEBBHGBKEPAA == nil {
		err := a.LoadData()
		if err != nil {
			return AtlasAvatarChangeInfo{}, err
		}
		a.GroupData()
	}
	return a._dataEBBHGBKEPAA[identifier], nil
}
