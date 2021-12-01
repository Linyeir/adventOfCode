var matrix = [
  { x: -1, y: -1 }, { x: 0, y: -1 }, { x: 1, y: -1 },
   { x: -1, y: 0 },   /* centre */    { x: 1, y: 0 },
   { x: -1, y: 1 },  { x: 0, y: 1 },  { x: 1, y: 1 }
]

function getFullAdjacentSeats(layout, row, column) {
  var ctr = 0;
  matrix.forEach(direction => {
    if (((row + direction.x) >= 0)
      && ((row + direction.x) < layout.length)
      && ((column + direction.y) >= 0)
      && ((column + direction.y) < layout[row].length)
      && !(((row + direction.x) == row) && ((column + direction.y) == column))
      && (layout[(row + direction.x)].charAt(column + direction.y) == '#')) {
      ctr++;
    }
  })
  return ctr;
}

function getFullVisibleSeats(layout, row, column) {
  var ctr = 0;
  matrix.forEach(direction => {
    var multiplier = 0;
    do {
      multiplier++;
      cx = (row + (multiplier * direction.x));
      cy = (column + (multiplier * direction.y));
      if ((cx >= 0) && (cx < layout.length)
        && (cy >= 0) && (cy < layout[row].length)
        && !((cx == row) && (cy == column))) {

        if ((layout[cx].charAt(cy) == '#')) {
          ctr++;
          multiplier = 0;
        } else if ((layout[cx].charAt(cy) == 'L')) {
          multiplier = 0;
        }

      } else {
        multiplier = 0;
      }
    } while (multiplier > 0)
  })
  return ctr;
}

function LayoutsAreEqual(one, two) {
  if (one.length != two.length) {
    console.log("arrays are of unequal length");
    return false;
  }
  var oneIsFalse = false;
  for (var i = 0; i < one.length; i++) {
    if ((one[i].localeCompare(two[i])) != 0) {
      oneIsFalse = true;
    }
  }
  return !oneIsFalse;
}

function execute() {
  var oldLayout = input.split("\n");
  var newLayout = oldLayout.map(() => "");
  var change = true;

  while (change) {
    newLayout = oldLayout.map(() => "");
    for (var row = 0; row < oldLayout.length; row++) {
      for (var column = 0; column < oldLayout[row].length; column++) {
        if (oldLayout[row].charAt(column) == 'L' && (getFullAdjacentSeats(oldLayout, row, column) == 0)) {
          newLayout[row] += ("#");
        } else if (oldLayout[row].charAt(column) == '#' && (getFullAdjacentSeats(oldLayout, row, column) >= 4)) {
          newLayout[row] += ("L");
        } else {
          newLayout[row] += (oldLayout[row].charAt(column));
        }
      }
    }
    var change = !LayoutsAreEqual(newLayout, oldLayout);
    oldLayout = [...newLayout];
  }

  var fullSeats = 0;

  newLayout.forEach(line => {
    fullSeats += line.split("#").length - 1;
  })

  document.getElementById("output").innerHTML = fullSeats;
}

function executeBonus() {
  var oldLayout = input.split("\n");
  var newLayout = oldLayout.map(() => "");
  var change = true;

  while (change) {
    newLayout = oldLayout.map(() => "");
    for (var row = 0; row < oldLayout.length; row++) {
      for (var column = 0; column < oldLayout[row].length; column++) {
        if (oldLayout[row].charAt(column) == 'L' && (getFullVisibleSeats(oldLayout, row, column) == 0)) {
          newLayout[row] += ("#");
        } else if (oldLayout[row].charAt(column) == '#' && (getFullVisibleSeats(oldLayout, row, column) >= 5)) {
          newLayout[row] += ("L");
        } else {
          newLayout[row] += (oldLayout[row].charAt(column));
        }
      }
    }
    var change = !LayoutsAreEqual(newLayout, oldLayout);
    oldLayout = [...newLayout];
    console.log(newLayout);
  }

  var fullSeats = 0;

  newLayout.forEach(line => {
    fullSeats += line.split("#").length - 1;
  })

  document.getElementById("outputBonus").innerHTML = fullSeats;
}
