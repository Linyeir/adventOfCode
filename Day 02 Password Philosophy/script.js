function checkList() {
    var count = 0;
    for (var i = 0; i < input.length; i++) {
        var occurence = checkPassword( input[i].password, input[i].letter);
        console.log(occurence);
        if(input[i].min <= occurence &&  occurence <= input[i].max){
            count++;
            console.log("Password #" + i + " (" + input[i].password + ") works.");
        }else{
            console.log("Password #" + i + " (" + input[i].password + ") doesnt work.");
        }
    }
    document.getElementById("output").innerHTML = count;
}

function checkPassword(password, letter) {
    return password.split(letter).length - 1;
}

function checkBonus(){
    var count = 0;

    for (var i = 0; i < input.length; i++) {
        var check = (input[i].password.charAt(input[i].min-1) == input[i].letter) + (input[i].password.charAt(input[i].max-1) == input[i].letter);
        console.log(check);
        if(check == 1){
            count++;
            console.log("Password #" + i + " (" + input[i].password + ") works.");
        }else{
            console.log("Password #" + i + " (" + input[i].password + ") doesnt work.");
        }
    }

    document.getElementById("outputBonus").innerHTML = count;

}