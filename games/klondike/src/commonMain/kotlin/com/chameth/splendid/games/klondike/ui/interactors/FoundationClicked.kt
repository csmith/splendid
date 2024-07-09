package com.chameth.splendid.games.klondike.ui.interactors

import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.actions.MoveTableauToFoundation
import com.chameth.splendid.games.klondike.actions.MoveWasteToFoundation
import com.chameth.splendid.games.klondike.ui.model.Selection

fun foundationClicked(
    state: State,
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
            invoke(MoveWasteToFoundation)
            return null
        }

        selection?.source == Selection.SelectionSource.Tableau -> {
            invoke(MoveTableauToFoundation(selection.tableau))
            return null
        }

        else ->
            return null
    }
}