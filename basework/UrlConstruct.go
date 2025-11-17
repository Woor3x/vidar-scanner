package basework

func UrlConstruct(urlStr string, filename string) []string {
	list, err := LoadFile2List(filename)
	CheckErr(err)
	finalpath := make([]string, 0, len(list))

	//base, err := url.Parse(urlStr)
	//if err != nil{
	//	log.Printf("error:%v",err)
	//	return nil
	//}

	for _, path := range list {
		finalUrl := urlStr + path
		finalpath = append(finalpath, finalUrl)
	}

	return finalpath
}
