
<div class="container">
    <div class="padding20">
        <div class="grid">
            <div class="row cells4">
                 <div class="cell colspan rowspan2">
                     <div class="image-container bordered image-format-square" style="width: 100%;" >
                         <div class="frame" style="height:200px;width:200px">
                         <div style="width: 100%; height: 180px; border-radius: 0px; background-image: url({{.UImage}}); background-size: cover; background-repeat: no-repeat;"></div>
                          </div>
                     </div>
                 </div>
                 <div class="cell colspan3">
                 <table class="table">
                    <thead>
                        <tr>
                            <td><h3>用户信息</h3></td>
                            <td></td>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td class="align-left">登录名：</td>
                            <td>{{.LoginName}}</td>
                        </tr>
                        <tr>
                            <td class="align-left">姓名：</td>
                            <td>{{.RealName}}</td>
                        </tr>
                       
                    </tbody>
                </table>
              </div>
            </div>
        </div>
    </div>
</div>