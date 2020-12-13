


function getIncorrectSum(preamble, XMAS) {
    for (var i = preamble; i < XMAS.length; i++) {
        var isPossible = false;
        for (var j = (i - preamble); j < i; j++) {
            for (var k = (i - preamble); k < i; k++) {
                if (j != k && (XMAS[j] + XMAS[k] == XMAS[i])) {
                    isPossible = true;
                }
            }
        }
        if (!isPossible) {
            return XMAS[i];
        }
    }
    return false;
}

function execute() {
    var XMAS = input.split("\n")
        .map(x => +x);
    var preamble = 25;
    console.log(XMAS);
    document.getElementById("output").innerHTML = getIncorrectSum(preamble, XMAS);
}

function executeBonus() {
    var XMAS = input.split("\n")
        .map(x => +x);
    var preamble = 25;
    var incorrectSum = getIncorrectSum(preamble, XMAS);

    for (var i = 0; i < XMAS.length; i++) {
        var sumOfContigousSet = XMAS[i];

        for (var j = i + 1; j < XMAS.length; j++) {
            sumOfContigousSet += XMAS[j];

            if (sumOfContigousSet > incorrectSum) {
                j = XMAS.length;
            }

            if (sumOfContigousSet == incorrectSum) {
                var toSearch = XMAS.slice(i, j + 1);
                toSearch.sort(function (a, b) { return a - b });
                document.getElementById("outputBonus").innerHTML = toSearch[0] + toSearch[toSearch.length - 1];
                return;
            }

        }
    }
}
