$(document).ready(function(){
    var phone = localStorage.getItem('phone');
    
    $('.main-profile__edit').hide();
    
    $('.sec1__edit').click(function(){
        $.ajax({
            url: 'http://127.0.0.1:3000/GetUserData',
            type: 'POST',
            data: {phone: phone},
            success:function(data){
                $('input[id="editNickname"]').val(data.login);
                $('input[id="editFullname"]').val(data.fullname);
                $('.main-profile__edit').show();
            },
            error:function(status, errorMsg){
                alert("Статус: " + status + " Ошибка: " + errorMsg);
            }
        });
    });
    
    $('#sendEditData').click(function(event){
        event.preventDefault(); 
        var login = $('input[id="editNickname"]').val();
        var fullname =$('input[id="editFullname"]').val();
        var avatar = null;
        if (document.getElementById('editAvatar').files.length !== 0) {
            avatar = document.getElementById('editAvatar').files[0];
        }
        var data = new FormData(); 
        data.append("phone", phone);
        data.append("login", login);
        data.append("fullname", fullname);
        data.append("avatar", avatar);

        $.ajax({ 
            type: 'POST', 
            url: 'http://127.0.0.1:3000/UpdateInfo', 
            contentType: false, 
            processData: false, 
            data: data, 
            success: function(){ 
                $('.sec1__nickname').text(login);
                $('.about__fullname').text(fullname);   
                if (avatar !== null){
                    var ref = "data:image/jpeg;base64," + avatar; 
                    $('#avatar').attr("src", ref);
                } 
                $('.main-profile__edit').hide();
            }, 
            error:function(){
                alert("Статус: " + status + " Ошибка: " + errorMsg);
            } 
        }); 
    });
    
    $('#closeEditData').click(function(){
        $('.main-profile__edit').hide();
        $('input[id="editNickname"]').val('');
        $('input[id="editFullname"]').val('');
        $('input[id="editAvatar"]').val('');
    });
});