package redis

func  Set(key, val string) error {
	client := Pool.Get()
	defer client.Close()
	_ , err := client.Do("SET" , key , val)
	return err
}

func SetEx(key, val string, exp int64) error {
	client := Pool.Get()
	defer client.Close()
	_ , err := client.Do("SETEX" , key , exp , val)
	return err
}

func Get(key string) (map[string]string , error) {
	client := Pool.Get()
	defer client.Close()
	rsp := make(map[string]string)
	res , err := client.Do("GET" , key)
	rsp[key] = string(res.([]byte))
	return rsp , err
}

func LPush(key, val string) error {
	client := Pool.Get()
	defer client.Close()
	_, err := client.Do("lpush", key, val)
	return err
}

func LPop(key string) (interface{}, error) {
	client := Pool.Get()
	defer client.Close()
	res, err := client.Do("lpop", key)
	return res, err
}

func RPush(key, val string) error {
	client := Pool.Get()
	defer client.Close()
	_, err := client.Do("rpush", key, val)
	return err
}

func RPop(key string) (interface{}, error) {
	client := Pool.Get()
	defer client.Close()
	res, err := client.Do("rpop", key)
	return res, err
}

func Del(key string) (error) {
	client := Pool.Get()
	defer client.Close()
	_ , err := client.Do("DEL" , key)
	return err
}