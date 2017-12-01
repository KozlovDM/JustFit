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
            if (data.avatar !== null){
                var ref = "data:image/jpeg;base64," + data.avatar; 
                $('#avatar').attr("src", ref);
            }
            if (data.publications !== 0){
                var base64, ref, block;
                for (var i = 1; i <= data.publications; i++){
                    base64 = data["file" + i]; 
                    ref = '"data:image/jpeg;base64,' + base64 + '"'; 
                    block = '<div class="publication"><a href="#"><img src=' + ref + '></a></div>'; 
                    $('.main-collage').append(block);
                }
            }
        },
        error:function(status, errorMsg){
            alert("Статус: " + status + " Ошибка: " + errorMsg);
        }
    });
    
    $('#clickbutton').on('click', function(event){ 
        event.preventDefault(); 
        var file = document.getElementById('uploadFile').files; 
        if (file.length !== 0){ 
            if (window.FormData !== undefined){ 
                var data = new FormData(); 
                data.append("file", file[0]); 
                data.append("phone", phone); 

                $.ajax({ 
                    type: 'POST', 
                    url: 'http://192.168.56.1:3000/Upload', 
                    contentType: false, 
                    processData: false, 
                    data: data, 
                    success: function(data){ 
                        var base64 = data.image; 
                        var ref = '"data:image/jpeg;base64,' + base64 + '"'; 
                        var block = '<div class="publication"><a href="#"><img src=' + ref + '></a></div>'; 
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