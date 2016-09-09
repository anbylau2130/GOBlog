(function() {
    usp.namespace("usp.admin.role.add");
    usp.admin.role.add.viewModel = {
        Name: ko.observable(""),
        Remark: ko.observable(""),
        Menus: ko.observable(""),
        Privileges: ko.observable("")
    };
    usp.admin.role.add.init = function() {
        ko.applyBindings(usp.admin.role.add.viewModel);
    };
    usp.admin.role.add.goBack = function() {
        window.location.href = "/admin/Role/List"
    };
    usp.admin.role.add.save = function() {
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
        usp.admin.role.add.viewModel.Menus = ko.observable(menus)
        usp.admin.role.add.viewModel.Privileges = ko.observable(privileges)
        $.ajax({
            url: "Add",
            data: ko.toJS(usp.admin.role.add.viewModel),
            async: false,
            success: function(data) {
                if (data.status) {
                    usp.Notify('系统信息', '保存成功', 'success')
                    usp.admin.role.add.goBack()
                } else {
                    usp.Notify('系统信息', data.info, 'alert')
                }
            }
        })
        return false;
    }
})(this);