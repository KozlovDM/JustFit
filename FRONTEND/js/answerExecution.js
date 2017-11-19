$(document).ready(function(){
    $("#regForm").submit(function(event) {
        var form = $(this),
            formData = form.serialize(),
            formUrl = form.attr('action'),
            formMethod = form.attr('method')
        
        $.ajax({
            url: formUrl,
            type: formMethod,
            data: formData,
            success:function(data){
                var responseData = jQuery.parseJSON(data),
                    }
        });
        alert(responseData);
    });
});