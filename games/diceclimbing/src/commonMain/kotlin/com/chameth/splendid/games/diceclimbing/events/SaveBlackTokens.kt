package com.chameth.splendid.games.diceclimbing.events

import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.Token
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data class SaveBlackTokens(val token: Token) : Event<State> {

    override fun resolve(state: State) = state.copy(
        board = state.board.map { column ->
            if (Token.Black in column.tokens) {
                column.copy(tokens = column.tokens + (token to (column.tokens[Token.Black] ?: 0)))
            } else {
                column
            }
        }
    )

}
