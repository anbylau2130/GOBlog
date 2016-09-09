(function() {
    usp.namespace("usp.admin.privilege.add");
    usp.admin.privilege.add.viewModel = {
        ID: ko.observable(""),
        Parent: ko.observable(""),
        Menu: ko.observable(""),
        Name: ko.observable(""),
        Clazz: ko.observable(""),
        Area: ko.observable(""),
        Controller: ko.observable(""),
        Method: ko.observable(""),
        Parameter: ko.observable(""),
        Url: ko.observable(""),
        Reserve: ko.observable(""),
        Remark: ko.observable(""),
        Creator: ko.observable(""),
        CreateTime: ko.observable(""),
        Auditor: ko.observable(""),
        AuditTime: ko.observable(""),
        Canceler: ko.observable(""),
        CancelTime: ko.observable(""),
    };
    usp.admin.privilege.add.init = function() {
        ko.applyBindings(usp.admin.privilege.add.viewModel);
    };
    usp.admin.privilege.add.goBack = function() {
        window.location.href = "/admin/privilege/List"
    };
    usp.admin.privilege.add.save = function() {
        $.ajax({
            url: "Add",
            data: ko.toJS(usp.admin.privilege.add.viewModel),
            async: false,
            success: function(data) {
                if (data.status) {
                    usp.Notify('系统信息', '保存成功', 'success')
                    usp.admin.privilege.add.goBack()
                } else {
                    usp.Notify('系统信息', data.info, 'alert')
                }
            }
        })
        return false;
    }
})(this);