package events

import (
	coremodel "github.com/csmith/splendid/backend/model"
)

func specifier(name string) coremodel.Specifier {
	return coremodel.Specifier{
		Domain:   "dyhat",
		Category: "e",
		Name:     name,
	}
}

func init() {
	coremodel.RegisterType(&CardDealtEvent{})
	coremodel.RegisterType(&CardDiscardedEvent{})
	coremodel.RegisterType(&CardRemovedEvent{})
	coremodel.RegisterType(&CardsSetAsideEvent{})
	coremodel.RegisterType(&CurrentPlayerUpdatedEvent{})
	coremodel.RegisterType(&DeckUpdatedEvent{})
	coremodel.RegisterType(&HandRevealedEvent{})
	coremodel.RegisterType(&HandsSwappedEvent{})
	coremodel.RegisterType(&LastRoundWinnersUpdatedEvent{})
	coremodel.RegisterType(&PhaseUpdatedEvent{})
	coremodel.RegisterType(&PlayerAddedEvent{})
	coremodel.RegisterType(&PlayerDiscardPileClearedEvent{})
	coremodel.RegisterType(&PlayerEliminatedEvent{})
	coremodel.RegisterType(&PlayerHandClearedEvent{})
	coremodel.RegisterType(&PlayerProtectionClearedEvent{})
	coremodel.RegisterType(&PlayerProtectionGrantedEvent{})
	coremodel.RegisterType(&PlayerRestoredEvent{})
	coremodel.RegisterType(&PlayerTokenAwardedEvent{})
	coremodel.RegisterType(&RemovedCardDealtEvent{})
	coremodel.RegisterType(&RoundUpdatedEvent{})
	coremodel.RegisterType(&WinnerDeclaredEvent{})
}
