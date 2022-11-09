// 包含所有与后端请求交互的方法

let userId = 0
let HostAddr = document.getElementById("hostAddr").innerText
let pageAuthorName = document.getElementById("authorName").innerText
let userName = document.getElementById("userName").innerText
let pageName = document.getElementById("pageName").innerText
let pages  = []

// let DATA

GetUserInfo()


function AddPage(newPageName) {
    var url = "http://"+HostAddr+'/backend/page/add?username='+userName+'&pagename='+newPageName
    $.ajax({
        type: 'GET',
        url: url,
        processData: false,
        contentType: false
    }).done(function (data) {
        // 服务器返回的数据
        console.log(data)
        if (data.status_code == 0) {
            // 更新pagelist
            pages.push(data)
            // 通过js添加page部分
            console.log("Add page: ", pages)
            UpdatePageList(pages)
        } else {
            alert("添加页面出错，该pagename已经存在！")
        }
    })
}

function GetUserInfo() {
    var url = "http://"+HostAddr+'/backend/user/add?username='+userName
    $.ajax({
        type: 'GET',
        url: url,
        processData: false,
        contentType: false
    }).done(function (data) {
        // 服务器返回的数据
        console.log(data)
        userName = data.user_name
        userId = data.user_id
        console.log(userId)
        InitPage()
        ws ()   // 直接建立websocket连接
        GetPageList()
    })
}

function GetPageList() {
    var url = "http://"+HostAddr+'/backend/page/list?author=' + pageAuthorName
    $.ajax({
        type: 'GET',
        url: url,
        processData: false,
        contentType: false
    }).done(function (data) {
        // 服务器返回的数据
        console.log("PageList", data.list.length)
        pages = data.list
        // 通过js添加page部分
        UpdatePageList(pages)

    })
}

function UpdatePageList(pages) {
    // 获取ul元素
    var ul = document.getElementById("pagelist")
    // 首先清除ul的child
    ul.innerHTML = ''
    console.log("UpdatePageList: ", pages)
    for (i =0; i< pages.length; i++) {
        var li = document.createElement("li")
        a = document.createElement("a")
        a.href = "/qiniu?author="+pageAuthorName+"&page="+pages[i].page_name
        a.innerHTML = pages[i].page_name
        li.insertBefore(a, null)
        ul.insertBefore(li, ul.lastChild)
    }
}

function UploadSVG() {
    if (readonly && userName != pageAuthorName) {
        console.log("只读模式、禁止修改")
        return
    }
    drawandnosave = false
    var svgSource = svg.outerHTML
    var blob = new Blob(['<?xml version="1.0" encoding="utf-8"?>', svgSource], { type: "image/xml+svg" })

    var fd = new FormData();
    fd.append('username', userName)
    fd.append('pagename', pageName)
    fd.append('data', blob)
    $.ajax({
        type: 'POST',
        url: "http://"+HostAddr+'/backend/page/upload',
        data: fd,
        processData: false,
        contentType: false
    }).done(function (data) {
        // 服务器返回的数据
        console.log(data)
    })
}

function InitPage() {
    var url = "http://"+HostAddr+'/backend/page/get?username='+pageAuthorName+'&pagename='+pageName
    $.ajax({
        type: 'GET',
        url: url,
        processData: false,
        contentType: false
    }).done(function (data) {
        // 服务器返回的数据
        console.log(data)
        if (data.status_code == 0) {
            var svgpath = data.svg_path
            $.ajax({
                type:"GET",
                url: svgpath,
                success:function(resp, status){
                    if(status == "success"){
                        console.log(resp)
                        svgparent.innerHTML = resp.firstChild.outerHTML
                        svg = document.querySelector('.svg')
                        // 设置多个图形的ID
                        // DATA = data
                        LoadIds()
                    }
                },
                complete:function(){
                    console.log("Syncws Reading...")
                    syncws() //
                }
            })

        } else {
            alert("该页面还未创建！")
        }
    })
}