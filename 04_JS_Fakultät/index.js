function faculty(number) {
  let num = 1;

  for (let i = 1; i <= number; i++) {
    num *= i;
  }

  return num;
}

console.log(faculty(4));
console.log(facultyRecursive(4));

function facultyRecursive(index, number = 1) {
  number *= index;
  index--;

  if (index == 0) return number;

  return facultyRecursive(index, number);
}
