
function login() {
    let postData={
        user_name:$('#user_name').val(),
        password:$('#password').val()
    }

    $.ajax({
        type: "POST",
        url: "/login",
        data: JSON.stringify(postData),
        processData: false,
        contentType: false,
        dataType: "json",
        success: function (data) {
            console.log('login successfully!');
            if (data.code == 0) {
                window.location.href = '/home';
            } else {
                alert(data.msg)
                console.log('login failed!');
            }
        },
        error: function () {
            alert('submit failed!');
        },
    });
}