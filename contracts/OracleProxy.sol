// SPDX-License-Identifier: MIT

pragma solidity ^0.8.10;

contract OracleProxy {

    constructor() {

    }

    event dataRequest(bytes encryptedKey, bytes tau, bytes rSignature);
    event dataResponse(bytes encryptedData, bytes encryptedKey, bytes signature);

    function requireOracleData(bytes calldata encryptedKey, bytes calldata tau, bytes calldata rSignature) external {
        emit dataRequest(encryptedKey, tau, rSignature);
    }

    function respondOracleData(bytes calldata encryptedData, bytes calldata encryptedKey, bytes calldata signature) external {
        emit dataResponse(encryptedData, encryptedKey, signature);
    }

}