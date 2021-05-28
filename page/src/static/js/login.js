$(document).ready(function() {

    $("#submit").click(function () {
        $.ajax({
            type:"POST",
            url:"http://localhost:12212/api/v1/loginreq",
            data: {
                "login":$("#login").val(),
                "password":$("#password").val()
            },
            success:function (data) {
                alert("logged in " + data.token.value)
            }
        });
    });

});