// SPDX-License-Identifier: MIT

pragma solidity ^0.8.10;

contract OracleProxy {

    constructor() {

    }

    event dataRequest(bytes reqSource);
    event dataResponse(bytes sourceData, bytes reqSource);

    function requireOracleData(bytes calldata reqSource) external {
        emit dataRequest(reqSource);
    }

    function respondOracleData(bytes calldata sourceData, bytes calldata reqSource) external {
        emit dataResponse(sourceData, reqSource);
    }

}