<script>
  import _ from "lodash";
  import GemCounter from "./GemCounter.svelte";
  import CardBack from "./CardBack.svelte";
  import CardPlaceholder from "./CardPlaceholder.svelte";
  import Noble from "./Noble.svelte";

  /** @type {import('../../../games/splendid/state.js').default} */
  export let state;
</script>

<style>
  section {
    text-align: center;
  }

  section > ul > li {
    border: 1px solid var(--border);
    margin-bottom: 10px;
  }

  .tokens {
    display: flex;
    flex-direction: row;
    justify-content: center;
    gap: 8px;
    margin: 10px 0;
  }

  .discount {
    display: inline-block;
    border: 2px solid var(--border);
    width: 2.1em;
    height: 2.1em;
    line-height: 1;
    padding: 0.5em;
    text-align: center;
    font-weight: bold;
    margin: 5px;
    position: relative;
  }

  ul,
  li {
    margin: 0;
    padding: 0;
    list-style-type: none;
  }

  section > ul > li {
    padding: 5px;
  }

  h3 {
    margin: 0 0 5px 0;
  }

  .points {
    display: inline-block;
    position: relative;
    background-color: var(--background);
    border: 1px solid var(--border);
    width: 1.4em;
    height: 1.4em;
    padding: 0.15em;
    line-height: 1;
  }

  .reserve,
  .nobles {
    display: flex;
    flex-direction: row;
    justify-content: center;
    gap: 8px;
    margin: 5px 0;
    height: 40px;
  }
</style>

<section>
  <ul>
    {#each _.sortBy(Object.values(state.players), "order") as player}
      <li id="player-{player.details.id}">
        <h3>
          {#if state.turn === player.details.id}⏵{/if}
          {player.details.name} <span class="points">{player.points}</span>
        </h3>

        <h4>{_.sum(Object.values(player.tokens))} / 10 tokens</h4>
        <div class="tokens">
          {#each Object.entries(player.tokens) as pair}
            <GemCounter type={pair[0]} amount={pair[1]} size="small" />
          {/each}
        </div>
        <h4>{_.sum(Object.values(player.bonuses))} owned developments</h4>
        <div>
          {#each Object.entries(player.bonuses) as pair}
            <span class="discount {pair[0]}" title={pair[0]}>{pair[1]}</span>
          {/each}
        </div>
        <h4>{player.reserved.length} reserved cards</h4>
        <div class="reserve">
          {#each [0, 1, 2] as index}
            {#if player.reserved[index]}
              <CardBack level={player.reserved[index].level} />
            {:else}
              <CardPlaceholder />
            {/if}
          {/each}
        </div>
        <h4>{player.nobles.length} nobles received</h4>
        <div class="nobles">
          {#each player.nobles as noble}
            <Noble {noble} />
          {/each}
        </div>
      </li>
    {/each}
  </ul>
</section>
