package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan int)

func Deposit(amount int) {
    deposits <- amount
}

func Balance() int {
    return <-balances
}

func WithDraw(amount int) {
    withdraws <- amount
}

func teller() {
    var balance int
    for {
        select {
        case balances <- balance:
        case amount := <-deposits:
            balance += amount
        }
    }

}
func init() {
    go teller()
}
