(function() {
    usp.namespace("usp.admin.user.add");
    usp.admin.user.add.viewModel = {
        LoginName: ko.observable(""),
        RealName: ko.observable(""),
        Password: ko.observable(""),
        PasswordConfirm: ko.observable(""),
        Role: ko.observable("")
    };
    usp.admin.user.add.init = function() {
        ko.applyBindings(usp.admin.user.add.viewModel);
        usp.bindSelect("#role", "/Service/RoleSelect")
    };
    usp.admin.user.add.goBack = function() {
        window.location.href = "/admin/User/List"
    };
    usp.admin.user.add.save = function() {
        var selectedRoles = "";
        var selectedArr = $("#role").select2("data")
        $.each(selectedArr, function(i, value) {
            if (i == selectedArr.length - 1)
                selectedRoles += value.id
            else
                selectedRoles += value.id + ','
        })
        console.log(ko.toJS(usp.admin.user.add.viewModel))
        usp.admin.user.add.viewModel.Role = ko.observable(selectedRoles);
        $.ajax({
            url: "Add",
            data: ko.toJS(usp.admin.user.add.viewModel),
            async: false,
            success: function(data) {
                if (data.status) {
                    usp.Notify('系统信息', '保存成功', 'success')
                    usp.admin.user.add.goBack()
                } else {
                    usp.Notify('系统信息', data.info, 'alert')
                }
            }
        })
        return false;
    }
})(this);