package com.chameth.splendid.shared.ui

// import androidx.compose.desktop.ui.tooling.preview.Preview
import androidx.compose.foundation.Image
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import com.chameth.splendid.shared.playingcards.Card
import com.chameth.splendid.shared.playingcards.Rank
import com.chameth.splendid.shared.playingcards.Suit
import org.jetbrains.compose.resources.painterResource
import splendid.shared.generated.resources.*

@Composable
fun PlayingCard(
    card: Card?,
    modifier: Modifier = Modifier,
) {
    if (card == null) {
        CardPlaceholder(modifier = modifier)
    } else if (card.visible) {
        Card(
            modifier = modifier.border(1.dp, Color.Black, RoundedCornerShape(8.dp))
        ) {
            Image(
                modifier = Modifier.fillMaxSize(),
                painter = painterResource(
                    when (card.suit) {
                        Suit.Hearts ->
                            when (card.rank) {
                                Rank.Ace -> Res.drawable.hearts_A
                                Rank.Two -> Res.drawable.hearts_2
                                Rank.Three -> Res.drawable.hearts_3
                                Rank.Four -> Res.drawable.hearts_4
                                Rank.Five -> Res.drawable.hearts_5
                                Rank.Six -> Res.drawable.hearts_6
                                Rank.Seven -> Res.drawable.hearts_7
                                Rank.Eight -> Res.drawable.hearts_8
                                Rank.Nine -> Res.drawable.hearts_9
                                Rank.Ten -> Res.drawable.hearts_T
                                Rank.Jack -> Res.drawable.hearts_J
                                Rank.Queen -> Res.drawable.hearts_Q
                                Rank.King -> Res.drawable.hearts_K
                            }
                        Suit.Diamonds ->
                            when (card.rank) {
                                Rank.Ace -> Res.drawable.diamonds_A
                                Rank.Two -> Res.drawable.diamonds_2
                                Rank.Three -> Res.drawable.diamonds_3
                                Rank.Four -> Res.drawable.diamonds_4
                                Rank.Five -> Res.drawable.diamonds_5
                                Rank.Six -> Res.drawable.diamonds_6
                                Rank.Seven -> Res.drawable.diamonds_7
                                Rank.Eight -> Res.drawable.diamonds_8
                                Rank.Nine -> Res.drawable.diamonds_9
                                Rank.Ten -> Res.drawable.diamonds_T
                                Rank.Jack -> Res.drawable.diamonds_J
                                Rank.Queen -> Res.drawable.diamonds_Q
                                Rank.King -> Res.drawable.diamonds_K
                            }
                        Suit.Spades ->
                            when (card.rank) {
                                Rank.Ace -> Res.drawable.spades_A
                                Rank.Two -> Res.drawable.spades_2
                                Rank.Three -> Res.drawable.spades_3
                                Rank.Four -> Res.drawable.spades_4
                                Rank.Five -> Res.drawable.spades_5
                                Rank.Six -> Res.drawable.spades_6
                                Rank.Seven -> Res.drawable.spades_7
                                Rank.Eight -> Res.drawable.spades_8
                                Rank.Nine -> Res.drawable.spades_9
                                Rank.Ten -> Res.drawable.spades_T
                                Rank.Jack -> Res.drawable.spades_J
                                Rank.Queen -> Res.drawable.spades_Q
                                Rank.King -> Res.drawable.spades_K
                            }
                        Suit.Clubs ->
                            when (card.rank) {
                                Rank.Ace -> Res.drawable.clubs_A
                                Rank.Two -> Res.drawable.clubs_2
                                Rank.Three -> Res.drawable.clubs_3
                                Rank.Four -> Res.drawable.clubs_4
                                Rank.Five -> Res.drawable.clubs_5
                                Rank.Six -> Res.drawable.clubs_6
                                Rank.Seven -> Res.drawable.clubs_7
                                Rank.Eight -> Res.drawable.clubs_8
                                Rank.Nine -> Res.drawable.clubs_9
                                Rank.Ten -> Res.drawable.clubs_T
                                Rank.Jack -> Res.drawable.clubs_J
                                Rank.Queen -> Res.drawable.clubs_Q
                                Rank.King -> Res.drawable.clubs_K
                            }
                    }
                ),
                contentDescription = card.toString()
            )
        }
    } else {
        CardBack(modifier = modifier)
    }
}

// @Preview
@Composable
internal fun PlayingCard_Preview() {
    MaterialTheme {
        Row(horizontalArrangement = Arrangement.spacedBy(16.dp)) {
            PlayingCard(
                modifier = Modifier.height(120.dp),
                card = Card(Suit.Clubs, Rank.Ace, visible = true)
            )

            PlayingCard(
                modifier = Modifier.height(120.dp),
                card = Card(Suit.Diamonds, Rank.Ace, visible = true)
            )

            PlayingCard(
                modifier = Modifier.height(120.dp),
                card = Card(Suit.Clubs, Rank.Ace, visible = false)
            )

            PlayingCard(
                modifier = Modifier.height(120.dp),
                card = null
            )
        }
    }
}
