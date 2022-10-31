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
    // let startPos = mousePos(svg)   // 获取鼠标相对svg的位置
    //
    // function drawEllipse() {
    //     let currPos = mousePos(svg)
    //     let cx = (startPos.x + currPos.x) / 2
    //     let cy = (startPos.y + currPos.y) / 2
    //     ellipse.setAttribute('cx', cx)
    //     ellipse.setAttribute('cy', cy)
    //     let rx = Math.abs(startPos.x - currPos.x) / 2
    //     let ry = Math.abs(startPos.y - currPos.y) / 2
    //     ellipse.setAttribute('rx', rx)
    //     ellipse.setAttribute('ry', ry)
    // }
    //
    // document.addEventListener('mousemove', drawEllipse)   // 持续运行
    //
    // document.addEventListener('mouseup', function once() {  // 解绑
    //     document.removeEventListener('mouseup', once)
    //     document.removeEventListener('mousemove', drawEllipse)
    // })
}

function elementMousedown(evt) {
    mousedownonelement = true;
}