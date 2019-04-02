pragma solidity ^0.5.0;

contract Friendship {
    mapping(address => address[]) internal friends;
    mapping(address => mapping(address => bool)) public friendship;

    constructor() public {}

    function MyFriends() public view returns (address[] memory) {
        return friends[msg.sender];
    }

    function AddFriend(address _friend) public {
        if (friendship[msg.sender][_friend]) {
            return;
        }
        friends[msg.sender].push(_friend);
    }

    function AddFriendList(address[] memory list) public {
        for (uint i = 0; i < list.length; ++i) {
            address cur = list[i];
            if (friendship[msg.sender][cur]) {
                continue;
            }
            friends[msg.sender].push(cur);
            friendship[msg.sender][cur] = true;
        }
    }

    function DelFriend(address _friend) public {
        if (!friendship[msg.sender][_friend]) {
            return;
        }
        delete friendship[msg.sender][_friend];
        for (uint i = 0; i < friends[msg.sender].length; ++i) {
            if (_friend != friends[msg.sender][i]) {
                continue;
            }
            if (i == friends[msg.sender].length) {
                friends[msg.sender].length--;
                return;
            }
            friends[msg.sender][i] = friends[msg.sender][friends[msg.sender].length - 1];
            friends[msg.sender].length--;
        }
    }
}

contract ERC20Interface {
    function totalSupply() public view returns (uint);
    function balanceOf(address tokenOwner) public view returns (uint balance);
    function allowance(address tokenOwner, address spender)
        public
        view
        returns (uint remaining);
    function transfer(address to, uint tokens) public returns (bool success);
    function approve(address spender, uint tokens)
        public
        returns (bool success);
    function transferFrom(address from, address to, uint tokens)
        public
        returns (bool success);

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(
        address indexed tokenOwner,
        address indexed spender,
        uint tokens
    );
}

// DON'T implements ERC20Interface!
contract TokenReciever {
    bool public locked = false;
    address owner;

    constructor() public {
        owner = msg.sender;
    }

    modifier IsLocked {
        require(!locked);
        _;
    }

    function transfer(address token, address to, uint value) public IsLocked {
        ERC20Interface(token).transfer(to, value);
    }

    function transferFrom(address token, address from, address to, uint tokens)
        public
        IsLocked
    {
        require(!locked);
        ERC20Interface(token).transferFrom(from, to, tokens);
    }

    function balanceOf(address token) public view returns (uint256) {
        return ERC20Interface(token).balanceOf(address(this));
    }

    function SetUnlock(bool _locked) public {
        require(msg.sender == owner);
        locked = _locked;
    }
}

contract Token {
    mapping(address => TokenReciever) public tokens;

    function NewTokenReceiver() public {
        require(
            address(tokens[msg.sender]) == address(0x0),
            "Has been created receiver address"
        );
        TokenReciever token = new TokenReciever();
        tokens[msg.sender] = token;
    }

    function WithdrawToken(address token, address _to, uint256 value) public {
        tokens[msg.sender].transfer(token, _to, value);
    }

    function Balance(address token) public view returns (uint256) {
        return tokens[msg.sender].balanceOf(token);
    }

    function SetUnlock(address user, bool locked) internal {
        tokens[user].SetUnlock(locked);
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
        address[] grabbed;
    }

    mapping(string => Record) records;

    constructor() public {}

    function getRecord(string memory word)
        public
        view
        returns (
        address,
        bool,
        bool,
        address,
        uint256,
        uint256,
        uint256,
        uint256,
        uint256,
        uint256,
        address[] memory
    )
    {
        Record memory r = records[word];
        return (r.owner, r.equalDivision, r.onlyFriend, r.token, r.amount, r.remainAmount, r.size, r.remainSize, r.timestamp, r.expired, r.grabbed);
    }

    function IsWordExists(string memory word) public view returns (bool) {
        Record memory r = records[word];
        return r.owner != address(0x0);
    }

    // Giving gives out ETH.
    function Giving(
        string memory word,
        bool equalDivision,
        bool onlyFriend,
        uint256 size,
        uint256 expireDays
    ) public payable {
        require(
            size > 0 && msg.value > 0 && msg.value > size && expireDays > 0,
            "invalid data provided"
        );
        require(!IsWordExists(word), "Red package exists");

        if (equalDivision) {
            require(
                msg.value % size == 0,
                "Invalid value and size for equal division mode"
            );
        }

        address[] memory grabbed;
        records[word] = Record(
            msg.sender,
            equalDivision,
            onlyFriend,
            address(0x0),
            msg.value,
            msg.value,
            size,
            size,
            block.timestamp,
            block.timestamp + expireDays * 1 days,
            grabbed
        );
    }

    function Giving(
        string memory word,
        address token,
        uint256 value,
        bool equalDivision,
        bool onlyFriend,
        uint256 size,
        uint256 expireDays
    ) public {
        uint256 balance = tokens[msg.sender].balanceOf(token);
        require(balance >= value, "not sufficient funds");
        require(
            size > 0 && value > 0 && value > size && expireDays > 0,
            "invalid data provided"
        );
        require(!IsWordExists(word), "Red package exists");

        if (equalDivision) {
            require(
                value % size == 0,
                "Invalid value and size for equal division mode"
            );
        }

        address[] memory grabbed;
        records[word] = Record(
            msg.sender,
            equalDivision,
            onlyFriend,
            token,
            value,
            value,
            size,
            size,
            block.timestamp,
            block.timestamp + expireDays * 1 days,
            grabbed
        );
        SetUnlock(msg.sender, true);
    }

    function Grabbing(string memory word) public {
        Record storage r = records[word];
        require(r.owner != address(0x0), "Red package not exists");
        require(r.expired >= block.timestamp, "Red package expired");

        if (r.onlyFriend) {
            require(friendship[r.owner][msg.sender], "Only friend can grab");
        }

        for (uint256 i = 0; i < r.grabbed.length; ++i) {
            require(msg.sender != r.grabbed[i], "can't grabbed twice");
        }

        uint256 value = 0;
        if (r.equalDivision) {
            value = uint256(r.amount) / uint256(r.size);
        } else if (r.remainSize == 1) {
            value = r.remainAmount;
        } else {
            bytes memory entropy = abi.encode(
                blockhash(block.number - 1),
                msg.sender,
                r.remainAmount,
                r.remainSize,
                block.timestamp
            );
            uint256 val = uint256(keccak256(entropy)) % r.remainAmount;
            uint256 max = uint256(r.remainAmount) / uint256(r.remainSize);
            if (val == 0) {
                value = 1;
            } else if (val > max) {
                value = max;
            } else {
                value = val;
            }
        }

        // owner of giving ETH can grab but owner of giving ERC20 token can't grab
        if (r.token == address(0x0)) {
            msg.sender.transfer(value);
        } else {
            tokens[r.owner].transfer(r.token, msg.sender, value);
        }

        r.remainAmount -= value;
        r.remainSize--;
        if (r.remainSize == 0) {
            delete records[word];
            if (r.token != address(0x0)) {
                SetUnlock(r.owner, false);
            }
            return;
        }
        r.grabbed.push(msg.sender);
    }
}

// Simple ERC20 token for testing
contract AdeToken {
    string public name = "AdeToken";
    string public symbol = "ADT";
    uint8 public decimals = 18;
    uint256 public constant totalSupply = 10 ** 18 * 10 ** 8;

    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    constructor() public {
        balanceOf[msg.sender] = totalSupply;
        emit Transfer(address(0x0), msg.sender, totalSupply);
    }

    function transfer(address _to, uint256 _value)
        public
        returns (bool success)
    {
        require(balanceOf[msg.sender] >= _value);
        require(balanceOf[_to] + _value >= balanceOf[_to]);
        balanceOf[msg.sender] -= _value;
        balanceOf[_to] += _value;
        emit Transfer(msg.sender, _to, _value);
        return true;
    }

    function transferFrom(address _from, address _to, uint256 _value)
        public
        returns (bool success)
    {
        require(balanceOf[_from] >= _value);
        require(balanceOf[_to] + _value >= balanceOf[_to]);
        require(allowance[_from][msg.sender] >= _value);
        balanceOf[_to] += _value;
        balanceOf[_from] -= _value;
        allowance[_from][msg.sender] -= _value;
        emit Transfer(_from, _to, _value);
        return true;
    }

    function approve(address _spender, uint256 _value)
        public
        returns (bool success)
    {
        require(_value == 0 || allowance[msg.sender][_spender] == 0);
        allowance[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value);
        return true;
    }

    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(
        address indexed _owner,
        address indexed _spender,
        uint256 _value
    );
}
