let userId = 0
let userName = ""
let pageName = document.getElementById("pageName").innerText
var httpRequest = new XMLHttpRequest()
var url = 'http://127.0.0.1:8080/backend/user/add?username=unnamed'
httpRequest.open('GET', url, true)
httpRequest.send()
httpRequest.onreadystatechange = function () {
    console.log(httpRequest)
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
        var data = httpRequest.responseText
        var json = JSON.parse(data)
        userName = json.user_name
        userId = json.user_id
        console.log(userId)
        ws ()  // 直接建立websocket连接
    }
};
