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
                success:function(data){  
                    window.location.href = "UserPage.html";
                },
                error:function(jqxhr, status, errorMsg){
                    alert("Status: " + status + " Error: " + errorMsg);
                }
            });
        }
        return false;
    });
});