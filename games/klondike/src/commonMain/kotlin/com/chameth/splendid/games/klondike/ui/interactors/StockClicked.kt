package com.chameth.splendid.games.klondike.ui.interactors

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.actions.DrawFromStock
import com.chameth.splendid.games.klondike.actions.ResetStock
import com.chameth.splendid.games.klondike.ui.model.Selection
import com.chameth.splendid.shared.engine.Action

fun stockClicked(
    state: State,
    actor: String,
    selection: Selection?,
    invoke: (Action<State>) -> Unit
): Selection? {
    when {
        state.stock.isNotEmpty() -> invoke(DrawFromStock(actor))
        state.waste.isNotEmpty() -> invoke(ResetStock(actor))
    }
    return null
}