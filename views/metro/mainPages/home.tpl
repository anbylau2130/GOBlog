<style>
html, body {
    height: 100%;
}
body {
}
.page-content {
    padding-top: 3.125rem;
    min-height: 100%;
    height: 100%;
}

</style>

<div class="app-bar fixed-top darcula" data-role="appbar">
        <a class="app-bar-element branding">{{.CorpName}}</a>
        <span class="app-bar-divider"></span>
        <ul class="app-bar-menu">
                {{range  .menuHorizontal}}
                    <li>
                       <a href=""  class="dropdown-toggle" >{{.Name}}</a>
                        <ul class="d-menu" data-role="dropdown">
                             {{range .ChildrenNode}}
                                <li><a href="javascript:getMenuVertical('{{.ID}}')">{{.Name}}</a></li>
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
                <h2 class="text-light">Quick settings</h2>
                <ul class="unstyled-list fg-dark">
                    <li><a href="javascript:openIframe('admin/User/Profile')" class="fg-white1 fg-hover-yellow">Profile</a></li>
                    <li><a href="" class="fg-white2 fg-hover-yellow">Security</a></li>
                    <li><a id="exit" href="javascript:Logout()" class="fg-white3 fg-hover-yellow">Exit</a></li>
                </ul>
            </div>
        </div>
    </div>

    <div class="page-content">
        <div class="flex-grid no-responsive-future" style="height: 100%;">
            <div class="row" style="height: 100%">
                <div class="cell size-x200" id="cell-sidebar" style="background-color: #71b1d1; height: 100%">
                    <ul id="menusVertical" class="sidebar" >
                      
                    </ul>
                </div>
                <div class="cell auto-size padding20 bg-white" id="cell-content">
                    <iframe id="iframe" scrolling="auto" frameborder="0"  src="" height="100%" width="100%" onload="javascript:usp.resizeIframe(this);"></iframe>
                   
                </div>
            </div>
        </div>
    </div>

    <script>
        $(function(){
          
        })

         function Logout(){
                 $.ajax({
                    type: "POST",
                    url: "/Logout",
                    success: function(rsp){
                       if(rsp.success){
                          window.location.href=rsp.data
                       }
                    }
                })
            }
          
        function openIframe(url){
            $("#iframe").attr("src",url)
            $("#iframe").attr("height",$("#cell-content").height())
          
        }
        function getMenuVertical(id){
            $.ajax({
                type: "POST",
                url: "main/GetMenusVertical",
                data:{ ID:id },
                success: function(rsp){
                    if(rsp.success){
                      var htmltag=''
                      for(i=0;i<rsp.data.length;i++)
                        htmltag+='<li><a href=javascript:openIframe(\''+rsp.data[i].URL+'\')> '+
                            '<span class=\''+rsp.data[i].Icon+'\'></span>'+
                            '<span class=\'title\'>'+rsp.data[i].Name+'</span>'+
                            '<span class=\'counter\'>Message:0</span>'+
                            '</a></li>'
                    }
                     $("#menusVertical").html(htmltag)
                }
            });
        }
    </script>