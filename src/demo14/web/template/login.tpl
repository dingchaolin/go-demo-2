<html>
<body>
<form method="post" action="checklogin">
    <p>用户名:<input type="text" name="user"/></p>
    <p>密码:<input type="password" name="password"/></p>
    <p><input type="submit"/></p>

</form>
{{ if . }}
<p style="color:red">{{.}}</p>
{{end}}
</body>
</html>