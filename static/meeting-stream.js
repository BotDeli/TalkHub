const streamWindowController = new StreamWindowController();

let webrtcConfig;
function getWebrtcConfig() {
    fetch('/webrtcConfig', {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
        }
    })
        .then(response => {
          if (response.status === 200) {
              let json = response.json();
              json.then(data => {
                  if (data) {
                      webrtcConfig = data;
                  }
              });
          }
        })
}

getWebrtcConfig();

const offerOptions = {
    offerToReceiveVideo: 1,
    offerToReceiveAudio: 1,
};

let peers = {};

let localStream;
let socket;
class StreamChannel {
    constructor() {}
    initSocket() {
        start().then(() => {
            socket = new WebSocket(`ws://${window.location.host}${window.location.pathname}/stream`)

            socket.onopen = (event) => {
                console.log('socket opened');
            };

            socket.onclose = (event) => {
                window.location.reload();
            };

            socket.onmessage = (event) => {
                const msg = JSON.parse(event.data);
                if (msg.action === '0') {
                    console.log(`remove ${msg.sender} user`);
                    peers[msg.sender] = null;
                    streamWindowController.endStream(msg.sender);
                } else if (msg.action === '1') {
                    console.log(`getting ${msg.sender} user request code 1`);
                    const peer = new RTCPeerConnection(webrtcConfig);
                    peer.onicecandidate = (event) => {
                        socket.send(JSON.stringify({
                            'sender': userID,
                            'username': username,
                            'recipient': '',
                            'action': '4',
                            'data': JSON.stringify(event.candidate),
                        }));
                    };

                    localStream.getTracks().forEach(track => peer.addTrack(track, localStream));

                    peer.ontrack = (event) => {
                         console.log('track');
                        streamWindowController.startNewStream(msg.sender, msg.username, event.streams[0]);
                    };

                    peers[msg.sender] = peer;

                    sendOffer(peer, msg.sender);
                } else if (msg.action === '2') {
                    const offer = JSON.parse(msg.data);
                    console.log(`getting offer from ${msg.sender} user request code 2`);

                    const peer = new RTCPeerConnection(webrtcConfig);
                    peer.onicecandidate = (event) => {
                        socket.send(JSON.stringify({
                            'sender': userID,
                            'username': username,
                            'recipient': '',
                            'action': '4',
                            'data': JSON.stringify(event.candidate),
                        }));
                    };

                    localStream.getTracks().forEach(track => peer.addTrack(track, localStream));

                    peer.ontrack = (event) => {
                        console.log('track');
                        streamWindowController.startNewStream(msg.sender, msg.username, event.streams[0]);
                    };

                    peers[msg.sender] = peer;


                    sendAnswer(peer, msg.sender, offer);
                } else if (msg.action === '3') {
                    const answer = JSON.parse(msg.data);
                    console.log(`getting answer from ${msg.sender} user request code 3`);
                    let peer = peers[msg.sender];
                    peer.setRemoteDescription(answer);
                } else if (msg.action === '4') {
                    const candidate = JSON.parse(msg.data);
                    console.log(`getting candidate from ${msg.sender} user request code 4`);
                    let peer = peers[msg.sender];
                    if (peer !== undefined) {
                        addIceCandidate(peer, candidate);
                    }
                }
            };
        });
    }
}

async function start() {
    const stream = await navigator.mediaDevices.getUserMedia({video: true, audio: false});
    localStream = stream;
    streamWindowController.startNewStream(userID, username, stream);
}

async function addIceCandidate(peer, candidate) {
    return await peer.addIceCandidate(candidate);
}

async function sendOffer(peer, recipient) {
    console.log(`send offer to ${recipient} user`);
    const offer = await peer.createOffer(offerOptions);
    await peer.setLocalDescription(offer);

    socket.send(JSON.stringify({
        'sender': userID,
        'username': username,
        'recipient': recipient,
        'action': '2',
        'data': JSON.stringify(offer),
    }));
}

async function sendAnswer(peer, recipient, offer) {
    await peer.setRemoteDescription(offer);

    const answer = await peer.createAnswer();
    await peer.setLocalDescription(answer);

    console.log(`send answer to ${recipient} user`);
    socket.send(JSON.stringify({
        'sender': userID,
        'username': username,
        'recipient': recipient,
        'action': '3',
        'data': JSON.stringify(answer),
    }));
}