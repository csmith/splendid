package com.chameth.splendid.games.diceclimbing.ui.components

import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import com.chameth.splendid.shared.Die

@Composable
fun Die(die: Die, modifier: Modifier = Modifier) {
    Box(
        modifier = modifier
            .size(48.dp)
            .border(1.dp, Color.Black, RoundedCornerShape(4.dp))
    ) {
        Text(
            modifier = Modifier.align(Alignment.Center),
            text = die.value.toString(),
            style = MaterialTheme.typography.labelLarge
        )
    }
}