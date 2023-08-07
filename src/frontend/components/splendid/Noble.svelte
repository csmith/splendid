<script>
    import {canReceiveNoble} from "../../../splendid/util.js";
    import {createEventDispatcher} from "svelte";

    const dispatch = createEventDispatcher();

    export let noble;
    export let player;
    export let interactive = false;
</script>


<style>
    .noble {
        container-type: inline-size;
        display: inline-block;
        height: 150px;
        max-height: 100%;
        max-width: 100%;
        aspect-ratio: 23/29;
        border: 1px solid var(--border);
        position: relative;
        background: linear-gradient(90deg, rgba(255, 255, 255, 0.6) 0%, rgba(255, 255, 255, 0.6) 25%, rgba(0, 0, 0, 0) 25%), var(--image), rebeccapurple;
        background-size: cover;
    }

    .score {
        position: absolute;
        top: 0;
        left: 0;
        font-size: 2em;
        line-height: 1;
        padding: 5px;
        color: white;
        text-shadow: 0 0 3px black, 0 0 1px black;
    }

    ul, li {
        margin: 0;
        padding: 0;
        list-style-type: none;
    }

    .costs {
        display: flex;
        flex-direction: column;
        gap: 2px;
        width: 30px;
        align-items: center;
        position: absolute;
        bottom: 5px;
        left: 0;
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
        border: 1px solid black;
    }

    .selectable {
        cursor: pointer;
        box-shadow: 0 0 5px red;
    }

    @container (max-width: 50px) {
        .score, .costs {
            display: none;
        }
    }
</style>

<div
        class="noble"
        id="noble-{noble.id}"
        class:selectable={interactive}
        style="--image: url('/splendid/nobles/{noble.id}.avif');"
        on:click={() => dispatch('click')}>
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