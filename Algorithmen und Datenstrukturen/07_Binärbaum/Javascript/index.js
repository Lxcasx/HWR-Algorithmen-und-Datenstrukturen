var values = [1, 42, 82, 69, 24, 9, 32];

// sort array
values = values.sort((a, b) => a - b);

var middle = values[Math.round((values.length - 1) / 2)];

var tree = {
  value: middle,
  left: null,
  right: null,
};

tree.left = {
  value: 42,
  left: 1,
  right: 9,
};

tree.right = {
  value: 69,
  left: 42,
  right: 82,
};

printTree(tree);

// TODO: Werte löschen

function printTree(tree) {
  console.log(tree.value);

  if (typeof tree.left === "object") {
    printTree(tree.left);
  } else {
    console.log(tree.left);
  }
}
