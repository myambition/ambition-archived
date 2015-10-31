//var serverRoot = "http://localhost:3000"

$('.action').fastClick(
    function (that) {
        that = that.currentTarget;
        $(that).addClass('active');
        test = that;
        var data = { time: new Date().toISOString() };
        $.ajax({
            type: "POST",
            url: "/actions/" + $(that).attr("dbid"),
            data: JSON.stringify(data),
            success: function () { console.log("woo");},
            error: function(e) { console.log("noooooo: ", e);}
        })
        setTimeout(function(){
            $(".action").removeClass("active")
        }, 120)

    }
);
