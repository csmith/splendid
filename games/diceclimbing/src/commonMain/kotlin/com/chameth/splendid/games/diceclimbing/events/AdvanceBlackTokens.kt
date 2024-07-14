package com.chameth.splendid.games.diceclimbing.events

import com.chameth.splendid.games.diceclimbing.Phase
import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.Token
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable
import kotlin.math.min

@Serializable
data class AdvanceBlackTokens(
    val token: Token,
    val rolls: List<Int>
) : Event<State> {

    override fun resolve(state: State) = state.copy(
        board = state.board.map { column ->
            val count = rolls.count { it == column.rollRequired }
            if (count > 0) {
                if (Token.Black in column.tokens) {
                    column.copy(tokens = column.tokens + (Token.Black to min((column.tokens[Token.Black] ?: 0) + count, column.height-1)))
                } else {
                    column.copy(tokens = column.tokens + (Token.Black to min((column.tokens[token] ?: -1) + count, column.height-1)))
                }
            } else {
                column
            }
        },
        currentRoll = emptyList(),
        phase = Phase.WaitingForDecision
    )

}
