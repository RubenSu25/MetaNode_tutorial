
// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract SimpleERC20 {
    string public constant name = "SimpleToken";
    string public constant symbol = "STK";
    uint8 public constant decimals = 18;
    
    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;
    
    uint256 private _totalSupply;
    address public owner;
    
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }
    
    constructor(uint256 initialSupply) {
        owner = msg.sender;
        _mint(msg.sender, initialSupply);
    }
    
    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }
    
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }
    
    function transfer(address to, uint256 amount) public returns (bool) {
        _transfer(msg.sender, to, amount);
        return true;
    }
    
    function approve(address spender, uint256 amount) public returns (bool) {
        _approve(msg.sender, spender, amount);
        return true;
    }
    
    function allowance(address ownerAddr, address spender) public view returns (uint256) {
        return _allowances[ownerAddr][spender];
    }
    
    function transferFrom(address from, address to, uint256 amount) public returns (bool) {
        _spendAllowance(from, msg.sender, amount);
        _transfer(from, to, amount);
        return true;
    }
    
    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }
    
    function _transfer(address from, address to, uint256 amount) internal {
        require(from != address(0), "ERC20: transfer from the zero address");
        require(to != address(0), "ERC20: transfer to the zero address");
        require(_balances[from] >= amount, "ERC20: transfer amount exceeds balance");
        
        _balances[from] -= amount;
        _balances[to] += amount;
        emit Transfer(from, to, amount);
    }
    
    function _mint(address to, uint256 amount) internal {
        require(to != address(0), "ERC20: mint to the zero address");
        
        _totalSupply += amount;
        _balances[to] += amount;
        emit Transfer(address(0), to, amount);
    }
    
    function _approve(address ownerAddr, address spender, uint256 amount) internal {
        require(ownerAddr != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");
        
        _allowances[ownerAddr][spender] = amount;
        emit Approval(ownerAddr, spender, amount);
    }
    
    function _spendAllowance(address ownerAddr, address spender, uint256 amount) internal {
        uint256 currentAllowance = allowance(ownerAddr, spender);
        require(currentAllowance >= amount, "ERC20: insufficient allowance");
        
        _approve(ownerAddr, spender, currentAllowance - amount);
    }
}
