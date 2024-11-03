package sync

var t *TickerCap

// InitSync 初始化sync并发操作
func InitSync() error {
	t = NewTickerCap()
	t.Start()
	return nil
}

func Close() {
	t.Stop()
}
