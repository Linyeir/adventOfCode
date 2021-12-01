function execute() {
    var count = 0;
    for (var i = 0; i < input.length; i++) {
        var isValid =
            input[i].hasOwnProperty('byr') &&
            input[i].hasOwnProperty('iyr') &&
            input[i].hasOwnProperty('eyr') &&
            input[i].hasOwnProperty('hgt') &&
            input[i].hasOwnProperty('hcl') &&
            input[i].hasOwnProperty('ecl') &&
            input[i].hasOwnProperty('pid');

        console.log(input[i].pid + ": " + isValid);

        if (isValid) { count++; }
    }
    document.getElementById("output").innerHTML = count;
}

function executeBonus() {
    var i = 0;
    var count = 0;

    function checkYr(field, min, max) {
        return min <= field && field <= max;
    }

    function checkHcl(field) {
        return /#[a-f\d]{6}$/.test(field);
    }

    function checkEcl(field) {
        return ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'].includes(field);
    }

    function checkHgt(field) {
        hgt = parseInt(field);
        if (/cm/.test(field)) {
            return 150 <= hgt && hgt <= 193;
        } else if (/in/.test(field)) {
            return 59 <= hgt && hgt <= 76;
        };
    }

    function checkPid(field) {
        return /^\d{9}$/.test(field);
    }

    for (i = 0; i < input.length; i++) {

        var isValid = 
            checkYr(input[i].byr, 1920, 2002) &&
            checkYr(input[i].iyr, 2010, 2020) &&
            checkYr(input[i].eyr, 2020, 2030) &&
            checkHgt(input[i].hgt) &&
            checkHcl(input[i].hcl) &&
            checkEcl(input[i].ecl) &&
            checkPid(input[i].pid);

        console.log(input[i].pid + ": " + isValid);

        if (isValid) { count++; }

    }

    document.getElementById("outputBonus").innerHTML = count;
}
