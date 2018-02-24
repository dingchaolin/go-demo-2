<html>
<body>
<table border="1" cellpadding="0">
    <tr>
        <td>ID<td/>
        <td>Name<td/>
        <td>Note<td/>
        <td>删除<td/>
        <td>更新<td/>
    </tr>
    <!--
    . 代表当前元素 会随上下文改变而改变
    -->
    {{range .}}
    <tr>
        <td>{{.Id}}<td/>
        <td>{{.Name}}<td/>
        <td>{{.Note}}<td/>
        <td><a href="/delete?id={{.Id}}">删除</a><td/>
        <td><a href="/update?id={{.Id}}">更新</a><td/>
    </tr>
    {{end}}
</table>
<a href="/add">新增</a>
</body>
</html>