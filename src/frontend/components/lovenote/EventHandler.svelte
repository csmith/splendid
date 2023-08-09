<script>
  import { createEventDispatcher } from "svelte";

  export let state;
  export let playerId;
  export let nextEvent = undefined;

  const dispatch = createEventDispatcher();

  $: if (nextEvent) {
    dispatch("eventProcessed");
  }
</script>
