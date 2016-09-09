(function() {
    usp.namespace("usp.admin.privilege.edit");
    usp.admin.privilege.edit.viewModel = null
    usp.admin.privilege.edit.init = function(id) {
        $.ajax({
            url: "GetModel",
            data: {
                id: id
            },
            async: false,
            success: function(response) {
                if (response.success) {
                    usp.admin.privilege.edit.viewModel = ko.observable(response.data)
                }
            }
        });
        ko.applyBindings(usp.admin.privilege.edit.viewModel);
    };
    usp.admin.privilege.edit.goBack = function() {
        window.location.href = "/admin/privilege/List"
    }
    usp.admin.privilege.edit.save = function() {
        $.ajax({
            url: "Edit",
            data: ko.toJS(usp.admin.privilege.edit.viewModel),
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