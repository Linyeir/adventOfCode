/**
var input = [
  //should return 127
  35,
  20,
  15,
  25,
  47,
  40,
  62,
  55,
  65,
  95,
  102,
  117,
  150,
  182,
  127,
  219,
  299,
  277,
  309,
  576,
];
*/

function getIncorrectSum() {
    var preamble = 25;

    for (var i = preamble; i < input.length; i++) {
        var isPossible = false;
        for (var j = (i - preamble); j < i; j++) {
            for (var k = (i - preamble); k < i; k++) {
                if (j != k && (input[j] + input[k] == input[i])) {
                    isPossible = true;
                }
            }
        }
        if (!isPossible) {
            return input[i];
        }
    }
    return false;
}

function execute() {
    document.getElementById("output").innerHTML = getIncorrectSum();
}

function executeBonus() {
    var incorrectSum = getIncorrectSum();

    for (var i = 0; i < input.length; i++) {
        var sumOfContigousSet = input[i];

        for (var j = i + 1; j < input.length; j++) {
            sumOfContigousSet += input[j];

            if (sumOfContigousSet > incorrectSum) {
                j = input.length;
            }

            if (sumOfContigousSet == incorrectSum) {
                var toSearch = input.slice(i, j + 1);
                toSearch.sort(function (a, b) { return a - b });
                document.getElementById("outputBonus").innerHTML = toSearch[0] + toSearch[toSearch.length - 1];
                return;
            }

        }
    }
}
