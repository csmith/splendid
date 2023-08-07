<script>
    import {createEventDispatcher} from "svelte";
    import Players from "./splendid/Players.svelte";
    import Grid from "./splendid/Grid.svelte";
    import Gems from "./splendid/Gems.svelte";
    import Nobles from "./splendid/Nobles.svelte";
    import DiscardTokens from "./splendid/DiscardTokens.svelte";
    import Banner from "./splendid/Banner.svelte";
    import Events from "./splendid/Events.svelte";
    import EventHandler from "./splendid/EventHandler.svelte";

    const dispatch = createEventDispatcher();

    export let actions = [];
    export let events = [];
    export let state = {};
    export let playerId = '';
    export let nextEvent = undefined;

    $: showJoin = actions.includes('join');
    $: showStart = actions.includes('start');
    $: selectCard = actions.includes('buy-card') || actions.includes('reserve-card');
    $: takeTokens = actions.includes('take-tokens');
    $: receiveNoble = actions.includes('receive-noble');
    $: discardTokens = actions.includes('discard-tokens');

    const handleJoinClick = () => dispatch('action', {name: 'join'});
    const handleStartClick = () => dispatch('action', {name: 'start'});
    const handleSelectedGems = ({detail: selection}) => dispatch('action', {
        name: 'take-tokens',
        args: {tokens: selection}
    })
    const handleBuyCard = ({detail}) => dispatch('action', {name: 'buy-card', args: {card: detail}});
    const handleReserveCard = ({detail}) => dispatch('action', {name: 'reserve-card', args: {card: detail}});
    const handleReserveFromDeck = ({detail: level}) => dispatch('action', {
        name: 'reserve-card-from-deck',
        args: {level}
    });

    const handleReceiveNoble = ({detail}) => dispatch('action', {name: 'receive-noble', args: {noble: detail}});
    const handleTokensDiscarded = ({detail}) => dispatch('action', {name: 'discard-tokens', args: {tokens: detail}});
</script>

<style>
    :root {
        --emerald-colour: #9defd8;
        --diamond-colour: #f2fbfe;
        --sapphire-colour: #c1ccf9;
        --onyx-colour: #babbc4;
        --ruby-colour: #fecccf;
        --gold-colour: #fbf1b5;
    }

    :global(.emerald) {
        background-color: var(--emerald-colour);
        color: black;
    }

    :global(.diamond) {
        background-color: var(--diamond-colour);
        color: black;
    }

    :global(.onyx) {
        background-color: var(--onyx-colour);
        color: black;
    }

    :global(.ruby) {
        background-color: var(--ruby-colour);
        color: black;
    }

    :global(.sapphire) {
        background-color: var(--sapphire-colour);
        color: black;
    }

    :global(.gold) {
        background-color: var(--gold-colour);
        color: black;
    }

    .board {
        display: grid;
        grid-template-columns: 1fr auto;
        grid-gap: 10px;
    }
</style>

{#if showJoin}
    <div class="single-action">
        <button on:click={handleJoinClick}>Join</button>
    </div>
{:else if showStart}
    <div class="single-action">
        <button on:click={handleStartClick}>Start</button>
    </div>
{/if}

<section>
    <Banner state={state} playerId={playerId}/>
    <div class="board">
        <Grid
                state={state}
                canSelect={selectCard}
                player={state.players[playerId]}
                on:buyCard={handleBuyCard}
                on:reserveCard={handleReserveCard}
                on:reserveFromDeck={handleReserveFromDeck}/>
        <Players state={state}/>
        <Gems state={state} canTake={takeTokens} on:selected={handleSelectedGems}/>
        <Nobles
                state={state}
                canSelect={receiveNoble}
                player={state.players[playerId]}
                on:receiveNoble={handleReceiveNoble}/>
        {#if discardTokens}
            <DiscardTokens
                    state={state}
                    player={state.players[playerId]}
                    on:tokensDiscarded={handleTokensDiscarded}/>
        {/if}
    </div>
    <Events state={state} events={events}/>
    <EventHandler
            state={state}
            nextEvent={nextEvent}
            playerId={playerId}
            on:eventProcessed={() => dispatch('eventProcessed')}/>
</section>