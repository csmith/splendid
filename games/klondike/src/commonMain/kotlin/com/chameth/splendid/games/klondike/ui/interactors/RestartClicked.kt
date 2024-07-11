package com.chameth.splendid.games.klondike.ui.interactors

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.actions.NewGame
import com.chameth.splendid.games.klondike.ui.model.Selection
import com.chameth.splendid.shared.engine.Action

fun restartClicked(
    state: State,
    actor: String,
    selection: Selection?,
    invoke: (Action<State>) -> Unit
): Selection? {
    invoke(NewGame(actor))
    return null
}