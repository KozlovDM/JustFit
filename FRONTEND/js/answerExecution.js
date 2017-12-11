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
                    localStorage.setItem('phone', $('#regForm input[name="phone"]').val());
                    window.location.href = "UserPage.html";
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
                    localStorage.setItem('phone', $('#enterForm input[name="phone"]').val());
                    window.location.href = "UserPage.html";
                },
                error:function(){
                    alert("Неправильный номер или пароль!");
                }
            });
        }
        return false;
    });
});