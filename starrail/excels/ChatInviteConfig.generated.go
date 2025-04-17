package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ChatInviteConfig struct {
	ChatNoticeType string    `json:"ChatNoticeType"`
	ExpireTime     float64   `json:"ExpireTime"`
	ID             float64   `json:"ID"`
	InviteContent  hash.Hash `json:"InviteContent"`
	InviteGo       hash.Hash `json:"InviteGo"`
	InviteInvalid  hash.Hash `json:"InviteInvalid"`
	InviteTitle    hash.Hash `json:"InviteTitle"`
	NoticeDesc     hash.Hash `json:"NoticeDesc"`
	NoticeTime     float64   `json:"NoticeTime"`
	PicPath        string    `json:"PicPath"`
	SendDesc       hash.Hash `json:"SendDesc"`
}
type ChatInviteConfigAccessor struct {
	_data        []ChatInviteConfig
	_dataID      map[float64]ChatInviteConfig
	_dataPicPath map[string]ChatInviteConfig
}

// LoadData retrieves the data. Must be called before ChatInviteConfig.GroupData
func (a *ChatInviteConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChatInviteConfig.json")
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
func (a *ChatInviteConfigAccessor) Raw() ([]ChatInviteConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChatInviteConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ChatInviteConfigAccessor.LoadData to preload everything
func (a *ChatInviteConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataPicPath[d.PicPath] = d
	}
}

// ByID returns the ChatInviteConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ChatInviteConfigAccessor) ByID(identifier float64) (ChatInviteConfig, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return ChatInviteConfig{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByPicPath returns the ChatInviteConfig uniquely identified by PicPath
//
// Error is only non-nil if the source errors out
func (a *ChatInviteConfigAccessor) ByPicPath(identifier string) (ChatInviteConfig, error) {
	if a._dataPicPath == nil {
		err := a.LoadData()
		if err != nil {
			return ChatInviteConfig{}, err
		}
		a.GroupData()
	}
	return a._dataPicPath[identifier], nil
}
