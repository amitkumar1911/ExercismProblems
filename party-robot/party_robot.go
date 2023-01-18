package partyrobot
import "fmt"
// import "strconv"

// Welcome greets a person by name.
func Welcome(name string) string {
	// panic("Please implement the Welcome function")
    // var temp string=",";

    return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	// panic("Please implement the HappyBirthday function")

    // return "Happy birthday"+" "+name+"!"+" "+"You are now"+" "+age+" "+"years old!"
    return fmt.Sprintf("Happy birthday %s! You are now %d years old!",name,age)
    
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// panic("Please implement the AssignTable function")


    var ans string 
    // var temp string

    ans=fmt.Sprintf("Welcome to my party, %s!", name)
    ans=ans+"\n"

    // temp, _:=strconv.Atoi(table);

    // if len(temp)<3{
    //     for i:=0;i<3-len(temp);i++{
            
    //         temp="0"+temp
            
    //     }
    // }

    ans=ans+ fmt.Sprintf("You have been assigned to table %.3d. Your table is %s, exactly %.1f meters from here.",table,direction, distance)

    ans=ans+"\n"

    ans=ans+ fmt.Sprintf("You will be sitting next to %s.", neighbor);

    return ans

}
