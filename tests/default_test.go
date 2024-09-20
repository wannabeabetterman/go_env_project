package tests

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	redisPkg "github.com/zeromicro/go-zero/core/stores/redis"
)

func init() {

}

type User struct {
	Id   uint64
	Name string
	Age  uint64
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "test_user"
}

func TestRedis(t *testing.T) {
	redis := redisPkg.New("8.136.104.221:6379", func(r *redisPkg.Redis) {
		r.Type = "node"
		r.Pass = "yvJR0jf8PH"
	})
	value, err := redis.Llen("key12323344")
	if err != nil {
		return
	}
	if value == 0 {
		_, err = redis.Lpush("key12323344", time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			return
		}
	}
	fmt.Println(value)
	value, err = redis.Llen("key12323344")
	if err != nil {
		return
	}
	if value > 0 {
		for i := 0; i <= 10; i++ {
			_, err := redis.Lpush("key12323344", i)
			if err != nil {
				return
			}
		}
	}
	dataTimeString, err := redis.Lrange("key12323344", 0, -2)
	if err != nil {
		return
	}
	fmt.Println(dataTimeString)
	storeDataStrings, err := redis.Lrange("key12323344", -1, -1)
	if err != nil {
		return
	}
	fmt.Println(storeDataStrings)
	redis.Del("key12323344")
}

func addUser(user *[]User) {

	*user = append(*user, User{
		Id:   1,
		Name: "哈哈",
		Age:  2,
	})
}

func TestLongitude(t *testing.T) {
	var mySlice []interface{}

	// 向切片添加不同类型的元素
	mySlice = append(mySlice, 42, "Hello", 3.14, true)

	// 打印切片的内容
	fmt.Println("切片的内容：", mySlice)

	// 遍历切片并使用类型断言获取元素的实际类型
	for _, v := range mySlice {
		switch value := v.(type) {
		case int:
			fmt.Println("这是一个整数:", value)
		case string:
			fmt.Println("这是一个字符串:", value)
		case float64:
			fmt.Println("这是一个浮点数:", value)
		case bool:
			fmt.Println("这是一个布尔值:", value)
		default:
			fmt.Println("未知类型")
		}
	}
	//var listString [][]string
	//newRow := []string{"NewPerson1", "NewPerson2", "NewPerson3"}
	//listString = append(listString, newRow)
	//fmt.Println(listString[0][0])
	//
	//var dataList []interface{}
	//dataList = append(dataList, newRow)
	//fmt.Println(dataList[0])
	//stringStr := ""
	//
	//c, _ := strconv.Atoi(stringStr)
	//fmt.Sprint(c)
	//var user []User
	//addUser(&user)
	//
	//_, ok := linq.From(user).FirstWithT(func(x User) bool {
	//	return x.Id == 1
	//}).(User)
	//
	//if ok {
	//	fmt.Println("1112323")
	//}
	//dataTime, _ := time.Parse("2006-01-02 15:04:05", "2023-11-17 11:13:03")
	//nowTimeString := time.Now().Format("2006-01-02 15:04:05")
	//nowTime, _ := time.Parse("2006-01-02 15:04:05", nowTimeString)
	//fmt.Println(dataTime.Before(nowTime))

	//var e, f, j string
	//e, f, j = "", "", ""
}

func TestKafka(t *testing.T) {
	brokerAddress := "101.37.25.242:9093"
	topic := "dbss_iot"
	username := "alikafka_post-cn-st21vo8lh004"
	password := "C2PFNLXEQXUHixAKiBtsakuYCXE48Ems"

	// Set up SASL authentication configuration
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		SASLMechanism: &plain.Mechanism{
			Username: username,
			Password: password,
		},
		TLS: &tls.Config{InsecureSkipVerify: true},
	}

	// Create a new Kafka writer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		Dialer:   dialer,
	})

	defer writer.Close()
	// Create a new Kafka message
	message := kafka.Message{
		Key:   []byte("key"),
		Value: []byte("{\"tenantCode\":\"T1713812628567126016\",\"name\":\"510101001006\",\"id\":\"510101001006\",\"events\":[{\"extension\":{\"replyId\":\"PLATFORM_REPLY_PILE_STATE_V2_TO_PILE_ID\",\"timestamp\":1706062170575},\"content\":{\"connectorNo\":1,\"equipmentCode\":\"510101001006\",\"connectorWorkStatus\":\"charging\",\"alarmCode\":0,\"connectorInsertStatus\":\"inserted\"},\"name\":\"[设备 => 枪状态信息 => 平台]\",\"id\":\"CONNECTOR_STATUS_TO_PLATFORM_ID\"},{\"extension\":{\"protocolName\":\"bc-v30\",\"timestamp\":1706062170575},\"content\":{\"elecMeterDegree\":2412,\"cardBalanceBeforeCharge\":999999,\"doorState\":0,\"soc\":49,\"totalDegree\":315,\"plugTemperature\":8,\"acaVolt\":353,\"chargeOrderNo\":\"C1749975465201840128\",\"chargeBootMode\":1,\"plugWorkState\":2,\"cardNo\":\"\",\"acbVolt\":353,\"beforeChargeElecMeterDegree\":2097,\"ele\":68.2,\"plugInCar\":1,\"chargePower\":241,\"orderNo\":\"C1749975465201840128\",\"bmsNeedCurrent\":686,\"totalTime\":504,\"degree\":0.315,\"plugNo\":1,\"cumulativeChargeAmount\":630,\"reserveFlag\":0,\"outletTemperature\":-50,\"pileSn\":\"510101001006\",\"acbCurrent\":68,\"sysVar3\":0,\"remainChargeTime\":52,\"sysVar4\":0,\"reportTime\":\"20240124100930\",\"bmsChargeMode\":0,\"carConnectState\":2,\"duration\":504,\"parameterVar\":0,\"vol\":353.6,\"currentMaxAlertCode\":0,\"currentSoc\":49,\"startTime\":\"20240124100104\",\"chargeStrategy\":0,\"power\":24.1,\"bmsNeedVolt\":3796,\"dcChargeVoltage\":3536,\"accVolt\":353,\"reserveChargeTime\":\"20240124100104\",\"ambientTemperature\":21,\"plugCount\":2,\"internalVar2\":0,\"internalVar3\":0,\"dcChargeCurrent\":682,\"plugType\":1,\"chargeStrategyParam\":0,\"carVinCode\":\"\",\"acaCurrent\":68,\"accCurrent\":68,\"reserveTimeout\":0},\"name\":\"[设备 => 充电中信息 => 平台]\",\"id\":\"CHARGING_DATA_TO_PLATFORM_ID\"}]}"),
	}

	// Produce the message to Kafka
	err := writer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Fatalf("Error writing message to Kafka: %v", err)
	}

	fmt.Println("Message sent to Kafka successfully")

}

func TestXxx(t *testing.T) {
	var student Student
	buildStudent(&student)
	fmt.Println(student)
}

func buildStudent(student *Student) {
	student.Name = "张三"
	student.Sex = "男"

}

type Student struct {
	Name string
	Sex  string
}
