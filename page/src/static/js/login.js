$(document).ready(function() {

    $("#submit").click(function () {
        obj = {
            "login":$("#login").val(),
            "password":$("#password").val()
        }
        $ajax({
            method:"POST",
            url:"http://localhost:12212/api/v1/loginreq",
            data:obj,
            success:function (data) {
                alert("logged in" + data.token)
            }
        })
    });

});