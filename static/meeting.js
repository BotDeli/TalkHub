const streamUsersList = document.getElementById('stream-users-list');

const userID = '1';

Synchronizer = new StreamSynchronizer();
// for (var i = 0; i < 12; i++) {
    const myVideoOutput = document.createElement('video');
    myVideoOutput.autoplay = true;
    myVideoOutput.className = 'stream-user';
    streamUsersList.appendChild(myVideoOutput);
    Synchronizer.addVideoOutput(userID, myVideoOutput);
// }

Synchronizer.synchronize();

window.addEventListener('resize', () => {
    Synchronizer.synchronize();
});

const btnStreamMicrophone = document.getElementById('btn-stream-microphone');

changerAudioStream = function() {
    let activated = false;
    return () => {
        if (activated) {
            Synchronizer.closeAudioStream(userID);
            btnStreamMicrophone.className = 'stream-functionality-button stream-functionality-microphone-off'
        } else {
            Synchronizer.openAudioStream(userID);
            btnStreamMicrophone.className = 'stream-functionality-button stream-functionality-microphone-on'
        }
        activated = !activated;
    }
}();

const btnStreamWebCamera = document.getElementById('btn-stream-web-camera');

changerVideoStream = function() {
    let activated = false;
    return () => {
        if (activated) {
            Synchronizer.closeVideoStream(userID);
            btnStreamWebCamera.className = 'stream-functionality-button stream-functionality-web-camera-off'
        } else {
            Synchronizer.openVideoStream(userID);
            btnStreamWebCamera.className = 'stream-functionality-button stream-functionality-web-camera-on'
        }
        activated = !activated;
    }
}();

btnStreamMicrophone.addEventListener('click', () => {
    changerAudioStream();
})

btnStreamWebCamera.addEventListener('click', () => {
    changerVideoStream();
})


const chatMessagesList = document.getElementById('chat-messages-list');
const Chat = new ChatChannel(chatMessagesList);
// Chat.sendMessage("testSender)))", "HELLO WORLD!!!");

const chatMessageInput = document.getElementById('chat-message-input');

document.getElementById('btn-chat-message-send-message').addEventListener('click', () => {
    if (chatMessageInput.value !== "") {
        Chat.sendMessage("Вы", chatMessageInput.value);
        chatMessageInput.value = "";
    }
});