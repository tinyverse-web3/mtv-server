#!/bin/sh

# solc version ^8-9
# npm install @openzeppelin/contracts
# npm install solc
# git clone https://github.com/ethereum/go-ethereum.git
# cd go-ethereum/cmd/abigen
# go build main.go
# cp ./abigen /usr/local/bin/.

solcjs --bin --include-path node_modules/ --base-path . UserStorage.sol
solcjs --abi --include-path node_modules/ --base-path . UserStorage.sol
# solc --abi UserStorage.sol -o .
# solc --bin UserStorage.sol -o .
abigen --bin=UserStorage_sol_UserStorage.bin --abi=UserStorage_sol_UserStorage.abi --pkg=UserStorage --out=UserStorage.go