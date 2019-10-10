package repos

type Repo interface {
	FindFromDB(map[string]interface{}) (interface{}, error)
	FindFromCache(map[string]interface{}) (interface{}, error)
	SaveCache(interface{})
}

func Find(repo Repo, m map[string]interface{}) (ret interface{}, err error) {
	ret, err = FindCache(repo, m)
	if ret != nil {
		return
	}
	ret, err = FindDB(repo, m)
	if err != nil {
		return
	}
	repo.SaveCache(ret)
	return
}
func FindCache(repo Repo, m map[string]interface{}) (ret interface{}, err error) {
	ret, err = repo.FindFromCache(m)
	return

}
func FindDB(repo Repo, m map[string]interface{}) (ret interface{}, err error) {
	ret, err = repo.FindFromDB(m)
	return
}
