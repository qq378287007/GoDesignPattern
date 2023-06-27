package main

import "fmt"

//定义一次性密码接口
type IOtp interface {
	GenRandomOTP(int) string
	SaveOTPCache(string)
	GetMessage(string) string
	SendNotification(string) error
	Publish()
}

//定义一次性密码类
type Otp struct {
	//IOtp IOtp
	IOtp
}

//生成验证码并发送
func (o *Otp) GenAndSendOTP(otpLength int) error {
	//生成随机验证码
	otp := o.IOtp.GenRandomOTP(otpLength)
	o.IOtp.SaveOTPCache(otp)
	message := o.IOtp.GetMessage(otp)
	err := o.IOtp.SendNotification(message)
	if err != nil {
		return err
	}
	o.Publish() //o.IOtp.Publish()
	return nil
}

//短信类
type Sms struct {
	Otp
}

func (s *Sms) GenRandomOTP(len int) string {
	randomOTP := "1688"
	fmt.Printf("SMS: 生成随机验证码：%s\n", randomOTP)
	return randomOTP
}

func (s *Sms) SaveOTPCache(otp string) {
	fmt.Printf("SMS: 保存验证码：%s 到缓存\n", otp)
}

func (s *Sms) GetMessage(otp string) string {
	return "登录的短信验证码是：" + otp
}

func (s *Sms) SendNotification(message string) error {
	fmt.Printf("SMS: 发送消息：%s\n", message)
	return nil
}

func (s *Sms) Publish() {
	fmt.Printf("SMS: 发布完成\n")
}

//邮箱类
type Email struct {
	Otp
}

func (s *Email) GenRandomOTP(len int) string {
	randomOTP := "3699"
	fmt.Printf("EMAIL: 生成随机验证码：%s\n", randomOTP)
	return randomOTP
}

func (s *Email) SaveOTPCache(otp string) {
	fmt.Printf("EMAIL: 保存验证码：%s 到缓存\n", otp)
}

func (s *Email) GetMessage(otp string) string {
	return "登录的短信验证码是：" + otp
}

func (s *Email) SendNotification(message string) error {
	fmt.Printf("EMAIL: 发送消息：%s\n", message)
	return nil
}

func (s *Email) Publish() {
	fmt.Printf("EMAIL:发布完成\n")
}

func main() {
	//创建短信对象
	smsOTP := &Sms{}
	o := Otp{IOtp: smsOTP}
	//生成短信验证码并发送
	o.GenAndSendOTP(4)

	fmt.Println()

	//创建邮件对象
	EmailOTP := &Email{}
	o = Otp{IOtp: EmailOTP}
	//生成邮件验证码并发送
	o.GenAndSendOTP(4)
}
