var array = [];

console.log(array);
fillArray();
console.log(quicksort(array));

function quicksort(array) {
  if (array.length <= 1) {
    return array;
  }

  var pivot = array[0];

  var left = [];
  var right = [];

  for (var i = 1; i < array.length; i++) {
    array[i] < pivot ? left.push(array[i]) : right.push(array[i]);
  }

  return quicksort(left).concat(pivot, quicksort(right));
}

function fillArray() {
  for (var i = 0; i < 50; i++) {
    array.push(randomIntFromInterval(0, 255));
  }
}

function randomIntFromInterval(min, max) {
  return Math.floor(Math.random() * (max - min + 1) + min);
}
