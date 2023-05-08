package factory

type IMinion interface {
	SetAttackPower(attackPower int)
	GetAttackPower() int
	SetAttackRange(attackRange int)
	GetAttackRange() int
	SetSpeed(speed int)
	GetSpeed() int
	GetDetails()
}
