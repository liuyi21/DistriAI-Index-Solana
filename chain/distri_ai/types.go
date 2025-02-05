// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package distri_ai

import ag_binary "github.com/gagliardetto/binary"

type MachineStatus ag_binary.BorshEnum

const (
	MachineStatusIdle MachineStatus = iota
	MachineStatusForRent
	MachineStatusRenting
)

func (value MachineStatus) String() string {
	switch value {
	case MachineStatusIdle:
		return "Idle"
	case MachineStatusForRent:
		return "ForRent"
	case MachineStatusRenting:
		return "Renting"
	default:
		return ""
	}
}

type OrderStatus ag_binary.BorshEnum

const (
	OrderStatusTraining OrderStatus = iota
	OrderStatusCompleted
	OrderStatusFailed
	OrderStatusRefunded
)

func (value OrderStatus) String() string {
	switch value {
	case OrderStatusTraining:
		return "Training"
	case OrderStatusCompleted:
		return "Completed"
	case OrderStatusFailed:
		return "Failed"
	case OrderStatusRefunded:
		return "Refunded"
	default:
		return ""
	}
}
