<h2>新增权限 <button id="btnBack" data-bind="click: usp.admin.privilege.add.goBack" class="button place-right primary"><span class="mif-keyboard-return "></span></button></h2>
<hr class="thin bg-grayLighter">
<form data-role="validator" data-hint-mode="hint" data-on-submit="return usp.admin.privilege.add.save()">
    <div class="grid">
        <div class="row cells2">
            <div class="cell">
                <label>ID</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:ID" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Parent</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Parent" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Menu</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Menu" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Name</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Name" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Clazz</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Clazz" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Area</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Area" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Controller</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Controller" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Method</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Method" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Parameter</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Parameter" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Url</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Url" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
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
    </div>
    <button data-bind='click:usp.admin.privilege.add.save' class="button place-right primary">保存</button>
</form>
<script src="/static/js/usp.admin.privilege.add.js"></script>
<script>
    $(function() {
        usp.init()
        usp.admin.privilege.add.init()
    })
</script>