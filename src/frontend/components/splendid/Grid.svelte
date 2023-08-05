<script>
    import {canAffordCard} from "../../../splendid/util.js";
    import {createEventDispatcher} from "svelte";

    const dispatch = createEventDispatcher();

    /** @type {import('../../../splendid/state.js').default} */
    export let state;

    export let player;
    export let canSelect = false;

    let buyDialog;
    let selectedCard = null;

    const levels = [2, 1, 0];

    const selectDeck = (level) => {
        if (canSelect) {
            dispatch('reserveFromDeck', level+1);
        }
    };

    const selectCard = (card) => {
        if (canSelect) {
            selectedCard = card;
            buyDialog.showModal();
        }
    };

    const buyCard = () => {
        if (canSelect) {
            buyDialog.close();
            dispatch('buyCard', selectedCard);
            selectedCard = null;
        }
    };

    const reserveCard = () => {
        if (canSelect) {
            buyDialog.close();
            dispatch('reserveCard', selectedCard);
            selectedCard = null;
        }
    };

    const handleDialogClick = (e) => {
        if (e.target === e.currentTarget) {
            buyDialog.close();
        }
    };

    const handleBuyDialogClose = () => {
        selectedCard = null;
    };
</script>

<style>
    section {
        display: grid;
        grid-template-columns: repeat(5, 100px);
        grid-gap: 10px;
    }

    .canSelect .card {
        cursor: pointer;
        box-shadow: 0 0 5px red;
    }

    .card, .placeholder {
        width: 100px;
        height: 150px;
        border: 1px solid black;
        position: relative;
    }

    .card.level0 {
        background-color: #8fbda1;
    }

    .card.level1 {
        background-color: #bdb78f;
    }

    .card.level2 {
        background-color: #8f9cbd;
    }

    .back.level0 {
        background-color: #5a7263;
    }

    .back.level1 {
        background-color: #706e5c;
    }

    .back.level2 {
        background-color: #545b6b;
    }

    .back {
        text-align: center;
        color: white;
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

    .bonus {
        position: absolute;
        top: 0;
        right: 0;
        color: white;
        padding: 2px 2px 2px 4px;
        border-bottom: 1px solid black;
        border-left: 1px solid black;
        border-bottom-left-radius: 5px;
    }

    .costs {
        position: absolute;
        bottom: 0;
        left: 0;
    }

    .affordable {
        position: absolute;
        bottom: 0;
        right: 0;
        font-size: 1.5em;
        line-height: 1;
        padding: 5px;
    }

    .cost {
        color: white;
        padding: 2px 4px;
        display: block;
        border-top: 1px solid black;
        border-right: 1px solid black;
    }

    .header {
        color: white;
    }

    ul, li {
        margin: 0;
        padding: 0;
        list-style-type: none;
    }

    dialog > div {
        display: grid;
        grid-template-columns: auto 1fr;
        grid-gap: 20px;
    }

    dialog h2, dialog nav {
        grid-column: 1 / span 2;
        text-align: center;
        margin: 0;
    }

    nav {
        display: flex;
        justify-content: space-around;
    }

    table {
        border-collapse: collapse;
    }

    td, th {
        border: 1px solid black;
        padding: 2px;
        text-align: center;
    }

    th {
        font-weight: bold;
    }

    td.satisfied {
        background-color: lightgreen;
    }

    .tokens {
        grid-column: 1 / span 2;
    }

    .reserve {
        grid-column: 3 / span 3;
    }
</style>

<section class:canSelect="{canSelect}">
    {#each levels as level}
        {#if state.decks[level].length > 0}
            <div class="card back level{level}" id="deck{level}" on:click={() => selectDeck(level)}>
                {state.decks[level].length}
            </div>
        {:else}
            <div class="placeholder"></div>
        {/if}
        {#each state.cards[level] as card, i}
            {#if card}
                <div id="card-{level}-{i}" class="card level{level}" on:click={() => selectCard(card)}>
                    {#if card.points > 0}
                        <span class="score">{card.points}</span>
                    {/if}
                    <span class="bonus {card.bonus}">{card.bonus}</span>
                    <ul class="costs">
                        {#each Object.entries(card.cost) as entry}
                            <li><span class="cost {entry[0]}">{entry[1]}</span></li>
                        {/each}
                    </ul>
                    {#if canAffordCard(player, card)}
                        <span class="affordable">✅</span>
                    {/if}
                </div>
            {:else}
                <div class="placeholder" id="placeholder-{level}-{i}"></div>
            {/if}
        {/each}
    {/each}
    {#if player}
        <div class="tokens">Your tokens</div>
        <div class="reserve">Your reserve</div>
        <div class="tokens">
            Look in the sidebar for now 🙃
        </div>
        {#each [0, 1, 2] as index}
            {#if player.reserved.length > index}
                <div class="card level{player.reserved[index].level-1}"
                     on:click={() => selectCard(player.reserved[index])}>
                    {#if player.reserved[index].points > 0}
                        <span class="score">{player.reserved[index].points}</span>
                    {/if}
                    <span class="bonus {player.reserved[index].bonus}">{player.reserved[index].bonus}</span>
                    <ul class="costs">
                        {#each Object.entries(player.reserved[index].cost) as entry}
                            <li><span class="cost {entry[0]}">{entry[1]}</span></li>
                        {/each}
                    </ul>
                    {#if canAffordCard(player, player.reserved[index])}
                        <span class="affordable">✅</span>
                    {/if}
                </div>
            {:else}
                <div class="placeholder" id="reserve-{index}"></div>
            {/if}
        {/each}
    {/if}
</section>

<dialog bind:this={buyDialog} on:click={handleDialogClick} on:close={handleBuyDialogClose}>
    {#if selectedCard}
        <div>
            <h2>Selected card</h2>
            <div class="card level{selectedCard.level-1}" on:click={() => selectCard(selectedCard)}>
                {#if selectedCard.points > 0}
                    <span class="score">{selectedCard.points}</span>
                {/if}
                <span class="bonus {selectedCard.bonus}">{selectedCard.bonus}</span>
                <ul class="costs">
                    {#each Object.entries(selectedCard.cost) as entry}
                        <li><span class="cost {entry[0]}">{entry[1]}</span></li>
                    {/each}
                </ul>
                {#if canAffordCard(player, selectedCard)}
                    <span class="affordable">✅</span>
                {/if}
            </div>
            <table>
                <tr>
                    <th>Gem</th>
                    {#each Object.keys(selectedCard.cost) as key}
                        <th class="header {key}">{key}</th>
                    {/each}
                </tr>
                <tr>
                    <th>Cost</th>
                    {#each Object.values(selectedCard.cost) as value}
                        <td>{value}</td>
                    {/each}
                </tr>
                <tr>
                    <th>Bonuses</th>
                    {#each Object.keys(selectedCard.cost) as key}
                        <td>{-player.bonuses[key]}</td>
                    {/each}
                </tr>
                <tr>
                    <th>Your gems</th>
                    {#each Object.entries(selectedCard.cost) as pair}
                        <td>{-Math.min(Math.max(0, pair[1] - player.bonuses[pair[0]]), player.tokens[pair[0]])}</td>
                    {/each}
                </tr>
                <tr>
                    <th>Remaining</th>
                    {#each Object.entries(selectedCard.cost) as pair}
                        <td class:satisfied={Math.max(0, pair[1] - player.bonuses[pair[0]] - player.tokens[pair[0]]) === 0}>
                            {Math.max(0, pair[1] - player.bonuses[pair[0]] - player.tokens[pair[0]])}
                        </td>
                    {/each}
                </tr>
            </table>
            <nav>
                {#if canAffordCard(player, selectedCard)}
                    <button on:click={buyCard}>Buy card</button>
                {/if}
                {#if player.reserved.length < 3}
                    <button on:click={reserveCard}>Reserve card</button>
                {/if}
            </nav>
        </div>
    {/if}
</dialog>