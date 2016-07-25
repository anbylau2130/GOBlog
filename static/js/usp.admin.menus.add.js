(function() {
    usp.namespace("usp.admin.menus.add");
    usp.admin.menus.add.viewModel = {
        ID: ko.observable(""),
        Direction: ko.observable(""),
        Parent: ko.observable(""),
        Name: ko.observable(""),
        Icon: ko.observable(""),
        Clazz: ko.observable(""),
        Area: ko.observable(""),
        Controller: ko.observable(""),
        Method: ko.observable(""),
        Parameter: ko.observable(""),
        URL: ko.observable(""),
        Reserve: ko.observable(""),
        Remark: ko.observable(""),
        Creator: ko.observable(""),
        CreateTime: ko.observable(""),
        Auditor: ko.observable(""),
        AuditTime: ko.observable(""),
        Canceler: ko.observable(""),
        CancelTime: ko.observable(""),
    };
    usp.admin.menus.add.init = function() {
        ko.applyBindings(usp.admin.menus.add.viewModel);
    };
    usp.admin.menus.add.goBack = function() {
        window.location.href = "/admin/Menus/List"
    };
    usp.admin.menus.add.save = function() {
        $.ajax({
            url: "Add",
            data: ko.toJS(usp.admin.menus.add.viewModel),
            async: false,
            success: function(data) {
                if (data.status) {
                    usp.Notify('系统信息', '保存成功', 'success')
                    usp.admin.menus.add.goBack()
                } else {
                    usp.Notify('系统信息', data.info, 'alert')
                }
            }
        })
    }
})(this);