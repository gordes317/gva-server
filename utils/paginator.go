package utils

//@author: wuchengjiang
//@function: Paginator
//@description: 分页
//@param: pg int ,pgSize int ,totalResult[]interface{}
//@return: pagiResult interface{}, err error
func Paginator(pg, pgSize int, totalResult []interface{}) (pagiResult interface{}, err error) {

	if len(totalResult) >= pgSize*(pg-1) {
		if pgSize*pg <= len(totalResult) { //中间页
			pagiResult = totalResult[pgSize*(pg-1) : pgSize*pg]
		} else if pg-1 == 0 { //第一页
			pagiResult = totalResult[:]
		} else if pgSize*pg > len(totalResult) { //最后一页
			pagiResult = totalResult[pgSize*(pg-1):]
		}
		return pagiResult, err
	} else {

		return totalResult, err

	}

}
