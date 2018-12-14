package camera

import "fmt"

type Camera struct{}

func (camera *Camera) TakePhoto() {
	fmt.Println("照相了，茄子")
}

type Photo struct{}

func (photo *Photo) Call() {
	fmt.Println("大家一起来欣赏刚才照的照片了")
}

type CameraPhoto struct {
	Camera
	Photo
}

func CameraTest() {
	cameraPhoto := new(CameraPhoto)
	cameraPhoto.TakePhoto()
	cameraPhoto.Call()
}
