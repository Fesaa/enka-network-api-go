package localization

import (
	"log/slog"
	"testing"
)

func TestLocalization(t *testing.T) {
	Init(slog.Default())
	l := Get()
	if len(l.genshinLocalizationCache) == 0 {
		t.Errorf("Genshin localization cache is empty")
	}
	if len(l.honkaiLocalizationCache) == 0 {
		t.Errorf("Honkai localization cache is empty")
	}
}
