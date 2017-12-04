$(document).ready(function(){ 
    var phone = location.search.substring(1);
    $('.window').hide();    
    
    $("#collage").on("click","a",function(){
        $("body").append("<div id='overlay'></div>");
        $("#overlay").height($(document).height());
        
        $("html,body").css("overflow","hidden");
        
        $("#scalePhoto").attr("src",$(this).find('img').attr('src'));
        $('.window').show();
        
        var nameimage = $(this).find('img').attr('alt');
        $('#send').on('click', function(event){ 
            event.preventDefault(); 
            var comment = $('input[class="comment"]').val(); 
            if (comment !== null){ 
                var data = new FormData(); 
                data.append("nameimage", nameimage); 
                data.append("phone", phone);
                data.append("comment", comment);

                $.ajax({ 
                    type: 'POST', 
                    url: 'http://127.0.0.1:3000/Comment', 
                    contentType: false, 
                    processData: false, 
                    data: data, 
                    success: function(data){ 
                        var block = '<div class="window-comments__all__pack"><div class="window-comments__all__pack__nickname">' + data.login + '</div><div class="window-comments__all__pack__message">' + comment + '</div></div>'; 
                        $('.window-comments__all').append(block);
                    }, 
                    error:function(status, errorMsg){
                        alert("Статус: " + status + " Ошибка: " + errorMsg);
                    } 
                }); 
            } 
        });
        
        return false;
    });
    
    $('.close').click(function(){
        $('.window').hide();
        $('.window-comments__all').html('');
        $('.heart').find('img').attr("src","css/images/heart.png");
        $("#overlay").remove();
        $("html,body").css("overflow","auto");
        
        return false;
    });
    
    /*$('.heart').click(function(){
        $(this).find('img').attr("src", "css/images/redheart.png");
    });*/
});