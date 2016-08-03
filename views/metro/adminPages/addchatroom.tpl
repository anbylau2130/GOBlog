<h2>chatroomAdd <button id="btnBack" data-bind="click: usp.admin.chatroom.add.goBack" class="button place-right primary"><span class="mif-keyboard-return "></span></button></h2>
<hr class="thin bg-grayLighter">
<form data-role="validator" data-hint-mode="hint" data-on-submit="return usp.admin.chatroom.add.save()">
    <div class="grid">

        <div class="row cells">
            <div class="cell">
                <label>Name</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Name" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells">
            <div class="cell">
                <label>Title</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Title" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells">
            <div class="cell">
                <label>Remark</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Remark" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells">
            <div class="cell">
                <label>PicURL</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:PicURL" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
    </div>
    </div>
    <button data-bind='click:usp.admin.chatroom.add.save' class="button place-right primary">保存</button>
</form>
<script src="/static/js/usp.admin.chatroom.add.js"></script>
<script>
    $(function() {
        usp.init()
        usp.admin.chatroom.add.init()
    })
</script>