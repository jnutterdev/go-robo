package main

import (
        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/firmata"
)

func main() {
        firmataAdaptor := firmata.NewAdaptor("/dev/tty.usbmodem1411")

        button := gpio.NewButtonDriver(firmataAdaptor, "5")
        led := gpio.NewLedDriver(firmataAdaptor, "13")

        work := func() {
                button.On(gpio.ButtonPush, func(data interface{}) {
                        led.On()
                })
                button.On(gpio.ButtonRelease, func(data interface{}) {
                        led.Off()
                })
        }

        robot := gobot.NewRobot("buttonBot",
                []gobot.Connection{firmataAdaptor},
                []gobot.Device{button, led},
                work,
        )

        robot.Start()
}
