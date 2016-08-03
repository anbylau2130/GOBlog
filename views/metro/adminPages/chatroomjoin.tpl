<!--<div class="container">
    <input id="uuid" type="hidden" value="{{.ChatRoomUUID}}" />
    <h3>用户名: <span id="uname">{{.UserName}}</span></h3>
    <h4>sendbox:</h4>
    <form class="form-inline">
        <div class="col-md-6 form-group">
            <input id="sendbox" type="text" class="form-control" onkeydown="if(event.keyCode==13)return false;" required>
        </div>
        <button id="sendbtn" type="button" class="btn btn-default">发送</button>
    </form>
</div>

<div class="container">
    <h3>历史记录:</h3>
    <ul id="chatbox">
        <li>欢迎光临：</li>
    </ul>
</div>-->
<div class="padding10 full-size" style="height:100%">
    <div class="bg-lime" style="height:100%;width:48%;float:left">
        <div data-role="video" style="height:100%">
            <video>
                <source src="http://vjs.zencdn.net/v/oceans.mp4" type='video/mp4' />
            </video>
        </div>
    </div>
    <div class="bg-grayLighter" style="height:100%;width:48%:float:right">
        <input id="uuid" type="hidden" value="{{.ChatRoomUUID}}" />
        <div class="bg-taupe" style="height:5%">

        </div>
        <div class="bg-grayLighter" style="height:70%;">
            <h3>历史记录:</h3>
            <ul id="chatbox">
                <li>欢迎光临：</li>
            </ul>
        </div>
        <div class="bg-gray " style="height:5%">
        </div>
        <div class="bg-grayLighter" style="height:20%;width:100%;">
            <div class="input-control textarea full-size" data-role="input" data-text-auto-resize="true">
                <textarea id="sendbox"></textarea>
            </div>
        </div>
    </div>
</div>

<script src="/static/js/usp.admin.chatroom.js"></script>
<script>
    $(document).ready(function() {
        usp.init()
        $("#video").height($("body").height() * 0.4)
        $("#chatroom").height($("body").height() * 0.4)
        $("#log").height($("body").height() * 0.2)
            //usp.admin.chatroom.init()
    })
</script>