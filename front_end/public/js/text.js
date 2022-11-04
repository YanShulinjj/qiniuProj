// 点击画板生成文本

window.getlocalmousecoord = function (svg, evt) {
    var pt = svg.createSVGPoint();
    pt.x = evt.clientX;
    pt.y = evt.clientY;
    var localpoint = pt.matrixTransform(svg.getScreenCTM().inverse());
    localpoint.x = Math.round(localpoint.x);
    localpoint.y = Math.round(localpoint.y);
    return localpoint;
};

// window.createtext = function (localpoint, svg) {
//     var myforeign = document.createElementNS('http://www.w3.org/2000/svg', 'foreignObject')
//     var textdiv = document.createElement("div");
//     var textnode = document.createTextNode("click");
//     textdiv.appendChild(textnode);
//     textdiv.setAttribute("contentEditable", "true");
//     textdiv.setAttribute("width", "auto");
//     myforeign.setAttribute("width", "100%");
//     myforeign.setAttribute("height", "100%");
//     myforeign.classList.add("foreign"); //to make div fit text
//     textdiv.classList.add("insideforeign"); //to make div fit text
//     textdiv.addEventListener("mousedown", elementMousedown, false);
//     myforeign.setAttributeNS(null, "transform", "translate(" + localpoint.x + " " + localpoint.y + ")");
//     svg.appendChild(myforeign);
//     myforeign.appendChild(textdiv);
//
// };

window.createtext = function (localpoint, svg) {
    let text = document.createElementNS("http://www.w3.org/2000/svg", 'text')   // 创建圆的标签

    // 然后加一堆属性
    content=prompt("请输入文本：");//弹出框执行的优先级别要高于返回值：2
    console.log(content);//4
    // 如果取消 直接返回
    if (content == null) {
        return
    }

    text.setAttribute('fill', colorInput.value)
    text.setAttribute('x', localpoint.x)
    text.setAttribute('y', localpoint.y)
    text.setAttribute('font-size', 12*widthInput.value)
    var textnode = document.createTextNode(content);

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
    // 多人协作
    let msg = {
        type: TextType,
        Attr: {
            x: localpoint.x,
            y: localpoint.y,
            color: colorInput.value,
            font_size: text.getAttribute('font-size'),
            fillValue: text.getAttribute('fill'),
            content: content
        }
    }
    client.send(JSON.stringify(msg))
}

function elementMousedown(evt) {
    mousedownonelement = true;
}