new ControllerPopupFromInputCells('input-cell input-cell-registration', 'input-cell-popup-text input-cell-popup-text-registration');

const acceptPolicy = document.getElementById('accept-policy');
const acceptPolicyText = document.getElementById('accept-policy-text');

acceptPolicy.addEventListener('change', () => {
    if (acceptPolicy.checked) {
        acceptPolicyText.style.color = 'black';
    } else {
        acceptPolicyText.style.color = 'red';
    }
});

const firstName = document.getElementById('first-name-input');
const lastName = document.getElementById('last-name-input');
const email = document.getElementById('email-input');
const password = document.getElementById('password-input');
const confirmPassword = document.getElementById('confirm-password-input');

const formErrorOutput = document.getElementById('form-error-output');
formErrorOutput.style.display = 'none';

document.getElementById('btn-go-account').addEventListener('click', () => {
    formErrorOutput.style.display = 'none';
    if (checkCorrectInputData()) {
        fetch('/createAccount', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            },
            body: JSON.stringify({
                'first_name': firstName.value,
                'last_name': lastName.value,
                'email': email.value,
                'password': password.value,
            })
        }).then(response => {
            if (response.status === 201) {
                window.location.reload();
            } else {
                let json = response.json();
                json.then((data) => {
                    formErrorOutput.innerText = data.error;
                    formErrorOutput.style.display = 'unset';
                })
            }
        });
    }
});

function checkCorrectInputData() {
    return acceptPolicy.checked &&
        checkEqualsPassword() &&
        checkCorrectEmail(email.value) &&
        checkCorrectPassword(password.value) &&
        onlyLetters(firstName.value) &&
        onlyLetters(lastName.value);
}

function checkEqualsPassword() {
    return password.value === confirmPassword.value;
}