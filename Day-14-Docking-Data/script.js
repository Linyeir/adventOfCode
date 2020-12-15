function execute() {
  var instructions = input.split("\n");

  var currentMask;
  var memory = new Map();

  instructions.forEach(x => {
    if (/^mask/.test(x)) {
      currentMask = /[X10]*$/.exec(x)[0];
    } else if (/^mem/.test(x)) {
      var address = Number(/\[(\d*)\]/.exec(x)[1]);
      var value = Number(/(\d*)$/.exec(x)[0])
        .toString(2)
        .padStart(36, "0")
        .split("")
        .map((val, i) => {
          return currentMask.charAt(i) == 'X' ? val : currentMask.charAt(i)
        })
        .join("");
      value = parseInt(value, 2);
      memory.set(address, value);
    }
  })

  document.getElementById("output").innerHTML = [...memory.values()].reduce((sum, val) => sum + val, 0);
}

function executeBonus() {
  var instructions = input.split("\n");


  var currentMask;
  var memory = new Map();

  instructions.forEach(x => {
    var allAddresses = [[]];

    if (/^mask/.test(x)) {
      currentMask = /[X10]*$/.exec(x)[0].split("");
    } else if (/^mem/.test(x)) {
      var address = Number(/\[(\d*)\]/.exec(x)[1])
        .toString(2)
        .padStart(36, "0")
        .split("")
        .map((val, i) => {
          if (currentMask[i] == "1") { return "1" }
          else if (currentMask[i] == "X") { return "X" }
          else return val
        });
        
      var value = Number(/(\d*)$/.exec(x)[0]);

      for (var i = 0; i < address.length; i++) {
        switch (address[i]) {
          case "0":
            allAddresses.forEach(x => x.push("0"));
            break;
          case "1":
            allAddresses.forEach(x => x.push("1"));
            break;
          case "X":
            allAddresses = allAddresses.concat(JSON.parse(JSON.stringify(allAddresses)));
            allAddresses.forEach((x, i) => x.push(Number(i < allAddresses.length / 2)));
            break;
          default:
            alert("error")
            break;
        }
      }

      allAddresses.forEach(x =>{
        x = x.join("");
        x = parseInt(x,2);
        memory.set(x, value);
      })

    }
  })

  document.getElementById("outputBonus").innerHTML =  [...memory.values()].reduce((sum, val) => sum + val, 0);
}
