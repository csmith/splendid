<script>
  import { canAffordCard } from "../../../games/splendid/util.js";
  import { createEventDispatcher } from "svelte";
  import Gem from "./Gem.svelte";
  import Card from "./Card.svelte";
  import CardBack from "./CardBack.svelte";
  import CardPlaceholder from "./CardPlaceholder.svelte";

  const dispatch = createEventDispatcher();

  /** @type {import('../../../games/splendid/state.js').default} */
  export let state;

  export let player;
  export let canSelect = false;

  let buyDialog;
  let selectedCard = null;

  const levels = [2, 1, 0];

  const selectDeck = (level) => {
    if (canSelect) {
      dispatch("reserveFromDeck", level + 1);
    }
  };

  const selectCard = (card) => {
    if (player) {
      selectedCard = card;
      buyDialog.showModal();
    }
  };

  const buyCard = () => {
    if (canSelect) {
      buyDialog.close();
      dispatch("buyCard", selectedCard);
      selectedCard = null;
    }
  };

  const reserveCard = () => {
    if (canSelect) {
      buyDialog.close();
      dispatch("reserveCard", selectedCard);
      selectedCard = null;
    }
  };

  const handleDialogClick = (e) => {
    if (e.target === e.currentTarget) {
      buyDialog.close();
    }
  };

  const handleBuyDialogClose = () => {
    selectedCard = null;
  };
</script>

<style>
  section {
    display: grid;
    grid-template-columns: repeat(5, 100px);
    grid-template-rows: repeat(3, 150px) 20px 150px;
    grid-gap: 10px;
  }

  .header {
    color: white;
  }

  ul,
  li {
    margin: 0;
    padding: 0;
    list-style-type: none;
  }

  dialog > div {
    display: grid;
    grid-template-columns: auto 1fr;
    grid-gap: 20px;
  }

  dialog h2,
  dialog nav {
    grid-column: 1 / span 2;
    text-align: center;
    margin: 0;
  }

  nav {
    display: flex;
    justify-content: space-around;
  }

  table {
    border-collapse: collapse;
  }

  td,
  th {
    border: 1px solid black;
    padding: 2px 10px;
    text-align: center;
  }

  th.header {
    width: 40px;
  }

  th {
    font-weight: bold;
  }

  td.satisfied {
    background-color: lightgreen;
  }

  .tokens {
    grid-column: 1 / span 2;
  }

  .reserve {
    grid-column: 3 / span 3;
  }
</style>

<section class:canSelect>
  {#each levels as level}
    {#if state.decks[level].length > 0}
      <CardBack level={level + 1} count={state.decks[level].length} on:click={() => selectDeck(level)} />
    {:else}
      <CardPlaceholder id="deck{level}" />
    {/if}
    {#each state.cards[level] as card, i}
      {#if card}
        <Card {card} {player} on:click={() => selectCard(card)} />
      {:else}
        <CardPlaceholder classes="placeholder level{level}" />
      {/if}
    {/each}
  {/each}
  {#if player}
    <div class="tokens"></div>
    <div class="reserve">Your reserve</div>
    <div class="tokens"></div>
    {#each [0, 1, 2] as index}
      {#if player.reserved.length > index}
        <Card card={player.reserved[index]} {player} on:click={() => selectCard(player.reserved[index])} />
      {:else}
        <CardPlaceholder id="reserve-{index}" />
      {/if}
    {/each}
  {/if}
</section>

<dialog bind:this={buyDialog} on:click={handleDialogClick} on:close={handleBuyDialogClose}>
  {#if selectedCard}
    <div>
      <h2>Selected card</h2>
      <Card card={selectedCard} />
      <table>
        <tbody>
          <tr>
            <th>Gem</th>
            {#each Object.keys(selectedCard.cost) as key}
              <th class="header {key}">
                <Gem type={key} />
              </th>
            {/each}
          </tr>
          <tr>
            <th>Cost</th>
            {#each Object.values(selectedCard.cost) as value}
              <td>{value}</td>
            {/each}
          </tr>
          <tr>
            <th>Bonuses</th>
            {#each Object.keys(selectedCard.cost) as key}
              <td>{-player.bonuses[key]}</td>
            {/each}
          </tr>
          <tr>
            <th>Your gems</th>
            {#each Object.entries(selectedCard.cost) as pair}
              <td>{-Math.min(Math.max(0, pair[1] - player.bonuses[pair[0]]), player.tokens[pair[0]])}</td>
            {/each}
          </tr>
          <tr>
            <th>Remaining</th>
            {#each Object.entries(selectedCard.cost) as pair}
              <td class:satisfied={Math.max(0, pair[1] - player.bonuses[pair[0]] - player.tokens[pair[0]]) === 0}>
                {Math.max(0, pair[1] - player.bonuses[pair[0]] - player.tokens[pair[0]])}
              </td>
            {/each}
          </tr>
        </tbody>
      </table>
      <nav>
        <button on:click={buyCard} disabled={!canSelect || !canAffordCard(player, selectedCard)}> Buy card </button>
        <button
          on:click={reserveCard}
          disabled={!canSelect ||
            player.reserved.length >= 3 ||
            player.reserved.some((card) => card.id === selectedCard.id)}
        >
          Reserve card
        </button>
      </nav>
    </div>
  {/if}
</dialog>
