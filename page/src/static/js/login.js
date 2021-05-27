$(document).ready(function() {

    $("#submit").click(function () {
        obj = {
            "login":$("#login").val(),
            "password":$("#password").val()
        }
        $ajax({
            method:"POST",
            url:"http://localhost:2387/api/v1/login",
            data:obj,
            success:function (data) {
                alert("logged in" + data.token)
            }
        })
    });

});