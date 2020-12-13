function execute() {
  var notes = input.split("\n");
  var time = notes[0];
  notes = notes[1].split(",").filter(x => x != "x").map(x => Number(x));
  var nextBus = notes[0];
  console.log(notes);

  notes.forEach(x => {
    if ((x - (time % x)) < (nextBus - (time % nextBus))) {
      nextBus = x;
    }
  })

  document.getElementById("output").innerHTML = nextBus * (nextBus - (time % nextBus));
}

function executeBonus() {
  var notes = input.split("\n");
  notes = notes[1].split(",").map((x, i) => ({
    index: i,
    value: Number(x)
  })).filter(x =>
    !Number.isNaN(x.value)
  ).sort(function (a, b) {
    return b.value - a.value;
  });

  var stepsize = 1;
  var time = 1;

  notes.forEach(x => {
    while ((time + x.index) % x.value) {
      time += stepsize;
    }
    stepsize *= x.value;
  })

  document.getElementById("outputBonus").innerHTML = time;

}
