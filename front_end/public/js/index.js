/* ----------------------- 变量定义区 -------------------------------- */
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
let rect = document.querySelector('#icon-qiniu-square')
let eraser = document.querySelector('#icon-qiniu-eraser')
let roundrect = document.querySelector('#icon-qiniu-rectangle')
let textedit = document.querySelector('#icon-qiniu-text')
let tips = document.querySelector(".tips")


//按钮
let addPage = document.querySelector('.add')
let rwMode = document.querySelector('.rwMode')
let undo = document.querySelector('.undo')
let redo = document.querySelector('.redo')
let share = document.querySelector('.share')
let saveFile = document.querySelector('.save')
let openFile = document.querySelector('.open')
let fileInput = document.querySelector('#fileInput')
let clear = document.querySelector('.clear')
let hidden = document.querySelector('.hidden')
let pagelist = document.querySelector('.icon-qiniu-list')


let drawandnosave = false   // 没有修改过画面。为true  // 用来判断是否需要保存
let isDrawPolygon = false  // 定义当前是否有多边形正在画
let ishidden = false       // 定义是否隐藏按钮
let ispagelist = false     // 定义是否点击了pagelist
let mousedownonelement = false



// type 编号
const PolylineType = 0
const DotType = 1
const LineType = 2
const CircleType = 3
const RoundRectangleType = 4
const RetangleType = 5
const TextType = 6
const UndoType = 7
const RedoType = 8
const ClearType = 9
const LoadType = 10
const CommonType = 11
const ModeChangeType = 12
const NeedSyncType = 13
const InitializeType = 14
const PageChangeType = 15

// 记录各种图形的id
let polylineNum = 1
let ellipseNum = 1
let rectangleNum = 1
let lineNum = 1

// 页面名称
let saveFileName = 'xxxx.svg'

// 用于撤销和反撤销
let elementQueue = []
let queueSize = 0
let maxSize = 10
let index = -1;

// PC 和移动端兼容
let down = {true: 'mousedown', false: 'touchstart'}
let move = {true: 'mousemove', false: 'touchmove'}
let up = {true: 'mouseup', false: 'touchend'}
let platform = IsPC()

if (!platform) {
    alert("移动端建议使用横屏")
}

/* -----------------------事件定义区 -------------------------------- */
svgContainer.addEventListener(down[platform], (e) => {

    // 如果当前只读模式，并且登录者不是该页面的所有者
    if (readonly && userName != pageAuthorName) {
        console.log("只读模式、禁止修改")
        return
    }
    //背景色改变
    if (labelFill.contains(e.target)) {
        svgContainer.addEventListener(up[platform], (fillE) => {
            labelFill.style.backgroundColor = fillColor.value
        })
    }

    // 下面是改变input颜色选择器的外面label的颜色也跟着改变
    if (laberColors.contains(e.target)) {    // 这个是修改颜色，把html那个颜色样式也修改
        svgContainer.addEventListener(up[platform] ,(colosE) => {
            laberColors.style.backgroundColor = colorInput.value
        })

    }

    // 下面是改变滑动条的数字显示
    if (range.contains(e.target)) {     // 修改滑动条的那个文字
        svgContainer.addEventListener(move[platform], function rangeE () {
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
            polyline.setAttribute("id", polylineNum)
            polylineNum ++
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
            console.log(2)

            let points = `${penStartPos.x} ${penStartPos.y} `   // 鼠标坐标
            polyline.setAttribute('points', points)   // 位置属性

            // 多人协作
            let msg = {
                type: PolylineType,
                Attr: {
                    id: polyline.getAttribute('id'),
                    points: polyline.getAttribute('points'),
                    color: polyline.getAttribute('stroke'),
                    stroke_width: polyline.getAttribute('stroke-width'),
                    fillValue: polyline.getAttribute('fill')
                }
            }
            client.send(JSON.stringify(msg))


            function drawDot(e) {   // 鼠标持续移动，持续触发这函数，持续画线
                let penMovePos = mousePos(svg)
                // let line = document.createElementNS("http://www.w3.org/2000/svg", 'line')
                points += `${penMovePos.x} ${penMovePos.y} `
                polyline.setAttribute('points', points)
                // 多人协作
                let msg = {
                    type: DotType,
                    Attr: {
                        id: polyline.getAttribute('id'),
                        appendPoint: `${penMovePos.x} ${penMovePos.y} `
                    }
                }
                client.send(JSON.stringify(msg))
            }

            svgContainer.addEventListener(move[platform], drawDot)
            svgContainer.addEventListener(up[platform],  function onceDot(){

                svgContainer.removeEventListener(up[platform], onceDot)
                svgContainer.removeEventListener(move[platform], drawDot)
            })
        }

        // 下面是画圆形,椭圆
        if (circle.checked) {
            let ellipse = document.createElementNS("http://www.w3.org/2000/svg", 'ellipse')   // 创建圆的标签
            svg.append(ellipse)  // 插入svg
            // 然后加一堆属性
            ellipse.setAttribute('stroke', colorInput.value)
            ellipse.setAttribute('stroke-width', widthInput.value)
            ellipse.setAttribute('id', 'ellipse_' + ellipseNum)
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
                // 多人协作
                let msg = {
                    type: CircleType,
                    Attr: {
                        id: ellipseNum,
                        cx: cx,
                        cy: cy,
                        rx: rx,
                        ry: ry,
                        color: colorInput.value,
                        stroke_width: widthInput.value,
                        fillValue: ellipse.getAttribute('fill')
                    }
                }
                client.send(JSON.stringify(msg))
            }
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

            document.addEventListener(move[platform], drawEllipse)   // 持续运行

            document.addEventListener(up[platform], function once() {  // 解绑
                // 多人协作, 发送结束标记
                let msg = {
                    type: CircleType,
                    Attr: {
                        id: ellipseNum,
                        isEnd: true,
                    },
                }
                client.send(JSON.stringify(msg))
                ellipseNum ++
                document.removeEventListener(up[platform], once)
                document.removeEventListener(move[platform], drawEllipse)
            })
        }
        // 矩形和圆角矩形
        if (rect.checked || roundrect.checked) {
            // cosole.log(111)
            let rect = document.createElementNS("http://www.w3.org/2000/svg", 'rect')
            svg.append(rect)
            let rectStartPos = mousePos(svg)   // 鼠标位置

            rect.setAttribute('id', 'rect_'+rectangleNum)
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

                // 多人协作
                let msg = {
                    type: RetangleType,
                    Attr: {
                        id: rectangleNum,
                        x: x,
                        y: y,
                        rx: rx,
                        ry: ry,
                        color: colorInput.value,
                        stroke_width: widthInput.value,
                        width: width,
                        height: height,
                        fillValue: rect.getAttribute('fill')
                    }
                }
                client.send(JSON.stringify(msg))

            }
            // 将此操作加入队列
            // 添加之前删除index 后面所有的elem
            elementQueue.splice(index+1, elementQueue.length-index-1)
            elementQueue.push({type: CommonType, value: rect})
            if (elementQueue.length > maxSize) {
                // 移除队头
                elementQueue.shift()
            } else {
                index ++
            }

            svgContainer.addEventListener(move[platform], drawRect)
            svgContainer.addEventListener(up[platform], function onceRect(){
                // 多人协作, 发送结束标记
                let msg = {
                    type: RetangleType,
                    Attr: {
                        id: rectangleNum,
                        isEnd: true,
                    },
                }
                client.send(JSON.stringify(msg))
                rectangleNum ++
                console.log("完成矩形勾画")
                svgContainer.removeEventListener(move[platform], drawRect)
                svgContainer.removeEventListener(up[platform], onceRect)
            })

        }
        // 直线
        if (linear.checked) {
            let linear = document.createElementNS("http://www.w3.org/2000/svg", 'line')
            svg.append(linear)
            let linearStartPos = mousePos(svg)
            linear.setAttribute('id', 'line_'+lineNum)
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

                // 多人协作
                let msg = {
                    type: LineType,
                    Attr: {
                        id: lineNum,
                        x1: linear.getAttribute('x1'),
                        y1: linear.getAttribute('y1'),
                        x2: linear.getAttribute('x2'),
                        y2: linear.getAttribute('y2'),
                        color: colorInput.value,
                        stroke_width: widthInput.value,
                        fillValue: linear.getAttribute('fill')
                    }
                }
                client.send(JSON.stringify(msg))

            }
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

            svgContainer.addEventListener(move[platform], drawLinear)

            svgContainer.addEventListener(up[platform], function onceLinear() {
                // 多人协作, 发送结束标记
                let msg = {
                    type: LineType,
                    Attr: {
                        id: lineNum,
                        isEnd: true,
                    },
                }
                client.send(JSON.stringify(msg))
                lineNum++
                svgContainer.removeEventListener(up[platform], onceLinear)
                svgContainer.removeEventListener(move[platform], drawLinear)
            })
        }
        // TODO: 文本
        if (textedit.checked) {
            var localpoint = mousePos(svg);
            if (!mousedownonelement) {
                createtext(localpoint, svg);
            } else {
                mousedownonelement = false;
            }
        }

    }
})
// 点击显示pagelist
pagelist.addEventListener('click', (e) => {
    //
    console.log("点击了pagelist", ispagelist)
    ispagelist = !ispagelist
    if (ispagelist)  {
        // 显示pagelist
        var element = document.getElementById("page-list-open")

        element.setAttribute("style","margin-top: 10px;display: block;width: " +
            "80px;background-color:rgb(94 170 255);border-radius: 10px;" +
            "border: solid 2px rgb(255, 255, 255);")
    } else {
        // 显示pagelist
        var element = document.getElementById("page-list-open")
        element.setAttribute("style","display: node;")
        console.log("隐藏pagelist")
        // // 实现hover效果
        // // 为element注册鼠标进入事件
        // element.onmouseover = function () {
        //     element.setAttribute("style","margin-top: 10px;display: block;width: " +
        //         "80px;background-color:rgb(94 170 255);border-radius: 10px;" +
        //         "border: solid 2px rgb(255, 255, 255);")
        // };
        // // 为element注册鼠标离开事件
        // element.onmouseout = function () {
        //     element.setAttribute("style","display: node;")
        // };
    }
})


// 清除所画内容
clear.addEventListener('click', (clickE) => {
    // 将此操作加入队列
    // 添加之前删除index 后面所有的elem
    elementQueue.splice(index+1, elementQueue.length-index-1)
    elementQueue.push({type: ClearType, value: svg.innerHTML})
    if (elementQueue.length > maxSize) {
        elementQueue.shift()
    } else {
        index ++
    }
    svg.innerHTML = ''
    // 多人协作
    let msg = {
        type: ClearType,
    }
    client.send(JSON.stringify(msg))
})
// 打开一个新文件
openFile.addEventListener('click', function openLoaclFile () {
    if (readonly && userName != pageAuthorName) {
        console.log("只读模式、禁止修改")
        return
    }
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
        LoadIds()
        let msg = {
            type: LoadType,
            Attr: {
                content: svgFileContent
            }
        }
        client.send(JSON.stringify(msg))
    })
    fr.readAsText(svgFile)
})

// 撤销与反撤销
document.addEventListener("keydown", function (event) {
    if (readonly && userName != pageAuthorName) {
        console.log("只读模式、禁止修改")
        return
    }
    if (event.code == 'KeyS' && event.shiftKey) {
        // 保存页面
        UploadSVG()
        showtips("Save")

    }else if (event.code == "KeyZ" && event.ctrlKey && event.shiftKey) {
        // console.log("redo, ", index, elementQueue.length)
        showtips("Redo")
        if (elementQueue.length > 0 && index < elementQueue.length-1) {
            console.log("redo#, ", index, elementQueue.length)
            index ++
            if (elementQueue[index].type == CommonType) {
                svg.append(elementQueue[index].value)
            } else if (elementQueue[index].type == ClearType) {
                svg.innerHTML = ''
            }
            let msg = {
                type: RedoType,
            }
            client.send(JSON.stringify(msg))
        }
    } else if (event.code == "KeyZ" && (event.ctrlKey || event.metaKey)) {
        // 撤销
        showtips("Undo")
        if (elementQueue.length > 0 && index >= 0) {
            if (elementQueue[index].type == CommonType) {
                svg.lastChild.remove()
            } else if (elementQueue[index].type == ClearType) {
                svg.innerHTML = elementQueue[index].value
            }
            index --
            let msg = {
                type: UndoType,
            }
            client.send(JSON.stringify(msg))
        }
    }
})

// 下面是退出浏览器提示
window.onbeforeunload = function (e) {
    if (drawandnosave && !readonly)  {
        // TODO: 持久化此页面
        e = e || window.event
        if(e) {
            e.returnValue = '关闭提示'
        }
        alert("页面还没有保存")
        console.log("点击了退出按钮")
    }
}

// 将图画保存到桌面
saveFile.addEventListener('click' ,function save(){
    // drawandnosave = false

    var svgSource = svg.outerHTML
    var blob = new Blob(['<?xml version="1.0" encoding="utf-8"?>', svgSource], { type: "image/xml+svg" })
    var url = URL.createObjectURL(blob)
    var anchor = document.createElement("a")
    anchor.href = url
    anchor.download = saveFileName
    anchor.click()  // 点击button 也触发a标签点击

    // 发起POST请求
    UploadSVG()
    // TODO：持久化此页面
})
addPage.addEventListener("click", function add(){
    // TODO: 添加新页面
    console.log("点击了add")
    content=prompt("请输入页面名称：");
    console.log(content);//4
    // 如果取消 直接返回
    if (content == null) {
        return
    }
    // 发起addpage请求
    AddPage(content)

})

// 撤销上一步
undo.addEventListener("click", function (e) {
    if (elementQueue.length > 0 && index >= 0) {
        if (elementQueue[index].type == CommonType) {
            svg.lastChild.remove()
        } else if (elementQueue[index].type == ClearType) {
            svg.innerHTML = elementQueue[index].value
        }
        index --
        let msg = {
            type: UndoType,
        }
        client.send(JSON.stringify(msg))
    }
})

// TODO redo
redo.addEventListener("click", function (e) {
    if (elementQueue.length > 0 && index < elementQueue.length-1) {
        index ++
        if (elementQueue[index].type == CommonType) {
            svg.append(elementQueue[index].value)
        } else if (elementQueue[index].type == ClearType) {
            svg.innerHTML = ''
        }
        let msg = {
            type: RedoType,
        }
        client.send(JSON.stringify(msg))
    }
})

// 生成链接
share.addEventListener("click", function (e){
    console.log("share click")
    let area = document.createElement("textarea")
    area.innerHTML = "http://" + HostAddr + "/qiniu?page="+pageName+"&author="+pageAuthorName
    document.body.appendChild(area)
    area.select()
    document.execCommand("Copy")
    document.body.removeChild(area)
    var r=confirm("分享链接已经复制到粘贴板上啦");
})

// TODO rwMode
rwMode.addEventListener("click", function (e){
    readonly = !readonly
    // 将按钮图标更换
    if (readonly) {
        document.getElementsByClassName('icon-qiniu-readonly')[0].setAttribute('data-before', 'R')
    } else {
        document.getElementsByClassName('icon-qiniu-readonly')[0].setAttribute('data-before', 'W')
    }
    let msg = {
        type: ModeChangeType,
        Attr: readonly,
    }
    client.send(JSON.stringify(msg))
})


// 鼠标相对于元素的位置
function mousePos(node) {
    var box = node.getBoundingClientRect()

    if (platform) {
        return {
            x: window.event.clientX - box.x,
            y: window.event.clientY - box.y,
        }
    } else {
        console.log("mousePos: ", window.event)
        return {
            x: window.event.changedTouches[0].pageX - box.x,
            y: window.event.changedTouches[0].pageY - box.y,
        }
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


// 判断是否是pc设备
function IsPC() {
    var userAgentInfo = navigator.userAgent;
    var Agents = ["Android", "iPhone","SymbianOS", "Windows Phone", "iPod"];
    var flag = true;
    console.log("Agent: ", userAgentInfo)
    for (var v = 0; v < Agents.length; v++) {
        if (userAgentInfo.indexOf(Agents[v]) > 0) {
            flag = false;
            break;
        }
    }
    console.log("ISPC()", flag)
    return flag;
}



// 更加当前svg的所有元素，设置各个图形的id起始
// 主要用于加载图形后使用
function LoadIds() {
    for (i = 0; i < svg.children.length; i++) {
        var tag = svg.children[i].nodeName
        var id = svg.children[i].getAttribute("id")
        console.log("LoadIds:", tag, id)

        if (tag == 'polyline') {
            polylineNum = max(polylineNum, parseInt(id))
        } else if (tag == 'ellipse') {

            ellipseNum = max(ellipseNum, id.split('_')[1])
        } else if (tag == 'rectangle') {
            rectangleNum = max(rectangleNum, id.split('_')[1])
        } else if (tag == 'line') {
            lineNum = max(lineNum, id.split('_')[1])
        }
    }
    polylineNum ++
    ellipseNum ++
    rectangleNum ++
    lineNum ++
}

function max(a, b) {
    if (a > b) {
        return a
    }
    return b
}