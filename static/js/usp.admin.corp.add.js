(function() {
    usp.namespace("usp.admin.corp.add");
    usp.admin.corp.add.viewModel = {
        ID: ko.observable(""),
        Parent: ko.observable(""),
        Priority: ko.observable(""),
        Name: ko.observable(""),
        BriefName: ko.observable(""),
        Certificate: ko.observable(""),
        CertificateNumber: ko.observable(""),
        Ceo: ko.observable(""),
        Postcode: ko.observable(""),
        Faxcode: ko.observable(""),
        Linkman: ko.observable(""),
        Mobile: ko.observable(""),
        Email: ko.observable(""),
        Qq: ko.observable(""),
        Wechat: ko.observable(""),
        Weibo: ko.observable(""),
        VirtualIntegral: ko.observable(""),
        RealIntegral: ko.observable(""),
        FeeType: ko.observable(""),
        PrepayValve: ko.observable(""),
        Balance: ko.observable(""),
        FrozenBalance: ko.observable(""),
        IncomingBalance: ko.observable(""),
        Commission: ko.observable(""),
        Discount: ko.observable(""),
        Province: ko.observable(""),
        Area: ko.observable(""),
        County: ko.observable(""),
        Community: ko.observable(""),
        Address: ko.observable(""),
        Status: ko.observable(""),
        Type: ko.observable(""),
        Grade: ko.observable(""),
        Vocation: ko.observable(""),
        Reserve: ko.observable(""),
        Remark: ko.observable(""),
        Creator: ko.observable(""),
        CreateTime: ko.observable(""),
        Auditor: ko.observable(""),
        AuditTime: ko.observable(""),
        Canceler: ko.observable(""),
        CancelTime: ko.observable(""),
        LogoIcon: ko.observable(""),
        LogoUrl: ko.observable(""),
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