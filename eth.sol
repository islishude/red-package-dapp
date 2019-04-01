pragma solidity ^0.5.0;

contract ERC20Interface {
    function totalSupply() public view returns (uint);
    function balanceOf(address tokenOwner) public view returns (uint balance);
    function allowance(address tokenOwner, address spender) public view returns (uint remaining);
    function transfer(address to, uint tokens) public returns (bool success);
    function approve(address spender, uint tokens) public returns (bool success);
    function transferFrom(address from, address to, uint tokens) public returns (bool success);

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}


contract Ownable {
    address payable public owner;
    
    constructor() public {
        owner = msg.sender;
    }
    
    modifier OnlyOwner {
        require(msg.sender == owner);
        _;
    }
    
    function transferOwnership(address payable _owner) public OnlyOwner {
        owner = _owner;
    }
}

contract Friendship is Ownable {
    address payable[] internal friends;
    mapping(address => bool) public friendship;
    mapping(address => bool) public blacklist; 
    
    function MyFriends() public view returns (address payable[] memory) {
        return friends;
    }
    
    function AddFriend(address payable _friend) public OnlyOwner {
        if(friendship[_friend]){
            return;
        }
        friends.push(_friend);
    }
    
    function AddFriendList(address payable[] memory list) public OnlyOwner {
        for(uint i = 0; i < list.length; ++i){
            address payable cur = list[i];
            if(friendship[cur]){
                continue;
            }
            friends.push(cur);
            friendship[cur] = true;
        }
    }
    
    function DelFriend(address payable _friend) public OnlyOwner{
        if(!friendship[_friend]){
           return;
        }
        delete friendship[_friend];
        for(uint i = 0; i < friends.length; ++i){
            if(_friend != friends[i]){
                continue;
            }
            if(i == friends.length){
                friends.length--;
                return;
            }
            friends[i] = friends[friends.length-1];
            friends.length--;
        }
    }
    
    modifier OnlyUser {
        require(!blacklist[msg.sender]);
        _;
    }
    
    function LockUser(address _user) public OnlyOwner {
        blacklist[_user] = true;
    }
    
    function UnlockUser(address _user) public OnlyOwner {
        // blacklist[_user] = false; // 20960
        delete blacklist[_user]; // 20872
    }
}

contract Tokens {
    
}

contract Words is Friendship {
    struct Record {
       bool equalDivision;
       bool onlyFriend;
       address token; // 0x0 for ETH
       uint256 value;
       uint256 size;
       uint256 timestamp;
    }

    mapping(string => Record) records;
    
    function Giving(string memory passwd,
                    bool equalDivision,
                    bool onlyFriend,
                    address token,
                    uint256 value,
                    uint256 size) public OnlyOwner {
        require(value > 0 && size > 0 && value >= size, "giving value and package size should be greater than 0");
        if(equalDivision){
            require(value % size == 0);
        }
        if(token == address(0x0)){
            require(address(this).balance >= value, "not sufficient funds");
        } else {
            require(ERC20Interface(token).balanceOf(address(this))>=value, "not sufficient funds");
        }
        records[passwd] =  Record(onlyFriend, equalDivision, token, value, size, block.timestamp);
    }
    
    function Revoke(string memory passwd) public OnlyOwner {
        delete records[passwd];
    }
    
    function Grabbing(string memory passwd) public {
        require(!blacklist[msg.sender]);
        Record storage r = records[passwd];
        require(r.value != 0 && r.timestamp + 1 days <= block.timestamp, "invalid red package password or expired");
        if(r.onlyFriend) {
            require(friendship[msg.sender], "only friend can grab this red package");
        }
        
        // uint256 value = r.equalDivision ? r.value / r.size : 0;
        
        // TODO
        if(r.token == address(0x0)){
            msg.sender.transfer(0);
        } else {
            ERC20Interface(r.token).transfer(msg.sender, 0);
        }
    }
}

contract RedPacket is Words {
    constructor() public payable {}
    
    function() external payable {}

  
    function destructor() public OnlyOwner {
        selfdestruct(owner);
    }
}

