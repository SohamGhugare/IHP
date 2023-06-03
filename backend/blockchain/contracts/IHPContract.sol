// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

contract IHPContract {
    struct Profile {
        string uri;
        string name;
    }
    
    mapping(uint256 => Profile) private profiles;
    
    function storeProfile(uint256 uhpId, string memory uri, string memory name) public {
        Profile storage profile = profiles[uhpId];
        profile.uri = uri;
        profile.name = name;
    }
    
    function getProfile(uint256 uhpId) public view returns (string memory, string memory) {
        Profile storage profile = profiles[uhpId];
        return (profile.uri, profile.name);
    }
}