package constant

/**
 * @Author nico
 * @Date 2024-10-17
 * @File: enum.go
 * @Description:
 */

const (
	STATUS_ENABLE = iota
	STATUS_DISABLE
)

var StatusValueMapKey = map[int]string{
	STATUS_ENABLE:  "Enable",
	STATUS_DISABLE: "Disable",
}

const (
	ROLE_GENERAL = iota + 1
	ROLE_ADMIN
)

const (
	Process_WAITING = iota
	PROCESS_RUNNING
	PROCESS_FINISHED
	PROCESS_ABORTED
)
