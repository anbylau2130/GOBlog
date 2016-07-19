(function() {
    usp.namespace("usp.admin.menus.edit");
    usp.admin.menus.edit.viewModel = null
    usp.admin.menus.edit.init = function(id) {
        $.ajax({
            url: "GetModel",
            data: {
                id: id
            },
            async: false,
            success: function(response) {
                if (response.success) {
                    usp.admin.menus.edit.viewModel = ko.observable(response.data)
                }
            }
        });
        ko.applyBindings(usp.admin.menus.edit.viewModel);
    };
    usp.admin.menus.edit.goBack = function() {
        window.location.href = "/admin/Menus/List"
    }
    usp.admin.menus.edit.save = function() {
        console.log(ko.toJS(usp.admin.menus.edit.viewModel))
        $.ajax({
            url: "Edit",
            data: ko.toJS(usp.admin.menus.edit.viewModel),
            async: false,
            success: function(data) {
                if (data.status) {
                    usp.Notify('系统信息', '保存成功', 'success')
                } else {
                    usp.Notify('系统信息', data.info, 'alert')
                }
            }
        })
    }

})(this);