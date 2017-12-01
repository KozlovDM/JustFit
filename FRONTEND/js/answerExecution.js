$(document).ready(function(){
    $('#regForm').submit(function(){
        if($(this).data('formstatus') !== 'submitting'){
            var form = $(this),
                formData = form.serialize(),
                formUrl = form.attr('action'),
                formMethod = form.attr('method');
                
            form.data('formstatus','submitting');
            $.ajax({
                url: formUrl,
                type: formMethod,
                data: formData,
                success:function(){
                    window.location.href = "UserPage.html" + "?" + $('#regForm input[name="phone"]').val();
                },
                error:function(){
                    alert("Такой пользователь уже существует!");            
                }
            });
        }
        return false;
    });
    
    $('#enterForm').submit(function(){
        if($(this).data('formstatus') !== 'submitting'){
            var form = $(this),
                formData = form.serialize(),
                formUrl = form.attr('action'),
                formMethod = form.attr('method');
            
            form.data('formstatus','submitting');
            $.ajax({
                url: formUrl,
                type: formMethod,
                data: formData,
                success:function(){
                    window.location.href = "UserPage.html" + "?" + $('#enterForm input[name="phone"]').val();
                },
                error:function(){
                    alert("Неправильный номер или пароль!");
                }
            });
        }
        return false;
    });
});