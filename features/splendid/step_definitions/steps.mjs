import {Before, Given, Then, When} from "@cucumber/cucumber";
import state from "../../../src/splendid/state.js";
import Engine from "../../../src/common/engine.js";
import phases from "../../../src/splendid/phases.js";
import {newPlayer} from "../../../src/common/player.js";
import {findPlayerByName} from "../../../src/common/state.js";
import {replaceNth} from "../../../src/common/util.js";
import _ from "lodash";
import assert from "assert";

Before(function () {
    this.setState = function (state) {
        this.engine = new Engine(state, phases, (e) => e, 'Splendid')
    };

    this.playerState = function (name) {
        return findPlayerByName(this.engine.state, name);
    }

    this.playerDetails = function (name) {
        return this.playerState(name).details;
    }

    this.perform = function (name, playerName, args) {
        try {
            this.engine.perform(name, this.playerDetails(playerName), args);
            this.error = null;
        } catch (e) {
            this.error = e;
        }
    }

    this.parseCosts = function (str) {
        return Object.fromEntries(
            ['emerald', 'sapphire', 'ruby', 'diamond', 'onyx'].map((type, index) => [type, parseInt(str[index])])
        )
    }

    this.parseCard = function (str) {
        const parts = str.split('/');
        return {
            level: parseInt(parts[0]),
            points: parseInt(parts[1]),
            bonus: parts[2],
            cost: this.parseCosts(parts[3]),
        }
    }

    this.parseNoble = function (str) {
        return {
            cost: this.parseCosts(str)
        }
    }

    this.error = null;

    this.setState(state);
});

Given(/^the following players joined the game:$/, function (dataTable) {
    dataTable.rows().forEach(row => {
        this.engine.perform('join', newPlayer(row[0]))
    })
});

Given(/^the game was started by (.*?)$/, function (playerName) {
    this.perform('start', playerName);
});

Given(/^it was (.*?)'s turn$/, function (playerName) {
    this.setState({
        ...this.engine.state,
        turn: this.playerDetails(playerName).id
    })
});

Given(/^the following tokens were available:$/, function (dataTable) {
    let tokens = _.mapValues(this.engine.state.tokens, () => 0);
    dataTable.hashes().forEach(row => {
        tokens[row.type] = parseInt(row.amount);
    });

    this.setState({
        ...this.engine.state,
        tokens: tokens,
    })
});

Given(/^(.*?) had the following tokens:$/, function (playerName, dataTable) {
    let tokens = _.mapValues(this.engine.state.tokens, () => 0);
    dataTable.hashes().forEach(row => {
        tokens[row.type] = parseInt(row.amount);
    });

    const playerState = this.playerState(playerName);

    this.setState({
        ...this.engine.state,
        players: {
            ...this.engine.state.players,
            [playerState.details.id]: {
                ...playerState,
                tokens: tokens
            }
        }
    })
});

Given(/^(.*?) had the following bonuses:$/, function (playerName, dataTable) {
    let bonuses = _.mapValues(this.engine.state.tokens, () => 0);
    dataTable.hashes().forEach(row => {
        bonuses[row.type] = parseInt(row.amount);
    });

    const playerState = this.playerState(playerName);

    this.setState({
        ...this.engine.state,
        players: {
            ...this.engine.state.players,
            [playerState.details.id]: {
                ...playerState,
                bonuses
            }
        }
    })
});


Given(/^the following cards were visible:$/, function (dataTable) {
    const cards = _.reverse(dataTable.raw().map(row => row.map(card => this.parseCard(card))));

    this.setState({
        ...this.engine.state,
        cards: cards,
    })
});

Given(/^the top card of deck (\d+) was (.*?)$/, function (deck, card) {
    this.setState({
        ...this.engine.state,
        decks: replaceNth(this.engine.state.decks, deck - 1, () => [this.parseCard(card)])
    })
});

Given(/^(.*?) had (\d+) points?$/, function (playerName, points) {
    const playerState = this.playerState(playerName);
    this.setState({
        ...this.engine.state,
        players: {
            ...this.engine.state.players,
            [playerState.details.id]: {
                ...playerState,
                points
            }
        }
    })
});

Given(/^the turn order was:$/, function (dataTable) {
    const order = dataTable.raw().map((row) => row[0]);
    this.setState({
        ...this.engine.state,
        players: _.mapValues(this.engine.state.players, (v) => ({
            ...v,
            order: order.indexOf(v.details.name)
        }))
    })
});

Given(/^this was the final round$/, function () {
    this.setState({
        ...this.engine.state,
        finalRound: true,
    })
});

Given(/^(.*?) had the following reserved cards:$/, function (playerName, dataTable) {
    const cards = dataTable.raw().map(r => this.parseCard(r[0]));
    const playerState = this.playerState(playerName);
    this.setState({
        ...this.engine.state,
        players: {
            ...this.engine.state.players,
            [playerState.details.id]: {
                ...playerState,
                reserved: cards,
            }
        }
    })
});

Given(/the following nobles were available:$/, function (dataTable) {
    this.setState({
        ...this.engine.state,
        nobles: dataTable.raw().map((r) => this.parseNoble(r[0])),
    })
});

Given(/^the game phase was (.*?)$/, function (phase) {
    this.setState({
        ...this.engine.state,
        phase,
    });
})

When(/^(.*?) draws the following tokens:$/, function (playerName, dataTable) {
    this.perform(
        'take-tokens',
        playerName,
        {tokens: Object.fromEntries(dataTable.hashes().map(row => [row.type, parseInt(row.amount)]))}
    );
});

When(/^(.*?) buys the card (.*?)$/, function (playerName, card) {
    this.perform(
        'buy-card',
        playerName,
        {card: this.parseCard(card)}
    )
});

When(/^(.*?) reserves the card (.*?)$/, function (playerName, card) {
    this.perform(
        'reserve-card',
        playerName,
        {card: this.parseCard(card)}
    )
});

When(/^(.*?) receives the noble (.*?)$/, function (playerName, noble) {
    this.perform(
        'receive-noble',
        playerName,
        {noble: this.parseNoble(noble)}
    )
});

Then(/^(.*?) will have the following bonuses:$/, function (playerName, dataTable) {
    const bonuses = this.playerState(playerName).bonuses;

    dataTable.hashes().forEach(row => {
        assert.equal(bonuses[row.type], parseInt(row.amount));
    });
});

Then(/^(.*?) will have the following tokens:$/, function (playerName, dataTable) {
    const tokens = this.playerState(playerName).tokens;

    dataTable.hashes().forEach(row => {
        assert.equal(tokens[row.type], parseInt(row.amount));
    });
});

Then(/^the following tokens will be available:$/, function (dataTable) {
    const tokens = this.engine.state.tokens;

    dataTable.hashes().forEach(row => {
        assert.equal(tokens[row.type], parseInt(row.amount), row.type);
    });
});

Then(/^an? "(.*?)" error will occur$/, function (message) {
    assert.equal(message, this.error?.message);
});

Then(/^the game phase will be (.*?)$/, function (phase) {
    assert.equal(this.engine.state.phase, phase);
});

Then(/^it will be (.*?)'s turn(?: still)?$/, function (playerName) {
    const details = this.playerDetails(playerName);
    assert.equal(this.engine.currentPlayer, details.id);
});

Then(/^there will be (\d+) nobles available?$/, function (count) {
    assert.equal(this.engine.state.nobles.length, count);
});

Then(/^the card in row (\d+) column (\d+) will be (.*?)$/, function (row, column, card) {
    const actual = this.engine.state.cards[3 - row][column - 1];
    const expected = this.parseCard(card);
    assert.ok(_.isEqual(actual, expected));
});

Then(/^(.*?) will have (\d+) points?$/, function (playerName, score) {
    const playerState = this.playerState(playerName);
    assert.equal(playerState.points, score);
});

Then(/^this will be the final round$/, function () {
    assert.ok(this.engine.state.finalRound);
});

Then(/^(.*?) will have the following reserved cards:$/, function (playerName, dataTable) {
    const actual = this.playerState(playerName).reserved;
    const expected = dataTable.raw().map((row) => this.parseCard(row[0]));
    assert.ok(_.isEqual(actual, expected));
});