package valueobjects


type AccountType string

const (
	TypeAsset     AccountType = "ASSET"
	TypeLiability AccountType = "LIABILITY"
	TypeEquity    AccountType = "EQUITY"
	TypeExpense   AccountType = "EXPENSE"
	TypeRevenue   AccountType = "REVENUE"
)
