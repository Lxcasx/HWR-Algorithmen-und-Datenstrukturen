var array = [];

// fill array
for (var i = 0; i < 50; i++) {
  array.push(randomIntFromInterval(0, 255));
}

console.log(array);

// sort array
for (var i = 0; i < array.length; i++) {
  var min = i;

  for (var k = i + 1; k < array.length; k++) {
    if (array[k] < array[min]) {
      min = k;
    }
  }

  if (min != i) {
    var tmp = array[i];
    array[i] = array[min];
    array[min] = tmp;
  }
}

console.log(array);

function randomIntFromInterval(min, max) {
  return Math.floor(Math.random() * (max - min + 1) + min);
}
