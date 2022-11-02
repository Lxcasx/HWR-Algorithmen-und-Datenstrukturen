var s3 = [];
var s2 = ["B", "B", "A", "B", "A", "B", "B", "A"];
var s1 = ["A", "A", "A", "B", "A", "B"];

for (var i = 0; i < s2.length; i++) {
  var element = s2.pop();

  if (element == "B") {
    s3.push(element);
  } else {
    s1.push(element);
  }
}

console.log(s3);
console.log(s2);
console.log(s1);
