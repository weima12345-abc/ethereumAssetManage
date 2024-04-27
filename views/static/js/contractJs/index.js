
$.ajax({
    url: "http://localhost:8080/dai/getSaveAddress",
    type: "get",
    dataType: "json",
    success: function (data) {
        adminAddress = data.data
        $("#admin").html("欢迎" + adminAddress);

        $("#admin").html("欢迎" + adminAddress);
        // alert("成功获取保存的session:" + adminAddress)
        //管理员自身资产
        $.ajax({
            url: "http://localhost:8080/dai/getallAssets",
            type: "post",
            data: {"initiator": adminAddress},
            dataType: "json",
            success: function (data) {
// console.log(document.getElementsByClassName("row")[2].innerHTML)
                console.log("管理员自身资产:", data)
                $(".view").html("")
                for (var i = 0; i < data.data.length; i++) {

                    let inserString =
                        `
                 <tr class="odd gradeX">
                  <td style="text-align: center;">  ${data.data[i].Id}</td>
                  <td style="text-align: center;"> ${data.data[i].Name}</td>
                  <td style="text-align: center;"> ${data.data[i].Description}</td>
                  <td style="text-align: center;">  ${data.data[i].OwnerAddress}</td>
               
               
                </tr>
      `
                    let inserStrings = decodeURIComponent(inserString);
                    $.ajax({
                        url: 'http://localhost:8080/table',
                        type: 'get',
                        datatype: 'json',
                        success: function (res) {
                            $(".view").html(inserStrings + document.getElementsByClassName("view")[0].innerHTML)
                        }
                    })

                }

            },
            error: function (data) {
                console.log(data)

            }
        })

    }
})

