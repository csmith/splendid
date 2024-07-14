package com.chameth.splendid.games.diceclimbing.ui.components

import androidx.compose.foundation.Canvas
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import com.chameth.splendid.games.diceclimbing.Token

@Composable
fun Token(type: Token, modifier: Modifier = Modifier) {
    Canvas(modifier = modifier.shadow(4.dp, shape = CircleShape)) {
        drawCircle(when (type) {
            Token.Red -> Color.Red
            Token.Green -> Color.Green
            Token.Blue -> Color.Blue
            Token.Yellow -> Color.Yellow
            Token.Black -> Color.Black
        })
    }
}