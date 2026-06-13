package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/zeromicro/go-zero/zrpc"
	// 👈 替换成你项目中生成的 pb 包路径
	"repaycal/apps/calculator/rpc/protoc/calculator"
)

func TestCalculator_Calculate(t *testing.T) {
	// 1. 建立与本地 go-zero rpc 服务的连接
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Endpoints: []string{"localhost:8080"}, // 如果在K8S测试，可转发端口到本地
	})
	if err != nil {
		t.Fatalf("连接失败: %v", err)
	}

	// 2. 绑定生成的 pb 客户端
	conn := client.Conn()
	rpcClient := calculator.NewCalculateClient(conn)

	// 3. 发起请求测试
	resp, err := rpcClient.CalculatePMT(context.Background(), &calculator.PMTRequest{
		LoanAmount:       100000,
		InterestRate:     45, // 比如代表 4.5%
		NumberOfPayments: 12,
	})
	if err != nil {
		t.Fatalf("接口报错: %v", err)
	}

	// 4. 打印返回结果
	fmt.Printf("🎯 测试成功！每月还款额为: %v\n", resp.MonthlyRepayment)
}
