package bridge

// Abstraction
type RemoteControl interface {
	togglePower()
	volumnDown()
	volumnUp()
	previousChannel()
	nextChannel()
}

// Refined Abstraction
type XiaoMiControl struct {
	device Device
}

func (x *XiaoMiControl) BindDevice(device Device) {
	x.device = device
}

func (x *XiaoMiControl) togglePower() {
	if x.device.isEnable() {
		x.device.disable()
	} else {
		x.device.enable()
	}
}

func (x *XiaoMiControl) volumnDown() {
	x.device.setVolumn(x.device.getVolumn() - 10)
}

func (x *XiaoMiControl) volumnUp() {
	x.device.setVolumn(x.device.getVolumn() + 10)
}

func (x *XiaoMiControl) previousChannel() {
	x.device.setChannel(x.device.getChannel() - 1)
}

func (x *XiaoMiControl) nextChannel() {
	x.device.setChannel(x.device.getChannel() + 1)
}

var _ RemoteControl = &XiaoMiControl{}

// Implementation
type Device interface {
	isEnable() bool
	enable()
	disable()
	getVolumn() int
	setVolumn(int)
	getChannel() int
	setChannel(int)
}

// Concrete Implementations
type TV struct {
	state   bool
	volumn  int
	channel int
}

func (t *TV) isEnable() bool {
	return t.state
}

func (t *TV) enable() {
	t.state = true
}

func (t *TV) disable() {
	t.state = false
}

func (t *TV) getVolumn() int {
	return t.volumn
}

func (t *TV) setVolumn(volumn int) {
	t.volumn = volumn
}

func (t *TV) getChannel() int {
	return t.channel
}

func (t *TV) setChannel(channel int) {
	t.channel = channel
}

var _ Device = &TV{}

type Radio struct {
	state   bool
	volumn  int
	channel int
}

func (t *Radio) isEnable() bool {
	return t.state
}

func (t *Radio) enable() {
	t.state = true
}

func (t *Radio) disable() {
	t.state = false
}

func (t *Radio) getVolumn() int {
	return t.volumn
}

func (t *Radio) setVolumn(volumn int) {
	t.volumn = volumn
}

func (t *Radio) getChannel() int {
	return t.channel
}

func (t *Radio) setChannel(channel int) {
	t.channel = channel
}

var _ Device = &Radio{}

// Client
func CheckMyRemoteController() {
	tv := &TV{}
	remote := &XiaoMiControl{}
	remote.BindDevice(tv)

	radio := &Radio{}
	remote.BindDevice(radio)
}
