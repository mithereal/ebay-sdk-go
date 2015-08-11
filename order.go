package ebay


type GetOrdersRequest struct {
		CreateTimeFrom   string
		CreateTimeTo   string
		IncludeFinalValueFee    string        
		ListingType         string
		ModTimeFrom    string
		ModTimeTo string
		NumberOfDays   string
		OrderRole   string
		OrderStatus   string
		SortingOrder   string
		DetailLevel string
		ErrorLanguage string
		MessageID string
		OutputSelector string
		Version string
		WarningLevel string
		
		
	} 
	
type	RequesterCredentials struct {
		eBayAuthToken string
		}

type	Pagination   struct {
		EntriesPerPage string
		PageNumber string
		}
		
type	OrderIDArray  struct {
		OrderID string
		}	
