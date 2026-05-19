CREATE TABLE `admin_account` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `account` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '管理员账号',
  `password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '密码（加密存储）',
  `real_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `nick_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '昵称',
  `email` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '手机号',
  `avatar` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '头像地址',
  `sex` TINYINT NOT NULL DEFAULT 0 COMMENT '性别：0未知 1男 2女',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1启用 0禁用',
  `is_admin` TINYINT NOT NULL DEFAULT 0 COMMENT '是否超级管理员：1是 0否',
  `last_login_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后登录时间-默认注册时间',
  `last_login_ip` char(64) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `create_by` BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
  `update_by` BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` DATETIME DEFAULT NULL COMMENT '删除时间（软删除）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_account` (`account`),
  KEY `idx_status` (`status`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_delete_time` (`delete_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员账号表';

CREATE TABLE `admin_role` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '角色名称',
  `description` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '角色描述',
  `sort` INT NOT NULL DEFAULT 0 COMMENT '排序号',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1启用 0禁用',
  `parent_id` BIGINT NOT NULL DEFAULT 0 COMMENT '父角色ID',
  `is_default` TINYINT NOT NULL DEFAULT 0 COMMENT '是否默认角色：1是 0否',
  `create_by` BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
  `update_by` BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` DATETIME DEFAULT NULL COMMENT '删除时间（软删除）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_code` (`role_code`),
  KEY `idx_status` (`status`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_delete_time` (`delete_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

CREATE TABLE `admin_menu` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `menu_type` TINYINT NOT NULL DEFAULT 1 COMMENT '类型：1目录 2菜单 3按钮',
  `permission_code` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '权限标识（如 system:user:list）',
  `parent_id` BIGINT NOT NULL DEFAULT 0 COMMENT '父菜单ID',
  `icon` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '图标',
  `route_path` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '路由路径',
  `route_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '路由名称',
  `component` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '组件路径',
  `query_params` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '路由参数',
  `is_frame` TINYINT NOT NULL DEFAULT 1 COMMENT '是否内嵌：1是 0否',
  `is_cache` TINYINT NOT NULL DEFAULT 1 COMMENT '是否缓存：1缓存 0不缓存',
  `visible` TINYINT NOT NULL DEFAULT 1 COMMENT '是否显示：1显示 0隐藏',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1启用 0禁用',
  `sort` INT NOT NULL DEFAULT 0 COMMENT '排序号',
  `remark` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '备注',
  `create_by` BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
  `update_by` BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_permission_code` (`permission_code`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_menu_type` (`menu_type`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单权限表';

CREATE TABLE `admin_account_role` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `account_id` BIGINT NOT NULL DEFAULT 0 COMMENT '管理员ID',
  `role_id` BIGINT NOT NULL DEFAULT 0 COMMENT '角色ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_account_role` (`account_id`, `role_id`),
  KEY `idx_role_id` (`role_id`),
  FOREIGN KEY (`account_id`) REFERENCES `admin_account` (`id`) ON DELETE CASCADE,
  FOREIGN KEY (`role_id`) REFERENCES `admin_role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员角色关联表';

CREATE TABLE `admin_role_menu` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_id` BIGINT NOT NULL DEFAULT 0 COMMENT '角色ID',
  `menu_id` BIGINT NOT NULL DEFAULT 0 COMMENT '菜单ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_menu` (`role_id`, `menu_id`),
  KEY `idx_menu_id` (`menu_id`),
  FOREIGN KEY (`role_id`) REFERENCES `admin_role` (`id`) ON DELETE CASCADE,
  FOREIGN KEY (`menu_id`) REFERENCES `admin_menu` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色菜单关联表';


CREATE TABLE `admin_dept` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `dept_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '部门名称',
  `parent_id` BIGINT NOT NULL DEFAULT 0 COMMENT '父部门ID',
  `leader` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '负责人',
  `phone` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '联系电话',
  `email` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '邮箱',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1启用 0禁用',
  `sort` INT NOT NULL DEFAULT 0 COMMENT '排序号',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` DATETIME DEFAULT NULL COMMENT '删除时间（软删除）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_dept_name` (`dept_name`),
  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门表';

CREATE TABLE `admin_account_dept` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `account_id` BIGINT NOT NULL DEFAULT 0 COMMENT '管理员ID',
  `dept_id` BIGINT NOT NULL DEFAULT 0 COMMENT '部门ID',
  `is_main` TINYINT NOT NULL DEFAULT 0 COMMENT '是否主部门：1是 0否',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_account_dept` (`account_id`, `dept_id`),
  KEY `idx_dept_id` (`dept_id`),
  FOREIGN KEY (`account_id`) REFERENCES `admin_account` (`id`) ON DELETE CASCADE,
  FOREIGN KEY (`dept_id`) REFERENCES `admin_dept` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员部门关联表';

CREATE TABLE `admin_operation_log` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `account_id` BIGINT NOT NULL DEFAULT 0 COMMENT '操作人ID',
  `account_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '操作人账号',
  `module` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '操作模块',
  `operation_type` TINYINT NOT NULL DEFAULT 0 COMMENT '操作类型：1新增 2修改 3删除 4查询 5导出 6导入 9999其他',
  `title` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '操作标题',
  `url` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '请求URL',
  `method` VARCHAR(16) NOT NULL DEFAULT '' COMMENT '请求方式',
  `request_params` TEXT COMMENT '请求参数',
  `response_data` TEXT COMMENT '返回数据',
  `ip` VARCHAR(64) NOT NULL DEFAULT '' COMMENT 'IP地址',
  `location` VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'IP归属地',
  `duration` INT NOT NULL DEFAULT 0 COMMENT '耗时(毫秒)',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1成功 0失败',
  `error_msg` VARCHAR(512) NOT NULL DEFAULT '' COMMENT '错误消息',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `idx_account_id` (`account_id`),
  KEY `idx_module` (`module`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_operation_type` (`operation_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员操作日志表';






