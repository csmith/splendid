<script>
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  export let card;
</script>

<style>
  .card {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-template-rows: 40px 1fr 40px;
    grid-column-gap: 5px;
    padding: 0;
    background: #d0b6ba;
    border: 1px solid black;
    border-radius: 5px;
    box-shadow: 3px 2px 2px rgba(0, 0, 0, 0.3);
    cursor: pointer;
    width: 100px;
    height: 150px;
  }

  .score {
    padding-left: 10px;
    font-size: 25px;
    justify-self: start;
    align-self: center;
    color: white;
    font-weight: bold;
    text-shadow:
      0 0 5px black,
      0 0 2px black,
      0 0 1px black;
  }

  .title {
    grid-row: 3;
    grid-column: 1 / span 2;
    text-align: center;
    align-self: center;
    color: black;
    font-weight: bold;
  }
</style>

<div id="card-{card.id}" class="card" on:click={() => dispatch("click")}>
  <span class="score">{card.closeness}</span>
  <span class="title">{card.type}</span>
</div>
