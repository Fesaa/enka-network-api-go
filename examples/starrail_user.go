package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	enkanetworkapigo "github.com/Fesaa/enka-network-api-go"
	"github.com/Fesaa/enka-network-api-go/cache"
	"github.com/Fesaa/enka-network-api-go/starrail"
)

const SRUID = "714656501"

var _log *slog.Logger = slog.Default()

func main() {

	api, err := enkanetworkapigo.New("enka-network-api-go example starrail_user.go", cache.MEMORY)
	if err != nil {
		// Use proper error handling in a real program
		panic(err)
	}

	api.FetchHonkaiUser(SRUID,
		func(rhu *starrail.RawHonkaiUser) {
			user := rhu.ToUser()

			if user == nil {
				_log.Info("No user found")
				return
			}

			selectedCharacters := user.Characters

			if len(selectedCharacters) == 0 {
				_log.Info("No characters found")
				return
			}
			_log.Info("User: " + user.NickName)
			_log.Info(fmt.Sprintf("They are level %d, and have unlocked %d planets in Herta's Simulated Universe", user.Level, user.SimulatedUniverse))
			for _, character := range selectedCharacters {
				gameData := api.GetStarRailCharacterData(&character)

				if gameData == nil {
					_log.Info(fmt.Sprintf("No game data found for %d", character.AvatarId))
					return
				}

				_log.Info(fmt.Sprintf("%s is level %d, and has %d stars", gameData.Name(), character.Level, gameData.Star))
				_log.Info(fmt.Sprintf("Find the icon for the character by surfing to %s", api.GetStarRailIcon(gameData.AvatarSideIconPath)))
				_log.Info("\n")
			}

		},
		func(err error) {
			_log.Error(err.Error())
		})

	// Make the program not shut down before the goroutines are done
	_log.Info("Press Ctrl+C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}
