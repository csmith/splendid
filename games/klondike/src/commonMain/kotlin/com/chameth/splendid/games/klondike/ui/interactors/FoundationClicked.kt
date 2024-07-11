package com.chameth.splendid.games.klondike.ui.interactors

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.actions.MoveTableauToFoundation
import com.chameth.splendid.games.klondike.actions.MoveWasteToFoundation
import com.chameth.splendid.games.klondike.ui.model.Selection
import com.chameth.splendid.shared.engine.Action

fun foundationClicked(
    state: State,
    actor: String,
    selection: Selection?,
    foundation: Int,
    invoke: (Action<State>) -> Unit
): Selection? {
    when {
        selection == null && state.foundations[foundation].isNotEmpty() ->
            return Selection(
                source = Selection.SelectionSource.Foundation,
                card = state.foundations[foundation].last(),
            )

        selection?.source == Selection.SelectionSource.Waste ->  {
            invoke(MoveWasteToFoundation(actor))
            return null
        }

        selection?.source == Selection.SelectionSource.Tableau -> {
            invoke(MoveTableauToFoundation(actor, selection.tableau))
            return null
        }

        else ->
            return null
    }
}