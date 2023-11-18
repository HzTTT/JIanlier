package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/db/util"
)

func createRandomAccount(t *testing.T) Account{
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account,err := testQueries.CreateAccount(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,account)

	require.Equal(t,arg.Owner,account.Owner)
	require.Equal(t,arg.Balance,account.Balance)
	require.Equal(t,arg.Currency,account.Currency)

	require.NotZero(t,account.ID)
	require.NotZero(t,account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T){
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T){
	//注意各个测试相互独立，解耦-----》每个测试都应该有自己的账户记录
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(),account1.ID)	//根据ID获取用户信息
	require.NoError(t, err)		//确保错误为nil，否则无法执行
	require.NotEmpty(t,account2)	//要求account2不为空

	require.Equal(t,account1.ID,account2.ID)
	require.Equal(t,account1.Owner,account2.Owner)
	require.Equal(t,account1.Balance,account2.Balance)
	require.Equal(t,account1.Currency,account2.Currency)
	//测试断言----->检查 2 个时间戳是否最多相差某个增量持续时间,此处间隔不超过time.Second（1 秒）
	require.WithinDuration(t, account1.CreatedAt,account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T){
	account1 := createRandomAccount(t)

	//声明变量，这是一个UpdateAccountParams 对象
	arg := UpdateAccountParams{
		ID : account1.ID,		///此处用：
		Balance  : util.RandomMoney(),
	}
	//(CURD)调用 testQueries.UpdateAccount(),传入后台上下文和更新参数
	account2,err := testQueries.UpdateAccount(context.Background(),arg)
	//要求不返回错误和更新后的account2对象不为空
	require.NoError(t, err)
	require.NotEmpty(t,account2)

	require.Equal(t,account1.ID,account2.ID)
	require.Equal(t,account1.Owner,account2.Owner)
	require.Equal(t,arg.Balance,account2.Balance)
	require.Equal(t,account1.Currency,account2.Currency)
	require.WithinDuration(t, account1.CreatedAt,account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T){
	//先创建在删除
	account1 := createRandomAccount(t)
	err :=testQueries.DeleteAccount(context.Background(),account1.ID)
	require.NoError(t,err)

	account2,err := testQueries.GetAccount(context.Background(),account1.ID)
	require.Error(t,err)
	//确保有错误和account2查不到
	require.EqualError(t,err,sql.ErrNoRows.Error())
	require.Empty(t,account2)
}

func TestListAccount(t *testing.T){
	for i := 1;i < 10;i++{
		createRandomAccount(t)
	}
	//限制是 5偏移量为5，这意味着跳过前 5 条记录，并返回后 5 条记录
	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts,err :=testQueries.ListAccounts(context.Background(),arg)
	require.NoError(t,err)
	require.Len(t,accounts,5)

	for _,account := range accounts{
		require.NotEmpty(t,account)
	}
}