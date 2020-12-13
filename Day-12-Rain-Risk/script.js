function mod(value, modulo) {
  return ((value % modulo) + modulo) % modulo;
}

function execute() {
  var instructions = input.split("\n");

  var ferry = {
    x: 0,
    y: 0,
    dir: 0
  }

  instructions.forEach(instruction => {
    var value = Number(instruction.substring(1));
    switch (instruction.charAt(0)) {
      case "N":
        ferry.y += value;
        break;
      case "S":
        ferry.y -= value;
        break;
      case "E":
        ferry.x += value;
        break;
      case "W":
        ferry.x -= value;
        break;
      case "L":
        ferry.dir = mod(ferry.dir - value, 360);
        break;
      case "R":
        ferry.dir = mod(ferry.dir + value, 360);
        break;
      case "F":
        var angle = ferry.dir / 180 * Math.PI;
        ferry.x += Math.cos(angle) * value;
        ferry.y -= Math.sin(angle) * value;
        break;
    }
    console.log(ferry);
  })

  document.getElementById("output").innerHTML = Math.abs(ferry.x) + Math.abs(ferry.y);
}


function executeBonus() {

  function turnWaypoint(value) {
    var angle = value / 180 * Math.PI;
    x_ = waypoint.x * Math.cos(angle) - waypoint.y * Math.sin(angle);
    y_ = waypoint.x * Math.sin(angle) + waypoint.y * Math.cos(angle);
    waypoint.x = Math.round(x_);
    waypoint.y = Math.round(y_);
  }

  var instructions = input.split("\n");

  var waypoint = {
    x: 10,
    y: 1
  }

  var ferry = {
    x: 0,
    y: 0
  }

  instructions.forEach(instruction => {
    var value = Number(instruction.substring(1));
    switch (instruction.charAt(0)) {
      case "N":
        waypoint.y += value;
        break;
      case "S":
        waypoint.y -= value;
        break;
      case "E":
        waypoint.x += value;
        break;
      case "W":
        waypoint.x -= value;
        break;
      case "L":
        turnWaypoint(value);
        break;
      case "R":
        turnWaypoint(-value);
        break;
      case "F":
        ferry.x += waypoint.x * Number(instruction.substring(1));
        ferry.y += waypoint.y * Number(instruction.substring(1));
        break;
    }
  })

  document.getElementById("outputBonus").innerHTML = Math.abs(ferry.x) + Math.abs(ferry.y);
}