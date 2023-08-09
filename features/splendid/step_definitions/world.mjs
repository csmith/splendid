import { Before } from "@cucumber/cucumber";

Before(function () {
  this.parseSplendidCosts = function (str) {
    return Object.fromEntries(
      ["emerald", "sapphire", "ruby", "diamond", "onyx"].map((type, index) => [type, parseInt(str[index])]),
    );
  };

  this.parseSplendidCard = function (str) {
    const parts = str.split("/");
    return {
      level: parseInt(parts[0]),
      points: parseInt(parts[1]),
      bonus: parts[2],
      cost: this.parseSplendidCosts(parts[3]),
      id: str,
    };
  };

  this.parseSplendidNoble = function (str) {
    return {
      cost: this.parseSplendidCosts(str),
      id: str,
    };
  };
});
