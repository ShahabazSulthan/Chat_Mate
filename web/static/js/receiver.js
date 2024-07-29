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
        const senderName = $('#sender-name').val();
        const receiverName = $('#receiver-name').val();
        const senderId = $('#sender-id').val();
        const receiverId = $('#receiver-id').val();
        const content = $('#message-content').val();

        const messageData = {
            sender_id: senderId,
            receiver_id: receiverId,
            content: content,
            timestamp: Date.now()
        };

        $.ajax({
            url: '/messages',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(messageData),
            success: function() {
                appendMessage(senderName, content, 'sender');
                $('#message-content').val('');

                // Load updated messages between sender and receiver
                $.get(`/messages/${senderId}/${receiverId}`, function(messages) {
                    chatContainer.empty();
                    messages.forEach(function(message) {
                        const type = message.sender_id === senderId ? 'sender' : 'receiver';
                        const name = type === 'sender' ? senderName : receiverName;
                        appendMessage(name, message.content, type);
                    });
                });
            }
        });
    });
});
