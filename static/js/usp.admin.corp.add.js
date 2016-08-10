(function() {
    usp.namespace("usp.admin.corp.add");
    usp.admin.corp.add.viewModel = {
        Name: ko.observable(""),
        Type: ko.observable(""),
        adminLoginName: ko.observable(""),
        adminPassWord: ko.observable(""),
    };
    usp.admin.corp.add.init = function() {
        ko.applyBindings(usp.admin.corp.add.viewModel);
    };
    usp.admin.corp.add.goBack = function() {
        window.location.href = "/admin/Corp/List"
    };
    usp.admin.corp.add.save = function() {
        $.ajax({
            url: "Add",
            data: ko.toJS(usp.admin.corp.add.viewModel),
            async: false,
            success: function(data) {
                if (data.status) {
                    usp.Notify('系统信息', '保存成功', 'success')
                    usp.admin.corp.add.goBack()
                } else {
                    usp.Notify('系统信息', data.info, 'alert')
                }
            }
        })
        return false;
    }
})(this);