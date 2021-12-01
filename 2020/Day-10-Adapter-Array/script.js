


function execute() {
  var adapters = input.split("\n")
    .map(x => +x)
    .sort((a, b) => a - b);
  adapters.unshift(0);

  console.log(adapters);
  //sort array, add ground and 2 counter for the jolt-differences 
  var oneJoltJumps = 0;
  var threeJoltJumps = 1; //one, because the laptop is always 3 jolts higher than the last adapter

  //go through all adapters and add to the counter if the distance to the one below is relevant
  for (var i = 1; i < adapters.length; i++) {
    switch (adapters[i] - adapters[i - 1]) {
      case 1:
        oneJoltJumps++;
        break;
      case 2:
        break;
      case 3:
        threeJoltJumps++;
        break;
      default:
        console.log("error");
        return;
    }
  }

  document.getElementById("output").innerHTML = oneJoltJumps * threeJoltJumps;
}

function executeBonus() {
  //sort input and add "ground" and "charger"
  var adapters = input.split("\n")
    .map(x => +x)
    .sort((a, b) => a - b);

  adapters.unshift(0);
  adapters.push(adapters[adapters.length - 1] + 3);

  //get an array for _dynamic programming_
  var dp = adapters.map((x, i) => 0);
  dp[dp.length - 1] = 1;

  function possibleWays(pos) {
    //if answer is already in array, just return that
    if (dp[pos] != 0) {
      return dp[pos];
    }

    //or compute the sum of the following possible ways
    var sumOfAllFollowingWays = 0;
    for (var i = (pos + 1); i <= (pos + 3); i++) {
      if ((adapters[i] - adapters[pos]) <= 3) {
        sumOfAllFollowingWays += possibleWays(i);
      }
    }

    //save that sum in the array and return it
    dp[pos] = sumOfAllFollowingWays;
    return sumOfAllFollowingWays;
  }

  document.getElementById("outputBonus").innerHTML = possibleWays(0);
}
