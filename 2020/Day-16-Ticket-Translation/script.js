function execute() {
  var parts = input.split("\n\n");
  var fields = parts[0].split("\n");
  console.log(fields);
  fields = fields.map(x => {
    var extract = /(\d*)-(\d*) or (\d*)-(\d*)/.exec(x)
    extract.shift();
    return extract.map(y => {
      return Number(y);
    })
  })
  console.log("fields");
  console.log(fields);

  var errorRate = 0;

  var otherTickets = parts[2].split("\n");
  otherTickets.shift();
  otherTickets = otherTickets.map(x => {
    y = x.split(",");
    return y.map(y => {
      return Number(y);
    })
  });
  console.log("otherTickets");
  console.log(otherTickets);

  otherTickets.forEach(ticket => {
    console.log(ticket);
    ticket.forEach((x, i) => {
      var isPossible = false;
      console.log(x, i);
      fields.forEach(rule => {
        if ((rule[0] <= x && x <= rule[1]) || (rule[2] <= x && x <= rule[3])) {
          isPossible = true;
        }
      })
      if (!isPossible) {
        console.log(x);
        errorRate += x;
      }
    })
  })

  document.getElementById("output").innerHTML = errorRate;
}

function executeBonus() {
  var parts = inputTest2.split("\n\n");
  var fields = parts[0].split("\n");

  fields = fields.map(x => {
    var extract = /([\w\s]*): (\d*)-(\d*) or (\d*)-(\d*)/.exec(x)

    extract.shift();
    return extract.map((y, i) => {
      if (i > 0) {
        return Number(y);
      }
      return y;
    })
  })
  console.log("fields");
  console.log(fields);

  var otherTickets = parts[2].split("\n")
    .map(x => {
      y = x.split(",");
      return y.map(y => {
        return Number(y);
      })
    })
    .filter(ticket => {
      return ticket.every(entry => fields.some(rule =>
        (rule[1] <= entry && entry <= rule[2]) || (rule[3] <= entry && entry <= rule[4])));
    });
    
  console.log(otherTickets);

  var fieldValues = "s";



  var myTicket = parts[1].split("\n");
  myTicket = myTicket[1].split(",");
  var myTicket = myTicket.map(x => {
    return Number(x);
  });
  console.log("myTicket");
  console.log(myTicket);


  document.getElementById("outputBonus").innerHTML = "finish";
}
