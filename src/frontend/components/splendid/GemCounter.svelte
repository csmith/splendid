<script>
  import Gem from "./Gem.svelte";
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  export let type;
  export let amount;
  export let interactive;
  export let size = "normal";
</script>

<style>
  div {
    position: relative;
  }

  .small .token {
    font-size: 1em;
  }

  .small .amount {
    font-size: 0.75em;
  }

  .amount {
    position: absolute;
    right: -0.5em;
    bottom: -0.5em;
    background-color: var(--background);
    border: 2px solid var(--border);
    border-radius: 100%;
    width: 1.6em;
    height: 1.6em;
    text-align: center;
  }

  .empty {
    filter: grayscale(100%) opacity(0.4);
  }

  .token {
    display: block;
    font-size: 1.5em;
    width: 2.1em;
    height: 2.1em;
    border-radius: 50%;
    text-align: center;
    padding: 0.3em;
    border: 2px solid var(--border);
  }

  .selectable .token {
    cursor: pointer;
    box-shadow: 0 0 5px var(--highlight);
  }
</style>

<div
  class={size}
  class:empty={amount === 0}
  class:selectable={interactive}
  on:click|preventDefault={() => dispatch("click")}
>
  <span class="token {type}">
    <Gem {type} />
  </span>
  <span class="amount">{amount}</span>
</div>
