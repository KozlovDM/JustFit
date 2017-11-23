$(document).ready(function(){
    $.ajax({
        url: 'http://192.168.56.1:3000/GetUserData',
        type: 'POST',
        data: window.Storage.phone,
        success:function(data){
            var userData = jQuery.parseJSON(data);
            
            $('.sec1__nickname').text(data.login);
            $('.publication_amount').text('Публикации: ' + data.publications);
            $('.subscriber_amount').text('Подписчики: ' + data.subscribers);
            $('.subscription_amount').text('Подписки: ' + data.subscriptions);
            $('.about__fullname').text(data.fullname);
            if (data.info != null){
                $('.about_me').text(data.info);
            }    
            
            $.ajax({
                url: 'http://192.168.56.1:3000/GetUserAvatar',
                type: 'POST',
                data: window.Storage.phone,
                success:function(data){
                    /*var userAvatar = ;*/  
                    
                    
                },
                error:function(status, errorMsg){
                    alert("Статус: " + status " Ошибка: " + errorMsg);
                }
            });
        },
        error:function(status, errorMsg){
            alert("Статус: " + status " Ошибка: " + errorMsg);
        }
    });
});