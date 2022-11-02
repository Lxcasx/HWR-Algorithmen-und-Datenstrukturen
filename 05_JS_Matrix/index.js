function createMatrix(columns, rows) {
  var matrix = [];
  var index = 0;

  for (var i = 0; i < rows; i++) {
    var row = [];

    for (k = 0; k < columns; k++) {
      index++;
      row.push(index);
    }

    matrix.push(row);
  }

  return matrix;
}

var matrix = createMatrix(4, 4);
console.log(matrix);
