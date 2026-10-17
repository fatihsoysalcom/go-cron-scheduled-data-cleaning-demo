# Go Cron Scheduled Data Cleaning Demo

This example demonstrates how to schedule a recurring task in Go using the `robfig/cron/v3` library. It simulates a "payment data cleaning" job that runs every few seconds, logging its start, end, and simulated processing duration. This illustrates the core concept of automating periodic tasks, which is central to the article's discussion on cron schedulers for data management and their impact on latency and cost.

## Language

`go`

## How to Run

1. Save the code as `main.go`.
2. Run `go mod init example && go mod tidy` to initialize the module and download dependencies.
3. Execute `go run main.go`.
4. Observe the scheduled tasks running in the console. Press `Ctrl+C` to stop.

## Original Article

This example accompanies the Turkish article: [Go Planlı Ödeme Verisi Temizliği: Dört Cron Scheduler Gecikme ve Maliyet Testlerinin Karşılaştırılması](https://fatihsoysal.com/blog/go-planli-odeme-verisi-temizligi-dort-cron-scheduler-gecikme-ve-maliyet-testlerinin-karsilastirilmasi/).

## License

MIT — see [LICENSE](LICENSE).
