(function() {
    usp.namespace("usp.admin.chatrooms");

    usp.admin.chatrooms.init = function() {
        ko.applyBindings();
    };
    usp.admin.chatrooms.add = function() {
        window.location.href = "/admin/ChatRoom/Create"
    };
    usp.admin.chatrooms.enter = function() {
        uuid = $("#uuid").val()
        window.location.href = "/admin/ChatRoom/Join?uuid=" + $("#uuid").val()
    };
})(this);