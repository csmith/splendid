<script>
  import diamond from "./gems/Diamond.svelte";
  import emerald from "./gems/Emerald.svelte";
  import gold from "./gems/Gold.svelte";
  import onyx from "./gems/Onyx.svelte";
  import ruby from "./gems/Ruby.svelte";
  import sapphire from "./gems/Sapphire.svelte";

  const types = {
    diamond,
    emerald,
    gold,
    onyx,
    ruby,
    sapphire,
  };

  export let type;
</script>

<svelte:component this={types[type]} />
