<script>
    import {createEventDispatcher} from "svelte";
    import {canReceiveNoble} from "../../../splendid/util.js";

    const dispatch = createEventDispatcher();

    export let player;

    /** @type {import('../../../splendid/state.js').default} */
    export let state;

    export let canSelect;

    const receiveNoble = (noble) => {
        if (canSelect && canReceiveNoble(player, noble)) {
            dispatch('receiveNoble', noble);
        }
    }
</script>

<style>
    .noble {
        display: inline-block;
        width: 100px;
        height: 100px;
        border: 1px solid black;
        margin: 5px;
        position: relative;
        background-color: rebeccapurple;
    }

    .score {
        position: absolute;
        top: 0;
        left: 0;
        font-size: 2em;
        line-height: 1;
        padding: 5px;
        color: white;
        text-shadow: 0 0 5px black;
    }

    ul, li {
        margin: 0;
        padding: 0;
        list-style-type: none;
    }

    .costs {
        position: absolute;
        bottom: 0;
        right: 0;
    }

    .affordable {
        position: absolute;
        bottom: 0;
        left: 0;
        font-size: 1.5em;
        line-height: 1;
        padding: 5px;
    }

    .cost {
        padding: 2px 4px;
        display: block;
        border-top: 1px solid black;
        border-left: 1px solid black;
    }

    .selectable {
        cursor: pointer;
        box-shadow: 0 0 5px red;
    }
</style>

<section>
    {#each state.nobles as noble}
        <div class="noble" id="noble-{noble.id}" class:selectable={canSelect && canReceiveNoble(player, noble)} on:click={() => receiveNoble(noble)}>
            <span class="score">3</span>
            <ul class="costs">
                {#each Object.entries(noble.cost) as entry}
                    <li><span class="cost {entry[0]}">{entry[1]}</span></li>
                {/each}
            </ul>
            {#if player && canReceiveNoble(player, noble)}
                <span class="affordable">✅</span>
            {/if}
        </div>
    {/each}
</section>