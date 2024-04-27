
function getOptions() {
    var position = $('#notification-position').val();
    var closeTimeout = $('#close-timeout').val();
    var animation = $('#animation').val();
    var showButtons = $('#show-buttons').get(0).checked;
    var showProgressBar = $('#show-progress-bar').get(0).checked;
    var animationOptions = {
        open: animation + '-in',
        close: animation + '-out'
    };

    if (animation === 'none') {
        animationOptions = {
            open: false,
            close: false
        };
    }

    return options = {
        description: 'I am a success notification',
        position: position,
        closeTimeout: closeTimeout,
        closeWith: ['click'],
        animation: animationOptions,
        showButtons: showButtons,
        buttons: {
            action: {
                callback: function(notification) {
                    console.log('action button');
                }
            }
        },
        showProgress: showProgressBar
    };
}
//登入

function checkname() {
    // var password=$("#exampleInputPassword").val() //得到前端输入的密码
    var username = $("#exampleInputEmail").val()//得到前端输入的账户

    if (username == undefined) {
        var options = getOptions();
        options.title = '信息';
        options.description = '请填写用户名！';
        options.width = 300;
        options.type = 'error';
        options.closeTimeout = 1000;
        GrowlNotification.notify(options);

        return false;
    } else if (username != undefined && (username.length > 11 || username.length < 6)) {
        var options = getOptions();
        options.title = '信息';
        options.description = '用户名长度应小于11位，大于6位！';
        options.width = 300;
        options.type = 'error';
        options.closeTimeout = 1000;
        GrowlNotification.notify(options);

        return false;
    }
}

function checkpwd() {
    var password = $("#exampleInputPassword").val() //得到前端输入的密码
    // var username =$("#exampleInputEmail").val()//得到前端输入的账户

    if (password == undefined) {
        var options = getOptions();
        options.title = '信息';
        options.description = '请填写密码！';
        options.width = 300;
        options.type = 'error';
        options.closeTimeout = 1000;
        GrowlNotification.notify(options);

        return false;
    } else if (password != undefined && (password.length > 11 || password.length < 6)) {
        var options = getOptions();
        options.title = '信息';
        options.description = '密码长度应小于11位，大于6位！！';
        options.width = 300;
        options.type = 'error';
        options.closeTimeout = 1000;
        GrowlNotification.notify(options);

        return false;
    }
}
var adminAddress
function check() {


    var password = $("#exampleInputPassword").val() //得到前端输入的密码
    var username = $("#exampleInputEmail").val()//得到前端输入的账户
    // alert("登录名:"+username)
    $.ajax({//通过ajax触发post请求
        url: " http://localhost:8080/dai/login",//post请求路径
        type: "POST",
        data: {username: username, password: password}, //对应相对数据
        dataType:"json",
        success: function (data) {
            console.log(data)
            // alert(data.data)
            if (data.data == 1) {
                //保存登录的用户名
                // alert("登录名:"+adminAddress)
                $.ajax({//通过ajax触发post请求
                    url: " http://localhost:8080/dai/saveAddress",//post请求路径
                    type: "POST",
                    data: {adminAddress: username}, //对应相对数据
                    dataType: "json",
                    success: function (data) {

                        // alert("登录名保存成功:"+data.data)

                    }
                })
                if (username == "0x5Ca22f6C1162Da1db405f4AFBD670A50c15c093a" && password == "1") {
                    var options = getOptions();
                    options.title = '信息';
                    options.description = '成功登入后台';
                    options.width = 300;
                    options.type = 'success';
                    options.closeTimeout = 1000;
                    GrowlNotification.notify(options);


                    window.location.href = "http://localhost:8080/table" //登入成功就切换到管理页面
                } else {

                    var options = getOptions();
                    options.title = '信息';
                    options.description = '成功登入个人主页';
                    options.width = 300;
                    options.type = 'success';
                    options.closeTimeout = 1000;
                    GrowlNotification.notify(options);

                    window.location.href = "http://localhost:8080/index" //登入成功就切换到主页面


                }
            }else {
                    var options = getOptions();
                    options.title = '信息';
                    options.description = '密码错误';
                    options.width = 300;
                    options.type = 'error';
                    options.closeTimeout = 1000;
                    GrowlNotification.notify(options);
                    alert("登录失败")
                }
            },

        error: function (data) {
            console.log(data);
        }
    })
}








