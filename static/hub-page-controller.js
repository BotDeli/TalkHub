const pagesList = document.getElementById('description-pages-list');
const pageInfo = document.getElementById('description-pages-show-out-info');
const pageShowOut = document.getElementById('description-pages-show-out-meeting');

const pageShowOutId = document.getElementById('description-settings-meeting-id');
const pageShowOutName = document.getElementById('description-settings-meeting-name');
const pageShowOutDate = document.getElementById('description-settings-meeting-date');
const pageShowOutTime = document.getElementById('description-settings-meeting-time');


const btnStartMeeting = document.getElementById('btn-settings-meeting-start-meeting');
const btnCancelMeeting = document.getElementById('btn-settings-meeting-cancel-meeting');
document.getElementById('description-page-info').addEventListener("click", () => {
    pageShowOut.style.display = "none";
    pageInfo.style.display = "flex";
});

let openedPages = {};

class PageController {
    constructor() {}

    addPageMeeting(id, name, datetime) {
        if (openedPages[id]) {
            return;
        }
        openedPages[id] = true;
        const page = document.createElement('div');
        page.className = 'description-page description-page-meeting';
        page.addEventListener("click", () => {
            pageShowOutId.innerText = '#' + id;
            pageShowOutName.innerText = name;

            setDateText(pageShowOutDate, datetime);
            setTimeText(pageShowOutTime, datetime);

            pageShowOut.style.display = "flex";
            pageInfo.style.display = "none";
        });

        btnStartMeeting.onclick = () => {
            fetch("/startMeeting", {
                method: "UPDATE",
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json'
                },
                body: JSON.stringify({
                    'id': id,
                })
            }).then((response) => {
                if (response.status === 202) {
                    redirectToMeetingCode(id);
                } else {
                    alert(`Error updating meeting, status ${response.status}`);
                }
            })
        }

        btnCancelMeeting.onclick = () => {
            console.log("end", id);
        }

        let div = document.createElement('div');
        div.innerText = name;
        page.appendChild(div);

        div = document.createElement('div');
        div.className = 'btn-description-page-close';
        div.addEventListener("click", (e) => {
            e.stopPropagation();
            page.remove();
            openedPages[id] = false;
        });
        page.appendChild(div);

        pagesList.appendChild(page);
    }
}

function setDateText(node, datetime) {
    let year = datetime.getFullYear();
    let month = ("0" + (datetime.getMonth() + 1)).slice(-2);
    let day = ("0" + datetime.getDate()).slice(-2);
    node.innerText = `${day}-${month}-${year}`;
}

function setTimeText(node, datetime) {
    let hours = ("0" + datetime.getHours()).slice(-2);
    let minutes = ("0" + (datetime.getMinutes())).slice(-2);
    node.innerText = `${hours}:${minutes}`;
}