package com.chameth.splendid.games.diceclimbing.ui.components

import androidx.compose.ui.geometry.Size
import androidx.compose.ui.graphics.Outline
import androidx.compose.ui.graphics.Path
import androidx.compose.ui.graphics.Shape
import androidx.compose.ui.unit.Density
import androidx.compose.ui.unit.LayoutDirection

class DiamondShape : Shape {
    override fun createOutline(size: Size, layoutDirection: LayoutDirection, density: Density) =
        Outline.Generic(Path().apply {
            moveTo(0f, size.height / 2)
            lineTo(size.width / 2, 0f)
            lineTo(size.width, size.height/2)
            lineTo(size.width / 2, size.height)
            close()
        })

}