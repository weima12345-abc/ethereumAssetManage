// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract AssetManagement {
    address public owner;//用于存储合约的拥有者地址。

    struct Asset {//表示资产的基本信息，包括资产 ID、名称、描述、拥有者地址和是否获得批准。
        uint256 id;
        string name;//名称
        string description;//描述
        address ownerAddress;//拥有者
        address isApprovedOwner;//资产转让转让接收者

    }
    struct Request {//表示资产的转移信息，包括资产 ID、名称、描述、拥有者地址和是否获得批准。
        uint256 id;
        string name;//名称
        address ownerAddress;//拥有者
        address isApprovedOwner;//资产转让转让接收者
        uint256 anwser;
    }
    struct Record {

        string name1;//名称
        string description1;//描述
        address ownerAddress1;//
        uint256 typeOf;
    }
    constructor() {
        owner = msg.sender;
        admins[msg.sender] = true;
    }

    mapping(address => bool) public admins;//存储管理员地址及其权限的布尔值。
    mapping(address => mapping(uint256 => bool))  approvalRequests;
    //approvalRequests : 使用嵌套映射存储转让请求，其中外层映射键是请求接收者的地址，内层映射键是资产 ID。
    mapping(uint256 => Asset)  public assets;//使用映射存储资产，其中键是资产 ID，值是资产的详细信息。
    mapping(uint256 => Asset[]) public AssetGroup;
    uint256 public totalAssets ;//存储合约中总资产数量的公共变量
    mapping(uint256  =>Request) oneRequest;
    Request[]  public requests;//储存所有用户的资产转移信息
    mapping(address => mapping(uint256 =>Asset )  )  personAssets;//储存用户所拥有的某一个资产。
    mapping(address => Asset[] )  personAllAsset;//储存用户所拥有的所有资产

    mapping(address => Request[]) personRequests  ;//存储当前用户的资产转移请求

    mapping(uint256 => Record)  public records;
    mapping(uint256 =>  Record[]) assetOperates;

    event AssetCreated(uint256 indexed id, string name, string description, address indexed owner);//定义了资产创建时触发的事件。
    event AssetTransferRequested(uint256 indexed id, address indexed from, address indexed to);//定义了资产转让请求时触发的事件。
    event AssetTransferred(uint256 indexed id, address indexed from, address indexed to);//定义了资产转让完成时触发的事件。
    event AssetApproved(uint256 indexed id);//定义了资产转让获得批准时触发的事件。

    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can call this function");
        _;
    }

    modifier onlyAdmin() {
        require(admins[msg.sender], "Only admins can call this function");
        _;
    }

//合约拥有者添加管理员的函数
    function addAdmin(address _admin) public onlyOwner {
        admins[_admin] = true;
    }
    uint256 totalRecords=0;
    uint256 totalRequests=0;
//创建资产的函数，生成新的资产并触发 AssetCreated 事件。
    function createAsset(address initiator,string memory _name, string memory _description) public returns(Asset memory) {
        totalAssets++;
        address a=0x0000000000000000000000000000000000000000;
        totalRecords++;

        assets[totalAssets] = Asset(totalAssets, _name, _description, initiator, a);
        AssetGroup[1].push(assets[totalAssets]);


        personAssets[initiator][totalAssets]=assets[totalAssets];
        personAllAsset[initiator].push(assets[totalAssets]);

        records[totalRecords] = Record( _name, _description, initiator,1);
        assetOperates[1].push(records[totalRecords]);
        emit AssetCreated(totalAssets, _name, _description, initiator);
        return assets[totalAssets];
    }
//请求资产转让的函数，生成转让请求并触发 AssetTransferRequested 事件。
    function requestAssetTransfer(address initiator,uint256 _assetId, address _to) public {
        Asset storage asset = assets[_assetId];
        asset.isApprovedOwner=_to;
        require(asset.ownerAddress == initiator, "Only the asset owner can request transfer");
        approvalRequests[_to][_assetId] = true;
        totalRequests++;


        Request memory r=Request(_assetId,asset.name,asset.ownerAddress,_to,0);
        oneRequest[totalRequests]=r;
        personRequests[initiator].push(r);
        requests.push(r);

        totalRecords++;
        records[totalRecords] = Record(assets[_assetId].name, assets[_assetId].description, assets[_assetId].ownerAddress,2);
        assetOperates[1].push(records[totalRecords]);
        emit AssetTransferRequested(_assetId,initiator, _to);
    }

//批准资产转让的函数，仅允许管理员调用，更新资产拥有者并触发相应事件。
    function approveAssetTransfer(address initiator, uint256 _assetId) public  returns(bool) {
        require(approvalRequests[assets[_assetId].isApprovedOwner][_assetId], "No transfer request found");
        require(admins[initiator], "Only admins can call this function");
        Asset storage asset = assets[_assetId];
        address oldOwner = asset.ownerAddress;

        asset.ownerAddress = asset.isApprovedOwner;
        address revicer =asset.isApprovedOwner;
        asset.isApprovedOwner = address(0);
        personAllAsset[revicer].push(asset);




        delete AssetGroup[1];



// 将 assets 数组中的元素添加到 Assets[1] 数组中
        for (uint j = 1; j <=totalAssets; j++) {
            AssetGroup[1].push(assets[j]);
        }
        totalRecords++;
        records[totalRecords] = Record( assets[_assetId].name, assets[_assetId].description, assets[_assetId].ownerAddress,3);
        assetOperates[1].push(records[totalRecords]);


        uint256 requestL=requests.length;
        oneRequest[_assetId].anwser=1;
        delete requests;
        for (uint j = 1; j <=requestL; j++) {
            requests.push(oneRequest[j]);
        }


        uint l = personAllAsset[asset.isApprovedOwner].length;

        // 修复数组长度问题
        if (l == 0) {
            // personAssets[asset.isApprovedOwner].push(asset);
            personAllAsset[asset.isApprovedOwner].push(asset);
        }

        if (personAllAsset[oldOwner].length > 1) {
            for (uint i = 0; i < personAllAsset[oldOwner].length; i++) {
                if (personAllAsset[oldOwner][i].id == _assetId) {
                    // 将找到的元素删除
                    delete personAllAsset[oldOwner][i];

                    // 将数组末尾的元素移到当前位置，然后删除末尾元素
                    personAllAsset[oldOwner][i] = personAllAsset[oldOwner][personAllAsset[oldOwner].length - 1];
                    personAllAsset[oldOwner].pop();

                    // 退出循环，因为我们已经找到并处理了要删除的元素
                    break;
                }
            }


            // 删除卖家（oldOwner）的资产
            delete personAssets[oldOwner][_assetId];
            emit AssetTransferred(_assetId, asset.ownerAddress, initiator);
            emit AssetApproved(_assetId);

            return true;


        } else if (personAllAsset[oldOwner].length == 1) {
            // 如果只有一个资产，直接删除
            delete personAllAsset[oldOwner][0];
            personAllAsset[oldOwner].pop();
            delete approvalRequests[asset.isApprovedOwner][_assetId];
            delete personAssets[oldOwner][_assetId];

            emit AssetTransferred(_assetId, asset.ownerAddress, initiator);
            emit AssetApproved(_assetId);

            return true;
        }


        return false;

    }

    function refuseAssetRequest(address initiator, uint256 _assetId) public  returns(bool){
        require(approvalRequests[assets[_assetId].isApprovedOwner][_assetId], "No transfer request found");
        require(admins[initiator], "Only admins can call this function");
        uint256 requestL=requests.length;
        oneRequest[_assetId].anwser=2;
        delete requests;
        for (uint j = 1; j <=requestL; j++) {
            requests.push(oneRequest[j]);
        }
        return false;
    }



//获取用户所拥有的某一个资产详细信息的函数，允许外部调用并返回资产结构。
    function getAssetDetails(uint256 _assetId,address _address) public view returns (Asset memory) {
        return personAssets[_address][_assetId];
    }
    //获取用户所有资产详细信息的函数，允许外部调用并返回资产结构。
    function getpersonAssets(address _address) public view returns (Asset[] memory) {
        return personAllAsset[_address];
    }

    //获取当前用户发起的资产转移请求，允许外部调用并返回资产结构。
    function getpersonRequests(address _address) public view returns (Request[] memory) {
        return personRequests[_address];
    }
    //获取所有发起的资产转移请求，允许外部调用并返回资产结构。
    function getAllRequests() public view returns (Request[] memory) {
        return requests;
    }

    //获取所有发起的资产转移请求，允许外部调用并返回资产结构。
    function getAllAssets() public view returns (Asset[] memory) {


        return AssetGroup[1];
    }
    function getAllRecords() public view returns (Record[] memory) {


        return assetOperates[1];
    }

}

