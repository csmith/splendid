package com.chameth.splendid.games.diceclimbing

import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import com.chameth.splendid.games.diceclimbing.actions.*
import com.chameth.splendid.games.diceclimbing.events.*
import com.chameth.splendid.games.diceclimbing.ui.Board
import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.engine.GameType
import kotlinx.serialization.modules.SerializersModule
import kotlinx.serialization.modules.polymorphic
import kotlinx.serialization.modules.subclass

object DiceClimbing : GameType<State> {

    override val stateFactory = { State() }

    override val actionsGenerator = { state: State ->
        buildList {
            addAll(AdvanceColumns.generate(state))
            addAll(GoBust.generate(state))
            addAll(RollDice.generate(state))
            addAll(SitDown.generate(state))
            addAll(StartGame.generate(state))
            addAll(Stop.generate(state))
        }
    }

    override val uiRoot: @Composable (State, (Action<State>) -> Unit, Modifier) -> Unit
        get() = { state, actionHandler, modifier ->
            Board(state, actionHandler, modifier)
        }

    override val version: Int = 1

    override val serializersModule = SerializersModule {
        polymorphic(Event::class) {
            subclass(AddPlayer::class)
            subclass(AdvanceBlackTokens::class)
            subclass(GameOver::class)
            subclass(RemoveBlackTokens::class)
            subclass(SaveBlackTokens::class)
            subclass(SetDice::class)
            subclass(SetTurn::class)
            subclass(SetTurnOrder::class)
            subclass(SitPlayer::class)
        }

        polymorphic(Action::class) {
            subclass(AdvanceColumns::class)
            subclass(GoBust::class)
            subclass(RollDice::class)
            subclass(SitDown::class)
            subclass(StartGame::class)
            subclass(Stop::class)
        }
    }

    override val name = "Dice Climbing"

    override fun newAddPlayerEvent(playerId: String) = AddPlayer(playerId)

    override fun mask(state: State, event: Event<*>, actor: String) = event

}