console.log(faculty(5));
console.log(facultyRecursive(5));

function faculty(number) {
  let num = 1;

  for (let i = 1; i <= number; i++) {
    num *= i;
  }

  return num;
}

function facultyRecursive(number) {
  if (number < 1) {
    return 1;
  }

  return number * facultyRecursive(number - 1);
}
