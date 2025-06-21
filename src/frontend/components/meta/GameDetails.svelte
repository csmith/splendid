<script>
  import { createEventDispatcher } from 'svelte';
  
  export let game;
  
  const dispatch = createEventDispatcher();
  
  let selectedOptions = {};
  
  // Initialize selected options with defaults
  if (game.options && game.options.length > 0) {
    game.options.forEach(option => {
      selectedOptions[option.key] = option.default;
    });
  }
  
  const handleStartGame = () => {
    dispatch('startGame', {
      name: game.name,
      options: selectedOptions
    });
  };
</script>

<style>
  section {
    background-color: var(--background-contrast);
    border-radius: 8px;
    padding: 1.5rem;
    margin: 1rem 0;
  }
  
  h3 {
    margin: 0 0 0.8em 0;
  }
  
  p {
    margin: 0.8em 0;
  }
  
  details {
    margin: 1rem 0;
  }
  
  details summary {
    cursor: pointer;
    font-weight: bold;
  }
  
  .option-group {
    margin: 0.5rem 1rem;
  }
  
  .option-group h4 {
    margin: 0 0 0.5rem 0;
  }
  
  .option-group > p {
    margin: 0.5rem 0 1rem 0;
    opacity: 0.8;
  }
  
  .radio-option {
    margin: 0.75rem 0;
  }
  
  .radio-option label {
    display: flex;
    align-items: flex-start;
    gap: 0.5rem;
  }
  
  .radio-option input[type="radio"] {
    margin-top: 0.2rem;
    flex-shrink: 0;
  }
  
  .radio-option-content {
    flex: 1;
  }
  
  .radio-option-content .option-description {
    font-size: 0.9em;
    opacity: 0.8;
    margin: 0.2rem 0 0 0;
  }
  
  button {
    font-size: 1.1em;
    padding: 0.75rem 1.25rem;
  }
</style>

<section>
  <h3>{game.name}</h3>
  <p>{game.description}</p>
  <p class="stats">
    Players: {game.players.min}&ndash;{game.players.max}. Based on:
    <a href={game.based_on.link}>{game.based_on.game} by {game.based_on.creator}</a>. If you enjoy this game, please
    consider <a href={game.based_on.purchase}>purchasing a physical copy</a>.
  </p>
  
  {#if game.options && game.options.length > 0}
    <details>
      <summary>Game options</summary>
      {#each game.options as option}
        <div class="option-group">
          <h4>{option.name}</h4>
          <p>{option.description}</p>
          {#each option.options as choice}
            <div class="radio-option">
              <label>
                <input 
                  type="radio" 
                  name="{game.name}-{option.key}"
                  value={choice.value}
                  bind:group={selectedOptions[option.key]}
                />
                <div class="radio-option-content">
                  <span>{choice.name}</span>
                  <p class="option-description">{choice.description}</p>
                </div>
              </label>
            </div>
          {/each}
        </div>
      {/each}
    </details>
  {/if}
  
  <button on:click={handleStartGame}>Start a new game of {game.name}</button>
</section>