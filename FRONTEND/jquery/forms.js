$(document).ready(function(){
    $('.form-enter').hide();  
    $('.form-reg').hide();
    
    $('.header-menu__block:last').click(function(){
        $('.form-reg').hide();
        $('.form-enter').toggle();
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
});