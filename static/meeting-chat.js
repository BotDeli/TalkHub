const chatMessagesList = document.getElementById('chat-messages-list');

function addMessageToMessagesList(sender, text) {
    let message = document.createElement('div');
    message.className = "chat-message";
    message.innerHTML = `<strong>${sender}</strong>: ${text}`;
    chatMessagesList.appendChild(message);
}

class ChatChannel {
    constructor() {
        this.socket = null;
    }

    initSocket() {
        const socket = new WebSocket(`ws://${window.location.host}${window.location.pathname}/chat`);

        socket.onmessage = (event) => {
            let data = JSON.parse(event.data)
            addMessageToMessagesList(data.sender, data.text);
        };

        socket.onclose = () => {
            window.location.reload();
        };

        this.socket = socket;
    }

    sendMessage(sender, text) {
        if (this.socket === null) {
            return;
        }

        this.socket.send(JSON.stringify({
            "sender": sender,
            "text": text,
        }));
    }
}