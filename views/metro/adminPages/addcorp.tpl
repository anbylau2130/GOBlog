<h2>corpAdd <button id="btnBack" data-bind="click: usp.admin.corp.add.goBack" class="button place-right primary"><span class="mif-keyboard-return "></span></button></h2>
<hr class="thin bg-grayLighter">
<form data-role="validator" data-hint-mode="hint" data-on-submit="return usp.admin.corp.add.save()">
    <div class="grid">
        <div class="row cells2">
            <div class="cell">
                <label>ID</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:ID" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Parent</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Parent" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Priority</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Priority" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Name</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Name" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>BriefName</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:BriefName" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Certificate</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Certificate" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>CertificateNumber</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:CertificateNumber" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Ceo</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Ceo" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Postcode</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Postcode" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Faxcode</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Faxcode" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Linkman</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Linkman" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Mobile</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Mobile" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Email</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Email" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Qq</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Qq" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Wechat</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Wechat" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Weibo</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Weibo" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>VirtualIntegral</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:VirtualIntegral" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>RealIntegral</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:RealIntegral" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>FeeType</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:FeeType" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>PrepayValve</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:PrepayValve" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Balance</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Balance" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>FrozenBalance</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:FrozenBalance" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>IncomingBalance</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:IncomingBalance" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Commission</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Commission" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Discount</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Discount" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Province</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Province" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Area</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Area" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>County</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:County" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Community</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Community" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Address</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Address" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Status</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Status" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Type</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Type" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Grade</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Grade" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Vocation</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Vocation" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Reserve</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Reserve" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>Remark</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Remark" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Creator</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Creator" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>CreateTime</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:CreateTime" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Auditor</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Auditor" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>AuditTime</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:AuditTime" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>Canceler</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:Canceler" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>CancelTime</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:CancelTime" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
        <div class="row cells2">
            <div class="cell">
                <label>LogoIcon</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:LogoIcon" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
            <div class="cell">
                <label>LogoUrl</label>
                <div class="input-control text full-size">
                    <input type="text" data-bind="value:LogoUrl" data-validate-func="required" data-validate-hint-position="top" data-validate-hint="不能为空!" />
                    <span class="input-state-error mif-warning"></span>
                    <span class="input-state-success mif-checkmark"></span>
                </div>
            </div>
        </div>
    </div>
    <button data-bind='click:usp.admin.corp.add.save' class="button place-right primary">保存</button>
</form>
<script src="/static/js/usp.admin.corp.add.js"></script>
<script>
    $(function() {
        usp.init()
        usp.admin.corp.add.init()
    })
</script>