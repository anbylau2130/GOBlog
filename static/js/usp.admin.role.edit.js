(function() {
    usp.namespace("usp.admin.role.edit");
    usp.admin.role.edit.viewModel = {
        ID: ko.observable(""),
        Name: ko.observable(""),
        Remark: ko.observable(""),
        Menus: ko.observable(""),
        Privileges: ko.observable("")
    }
    usp.admin.role.edit.init = function(id) {
        $.ajax({
            url: "GetModel",
            data: {
                id: id
            },
            async: false,
            success: function(response) {
                if (response.success) {
                    usp.admin.role.edit.viewModel.ID = ko.observable(response.data.ID)
                    usp.admin.role.edit.viewModel.Name = ko.observable(response.data.Name)
                    usp.admin.role.edit.viewModel.Remark = ko.observable(response.data.Remark)
                    usp.admin.role.edit.viewModel.Menus = ko.observable("")
                    usp.admin.role.edit.viewModel.Privileges = ko.observable("")
                }
            }
        });
        ko.applyBindings(usp.admin.role.edit.viewModel);
    };
    usp.admin.role.edit.goBack = function() {
        window.location.href = "/admin/Role/List"
    }
    usp.admin.role.edit.save = function() {
        var data = $('#container').jstree('get_selected', true);
        var menus = "";
        var privileges = "";
        $.each(data, function(index, item) {
            //if (item.original.IsLeaf == "true") {
            if (item.original.type == "menu") {
                menus += item.original.id.replace("m_", "") + ","
            } else if (item.original.type == "privilege") {
                privileges += item.original.id.replace("p_", "") + ","
            }
            //}
        });
        usp.admin.role.edit.viewModel.Menus = ko.observable(menus)
        usp.admin.role.edit.viewModel.Privileges = ko.observable(privileges)
        $.ajax({
            url: "Edit",
            data: ko.toJS(usp.admin.role.edit.viewModel),
            async: false,
            success: function(data) {
                if (data.status) {
                    usp.Notify('系统信息', '保存成功', 'success')
                    usp.admin.role.edit.goBack()
                } else {
                    usp.Notify('系统信息', data.info, 'alert')
                }
            }
        })
        return false;
    }

})(this);