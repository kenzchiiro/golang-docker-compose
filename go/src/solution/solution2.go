
package solution

import (
    "strconv"
    "strings"
    "math"
)

func Month(s string) string{
    months := []string{"มกราคม","กุมภาพันธ์","มีนาคม","เมษายน","พฤษภาคม","มิถุนายน","กรกฎาคม","สิงหาคม","กันยายน","ตุลาคม","พฤศจิกายน","ธันวาคม"};
    m,_ := strconv.Atoi(s) 
    return months[m-1]
}

func Year(s string) string{
    y,_ := strconv.Atoi(s) 
    return strconv.Itoa(y+543) 
}

func DayOfWeek(d string,m string,y string) string{
    D,_ := strconv.ParseFloat(d,64)
    M,_ := strconv.ParseFloat(m,64)
    Y,_ := strconv.ParseFloat(y,64)
    mon := 0.00
    if mon = M; M>2{
       mon = 12+M
       Y--
    }
    Y = math.Mod(Y,100)
    C := Y/100.00
    W := (D + math.Floor((13*(mon+1))/5) + Y + math.Floor(Y/4) + math.Floor(C/4) + (5*C));
    W = math.Mod(W,7)
    DayOfWeeks := []string{"อาทิตย์","จันทร์","อังคาร","พุธ","พฤหัสบดี","ศุกร์","เสาร์"}
    return DayOfWeeks[int(W)]
}

func DateFormatter(date string) string{
    
    //date:="29/02/2020" 
    s := strings.Split(date, "/")
    result:= "วัน"+DayOfWeek(s[0],s[1],s[2])+"ที่ "+s[0]+" "+Month(s[1])+" "+ Year(s[2])

    return result
    
}

