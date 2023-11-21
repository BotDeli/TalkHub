
// Synchronizer = new StreamSynchronizer();
// for (var i = 0; i < 12; i++) {

    // Synchronizer.addVideoOutput(userID, myVideoOutput);
// }

// Synchronizer.synchronize();

// window.addEventListener('resize', () => {
    // Synchronizer.synchronize();
// });


const Chat = new ChatChannel();

const chatMessageInput = document.getElementById('chat-message-input');

const sendMessageFromChat = () => {
    if (chatMessageInput.value !== "") {
        Chat.sendMessage(username, chatMessageInput.value);
        chatMessageInput.value = "";
    }
}

document.getElementById('btn-chat-message-send-message').addEventListener('click', sendMessageFromChat);

window.addEventListener('keypress', (e) => {
    if (e.key === "Enter" && chatMessageInput === document.activeElement) {
        sendMessageFromChat();
    }
})


const Stream = new StreamChannel();

document.getElementById('btn-connect-to-meeting').addEventListener('click', () => {
    Chat.initSocket();
    Stream.initSocket();
    document.getElementById('pre-connect-panel').style.display = 'none';
});
