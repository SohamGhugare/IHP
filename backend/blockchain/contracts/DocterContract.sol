// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

contract DoctorContract {
    struct Profile {
        string name;
        uint256 phone;
        string email;
        uint256 pin; // needs encryption

    }
    
    mapping(uint256 => Profile) private profiles;
    
    function storeProfile(uint256 license, string memory name, string memory email, uint256 pin) public {
        Profile storage profile = profiles[license];
        profile.name = name;
        profile.email = email;
        profile.pin = pin;
    }
    
    function getProfile(uint256 license) public view returns (string memory, string memory, uint256) {
        Profile storage profile = profiles[license];
        return (profile.name, profile.email, profile.pin);
    }
}
