<h2>修改角色 <button id="return" data-bind="click: usp.admin.role.edit.goBack" class="button place-right primary"><span class="mif-keyboard-return "></span></button></h2>
<hr class="thin bg-grayLighter">
<form data-role="validator" data-hint-mode="hint" data-on-submit="return usp.admin.role.edit.save()">
    <!--<div class="grid padding20">
        <div class="row cells2">
            <div class="cell">
                <label>ID</label>
                <div class="input-control text full-size">
                    <input type="hidden" data-bind="value:ID" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Corp</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Corp" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Name</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Name" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Type</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Type" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Reserve</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Reserve" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Remark</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Remark" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Creator</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Creator" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>CreateTime</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:CreateTime" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Auditor</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Auditor" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>AuditTime</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:AuditTime" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Canceler</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Canceler" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>CancelTime</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:CancelTime" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <button class="button place-right primary">保存</button>
    </div>-->
    <div class="panel">
        <div class="heading">
            <span class="title">基本信息</span>
        </div>
        <div class="content padding10">
            <div class="grid">
                <div class="row ">
                    <div class="cell">
                        <label> 角色名称</label>
                        <div class="input-control text full-size">
                            <input type="hidden" data-bind="value:Menus" />
                            <input type="hidden" data-bind="value:Privileges" />
                            <input type="hidden" data-bind="value:ID" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                            <input type="text" data-bind="value:Name" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                            <span class="input-state-error mif-warning"></span>
                            <span class="input-state-success mif-checkmark"></span>
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="cell">
                        <label>角色说明</label>
                        <div class="input-control text full-size">
                            <input type="text" data-bind="value:Remark" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                            <span class="input-state-error mif-warning"></span>
                            <span class="input-state-success mif-checkmark"></span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="panel">
        <div class="heading">
            <span class="title">权限分配</span>
        </div>
        <div class="content  padding10">
            <div id="container"></div>
        </div>
    </div>
    <button data-bind='click:usp.admin.role.edit.save' class="button place-right primary">保存</button>
</form>
</div>
<script src="/static/js/usp.admin.role.edit.js"></script>
<script>
    $(function() {
        var id = '{{.Model.ID}}'
        usp.init()
        usp.admin.role.edit.init(id)
        usp.bindJsTree("#container", "/Service/RoleTreeInit?Role=" + id, true, function() {})

    })
</script>