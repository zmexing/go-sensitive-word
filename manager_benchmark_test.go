package go_sensitive_word

// 压力测试
// func BenchmarkIsSensitive(b *testing.B) {
// 	filter, err := NewFilter(
// 		StoreOption{Type: StoreMemory},
// 		FilterOption{Type: FilterDfa},
// 	)
//
// 	if err != nil {
// 		log.Fatalf("敏感词服务启动失败, err:%v", err)
// 		return
// 	}
//
// 	// 加载敏感词库
// 	err = filter.LoadDictEmbed(
// 		DictCovid19,
// 		DictOther,
// 		DictReactionary,
// 		DictViolence,
// 		DictPeopleLife,
// 		DictPornography,
// 		DictAdditional,
// 		DictCorruption,
// 		DictTemporaryTencent,
// 	)
// 	if err != nil {
// 		log.Fatalf("加载词库发生了错误, err:%v", err)
// 		return
// 	}
//
// 	err = filter.Store.AddWord("测试1", "测试2")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
//
// 	sensitiveText := "小明微笑着对毒品销售说，我认为台湾国的人有点意思"
//
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_ = filter.IsSensitive(sensitiveText)
// 	}
// }

// 压力测试
// func BenchmarkReplace(b *testing.B) {
// 	filter, err := NewFilter(
// 		StoreOption{Type: StoreMemory},
// 		FilterOption{Type: FilterDfa},
// 	)
//
// 	if err != nil {
// 		log.Fatalf("敏感词服务启动失败, err:%v", err)
// 		return
// 	}
//
// 	// 加载敏感词库
// 	err = filter.LoadDictEmbed(
// 		DictCovid19,
// 		DictOther,
// 		DictReactionary,
// 		DictViolence,
// 		DictPeopleLife,
// 		DictPornography,
// 		DictAdditional,
// 		DictCorruption,
// 		DictTemporaryTencent,
// 	)
// 	if err != nil {
// 		log.Fatalf("加载词库发生了错误, err:%v", err)
// 		return
// 	}
//
// 	err = filter.Store.AddWord("测试1", "测试2")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
//
// 	sensitiveText := "小明微笑着对毒品销售说，我认为台湾国的人有点意思"
//
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_ = filter.Replace(sensitiveText, '*')
// 	}
// }
