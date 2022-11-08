var array = [];

fillArray();

console.log(array);

for (var i = array.length; i > 1; i--) {
  for (var k = 0; k < i - 1; k++) {
    var num1 = array[k];
    var num2 = array[k + 1];

    if (num1 > num2) {
      array[k] = num2;
      array[k + 1] = num1;
    }
  }
}

console.log(array);

function fillArray() {
  // fill array
  for (var i = 0; i < 50; i++) {
    array.push(randomIntFromInterval(0, 255));
  }
}

function randomIntFromInterval(min, max) {
  return Math.floor(Math.random() * (max - min + 1) + min);
}
