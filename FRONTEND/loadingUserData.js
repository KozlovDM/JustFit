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
    
    $('#uploadForm').on('submit', function(event){
        if ($(this).data('formstatus') !== 'submitting' &&  window.FormData !== undefined) {
            event.preventDefault();
            var form = $(this),
                formData = new FormData(form.get(0));
            formData.append("phone", phone);
            
            $(this).data('formstatus','submitting');
            $.ajax({
                url: 'http://192.168.56.1:3000/download',
                type: form.attr('method'),
                contentType: false,
                processData: false,
                data: formData,
                dataType: 'json',
                success: function(data){
                    var ref = '"' + data[0] + '"';
                    var block = '<div class="publication"><a href=' + ref + ' target="_blank"><img src=' + ref + '></a></div>';
                    $('.header-menu__link:eq(1)').click(function(){
                        $('.main-collage').append(block);
                    });
                },
                error: function() {
                    alert("Файл не может быть загружен!");
                }
            });
        }
        return false;
    });
});