package dao

import (
	"fmt"
	"server/common/mtime"
)

type LogsModel struct {
}

func NewLogsModel() {
}

func (logsModel *LogsModel) createTable() error {
	index := mtime.IntToString(mtime.GetTime(), mtime.TimeTemplate8)
	//language=sql
	_, err := dao.SqlxExec(fmt.Sprintf("create table if not exists `m_logs_%s`(`type` varchar(10) not null comment '类型', `module` varchar(255) not null collate utf8mb4_bin comment '模块', `args` varchar(255) not null collate utf8mb4_bin comment '参数', `create_time` bigint not null primary key comment '创建时间', `uid` bigint not null comment '用户id')engine=InnoDB default charset=utf8mb4 collate=utf8mb4_bin comment '日志表';", index))
	if err != nil {
		return err
	}
	return nil
}

func (logsModel *LogsModel) InsertList(data []MLogsTable) error {
	//language=sql
	_, err := dao.SqlxExec("")
	if err != nil {
		return err
	}
	return nil
}
