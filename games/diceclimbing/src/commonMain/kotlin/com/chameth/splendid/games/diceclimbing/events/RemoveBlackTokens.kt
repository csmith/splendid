package com.chameth.splendid.games.diceclimbing.events

import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.Token
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data object RemoveBlackTokens : Event<State> {

    override fun resolve(state: State) = state.copy(
        board = state.board.map { column ->
            column.withoutToken(Token.Black)
        }
    )
    
}
