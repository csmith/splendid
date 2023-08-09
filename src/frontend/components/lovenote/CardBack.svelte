<script>
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  export let count;
  export let id;
</script>

<style>
  .back {
    container-type: inline-size;
    border: 1px solid black;
    border-radius: 5px;
    box-shadow: 3px 2px 2px rgba(0, 0, 0, 0.3);
    display: grid;
    grid-template-rows: 10% 1fr 10%;
    justify-items: center;
    color: white;
    background-color: lightpink;
    padding: 10px 0;
    aspect-ratio: 2 / 3;
    width: 100px;
    height: 150px;
  }

  .count {
    grid-row: 2;
    align-self: center;
    font-size: x-large;
  }

  .many {
    box-shadow:
      3px 2px 1px rgba(0, 0, 0, 0.3),
      2px 1px 1px rgba(0, 0, 0, 0.3);
  }

  .many-many {
    box-shadow:
      5px 4px 1px rgba(0, 0, 0, 0.3),
      3px 2px 1px rgba(0, 0, 0, 0.3),
      2px 1px 1px rgba(0, 0, 0, 0.6);
  }
</style>

<div
  class="back"
  id="card-{id}"
  class:many={count && count > 2}
  class:many-many={count && count > 10}
  on:click={() => dispatch("click")}
>
  {#if count}
    <span class="count">{count}</span>
  {/if}
</div>
