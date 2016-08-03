<h2>聊天室列表 <button data-bind="click:  usp.admin.chatrooms.add" class="button place-right primary">创建</button></h2>
<hr class="thin bg-grayLighter">
<div class="grid">
    <div class="row cells6">
        {{range $index, $elem := .Rooms}}
        <div class="cell" data-bind="click: usp.admin.chatrooms.enter">
            <div class="image-container image-format-square" style="width: 100%;height:auto">
                <div class="frame">
                    <input type="hidden" id="uuid" value="{{$elem.UUID}}" />
                    <img src="{{$elem.PicURL}}" />
                </div>
                <div class="image-overlay">
                    <h1>{{$elem.Name}}</h1>
                    <p></p>
                </div>
            </div>
        </div>
        {{end}}
    </div>

</div>
<script src="/static/js/usp.admin.chatrooms.js"></script>
<script>
    $(function() {
        usp.init()
        usp.admin.chatrooms.init()
    })
</script>