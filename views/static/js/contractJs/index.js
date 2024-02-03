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
const abi = [
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_admin",
                "type": "address"
            }
        ],
        "name": "addAdmin",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "initiator",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "_assetId",
                "type": "uint256"
            }
        ],
        "name": "approveAssetTransfer",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "stateMutability": "nonpayable",
        "type": "constructor"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "id",
                "type": "uint256"
            }
        ],
        "name": "AssetApproved",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "id",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "name",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "description",
                "type": "string"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "owner",
                "type": "address"
            }
        ],
        "name": "AssetCreated",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "id",
                "type": "uint256"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "to",
                "type": "address"
            }
        ],
        "name": "AssetTransferRequested",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "id",
                "type": "uint256"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "to",
                "type": "address"
            }
        ],
        "name": "AssetTransferred",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "initiator",
                "type": "address"
            },
            {
                "internalType": "string",
                "name": "_name",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "_description",
                "type": "string"
            }
        ],
        "name": "createAsset",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "initiator",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "_assetId",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "_to",
                "type": "address"
            }
        ],
        "name": "requestAssetTransfer",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "name": "admins",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "getAllRequests",
        "outputs": [
            {
                "components": [
                    {
                        "internalType": "uint256",
                        "name": "id",
                        "type": "uint256"
                    },
                    {
                        "internalType": "string",
                        "name": "name",
                        "type": "string"
                    },
                    {
                        "internalType": "address",
                        "name": "ownerAddress",
                        "type": "address"
                    },
                    {
                        "internalType": "address",
                        "name": "isApprovedOwner",
                        "type": "address"
                    }
                ],
                "internalType": "struct AssetManagement.Request[]",
                "name": "",
                "type": "tuple[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "_assetId",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "_address",
                "type": "address"
            }
        ],
        "name": "getAssetDetails",
        "outputs": [
            {
                "components": [
                    {
                        "internalType": "uint256",
                        "name": "id",
                        "type": "uint256"
                    },
                    {
                        "internalType": "string",
                        "name": "name",
                        "type": "string"
                    },
                    {
                        "internalType": "string",
                        "name": "description",
                        "type": "string"
                    },
                    {
                        "internalType": "address",
                        "name": "ownerAddress",
                        "type": "address"
                    },
                    {
                        "internalType": "address",
                        "name": "isApprovedOwner",
                        "type": "address"
                    }
                ],
                "internalType": "struct AssetManagement.Asset",
                "name": "",
                "type": "tuple"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_address",
                "type": "address"
            }
        ],
        "name": "getpersonAssets",
        "outputs": [
            {
                "components": [
                    {
                        "internalType": "uint256",
                        "name": "id",
                        "type": "uint256"
                    },
                    {
                        "internalType": "string",
                        "name": "name",
                        "type": "string"
                    },
                    {
                        "internalType": "string",
                        "name": "description",
                        "type": "string"
                    },
                    {
                        "internalType": "address",
                        "name": "ownerAddress",
                        "type": "address"
                    },
                    {
                        "internalType": "address",
                        "name": "isApprovedOwner",
                        "type": "address"
                    }
                ],
                "internalType": "struct AssetManagement.Asset[]",
                "name": "",
                "type": "tuple[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_address",
                "type": "address"
            }
        ],
        "name": "getpersonRequests",
        "outputs": [
            {
                "components": [
                    {
                        "internalType": "uint256",
                        "name": "id",
                        "type": "uint256"
                    },
                    {
                        "internalType": "string",
                        "name": "name",
                        "type": "string"
                    },
                    {
                        "internalType": "address",
                        "name": "ownerAddress",
                        "type": "address"
                    },
                    {
                        "internalType": "address",
                        "name": "isApprovedOwner",
                        "type": "address"
                    }
                ],
                "internalType": "struct AssetManagement.Request[]",
                "name": "",
                "type": "tuple[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "owner",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "name": "requests",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "id",
                "type": "uint256"
            },
            {
                "internalType": "string",
                "name": "name",
                "type": "string"
            },
            {
                "internalType": "address",
                "name": "ownerAddress",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "isApprovedOwner",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "totalAssets",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    }
]
let web3 = new Web3('ws://localhost:8546');
// var myContract = new web3.eth.Contract(abi, '0x33975e9dD0a6F4BB3DF7DCbDaDb2c9F4EdC2961C');
// alert(daiaccounts)
let daiaccounts = [];
console.log("web3:",web3)
web3.eth.getAccounts().then(function(accounts) {

    daiaccounts[0] = accounts[0];
    alert(daiaccounts)
    console.log( daiaccounts[0])
    $.ajax({
        url: "http://localhost:8080/dai/getallAssets",
        type: "post",
        data: {"initiator": daiaccounts[0]},
        dataType: "json",
        success: function (data) {
// console.log(document.getElementsByClassName("row")[2].innerHTML)
            console.log(data)

            for (var i = 0; i < data.data.length; i++) {


                let inserString =
                    `
                 <tr class="odd gradeX">
                  <td>  ${data.data[i].Id}</td>
                  <td> ${data.data[i].Id}</td>
                  <td> ${data.data[i].Id}</td>
                  <td class="center">  ${data.data[i].Id}</td>
                  <td class="center"> ${data.data[i].Id}</td>
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

})

ethereum.on('accountsChanged', function(account) {
    // Time to reload your interface with accounts[0]!


    daiaccounts[0]=account[0]
    alert(daiaccounts)
    $(".showAccount").html(account);
    $.ajax({
        url: "http://localhost:8080/dai/getallAssets",
        type: "post",
        data: {"initiator": daiaccounts[0]},
        dataType: "json",
        success: function (data) {
// console.log(document.getElementsByClassName("row")[2].innerHTML)
            console.log(data)

            for (var i = 0; i < data.data.length; i++) {


                let inserString =
                    `
                 <tr class="odd gradeX">
                  <td>  ${data.data[i].Title}</td>
                  <td>Internet Explorer 4.0</td>
                  <td>Win 95+</td>
                  <td class="center"> 4</td>
                  <td class="center">X</td>
                </tr>
      `
                let inserStrings = decodeURIComponent(inserString);
                $.ajax({
                    url: 'http://localhost:8080/index',
                    type: 'get',
                    datatype: 'json',
                    success: function (res) {
                        $(".add").html(inserStrings + document.getElementsByClassName("add")[0].innerHTML)
                    }
                })

            }

        },
        error: function (data) {
            console.log(data)

        }
    })

})


//ajax调用合约 返回所有公益项目的具体详情
$("#name").html(j_query)
// console.log( daiaccounts[0])


//ajax调用合约 返回某个公益项目的具体详情

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
