package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.events.MoveWasteToStock
import com.chameth.splendid.shared.engine.Action
import kotlinx.serialization.Serializable

@Serializable
data class ResetStock(override val actor: String) : Action<State> {

    override fun resolve(state: State) = listOf(MoveWasteToStock)

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.WaitingForMove && state.stock.isEmpty() && state.waste.isNotEmpty()) {
                state.players.forEach {
                    add(ResetStock(it))
                }
            }
        }
    }

}