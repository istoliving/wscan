/**
* @Author: shaochuyu
* @Date: 5/7/2022 11:30
 */
package collections

import (
	"container/list"
	"sync"
)

type Queue struct {
	l *list.List
	m sync.RWMutex
}
