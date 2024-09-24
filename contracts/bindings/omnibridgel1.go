// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// OmniBridgeL1MetaData contains all meta data concerning the OmniBridgeL1 contract.
var OmniBridgeL1MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ACTION_BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ACTION_WITHDRAW\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KeyPauseAll\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XCALL_WITHDRAW_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"bridgeFee\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Bridge\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516200177238038062001772833981016040819052610031916100fe565b6001600160a01b03811660805261004661004c565b5061012e565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561009c5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100fb5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60006020828403121561011057600080fd5b81516001600160a01b038116811461012757600080fd5b9392505050565b60805161160c62000166600039600081816103a60152818161048901528181610a6f01528181610eb1015261106b015261160c6000f3fe6080604052600436106101095760003560e01c80638456cb5911610095578063c3de453d11610064578063c3de453d14610321578063ed56531a14610334578063f2fde38b14610354578063f3fef3a314610374578063fc0c546a1461039457600080fd5b80638456cb591461027d5780638da5cb5b146102925780638fdcb4c9146102cf578063a10ac97a146102ff57600080fd5b80633794999d116100dc5780633794999d146101db57806339acf9f1146101fb5780633f4ba83a14610233578063485cc95514610248578063715018a61461026857600080fd5b806309839a931461010e578063241b71bb1461015557806325d70f78146101855780632f4dae9f146101b9575b600080fd5b34801561011a57600080fd5b506101427f0683d1c283a672fc58eb7940a0dba83ea98b96966a9ca1b030dec2c60cea4d1e81565b6040519081526020015b60405180910390f35b34801561016157600080fd5b506101756101703660046112f9565b6103c8565b604051901515815260200161014c565b34801561019157600080fd5b506101427f855511cc3694f64379908437d6d64458dc76d02482052bfb8a5b33a72c054c7781565b3480156101c557600080fd5b506101d96101d43660046112f9565b6103d9565b005b3480156101e757600080fd5b506101426101f6366004611327565b6103ed565b34801561020757600080fd5b5060005461021b906001600160a01b031681565b6040516001600160a01b03909116815260200161014c565b34801561023f57600080fd5b506101d96105bf565b34801561025457600080fd5b506101d9610263366004611368565b6105d1565b34801561027457600080fd5b506101d96106fd565b34801561028957600080fd5b506101d961070f565b34801561029e57600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661021b565b3480156102db57600080fd5b506102e66201388081565b60405167ffffffffffffffff909116815260200161014c565b34801561030b57600080fd5b506101426000805160206115b783398151915281565b6101d961032f3660046113a1565b61071f565b34801561034057600080fd5b506101d961034f3660046112f9565b6107a0565b34801561036057600080fd5b506101d961036f3660046113cd565b6107b1565b34801561038057600080fd5b506101d961038f3660046113a1565b6107ec565b3480156103a057600080fd5b5061021b7f000000000000000000000000000000000000000000000000000000000000000081565b60006103d382610b26565b92915050565b6103e1610ba5565b6103ea81610c00565b50565b600080546040805163110ff5f160e01b815290516001600160a01b0390921691638dd9523c91839163110ff5f1916004808201926020929091908290030181865afa158015610440573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104649190611407565b6040516370a0823160e01b815230600482015287908790879081906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa1580156104d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f49190611422565b6104fe919061143b565b6040516001600160a01b0394851660248201529390921660448401526064830152608482015260a40160408051601f198184030181529181526020820180516001600160e01b0316631effa54360e21b179052516001600160e01b031960e085901b16815261057692919062013880906004016114a2565b602060405180830381865afa158015610593573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b79190611422565b949350505050565b6105c7610ba5565b6105cf610cb8565b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156106175750825b905060008267ffffffffffffffff1660011480156106345750303b155b905081158015610642575080155b156106605760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561068a57845460ff60401b1916600160401b1785555b61069387610ccf565b600080546001600160a01b0319166001600160a01b03881617905583156106f457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b610705610ba5565b6105cf6000610ce0565b610717610ba5565b6105cf610d51565b7f0683d1c283a672fc58eb7940a0dba83ea98b96966a9ca1b030dec2c60cea4d1e61074981610b26565b156107905760405162461bcd60e51b815260206004820152601260248201527113db5b9a509c9a5919d94e881c185d5cd95960721b60448201526064015b60405180910390fd5b61079b338484610d68565b505050565b6107a8610ba5565b6103ea816111f0565b6107b9610ba5565b6001600160a01b0381166107e357604051631e4fbdf760e01b815260006004820152602401610787565b6103ea81610ce0565b7f855511cc3694f64379908437d6d64458dc76d02482052bfb8a5b33a72c054c7761081681610b26565b156108585760405162461bcd60e51b815260206004820152601260248201527113db5b9a509c9a5919d94e881c185d5cd95960721b6044820152606401610787565b6000805460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa1580156108a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c491906114d9565b6000549091506001600160a01b031633146109195760405162461bcd60e51b815260206004820152601560248201527413db5b9a509c9a5919d94e881b9bdd081e18d85b1b605a1b6044820152606401610787565b60208101516001600160a01b0316600262048789608a1b01146109775760405162461bcd60e51b81526020600482015260166024820152754f6d6e694272696467653a206e6f742062726964676560501b6044820152606401610787565b60008054906101000a90046001600160a01b03166001600160a01b031663110ff5f16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ec9190611407565b67ffffffffffffffff16816000015167ffffffffffffffff1614610a495760405162461bcd60e51b81526020600482015260146024820152734f6d6e694272696467653a206e6f74206f6d6e6960601b6044820152606401610787565b60405163a9059cbb60e01b81526001600160a01b038581166004830152602482018590527f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303816000875af1158015610ab8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610adc9190611544565b50836001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a942436484604051610b1891815260200190565b60405180910390a250505050565b6000805160206115b783398151915260009081527fff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd69340060208190527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff1680610b9e575060008381526020829052604090205460ff165b9392505050565b33610bd77f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146105cf5760405163118cdaa760e01b8152336004820152602401610787565b60008181527fff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400602081905260409091205460ff16610c775760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610787565b600082815260208290526040808220805460ff191690555183917fd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a91a25050565b6105cf6000805160206115b7833981519152610c00565b610cd76112a8565b6103ea816112f1565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b6105cf6000805160206115b78339815191526111f0565b60008111610db85760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e694272696467653a20616d6f756e74206d757374206265203e203000006044820152606401610787565b6001600160a01b038216610e0e5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e694272696467653a206e6f2062726964676520746f207a65726f0000006044820152606401610787565b60008060009054906101000a90046001600160a01b03166001600160a01b031663110ff5f16040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e62573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e869190611407565b6040516370a0823160e01b815230600482015290915060009085908590859081906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015610ef8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f1c9190611422565b610f26919061143b565b6040516001600160a01b0394851660248201529390921660448401526064830152608482015260a40160408051601f198184030181529181526020820180516001600160e01b0316631effa54360e21b1790526000549051632376548f60e21b81529192506001600160a01b031690638dd9523c90610fb0908590859062013880906004016114a2565b602060405180830381865afa158015610fcd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff19190611422565b341461103f5760405162461bcd60e51b815260206004820152601960248201527f4f6d6e694272696467653a20696e636f727265637420666565000000000000006044820152606401610787565b6040516323b872dd60e01b81526001600160a01b038681166004830152306024830152604482018590527f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af11580156110b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d89190611544565b6111245760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e694272696467653a207472616e73666572206661696c656400000000006044820152606401610787565b60005460405163c21dda4f60e01b81526001600160a01b039091169063c21dda4f90349061116a908690600490600262048789608a1b0190889062013880908401611566565b6000604051808303818588803b15801561118357600080fd5b505af1158015611197573d6000803e3d6000fd5b5050505050836001600160a01b0316856001600160a01b03167f59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422856040516111e191815260200190565b60405180910390a35050505050565b60008181527fff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400602081905260409091205460ff16156112645760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610787565b600082815260208290526040808220805460ff191660011790555183917f0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba3191a25050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166105cf57604051631afcd79f60e31b815260040160405180910390fd5b6107b96112a8565b60006020828403121561130b57600080fd5b5035919050565b6001600160a01b03811681146103ea57600080fd5b60008060006060848603121561133c57600080fd5b833561134781611312565b9250602084013561135781611312565b929592945050506040919091013590565b6000806040838503121561137b57600080fd5b823561138681611312565b9150602083013561139681611312565b809150509250929050565b600080604083850312156113b457600080fd5b82356113bf81611312565b946020939093013593505050565b6000602082840312156113df57600080fd5b8135610b9e81611312565b805167ffffffffffffffff8116811461140257600080fd5b919050565b60006020828403121561141957600080fd5b610b9e826113ea565b60006020828403121561143457600080fd5b5051919050565b808201808211156103d357634e487b7160e01b600052601160045260246000fd5b6000815180845260005b8181101561148257602081850181015186830182015201611466565b506000602082860101526020601f19601f83011685010191505092915050565b600067ffffffffffffffff8086168352606060208401526114c6606084018661145c565b9150808416604084015250949350505050565b6000604082840312156114eb57600080fd5b6040516040810181811067ffffffffffffffff8211171561151c57634e487b7160e01b600052604160045260246000fd5b604052611528836113ea565b8152602083015161153881611312565b60208201529392505050565b60006020828403121561155657600080fd5b81518015158114610b9e57600080fd5b600067ffffffffffffffff808816835260ff8716602084015260018060a01b038616604084015260a060608401526115a160a084018661145c565b9150808416608084015250969550505050505056fe76e8952e4b09b8d505aa08998d716721a1dbf0884ac74202e33985da1ed005e9a264697066735822122006323b706a013abf28420cff6faf5ae1928abfb2d32256217dbd10a5da3f6ece64736f6c63430008180033",
}

// OmniBridgeL1ABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniBridgeL1MetaData.ABI instead.
var OmniBridgeL1ABI = OmniBridgeL1MetaData.ABI

// OmniBridgeL1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniBridgeL1MetaData.Bin instead.
var OmniBridgeL1Bin = OmniBridgeL1MetaData.Bin

// DeployOmniBridgeL1 deploys a new Ethereum contract, binding an instance of OmniBridgeL1 to it.
func DeployOmniBridgeL1(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address) (common.Address, *types.Transaction, *OmniBridgeL1, error) {
	parsed, err := OmniBridgeL1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniBridgeL1Bin), backend, token_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniBridgeL1{OmniBridgeL1Caller: OmniBridgeL1Caller{contract: contract}, OmniBridgeL1Transactor: OmniBridgeL1Transactor{contract: contract}, OmniBridgeL1Filterer: OmniBridgeL1Filterer{contract: contract}}, nil
}

// OmniBridgeL1 is an auto generated Go binding around an Ethereum contract.
type OmniBridgeL1 struct {
	OmniBridgeL1Caller     // Read-only binding to the contract
	OmniBridgeL1Transactor // Write-only binding to the contract
	OmniBridgeL1Filterer   // Log filterer for contract events
}

// OmniBridgeL1Caller is an auto generated read-only Go binding around an Ethereum contract.
type OmniBridgeL1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniBridgeL1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniBridgeL1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniBridgeL1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniBridgeL1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniBridgeL1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniBridgeL1Session struct {
	Contract     *OmniBridgeL1     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniBridgeL1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniBridgeL1CallerSession struct {
	Contract *OmniBridgeL1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OmniBridgeL1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniBridgeL1TransactorSession struct {
	Contract     *OmniBridgeL1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OmniBridgeL1Raw is an auto generated low-level Go binding around an Ethereum contract.
type OmniBridgeL1Raw struct {
	Contract *OmniBridgeL1 // Generic contract binding to access the raw methods on
}

// OmniBridgeL1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniBridgeL1CallerRaw struct {
	Contract *OmniBridgeL1Caller // Generic read-only contract binding to access the raw methods on
}

// OmniBridgeL1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniBridgeL1TransactorRaw struct {
	Contract *OmniBridgeL1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniBridgeL1 creates a new instance of OmniBridgeL1, bound to a specific deployed contract.
func NewOmniBridgeL1(address common.Address, backend bind.ContractBackend) (*OmniBridgeL1, error) {
	contract, err := bindOmniBridgeL1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1{OmniBridgeL1Caller: OmniBridgeL1Caller{contract: contract}, OmniBridgeL1Transactor: OmniBridgeL1Transactor{contract: contract}, OmniBridgeL1Filterer: OmniBridgeL1Filterer{contract: contract}}, nil
}

// NewOmniBridgeL1Caller creates a new read-only instance of OmniBridgeL1, bound to a specific deployed contract.
func NewOmniBridgeL1Caller(address common.Address, caller bind.ContractCaller) (*OmniBridgeL1Caller, error) {
	contract, err := bindOmniBridgeL1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1Caller{contract: contract}, nil
}

// NewOmniBridgeL1Transactor creates a new write-only instance of OmniBridgeL1, bound to a specific deployed contract.
func NewOmniBridgeL1Transactor(address common.Address, transactor bind.ContractTransactor) (*OmniBridgeL1Transactor, error) {
	contract, err := bindOmniBridgeL1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1Transactor{contract: contract}, nil
}

// NewOmniBridgeL1Filterer creates a new log filterer instance of OmniBridgeL1, bound to a specific deployed contract.
func NewOmniBridgeL1Filterer(address common.Address, filterer bind.ContractFilterer) (*OmniBridgeL1Filterer, error) {
	contract, err := bindOmniBridgeL1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1Filterer{contract: contract}, nil
}

// bindOmniBridgeL1 binds a generic wrapper to an already deployed contract.
func bindOmniBridgeL1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniBridgeL1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniBridgeL1 *OmniBridgeL1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniBridgeL1.Contract.OmniBridgeL1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniBridgeL1 *OmniBridgeL1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.OmniBridgeL1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniBridgeL1 *OmniBridgeL1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.OmniBridgeL1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniBridgeL1 *OmniBridgeL1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniBridgeL1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniBridgeL1 *OmniBridgeL1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniBridgeL1 *OmniBridgeL1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.contract.Transact(opts, method, params...)
}

// ACTIONBRIDGE is a free data retrieval call binding the contract method 0x09839a93.
//
// Solidity: function ACTION_BRIDGE() view returns(bytes32)
func (_OmniBridgeL1 *OmniBridgeL1Caller) ACTIONBRIDGE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "ACTION_BRIDGE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACTIONBRIDGE is a free data retrieval call binding the contract method 0x09839a93.
//
// Solidity: function ACTION_BRIDGE() view returns(bytes32)
func (_OmniBridgeL1 *OmniBridgeL1Session) ACTIONBRIDGE() ([32]byte, error) {
	return _OmniBridgeL1.Contract.ACTIONBRIDGE(&_OmniBridgeL1.CallOpts)
}

// ACTIONBRIDGE is a free data retrieval call binding the contract method 0x09839a93.
//
// Solidity: function ACTION_BRIDGE() view returns(bytes32)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) ACTIONBRIDGE() ([32]byte, error) {
	return _OmniBridgeL1.Contract.ACTIONBRIDGE(&_OmniBridgeL1.CallOpts)
}

// ACTIONWITHDRAW is a free data retrieval call binding the contract method 0x25d70f78.
//
// Solidity: function ACTION_WITHDRAW() view returns(bytes32)
func (_OmniBridgeL1 *OmniBridgeL1Caller) ACTIONWITHDRAW(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "ACTION_WITHDRAW")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACTIONWITHDRAW is a free data retrieval call binding the contract method 0x25d70f78.
//
// Solidity: function ACTION_WITHDRAW() view returns(bytes32)
func (_OmniBridgeL1 *OmniBridgeL1Session) ACTIONWITHDRAW() ([32]byte, error) {
	return _OmniBridgeL1.Contract.ACTIONWITHDRAW(&_OmniBridgeL1.CallOpts)
}

// ACTIONWITHDRAW is a free data retrieval call binding the contract method 0x25d70f78.
//
// Solidity: function ACTION_WITHDRAW() view returns(bytes32)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) ACTIONWITHDRAW() ([32]byte, error) {
	return _OmniBridgeL1.Contract.ACTIONWITHDRAW(&_OmniBridgeL1.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_OmniBridgeL1 *OmniBridgeL1Caller) KeyPauseAll(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "KeyPauseAll")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_OmniBridgeL1 *OmniBridgeL1Session) KeyPauseAll() ([32]byte, error) {
	return _OmniBridgeL1.Contract.KeyPauseAll(&_OmniBridgeL1.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) KeyPauseAll() ([32]byte, error) {
	return _OmniBridgeL1.Contract.KeyPauseAll(&_OmniBridgeL1.CallOpts)
}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_OmniBridgeL1 *OmniBridgeL1Caller) XCALLWITHDRAWGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "XCALL_WITHDRAW_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_OmniBridgeL1 *OmniBridgeL1Session) XCALLWITHDRAWGASLIMIT() (uint64, error) {
	return _OmniBridgeL1.Contract.XCALLWITHDRAWGASLIMIT(&_OmniBridgeL1.CallOpts)
}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) XCALLWITHDRAWGASLIMIT() (uint64, error) {
	return _OmniBridgeL1.Contract.XCALLWITHDRAWGASLIMIT(&_OmniBridgeL1.CallOpts)
}

// BridgeFee is a free data retrieval call binding the contract method 0x3794999d.
//
// Solidity: function bridgeFee(address payor, address to, uint256 amount) view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1Caller) BridgeFee(opts *bind.CallOpts, payor common.Address, to common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "bridgeFee", payor, to, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeFee is a free data retrieval call binding the contract method 0x3794999d.
//
// Solidity: function bridgeFee(address payor, address to, uint256 amount) view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1Session) BridgeFee(payor common.Address, to common.Address, amount *big.Int) (*big.Int, error) {
	return _OmniBridgeL1.Contract.BridgeFee(&_OmniBridgeL1.CallOpts, payor, to, amount)
}

// BridgeFee is a free data retrieval call binding the contract method 0x3794999d.
//
// Solidity: function bridgeFee(address payor, address to, uint256 amount) view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) BridgeFee(payor common.Address, to common.Address, amount *big.Int) (*big.Int, error) {
	return _OmniBridgeL1.Contract.BridgeFee(&_OmniBridgeL1.CallOpts, payor, to, amount)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 action) view returns(bool)
func (_OmniBridgeL1 *OmniBridgeL1Caller) IsPaused(opts *bind.CallOpts, action [32]byte) (bool, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "isPaused", action)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 action) view returns(bool)
func (_OmniBridgeL1 *OmniBridgeL1Session) IsPaused(action [32]byte) (bool, error) {
	return _OmniBridgeL1.Contract.IsPaused(&_OmniBridgeL1.CallOpts, action)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 action) view returns(bool)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) IsPaused(action [32]byte) (bool, error) {
	return _OmniBridgeL1.Contract.IsPaused(&_OmniBridgeL1.CallOpts, action)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Caller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Session) Omni() (common.Address, error) {
	return _OmniBridgeL1.Contract.Omni(&_OmniBridgeL1.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) Omni() (common.Address, error) {
	return _OmniBridgeL1.Contract.Omni(&_OmniBridgeL1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Session) Owner() (common.Address, error) {
	return _OmniBridgeL1.Contract.Owner(&_OmniBridgeL1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) Owner() (common.Address, error) {
	return _OmniBridgeL1.Contract.Owner(&_OmniBridgeL1.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Session) Token() (common.Address, error) {
	return _OmniBridgeL1.Contract.Token(&_OmniBridgeL1.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) Token() (common.Address, error) {
	return _OmniBridgeL1.Contract.Token(&_OmniBridgeL1.CallOpts)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Bridge(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "bridge", to, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Bridge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Bridge(&_OmniBridgeL1.TransactOpts, to, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Bridge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Bridge(&_OmniBridgeL1.TransactOpts, to, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address omni_) returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "initialize", owner_, omni_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address omni_) returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Initialize(owner_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Initialize(&_OmniBridgeL1.TransactOpts, owner_, omni_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address omni_) returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Initialize(owner_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Initialize(&_OmniBridgeL1.TransactOpts, owner_, omni_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Pause() (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Pause(&_OmniBridgeL1.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Pause() (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Pause(&_OmniBridgeL1.TransactOpts)
}

// Pause0 is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 action) returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Pause0(opts *bind.TransactOpts, action [32]byte) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "pause0", action)
}

// Pause0 is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 action) returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Pause0(action [32]byte) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Pause0(&_OmniBridgeL1.TransactOpts, action)
}

// Pause0 is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 action) returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Pause0(action [32]byte) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Pause0(&_OmniBridgeL1.TransactOpts, action)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) RenounceOwnership() (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.RenounceOwnership(&_OmniBridgeL1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.RenounceOwnership(&_OmniBridgeL1.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.TransferOwnership(&_OmniBridgeL1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.TransferOwnership(&_OmniBridgeL1.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 action) returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Unpause(opts *bind.TransactOpts, action [32]byte) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "unpause", action)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 action) returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Unpause(action [32]byte) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Unpause(&_OmniBridgeL1.TransactOpts, action)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 action) returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Unpause(action [32]byte) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Unpause(&_OmniBridgeL1.TransactOpts, action)
}

// Unpause0 is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Unpause0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "unpause0")
}

// Unpause0 is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Unpause0() (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Unpause0(&_OmniBridgeL1.TransactOpts)
}

// Unpause0 is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Unpause0() (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Unpause0(&_OmniBridgeL1.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Withdraw(&_OmniBridgeL1.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Withdraw(&_OmniBridgeL1.TransactOpts, to, amount)
}

// OmniBridgeL1BridgeIterator is returned from FilterBridge and is used to iterate over the raw logs and unpacked data for Bridge events raised by the OmniBridgeL1 contract.
type OmniBridgeL1BridgeIterator struct {
	Event *OmniBridgeL1Bridge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OmniBridgeL1BridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1Bridge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OmniBridgeL1Bridge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OmniBridgeL1BridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1BridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1Bridge represents a Bridge event raised by the OmniBridgeL1 contract.
type OmniBridgeL1Bridge struct {
	Payor  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridge is a free log retrieval operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterBridge(opts *bind.FilterOpts, payor []common.Address, to []common.Address) (*OmniBridgeL1BridgeIterator, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "Bridge", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1BridgeIterator{contract: _OmniBridgeL1.contract, event: "Bridge", logs: logs, sub: sub}, nil
}

// WatchBridge is a free log subscription operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchBridge(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1Bridge, payor []common.Address, to []common.Address) (event.Subscription, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "Bridge", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1Bridge)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "Bridge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridge is a log parse operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseBridge(log types.Log) (*OmniBridgeL1Bridge, error) {
	event := new(OmniBridgeL1Bridge)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "Bridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeL1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OmniBridgeL1 contract.
type OmniBridgeL1InitializedIterator struct {
	Event *OmniBridgeL1Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OmniBridgeL1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OmniBridgeL1Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OmniBridgeL1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1Initialized represents a Initialized event raised by the OmniBridgeL1 contract.
type OmniBridgeL1Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterInitialized(opts *bind.FilterOpts) (*OmniBridgeL1InitializedIterator, error) {

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1InitializedIterator{contract: _OmniBridgeL1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1Initialized) (event.Subscription, error) {

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1Initialized)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseInitialized(log types.Log) (*OmniBridgeL1Initialized, error) {
	event := new(OmniBridgeL1Initialized)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeL1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OmniBridgeL1 contract.
type OmniBridgeL1OwnershipTransferredIterator struct {
	Event *OmniBridgeL1OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OmniBridgeL1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OmniBridgeL1OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OmniBridgeL1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1OwnershipTransferred represents a OwnershipTransferred event raised by the OmniBridgeL1 contract.
type OmniBridgeL1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OmniBridgeL1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1OwnershipTransferredIterator{contract: _OmniBridgeL1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1OwnershipTransferred)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseOwnershipTransferred(log types.Log) (*OmniBridgeL1OwnershipTransferred, error) {
	event := new(OmniBridgeL1OwnershipTransferred)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeL1PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OmniBridgeL1 contract.
type OmniBridgeL1PausedIterator struct {
	Event *OmniBridgeL1Paused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OmniBridgeL1PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1Paused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OmniBridgeL1Paused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OmniBridgeL1PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1Paused represents a Paused event raised by the OmniBridgeL1 contract.
type OmniBridgeL1Paused struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterPaused(opts *bind.FilterOpts, key [][32]byte) (*OmniBridgeL1PausedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "Paused", keyRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1PausedIterator{contract: _OmniBridgeL1.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1Paused, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "Paused", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1Paused)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParsePaused(log types.Log) (*OmniBridgeL1Paused, error) {
	event := new(OmniBridgeL1Paused)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeL1UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OmniBridgeL1 contract.
type OmniBridgeL1UnpausedIterator struct {
	Event *OmniBridgeL1Unpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OmniBridgeL1UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1Unpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OmniBridgeL1Unpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OmniBridgeL1UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1Unpaused represents a Unpaused event raised by the OmniBridgeL1 contract.
type OmniBridgeL1Unpaused struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterUnpaused(opts *bind.FilterOpts, key [][32]byte) (*OmniBridgeL1UnpausedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "Unpaused", keyRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1UnpausedIterator{contract: _OmniBridgeL1.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1Unpaused, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "Unpaused", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1Unpaused)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseUnpaused(log types.Log) (*OmniBridgeL1Unpaused, error) {
	event := new(OmniBridgeL1Unpaused)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeL1WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the OmniBridgeL1 contract.
type OmniBridgeL1WithdrawIterator struct {
	Event *OmniBridgeL1Withdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OmniBridgeL1WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1Withdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OmniBridgeL1Withdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OmniBridgeL1WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1Withdraw represents a Withdraw event raised by the OmniBridgeL1 contract.
type OmniBridgeL1Withdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*OmniBridgeL1WithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1WithdrawIterator{contract: _OmniBridgeL1.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1Withdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1Withdraw)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseWithdraw(log types.Log) (*OmniBridgeL1Withdraw, error) {
	event := new(OmniBridgeL1Withdraw)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
