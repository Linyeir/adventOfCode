/**
 var input = [
  28,
  33,
  18,
  42,
  31,
  14,
  46,
  20,
  48,
  47,
  24,
  23,
  49,
  45,
  19,
  38,
  39,
  11,
  1,
  32,
  25,
  35,
  8,
  17,
  7,
  9,
  4,
  2,
  34,
  10,
  3,
];
*/

function execute() {
  //sort array, add ground and 2 counter for the jolt-differences 
  input.sort((a, b) => a - b);
  input.unshift(0);
  var oneJoltJumps = 0;
  var threeJoltJumps = 1; //one, because the laptop is always 3 jolts higher than the last adapter

  //go through all adapters and add to the counter if the distance to the one below is relevant
  for (var i = 1; i < input.length; i++) {
    switch (input[i] - input[i - 1]) {
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
  input.sort((a, b) => a - b);
  input.unshift(0);
  input.push(input[input.length-1] + 3);

  //get an array for _dynamic programming_
  var dp = input.map((x,i) => 0);
  dp[dp.length -1] = 1;

  function possibleWays(pos) {
    //if answer is already in array, just return that
    if (dp[pos] != 0) {
      return dp[pos];
    }

    //or compute the sum of the following possible ways
    var sumOfAllFollowingWays = 0;
    for (var i = (pos + 1); i <= (pos+3); i++) {
      if ((input[i] - input[pos]) <= 3) {
        sumOfAllFollowingWays += possibleWays(i);
      }
    }

    //save that sum in the array and return it
    dp[pos] = sumOfAllFollowingWays;
    return sumOfAllFollowingWays;
  }

  document.getElementById("outputBonus").innerHTML = possibleWays(0);
}
