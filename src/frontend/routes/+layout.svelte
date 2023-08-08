<script>
  import "../style/global.scss";
</script>

<svelte:head>
  <title>Splendid!</title>
</svelte:head>

<header id="site-header">
  <h1><img src="/splendid.png" alt="A splendid looking meeple" /> Splendid!</h1>
</header>

<main>
  <slot />
</main>

<footer id="site-footer">Look at me, I'm a footer!</footer>
