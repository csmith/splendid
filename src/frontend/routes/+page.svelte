<script>
  import { onMount } from "svelte";
  import Client from "../../client/Client.js";
  import { goto } from "$app/navigation";

  const client = new Client();

  const hasPlayer = client.hasPlayer;
  const isConnected = client.isConnected;
  const isInGame = client.isInGame;

  let displayName = "";
  let joinGameId = "";

  onMount(() => {
    client.connect();
    client.on("error", (message) => alert(message));
    client.on("game-joined", (e) => goto(`/play/${e.id}`));
  });

  const startNewGame = () => {
    client.startGame("Splendid");
  };

  const joinExistingGame = () => {
    client.joinGame(joinGameId);
  };

  const selectDisplayName = () => {
    client.createPlayer(displayName);
  };
</script>

{#if !$hasPlayer}
  <form on:submit|preventDefault={selectDisplayName}>
    <input type="text" bind:value={displayName} placeholder="Display name" />
    <input type="submit" value="Set display name" />
  </form>
{:else if !$isConnected}
  Connecting...
{:else if !$isInGame}
  <button on:click={startNewGame}>Start new game</button>
  or
  <form on:submit|preventDefault={joinExistingGame}>
    <input type="text" bind:value={joinGameId} placeholder="Game ID" />
    <input type="submit" value="Join existing game" />
  </form>
{/if}
