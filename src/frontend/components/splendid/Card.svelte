<script>
    import {canAffordCard} from "../../../splendid/util.js";
    import Gem from "./Gem.svelte";
    import {createEventDispatcher} from "svelte";

    const dispatch = createEventDispatcher();

    export let card;
    export let player;
</script>

<style>
    .level1 {
        --colour: #8fbda1;
    }

    .level2 {
        --colour: #bdb78f;
    }

    .level3 {
        --colour: #8f9cbd;
    }

    .card {
        display: grid;
        grid-template-columns: 1fr 1fr;
        grid-template-rows: 40px 1fr;
        grid-column-gap: 5px;
        padding: 0;
        background: linear-gradient(rgba(255, 255, 255, 0.4) 0%, rgba(255, 255, 255, 0.4) 40px, rgba(0, 0, 0, 0) 40px), var(--colour);
        border: 1px solid black;
        border-radius: 5px;
        box-shadow: 3px 2px 2px rgba(0, 0, 0, 0.3);
        cursor: pointer;
    }

    .score {
        padding-left: 10px;
        font-size: 25px;
        justify-self: start;
        align-self: center;
        color: white;
        font-weight: bold;
        text-shadow: 0 0 5px black,  0 0 2px black, 0 0 1px black;
    }

    .bonus {
        height: 30px;
        width: 30px;
        padding-right: 5px;
        justify-self: end;
        align-self: center;
        grid-column: 2;
    }

    .affordable {
        justify-self: end;
        align-self: end;
        font-size: 25px;
        padding: 2px 5px;
    }

    .costs {
        list-style: none;
        padding: 0;
        margin: 0 0 5px 0;
        display: grid;
        align-self: end;
        grid-template-columns: 36px;
        grid-auto-rows: 22px;
        grid-row-gap: 2px;
    }

    .costs li {
        display: contents;
    }

    .cost {
        display: grid;
        grid-template-columns: 12px auto;
        grid-column-gap: 6px;
        border-right: 1px solid black;
        border-top: 1px solid black;
        border-bottom: 1px solid black;
        border-bottom-right-radius: 4px;
        border-top-right-radius: 4px;
        margin: 0;
        padding: 0 5px;
        align-self: stretch;
        justify-self: stretch;
    }

    .icon {
        align-self: center;
    }

    .amount {
        justify-self: end;
        align-self: center;
        font-size: small;
        line-height: 1;
        color: black;
    }
</style>

<div id="card-{card.id}" class="card level{card.level}" on:click={() => dispatch('click')}>
    {#if card.points > 0}
        <span class="score" title="Buying this development will gain you {card.points} prestige points">{card.points}</span>
    {/if}
    <span class="bonus" title="Owning this development gives a bonus to: {card.bonus}"><Gem type={card.bonus}/></span>
    <ul class="costs">
        {#each Object.entries(card.cost) as entry}
            <li><span class="cost {entry[0]}">
                <span class="icon"><Gem type={entry[0]}/></span>
                <span class="amount">{entry[1]}</span>
            </span></li>
        {/each}
    </ul>
    {#if player && canAffordCard(player, card)}
        <span class="affordable" title="You can afford to buy this development">✅</span>
    {/if}
</div>