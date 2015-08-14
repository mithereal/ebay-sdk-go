package ebay


type GetOrdersRequest struct {
		Xmlns           string `xml:"xmlns,attr"`
	
		RequesterCredentials RequesterCredentials 

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
		
		Pagination Pagination
		OrderIDArray OrderIDArray
	} 
	

type	Pagination   struct {
		EntriesPerPage string
		PageNumber string
		}
		
type	OrderIDArray  struct {
		OrderID string
		}	

func (p *Pagination) Setpagination(entriesperpage,pagenumber string) {
    p.EntriesPerPage = entriesperpage
    p.PageNumber = pagenumber
}

func (o *OrderIDArray) Getorderd(orderid string) {
    o.OrderID = orderid
    
}


type	RequesterCredentials struct {
		eBayAuthToken string 
		}
		
func (r *RequesterCredentials) SetToken(token string) {
    r.eBayAuthToken = token
}

