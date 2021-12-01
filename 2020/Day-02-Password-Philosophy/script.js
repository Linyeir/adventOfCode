function checkList() {
    var count = 0;
    input.forEach( entry =>{
        var occurence = checkPassword(entry.password, entry.letter);
        if(entry.min <= occurence &&  occurence <= entry.max){
            count++;
        }
})
    document.getElementById("output").innerHTML = count;
}

function checkPassword(password, letter) {
    return password.split(letter).length - 1;
}

function checkBonus(){
    var count = 0;

    input.forEach( entry =>{
        if((entry.password.charAt(entry.min-1) == entry.letter) != (entry.password.charAt(entry.max-1) == entry.letter)){
            count++;
        }
    })

    document.getElementById("outputBonus").innerHTML = count;

}