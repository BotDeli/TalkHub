const awaitMeetings = document.getElementById('await-meetings-list');
const Pages = new PageController();

class MeetingAwaitMeetingsList {
    constructor() {
        awaitMeetings.innerHTML = "";
        this.meetings = new Map();
        this.meetingsName = new Map();
        this.meetingsDate = new Map();
        this.meetingsTime = new Map();
    }

    addMeeting(id, name, datetime, started) {
        if (this.meetings.has(id)) {
            return;
        }

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
        this.meetingsName.set(id, div)

        row = document.createElement('div');
        row.className = "await-meeting-row";

        div = document.createElement('div');
        setDateText(div, datetime)
        row.appendChild(div);
        this.meetingsDate.set(id, div)

        div = document.createElement('div');
        setTimeText(div, datetime);
        row.appendChild(div);
        meeting.appendChild(row);
        this.meetingsTime.set(id, div)

        meeting.addEventListener('click', () => {
            if (started) {
                redirectToMeetingCode(id);
            } else {
                Pages.addPageMeeting(id, name, datetime);
            }
        });

        awaitMeetings.appendChild(meeting);
        this.meetings.set(id, meeting)
        return meeting;
    }

    removeMeeting(id) {
        if (!this.meetings.has(id)) {
            return;
        }

        const meeting = this.meetings.get(id);
        meeting.remove();
        this.meetings.delete(id);
        this.meetingsName.delete(id);
        this.meetingsDate.delete(id);
        this.meetingsTime.delete(id);
    }
}