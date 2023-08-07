<script>
    import {createEventDispatcher} from "svelte";
    import {canReceiveNoble} from "../../../splendid/util.js";
    import Noble from "./Noble.svelte";

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

<section>
    <h3>Wandering nobles</h3>
    {#each state.nobles as noble}
        <Noble {noble} {player} interactive={canSelect} on:click={() => receiveNoble(noble)} />
    {/each}
</section>