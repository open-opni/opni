package testbench

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"log/slog"

	"emperror.dev/errors"
	"github.com/golang/snappy"
	"github.com/open-panoptes/opni/internal/bench"
	"github.com/open-panoptes/opni/pkg/logger"
	"github.com/open-panoptes/opni/pkg/test"
	"github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/prompb"
	"github.com/prometheus/prometheus/storage/remote"
)

// Adapted from cortex benchtool

type BenchRunner struct {
	client    remote.WriteClient
	logger    *slog.Logger
	batchChan chan bench.BatchReq
}

func NewBenchRunner(env *test.Environment, agent string, desc bench.WorkloadDesc) (*BenchRunner, error) {
	runningAgent := env.GetAgent(agent)
	if runningAgent.Agent == nil {
		return nil, fmt.Errorf("agent %s not found", agent)
	}
	u, err := url.Parse(fmt.Sprintf("http://%s/api/agent/push", runningAgent.Agent.ListenAddress()))
	if err != nil {
		return nil, err
	}
	writeClient, err := remote.NewWriteClient(agent, &remote.ClientConfig{
		URL:              &config.URL{URL: u},
		HTTPClientConfig: config.HTTPClientConfig{},
		RetryOnRateLimit: true,
		Timeout:          model.Duration(10 * time.Second),
	})
	if err != nil {
		return nil, errors.Wrap(err, "error creating remote write client")
	}

	workload := bench.NewWriteWorkload(desc, nil)
	batchChan := make(chan bench.BatchReq, 100)
	go func() {
		if err := workload.GenerateWriteBatch(env.Context(), agent, 100, batchChan); err != nil {
			env.Logger.Error(fmt.Sprintf("error generating write batch: %v", err))
		}
	}()

	return &BenchRunner{
		client:    writeClient,
		logger:    env.Logger.WithGroup("bench").With("agent", agent),
		batchChan: batchChan,
	}, nil
}

func (b *BenchRunner) StartWorker(ctx context.Context) {
	go func() {
		b.logger.Info("worker started")
		for {
			select {
			case batchReq := <-b.batchChan:
				if err := b.sendBatch(ctx, batchReq.Batch); err != nil {
					b.logger.With(
						logger.Err(err),
					).Error("failed to send batch")
				}
				batchReq.PutBack <- batchReq.Batch
				batchReq.Wg.Done()
			case <-ctx.Done():
				b.logger.With(
					logger.Err(ctx.Err()),
				).Warn("worker stopped")
				return
			}
		}
	}()
}

func (b *BenchRunner) sendBatch(ctx context.Context, batch []prompb.TimeSeries) error {
	req := prompb.WriteRequest{
		Timeseries: batch,
	}

	data, err := req.Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal remote-write request")
	}

	compressed := snappy.Encode(nil, data)

	return b.client.Store(ctx, compressed)
}
