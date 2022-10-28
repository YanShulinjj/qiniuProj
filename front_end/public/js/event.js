var canvas = document.getElementById("cs");//获取画布
var context = canvas.getContext("2d");

function setLineWidth(e) {    // this 指向是就是该元素本身
    console.log("你点击了画笔：", e);
    console.log(e.value)
    context.lineWidth = e.value;
    document.getElementById("size").innerHTML = e.value + " px";
}

/* 用户绘制的动作，可以分解为如下操作：
    1.按下鼠标
    2.移动鼠标
    3.松开鼠标

   它们分别对应于鼠标的onmousedown、onmousemove和onmouseup事件。
   并且上述操作必然是有想后顺序的，因为人的操作必然是几个操作
   集合中的一种。所以我们需要来限定以下，过滤用户的无效操作，
   只对按照上诉顺序的操作进行响应。
*/
let isDowned = false;  // 是否按下鼠标，默认是false，如果为false，则不响应任何事件。

// 开始添加鼠标事件
canvas.onmousedown = function(e) {
    let x = e.clientX - canvas.offsetLeft;
    let y = e.clientY - canvas.offsetTop;
    isDowned = true;   // 设置isDowned为true，可以响应鼠标移动事件
    console.log("当前鼠标点击的坐标为：(", x + ", " + y + ")");
    context.strokeStyle = document.getElementById("cl").value;   // 设置颜色，大小已经设置完毕了
    context.beginPath();    // 开始一个新的路径
    context.moveTo(x, y);   // 移动画笔到鼠标的点击位置

    // 多人协作的逻辑
    let pos = {type: 0, x: x, y: y, color: context.strokeStyle, size: context.lineWidth}
    client.send(JSON.stringify(pos))

}

canvas.onmousemove = function(e) {
    if (!isDowned) {
        return ;
    }
    let x = e.clientX - canvas.offsetLeft;
    let y = e.clientY - canvas.offsetTop;
    console.log("当前鼠标的坐标为：(", x + ", " + y + ")");
    context.lineTo(x, y);    // 移动画笔绘制线条
    context.stroke();

    // 多人协作逻辑
    let pos = {type: 1, x: x, y: y, color: context.strokeStyle, size: context.lineWidth}
    client.send(JSON.stringify(pos))
}

canvas.onmouseup = function(e) {
    isDowned = false;
}


/*
    在按下鼠标移动的过程中，如果移出了画布，则无法触发鼠标松开事件，即onmouseup。
    所以需要在鼠标移出画布时，设置isDowned为false。
*/
canvas.onmouseout = function(e) {
    isDowned = false;
}