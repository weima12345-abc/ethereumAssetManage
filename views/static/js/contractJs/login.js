
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
    // alert(username,password)

    // if(username=="18174001013"&&password=="123"){
    //     window.location.href="http://localhost:8080/manage" //登入成功就切换到管理页面
    // }
    //  console.log(password ,username)//测试是否得到这些值
    $.ajax({//通过ajax触发post请求
        url: " http://localhost:8080/dai/login",//post请求路径
        type: "POST",
        data: {username: username, password: password}, //对应相对数据
        dataType:"json",
        success: function (data) {
            console.log(data)
            // alert(data.data)
            if (data.data == 1) {
                // if (username == "0x7fE3bC3FCc9d7d19254D4DE515eC395dA9768331" && password == "123456") {
                    var options = getOptions();
                    options.title = '信息';
                    options.description = '成功登入后台';
                    options.width = 300;
                    options.type = 'success';
                    options.closeTimeout = 1000;
                    GrowlNotification.notify(options);
                    //保存登录的用户名
                    $.ajax({//通过ajax触发post请求
                        url: " http://localhost:8080/dai/saveAddress",//post请求路径
                        type: "POST",
                        data: {adminAddress: username}, //对应相对数据
                        dataType: "json",
                        success: function (data) {

                            alert("登录名保存成功:"+data.data)

                        }
                    })

                    window.location.href = "http://localhost:8080/table" //登入成功就切换到管理页面
                // } else {
                //
                //     var options = getOptions();
                //     options.title = '信息';
                //     options.description = '成功登入个人主页';
                //     options.width = 300;
                //     options.type = 'success';
                //     options.closeTimeout = 1000;
                //     GrowlNotification.notify(options);
                //
                //     window.location.href = "http://localhost:8080/index/?username=" + username + "&&password="+password //登入成功就切换到主页面
                //
                //
                // }
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








