new ControllerPopupFromInputCells('input-cell input-cell-meeting-code', 'input-cell-popup-text input-cell-popup-text-meeting-code');

const meetingCode = document.getElementById('connect-to-meeting-input-code');

document.getElementById('btn-connect-to-meeting').addEventListener('click', () => {
    redirectToMeetingCode(meetingCode.value);
});

function redirectToMeetingCode(id) {
    window.location.replace("/meeting/"+id);
}