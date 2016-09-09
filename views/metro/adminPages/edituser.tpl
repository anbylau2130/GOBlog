<h2>修改用户信息 <button id="return" data-bind="click: usp.admin.user.edit.goBack" class="button place-right primary"><span class="mif-keyboard-return "></span></button></h2>
<hr class="thin bg-grayLighter">
<form data-role="validator" data-hint-mode="hint" data-on-submit="return usp.admin.user.edit.save()">
    <div class="grid padding20">

        <div class="row cells2">
            <div class="cell">
                <label>登录名</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:LoginName" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>真实姓名</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:RealName" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>公司名称</label>
                <div class="input-control text full-size">
                    <input type="hidden" data-bind="value:Password" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <input type="hidden" data-bind="value:ID" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <input type="text" data-bind="value:Corp" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>手机号</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Mobile" data-validate-func="required,minlength,maxlength" data-validate-arg=",11,11" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>身份证</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:IdCard" data-validate-func="required,minlength,maxlength" data-validate-arg=",18,18" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Email</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Email" data-validate-func="required,email" data-validate-hint-position="top" data-validate-hint="不能为空且必须为EMAIL格式！" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>微信</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:WechatOpenid" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>AlipayOpenid</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:AlipayOpenid" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Weibo</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Weibo" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>AvailableIP</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:AvailableIP" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>WeatherCode</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:WeatherCode" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>VirtualIntegral</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:VirtualIntegral" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>RealIntegral</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:RealIntegral" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Balance</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Balance" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>FrozenBalance</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:FrozenBalance" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>IncomingBalance</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:IncomingBalance" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Commission</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Commission" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Discount</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Discount" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Province</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Province" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Area</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Area" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>County</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:County" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Community</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Community" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Address</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Address" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Status</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Status" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Skin</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Skin" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Grade</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Grade" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Star</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Star" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Session</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Session" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>LoginTime</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:LoginTime" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>LoginIP</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:LoginIP" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>LoginAgent</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:LoginAgent" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>LoginCount</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:LoginCount" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>LoginErrorCount</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:LoginErrorCount" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>FrozenStartTime</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:FrozenStartTime" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>FrozenEndTime</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:FrozenEndTime" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Reserve</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Reserve" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Remark</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Remark" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Creator</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Creator" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>CreateTime</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:CreateTime" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Auditor</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Auditor" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>AuditTime</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:AuditTime" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Canceler</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Canceler" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>CancelTime</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:CancelTime" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">

            </div>
        </div>
        <button class="button place-right primary">保存</button>
    </div>
</form>
</div>
<script src="/static/js/usp.admin.user.edit.js"></script>
<script>
    $(function() {
        var id = '{{.Model.ID}}'
        usp.init()
        usp.admin.user.edit.init(id)
    })
</script>