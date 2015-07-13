// musicserver
package music

import (
	"fmt"
	"net/http"
)

/**
*音乐查询接口1
 */
func MusicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "音乐查询接口1")
}
