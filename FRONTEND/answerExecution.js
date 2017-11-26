$(document).ready(function(){
    window.Storage = {};
    $('#regForm').submit(function(){
        if($(this).data('formstatus') !== 'submitting'){
            var form = $(this),
                formData = form.serialize(),
                formUrl = form.attr('action'),
                formMethod = form.attr('method');
            
            localStorage.setItem("phone", $('input[name="phone"]').val());
                
            form.data('formstatus','submitting');
            $.ajax({
                url: formUrl,
                type: formMethod,
                data: formData,
                success:function(){
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
            
            localStorage.setItem("phone", $('input[name="phone"]').val());
            
            form.data('formstatus','submitting');
            $.ajax({
                url: formUrl,
                type: formMethod,
                data: formData,
                success:function(){
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