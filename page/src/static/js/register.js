$(document).ready(function() {

    $("#submit").click(function () {
        if ($("#password").val() != $("#retypePassword").val()) {
            alert("you typed different passwords");
            return
        }
        $.ajax({
            type:"POST",
            url:"http://localhost:12212/api/v1/registerreq",
            data: {
                "email": $("#email").val(),
                "username": $("#username").val(),
                "name": $("#name").val(),
                "password": $("#password").val()
            },
            statusCode: {
                400: function(data) {
                    alert(data.responseJSON.error);
                }
            },
            success:function (data) {
                window.location.href = "/";
            }
        });
    });

});