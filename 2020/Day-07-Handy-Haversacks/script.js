/**
//input for part 1, result should be 4
var input = [
    { outerBag: "light red", innerBags: [{ number: 1, color: "bright white"}, { number: 2, color: "muted yellow"}] },
    { outerBag: "dark orange", innerBags: [{ number: 3, color: "bright white"}, { number: 4, color: "muted yellow"}] },
    { outerBag: "bright white", innerBags: [{ number: 1, color: "shiny gold"}] },
    { outerBag: "muted yellow", innerBags: [{ number: 2, color: "shiny gold"}, { number: 9, color: "faded blue"}] },
    { outerBag: "shiny gold", innerBags: [{ number: 1, color: "dark olive"}, { number: 2, color: "vibrant plum"}] },
    { outerBag: "dark olive", innerBags: [{ number: 3, color: "faded blue"}, { number: 4, color: "dotted black"}] },
    { outerBag: "vibrant plum", innerBags: [{ number: 5, color: "faded blue"}, { number: 6, color: "dotted black"}] },
    { outerBag: "faded blue", innerBags: [] },
    { outerBag: "dotted black", innerBags: [] }]



//input for part 2, result should be 126

var input = [
    { outerBag: "shiny gold", innerBags: [{ number: 2, color: "dark red"}] },           // value = 126
    { outerBag: "dark red", innerBags: [{ number: 2, color: "dark orange"}] },          // value = 63
    { outerBag: "dark orange", innerBags: [{ number: 2, color: "dark yellow"}] },       // value = 31
    { outerBag: "dark yellow", innerBags: [{ number: 2, color: "dark green"}] },        // value = 15
    { outerBag: "dark green", innerBags: [{ number: 2, color: "dark blue"}] },          // value = 7
    { outerBag: "dark blue", innerBags: [{ number: 2, color: "dark violet"}] },         // value = 3
    { outerBag: "dark violet", innerBags: [] }]                                         // value = 1
*/

function execute() {
    var containingGold = ["shiny gold"];
    var gotNewBags = true;
    var workingContainingGold = [...containingGold];

    while(gotNewBags){

        input.forEach(bagRule => {
           bagRule.innerBags.forEach(innerBag => {
                if(!(workingContainingGold.includes(bagRule.outerBag))
                    && workingContainingGold.includes(innerBag.color)){
                        workingContainingGold.push(bagRule.outerBag);
                }
            })    
        })

        if(containingGold.length != workingContainingGold.length){
            containingGold = [...workingContainingGold];
        }else{
            gotNewBags = false;
        }
    }

    document.getElementById("output").innerHTML = containingGold.length -1;
}

function executeBonus() {

    function getBag(color){
        var toreturn;
        input.forEach(bagRule =>{
            if(bagRule.outerBag === color){
                toreturn = bagRule;
            }
        })
        return toreturn;
    }

    function getValueOfBag(color){
        var valueOfThisBag = 0;
        var thisBag = getBag(color);

        if(thisBag.innerBags.length == 0){
            return 1;
        }else{
            thisBag.innerBags.forEach(innerBag=>{
                valueOfThisBag += innerBag.number * getValueOfBag(innerBag.color);
            })
            return valueOfThisBag + 1;
        }

    } 

    document.getElementById("outputBonus").innerHTML = getValueOfBag("shiny gold") -1;
}
