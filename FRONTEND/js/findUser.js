$(document).ready(function(){ 
    var searhBlock = '<div class="foundUser"><span class="foundUser-notfound">Ничего не найдено</span></div>';
    
    $('.searchWindow').hide();
    
    $('input[class="find"]').on('click',function(){
        $('.searchWindow').show();
    });
    
    $(document).mouseup(function(e){
        var div = $('.searchWindow');
        if (!div.is(e.target) && div.has(e.target).length === 0){
            div.hide();
            $('input[class="find"]').val('');
            $('.searchWindow').html(searhBlock);
        }
    });
    
    $('input[class="find"]').on('input keyup',function(event){
        $('.searchWindow').show();
        event.preventDefault(); 
        var inputData = $('input[class="find"]').val(); 
        if (inputData !== null){ 
            var data = new FormData(); 
            data.append("login", inputData); 

            $.ajax({ 
                type: 'POST', 
                url: 'http://127.0.0.1:3000/Search', 
                contentType: false, 
                processData: false, 
                data: data, 
                success: function(data){ 
                    if (data.count !== 0){
                        $('.searchWindow').html('');
                        var ref = '"css/images/avatar.jpg"';
                        for (var i = 1; i <= data.count; i++){
                            if (data["user" + i].avatar !== null){
                                var base64 = data["user" + i].avatar; 
                                ref = '"data:image/jpeg;base64,' + base64 + '"';
                            }
                            var login = data["user" + i].login;
                            var block = '<div class="foundUser"><div class="foundUser-avatar"><img src=' + ref + '/></div><a href="#" class="foundUser-nickname" alt="' + login + '">' + login + '</a></div>'; 
                            $('.searchWindow').append(block);
                        }
                    }
                    else{
                        $('.searchWindow').html(searhBlock);    
                    }
                }, 
                error:function(){
                    alert("Некорректные данные!");
                } 
            }); 
        } 
        else{
            $('.searchWindow').html(searhBlock);
        }
    });
    
    $(".searchWindow").on("click","a",function(){                
        var login = $(this).attr('alt');
        
        $.ajax({ 
            type: 'POST', 
            url: 'http://127.0.0.1:3000/GetUserData',  
            data: {login: login}, 
            success: function(data){
                $('.searchWindow').hide();
                $('.sec1__nickname').text(data.login);
                $('.publication_amount').text('Публикации: ' + data.publications);
                $('.subscriber_amount').text('Подписчики: ' + data.subscribers);
                $('.subscription_amount').text('Подписки: ' + data.subscriptions);
                $('.about__fullname').text(data.fullname);
                if (data.info != null){
                    $('.about_me').text(data.info);
                }    
                var ref = 'css/images/avatar.jpg';
                if (data.avatar !== null){
                    ref = "data:image/jpeg;base64," + data.avatar; 
                }
                $('#avatar').attr("src", ref);
                $('#collage').html('');
            if (data.publications !== 0){
                var base64, ref, alt, block;
                for (var i = 1; i <= data.publications; i++){
                    base64 = data["file" + i].file; 
                    ref = '"data:image/jpeg;base64,' + base64 + '"'; 
                    alt = '"' + data["file" + i].nameimage + '" ';
                    block = '<div class="publication"><a href="#"><img id="userFile" alt=' + alt + 'src=' + ref + '></a></div>'; 
                    $('#collage').append(block);
                }
            }
            if ($('#avatar').attr("alt") !== login){
                $('.sec1__edit').hide();
                $('.main-profile__info__sec4').hide();
            }
            else{
                $('.sec1__edit').show();
                $('.main-profile__info__sec4').show();
            }
            $('input[class="find"]').val('');
        }, 
            error:function(status, errorMsg){
                alert("Статус: " + status + " Ошибка: " + errorMsg);
            } 
        });
    });
});