$(document).ready(function(){ 
    $('.window').hide();    
    
    $("#collage").on("click","a",function(){
        $("body").append("<div id='overlay'></div>");
        $("#overlay").height($(document).height());
        
        $("html,body").css("overflow","hidden");
        
        $("#scalePhoto").attr("src",$(this).find('img').attr('src'));
        $('.window').show();
        
        return false;
    });
    
    $('.close').click(function(){
        $('.heart').find('img').attr("src", "css/images/heart.png");
        $('.window').hide();
        $("#overlay").remove();
        $("html,body").css("overflow","auto");
        
        return false;
    });
    
    /*$('.heart').click(function(){
        $(this).find('img').attr("src", "css/images/redheart.png");
    });*/
});