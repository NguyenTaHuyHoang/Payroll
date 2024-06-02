// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EthTransfer {
    // Hàm để gửi ETH
    function sendEth(address payable recipient) public payable {
        require(msg.value > 0, "Must send some ETH");
        recipient.transfer(msg.value);
    }
}

const Web3 = require('web3');

// ABI của hợp đồng EthTransfer (lấy từ Remix sau khi compile)
const abi = [
	{
		"inputs": [
			{
				"internalType": "address payable",
				"name": "recipient",
				"type": "address"
			}
		],
		"name": "sendEth",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	}
]

// Địa chỉ hợp đồng sau khi triển khai
const contractAddress = '0x5aB8ea245192F5Db67F3cbDD01d64c4b4976d496'; // Thay bằng địa chỉ hợp đồng của bạn

async function sendEth(recipientAddress, amountInEth) {
    if (typeof window.ethereum !== 'undefined') {
        await window.ethereum.request({ method: 'eth_requestAccounts' });

        const web3 = new Web3(window.ethereum);
        const accounts = await web3.eth.getAccounts();
        const senderAddress = accounts[0];

        const contract = new web3.eth.Contract(abi, contractAddress);

        try {
            const tx = await contract.methods.sendEth(recipientAddress).send({
                from: senderAddress,
                value: web3.utils.toWei(amountInEth, 'ether') // Chuyển đổi số lượng ETH từ ether sang wei
            });

            console.log('Transaction sent:', tx);
        } catch (error) {
            console.error('Transaction failed:', error);
        }
    } else {
        console.log('MetaMask is not installed. Please install MetaMask and try again.');
    }
}

// Sử dụng hàm sendEth
sendEth('0xRecipientAddress', '0.1'); // Chuyển 0.1 ETH tới địa chỉ người nhận
