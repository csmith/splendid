<script>
    import {createEventDispatcher} from "svelte";
    import anime from "animejs";
    import _ from "lodash";

    export let state;
    export let playerId;
    export let nextEvent = undefined;

    const dispatch = createEventDispatcher();

    const wait = async (ms) => {
        await new Promise(resolve => setTimeout(resolve, ms));
    }

    const getOffset = (el) => {
        const rect = el.getBoundingClientRect();
        return {
            left: rect.left + window.scrollX,
            top: rect.top + window.scrollY
        };
    }

    const animateTokens = async (startContainer, endContainer, tokens) => {
        await Promise.all(Object.entries(tokens).filter(([, amount]) => amount > 0).map(([type, amount]) => {
            const source = document.querySelector(`${startContainer} .token.${type}`);
            const sourceOffset = getOffset(source);
            const target = document.querySelector(`${endContainer} .token.${type}`);
            const targetOffset = getOffset(target);
            const copy = source.cloneNode(true);
            copy.innerText = amount;
            copy.style.position = 'absolute';
            copy.style.top = `${sourceOffset.top}px`;
            copy.style.left = `${sourceOffset.left}px`;
            document.body.appendChild(copy);

            return anime({
                targets: copy,
                translateX: targetOffset.left - sourceOffset.left,
                translateY: targetOffset.top - sourceOffset.top,
                easing: 'easeInOutQuad',
                complete: () => document.body.removeChild(copy),
            }).finished;
        }))
    }

    const animateCard = async (card, endContainer) => {
        const index = _.findIndex(state.cards[card.level-1], (c) => _.isEqual(c, card));
        if (index === -1) {
            return;
        }

        const source = document.querySelector(`#card-${card.level-1}-${index}`)
        const sourceOffset = getOffset(source);
        const target = document.querySelector(endContainer);
        const targetOffset = getOffset(target);
        console.log(source, `#card-${card.level-1}-${index}`, target, endContainer);
        const copy = source.cloneNode(true);
        copy.style.position = 'absolute';
        copy.style.top = `${sourceOffset.top}px`;
        copy.style.left = `${sourceOffset.left}px`;
        document.body.appendChild(copy);
        source.style.visibility = 'hidden';

        return anime({
            targets: copy,
            translateX: targetOffset.left - sourceOffset.left,
            translateY: targetOffset.top - sourceOffset.top,
            easing: 'easeInOutQuad',
            complete: () => document.body.removeChild(copy),
        }).finished;
    }

    const animateDiscount = async (player, type) => {
        const el = document.querySelector(`#player-${player} .discount.${type}`);
        return anime({
            targets: el,
            scale: [
                {value: 2, duration: 400},
                {value: 2, duration: 200},
                {value: 1, duration: 400},
            ],
            innerText: [
                {value: el.innerText, duration: 500},
                {value: `${parseInt(el.innerText) + 1}`, duration: 1},
            ],
            duration: 1000,
        }).finished;
    }

    const animatePoints = async (player, amount) => {
        const el = document.querySelector(`#player-${player} .points`);
        return anime({
            targets: el,
            scale: [
                {value: 2, duration: 400},
                {value: 2, duration: 200},
                {value: 1, duration: 400},
            ],
            innerText: [
                {value: el.innerText, duration: 500},
                {value: `${parseInt(el.innerText) + amount}`, duration: 1},
            ],
            duration: 1000,
        }).finished;
    }

    const process = async (e) => {
        console.log('Processing animation for ', e)
        switch (e.event) {
            case 'take-tokens':
                await animateTokens('#token-supply', `#player-${e.playerId}`, e.tokens);
                break;
            case 'return-tokens':
                await animateTokens(`#player-${e.playerId}`, '#token-supply', e.tokens);
                break;
            case 'discard-card':
                await animateCard(e.card, `#player-${e.playerId}`);
                break;
            case 'add-bonus':
                await animateDiscount(e.playerId, e.type);
                break;
            case 'add-points':
                await animatePoints(e.playerId, e.points);
                break;
            default:
                await wait(250);
        }
    }

    $: if (nextEvent) {
        process(nextEvent)
            .then(() => setTimeout(() => dispatch('eventProcessed'), 1));
    }
</script>