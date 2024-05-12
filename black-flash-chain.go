package yuji

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Sudharsan-Deloitte-DBG/sukuna"
)

func ChainBlackFlash() string {
	i := rand.Intn(10)
	fmt.Println("Black flash chained " + strconv.Itoa(i) + "times")
	return sukuna.DefeatSukuna(i)
}
