package com.chameth.splendid.games.diceclimbing.ui.components

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.horizontalScroll
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.aspectRatio
import androidx.compose.foundation.rememberScrollState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.key
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import com.chameth.splendid.games.diceclimbing.Column

@Composable
fun Tableau(
    columns: List<Column>,
    modifier: Modifier = Modifier,
) {
    Row(
        modifier = modifier
            .horizontalScroll(rememberScrollState())
            .aspectRatio(1f)
            .background(
                brush = Brush.radialGradient(listOf(Color(0x99, 0xA3, 0xD4), Color(0x89, 0x93, 0xC4))),
                shape = DiamondShape()
            )
            .border(
                width = 2.dp,
                color = Color(0x16, 0x0C, 0x40),
                shape = DiamondShape()
            ),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(16.dp, Alignment.CenterHorizontally),
    ) {
        columns.forEach { column ->
            key("column-${column.rollRequired}") {
                TableauColumn(column)
            }
        }
    }
}