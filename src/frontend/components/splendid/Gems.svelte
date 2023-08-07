<script>
    import _ from "lodash";
    import {createEventDispatcher} from "svelte";
    import Gem from "./Gem.svelte";

    const dispatch = createEventDispatcher();

    /** @type {import('../../../splendid/state.js').default} */
    export let state;

    export let canTake;

    let selected = {};
    let canSelect = {};
    let hasFullSelection = false;

    $: canSelect = _.mapValues(state.tokens, (v, k) => {
        const count = _.sum(Object.values(selected));
        const double = Object.values(selected).some(v => v > 1);
        const selectedCount = selected[k] || 0;

        return canTake
            // Can't ever select gold
            && k !== 'gold'
            // Can't select if there's none in stock
            && v - selectedCount > 0
            // Can't select if you've already picked a double
            && !double
            // Can't select if you've picked three
            && count < 3
            // Can't double up unless it's the only selection
            && (selectedCount === 0 || selectedCount === count)
            // Can't double up if there are less than 4 left
            && (selectedCount < 1 || v >= 4);
    });

    $: hasFullSelection = _.sum(Object.values(selected)) === 3 || Object.values(selected).some(v => v > 1);

    const selectGem = (type) => {
        if (canSelect[type]) {
            selected[type] = (selected[type] || 0) + 1;
        }
    };

    const unselectGem = (type) => {
        selected[type] = (selected[type] || 0) - 1;
    }

    const submitSelection = () => {
        dispatch('selected', selected);
        selected = {};
    };
</script>

<style>
    ul, li {
        list-style-type: none;
        margin: 0;
        padding: 0;
    }

    ul {
        display: flex;
        gap: 20px;
        align-items: center;
    }

    .holder {
        position: relative;
    }

    .amount {
        position: absolute;
        top: 2.1em;
        left: 2.1em;
        background-color: white;
        border: 2px solid black;
        border-radius: 100%;
        width: 1.5em;
        height: 1.5em;
        text-align: center;
    }

    .token {
        display: block;
        font-size: 1.5em;
        width: 1.5em;
        height: 1.5em;
        border-radius: 50%;
        text-align: center;
        padding: 0.3em;
        border: 2px solid black;
    }

    .selectable .token, .take .token {
        cursor: pointer;
        box-shadow: 0 0 5px red;
    }

    button {
        margin-top: 20px;
    }
</style>

<section>
    <h3>Gem supply</h3>
    <ul id="token-supply">
        {#each Object.entries(state.tokens) as pair}
            <li class="holder" class:selectable={canSelect[pair[0]]}
                on:click|preventDefault={() => selectGem(pair[0])}>
                <span class="token {pair[0]}">
                    <Gem type={pair[0]}/>
                </span>
                <span class="amount">{pair[1]}</span>
            </li>
        {/each}
    </ul>
    {#if _.sum(Object.values(selected)) > 0}
        <h3>You will take</h3>
        <ul id="token-selection">
            {#each Object.entries(selected) as pair}
                {#if pair[1] > 0}
                    <li class="holder take" on:click|preventDefault={() => unselectGem(pair[0])}>
                        <span class="token {pair[0]}">
                            <Gem type={pair[0]}/>
                        </span>
                        <span class="amount">{pair[1]}</span>
                    </li>
                {/if}
            {/each}
        </ul>
        <button on:click={submitSelection} disabled={!hasFullSelection}>Take</button>
    {/if}
</section>