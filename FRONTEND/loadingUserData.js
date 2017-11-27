$(document).ready(function(){
    var phone = location.search.substring(1);
    $.ajax({
        url: 'http://192.168.56.1:3000/GetUserData',
        type: 'POST',
        data: {phone: phone},
        success:function(data){
            $('.sec1__nickname').text(data.login);
            $('.publication_amount').text('Публикации: ' + data.publications);
            $('.subscriber_amount').text('Подписчики: ' + data.subscribers);
            $('.subscription_amount').text('Подписки: ' + data.subscriptions);
            $('.about__fullname').text(data.fullname);
            if (data.info != null){
                $('.about_me').text(data.info);
            }    
            
            /*$.ajax({
                url: 'http://192.168.56.1:3000/GetUserAvatar',
                type: 'POST',
                data: phone,
                success:function(data){
                    var userAvatar = ; 
                       
                },
                error:function(status, errorMsg){
                    alert("Статус: " + status + " Ошибка: " + errorMsg);
                }
            });*/
        },
        error:function(status, errorMsg){
            alert("Статус: " + status + " Ошибка: " + errorMsg);
        }
    });
    
    $('#clickbutton').on('click', function(event){
        event.preventDefault();
        var file = document.getElementById('uploadFile').files;
        if (files.length !== 0){
            if (window.FormData !== undefined){
                var data = new FormData();
                data.append("file", files[0]);
                data.append("phone", phone);
                
                $.ajax({
                    type: 'POST',
                    url: 'http://192.168.56.1:3000/Upload',
                    contentType: false,
                    processData: false,
                    data: data,
                    success: function(data){
                        var ref = '"' + data + '"';
                        var block = '<div class="publication"><a href=' + ref + ' target="_blank"><img src=' + ref + '></a></div>';
                        $('.main-collage').append(block);
                    },
                    error: function(){
                        alert("Файл не может быть загружен!");
                    }
                });
            }
        }
    });       
});