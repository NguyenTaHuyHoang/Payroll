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
sendEth('0x3942B0Dd1B21D206512216B482697dE4915Fd72F', '0.1'); // Chuyển 0.1 ETH tới địa chỉ người nhận

async function sendETH(_to, _amount, callback) {
    // Lấy instance của contract
    var contractInstance = web3.eth.contract(abi).at(contractAddress);

    // Gọi hàm chuyển tiền trên contract
    contractInstance.transferEth(_to, _amount, function(error, result) {
        if (!error) {
            // Gọi callback nếu không có lỗi
            callback(null, result);
        } else {
            // Gọi callback với lỗi nếu có lỗi xảy ra
            callback(error, null);
        }
    });
}
// Gọi hàm transferEth của smart contract để chuyển Ether
web3.eth.getAccounts(function(err, accounts) {
    if (err) {
        console.error('Error fetching accounts:', err);
        return;
    }
    var owner = accounts[0]; // Địa chỉ của chủ sở hữu hợp đồng
    contractInstance.transferEth.sendEth(toAddress, amount, { from: owner, value: amount, gas: 200000 }, function(error, transactionHash) {
        if (!error) {
            console.log('Transaction hash:', transactionHash);
        } else {
            console.error('Transaction error:', error);
        }
    });
});
