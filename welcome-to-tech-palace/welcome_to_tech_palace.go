package techpalace


import "fmt"

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	// panic("Please implement the WelcomeMessage() function")


     customer= strings.ToUpper(customer);

    return "Welcome to the Tech Palace"+","+" "+customer
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// panic("Please implement the AddBorder() function") 

    var ans string
    for i:=0;i<numStarsPerLine;i++{
        ans=ans+"*";
    }

    ans=ans+"\n" 
    ans=ans+welcomeMsg 
    ans=ans+"\n" ;

      for i:=0;i<numStarsPerLine;i++{
        ans=ans+"*";
    }

    return ans
    
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// panic("Please implement the CleanupMessage() function")

    var ans string 

    ans=strings.Trim(oldMsg, "*");
    ans=strings.Trim(ans, "\n");
    // ans=strings.TrimLeft(ans," ");
    // ans=strings.TrimRight(ans, " ");
     ans=strings.TrimLeft(ans,"*");
    ans=strings.TrimRight(ans, "*");
    ans=strings.TrimLeft(ans," ");
    ans=strings.TrimRight(ans, " ");
    

    // ans=strings.TrimLeft(oldMsg, " ");

    fmt.Println(ans)

    return ans;

    
}
