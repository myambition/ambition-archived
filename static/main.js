//var serverRoot = "http://localhost:3000"

$('.list-group-item').on({
        'mousedown' : function() { activate(this); },
        'touchstart' : function() { activate(this); },
        'mouseup' : function() { deactivate(this); },
        'touchend' : function() { deactivate(this); }
});

function activate(that) {
    $('#main').find('.active').removeClass('active');
    $(that).addClass('active');
    var data = { time: new Date().toISOString() };
    $.ajax({
        type: "POST",
        url: "/actions/" + $(that).attr("dbid"),
        data: JSON.stringify(data),
        sucess: function () { alert("woo");}
    })
}

function deactivate(that) {
    $(that).removeClass('active');
}
