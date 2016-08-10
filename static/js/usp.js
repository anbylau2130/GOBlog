(function() {
    Date.prototype.format = function(format) //author: meizz 
        {
            var o = {
                'M+': this.getMonth() + 1, //month 
                'd+': this.getDate(), //day 
                'h+': this.getHours(), //hour 
                'm+': this.getMinutes(), //minute 
                's+': this.getSeconds(), //second 
                'q+': Math.floor((this.getMonth() + 3) / 3), //quarter 
                'S': this.getMilliseconds() //millisecond 
            }
            if (/(y+)/.test(format)) format = format.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
            for (var k in o)
                if (new RegExp('(' + k + ')').test(format))
                    format = format.replace(RegExp.$1,
                        RegExp.$1.length == 1 ? o[k] :
                        ('00' + o[k]).substr(('' + o[k]).length));
            return format;
        };
    usp = window.usp || (window.usp = {});
    usp.namespace = function() {
        var a = arguments,
            o = null,
            i, j, d, rt;
        for (i = 0; i < a.length; ++i) {
            d = a[i].split('.');
            rt = d[0];
            eval('if (typeof ' + rt + ' == "undefined"){' + rt + ' = {};} o = ' + rt + ';');
            for (j = 1; j < d.length; ++j) {
                o[d[j]] = o[d[j]] || {};
                o = o[d[j]];
            }
        }
    };
    usp.ParseUTCDate = function(milliseconds) {
        //alert(milliseconds);
        if (milliseconds != null) {
            var datetime = milliseconds.replace(new RegExp(/\//g), '');
            eval('var datetime = new ' + datetime);
            return datetime.format('yyyy年MM月dd日 hh:mm:ss');
        } else {
            return "";
        }
    };
    usp.ParseShortDate = function(milliseconds) {
        //alert(milliseconds);
        if (milliseconds != null) {
            var datetime = milliseconds.replace(new RegExp(/\//g), '');
            eval('var datetime = new ' + datetime);
            return datetime.format('yyyy年MM月dd日');
        } else {
            return "";
        }
    };

    usp.ParseUTCDateToDate = function(milliseconds) {
        //alert(milliseconds);
        if (milliseconds != null) {
            var temp = milliseconds.replace(new RegExp(/\//g), '');
            eval('var temp = new ' + temp);
            return temp;
        } else {
            return '';
        }
    };
    //获取某一月的第一天
    usp.GetMonthFirstDay = function(iYear, iMonth) {
            iMonth = iMonth - 1;
            return (new Date(iYear, iMonth, 1)).format("yyyy-MM-dd");
        }
        //获取某一月的最后一天
    usp.GetMonthLastDay = function(iYear, iMonth) {
        iMonth = iMonth - 1;

        var MonthNextFirstDay = new Date(iYear, parseInt(iMonth) + 1, 1);
        return (new Date(MonthNextFirstDay - 86400000)).format("yyyy-MM-dd");
    }

    //获取某一季度的第一天
    usp.GetQuarterFirstDay = function(iYear, iQuarter) {
            return getMonthFirstDay(iYear, (parseInt(iQuarter) - 1) * 3);
        }
        //获取某一季度的最后一天
    usp.GetQuarterLastDay = function(iYear, iQuarter) {
            return getMonthLastDay(iYear, (parseInt(iQuarter) - 1) * 3 + 2);
        }
        //获取某一年的第一天
    usp.GetYearFirstDay = function(iYear) {
            return (new Date(iYear, 0, 1)).format("yyyy-MM-dd");
        }
        //获取某一年的最后一天
    usp.GetYearLastDay = function(iYear) {
        var YearNextFirstDay = new Date(parseInt(iYear) + 1, 0, 1);
        return (new Date(YearNextFirstDay - 86400000)).format("yyyy-MM-dd");
    }
    usp.resizeIframe = function(obj) {
        //alert(obj.contentWindow.document.body.scrollHeight);
        //obj.style.height = obj.contentWindow.document.body.scrollHeight + 'px';
        // $('#iframe', parent.document).perfectScrollbar();
    }

    usp.addTab = function(tab_container, tab_icon, tab_title, tab_href) {
        var container = tab_container || usp.tabContainer;
        if (container.tabs('exists', tab_title)) {
            container.tabs('select', tab_title);
        } else {
            if (tab_href) {
                var content = '<iframe scrolling="auto"   frameborder="0"  src="' + tab_href + '" style="overflow:hidden;height:100%;width:100%;position:relative" height="100%" width="100%" onload="javascript:usp.resizeIframe(this);"></iframe>';
            } else {
                content = '';
            }
            if (tab_icon) {
                container.tabs('add', {
                    icon: tab_icon,
                    title: tab_title,
                    closable: true,
                    content: content
                        //width: container.parent().width(),
                        //height: container.parent().height()
                });
            } else {
                container.tabs('add', {
                    title: tab_title,
                    closable: true,
                    content: content
                        //width: container.parent().width(),
                        //height: container.parent().height()
                });
            }
            container.tabs('select', tab_title);
        }
    };

    usp.bindSelect = function(selectid, url) {
            var control = $(selectid);
            //设置Select2的处理
            control.select2({
                allowClear: true,
                placeholder: "请选择...",
                // formatResult: formatResult,
                // formatSelection: formatSelection,
                escapeMarkup: function(m) {
                    return m;
                }
            });

            //绑定Ajax的内容
            $.getJSON(url, function(data) {
                control.empty(); //清空下拉框
                $.each(data, function(i, item) {

                    control.append("<option value='" + item.Value + "'>&nbsp;" + item.Text + "</option>");

                });
            });
        }
        // modal: false,
        // overlay: false,
        // overlayColor: 'default',
        // overlayClickClose: false,
        // type: 'default', // success, alert, warning, info
        // place: 'center', // center, top-left, top-center, top-right, center-left, center-right, bottom-left, bottom-center, bottom-right
        // position: 'default',
        // content: false,
        // hide: false,
        // width: 'auto',
        // height: 'auto',
        // background: 'default',
        // color: 'default',
        // closeButton: false,
        // windowsStyle: false,
        // show: false,
        // href: false,
        // contentType: 'default', // video
        // _interval: undefined,
        // _overlay: undefined,
        // onDialogOpen: function(dialog){},
        // onDialogClose: function(dialog){}
    usp.showDialog = function(options) {
        temp = $("<div data-role='dialog' class='padding20' ></div>")
            .appendTo("body")
            .dialog(options);
        temp.dialog("open")
        return temp;
    }

    usp.Notify = function(title, content, type) {
        // , icon, timeout, keepOpen, shadow, style
        $.Notify({
            caption: title,
            content: content,
            type: type //,
                // icon: icon,
                // timeout: timeout,
                // keepOpen: keepOpen,
                // shadow: shadow,
                // style: style
        });
    }

    usp.init = function() {
        // if (typeof (toastr) == 'object') {
        //     toastr.options.positionClass = 'toast-top-center';
        //     toastr.options.timeOut = '2000';
        //     toastr.options.extendedTimeOut = '1000';
        // }
    };



})(this)