package com.chameth.splendid.games.klondike.ui.util

import androidx.compose.foundation.border
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import com.chameth.splendid.games.klondike.ui.SelectableCard

fun Modifier.applySelectableState(card: SelectableCard?): Modifier {
    return if (card?.selected == true)
        this.border(2.dp, Color.Red, RoundedCornerShape(8.dp))
    else
        this
}