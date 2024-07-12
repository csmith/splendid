package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.events.BuildFoundationFromWaste
import com.chameth.splendid.games.klondike.events.SetPhase
import com.chameth.splendid.games.klondike.rules.canBuildFoundationWithCard
import com.chameth.splendid.games.klondike.rules.willWin
import com.chameth.splendid.shared.engine.Action
import kotlinx.serialization.Serializable

@Serializable
data class MoveWasteToFoundation(override val actor: String) : Action<State> {

    override fun resolve(state: State) = buildList {
        add(BuildFoundationFromWaste)

        if (state.willWin()) {
            add(SetPhase(phase = Phase.Finished))
        }
    }

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.WaitingForMove
                && state.waste.isNotEmpty()
                && state.canBuildFoundationWithCard(state.waste.flatten().last())
            ) {
                state.players.forEach {
                    add(MoveWasteToFoundation(it))
                }
            }
        }
    }
}