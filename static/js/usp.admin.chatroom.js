(function() {
    usp.namespace("usp.admin.chatroom");
    usp.admin.chatroom.init = function() {
        socket = new WebSocket('ws://' + window.location.host + '/admin/ChatRoom/SendMsg?uuid=' + $('#uuid').val());
        // Message received on the socket
        socket.onmessage = function(event) {
            var data = JSON.parse(event.data);
            //console.log(data);
            switch (data.Type) {
                case 0: // JOIN
                    if (data.User == $('#uname').text()) {
                        $("#chatbox li").first().before("<li>您已经进入房间.</li>");
                    } else {
                        $("#chatbox li").first().before("<li>" + data.User.RealName + " 进入房间.</li>");
                    }
                    break;
                case 1: // LEAVE
                    $("#chatbox li").first().before("<li>" + data.User.RealName + " 离开房间.</li>");
                    break;
                case 2: // MESSAGE
                    $("#chatbox li").first().before("<li><b>" + data.User.RealName + "</b>: " + data.Content + "</li>");
                    break;
            }
        };

        // Send messages.
        var postConecnt = function() {
            var uname = $('#uname').text();
            var content = $('#sendbox').val();
            socket.send(content);
            $('#sendbox').val("");
        }

        $('#sendbtn').click(function() {
            postConecnt();
        });
    };
})(this);