create table users(
                      id bigint not null auto_increment comment 'id,用于点赞接口',
                      uid bigint not null comment '用户id',
                      username varchar(50) not null comment '用户名',
                      password varchar(50) not null comment '密码',
                      email varchar(100) not null comment '邮箱',
                      create_time timestamp null default current_timestamp,
                      update_time timestamp null default current_timestamp on update current_timestamp,
                      gender smallint NOT NULL DEFAULT 0 COMMENT '性别，未设定时为0，男性为1，女性为2',
                      introduction varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '简介',
                      primary key (id),
                      unique key `idx_uid` (uid) using btree ,
                      unique key `idx_username` (username) using btree
)engine =InnoDB character set =utf8mb4 collate =utf8mb4_general_ci;

create table topics(
                       tid int not null comment '话题id',
                       topic_name varchar(128) not null comment '话题名称',
                       introduction varchar(256) not null comment '话题简介',
                       create_time timestamp null default current_timestamp,
                       update_time timestamp null default current_timestamp on update current_timestamp,
                       primary key (tid)
)collate =utf8mb4_general_ci;

insert into topics (tid, topic_name, introduction) VALUES (1,'golang','golang');
insert into topics (tid, topic_name, introduction) values (2,'leetcode','算法');
insert into topics(tid, topic_name, introduction) VALUES (3,'csgo','csgo');
insert into topics (tid, topic_name, introduction) VALUES (4,'lol','英雄联盟');

create table posts(
                      pid bigint not null comment '帖子id',
                      type smallint not null comment '帖子类别,1为问题，2为文章',
                      title varchar(128) not null comment '帖子标题',
                      content varchar(8192) not null comment '帖子内容',
                      author_id bigint not null comment '作者id',
                      topic_id int not null comment '所属话题id',
                      stars bigint NOT NULL DEFAULT 0 COMMENT '点赞数',
                      create_time timestamp null default current_timestamp,
                      update_time timestamp null default current_timestamp on update current_timestamp,
                      primary key (pid),
                    key `idx_topic_id` (topic_id) using btree ,
                      key `idx_author_id` (author_id) using btree,
                      key `idx_type_id` (type) using btree
)collate = utf8mb4_general_ci;

create table comments(
                         cid bigint not null comment '评论id',
                         author_id bigint not null comment '评论人id',
                         post_id bigint comment '评论的帖子id',
                         parent_id bigint not null default 0 comment '父级评论id,若为最高级评论则为0',
                         root_id bigint not null default 0 comment '根级评论id，若为最高级评论则为0',
                         commented_uid bigint not null default 0 comment '被回复的人的id',
                         content varchar(256) not null comment '评论内容',
                         stars bigint NOT NULL DEFAULT 0 COMMENT '点赞数',
                         create_time timestamp null default current_timestamp,
                         primary key (cid),
                         key `idx_author_id` (author_id) using btree ,
                         key `idx_post_id` (post_id) using btree ,
                         key `idx_commented_uid` (commented_uid) using btree
)engine =InnoDB character set =utf8mb4 collate =utf8mb4_general_ci;