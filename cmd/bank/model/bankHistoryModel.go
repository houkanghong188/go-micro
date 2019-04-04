package model

type HistoryModel struct {
	Id                    int32
	Type                  int32
	TransactionMoney      int32
	Detail                string
	AfterTransactionMoney int32
	Status                int32
	BankId                int32
	HandlingFee           int32
	SiteId                int
	IsDelete              int
	CreateTime            string
	UpdateTime            string
	IsInvoice             int32
}

func NewBankHistoryModel() *HistoryModel {
	return &HistoryModel{}
}

func (m *HistoryModel) TableName() string {
	return "platv5_website_bank_history"
}
