pragma solidity ^0.4.26;

contract RandomAddressGenerator {
    
    event AddressGenerated(address indexed generatedAddress);

    function generateRandomAddress() public {
        // Create a pseudo-random hash using block attributes and sender address
        bytes32 hash = keccak256(abi.encodePacked(block.timestamp, block.difficulty, msg.sender));
        
        // Convert the hash to an address by taking the lower 160 bits
        address randomAddress = address(uint160(uint256(hash)));
        emit AddressGenerated(randomAddress);
        
    }
}
