<script>
    import _ from "lodash";
    import {createEventDispatcher} from "svelte";
    import Gem from "./Gem.svelte";
    import GemCounter from "./GemCounter.svelte";

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
        margin: 15px 0;
    }
</style>

<section>
    <h3>Gem supply</h3>
    <ul id="token-supply">
        {#each Object.entries(state.tokens) as pair}
            <li>
                <GemCounter type={pair[0]} amount={pair[1]} interactive={canSelect[pair[0]]} on:click={() => selectGem(pair[0])}/>
            </li>
        {/each}
    </ul>
    {#if _.sum(Object.values(selected)) > 0}
        <h3>You will take</h3>
        <ul id="token-selection">
            {#each Object.entries(selected) as pair}
                <li>
                    <GemCounter type={pair[0]} amount={pair[1]} interactive={true} on:click={() => unselectGem(pair[0])}/>
                </li>
            {/each}
        </ul>
        <button on:click={submitSelection} disabled={!hasFullSelection}>Take</button>
    {/if}
</section>