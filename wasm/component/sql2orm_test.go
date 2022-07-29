package component

import (
	"fmt"
	"testing"
)

func TestSql2Gorm(t *testing.T) {
	sql := "CREATE TABLE `company` (\n  `id` int NOT NULL AUTO_INCREMENT,\n  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',\n  `expire_time` int DEFAULT NULL,\n  `symbol` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '简写',\n  PRIMARY KEY (`id`) USING BTREE,\n  UNIQUE KEY `name` (`name`)\n) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='公司表';"
	res, err := Sql2Gorm(sql)
	fmt.Println(res, err)
}
