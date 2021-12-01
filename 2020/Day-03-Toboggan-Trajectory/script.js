/*
var input = ["..##.......",
              "#...#...#..",
              ".#....#..#.",
              "..#.#...#.#",
              ".#...##..#.",
              "..#.##.....",
              ".#.#.#....#",
              ".#........#",
              "#.##...#...",
              "#...##....#",
              ".#..#...#.#"]
*/

function execute(){
    document.getElementById("output").innerHTML = checkDiagonal(3,1);
}

function executeBonus(){
    
    var product = checkDiagonal(1,1) * checkDiagonal(3,1)* checkDiagonal(5,1)* checkDiagonal(7,1)* checkDiagonal(1,2);

    document.getElementById("outputBonus").innerHTML = product;
}

function checkDiagonal(right, down){
    var count = 0;
    for(var i = 0; i < input.length; i++){
        if(i % down == 0){
            if(input[i].charAt(((i/down)*right) % input[i].length) == "#"){
                count++;
            }
        }
    }
    return count;
}