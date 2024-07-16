package com.chameth.splendid.games.diceclimbing.ui.components

import androidx.compose.foundation.Canvas
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import com.chameth.splendid.shared.Die

@Composable
fun Die(die: Die, modifier: Modifier = Modifier) {
    Box(
        modifier = modifier
            .size(48.dp)
            .border(1.dp, Color.Black, RoundedCornerShape(8.dp))
            .shadow(4.dp, RoundedCornerShape(8.dp))
            .background(Brush.radialGradient(0.5f to Color(0xEE, 0xEE, 0xEE), 1f to Color(0xDD, 0xDD, 0xDD)))
    ) {
        Canvas(modifier = Modifier.fillMaxSize()) {
            if (die.value == 1 || die.value == 3 || die.value == 5) {
                // Center pip

                drawCircle(
                    color = Color.Black,
                    center = Offset(size.width / 2, size.height / 2),
                    radius = size.width / 10
                )
            }

            if (die.value == 2 || die.value == 3 || die.value == 4 || die.value == 5 || die.value == 6) {
                // Bottom-left + top-right
                drawCircle(
                    color = Color.Black,
                    center = Offset(size.width / 4, 3 * size.height / 4),
                    radius = size.width / 10
                )
                drawCircle(
                    color = Color.Black,
                    center = Offset(3 * size.width / 4, size.height / 4),
                    radius = size.width / 10
                )
            }

            if (die.value == 4 || die.value == 5 || die.value == 6) {
                // Top-left + Bottom-right

                drawCircle(
                    color = Color.Black,
                    center = Offset(size.width / 4, size.height / 4),
                    radius = size.width / 10
                )

                drawCircle(
                    color = Color.Black,
                    center = Offset(3 * size.width / 4, 3 * size.height / 4),
                    radius = size.width / 10
                )
            }

            if (die.value == 6) {
                // Center pair

                drawCircle(
                    color = Color.Black,
                    center = Offset(size.width / 4, size.height / 2),
                    radius = size.width / 10
                )

                drawCircle(
                    color = Color.Black,
                    center = Offset(3 * size.width / 4, size.height / 2),
                    radius = size.width / 10
                )
            }
        }
    }
}