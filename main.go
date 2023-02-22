package main

import (
	"encoding/binary"
	"encoding/csv"
	"math/big"
	"net"
	"net/http"
	"os"
	"sort"

	"github.com/gin-gonic/gin"
)

func Ip2Int(ip net.IP) *big.Int {
	i := big.NewInt(0)
	i.SetBytes(ip)
	return i
}

func ReadAndGet(fn string) []IpItem {
	var ip_items []IpItem = []IpItem{}
	f, _ := os.Open(fn)
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err != nil {
			break
		}
		start := new(big.Int)
		start, _ = start.SetString(record[0], 10)
		end := new(big.Int)
		end, _ = end.SetString(record[1], 10)
		ip_items = append(ip_items, IpItem{start, end, record[2]})
	}
	f.Close()
	return ip_items
}

func main() {
	var ip_items []IpItem = []IpItem{}
	ip_items = append(ip_items, ReadAndGet("ipv4_db.csv")...)
	ip_items = append(ip_items, ReadAndGet("ipv6_db.csv")...)

	sort.Slice(ip_items, func(i, j int) bool {
		return ip_items[i].start.Cmp(ip_items[j].start) == -1
	})

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/getIpInfo", func(c *gin.Context) {
		addr := net.ParseIP(c.Query("addr"))
		if addr != nil {
			ip_num := big.NewInt(0)

			if addr.To4() != nil {
				ip_num = new(big.Int).SetUint64(uint64(binary.BigEndian.Uint32(addr.To4())))
			} else {
				ip_num.SetBytes(addr)
			}
			idx, _ := Binary(ip_items, ip_num, 0, len(ip_items))
			if idx != -1 && ip_num.Cmp(big.NewInt(0)) != 0 {
				c.JSON(http.StatusOK, gin.H{
					"ok":      true,
					"country": ip_items[idx].country,
				})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"ok": false,
		})
	})
	router.NoMethod(catchAll)
	router.NoRoute(catchAll)
	router.Run(":8080")
}

func catchAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": false,
	})
}
