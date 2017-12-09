$(document).ready(function(){ 
    var windowBlock = '<div class="window-publication"><div class="window-publication__file"><img id="scalePhoto" src="css/images/avatar.jpg"/></div><div class="window-publication__likes"><div class="window-publication__likes__amount">Нравится: 0</div><div class="window-publication__likes__picture"><a class="heart" href="#"><img src="css/images/heart.png"/></a></div></div></div><div class="window-comments"><div class="window-comments__head"><div class="window-comments__head__title">Комментарии</div><div class="window-comments__head__close"><a class="close" href="#"><img src="css/images/close2.png"/></a></div></div><div class="window-comments__all"></div><div class="window-comments__send"><input class="comment" id="uploadComment" type="text" placeholder="Оставить комментарий..." maxLength=100/><button id="send">Отправить</button></div></div>';
    
    var phone = location.search.substring(1);
    var like;
    $('.window').hide();    
    
    $("#collage").on("click","a",function(){
        $("body").append("<div id='overlay'></div>");
        $("#overlay").height($(document).height());
        
        $("html,body").css("overflow","hidden");
        
        $("#scalePhoto").attr("src",$(this).find('img').attr('src'));
        
        var nameimage = $(this).find('img').attr('alt');
        
        $.ajax({ 
            type: 'POST', 
            url: 'http://10.254.4.178:3000/ImageInfo',  
            data: {nameimage: nameimage, phone: phone}, 
            success: function(data){
                like = data.islike;
                if (data.count !== 0){
                    for (var i = 1; i <= data.count; i++){
                        var block = '<div class="window-comments__all__pack"><div class="window-comments__all__pack__nickname">' + data.comment["user" + i] + '</div><div class="window-comments__all__pack__message">' + data.comment["comment" + i] + '</div></div>'; 
                        $('.window-comments__all').append(block);
                    }
                }
                if (like){
                    $(".heart").find('img').attr("src", "css/images/redheart.png");
                }
                $('.window-publication__likes__amount').text('Нравится: ' + data.like);
                $('.window').show();
            }, 
            error:function(status, errorMsg){
                alert("Статус: " + status + " Ошибка: " + errorMsg);
            } 
        });
        
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
                    url: 'http://10.254.4.178:3000/Comment', 
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
        
        $('.heart').click(function(){
            if (like){
                $(this).find('img').attr("src", "css/images/heart.png");
                like = false;
            }
            else{
                $(this).find('img').attr("src", "css/images/redheart.png");
                like = true;
            }
            
            $.ajax({ 
                type: 'POST', 
                url: 'http://10.254.4.178:3000/Like',  
                data: {nameimage: nameimage, phone: phone}, 
                success: function(data){ 
                    $('.window-publication__likes__amount').text('Нравится: ' + data.like);
                }, 
                error:function(status, errorMsg){
                    alert("Статус: " + status + " Ошибка: " + errorMsg);
                } 
            });
        });
        
        $('.close').click(function(){
            $('.window').hide();
            $('.window').html(windowBlock);
            $("#overlay").remove();
            $("html,body").css("overflow","auto");
        });
        
        return false;
    });
});