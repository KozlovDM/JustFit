$(document).ready(function(){
    $("#regForm").submit(function(event) {
        if($(this).data('formstatus') !== 'submitting'){
            var form = $(this),
                formData = form.serialize(),
                formUrl = form.attr('action'),
                formMethod = form.attr('method'),
                responseMsg = $('#signup-response');
        
            $.ajax({
                url: formUrl,
                type: formMethod,
                data: formData,
                success:function(data){
                    var responseData = jQuery.parseJSON(data),
                        klass = 'message';
                }
                
                responseMsg.fadeOut(200,function(){
                    $(this).addClass(klass)
                    .text(responseData.message)
                    .fadeIn(200,function());
                });
            });
        }
        return false;
    });
});