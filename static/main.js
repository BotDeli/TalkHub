const inputCells = document.getElementsByClassName('input-cell');
const popupsText = document.getElementsByClassName('authorization-data-popup-text');

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
                if (node.value === '') {
                    popup.style.animation = 'flyAwayReversed 0.4s forwards';
                    popup.style.color = 'red';
                    inputCells[i].style.borderBottomColor = 'red';
                } else {
                    inputCells[i].style.borderBottomColor = 'black';
                    popup.style.color = 'black';
                }
            });
        }
    }
}

const email = document.getElementById('email-input');
const password = document.getElementById('password-input');

const authorizationError = document.getElementById('authorization-error');
let authorizationErrorSwitcher = new HideSwitcher(authorizationError);

btnGoAccount.addEventListener('click', () => {
    authorizationErrorSwitcher.hide();
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
        if (response.status === 200) {
                document.location.reload();
        } else {
            InvalidInputData();
        }
    });
});

function InvalidInputData() {
    authorizationError.innerText = "Invalid email or password";
    authorizationErrorSwitcher.show();
}
