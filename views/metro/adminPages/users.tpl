<h2>SysOperatorList<span class="mif-drive-eta place-right"></span></h2>
<hr class="thin bg-grayLighter">
<div class="toolbar">
    <div class="toolbar-section">
        <!--<button class='toolbar-button' data-bind="click: usp.admin.menus.list.add">新增</button>
        <button class='toolbar-button' data-bind="click: usp.admin.user.list.edit">修改</button>
        <button class='toolbar-button' data-bind="click: usp.admin.user.list.remove">删除</button>-->
    </div>
</div>
<hr class="thin bg-grayLighter">
<table id="List" class="display table striped hovered cell-hovered border bordered nowrap" cellspacing="0" width="100%">
</table>
<script src="/static/js/usp.admin.user.list.js"></script>
<script>
    $(document).ready(function() {
        usp.init()
        usp.admin.user.list.init("#List", "GetList")
    })
</script>