// 客户端与服务端建立websocket链接
function ws () {
    client = new WebSocket("ws://"+HostAddr+"/ws/wedraw?username="+userName+"&page="+pageName+"&author="+pageAuthorName);    //连接服务器
    // client.onopen = function(e){
    //     alert('连接服务器成功！');
    // };

    // 从服务器接收数据
    client.onmessage = function (e) {
        let data = e.data
        let msg = JSON.parse(data)

        // console.log("接受到的消息：" + data)

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
                polylineNum = parseInt(msg.Attr.id) + 1
                // 将此操作加入队列
                // 添加之前删除index 后面所有的elem
                elementQueue.splice(index+1, elementQueue.length-index-1)
                elementQueue.push({type: CommonType, value: polyline})
                if (elementQueue.length > maxSize) {
                    // 移除队头
                    elementQueue.shift()
                } else {
                    index ++
                }
                break
            case DotType:
                var polyline = document.getElementById(msg.Attr.id)

                points = polyline.getAttribute("points")
                points += msg.Attr.appendPoint
                polyline.setAttribute("points", points)
                break
            case LineType:
                var linear = document.getElementById('line_'+msg.Attr.id)

                if (msg.Attr.isEnd == true) {
                    // 将此操作加入队列
                    // 添加之前删除index 后面所有的elem
                    elementQueue.splice(index+1, elementQueue.length-index-1)
                    elementQueue.push({type: CommonType, value: linear})
                    if (elementQueue.length > maxSize) {
                        // 移除队头
                        elementQueue.shift()
                    } else {
                        index ++
                    }
                    break
                }

                if (linear == null) {
                    linear = document.createElementNS("http://www.w3.org/2000/svg", 'line')
                    linear.setAttribute('id', 'line_' + msg.Attr.id)
                    lineNum = parseInt(msg.Attr.id) + 1
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

                if (msg.Attr.isEnd == true) {
                    // 将此操作加入队列
                    // 添加之前删除index 后面所有的elem
                    elementQueue.splice(index+1, elementQueue.length-index-1)
                    elementQueue.push({type: CommonType, value: ellipse})
                    if (elementQueue.length > maxSize) {
                        // 移除队头
                        elementQueue.shift()
                    } else {
                        index ++
                    }
                    break
                }

                if (ellipse == null) {
                    ellipse = document.createElementNS("http://www.w3.org/2000/svg", 'ellipse')
                    ellipse.setAttribute('id', 'ellipse_' + msg.Attr.id)
                    ellipseNum = parseInt(msg.Attr.id) + 1
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

                if (msg.Attr.isEnd === true) {
                    // 将此操作加入队列
                    // 添加之前删除index 后面所有的elem
                    console.log("attr is false")
                    elementQueue.splice(index+1, elementQueue.length-index-1)
                    elementQueue.push({type: CommonType, value: rect})
                    if (elementQueue.length > maxSize) {
                        // 移除队头
                        elementQueue.shift()
                    } else {
                        index ++
                    }
                    break
                }

                if (rect == null) {
                    rect = document.createElementNS("http://www.w3.org/2000/svg", 'rect')
                    rect.setAttribute('id', 'rect_' + msg.Attr.id)
                    rectangleNum = parseInt(msg.Attr.id) + 1
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
                // 将此操作加入队列
                // 添加之前删除index 后面所有的elem
                elementQueue.splice(index+1, elementQueue.length-index-1)
                elementQueue.push({type: CommonType, value: text})
                if (elementQueue.length > maxSize) {
                    // 移除队头
                    elementQueue.shift()
                } else {
                    index ++
                }
                break
            case UndoType:
                if (elementQueue.length > 0 && index >= 0) {
                    if (elementQueue[index].type == CommonType) {
                        svg.lastChild.remove()
                    } else if (elementQueue[index].type == ClearType) {
                        svg.innerHTML = elementQueue[index].value
                    }
                    index --
                }
                break
            case RedoType:
                // TODO
                if (elementQueue.length > 0 && index < elementQueue.length -1) {
                    index ++
                    if (elementQueue[index].type == CommonType) {
                        svg.append(elementQueue[index].value)
                    } else if (elementQueue[index].type == ClearType) {
                        svg.innerHTML = ''
                    }
                }
                break
            case ClearType:
                // 将此操作加入队列
                // 添加之前删除index 后面所有的elem
                elementQueue.splice(index+1, elementQueue.length-index-1)
                elementQueue.push({type: ClearType, value: svg.innerHTML})
                if (elementQueue.length > maxSize) {
                    // 移除队头
                    elementQueue.shift()
                } else {
                    index ++
                }
                svg.innerHTML = ''
                break
            case LoadType:
                svgparent.innerHTML = msg.Attr.content
                svg = document.querySelector('.svg')
                break
            case ModeChangeType:
                readonly = msg.Attr
                // 将按钮图标更换
                if (readonly) {
                    if (pageAuthorName != userName) {
                        // 直接隐藏
                        hiddenElems()
                    }
                    document.getElementsByClassName('icon-qiniu-readonly')[0].setAttribute('data-before', 'R')
                } else {
                    if (pageAuthorName != userName) {
                        displayElems()
                    }
                    document.getElementsByClassName('icon-qiniu-readonly')[0].setAttribute('data-before', 'W')
                }
                break
            default:
                console.log("不支持的消息类型")
        }
    }

    client.onclose = function(e){
        alert("连接失败, 检测是否重复建立连接：" + this.readyState);
    };

    client.onerror = function(e){
        alert("连接失败, 检测是否重复建立连接！");
    };
}

function syncws() {
    syncclient = new WebSocket("ws://"+HostAddr+"/ws/statue?username="+userName+"&page="+pageName);    //连接服务器
    // client.onopen = function(e){
    //     alert('连接服务器成功！');
    // };

    // 从服务器接收数据
    syncclient.onmessage = function (e) {
        let data = e.data
        let msg = JSON.parse(data)
        console.log(" ***** 接受到的消息：" + data)
        if (msg.type == NeedSyncType) {
            // 发送自己面板上的全部元素
            syncMsg()
        } else if (msg.type == InitializeType) {
            svg.innerHTML = msg.Attr.content
            // 为了实现能够同步undo,redo功能
            // 需要将末尾N个入队列
            startIdx = svg.children.length - maxSize < 0? 0: svg.children.length - maxSize
            for (i=startIdx; i<svg.children.length; i++) {
                elementQueue.push({type: CommonType, value: svg.children[i]})
                index ++
            }
            console.log("Queue:", elementQueue)

            // 如果不是作者，关闭连接
            if (pageAuthorName != userName) {
                syncclient.close()
            }
        } else {
            console.log("不支持的消息类型")
        }

    }
}

function syncMsg() {
    // 如果是作者
    // 发送同步消息
    if (pageAuthorName == userName) {
        let syncmsg = {
            type: InitializeType,
            Attr: {
                content: svg.innerHTML
            }
        }
        console.log("send syncmsg")
        syncclient.send(JSON.stringify(syncmsg))
    }
}