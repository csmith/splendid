<script>
  import { onMount } from "svelte";
  import Client from "../../client/Client.js";
  import { goto } from "$app/navigation";
  import games from "../../games.js";
  import GameDetails from "../components/meta/GameDetails.svelte";

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

  const startNewGame = (name, options = {}) => {
    client.startGame(name, options);
  };

  const joinExistingGame = () => {
    client.joinGame(joinGameId);
  };

  const selectDisplayName = () => {
    client.createPlayer(displayName);
  };
</script>

<style>
  h2 {
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
    <GameDetails {game} on:startGame={(e) => startNewGame(e.detail.name, e.detail.options)} />
  {/each}
{/if}
