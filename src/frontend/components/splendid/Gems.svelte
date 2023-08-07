<script>
    import _ from "lodash";
    import {createEventDispatcher} from "svelte";
    import Gem from "./Gem.svelte";
    import GemCounter from "./GemCounter.svelte";

    const dispatch = createEventDispatcher();

    /** @type {import('../../../splendid/state.js').default} */
    export let state;

    export let canTake;

    const noSelectedGems = {
        emerald: 0,
        diamond: 0,
        sapphire: 0,
        onyx: 0,
        ruby: 0,
    };

    let selected = _.cloneDeep(noSelectedGems);
    let canSelect = {};
    let hasFullSelection = false;

    $: canSelect = _.mapValues(state.tokens, (v, k) => {
        const count = _.sum(Object.values(selected));
        const double = Object.values(selected).some(v => v > 1);
        const selectedCount = selected[k];

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

    $: availableGems = _.sum(Object.keys(noSelectedGems).map(k => state.tokens[k]));

    $: hasFullSelection = _.sum(Object.values(selected)) === Math.min(availableGems, 3)
        || Object.values(selected).some(v => v > 1);

    const selectGem = (type) => {
        if (canSelect[type]) {
            selected[type]++;
        }
    };

    const unselectGem = (type) => {
        if (selected[type] > 0) {
            selected[type]--;
        }
    }

    const submitSelection = () => {
        dispatch('selected', selected);
        selected = _.cloneDeep(noSelectedGems);
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
        margin: 15px 0 30px 0;
    }

    .hidden {
        visibility: hidden;
    }
</style>

<section>
    <h3>Gem supply</h3>
    <ul id="token-supply">
        {#each Object.entries(state.tokens) as pair}
            <li>
                <GemCounter type={pair[0]} amount={pair[1]} interactive={canSelect[pair[0]]}
                            on:click={() => selectGem(pair[0])}/>
            </li>
        {/each}
    </ul>
    <div class="take" class:hidden={_.sum(Object.values(selected)) === 0}>
        <h3>You will take:</h3>
        <ul id="token-selection">
            {#each Object.entries(selected) as pair}
                <li>
                    <GemCounter type={pair[0]} amount={pair[1]} interactive={pair[1] > 0}
                                on:click={() => unselectGem(pair[0])}/>
                </li>
            {/each}
        </ul>
        <button on:click={submitSelection} disabled={!hasFullSelection}>Take</button>
    </div>
</section>