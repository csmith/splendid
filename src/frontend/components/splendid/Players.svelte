<script>
    import _ from 'lodash';

    /** @type {import('../../../splendid/state.js').default} */
    export let state;
</script>

<style>
    section {
        width: 300px;
        text-align: center;
    }

    section > ul > li {
        border: 1px solid black;
        margin-bottom: 10px;
    }

    .token {
        display: inline-block;
        border: 2px solid black;
        border-radius: 50%;
        width: 1em;
        height: 1em;
        line-height: 1;
        padding: 0.5em;
        text-align: center;
        color: white;
        font-weight: bold;
        margin: 5px;
    }

    .discount {
        display: inline-block;
        border: 2px solid black;
        width: 1em;
        height: 1em;
        line-height: 1;
        padding: 0.5em;
        text-align: center;
        color: white;
        font-weight: bold;
        margin: 5px;
    }

    ul, li {
        margin: 0;
        padding: 0;
        list-style-type: none;
    }

    h3 {
        margin: 0;
    }
</style>

<section>
    <ul>
        {#each _.sortBy(Object.values(state.players), 'order') as player}
            <li>
                <h3>{#if state.turn === player.details.id}⏵{/if} {player.details.name} ({player.points})</h3>
                <ul>
                    <li>
                        {#each Object.entries(player.tokens) as pair}
                            <span class="token {pair[0]}" title="{pair[0]}">{pair[1]}</span>
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