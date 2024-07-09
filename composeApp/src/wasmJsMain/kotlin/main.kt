
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.MaterialTheme
import androidx.compose.ui.ExperimentalComposeUiApi
import androidx.compose.ui.Modifier
import androidx.compose.ui.window.ComposeViewport
import com.chameth.splendid.shared.engine.Game
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.actions.KlondikeActions
import com.chameth.splendid.games.klondike.ui.Board
import kotlinx.browser.document

@OptIn(ExperimentalComposeUiApi::class)
fun main() {
    ComposeViewport(document.body!!) {
        val game = Game(
            initialState = State(),
            id = "id",
            actionsGenerator = KlondikeActions::generate
        )

        MaterialTheme {
            Board(
                modifier = Modifier.fillMaxSize(),
                game = game
            )
        }
    }
}