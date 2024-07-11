package com.chameth.splendid.games.klondike.ui.interactors

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.actions.AutoSolve
import com.chameth.splendid.games.klondike.ui.model.Selection
import com.chameth.splendid.shared.engine.Action

fun autoSolveClicked(
    state: State,
    actor: String,
    selection: Selection?,
    invoke: (Action<State>) -> Unit
): Selection? {
    invoke(AutoSolve(actor))
    return null
}