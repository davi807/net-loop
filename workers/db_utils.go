package workers

import (
	"errors"
	"strconv"

	"github.com/prologic/bitcask"
)

func addCount(meta map[string]string, addition int) {
	countStr := meta["count"]
	count, _ := strconv.Atoi(countStr)
	count += addition
	countStr = strconv.Itoa(count)
	meta["count"] = countStr
}

func removeFild(db *bitcask.Bitcask, key string) error {
	if key == "_meta_" {
		println(errors.New("'__meta__' is not allowed to delete"))
	}

	err := db.Delete([]byte(key))

	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
