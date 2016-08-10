<div class="grid padding10">
    <div class="row cells2">
        <div class="cell">
            <div class="bordered">
                <div class="frame">
                    <div data-role="video" style="height:500px">
                        <video>
                            <source src="http://vjs.zencdn.net/v/oceans.mp4" type='video/mp4' />
                        </video>
                    </div>
                </div>
            </div>
        </div>
        <div class="cell">
            <div class="window" style="height:500px">
                <div class="window-caption bg-brown fg-white">
                    <span class="window-caption-icon"></span>
                    <span class="window-caption-title">历史记录</span>
                </div>
                <div class="window-content">
                    <input id="uuid" type="hidden" value="{{.ChatRoomUUID}}" />
                    <input type="hidden" id="uname" value="{{.UserName}}" />

                    <div style="height:350px;overflow:auto">
                        <ul id="chatbox">
                            <li>欢迎光临：</li>
                        </ul>
                    </div>
                    <hr />
                    <div style="height:150px">
                        <textarea id="sendbox" style="width:100%;height:50px"></textarea>
                        <button class="button place-right" id="sendbtn">发送</button>
                    </div>
                </div>
            </div>

        </div>

        <script src="/static/js/usp.admin.chatroom.js"></script>
        <script>
            $(document).ready(function() {
                usp.init()
                usp.admin.chatroom.init()
            })
        </script>