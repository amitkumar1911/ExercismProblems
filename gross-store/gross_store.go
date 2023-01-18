package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	// panic("Please implement the Units() function") 

    mp:=make(map[string]int)
    mp["quarter_of_a_dozen"]=3 
     mp["half_of_a_dozen"]=6
     mp["dozen"]=12
     mp["small_gross"]=120
     mp["gross"]=144
     mp["great_gross"]=1728

    return mp
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	// panic("Please implement the NewBill() function") 

    emptyBill:=make(map[string]int)
    return emptyBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// /panic("Please implement the AddItem() function")   

    

    _, ok:=units[unit] 
    if ok==false{
        return false
    }  

    bill[item]+=units[unit]
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the RemoveItem() function") 

    _, ok:=units[unit] 
    _, ok1:=bill[item] 
    newQuantity:=bill[item]-units[unit] 

    if ok==false || ok1==false || newQuantity<0 {
        return false
    } else if newQuantity==0{
    delete(bill, item) 
        return true
    }

    bill[item]-=units[unit] 
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	// panic("Please implement the GetItem() function") 

    _, ok:=bill[item] 

    if ok==false{
        return 0, false
    } 
return bill[item], true
}
