# Go API client for metafab

Complete MetaFab API references and guides can be found at: https://trymetafab.com

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.5.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://trymetafab.com](https://trymetafab.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import metafab "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), metafab.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), metafab.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), metafab.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), metafab.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.trymetafab.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ContractsApi* | [**CreateContract**](docs/ContractsApi.md#createcontract) | **Post** /v1/contracts | Create custom contract
*ContractsApi* | [**GetContracts**](docs/ContractsApi.md#getcontracts) | **Get** /v1/contracts | Get contracts
*ContractsApi* | [**ReadContract**](docs/ContractsApi.md#readcontract) | **Get** /v1/contracts/{contractId}/reads | Read contract data
*ContractsApi* | [**TransferContractOwnership**](docs/ContractsApi.md#transfercontractownership) | **Post** /v1/contracts/{contractId}/owners | Transfer contract ownership
*ContractsApi* | [**UpgradeContractTrustedForwarder**](docs/ContractsApi.md#upgradecontracttrustedforwarder) | **Post** /v1/contracts/{contractId}/forwarders | Upgrade contract trusted forwarder
*ContractsApi* | [**WriteContract**](docs/ContractsApi.md#writecontract) | **Post** /v1/contracts/{contractId}/writes | Write contract data
*CurrenciesApi* | [**BatchTransferCurrency**](docs/CurrenciesApi.md#batchtransfercurrency) | **Post** /v1/currencies/{currencyId}/batchTransfers | Batch transfer currency
*CurrenciesApi* | [**BurnCurrency**](docs/CurrenciesApi.md#burncurrency) | **Post** /v1/currencies/{currencyId}/burns | Burn currency
*CurrenciesApi* | [**CreateCurrency**](docs/CurrenciesApi.md#createcurrency) | **Post** /v1/currencies | Create currency
*CurrenciesApi* | [**GetCurrencies**](docs/CurrenciesApi.md#getcurrencies) | **Get** /v1/currencies | Get currencies
*CurrenciesApi* | [**GetCurrencyBalance**](docs/CurrenciesApi.md#getcurrencybalance) | **Get** /v1/currencies/{currencyId}/balances | Get currency balance
*CurrenciesApi* | [**GetCurrencyFees**](docs/CurrenciesApi.md#getcurrencyfees) | **Get** /v1/currencies/{currencyId}/fees | Get currency fees
*CurrenciesApi* | [**GetCurrencyRole**](docs/CurrenciesApi.md#getcurrencyrole) | **Get** /v1/currencies/{currencyId}/roles | Get currency role
*CurrenciesApi* | [**GrantCurrencyRole**](docs/CurrenciesApi.md#grantcurrencyrole) | **Post** /v1/currencies/{currencyId}/roles | Grant currency role
*CurrenciesApi* | [**MintCurrency**](docs/CurrenciesApi.md#mintcurrency) | **Post** /v1/currencies/{currencyId}/mints | Mint currency
*CurrenciesApi* | [**RevokeCurrencyRole**](docs/CurrenciesApi.md#revokecurrencyrole) | **Delete** /v1/currencies/{currencyId}/roles | Revoke currency role
*CurrenciesApi* | [**SetCurrencyFees**](docs/CurrenciesApi.md#setcurrencyfees) | **Post** /v1/currencies/{currencyId}/fees | Set currency fees
*CurrenciesApi* | [**TransferCurrency**](docs/CurrenciesApi.md#transfercurrency) | **Post** /v1/currencies/{currencyId}/transfers | Transfer currency
*EcosystemsApi* | [**ApproveEcosystemGame**](docs/EcosystemsApi.md#approveecosystemgame) | **Post** /v1/ecosystems/{ecosystemId}/games | Approve ecosystem game
*EcosystemsApi* | [**AuthEcosystem**](docs/EcosystemsApi.md#authecosystem) | **Get** /v1/ecosystems/auth | Authenticate ecosystem
*EcosystemsApi* | [**AuthProfile**](docs/EcosystemsApi.md#authprofile) | **Get** /v1/profiles/auth | Authenticate profile
*EcosystemsApi* | [**AuthProfilePlayer**](docs/EcosystemsApi.md#authprofileplayer) | **Get** /v1/profiles/{profileId}/games/{gameId}/players/auth | Authenticate profile player
*EcosystemsApi* | [**CreateEcosystem**](docs/EcosystemsApi.md#createecosystem) | **Post** /v1/ecosystems | Create ecosystem
*EcosystemsApi* | [**CreateProfile**](docs/EcosystemsApi.md#createprofile) | **Post** /v1/profiles | Create profile
*EcosystemsApi* | [**CreateProfilePlayer**](docs/EcosystemsApi.md#createprofileplayer) | **Post** /v1/profiles/{profileId}/games/{gameId}/players | Create profile player
*EcosystemsApi* | [**GetEcosystem**](docs/EcosystemsApi.md#getecosystem) | **Get** /v1/ecosystems/{ecosystemId} | Get ecosystem
*EcosystemsApi* | [**GetEcosystemGame**](docs/EcosystemsApi.md#getecosystemgame) | **Get** /v1/ecosystems/{ecosystemId}/games/{gameId} | Get ecosystem game
*EcosystemsApi* | [**GetEcosystemGames**](docs/EcosystemsApi.md#getecosystemgames) | **Get** /v1/ecosystems/{ecosystemId}/games | Get ecosystem games
*EcosystemsApi* | [**GetProfileGame**](docs/EcosystemsApi.md#getprofilegame) | **Get** /v1/profiles/{profileId}/games/{gameId} | Get profile game
*EcosystemsApi* | [**GetProfileGames**](docs/EcosystemsApi.md#getprofilegames) | **Get** /v1/profiles/{profileId}/games | Get profile games
*EcosystemsApi* | [**UnapproveEcosystemGame**](docs/EcosystemsApi.md#unapproveecosystemgame) | **Delete** /v1/ecosystems/{ecosystemId}/games/{gameId} | Unapprove ecosystem game
*EcosystemsApi* | [**UpdateEcosystem**](docs/EcosystemsApi.md#updateecosystem) | **Patch** /v1/ecosystems/{ecosystemId} | Update ecosystem
*EcosystemsApi* | [**UpdateProfile**](docs/EcosystemsApi.md#updateprofile) | **Patch** /v1/profiles/{profileId} | Update profile
*EcosystemsApi* | [**UpdateProfilePlayer**](docs/EcosystemsApi.md#updateprofileplayer) | **Patch** /v1/profiles/{profileId}/games/{gameId}/players/{playerId} | Update profile player
*GamesApi* | [**AuthGame**](docs/GamesApi.md#authgame) | **Get** /v1/games/auth | Authenticate game
*GamesApi* | [**CreateGame**](docs/GamesApi.md#creategame) | **Post** /v1/games | Create game
*GamesApi* | [**GetGame**](docs/GamesApi.md#getgame) | **Get** /v1/games/{gameId} | Get game
*GamesApi* | [**UpdateGame**](docs/GamesApi.md#updategame) | **Patch** /v1/games/{gameId} | Update game
*ItemsApi* | [**BatchMintCollectionItems**](docs/ItemsApi.md#batchmintcollectionitems) | **Post** /v1/collections/{collectionId}/batchMints | Batch mint collection items
*ItemsApi* | [**BatchTransferCollectionItems**](docs/ItemsApi.md#batchtransfercollectionitems) | **Post** /v1/collections/{collectionId}/batchTransfers | Batch transfer collection items
*ItemsApi* | [**BurnCollectionItem**](docs/ItemsApi.md#burncollectionitem) | **Post** /v1/collections/{collectionId}/items/{collectionItemId}/burns | Burn collection item
*ItemsApi* | [**CreateCollection**](docs/ItemsApi.md#createcollection) | **Post** /v1/collections | Create collection
*ItemsApi* | [**CreateCollectionItem**](docs/ItemsApi.md#createcollectionitem) | **Post** /v1/collections/{collectionId}/items | Create collection item
*ItemsApi* | [**GetCollectionApproval**](docs/ItemsApi.md#getcollectionapproval) | **Get** /v1/collections/{collectionId}/approvals | Get collection approval
*ItemsApi* | [**GetCollectionItem**](docs/ItemsApi.md#getcollectionitem) | **Get** /v1/collections/{collectionId}/items/{collectionItemId} | Get collection item
*ItemsApi* | [**GetCollectionItemBalance**](docs/ItemsApi.md#getcollectionitembalance) | **Get** /v1/collections/{collectionId}/items/{collectionItemId}/balances | Get collection item balance
*ItemsApi* | [**GetCollectionItemBalances**](docs/ItemsApi.md#getcollectionitembalances) | **Get** /v1/collections/{collectionId}/balances | Get collection item balances
*ItemsApi* | [**GetCollectionItemSupplies**](docs/ItemsApi.md#getcollectionitemsupplies) | **Get** /v1/collections/{collectionId}/supplies | Get collection item supplies
*ItemsApi* | [**GetCollectionItemSupply**](docs/ItemsApi.md#getcollectionitemsupply) | **Get** /v1/collections/{collectionId}/items/{collectionItemId}/supplies | Get collection item supply
*ItemsApi* | [**GetCollectionItemTimelock**](docs/ItemsApi.md#getcollectionitemtimelock) | **Get** /v1/collections/{collectionId}/items/{collectionItemId}/timelocks | Get collection item timelock
*ItemsApi* | [**GetCollectionItems**](docs/ItemsApi.md#getcollectionitems) | **Get** /v1/collections/{collectionId}/items | Get collection items
*ItemsApi* | [**GetCollectionRole**](docs/ItemsApi.md#getcollectionrole) | **Get** /v1/collections/{collectionId}/roles | Get collection role
*ItemsApi* | [**GetCollections**](docs/ItemsApi.md#getcollections) | **Get** /v1/collections | Get collections
*ItemsApi* | [**GrantCollectionRole**](docs/ItemsApi.md#grantcollectionrole) | **Post** /v1/collections/{collectionId}/roles | Grant collection role
*ItemsApi* | [**MintCollectionItem**](docs/ItemsApi.md#mintcollectionitem) | **Post** /v1/collections/{collectionId}/items/{collectionItemId}/mints | Mint collection item
*ItemsApi* | [**RevokeCollectionRole**](docs/ItemsApi.md#revokecollectionrole) | **Delete** /v1/collections/{collectionId}/roles | Revoke collection role
*ItemsApi* | [**SetCollectionApproval**](docs/ItemsApi.md#setcollectionapproval) | **Post** /v1/collections/{collectionId}/approvals | Set collection approval
*ItemsApi* | [**SetCollectionItemTimelock**](docs/ItemsApi.md#setcollectionitemtimelock) | **Post** /v1/collections/{collectionId}/items/{collectionItemId}/timelocks | Set collection item timelock
*ItemsApi* | [**TransferCollectionItem**](docs/ItemsApi.md#transfercollectionitem) | **Post** /v1/collections/{collectionId}/items/{collectionItemId}/transfers | Transfer collection item
*LootboxesApi* | [**CreateLootboxManager**](docs/LootboxesApi.md#createlootboxmanager) | **Post** /v1/lootboxManagers | Create lootbox manager
*LootboxesApi* | [**GetLootboxManagerLootbox**](docs/LootboxesApi.md#getlootboxmanagerlootbox) | **Get** /v1/lootboxManagers/{lootboxManagerId}/lootboxes/{lootboxManagerLootboxId} | Get lootbox manager lootbox
*LootboxesApi* | [**GetLootboxManagerLootboxes**](docs/LootboxesApi.md#getlootboxmanagerlootboxes) | **Get** /v1/lootboxManagers/{lootboxManagerId}/lootboxes | Get lootbox manager lootboxes
*LootboxesApi* | [**GetLootboxManagers**](docs/LootboxesApi.md#getlootboxmanagers) | **Get** /v1/lootboxManagers | Get lootbox managers
*LootboxesApi* | [**OpenLootboxManagerLootbox**](docs/LootboxesApi.md#openlootboxmanagerlootbox) | **Post** /v1/lootboxManagers/{lootboxManagerId}/lootboxes/{lootboxManagerLootboxId}/opens | Open lootbox manager lootbox
*LootboxesApi* | [**RemoveLootboxManagerLootbox**](docs/LootboxesApi.md#removelootboxmanagerlootbox) | **Delete** /v1/lootboxManagers/{lootboxManagerId}/lootboxes/{lootboxManagerLootboxId} | Remove lootbox manager lootbox
*LootboxesApi* | [**SetLootboxManagerLootbox**](docs/LootboxesApi.md#setlootboxmanagerlootbox) | **Post** /v1/lootboxManagers/{lootboxManagerId}/lootboxes | Set lootbox manager lootbox
*PlayersApi* | [**AuthPlayer**](docs/PlayersApi.md#authplayer) | **Get** /v1/players/auth | Authenticate player
*PlayersApi* | [**CreatePlayer**](docs/PlayersApi.md#createplayer) | **Post** /v1/players | Create player
*PlayersApi* | [**GetPlayer**](docs/PlayersApi.md#getplayer) | **Get** /v1/players/{playerId} | Get player
*PlayersApi* | [**GetPlayerData**](docs/PlayersApi.md#getplayerdata) | **Get** /v1/players/{playerId}/data | Get player data
*PlayersApi* | [**GetPlayers**](docs/PlayersApi.md#getplayers) | **Get** /v1/players | Get players
*PlayersApi* | [**RemovePlayerConnectedWallet**](docs/PlayersApi.md#removeplayerconnectedwallet) | **Delete** /v1/players/{playerId}/wallets/{playerWalletId} | Remove player connected wallet
*PlayersApi* | [**SetPlayerConnectedWallet**](docs/PlayersApi.md#setplayerconnectedwallet) | **Post** /v1/players/{playerId}/wallets | Set player connected wallet
*PlayersApi* | [**SetPlayerData**](docs/PlayersApi.md#setplayerdata) | **Post** /v1/players/{playerId}/data | Set player data
*PlayersApi* | [**UpdatePlayer**](docs/PlayersApi.md#updateplayer) | **Patch** /v1/players/{playerId} | Update player
*ShopsApi* | [**CreateShop**](docs/ShopsApi.md#createshop) | **Post** /v1/shops | Create shop
*ShopsApi* | [**GetShopOffer**](docs/ShopsApi.md#getshopoffer) | **Get** /v1/shops/{shopId}/offers/{shopOfferId} | Get shop offer
*ShopsApi* | [**GetShopOffers**](docs/ShopsApi.md#getshopoffers) | **Get** /v1/shops/{shopId}/offers | Get shop offers
*ShopsApi* | [**GetShops**](docs/ShopsApi.md#getshops) | **Get** /v1/shops | Get shops
*ShopsApi* | [**RemoveShopOffer**](docs/ShopsApi.md#removeshopoffer) | **Delete** /v1/shops/{shopId}/offers/{shopOfferId} | Remove shop offer
*ShopsApi* | [**SetShopOffer**](docs/ShopsApi.md#setshopoffer) | **Post** /v1/shops/{shopId}/offers | Set shop offer
*ShopsApi* | [**UseShopOffer**](docs/ShopsApi.md#useshopoffer) | **Post** /v1/shops/{shopId}/offers/{shopOfferId}/uses | Use shop offer
*ShopsApi* | [**WithdrawFromShop**](docs/ShopsApi.md#withdrawfromshop) | **Post** /v1/shops/{shopId}/withdrawals | Withdraw from shop
*TransactionsApi* | [**GetTransaction**](docs/TransactionsApi.md#gettransaction) | **Get** /v1/transactions/{transactionId} | Get transaction
*WalletsApi* | [**CreateWalletSignature**](docs/WalletsApi.md#createwalletsignature) | **Post** /v1/wallets/{walletId}/signatures | Create wallet signature
*WalletsApi* | [**GetWallet**](docs/WalletsApi.md#getwallet) | **Get** /v1/wallets/{walletId} | Get wallet
*WalletsApi* | [**GetWalletBalances**](docs/WalletsApi.md#getwalletbalances) | **Get** /v1/wallets/{walletId}/balances | Get wallet balances
*WalletsApi* | [**GetWalletTransactions**](docs/WalletsApi.md#getwallettransactions) | **Get** /v1/wallets/{walletId}/transactions | Get wallet transactions


## Documentation For Models

 - [ApproveEcosystemGameRequest](docs/ApproveEcosystemGameRequest.md)
 - [AuthGame200Response](docs/AuthGame200Response.md)
 - [AuthGame200ResponseAllOf](docs/AuthGame200ResponseAllOf.md)
 - [AuthPlayer200Response](docs/AuthPlayer200Response.md)
 - [AuthPlayer200ResponseAllOf](docs/AuthPlayer200ResponseAllOf.md)
 - [AuthProfile200Response](docs/AuthProfile200Response.md)
 - [BatchMintCollectionItemsRequest](docs/BatchMintCollectionItemsRequest.md)
 - [BatchTransferCollectionItemsRequest](docs/BatchTransferCollectionItemsRequest.md)
 - [BatchTransferCurrencyRequest](docs/BatchTransferCurrencyRequest.md)
 - [BurnCollectionItemRequest](docs/BurnCollectionItemRequest.md)
 - [BurnCurrencyRequest](docs/BurnCurrencyRequest.md)
 - [CollectionItem](docs/CollectionItem.md)
 - [CollectionItemAttributesInner](docs/CollectionItemAttributesInner.md)
 - [CollectionItemAttributesInnerValue](docs/CollectionItemAttributesInnerValue.md)
 - [CollectionModel](docs/CollectionModel.md)
 - [ContractModel](docs/ContractModel.md)
 - [CreateCollection200Response](docs/CreateCollection200Response.md)
 - [CreateCollection200ResponseAllOf](docs/CreateCollection200ResponseAllOf.md)
 - [CreateCollection200ResponseAllOfContract](docs/CreateCollection200ResponseAllOfContract.md)
 - [CreateCollection200ResponseAllOfContractAllOf](docs/CreateCollection200ResponseAllOfContractAllOf.md)
 - [CreateCollectionItemRequest](docs/CreateCollectionItemRequest.md)
 - [CreateCollectionItemRequestAttributesInner](docs/CreateCollectionItemRequestAttributesInner.md)
 - [CreateCollectionRequest](docs/CreateCollectionRequest.md)
 - [CreateContractRequest](docs/CreateContractRequest.md)
 - [CreateCurrency200Response](docs/CreateCurrency200Response.md)
 - [CreateCurrencyRequest](docs/CreateCurrencyRequest.md)
 - [CreateEcosystemRequest](docs/CreateEcosystemRequest.md)
 - [CreateGameRequest](docs/CreateGameRequest.md)
 - [CreateLootboxManager200Response](docs/CreateLootboxManager200Response.md)
 - [CreateLootboxManagerRequest](docs/CreateLootboxManagerRequest.md)
 - [CreatePlayerRequest](docs/CreatePlayerRequest.md)
 - [CreateProfilePlayerRequest](docs/CreateProfilePlayerRequest.md)
 - [CreateProfileRequest](docs/CreateProfileRequest.md)
 - [CreateShop200Response](docs/CreateShop200Response.md)
 - [CreateShopRequest](docs/CreateShopRequest.md)
 - [CreateWalletSignatureRequest](docs/CreateWalletSignatureRequest.md)
 - [CurrencyModel](docs/CurrencyModel.md)
 - [EcosystemGameModel](docs/EcosystemGameModel.md)
 - [EcosystemModel](docs/EcosystemModel.md)
 - [GameModel](docs/GameModel.md)
 - [GetCollections200ResponseInner](docs/GetCollections200ResponseInner.md)
 - [GetCollections200ResponseInnerAllOf](docs/GetCollections200ResponseInnerAllOf.md)
 - [GetCurrencies200ResponseInner](docs/GetCurrencies200ResponseInner.md)
 - [GetCurrencyFees200Response](docs/GetCurrencyFees200Response.md)
 - [GetLootboxManagers200ResponseInner](docs/GetLootboxManagers200ResponseInner.md)
 - [GetPlayerData200Response](docs/GetPlayerData200Response.md)
 - [GetProfileGames200ResponseInner](docs/GetProfileGames200ResponseInner.md)
 - [GetShops200ResponseInner](docs/GetShops200ResponseInner.md)
 - [GrantCollectionRoleRequest](docs/GrantCollectionRoleRequest.md)
 - [GrantCurrencyRoleRequest](docs/GrantCurrencyRoleRequest.md)
 - [LootboxManagerLootbox](docs/LootboxManagerLootbox.md)
 - [LootboxManagerModel](docs/LootboxManagerModel.md)
 - [MintCollectionItemRequest](docs/MintCollectionItemRequest.md)
 - [MintCurrencyRequest](docs/MintCurrencyRequest.md)
 - [PlayerModel](docs/PlayerModel.md)
 - [ProfileModel](docs/ProfileModel.md)
 - [ProfilePermissionsValue](docs/ProfilePermissionsValue.md)
 - [PublicEcosystem](docs/PublicEcosystem.md)
 - [PublicGame](docs/PublicGame.md)
 - [PublicPlayer](docs/PublicPlayer.md)
 - [PublicPlayerCustodialWallet](docs/PublicPlayerCustodialWallet.md)
 - [PublicProfile](docs/PublicProfile.md)
 - [RemovePlayerConnectedWalletRequest](docs/RemovePlayerConnectedWalletRequest.md)
 - [RevokeCollectionRoleRequest](docs/RevokeCollectionRoleRequest.md)
 - [SetCollectionApprovalRequest](docs/SetCollectionApprovalRequest.md)
 - [SetCollectionItemTimelockRequest](docs/SetCollectionItemTimelockRequest.md)
 - [SetCurrencyFeesRequest](docs/SetCurrencyFeesRequest.md)
 - [SetLootboxManagerLootboxRequest](docs/SetLootboxManagerLootboxRequest.md)
 - [SetPlayerConnectedWallet200Response](docs/SetPlayerConnectedWallet200Response.md)
 - [SetPlayerConnectedWalletRequest](docs/SetPlayerConnectedWalletRequest.md)
 - [SetPlayerDataRequest](docs/SetPlayerDataRequest.md)
 - [SetShopOfferRequest](docs/SetShopOfferRequest.md)
 - [ShopModel](docs/ShopModel.md)
 - [ShopOffer](docs/ShopOffer.md)
 - [TransactionModel](docs/TransactionModel.md)
 - [TransferCollectionItemRequest](docs/TransferCollectionItemRequest.md)
 - [TransferContractOwnershipRequest](docs/TransferContractOwnershipRequest.md)
 - [TransferCurrencyRequest](docs/TransferCurrencyRequest.md)
 - [UpdateEcosystemRequest](docs/UpdateEcosystemRequest.md)
 - [UpdateGame200Response](docs/UpdateGame200Response.md)
 - [UpdateGame200ResponseAllOf](docs/UpdateGame200ResponseAllOf.md)
 - [UpdateGameRequest](docs/UpdateGameRequest.md)
 - [UpdatePlayer200Response](docs/UpdatePlayer200Response.md)
 - [UpdatePlayerRequest](docs/UpdatePlayerRequest.md)
 - [UpdateProfilePlayer200Response](docs/UpdateProfilePlayer200Response.md)
 - [UpdateProfilePlayer200ResponseAllOf](docs/UpdateProfilePlayer200ResponseAllOf.md)
 - [UpdateProfilePlayerRequest](docs/UpdateProfilePlayerRequest.md)
 - [UpdateProfileRequest](docs/UpdateProfileRequest.md)
 - [UpgradeContractTrustedForwarderRequest](docs/UpgradeContractTrustedForwarderRequest.md)
 - [WalletModel](docs/WalletModel.md)
 - [WithdrawFromShopRequest](docs/WithdrawFromShopRequest.md)
 - [WriteContractRequest](docs/WriteContractRequest.md)
 - [WriteContractRequestArgsInner](docs/WriteContractRequestArgsInner.md)


## Documentation For Authorization



### basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

metafabproject@gmail.com

