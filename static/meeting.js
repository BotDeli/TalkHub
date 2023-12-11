const chatMessageInput = document.getElementById('chat-message-input');

const sendMessageFromChat = () => {
    if (chatMessageInput.value !== "") {
        const text = chatMessageInput.value;
        socket.send(JSON.stringify({
            'recipient': "",
            'action': "-1",
            'data': JSON.stringify(text),
        }));
        addMessageToMessagesList("You", text);
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
const Actives = new VideoAudioController(true, true);

document.getElementById('btn-connect-to-meeting').addEventListener('click', () => {
    Stream.initSocket();
    document.getElementById('pre-connect-panel').style.display = 'none';
});

document.getElementById('btn-end-meeting').addEventListener('click', () => {
    fetch(window.location.pathname+'/endMeeting', {
        method: 'DELETE',
    }).then(response => {
        if (response.status === 202) {
            socket.send(JSON.stringify({
                'recipient': "",
                'action': "5",
                'data': "",
            }));
            window.location.replace('/hub');
        } else {
            alert("Only the owner can complete the meeting");
        }
    })
});