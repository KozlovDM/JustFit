$(document).ready(function(){
    $.ajax({
        url: 'http://192.168.56.1:3000/GetUserData',
        type: 'POST',
        data: localStorage.getItem("phone"),
        success:function(data){
            var userData = jQuery.parseJSON(data);
            $('.sec1__nickname').text(userData.login);
            $('.publication_amount').text('Публикации: ' + userData.publications);
            $('.subscriber_amount').text('Подписчики: ' + userData.subscribers);
            $('.subscription_amount').text('Подписки: ' + userData.subscriptions);
            $('.about__fullname').text(userData.fullname);
            if (data.info != null){
                $('.about_me').text(userData.info);
            }    
            
            /*$.ajax({
                url: 'http://192.168.56.1:3000/GetUserAvatar',
                type: 'POST',
                data: localStorage.getItem("phone"),
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
});