package com.chameth.splendid.games.diceclimbing.ui.components

import androidx.compose.foundation.Canvas
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.unit.dp
import com.chameth.splendid.games.diceclimbing.Token

@Composable
fun TableauCell(
    tokens: List<Token>,
    modifier: Modifier = Modifier,
) {
    Box(modifier = modifier) {
        Canvas(
            modifier = Modifier
                .shadow(4.dp, shape = CircleShape)
                .fillMaxSize()
        ) {
            drawCircle(Colours.trackBackground)
        }

        Row(
            modifier = Modifier.align(Alignment.BottomEnd),
            horizontalArrangement = Arrangement.spacedBy(-(18.dp))
        ) {
            tokens.forEach {
                Token(type = it, modifier = Modifier.size(24.dp))
            }
        }
    }
}