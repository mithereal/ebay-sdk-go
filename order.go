package ebay


GetOrdersRequest struct {
	
		RequesterCredentials = struct {
		eBayAuthToken
		}
		
		CreateTimeFrom  
		CreateTimeTo  
		IncludeFinalValueFee           
		ListingType        
		ModTimeFrom   
		ModTimeTo
		NumberOfDays  
		OrderRole  
		OrderStatus  
		SortingOrder  
		DetailLevel
		ErrorLanguage
		MessageID
		OutputSelector
		Version
		WarningLevel
		

		Pagination  = struct {
		EntriesPerPage
		PageNumber
		}
		
		OrderIDArray = struct {
		OrderID
		}	
		
	} 
