class openController {
    constructor() {
        this.opened = 0;
    }

    openOne() {
        if (this.opened === 1) {
            this.opened = 0;
            return true;
        }
        this.opened = 1;
        return false;
    }

    openTwo() {
        if (this.opened === 2) {
            this.opened = 0;
            return true;
        }
        this.opened = 2;
        return false;
    }
}

const authorization = document.getElementById("authorization");
const errorOutput = document.getElementById("error-output");

const hideSwitcherAuthorization = new HideSwitcher(authorization, "flex");
const hideSwitcherError = new HideSwitcher(errorOutput, "unset");

document.getElementById("close-authorization").addEventListener("click", (event) => {
    event.preventDefault();
    hideSwitcherAuthorization.hide();
})

const process = document.getElementById("process");
const sendLogin = document.getElementById("send-data");

const login = document.getElementById("login");

const password = document.getElementById("password");

const openControl = new openController();

document.getElementById("register-btn").addEventListener("click", (event) => {
    event.preventDefault();
    hideSwitcherAuthorization.hide();
    if (openControl.openOne()) {
        return
    }

    process.innerText = "Регистрация";
    sendLogin.value = "Создать";
    sendLogin.onclick = () => {
        sendAuthorizationRequest("/signUp");
    }
    hideSwitcherAuthorization.show();
});

document.getElementById("logIn-btn").addEventListener("click", (event) => {
    event.preventDefault();
    hideSwitcherAuthorization.hide();
    if (openControl.openTwo()) {
        return
    }

    process.innerText = "Вход";
    sendLogin.value = "Войти";
    sendLogin.onclick = () => {
        sendAuthorizationRequest("/signIn");
    }
    hideSwitcherAuthorization.show();
});


function sendAuthorizationRequest(urlRequest) {
    if (checkAuthorizationData()) {
        return;
    }
    fetch(urlRequest, {
        method: "POST",
        body: JSON.stringify({"login": login.value, "password": password.value})
    }).
    then(data => data.json()).
    then(response => {
        if(response.error === "") {
            document.location.reload();
            hideSwitcherError.hide();
        } else {
            errorOutput.innerText = response.error;
            hideSwitcherError.show();
        }
    });
}

function checkAuthorizationData() {
    if(isCorrectLogin() && isCorrectPassword()) {
        hideSwitcherError.hide();
        return false;

    }
    hideSwitcherError.show();
    errorOutput.innerText = "Логин и пароль должен состоять от 6 до 18 символов и не должен начинаться с цифры!";
    return true;
}

function isCorrectLogin() {
    return !(login.value.length < 6 && login.value.length <= 18 && isDigit(login.value.charAt(0)))
}

function isCorrectPassword() {
    return !(password.value.length < 6 && password.value.length <= 18 && isDigit(password.value.charAt(0)))
}

function isDigit(char) {
    return 0 <= char && char <= 9;
}