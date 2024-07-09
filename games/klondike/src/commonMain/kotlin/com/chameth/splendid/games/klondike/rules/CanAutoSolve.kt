package com.chameth.splendid.games.klondike.rules

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.playingcards.Card

fun State.canAutoSolve() =
    phase == Phase.WaitingForMove &&
            waste.isEmpty() &&
            stock.isEmpty() &&
            tableau.all { it.all(Card::visible) }