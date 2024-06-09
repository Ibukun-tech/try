package try

type Service interface {
	Add(Log) (string, error)
	List() (Logs, error)
}
