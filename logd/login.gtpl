<html>
<head>
<title></title>
</head>
<body>

<script type="text/javascript">
    function submitFunc(act) {
        loginForm.action=act;
        loginForm.submit();
    }
</script>

<form method="post" name="loginForm">
    User:<input type="text" name="username">
    Pass:<input type="password" name="password">
    <BR/>
    <input type="button" value="Login" onclick="submitFunc('/login')">
    <input type="button" value="Register" onclick="submitFunc('/register')">

</form>


</body>
</html>
