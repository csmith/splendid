<script>
    import _ from "lodash";
    import {createEventDispatcher, onMount} from "svelte";

    const dispatch = createEventDispatcher();

    export let player;

    let dialog;
    let selection = {};
    $: selectedEnough = _.sum(Object.values(player.tokens)) - _.sum(Object.values(selection)) === 10;

    onMount(() => {
        dialog.showModal();
    });

    const selectToken = (token) => {
        const newValue = (selection[token] || 0) + 1;
        if (newValue <= player.tokens[token]) {
            selection[token] = newValue;
        }
    };

    const deselectToken = (token) => {
        const newValue = (selection[token] || 0) - 1;
        if (newValue >= 0) {
            selection[token] = newValue;
        }
    };

    const handleSubmit = () => {
        dialog.close();
        dispatch('tokensDiscarded', selection);
    };
</script>

<style>
    dialog {
        text-align: center;
    }

    ul, li {
        margin: 0;
        padding: 0;
        list-style-type: none;
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
        cursor: pointer;
    }
</style>

<dialog bind:this={dialog}>
    <p>You have too many tokens! Select which to discard:</p>
    <ul class="tokens">
        {#each Object.entries(player.tokens) as pair}
            <li class="token {pair[0]}" on:click={() => selectToken(pair[0])}>{pair[1] - (selection[pair[0]] || 0)}</li>
        {/each}
    </ul>
    {#if _.sum(Object.values(selection)) > 0}
        <p>You will discard:</p>
        <ul class="tokens">
            {#each Object.entries(selection) as pair}
                {#if pair[1]}
                <li class="token {pair[0]}"
                    on:click={() => deselectToken(pair[0])}>{pair[1]}</li>
                {/if}
            {/each}
        </ul>
    {/if}
    {#if selectedEnough}
        <button on:click={handleSubmit}>Discard</button>
    {/if}
</dialog>