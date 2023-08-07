<script>
    import { calculateWinners } from '../../../splendid/util.js';
    export let state;

    export let playerId;

    $: pronoun = state.turn === playerId ? 'You' : 'They';

    const winnerNames = () => {
        const winners = calculateWinners(state).map((id) => state.players[id].details.name);
        if (winners.length === 1) {
            return `${winners[0]} wins!`
        } else {
            return `It's a tie! ${winners.join(' and ')} win!`
        }
    }
</script>

<style>
    section {
        width: 100%;
        text-align: center;
        font-size: 1.5em;
        background-color: var(--background-contrast);
        padding: 10px;
        margin-bottom: 10px;
    }

    .yours {
        background-color: var(--accent);
        color: var(--text-alt);
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
            Game over! {winnerNames()}
        {/if}
    {/if}
</section>