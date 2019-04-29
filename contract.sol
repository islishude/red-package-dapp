pragma solidity ^0.5.0;

import "./friendship.sol";
import "./token_receiver.sol";

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
        uint256 id;
    }

    event Send(
        address indexed sender,
        address indexed token,
        uint256 indexed value
    );
    event Receive(
        address indexed receiver,
        address indexed token,
        uint256 indexed value
    );

    mapping(bytes32 => Record) records;
    mapping(uint256 => mapping(address => bool)) grabbed;
    uint256 internal nonce = 0;

    constructor() public {}

    function getRecord(bytes32 word)
        public
        view
        returns (
        address owner,
        bool equalDivision,
        bool onlyFriend,
        address token,
        uint256 amount,
        uint256 remainAmount,
        uint256 size,
        uint256 remainSize,
        uint256 timestamp,
        uint256 expired
    )
    {
        bytes32 newWord = keccak256(abi.encode(word));
        Record memory r = records[newWord];
        return (r.owner, r.equalDivision, r.onlyFriend, r.token, r.amount, r.remainAmount, r.size, r.remainSize, r.timestamp, r.expired);
    }

    function IsWordExists(bytes32 newWord) internal view returns (bool) {
        Record memory r = records[newWord];
        return r.owner != address(0x0);
    }

    // Noitce for app implements
    // If word length is more than 32bytes
    // you can keccak256(word) at first

    // Giving gives out ETH.
    function Giving(
        bytes32 word,
        bool equalDivision,
        bool onlyFriend,
        uint256 size,
        uint256 expireDays
    ) public payable {
        require(
            size > 0 && msg.value > 0 && msg.value > size && expireDays > 0,
            "invalid data provided"
        );

        bytes32 newWord = keccak256(abi.encode(word));
        require(!IsWordExists(newWord), "Red package exists");

        if (equalDivision) {
            require(
                msg.value % size == 0,
                "Invalid value and size for equal division mode"
            );
        }

        records[newWord] = Record(
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
            nonce
        );
        nonce++;
        emit Send(msg.sender, address(0x0), msg.value);
    }

    function Giving(
        bytes32 word,
        address token,
        uint256 value,
        bool equalDivision,
        bool onlyFriend,
        uint256 size,
        uint256 expireDays
    ) public {
        uint256 balance = TokenReciever(tokens[msg.sender]).balanceOf(token);
        require(balance >= value, "not sufficient funds");
        require(
            size > 0 && value > 0 && value > size && expireDays > 0,
            "invalid data provided"
        );
        bytes32 newWord = keccak256(abi.encode(word));
        require(!IsWordExists(newWord), "Red package exists");

        if (equalDivision) {
            require(
                value % size == 0,
                "Invalid value and size for equal division mode"
            );
        }

        records[newWord] = Record(
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
            nonce
        );
        nonce++;
        SetLock(msg.sender, true);
        emit Send(msg.sender, token, value);
    }

    function Revoke(bytes32 word) public {
        bytes32 newWord = keccak256(abi.encode(word));
        Record storage r = records[newWord];
        require(
            r.owner == msg.sender,
            "Red package not exists or you're not the owner"
        );
        require(r.expired < block.timestamp, "Only revoke expired one");
        if (r.token == address(0x0)) {
            msg.sender.transfer(r.amount);
        } else {
            SetLock(msg.sender, false);
        }
        delete records[newWord];
    }

    function CanGrab(bytes32 word) public view returns (bool has) {
        bytes32 newWord = keccak256(abi.encode(word));
        Record storage r = records[newWord];
        require(r.owner != address(0x0), "Red package not exists");
        require(r.expired >= block.timestamp, "Red package expired");

        if (r.onlyFriend) {
            require(friendship[r.owner][msg.sender], "Only friend can grab");
        }
        return !grabbed[r.id][msg.sender];
    }

    function Grabbing(bytes32 word) public {
        bytes32 newWord = keccak256(abi.encode(word));
        Record storage r = records[newWord];
        require(r.owner != address(0x0), "Red package not exists");
        require(r.expired >= block.timestamp, "Red package expired");

        if (r.onlyFriend) {
            require(friendship[r.owner][msg.sender], "Only friend can grab");
        }

        require(!grabbed[r.id][msg.sender], "can't grabbed twice");

        uint256 value = 0;
        if (r.equalDivision) {
            value = uint256(r.amount) / uint256(r.size);
        } else if (r.remainSize == 1) {
            value = r.remainAmount;
        } else {
            bytes memory entropy = abi.encode(
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

        if (r.token == address(0x0)) {
            msg.sender.transfer(value);
        } else if (r.owner != msg.sender) {
            SendToken(r.token, r.owner, msg.sender, value);
        }

        r.remainAmount -= value;
        r.remainSize--;
        if (r.remainSize == 0) {
            delete records[newWord];
            if (r.token != address(0x0)) {
                SetLock(r.owner, false);
            }
        } else {
            grabbed[r.id][msg.sender] = true;
        }
        emit Receive(msg.sender, r.token, value);
    }
}
