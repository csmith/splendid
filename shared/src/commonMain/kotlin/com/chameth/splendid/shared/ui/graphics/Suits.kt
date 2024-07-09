package com.chameth.splendid.shared.ui.graphics

import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.SolidColor
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.graphics.vector.path
import androidx.compose.ui.unit.dp

object Suits {
    val hearts = ImageVector.Builder(
        name = "hearts",
        defaultWidth = 16.dp,
        defaultHeight = 16.dp,
        viewportWidth = 64f,
        viewportHeight = 64f
    ).apply {
        path(fill = SolidColor(Color.Red)) {
            moveTo(32f, 60f)
            curveTo(22f, 40f, 10f, 30f, 10f, 20f)
            curveTo(10f, 0f, 32f, 0f, 32f, 16f)
            curveTo(32f, 0f, 54f, 0f, 54f, 20f)
            curveTo(54f, 30f, 42f, 40f, 32f, 60f)
            close()
        }
    }.build()

    val diamonds = ImageVector.Builder(
        name = "diamonds",
        defaultWidth = 16.dp,
        defaultHeight = 16.dp,
        viewportWidth = 64f,
        viewportHeight = 64f
    ).apply {
        path(fill = SolidColor(Color.Red)) {
            moveTo(32f, 60f)
            curveTo(24f, 46f, 24f, 46f, 10f, 32f)
            curveTo(24f, 18f, 24f, 18f, 32f, 4f)
            curveTo(40f, 18f, 40f, 18f, 54f, 32f)
            curveTo(40f, 46f, 40f, 46f, 32f, 60f)
            close()
        }
    }.build()

    // TODO: Spades and clubs
}
