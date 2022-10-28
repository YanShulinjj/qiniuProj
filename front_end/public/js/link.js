// 客户端与服务端建立websocket链接
function link () {
    client = new WebSocket("ws://localhost:9999/ws/wedraw");    //连接服务器
    client.onopen = function(e){
        alert('连接服务器成功！');
    };

    client.onmessage = function (e) {
        let data = e.data
        let pos = JSON.parse(data)

        console.log("接受到的消息：" + data)

        context.strokeStyle = pos.color   // 设置颜色
        context.lineWidth = pos.size      // 设置线宽
        if (pos.type === 0) {             // 如果该点是移动画笔，则移动画笔
            context.beginPath()           // 开始一个新的路径
            context.moveTo(pos.x, pos.y)
        } else if (pos.type === 1) {      // 如果该点是画线，就画线
            context.lineTo(pos.x, pos.y);
            context.stroke();                  // 绘制点
        } else {
            console.log("不存在的情况，直接返回")
            return
        }
    }

    client.onclose = function(e){
        alert("已经与服务器断开连接\r\n当前连接状态：" + this.readyState);
    };

    client.onerror = function(e){
        alert("WebSocket异常！");
    };
}

function sendMsg(position){
    client.send(position);
}
// link ()  // 直接建立websocket连接