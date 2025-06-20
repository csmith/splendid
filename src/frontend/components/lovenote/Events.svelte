<script>
  import ChangePhase from "../../../games/shared/events/ChangePhase.js";
  import ChangePlayer from "../../../games/shared/events/ChangePlayer.js";
  import SetPlayerOrder from "../../../games/shared/events/SetPlayerOrder.js";
  import AddPoints from "../../../games/splendid/events/AddPoints.js";
  import DealCard from "../../../games/lovenote/events/DealCard.js";
  import EliminatePlayer from "../../../games/lovenote/events/EliminatePlayer.js";
  import DiscardCard from "../../../games/lovenote/events/DiscardCard.js";
  import HandRevealed from "../../../games/lovenote/events/HandRevealed.js";
  import RoundOver from "../../../games/lovenote/events/RoundOver.js";
  import AddPlayer from "../../../games/shared/events/AddPlayer.js";

  export let state;
  export let events = [];

  const pretty = (e) => {
    let player = "";
    if (e.playerId && state.players[e.playerId]) {
      player = state.players[e.playerId].details.name;
    }

    switch (e.event) {
      case AddPlayer.name:
        return `${e.details.name} joins the game`;
      case AddPoints.name:
        return `${player} gains ${e.points} points`;
      case ChangePlayer.name:
        if (player) {
          return `It is now ${player}'s turn`;
        } else {
          return `It is no-one's turn`;
        }
      case ChangePhase.name:
        return `The game is now in the "${e.phase}" phase`;
      case DealCard.name:
        return `${player} is dealt a card: ${e.card.type || "?"}`;
      case DiscardCard.name:
        return `${player} discards a card: ${e.card.type}`;
      case HandRevealed.name:
        return `${player} gets to see ${state.players[e.handPlayerId].details.name}'s hand: ${e.hand
          .map((c) => c.type || "?")
          .join(", ")}`;
      case EliminatePlayer.name:
        return `${player} is eliminated from the round: ${e.reason}`;
      case RoundOver.name:
        if (e.winningPlayerIds.length === 1) {
          return `Round over! ${state.players[e.winningPlayerIds[0]].details.name} wins and gets one token of affection.`;
        } else {
          const winners = e.winningPlayerIds.map((id) => state.players[id].details.name).join(", ");
          return `Round over! Tied round - ${winners} each get one token of affection.`;
        }
      case SetPlayerOrder.name:
        return `The turn order will be ${e.order.map((o) => state.players[o].details.name).join(", ")}`;
      case "setup":
        return `The game has been configured for ${Object.values(state.players).length} players`;
      default:
        return `Unknown event: ${JSON.stringify(e)}`;
    }
  };
</script>

<style>
  .round-end {
    border-width: 3px;
    background-color: black;
  }
</style>

<section>
  <h3>Event history</h3>
  <ul>
    {#each events.slice().reverse() as event}
      {#if event.event === RoundOver.name}
        <hr class="round-end" />
      {/if}
      <li>
        {pretty(event)}
      </li>
      {#if event.event === ChangePlayer.name && event.playerId}
        <hr class="turn-end" />
      {/if}
    {/each}
  </ul>
</section>
