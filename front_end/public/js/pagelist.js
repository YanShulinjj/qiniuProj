
// 通过js添加page部分
// 获取ul元素
var ul = document.getElementById("pagelist")
for (i =0; i<page.size; i++) {
    var li = document.createElement("li")
    a = document.createElement("a")
    a.href = "0.0.0.0:8080/qiniu?page="+page[i].page_name
    li.append(a)
    ul.append(li)
}