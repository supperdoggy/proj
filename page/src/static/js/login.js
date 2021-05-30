$(document).ready(function() {

    $("#submit").click(function () {
        $.ajax({
            type:"POST",
            url:"http://localhost:12212/api/v1/loginreq",
            data: {
                "login": $("#login").val(),
                "password": $("#password").val()
            },
            statusCode: {
                400: function(data) {
                    alert(data.responseJSON.error);
                    console.log(data)
                }
            },
            success:function (data) {
                window.location.href = "/";
            }
        });
    });

});