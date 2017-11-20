$(document).ready(function(){
    $('#regForm').submit(function(){
        if($(this).data('formstatus') !== 'submitting'){
            var form = $(this),
                formData = form.serialize(),
                formUrl = "http://192.168.56.1:3000/SignUp",
                formMethod = form.attr('method'),
                code = 0;
                
                
            form.data('formstatus','submitting');
            $.ajax({
                url: formUrl,
                type: formMethod,
                data: formData,
                success:function(data){
                    var responseData = jQuery.parseJSON(data);  
                    code = data.statusCode;
                }
            });
            
            switch(code){
                case 200:
                    window.location.href = form.attr('action');
                    break;
                default:
                    alert("Такой пользователь уже зарегистрирован!");
                    break;
        }
        return false;
    });
});