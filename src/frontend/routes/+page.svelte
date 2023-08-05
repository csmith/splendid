<script>
    import {onMount} from "svelte";
    import Client from '../../client/Client.js';
    import Splendid from "../components/Splendid.svelte";

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

    let displayName = '';
    let joinGameId = '';

    onMount(() => {
        client.connect();
        client.on('error', (message) => alert(message))
    })

    const startNewGame = () => {
        client.startGame('Splendid');
    };

    const joinExistingGame = () => {
        client.joinGame(joinGameId);
    }

    const selectDisplayName = () => {
        client.createPlayer(displayName);
    }

    const onGameAction = ({detail: {name, args}}) => {
        client.perform(name, args);
    }

    const onEventProcessed = () => {
        client.advanceEvents();
    }
</script>

{#if !$hasPlayer}
    <form on:submit|preventDefault={selectDisplayName}>
        <input type="text" bind:value={displayName} placeholder="Display name">
        <input type="submit" value="Set display name">
    </form>
{:else if !$isConnected}
    Connecting...
{:else if !$isInGame}
    <button on:click={startNewGame}>Start new game</button>
    or
    <form on:submit|preventDefault={joinExistingGame}>
        <input type="text" bind:value={joinGameId} placeholder="Game ID">
        <input type="submit" value="Join existing game">
    </form>
{:else}
    <h1>Game ID: {$gameId}</h1>
    <hr>
    {#if $gameType === 'Splendid'}
        <Splendid
                actions={$actions}
                state={$state}
                playerId={$playerId}
                events={$events}
                nextEvent={$nextEvent}
                on:action={onGameAction}
                on:eventProcessed={onEventProcessed}/>
    {/if}
    <hr>
    Available actions:
    {JSON.stringify($actions)}
    <hr>
    State:
    {JSON.stringify($state)}
{/if}
