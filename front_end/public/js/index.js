let svgContainer = document.querySelector('div.svgContainer')
let svg = document.querySelector('.svg')
let svgparent = document.querySelector('.svgparent')
//  下面这块是左边工具的变量名
let toolContainer = document.querySelector('.toolContainer')   // 中间工具栏那块，把svgContainer挡住了，所以事件监听再判断一个
let colorContainer = document.querySelector('.colorContainer')  // 屏幕右边颜色调和colorinput的容器
let toolRangeBox = document.querySelector('.tool-rangeBox')   // 左边所有工具的容器
let range = document.querySelector('.range')   // 滑动条
let rangeText = document.querySelector('em')  // 滑动条的数字显示
let fillColor = document.querySelector('input[id="fill-input"]')   // 背景色的input
let labelFill = document.querySelector('.fill')    //背景色的input关联的label
let fillStart = document.querySelector('input[id="fillstart-input"]')  // 启动fill的按钮
// 下面这块是右边工具的变量名
let colorSelectUl = document.querySelector('.colorSelect') // 右边那一条的颜色框
let colorSelectLis = colorSelectUl.children     /// 这个是颜色框里面的每一个颜色
let laberColors = document.querySelector('label.colors')      // input 外面的label
let colorInput = document.querySelector('input[id="color"]')  // input颜色选择器
let widthInput = document.querySelector('input[type="range"]')  //笔画粗细
// 各种图画工具的变量名
let pen = document.querySelector('#icon-qiniu-huabi')
let linear = document.querySelector('#icon-qiniu-line')
let circle = document.querySelector('#icon-qiniu-ellipse')
let rect = document.querySelector('#icon-qiniu-rectangle')
let eraser = document.querySelector('#icon-qiniu-eraser')
let roundrect = document.querySelector('#icon-qiniu-square')
let textedit = document.querySelector('#icon-qiniu-text')


//按钮
let addPage = document.querySelector('.add')
let undo = document.querySelector('.undo')
let saveFile = document.querySelector('.save')
let openFile = document.querySelector('.open')
let fileInput = document.querySelector('#fileInput')
let clear = document.querySelector('.clear')
let hidden = document.querySelector('.hidden')
let drawandnosave = false   // 没有修改过画面。为true  // 用来判断是否需要保存
let isDrawPolygon = false  // 定义当前是否有多边形正在画
let ishidden = false       // 定义是否隐藏按钮
let mousedownonelement = false

svgContainer.addEventListener('mousedown', (e) => {
    //背景色改变
    if (labelFill.contains(e.target)) {
        svgContainer.addEventListener('mouseup', (fillE) => {
            labelFill.style.backgroundColor = fillColor.value
        })
    }

    // 下面是改变input颜色选择器的外面label的颜色也跟着改变
    if (laberColors.contains(e.target)) {    // 这个是修改颜色，把html那个颜色样式也修改
        svgContainer.addEventListener('mouseup' ,(colosE) => {
            laberColors.style.backgroundColor = colorInput.value
        })

    }

    // 下面是改变滑动条的数字显示
    if (range.contains(e.target)) {     // 修改滑动条的那个文字
        svgContainer.addEventListener('mousemove', function rangeE () {
            rangeText.textContent = range.lastElementChild.value
        })
    }

    // 下面是点击右边几个默认颜色，上面的input颜色选择器也跟改，并且svg画图颜色也更改
    if (colorSelectUl.contains(e.target)) {
        if (colorSelectLis[0].contains(e.target)) {
            laberColors.style.backgroundColor =  colorSelectLis[0].style.backgroundColor
            colorInput.value = colorSelectLis[0].style.backgroundColor.colorHex()
        }
        if (colorSelectLis[1].contains(e.target)) {
            laberColors.style.backgroundColor =  colorSelectLis[1].style.backgroundColor
            colorInput.value = colorSelectLis[1].style.backgroundColor.colorHex()
        }
        if (colorSelectLis[2].contains(e.target)) {
            laberColors.style.backgroundColor =  colorSelectLis[2].style.backgroundColor
            colorInput.value = colorSelectLis[2].style.backgroundColor.colorHex()
        }
        if (colorSelectLis[3].contains(e.target)) {
            laberColors.style.backgroundColor =  colorSelectLis[3].style.backgroundColor
            colorInput.value = colorSelectLis[3].style.backgroundColor.colorHex()
        }
        if (colorSelectLis[4].contains(e.target)) {
            laberColors.style.backgroundColor =  colorSelectLis[4].style.backgroundColor
            colorInput.value = colorSelectLis[4].style.backgroundColor.colorHex()
        }
        if (colorSelectLis[5].contains(e.target)) {
            laberColors.style.backgroundColor =  colorSelectLis[5].style.backgroundColor
            colorInput.value = colorSelectLis[5].style.backgroundColor.colorHex()
        }
        if (colorSelectLis[6].contains(e.target)) {
            laberColors.style.backgroundColor =  colorSelectLis[6].style.backgroundColor
            colorInput.value = colorSelectLis[6].style.backgroundColor.colorHex()
        }
        if (colorSelectLis[7].contains(e.target)) {
            laberColors.style.backgroundColor =  colorSelectLis[7].style.backgroundColor
            colorInput.value = colorSelectLis[7].style.backgroundColor.colorHex()
        }

    }

    // 下面这个判断的意思是，是中间那块大的toolContainer 但不是两边小的工具容器和颜色容易，就是两边两块工具, 所有工具实现都在这个if里面
    if ( (svg.contains(e.target) || toolContainer.contains(e.target))   && !(colorContainer.contains(e.target) || (toolRangeBox.contains(e.target)))){

        drawandnosave = true   // 改过了画面

        // 这个是画笔工具 或者 橡皮擦工具
        if (pen.checked || eraser.checked) {
            let penStartPos = mousePos(svg)   // 拿到鼠标开始画的位置
            let polyline = document.createElementNS("http://www.w3.org/2000/svg", 'polyline')
            // 添加一堆Attribute
            if (pen.checked){
                // 当笔选中时
                polyline.setAttribute('stroke', colorInput.value)
                polyline.setAttribute('stroke-width', widthInput.value)
            }else if (eraser.checked) {
                polyline.setAttribute('stroke', '#ffffff')
                // 如果是橡皮擦选中时，将宽度变成10倍
                polyline.setAttribute('stroke-width', 10*widthInput.value)
            }


            polyline.setAttribute('stroke-linecap', 'round')
            polyline.setAttribute('stroke-linejoin', 'round')
            if (fillStart.checked){
                polyline.setAttribute('fill', fillColor.value)
            }else {
                polyline.setAttribute('fill', 'none')
            }

            svg.append(polyline)
            console.log(2)
            let points = `${penStartPos.x} ${penStartPos.y} `   // 鼠标坐标
            polyline.setAttribute('points', points)   // 位置属性

            function drawDot(e) {   // 鼠标持续移动，持续触发这函数，持续画线
                let penMovePos = mousePos(svg)
                let line = document.createElementNS("http://www.w3.org/2000/svg", 'line')
                points += `${penMovePos.x} ${penMovePos.y} `
                polyline.setAttribute('points', points)
            }

            svgContainer.addEventListener('mousemove', drawDot)

            svgContainer.addEventListener('mouseup',  (once) => {
                svgContainer.removeEventListener('mouseup', once)
                svgContainer.removeEventListener('mousemove', drawDot)
            })
        }

        // 下面是画圆形,椭圆
        if (circle.checked) {
            let ellipse = document.createElementNS("http://www.w3.org/2000/svg", 'ellipse')   // 创建圆的标签
            svg.append(ellipse)  // 插入svg
            // 然后加一堆属性
            ellipse.setAttribute('stroke', colorInput.value)
            ellipse.setAttribute('stroke-width', widthInput.value)
            if (fillStart.checked){
                ellipse.setAttribute('fill', fillColor.value)
            }else {
                ellipse.setAttribute('fill', 'none')
            }

            let startPos = mousePos(svg)   // 获取鼠标相对svg的位置

            function drawEllipse() {
                let currPos = mousePos(svg)
                let cx = (startPos.x + currPos.x) / 2
                let cy = (startPos.y + currPos.y) / 2
                ellipse.setAttribute('cx', cx)
                ellipse.setAttribute('cy', cy)
                let rx = Math.abs(startPos.x - currPos.x) / 2
                let ry = Math.abs(startPos.y - currPos.y) / 2
                ellipse.setAttribute('rx', rx)
                ellipse.setAttribute('ry', ry)
            }

            document.addEventListener('mousemove', drawEllipse)   // 持续运行

            document.addEventListener('mouseup', function once() {  // 解绑
                document.removeEventListener('mouseup', once)
                document.removeEventListener('mousemove', drawEllipse)
            })
        }
        // 矩形和圆角矩形
        if (rect.checked || roundrect.checked) {
            // cosole.log(111)
            let rect = document.createElementNS("http://www.w3.org/2000/svg", 'rect')
            svg.append(rect)
            let rectStartPos = mousePos(svg)   // 鼠标位置

            rect.setAttribute('stroke', colorInput.value)
            rect.setAttribute('stroke-width', widthInput.value)
            rect.setAttribute('x', rectStartPos.x)
            rect.setAttribute('y', rectStartPos.y)
            if (fillStart.checked){
                rect.setAttribute('fill', fillColor.value)
            }else {
                rect.setAttribute('fill', 'none')
            }
            function drawRect () {
                let rectMovePos = mousePos(svg)
                let x = rectStartPos.x
                let y = rectStartPos.y
                let rx,ry
                // 右下角往左上角画矩形

                if (rectMovePos.x < rectStartPos.x) {
                    x = rectMovePos.x
                }
                if (rectMovePos.y < rectStartPos.y) {
                    y = rectMovePos.y
                }

                let width = Math.abs(rectStartPos.x - rectMovePos.x)   // 一正一负，而宽高只能是正
                let height = Math.abs(rectStartPos.y - rectMovePos.y)

                // 圆角矩形
                if (roundrect.checked) {
                    if (width < 50 && height < 50) {
                        rx = (width + height) / 10
                        ry = (width + height) / 10
                    } else {
                        rx = (Math.min(width, height) - 50) / 50 + 10
                        ry = (Math.min(width, height) - 50) / 50 + 10
                    }
                }

                rect.setAttribute('width', width)
                rect.setAttribute('height', height)
                rect.setAttribute('x', x)
                rect.setAttribute('y', y)
                rect.setAttribute('rx', rx)
                rect.setAttribute('ry', ry)

            }
            svgContainer.addEventListener('mousemove', drawRect)

            svgContainer.addEventListener('mouseup', (rectE) => {
                svgContainer.removeEventListener('mousemove', drawRect)
                svgContainer.removeEventListener('mouseuo', rectE)
            })
        }
        // 直线
        if (linear.checked) {
            let linear = document.createElementNS("http://www.w3.org/2000/svg", 'line')
            svg.append(linear)
            let linearStartPos = mousePos(svg)

            linear.setAttribute('stroke', colorInput.value)
            linear.setAttribute('stroke-width', widthInput.value)
            linear.setAttribute('stroke-linecap', 'round')
            linear.setAttribute('stroke-linejoin', 'round')
            linear.setAttribute('x1', linearStartPos.x)
            linear.setAttribute('y1', linearStartPos.y)
            linear.setAttribute('x2', linearStartPos.x)
            linear.setAttribute('y2', linearStartPos.y)
            if (fillStart.checked){
                linear.setAttribute('fill', fillColor.value)
            }else {
                linear.setAttribute('fill', 'none')
            }

            function drawLinear () {
                let linearMoverPos = mousePos(svg)

                linear.setAttribute('x2', linearStartPos.x + (linearMoverPos.x - linearStartPos.x)) // 鼠标初始位置加上终点位置减去起始位置
                linear.setAttribute('y2', linearStartPos.y + (linearMoverPos.y - linearStartPos.y))
            }
            svgContainer.addEventListener('mousemove', drawLinear)

            svgContainer.addEventListener('mouseup', (linerE) => {
                svgContainer.removeEventListener('mouseup', linerE)
                svgContainer.removeEventListener('mousemove', drawLinear)
            })
        }
        // TODO: 文本
        if (textedit.checked) {
            var localpoint = window.getlocalmousecoord(svg, e);
            if (!mousedownonelement) {
                createtext(localpoint, svg);
            } else {
                mousedownonelement = false;
            }
        }



    }
})


// 将所有按钮隐藏
hidden.addEventListener('click', (e) => {
    //
    console.log("点击了隐藏")
    ishidden = !ishidden
    if (ishidden)  {
        // 左右侧隐藏
        document.getElementsByClassName("toolContainer")[0].style.display="none"
        var elements = document.getElementsByClassName("hiddenbutton")
        var i
        for (i = 0; i < elements.length; i++) {
            elements[i].style.display = "none";
        }
        // 将顶部按钮宽度更新
        document.getElementsByClassName("addundobar")[0].style.width = "80px"
    } else {
        // 左右侧显示
        document.getElementsByClassName("toolContainer")[0].style.display="flex"
        var elements = document.getElementsByClassName("hiddenbutton")
        var i
        for (i = 0; i < elements.length; i++) {
            elements[i].style.display = "block";
        }
        // 将顶部按钮宽度更新
        document.getElementsByClassName("addundobar")[0].style.width = "400px"

    }

})


// 清除所画内容
clear.addEventListener('click', (clickE) => {
    svg.innerHTML = ''
})
// 打开一个新文件
openFile.addEventListener('click', function openLoaclFile () {
    if (drawandnosave) {
        var answer = confirm('当前绘画未保存，确定要打开新文件吗？')
        if (answer == false) {
            return
        }
    }
    fileInput.click()   // 如果用户点击是
})

fileInput.addEventListener('change', e => {
    var svgFile = fileInput.files[0]
    var fr = new FileReader()

    // 读取svg文件
    fr.addEventListener('load', () => {
        var svgFileContent = fr.result
        svgparent.innerHTML = svgFileContent
        svg = document.querySelector('.svg')
    })
    fr.readAsText(svgFile)
})

// 撤销
document.addEventListener("keydown", function (event) {
    if (event.code == "KeyZ" && (event.ctrlKey || event.metaKey)) {
        if (svg.lastChild) svg.lastChild.remove()
    }
})

// 下面是退出浏览器提示
window.onbeforeunload = function () {
  if (drawandnosave) {
      // TODO: 持久化此页面
      console.log("点击了退出按钮")
  }
}

// 将图画保存到桌面
saveFile.addEventListener('click' ,function save(){
    drawandnosave = false

    var svgSource = svg.outerHTML
    var blob = new Blob(['<?xml version="1.0" encoding="utf-8"?>', svgSource], { type: "image/xml+svg" })
    var url = URL.createObjectURL(blob)
    var anchor = document.createElement("a")
    anchor.href = url
    anchor.download = "xxxx.svg"
    anchor.click()  // 点击button 也触发a标签点击
    // TODO：持久化此页面
})
addPage.addEventListener("click", function add(){
    // TODO: 添加新页面
    console.log("点击了add")
})

// 撤销上一步
undo.addEventListener("click", function (e) {
    if (svg.lastChild) svg.lastChild.remove()
})

// TODO redo

// 鼠标相对于元素的位置
function mousePos(node) {
    var box = node.getBoundingClientRect()

    return {
        x: window.event.clientX - box.x,
        y: window.event.clientY - box.y,
    }
}

String.prototype.colorHex = function () {
    // RGB颜色值的正则
    var reg = /^(rgb|RGB)/;
    var color = this;
    if (reg.test(color)) {
        var strHex = "#";
        // 把RGB的3个数值变成数组
        var colorArr = color.replace(/(?:\(|\)|rgb|RGB)*/g, "").split(",");
        // 转成16进制
        for (var i = 0; i < colorArr.length; i++) {
            var hex = Number(colorArr[i]).toString(16);
            if (hex === "0") {
                hex += hex;
            }
            strHex += hex;
        }
        return strHex;
    } else {
        return String(color);
    }
}