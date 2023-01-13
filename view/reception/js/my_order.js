orderType="<tr><td><font color=\"#ff4e00\">2015092823056</font></td><td>2015-09-26   16:45:20</td><td>未确认，未付款，未发货</td><td type=\"botton\" ><input type=\"button\" value=\"取消订单\" oid=\"\" class=\"log_btn\" onclick=\"login()\" /></td></tr>"

var sta = ["未确认","未付款","未发货","已发货","已取消"]

function getOrder() {
    $.ajax({
        type: "GET",
        url: "/order",
        processData: false,
        contentType: false,
        dataType: "json",
        success: function (data) {
            console.log('login successfully!');
            if (data.code == 0) {
                orderData="<tr><td width=\"20%\">订单号</td><td width=\"25%\">下单时间</td><td width=\"25%\">订单状态</td></tr><br/>"
                for(i=0;i<data.data.length;i++){
                    orderData+="<tr><td><font color=\"#ff4e00\">"+data.data[i].id+"</font></td><td>"+data.data[i].start_time+"</td><td>"+sta[data.data[i].status]+"</td><td type=\"botton\" ><input type=\"button\" value=\"取消订单\" id=\""+data.data[i].id+"\" class=\"log_btn\" onclick=\"login(this.id)\" /></td></tr><br/>"
                    document.getElementById("order_list").innerHTML = orderData;
                }
            } else {
                alert(data.msg)
                console.log('获取订单错误');
            }
        },
        error: function () {
            alert('submit failed!');
        },
    });
}

function login(id) {
    let postData={
        oid:parseInt(id),
    }

    $.ajax({
        type: "POST",
        url: "/cancel_order",
        data: JSON.stringify(postData),
        processData: false,
        contentType: false,
        dataType: "json",
        success: function (data) {
            if (data.code == 0) {
                window.location.href = '/my_order';
            } else {
                alert(data.msg)
            }
        },
        error: function () {
            alert('submit failed!');
        },
    });
}

getOrder()