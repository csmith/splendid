
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.MaterialTheme
import androidx.compose.ui.ExperimentalComposeUiApi
import androidx.compose.ui.Modifier
import androidx.compose.ui.window.ComposeViewport
import com.chameth.splendid.allGames
import com.chameth.splendid.shared.engine.Game
import com.chameth.splendid.shared.util.now
import com.chameth.splendid.ui.GameSelector
import kotlinx.browser.document
import kotlinx.datetime.LocalDateTime

@OptIn(ExperimentalComposeUiApi::class)
fun main() {
    ComposeViewport(document.body!!) {
        MaterialTheme {
            GameSelector(
                modifier = Modifier.fillMaxSize(),
                types = allGames,
                createGame = { Game(it, LocalDateTime.now().toString()) },
            )
        }
    }
}