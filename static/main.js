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

const authorizationError = document.getElementById('authorization-error');
let authorizationErrorSwitcher = new HideSwitcher(authorizationError);

btnCreateAccount.addEventListener('click', () => {
    authorizationErrorSwitcher.hide();
    fetch('/goToAccount', {
        method: 'POST',
        headers: {
            'Accept': 'application/json'
        },
        body: JSON.stringify({
            'email': email,
            'password': password,
        })
    }).
    then(data => data.json()).
    then(response => {
        if(response.error === "") {
            document.location.reload();
        } else {
            authorizationErrorSwitcher.innerText = response.error;
            authorizationErrorSwitcher.show();
        }
    });
});
