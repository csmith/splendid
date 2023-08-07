<script>
    import _ from 'lodash';
    import GemCounter from "./GemCounter.svelte";

    /** @type {import('../../../splendid/state.js').default} */
    export let state;
</script>

<style>
    section {
        text-align: center;
    }

    section > ul > li {
        border: 1px solid var(--border);
        margin-bottom: 10px;
    }

    .tokens {
        display: flex;
        flex-direction: row;
        justify-content: center;
        gap: 8px;
        margin: 10px 0;
    }

    .discount {
        display: inline-block;
        border: 2px solid var(--border);
        width: 2.1em;
        height: 2.1em;
        line-height: 1;
        padding: 0.5em;
        text-align: center;
        font-weight: bold;
        margin: 5px;
        position: relative;
    }

    ul, li {
        margin: 0;
        padding: 0;
        list-style-type: none;
    }

    section > ul > li {
        padding: 5px;
    }

    h3 {
        margin: 0;
    }

    .points {
        display: inline-block;
        position: relative;
        background-color: var(--background);
        border: 1px solid var(--border);
        width: 1.4em;
        height: 1.4em;
        padding: 0.15em;
        line-height: 1;
    }
</style>

<section>
    <ul>
        {#each _.sortBy(Object.values(state.players), 'order') as player}
            <li id="player-{player.details.id}">
                <h3>{#if state.turn === player.details.id}⏵{/if} {player.details.name} <span class="points">{player.points}</span></h3>
                <ul>
                    <li class="tokens">
                        {#each Object.entries(player.tokens) as pair}
                            <GemCounter type={pair[0]} amount="{pair[1]}" size="small" />
                        {/each}
                    </li>
                    <li>
                        {#each Object.entries(player.bonuses) as pair}
                            <span class="discount {pair[0]}" title="{pair[0]}">{pair[1]}</span>
                        {/each}
                    </li>
                    <li>Reserved cards: {player.reserved.length}</li>
                    <li>Nobles: {player.nobles.length}</li>
                </ul>
            </li>
        {/each}
    </ul>
</section>