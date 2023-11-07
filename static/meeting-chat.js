class ChatChannel {
    constructor(messagesOut) {
        this.messagesOut = messagesOut;
    }

    sendMessage(sender, message) {
        // TODO: отправка сообщения на сервер

        let msg = document.createElement('div');
        msg.className = "chat-message";
        msg.innerHTML = "<strong>"+sender+"</strong>"+": "+message;
        this.messagesOut.appendChild(msg);
    }
}