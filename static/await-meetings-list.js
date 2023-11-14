const awaitMeetings = document.getElementById('await-meetings-list');
const Pages = new PageController();

class AwaitMeetingsList {
    constructor() {
        awaitMeetings.innerHTML = "";
    }

    addMeeting(id, name, datetime, started) {
        const meeting = document.createElement('div');
        if (started) {
            meeting.className = "await-meeting await-meeting-started";
        } else {
            meeting.className = "await-meeting";
        }

        let div = document.createElement('div');
        div.className = "await-meeting-id";
        div.innerText = '#' + id;
        meeting.appendChild(div);

        let row = document.createElement('div');
        row.className = "await-meeting-row";

        div = document.createElement('div');
        div.className = "await-meeting-name";
        div.innerText = name;
        row.appendChild(div);
        meeting.appendChild(row);

        row = document.createElement('div');
        row.className = "await-meeting-row";

        div = document.createElement('div');
        setDateText(div, datetime)
        row.appendChild(div);

        div = document.createElement('div');
        setTimeText(div, datetime);
        row.appendChild(div);
        meeting.appendChild(row);

        meeting.addEventListener('click', () => {
            Pages.addPageMeeting(id, name, datetime);
        });

        awaitMeetings.appendChild(meeting);
    }
}