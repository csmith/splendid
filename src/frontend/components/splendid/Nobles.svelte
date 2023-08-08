<script>
  import { createEventDispatcher } from "svelte";
  import { canReceiveNoble } from "../../../games/splendid/util.js";
  import Noble from "./Noble.svelte";

  const dispatch = createEventDispatcher();

  export let player;

  /** @type {import('../../../games/splendid/state.js').default} */
  export let state;

  export let canSelect;

  const receiveNoble = (noble) => {
    if (canSelect && canReceiveNoble(player, noble)) {
      dispatch("receiveNoble", noble);
    }
  };
</script>

<style>
  div {
    margin: 10px 0;
    display: flex;
    gap: 10px;
  }
</style>

<section>
  <h3>Wandering nobles</h3>
  <div>
    {#each state.nobles as noble}
      <Noble {noble} {player} interactive={canSelect} on:click={() => receiveNoble(noble)} />
    {/each}
  </div>
</section>
