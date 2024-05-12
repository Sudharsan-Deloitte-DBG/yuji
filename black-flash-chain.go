package yuji

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Sudharsan-Deloitte-DBG/sukuna"
)

func ChainBlackFlash() string {
	i := rand.Intn(9)
	fmt.Println("Black flash chained " + strconv.Itoa(i) + "times")
	return sukuna.DefeatSukuna(i)
}
