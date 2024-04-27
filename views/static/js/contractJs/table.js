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
                callback: function (notification) {
                    console.log('action button');
                }
            }
        },
        showProgress: showProgressBar
    };
}
let query = window.location.search;
// alert(query)
var j_query=query.substr(10,1)
// alert(j_query)
// const abi =[
//     {
//         "inputs": [
//             {
//                 "internalType": "address",
//                 "name": "initiator",
//                 "type": "address"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "_assetId",
//                 "type": "uint256"
//             }
//         ],
//         "name": "approveAsset",
//         "outputs": [
//             {
//                 "internalType": "bool",
//                 "name": "",
//                 "type": "bool"
//             }
//         ],
//         "stateMutability": "nonpayable",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "address",
//                 "name": "initiator",
//                 "type": "address"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "_assetId",
//                 "type": "uint256"
//             }
//         ],
//         "name": "approveAssetTransfer",
//         "outputs": [
//             {
//                 "internalType": "bool",
//                 "name": "",
//                 "type": "bool"
//             }
//         ],
//         "stateMutability": "nonpayable",
//         "type": "function"
//     },
//     {
//         "inputs": [],
//         "stateMutability": "nonpayable",
//         "type": "constructor"
//     },
//     {
//         "anonymous": false,
//         "inputs": [
//             {
//                 "indexed": true,
//                 "internalType": "uint256",
//                 "name": "id",
//                 "type": "uint256"
//             }
//         ],
//         "name": "AssetApproved",
//         "type": "event"
//     },
//     {
//         "anonymous": false,
//         "inputs": [
//             {
//                 "indexed": true,
//                 "internalType": "uint256",
//                 "name": "id",
//                 "type": "uint256"
//             },
//             {
//                 "indexed": false,
//                 "internalType": "string",
//                 "name": "name",
//                 "type": "string"
//             },
//             {
//                 "indexed": false,
//                 "internalType": "string",
//                 "name": "description",
//                 "type": "string"
//             },
//             {
//                 "indexed": true,
//                 "internalType": "address",
//                 "name": "owner",
//                 "type": "address"
//             }
//         ],
//         "name": "AssetCreated",
//         "type": "event"
//     },
//     {
//         "anonymous": false,
//         "inputs": [
//             {
//                 "indexed": true,
//                 "internalType": "uint256",
//                 "name": "id",
//                 "type": "uint256"
//             },
//             {
//                 "indexed": true,
//                 "internalType": "address",
//                 "name": "from",
//                 "type": "address"
//             },
//             {
//                 "indexed": true,
//                 "internalType": "address",
//                 "name": "to",
//                 "type": "address"
//             }
//         ],
//         "name": "AssetTransferRequested",
//         "type": "event"
//     },
//     {
//         "anonymous": false,
//         "inputs": [
//             {
//                 "indexed": true,
//                 "internalType": "uint256",
//                 "name": "id",
//                 "type": "uint256"
//             },
//             {
//                 "indexed": true,
//                 "internalType": "address",
//                 "name": "from",
//                 "type": "address"
//             },
//             {
//                 "indexed": true,
//                 "internalType": "address",
//                 "name": "to",
//                 "type": "address"
//             }
//         ],
//         "name": "AssetTransferred",
//         "type": "event"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "address",
//                 "name": "initiator",
//                 "type": "address"
//             },
//             {
//                 "internalType": "string",
//                 "name": "_name",
//                 "type": "string"
//             },
//             {
//                 "internalType": "string",
//                 "name": "_description",
//                 "type": "string"
//             }
//         ],
//         "name": "createAsset",
//         "outputs": [
//             {
//                 "components": [
//                     {
//                         "internalType": "uint256",
//                         "name": "id",
//                         "type": "uint256"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "name",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "description",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "ownerAddress",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "isApprovedOwner",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "uint256",
//                         "name": "status",
//                         "type": "uint256"
//                     }
//                 ],
//                 "internalType": "struct AssetManagement.Asset",
//                 "name": "",
//                 "type": "tuple"
//             }
//         ],
//         "stateMutability": "nonpayable",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "address",
//                 "name": "initiator",
//                 "type": "address"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "_id",
//                 "type": "uint256"
//             }
//         ],
//         "name": "personReceiveAsset",
//         "outputs": [
//             {
//                 "components": [
//                     {
//                         "internalType": "uint256",
//                         "name": "id",
//                         "type": "uint256"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "name",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "description",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "receiveAddress",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "uint256",
//                         "name": "result",
//                         "type": "uint256"
//                     }
//                 ],
//                 "internalType": "struct AssetManagement.receiveAsset",
//                 "name": "",
//                 "type": "tuple"
//             }
//         ],
//         "stateMutability": "nonpayable",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "address",
//                 "name": "initiator",
//                 "type": "address"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "_id",
//                 "type": "uint256"
//             }
//         ],
//         "name": "personReturnAsset",
//         "outputs": [
//             {
//                 "internalType": "bool",
//                 "name": "",
//                 "type": "bool"
//             }
//         ],
//         "stateMutability": "nonpayable",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "address",
//                 "name": "initiator",
//                 "type": "address"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "_assetId",
//                 "type": "uint256"
//             }
//         ],
//         "name": "refuseAssetRequest",
//         "outputs": [
//             {
//                 "internalType": "bool",
//                 "name": "",
//                 "type": "bool"
//             }
//         ],
//         "stateMutability": "nonpayable",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "address",
//                 "name": "initiator",
//                 "type": "address"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "_assetId",
//                 "type": "uint256"
//             }
//         ],
//         "name": "rejectAsset",
//         "outputs": [
//             {
//                 "internalType": "bool",
//                 "name": "",
//                 "type": "bool"
//             }
//         ],
//         "stateMutability": "nonpayable",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "address",
//                 "name": "initiator",
//                 "type": "address"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "_assetId",
//                 "type": "uint256"
//             },
//             {
//                 "internalType": "address",
//                 "name": "_to",
//                 "type": "address"
//             }
//         ],
//         "name": "requestAssetTransfer",
//         "outputs": [],
//         "stateMutability": "nonpayable",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "address",
//                 "name": "",
//                 "type": "address"
//             }
//         ],
//         "name": "admins",
//         "outputs": [
//             {
//                 "internalType": "bool",
//                 "name": "",
//                 "type": "bool"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "uint256",
//                 "name": "",
//                 "type": "uint256"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "",
//                 "type": "uint256"
//             }
//         ],
//         "name": "AssetGroup",
//         "outputs": [
//             {
//                 "internalType": "uint256",
//                 "name": "id",
//                 "type": "uint256"
//             },
//             {
//                 "internalType": "string",
//                 "name": "name",
//                 "type": "string"
//             },
//             {
//                 "internalType": "string",
//                 "name": "description",
//                 "type": "string"
//             },
//             {
//                 "internalType": "address",
//                 "name": "ownerAddress",
//                 "type": "address"
//             },
//             {
//                 "internalType": "address",
//                 "name": "isApprovedOwner",
//                 "type": "address"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "status",
//                 "type": "uint256"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "uint256",
//                 "name": "",
//                 "type": "uint256"
//             }
//         ],
//         "name": "assets",
//         "outputs": [
//             {
//                 "internalType": "uint256",
//                 "name": "id",
//                 "type": "uint256"
//             },
//             {
//                 "internalType": "string",
//                 "name": "name",
//                 "type": "string"
//             },
//             {
//                 "internalType": "string",
//                 "name": "description",
//                 "type": "string"
//             },
//             {
//                 "internalType": "address",
//                 "name": "ownerAddress",
//                 "type": "address"
//             },
//             {
//                 "internalType": "address",
//                 "name": "isApprovedOwner",
//                 "type": "address"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "status",
//                 "type": "uint256"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [],
//         "name": "getAllAssets",
//         "outputs": [
//             {
//                 "components": [
//                     {
//                         "internalType": "uint256",
//                         "name": "id",
//                         "type": "uint256"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "name",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "description",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "ownerAddress",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "isApprovedOwner",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "uint256",
//                         "name": "status",
//                         "type": "uint256"
//                     }
//                 ],
//                 "internalType": "struct AssetManagement.Asset[]",
//                 "name": "",
//                 "type": "tuple[]"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [],
//         "name": "getAllRecords",
//         "outputs": [
//             {
//                 "components": [
//                     {
//                         "internalType": "string",
//                         "name": "name1",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "description1",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "ownerAddress1",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "uint256",
//                         "name": "typeOf",
//                         "type": "uint256"
//                     }
//                 ],
//                 "internalType": "struct AssetManagement.Record[]",
//                 "name": "",
//                 "type": "tuple[]"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [],
//         "name": "getAllRequests",
//         "outputs": [
//             {
//                 "components": [
//                     {
//                         "internalType": "uint256",
//                         "name": "id",
//                         "type": "uint256"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "name",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "ownerAddress",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "isApprovedOwner",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "uint256",
//                         "name": "anwser",
//                         "type": "uint256"
//                     }
//                 ],
//                 "internalType": "struct AssetManagement.Request[]",
//                 "name": "",
//                 "type": "tuple[]"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "address",
//                 "name": "_address",
//                 "type": "address"
//             }
//         ],
//         "name": "getpersonAssets",
//         "outputs": [
//             {
//                 "components": [
//                     {
//                         "internalType": "uint256",
//                         "name": "id",
//                         "type": "uint256"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "name",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "description",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "ownerAddress",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "isApprovedOwner",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "uint256",
//                         "name": "status",
//                         "type": "uint256"
//                     }
//                 ],
//                 "internalType": "struct AssetManagement.Asset[]",
//                 "name": "",
//                 "type": "tuple[]"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "address",
//                 "name": "_address",
//                 "type": "address"
//             }
//         ],
//         "name": "getpersonRequests",
//         "outputs": [
//             {
//                 "components": [
//                     {
//                         "internalType": "uint256",
//                         "name": "id",
//                         "type": "uint256"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "name",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "ownerAddress",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "isApprovedOwner",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "uint256",
//                         "name": "anwser",
//                         "type": "uint256"
//                     }
//                 ],
//                 "internalType": "struct AssetManagement.Request[]",
//                 "name": "",
//                 "type": "tuple[]"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [],
//         "name": "getReceiveAsset",
//         "outputs": [
//             {
//                 "components": [
//                     {
//                         "internalType": "uint256",
//                         "name": "id",
//                         "type": "uint256"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "name",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "string",
//                         "name": "description",
//                         "type": "string"
//                     },
//                     {
//                         "internalType": "address",
//                         "name": "receiveAddress",
//                         "type": "address"
//                     },
//                     {
//                         "internalType": "uint256",
//                         "name": "result",
//                         "type": "uint256"
//                     }
//                 ],
//                 "internalType": "struct AssetManagement.receiveAsset[]",
//                 "name": "",
//                 "type": "tuple[]"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [],
//         "name": "owner",
//         "outputs": [
//             {
//                 "internalType": "address",
//                 "name": "",
//                 "type": "address"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "uint256",
//                 "name": "",
//                 "type": "uint256"
//             }
//         ],
//         "name": "receiveRequest",
//         "outputs": [
//             {
//                 "internalType": "uint256",
//                 "name": "id",
//                 "type": "uint256"
//             },
//             {
//                 "internalType": "string",
//                 "name": "name",
//                 "type": "string"
//             },
//             {
//                 "internalType": "string",
//                 "name": "description",
//                 "type": "string"
//             },
//             {
//                 "internalType": "address",
//                 "name": "receiveAddress",
//                 "type": "address"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "result",
//                 "type": "uint256"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "uint256",
//                 "name": "",
//                 "type": "uint256"
//             }
//         ],
//         "name": "records",
//         "outputs": [
//             {
//                 "internalType": "string",
//                 "name": "name1",
//                 "type": "string"
//             },
//             {
//                 "internalType": "string",
//                 "name": "description1",
//                 "type": "string"
//             },
//             {
//                 "internalType": "address",
//                 "name": "ownerAddress1",
//                 "type": "address"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "typeOf",
//                 "type": "uint256"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [
//             {
//                 "internalType": "uint256",
//                 "name": "",
//                 "type": "uint256"
//             }
//         ],
//         "name": "requests",
//         "outputs": [
//             {
//                 "internalType": "uint256",
//                 "name": "id",
//                 "type": "uint256"
//             },
//             {
//                 "internalType": "string",
//                 "name": "name",
//                 "type": "string"
//             },
//             {
//                 "internalType": "address",
//                 "name": "ownerAddress",
//                 "type": "address"
//             },
//             {
//                 "internalType": "address",
//                 "name": "isApprovedOwner",
//                 "type": "address"
//             },
//             {
//                 "internalType": "uint256",
//                 "name": "anwser",
//                 "type": "uint256"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [],
//         "name": "totalAssets",
//         "outputs": [
//             {
//                 "internalType": "uint256",
//                 "name": "",
//                 "type": "uint256"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "inputs": [],
//         "name": "totalreceiveRequest",
//         "outputs": [
//             {
//                 "internalType": "uint256",
//                 "name": "",
//                 "type": "uint256"
//             }
//         ],
//         "stateMutability": "view",
//         "type": "function"
//     }
// ]

var  adminAddress;



$.ajax({
    url: "http://localhost:8080/dai/getSaveAddress",
    type: "get",
    dataType: "json",
    success: function (data) {
        adminAddress=data.data
        $("#admin").html("欢迎"+adminAddress);
        // alert("成功获取保存的session:" +adminAddress)
        //管理员自身资产
        $.ajax({
            url: "http://localhost:8080/dai/getallAssets",
            type: "post",
            data: {"initiator": adminAddress},
            dataType: "json",
            success: function (data) {
// console.log(document.getElementsByClassName("row")[2].innerHTML)
                console.log("管理员自身资产:",data)
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

        //所有的资产领用请求
        $.ajax({
            url: "http://localhost:8080/dai/getAssetReceive",
            type: "post",
            data: {"initiator":adminAddress},
            dataType: "json",
            success: function (data) {
// console.log(document.getElementsByClassName("row")[2].innerHTML)
                console.log("所有的资产领用请求:",data)

                for (var i = 0; i < data.data.length; i++) {

                    if (data.data[i].Result==0){
                    let inserString =
                        `
                 <tr>
                  <td style="text-align: center;" >  ${data.data[i].Id}</td>
                  <td style="text-align: center;" > ${data.data[i].Name}</td>
                  <td style="text-align: center;" >  ${data.data[i].Description}</td>
                     <td style="text-align: center;"> ${data.data[i].ReceiveAddress}</td>
                   <td  style="text-align: center;"  ><button onclick="receiveApprove(event)" style="color: green">批准</button>  <button onclick="rejectApprove(event)" style="color: red">拒绝</button></td>
                   
                </tr>
      `

                    let inserStrings = decodeURIComponent(inserString);
                    $.ajax({
                        url: 'http://localhost:8080/table',
                        type: 'get',
                        datatype: 'json',
                        success: function (res) {
                            $(".viewReceive").html(inserStrings + document.getElementsByClassName("viewReceive")[0].innerHTML)
                        }
                    })


                }else if (data.data[i].Result==1){
                        let inserString =
                            `
                 <tr>
                  <td style="text-align: center;" >  ${data.data[i].Id}</td>
                  <td style="text-align: center;" > ${data.data[i].Name}</td>
                  <td style="text-align: center;" >  ${data.data[i].Description}</td>
                     <td style="text-align: center;"> ${data.data[i].ReceiveAddress}</td>
                   <td  style="text-align: center;"  ><button  style="color: green">已批准</button></td>
                   
                </tr>
      `

                        let inserStrings = decodeURIComponent(inserString);
                        $.ajax({
                            url: 'http://localhost:8080/table',
                            type: 'get',
                            datatype: 'json',
                            success: function (res) {
                                $(".viewReceive").html(inserStrings + document.getElementsByClassName("viewReceive")[0].innerHTML)
                            }
                        })


                    }
                    else if (data.data[i].Result==2){
                        let inserString =
                            `
                 <tr>
                  <td style="text-align: center;" >  ${data.data[i].Id}</td>
                  <td style="text-align: center;" > ${data.data[i].Name}</td>
                  <td style="text-align: center;" >  ${data.data[i].Description}</td>
                     <td style="text-align: center;"> ${data.data[i].ReceiveAddress}</td>
                   <td  style="text-align: center;"  ><button  style="color: red">已拒绝</button></td>
                   
                </tr>
      `

                        let inserStrings = decodeURIComponent(inserString);
                        $.ajax({
                            url: 'http://localhost:8080/table',
                            type: 'get',
                            datatype: 'json',
                            success: function (res) {
                                $(".viewReceive").html(inserStrings + document.getElementsByClassName("viewReceive")[0].innerHTML)
                            }
                        })


                    }
                }

            },
            error: function (data) {
                console.log(data)

            }
        })

        //所有的资产转移请求
        $.ajax({
            url: "http://localhost:8080/dai/getAllRequests",
            type: "post",
            data: {"initiator":adminAddress},
            dataType: "json",
            success: function (data) {
// console.log(document.getElementsByClassName("row")[2].innerHTML)
                console.log("所有的资产转移请求:",data)

                for (var i = 0; i < data.data.length; i++) {


                    if (data.data[i].Anwser==0){
                        // alert("未批准")


                    let inserString =
                        `
                 <tr>
                  <td style="text-align: center;" >  ${data.data[i].Id}</td>
                  <td style="text-align: center;" > ${data.data[i].Name}</td>
                  <td style="text-align: center;" >  ${data.data[i].OwnerAddress}</td>
                     <td style="text-align: center;"> ${data.data[i].IsApprovedOwner}</td>
                   <td  style="text-align: center;"  ><button onclick="requestApprove(event)" style="color: green">批准</button>  <button onclick="requestRefuse(event)" style="color: red">拒绝</button></td>
                   
                </tr>
      `

                    let inserStrings = decodeURIComponent(inserString);
                    $.ajax({
                        url: 'http://localhost:8080/table',
                        type: 'get',
                        datatype: 'json',
                        success: function (res) {
                            $(".viewRequest").html(inserStrings + document.getElementsByClassName("viewRequest")[0].innerHTML)
                        }
                    })

                }else if (data.data[i].Anwser==1){
                        // alert("已经批准")
                        let inserString =
                            `
                 <tr>
                  <td style="text-align: center;" >  ${data.data[i].Id}</td>
                  <td style="text-align: center;" > ${data.data[i].Name}</td>
                  <td style="text-align: center;" >  ${data.data[i].OwnerAddress}</td>
                     <td style="text-align: center;"> ${data.data[i].IsApprovedOwner}</td>
                   <td  style="text-align: center;"  ><button style="color: green">已批准</button></td>
                   
                </tr>
      `

                        let inserStrings = decodeURIComponent(inserString);
                        $.ajax({
                            url: 'http://localhost:8080/table',
                            type: 'get',
                            datatype: 'json',
                            success: function (res) {
                                $(".viewRequest").html(inserStrings + document.getElementsByClassName("viewRequest")[0].innerHTML)
                            }
                        })
                    }
                    else if (data.data[i].Anwser==2){
                        // alert("已经批准")
                        let inserString =
                            `
                 <tr>
                  <td style="text-align: center;" >  ${data.data[i].Id}</td>
                  <td style="text-align: center;" > ${data.data[i].Name}</td>
                  <td style="text-align: center;" >  ${data.data[i].OwnerAddress}</td>
                     <td style="text-align: center;"> ${data.data[i].IsApprovedOwner}</td>
                   <td  style="text-align: center;"  ><button style="color: red">已拒绝</button></td>
                   
                </tr>
      `

                        let inserStrings = decodeURIComponent(inserString);
                        $.ajax({
                            url: 'http://localhost:8080/table',
                            type: 'get',
                            datatype: 'json',
                            success: function (res) {
                                $(".viewRequest").html(inserStrings + document.getElementsByClassName("viewRequest")[0].innerHTML)
                            }
                        })
                    }
                    }

            },
            error: function (data) {
                console.log(data)

            }
        })

        //所有的资产库存情况
        $.ajax({
            url: "http://localhost:8080/dai/getAllGroup",
            type: "post",
            dataType: "json",
            success: function (data) {
// console.log(document.getElementsByClassName("row")[2].innerHTML)
                console.log("所有的资产库存请求:",data)

                for (var i = 0; i < data.data.length; i++) {


                    let inserString =
                        `
   <tr>
                             
                                <td style="text-align: center;">${data.data[i].Name}</td>
                                <td style="text-align: center;">${data.data[i].Description}</td>
                                   <td style="text-align: center;">${data.data[i].OwnerAddress}</td>

                            </tr>
                            
             
      `
                    let inserStrings = decodeURIComponent(inserString);
                    $.ajax({
                        url: 'http://localhost:8080/table',
                        type: 'get',
                        datatype: 'json',
                        success: function (res) {
                            $(".viewStock").html(inserStrings + document.getElementsByClassName("viewStock")[0].innerHTML)
                        }
                    })

                }

            },
            error: function (data) {
                console.log(data)

            }
        })
        var r;
        //所有的资产操作情况
        $.ajax({
            url: "http://localhost:8080/dai/getRecords",
            type: "post",
            dataType: "json",
            success: function (data) {
// console.log(document.getElementsByClassName("row")[2].innerHTML)
                console.log("所有的资产操作记录请求:",data)



                for (var i = 0; i < data.data.length; i++) {
                    if (data.data[i].TypeOf=="1"){
                        r="用户领用资产"
                    }else if(data.data[i].TypeOf=="2"){
                        r="用户申请转移资产"
                    }else if(data.data[i].TypeOf=="3"){
                        r="管理员批准转移资产"
                    }else if(data.data[i].TypeOf=="4"){
                        r="管理员拒绝转移资产"
                    }


                    let inserString =

                        `
   <tr className="gradeX">
                             
                                <td style="text-align: center;">${data.data[i].Name1}</td>
                                <td style="text-align: center;">${data.data[i].Description1}</td>
                                   <td style="text-align: center;">${data.data[i].OwnerAddress1}</td>
 <td style="text-align: center;">${r}</td>
                            </tr>
                            
             
      `
                    let inserStrings = decodeURIComponent(inserString);
                    $.ajax({
                        url: 'http://localhost:8080/table',
                        type: 'get',
                        datatype: 'json',
                        success: function (res) {
                            $(".viewRecord").html(inserStrings + document.getElementsByClassName("viewRecord")[0].innerHTML)
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


// })

// ethereum.on('accountsChanged', function(account) {
//     // Time to reload your interface with accounts[0]!
//
//
//     daiaccounts[0]=account[0]
//     alert(daiaccounts)
//     $(".showAccount").html(account);
//     $.ajax({
//         url: "http://localhost:8080/dai/getallAssets",
//         type: "post",
//         data: {"initiator":adminAddress},
//         dataType: "json",
//         success: function (data) {
// // console.log(document.getElementsByClassName("row")[2].innerHTML)
// //             console.log(data)
//
//             for (var i = 0; i < data.data.length; i++) {
//
//
//                 let inserString =
//                     `
//                  <tr class="odd gradeX">
//                   <td>  ${data.data[i].Title}</td>
//                   <td>Internet Explorer 4.0</td>
//                   <td>Win 95+</td>
//                   <td class="center"> 4</td>
//                   <td class="center">X</td>
//                 </tr>
//       `
//                 let inserStrings = decodeURIComponent(inserString);
//                 $.ajax({
//                     url: 'http://localhost:8080/index',
//                     type: 'get',
//                     datatype: 'json',
//                     success: function (res) {
//                         $(".add").html(inserStrings + document.getElementsByClassName("add")[0].innerHTML)
//                     }
//                 })
//
//             }
//
//         },
//         error: function (data) {
//             console.log(data)
//
//         }
//     })
//
// })
//

$("#name").html(j_query);

function newAsset(name,description,adminAddress){
    $.ajax({
        url: "http://localhost:8080/dai/newAsset",
        type: "post",
        data: {"initiator": adminAddress, "name": name, "description": description},
        dataType: "json",
        success: function (data) {
            // alert(data.data)
            console.log(data.data)


        },
        error: function (data) {
            console.log(data)

        }
    })
}

//资产转让
function requestAssetTransfer(id,to){

    // alert(daiaccounts[0])
    // alert("id:"+id)
    // alert("to:"+to)
    $.ajax({
        url: "http://localhost:8080/dai/requestAssetTransfer",
        type: "post",
        data: {"initiator":adminAddress,"id":id,"to":to},
        dataType: "json",
        success: function (data) {
            console.log(data.data)
            alert("资产转让成功")

        },
        error: function (data) {
            console.log(data)

        }
    })

}


//管理员批准资产转移
function requestApprove(event){
    // 获取当前按钮所在的 <td> 元素
    var tdElement = event.target.parentNode;

    // 获取同级的第一个 <td> 的值
    var firstTdValue = tdElement.parentNode.querySelector('td').innerText;

    $.ajax({
        url: "http://localhost:8080/dai/getSaveAddress",
        type: "get",
        dataType: "json",
        success: function (data) {
            adminAddress=data.data
        }
    })
    // alert("firstTdValue:" +firstTdValue);
    // alert("appRover:" +adminAddress);
    $.ajax({
        url: "http://localhost:8080/dai/approveAssetTransfer",
        type: "post",
        data: {"initiator":adminAddress,"id":firstTdValue},
        dataType: "json",
        success: function (data) {
            console.log(data.data)
            alert("管理员批准资产转让成功")

        },
        error: function (data) {
            console.log(data)

        }
        })
}
//管理员拒绝资产转移
function requestRefuse(event){
    // 获取当前按钮所在的 <td> 元素
    var tdElement = event.target.parentNode;

    // 获取同级的第一个 <td> 的值
    var firstTdValue = tdElement.parentNode.querySelector('td').innerText;

    $.ajax({
        url: "http://localhost:8080/dai/getSaveAddress",
        type: "get",
        dataType: "json",
        success: function (data) {
            adminAddress=data.data
        }
    })
    // alert("firstTdValue:" +firstTdValue);
    // alert("appRover:" +adminAddress);
    $.ajax({
        url: "http://localhost:8080/dai/refuseAssetTransfer",
        type: "post",
        data: {"initiator":adminAddress,"id":firstTdValue},
        dataType: "json",
        success: function (data) {
            console.log(data.data)
            alert("管理员拒绝资产转让成功")

        },
        error: function (data) {
            console.log(data)

        }
    })
}
//管理员批准资产领用
function receiveApprove(event){
    // 获取当前按钮所在的 <td> 元素
    var tdElement = event.target.parentNode;

    // 获取同级的第一个 <td> 的值
    var firstTdValue = tdElement.parentNode.querySelector('td').innerText;

    $.ajax({
        url: "http://localhost:8080/dai/getSaveAddress",
        type: "get",
        dataType: "json",
        success: function (data) {
            adminAddress=data.data
        }
    })
    // alert("firstTdValue:" +firstTdValue);
    // alert("appRover:" +adminAddress);
    $.ajax({
        url: "http://localhost:8080/dai/receiveApprove",
        type: "post",
        data: {"initiator":adminAddress,"id":firstTdValue},
        dataType: "json",
        success: function (data) {
            console.log(data.data)
            alert("管理员批准资产领用成功")

        },
        error: function (data) {
            console.log(data)

        }
    })
}
//管理员拒绝资产领用
function rejectApprove(event){
    // 获取当前按钮所在的 <td> 元素
    var tdElement = event.target.parentNode;

    // 获取同级的第一个 <td> 的值
    var firstTdValue = tdElement.parentNode.querySelector('td').innerText;

    $.ajax({
        url: "http://localhost:8080/dai/getSaveAddress",
        type: "get",
        dataType: "json",
        success: function (data) {
            adminAddress=data.data
        }
    })
    // alert("firstTdValue:" +firstTdValue);
    // alert("appRover:" +adminAddress);
    $.ajax({
        url: "http://localhost:8080/dai/rejectApprove",
        type: "post",
        data: {"initiator":adminAddress,"id":firstTdValue},
        dataType: "json",
        success: function (data) {
            console.log(data.data)
            alert("管理员拒绝资产领用成功")

        },
        error: function (data) {
            console.log(data)

        }
    })
}





var id;

function get() {

    var arr = document.getElementsByTagName('img');
    for (var i = 0; i < arr.length; i++) {
        arr[i].onclick = function () {
            id = this.id;

            window.location.href = "/d/?" + id;

        }

    }


}

//项目加入收藏
function ins() {
    var a;
    var arr = document.getElementsByTagName('li');

    for (var i = 0; i < arr.length; i++) {
        arr[i].onclick = function () {
            a = this.id;
            console.log(a)
            $.ajax({
                url: "http://localhost:8080/dai/insFunding",
                type: "post",
                dataType: "json",
                data: {id: a},
                success: function (data) {
                    console.log(data.data)
                    var options = getOptions();
                    options.title = '信息';
                    options.description = '项目加入收藏成功，您可以在收藏页面查看您收藏的公益';
                    options.width = 300;
                    options.type = 'success';
                    options.closeTimeout = 3000;
                    GrowlNotification.notify(options);


                }
            })
        }


    }


}



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
                callback: function (notification) {
                    console.log('action button');
                }
            }
        },
        showProgress: showProgressBar
    };
}
