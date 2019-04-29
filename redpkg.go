// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"word\",\"type\":\"bytes32\"},{\"name\":\"equalDivision\",\"type\":\"bool\"},{\"name\":\"onlyFriend\",\"type\":\"bool\"},{\"name\":\"size\",\"type\":\"uint256\"},{\"name\":\"expireDays\",\"type\":\"uint256\"}],\"name\":\"Giving\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MyToken\",\"outputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_friend\",\"type\":\"address\"}],\"name\":\"AddFriend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"word\",\"type\":\"bytes32\"}],\"name\":\"getRecord\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"equalDivision\",\"type\":\"bool\"},{\"name\":\"onlyFriend\",\"type\":\"bool\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"remainAmount\",\"type\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\"},{\"name\":\"remainSize\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"expired\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"word\",\"type\":\"bytes32\"}],\"name\":\"Revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"Balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"list\",\"type\":\"address[]\"}],\"name\":\"AddFriendList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MyFriends\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_friend\",\"type\":\"address\"}],\"name\":\"DelFriend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"NewTokenReceiver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"friendship\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"word\",\"type\":\"bytes32\"}],\"name\":\"Grabbing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"word\",\"type\":\"bytes32\"}],\"name\":\"CanGrab\",\"outputs\":[{\"name\":\"has\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"word\",\"type\":\"bytes32\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"equalDivision\",\"type\":\"bool\"},{\"name\":\"onlyFriend\",\"type\":\"bool\"},{\"name\":\"size\",\"type\":\"uint256\"},{\"name\":\"expireDays\",\"type\":\"uint256\"}],\"name\":\"Giving\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Send\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Receive\",\"type\":\"event\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
const ContractBin = `6080604052600060055534801561001557600080fd5b5061229a806100256000396000f3fe6080604052600436106100e85760003560e01c80632bba4dde1161008a57806366031fdf1161005957806366031fdf1461042757806386060c02146104765780638749a6fa146104a0578063d3a04056146104ca576100e8565b80632bba4dde146102ca57806337b371601461037a5780635663132e146103df5780635daccf5114610412576100e8565b806313962258116100c6578063139622581461019c578063213681cd146101cf57806322b3f5491461025b578063239fcf0f14610285576100e8565b806303723885146100ed57806306758f0214610132578063089327de1461016b575b600080fd5b3480156100f957600080fd5b506101306004803603606081101561011057600080fd5b506001600160a01b03813581169160208101359091169060400135610525565b005b610130600480360360a081101561014857600080fd5b5080359060208101351515906040810135151590606081013590608001356105a9565b34801561017757600080fd5b506101806108aa565b604080516001600160a01b039092168252519081900360200190f35b3480156101a857600080fd5b50610130600480360360208110156101bf57600080fd5b50356001600160a01b03166108c7565b3480156101db57600080fd5b506101f9600480360360208110156101f257600080fd5b50356109eb565b604080516001600160a01b039b8c16815299151560208b015297151589890152959098166060880152608087019390935260a086019190915260c085015260e08401526101008301939093526101208201929092529051908190036101400190f35b34801561026757600080fd5b506101306004803603602081101561027e57600080fd5b5035610b7f565b34801561029157600080fd5b506102b8600480360360208110156102a857600080fd5b50356001600160a01b0316610d09565b60408051918252519081900360200190f35b3480156102d657600080fd5b50610130600480360360208110156102ed57600080fd5b81019060208101813564010000000081111561030857600080fd5b82018360208201111561031a57600080fd5b8035906020019184602083028401116401000000008311171561033c57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610d94945050505050565b34801561038657600080fd5b5061038f610e69565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103cb5781810151838201526020016103b3565b505050509050019250505060405180910390f35b3480156103eb57600080fd5b506101306004803603602081101561040257600080fd5b50356001600160a01b0316610ed2565b34801561041e57600080fd5b50610130611068565b34801561043357600080fd5b506104626004803603604081101561044a57600080fd5b506001600160a01b0381358116916020013516611118565b604080519115158252519081900360200190f35b34801561048257600080fd5b506101306004803603602081101561049957600080fd5b5035611138565b3480156104ac57600080fd5b50610462600480360360208110156104c357600080fd5b503561158d565b3480156104d657600080fd5b50610130600480360360e08110156104ed57600080fd5b508035906001600160a01b03602082013516906040810135906060810135151590608081013515159060a08101359060c00135611731565b33600090815260026020526040808220548151600160e11b6349f59e310281526001600160a01b038781166004830152868116602483015260448201869052925192909116926393eb3c629260648084019382900301818387803b15801561058c57600080fd5b505af11580156105a0573d6000803e3d6000fd5b50505050505050565b6000821180156105b95750600034115b80156105c457508134115b80156105d05750600081115b6106245760408051600160e51b62461bcd02815260206004820152601560248201527f696e76616c696420646174612070726f76696465640000000000000000000000604482015290519081900360640190fd5b604080516020808201889052825180830382018152918301909252805191012061064d81611b14565b1561069a5760408051600160e51b62461bcd0281526020600482015260126024820152600160701b71526564207061636b6167652065786973747302604482015290519081900360640190fd5b84156106e9578234816106a957fe5b06156106e957604051600160e51b62461bcd02815260040180806020018281038252602e8152602001806121d1602e913960400191505060405180910390fd5b604051806101600160405280336001600160a01b031681526020018615158152602001851515815260200160006001600160a01b03168152602001348152602001348152602001848152602001848152602001428152602001836201518002420181526020016005548152506003600083815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160000160146101000a81548160ff02191690831515021790555060408201518160000160156101000a81548160ff02191690831515021790555060608201518160010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506080820151816002015560a0820151816003015560c0820151816004015560e082015181600501556101008201518160060155610120820151816007015561014082015181600801559050506005600081548092919060010191905055503460006001600160a01b0316336001600160a01b03167f93eb3c629eb575edaf0252e4f9fc0c5ccada50496f8c1d32f0f93a65a8257eb560405160405180910390a4505050505050565b336000908152600260205260409020546001600160a01b03165b90565b6001600160a01b03811633141561091257604051600160e51b62461bcd02815260040180806020018281038252602181526020018061224e6021913960400191505060405180910390fd5b3360009081526001602090815260408083206001600160a01b038516845290915290205460ff161561098e5760408051600160e51b62461bcd02815260206004820152601460248201527f467269656e6420616c7265616479206164646564000000000000000000000000604482015290519081900360640190fd5b336000818152602081815260408083208054600180820183559185528385200180546001600160a01b0319166001600160a01b0397909716968717905593835283825280832094835293905291909120805460ff19169091179055565b60008060008060008060008060008060008b60405160200180828152602001915050604051602081830303815290604052805190602001209050610a2d611d0f565b60036000838152602001908152602001600020604051806101600160405290816000820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016000820160149054906101000a900460ff161515151581526020016000820160159054906101000a900460ff161515151581526020016001820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815250509050806000015181602001518260400151836060015184608001518560a001518660c001518760e001518861010001518961012001519b509b509b509b509b509b509b509b509b509b5050509193959799509193959799565b60408051602080820184905282518083038201815291830183528151918101919091206000818152600390925291902080546001600160a01b03163314610bfa57604051600160e51b62461bcd02815260040180806020018281038252602e8152602001806121ff602e913960400191505060405180910390fd5b42816007015410610c555760408051600160e51b62461bcd02815260206004820152601760248201527f4f6e6c79207265766f6b652065787069726564206f6e65000000000000000000604482015290519081900360640190fd5b60018101546001600160a01b0316610c9d576002810154604051339180156108fc02916000818181858888f19350505050158015610c97573d6000803e3d6000fd5b50610ca8565b610ca8336000611bcf565b506000908152600360208190526040822080546001600160b01b03191681556001810180546001600160a01b031916905560028101839055908101829055600481018290556005810182905560068101829055600781018290556008015550565b336000908152600260209081526040808320548151600160e01b6370a082310281526001600160a01b038681166004830152925192909116926370a0823192602480840193829003018186803b158015610d6257600080fd5b505afa158015610d76573d6000803e3d6000fd5b505050506040513d6020811015610d8c57600080fd5b505192915050565b60005b8151811015610e65576000828281518110610dae57fe5b6020908102919091018101513360009081526001835260408082206001600160a01b0384168352909352919091205490915060ff1680610df657506001600160a01b03811633145b15610e015750610e5d565b336000818152602081815260408083208054600180820183559185528385200180546001600160a01b0319166001600160a01b0397909716968717905593835283825280832094835293905291909120805460ff191690911790555b600101610d97565b5050565b3360009081526020818152604091829020805483518184028101840190945280845260609392830182828015610ec857602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610eaa575b5050505050905090565b3360009081526001602090815260408083206001600160a01b038516845290915290205460ff16610f0257611065565b3360009081526001602090815260408083206001600160a01b03851684529091528120805460ff191690555b33600090815260208190526040902054811015610e6557336000908152602081905260409020805482908110610f6057fe5b6000918252602090912001546001600160a01b03838116911614610f835761105d565b33600090815260208190526040902054811415610fc157336000908152602081905260409020805490610fba906000198301611d7f565b5050611065565b33600090815260208190526040902080546000198101908110610fe057fe5b600091825260208083209091015433835290829052604090912080546001600160a01b03909216918390811061101257fe5b600091825260208083209190910180546001600160a01b0319166001600160a01b039490941693909317909255338152908190526040902080549061105b906000198301611d7f565b505b600101610f2e565b50565b336000908152600260205260409020546001600160a01b0316156110c057604051600160e51b62461bcd02815260040180806020018281038252602181526020018061222d6021913960400191505060405180910390fd5b6040516110cc90611da3565b604051809103906000f0801580156110e8573d6000803e3d6000fd5b5033600090815260026020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b600160209081526000928352604080842090915290825290205460ff1681565b60408051602080820184905282518083038201815291830183528151918101919091206000818152600390925291902080546001600160a01b03166111c75760408051600160e51b62461bcd02815260206004820152601660248201527f526564207061636b616765206e6f742065786973747300000000000000000000604482015290519081900360640190fd5b42816007015410156112235760408051600160e51b62461bcd02815260206004820152601360248201527f526564207061636b616765206578706972656400000000000000000000000000604482015290519081900360640190fd5b8054600160a81b900460ff16156112b05780546001600160a01b0316600090815260016020908152604080832033845290915290205460ff166112b05760408051600160e51b62461bcd02815260206004820152601460248201527f4f6e6c7920667269656e642063616e2067726162000000000000000000000000604482015290519081900360640190fd5b6008810154600090815260046020908152604080832033845290915290205460ff16156113275760408051600160e51b62461bcd02815260206004820152601360248201527f63616e2774206772616262656420747769636500000000000000000000000000604482015290519081900360640190fd5b8054600090600160a01b900460ff161561135457816004015482600201548161134c57fe5b049050611402565b81600501546001141561136c57506003810154611402565b6003820154600583015460408051336020808301919091528183018590526060820193909352426080808301919091528251808303909101815260a09091019091528051918101919091209091600091816113c357fe5b069050600084600501548560030154816113d957fe5b049050816113ea57600193506113fe565b808211156113fa578093506113fe565b8193505b5050505b60018201546001600160a01b031661144757604051339082156108fc029083906000818181858888f19350505050158015611441573d6000803e3d6000fd5b50611477565b81546001600160a01b031633146114775760018201548254611477916001600160a01b0390811691163384611c87565b60038201805482900390556005820180546000190190819055611521576000838152600360208190526040822080546001600160b01b0319168155600180820180546001600160a01b03191690556002820184905591810183905560048101839055600581018390556006810183905560078101839055600801919091558201546001600160a01b03161561151c57815461151c906001600160a01b03166000611bcf565b61154b565b600882015460009081526004602090815260408083203384529091529020805460ff191660011790555b600182015460405182916001600160a01b03169033907ffd19781f43410d9594fd4c02dd1d98dbe99099cbd222d5851a6e183c468d33ca90600090a450505050565b60408051602080820184905282518083038201815291830183528151918101919091206000818152600390925291812080549192916001600160a01b031661161f5760408051600160e51b62461bcd02815260206004820152601660248201527f526564207061636b616765206e6f742065786973747300000000000000000000604482015290519081900360640190fd5b428160070154101561167b5760408051600160e51b62461bcd02815260206004820152601360248201527f526564207061636b616765206578706972656400000000000000000000000000604482015290519081900360640190fd5b8054600160a81b900460ff16156117085780546001600160a01b0316600090815260016020908152604080832033845290915290205460ff166117085760408051600160e51b62461bcd02815260206004820152601460248201527f4f6e6c7920667269656e642063616e2067726162000000000000000000000000604482015290519081900360640190fd5b60080154600090815260046020908152604080832033845290915290205460ff16159392505050565b336000908152600260209081526040808320548151600160e01b6370a082310281526001600160a01b038b81166004830152925192909116926370a0823192602480840193829003018186803b15801561178a57600080fd5b505afa15801561179e573d6000803e3d6000fd5b505050506040513d60208110156117b457600080fd5b50519050858110156118105760408051600160e51b62461bcd02815260206004820152601460248201527f6e6f742073756666696369656e742066756e6473000000000000000000000000604482015290519081900360640190fd5b6000831180156118205750600086115b801561182b57508286115b80156118375750600082115b61188b5760408051600160e51b62461bcd02815260206004820152601560248201527f696e76616c696420646174612070726f76696465640000000000000000000000604482015290519081900360640190fd5b6040805160208082018b905282518083038201815291830190925280519101206118b481611b14565b156119015760408051600160e51b62461bcd0281526020600482015260126024820152600160701b71526564207061636b6167652065786973747302604482015290519081900360640190fd5b85156119505783878161191057fe5b061561195057604051600160e51b62461bcd02815260040180806020018281038252602e8152602001806121d1602e913960400191505060405180910390fd5b604051806101600160405280336001600160a01b0316815260200187151581526020018615158152602001896001600160a01b03168152602001888152602001888152602001858152602001858152602001428152602001846201518002420181526020016005548152506003600083815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160000160146101000a81548160ff02191690831515021790555060408201518160000160156101000a81548160ff02191690831515021790555060608201518160010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506080820151816002015560a0820151816003015560c0820151816004015560e08201518160050155610100820151816006015561012082015181600701556101408201518160080155905050600560008154809291906001019190505550611ad1336001611bcf565b60405187906001600160a01b038a169033907f93eb3c629eb575edaf0252e4f9fc0c5ccada50496f8c1d32f0f93a65a8257eb590600090a4505050505050505050565b6000611b1e611d0f565b505060009081526003602081815260409283902083516101608101855281546001600160a01b0380821680845260ff600160a01b84048116151596850196909652600160a81b90920490941615159582019590955260018201549092166060830152600281015460808301529182015460a0820152600482015460c0820152600582015460e08201526006820154610100820152600782015461012082015260089091015461014090910152151590565b6001600160a01b038083166000908152600260205260409020541681611c4757806001600160a01b03166370e3fffe6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611c2a57600080fd5b505af1158015611c3e573d6000803e3d6000fd5b50505050611c82565b806001600160a01b03166346620e396040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561058c57600080fd5b505050565b6001600160a01b03808416600090815260026020526040808220548151600160e11b6349f59e31028152888516600482015286851660248201526044810186905291519316926393eb3c629260648084019391929182900301818387803b158015611cf157600080fd5b505af1158015611d05573d6000803e3d6000fd5b5050505050505050565b60405180610160016040528060006001600160a01b0316815260200160001515815260200160001515815260200160006001600160a01b03168152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b815481835581811115611c8257600083815260209020611c82918101908301611db0565b61040280611dcf83390190565b6108c491905b80821115611dca5760008155600101611db6565b509056fe60806040526000805460ff1916905534801561001a57600080fd5b50600080546101003302610100600160a81b03199091161790556103bf806100436000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806346620e391461006757806359a3ccb81461007157806370a08231146100ad57806370e3fffe146100e557806393eb3c62146100ed578063cf30901214610123575b600080fd5b61006f61013f565b005b61006f6004803603608081101561008757600080fd5b506001600160a01b0381358116916020810135821691604082013516906060013561016a565b6100d3600480360360208110156100c357600080fd5b50356001600160a01b0316610226565b60408051918252519081900360200190f35b61006f6102a5565b61006f6004803603606081101561010357600080fd5b506001600160a01b038135811691602081013590911690604001356102cd565b61012b61038a565b604080519115158252519081900360200190f35b60005461010090046001600160a01b0316331461015b57600080fd5b6000805460ff19166001179055565b60005461010090046001600160a01b0316331461018657600080fd5b60005460ff161561019657600080fd5b60408051600160e01b6323b872dd0281526001600160a01b0385811660048301528481166024830152604482018490529151918616916323b872dd916064808201926020929091908290030181600087803b1580156101f457600080fd5b505af1158015610208573d6000803e3d6000fd5b505050506040513d602081101561021e57600080fd5b505050505050565b60408051600160e01b6370a0823102815230600482015290516000916001600160a01b038416916370a0823191602480820192602092909190829003018186803b15801561027357600080fd5b505afa158015610287573d6000803e3d6000fd5b505050506040513d602081101561029d57600080fd5b505192915050565b60005461010090046001600160a01b031633146102c157600080fd5b6000805460ff19169055565b60005461010090046001600160a01b031633146102e957600080fd5b60005460ff16156102f957600080fd5b826001600160a01b031663a9059cbb83836040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561035957600080fd5b505af115801561036d573d6000803e3d6000fd5b505050506040513d602081101561038357600080fd5b5050505050565b60005460ff168156fea165627a7a723058205db0c3d653e56e3f7ea83d589929d8b6a985359fadf17168fbae3e4797334c080029496e76616c69642076616c756520616e642073697a6520666f7220657175616c206469766973696f6e206d6f6465526564207061636b616765206e6f7420657869737473206f7220796f75277265206e6f7420746865206f776e6572486173206265656e2063726561746564207265636569766572206164647265737343616e27742061646420796f757273656c6620746f20667269656e64206c697374a165627a7a7230582099fd806c708a7cfe6a93f4196af57a436dd634c8342edc259ef85e75d5df32bb0029`

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0x239fcf0f.
//
// Solidity: function Balance(address token) constant returns(uint256)
func (_Contract *ContractCaller) Balance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "Balance", token)
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0x239fcf0f.
//
// Solidity: function Balance(address token) constant returns(uint256)
func (_Contract *ContractSession) Balance(token common.Address) (*big.Int, error) {
	return _Contract.Contract.Balance(&_Contract.CallOpts, token)
}

// Balance is a free data retrieval call binding the contract method 0x239fcf0f.
//
// Solidity: function Balance(address token) constant returns(uint256)
func (_Contract *ContractCallerSession) Balance(token common.Address) (*big.Int, error) {
	return _Contract.Contract.Balance(&_Contract.CallOpts, token)
}

// CanGrab is a free data retrieval call binding the contract method 0x8749a6fa.
//
// Solidity: function CanGrab(bytes32 word) constant returns(bool has)
func (_Contract *ContractCaller) CanGrab(opts *bind.CallOpts, word [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "CanGrab", word)
	return *ret0, err
}

// CanGrab is a free data retrieval call binding the contract method 0x8749a6fa.
//
// Solidity: function CanGrab(bytes32 word) constant returns(bool has)
func (_Contract *ContractSession) CanGrab(word [32]byte) (bool, error) {
	return _Contract.Contract.CanGrab(&_Contract.CallOpts, word)
}

// CanGrab is a free data retrieval call binding the contract method 0x8749a6fa.
//
// Solidity: function CanGrab(bytes32 word) constant returns(bool has)
func (_Contract *ContractCallerSession) CanGrab(word [32]byte) (bool, error) {
	return _Contract.Contract.CanGrab(&_Contract.CallOpts, word)
}

// MyFriends is a free data retrieval call binding the contract method 0x37b37160.
//
// Solidity: function MyFriends() constant returns(address[])
func (_Contract *ContractCaller) MyFriends(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "MyFriends")
	return *ret0, err
}

// MyFriends is a free data retrieval call binding the contract method 0x37b37160.
//
// Solidity: function MyFriends() constant returns(address[])
func (_Contract *ContractSession) MyFriends() ([]common.Address, error) {
	return _Contract.Contract.MyFriends(&_Contract.CallOpts)
}

// MyFriends is a free data retrieval call binding the contract method 0x37b37160.
//
// Solidity: function MyFriends() constant returns(address[])
func (_Contract *ContractCallerSession) MyFriends() ([]common.Address, error) {
	return _Contract.Contract.MyFriends(&_Contract.CallOpts)
}

// MyToken is a free data retrieval call binding the contract method 0x089327de.
//
// Solidity: function MyToken() constant returns(address token)
func (_Contract *ContractCaller) MyToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "MyToken")
	return *ret0, err
}

// MyToken is a free data retrieval call binding the contract method 0x089327de.
//
// Solidity: function MyToken() constant returns(address token)
func (_Contract *ContractSession) MyToken() (common.Address, error) {
	return _Contract.Contract.MyToken(&_Contract.CallOpts)
}

// MyToken is a free data retrieval call binding the contract method 0x089327de.
//
// Solidity: function MyToken() constant returns(address token)
func (_Contract *ContractCallerSession) MyToken() (common.Address, error) {
	return _Contract.Contract.MyToken(&_Contract.CallOpts)
}

// Friendship is a free data retrieval call binding the contract method 0x66031fdf.
//
// Solidity: function friendship(address , address ) constant returns(bool)
func (_Contract *ContractCaller) Friendship(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "friendship", arg0, arg1)
	return *ret0, err
}

// Friendship is a free data retrieval call binding the contract method 0x66031fdf.
//
// Solidity: function friendship(address , address ) constant returns(bool)
func (_Contract *ContractSession) Friendship(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Contract.Contract.Friendship(&_Contract.CallOpts, arg0, arg1)
}

// Friendship is a free data retrieval call binding the contract method 0x66031fdf.
//
// Solidity: function friendship(address , address ) constant returns(bool)
func (_Contract *ContractCallerSession) Friendship(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Contract.Contract.Friendship(&_Contract.CallOpts, arg0, arg1)
}

// GetRecord is a free data retrieval call binding the contract method 0x213681cd.
//
// Solidity: function getRecord(bytes32 word) constant returns(address owner, bool equalDivision, bool onlyFriend, address token, uint256 amount, uint256 remainAmount, uint256 size, uint256 remainSize, uint256 timestamp, uint256 expired)
func (_Contract *ContractCaller) GetRecord(opts *bind.CallOpts, word [32]byte) (struct {
	Owner         common.Address
	EqualDivision bool
	OnlyFriend    bool
	Token         common.Address
	Amount        *big.Int
	RemainAmount  *big.Int
	Size          *big.Int
	RemainSize    *big.Int
	Timestamp     *big.Int
	Expired       *big.Int
}, error) {
	ret := new(struct {
		Owner         common.Address
		EqualDivision bool
		OnlyFriend    bool
		Token         common.Address
		Amount        *big.Int
		RemainAmount  *big.Int
		Size          *big.Int
		RemainSize    *big.Int
		Timestamp     *big.Int
		Expired       *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "getRecord", word)
	return *ret, err
}

// GetRecord is a free data retrieval call binding the contract method 0x213681cd.
//
// Solidity: function getRecord(bytes32 word) constant returns(address owner, bool equalDivision, bool onlyFriend, address token, uint256 amount, uint256 remainAmount, uint256 size, uint256 remainSize, uint256 timestamp, uint256 expired)
func (_Contract *ContractSession) GetRecord(word [32]byte) (struct {
	Owner         common.Address
	EqualDivision bool
	OnlyFriend    bool
	Token         common.Address
	Amount        *big.Int
	RemainAmount  *big.Int
	Size          *big.Int
	RemainSize    *big.Int
	Timestamp     *big.Int
	Expired       *big.Int
}, error) {
	return _Contract.Contract.GetRecord(&_Contract.CallOpts, word)
}

// GetRecord is a free data retrieval call binding the contract method 0x213681cd.
//
// Solidity: function getRecord(bytes32 word) constant returns(address owner, bool equalDivision, bool onlyFriend, address token, uint256 amount, uint256 remainAmount, uint256 size, uint256 remainSize, uint256 timestamp, uint256 expired)
func (_Contract *ContractCallerSession) GetRecord(word [32]byte) (struct {
	Owner         common.Address
	EqualDivision bool
	OnlyFriend    bool
	Token         common.Address
	Amount        *big.Int
	RemainAmount  *big.Int
	Size          *big.Int
	RemainSize    *big.Int
	Timestamp     *big.Int
	Expired       *big.Int
}, error) {
	return _Contract.Contract.GetRecord(&_Contract.CallOpts, word)
}

// AddFriend is a paid mutator transaction binding the contract method 0x13962258.
//
// Solidity: function AddFriend(address _friend) returns()
func (_Contract *ContractTransactor) AddFriend(opts *bind.TransactOpts, _friend common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "AddFriend", _friend)
}

// AddFriend is a paid mutator transaction binding the contract method 0x13962258.
//
// Solidity: function AddFriend(address _friend) returns()
func (_Contract *ContractSession) AddFriend(_friend common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddFriend(&_Contract.TransactOpts, _friend)
}

// AddFriend is a paid mutator transaction binding the contract method 0x13962258.
//
// Solidity: function AddFriend(address _friend) returns()
func (_Contract *ContractTransactorSession) AddFriend(_friend common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddFriend(&_Contract.TransactOpts, _friend)
}

// AddFriendList is a paid mutator transaction binding the contract method 0x2bba4dde.
//
// Solidity: function AddFriendList(address[] list) returns()
func (_Contract *ContractTransactor) AddFriendList(opts *bind.TransactOpts, list []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "AddFriendList", list)
}

// AddFriendList is a paid mutator transaction binding the contract method 0x2bba4dde.
//
// Solidity: function AddFriendList(address[] list) returns()
func (_Contract *ContractSession) AddFriendList(list []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddFriendList(&_Contract.TransactOpts, list)
}

// AddFriendList is a paid mutator transaction binding the contract method 0x2bba4dde.
//
// Solidity: function AddFriendList(address[] list) returns()
func (_Contract *ContractTransactorSession) AddFriendList(list []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddFriendList(&_Contract.TransactOpts, list)
}

// DelFriend is a paid mutator transaction binding the contract method 0x5663132e.
//
// Solidity: function DelFriend(address _friend) returns()
func (_Contract *ContractTransactor) DelFriend(opts *bind.TransactOpts, _friend common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "DelFriend", _friend)
}

// DelFriend is a paid mutator transaction binding the contract method 0x5663132e.
//
// Solidity: function DelFriend(address _friend) returns()
func (_Contract *ContractSession) DelFriend(_friend common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DelFriend(&_Contract.TransactOpts, _friend)
}

// DelFriend is a paid mutator transaction binding the contract method 0x5663132e.
//
// Solidity: function DelFriend(address _friend) returns()
func (_Contract *ContractTransactorSession) DelFriend(_friend common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DelFriend(&_Contract.TransactOpts, _friend)
}

// Giving is a paid mutator transaction binding the contract method 0xd3a04056.
//
// Solidity: function Giving(bytes32 word, address token, uint256 value, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) returns()
func (_Contract *ContractTransactor) Giving(opts *bind.TransactOpts, word [32]byte, token common.Address, value *big.Int, equalDivision bool, onlyFriend bool, size *big.Int, expireDays *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "Giving", word, token, value, equalDivision, onlyFriend, size, expireDays)
}

// Giving is a paid mutator transaction binding the contract method 0xd3a04056.
//
// Solidity: function Giving(bytes32 word, address token, uint256 value, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) returns()
func (_Contract *ContractSession) Giving(word [32]byte, token common.Address, value *big.Int, equalDivision bool, onlyFriend bool, size *big.Int, expireDays *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Giving(&_Contract.TransactOpts, word, token, value, equalDivision, onlyFriend, size, expireDays)
}

// Giving is a paid mutator transaction binding the contract method 0xd3a04056.
//
// Solidity: function Giving(bytes32 word, address token, uint256 value, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) returns()
func (_Contract *ContractTransactorSession) Giving(word [32]byte, token common.Address, value *big.Int, equalDivision bool, onlyFriend bool, size *big.Int, expireDays *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Giving(&_Contract.TransactOpts, word, token, value, equalDivision, onlyFriend, size, expireDays)
}

// Grabbing is a paid mutator transaction binding the contract method 0x86060c02.
//
// Solidity: function Grabbing(bytes32 word) returns()
func (_Contract *ContractTransactor) Grabbing(opts *bind.TransactOpts, word [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "Grabbing", word)
}

// Grabbing is a paid mutator transaction binding the contract method 0x86060c02.
//
// Solidity: function Grabbing(bytes32 word) returns()
func (_Contract *ContractSession) Grabbing(word [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Grabbing(&_Contract.TransactOpts, word)
}

// Grabbing is a paid mutator transaction binding the contract method 0x86060c02.
//
// Solidity: function Grabbing(bytes32 word) returns()
func (_Contract *ContractTransactorSession) Grabbing(word [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Grabbing(&_Contract.TransactOpts, word)
}

// NewTokenReceiver is a paid mutator transaction binding the contract method 0x5daccf51.
//
// Solidity: function NewTokenReceiver() returns()
func (_Contract *ContractTransactor) NewTokenReceiver(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "NewTokenReceiver")
}

// NewTokenReceiver is a paid mutator transaction binding the contract method 0x5daccf51.
//
// Solidity: function NewTokenReceiver() returns()
func (_Contract *ContractSession) NewTokenReceiver() (*types.Transaction, error) {
	return _Contract.Contract.NewTokenReceiver(&_Contract.TransactOpts)
}

// NewTokenReceiver is a paid mutator transaction binding the contract method 0x5daccf51.
//
// Solidity: function NewTokenReceiver() returns()
func (_Contract *ContractTransactorSession) NewTokenReceiver() (*types.Transaction, error) {
	return _Contract.Contract.NewTokenReceiver(&_Contract.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0x22b3f549.
//
// Solidity: function Revoke(bytes32 word) returns()
func (_Contract *ContractTransactor) Revoke(opts *bind.TransactOpts, word [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "Revoke", word)
}

// Revoke is a paid mutator transaction binding the contract method 0x22b3f549.
//
// Solidity: function Revoke(bytes32 word) returns()
func (_Contract *ContractSession) Revoke(word [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Revoke(&_Contract.TransactOpts, word)
}

// Revoke is a paid mutator transaction binding the contract method 0x22b3f549.
//
// Solidity: function Revoke(bytes32 word) returns()
func (_Contract *ContractTransactorSession) Revoke(word [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Revoke(&_Contract.TransactOpts, word)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x03723885.
//
// Solidity: function WithdrawToken(address token, address _to, uint256 value) returns()
func (_Contract *ContractTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, _to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "WithdrawToken", token, _to, value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x03723885.
//
// Solidity: function WithdrawToken(address token, address _to, uint256 value) returns()
func (_Contract *ContractSession) WithdrawToken(token common.Address, _to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawToken(&_Contract.TransactOpts, token, _to, value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x03723885.
//
// Solidity: function WithdrawToken(address token, address _to, uint256 value) returns()
func (_Contract *ContractTransactorSession) WithdrawToken(token common.Address, _to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawToken(&_Contract.TransactOpts, token, _to, value)
}

// ContractReceiveIterator is returned from FilterReceive and is used to iterate over the raw logs and unpacked data for Receive events raised by the Contract contract.
type ContractReceiveIterator struct {
	Event *ContractReceive // Event containing the contract specifics and raw log

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
func (it *ContractReceiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractReceive)
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
		it.Event = new(ContractReceive)
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
func (it *ContractReceiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractReceiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractReceive represents a Receive event raised by the Contract contract.
type ContractReceive struct {
	Receiver common.Address
	Token    common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceive is a free log retrieval operation binding the contract event 0xfd19781f43410d9594fd4c02dd1d98dbe99099cbd222d5851a6e183c468d33ca.
//
// Solidity: event Receive(address indexed receiver, address indexed token, uint256 indexed value)
func (_Contract *ContractFilterer) FilterReceive(opts *bind.FilterOpts, receiver []common.Address, token []common.Address, value []*big.Int) (*ContractReceiveIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Receive", receiverRule, tokenRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &ContractReceiveIterator{contract: _Contract.contract, event: "Receive", logs: logs, sub: sub}, nil
}

// WatchReceive is a free log subscription operation binding the contract event 0xfd19781f43410d9594fd4c02dd1d98dbe99099cbd222d5851a6e183c468d33ca.
//
// Solidity: event Receive(address indexed receiver, address indexed token, uint256 indexed value)
func (_Contract *ContractFilterer) WatchReceive(opts *bind.WatchOpts, sink chan<- *ContractReceive, receiver []common.Address, token []common.Address, value []*big.Int) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Receive", receiverRule, tokenRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractReceive)
				if err := _Contract.contract.UnpackLog(event, "Receive", log); err != nil {
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

// ContractSendIterator is returned from FilterSend and is used to iterate over the raw logs and unpacked data for Send events raised by the Contract contract.
type ContractSendIterator struct {
	Event *ContractSend // Event containing the contract specifics and raw log

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
func (it *ContractSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSend)
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
		it.Event = new(ContractSend)
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
func (it *ContractSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSend represents a Send event raised by the Contract contract.
type ContractSend struct {
	Sender common.Address
	Token  common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSend is a free log retrieval operation binding the contract event 0x93eb3c629eb575edaf0252e4f9fc0c5ccada50496f8c1d32f0f93a65a8257eb5.
//
// Solidity: event Send(address indexed sender, address indexed token, uint256 indexed value)
func (_Contract *ContractFilterer) FilterSend(opts *bind.FilterOpts, sender []common.Address, token []common.Address, value []*big.Int) (*ContractSendIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Send", senderRule, tokenRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &ContractSendIterator{contract: _Contract.contract, event: "Send", logs: logs, sub: sub}, nil
}

// WatchSend is a free log subscription operation binding the contract event 0x93eb3c629eb575edaf0252e4f9fc0c5ccada50496f8c1d32f0f93a65a8257eb5.
//
// Solidity: event Send(address indexed sender, address indexed token, uint256 indexed value)
func (_Contract *ContractFilterer) WatchSend(opts *bind.WatchOpts, sink chan<- *ContractSend, sender []common.Address, token []common.Address, value []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Send", senderRule, tokenRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSend)
				if err := _Contract.contract.UnpackLog(event, "Send", log); err != nil {
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
