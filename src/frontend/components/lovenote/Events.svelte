<script>
  import ChangePhase from "../../../games/shared/events/ChangePhase.js";
  import ChangePlayer from "../../../games/shared/events/ChangePlayer.js";
  import SetPlayerOrder from "../../../games/shared/events/SetPlayerOrder.js";

  export let state;
  export let events = [];

  const pretty = (e) => {
    let player = "";
    if (e.playerId && state.players[e.playerId]) {
      player = state.players[e.playerId].details.name;
    }

    switch (e.event) {
      case "add-player":
        return `${e.details.name} joins the game`;
      case "add-points":
        return `${player} gains ${e.points} points`;
      case ChangePlayer.name:
        if (player) {
          return `It is now ${player}'s turn`;
        } else {
          return `It is no-one's turn`;
        }
      case ChangePhase.name:
        return `The game is now in the "${e.phase}" phase`;
      case "deal-card":
        return `${player} is dealt a card: ${e.card.type || "?"}`;
      case "discard-card":
        return `${player} discards a card: ${e.card.type}`;
      case "hand-revealed":
        return `${player} gets to see ${state.players[e.handPlayerId].details.name}'s hand: ${e.hand
          .map((c) => c.type || "?")
          .join(", ")}`;
      case "eliminate-player":
        return `${player} is eliminated from the round: ${e.reason}`;
      case "round-over":
        return `Round over! ${state.players[e.winningPlayerId].details.name} wins and gets one token of affection.`;
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
      {#if event.event === "round-over"}
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
