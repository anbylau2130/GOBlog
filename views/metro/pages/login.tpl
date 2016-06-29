
    <style>
        .login-form {
            width: 25rem;
            height: 18.75rem;
            position: fixed;
            top: 50%;
            margin-top: -9.375rem;
            left: 50%;
            margin-left: -12.5rem;
            background-color: #ffffff;
            opacity: 0;
            -webkit-transform: scale(.8);
            transform: scale(.8);
        }
    </style>

    <script>

        /*
        * Do not use this is a google analytics fro Metro UI CSS
        * */
        if (window.location.hostname !== 'localhost') {

            (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
                (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
                    m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
            })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

            ga('create', 'UA-58849249-3', 'auto');
            ga('send', 'pageview');

        }
        //这个就是键盘触发的函数
    var SubmitOrHidden = function(evt){
        evt = window.event || evt;
        if(evt.keyCode==13){//如果取到的键值是回车
            formSubmit();       
        }
                    
    }
    
    function formSubmit(){
        $.ajax({
            type: 'post',
            url: 'Login?isajax=1',
            data: $("form").serialize(),
            success: function(data) {
                // var data = $.parseJSON(data);
                if(data.status){
                    location.href = "/"
                }else{
                    $.Notify({
                        caption: "登录失败",
                        content: data.info,
                        type:"alert"
                    });
                }
            }
        });
        return false;
    }

        $(function(){
            window.document.onkeydown=SubmitOrHidden;//当有键按下时执行函数
            var form = $(".login-form");
            form.css({
                opacity: 1,
                "-webkit-transform": "scale(1)",
                "transform": "scale(1)",
                "-webkit-transition": ".5s",
                "transition": ".5s"
            });
        });
      
    </script>

     <div class="login-form padding20 block-shadow">
        <form id="loginForm" data-role="validator"  data-on-submit="formSubmit" data-hint-mode="hint">
            <h1 class="text-light">登录</h1>
            <hr class="thin"/>
            <br />
            <div class="input-control text full-size" data-role="input">
                <label for="username">用户名:</label>
                <input type="text"  name="username" id="username" 
                 data-validate-func="required,minlength"
                 data-validate-arg=",5"
                data-validate-hint="不能为空 且 最少5个字符"
                value="admin"
                />
                <button class="button helper-button clear"><span class="mif-cross"></span></button>
            </div>
            <br />
            <br />
            <div class="input-control password full-size" data-role="input">
               
                <label for="userpassword"> 密&nbsp;&nbsp;&nbsp;码 :</label>
                <input type="password" 
                data-validate-func="required,minlength"
                data-validate-arg=",6"
                data-validate-hint="不能为空 且 最少6个字符"
                name="password" id="password" 
                value="123456"
                />
                 <span class="input-state-error mif-warning"></span>
                 <span class="input-state-success mif-checkmark"></span>
                <button class="button helper-button reveal"><span class="mif-looks"></span></button>
            </div>
            <br />
            <br />
            <div class="form-actions">
                <button  class="button primary success">Login&nbsp;.&nbsp;.&nbsp;.</button>
                <button type="button" class="button link">Cancel</button>
            </div>
        </form>
       

    </div>