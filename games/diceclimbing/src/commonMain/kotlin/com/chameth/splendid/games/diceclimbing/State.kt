package com.chameth.splendid.games.diceclimbing

import com.chameth.splendid.shared.Die
import com.chameth.splendid.shared.engine.State

data class State(
    val phase: Phase = Phase.AssemblingPlayers,
    val lobbyMembers: Set<String> = emptySet(),
    val players: Map<String, Token> = emptyMap(),
    val currentPlayer: String? = null,
    val turnOrder: List<String> = emptyList(),
    val currentRoll: List<Die> = emptyList(),
    val board: List<Column> = listOf(
        Column(2, 3),
        Column(3, 5),
        Column(4, 7),
        Column(5, 9),
        Column(6, 11),
        Column(7, 13),
        Column(8, 11),
        Column(9, 9),
        Column(10, 7),
        Column(11, 5),
        Column(12, 3),
    ),
    val winner: String? = null
) : State {

    val blackTokenCount: Int
        get() = board.count { it.hasBlackToken }

    val rollPermutations: List<Pair<Int, Int>>
        get() = buildList {
            if (currentRoll.isNotEmpty()) {
                add((currentRoll[0].value + currentRoll[1].value) to (currentRoll[2].value + currentRoll[3].value))
                add((currentRoll[0].value + currentRoll[2].value) to (currentRoll[1].value + currentRoll[3].value))
                add((currentRoll[0].value + currentRoll[3].value) to (currentRoll[1].value + currentRoll[2].value))
            }
        }

    val currentPlayerToken: Token?
        get() = currentPlayer?.let { players[it] }

    fun columnForRoll(roll: Int) = board[roll - 2]

}