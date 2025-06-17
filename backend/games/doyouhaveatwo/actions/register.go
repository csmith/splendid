package actions

import (
	coremodel "github.com/csmith/splendid/backend/model"
)

func specifier(name string) coremodel.Specifier {
	return coremodel.Specifier{
		Domain:   "dyhat",
		Category: "a",
		Name:     name,
	}
}

func init() {
	coremodel.RegisterType(&AddPlayerAction{})
	coremodel.RegisterType(&DrawCardAction{})
	coremodel.RegisterType(&StartGameAction{})
	coremodel.RegisterType(&PlayCardWithGuessAction{})
	coremodel.RegisterType(&PlayCardTargetOthersAction{})
	coremodel.RegisterType(&PlayCardTargetAnyAction{})
	coremodel.RegisterType(&PlayCardNoTargetAction{})
	coremodel.RegisterType(&DiscardCardAction{})
}
