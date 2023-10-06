package imageControllers

import (
	//"flag"
	"bufio"
	//"fmt"
	"image"
	"image/png"

	//"image/png"
	"log"
	"os"
	"whiteWall/app/services/studentServices/articalServices"
	"whiteWall/app/utils"

	"path/filepath"

	"github.com/gin-gonic/gin"
	//uuid "github.com/satori/go.uuid"
	//"github.com/gin-gonic/gin"
)

// "fmt"
// "io"
// "io/ioutil"
// "os"
// "path"
// "net/http"
// "errors"
// "strings"
//
//	type Client struct {
//		id     string
//		socket *websocket.Conn
//		send   chan []byte
//	}
//
//	type Image interface {
//		ColorModel() color.Model // 返回图片的颜色模型
//		Bounds() Rectangle       // 返回图片的长宽
//		At(x, y int) color.Color // 返回(x,y)像素点的颜色
//	  }
type CreateImageData struct {
	UserID  uint     `json:"user_id" `
	ImageIn *os.File `json:"image_in" type:"file"`
}

func CreateImage(c *gin.Context) {
	var data CreateImageData
	err := c.ShouldBind(&data)
	if err != nil {
		log.Println(err)
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//打开图片文件并获取参数
	f, err := os.Open(data.ImageIn.Name())
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	defer f.Close()
	img, _, err := image.Decode(data.ImageIn) //忽略了文件格式
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//获取参数
	// fmt.Println(formatName)
	// fmt.Println(img.Bounds())
	// fmt.Println(img.ColorModel())
	//保存图片
	outFile, err := os.Create("gopher2.png")
	defer outFile.Close()
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, img)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//获取文件路径
	imageAddr, err := filepath.Abs(filepath.Dir(outFile.Name()))
	if err != nil {
		log.Println(err)
		utils.JsonErrorResponse(c, 200405, "获取文件路径失败")
	}
	// //重命名文件
	// image.Filename = strconv.FormatInt(time.Now().UnixMilli(), 10) /*当前时间戳*/ + "-" + strconv.FormatUint(claims.Id, 10) /*账户id*/ + "." + strings.Split(cType, "/")[1] /*扩展名*/

	// //上传文件到oss
	// image, err := service.UploadFileToOSS("alumni-circle", "development", file) //开发阶段，此处目录为development
	// if err != nil {
	// 	global.Logger.Debug(err)
	// 	model.Server(c)
	// 	return

	// // 创建uuid
	// u1, _ := uuid.NewV4()
	// fmt.Printf("UUIDv4: %s\n", u1)

	// // 解析
	// u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	// if err != nil {
	//     fmt.Printf("Something gone wrong: %s", err)
	//     return
	// }

	//client := &Client{id: uuid.NewV4().String(), socket: conn, send: make(chan []byte)}

	//获取地址
	// addr := flag.String("fpath", "test.txt", "file path to read from")
	// flag.Parse()
	// contents, err := os.ReadFile(*fptr)//读取路径
	// if err != nil {
	//     fmt.Println("File reading error", err)
	//     return
	// }
	// fmt.Println("Contents of file:", string(contents))

	//创建新文件
	// file, err := os.Create("myfile.txt")
	// if err != nil {
	// 	log.Println(err)
	// 	utils.JsonInternalServerErrorResponse(c)
	// 	return
	// }

	//将图片存入新文件

	//获取图片路径，并存入数据库

	//os.rename(source_file, target_file)time.go

	//存入数据库
	err = articalServices.UpdateImage(data.UserID, imageAddr)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}
