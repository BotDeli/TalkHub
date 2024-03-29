const timeAnimation = '0.2s';

class ControllerPopupFromInputCells {
    constructor(inputCellsClassName, popupTextClassName) {
        this.inputCells = document.getElementsByClassName(inputCellsClassName);
        this.popupsText = document.getElementsByClassName(popupTextClassName);

        if (this.inputCells.length !== this.popupsText.length) {
            console.error("Count input cells not equals popups text");
            return
        }

        for (let i = 0; i < this.inputCells.length; i++) {
            let inputCellFor = this.inputCells[i].getAttributeNode("for");
            if (inputCellFor === undefined) {
                continue;
            }

            let targetInput = document.getElementById(inputCellFor.value);
            let popup = this.popupsText[i];
            if (targetInput === undefined || popup === undefined) {
                continue;
            }

            let inputType = targetInput.getAttribute("type");
            if (isTypeDatetime(inputType)) {
                popup.style.zIndex = '2';
            }

            if (popup.style.zIndex === '2') {
                popup.style.animation = `flyAway ${timeAnimation} forwards`;
            }

            targetInput.addEventListener('focusin', () => {
                this.inputCells[i].style.borderBottomColor = 'blueviolet';
                popup.style.color = 'black';
                if (popup.style.zIndex !== '2') {
                    popup.style.animation = `flyAway ${timeAnimation} forwards`;
                }
            });

            const checkFunction = selectCheckFunction(popup);


            targetInput.addEventListener('focusout', () => {
                if (targetInput.value === "") {
                    this.inputCells[i].style.borderBottomColor = 'black';
                    popup.style.color = 'black';
                    if (popup.style.zIndex !== '2') {
                        popup.style.animation = `flyAwayReversed ${timeAnimation} forwards`;
                    }
                } else if (checkFunction(targetInput.value)) {
                    this.inputCells[i].style.borderBottomColor = 'green';
                    popup.style.color = 'green';
                } else {
                    this.inputCells[i].style.borderBottomColor = 'red';
                    popup.style.color = 'red';
                }
            });
        }
    }

     defaultSettings() {
        Object.values(this.inputCells).forEach(value => {
            value.style.borderBottomColor = 'black';
        });
        Object.values(this.popupsText).forEach(value => {
            value.style.color = 'black';

            if (value.style.zIndex !== '2') {
                value.style.animation = `flyAwayReversed ${timeAnimation} forwards`;
            }
        })
     }
}

function isTypeDatetime(type) {
    return type === 'datetime-local';
}

function selectCheckFunction(popup) {
    const text = popup.innerText;
    if (isEmail(text)) {
        return checkCorrectEmail;
    } else if (isPassword(text)) {
        return checkCorrectPassword;
    } else if (isMeetingCode(text)) {
        return onlyEnglishLettersAndDigits
    } else if (isDatetime(text)) {
        return datetimeInTheFuture
    }
    return onlyLetters;
}

function isEmail(text) {
    return text === "Email" || text === "Электронная почта";
}

function checkCorrectEmail(value) {
    return /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(value);
}

function isPassword(text) {
    return text === "Password" || text === "Confirm Password" || text === "New Password" || text === "Пароль" || text === "Подтвердите пароль" || text === "Новый пароль";
}

function checkCorrectPassword(value) {
    return value.length >= 8 && onlyEnglishLetters(value.charAt(0)) && onlyEnglishLettersAndDigits(value);
}

function datetimeInTheFuture(value) {
    return new Date() < new Date(value);
}

function onlyLetters(value) {
    return /^[a-zA-Zа-яА-Я]+$/.test(value);
}

function onlyEnglishLetters(value) {
    return /^[a-zA-Z]+$/.test(value);
}

function onlyEnglishLettersAndDigits(value) {
    return /^[a-zA-Z0-9]+$/.test(value);
}

function isMeetingCode(text) {
    return text === "Enter the meeting code" || text === "Введите код встречи";
}

function isDatetime(text) {
    return text === "Datetime" || text === "Дата и время";
}