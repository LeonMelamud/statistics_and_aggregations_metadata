package cmd

import "sort"

func mostFrequentExt(m map[string]int) ExtInfo {
	var maping []kv
	for k, v := range m {
		maping = append(maping, kv{k, v})
	}

	sort.Slice(maping, func(i, j int) bool {
		return maping[i].Value > maping[j].Value
	})
	extInfo := ExtInfo{}
	extInfo.Extension = maping[0].Key
	extInfo.NumOccurrences = int64(maping[0].Value)
	return extInfo

}

type kv struct {
	Key   string
	Value int
}
