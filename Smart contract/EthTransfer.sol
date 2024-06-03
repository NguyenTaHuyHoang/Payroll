pragma solidity ^0.8.0;

contract EthTransfer {
    // Hàm để gửi ETH
    function sendEth(address payable recipient) public payable {
        require(msg.value > 0, "Must send some ETH");
        require(address(this).balance > msg.value );
        recipient.transfer(msg.value);
    }
}
