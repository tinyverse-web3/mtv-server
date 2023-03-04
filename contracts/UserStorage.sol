// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

import "@openzeppelin/contracts/utils/cryptography/SignatureChecker.sol";

contract UserStorage {
    address public owner;
    struct User {
        string dirHash;
    }
    mapping(address => User) userList;

    constructor() {
        owner = msg.sender; 
    }

    event SetUserSucc(address indexed sender, address indexed userAddress);

    modifier onlyOwner {
        require(msg.sender == owner, "Not owner");
        _;
    }

    modifier validAddress(address _addr) {
        require(_addr != address(0), "Not valid address");
        _;
    }

    function getSender() public view returns (address) {
        return msg.sender;
    }

    function changeOwner(address _newOwner) public onlyOwner validAddress(_newOwner) {
        owner = _newOwner;
    }

    function setUser(address _userAddress, string calldata _dirHash, bytes calldata _sig) public onlyOwner {
        bool verifyResult = verify(_userAddress, _dirHash, _sig);
        require(verifyResult, "Verification fail");
        userList[_userAddress] = User(_dirHash);
        emit SetUserSucc(msg.sender, _userAddress);
    }

    function getUser(address _userAddress) public view returns (string memory)  {
        return userList[_userAddress].dirHash;
    }

    function verify(address _userAddress, string calldata _dirHash, bytes calldata _sig) public returns (bool) {
        bytes32 hash = keccak256(abi.encodePacked(_userAddress, _dirHash));
        bool result = _isValidSignatureNow(_userAddress, hash, _sig);
        return result;
    }

    event VertifySignFail(address indexed userAddress, bytes32 hash, bytes sig);
    
    function _isValidSignatureNow(address _signer, bytes32 _hash, bytes calldata _sig) private returns (bool) {
        bool result = SignatureChecker.isValidSignatureNow(_signer, _hash, _sig);
        if (!result) {
            emit VertifySignFail(_signer, _hash, _sig);
        }
        return result;
    }

    fallback() external payable {}

    receive() external payable {}
}
