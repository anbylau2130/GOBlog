(function() {
    usp.namespace("usp.admin.user.list");

    usp.admin.user.list.add = function name(params) {
        window.location.href = "Add";
    }
    usp.admin.user.list.edit = function name(params) {
        if (usp.admin.user.list.table.row('.selected').data())
            window.location.href = "Edit?id=" + usp.admin.user.list.table.row('.selected').data().ID;
    }

    usp.admin.user.list.remove = function name(params) {
        if (usp.admin.user.list.table.row('.selected').data()) {
            $.ajax({
                url: "Del",
                data: {
                    id: usp.admin.user.list.table.row('.selected').data().ID
                },
                async: false,
                success: function(response) {
                    if (response.status) {
                        usp.admin.user.list.table.row('.selected').remove().draw(false)
                        usp.Notify('系统信息', response.info, 'success')
                    } else {
                        usp.Notify('系统信息', response.info, 'alert')
                    }
                }
            });
        }
    }
    usp.admin.user.list.table = null;
    usp.admin.user.list.init = function(id, url) {
        ko.applyBindings();
        usp.admin.user.list.table = $(id).DataTable({
            dom: 'Bfrtip',
            iDisplayLength: 10,
            "bServerSide": true,
            "bPaginate": true, //是否允许分页
            "bLengthChange": true, //是否显示每页显示条数
            "bFilter": true, //是否启用条件查询
            "bSort": true, //是否启用列排序
            "bInfo": true, //是否显示分页信息 
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
                title: "Corp",
                name: "Corp",
                targets: 1,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "LoginName",
                name: "LoginName",
                targets: 2,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "RealName",
                name: "RealName",
                targets: 3,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Password",
                name: "Password",
                targets: 4,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Mobile",
                name: "Mobile",
                targets: 5,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "IdCard",
                name: "IdCard",
                targets: 6,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Email",
                name: "Email",
                targets: 7,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "WechatOpenid",
                name: "WechatOpenid",
                targets: 8,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "AlipayOpenid",
                name: "AlipayOpenid",
                targets: 9,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Weibo",
                name: "Weibo",
                targets: 10,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "AvailableIP",
                name: "AvailableIP",
                targets: 11,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "WeatherCode",
                name: "WeatherCode",
                targets: 12,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "VirtualIntegral",
                name: "VirtualIntegral",
                targets: 13,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "RealIntegral",
                name: "RealIntegral",
                targets: 14,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Balance",
                name: "Balance",
                targets: 15,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "FrozenBalance",
                name: "FrozenBalance",
                targets: 16,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "IncomingBalance",
                name: "IncomingBalance",
                targets: 17,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Commission",
                name: "Commission",
                targets: 18,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Discount",
                name: "Discount",
                targets: 19,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Province",
                name: "Province",
                targets: 20,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Area",
                name: "Area",
                targets: 21,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "County",
                name: "County",
                targets: 22,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Community",
                name: "Community",
                targets: 23,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Address",
                name: "Address",
                targets: 24,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Status",
                name: "Status",
                targets: 25,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Skin",
                name: "Skin",
                targets: 26,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Grade",
                name: "Grade",
                targets: 27,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Star",
                name: "Star",
                targets: 28,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Session",
                name: "Session",
                targets: 29,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "LoginTime",
                name: "LoginTime",
                targets: 30,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "LoginIP",
                name: "LoginIP",
                targets: 31,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "LoginAgent",
                name: "LoginAgent",
                targets: 32,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "LoginCount",
                name: "LoginCount",
                targets: 33,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "LoginErrorCount",
                name: "LoginErrorCount",
                targets: 34,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "FrozenStartTime",
                name: "FrozenStartTime",
                targets: 35,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "FrozenEndTime",
                name: "FrozenEndTime",
                targets: 36,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Reserve",
                name: "Reserve",
                targets: 37,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Remark",
                name: "Remark",
                targets: 38,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Creator",
                name: "Creator",
                targets: 39,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "CreateTime",
                name: "CreateTime",
                targets: 40,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Auditor",
                name: "Auditor",
                targets: 41,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "AuditTime",
                name: "AuditTime",
                targets: 42,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Canceler",
                name: "Canceler",
                targets: 43,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "CancelTime",
                name: "CancelTime",
                targets: 44,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, ],
            columns: [{
                data: "ID"
            }, {
                data: "Corp"
            }, {
                data: "LoginName"
            }, {
                data: "RealName"
            }, {
                data: "Password"
            }, {
                data: "Mobile"
            }, {
                data: "IdCard"
            }, {
                data: "Email"
            }, {
                data: "WechatOpenid"
            }, {
                data: "AlipayOpenid"
            }, {
                data: "Weibo"
            }, {
                data: "AvailableIP"
            }, {
                data: "WeatherCode"
            }, {
                data: "VirtualIntegral"
            }, {
                data: "RealIntegral"
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
                data: "Skin"
            }, {
                data: "Grade"
            }, {
                data: "Star"
            }, {
                data: "Session"
            }, {
                data: "LoginTime"
            }, {
                data: "LoginIP"
            }, {
                data: "LoginAgent"
            }, {
                data: "LoginCount"
            }, {
                data: "LoginErrorCount"
            }, {
                data: "FrozenStartTime"
            }, {
                data: "FrozenEndTime"
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
            }, ],
            select: true,
            buttons: [{
                text: '新增',
                action: function(e, dt, node, config) {
                    usp.admin.user.list.add()
                }
            }, {
                text: '修改',
                action: function(e, dt, node, config) {
                    usp.admin.user.list.edit()
                }
            }, {
                text: '删除',
                action: function(e, dt, node, config) {
                    usp.admin.user.list.remove()
                }
            }, 'copy', /*'csv', */ 'excel', 'pdf', 'print']

        });
    }
})(this)