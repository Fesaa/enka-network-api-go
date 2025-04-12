package localization

import (
	"github.com/rs/zerolog"
	"testing"
)

func TestLocalization(t *testing.T) {
	Init(zerolog.Nop())
	l := Get()
	if len(l.genshinCache) == 0 {
		t.Errorf("Genshin localization cache is empty")
	}
	if len(l.hsrCache) == 0 {
		t.Errorf("Honkai localization cache is empty")
	}
}
