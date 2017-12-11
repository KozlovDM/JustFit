$(document).ready(function(){
    var phone = localStorage.getItem('phone');
    $.ajax({
        url: 'http://127.0.0.1:3000/GetUserData',
        type: 'POST',
        data: {phone: phone},
        success:function(data){
            $('.sec1__nickname').text(data.login);
            $('#avatar').attr("alt", data.login);
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
                var base64, ref, alt, block;
                for (var i = 1; i <= data.publications; i++){
                    base64 = data["file" + i].file; 
                    ref = '"data:image/jpeg;base64,' + base64 + '"'; 
                    alt = '"' + data["file" + i].nameimage + '" ';
                    block = '<div class="publication"><a href="#"><img id="userFile" alt=' + alt + 'src=' + ref + '></a></div>'; 
                    $('#collage').append(block);
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
                    url: 'http://127.0.0.1:3000/Upload', 
                    contentType: false, 
                    processData: false, 
                    data: data, 
                    success: function(data){ 
                        var base64 = data.file; 
                        var ref = '"data:image/jpeg;base64,' + base64 + '"'; 
                        var alt = '"' + data.nameimage + '" ';
                        var block = '<div class="publication"><a href="#"><img id="userFile" alt=' + alt + 'src=' + ref + '></a></div>'; 
                        $('#collage').append(block);
                        $('.publication_amount').text('Публикации: ' + data.publications);
                    }, 
                    error: function(){ 
                        alert("Файл не может быть загружен!"); 
                    } 
                }); 
            } 
        } 
        $('input[name="picture"]').val('');
    });       
});