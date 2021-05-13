package advancedGo

type Interfacer interface {
	Do()
}

func Work(worker Interfacer) {
	worker.Do()
}
