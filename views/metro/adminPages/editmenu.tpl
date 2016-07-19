<h2>编辑菜单 <button id="return"  data-bind="click: usp.admin.menus.edit.goBack" class="button place-right primary" ><span class="mif-keyboard-return "></span></button></h2>
<hr class="thin bg-grayLighter">
<div class="grid">
    <div class="row cells3">
        <div class="cell">
            <label>ID  </label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:ID" />
            </div>
        </div>
        <div class="cell">
            <label>Direction  </label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:Direction" />
            </div>
        </div>
        <div class="cell">
            <label>Parent     </label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:Parent" />
            </div>
        </div>
    </div>
    <div class="row cells3">

        <div class="cell">
            <label>Name       </label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:Name" />
            </div>
        </div>
        <div class="cell">
            <label>Icon       </label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:Icon" />
            </div>
        </div>
        <div class="cell">
            <label>Clazz</label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:Clazz" />
            </div>
        </div>
    </div>
    <div class="row cells3">

        <div class="cell">
            <label>Area</label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:Area" />
            </div>
        </div>
        <div class="cell">
            <label>Controller </label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:Controller" />
            </div>
        </div>
        <div class="cell">
            <label>Method</label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:Method" />
            </div>
        </div>
    </div>
    <div class="row cells3">
        <div class="cell">
            <label>Parameter</label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:Parameter" />
            </div>
        </div>
        <div class="cell">
            <label>URL</label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:URL" />
            </div>
        </div>
        <div class="cell">
            <label>Reserve</label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:Reserve" />
            </div>
        </div>
    </div>
    <div class="row cells3">
        <div class="cell">
            <label>Remark</label>
            <div class="input-control text full-size">
                <input type="text" data-bind="value:Remark" />
            </div>
        </div>
        <div class="cell"></div>
        <div class="cell"></div>
    </div>
</div>
<button data-bind='click:usp.admin.menus.edit.save' class="button place-right primary">保存</button>
</div>
<script src="/static/js/usp.admin.menus.edit.js"></script>
<script>
    $(function() {
        var id = {{.Model.ID}}
        usp.init()
        usp.admin.menus.edit.init(id)
    })
</script>