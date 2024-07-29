/* /web/static/js/scripts.js */
$(document).ready(function() {
    $('#message-form').on('submit', function(e) {
        e.preventDefault();

        let senderId = $('#sender-id').val();
        let receiverId = $('#receiver-id').val();
        let messageContent = $('#message-content').val();

        $.ajax({
            url: '/messages',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({
                sender_id: senderId,
                receiver_id: receiverId,
                content: messageContent
            }),
            success: function(data) {
                $('#message-content').val('');
                appendMessage(data.sender_id, data.receiver_id, data.content, true);
            }
        });
    });

    function appendMessage(senderId, receiverId, content, isSender) {
        let messageClass = isSender ? 'sender' : 'receiver';
        let messageHtml = `
            <div class="chat-message ${messageClass}">
                <div class="message-content">${content}</div>
                <div class="message-time">${new Date().toLocaleTimeString()}</div>
            </div>
        `;
        $('#chat-container').append(messageHtml);
        $('#chat-container').scrollTop($('#chat-container')[0].scrollHeight);
    }

    // Load initial messages
    function loadMessages(senderId, receiverId) {
        $.ajax({
            url: `/messages/${senderId}/${receiverId}`,
            method: 'GET',
            success: function(data) {
                data.forEach(function(message) {
                    appendMessage(message.sender_id, message.receiver_id, message.content, message.sender_id === senderId);
                });
            }
        });
    }

    // Example usage
    let currentUserId = '1';
    let chattingWithUserId = '2';
    loadMessages(currentUserId, chattingWithUserId);
});
