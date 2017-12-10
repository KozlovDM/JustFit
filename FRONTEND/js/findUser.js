$(document).ready(function(){ 
    var searhBlock = '<div class="foundUser"><span class="foundUser-notfound">Ничего не найдено</span></div>';
    
    $('.searchWindow').hide();
    
    $('input[class="find"]').on('click',function(){
        $("html,body").css("overflow","hidden");
        $('.searchWindow').show();
    });
    
    $('input[class="find"]').focusout(function(){
        $('input[class="find"]').val('');
        $('.searchWindow').hide();
        $('.searchWindow').html(searhBlock);
        $("html,body").css("overflow","auto");
    });
    
    $('input[class="find"]').on('input keyup',function(event){
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
                        $('.foundUser').html('');
                        var ref = '"css/images/avatar.jpg"';
                        for (var i = 1; i <= data.count; i++){
                            if (data["avatar" + i] !== null){
                                var base64 = data["avatar" + i]; 
                                ref = '"data:image/jpeg;base64,' + base64 + '"';
                            }
                            var block = '<div class="foundUser"><div class="foundUser-avatar"><img src=' + ref + '/></div><a href="#" class="foundUser-nickname">' + data["name" + i] + '</a></div>'; 
                            $('.foundUser').append(block);
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
    });
});