package com.chameth.splendid.games.diceclimbing

data class Column(
    val rollRequired: Int,
    val height: Int,
    val tokens: Map<Token, Int> = emptyMap()
) {
    val completed: Boolean
        get() = tokens.any { (type, pos) -> type != Token.Black && height == pos+1 }

    val owner: Token?
        get() = tokens.keys.firstOrNull()

    val hasBlackToken: Boolean
        get() = Token.Black in tokens

    fun withoutToken(token: Token) =
        copy(tokens = tokens.filterNot { it.key == token })
}
