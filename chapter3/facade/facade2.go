package main

import (
	"fmt"
	"log"
)

// 账户
type Account struct {
	name string
}

// 创建账户
func NewAccount(accountName string) *Account {
	return &Account{name: accountName}
}

// 检查账户
func (a *Account) CheckAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("%s", "账户名不正确～")
	}
	fmt.Println("账户验证通过～")
	return nil
}

// 分类帐
type Ledger struct {
}

// 生成分类帐条目
func (s *Ledger) MakeEntry(accountID, txnType string, amount int) {
	fmt.Printf("为账户：%s 生成分类帐条目，账目类型为：%s，金额为：%d\n", accountID, txnType, amount)
}

// 通知
type Notification struct {
}

// 发送信用通知
func (n *Notification) SendWalletCreditNotification() {
	fmt.Println("发送钱包信用通知...")
}

// 发送借款通知
func (n *Notification) SendWalletDebitNotification() {
	fmt.Println("发送钱包借款通知...")
}

// 验证码
type VerificationCode struct {
	code int
}

// 创建验证码
func NewVerificationCode(code int) *VerificationCode {
	return &VerificationCode{code: code}
}

// 检查验证码
func (s *VerificationCode) CheckCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("%s", "验证码不正确")
	}
	fmt.Println("验证通过～")
	return nil
}

// 钱包
type Wallet struct {
	balance int
}

// 创建钱包
func NewWallet() *Wallet {
	return &Wallet{balance: 0}
}

// 添加金额
func (w *Wallet) AddBalance(amount int) {
	w.balance += amount
	fmt.Println("添加钱包金额成功～")
}

// 借款金额
func (w *Wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("%s", "金额无效～")
	}
	fmt.Println("钱包金额足够～")
	w.balance -= amount
	return nil
}

// 定义钱包的外观类
type WalletFacade struct {
	Account          *Account
	Wallet           *Wallet
	VerificationCode *VerificationCode
	Notification     *Notification
	Ledger           *Ledger
}

// 创建钱包的外观类
func NewWalletFacade(accountID string, code int) *WalletFacade {
	WalletFacacde := &WalletFacade{
		Account:          NewAccount(accountID),
		VerificationCode: NewVerificationCode(code),
		Wallet:           NewWallet(),
		Notification:     &Notification{},
		Ledger:           &Ledger{},
	}
	return WalletFacacde
}

// 添加钱到钱包
func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("添加钱到钱包")
	//1.检查账户
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	//2.检查验证码
	err = w.VerificationCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	//3.添加金额
	w.Wallet.AddBalance(amount)
	//4.发送信用通知
	w.Notification.SendWalletCreditNotification()
	w.Ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

// 从钱包里扣款
func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("从钱包里扣款")
	//1.检查账户
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	//2.检查验证码
	err = w.VerificationCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	//3.借款金额
	err = w.Wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	//4.发送借款通知
	w.Notification.SendWalletDebitNotification()
	w.Ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func main() {
	//实例化外观模式
	WalletFacade := NewWalletFacade("barry", 1688)
	fmt.Println()

	//添加16元到钱包
	err := WalletFacade.AddMoneyToWallet("barry", 1688, 16)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	//从钱包取出5元
	err = WalletFacade.DeductMoneyFromWallet("barry", 1688, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
