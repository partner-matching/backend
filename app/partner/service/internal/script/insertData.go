package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/partner-matching/backend/app/partner/service/internal/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/partner_matching?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	return db
}

// GetUserName 获取姓名
func GetUserName() (string, error) {
	resp, err := http.Get("https://v.api.aa1.cn/api/api-xingming/index.php")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	res := map[string]interface{}{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}

	return res["xingming"].(string), nil

}

// GetUserAccount 获取姓名
func GetUserAccount(name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://v.api.aa1.cn/api/api-fanyi-yd/index.php?msg=%s&type=1", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	res := map[string]interface{}{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}

	return strings.Replace(res["text"].(string), " ", "", -1), nil
}

// GetUserAvatar 获取头像
func GetUserAvatar(num int32) (string, error) {
	var gender string
	if num == 0 {
		gender = "男"
	} else {
		gender = "女"
	}

	resp, err := http.Get(fmt.Sprintf("https://api.uomg.com/api/rand.avatar?sort=动漫%s&format=json", gender))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	res := map[string]interface{}{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}

	return res["imgurl"].(string), nil
}

// GetUserSign 获取个性签名
func GetUserSign() (string, error) {
	resp, err := http.Get("https://v.api.aa1.cn/api/pyq/index.php?aa1=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	res := map[string]interface{}{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}

	return res["pyq"].(string), nil
}

// GetRandomGender 获取性别
func GetRandomGender() int32 {
	num := rand.Int()
	if num%2 == 0 {
		return 0
	}
	return 1
}

// GetRandomTags 获取标签
func GetRandomTags() string {
	tags := []string{
		"{\"Java\":1,\"C++\":1,\"Python\":1}",
		"{\"Golang\":1,\"Docker\":1,\"Rust\":1}",
		"{\"大一\":1,\"C++\":1,\"emo\":1}",
		"{\"随和\":1,\"程序员\":1,\"Python\":1}",
		"{\"伤心\":1,\"萌妹子\":1,\"Web3\":1}",
		"{\"开心\":1,\"萌妹子\":1,\"Docker\":1}",
		"{\"emo\":1,\"乐子人\":1,\"C++\":1}",
	}
	return tags[rand.Intn(7)]
}

func passwordMD5Hash(userAccount string) string {
	m := md5.New()
	m.Write([]byte(userAccount))
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	// 这里输入需要注入多少条数据
	number := 1000
	group := &sync.WaitGroup{}
	list := make([]*data.User, 0)
	rand.Seed(time.Now().UnixNano())

	db := NewDB()

	for i := 0; i < number; i++ {
		group.Add(1)
		go func() {
			defer group.Done()

			// 获取姓名
			name, err := GetUserName()
			if err != nil {
				log.Errorf("GetUserName: %s", err.Error())
				return
			}

			// 获取账户
			account, err := GetUserAccount(name)
			if err != nil {
				log.Errorf("GetUserAccount: %s, %s", err.Error(), name)
				return
			}

			gender := GetRandomGender()

			// 获取用户随机头像
			avatar, err := GetUserAvatar(gender)
			if err != nil {
				log.Errorf("GetUserAvatar: %s", err.Error())
				return
			}

			// 获取用户个性签名
			sign, err := GetUserSign()
			if err != nil {
				log.Errorf("GetUserSign: %s", err.Error())
				return
			}

			tags := GetRandomTags()

			list = append(list, &data.User{
				UserName:     name,
				UserPassword: passwordMD5Hash("12345678"),
				Gender:       gender,
				UserAccount:  account,
				AvatarUrl:    avatar,
				Profile:      sign,
				Tags:         tags,
				CreateTime:   time.Now(),
				UpdateTime:   time.Now(),
			})
		}()
	}
	group.Wait()

	// 数据批量插入
	err := db.Create(&list).Error
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("data insert successfully")

}
