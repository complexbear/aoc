const fs = require('node:fs');
const _ = require('underscore');

const main = (inputName) => {

  const sum = (items) => _.reduce(items, (memo, i) => memo + i, 0);

  const calc_distances = (items) => {
      const [left, right] = items.map(list => list.map(Number).sort());
      const distances = left.map((l, i) => Math.abs(l - right[i]));
      console.log(inputName, "distance", sum(distances));
  };

  const calc_similarity = (items) => {
    // map of count of items in right list
    const [left, right] = items;
    const rightMap = {};
    right.forEach(item => {
      rightMap[item] = (rightMap[item] ?? 0) + 1;
    });
    const scores = left.map(item => item * (rightMap[item] ?? 0));
    console.log(inputName, "similarity", sum(scores));
  };

  fs.readFile(inputName, 'utf8', (err, data) => {
      if (err) {
        console.error(err);
        return;
      }
      const items =  _.unzip(data.split("\n").map(row => row.trim().split(/\s+/)));
      calc_distances(items);  
      calc_similarity(items);
    });
};

main("example.txt");
main("input.txt");

