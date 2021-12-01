function execute() {
    //loop twice nested through every value
    input.forEach(firstNumber =>{
        input.forEach(secondNumber =>{
            // if the sum of the two nested values is 2020, print their product
            if (firstNumber + secondNumber == 2020) {
                document.getElementById("output").innerHTML = firstNumber * secondNumber;
                return;
            }
        })
    })
}

function executeBonus() {
    //loop three times nested through every value
    input.forEach(firstNumber =>{
        input.forEach(secondNumber =>{
            // to reduce computing you can stop if the some of the first two values is already bigger then 2020
            if (firstNumber + secondNumber < 2020) {
                
                input.forEach(thirdNumber =>{
                    // if the sum of the three nested values is 2020, print their product
                    if (firstNumber + secondNumber + thirdNumber== 2020) {
                        document.getElementById("outputBonus").innerHTML = firstNumber * secondNumber *thirdNumber;
                        return;
                    }
                })
            }
        })
    })
}
