new ControllerPopupFromInputCells('input-cell input-cell-create-meeting', 'input-cell-popup-text input-cell-popup-text-create-meeting');

const meetingName = document.getElementById('create-meeting-input-name');
const meetingDatetime = document.getElementById('create-meeting-input-datetime');

const awaitMeetingsList = new AwaitMeetingsList();
meetingDatetime.value = getCurrentFormattedDate();

const errorOut = document.getElementById('create-meeting-error-out');

awaitMeetingsList.addMeeting('id222', "meetingName.value", new Date())
awaitMeetingsList.addMeeting('id11111', "meetingName2222", new Date())

document.getElementById('btn-create-meeting').addEventListener('click', () => {
    errorOut.display = "none";
    if (onlyLetters(meetingName.value) && datetimeInTheFuture(meetingDatetime.value)) {
        fetch('/createMeeting', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            },
            body: JSON.stringify({
                'name': meetingName.value,
                'datetime': new Date(meetingDatetime.value),
            })
        }).then(response => {
            if (response.status === 201) {
                let json = response.json();
                json.then(data => {
                    awaitMeetingsList.addMeeting(data.id, meetingName.value, new Date(meetingDatetime.value))
                })
            } else {
                errorOut.innerText = "Error create meeting";
                errorOut.display = "unset";
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