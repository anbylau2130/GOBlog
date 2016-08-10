<h2>userAdd <button id="btnBack" data-bind="click: usp.admin.user.add.goBack" class="button place-right primary"><span class="mif-keyboard-return "></span></button></h2>
<hr class="thin bg-grayLighter">
<form data-role="validator" data-hint-mode="hint" data-on-submit="return usp.admin.user.add.save()">
    <div class="grid no-padding-top padding100">
        <div class="row">
            <div class="cell">
                <label>LoginName</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:LoginName" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="cell">
                <label>RealName</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:RealName" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="cell">
                <label>Password</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Password" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="cell">
                <label>PasswordConfirm</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:PasswordConfirm" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="cell">
                <label>Role</label>
                <div class="input-control text full-size required">
                    <select id="role" multiple="multiple" data-validate-func="selected" data-validate-hint-position="top" data-validate-hint="请选择角色!"></select>
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
    </div>

    <button class="button place-right primary">保存</button>
</form>
<script src="/static/js/usp.admin.user.add.js"></script>
<script>
    $(function() {
        usp.init()
        usp.admin.user.add.init()

    })
</script>