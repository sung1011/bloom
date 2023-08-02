package svc

const Key_Meta = "tk:meta"

type Meta interface {
	// Get(key string) interface{}
	// Load(key string, val interface{}) error
	Get() *Data
}

type Data struct {
	Env string
	App struct {
		Log struct {
			Level     string
			Format    string
			Output    []string
			ErrOutput []string
		}
		Storage struct {
			Mongodb []struct {
				Name string
				URI  string
				DB   string
			}
			Redis []struct {
				Name      string
				URI       string
				DB        int
				IsCluster bool
			}
			SharedRedis []struct {
				Name string
				URLs []string
				DBs  []int
			}
		}
	}
}
