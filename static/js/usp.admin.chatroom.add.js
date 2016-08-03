(function() {
    usp.namespace("usp.admin.chatroom.add");
    usp.admin.chatroom.add.viewModel = {
        ID: ko.observable(""),
        UUID: ko.observable(""),
        PicURL: ko.observable(""),
        Name: ko.observable(""),
        Title: ko.observable(""),
        Reserve: ko.observable(""),
        Remark: ko.observable(""),
        Creator: ko.observable(""),
        CreateTime: ko.observable(""),
        Auditor: ko.observable(""),
        AuditTime: ko.observable(""),
        Canceler: ko.observable(""),
        CancelTime: ko.observable(""),
    };
    usp.admin.chatroom.add.init = function() {
        ko.applyBindings(usp.admin.chatroom.add.viewModel);
    };
    usp.admin.chatroom.add.goBack = function() {
        window.location.href = "/admin/ChatRoom/ChatRooms"
    };
    usp.admin.chatroom.add.save = function() {
        $.ajax({
            url: "Create",
            data: ko.toJS(usp.admin.chatroom.add.viewModel),
            async: false,
            success: function(data) {
                if (data.status) {
                    usp.Notify('系统信息', '保存成功', 'success')
                    usp.admin.chatroom.add.goBack()
                } else {
                    usp.Notify('系统信息', data.info, 'alert')
                }
            }
        })
        return false;
    }
})(this);