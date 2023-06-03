// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

contract DoctorContract {
    struct Profile {
        string name;
        int phone;
        string email;
        string password;

    }
    
    mapping(uint256 => Profile) private profiles;
    
    function storeProfile(uint256 license, string memory name, string memory email, string memory password) public {
        Profile storage profile = profiles[license];
        profile.name = name;
        profile.email = email;
        profile.password = password;
    }
    
    function getProfile(uint256 license) public view returns (string memory, string memory, string memory) {
        Profile storage profile = profiles[license];
        return (profile.name, profile.email, profile.password);
    }
}
