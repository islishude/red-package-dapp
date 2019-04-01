pragma solidity ^0.5.0;

contract Friendship {
    mapping(address => address payable[]) internal friends;
    mapping(address => mapping(address => bool)) public friendship;
    
    function MyFriends() public view returns (address payable[] memory) {
        return friends[msg.sender];
    }
    
    function AddFriend(address payable _friend) public {
        if(friendship[msg.sender][_friend]){
            return;
        }
        friends[msg.sender].push(_friend);
    }
    
    function AddFriendList(address payable[] memory list) public {
        for(uint i = 0; i < list.length; ++i){
            address payable cur = list[i];
            if(friendship[msg.sender][cur]){
                continue;
            }
            friends[msg.sender].push(cur);
            friendship[msg.sender][cur] = true;
        }
    }
    
    function DelFriend(address payable _friend) public {
        if(!friendship[msg.sender][_friend]){
           return;
        }
        delete friendship[msg.sender][_friend];
        for(uint i = 0; i < friends[msg.sender].length; ++i){
            if(_friend != friends[msg.sender][i]){
                continue;
            }
            if(i == friends[msg.sender].length){
                friends[msg.sender].length--;
                return;
            }
            friends[msg.sender][i] = friends[msg.sender][friends[msg.sender].length-1];
            friends[msg.sender].length--;
        }
    }
}

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

// DON'T implements ERC20Interface!
contract TokenReciever {
    constructor() public {}
    
    function transfer(address token, address to, uint value) public {
        ERC20Interface(token).transfer(to, value);
    }
    
    function transferFrom(address token, address from, address to, uint tokens) public {
        ERC20Interface(token).transferFrom(from, to, tokens);
    }
    
    function balanceOf(address token) public view returns (uint256){
        return ERC20Interface(token).balanceOf(address(this));
    }
    
    function allowance(address token, address spender) public view returns (uint256) {
        return ERC20Interface(token).allowance(address(this), spender);
    }
    
    function approve(address token, address spender, uint256 value) public returns (bool) {
        return ERC20Interface(token).approve(spender, value);
    }
}

contract Token {
    mapping(address => TokenReciever) tokens;
    constructor() public {}
    
    function NewTokenReceiver() public {
        require(address(tokens[msg.sender]) == address(0x0), "Has been created receiver address");
        TokenReciever token = new TokenReciever();
        tokens[msg.sender] = token;
    }
    
    function WithdrawToken(address token, address _to, uint256 value) public {
        tokens[msg.sender].transfer(token, _to, value);
    }
}

contract RedPackage is Friendship, Token {
     struct Record {
       address owner;
       bool equalDivision;
       bool onlyFriend;
       address token;
       uint256 amount;
       uint256 remainAmount;
       uint256 size;
       uint256 remainSize;
       uint256 timestamp;
       uint256 expired;
       address payable[] grabbed;
    }
    
    mapping(string => Record) records;
    
    function IsWordExists(string memory word) public view returns (bool) {
         Record memory r = records[word];
         return r.owner != address(0x0);
    }
    
    function Giving(string memory word, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) public payable {
        require(size > 0 && msg.value > 0 && msg.value > size && expireDays > 0, "invalid data provided");
        require(!IsWordExists(word), "Red package exists");
        
        if(equalDivision) {
            require(msg.value % size == 0, "Invalid value and size for equal division mode");
        }
        
        address payable[] memory grabbed;
        records[word] = Record(msg.sender, equalDivision, onlyFriend, address(0x0), msg.value, msg.value, size, size, now, now + expireDays * 1 days, grabbed);
    }
    
    function Giving(string memory word, address token, uint256 value, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) public {
        uint256 balance = tokens[msg.sender].balanceOf(token);
        require(balance >= value, "not sufficient funds");
        require(size > 0 && value > 0 && value > size && expireDays > 0, "invalid data provided");
        require(!IsWordExists(word), "Red package exists");
        
        if(equalDivision) {
            require(value % size == 0, "Invalid value and size for equal division mode");
        }
        
        address payable[] memory grabbed;
        records[word] = Record(msg.sender, equalDivision, onlyFriend, token, value, value, size, size, now, now + expireDays * 1 days, grabbed);
    }
    
    function Grabbing(string memory word) public {
        Record storage r = records[word];
        require(r.owner != address(0x0), "Red package not exists");
        require(r.expired <= now, "Red package expired");
        
        if(r.onlyFriend) {
            require(friendship[r.owner][msg.sender], "Only friend can grab");
        }
        
        for(uint256 i = 0; i < r.grabbed.length; ++i){
            require(msg.sender != r.grabbed[i], "can't grabbed twice");
        }
        
        uint256 value = 0;
        if(r.equalDivision) {
            value = uint256(r.amount) / uint256(r.size);
        } else  if (r.remainSize == 1){
            value = r.remainAmount;
        } else {
            bytes memory entropy = abi.encode(blockhash(block.number-1), msg.sender, r.remainAmount, r.remainSize, now);
            bytes32 random = keccak256(entropy);
            uint256 tmp = uint256(random) % r.remainAmount;
            uint256 max = uint256(r.remainAmount) / uint256(r.remainSize);
            if(tmp == 0) {
                value = 1;
            } else if (tmp > max){
                value = max;
            } else {
                value = tmp;
            }
        }
        
        if(r.token == address(0x0)){
            msg.sender.transfer(value);
        } else {
            tokens[r.owner].transfer(r.token, msg.sender, value);
        }


        r.remainAmount -= value;
        r.remainSize--;
        if(r.remainSize == 0){
            delete records[word];
            return;
        }
        r.grabbed.push(msg.sender);
     }
}
