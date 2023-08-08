<script>
  import { createEventDispatcher } from "svelte";
  import anime from "animejs";

  export let state;
  export let playerId;
  export let nextEvent = undefined;

  const dispatch = createEventDispatcher();

  const wait = async (ms) => {
    await new Promise((resolve) => setTimeout(resolve, ms));
  };

  const getOffset = (el) => {
    const rect = el.getBoundingClientRect();
    return {
      left: rect.left + window.scrollX,
      top: rect.top + window.scrollY,
    };
  };

  /**
   * Creates a copy of the element, then animates it over to the target
   * element and removes it.
   *
   * @param {boolean?} opts.hideSource If truthy, hide the source element
   * @param {string?} opts.innerText If set, replace the inner text of the copy
   * @param {number?} opts.duration If set, the duration of the animation
   */
  const copyAndMoveElement = async (source, target, opts) => {
    const sourceOffset = getOffset(source);
    const targetOffset = getOffset(target);
    const copy = source.cloneNode(true);
    copy.style.position = "absolute";
    copy.style.top = `${sourceOffset.top}px`;
    copy.style.left = `${sourceOffset.left}px`;
    copy.style.width = `${source.offsetWidth}px`;
    copy.style.height = `${source.offsetHeight}px`;
    copy.style.zIndex = 1000;
    if (opts.innerText) {
      copy.innerText = opts.innerText;
    }
    document.body.appendChild(copy);
    if (opts.hideSource) {
      source.style.visibility = "hidden";
    }

    await anime({
      targets: copy,
      translateX: targetOffset.left - sourceOffset.left,
      translateY: targetOffset.top - sourceOffset.top,
      easing: "easeInOutQuad",
      complete: () => {
        document.body.removeChild(copy);
        if (opts.hideSource) {
          // Make sure it becomes visible again, if nothing else changes it
          setTimeout(() => {
            source.style.visibility = "visible";
          }, 100);
        }
      },
      duration: opts.duration || 500,
    }).finished;
  };

  const animateTokens = async (startContainer, endContainer, tokens) => {
    await Promise.all(
      Object.entries(tokens)
        .filter(([, amount]) => amount > 0)
        .map(([type, amount]) => {
          const source = document.querySelector(`${startContainer} .token.${type}`);
          const target = document.querySelector(`${endContainer} .token.${type}`);
          return copyAndMoveElement(source, target, { innerText: amount });
        }),
    );
  };

  const animateCard = async (card, endContainer) => {
    const source = document.querySelector(`#card-${card.id}`);
    const target = document.querySelector(endContainer);
    return copyAndMoveElement(source, target, { hideSource: true });
  };

  const animateDeal = async (e) => {
    const source = document.querySelector(`#deck${e.level}`);
    const target =
      e.reason === "reserve"
        ? e.playerId === playerId
          ? document.querySelector(`#reserve-${state.players[e.playerId].reserved.length}`)
          : document.querySelector(`#player-${e.playerId}`)
        : document.querySelector(`.placeholder.level${e.level - 1}`);

    return copyAndMoveElement(source, target, { innerText: "", duration: 100 });
  };

  /**
   * Scales an element up, then replaces the text with the new value, before
   * scaling it back down.
   */
  const highlightChange = async (el, newValue) => {
    let updated = false;
    const originZIndex = el.style.zIndex;
    await anime({
      targets: el,
      scale: [
        { value: 2, duration: 400 },
        { value: 2, duration: 200 },
        { value: 1, duration: 400 },
      ],
      duration: 1000,
      begin: () => (el.style.zIndex = 1000),
      complete: () => (el.style.zIndex = originZIndex),
      update: (anim) => {
        if (anim.progress > 40 && !updated) {
          updated = true;
          el.innerText = newValue;
        }
      },
    }).finished;
  };

  const animateDiscount = async (player, type) => {
    const el = document.querySelector(`#player-${player} .discount.${type}`);
    await highlightChange(el, parseInt(el.innerText) + 1);
  };

  const animatePoints = async (player, amount) => {
    const el = document.querySelector(`#player-${player} .points`);
    await highlightChange(el, parseInt(el.innerText) + amount);
  };

  const process = async (e) => {
    console.log("Processing animation for ", e);
    switch (e.event) {
      case "take-tokens":
        await animateTokens("#token-supply", `#player-${e.playerId}`, e.tokens);
        break;
      case "return-tokens":
        await animateTokens(`#player-${e.playerId}`, "#token-supply", e.tokens);
        break;
      case "discard-card":
        if (e.reason === "reserve" && e.playerId === playerId) {
          await animateCard(e.card, `#reserve-${state.players[e.playerId].reserved.length}`);
        } else {
          await animateCard(e.card, `#player-${e.playerId}`);
        }
        break;
      case "add-bonus":
        await animateDiscount(e.playerId, e.type);
        break;
      case "add-points":
        await animatePoints(e.playerId, e.points);
        break;
      case "remove-card-from-deck":
        await animateDeal(e);
        break;
      case "discard-reserve":
        if (e.playerId === playerId) {
          await copyAndMoveElement(
            document.querySelector(`#card-${e.card.id}`),
            document.querySelector(`#player-${e.playerId}`),
            { hideSource: true },
          );
        }
        break;
      case "receive-noble":
        await copyAndMoveElement(
          document.querySelector(`#noble-${e.noble.id}`),
          document.querySelector(`#player-${e.playerId}`),
          { hideSource: true },
        );
        break;
      case "place-card":
        // Do nothing – we've already animated the deal
        break;
      default:
        await wait(250);
    }
  };

  $: if (nextEvent) {
    process(nextEvent).then(() => setTimeout(() => dispatch("eventProcessed"), 1));
  }
</script>
