package index

import "fmt"

// responses
type MasterchainInfo struct {
	Last  *Block `json:"last"`
	First *Block `json:"first"`
} // @name MasterchainInfo

type BlocksResponse struct {
	Blocks []Block `json:"blocks"`
} // @name BlocksResponse

type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
	AddressBook  AddressBook   `json:"address_book"`
} // @name TransactionsResponse

type MessagesResponse struct {
	Messages    []Message   `json:"messages"`
	AddressBook AddressBook `json:"address_book"`
} // @name MessagesResponse

type AccountStatesResponse struct {
	Accounts    []AccountStateFull `json:"accounts"`
	AddressBook AddressBook        `json:"address_book"`
} // @name AccountStatesResponse

type WalletStatesResponse struct {
	Wallets     []WalletState `json:"wallets"`
	AddressBook AddressBook   `json:"address_book"`
} // @name WalletStatesResponse

type JettonMastersResponse struct {
	Masters     []JettonMaster `json:"jetton_masters"`
	AddressBook AddressBook    `json:"address_book"`
} // @name JettonMastersResponse

type JettonWalletsResponse struct {
	Wallets     []JettonWallet `json:"jetton_wallets"`
	AddressBook AddressBook    `json:"address_book"`
} // @name JettonWalletsResponse

type JettonTransfersResponse struct {
	Transfers   []JettonTransfer `json:"jetton_transfers"`
	AddressBook AddressBook      `json:"address_book"`
} // @name JettonTransfersResponse

type JettonBurnsResponse struct {
	Burns       []JettonBurn `json:"jetton_burns"`
	AddressBook AddressBook  `json:"address_book"`
} // @name JettonBurnsResponse

type EventsResponse struct {
	Events      []Event     `json:"events"`
	AddressBook AddressBook `json:"address_book"`
} // @name EventsResponse

type ActionsResponse struct {
	Actions     []Action    `json:"actions"`
	AddressBook AddressBook `json:"address_book"`
} // @name ActionsResponse

// errors
type RequestError struct {
	Message string `json:"error"`
	Code    int    `json:"code"`
} // @name RequestError

func (r RequestError) Error() string {
	return fmt.Sprintf("Error %d: %s", r.Code, r.Message)
}
