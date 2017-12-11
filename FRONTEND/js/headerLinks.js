$(document).ready(function(){
    var phone = location.search.substring(1);
    
    $('.header-menu__link:first').click(function(){
        window.location.href = "UserPage.html" + "?" + phone;
    });
});