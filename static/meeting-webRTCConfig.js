let webRTCConfig;
getWebrtcConfig();

function getWebrtcConfig() {
    fetch('/webrtcConfig', {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
        }
    }).then(response => {
        if (response && isStatusOK(response)) {
            response.json().then(data => {
                SetWebRTCConfigFromJSONData(data);
            });
        }
    });
}

const statusOK = 200;

function isStatusOK(response) {
    return response.status === statusOK;
}

function SetWebRTCConfigFromJSONData(data) {
    webRTCConfig = {
        iceServers: [
            {urls: data.stun},
            {urls: data.turn, username: data.username, credential: data.credential},
        ],
        codecs: [
            {mimeType: data.audio},
            {mimeType: data.video},
        ]
    };
}