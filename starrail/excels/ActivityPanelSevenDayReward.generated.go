package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityPanelSevenDayReward struct {
	BGImgPath          string  `json:"BGImgPath"`
	CardWidget0ImgPath string  `json:"CardWidget0ImgPath"`
	CardWidget1ImgPath string  `json:"CardWidget1ImgPath"`
	CardWidget2ImgPath string  `json:"CardWidget2ImgPath"`
	ID                 float64 `json:"ID"`
	PicImgPath         string  `json:"PicImgPath"`
}
type ActivityPanelSevenDayRewardAccessor struct {
	_data                   []ActivityPanelSevenDayReward
	_dataBGImgPath          map[string]ActivityPanelSevenDayReward
	_dataCardWidget0ImgPath map[string]ActivityPanelSevenDayReward
	_dataCardWidget1ImgPath map[string]ActivityPanelSevenDayReward
	_dataCardWidget2ImgPath map[string]ActivityPanelSevenDayReward
	_dataID                 map[float64]ActivityPanelSevenDayReward
	_dataPicImgPath         map[string]ActivityPanelSevenDayReward
}

// LoadData retrieves the data. Must be called before ActivityPanelSevenDayReward.GroupData
func (a *ActivityPanelSevenDayRewardAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityPanelSevenDayReward.json")
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
func (a *ActivityPanelSevenDayRewardAccessor) Raw() ([]ActivityPanelSevenDayReward, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityPanelSevenDayReward{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityPanelSevenDayRewardAccessor.LoadData to preload everything
func (a *ActivityPanelSevenDayRewardAccessor) GroupData() {
	a._dataBGImgPath = map[string]ActivityPanelSevenDayReward{}
	a._dataCardWidget0ImgPath = map[string]ActivityPanelSevenDayReward{}
	a._dataCardWidget1ImgPath = map[string]ActivityPanelSevenDayReward{}
	a._dataCardWidget2ImgPath = map[string]ActivityPanelSevenDayReward{}
	a._dataID = map[float64]ActivityPanelSevenDayReward{}
	a._dataPicImgPath = map[string]ActivityPanelSevenDayReward{}
	for _, d := range a._data {
		a._dataBGImgPath[d.BGImgPath] = d
		a._dataCardWidget0ImgPath[d.CardWidget0ImgPath] = d
		a._dataCardWidget1ImgPath[d.CardWidget1ImgPath] = d
		a._dataCardWidget2ImgPath[d.CardWidget2ImgPath] = d
		a._dataID[d.ID] = d
		a._dataPicImgPath[d.PicImgPath] = d
	}
}

// ByBGImgPath returns the ActivityPanelSevenDayReward uniquely identified by BGImgPath
//
// Error is only non-nil if the source errors out
func (a *ActivityPanelSevenDayRewardAccessor) ByBGImgPath(identifier string) (ActivityPanelSevenDayReward, error) {
	if a._dataBGImgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityPanelSevenDayReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBGImgPath[identifier], nil
}

// ByCardWidget0ImgPath returns the ActivityPanelSevenDayReward uniquely identified by CardWidget0ImgPath
//
// Error is only non-nil if the source errors out
func (a *ActivityPanelSevenDayRewardAccessor) ByCardWidget0ImgPath(identifier string) (ActivityPanelSevenDayReward, error) {
	if a._dataCardWidget0ImgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityPanelSevenDayReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCardWidget0ImgPath[identifier], nil
}

// ByCardWidget1ImgPath returns the ActivityPanelSevenDayReward uniquely identified by CardWidget1ImgPath
//
// Error is only non-nil if the source errors out
func (a *ActivityPanelSevenDayRewardAccessor) ByCardWidget1ImgPath(identifier string) (ActivityPanelSevenDayReward, error) {
	if a._dataCardWidget1ImgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityPanelSevenDayReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCardWidget1ImgPath[identifier], nil
}

// ByCardWidget2ImgPath returns the ActivityPanelSevenDayReward uniquely identified by CardWidget2ImgPath
//
// Error is only non-nil if the source errors out
func (a *ActivityPanelSevenDayRewardAccessor) ByCardWidget2ImgPath(identifier string) (ActivityPanelSevenDayReward, error) {
	if a._dataCardWidget2ImgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityPanelSevenDayReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCardWidget2ImgPath[identifier], nil
}

// ByID returns the ActivityPanelSevenDayReward uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ActivityPanelSevenDayRewardAccessor) ByID(identifier float64) (ActivityPanelSevenDayReward, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityPanelSevenDayReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByPicImgPath returns the ActivityPanelSevenDayReward uniquely identified by PicImgPath
//
// Error is only non-nil if the source errors out
func (a *ActivityPanelSevenDayRewardAccessor) ByPicImgPath(identifier string) (ActivityPanelSevenDayReward, error) {
	if a._dataPicImgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityPanelSevenDayReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPicImgPath[identifier], nil
}
