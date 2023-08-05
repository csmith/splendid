<script>
    export let state;

    export let playerId;

    $: pronoun = state.turn === playerId ? 'You' : 'They';
</script>

<style>
    section {
        width: 100%;
        text-align: center;
        font-size: 1.5em;
        background-color: lightgray;
        padding: 10px;
        margin-bottom: 10px;
    }

    .yours {
        background-color: lightcoral;
    }
</style>

<section class:yours={state.turn === playerId}>
    {#if !state.players[playerId]}
        You are spectating.
    {:else}
        {#if state.turn === playerId}
            It is your turn.
        {:else if state.turn}
            It is {state.players[state.turn].details.name}'s turn.
        {/if}

        {#if state.phase === 'setup'}
            The game has not yet begun!
        {:else if state.phase === 'play'}
            {pronoun} must take tokens, reserve a card, or buy a card.
        {:else if state.phase === 'discard'}
            {pronoun} must discard excess tokens.
        {:else if state.phase === 'noble'}
            {pronoun} must select a noble to receive.
        {:else if state.phase === 'end'}
            Game over!
        {/if}
    {/if}
</section>