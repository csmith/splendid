package com.chameth.splendid.games.klondike.ui.interactors

import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.Variant
import com.chameth.splendid.games.klondike.actions.StartGame
import com.chameth.splendid.games.klondike.ui.model.Selection

fun startClicked(
    state: State,
    selection: Selection?,
    variant: Variant,
    invoke: (Action<State>) -> Unit
): Selection? {
    invoke(StartGame(variant))
    return null
}