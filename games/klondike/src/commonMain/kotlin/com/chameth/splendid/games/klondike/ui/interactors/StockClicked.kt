package com.chameth.splendid.games.klondike.ui.interactors

import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.actions.DrawFromStock
import com.chameth.splendid.games.klondike.actions.ResetStock
import com.chameth.splendid.games.klondike.ui.model.Selection

fun stockClicked(
    state: State,
    selection: Selection?,
    invoke: (Action<State>) -> Unit
): Selection? {
    when {
        state.stock.isNotEmpty() -> invoke(DrawFromStock)
        state.waste.isNotEmpty() -> invoke(ResetStock)
    }
    return null
}