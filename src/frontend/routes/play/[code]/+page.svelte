<script>
  import { onMount } from "svelte";
  import Client from "../../../../client/Client.js";
  import Splendid from "../../../components/Splendid.svelte";

  const client = new Client();

  const hasPlayer = client.hasPlayer;
  const isConnected = client.isConnected;
  const isInGame = client.isInGame;
  const actions = client.actions;
  const gameId = client.gameId;
  const gameType = client.gameType;
  const state = client.gameState;
  const events = client.gameEvents;
  const playerId = client.playerId;
  const nextEvent = client.nextEvent;

  export let data;
  let displayName = "";

  onMount(() => {
    client.game = data.code;
    client.on("error", (message) => alert(message));
    client.connect();
  });

  const selectDisplayName = () => {
    client.createPlayer(displayName);
  };

  const onGameAction = ({ detail: { name, args } }) => {
    client.perform(name, args);
  };

  const onEventProcessed = () => {
    client.advanceEvents();
  };
</script>

<svelte:head>
  <title>
    {$state.turn === $playerId ? "**YOUR TURN**" : ""}
    {$gameType} ({data.code}) on Splendid!
  </title>
</svelte:head>

{#if !$hasPlayer}
  <form on:submit|preventDefault={selectDisplayName}>
    <input type="text" bind:value={displayName} placeholder="Display name" />
    <input type="submit" value="Set display name" />
  </form>
{:else if !$isConnected || !$isInGame}
  Connecting...
{:else}
  <h2>Game ID: {$gameId}</h2>
  <hr />
  {#if $gameType === "Splendid"}
    <Splendid
      actions={$actions}
      state={$state}
      playerId={$playerId}
      events={$events}
      nextEvent={$nextEvent}
      on:action={onGameAction}
      on:eventProcessed={onEventProcessed}
    />
  {/if}

  <hr />

  <details>
    <summary>Debugging information</summary>
    <details>
      <summary>Available actions</summary>
      <pre>{JSON.stringify($actions)}</pre>
    </details>

    <details>
      <summary>Current state</summary>
      <pre>{JSON.stringify($state, null, 4)}</pre>
    </details>

    <details>
      <summary>Events</summary>
      <pre>{JSON.stringify($events, null, 4)}</pre>
    </details>
  </details>
{/if}
