package com.chameth.splendid.games.klondike

import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import com.chameth.splendid.games.klondike.actions.*
import com.chameth.splendid.games.klondike.events.*
import com.chameth.splendid.games.klondike.ui.Board
import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.engine.GameType
import com.chameth.splendid.shared.playingcards.Card
import com.chameth.splendid.shared.playingcards.Rank
import com.chameth.splendid.shared.playingcards.Suit
import kotlinx.serialization.modules.SerializersModule
import kotlinx.serialization.modules.polymorphic
import kotlinx.serialization.modules.subclass

data object Klondike : GameType<State> {

    override val uiRoot: @Composable (State, (Action<State>) -> Unit, Modifier) -> Unit = { state, action, modifier ->
        Board(gameState = state, action = action, modifier = modifier)
    }

    override val stateFactory = { State() }

    override val actionsGenerator = { state: State ->
        buildList {
            addAll(AutoSolve.generate(state))
            addAll(DrawFromStock.generate(state))
            addAll(MoveFoundationToTableau.generate(state))
            addAll(MoveTableauToFoundation.generate(state))
            addAll(MoveTableauToTableau.generate(state))
            addAll(MoveWasteToFoundation.generate(state))
            addAll(MoveWasteToTableau.generate(state))
            addAll(NewGame.generate(state))
            addAll(ResetStock.generate(state))
            addAll(StartGame.generate(state))
        }
    }

    override val version = 1

    override val name = "Klondike"

    override val serializersModule = SerializersModule {
        polymorphic(Event::class) {
            subclass(AddPlayer::class)
            subclass(BuildFoundationFromTableau::class)
            subclass(BuildFoundationFromWaste::class)
            subclass(BuildTableauFromFoundation::class)
            subclass(BuildTableauFromWaste::class)
            subclass(DealToTableau::class)
            subclass(DealToWaste::class)
            subclass(MoveCardsWithinTableau::class)
            subclass(MoveWasteToStock::class)
            subclass(ResetState::class)
            subclass(RevealCardInTableau::class)
            subclass(SetPhase::class)
            subclass(SetVariant::class)
        }

        polymorphic(Action::class) {
            subclass(AutoSolve::class)
            subclass(DrawFromStock::class)
            subclass(MoveFoundationToTableau::class)
            subclass(MoveTableauToFoundation::class)
            subclass(MoveTableauToTableau::class)
            subclass(MoveWasteToFoundation::class)
            subclass(MoveWasteToTableau::class)
            subclass(NewGame::class)
            subclass(ResetStock::class)
            subclass(StartGame::class)
        }
    }

    override fun mask(state: State, event: Event<*>, actor: String): Event<*> {
        return when (event) {
            is ResetState -> event.copy(
                stock = event.stock.map { Card(suit = Suit.Hearts, rank = Rank.Ace, visible = false) }
            )

            is DealToTableau -> event.copy(
                card = if (event.card.visible) event.card else Card(
                    suit = Suit.Hearts,
                    rank = Rank.Ace,
                    visible = false
                )
            )

            else -> event
        }
    }

    override fun newAddPlayerEvent(playerId: String): Event<*> {
        return AddPlayer(playerId)
    }
}