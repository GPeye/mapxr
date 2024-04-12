package main

import "github.com/micmonay/keybd_event"

var tapmap map[uint8]int

func initTapMap() {
	tapmap = make(map[uint8]int)

	tapmap[1] = keybd_event.VK_A
	tapmap[2] = keybd_event.VK_E
	tapmap[3] = keybd_event.VK_N
	tapmap[4] = keybd_event.VK_I
	tapmap[5] = keybd_event.VK_D
	tapmap[6] = keybd_event.VK_T
	//tapmap[7]= keybd_event. shift mode
	tapmap[8] = keybd_event.VK_O
	tapmap[9] = keybd_event.VK_K
	tapmap[10] = keybd_event.VK_M
	tapmap[11] = keybd_event.VK_F
	tapmap[12] = keybd_event.VK_L
	tapmap[13] = keybd_event.VK_G
	tapmap[14] = keybd_event.VK_BACKSPACE
	tapmap[15] = keybd_event.VK_R
	tapmap[16] = keybd_event.VK_U
	tapmap[17] = keybd_event.VK_Y
	tapmap[18] = keybd_event.VK_B
	tapmap[19] = keybd_event.VK_P
	tapmap[20] = keybd_event.VK_Z
	tapmap[21] = keybd_event.VK_W
	tapmap[22] = keybd_event.VK_Q
	tapmap[23] = keybd_event.VK_J
	tapmap[24] = keybd_event.VK_S
	tapmap[25] = keybd_event.VK_ENTER
	tapmap[26] = keybd_event.VK_X
	tapmap[27] = keybd_event.VK_V
	//tapmap[28]= keybd_event. switch mode
	tapmap[29] = keybd_event.VK_C
	tapmap[30] = keybd_event.VK_H
	tapmap[31] = keybd_event.VK_SPACE
}
