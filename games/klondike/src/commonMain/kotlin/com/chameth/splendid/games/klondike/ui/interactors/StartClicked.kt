package com.chameth.splendid.games.klondike.ui.interactors

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.Variant
import com.chameth.splendid.games.klondike.actions.StartGame
import com.chameth.splendid.games.klondike.ui.model.Selection
import com.chameth.splendid.shared.engine.Action

fun startClicked(
    state: State,
    actor: String,
    selection: Selection?,
    variant: Variant,
    invoke: (Action<State>) -> Unit
): Selection? {
    invoke(StartGame(actor, variant))
    return null
}