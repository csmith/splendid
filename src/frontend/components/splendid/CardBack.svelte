<script>

    import {createEventDispatcher} from "svelte";

    const dispatch = createEventDispatcher();

    export let level;
    export let count;
</script>

<style>
    .back {
        container-type: inline-size;
        border: 1px solid black;
        border-radius: 5px;
        box-shadow: 3px 2px 2px rgba(0, 0, 0, 0.3);
        display: grid;
        grid-template-rows: 10% 1fr 10%;
        justify-items: center;
        color: white;
        padding: 10px 0;
        aspect-ratio: 2 / 3;
    }

    .count {
        grid-row: 2;
        align-self: center;
        font-size: x-large;
    }

    .blobs {
        grid-row: 3;
        font-size: xx-small;
    }

    @container (max-width: 50px) {
        .blobs {
            font-size: 3px;
        }
    }

    .many {
        box-shadow: 3px 2px 1px rgba(0, 0, 0, 0.3), 2px 1px 1px rgba(0, 0, 0, 0.3);
    }

    .many-many {
        box-shadow: 5px 4px 1px rgba(0, 0, 0, 0.3), 3px 2px 1px rgba(0, 0, 0, 0.3), 2px 1px 1px rgba(0, 0, 0, 0.6);
    }

    .level1 {
        background-color: #5a7263;
    }

    .level2 {
        background-color: #706e5c;
    }

    .level3 {
        background-color: #545b6b;
    }

    .deck {
        cursor: pointer;
    }
</style>

<div class="back level{level}" id="deck{level}" class:deck={!!count} class:many={count && count > 2} class:many-many={count && count > 10} on:click={() => dispatch('click')}>
    {#if count}
        <span class="count">{count}</span>
    {/if}
    <span class="blobs">
        {#each Array(level) as _}
            ◯&nbsp;
        {/each}
    </span>
</div>