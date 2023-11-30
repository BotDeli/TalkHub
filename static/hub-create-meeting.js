const createMeetingController = new ControllerPopupFromInputCells('input-cell input-cell-create-meeting', 'input-cell-popup-text input-cell-popup-text-create-meeting');

const createMeetingName = document.getElementById('create-meeting-input-name');
const createMeetingDatetime = document.getElementById('create-meeting-input-datetime');

awaitMeetingsList = new MeetingAwaitMeetingsList();
createMeetingDatetime.value = getCurrentFormattedDate();

const createMeetingErrorOut = document.getElementById('create-meeting-error-out');

document.getElementById('btn-create-meeting').addEventListener('click', () => {
    createMeetingErrorOut.display = "none";
    if (onlyLetters(createMeetingName.value) && datetimeInTheFuture(createMeetingDatetime.value)) {
        fetch('/createMeeting', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            },
            body: JSON.stringify({
                'name': createMeetingName.value,
                'datetime': new Date(createMeetingDatetime.value),
            })
        }).then(response => {
            if (response.status === 201) {
                let json = response.json();
                json.then(data => {
                    const meeting = awaitMeetingsList.addMeeting(data.id, createMeetingName.value, new Date(createMeetingDatetime.value), false)
                    createMeetingName.value = "";
                    createMeetingDatetime.value = getCurrentFormattedDate();
                    createMeetingController.defaultSettings();
                    meeting.click();
                    Pages.openPageFromID(data.id);
                })
            } else {
                createMeetingErrorOut.innerText = "Error create meeting";
                createMeetingErrorOut.display = "unset";
            }
        })
    }
})

function getCurrentFormattedDate() {
    let currentDate = new Date();

    let year = currentDate.getFullYear();
    let month = ("0" + (currentDate.getMonth() + 1)).slice(-2);
    let day = ("0" + currentDate.getDate()).slice(-2);
    let hours = ("0" + currentDate.getHours()).slice(-2);
    let minutes = ("0" + (currentDate.getMinutes())).slice(-2);

    return year + "-" + month + "-" + day + "T" + hours + ":" + minutes;
}