var inputTest =
    [["abc"],
    ["a",
        "b",
        "c"],
    ["ab",
        "ac"],
    ["a",
        "a",
        "a",
        "a"],
    ["b"]]




function execute() {
    var sum = 0;
    input.forEach(answerGroup =>{
        sum += (new Set(answerGroup.join('').split(""))).size;
    })
    document.getElementById("output").innerHTML = sum;
}

function executeBonus() {

    var commonAnswersCount = 0;
    
    input.forEach(answerGroup =>{

        var allAnswers = [... new Set(answerGroup.join('').split(""))];

        allAnswers.forEach( answer => {
            if (answerGroup.every(answers => answers.includes(answer))) commonAnswersCount++;            
        })

    })

    document.getElementById("output").innerHTML = commonAnswersCount;
}
