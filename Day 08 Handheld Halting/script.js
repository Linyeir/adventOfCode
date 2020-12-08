/*
var input = [ //should return 5
    { operation: "nop", argument: '+0' },
    { operation: "acc", argument: '+1' },
    { operation: "jmp", argument: '+4' },
    { operation: "acc", argument: '+3' },
    { operation: "jmp", argument: '-3' },
    { operation: "acc", argument: '-99' },
    { operation: "acc", argument: '+1' },
    { operation: "jmp", argument: '-4' },
    { operation: "acc", argument: '+6' }]
*/

function execute() {
    var accumulator = 0;
    var pos = 0;
    var visited = [];

    while (!(visited.includes(pos))) {
        visited.push(pos);

        switch (input[pos].operation) {
            case "acc":
                accumulator += parseInt(input[pos].argument);
                pos++;
                break;
            case "jmp":
                pos += parseInt(input[pos].argument);
                break;
            case "nop":
                pos++;
                break;
            default:
                console.log("error");
        }
    }

    document.getElementById("output").innerHTML = accumulator;

}

function executeBonus() {

    function tryNow() {
        var accumulator = 0;
        var pos = 0;
        var visited = [];

        while (!(visited.includes(pos))) {
            if (pos == input.length) {
                document.getElementById("outputBonus").innerHTML = accumulator;
                return true;
            }else{
                visited.push(pos);

                switch (input[pos].operation) {
                    case "acc":
                        accumulator += parseInt(input[pos].argument);
                        pos++;
                        break;
                    case "jmp":
                        pos += parseInt(input[pos].argument);
                        break;
                    case "nop":
                        pos++;
                        break;
                    default:
                        console.log("error");
                }
            }
        }
        return false;
    }

    // since it's give that exactly one change will suffice, we can skip backtracking and just iterate through the whole array

    for (var i = 0; i < input.length; i++) {
        console.log("testing pos " + i);
        if (input[i].operation == "jmp") {
            input[i].operation = "nop";
            if(tryNow()){return}
            input[i].operation = "jmp";
        }
        if (input[i].operation == "nop") {
            input[i].operation = "jmp";
            if(tryNow()){return}
            input[i].operation = "nop";
        }
    }


}
