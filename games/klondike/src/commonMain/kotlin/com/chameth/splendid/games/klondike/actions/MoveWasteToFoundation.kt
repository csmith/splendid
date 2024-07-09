package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.events.BuildFoundationFromWaste
import com.chameth.splendid.games.klondike.events.SetPhase
import com.chameth.splendid.games.klondike.rules.canBuildFoundationWithCard
import com.chameth.splendid.games.klondike.rules.willWin

data object MoveWasteToFoundation : Action<State> {

    override fun resolve(state: State) = buildList {
        add(BuildFoundationFromWaste)

        if (state.willWin()) {
            add(SetPhase(phase = Phase.Finished))
        }
    }

    fun generate(state: State) = buildList {
        if (state.phase == Phase.WaitingForMove
            && state.waste.isNotEmpty()
            && state.canBuildFoundationWithCard(state.waste.flatten().last())
        ) {
            add(MoveWasteToFoundation)
        }
    }
}