-- 创建用户

CREATE USER IF NOT EXISTS virzz@localhost identified by 'virzz9999';

-- 创建数据库

CREATE DATABASE
    IF NOT EXISTS `virzz_platform` CHARACTER SET = 'utf8mb4' COLLATE = 'utf8mb4_general_ci';

-- 授权

GRANT ALL ON `virzz_platform`.* to virzz@localhost;

-- 刷新权限

FLUSH PRIVILEGES;