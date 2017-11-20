$(document).ready(function(){
    $('#regForm').submit(function(){
        if($(this).data('formstatus') !== 'submitting'){
            var form = $(this),
                formData = form.serialize(),
                formUrl = "http://192.168.56.1:3000/SignUp",
                formMethod = form.attr('method');
                
            form.data('formstatus','submitting');
            $.ajax({
                url: formUrl,
                type: formMethod,
                data: formData,
                success:function(data){
                    var responseData = jQuery.parseJSON(data);
                    alert(data);     
                }
            });
            window.location.href = form.attr('action');
        }
        return false;
    });
});