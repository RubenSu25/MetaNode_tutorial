// SPDX-License-Identifier: MIT
pragma solidity ~0.8.0;
contract BeggingContract{
// 合约应包含以下功能：
// 一个 mapping 来记录每个捐赠者的捐赠金额。
// 一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
// 一个 withdraw 函数，允许合约所有者提取所有资金。
// 一个 getDonation 函数，允许查询某个地址的捐赠金额。
// 使用 payable 修饰符和 address.transfer 实现支付和提款。
    mapping(address => uint) beggingmap;
    address public owner;
    constructor(){
        owner = msg.sender;
    }
    modifier onlyOwner(){
        require(msg.sender == owner,"not owner");
        _;
    }


    function donate() public payable{
        beggingmap[msg.sender] += msg.value;
    }

    function totalAmount() public view returns (uint){
        return address(this).balance;
    }
    function withdraw() public payable onlyOwner {
        payable(msg.sender).transfer(address(this).balance);
    }
    function getDonation(address _address) public view returns(uint){
        return beggingmap[_address];
    }
}