let num = 27;

function isNumPrime(p) {
  for (var i = 2; i < p; i++) {
    if (p % i == 0) {
      return false;
    }
  }

  return true;
}

function getPrimes(n) {
  var primes = [];
  var i = 1;

  while (primes.length != n) {
    if (isNumPrime(i)) primes.push(i);

    i++;
  }

  return primes;
}

console.log(getPrimes(5));
