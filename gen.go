package main

import "fmt"

func main() {
	var R = []string {"ID", "PID", "SID", "SN","PN","MID","MN"}
	for i,s:=range R{
		for j,v:=range R{
			if i!=j{
				PrintSQL(s,v)
			}
		}
	}
}

func PrintSQL(Att1 string, Att2 string){
	fmt.Printf("SELECT 'Project: %v --> %v' AS FD,CASE WHEN COUNT(*)=0 THEN 'MAY HOLD' ELSE 'does not hold' END AS VALIDTY FROM (SELECT P.%v FROM Project P GROUP BY P.%v HAVING COUNT(DISTINCT P.%v)>1)X;",Att1,Att2,Att1,Att1,Att2)
}
