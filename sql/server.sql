
create table if not exists `m_user`
(
    `uid` bigint not null comment 'uid',
    `username` varchar(20) not null collate utf8mb4_bin comment '账户名',
    `password` varchar(32) not null collate utf8mb4_bin comment '密码',
    `name` varchar(50) collate utf8mb4_bin comment '名称',
    `email` varchar(255) collate utf8mb4_bin comment '邮件',
    `create_time` bigint not null comment '创建时间',
    `last_login_time` bigint comment '上次登录时间',
    `state` tinyint not null comment '状态',
    primary key (`uid`),
    unique index (`username`)
)engine=InnoDB default charset=utf8mb4 collate=utf8mb4_bin comment '用户表';

create table if not exists  `m_role`
(
    `rid` bigint not null comment 'rid',
    `rname` varchar(20) not null collate utf8mb4_bin comment '角色名',
    primary key (`rid`)
)engine=InnoDB default charset=utf8mb4 collate=utf8mb4_bin comment '角色表';

create table if not exists `m_user_role_relation`
(
    `uid` bigint not null comment 'uid',
    `rid` bigint not null comment 'rid',
    index(`uid`),
    index(`rid`),
    unique index (uid, rid)
)engine=InnoDB default charset=utf8mb4 collate=utf8mb4_bin comment '用户角色关系表';

create table if not exists `m_user_permission_relation`
(
    `uid` bigint not null comment 'uid',
    `pid` bigint not null comment '权限id',
    index (`uid`),
    index (`pid`),
    unique index (`uid`, `pid`)
)engine=InnoDB default charset=utf8mb4 collate=utf8mb4_bin comment '用户权限关系表';

create table if not exists `m_role_permission_relation`
(
    `rid` bigint not null comment 'rid',
    `pid` bigint not null comment '权限id',
    index (`rid`),
    index (`pid`),
    unique index (`rid`, `pid`)
)engine=InnoDB default charset=utf8mb4 collate=utf8mb4_bin comment '角色权限关系表';
