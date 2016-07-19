(function() {
    usp.namespace("usp.admin.menus.add");
    usp.admin.menus.add.viewModel = {
            ID         :ko.observable("ID"),
            Direction  :ko.observable("Direction"),
            Parent     :ko.observable("Parent"),
            Name       :ko.observable("Name"),
            Icon       :ko.observable("Icon"),
            Clazz      :ko.observable("Clazz"),
            Area       :ko.observable("Area"),
            Controller :ko.observable("Controller"),
            Method     :ko.observable("Method"),
            Parameter  :ko.observable("Parameter"),
            URL        :ko.observable("URL"),     
            Reserve    :ko.observable("Reserve"),   
            Remark     :ko.observable("Remark"),
            Creator    :ko.observable("Creator"),
            CreateTime :ko.observable("CreateTime"),
            Auditor    :ko.observable("Auditor"),
            AuditTime  :ko.observable("AuditTime"),
            Canceler   :ko.observable("Canceler"),
            CancelTime :ko.observable("CancelTime"),
        };
    usp.admin.menus.add.init = function() {
      
    };
    usp.admin.menus.add.goBack = function() {
        window.location.href = "/admin/Menus/List"
    };
    usp.admin.menus.add.save = function() {
        console.log(ko.toJS(usp.admin.menus.add.viewModel))
        $.ajax({
            url: "Add",
            data: ko.toJS(usp.admin.menus.add.viewModel),
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