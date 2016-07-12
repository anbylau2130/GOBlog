(function() {
    usp.namespace("usp.home");
    usp.home.TimerPID = 0;
    usp.home.CheckSsoPID = 0;
    usp.home.Timer = function(timeSpan) {
        var year, month, day, seconds, minutes, hours;
        year = parseInt(usp.home.yearContainer.text(), 10);
        month = parseInt(usp.home.monthContainer.text(), 10);
        day = parseInt(usp.home.dayContainer.text(), 10);
        hours = parseInt(usp.home.hourContainer.text(), 10);
        minutes = parseInt(usp.home.minuteContainer.text(), 10);
        seconds = parseInt(usp.home.secondContainer.text(), 10);
        var dateTime = new Date(year, month, day, hours, minutes, seconds, timeSpan);
        if (dateTime.getSeconds() < 10)
            seconds = "0" + dateTime.getSeconds();
        else
            seconds = dateTime.getSeconds();
        if (dateTime.getMinutes() < 10)
            minutes = "0" + dateTime.getMinutes();
        else
            minutes = dateTime.getMinutes();
        if (dateTime.getHours() < 10)
            hours = "0" + dateTime.getHours();
        else
            hours = dateTime.getHours();
        usp.home.yearContainer.text(year);
        usp.home.monthContainer.text(month);
        usp.home.dayContainer.text(day);
        usp.home.secondContainer.text(seconds);
        usp.home.minuteContainer.text(minutes);
        usp.home.hourContainer.text(hours);
    };
    usp.home.CheckSSO = function(checkSsoUrl) {
        $.post(checkSsoUrl, null, function(json) {
            if (json.flag == false) {
                clearInterval(usp.home.CheckSsoPID);
                usp.home.loginFlag = false;
                clearInterval(usp.home.checkSSO_timer);
                $.messager.confirm('"警告"', json.message, function(r) {
                    window.location.href = '/';
                });
                window.location.href = '/';
            } else {
                var year, month, day, seconds, minutes, hours;
                var dateTime = json.dateTime.replace(new RegExp(/\//g), "");
                eval('var dateTime = new ' + dateTime);
                year = dateTime.getFullYear();
                month = dateTime.getMonth() + 1;
                day = dateTime.getDate();
                if (dateTime.getSeconds() < 10)
                    seconds = "0" + dateTime.getSeconds();
                else
                    seconds = dateTime.getSeconds();
                if (dateTime.getMinutes() < 10)
                    minutes = "0" + dateTime.getMinutes();
                else
                    minutes = dateTime.getMinutes();
                if (dateTime.getHours() < 10)
                    hours = "0" + dateTime.getHours();
                else
                    hours = dateTime.getHours();
                usp.home.yearContainer.text(year);
                usp.home.monthContainer.text(month);
                usp.home.dayContainer.text(day);
                usp.home.secondContainer.text(seconds);
                usp.home.minuteContainer.text(minutes);
                usp.home.hourContainer.text(hours);
            }
        });
    }

    usp.home.logout = function(url) {
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
    
    usp.home.showDialog=function (id){
        var dialog = $(id).data('dialog');
        if (!dialog.element.data('opened')) {
            dialog.open();
        } else {
            dialog.close();
        }
    }
    usp.home.openIframe=function (url){
            $("#iframe").attr("src",url)
            $("#iframe").attr("height",$("#cell-content").height())
          
        }
    usp.home.getMenuVertical=function (id){
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
    usp.home.init = function(tabContainer,yearContainer, monthContainer, dayContainer, hourContainer, minuteContainer, secondContainer, checkSsoUrl) {
        this.yearContainer = yearContainer;
        this.monthContainer = monthContainer;
        this.dayContainer = dayContainer;
        this.hourContainer = hourContainer;
        this.minuteContainer = minuteContainer;
        this.secondContainer = secondContainer;
        this.checkSsoUrl = checkSsoUrl;
        var interval = 1000;
        usp.home.TimerPID = setInterval('usp.home.Timer(' + interval + ')', interval);
        usp.home.CheckSsoPID = setInterval('usp.home.CheckSSO("' + checkSsoUrl + '")', 60 * interval);
        // usp.tabContainer = tabContainer;
        // console.log(tabContainer)
        // usp.tabContainer.tabs({
        //     width: tabContainer.parent().width(),
        //     height: "auto"
        // });
    };
})(this);