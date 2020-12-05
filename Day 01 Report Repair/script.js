function execute() {
    //loop twice nested through every value
    for (var i = 0; i < input.length; i++) {
        for (var j = 0; j < input.length; j++) {
            // if the sum of the two nested values is 2020, print their product
            if (input[i] + input[j] == 2020) {
                document.getElementById("output").innerHTML = input[i] * input[j];
                return;
            }
        }
    }
}

function executeBonus() {
    //loop three times nested through every value
    for (var i = 0; i < input.length; i++) {
        for (var j = 0; j < input.length; j++) {
            // to reduce computing you can stop if the some of the first two values is already bigger then 2020
            if (input[i] + input[j] < 2020) {
                for (var k = 0; k < input.length; k++) {
                    // if the sum of the three nested values is 2020, print their product
                    if (input[i] + input[j] + input[k] == 2020) {
                        document.getElementById("outputBonus").innerHTML = input[i] * input[j] * input[k];
                        return;
                    }
                }
            }
        }
    }
}
