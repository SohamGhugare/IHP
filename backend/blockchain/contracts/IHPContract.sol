// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

contract IHPContract {
    struct Profile {
        string uri;
        string name;
        string[] records;
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

    function addRecord(uint256 uhpId, string memory record_cid) public {
        Profile storage profile = profiles[uhpId];
        profile.records.push(record_cid);
    }

    function getRecord(uint256 uhpId, uint256 index) public view returns (string memory){
        require(index < profiles[uhpId].records.length, "Invalid Index");
        return profiles[uhpId].records[index];
    }
}