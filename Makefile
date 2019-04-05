build: lint
	solc --overwrite --abi --bin --optimize -o internal/contract/build internal/contract/contract.sol
	abigen -abi internal/contract/build/RedPackage.abi --bin internal/contract/build/RedPackage.bin --pkg contract --out internal/contract/redpkg.go
lint:
	prettier --write internal/contract/*.sol
clean:
	rm -rf internal/contract/build
	rm internal/contract/redpkg.go
deps:
	brew tap ethereum/ethereum
	brew install solidity ethereum node
	npm install -g prettier prettier-plugin-solidity --registry=https://registry.npm.taobao.org
