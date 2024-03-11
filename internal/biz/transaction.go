package biz

import "context"

// Transaction 新增事务接口方法
type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}
