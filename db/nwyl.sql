/*
 Navicat Premium Data Transfer

 Source Server         : local_mysql
 Source Server Type    : MySQL
 Source Server Version : 80401
 Source Host           : 127.0.0.1:3306
 Source Schema         : nwyl

 Target Server Type    : MySQL
 Target Server Version : 80401
 File Encoding         : 65001

 Date: 12/08/2024 18:23:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_allow
-- ----------------------------
DROP TABLE IF EXISTS `app_allow`;
CREATE TABLE `app_allow`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `app_id` bigint NULL DEFAULT NULL COMMENT '应用id',
  `service_id` int NULL DEFAULT NULL COMMENT '服务id',
  `service_allow` datetime(3) NULL DEFAULT NULL COMMENT '服务许可时间',
  `service_exp` datetime(3) NULL DEFAULT NULL COMMENT '服务过期时间',
  `use_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '使用方式(1次数 2流量 3期限)',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示过期2禁用3正常4暂停',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '业务应用许可的服务项目和期限' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for app_auth
-- ----------------------------
DROP TABLE IF EXISTS `app_auth`;
CREATE TABLE `app_auth`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `app_id` bigint NULL DEFAULT NULL COMMENT '应用id',
  `access_key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '授权码',
  `key_exp` datetime(3) NULL DEFAULT NULL COMMENT '过期时间',
  `start_time` datetime(3) NULL DEFAULT NULL COMMENT '启动时间',
  `register_time` datetime(3) NULL DEFAULT NULL COMMENT '登记时间',
  `contract` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '合同地址',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `app_id_index`(`app_id` ASC) USING BTREE COMMENT 'id不允许重复'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '业务应用的授权码记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for app_direct
-- ----------------------------
DROP TABLE IF EXISTS `app_direct`;
CREATE TABLE `app_direct`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `app_id` bigint NULL DEFAULT NULL COMMENT '业务引用id',
  `direct_id` bigint NULL DEFAULT NULL COMMENT '数据目录id',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1正常访问 2禁止',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for app_info
-- ----------------------------
DROP TABLE IF EXISTS `app_info`;
CREATE TABLE `app_info`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `app_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '业务应用名称',
  `app_user` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '应用联系人',
  `app_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `app_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '联系电话',
  `app_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '业务应用类型(1文旅2电商3资产4交易)',
  `app_public` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '应用公钥',
  `app_public_md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '应用公钥md5',
  `app_private` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '应用私钥',
  `app_private_md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '应用私钥md5',
  `algo` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '算法',
  `store_key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '应用私钥',
  `start_time` datetime(3) NULL DEFAULT NULL COMMENT '开始时间',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示正常 2禁用 3暂停',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `app_name_index`(`app_name` ASC) USING BTREE COMMENT '应用名称唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '业务应用的基本信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for app_log
-- ----------------------------
DROP TABLE IF EXISTS `app_log`;
CREATE TABLE `app_log`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `app_id` bigint NULL DEFAULT NULL COMMENT 'app用户id',
  `service_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '服务类别',
  `req_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求路径',
  `method` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求方法',
  `remote_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '访问ip',
  `req_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '请求内容',
  `res_msg` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '响应消息',
  `res_size` int NULL DEFAULT NULL COMMENT '响应字节数',
  `res_err` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '响应错误',
  `cost` bigint NULL DEFAULT NULL COMMENT '耗时(ms)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 420 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '服务的访问日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for app_token
-- ----------------------------
DROP TABLE IF EXISTS `app_token`;
CREATE TABLE `app_token`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `app_id` bigint NULL DEFAULT NULL COMMENT '业务应用id',
  `token` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT 'token',
  `token_exp` datetime(3) NULL DEFAULT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for data_direct
-- ----------------------------
DROP TABLE IF EXISTS `data_direct`;
CREATE TABLE `data_direct`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `direct_id` bigint NULL DEFAULT NULL COMMENT '目录id',
  `parent_id` bigint NULL DEFAULT NULL COMMENT '父目录id',
  `direct_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目录名称',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标签',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '数据目录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for data_direct1
-- ----------------------------
DROP TABLE IF EXISTS `data_direct1`;
CREATE TABLE `data_direct1`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `direct_id1` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目录标识',
  `direct_id1_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目录名称',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标签',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for data_direct2
-- ----------------------------
DROP TABLE IF EXISTS `data_direct2`;
CREATE TABLE `data_direct2`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `direct_id2` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目录id(00000000进行表示)',
  `direct_id2_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目录名称',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标签',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for data_direct3
-- ----------------------------
DROP TABLE IF EXISTS `data_direct3`;
CREATE TABLE `data_direct3`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `direct_id3` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目录id(00000000进行表示)',
  `direct_id3_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目录名称',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标签',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for data_enc
-- ----------------------------
DROP TABLE IF EXISTS `data_enc`;
CREATE TABLE `data_enc`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据Id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `file_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文件名称',
  `file_size` bigint NULL DEFAULT NULL COMMENT '文件大小',
  `file_md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文件hash',
  `file_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '对称密钥',
  `file_algo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '加密算法',
  `file_encrypted_hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '加密之后的hash',
  `file_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '可以进行文件下载的URL',
  `node_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for data_info
-- ----------------------------
DROP TABLE IF EXISTS `data_info`;
CREATE TABLE `data_info`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `use_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `data_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '数据散列',
  `data_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '数据名称',
  `data_diectory_id` int NULL DEFAULT NULL COMMENT '数据目录',
  `data_type` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示文档2表示图片3表示视频4表示关系数据',
  `is_open` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示开放2表示关闭',
  `data_size` bigint NULL DEFAULT NULL COMMENT '数据大小',
  `data_pos` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '数据位置',
  `node_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点id',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '介绍',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '数据是否共享 是否需要审核 数据价值' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for data_safe
-- ----------------------------
DROP TABLE IF EXISTS `data_safe`;
CREATE TABLE `data_safe`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `data_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '数据id',
  `safe_level` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '安全登记(1表示确权2授权3密码)',
  `safe_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '授权信息确权信息密码信息',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '记录数据安全相关信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for data_share
-- ----------------------------
DROP TABLE IF EXISTS `data_share`;
CREATE TABLE `data_share`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `buyer_node` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '买家节点',
  `seller_node` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '卖家节点',
  `buyer_id` bigint NULL DEFAULT NULL COMMENT '买家id',
  `seller_id` bigint NULL DEFAULT NULL COMMENT '卖家id',
  `trade_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '交易编号',
  `trade_price` double NULL DEFAULT NULL COMMENT '交易价格',
  `seller_approve` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示卖家审核 0表示卖家为审核',
  `data_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '数据id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '请求共享和审核来自哪个业务应用' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for data_sharelog
-- ----------------------------
DROP TABLE IF EXISTS `data_sharelog`;
CREATE TABLE `data_sharelog`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `buyer_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '买家名字',
  `seller_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '卖家名字',
  `data_ids` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '数据集合',
  `price` double NULL DEFAULT NULL COMMENT '价格',
  `data_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '数据名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '成功共享的信息记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for node_hard
-- ----------------------------
DROP TABLE IF EXISTS `node_hard`;
CREATE TABLE `node_hard`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `hard_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '硬件id',
  `hard_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '硬件名称',
  `hard_model` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '硬件型号',
  `hard_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '硬件类别(机架,存储,GPU,CPU)',
  `hard_brand` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '硬件品牌',
  `hard_install` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '安装信息',
  `hard_params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '硬件参数',
  `hard_explain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '硬件介绍',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '超级节点的硬件信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for node_info
-- ----------------------------
DROP TABLE IF EXISTS `node_info`;
CREATE TABLE `node_info`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `node_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点编号',
  `node_unit` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点单位',
  `node_deploy` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点部署地址',
  `node_operation` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点运营商',
  `login_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点登录码',
  `node_public` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '节点公钥',
  `node_public_md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '公钥md5',
  `node_private` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '节点私钥',
  `node_private_md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '私钥md5',
  `node_algo` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点公私要生成算法',
  `node_store_key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点sm2私钥保存算法',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '超级节点的登记信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for node_log
-- ----------------------------
DROP TABLE IF EXISTS `node_log`;
CREATE TABLE `node_log`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `node_id_sender` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '发送的节点id',
  `node_id_reader` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '接受的节点id',
  `msg_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '消息内容',
  `err_msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '发送消息是否成功',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '记录节点的通信记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for node_soft
-- ----------------------------
DROP TABLE IF EXISTS `node_soft`;
CREATE TABLE `node_soft`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `node_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点id',
  `soft_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件id',
  `soft_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件url',
  `sort_size` bigint NULL DEFAULT NULL COMMENT '软件大小(单位字节)',
  `soft_md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件md5',
  `soft_version` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件版本',
  `soft_language` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件语言',
  `soft_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件类型',
  `soft_explain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件说明',
  `device_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '所在设备id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '超级节点的软件信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for node_user
-- ----------------------------
DROP TABLE IF EXISTS `node_user`;
CREATE TABLE `node_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `node_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点id',
  `node_contact` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点联系人',
  `node_phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点联系人电话',
  `contact_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示可登录人员 2纯管理人员 3投资人',
  `user_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点人员账号',
  `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点人员密码',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示正常 2表示暂停 3表示禁用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '超级节点的联系人信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for service_cfg
-- ----------------------------
DROP TABLE IF EXISTS `service_cfg`;
CREATE TABLE `service_cfg`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `service_id` bigint NULL DEFAULT NULL COMMENT '服务id',
  `service_config` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '服务代理的配置',
  `start_time` datetime(3) NULL DEFAULT NULL COMMENT '启动时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'json描述的配置信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for service_info
-- ----------------------------
DROP TABLE IF EXISTS `service_info`;
CREATE TABLE `service_info`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `service_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '服务名称',
  `service_exec_dir` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '程序运行目录',
  `service_version` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '服务版本',
  `service_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '服务类别(1水印 2智能 3数权)',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示正常 2表示异常(不能用)',
  `visit_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '访问地址',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '备注',
  `node_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '服务代理的基本信息,包含程序' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority`;
CREATE TABLE `sys_authority`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `authority_id` bigint NULL DEFAULT NULL COMMENT '角色id',
  `authority_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '角色名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统角色' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_authority_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menu`;
CREATE TABLE `sys_authority_menu`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '删除时间',
  `authority_id` int NULL DEFAULT NULL COMMENT '角色ID',
  `menu_id` int NULL DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色菜单绑定表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `sys_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '系统编号',
  `key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'key',
  `value` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'value',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '主要配置通信' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_download
-- ----------------------------
DROP TABLE IF EXISTS `sys_download`;
CREATE TABLE `sys_download`  (
  `id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '记录软件的下载信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `menu_level` int NULL DEFAULT NULL COMMENT '菜单级别',
  `parent_id` bigint NULL DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '路由path',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '路由名称',
  `hidden` tinyint(1) NULL DEFAULT NULL COMMENT '是否隐藏',
  `compoment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '组件',
  `sort` bigint NULL DEFAULT NULL COMMENT '排序',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '图标',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标题',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色菜单' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_msg
-- ----------------------------
DROP TABLE IF EXISTS `sys_msg`;
CREATE TABLE `sys_msg`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `msg_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示平台通知 2表示电话通知',
  `msg_body` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '通知内容',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '通知各种信息，有短信和平台通知两种' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_operlog
-- ----------------------------
DROP TABLE IF EXISTS `sys_operlog`;
CREATE TABLE `sys_operlog`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求ip',
  `method` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求方法',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求路径',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求代理',
  `err_msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '错误信息',
  `body` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '响应body',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '管理端人员的操作记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_policy
-- ----------------------------
DROP TABLE IF EXISTS `sys_policy`;
CREATE TABLE `sys_policy`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `node_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点id',
  `soft_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件id',
  `start_upgrade` datetime(3) NULL DEFAULT NULL COMMENT '开始升级时间',
  `end_upgrade` datetime(3) NULL DEFAULT NULL COMMENT '升级完成时间',
  `upgrade_err` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '升级错误',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '描述各种升级策略' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_soft
-- ----------------------------
DROP TABLE IF EXISTS `sys_soft`;
CREATE TABLE `sys_soft`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `soft_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件id',
  `soft_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件名称',
  `soft_version` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件版本',
  `soft_package_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件包url',
  `soft_package_md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '软件包md5',
  `upload_user` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '上传用户',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '软件信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户名称',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户密码',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户 昵称',
  `side_mode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户侧边主题',
  `header_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户头像',
  `base_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '基础颜色',
  `authority_id` bigint NULL DEFAULT NULL COMMENT '用户角色id',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '电话号码',
  `email` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `enable` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1正常 2禁用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统用户' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user_account
-- ----------------------------
DROP TABLE IF EXISTS `user_account`;
CREATE TABLE `user_account`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `account` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT 'json数组用户绑定',
  `approve` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '记录验证账号信息',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_id_index`(`user_id` ASC) USING BTREE COMMENT '用户id唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_bank
-- ----------------------------
DROP TABLE IF EXISTS `user_bank`;
CREATE TABLE `user_bank`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `user_bank_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '开户行名称',
  `user_bank` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '银行卡号',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1审核通过 0未审核',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_id_index`(`user_id` ASC) USING BTREE COMMENT 'id 不重复'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户用于支付信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user_ca
-- ----------------------------
DROP TABLE IF EXISTS `user_ca`;
CREATE TABLE `user_ca`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint NULL DEFAULT NULL COMMENT '创建时间',
  `public` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '公钥d',
  `public_md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '公钥md5',
  `private` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '私钥',
  `private_md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '私钥md5',
  `algo` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '算法',
  `store_key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '存储密码',
  `time_stamp` bigint NULL DEFAULT NULL COMMENT '时间戳',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用用户公钥证书 日期 算法 存储密钥' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user_fee
-- ----------------------------
DROP TABLE IF EXISTS `user_fee`;
CREATE TABLE `user_fee`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `order_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '订单id',
  `pay_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '支付方式',
  `pay_amount` double NULL DEFAULT NULL COMMENT '支付金额',
  `currency` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '币种',
  `opposite_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '对方账户',
  `trade_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '交易类型(数据买卖 扩容等)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户缴费充值有关记录信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户名称',
  `nick_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户昵称',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户密码',
  `user_level` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户级别',
  `user_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户手机号',
  `user_source_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示超级节点 2表示平台注册 3业务应用',
  `user_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '存储用户来源的详细信息',
  `user_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示个人 2表示企业 3表示政府部门',
  `did` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户did',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户的详细信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user_log
-- ----------------------------
DROP TABLE IF EXISTS `user_log`;
CREATE TABLE `user_log`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `login_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '登录ip',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '登录代理',
  `response` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '失败返回错误消息进行记录',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户登录信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user_other
-- ----------------------------
DROP TABLE IF EXISTS `user_other`;
CREATE TABLE `user_other`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `id_card` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '身份证号',
  `addr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '地址',
  `business_license` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '营业执照',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '不同类别用户补充信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user_pdc
-- ----------------------------
DROP TABLE IF EXISTS `user_pdc`;
CREATE TABLE `user_pdc`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `space_total` int NULL DEFAULT NULL COMMENT '默认单位MB',
  `space_use` double NULL DEFAULT NULL COMMENT '使用空间单位字节',
  `space_available` double NULL DEFAULT NULL COMMENT '剩余空间单位字节',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1正常 2禁用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户数据空间变化记录表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user_prove
-- ----------------------------
DROP TABLE IF EXISTS `user_prove`;
CREATE TABLE `user_prove`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `account_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号类型(qq等da)',
  `account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号密码',
  `app_id` bigint NULL DEFAULT NULL COMMENT '应用id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户认证信息' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
