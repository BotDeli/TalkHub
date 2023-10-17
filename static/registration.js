const inputCells = document.getElementsByClassName('input-cell');
const popupsText = document.getElementsByClassName('registration-data-popup-text');

const registrationError = document.getElementById('registration-error');
let registrationErrorSwitcher = new HideSwitcher(registrationError);

for (let i = 0; i < inputCells.length; i++) {
    let nodeFor = inputCells[i].getAttributeNode("for")
    if (nodeFor !== null) {
        let node = document.getElementById(nodeFor.value);
        let popup = popupsText[i];
        if (node !== null && popup != null) {
            node.addEventListener('focusin', () => {
                inputCells[i].style.borderBottomColor = 'blueviolet';
                popup.style.animation = 'flyAway 0.4s forwards';
            });
            node.addEventListener('focusout', () => {
                registrationErrorSwitcher.hide();
                inputCells[i].style.borderBottomColor = 'black';
                popup.style.color = 'black';
                [correct, error] = correctInput(node.value, i);
                if (!correct) {
                    popup.style.color = 'red';
                    inputCells[i].style.borderBottomColor = 'red';
                    if (node.value === '') {
                        popup.style.animation = 'flyAwayReversed 0.4s forwards';
                    }

                    if (error.length > 0) {
                        registrationError.innerText = error;
                        registrationErrorSwitcher.show();
                    }
                }
            });
        }
    }
}

const minLengthInputData = [1, 1, 1, 8, 8];

function correctInput(value, i) {
    if (i < 2) {
        if (value.length >= minLengthInputData[i]) {
            if (onlyLetters(value)) {
                return [true, ''];
            }
            return [false, 'Name must contain only letters'];
        }
        return [false, ''];
    } else if (i === 2) {
        if (value.length >= minLengthInputData[i]) {
            if (checkEmail(value)) {
                return [true, ''];
            }
            return [false, 'Dont correct email'];
        }
        return [false, ''];
    } else if (value.length >= minLengthInputData[i]) {
        if (onlyLetters(value.charAt(0))) {
            return [true, ''];

        }
        return [false, 'Password must start with letter'];
    }
    return [false, 'Password must contain at least 8 characters.'];
}

function onlyLetters(value) {
    return /^[a-zA-Zа-яА-Я]+$/.test(value);
}

function checkEmail(value) {
    return /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(value);
}

const firstName = document.getElementById('first-name-input');
const lastName = document.getElementById('last-name-input');
const email = document.getElementById('email-input');
const password = document.getElementById('password-input');
const confirmPassword = document.getElementById('confirm-password-input');


const acceptPolicy = document.getElementById('accept-policy');
const acceptPolicyText = document.getElementById('accept-policy-text');

acceptPolicy.addEventListener('change', () => {
    if (acceptPolicyText.style.color === 'red') {
        acceptPolicyText.style.color = 'black';
    } else {
        acceptPolicyText.style.color = 'red';
    }
    registrationErrorSwitcher.hide();
});

btnCreateAccount.addEventListener('click', () => {
    registrationErrorSwitcher.hide();
    if (acceptPolicy.checked) {
        if (correctPassword()) {
            fetch('/createAccount', {
                method: 'POST',
                headers: {
                    'Accept': 'application/json'
                },
                body: JSON.stringify({
                    'first-name': firstName,
                    'last-name': lastName,
                    'email': email,
                    'password': password,
                })
            }).
                then(data => data.json()).
                then(response => {
                    if(response.error === "") {
                        document.location.reload();
                    } else {
                        registrationError.innerText = response.error;
                        registrationErrorSwitcher.show();
                    }
                });
        } else {
            registrationError.innerText = 'password and confirm password dont equals';
            registrationErrorSwitcher.show();
        }
    }
});

function correctPassword() {
    return password.value === confirmPassword.value;
}