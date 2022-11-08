var array = [];

// fill array
for (var i = 0; i < 50; i++) {
  array.push(randomIntFromInterval(0, 255));
}

console.log(array);

// sort array
for (var i = 0; i < array.length; i++) {
  var num1 = 0;
  var num2 = 0;

  for (var k = i; k > 0; k--) {
    num1 = array[k - 1];
    num2 = array[k];

    if (num1 > num2) {
      array[k - 1] = num2;
      array[k] = num1;
    }
  }
}

console.log(array);

function randomIntFromInterval(min, max) {
  return Math.floor(Math.random() * (max - min + 1) + min);
}
