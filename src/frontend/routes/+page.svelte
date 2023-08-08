<script>
  import { onMount } from "svelte";
  import Client from "../../client/Client.js";
  import { goto } from "$app/navigation";
  import games from "../../games.js";

  const client = new Client();

  const hasPlayer = client.hasPlayer;
  const isConnected = client.isConnected;

  let displayName = "";
  let joinGameId = "";

  onMount(() => {
    client.connect();
    client.on("error", (message) => alert(message));
    client.on("game-joined", (e) => goto(`/play/${e.id}`));
  });

  const startNewGame = (name) => {
    client.startGame(name);
  };

  const joinExistingGame = () => {
    client.joinGame(joinGameId);
  };

  const selectDisplayName = () => {
    client.createPlayer(displayName);
  };
</script>

<style>
  h2,
  h3,
  p {
    margin: 0.8em 0;
  }
</style>

{#if !$hasPlayer}
  <form on:submit|preventDefault={selectDisplayName}>
    <input type="text" bind:value={displayName} placeholder="Display name" />
    <input type="submit" value="Set display name" />
  </form>
{:else if !$isConnected}
  Connecting...
{:else}
  <h2>Join an existing game</h2>
  <form on:submit|preventDefault={joinExistingGame}>
    <input type="text" bind:value={joinGameId} placeholder="Game ID" />
    <input type="submit" value="Join existing game" />
  </form>
  <h2>Start a new game</h2>
  {#each Object.values(games) as game}
    <h3>{game.name}</h3>
    <p>{game.description}</p>
    <p class="stats">
      Players: {game.players.min}&ndash;{game.players.max}. Based on:
      <a href={game.based_on.link}>{game.based_on.game} by {game.based_on.creator}</a>. If you enjoy this game, please
      consider <a href={game.based_on.purchase}>purchasing a physical copy</a>.
    </p>
    <button on:click={() => startNewGame(game.name)}>Start new game</button>
  {/each}
{/if}
