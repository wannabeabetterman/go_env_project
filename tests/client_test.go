package tests

//import (
//	"context"
//	"fmt"
//	"google.golang.org/grpc"
//	"testing"
//)
//
//func TestClient(t *testing.T) {
//	svcCfg := fmt.Sprintf(`{"loadBalancingPolicy":"%s"}`, "round_robin")
//	conn, err := grpc.Dial("192.168.63.24:9081", grpc.WithInsecure(), grpc.WithDefaultServiceConfig(svcCfg))
//	defer conn.Close()
//	if err != nil {
//		t.Fatal(err)
//		return
//	}
//	client := pb.NewUserServiceClient(conn)
//	reply, err := client.UserAuth(context.Background(), &pb.UserAuthReq{
//		Token:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2OCwiVXNlcm5hbWUiOiIxOTkwOTA5MDgwOCIsIlRlbmFudENvZGUiOiJUNDE3Mzc1MDc5NDA0NDA1MDA4IiwiTGFuZyI6InpoIiwiZXhwIjoxNjU4ODkxOTkyLCJuYmYiOjE2NTg4MDQ1OTJ9.1SuKM8OACvNR-lHg8sXpLePwHUybkWHyPjf83CjyWQM",
//		Url:         "/user/api/v1/menu/getMenu",
//		Method:      "GET",
//		ProductCode: "",
//	})
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Println(reply)
//}
//
////func Test2(t *testing.T) {
////	// 读取 Excel 文件
////	excelFileName := "C:\\Users\\23231\\Desktop\\test.xlsx"
////	xlFile, err := xlsx.OpenFile(excelFileName)
////	if err != nil {
////		fmt.Println("无法打开 Excel 文件:", err)
////		return
////	}
////
////	// 遍历每个工作表
////	for _, sheet := range xlFile.Sheets {
////		fmt.Println("工作表名称:", sheet.Name)
////		// 遍历每行数据
////		for _, row := range sheet.Rows {
////			// 遍历每个单元格
////			for _, cell := range row.Cells {
////				text := cell.String()
////				fmt.Printf("%s\t", text)
////			}
////			fmt.Println() // 换行
////		}
////	}
////
////}
