package actions

import (
	coremodel "github.com/csmith/splendid/backend/model"
	"github.com/csmith/splendid/backend/serialization"
)

func specifier(name string) coremodel.Specifier {
	return coremodel.Specifier{
		Domain:   "dyhat",
		Category: "a",
		Name:     name,
	}
}

func init() {
	serialization.RegisterType(&AddPlayerAction{})
	serialization.RegisterType(&DrawCardAction{})
	serialization.RegisterType(&StartGameAction{})
	serialization.RegisterType(&PlayCardWithGuessAction{})
	serialization.RegisterType(&PlayCardTargetOthersAction{})
	serialization.RegisterType(&PlayCardTargetAnyAction{})
	serialization.RegisterType(&PlayCardNoTargetAction{})
	serialization.RegisterType(&DiscardCardAction{})
}
