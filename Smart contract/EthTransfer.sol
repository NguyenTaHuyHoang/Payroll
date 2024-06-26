// SPDX-License-Identifier: GPL-3.0
pragma solidity >= 0.8.2 <0.9.0;

contract EthTransfer {
    address public owner;

    // Hàm khởi tạo hợp đồng, lưu địa chỉ của người triển khai
    constructor() {
        owner = msg.sender;
    }

    // Hàm để gửi ETH
    function sendEth(address payable recipient) public payable {
        require(msg.value > 0, "Must send some ETH");
        recipient.transfer(msg.value);
    }

    // Hàm để lấy số dư của người triển khai hợp đồng
    function getOwnerBalance() public view returns (uint256){
        return owner.balance;
    }
}
