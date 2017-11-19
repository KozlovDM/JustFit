$(document).ready(function(){
    $("#regForm").submit(function(event) {
        var form = $(this),
            formUrl = form.attr('action')
        $.post(formUrl, function(data) {
            alert(data);
        });
    });
});