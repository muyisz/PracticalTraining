
function getCookie(name) {
    var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
    if (arr = document.cookie.match(reg))
        return unescape(arr[2]);
    else
        return null;

}
window.onload = function () {
    user_name = getCookie("user_name")
    if (user_name != null) {
        document.getElementById("hello").innerText = user_name
        document.getElementById("hello").setAttribute("class","f1")
    }
}