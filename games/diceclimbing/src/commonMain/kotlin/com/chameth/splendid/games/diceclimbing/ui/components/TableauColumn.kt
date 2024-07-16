package com.chameth.splendid.games.diceclimbing.ui.components

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import com.chameth.splendid.games.diceclimbing.Column

@Composable
fun TableauColumn(
    column: Column,
    modifier: Modifier = Modifier,
) {
    Column(
        modifier = modifier
            .background(Brush.horizontalGradient(
                0f to Color.Transparent,
                0.475f to Color.Transparent,
                0.475f to Colours.trackBackground,
                0.525f to Colours.trackBackground,
                0.525f to Color.Transparent,
                1f to Color.Transparent
            )),
        horizontalAlignment = Alignment.CenterHorizontally,
        verticalArrangement = Arrangement.spacedBy(8.dp)
    ) {
        for (i in 1..column.height) {
            if (column.completed) {
                Token(type = column.owner!!, modifier = Modifier.size(36.dp))
            } else {
                val index = column.height - i
                val tokens = column.tokens.filter { it.value == index }.map { it.key }.toList()
                TableauCell(tokens = tokens, modifier = Modifier.size(36.dp))
            }
        }

        Box(
            modifier = Modifier
                .size(24.dp)
                .background(Colours.trackBackground, shape = CircleShape),
        ) {
            Text(
                modifier = Modifier.align(Alignment.Center),
                text = column.rollRequired.toString(),
                style = MaterialTheme.typography.labelSmall,
            )
        }
    }
}