// 客户端与服务端建立websocket链接
function link () {
    client = new WebSocket("ws://localhost:9999/ws/wedraw");    //连接服务器
    client.onopen = function(e){
        alert('连接服务器成功！');
    };

    // 从服务器接受数据
    client.onmessage = function (e) {
        let data = e.data
        let msg = JSON.parse(data)

        console.log("接受到的消息：" + data)

        switch (msg.type) {
            case PolylineType:
                var polyline = document.createElementNS("http://www.w3.org/2000/svg", 'polyline')
                polyline.setAttribute('id', msg.Attr.id)
                polyline.setAttribute('stroke', msg.Attr.color)
                polyline.setAttribute('stroke-width', msg.Attr.stroke_width)
                polyline.setAttribute('stroke-linecap', 'round')
                polyline.setAttribute('stroke-linejoin', 'round')
                polyline.setAttribute('fill', msg.Attr.fillValue)
                polyline.setAttribute('points', msg.Attr.points)
                svg.append(polyline)
                break
            case DotType:
                var polyline = document.getElementById(msg.Attr.id)
                polylineNum = msg.Attr.id + 1
                points = polyline.getAttribute("points")
                points += msg.Attr.appendPoint
                polyline.setAttribute("points", points)
                break
            case LineType:
                var linear = document.getElementById('line_'+msg.Attr.id)
                if (linear == null) {
                    linear = document.createElementNS("http://www.w3.org/2000/svg", 'line')
                    linear.setAttribute('id', 'line_' + msg.Attr.id)
                    lineNum = msg.Attr.id+1
                }
                linear.setAttribute('stroke', msg.Attr.color)
                linear.setAttribute('stroke-width', msg.Attr.stroke_width)
                linear.setAttribute('stroke-linecap', 'round')
                linear.setAttribute('stroke-linejoin', 'round')
                linear.setAttribute('x1', msg.Attr.x1)
                linear.setAttribute('y1', msg.Attr.y1)
                linear.setAttribute('x2', msg.Attr.x2)
                linear.setAttribute('y2', msg.Attr.y2)
                linear.setAttribute('fill', msg.Attr.fillValue)
                svg.append(linear)
                break
            case CircleType:
                var ellipse = document.getElementById('ellipse_' + msg.Attr.id)
                if (ellipse == null) {
                    ellipse = document.createElementNS("http://www.w3.org/2000/svg", 'ellipse')
                    ellipse.setAttribute('id', 'ellipse_' + msg.Attr.id)
                    ellipseNum = msg.Attr.id+1
                }

                // 然后加一堆属性
                ellipse.setAttribute('stroke', msg.Attr.color)
                ellipse.setAttribute('stroke-width', msg.Attr.stroke_width)
                ellipse.setAttribute('fill', msg.Attr.fillValue)
                ellipse.setAttribute('cx', msg.Attr.cx)
                ellipse.setAttribute('cy', msg.Attr.cy)
                ellipse.setAttribute('rx', msg.Attr.rx)
                ellipse.setAttribute('ry', msg.Attr.ry)
                svg.append(ellipse)
                break
            case RetangleType:
                var rect = document.getElementById('rect_'+msg.Attr.id)
                console.log("hhh")
                if (rect == null) {
                    rect = document.createElementNS("http://www.w3.org/2000/svg", 'rect')
                    rect.setAttribute('id', 'rect_' + msg.Attr.id)
                    rectangleNum = msg.Attr.id+1
                }
                rect.setAttribute('stroke', msg.Attr.color)
                rect.setAttribute('stroke-width', msg.Attr.stroke_width)
                rect.setAttribute('fill', msg.Attr.fillValue)
                rect.setAttribute('width', msg.Attr.width)
                rect.setAttribute('height', msg.Attr.height)
                rect.setAttribute('x', msg.Attr.x)
                rect.setAttribute('y', msg.Attr.y)
                rect.setAttribute('rx', msg.Attr.rx)
                rect.setAttribute('ry', msg.Attr.ry)
                svg.append(rect)
                break
            case TextType:
                var text = document.createElementNS("http://www.w3.org/2000/svg", 'text')   // 创建圆的标签
                text.setAttribute('fill', msg.Attr.fillValue)
                text.setAttribute('x', msg.Attr.x)
                text.setAttribute('y', msg.Attr.y)
                text.setAttribute('font-size', msg.Attr.font_size)
                var textnode = document.createTextNode(msg.Attr.content);

                text.appendChild(textnode)
                svg.append(text)  // 插入svg
                break
            case UndoType:
                if (svg.lastChild) {
                    svg.lastChild.remove()
                }
                break
            case RedoType:
                // TODO
                break
            case ClearType:
                svg.innerHTML = ''
                break
            default:
                console.log("不支持的消息类型")
        }
    }

    client.onclose = function(e){
        alert("已经与服务器断开连接\r\n当前连接状态：" + this.readyState);
    };

    client.onerror = function(e){
        alert("WebSocket异常！");
    };
}

link ()  // 直接建立websocket连接