pragma solidity ^0.5.0;

import {ERC20Interface} from "./ERC20.sol";

contract TokenReciever {
    bool public locked = false;
    address owner;

    constructor() public {
        owner = msg.sender;
    }

    modifier OnlyUnlock {
        require(!locked);
        _;
    }

    modifier OnlyOwner {
        require(msg.sender == owner);
        _;
    }

    function Send(address token, address to, uint value)
        public
        OnlyOwner
        OnlyUnlock
    {
        ERC20Interface(token).transfer(to, value);
    }

    function SendFrom(address token, address from, address to, uint tokens)
        public
        OnlyOwner
        OnlyUnlock
    {
        ERC20Interface(token).transferFrom(from, to, tokens);
    }

    function balanceOf(address token) public view returns (uint256) {
        return ERC20Interface(token).balanceOf(address(this));
    }

    function Unlock() public OnlyOwner {
        locked = false;
    }

    function Lock() public OnlyOwner {
        locked = true;
    }
}

contract Token {
    mapping(address => address) internal tokens;

    constructor() public {}

    function NewTokenReceiver() public {
        require(
            address(tokens[msg.sender]) == address(0x0),
            "Has been created receiver address"
        );
        // TokenReciever token = new TokenReciever();
        tokens[msg.sender] = address(new TokenReciever());
    }

    function WithdrawToken(address token, address _to, uint256 value) public {
        TokenReciever(tokens[msg.sender]).Send(token, _to, value);
    }

    function SendToken(address token, address from, address to, uint256 value)
        internal
    {
        TokenReciever(tokens[from]).Send(token, to, value);
    }

    function Balance(address token) public view returns (uint256) {
        return TokenReciever(tokens[msg.sender]).balanceOf(token);
    }

    function SetLock(address user, bool locked) internal {
        TokenReciever tr = TokenReciever(tokens[user]);
        return locked ? tr.Lock() : tr.Unlock();
    }

    function MyToken() public view returns (address token) {
        return address(tokens[msg.sender]);
    }
}
