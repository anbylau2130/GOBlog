(function() {
    usp.namespace("usp.admin.corp.list");

    usp.admin.corp.list.add = function name(params) {
        window.location.href = "Add";
    }
    usp.admin.corp.list.edit = function name(params) {
        if (usp.admin.corp.list.table.row('.selected').data())
            window.location.href = "Edit?id=" + usp.admin.corp.list.table.row('.selected').data().ID;
    }

    usp.admin.corp.list.remove = function name(params) {
        if (usp.admin.corp.list.table.row('.selected').data()) {
            $.ajax({
                url: "Del",
                data: {
                    id: usp.admin.corp.list.table.row('.selected').data().ID
                },
                async: false,
                success: function(response) {
                    if (response.status) {
                        usp.admin.corp.list.table.row('.selected').remove().draw(false)
                        usp.Notify('系统信息', response.info, 'success')
                    } else {
                        usp.Notify('系统信息', response.info, 'alert')
                    }
                }
            });
        }
    }
    usp.admin.corp.list.table = null;
    usp.admin.corp.list.init = function(id, url) {
        ko.applyBindings();
        usp.admin.corp.list.table = $(id).DataTable({
            dom: 'Bfrtip',
            iDisplayLength: 10,
            "bServerSide": true,
            "bPaginate": true, //是否允许分页
            "bLengthChange": true, //是否显示每页显示条数
            "bFilter": true, //是否启用条件查询
            "bSort": true, //是否启用列排序
            "bInfo": true, //是否显示分页信息 
            "bScrollCollapse": false, //当显示的数据不足以支撑表格的默认的高度时，依然显示纵向的滚动条。(默认是false)      
            "language": {
                buttons: {
                    //'copy', 'csv', 'excel', 'pdf', 'print'
                    // colvis: 'Change columns',
                    copy: "拷贝",
                    excel: "导出Excel",
                    pdf: "导出PDF",
                    print: '打印'
                },
                "sProcessing": "处理中...",
                "sLengthMenu": "显示 _MENU_ 项结果",
                "sZeroRecords": "没有匹配结果",
                "sInfo": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
                "sInfoEmpty": "显示第 0 至 0 项结果，共 0 项",
                "sInfoFiltered": "(由 _MAX_ 项结果过滤)",
                "sInfoPostFix": "",
                "sSearch": "搜索:",
                "sUrl": "",
                "sEmptyTable": "表中数据为空",
                "sLoadingRecords": "载入中...",
                "sInfoThousands": ",",
                "oPaginate": {
                    "sFirst": "首页",
                    "sPrevious": "上页",
                    "sNext": "下页",
                    "sLast": "末页"
                },
                "oAria": {
                    "sSortAscending": ": 以升序排列此列",
                    "sSortDescending": ": 以降序排列此列"
                }
            },
            ajax: url,

            columnDefs: [{
                title: "ID",
                name: "ID",
                targets: 0,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "父公司",
                name: "Parent",
                targets: 1,
                orderable: true,
                "width": "200px",
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "公司优先级(0-9)",
                name: "Priority",
                targets: 2,
                "width": "20%",
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "公司名称",
                name: "Name",
                targets: 3,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "公司简称",
                name: "BriefName",
                targets: 4,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "证书名称",
                name: "Certificate",
                targets: 5,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "证书号码",
                name: "CertificateNumber",
                targets: 6,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "法人代表",
                name: "Ceo",
                targets: 7,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "邮编",
                name: "Postcode",
                targets: 8,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "传真号",
                name: "Faxcode",
                targets: 9,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "联系人",
                name: "Linkman",
                targets: 10,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "联系手机",
                name: "Mobile",
                targets: 11,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Email",
                name: "Email",
                targets: 12,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "联系QQ",
                name: "Qq",
                targets: 13,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "微信",
                name: "Wechat",
                targets: 14,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "微博",
                name: "Weibo",
                targets: 15,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "虚拟积分",
                name: "VirtualIntegral",
                targets: 16,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "真实积分",
                name: "RealIntegral",
                targets: 17,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "付费类型",
                name: "FeeType",
                targets: 18,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "预付阀值",
                name: "PrepayValve",
                targets: 19,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "余额",
                name: "Balance",
                targets: 20,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "冻结余额",
                name: "FrozenBalance",
                targets: 21,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "在途余额",
                name: "IncomingBalance",
                targets: 22,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "佣金比例",
                name: "Commission",
                targets: 23,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "折扣比例",
                name: "Discount",
                targets: 24,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "省份",
                name: "Province",
                targets: 25,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "地区",
                name: "Area",
                targets: 26,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "县",
                name: "County",
                targets: 27,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "社区",
                name: "Community",
                targets: 28,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "地址",
                name: "Address",
                targets: 29,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "状态",
                name: "Status",
                targets: 30,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "类型",
                name: "Type",
                targets: 31,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "级别",
                name: "Grade",
                targets: 32,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "行业类型",
                name: "Vocation",
                targets: 33,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Reserve",
                name: "Reserve",
                targets: 34,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Remark",
                name: "Remark",
                targets: 35,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "创建人",
                name: "Creator",
                targets: 36,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "创建时间",
                name: "CreateTime",
                targets: 37,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "审核人",
                name: "Auditor",
                targets: 38,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "审核时间",
                name: "AuditTime",
                targets: 39,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "注销人",
                name: "Canceler",
                targets: 40,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "注销时间",
                name: "CancelTime",
                targets: 41,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "公司Icon",
                name: "LogoIcon",
                targets: 42,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "公司logo",
                name: "LogoUrl",
                targets: 43,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, ],

            columns: [{
                data: "ID",

            }, {
                data: "Parent"
            }, {
                data: "Priority"
            }, {
                data: "Name"
            }, {
                data: "BriefName"
            }, {
                data: "Certificate"
            }, {
                data: "CertificateNumber"
            }, {
                data: "Ceo"
            }, {
                data: "Postcode"
            }, {
                data: "Faxcode"
            }, {
                data: "Linkman"
            }, {
                data: "Mobile"
            }, {
                data: "Email"
            }, {
                data: "Qq"
            }, {
                data: "Wechat"
            }, {
                data: "Weibo"
            }, {
                data: "VirtualIntegral"
            }, {
                data: "RealIntegral"
            }, {
                data: "FeeType"
            }, {
                data: "PrepayValve"
            }, {
                data: "Balance"
            }, {
                data: "FrozenBalance"
            }, {
                data: "IncomingBalance"
            }, {
                data: "Commission"
            }, {
                data: "Discount"
            }, {
                data: "Province"
            }, {
                data: "Area"
            }, {
                data: "County"
            }, {
                data: "Community"
            }, {
                data: "Address"
            }, {
                data: "Status"
            }, {
                data: "Type"
            }, {
                data: "Grade"
            }, {
                data: "Vocation"
            }, {
                data: "Reserve"
            }, {
                data: "Remark"
            }, {
                data: "Creator"
            }, {
                data: "CreateTime"
            }, {
                data: "Auditor"
            }, {
                data: "AuditTime"
            }, {
                data: "Canceler"
            }, {
                data: "CancelTime"
            }, {
                data: "LogoIcon"
            }, {
                data: "LogoUrl"
            }, ],
            select: true,
            buttons: [{
                text: '新增',
                action: function(e, dt, node, config) {
                    usp.admin.corp.list.add()
                }
            }, {
                text: '修改',
                action: function(e, dt, node, config) {
                    usp.admin.corp.list.edit()
                }
            }, {
                text: '删除',
                action: function(e, dt, node, config) {
                    usp.admin.corp.list.remove()
                }
            }, 'copy', /*'csv', */ 'excel', 'pdf', 'print']

        });

    }
})(this)