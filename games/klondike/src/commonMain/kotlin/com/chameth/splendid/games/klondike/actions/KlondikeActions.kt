package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.games.klondike.State

object KlondikeActions {

    fun generate(state: State) = buildList {
        addAll(StartGame.generate(state))
        addAll(MoveTableauToFoundation.generate(state))
        addAll(MoveWasteToFoundation.generate(state))
        addAll(DrawFromStock.generate(state))
        addAll(ResetStock.generate(state))
        addAll(MoveWasteToTableau.generate(state))
        addAll(MoveTableauToTableau.generate(state))
        addAll(MoveFoundationToTableau.generate(state))
        addAll(NewGame.generate(state))
        addAll(AutoSolve.generate(state))
    }

}