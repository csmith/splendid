import { Before } from "@cucumber/cucumber";

Before(function () {
  this.parseCosts = function (str) {
    return Object.fromEntries(
      ["emerald", "sapphire", "ruby", "diamond", "onyx"].map((type, index) => [type, parseInt(str[index])]),
    );
  };

  this.parseCard = function (str) {
    const parts = str.split("/");
    return {
      level: parseInt(parts[0]),
      points: parseInt(parts[1]),
      bonus: parts[2],
      cost: this.parseCosts(parts[3]),
      id: str,
    };
  };

  this.parseNoble = function (str) {
    return {
      cost: this.parseCosts(str),
      id: str,
    };
  };

  this.error = null;
});
