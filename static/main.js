new ControllerPopupFromInputCells('input-cell', 'input-cell-popup-text');

const email = document.getElementById('email-input');
const password = document.getElementById('password-input');

const formErrorOutput = document.getElementById('form-error-output');
formErrorOutput.style.display = 'none';

document.getElementById('btn-go-account').addEventListener('click', () => {
    formErrorOutput.style.display = 'none';
    if (email.value.length < 5 || password.value.length < 8) {
        InvalidInputData();
        return;
    }
    fetch('/goToAccount', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        },
        body: JSON.stringify({
            'email': email.value,
            'password': password.value,
        })
    }).
    then(response => {
        if (response.status === 308) {
                window.location.reload();
        } else {
            InvalidInputData();
        }
    })
});

function InvalidInputData() {
    formErrorOutput.style.display = 'unset';
    formErrorOutput.innerText = "Invalid email or password";
}
