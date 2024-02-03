pragma solidity ^0.8.0;

contract LogisticsContract {
    struct Product {
        address currentOwner;
        string[] logisticsHistory;  // 用于记录轨迹的字符串数组
        bool isVerified;
    }

    mapping(string => Product) public products;    // 产品的唯一标识符映射到产品结构体

    event LogisticsUpdated(string indexed productId, address indexed newOwner, string newLogistics);  // 事件，用于记录物流信息变更

    //修饰符
    modifier onlyProductOwner(string memory productId) {
        require(msg.sender == products[productId].currentOwner, "你不是当前物品的运送人，不能调用这个函数");
        _;
    }
    constructor() {}

    // 添加新产品物流信息
    function addProduct(string memory productId, string memory initialLogistics) public {
        require(products[productId].currentOwner == address(0), "物品不存在");

        string[] memory history;
        history[0] = initialLogistics;

        products[productId] = Product(msg.sender, history, false);
        emit LogisticsUpdated(productId, msg.sender, initialLogistics);
    }

    // 更新当前物品的物流信息
    function updateLogistics(string memory productId, string memory newLogistics) public {
        require(products[productId].currentOwner == msg.sender, "只有当前的运送者才能更新该物品物流信息");
        require(!products[productId].isVerified, "不能更新以及被验证过的物品物流信息");

        products[productId].logisticsHistory.push(newLogistics);
        emit LogisticsUpdated(productId, msg.sender, newLogistics);
    }

    // 验证产品物流
    function verifyProduct(string memory productId) public {
        require(products[productId].currentOwner == msg.sender, "只有当前的运送者才能验证该物品");
        require(!products[productId].isVerified, "物品已经被验证过");

        products[productId].isVerified = true;
    }

    //获取当前物品从开始到结束的所有轨迹记录
    function getLogisticsHistory(string memory productId) public view returns (string[] memory) {
        return products[productId].logisticsHistory;
    }
}

//addProduct: 添加新的产品到合约中，指定初始的物流信息。
//updateLogistics: 仅允许当前产品持有者更新物流信息。
//verifyProduct: 验证产品的物流信息，只有当前持有者可以进行验证。
