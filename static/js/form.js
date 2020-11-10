$(document).ready(function () {
    $('.ajax-form').ajaxForm({
        dataType: "json",
        error: function (a, b, c) {
            notify("请求错误，请重试。响应码：" + a.status, "danger")
        },
        success: function (data) {
            if (data.code == 200) {
                notify(data.msg || "操作成功！", "success")
                if (data.url) {
                    setTimeout(function () {
                        location = data.url
                    }, 1500)
                }
            } else {
                notify(data.msg || "操作失败，请重试。", "warning")
            }
        }
    })
});