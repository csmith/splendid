package com.chameth.splendid.games.klondike.ui.interactors

import com.chameth.splendid.shared.playingcards.Card
import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.actions.MoveFoundationToTableau
import com.chameth.splendid.games.klondike.actions.MoveTableauToFoundation
import com.chameth.splendid.games.klondike.actions.MoveTableauToTableau
import com.chameth.splendid.games.klondike.actions.MoveWasteToTableau
import com.chameth.splendid.games.klondike.ui.model.Selection

fun tableauClicked(
    state: State,
    selection: Selection?,
    tableau: Int,
    card: Card?,
    invoke: (Action<State>) -> Unit
): Selection? {
    when {
        selection == null && card?.visible == true ->
            return Selection(
                source = Selection.SelectionSource.Tableau,
                card = card,
                tableau = tableau,
            )

        card != null && selection?.card == card -> {
            invoke(MoveTableauToFoundation(tableau))
            return null
        }

        selection?.source == Selection.SelectionSource.Waste -> {
            invoke(MoveWasteToTableau(tableau))
            return null
        }

        selection?.source == Selection.SelectionSource.Foundation -> {
            invoke(MoveFoundationToTableau(foundation = selection.card.suit.ordinal, tableau = tableau))
            return null
        }

        selection?.source == Selection.SelectionSource.Tableau -> {
            val from = state.tableau[selection.tableau]
            val count = from.size - from.indexOf(selection.card)
            invoke(MoveTableauToTableau(from = selection.tableau, to = tableau, count = count))
            return null
        }

        else ->
            return null
    }
}