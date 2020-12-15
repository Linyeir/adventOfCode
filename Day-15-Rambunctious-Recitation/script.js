function findNthNumber(startingNumber, n){

  var lastSpokenNumbers = new Map();

  startingNumber.forEach((x, i) => {
    if (i != startingNumber.length -1) {
      lastSpokenNumbers.set(Number(x), i);
    }
  })

  var lastNumber = Number(startingNumber[startingNumber.length - 1]);

  for (var i = startingNumber.length; i < n; i++) {
    if(lastSpokenNumbers.has(lastNumber)){
      distance = i -1 - lastSpokenNumbers.get(lastNumber);
      lastSpokenNumbers.set(lastNumber, i-1);
      lastNumber = distance;
    }else{
      lastSpokenNumbers.set(lastNumber, i-1);
      lastNumber = 0;
    }
  }

  return lastNumber;
}


function execute() {
  document.getElementById("output").innerHTML = findNthNumber( input.split(","), 2020);
}

function executeBonus() {
  document.getElementById("outputBonus").innerHTML = findNthNumber( input.split(","), 30000000);
}
