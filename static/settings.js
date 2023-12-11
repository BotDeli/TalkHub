controllerChangePassword = new ControllerPopupFromInputCells('input-cell input-cell-change-password', 'input-cell-popup-text input-cell-popup-text-change-password');

const passwordInput = document.getElementById('password-input');
const newPasswordInput = document.getElementById('new-password-input');

const statusChangePassword = document.getElementById('change-password-status');

document.getElementById('btn-settings-change-password').addEventListener('click', (e) => {
    e.preventDefault();
    statusChangePassword.innerText = "";
    const password = passwordInput.value;
    const newPassword = newPasswordInput.value;
    if (checkCorrectPassword(password) && checkCorrectPassword(newPassword)) {
        passwordInput.value = "";
        newPasswordInput.value = "";
        controllerChangePassword.defaultSettings();
        fetch('/changePassword', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                'password': password,
                'new_password': newPassword,
            }),
        }).then((response) => {
            if (response.status === 200) {
                statusChangePassword.style.color = 'green';
                statusChangePassword.innerText = "Ваш пароль успешно изменен!";
            } else {
                statusChangePassword.style.color = 'red';
                statusChangePassword.innerText = "Некорректный пароль.";
            }
        });
    }
});