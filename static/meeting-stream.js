const streamWindowController = new StreamWindowController();

const offerOptions = {
    offerToReceiveVideo: 1,
    offerToReceiveAudio: 1,
};

const wsStreamURL = `ws://${window.location.host}${window.location.pathname}/stream`;

const peers = new Map();

let localStream;
navigator.mediaDevices.getUserMedia({video: true, audio: true})
    .then(stream => {
        localStream = stream;
    })
    .catch(err => {
    localStream = new MediaStream();
    console.error(err);
})

let socket;

class StreamChannel {
    constructor() {}
    initSocket() {
        streamWindowController.startNewStream(userID, username, localStream, true);

        socket = new WebSocket(wsStreamURL);

        socket.onclose = () => {
            window.location.reload();
        };

        socket.onmessage = (event) => {
            const msg = JSON.parse(event.data);

            let peer = peers.get(msg.sender);
            if (!peer) {
                peer = makeNewPeerConnection(msg.sender, msg.username);
            }

            if (msg.action === '-1') {
                addMessageToMessagesList(msg.username, msg.data.slice(1, -1));
            } else if (msg.action === '0') {
                removeUser(msg.sender);
            } else if (msg.action === '1') {
                sendOffer(peer, msg.sender);
            } else if (msg.action === '2') {
                const offer = JSON.parse(msg.data);

                sendAnswer(peer, msg.sender, offer);
            } else if (msg.action === '3') {
                const answer = JSON.parse(msg.data);

                peer.setRemoteDescription(answer);
            } else if (msg.action === '4') {
                const candidate = JSON.parse(msg.data);

                if (peer !== undefined) {
                    addIceCandidate(peer, candidate);
                }
            } else if (msg.action === '5') {
                window.location.replace('/hub');
            }
        };
    }
}

function removeUser(sender) {
    peers.delete(sender);
    streamWindowController.endStream(sender);
}

function makeNewPeerConnection(sender, senderUsername) {
    const peer = new RTCPeerConnection(webRTCConfig);

    peer.onicecandidate = (event) => {
        sendStringifyDataToSocket('4', event.candidate);
    };

    addLocalStreamTracksToPeer(peer);

    peer.ontrack = (event) => {
        streamWindowController.startNewStream(sender, senderUsername, event.streams[0]);
    };

    streamWindowController.startNewStream(sender, senderUsername, new MediaStream());

    peers.set(sender, peer);
    return peer;
}

function sendStringifyDataToSocket(action, data = '', recipient = '') {
    socket.send(JSON.stringify({
        'recipient': recipient,
        'action': action,
        'data': JSON.stringify(data),
    }));
}

function addLocalStreamTracksToPeer(peer) {
    localStream.getTracks().forEach(track => peer.addTrack(track, localStream));
}

async function addIceCandidate(peer, candidate) {
    await peer.addIceCandidate(candidate);
}

async function sendOffer(peer, recipient) {
    const offer = await peer.createOffer(offerOptions);
    await peer.setLocalDescription(offer);

    sendStringifyDataToSocket('2', offer, recipient);
}

async function sendAnswer(peer, recipient, offer) {
    await peer.setRemoteDescription(offer);

    const answer = await peer.createAnswer();
    await peer.setLocalDescription(answer);

    sendStringifyDataToSocket('3', answer, recipient);
}

function updateTrackTransfer(kind, enabled) {
    peers.forEach((key) => {
        const sender = key.getSenders().find(sender => sender.track.kind === kind);
        if (sender) {
            sender.track.enabled = enabled;
        }
    });
}