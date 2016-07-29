(function() {
    usp.namespace("usp.admin.corp.edit");
    usp.admin.corp.edit.viewModel = null
    usp.admin.corp.edit.init = function(id) {
        $.ajax({
            url: "GetModel",
            data: {
                id: id
            },
            async: false,
            success: function(response) {
                if (response.success) {
                    usp.admin.corp.edit.viewModel = ko.observable(response.data)
                }
            }
        });
        ko.applyBindings(usp.admin.corp.edit.viewModel);
    };
    usp.admin.corp.edit.goBack = function() {
        window.location.href = "/admin/corp/List"
    }
    usp.admin.corp.edit.save = function() {
        $.ajax({
            url: "Edit",
            data: ko.toJS(usp.admin.corp.edit.viewModel),
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