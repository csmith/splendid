package com.chameth.splendid.shared.util

import org.jetbrains.compose.resources.getStringArray
import splendid.shared.generated.resources.Res
import splendid.shared.generated.resources.adjectives
import splendid.shared.generated.resources.animals
import splendid.shared.generated.resources.colours

suspend fun generateAca() = buildString {
    append(getStringArray(Res.array.adjectives).random())
    append('-')
    append(getStringArray(Res.array.colours).random())
    append('-')
    append(getStringArray(Res.array.animals).random())
}