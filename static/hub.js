document.addEventListener("DOMContentLoaded", () => {
   fetch("/getMyMeetings", {
       method: "GET",
       headers: {
           'Content-Type': 'application/json',
           'Accept': 'application/json'
       },
   }).then(response => {
       if (response.status === 200) {
           let json = response.json();
           json.then(data => {
               data.meetings.forEach(meeting => {
                   awaitMeetingsList.addMeeting(meeting.id, meeting.name, new Date(meeting.date));
               });
           })
       }
   })
});