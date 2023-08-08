<script>
  import _ from "lodash";
  import { createEventDispatcher, onMount } from "svelte";
  import GemCounter from "./GemCounter.svelte";

  const dispatch = createEventDispatcher();

  export let player;

  let dialog;
  let selection = {
    emerald: 0,
    diamond: 0,
    sapphire: 0,
    onyx: 0,
    ruby: 0,
    gold: 0,
  };

  $: totalTokens = _.sum(Object.values(player.tokens));
  $: selectedTokens = _.sum(Object.values(selection));
  $: numberToDiscard = totalTokens - 10;
  $: numberRemaining = numberToDiscard - selectedTokens;
  $: selectedEnough = totalTokens - selectedTokens === 10;

  onMount(() => {
    dialog.showModal();
  });

  const selectToken = (token) => {
    const newValue = (selection[token] || 0) + 1;
    if (newValue <= player.tokens[token]) {
      selection[token] = newValue;
    }
  };

  const deselectToken = (token) => {
    const newValue = (selection[token] || 0) - 1;
    if (newValue >= 0) {
      selection[token] = newValue;
    }
  };

  const handleSubmit = () => {
    dialog.close();
    dispatch("tokensDiscarded", selection);
  };
</script>

<style>
  dialog {
    text-align: center;
  }

  ul {
    display: flex;
    gap: 15px;
    margin-bottom: 25px;
    justify-content: center;
  }

  p {
    margin: 15px 0;
  }
</style>

<dialog bind:this={dialog}>
  <p>You have too many tokens! Select which to discard:</p>
  <ul>
    {#each Object.entries(player.tokens) as pair}
      <li>
        <GemCounter
          type={pair[0]}
          amount={pair[1] - selection[pair[0]]}
          interactive={pair[1] - selection[pair[0]] > 0}
          on:click={() => selectToken(pair[0])}
        />
      </li>
    {/each}
  </ul>
  <p>&darr;</p>
  <ul>
    {#each Object.entries(selection) as pair}
      <li>
        <GemCounter type={pair[0]} amount={pair[1]} interactive={pair[1] > 0} on:click={() => deselectToken(pair[0])} />
      </li>
    {/each}
  </ul>
  <p>You have {totalTokens} tokens, so must discard {numberToDiscard} to bring your total down to 10.</p>
  <p>
    You have selected {selectedTokens}.
    {#if numberRemaining === 0}
      Good job!
    {:else}
      Select
      {#if numberRemaining > 0}
        {numberRemaining} more.
      {:else}
        {numberRemaining * -1} less.
      {/if}
    {/if}
  </p>
  <button on:click={handleSubmit} disabled={!selectedEnough}>Discard</button>
</dialog>
