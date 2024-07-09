package com.chameth.splendid.games.klondike

import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import com.chameth.splendid.games.klondike.actions.KlondikeActions
import com.chameth.splendid.games.klondike.events.*
import com.chameth.splendid.games.klondike.ui.Board
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.engine.Game
import com.chameth.splendid.shared.engine.GameType
import kotlinx.serialization.modules.SerializersModule
import kotlinx.serialization.modules.polymorphic
import kotlinx.serialization.modules.subclass

data object Klondike : GameType<State> {
    @Suppress("UNCHECKED_CAST")
    override val uiRoot: @Composable (Game<*>, Modifier) -> Unit = { game, modifier ->
        Board(game = game as Game<State>, modifier = modifier)
    }

    override val stateFactory = { State() }

    override val actionsGenerator = KlondikeActions::generate

    override val version = 1

    override val name = "klondike"

    override val serializersModule: SerializersModule
        get() = SerializersModule {
            polymorphic(Event::class) {
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
        }
}