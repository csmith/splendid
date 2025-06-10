package events

import "github.com/csmith/splendid/backend/serialization"

func specifier(name string) serialization.Specifier {
	return serialization.Specifier{
		Domain:   "dyhat",
		Category: "e",
		Name:     name,
	}
}

func init() {
	serialization.RegisterType(&CardDealtEvent{})
	serialization.RegisterType(&CardDiscardedEvent{})
	serialization.RegisterType(&CardRemovedEvent{})
	serialization.RegisterType(&CardsSetAsideEvent{})
	serialization.RegisterType(&CurrentPlayerUpdatedEvent{})
	serialization.RegisterType(&DeckUpdatedEvent{})
	serialization.RegisterType(&HandRevealedEvent{})
	serialization.RegisterType(&HandsSwappedEvent{})
	serialization.RegisterType(&LastRoundWinnersUpdatedEvent{})
	serialization.RegisterType(&PhaseUpdatedEvent{})
	serialization.RegisterType(&PlayerActionCompletedEvent{})
	serialization.RegisterType(&PlayerActionStartedEvent{})
	serialization.RegisterType(&PlayerActionUpdatedEvent{})
	serialization.RegisterType(&PlayerAddedEvent{})
	serialization.RegisterType(&PlayerDiscardPileClearedEvent{})
	serialization.RegisterType(&PlayerEliminatedEvent{})
	serialization.RegisterType(&PlayerHandClearedEvent{})
	serialization.RegisterType(&PlayerProtectionClearedEvent{})
	serialization.RegisterType(&PlayerProtectionGrantedEvent{})
	serialization.RegisterType(&PlayerRestoredEvent{})
	serialization.RegisterType(&PlayerTokenAwardedEvent{})
	serialization.RegisterType(&RemovedCardDealtEvent{})
	serialization.RegisterType(&RoundUpdatedEvent{})
	serialization.RegisterType(&WinnerDeclaredEvent{})
}
