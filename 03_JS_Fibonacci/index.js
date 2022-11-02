let nums = [0, 1];

console.log(getFibonacciNumAt(5));

function getFibonacciNumAt(index) {
  let last = 0;
  let last2 = 1;

  for (let i = 0; i < index + 1; i++) {
    if (i <= 1) {
      continue;
    }

    nums[i] = last + last2;
    last = last2;
    last2 = nums[i];
  }

  return nums[index];
}
