package workload_tracer

import (
	"github.com/Sirupsen/logrus"
	. "github.com/google/syzkaller/tools/syz-strace/domain"
	"github.com/google/syzkaller/tools/syz-strace/ssh"
	"github.com/google/syzkaller/tools/syz-strace/config"
)

type DefaultTracer struct {
	executor Executor
	workloads []WorkloadConfig
}

func NewDefaultTracer(config config.CorpusGenConfig) (tracer Tracer) {
	var executor Executor
	switch config.Executor {
	case "ssh":
		executor = syz_ssh.NewClient(config.SSHConfig, "127.0.0.1")
	default:
		panic("Only ssh executor supported\n")
	}
	workloads := readWorkload(config.ConfigPath)
	tracer = &DefaultTracer{executor: executor, workloads: workloads}
	return
}


func (dt *DefaultTracer) GenerateCorpus() (err error) {
	//ctx := context.Background()
	//client, err := storage.NewClient(ctx)
	if err != nil {
		logrus.Fatalf("Unable to generate client config: %s", err.Error())
	}
	for _, wc := range dt.workloads {
		if err = dt.executor.RunStrace(wc); err != nil {
			return err
		}
	}
	return
}
