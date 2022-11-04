


// 将所有按钮隐藏
hidden.addEventListener('click', (e) => {
    //
    console.log("点击了隐藏")
    // 如果是只读模式，
    if (readonly && pageAuthorName != userName) {
        return
    }
    ishidden = !ishidden
    if (ishidden)  {
        hiddenElems()
    } else {
        displayElems()
    }

})

function hiddenElems() {
    // 左右侧隐藏
    document.getElementsByClassName("toolContainer")[0].style.display="none"
    var elements = document.getElementsByClassName("hiddenbutton")
    var i
    for (i = 0; i < elements.length; i++) {
        elements[i].style.display = "none";
    }
    // 将顶部按钮宽度更新
    document.getElementsByClassName("addundobar")[0].style.width = "80px"
}

function displayElems() {
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


function showtips(text) {
    // 先将tips中的文本替换

    tips.innerHTML = text
    tips.style.display = "block"

    $(".tips").fadeOut(300);
}