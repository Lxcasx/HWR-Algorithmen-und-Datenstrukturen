function fakultät(number) {
  let num = 1;

  for (let i = 1; i <= number; i++) {
    num *= i;
  }

  return num;
}

console.log(fakultät(4));
console.log(fakultätReko(4));

function fakultätReko(index, number = 1) {
  number *= index;
  index--;

  if (index == 0) return number;

  return fakultätReko(index, number);
}

/*
let nums = [8, 10, 11, 15];

nums.forEach((num) => {
  console.log(exer(num));
});

function exer(n) {
  if (n > 10) {
    return exer(n - 2);
  } else {
    return n - 1;
  }
}
*/
