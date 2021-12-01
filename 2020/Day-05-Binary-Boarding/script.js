/**
var input = "BFFFBBFRRR"; // should return 567
var input = "FFFBBBFRRR"; // should return 119
var input = "BBFFBBFRLL"; // should return 820
**/

function getSeatNumber(bsp) {
    return 8 * (
        64 * (bsp.charAt(0) == 'B') +
        32 * (bsp.charAt(1) == 'B') +
        16 * (bsp.charAt(2) == 'B') +
        8 * (bsp.charAt(3) == 'B') +
        4 * (bsp.charAt(4) == 'B') +
        2 * (bsp.charAt(5) == 'B') +
        (bsp.charAt(6) == 'B')
    ) + (
            4 * (bsp.charAt(7) == 'R') +
            2 * (bsp.charAt(8) == 'R') +
            (bsp.charAt(9) == 'R')
        )
}

function execute() {
    var largestSeatNumber = 0;
    for (var i = 0; i < input.length; i++) {
        var thisSetNumber = getSeatNumber(input[i]);
        
        if (thisSetNumber > largestSeatNumber){
            largestSeatNumber = thisSetNumber;
        }
    }
    document.getElementById("output").innerHTML = largestSeatNumber;
}

function executeBonus() {

    // convert all seat codes into numbers
    for (var i = 0; i < input.length; i++) {
        input[i] = getSeatNumber(input[i]);
    }

    // numerically sort the array
    input.sort(function(a, b){return a - b}); 

    var mySeat;

    // go through all seats again
    for (var i = 0; i < input.length-1; i++) {
        // when the Number of the next seat is 2 larger then the current one, the one between those is your seat
        if(input[i+1] == input[i]+2){
            mySeat = input[i]+1;
        }
    }

    document.getElementById("output").innerHTML = mySeat;
}
