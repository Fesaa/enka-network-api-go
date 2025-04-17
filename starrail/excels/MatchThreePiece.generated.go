package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MatchThreePiece struct {
	ImagePath      string  `json:"ImagePath"`
	PieceID        float64 `json:"PieceID"`
	RowBombPath    string  `json:"RowBombPath"`
	SquareBombPath string  `json:"SquareBombPath"`
}
type MatchThreePieceAccessor struct {
	_data          []MatchThreePiece
	_dataImagePath map[string]MatchThreePiece
	_dataPieceID   map[float64]MatchThreePiece
}

// LoadData retrieves the data. Must be called before MatchThreePiece.GroupData
func (a *MatchThreePieceAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MatchThreePiece.json")
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
func (a *MatchThreePieceAccessor) Raw() ([]MatchThreePiece, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MatchThreePiece{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MatchThreePieceAccessor.LoadData to preload everything
func (a *MatchThreePieceAccessor) GroupData() {
	for _, d := range a._data {
		a._dataImagePath[d.ImagePath] = d
		a._dataPieceID[d.PieceID] = d
	}
}

// ByImagePath returns the MatchThreePiece uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *MatchThreePieceAccessor) ByImagePath(identifier string) (MatchThreePiece, error) {
	if a._dataImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreePiece{}, err
			}
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}

// ByPieceID returns the MatchThreePiece uniquely identified by PieceID
//
// Error is only non-nil if the source errors out
func (a *MatchThreePieceAccessor) ByPieceID(identifier float64) (MatchThreePiece, error) {
	if a._dataPieceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreePiece{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPieceID[identifier], nil
}
