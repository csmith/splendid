<script>
  import { createEventDispatcher } from "svelte";
  import EventHandler from "./lovenote/EventHandler.svelte";
  import CardBack from "./lovenote/CardBack.svelte";
  import Card from "./lovenote/Card.svelte";
  import _ from "lodash";
  import Events from "./lovenote/Events.svelte";
  import { areAllProtected } from "../../games/lovenote/util.js";

  const dispatch = createEventDispatcher();

  export let actions = [];
  export let events = [];
  export let state = {};
  export let playerId = "";
  export let nextEvent = undefined;

  $: showJoin = actions.includes("join");
  $: showStart = actions.includes("start");
  $: canPlay = actions.includes("play-card");

  $: players = Object.values(state.players);
  $: orderedPlayers = _.sortBy(players, "order");

  const needsTarget = (type) => ["Guard", "Priest", "Baron", "Prince", "King"].includes(type);
  const maySelectSelf = (type) => type === "Prince";
  const needsType = (type) => type === "Guard";

  let selectedCard = undefined;
  let selectTarget = false;
  let selectedTarget = undefined;
  let selectTypeDialog = undefined;

  const handleJoinClick = () => dispatch("action", { name: "join" });
  const handleStartClick = () => dispatch("action", { name: "start" });
  const handleCardClick = (card) => {
    console.log(card);
    if (canPlay) {
      if (needsTarget(card.type) && !areAllProtected(state, playerId)) {
        selectedCard = card;
        selectTarget = true;
      } else {
        dispatch("action", { name: "play-card", args: { cardId: card.id } });
      }
    }
  };

  const handleTargetClick = (player) => {
    if (canPlay) {
      if (needsType(selectedCard.type)) {
        selectedTarget = player;
        selectTypeDialog.showModal();
      } else {
        dispatch("action", { name: "play-card", args: { cardId: selectedCard.id, targetPlayerId: player.details.id } });
        selectedCard = undefined;
        selectTarget = false;
      }
    } else {
      selectedCard = undefined;
      selectTarget = false;
    }
  };

  const handleTypeClick = (type) => {
    dispatch("action", {
      name: "play-card",
      args: { cardId: selectedCard.id, targetPlayerId: selectedTarget.details.id, guessedType: type },
    });

    selectedCard = undefined;
    selectTarget = false;
    selectTypeDialog.close();
  };
</script>

<style>
  section {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-template-rows: repeat(3, 1fr);
  }

  #shared {
    grid-column: 2;
    grid-row: 2;
    display: flex;
  }

  .cards,
  .discards {
    display: flex;
    flex-direction: row;
    height: 150px;
  }

  #player-0 {
    grid-column: 2;
    grid-row: 1;
  }

  #player-1 {
    grid-column: 3;
    grid-row: 2;
  }

  #player-2 {
    grid-column: 2;
    grid-row: 3;
  }

  #player-3 {
    grid-column: 1;
    grid-row: 2;
  }

  .eliminated {
    text-decoration: line-through;
  }

  #quickref {
    border: 1px solid var(--accent);
    background-color: var(--background-contrast);
    border-radius: 10px;
    margin: 40px;
    padding: 20px;
    font-size: small;
  }
</style>

{#if showJoin}
  <div class="single-action">
    <button on:click={handleJoinClick}>Join</button>
  </div>
{:else if showStart}
  <div class="single-action">
    <button on:click={handleStartClick}>Start</button>
  </div>
{/if}

<section>
  <div id="shared">
    <CardBack id="deck" count={state.deck.length} />
  </div>
  <div id="quickref">
    <h3>Quick reference</h3>
    <ul>
      <li>
        1: Guard (5): Player may choose another player and name a card other than Guard. If the chosen player's hand
        contains the same card as named, that player is eliminated from the round.
      </li>
      <li>2: Priest (2): Player may privately see another player's hand.</li>
      <li>
        3: Baron (2): Player may choose another player and privately compare hands. The player with the lower-value card
        is eliminated from the round.
      </li>
      <li>4: Handmaid (2): Player cannot be affected by any other player's cards until their next turn.</li>
      <li>
        5: Prince (2): Player may choose any player (including themselves) to discard their hand and draw a new one.
      </li>
      <li>6: King (1): Player may trade hands with another player.</li>
      <li>
        7: Countess (1): Does nothing when played, but if the player has this card and either the King or the Prince,
        this card must be played immediately.
      </li>
      <li>
        8: Princess (1): If the player plays or discards this card for any reason, they are eliminated from the round.
      </li>
    </ul>
  </div>
  {#each orderedPlayers as player, i}
    <div id="player-{i}">
      <h3 class:eliminated={player.eliminated}>
        {player.details.name}
        {#if player.protected}
          🛡️
        {/if}
        {#if selectTarget && (player.details.id !== playerId || maySelectSelf(selectedCard.type)) && !player.protected && !player.eliminated}
          <button on:click={() => handleTargetClick(player)}>Target this player</button>
        {/if}
      </h3>
      {#each Array(player.points) as i}💖{/each}
      <h4>Hand</h4>
      <div class="cards">
        {#each player.hand as card}
          {#if card.type}
            <Card {card} on:click={() => handleCardClick(card)} />
          {:else}
            <CardBack id={card.id} />
          {/if}
        {/each}
      </div>
      <h4>Discards</h4>
      <div class="discards">
        {#each player.discards as card}
          <Card {card} />
        {/each}
      </div>
    </div>
  {/each}
</section>

<EventHandler {state} {nextEvent} {playerId} on:eventProcessed={() => dispatch("eventProcessed")} />

<dialog bind:this={selectTypeDialog}>
  <h3>Select a card type</h3>
  <ul>
    {#each ["Priest", "Baron", "Handmaid", "Prince", "King", "Countess", "Princess"] as type}
      <li>
        <button on:click={() => handleTypeClick(type)}>{type}</button>
      </li>
    {/each}
  </ul>
</dialog>

<hr />
<Events {state} {events} />
