type Bank struct {
    balance []int64
}


func Constructor(balance []int64) Bank {
    return Bank {
        balance: balance,
    }
}


func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
    if account1 > len(this.balance) || account2 > len(this.balance) {
        return false
    }
    if account1 == 0 || account2 == 0 {
        return false
    }
    bal1 := this.balance[account1 -1]
    if bal1 >= money {
            this.balance[account1 - 1] = this.balance[account1 - 1] - money
            this.balance[account2 - 1] = this.balance[account2 - 1] + money
    } else {
            return false
    }
    return true
}


func (this *Bank) Deposit(account int, money int64) bool {
    if account > len(this.balance) {
        return false
    }
    if account == 0 {
        return false
    }
    this.balance[account - 1] = this.balance[account - 1] + money
    return true
}


func (this *Bank) Withdraw(account int, money int64) bool {
    if account > len(this.balance) {
        return false
    }
    if account == 0 {
        return false
    }
    bal := this.balance[account - 1]
    if bal >= money {
        bal = bal -money 
        this.balance[account - 1] = bal
    } else {
        return false
    }
    return true
}


/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */