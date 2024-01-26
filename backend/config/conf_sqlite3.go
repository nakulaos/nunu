/**
 ******************************************************************************
 * @file           : conf_sqlite3.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/26
 ******************************************************************************
 */

package config

import "fmt"

type Sqlite3Conf struct {
	Path   string
	Enable bool
}

func (s *Sqlite3Conf) Dsn(driverName string) string {
	//_foreign_keys=1: 启用外键约束。外键约束用于维护表与表之间的关系，确保引用完整性。
	//_journal_mode=WAL: 使用 Write-Ahead Logging（WAL）日志模式。WAL 日志模式提供更好的并发性能和事务处理效率。
	//_synchronous=NORMAL: 设置同步模式为 NORMAL。同步模式影响 SQLite 如何将数据写入磁盘。NORMAL 模式表示数据库引擎在提交事务后会等待数据写入磁盘。
	//_busy_timeout=8000: 设置忙时超时为 8000 毫秒。当数据库引擎因为其他连接正在执行而被标记为忙时，超时设置表示等待其他连接释放锁的时间
	pragmas := "_foreign_keys=1&_journal_mode=WAL&_synchronous=NORMAL&_busy_timeout=8000"
	if driverName == "sqlite" {
		pragmas = "_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(8000)&_pragma=journal_size_limit(100000000)"
	}
	return fmt.Sprintf("file:%s?%s", s.Path, pragmas)
}
