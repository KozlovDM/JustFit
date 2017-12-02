$(document).ready(function(){     
    $('.window').hide();    
    
    $('#userFile').click(function(){
        $("body").append("<div id='overlay'></div>");
        $("#overlay").height($(document).height());
        
        $("html,body").css("overflow","hidden");
        
        $("#scalePhoto").attr("src",$(this).attr('src'));
        $('.window').show();
        
        return false;
    });
    
    $('.close').click(function(){
        $('.window').hide();
        $("#overlay").remove();
        $("html,body").css("overflow","auto");
        
        return false;
    });
});