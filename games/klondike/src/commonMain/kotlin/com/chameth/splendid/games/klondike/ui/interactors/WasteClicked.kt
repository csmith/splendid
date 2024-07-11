package com.chameth.splendid.games.klondike.ui.interactors

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.actions.MoveWasteToFoundation
import com.chameth.splendid.games.klondike.ui.model.Selection
import com.chameth.splendid.shared.engine.Action

fun wasteClicked(
    state: State,
    actor: String,
    selection: Selection?,
    invoke: (Action<State>) -> Unit
): Selection? {
    when {
        selection?.source == Selection.SelectionSource.Waste -> {
            invoke(MoveWasteToFoundation(actor))
            return null
        }

        state.waste.isNotEmpty() ->
            return Selection(
                source = Selection.SelectionSource.Waste,
                card = state.waste.flatten().last(),
            )

        else ->
            return null
    }
}