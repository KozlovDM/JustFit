$(document).ready(function(){
    $('#regForm').submit(function(){
        if($(this).data('formstatus') !== 'submitting'){
            var form = $(this),
                formData = form.serialize(),
                formUrl = form.attr('action'),
                formMethod = "http://192.168.56.1:3000/SignUp";
                
            form.data('formstatus','submitting');
            $.ajax({
                url: formUrl,
                type: formMethod,
                data: formData,
                statusCode: {
                    200: function() {                    
                        window.location.href = form.attr('action');   
                    },
                    404: function(status, errorMsg) {
					    alert("Статус: " + status + " Ошибка: " + errorMsg);
                    }
                }
            });
        }
        return false;
    });
});