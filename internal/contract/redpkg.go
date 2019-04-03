// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package redpkg

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

// RedpkgABI is the input ABI used to generate the binding from.
const RedpkgABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"word\",\"type\":\"string\"}],\"name\":\"getRecord\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_friend\",\"type\":\"address\"}],\"name\":\"AddFriend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"word\",\"type\":\"string\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"equalDivision\",\"type\":\"bool\"},{\"name\":\"onlyFriend\",\"type\":\"bool\"},{\"name\":\"size\",\"type\":\"uint256\"},{\"name\":\"expireDays\",\"type\":\"uint256\"}],\"name\":\"Giving\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"word\",\"type\":\"string\"}],\"name\":\"Grabbing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"Balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"word\",\"type\":\"string\"},{\"name\":\"equalDivision\",\"type\":\"bool\"},{\"name\":\"onlyFriend\",\"type\":\"bool\"},{\"name\":\"size\",\"type\":\"uint256\"},{\"name\":\"expireDays\",\"type\":\"uint256\"}],\"name\":\"Giving\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"list\",\"type\":\"address[]\"}],\"name\":\"AddFriendList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MyFriends\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_friend\",\"type\":\"address\"}],\"name\":\"DelFriend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"NewTokenReceiver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"friendship\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"word\",\"type\":\"string\"}],\"name\":\"IsWordExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// RedpkgBin is the compiled bytecode used for deploying new contracts.
const RedpkgBin = `608060405234801561001057600080fd5b5061221f806100206000396000f3fe6080604052600436106100dd5760003560e01c80632bba4dde1161007f5780635daccf51116100595780635daccf51146106a157806366031fdf146106b6578063ce7d23d814610705578063e4860339146107b6576100dd565b80632bba4dde1461055b57806337b37160146106095780635663132e1461066e576100dd565b806314313999116100bb57806314313999146102d15780631f1dc595146103ac578063239fcf0f1461045d5780632a43e999146104a2576100dd565b806303723885146100e257806311dd884514610127578063139622581461029e575b600080fd5b3480156100ee57600080fd5b506101256004803603606081101561010557600080fd5b506001600160a01b03813581169160208101359091169060400135610805565b005b34801561013357600080fd5b506101d86004803603602081101561014a57600080fd5b810190602081018135600160201b81111561016457600080fd5b82018360208201111561017657600080fd5b803590602001918460018302840111600160201b8311171561019757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610889945050505050565b604051808c6001600160a01b03166001600160a01b031681526020018b1515151581526020018a151515158152602001896001600160a01b03166001600160a01b0316815260200188815260200187815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015610280578181015183820152602001610268565b505050509050019c5050505050505050505050505060405180910390f35b3480156102aa57600080fd5b50610125600480360360208110156102c157600080fd5b50356001600160a01b0316610a57565b3480156102dd57600080fd5b50610125600480360360e08110156102f457600080fd5b810190602081018135600160201b81111561030e57600080fd5b82018360208201111561032057600080fd5b803590602001918460018302840111600160201b8311171561034157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550506001600160a01b03833516935050506020810135906040810135151590606081013515159060808101359060a00135610ac2565b3480156103b857600080fd5b50610125600480360360208110156103cf57600080fd5b810190602081018135600160201b8111156103e957600080fd5b8201836020820111156103fb57600080fd5b803590602001918460018302840111600160201b8311171561041c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610e75945050505050565b34801561046957600080fd5b506104906004803603602081101561048057600080fd5b50356001600160a01b03166113ae565b60408051918252519081900360200190f35b610125600480360360a08110156104b857600080fd5b810190602081018135600160201b8111156104d257600080fd5b8201836020820111156104e457600080fd5b803590602001918460018302840111600160201b8311171561050557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505050803515159150602081013515159060408101359060600135611439565b34801561056757600080fd5b506101256004803603602081101561057e57600080fd5b810190602081018135600160201b81111561059857600080fd5b8201836020820111156105aa57600080fd5b803590602001918460208302840111600160201b831117156105cb57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506116f4945050505050565b34801561061557600080fd5b5061061e6117b7565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561065a578181015183820152602001610642565b505050509050019250505060405180910390f35b34801561067a57600080fd5b506101256004803603602081101561069157600080fd5b50356001600160a01b0316611821565b3480156106ad57600080fd5b506101256119af565b3480156106c257600080fd5b506106f1600480360360408110156106d957600080fd5b506001600160a01b0381358116916020013516611a62565b604080519115158252519081900360200190f35b34801561071157600080fd5b506106f16004803603602081101561072857600080fd5b810190602081018135600160201b81111561074257600080fd5b82018360208201111561075457600080fd5b803590602001918460018302840111600160201b8311171561077557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611a82945050505050565b3480156107c257600080fd5b506107e9600480360360208110156107d957600080fd5b50356001600160a01b0316611bf1565b604080516001600160a01b039092168252519081900360200190f35b33600090815260026020526040808220548151600160e31b6317d575990281526001600160a01b0387811660048301528681166024830152604482018690529251929091169263beabacc89260648084019382900301818387803b15801561086c57600080fd5b505af1158015610880573d6000803e3d6000fd5b50505050505050565b60008060008060008060008060008060606108a2611c83565b60038d6040518082805190602001908083835b602083106108d45780518252601f1990920191602091820191016108b5565b518151600019602094850361010090810a9190910191821691199290921617909152939091019586526040805196879003820187206101608801825280546001600160a01b038082168a5260ff600160a01b8304811615158b870152600160a81b909204909116151589840152600182015416606089015260028101546080890152600381015460a0890152600481015460c0890152600581015460e08901526006810154948801949094526007840154610120880152600884018054825181850281018501909352808352949650610140880195509093909250908301828280156109e957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116109cb575b5050505050815250509050806000015181602001518260400151836060015184608001518560a001518660c001518760e001518861010001518961012001518a61014001518090509b509b509b509b509b509b509b509b509b509b509b505091939597999b90929496989a50565b3360009081526001602090815260408083206001600160a01b038516845290915290205460ff1615610a8857610abf565b336000908152602081815260408220805460018101825590835291200180546001600160a01b0319166001600160a01b0383161790555b50565b336000908152600260209081526040808320548151600160e01b6370a082310281526001600160a01b038b81166004830152925192909116926370a0823192602480840193829003018186803b158015610b1b57600080fd5b505afa158015610b2f573d6000803e3d6000fd5b505050506040513d6020811015610b4557600080fd5b5051905085811015610ba15760408051600160e51b62461bcd02815260206004820152601460248201527f6e6f742073756666696369656e742066756e6473000000000000000000000000604482015290519081900360640190fd5b600083118015610bb15750600086115b8015610bbc57508286115b8015610bc85750600082115b1515610c1e5760408051600160e51b62461bcd02815260206004820152601560248201527f696e76616c696420646174612070726f76696465640000000000000000000000604482015290519081900360640190fd5b610c2788611a82565b15610c745760408051600160e51b62461bcd0281526020600482015260126024820152600160701b71526564207061636b6167652065786973747302604482015290519081900360640190fd5b8415610cc5578286811515610c8557fe5b0615610cc557604051600160e51b62461bcd02815260040180806020018281038252602e8152602001806121a5602e913960400191505060405180910390fd5b6060604051806101600160405280336001600160a01b0316815260200187151581526020018615158152602001896001600160a01b03168152602001888152602001888152602001858152602001858152602001428152602001846201518002420181526020018281525060038a6040518082805190602001908083835b60208310610d625780518252601f199092019160209182019101610d43565b518151600019602094850361010090810a919091019182169119929092161790915293909101958652604080519687900382019096208751815489840151988a01511515600160a81b02600160a81b60ff0219991515600160a01b02600160a01b60ff02196001600160a01b039485166001600160a01b0319948516171617999099169890981782556060890151600183018054919092169816979097179096556080870151600287015560a0870151600387015560c0870151600487015560e087015160058701559186015160068601555061012085015160078501556101408501518051610e5b9450600886019350910190611cf3565b50905050610e6a336001611c0c565b505050505050505050565b60006003826040518082805190602001908083835b60208310610ea95780518252601f199092019160209182019101610e8a565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922080549093506001600160a01b031615159150610f3c90505760408051600160e51b62461bcd02815260206004820152601660248201527f526564207061636b616765206e6f742065786973747300000000000000000000604482015290519081900360640190fd5b6007810154421115610f985760408051600160e51b62461bcd02815260206004820152601360248201527f526564207061636b616765206578706972656400000000000000000000000000604482015290519081900360640190fd5b8054600160a81b900460ff16156110275780546001600160a01b0316600090815260016020908152604080832033845290915290205460ff1615156110275760408051600160e51b62461bcd02815260206004820152601460248201527f4f6e6c7920667269656e642063616e2067726162000000000000000000000000604482015290519081900360640190fd5b60005b60088201548110156110ba576008820180548290811061104657fe5b6000918252602090912001546001600160a01b03163314156110b25760408051600160e51b62461bcd02815260206004820152601360248201527f63616e2774206772616262656420747769636500000000000000000000000000604482015290519081900360640190fd5b60010161102a565b508054600090600160a01b900460ff16156110ea57816004015482600201548115156110e257fe5b0490506111a9565b816005015460011415611102575060038101546111a9565b600382015460058301546040805160001943014060208083019190915233828401526060820185905260808201939093524260a0808301919091528251808303909101815260c0909101909152805191810191909120909160009181151561116657fe5b06905060008460050154856003015481151561117e57fe5b04905081151561119157600193506111a5565b808211156111a1578093506111a5565b8193505b5050505b60018201546001600160a01b031615156111f057604051339082156108fc029083906000818181858888f193505050501580156111ea573d6000803e3d6000fd5b50611278565b81546001600160a01b039081166000908152600260205260408082205460018601548251600160e31b6317d57599028152908516600482015233602482015260448101869052915193169263beabacc89260648084019391929182900301818387803b15801561125f57600080fd5b505af1158015611273573d6000803e3d6000fd5b505050505b600382018054829003905560058201805460001901908190551515611384576003836040518082805190602001908083835b602083106112c95780518252601f1990920191602091820191016112aa565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922080546001600160b01b03191681556001810180546001600160a01b0319169055600060028201819055600382018190556004820181905560058201819055600682018190556007820181905590925090506113526008830182611d58565b505060018201546001600160a01b03161561137d57815461137d906001600160a01b03166000611c0c565b5050610abf565b506008018054600181018255600091825260209091200180546001600160a01b0319163317905550565b336000908152600260209081526040808320548151600160e01b6370a082310281526001600160a01b038681166004830152925192909116926370a0823192602480840193829003018186803b15801561140757600080fd5b505afa15801561141b573d6000803e3d6000fd5b505050506040513d602081101561143157600080fd5b505192915050565b6000821180156114495750600034115b801561145457508134115b80156114605750600081115b15156114b65760408051600160e51b62461bcd02815260206004820152601560248201527f696e76616c696420646174612070726f76696465640000000000000000000000604482015290519081900360640190fd5b6114bf85611a82565b1561150c5760408051600160e51b62461bcd0281526020600482015260126024820152600160701b71526564207061636b6167652065786973747302604482015290519081900360640190fd5b831561155d57813481151561151d57fe5b061561155d57604051600160e51b62461bcd02815260040180806020018281038252602e8152602001806121a5602e913960400191505060405180910390fd5b6060604051806101600160405280336001600160a01b031681526020018615158152602001851515815260200160006001600160a01b0316815260200134815260200134815260200184815260200184815260200142815260200183620151800242018152602001828152506003876040518082805190602001908083835b602083106115fb5780518252601f1990920191602091820191016115dc565b518151600019602094850361010090810a919091019182169119929092161790915293909101958652604080519687900382019096208751815489840151988a01511515600160a81b02600160a81b60ff0219991515600160a01b02600160a01b60ff02196001600160a01b039485166001600160a01b0319948516171617999099169890981782556060890151600183018054919092169816979097179096556080870151600287015560a0870151600387015560c0870151600487015560e087015160058701559186015160068601555061012085015160078501556101408501518051610e6a9450600886019350910190611cf3565b60005b81518110156117b3576000828281518110151561171057fe5b60209081029091018101513360009081526001835260408082206001600160a01b0384168352909352919091205490915060ff161561174f57506117ab565b336000818152602081815260408083208054600180820183559185528385200180546001600160a01b0319166001600160a01b0397909716968717905593835283825280832094835293905291909120805460ff191690911790555b6001016116f7565b5050565b336000908152602081815260409182902080548351818402810184019094528084526060939283018282801561181657602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116117f8575b505050505090505b90565b3360009081526001602090815260408083206001600160a01b038516845290915290205460ff16151561185357610abf565b3360009081526001602090815260408083206001600160a01b03851684529091528120805460ff191690555b336000908152602081905260409020548110156117b3573360009081526020819052604090208054829081106118b157fe5b6000918252602090912001546001600160a01b038381169116146118d4576119a7565b3360009081526020819052604090205481141561190b5733600090815260208190526040902080549061137d906000198301611d76565b3360009081526020819052604090208054600019810190811061192a57fe5b600091825260208083209091015433835290829052604090912080546001600160a01b03909216918390811061195c57fe5b600091825260208083209190910180546001600160a01b0319166001600160a01b03949094169390931790925533815290819052604090208054906119a5906000198301611d76565b505b60010161187f565b336000908152600260205260409020546001600160a01b031615611a0757604051600160e51b62461bcd0281526004018080602001828103825260218152602001806121d36021913960400191505060405180910390fd5b6000604051611a1590611d9f565b604051809103906000f080158015611a31573d6000803e3d6000fd5b5033600090815260026020526040902080546001600160a01b0319166001600160a01b039290921691909117905550565b600160209081526000928352604080842090915290825290205460ff1681565b6000611a8c611c83565b6003836040518082805190602001908083835b60208310611abe5780518252601f199092019160209182019101611a9f565b518151600019602094850361010090810a9190910191821691199290921617909152939091019586526040805196879003820187206101608801825280546001600160a01b038082168a5260ff600160a01b8304811615158b870152600160a81b909204909116151589840152600182015416606089015260028101546080890152600381015460a0890152600481015460c0890152600581015460e0890152600681015494880194909452600784015461012088015260088401805482518185028101850190935280835294965061014088019550909390925090830182828015611bd357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611bb5575b50505091909252505090516001600160a01b03161515949350505050565b6002602052600090815260409020546001600160a01b031681565b6001600160a01b03808316600090815260026020526040808220548151600160e01b633db31e9f02815285151560048201529151931692633db31e9f9260248084019391929182900301818387803b158015611c6757600080fd5b505af1158015611c7b573d6000803e3d6000fd5b505050505050565b60405180610160016040528060006001600160a01b0316815260200160001515815260200160001515815260200160006001600160a01b03168152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001606081525090565b828054828255906000526020600020908101928215611d48579160200282015b82811115611d4857825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611d13565b50611d54929150611dac565b5090565b5080546000825590600052602060002090810190610abf9190611dd0565b815481835581811115611d9a57600083815260209020611d9a918101908301611dd0565b505050565b6103ba80611deb83390190565b61181e91905b80821115611d545780546001600160a01b0319168155600101611db2565b61181e91905b80821115611d545760008155600101611dd656fe60806040526000805460ff1916905534801561001a57600080fd5b50600080546101003302610100600160a81b0319909116179055610377806100436000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806315dacbea1461005c5780633db31e9f1461009a57806370a08231146100b9578063beabacc8146100f1578063cf30901214610127575b600080fd5b6100986004803603608081101561007257600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135610143565b005b610098600480360360208110156100b057600080fd5b503515156101f3565b6100df600480360360208110156100cf57600080fd5b50356001600160a01b0316610222565b60408051918252519081900360200190f35b6100986004803603606081101561010757600080fd5b506001600160a01b038135811691602081013590911690604001356102a1565b61012f610342565b604080519115158252519081900360200190f35b60005460ff161561015357600080fd5b60005460ff161561016357600080fd5b60408051600160e01b6323b872dd0281526001600160a01b0385811660048301528481166024830152604482018490529151918616916323b872dd916064808201926020929091908290030181600087803b1580156101c157600080fd5b505af11580156101d5573d6000803e3d6000fd5b505050506040513d60208110156101eb57600080fd5b505050505050565b60005461010090046001600160a01b0316331461020f57600080fd5b6000805460ff1916911515919091179055565b60408051600160e01b6370a0823102815230600482015290516000916001600160a01b038416916370a0823191602480820192602092909190829003018186803b15801561026f57600080fd5b505afa158015610283573d6000803e3d6000fd5b505050506040513d602081101561029957600080fd5b505192915050565b60005460ff16156102b157600080fd5b826001600160a01b031663a9059cbb83836040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561031157600080fd5b505af1158015610325573d6000803e3d6000fd5b505050506040513d602081101561033b57600080fd5b5050505050565b60005460ff168156fea165627a7a723058207cb627ab1cacb29d4d31b2fc0c69033e27112b1604859592186f5821e5b963bf0029496e76616c69642076616c756520616e642073697a6520666f7220657175616c206469766973696f6e206d6f6465486173206265656e20637265617465642072656365697665722061646472657373a165627a7a72305820b50707104b7c2239c5d4b711dcfad3f32f58086d25307afefb4eeff2c32a85320029`

// DeployRedpkg deploys a new Ethereum contract, binding an instance of Redpkg to it.
func DeployRedpkg(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Redpkg, error) {
	parsed, err := abi.JSON(strings.NewReader(RedpkgABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RedpkgBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Redpkg{RedpkgCaller: RedpkgCaller{contract: contract}, RedpkgTransactor: RedpkgTransactor{contract: contract}, RedpkgFilterer: RedpkgFilterer{contract: contract}}, nil
}

// Redpkg is an auto generated Go binding around an Ethereum contract.
type Redpkg struct {
	RedpkgCaller     // Read-only binding to the contract
	RedpkgTransactor // Write-only binding to the contract
	RedpkgFilterer   // Log filterer for contract events
}

// RedpkgCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedpkgCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedpkgTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedpkgTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedpkgFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedpkgFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedpkgSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedpkgSession struct {
	Contract     *Redpkg           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedpkgCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedpkgCallerSession struct {
	Contract *RedpkgCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RedpkgTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedpkgTransactorSession struct {
	Contract     *RedpkgTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedpkgRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedpkgRaw struct {
	Contract *Redpkg // Generic contract binding to access the raw methods on
}

// RedpkgCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedpkgCallerRaw struct {
	Contract *RedpkgCaller // Generic read-only contract binding to access the raw methods on
}

// RedpkgTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedpkgTransactorRaw struct {
	Contract *RedpkgTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedpkg creates a new instance of Redpkg, bound to a specific deployed contract.
func NewRedpkg(address common.Address, backend bind.ContractBackend) (*Redpkg, error) {
	contract, err := bindRedpkg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Redpkg{RedpkgCaller: RedpkgCaller{contract: contract}, RedpkgTransactor: RedpkgTransactor{contract: contract}, RedpkgFilterer: RedpkgFilterer{contract: contract}}, nil
}

// NewRedpkgCaller creates a new read-only instance of Redpkg, bound to a specific deployed contract.
func NewRedpkgCaller(address common.Address, caller bind.ContractCaller) (*RedpkgCaller, error) {
	contract, err := bindRedpkg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedpkgCaller{contract: contract}, nil
}

// NewRedpkgTransactor creates a new write-only instance of Redpkg, bound to a specific deployed contract.
func NewRedpkgTransactor(address common.Address, transactor bind.ContractTransactor) (*RedpkgTransactor, error) {
	contract, err := bindRedpkg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedpkgTransactor{contract: contract}, nil
}

// NewRedpkgFilterer creates a new log filterer instance of Redpkg, bound to a specific deployed contract.
func NewRedpkgFilterer(address common.Address, filterer bind.ContractFilterer) (*RedpkgFilterer, error) {
	contract, err := bindRedpkg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedpkgFilterer{contract: contract}, nil
}

// bindRedpkg binds a generic wrapper to an already deployed contract.
func bindRedpkg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RedpkgABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redpkg *RedpkgRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Redpkg.Contract.RedpkgCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redpkg *RedpkgRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redpkg.Contract.RedpkgTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redpkg *RedpkgRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redpkg.Contract.RedpkgTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redpkg *RedpkgCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Redpkg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redpkg *RedpkgTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redpkg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redpkg *RedpkgTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redpkg.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0x239fcf0f.
//
// Solidity: function Balance(address token) constant returns(uint256)
func (_Redpkg *RedpkgCaller) Balance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Redpkg.contract.Call(opts, out, "Balance", token)
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0x239fcf0f.
//
// Solidity: function Balance(address token) constant returns(uint256)
func (_Redpkg *RedpkgSession) Balance(token common.Address) (*big.Int, error) {
	return _Redpkg.Contract.Balance(&_Redpkg.CallOpts, token)
}

// Balance is a free data retrieval call binding the contract method 0x239fcf0f.
//
// Solidity: function Balance(address token) constant returns(uint256)
func (_Redpkg *RedpkgCallerSession) Balance(token common.Address) (*big.Int, error) {
	return _Redpkg.Contract.Balance(&_Redpkg.CallOpts, token)
}

// IsWordExists is a free data retrieval call binding the contract method 0xce7d23d8.
//
// Solidity: function IsWordExists(string word) constant returns(bool)
func (_Redpkg *RedpkgCaller) IsWordExists(opts *bind.CallOpts, word string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Redpkg.contract.Call(opts, out, "IsWordExists", word)
	return *ret0, err
}

// IsWordExists is a free data retrieval call binding the contract method 0xce7d23d8.
//
// Solidity: function IsWordExists(string word) constant returns(bool)
func (_Redpkg *RedpkgSession) IsWordExists(word string) (bool, error) {
	return _Redpkg.Contract.IsWordExists(&_Redpkg.CallOpts, word)
}

// IsWordExists is a free data retrieval call binding the contract method 0xce7d23d8.
//
// Solidity: function IsWordExists(string word) constant returns(bool)
func (_Redpkg *RedpkgCallerSession) IsWordExists(word string) (bool, error) {
	return _Redpkg.Contract.IsWordExists(&_Redpkg.CallOpts, word)
}

// MyFriends is a free data retrieval call binding the contract method 0x37b37160.
//
// Solidity: function MyFriends() constant returns(address[])
func (_Redpkg *RedpkgCaller) MyFriends(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Redpkg.contract.Call(opts, out, "MyFriends")
	return *ret0, err
}

// MyFriends is a free data retrieval call binding the contract method 0x37b37160.
//
// Solidity: function MyFriends() constant returns(address[])
func (_Redpkg *RedpkgSession) MyFriends() ([]common.Address, error) {
	return _Redpkg.Contract.MyFriends(&_Redpkg.CallOpts)
}

// MyFriends is a free data retrieval call binding the contract method 0x37b37160.
//
// Solidity: function MyFriends() constant returns(address[])
func (_Redpkg *RedpkgCallerSession) MyFriends() ([]common.Address, error) {
	return _Redpkg.Contract.MyFriends(&_Redpkg.CallOpts)
}

// Friendship is a free data retrieval call binding the contract method 0x66031fdf.
//
// Solidity: function friendship(address , address ) constant returns(bool)
func (_Redpkg *RedpkgCaller) Friendship(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Redpkg.contract.Call(opts, out, "friendship", arg0, arg1)
	return *ret0, err
}

// Friendship is a free data retrieval call binding the contract method 0x66031fdf.
//
// Solidity: function friendship(address , address ) constant returns(bool)
func (_Redpkg *RedpkgSession) Friendship(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Redpkg.Contract.Friendship(&_Redpkg.CallOpts, arg0, arg1)
}

// Friendship is a free data retrieval call binding the contract method 0x66031fdf.
//
// Solidity: function friendship(address , address ) constant returns(bool)
func (_Redpkg *RedpkgCallerSession) Friendship(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Redpkg.Contract.Friendship(&_Redpkg.CallOpts, arg0, arg1)
}

// GetRecord is a free data retrieval call binding the contract method 0x11dd8845.
//
// Solidity: function getRecord(string word) constant returns(address, bool, bool, address, uint256, uint256, uint256, uint256, uint256, uint256, address[])
func (_Redpkg *RedpkgCaller) GetRecord(opts *bind.CallOpts, word string) (common.Address, bool, bool, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, []common.Address, error) {
	var (
		ret0  = new(common.Address)
		ret1  = new(bool)
		ret2  = new(bool)
		ret3  = new(common.Address)
		ret4  = new(*big.Int)
		ret5  = new(*big.Int)
		ret6  = new(*big.Int)
		ret7  = new(*big.Int)
		ret8  = new(*big.Int)
		ret9  = new(*big.Int)
		ret10 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
		ret9,
		ret10,
	}
	err := _Redpkg.contract.Call(opts, out, "getRecord", word)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, *ret9, *ret10, err
}

// GetRecord is a free data retrieval call binding the contract method 0x11dd8845.
//
// Solidity: function getRecord(string word) constant returns(address, bool, bool, address, uint256, uint256, uint256, uint256, uint256, uint256, address[])
func (_Redpkg *RedpkgSession) GetRecord(word string) (common.Address, bool, bool, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, []common.Address, error) {
	return _Redpkg.Contract.GetRecord(&_Redpkg.CallOpts, word)
}

// GetRecord is a free data retrieval call binding the contract method 0x11dd8845.
//
// Solidity: function getRecord(string word) constant returns(address, bool, bool, address, uint256, uint256, uint256, uint256, uint256, uint256, address[])
func (_Redpkg *RedpkgCallerSession) GetRecord(word string) (common.Address, bool, bool, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, []common.Address, error) {
	return _Redpkg.Contract.GetRecord(&_Redpkg.CallOpts, word)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) constant returns(address)
func (_Redpkg *RedpkgCaller) Tokens(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Redpkg.contract.Call(opts, out, "tokens", arg0)
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) constant returns(address)
func (_Redpkg *RedpkgSession) Tokens(arg0 common.Address) (common.Address, error) {
	return _Redpkg.Contract.Tokens(&_Redpkg.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) constant returns(address)
func (_Redpkg *RedpkgCallerSession) Tokens(arg0 common.Address) (common.Address, error) {
	return _Redpkg.Contract.Tokens(&_Redpkg.CallOpts, arg0)
}

// AddFriend is a paid mutator transaction binding the contract method 0x13962258.
//
// Solidity: function AddFriend(address _friend) returns()
func (_Redpkg *RedpkgTransactor) AddFriend(opts *bind.TransactOpts, _friend common.Address) (*types.Transaction, error) {
	return _Redpkg.contract.Transact(opts, "AddFriend", _friend)
}

// AddFriend is a paid mutator transaction binding the contract method 0x13962258.
//
// Solidity: function AddFriend(address _friend) returns()
func (_Redpkg *RedpkgSession) AddFriend(_friend common.Address) (*types.Transaction, error) {
	return _Redpkg.Contract.AddFriend(&_Redpkg.TransactOpts, _friend)
}

// AddFriend is a paid mutator transaction binding the contract method 0x13962258.
//
// Solidity: function AddFriend(address _friend) returns()
func (_Redpkg *RedpkgTransactorSession) AddFriend(_friend common.Address) (*types.Transaction, error) {
	return _Redpkg.Contract.AddFriend(&_Redpkg.TransactOpts, _friend)
}

// AddFriendList is a paid mutator transaction binding the contract method 0x2bba4dde.
//
// Solidity: function AddFriendList(address[] list) returns()
func (_Redpkg *RedpkgTransactor) AddFriendList(opts *bind.TransactOpts, list []common.Address) (*types.Transaction, error) {
	return _Redpkg.contract.Transact(opts, "AddFriendList", list)
}

// AddFriendList is a paid mutator transaction binding the contract method 0x2bba4dde.
//
// Solidity: function AddFriendList(address[] list) returns()
func (_Redpkg *RedpkgSession) AddFriendList(list []common.Address) (*types.Transaction, error) {
	return _Redpkg.Contract.AddFriendList(&_Redpkg.TransactOpts, list)
}

// AddFriendList is a paid mutator transaction binding the contract method 0x2bba4dde.
//
// Solidity: function AddFriendList(address[] list) returns()
func (_Redpkg *RedpkgTransactorSession) AddFriendList(list []common.Address) (*types.Transaction, error) {
	return _Redpkg.Contract.AddFriendList(&_Redpkg.TransactOpts, list)
}

// DelFriend is a paid mutator transaction binding the contract method 0x5663132e.
//
// Solidity: function DelFriend(address _friend) returns()
func (_Redpkg *RedpkgTransactor) DelFriend(opts *bind.TransactOpts, _friend common.Address) (*types.Transaction, error) {
	return _Redpkg.contract.Transact(opts, "DelFriend", _friend)
}

// DelFriend is a paid mutator transaction binding the contract method 0x5663132e.
//
// Solidity: function DelFriend(address _friend) returns()
func (_Redpkg *RedpkgSession) DelFriend(_friend common.Address) (*types.Transaction, error) {
	return _Redpkg.Contract.DelFriend(&_Redpkg.TransactOpts, _friend)
}

// DelFriend is a paid mutator transaction binding the contract method 0x5663132e.
//
// Solidity: function DelFriend(address _friend) returns()
func (_Redpkg *RedpkgTransactorSession) DelFriend(_friend common.Address) (*types.Transaction, error) {
	return _Redpkg.Contract.DelFriend(&_Redpkg.TransactOpts, _friend)
}

// Giving is a paid mutator transaction binding the contract method 0x2a43e999.
//
// Solidity: function Giving(string word, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) returns()
func (_Redpkg *RedpkgTransactor) Giving(opts *bind.TransactOpts, word string, equalDivision bool, onlyFriend bool, size *big.Int, expireDays *big.Int) (*types.Transaction, error) {
	return _Redpkg.contract.Transact(opts, "Giving", word, equalDivision, onlyFriend, size, expireDays)
}

// Giving is a paid mutator transaction binding the contract method 0x2a43e999.
//
// Solidity: function Giving(string word, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) returns()
func (_Redpkg *RedpkgSession) Giving(word string, equalDivision bool, onlyFriend bool, size *big.Int, expireDays *big.Int) (*types.Transaction, error) {
	return _Redpkg.Contract.Giving(&_Redpkg.TransactOpts, word, equalDivision, onlyFriend, size, expireDays)
}

// Giving is a paid mutator transaction binding the contract method 0x2a43e999.
//
// Solidity: function Giving(string word, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) returns()
func (_Redpkg *RedpkgTransactorSession) Giving(word string, equalDivision bool, onlyFriend bool, size *big.Int, expireDays *big.Int) (*types.Transaction, error) {
	return _Redpkg.Contract.Giving(&_Redpkg.TransactOpts, word, equalDivision, onlyFriend, size, expireDays)
}

// Grabbing is a paid mutator transaction binding the contract method 0x1f1dc595.
//
// Solidity: function Grabbing(string word) returns()
func (_Redpkg *RedpkgTransactor) Grabbing(opts *bind.TransactOpts, word string) (*types.Transaction, error) {
	return _Redpkg.contract.Transact(opts, "Grabbing", word)
}

// Grabbing is a paid mutator transaction binding the contract method 0x1f1dc595.
//
// Solidity: function Grabbing(string word) returns()
func (_Redpkg *RedpkgSession) Grabbing(word string) (*types.Transaction, error) {
	return _Redpkg.Contract.Grabbing(&_Redpkg.TransactOpts, word)
}

// Grabbing is a paid mutator transaction binding the contract method 0x1f1dc595.
//
// Solidity: function Grabbing(string word) returns()
func (_Redpkg *RedpkgTransactorSession) Grabbing(word string) (*types.Transaction, error) {
	return _Redpkg.Contract.Grabbing(&_Redpkg.TransactOpts, word)
}

// NewTokenReceiver is a paid mutator transaction binding the contract method 0x5daccf51.
//
// Solidity: function NewTokenReceiver() returns()
func (_Redpkg *RedpkgTransactor) NewTokenReceiver(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redpkg.contract.Transact(opts, "NewTokenReceiver")
}

// NewTokenReceiver is a paid mutator transaction binding the contract method 0x5daccf51.
//
// Solidity: function NewTokenReceiver() returns()
func (_Redpkg *RedpkgSession) NewTokenReceiver() (*types.Transaction, error) {
	return _Redpkg.Contract.NewTokenReceiver(&_Redpkg.TransactOpts)
}

// NewTokenReceiver is a paid mutator transaction binding the contract method 0x5daccf51.
//
// Solidity: function NewTokenReceiver() returns()
func (_Redpkg *RedpkgTransactorSession) NewTokenReceiver() (*types.Transaction, error) {
	return _Redpkg.Contract.NewTokenReceiver(&_Redpkg.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x03723885.
//
// Solidity: function WithdrawToken(address token, address _to, uint256 value) returns()
func (_Redpkg *RedpkgTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, _to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Redpkg.contract.Transact(opts, "WithdrawToken", token, _to, value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x03723885.
//
// Solidity: function WithdrawToken(address token, address _to, uint256 value) returns()
func (_Redpkg *RedpkgSession) WithdrawToken(token common.Address, _to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Redpkg.Contract.WithdrawToken(&_Redpkg.TransactOpts, token, _to, value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x03723885.
//
// Solidity: function WithdrawToken(address token, address _to, uint256 value) returns()
func (_Redpkg *RedpkgTransactorSession) WithdrawToken(token common.Address, _to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Redpkg.Contract.WithdrawToken(&_Redpkg.TransactOpts, token, _to, value)
}
