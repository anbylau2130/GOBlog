(function() {
    usp.namespace("usp.admin.user.edit");
    usp.admin.user.edit.viewModel = null
    usp.admin.user.edit.init = function(id) {
        $.ajax({
            url: "GetModel",
            data: {
                id: id
            },
            async: false,
            success: function(response) {
                if (response.success) {
                    usp.admin.user.edit.viewModel = ko.observable(response.data)
                }
            }
        });
        ko.applyBindings(usp.admin.user.edit.viewModel);
    };
    usp.admin.user.edit.goBack = function() {
        window.location.href = "/admin/user/List"
    }
    usp.admin.user.edit.save = function() {
        $.ajax({
            url: "Edit",
            data: ko.toJS(usp.admin.user.edit.viewModel),
            async: false,
            success: function(data) {
                if (data.status) {
                    usp.Notify('系统信息', '保存成功', 'success')
                    usp.admin.menus.edit.goBack()
                } else {
                    usp.Notify('系统信息', data.info, 'alert')
                }
            }
        })
        return false;
    }

})(this);