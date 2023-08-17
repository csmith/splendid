<script>
  import _ from "lodash";
  import { isFirstPlayer } from "../../../common/state.js";
  import ChangePhase from "../../../games/shared/events/ChangePhase.js";
  import ChangePlayer from "../../../games/shared/events/ChangePlayer.js";
  import SetPlayerOrder from "../../../games/shared/events/SetPlayerOrder.js";
  import AddBonus from "../../../games/splendid/events/AddBonus.js";
  import AddPlayer from "../../../games/splendid/events/AddPlayer.js";
  import AddPoints from "../../../games/splendid/events/AddPoints.js";
  import DiscardCard from "../../../games/splendid/events/DiscardCard.js";
  import DiscardReserve from "../../../games/splendid/events/DiscardReserve.js";
  import FinalRound from "../../../games/splendid/events/FinalRound.js";
  import PlaceCard from "../../../games/splendid/events/PlaceCard.js";
  import ReceiveNoble from "../../../games/splendid/events/ReceiveNoble.js";
  import RemoveCardFromDeck from "../../../games/splendid/events/RemoveCardFromDeck.js";
  import ReserveCard from "../../../games/splendid/events/ReserveCard.js";
  import ReturnTokens from "../../../games/splendid/events/ReturnTokens.js";
  import Setup from "../../../games/splendid/events/Setup.js";
  import TakeTokens from "../../../games/splendid/events/TakeTokens.js";

  export let state;
  export let events = [];

  const pretty = (e) => {
    let player = "";
    if (e.playerId && state.players[e.playerId]) {
      player = state.players[e.playerId].details.name;
    }

    switch (e.event) {
      case AddBonus.name:
        return `${player} receives a ${e.type} bonus`;
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
      case DiscardCard.name:
        return `A card is removed from the board`;
      case DiscardReserve.name:
        return `${player} buys a card from their reserve`;
      case FinalRound.name:
        return `This is the final round!`;
      case PlaceCard.name:
        return `A card is dealt to the board`;
      case ReceiveNoble.name:
        return `${player} receives a visit from a noble`;
      case RemoveCardFromDeck.name:
        return `A card is removed from the deck`;
      case ReserveCard.name:
        return `${player} reserves a card`;
      case ReturnTokens.name:
        return `${player} returns tokens to the supply: ${_.map(e.tokens, (v, k) => `${v} ${k}`).join(", ")}`;
      case SetPlayerOrder.name:
        return `The turn order will be ${e.order.map((o) => state.players[o].details.name).join(", ")}`;
      case Setup.name:
        return `The game has been configured for ${Object.values(state.players).length} players`;
      case TakeTokens.name:
        return `${player} obtains tokens from the supply: ${_.map(e.tokens, (v, k) => `${v} ${k}`).join(", ")}`;
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
      <li>
        {pretty(event)}
      </li>
      {#if event.event === ChangePlayer.name && event.playerId}
        <hr class="turn-end" class:round-end={isFirstPlayer(state, event.playerId)} />
      {/if}
    {/each}
  </ul>
</section>
