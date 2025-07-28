package main

import "fmt"

func main() {
    // In ra màn hình dòng chữ chào mừng
    fmt.Println("Xin chào! Đây là chương trình Go đầu tiên của tôi!")
    
    // In thêm một vài ví dụ khác
    name := "Golang"
    fmt.Printf("Tôi đang học ngôn ngữ %s\n", name)
    
    // Sử dụng các kiểu in khác nhau
    fmt.Print("Dòng này không có xuống dòng")
    fmt.Print(" và tiếp tục trên cùng một dòng\n")
    
    // In với nhiều biến
    age := 2009
    fmt.Printf("%s ra đời vào năm %d\n", name, age)
}
