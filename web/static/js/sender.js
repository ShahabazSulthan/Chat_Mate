$(document).ready(function() {
    const chatContainer = $('#chat-container');

    function appendMessage(name, content, type) {
        const messageClass = type === 'sender' ? 'sender' : 'receiver';
        const messageHtml = `<div class="message ${messageClass}"><strong>${name}:</strong> ${content}</div>`;
        chatContainer.append(messageHtml);
        chatContainer.scrollTop(chatContainer[0].scrollHeight);
    }

    $('#message-form').on('submit', function(e) {
        e.preventDefault();
        const senderId = $('#sender-id').val();
        const receiverId = $('#receiver-id').val();

        // Load messages between sender and receiver
        $.get(`/messages/${senderId}/${receiverId}`, function(messages) {
            chatContainer.empty();
            messages.forEach(function(message) {
                const type = message.sender_id === senderId ? 'sender' : 'receiver';
                const name = type === 'sender' ? 'You' : 'Receiver';
                appendMessage(name, message.content, type);
            });
        });
    });
});

