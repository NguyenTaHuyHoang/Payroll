pragma solidity >=0.4.0 <0.6.0;

interface IERC20 {
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    function balanceOf(address account) external view returns (uint256);
}

contract TokenToETHExchange {
    address public admin;
    IERC20 public internalToken; // Địa chỉ của token nội bộ
    
    event Exchanged(address indexed employee, uint256 amountInInternalToken, uint256 amountInETH);

    constructor(address _internalTokenAddress) public {
        admin = msg.sender;
        internalToken = IERC20(_internalTokenAddress);
    }
    
    modifier onlyAdmin() {
        require(msg.sender == admin, "Only admin can perform this action");
        _;
    }

    // Đổi token nội bộ sang ETH
    function exchangeTokenToETH(address employee, uint256 amount) external onlyAdmin {
        // Kiểm tra số token nội bộ có sẵn trong hợp đồng
        require(internalToken.balanceOf(employee) >= amount, "Insufficient internal tokens");

        // Chuyển token nội bộ từ nhân viên tới hợp đồng
        require(internalToken.transferFrom(employee, address(this), amount), "Token transfer failed");

        // Kiểm tra số ETH có sẵn trong hợp đồng
        require(address(this).balance >= amount, "Insufficient ETH balance");

        // Chuyển ETH tương ứng tới ví của nhân viên
        (bool success, ) = employee.call.value(amount)("");
        require(success, "ETH transfer failed");

        emit Exchanged(employee, amount, amount);
    }

    // Cho phép admin nạp ETH vào hợp đồng
    function depositETH() external payable onlyAdmin {}

    // Cho phép admin rút ETH từ hợp đồng
    function withdrawETH(uint256 amount) external onlyAdmin {
        require(address(this).balance >= amount, "Insufficient ETH balance");
        (bool success, ) = admin.call.value(amount)("");
        require(success, "ETH transfer failed");
    }

    // Hàm nhận ETH gửi trực tiếp tới hợp đồng
    function() external payable {}
}
