$(document).ready(function(){
    $('.form-enter').hide();  
    $('.form-reg').hide();
    $('.header-about__text').hide();
    
    $('.header-menu__block:last').click(function(){
        $('.form-reg').hide();
        $('.header-title').hide();
        $('.header-about__text').hide();
        $('.form-enter').show();
    });
    
     $('.header-menu__block:eq(1)').click(function(){
        $('.form-reg').hide();
        $('.form-enter').hide();
        $('.header-title').hide();
        $('.header-about__text').show();
    });
    
    $('.header-menu__block:first').click(function(){
        $('.form-reg').hide();
        $('.form-enter').hide();
        $('.header-about__text').hide();
        $('.header-title').show();
    });
    
    $('.form-enter__close').mousemove(function(){
        $('.form-enter__close').css('border-style','inset');
    });
    $('.form-enter__close').mouseleave(function(){
        $('.form-enter__close').css('border-style','inherit');
    });
    $('.form-enter__close').click(function(){
        $('.form-enter').hide();
        $('.header-title').show();
    });
    
    $('.form-enter__regTitle').mousemove(function(){
        $('.form-enter__regTitle').css('color','red');
    });
    $('.form-enter__regTitle').mouseleave(function(){
        $('.form-enter__regTitle').css('color','blue');
    });  
    
    $('.form-enter__regTitle').click(function(){
        $('.form-enter').hide();
        $('.form-reg').show();
    });
    $('.form-reg__close').click(function(){
        $('.form-reg').hide();
        $('.header-title').show();
    });
    
    $('.form-reg__close').mousemove(function(){
        $('.form-reg__close').css('border-style','inset');
    });
    $('.form-reg__close').mouseleave(function(){
        $('.form-reg__close').css('border-style','inherit');
    });
    
    $('.form-reg__enterTitle').mousemove(function(){
        $('.form-reg__enterTitle').css('color','red');
    });
    $('.form-reg__enterTitle').mouseleave(function(){
        $('.form-reg__enterTitle').css('color','blue');
    });  
    
    $('.form-reg__enterTitle').click(function(){
        $('.form-reg').hide();
        $('.form-enter').show();
    });
});