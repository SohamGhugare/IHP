// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

contract IHPContract {
    struct Profile {
        string uri;
    }
    
    mapping(uint256 => Profile) private profiles;
    
    function storeProfile(uint256 uhpId, string memory uri) public {
        Profile storage profile = profiles[uhpId];
        profile.uri = uri;
    }
    
    function getProfile(uint256 uhpId) public view returns (string memory) {
        Profile storage profile = profiles[uhpId];
        return (profile.uri);
    }
}