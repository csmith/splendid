<script>
    import _ from 'lodash';

    export let state;
    export let events = [];

    const pretty = (e) => {
        let player = '';
        if (e.playerId && state.players[e.playerId]) {
            player = state.players[e.playerId].details.name;
        }

        switch (e.event) {
            case 'add-bonus':
                return `${player} receives a ${e.type} bonus`;
            case 'add-player':
                return `${e.details.name} joins the game`;
            case 'add-points':
                return `${player} gains ${e.points} points`;
            case 'change-player':
                if (player) {
                    return `It is now ${player}'s turn`;
                } else {
                    return `It is no-one's turn`;
                }
            case 'change-phase':
                return `The game is now in the "${e.phase}" phase`;
            case 'discard-card':
                return `A card is removed from the board`;
            case 'discard-reserve':
                return `${player} buys a card from their reserve`;
            case 'final-round':
                return `This is the final round!`;
            case 'place-card':
                return `A card is dealt to the board`;
            case 'receive-noble':
                return `${player} receives a visit from a noble`;
            case 'remove-card-from-deck':
                return `A card is removed from the deck`;
            case 'reserve-card':
                return `${player} reserves a card`;
            case 'return-tokens':
                return `${player} returns tokens to the supply: ${_.map(e.tokens, (v, k) => `${v} ${k}`).join(', ')}`;
            case 'set-player-order':
                return `The turn order will be ${e.order.map((o) => state.players[o].details.name).join(', ')}`;
            case 'setup':
                return `The game has been configured for ${Object.values(state.players).length} players`;
            case 'take-tokens':
                return `${player} obtains tokens from the supply: ${_.map(e.tokens, (v, k) => `${v} ${k}`).join(', ')}`;
            default:
                return `Unknown event: ${JSON.stringify(e)}`
        }
    }
</script>

<section>
    <ul>
        {#each events.slice().reverse() as event}
            <li>{pretty(event)}</li>
        {/each}
    </ul>
</section>