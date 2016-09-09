(function() {
    usp.namespace("usp.admin.privilege.list");

    usp.admin.privilege.list.add = function name(params) {
        window.location.href = "Add";
    }
    usp.admin.privilege.list.edit = function name(params) {
        if (usp.admin.privilege.list.table.row('.selected').data())
            window.location.href = "Edit?id=" + usp.admin.privilege.list.table.row('.selected').data().ID;
    }

    usp.admin.privilege.list.remove = function name(params) {
        if (usp.admin.privilege.list.table.row('.selected').data()) {
            $.ajax({
                url: "Del",
                data: {
                    id: usp.admin.privilege.list.table.row('.selected').data().ID
                },
                async: false,
                success: function(response) {
                    if (response.status) {
                        usp.admin.privilege.list.table.row('.selected').remove().draw(false)
                        usp.Notify('系统信息', response.info, 'success')
                    } else {
                        usp.Notify('系统信息', response.info, 'alert')
                    }
                }
            });
        }
    }
    usp.admin.privilege.list.table = null;
    usp.admin.privilege.list.init = function(id, url) {
        ko.applyBindings();
        usp.admin.privilege.list.table = $(id).DataTable({
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
                title: "Parent",
                name: "Parent",
                targets: 1,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Menu",
                name: "Menu",
                targets: 2,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Name",
                name: "Name",
                targets: 3,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Clazz",
                name: "Clazz",
                targets: 4,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Area",
                name: "Area",
                targets: 5,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Controller",
                name: "Controller",
                targets: 6,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Method",
                name: "Method",
                targets: 7,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Parameter",
                name: "Parameter",
                targets: 8,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Url",
                name: "Url",
                targets: 9,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Reserve",
                name: "Reserve",
                targets: 10,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Remark",
                name: "Remark",
                targets: 11,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Creator",
                name: "Creator",
                targets: 12,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "CreateTime",
                name: "CreateTime",
                targets: 13,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Auditor",
                name: "Auditor",
                targets: 14,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "AuditTime",
                name: "AuditTime",
                targets: 15,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "Canceler",
                name: "Canceler",
                targets: 16,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, {
                title: "CancelTime",
                name: "CancelTime",
                targets: 17,
                orderable: true,
                //"render": function(data, type, row) {
                //    return '<span class="' + data + '"></span>'
                //}
            }, ],
            columns: [{
                data: "ID"
            }, {
                data: "Parent"
            }, {
                data: "Menu"
            }, {
                data: "Name"
            }, {
                data: "Clazz"
            }, {
                data: "Area"
            }, {
                data: "Controller"
            }, {
                data: "Method"
            }, {
                data: "Parameter"
            }, {
                data: "Url"
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
                    usp.admin.privilege.list.add()
                }
            }, {
                text: '修改',
                action: function(e, dt, node, config) {
                    usp.admin.privilege.list.edit()
                }
            }, {
                text: '删除',
                action: function(e, dt, node, config) {
                    usp.admin.privilege.list.remove()
                }
            }, 'copy', /*'csv', */ 'excel', 'pdf', 'print']

        });
    }
})(this)