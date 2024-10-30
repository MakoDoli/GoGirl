function ifElse(a, b, max) {
  if (a + b < max) {
    return a + b;
  } else {
    console.log(` Sum of ${a} and ${b} is more than ${max}`);
  }
  return max;
}

console.log(ifElse(2, 3, 10), ifElse(6, 7, 11));
