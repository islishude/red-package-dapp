pragma solidity ^0.5.0;

contract Friendship {
    mapping(address => address[]) internal friends;
    mapping(address => mapping(address => bool)) public friendship;

    constructor() public {}

    function MyFriends() public view returns (address[] memory) {
        return friends[msg.sender];
    }

    function AddFriend(address _friend) public {
        require(_friend != msg.sender, "Can't add yourself to friend list");
        require(!friendship[msg.sender][_friend], "Friend already added");
        friends[msg.sender].push(_friend);
        friendship[msg.sender][_friend] = true;
    }

    function AddFriendList(address[] memory list) public {
        for (uint i = 0; i < list.length; ++i) {
            address cur = list[i];
            if (friendship[msg.sender][cur] || cur == msg.sender) {
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
