<h2>新增角色 <button id="btnBack" data-bind="click: usp.admin.role.add.goBack" class="button place-right primary"><span class="mif-keyboard-return "></span></button></h2>
<hr class="thin bg-grayLighter">
<form data-role="validator" data-hint-mode="hint" data-on-submit="return usp.admin.role.add.save()">
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

    <button data-bind='click:usp.admin.role.add.save' class="button place-right primary">保存</button>
</form>
<script src="/static/js/usp.admin.role.add.js"></script>
<script>
    $(function() {
        usp.init()
        usp.admin.role.add.init()
        usp.bindJsTree("#container", "/Service/RoleTree", true, function() {})
    })
</script>