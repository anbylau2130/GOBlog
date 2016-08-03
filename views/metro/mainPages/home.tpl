<style>
    html,
    body {
        height: 100%;
    }
    
    .page-content {
        padding-top: 3.125rem;
        min-height: 100%;
        height: 100%;
    }
</style>

<div class="app-bar fixed-top darcula" data-role="appbar">
    <a onclick="javascript:$('#menusVertical').toggleClass('compact');" class="app-bar-element branding">{{.CorpName}}</a>
    <span class="app-bar-divider"></span>
    <ul class="app-bar-menu">
        {{range .menuHorizontal}}
        <li>
            <a href="" class="dropdown-toggle">{{.Name}}</a>
            <ul class="d-menu" data-role="dropdown">
                {{range .ChildrenNode}}
                <li><a href="javascript:usp.home.getMenuVertical('{{.ID}}')">{{.Name}}</a></li>
                <li class="divider"></li>
                {{end}}
            </ul>
        </li>
        <li class="divider"></li>
        {{end}}
    </ul>

    <div class="app-bar-element place-right">
        <span class="dropdown-toggle"><span class="mif-cog"></span> {{.UserName}}</span>
        <div class="app-bar-drop-container padding10 place-right no-margin-top block-shadow fg-dark" data-role="dropdown" data-no-close="true" style="width: 220px">
            <h2 class="text-light">设 置</h2>
            <ul class="unstyled-list fg-dark">
                <li><a href="javascript:usp.home.openIframe('admin/User/Profile')" class="fg-white1 fg-hover-yellow">个人信息</a></li>
                <li><a href="" class="fg-white2 fg-hover-yellow">账号安全</a></li>
                <li><a id="exit" href="javascript:usp.home.openLogoutWindow()" class="fg-white3 fg-hover-yellow">账号注销</a></li>
                <li>
                    <span id="timerYear">{{.Now.Year}}</span>年
                    <span id="timerMonth">{{.Now.Month}}</span>月
                    <span id="timerDay">{{.Now.Day}}</span>日&nbsp;
                    <br/>
                    <span id="timerHour">{{.Now.Hour}}</span>:
                    <span id="timerMinute">{{.Now.Minute}}</span>:
                    <span id="timerSecond">{{.Now.Second}}</span>
                </li>
            </ul>
        </div>
    </div>
</div>

<div class="page-content">
    <div class="flex-grid no-responsive-future" style="height: 100%;">
        <div class="row" style="height: 100%">
            <div class="cell">
                <ul id="menusVertical" class="sidebar compact no-responsive-future" style="height:100%">
                </ul>
            </div>
            <div class="cell auto-size padding10 bg-white" id="cell-content">
                <iframe id="iframe" scrolling="auto" frameborder="0" src="" height="100%" width="100%" onload="javascript:usp.resizeIframe(this);"></iframe>
            </div>
        </div>
    </div>
</div>
<script src="/static/js/usp.home.js"></script>
<script>
    $(function() {
        if (window.top != window) {
            window.top.document.location.href = window.location.href;
        }
        usp.home.init($('#tabContainer'), $('#timerYear'), $('#timerMonth'), $('#timerDay'), $('#timerHour'), $('#timerMinute'), $('#timerSecond'), 'main/CheckSSO');
    })
</script>