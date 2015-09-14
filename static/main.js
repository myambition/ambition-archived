$('.list-group-item').on({
        'mousedown' : function() { activate(this); },
        'touchstart' : function() { activate(this); },
        'mouseup' : function() { deactivate(this); },
        'touchend' : function() { deactivate(this); }
});

function activate(that) {
    $('#main').find('.active').removeClass('active');
    $(that).addClass('active');
}

function deactivate(that) {
    $(that).removeClass('active');
}
