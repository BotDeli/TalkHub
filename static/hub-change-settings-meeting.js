new ControllerPopupFromInputCells('input-cell input-cell-change-settings-meeting', 'input-cell-popup-text input-cell-popup-text-change-settings-meeting');

const changeMeetingName = document.getElementById('change-settings-meeting-input-name');
changeMeetingName.value = changeMeetingName.defaultValue;

const changeMeetingDatetime = document.getElementById('change-settings-meeting-input-datetime');
changeMeetingDatetime.value = changeMeetingDatetime.defaultValue;

const btnChangeMeetingName = document.getElementById('btn-change-settings-meeting-name');
const btnChangeMeetingDatetime = document.getElementById('btn-change-settings-meeting-datetime');