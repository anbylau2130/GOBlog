<h2>新增公司 <button id="btnBack" data-bind="click: usp.admin.corp.add.goBack" class="button place-right primary"><span class="mif-keyboard-return "></span></button></h2>
<hr class="thin bg-grayLighter">
<form data-role="validator" data-hint-mode="hint" data-on-submit="return usp.admin.corp.add.save()">
    <div class="grid no-padding-top padding100">
        <div class="row ">
            <div class="cell">
                <label>公司名称</label>
                <div class="input-control text full-size" data-role="input">
                    <input type="text" name="Name" id="Name" data-bind="value:Name" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="cell">
                <label>公司类型</label>
                <div class="input-control full-size" data-role="select" data-placeholder="Select a state" data-allow-clear="true">
                    <select id="type"></select>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="cell ">
                <label>管理员账号</label>
                <div class="input-control text full-size " data-role="input ">
                    <input type="text " data-bind="value:adminLoginName " data-validate-func="required " data-validate-hint-position="top " data-validate-hint="不能为空! " />
                    <span class="input-state-error mif-warning "></span>
                    <span class="input-state-success mif-checkmark "></span>
                </div>
            </div>
        </div>
        <div class="row ">
            <div class="cell ">
                <label>管理员密码</label>
                <div class="input-control text full-size " data-role="input ">
                    <input type="password" data-bind="value:adminPassWord " data-validate-func="required " data-validate-hint-position="top " data-validate-hint="不能为空! " />
                    <span class="input-state-error mif-warning "></span>
                    <span class="input-state-success mif-checkmark "></span>
                </div>
            </div>
        </div>
    </div>
    <button class="button place-right primary ">保存</button>
</form>
<script src="/static/js/usp.admin.corp.add.js "></script>
<script>
    $(function() {
        usp.init()
        usp.admin.corp.add.init()
        usp.bindSelect("#type", "/Service/CorpTypeSelect")

    })
</script>