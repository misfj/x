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

 Date: 02/08/2024 18:37:01
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
-- Records of app_allow
-- ----------------------------
INSERT INTO `app_allow` VALUES (1, '2024-07-30 14:53:19.783', '2024-07-30 14:53:19.783', '2024-07-30 14:53:19.783', 1, 1, '2024-07-30 15:03:19.783', '2024-07-30 15:13:19.783', '1', '3');
INSERT INTO `app_allow` VALUES (2, '2024-07-30 14:54:00.250', '2024-07-30 14:54:00.250', '2024-07-30 14:54:00.250', 1, 1, '2024-07-30 15:04:00.250', '2024-07-30 15:14:00.250', '1', '3');
INSERT INTO `app_allow` VALUES (3, '2024-07-30 14:54:01.340', '2024-07-30 14:54:01.340', '2024-07-30 14:54:01.340', 1, 1, '2024-07-30 15:04:01.340', '2024-07-30 15:14:01.340', '1', '3');
INSERT INTO `app_allow` VALUES (4, '2024-07-30 14:58:21.188', '2024-07-30 14:58:21.188', '2024-07-30 14:58:21.188', 1, 1, '2024-07-30 15:08:21.188', '2024-07-30 15:18:21.188', '1', '3');
INSERT INTO `app_allow` VALUES (5, '2024-07-30 14:58:30.555', '2024-07-30 14:58:30.555', '2024-07-30 14:58:30.555', 1, 1, '2024-07-30 15:08:30.555', '2024-07-30 15:18:30.555', '1', '3');
INSERT INTO `app_allow` VALUES (6, '2024-07-30 14:58:31.543', '2024-07-30 14:58:31.543', '2024-07-30 14:58:31.543', 1, 1, '2024-07-30 15:08:31.543', '2024-07-30 15:18:31.543', '1', '3');
INSERT INTO `app_allow` VALUES (7, '2024-07-30 15:01:47.494', '2024-07-30 15:01:47.494', NULL, 1, 1, '2024-07-30 15:11:47.494', '2024-07-30 15:21:47.494', '1', '3');
INSERT INTO `app_allow` VALUES (8, '2024-07-30 15:05:29.189', '2024-07-30 15:05:29.189', NULL, 1, 1, '2024-07-30 15:15:29.189', '2024-07-30 15:25:29.189', '1', '3');
INSERT INTO `app_allow` VALUES (9, '2024-07-30 15:05:30.321', '2024-07-30 15:05:30.321', NULL, 1, 1, '2024-07-30 15:15:30.321', '2024-07-30 15:25:30.321', '1', '3');
INSERT INTO `app_allow` VALUES (10, '2024-07-30 15:05:31.184', '2024-07-30 15:05:31.184', NULL, 1, 1, '2024-07-30 15:15:31.184', '2024-07-30 15:25:31.184', '1', '3');
INSERT INTO `app_allow` VALUES (11, '2024-07-30 15:05:32.296', '2024-07-30 15:05:32.296', NULL, 1, 1, '2024-07-30 15:15:32.296', '2024-07-30 15:25:32.296', '1', '3');
INSERT INTO `app_allow` VALUES (12, '2024-07-30 15:05:33.177', '2024-07-30 15:05:33.177', NULL, 1, 1, '2024-07-30 15:15:33.177', '2024-07-30 15:25:33.177', '1', '3');
INSERT INTO `app_allow` VALUES (13, '2024-07-30 15:05:34.054', '2024-07-30 15:05:34.054', NULL, 1, 1, '2024-07-30 15:15:34.054', '2024-07-30 15:25:34.054', '1', '3');
INSERT INTO `app_allow` VALUES (14, '2024-07-31 17:55:39.500', '2024-07-31 17:55:39.500', NULL, 1, 1, '2024-07-31 18:05:39.500', '2024-07-31 18:15:39.500', '1', '3');
INSERT INTO `app_allow` VALUES (15, '2024-07-31 17:56:00.865', '2024-07-31 17:56:00.865', NULL, 1, 1, '2024-07-31 18:06:00.865', '2024-07-31 18:16:00.865', '1', '3');
INSERT INTO `app_allow` VALUES (16, '2024-07-31 17:56:19.973', '2024-07-31 17:56:19.973', NULL, 1, 1, '2024-07-31 18:06:19.973', '2024-07-31 18:16:19.973', '1', '3');

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '业务应用的授权码记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_auth
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '业务应用的基本信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_info
-- ----------------------------
INSERT INTO `app_info` VALUES (9, '2024-08-01 16:44:37.301', '2024-08-01 16:44:37.301', '2024-08-02 17:46:42.000', '文旅平台用户db', '小万', '$2a$04$1ubkwHn9ZjTDpnGkVnnuqe3P77e0Z8yX.RD3Pnzg7AIJGb1Q56C16', '12334556', '文旅平台', 'MIIBCgKCAQEAv5qs7oi8ox1oOWzXMmTb267tmvwU6Xy+Nr/N6XWfyw5oNJnAgA8o+UkpsCVw8Ul2pWzyJGgiZ79jRS/MHY8DpmGz5bv+yvjHk1bLH9TLW58P+syP0jpCx5qOD0XN/1YXR+hkp4vwbpMArsSN+BjuwxJ1QvWGP/7Zhda0UkZwCRrqH3+gumnVA3fqG+uNZ85ibY3uEQ2HwHaK/+/22PH2Z93IK8sKIROxcNbRPlvwnK1iPuRwRo9+pNUMxHZAmUR0WyOoeRJ2y4YnoAtLV+KkLTf8rd6wK7o1ShEgR3L5AQZsdHSYBYpGwuk9EWX4lzsIZuIAPJ18kyDfVV9UFTW/4wIDAQAB', '4bea3e9a1d29cc78245081f4b899252d', 'MIIEogIBAAKCAQEAv5qs7oi8ox1oOWzXMmTb267tmvwU6Xy+Nr/N6XWfyw5oNJnAgA8o+UkpsCVw8Ul2pWzyJGgiZ79jRS/MHY8DpmGz5bv+yvjHk1bLH9TLW58P+syP0jpCx5qOD0XN/1YXR+hkp4vwbpMArsSN+BjuwxJ1QvWGP/7Zhda0UkZwCRrqH3+gumnVA3fqG+uNZ85ibY3uEQ2HwHaK/+/22PH2Z93IK8sKIROxcNbRPlvwnK1iPuRwRo9+pNUMxHZAmUR0WyOoeRJ2y4YnoAtLV+KkLTf8rd6wK7o1ShEgR3L5AQZsdHSYBYpGwuk9EWX4lzsIZuIAPJ18kyDfVV9UFTW/4wIDAQABAoIBAGdRNYzOaMGgGtJUUVoYLsGv+unGolQg8sS8+oC373lkBGbQvY+QHS0G+ZbFKbjCP1sGey+Mg3jkJ9hn5K3Pr00phAMOM18pdRyLAa7hZOK7y7cWVTssNVRCj/XGU0UrHBSnSJdcC0Lku8qBrHgqoZzIuDhqAoFBn/PeiY+P8dvGG2S9+D/g7vEeLhscAZdfbQvHgZFz0YVstyRk/1t5cnnPzs3/VaSpv9CDu2KIf7VKhE/zgb21pWv1f3yh15j70dpmBRKZfKMYim3suDkYs+A8UUtGhw9YRcFqGS7hYfrQvH53iVwtxar8ILXVcsuWDV7HS/YmkUrPUrKIsRC0oyECgYEA0eHw4kGqdMPBjgm+3cyBcViCWyoPaAN5gEqejCsbOwfLlmlkwiso/ojgbuRvzPwLy9eKnWRSL0klYu60uXzpfYriNoqkaNIXIwozNujxLUUSjSI/p/JTfMy1XgX1LSnPiX2x4JGl3uAF2NkygbIAw6YBxZWproSWnMYpfKfr61MCgYEA6bSPtv1V3AHPeTC1LLPfW1v2Dmha+rRb5kUBERSRsEaBJ8Vpfisbhyvi+H2x2emt5e0hHnejsowqF9QqyH9ZuCSdRn26o0HJYm/Rjx0oWneDY10dgTsfeVhOSaQwPeBOze7I3GKiwOHPc1wqvOq2Z3YTVL4DiTLdaV4BY8WJ1zECgYBVyrqxcdAgfKhQzDEF7wN2doyKgU3c4Za310ip2TD/VzdkG17Bc+0E5qR/D8eXjeuwfkG+BnUFuSucHiEntSSoWa4QR2nfIwmsHA5FQsDbFWH46OgGWarA19jHPz4yIOYhuIyOVFruGLqrIlVf6IghcWlF/+BlYvNB0ErDFsz72wKBgAFomDEnJ7xSunhlhcIGN4Nqc3o0wM+gBKnHorECqbohQqgFtZQQGHNhHVYYE9lXYjfvjPKNEAyEccCBA9z/f3/voI8LNPPE8rzIJcAyEcsxyzXvGr1rqM3nbVleRvgogPEWzlkdsxNsiP6OAKfw1jabbJRnLMhH6TO4YFPEngcRAoGAaN0vQxYWb85P5SbAkoqLLN1dmD4adhABk/Rv4oeUOzKPAan4lXRuNMhYih9p1kFT+JQ+FxKlAKeAxqGt5KIx+EG4NWtAK4UHY1XdwexdOQtJMaMA0kEbslnlLbFHYE9PgYayct9cJ94CJxDrjG3K6CyLTH/Dsc5DYfE8CzTU1IU=', '7a61748679660a66a206323c8acf50b6', 'rsa', '', '2024-08-01 16:44:37.300', '1');
INSERT INTO `app_info` VALUES (10, '2024-08-01 16:55:11.586', '2024-08-01 16:55:11.586', NULL, '你好小艾', '小万', '$2a$04$nIykSMzakkZ9u1P8TDD0pu2iQ2/dIuagH/CY2NXQMo2spEgn1v3a2', '12334556', '文旅平台', 'MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB', '2c752aa85774005612c593c0f9156c3a', 'MIIEpAIBAAKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQABAoIBAQCDJtrMDJJYd8fgcaInoYR2UtbesWf9XqwjiYBfcC7QJkz1pmRYUI5iAKsre0JydhiM3zrxpM/xzi5bpf3HXQNMVPd3jBCvIb6K2pVvSE7EYOCFqMhWQjZA7gXBzq3L+llFMd3mcMZZhiE2xcLJLI+2W9qhl5dVqeqnqvSVn6szM1iTHGy5KkFOsKWjwIKu1OZQ/zInFjSbX46PNYgyzOfT/r4eUzmwrYZi60dqiAM9/VGVpb/TeilNakH11LBzDs1MdePKKvOoWS7XamjQzfnl73xwj4ae9rKj8Bm4iI47NT3P6JUNmif/ZznLaqNYNITE/6Nilus7j0Lwe3mWN2zVAoGBANXGFpHEiCMVjZBGkiFeu/M9SqNWALk4UNq9FvrQzE45r8FNc6sez77aJ4QA02idRXHEZgHzjnwuEt2kOlc/1UwU9lAGhDAN4+9bhyGVv7OOMptf/z9zJdNkqeM1yVe/k1WiCMF+QYJ0U3OsFo37k8Zz3hiZNYy954K9At8aYcRrAoGBAN2SSltJmpdp6vMmIdfz5LY8k3XudKFTWTgZiDZX7W5EqZMP+Eej77JoqsShdifL9W2u9K4sBbv+ac0IcBZdNOVCuoL7fGArdNq/oE5QxW9ob3sX6+UDPvho5xTrY9RyiYWB00svJYgRP0XeS0P3iWuOGf7KN5mPoEEXJEii3k6PAoGATsVoSZxGoc7uVAx9CkjTLuUV3KvzJk0HFiL5FcDcl7KP/cWhapu251eRmgWH4ROapDo/cFaj2URM9VjvmnjOCvPqc0CDcwWbSPbMC2cXtX0fTruH3YR7mrKMG0oVf/z6uGI1ob75Oq3MlmICCZ17LXRCkYPbBJpxdW5aYYViIzMCgYAXLd6pfiG2BVaW6T5RISqkBeJHTBLXbai/vawKQ1iLWJOE1v7aP8Qrndd/ioSZ1TRvrQXb0q8m446d7GwNauLFuR1Z5oU0aldnKvTSPJDq/rnq3q3G33gO9Kp3/cgLyY+pb9Ny4hPDEErWPFMTYSn3Hn6JsQ3v+p6BMLKTBrbK8QKBgQC1s50TuFab3KelmbgdvgsfazX5mZfnDSe4SD+5cyiJwkWNWNC96ZbSsd7zVAGkp4QQdc7Ux+Et+y2bBkV+P11XZUyp0n9BFMLfvty/fgAuEQthe039lUiZk+u0Ev6RKEkY9/auuDVRcPcNQJm2oJBYFUmp6Mj4qEKRgeHw454M+w==', 'a32fa7df4a0b418fdedec4c6455bf87e', 'rsa', '', '2024-08-01 16:55:11.586', '1');
INSERT INTO `app_info` VALUES (11, '2024-08-02 15:44:54.464', '2024-08-02 15:44:54.464', NULL, '你好小艾1', '小万', '$2a$04$EZAZNnGtonA/RrCiVBa/VO3TTT7FrQVN532yDODUcGM3bJqsGGJJC', '12334556', '文旅平台', 'MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB', '6cfb71885d97715d9817542bcc16784f', 'MIIEowIBAAKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQABAoIBAFXXeGwuKQO/0WYwnvezSdsxesd74a/YT3lSz5g8wr7S7CHMZ7SAWz0MsNmC1yTZ3Tacj4V/fr9O4hwx6GKzw5kNCgti2Do7LZjJKqCWBfGecX6F22oJz6EDLgUoJvLKEtChbGQi44crLUuCe9oKBreE2gL3wbJGEIOW7xI/UM00G4Sg3YkxqJ3IvOzMXgpMfKxRqFaIrKLR0xFIh3xBOZAu7n3E2I+nc4+vQu64SZt8x08f8a9gDZTa1B265jmmAaI5i9jmnwkiHlHAtOGW91h/nYJsfyhDwMZbfDEdYmW27Hb34KEDj66MZE6LxbtNBOY7VYmRvEtJcVl8fG2QviECgYEAzp68l72grbb+8vvLJys2eSR2eRFsdnI9vo42F2iF2UuONoLHKY03G2z7aVqx8Ed+KCOQp8CX+Dv926jtb8IulKxNYLvrqiPVskA7JSJF9oYD+z2V5N6kTt4yM5ikaGIBa6AYt3qQ9CxuuKbKL+D5cwXOrA++nQqdV9yRCdaVTf0CgYEA0Azmuq+6CUD7ltC/pQMwW8wKx1Cc9QPzAmsx0P7tUbOfOofA2BuBQJWbGhlHlaNEC0IX7wor+r4iIehiq7yybQG0bpIK7u1ryiGED8BcS/vMVsXtTyxn9iw/NGiwhsOhrnwTEePr5yPv3G+GRWoZf/2ALtNVnq+8upj0w+GgkokCgYBEk/sUgiPqhGpqS5hKD9Wsr6WC2X2ecpx7usfJTS/fx2wUO6PO7sfzmmB3v1p/3brJp8RDaXEHOyIu3gqwpUrAnc+w4658chhjaV44HQALqr93jCFMLXlrk+Qlq+wlmSHPjvulKlEiOf0l3HYPHiStQKgwA6nIhdVjbBnXohk8IQKBgQDO9vaUYnmT+RQOnMz472os95cFXhk12F7nWhGtkdwtuZ090Ywyr3Ht+KXZ2aoPnuHyvSYoAptrjbdQuedmZomszNRcZmX/9ymSY7MHJFnssKk+7Ri8+TTWfH/HuBlbhsORWxZxZqXnQyW6ySruZLTHUTwx2MRE3xfFHKHpNsvQ2QKBgA7lSHR19uZ1yw0v4YyG4KrHspfCdehQZ1/HmcKp2laxZWn/+Dl6qLayYRRQ/3uPWEFg1Ot3EuQg7BHKyiDFVaD7k/ww6gkq0TD58gVGxFHwHUzXP9nJffqdTjYe34EaAPIehsCXu5FZLnIxpWfobsdI014IaqYXOkEgAwqHgoRd', '96016fb3a78bdb4419389c7e8eaec3cc', 'rsa', '', '2024-08-02 15:44:54.464', '1');
INSERT INTO `app_info` VALUES (12, '2024-08-02 17:45:45.537', '2024-08-02 17:45:45.537', NULL, '你好小艾33331', '小万', '$2a$04$p.d8SgABivVrw4DUMIykNOVpaAQeBJMkh7AbC7goFG5lVsA.298gC', '12334556', '文旅平台', 'MIIBCgKCAQEAo+ly0HVlPTPxGW/Ox+19lKosf001alJeWhtRL/BZ/H7UBQ4pkrtsyhXs+6dV4QAZa4CaiZ6jMR7X88qh1KbtO3w1pRczUBROduiAMqVkE2kqQJW3ZoUiepO9PFxDMPmw4gSxDJw0BK/Zni4Zg0nXmDVIgZwLXjOeAlnQxU/lFGzg8R89NzjlBXzz4cM45EgRwsohnLubjnuaNt7p7EIDAMEu5rBnlz8k3zNYZKEXQ3q9UK3iGQdOfyJYQuQ3bcEhWW5GvcWZ9IntiETdijdJ6JNbbGgz7GulzsnbKjb0iJ43Tp7E0MtXGme+oB3RCIKdeRofP2qZZqqgNKHqR7q6tQIDAQAB', 'ff4f0402620a45b7805556147eff4390', 'MIIEpAIBAAKCAQEAo+ly0HVlPTPxGW/Ox+19lKosf001alJeWhtRL/BZ/H7UBQ4pkrtsyhXs+6dV4QAZa4CaiZ6jMR7X88qh1KbtO3w1pRczUBROduiAMqVkE2kqQJW3ZoUiepO9PFxDMPmw4gSxDJw0BK/Zni4Zg0nXmDVIgZwLXjOeAlnQxU/lFGzg8R89NzjlBXzz4cM45EgRwsohnLubjnuaNt7p7EIDAMEu5rBnlz8k3zNYZKEXQ3q9UK3iGQdOfyJYQuQ3bcEhWW5GvcWZ9IntiETdijdJ6JNbbGgz7GulzsnbKjb0iJ43Tp7E0MtXGme+oB3RCIKdeRofP2qZZqqgNKHqR7q6tQIDAQABAoIBAGt/fjoy09uoxhV9m2rJWZP3a0lt9HsvXAvSYJOFGS8caqaKHD/MRdPfP1Xn782piAJWDIP3E2XyqYSMv0e81lct5ezI43FAcBlR8EnG6HRGg1xqFCbbZ7pmTDRa+gITr6r7krYHd6IfRiSUbU0KT6fWQiAmm9oDFjzKdL4C4OVcWfoyvlcqCxehFks1U8swiSbwv5Lg+kH0Bne1ybL4pRKv5Mt3FI/1tAB46ZAOBn6sNva7MQHxQW4rMGzdwg0deIgcXBqFb6f3gfL7pY7El+dTmLUsl5cZBV6AhiP9xdc3MzPb+0y7QTZkrDx8aqC/3hK0tkNtwDjehDVd4JWtUiECgYEAxyg6iuZV/RWxMiot3L6XsqBrUo0ihpLxZJT/sgk2/MYWadKo6FPgORWNYUwV33L9gIwF8S+0KAxmZYAnENbTs6qoZ1jeC0tHQwDvLfOJ6GRAADhwHdznnY1J3nO0lCs4Aeuh5OFNoy2xBsBAF2Y9kiZMuS/i4kFzSBpBvVx82v0CgYEA0rH2xSljl6f2y8yQKowrAY+O/B2Aklv8U4WOFmo4mzOgi9f/kDtewScZr7xdB+9ZfzJ2TedN2ly3dzr96YEc78LKx3mfZddrW9z9+QyEqS/3UnRvtE+EPm+BctNfhzXyF2vLuVFXQB2vhoVoGNdpyFr5tNjuKCQlSORMr/TKOBkCgYEAiNhzoY2z+Ds3jqi86rlrsynXlNm1p5LAm8RmDgeW6QmTzRMbTMiVgaP1ia2mOevhlLqiOp/0wrACjCwKq99xiDrzjE/cAxau2LODpad6U1zPPVeKcnqgC/pRj7Ehm842pdsrWwIVPtLemSzw3SM7m3kFrxngZGdkt16TXXJ3uGECgYAL/35h891sAXCYYmnDQoAgcqBd0wBDVWGlD/HrbNkUXXhC/mXbPMZEkYlw7HUPwp2CFMmP3iAc0zLPY1iaN3QCY0FQ4qglEINYk9kSmZfkWorm8yJ5Th3lwEvK0iaaeIP0uXq8btldGVOpgWdQItQvSJ73oXLObmCBKh8D1hi94QKBgQCt5JyVdT+O07rEuALL4VnnqFMYtXQTNz8A1uBHQBnmyfVyKgzTYf+mVHXKbcLfwOJX/itXm0GoJ5D4Ok21JqbT8BAriK01slczUMiB+nzlKsgIWO/ElOACx4xJcPpfsPvgAVJjl09vh/V/e3lLxhmOkeIoI2q6TssZ6jCFM/UptA==', '1cf753e64c25f10a9604c410abff581d', 'rsa', '', '2024-08-02 17:45:45.536', '1');
INSERT INTO `app_info` VALUES (13, '2024-08-02 17:45:48.777', '2024-08-02 17:45:48.777', NULL, '你好小艾33331333', '小万', '$2a$04$wtv5Ku3UyrhNOdLY3gWMY.pTuLEZDU8lodMD/KKZQEIFfeqojnPrq', '12334556', '文旅平台', 'MIIBCgKCAQEA5JuN7ZOUaoEvNHNOW2e5YuQuouCSCdi3/vybn42SxGjv7GDZEmLgtIinyGFElJ54YdHO04HJ7seoG5LXxI1sVaFEhc3MEc29jJjD0eCDLI+0D7Qxcg1RR0LPsTjBsY5CUHl8hiWDanywqlgu2ZF2be3k4pNroyE7GT7qkGz89Q6ORDN6qB8fTAEu98i67jEfrd20feTCALgbqdtWqpKR5owN5tKzaj/IcuIsIlc4weCb0erUJaKdrY7SWQ3q4TRFi8LTxCaujtaQ2EYaFVFyZrdyt2iJeFvMsiozq0AloP59eCQUFmkfTSAfdRCxrebD366nWF1m+IZlAT2yay1akwIDAQAB', '179cac03e9e4a0860129c13d6a921c88', 'MIIEpQIBAAKCAQEA5JuN7ZOUaoEvNHNOW2e5YuQuouCSCdi3/vybn42SxGjv7GDZEmLgtIinyGFElJ54YdHO04HJ7seoG5LXxI1sVaFEhc3MEc29jJjD0eCDLI+0D7Qxcg1RR0LPsTjBsY5CUHl8hiWDanywqlgu2ZF2be3k4pNroyE7GT7qkGz89Q6ORDN6qB8fTAEu98i67jEfrd20feTCALgbqdtWqpKR5owN5tKzaj/IcuIsIlc4weCb0erUJaKdrY7SWQ3q4TRFi8LTxCaujtaQ2EYaFVFyZrdyt2iJeFvMsiozq0AloP59eCQUFmkfTSAfdRCxrebD366nWF1m+IZlAT2yay1akwIDAQABAoIBAEWCgtIlGHCV3RuOn9mtHTJTfVaq/9ycl72hY1RNQL4VxjXScM5lYDukfZew86BY2vOrTr8C7Lp83MGdkZvDAJi800/39j/HlgXlAU2UfW4UN8S6nno2UGthhjM1tbdeMQ21EsbbKy06wqDY9U3UK/Va60h1WLoeRZD0j2/go72C30pgRHcwrL0EDBo3aU9EeP2ieUjFP/kSrPRdhuUaJKrHyuT71ufi1+YJqCfjr+8iv4T7LQVvSd8jiDts1ET889w3HUqPjWIOAZEOzc+e1TD46ZSvv26eilyOnh6zkN+JvbPfEm6TjanwaW/Co+HBbB+4xF4PB3ka4Ni96fmrsBECgYEA6TbpLTdkQ2EOlgy5Wl9wq6F4eRv+/kp84FgKYYch5yku5tpZdDvVliZo4IWTen59nZxer/hSzGH0x/9L6fnymkfK4PXc7hmRONwdUXGKlfBKzpUVBhl5rooC9Fd38VuV0Lrt4nvFTRyVCeONda+cz1Znn+wvRaQ7wdAdf27ldNkCgYEA+vFrDR94thm51pe9yuvRRD5nBMr3IwTOtrNbQ4w7URhW86AgWQUaOjN6vpxXupQVBE1UuESjRFwDzdrrcE0UABDQA4lhHvpLoXSerm04e0P5KosdNTSXdvqfyz0ZNZSGZj2HgryYLb47Dl83UFZ0BHo1nsy7ekrRohrTn1q1t0sCgYEA26VX37o3uC4FKtU5DSCuYTdm28m+U4mcKz9AJ7K5/5kOD2a+sXZZJ11tkWi9CbVwYa0QkPN0KAJqJKMRwIAkOaI8NbvaJHX4DHFjO32QZL839XD7qA7+R6C8P5zR3oO/iiNEQPcJUCMHS5UBekG3kp9yYXB05UZU8aks3wR3H6ECgYEA1FTEDPWNMSAxc9cvW20Pw9u9VXvbg0EFG+hFF7GnWyXqJPhfhpNfrtFyUah+Prwf0GUSg1xofWYSPfO39WzyuF+BcwyiJhEJP5Mq1VJLUVmhkhLl4ugeJSlJXjHTi9ehpQ75A3FKnqHw3GdKqWqNfmBd0IFRPNfj63Mhjxu02GMCgYEA6MId68X3NTjlqeTqXd6/Q4tmEF7jAd0QlZh1SNFp/2Xyd+8Dx8ocLguSNO5U6K93wZMMDw6l9EbyNfjc5uOmLkD9aZEj7aFDZxtsI5/WDdmwyhd9tGDqPYPxT+39ABcKxUOp5S2oO3e9buuIkJrFVRefCH9RldCg27Dos2uPDus=', '0c07897c757c3df17642ba5ab1a25cca', 'rsa', '', '2024-08-02 17:45:48.776', '1');
INSERT INTO `app_info` VALUES (14, '2024-08-02 17:46:03.776', '2024-08-02 17:46:03.776', NULL, '东算细算', '小万', '$2a$04$DMXpOQ/I/6.3V.S5YvNWdeilqw7/Ogd7CB5NrR6EsjstcLeoqa.Qi', '12334556', '文旅平台', 'MIIBCgKCAQEA0taMMRDQ4hoof20XqTWwKRt9HDq0sDz6Y7z2zLfsmMxtbp3FYH5f2HKTdieudzNvfPSYONlpsQv1J8zhXB/oOSirMJ5pEFU7FASFH2Z2aeIAERm0XYVOpCe6v34qCwLSOE448UUZ66YuAyx6mME4bl9yWM48oxnGxXP4dJ9wldGbGYo28mRam/vH2nrtJ879KBOob0/q8FnTrBRtaTbwcAnX2Q0vsHqaLjOH3P7mm/MpF44n22x7oUm6KtosFbwDP/0ag/A1ErrNVaA6UO5RCEw9HwtkXbLq9Y7/oLPx30CS01MXi4FCUixU6L7TN2ZbWN5BNLAN4Ne0R19KMdZwBQIDAQAB', '3511a0d2eeb2dd49d7af2cd79f4831f4', 'MIIEpAIBAAKCAQEA0taMMRDQ4hoof20XqTWwKRt9HDq0sDz6Y7z2zLfsmMxtbp3FYH5f2HKTdieudzNvfPSYONlpsQv1J8zhXB/oOSirMJ5pEFU7FASFH2Z2aeIAERm0XYVOpCe6v34qCwLSOE448UUZ66YuAyx6mME4bl9yWM48oxnGxXP4dJ9wldGbGYo28mRam/vH2nrtJ879KBOob0/q8FnTrBRtaTbwcAnX2Q0vsHqaLjOH3P7mm/MpF44n22x7oUm6KtosFbwDP/0ag/A1ErrNVaA6UO5RCEw9HwtkXbLq9Y7/oLPx30CS01MXi4FCUixU6L7TN2ZbWN5BNLAN4Ne0R19KMdZwBQIDAQABAoIBAHC2uiPQ9ZqqmVPmImwXg6G2TGA3EjnVn5aKgvZWrWnSf/5O4iH/7YVtW2AjPqYDHWsT5/0cXeCEn/8zDJePhzpnf31ycoCE0ByXfgOXZ16br6V0tHP6vFwN0UxmSluwmmn3h4GUi5LrOFYCBVh8k4qOBDdtqgUX1qfgo8bVJf8xIqotlVwmMMeFu7mEU+GYvn2VOjA2/tGbyMEoaJ6r9xST7n3dc823+6acgBiYEloCISyF0yW+15kd+0NHUvKNiylZbiKCBVbFLjEJN7qfX3yu6UIe4n/naqyWT3UbJAOMJVzrMkhHg3aL4Ere9fEoCYAlZ9b93ylz8FtKcs09dwECgYEA81LCUZov1W9q5DByA0R90pAHTxbqJxGQ8p+q/TvKsdQRHfmhK8YHGquy7r+4CIkFIZ/LraJW3JcClc5d1uAQnv3LyxtKjxhxxakJRpRwqKbBLUhB596RBBJoFtovD58IKqAeaaHrrprJYmvFe2uGb+YzrkwgL2GmsqpdGWni71ECgYEA3dKHP7szWCbgp9AHlpOjs/kzfFlLfz2bh9cuTH91fyr8pjgRbapblJBEZlbD6JwWxxcka3dpSd9rYX7ACSWmQ62vfL2M89ywZi2pJ6yKF9LPUngel06Dh967Q/9YxZ4vKiSlh+vK3CBCE0LJ0FP4gm8FWnBWp+9QJi6Y7bUKEHUCgYEAicEt0CMZt503h+7BbYYmMcJm0OjpMaspQ6MPnARw0dJ8ylqcW3rmZLwWNkKGPxbLt+iO2EOTJ5m61fMaf18lKc4HxXoy8OZm6v+zB6lcSq71qNxCq7H+qa5+DYoCoLZpWjCCZ01Lb/Oj1bur9x694rYSRjhE9G41gOsggBXJa7ECgYEAwQIuLIX2Zv1s/JdGAv7zseVjpw8LYet5KPlSj71Xvw7oRrkCcSmKBfqBiYp1bDBkoVbBmcCNHHoKNQrUjSD1Sm4WA7PqsUWVN1MDnEHE5nXHbzqVY5ItYEutJb1IwKqi30iDv4CrQl3PIHM8pshpsxRm1AprXDQeyRpCdpLP7wECgYAmi5qsLP5I+1D++1Z83z3oAXzfXfVRRYXAW+0ePHpp8nclUaZ1eiWT+ZUFsK5wkevsIJO/WSo/Mv2O9JciQ2OL6o1O41c6r9zHGZbWLCekbQRXwJukP/oA/yjcooAKA8RsImy4UPa0jVYVHHnBazuWxb7RTApWnODPvPB9byy6QQ==', '461d4f0c7d6d924a226796e43857f945', 'rsa', '', '2024-08-02 17:46:03.775', '1');
INSERT INTO `app_info` VALUES (15, '2024-08-02 17:46:25.071', '2024-08-02 17:46:25.071', NULL, '东算细算1', '小万', '$2a$04$F5VPmxHVG02xYGbsKeQxf.u4d5cLMjTJULBonesh4iaPHulmwsD1i', '12334556', '文旅平台', 'MIIBCgKCAQEA5bdm2ql5VoTJEOrQBeQWq1ncYlvDryFYJzYZCWJPzhns7eYgNptfQqCY/VmzScc7jurt9XIgNLDQij7Xzmtt6cI6I8eheEDvBg/XXNJCAvsyoYkdSeUinC27SxQLSFOmHceqGjNIzKikuzi42oRntO67eCjOJtbwEFljEoBRns9PXmzUOo+qOYKHbE1T3WnyhY+A0oZX09TLrpu8Cd5bPnX3jONNZtUoQwmfKbI5fmvVgfikTmNuoApDkCbCUOADPuKMSyXNkPq5jIDPWtsTasSgGws94CHXxmqOEajAHK2kGSEwIfO3D+DdL9AW5w/tVEEbAIStR95ZWIzGHXx5CwIDAQAB', 'dec85f35cb7be717f42d46dea935a087', 'MIIEpAIBAAKCAQEA5bdm2ql5VoTJEOrQBeQWq1ncYlvDryFYJzYZCWJPzhns7eYgNptfQqCY/VmzScc7jurt9XIgNLDQij7Xzmtt6cI6I8eheEDvBg/XXNJCAvsyoYkdSeUinC27SxQLSFOmHceqGjNIzKikuzi42oRntO67eCjOJtbwEFljEoBRns9PXmzUOo+qOYKHbE1T3WnyhY+A0oZX09TLrpu8Cd5bPnX3jONNZtUoQwmfKbI5fmvVgfikTmNuoApDkCbCUOADPuKMSyXNkPq5jIDPWtsTasSgGws94CHXxmqOEajAHK2kGSEwIfO3D+DdL9AW5w/tVEEbAIStR95ZWIzGHXx5CwIDAQABAoIBAQDicjS2ALJYE3WRc6e05u9h4qdZXlQ8Y4ZewPlVQtCgvvXMXjoLGoaBS4cE2FPgusF/VR9WATa2OkaYwPJDzar/8DMX6kbGIFx/gUYvfiUMg0uZ6Wwx1+qSQpKKpCurFKqAWL7aLQvqk6Qv2u5+ml8RAPVUsgTmTKNBupvrNQtAnAWW8R5NpoPEjViDHQhaHvcDOIWn6di/I+s/QmKfw4YX/a5sflVHB7GuPRZ6tU1IP+w530M9WOvpVZNkVsJOa7waCZJk8VkBhpEVtAxLfNW/MvtyRlbQwQzQt/+3GfGEFBE6Ppo3uptupTg0DJhe1Bd1CZTQBUq+4XJlDrVThr/BAoGBAOm73jwfTUlTYGuz943NiTqMhssZGF7teWjcQDBuBwXl+Fj/41fMhvyc7TB8f5u3Hy39w7Ril0r/CBRA+pWD4XY5XCLS/CflB2/MKcHF1FNLlVCdFMFnJLDfu39S70iyCgirXb+Ad9jBwFS7CrKw7A4hMTO4cKtVJUuWa2EgCfJzAoGBAPuZjyDOKN6cvc17CQ9PZKMa4GIITZ50FrZrh43yM5ddU3Ka1ZGAg0xbQSOYDdv4KVZ3VplcDKU2hQLfe2aWMP4jkh+s/j2yi7odI8TgzX5wBN5UupVMRoBeovI3xHrlSIaHXM6E9zg4OB99W9uVTd3M1rU0KWULUmrSb6eFbYEJAoGASXYolRDpg6BW78LyjXkKJoLAYtkSVdOhkxmWeyI5xtrOwXo7g/7edksEKXcXzGMzw4q8ldde4c1iRqtdltyKYlQI51EUu506RyMOBL8eysc4uMuGGHtIhubq0UvJtTlv+eWSY3wHLvNGddsLs7nzl0VLZ4wSrWbnf+hY+3/MYI0CgYEAiLhYY65D3Tx1k5yBPtWYzV1Eg2EZtLpGwVhLcCImUOGBAC0NUyTvtGV+TzZIibkf1YHCCFJR+NFYlEgXCZciBmgT9qyN9773WOqOzmSOpl9+9cY7Ifgx8335tAAD//hqYdha4Uq2ANHkBZCcgwVuK9Is9oji2aq+XZf+3pyN9ekCgYAgRfaidTWGEnAfT4aFfS0QlBsH7hMA68POUYWw8ciy52EPgcyeoLem0OhGIrNaoXGfrAuiThuFX3t582WLMz+mqe7b8M5Ah9PhmTw8CAtp7hNS2OBK2SCLMe/hBD2/giC3BZzoktWtHDpE/G0sLeCD465/RLNWEj1WtesDC8gyEw==', '7d0775d2d7a753c18d5361309dec9362', 'rsa', '', '2024-08-02 17:46:25.071', '1');
INSERT INTO `app_info` VALUES (16, '2024-08-02 17:46:28.822', '2024-08-02 17:46:28.822', NULL, '东算细算2', '小万', '$2a$04$pp5SXN0m9EFAoSgs5a5p.ejZX5SCr0CWbjFKndik16AoVnHB864dW', '12334556', '文旅平台', 'MIIBCgKCAQEAuGyPwuavQIW/kpMV+3VVnCk3wWQxYFH1C6gW8YG82LOjBpwZhr4sgTuKMNIrmiMHPmUZU0tWTiMprsRie0IG/dOzklL/qrZ44vRti5bkJP+NH9vRp0n5cuKGzPpUojyvDFu/uP7CVLNoJVLHsaV3t1sECzYhjj7rK9B/IEXLSybEYN6A3waV8HiKJzVJgiG6tFKXRupCiYn6w3k4N9keQJ/4VhBEtQlY97Kt6vHnbam+6WDk7p1fASh284YBYIEoAOd9K66+7tomODgYxlW4Ijr4TjhXbZMrx77xSKX5GV+o9t7kZFTA/2t1FrH6Iym6mFaokhVWTEul1Md9TL6ZDQIDAQAB', '82e33e9a9e0e5b4e04dcd1edb9b00959', 'MIIEpQIBAAKCAQEAuGyPwuavQIW/kpMV+3VVnCk3wWQxYFH1C6gW8YG82LOjBpwZhr4sgTuKMNIrmiMHPmUZU0tWTiMprsRie0IG/dOzklL/qrZ44vRti5bkJP+NH9vRp0n5cuKGzPpUojyvDFu/uP7CVLNoJVLHsaV3t1sECzYhjj7rK9B/IEXLSybEYN6A3waV8HiKJzVJgiG6tFKXRupCiYn6w3k4N9keQJ/4VhBEtQlY97Kt6vHnbam+6WDk7p1fASh284YBYIEoAOd9K66+7tomODgYxlW4Ijr4TjhXbZMrx77xSKX5GV+o9t7kZFTA/2t1FrH6Iym6mFaokhVWTEul1Md9TL6ZDQIDAQABAoIBAGFWezIlS05EyW31T97eeQbpAlKCLUVnPGyglSt2jFPbzCnK7fYeHaM/55oI85Zw8nvOJGXAF5c6G9/wKwJfQQ1rVRXf1K0njfpLhB3u0A9vCe9MuAURmCU4xmkdntnaKZHeQvzDCZJm9orsfBWY/ey6jxHTyjYnozwG2XIRJFFPCU9cjkyYHJNXiAJYbjC9BEtrYyJWHivd5+M4JrIzu+rvqIwTehJJD4uxUVqCtAyToSPljLrrXWVE7ygaApG8nR3k2jctvkcLaebdzDOBhDdfq6cVLM4PXuP2mwYhxfT3H2HQo0vmYHsmPbD+639Ie+L0c44P3CVvwT/JO3W+2YkCgYEA0nRFqsp3QeeN3FkTBwLAr/WmPsLClfbLMSM+rVdb8guqIxeDW3rMjdcnWkIf4SDeVzI5vzqa1ALrnGSwXiIg6L9WLCxsvowxyccaOxxpYnlxqGWmW3gHNvoFmDPpiYNvVnBt/gNIIZrmg2mORm1eCHEfXh5iFKcRkHyFKkZDcJsCgYEA4FYmnKt2fygxRjjSuY5R83zZ6rtqj3kBCwM9HmA6C31Uii9/JOVekqMitqvLkRp0l/z3+8uagyxIwFnOJJXJoXpsFrEiIwu2lCtQFL89d1tX5FJ8q5NKOn8Gk8uBPf5fmdlRK8vc2pbXNQfqVznYGymYungsvBAZClIZD+BZU3cCgYEAjCq2r87TaQiQ2j5VRukVbtxcUhajRQAvStXfi9HQQkB0m5Swf0Aldx7ibeH2m5mpnEeuYGfm8OfVBr+lG7z9UCcHfLpF3w8+pdY+6Tm+EsEx6udwyiECJSKWkU2Z8GWixN5y06hZ5U/m0YV/Jfb3lwXIz4XK61hugV39zy8nAZsCgYEAhYNAX6W8VFk7WiufXkfRhjQcDZOVC+ZcNI2s5lMmuIfStJsn9HF+I49rrYP5F1bDZWFsxTKbts24TqIkCaoL4krV4XtycaEa/Hv6oR5gGpUGbiCZihOS7dWDuJo5RQgLJaksogAmZQG4+xrG4TrP3vqWpwO2vAZooLfTqps/DmkCgYEAtjyvJuu/AEYI9Xg6Ej0S05ch0aIsyuswnEAV0k7AN8kPXEKQ30bR5UAUTzMZPtEKMPGYKpDnB6wDsI3FLUa/vE65574huA2pvRdni04BCXztCgCYKNLgYA/seYsJzzq4i+lI+TSjzKsUUMjlhfIwJYmIJdbStla3HSRpHhzuZjs=', 'a38bc9153af5986370b308c03f48da71', 'rsa', '', '2024-08-02 17:46:28.821', '1');

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
  `req_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求内容',
  `res_msg` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '响应消息',
  `res_size` int NULL DEFAULT NULL COMMENT '响应字节数',
  `res_err` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '响应错误',
  `cost` bigint NULL DEFAULT NULL COMMENT '耗时(ms)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 92 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '服务的访问日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_log
-- ----------------------------
INSERT INTO `app_log` VALUES (58, '2024-08-02 10:58:33.749', '2024-08-02 10:58:33.749', NULL, 10, '业务应用用户账号存储', '/v1/app/account/bind', 'POST', '::1', '', '', -1, '', 956);
INSERT INTO `app_log` VALUES (59, '2024-08-02 14:18:15.227', '2024-08-02 14:18:15.227', NULL, 10, '业务应用用户账号存储', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"string\",\r\n  \"user_level\": \"string\",\r\n  \"user_name\": \"string\",\r\n  \"user_phone\": \"string\",\r\n  \"user_source\": \"string\",\r\n  \"user_source_type\": \"string\",\r\n  \"user_type\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"072b287765436146d41d35ae93881d6d\"}', 70, '', 109519);
INSERT INTO `app_log` VALUES (60, '2024-08-02 14:20:02.278', '2024-08-02 14:20:02.278', NULL, 10, '业务应用用户账号存储', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"17ae8c405bd8434043cc61c545c84ca1\"}', 70, '', 217676);
INSERT INTO `app_log` VALUES (61, '2024-08-02 14:20:04.283', '2024-08-02 14:20:04.283', NULL, 10, '业务应用用户账号存储', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"当前的昵称已经被占用,请重新填写\",\"data\":null}', 79, '', 99370);
INSERT INTO `app_log` VALUES (62, '2024-08-02 14:20:13.285', '2024-08-02 14:20:13.285', NULL, 10, '业务应用用户账号存储', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总1\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"4b1d8321355f36d6a99186124344179a\"}', 70, '', 324985);
INSERT INTO `app_log` VALUES (63, '2024-08-02 14:20:17.808', '2024-08-02 14:20:17.808', NULL, 10, '业务应用用户账号存储', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总1\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"当前的昵称已经被占用,请重新填写\",\"data\":null}', 79, '', 162032);
INSERT INTO `app_log` VALUES (64, '2024-08-02 14:20:22.582', '2024-08-02 14:20:22.582', NULL, 10, '业务应用用户账号存储', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总2\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"d6123b40f5483b1905637ed846f62c93\"}', 70, '', 324715);
INSERT INTO `app_log` VALUES (65, '2024-08-02 14:21:27.871', '2024-08-02 14:21:27.871', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总2\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"当前的昵称已经被占用,请重新填写\",\"data\":null}', 79, '', 90487);
INSERT INTO `app_log` VALUES (66, '2024-08-02 14:21:32.640', '2024-08-02 14:21:32.640', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总3\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"c0f15a79d896751f703186a8a867a187\"}', 70, '', 153559);
INSERT INTO `app_log` VALUES (67, '2024-08-02 14:21:37.699', '2024-08-02 14:21:37.699', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总4\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"33f9bb2808f4ca72fddb08d08c84f8d2\"}', 70, '', 128528);
INSERT INTO `app_log` VALUES (68, '2024-08-02 14:24:33.469', '2024-08-02 14:24:33.469', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总4\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"当前的昵称已经被占用,请重新填写\",\"data\":null}', 79, '', 255672);
INSERT INTO `app_log` VALUES (69, '2024-08-02 14:56:32.002', '2024-08-02 14:56:32.002', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总4\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"当前的昵称已经被占用,请重新填写\",\"data\":null}', 79, '', 121321);
INSERT INTO `app_log` VALUES (70, '2024-08-02 14:56:39.692', '2024-08-02 14:56:39.692', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总5\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"Error 1406 (22001): Data too long for column \'user_source\' at row 1\",\"data\":null}', 100, '', 38106);
INSERT INTO `app_log` VALUES (71, '2024-08-02 14:58:20.987', '2024-08-02 14:58:20.987', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总5\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"Error 1406 (22001): Data too long for column \'user_source\' at row 1\",\"data\":null}', 100, '', 175946);
INSERT INTO `app_log` VALUES (72, '2024-08-02 14:59:25.707', '2024-08-02 14:59:25.707', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总5\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"Error 1406 (22001): Data too long for column \'user_source\' at row 1\",\"data\":null}', 100, '', 222210);
INSERT INTO `app_log` VALUES (73, '2024-08-02 14:59:41.251', '2024-08-02 14:59:41.251', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总5\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"Error 1406 (22001): Data too long for column \'user_source\' at row 1\",\"data\":null}', 100, '', 151563);
INSERT INTO `app_log` VALUES (74, '2024-08-02 15:00:27.133', '2024-08-02 15:00:27.133', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总5\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"Error 1406 (22001): Data too long for column \'user_source\' at row 1\",\"data\":null}', 100, '', 136644);
INSERT INTO `app_log` VALUES (75, '2024-08-02 15:01:29.294', '2024-08-02 15:01:29.294', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总5\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"Error 1406 (22001): Data too long for column \'user_source\' at row 1\",\"data\":null}', 100, '', 71956);
INSERT INTO `app_log` VALUES (76, '2024-08-02 15:01:42.375', '2024-08-02 15:01:42.375', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总5\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"Error 1406 (22001): Data too long for column \'user_source\' at row 1\",\"data\":null}', 100, '', 27107);
INSERT INTO `app_log` VALUES (77, '2024-08-02 15:02:47.934', '2024-08-02 15:02:47.934', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总5\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"Error 1406 (22001): Data too long for column \'user_source\' at row 1\",\"data\":null}', 100, '', 62092);
INSERT INTO `app_log` VALUES (78, '2024-08-02 15:04:41.149', '2024-08-02 15:04:41.149', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总5\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"当前的昵称已经被占用,请重新填写\",\"data\":null}', 79, '', 78624);
INSERT INTO `app_log` VALUES (79, '2024-08-02 15:04:48.868', '2024-08-02 15:04:48.868', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总6\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"fede4b279877530a9ab9639ddedbf8a4\"}', 70, '', 78358);
INSERT INTO `app_log` VALUES (80, '2024-08-02 15:44:50.416', '2024-08-02 15:44:50.416', NULL, 10, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"你好小艾\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用已经存在\"}', 62, '', 534);
INSERT INTO `app_log` VALUES (81, '2024-08-02 15:44:54.467', '2024-08-02 15:44:54.467', NULL, 11, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"你好小艾1\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用注册成功\"}', 62, '', 43186);
INSERT INTO `app_log` VALUES (82, '2024-08-02 16:26:08.584', '2024-08-02 16:26:08.584', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总6\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"当前的昵称已经被占用,请重新填写\",\"data\":null}', 79, '', 73297);
INSERT INTO `app_log` VALUES (83, '2024-08-02 16:31:26.273', '2024-08-02 16:31:26.273', NULL, 10, '业务应用用户账号存储', '/v1/app/account/bind', 'POST', '::1', '{\r\n    \"app_name\":\"11\",\r\n    \"app_counts\": [\r\n    {\r\n      \"account\": \"1836447715@qq.com\",\r\n      \"password\": \"1836447715gzwG$\",\r\n      \"account_type\":\"QQ\"\r\n    }\r\n  ]\r\n}', '', -1, '', 0);
INSERT INTO `app_log` VALUES (84, '2024-08-02 16:39:42.536', '2024-08-02 16:39:42.536', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总6\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"当前的昵称已经被占用,请重新填写\",\"data\":null}', 79, '', 206305);
INSERT INTO `app_log` VALUES (85, '2024-08-02 16:39:48.217', '2024-08-02 16:39:48.217', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总7\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"ea7bf7ae5c26b4faacbc873ce8ff5ee8\"}', 70, '', 175121);
INSERT INTO `app_log` VALUES (86, '2024-08-02 16:43:44.438', '2024-08-02 16:43:44.438', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总7\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"当前的昵称已经被占用,请重新填写\",\"data\":null}', 79, '', 69656);
INSERT INTO `app_log` VALUES (87, '2024-08-02 16:43:50.735', '2024-08-02 16:43:50.735', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总8\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"785d59f44ab0af3a821226a9c0510730\"}', 70, '', 65596);
INSERT INTO `app_log` VALUES (88, '2024-08-02 16:45:33.594', '2024-08-02 16:45:33.594', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总8\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"当前的昵称已经被占用,请重新填写\",\"data\":null}', 79, '', 63936);
INSERT INTO `app_log` VALUES (89, '2024-08-02 16:45:34.207', '2024-08-02 16:45:34.207', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总8\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"当前的昵称已经被占用,请重新填写\",\"data\":null}', 79, '', 77673);
INSERT INTO `app_log` VALUES (90, '2024-08-02 16:45:34.859', '2024-08-02 16:45:34.859', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总8\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"当前的昵称已经被占用,请重新填写\",\"data\":null}', 79, '', 73089);
INSERT INTO `app_log` VALUES (91, '2024-08-02 16:45:38.757', '2024-08-02 16:45:38.757', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/user/register', 'POST', '::1', '{\r\n  \"nick_name\": \"万总9\",\r\n  \"user_level\": \"0\",\r\n  \"user_name\": \"万宝龙\",\r\n  \"user_phone\": \"11111111\",\r\n  \"user_source\": \"3\",\r\n  \"user_source_type\": \"小万商城\",\r\n  \"user_type\": \"1\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"fea15bca0c82996ef108942b6b9bc0f7\"}', 70, '', 57267);
INSERT INTO `app_log` VALUES (92, '2024-08-02 17:45:15.205', '2024-08-02 17:45:15.205', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":1,\r\n    \"page_size\":3\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"文旅平台用户db\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAv5qs7oi8ox1oOWzXMmTb267tmvwU6Xy+Nr/N6XWfyw5oNJnAgA8o+UkpsCVw8Ul2pWzyJGgiZ79jRS/MHY8DpmGz5bv+yvjHk1bLH9TLW58P+syP0jpCx5qOD0XN/1YXR+hkp4vwbpMArsSN+BjuwxJ1QvWGP/7Zhda0UkZwCRrqH3+gumnVA3fqG+uNZ85ibY3uEQ2HwHaK/+/22PH2Z93IK8sKIROxcNbRPlvwnK1iPuRwRo9+pNUMxHZAmUR0WyOoeRJ2y4YnoAtLV+KkLTf8rd6wK7o1ShEgR3L5AQZsdHSYBYpGwuk9EWX4lzsIZuIAPJ18kyDfVV9UFTW/4wIDAQAB\",\"app_public_md5\":\"4bea3e9a1d29cc78245081f4b899252d\",\"app_private_md5\":\"7a61748679660a66a206323c8acf50b6\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB\",\"app_public_md5\":\"2c752aa85774005612c593c0f9156c3a\",\"app_private_md5\":\"a32fa7df4a0b418fdedec4c6455bf87e\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB\",\"app_public_md5\":\"6cfb71885d97715d9817542bcc16784f\",\"app_private_md5\":\"96016fb3a78bdb4419389c7e8eaec3cc\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 1867, '', 3433);
INSERT INTO `app_log` VALUES (93, '2024-08-02 17:45:45.540', '2024-08-02 17:45:45.540', NULL, 12, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"你好小艾33331\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用注册成功\"}', 62, '', 157348);
INSERT INTO `app_log` VALUES (94, '2024-08-02 17:45:48.780', '2024-08-02 17:45:48.780', NULL, 13, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"你好小艾33331333\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用注册成功\"}', 62, '', 115473);
INSERT INTO `app_log` VALUES (95, '2024-08-02 17:46:03.778', '2024-08-02 17:46:03.778', NULL, 14, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"东算细算\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用注册成功\"}', 62, '', 235669);
INSERT INTO `app_log` VALUES (96, '2024-08-02 17:46:04.394', '2024-08-02 17:46:04.394', NULL, 14, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"东算细算\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用已经存在\"}', 62, '', 1592);
INSERT INTO `app_log` VALUES (97, '2024-08-02 17:46:04.997', '2024-08-02 17:46:04.997', NULL, 14, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"东算细算\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用已经存在\"}', 62, '', 1856);
INSERT INTO `app_log` VALUES (98, '2024-08-02 17:46:05.477', '2024-08-02 17:46:05.477', NULL, 14, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"东算细算\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用已经存在\"}', 62, '', 1432);
INSERT INTO `app_log` VALUES (99, '2024-08-02 17:46:06.011', '2024-08-02 17:46:06.011', NULL, 14, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"东算细算\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用已经存在\"}', 62, '', 1131);
INSERT INTO `app_log` VALUES (100, '2024-08-02 17:46:06.435', '2024-08-02 17:46:06.435', NULL, 14, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"东算细算\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用已经存在\"}', 62, '', 1496);
INSERT INTO `app_log` VALUES (101, '2024-08-02 17:46:06.831', '2024-08-02 17:46:06.831', NULL, 14, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"东算细算\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用已经存在\"}', 62, '', 1464);
INSERT INTO `app_log` VALUES (102, '2024-08-02 17:46:07.154', '2024-08-02 17:46:07.154', NULL, 14, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"东算细算\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用已经存在\"}', 62, '', 545);
INSERT INTO `app_log` VALUES (103, '2024-08-02 17:46:25.074', '2024-08-02 17:46:25.074', NULL, 15, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"东算细算1\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用注册成功\"}', 62, '', 81117);
INSERT INTO `app_log` VALUES (104, '2024-08-02 17:46:28.824', '2024-08-02 17:46:28.824', NULL, 16, '平台注册业务应用服务', '/v1/app/register', 'POST', '::1', '{\r\n  \"algo\": \"rsa\",\r\n  \"app_name\": \"东算细算2\",\r\n  \"app_password\": \"123456\",\r\n  \"app_phone\": \"12334556\",\r\n  \"app_type\": \"文旅平台\",\r\n  \"app_user\": \"小万\",\r\n  \"store_key\": \"string\"\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":\"业务应用注册成功\"}', 62, '', 110495);
INSERT INTO `app_log` VALUES (105, '2024-08-02 17:46:50.222', '2024-08-02 17:46:50.222', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":1,\r\n    \"page_size\":3\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"文旅平台用户db\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAv5qs7oi8ox1oOWzXMmTb267tmvwU6Xy+Nr/N6XWfyw5oNJnAgA8o+UkpsCVw8Ul2pWzyJGgiZ79jRS/MHY8DpmGz5bv+yvjHk1bLH9TLW58P+syP0jpCx5qOD0XN/1YXR+hkp4vwbpMArsSN+BjuwxJ1QvWGP/7Zhda0UkZwCRrqH3+gumnVA3fqG+uNZ85ibY3uEQ2HwHaK/+/22PH2Z93IK8sKIROxcNbRPlvwnK1iPuRwRo9+pNUMxHZAmUR0WyOoeRJ2y4YnoAtLV+KkLTf8rd6wK7o1ShEgR3L5AQZsdHSYBYpGwuk9EWX4lzsIZuIAPJ18kyDfVV9UFTW/4wIDAQAB\",\"app_public_md5\":\"4bea3e9a1d29cc78245081f4b899252d\",\"app_private_md5\":\"7a61748679660a66a206323c8acf50b6\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB\",\"app_public_md5\":\"2c752aa85774005612c593c0f9156c3a\",\"app_private_md5\":\"a32fa7df4a0b418fdedec4c6455bf87e\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB\",\"app_public_md5\":\"6cfb71885d97715d9817542bcc16784f\",\"app_private_md5\":\"96016fb3a78bdb4419389c7e8eaec3cc\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 1867, '', 1064);
INSERT INTO `app_log` VALUES (106, '2024-08-02 17:47:03.421', '2024-08-02 17:47:03.421', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":1,\r\n    \"page_size\":3\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"文旅平台用户db\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAv5qs7oi8ox1oOWzXMmTb267tmvwU6Xy+Nr/N6XWfyw5oNJnAgA8o+UkpsCVw8Ul2pWzyJGgiZ79jRS/MHY8DpmGz5bv+yvjHk1bLH9TLW58P+syP0jpCx5qOD0XN/1YXR+hkp4vwbpMArsSN+BjuwxJ1QvWGP/7Zhda0UkZwCRrqH3+gumnVA3fqG+uNZ85ibY3uEQ2HwHaK/+/22PH2Z93IK8sKIROxcNbRPlvwnK1iPuRwRo9+pNUMxHZAmUR0WyOoeRJ2y4YnoAtLV+KkLTf8rd6wK7o1ShEgR3L5AQZsdHSYBYpGwuk9EWX4lzsIZuIAPJ18kyDfVV9UFTW/4wIDAQAB\",\"app_public_md5\":\"4bea3e9a1d29cc78245081f4b899252d\",\"app_private_md5\":\"7a61748679660a66a206323c8acf50b6\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB\",\"app_public_md5\":\"2c752aa85774005612c593c0f9156c3a\",\"app_private_md5\":\"a32fa7df4a0b418fdedec4c6455bf87e\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB\",\"app_public_md5\":\"6cfb71885d97715d9817542bcc16784f\",\"app_private_md5\":\"96016fb3a78bdb4419389c7e8eaec3cc\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 1867, '', 1113);
INSERT INTO `app_log` VALUES (107, '2024-08-02 17:47:04.155', '2024-08-02 17:47:04.155', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":1,\r\n    \"page_size\":3\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"文旅平台用户db\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAv5qs7oi8ox1oOWzXMmTb267tmvwU6Xy+Nr/N6XWfyw5oNJnAgA8o+UkpsCVw8Ul2pWzyJGgiZ79jRS/MHY8DpmGz5bv+yvjHk1bLH9TLW58P+syP0jpCx5qOD0XN/1YXR+hkp4vwbpMArsSN+BjuwxJ1QvWGP/7Zhda0UkZwCRrqH3+gumnVA3fqG+uNZ85ibY3uEQ2HwHaK/+/22PH2Z93IK8sKIROxcNbRPlvwnK1iPuRwRo9+pNUMxHZAmUR0WyOoeRJ2y4YnoAtLV+KkLTf8rd6wK7o1ShEgR3L5AQZsdHSYBYpGwuk9EWX4lzsIZuIAPJ18kyDfVV9UFTW/4wIDAQAB\",\"app_public_md5\":\"4bea3e9a1d29cc78245081f4b899252d\",\"app_private_md5\":\"7a61748679660a66a206323c8acf50b6\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB\",\"app_public_md5\":\"2c752aa85774005612c593c0f9156c3a\",\"app_private_md5\":\"a32fa7df4a0b418fdedec4c6455bf87e\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB\",\"app_public_md5\":\"6cfb71885d97715d9817542bcc16784f\",\"app_private_md5\":\"96016fb3a78bdb4419389c7e8eaec3cc\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 1867, '', 1178);
INSERT INTO `app_log` VALUES (108, '2024-08-02 17:47:04.763', '2024-08-02 17:47:04.763', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":1,\r\n    \"page_size\":3\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"文旅平台用户db\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAv5qs7oi8ox1oOWzXMmTb267tmvwU6Xy+Nr/N6XWfyw5oNJnAgA8o+UkpsCVw8Ul2pWzyJGgiZ79jRS/MHY8DpmGz5bv+yvjHk1bLH9TLW58P+syP0jpCx5qOD0XN/1YXR+hkp4vwbpMArsSN+BjuwxJ1QvWGP/7Zhda0UkZwCRrqH3+gumnVA3fqG+uNZ85ibY3uEQ2HwHaK/+/22PH2Z93IK8sKIROxcNbRPlvwnK1iPuRwRo9+pNUMxHZAmUR0WyOoeRJ2y4YnoAtLV+KkLTf8rd6wK7o1ShEgR3L5AQZsdHSYBYpGwuk9EWX4lzsIZuIAPJ18kyDfVV9UFTW/4wIDAQAB\",\"app_public_md5\":\"4bea3e9a1d29cc78245081f4b899252d\",\"app_private_md5\":\"7a61748679660a66a206323c8acf50b6\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB\",\"app_public_md5\":\"2c752aa85774005612c593c0f9156c3a\",\"app_private_md5\":\"a32fa7df4a0b418fdedec4c6455bf87e\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB\",\"app_public_md5\":\"6cfb71885d97715d9817542bcc16784f\",\"app_private_md5\":\"96016fb3a78bdb4419389c7e8eaec3cc\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 1867, '', 1037);
INSERT INTO `app_log` VALUES (109, '2024-08-02 17:47:06.408', '2024-08-02 17:47:06.408', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":1,\r\n    \"page_size\":3\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"文旅平台用户db\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAv5qs7oi8ox1oOWzXMmTb267tmvwU6Xy+Nr/N6XWfyw5oNJnAgA8o+UkpsCVw8Ul2pWzyJGgiZ79jRS/MHY8DpmGz5bv+yvjHk1bLH9TLW58P+syP0jpCx5qOD0XN/1YXR+hkp4vwbpMArsSN+BjuwxJ1QvWGP/7Zhda0UkZwCRrqH3+gumnVA3fqG+uNZ85ibY3uEQ2HwHaK/+/22PH2Z93IK8sKIROxcNbRPlvwnK1iPuRwRo9+pNUMxHZAmUR0WyOoeRJ2y4YnoAtLV+KkLTf8rd6wK7o1ShEgR3L5AQZsdHSYBYpGwuk9EWX4lzsIZuIAPJ18kyDfVV9UFTW/4wIDAQAB\",\"app_public_md5\":\"4bea3e9a1d29cc78245081f4b899252d\",\"app_private_md5\":\"7a61748679660a66a206323c8acf50b6\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB\",\"app_public_md5\":\"2c752aa85774005612c593c0f9156c3a\",\"app_private_md5\":\"a32fa7df4a0b418fdedec4c6455bf87e\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB\",\"app_public_md5\":\"6cfb71885d97715d9817542bcc16784f\",\"app_private_md5\":\"96016fb3a78bdb4419389c7e8eaec3cc\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 1867, '', 1154);
INSERT INTO `app_log` VALUES (110, '2024-08-02 17:47:06.988', '2024-08-02 17:47:06.988', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":1,\r\n    \"page_size\":3\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"文旅平台用户db\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAv5qs7oi8ox1oOWzXMmTb267tmvwU6Xy+Nr/N6XWfyw5oNJnAgA8o+UkpsCVw8Ul2pWzyJGgiZ79jRS/MHY8DpmGz5bv+yvjHk1bLH9TLW58P+syP0jpCx5qOD0XN/1YXR+hkp4vwbpMArsSN+BjuwxJ1QvWGP/7Zhda0UkZwCRrqH3+gumnVA3fqG+uNZ85ibY3uEQ2HwHaK/+/22PH2Z93IK8sKIROxcNbRPlvwnK1iPuRwRo9+pNUMxHZAmUR0WyOoeRJ2y4YnoAtLV+KkLTf8rd6wK7o1ShEgR3L5AQZsdHSYBYpGwuk9EWX4lzsIZuIAPJ18kyDfVV9UFTW/4wIDAQAB\",\"app_public_md5\":\"4bea3e9a1d29cc78245081f4b899252d\",\"app_private_md5\":\"7a61748679660a66a206323c8acf50b6\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB\",\"app_public_md5\":\"2c752aa85774005612c593c0f9156c3a\",\"app_private_md5\":\"a32fa7df4a0b418fdedec4c6455bf87e\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB\",\"app_public_md5\":\"6cfb71885d97715d9817542bcc16784f\",\"app_private_md5\":\"96016fb3a78bdb4419389c7e8eaec3cc\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 1867, '', 1046);
INSERT INTO `app_log` VALUES (111, '2024-08-02 17:47:07.596', '2024-08-02 17:47:07.596', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":1,\r\n    \"page_size\":3\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"文旅平台用户db\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAv5qs7oi8ox1oOWzXMmTb267tmvwU6Xy+Nr/N6XWfyw5oNJnAgA8o+UkpsCVw8Ul2pWzyJGgiZ79jRS/MHY8DpmGz5bv+yvjHk1bLH9TLW58P+syP0jpCx5qOD0XN/1YXR+hkp4vwbpMArsSN+BjuwxJ1QvWGP/7Zhda0UkZwCRrqH3+gumnVA3fqG+uNZ85ibY3uEQ2HwHaK/+/22PH2Z93IK8sKIROxcNbRPlvwnK1iPuRwRo9+pNUMxHZAmUR0WyOoeRJ2y4YnoAtLV+KkLTf8rd6wK7o1ShEgR3L5AQZsdHSYBYpGwuk9EWX4lzsIZuIAPJ18kyDfVV9UFTW/4wIDAQAB\",\"app_public_md5\":\"4bea3e9a1d29cc78245081f4b899252d\",\"app_private_md5\":\"7a61748679660a66a206323c8acf50b6\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB\",\"app_public_md5\":\"2c752aa85774005612c593c0f9156c3a\",\"app_private_md5\":\"a32fa7df4a0b418fdedec4c6455bf87e\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB\",\"app_public_md5\":\"6cfb71885d97715d9817542bcc16784f\",\"app_private_md5\":\"96016fb3a78bdb4419389c7e8eaec3cc\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 1867, '', 2055);
INSERT INTO `app_log` VALUES (112, '2024-08-02 17:47:08.284', '2024-08-02 17:47:08.284', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":1,\r\n    \"page_size\":3\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"文旅平台用户db\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAv5qs7oi8ox1oOWzXMmTb267tmvwU6Xy+Nr/N6XWfyw5oNJnAgA8o+UkpsCVw8Ul2pWzyJGgiZ79jRS/MHY8DpmGz5bv+yvjHk1bLH9TLW58P+syP0jpCx5qOD0XN/1YXR+hkp4vwbpMArsSN+BjuwxJ1QvWGP/7Zhda0UkZwCRrqH3+gumnVA3fqG+uNZ85ibY3uEQ2HwHaK/+/22PH2Z93IK8sKIROxcNbRPlvwnK1iPuRwRo9+pNUMxHZAmUR0WyOoeRJ2y4YnoAtLV+KkLTf8rd6wK7o1ShEgR3L5AQZsdHSYBYpGwuk9EWX4lzsIZuIAPJ18kyDfVV9UFTW/4wIDAQAB\",\"app_public_md5\":\"4bea3e9a1d29cc78245081f4b899252d\",\"app_private_md5\":\"7a61748679660a66a206323c8acf50b6\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB\",\"app_public_md5\":\"2c752aa85774005612c593c0f9156c3a\",\"app_private_md5\":\"a32fa7df4a0b418fdedec4c6455bf87e\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB\",\"app_public_md5\":\"6cfb71885d97715d9817542bcc16784f\",\"app_private_md5\":\"96016fb3a78bdb4419389c7e8eaec3cc\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 1867, '', 1022);
INSERT INTO `app_log` VALUES (113, '2024-08-02 17:47:23.919', '2024-08-02 17:47:23.919', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":1,\r\n    \"page_size\":3\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"文旅平台用户db\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAv5qs7oi8ox1oOWzXMmTb267tmvwU6Xy+Nr/N6XWfyw5oNJnAgA8o+UkpsCVw8Ul2pWzyJGgiZ79jRS/MHY8DpmGz5bv+yvjHk1bLH9TLW58P+syP0jpCx5qOD0XN/1YXR+hkp4vwbpMArsSN+BjuwxJ1QvWGP/7Zhda0UkZwCRrqH3+gumnVA3fqG+uNZ85ibY3uEQ2HwHaK/+/22PH2Z93IK8sKIROxcNbRPlvwnK1iPuRwRo9+pNUMxHZAmUR0WyOoeRJ2y4YnoAtLV+KkLTf8rd6wK7o1ShEgR3L5AQZsdHSYBYpGwuk9EWX4lzsIZuIAPJ18kyDfVV9UFTW/4wIDAQAB\",\"app_public_md5\":\"4bea3e9a1d29cc78245081f4b899252d\",\"app_private_md5\":\"7a61748679660a66a206323c8acf50b6\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB\",\"app_public_md5\":\"2c752aa85774005612c593c0f9156c3a\",\"app_private_md5\":\"a32fa7df4a0b418fdedec4c6455bf87e\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB\",\"app_public_md5\":\"6cfb71885d97715d9817542bcc16784f\",\"app_private_md5\":\"96016fb3a78bdb4419389c7e8eaec3cc\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 1867, '', 1078);
INSERT INTO `app_log` VALUES (114, '2024-08-02 17:50:19.545', '2024-08-02 17:50:19.545', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":1,\r\n    \"page_size\":3\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"你好小艾\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB\",\"app_public_md5\":\"2c752aa85774005612c593c0f9156c3a\",\"app_private_md5\":\"a32fa7df4a0b418fdedec4c6455bf87e\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB\",\"app_public_md5\":\"6cfb71885d97715d9817542bcc16784f\",\"app_private_md5\":\"96016fb3a78bdb4419389c7e8eaec3cc\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾33331\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAo+ly0HVlPTPxGW/Ox+19lKosf001alJeWhtRL/BZ/H7UBQ4pkrtsyhXs+6dV4QAZa4CaiZ6jMR7X88qh1KbtO3w1pRczUBROduiAMqVkE2kqQJW3ZoUiepO9PFxDMPmw4gSxDJw0BK/Zni4Zg0nXmDVIgZwLXjOeAlnQxU/lFGzg8R89NzjlBXzz4cM45EgRwsohnLubjnuaNt7p7EIDAMEu5rBnlz8k3zNYZKEXQ3q9UK3iGQdOfyJYQuQ3bcEhWW5GvcWZ9IntiETdijdJ6JNbbGgz7GulzsnbKjb0iJ43Tp7E0MtXGme+oB3RCIKdeRofP2qZZqqgNKHqR7q6tQIDAQAB\",\"app_public_md5\":\"ff4f0402620a45b7805556147eff4390\",\"app_private_md5\":\"1cf753e64c25f10a9604c410abff581d\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 1864, '', 1547);
INSERT INTO `app_log` VALUES (115, '2024-08-02 17:50:40.589', '2024-08-02 17:50:40.589', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":1,\r\n    \"page_size\":10\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"你好小艾\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuQYqjDMyWJdChWDl6ewbWrJpCBMgPIWyuA5C30ORdP40QJT5+w5s5Qq5i+F1kccc7lkVyeT+yRj1eqdePC5qMuMpTjujXx9qLZoQgmSakXKcdogSBpgFcJoSe2Mec0YY+6iHbtGN+VtScI5ceZAJzZi+68V7O1zc2x1bYFK1kli2ozxwTTBnNtsbd6BGiQKui5qWe7ovr7ubuRzMxE3wdBdVe+XVkICaqllp5qqmiVxJljB03jHAu3n38MxXtmVekJs6qiHsm3Es58YJXpBv+Z9Adz77vbQ9594Svoqy+EkzZ+tDrkY70Aj67FOKYZeMqs73kLcNgd5GWp3VkT9RxQIDAQAB\",\"app_public_md5\":\"2c752aa85774005612c593c0f9156c3a\",\"app_private_md5\":\"a32fa7df4a0b418fdedec4c6455bf87e\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAp+ti5WvJdQO+ynV7tE8kqHdO2x8V2AWOubYtu34X6dqBN3SjQGRfIJLQwa81O7Up2p4z1nSXAnsEcp3ijks4qj+sj/caUQl7NmsMI8vBEKsP9FHGTxBy4fgPpsyNoKGq4inN/kSB8fIcrLfHbeLFWd0fVbVhzHtZSJ/fThpLHbbZdYjDQ4swPnps3cCPJjEH0fNcdTVbgNCqgKh1SBLCO5Gm9k7Ow/X2sp1hQklIko9sjzplKl5ztMEmG6PWm4rgzzd5TLa1y8ygUSSZ7Qg+wkQk8l+heDkvNgO64p5yDiKxzuBP83m4tU0o17/gEhWxpx6vCLLDJsT9yB04F4EGZQIDAQAB\",\"app_public_md5\":\"6cfb71885d97715d9817542bcc16784f\",\"app_private_md5\":\"96016fb3a78bdb4419389c7e8eaec3cc\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾33331\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAo+ly0HVlPTPxGW/Ox+19lKosf001alJeWhtRL/BZ/H7UBQ4pkrtsyhXs+6dV4QAZa4CaiZ6jMR7X88qh1KbtO3w1pRczUBROduiAMqVkE2kqQJW3ZoUiepO9PFxDMPmw4gSxDJw0BK/Zni4Zg0nXmDVIgZwLXjOeAlnQxU/lFGzg8R89NzjlBXzz4cM45EgRwsohnLubjnuaNt7p7EIDAMEu5rBnlz8k3zNYZKEXQ3q9UK3iGQdOfyJYQuQ3bcEhWW5GvcWZ9IntiETdijdJ6JNbbGgz7GulzsnbKjb0iJ43Tp7E0MtXGme+oB3RCIKdeRofP2qZZqqgNKHqR7q6tQIDAQAB\",\"app_public_md5\":\"ff4f0402620a45b7805556147eff4390\",\"app_private_md5\":\"1cf753e64c25f10a9604c410abff581d\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"你好小艾33331333\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEA5JuN7ZOUaoEvNHNOW2e5YuQuouCSCdi3/vybn42SxGjv7GDZEmLgtIinyGFElJ54YdHO04HJ7seoG5LXxI1sVaFEhc3MEc29jJjD0eCDLI+0D7Qxcg1RR0LPsTjBsY5CUHl8hiWDanywqlgu2ZF2be3k4pNroyE7GT7qkGz89Q6ORDN6qB8fTAEu98i67jEfrd20feTCALgbqdtWqpKR5owN5tKzaj/IcuIsIlc4weCb0erUJaKdrY7SWQ3q4TRFi8LTxCaujtaQ2EYaFVFyZrdyt2iJeFvMsiozq0AloP59eCQUFmkfTSAfdRCxrebD366nWF1m+IZlAT2yay1akwIDAQAB\",\"app_public_md5\":\"179cac03e9e4a0860129c13d6a921c88\",\"app_private_md5\":\"0c07897c757c3df17642ba5ab1a25cca\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"东算细算\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEA0taMMRDQ4hoof20XqTWwKRt9HDq0sDz6Y7z2zLfsmMxtbp3FYH5f2HKTdieudzNvfPSYONlpsQv1J8zhXB/oOSirMJ5pEFU7FASFH2Z2aeIAERm0XYVOpCe6v34qCwLSOE448UUZ66YuAyx6mME4bl9yWM48oxnGxXP4dJ9wldGbGYo28mRam/vH2nrtJ879KBOob0/q8FnTrBRtaTbwcAnX2Q0vsHqaLjOH3P7mm/MpF44n22x7oUm6KtosFbwDP/0ag/A1ErrNVaA6UO5RCEw9HwtkXbLq9Y7/oLPx30CS01MXi4FCUixU6L7TN2ZbWN5BNLAN4Ne0R19KMdZwBQIDAQAB\",\"app_public_md5\":\"3511a0d2eeb2dd49d7af2cd79f4831f4\",\"app_private_md5\":\"461d4f0c7d6d924a226796e43857f945\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"东算细算1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEA5bdm2ql5VoTJEOrQBeQWq1ncYlvDryFYJzYZCWJPzhns7eYgNptfQqCY/VmzScc7jurt9XIgNLDQij7Xzmtt6cI6I8eheEDvBg/XXNJCAvsyoYkdSeUinC27SxQLSFOmHceqGjNIzKikuzi42oRntO67eCjOJtbwEFljEoBRns9PXmzUOo+qOYKHbE1T3WnyhY+A0oZX09TLrpu8Cd5bPnX3jONNZtUoQwmfKbI5fmvVgfikTmNuoApDkCbCUOADPuKMSyXNkPq5jIDPWtsTasSgGws94CHXxmqOEajAHK2kGSEwIfO3D+DdL9AW5w/tVEEbAIStR95ZWIzGHXx5CwIDAQAB\",\"app_public_md5\":\"dec85f35cb7be717f42d46dea935a087\",\"app_private_md5\":\"7d0775d2d7a753c18d5361309dec9362\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"东算细算2\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuGyPwuavQIW/kpMV+3VVnCk3wWQxYFH1C6gW8YG82LOjBpwZhr4sgTuKMNIrmiMHPmUZU0tWTiMprsRie0IG/dOzklL/qrZ44vRti5bkJP+NH9vRp0n5cuKGzPpUojyvDFu/uP7CVLNoJVLHsaV3t1sECzYhjj7rK9B/IEXLSybEYN6A3waV8HiKJzVJgiG6tFKXRupCiYn6w3k4N9keQJ/4VhBEtQlY97Kt6vHnbam+6WDk7p1fASh284YBYIEoAOd9K66+7tomODgYxlW4Ijr4TjhXbZMrx77xSKX5GV+o9t7kZFTA/2t1FrH6Iym6mFaokhVWTEul1Md9TL6ZDQIDAQAB\",\"app_public_md5\":\"82e33e9a9e0e5b4e04dcd1edb9b00959\",\"app_private_md5\":\"a38bc9153af5986370b308c03f48da71\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 4302, '', 1022);
INSERT INTO `app_log` VALUES (116, '2024-08-02 17:50:49.714', '2024-08-02 17:50:49.714', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":2,\r\n    \"page_size\":8\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[]}', 38, '', 1991);
INSERT INTO `app_log` VALUES (117, '2024-08-02 17:51:07.463', '2024-08-02 17:51:07.463', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":3,\r\n    \"page_size\":2\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"东算细算\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEA0taMMRDQ4hoof20XqTWwKRt9HDq0sDz6Y7z2zLfsmMxtbp3FYH5f2HKTdieudzNvfPSYONlpsQv1J8zhXB/oOSirMJ5pEFU7FASFH2Z2aeIAERm0XYVOpCe6v34qCwLSOE448UUZ66YuAyx6mME4bl9yWM48oxnGxXP4dJ9wldGbGYo28mRam/vH2nrtJ879KBOob0/q8FnTrBRtaTbwcAnX2Q0vsHqaLjOH3P7mm/MpF44n22x7oUm6KtosFbwDP/0ag/A1ErrNVaA6UO5RCEw9HwtkXbLq9Y7/oLPx30CS01MXi4FCUixU6L7TN2ZbWN5BNLAN4Ne0R19KMdZwBQIDAQAB\",\"app_public_md5\":\"3511a0d2eeb2dd49d7af2cd79f4831f4\",\"app_private_md5\":\"461d4f0c7d6d924a226796e43857f945\",\"app_algo\":\"\",\"store_key\":\"\"},{\"app_name\":\"东算细算1\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEA5bdm2ql5VoTJEOrQBeQWq1ncYlvDryFYJzYZCWJPzhns7eYgNptfQqCY/VmzScc7jurt9XIgNLDQij7Xzmtt6cI6I8eheEDvBg/XXNJCAvsyoYkdSeUinC27SxQLSFOmHceqGjNIzKikuzi42oRntO67eCjOJtbwEFljEoBRns9PXmzUOo+qOYKHbE1T3WnyhY+A0oZX09TLrpu8Cd5bPnX3jONNZtUoQwmfKbI5fmvVgfikTmNuoApDkCbCUOADPuKMSyXNkPq5jIDPWtsTasSgGws94CHXxmqOEajAHK2kGSEwIfO3D+DdL9AW5w/tVEEbAIStR95ZWIzGHXx5CwIDAQAB\",\"app_public_md5\":\"dec85f35cb7be717f42d46dea935a087\",\"app_private_md5\":\"7d0775d2d7a753c18d5361309dec9362\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 1252, '', 1059);
INSERT INTO `app_log` VALUES (118, '2024-08-02 17:51:16.647', '2024-08-02 17:51:16.647', NULL, 10, '业务应用代理用户进行主权账户注册', '/v1/app/info', 'POST', '::1', '{\r\n    \"page_num\":3,\r\n    \"page_size\":3\r\n}', '{\"code\":200,\"msg\":\"success\",\"data\":[{\"app_name\":\"东算细算2\",\"app_user\":\"小万\",\"app_phone\":\"12334556\",\"app_type\":\"文旅平台\",\"app_public\":\"MIIBCgKCAQEAuGyPwuavQIW/kpMV+3VVnCk3wWQxYFH1C6gW8YG82LOjBpwZhr4sgTuKMNIrmiMHPmUZU0tWTiMprsRie0IG/dOzklL/qrZ44vRti5bkJP+NH9vRp0n5cuKGzPpUojyvDFu/uP7CVLNoJVLHsaV3t1sECzYhjj7rK9B/IEXLSybEYN6A3waV8HiKJzVJgiG6tFKXRupCiYn6w3k4N9keQJ/4VhBEtQlY97Kt6vHnbam+6WDk7p1fASh284YBYIEoAOd9K66+7tomODgYxlW4Ijr4TjhXbZMrx77xSKX5GV+o9t7kZFTA/2t1FrH6Iym6mFaokhVWTEul1Md9TL6ZDQIDAQAB\",\"app_public_md5\":\"82e33e9a9e0e5b4e04dcd1edb9b00959\",\"app_private_md5\":\"a38bc9153af5986370b308c03f48da71\",\"app_algo\":\"\",\"store_key\":\"\"}]}', 645, '', 1054);

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of app_token
-- ----------------------------
INSERT INTO `app_token` VALUES (1, '2024-08-02 15:50:23.210', '2024-08-02 15:50:23.210', NULL, 11, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfbmFtZSI6IuS9oOWlveWwj-iJvjEiLCJpc3MiOiJjb3JlZCIsImV4cCI6MTczMDM2MTAyMywiaWF0IjoxNzIyNTg1MDIzfQ.QwdwF2O_vf9-XtSEjvc7eF28sOt6bNqi4_UluaxHbG0', '2024-10-31 15:50:23.208');

-- ----------------------------
-- Table structure for data_direct
-- ----------------------------
DROP TABLE IF EXISTS `data_direct`;
CREATE TABLE `data_direct`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '数据id',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `direct_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目录id(00000000进行表示)',
  `direct_id_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目录名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '数据目录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of data_direct
-- ----------------------------

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of data_direct1
-- ----------------------------

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of data_direct2
-- ----------------------------

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of data_direct3
-- ----------------------------

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '数据是否共享 是否需要审核 数据价值' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of data_info
-- ----------------------------

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '记录数据安全相关信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of data_safe
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '请求共享和审核来自哪个业务应用' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of data_share
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '成功共享的信息记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of data_sharelog
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '超级节点的硬件信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of node_hard
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '超级节点的登记信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of node_info
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '记录节点的通信记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of node_log
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '超级节点的软件信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of node_soft
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '超级节点的联系人信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of node_user
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'json描述的配置信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of service_cfg
-- ----------------------------

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
  `node_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '节点id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '服务代理的基本信息,包含程序' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of service_info
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统角色' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_authority
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色菜单绑定表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_authority_menu
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '主要配置通信' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_config
-- ----------------------------

-- ----------------------------
-- Table structure for sys_download
-- ----------------------------
DROP TABLE IF EXISTS `sys_download`;
CREATE TABLE `sys_download`  (
  `id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '记录软件的下载信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_download
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色菜单' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '通知各种信息，有短信和平台通知两种' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_msg
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '管理端人员的操作记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_operlog
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '描述各种升级策略' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_policy
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '软件信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_soft
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统用户' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_account
-- ----------------------------
INSERT INTO `user_account` VALUES (1, '2024-08-02 16:55:51.819', '2024-08-02 16:55:51.819', NULL, 11, '[{\"account\":\"1836447715@qq.com\",\"password\":\"1836447715gzwG$\"},{\"account\":\"1836447715@qq.com\",\"password\":\"1836447715gzwG$\"}]', '');

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户用于支付信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_bank
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用用户公钥证书 日期 算法 存储密钥' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_ca
-- ----------------------------
INSERT INTO `user_ca` VALUES (1, '2024-08-02 15:04:48.865', '2024-08-02 15:04:48.865', NULL, 0, 'MIIBCgKCAQEAocsPXEfQyow/uzteyZ4ySbNJ/nbDV9BNB5Z1Fee+DDTZBM+m9CDWOvCvgQtruvSYL38rx0Vap8YqnNfngyHPmsBOu1yC/5MWMM+PUPKqC0gSwZo0ac3d+vnhrdAGRC6Con5ypfUAF//cARM7Z0jOJp2vWCGtAwrZODMG7BN3tzKKor5y71fCQ/860CPNj+UO/nc9hPp8CMZeYR+/BSUbN3gMd2hwU253PPZGKbzYpyODUkBW81moOaYFYUR8ZLDFWwv+O8Qhk9sPXZBvBtdg4qdgbOlR8OGot9dlZp9ppULBP3L8qzoUypn6PJ9BuxZD0OZPUt8rRPgWSC1NBn40JwIDAQAB', 'f0bce38a74e8b0c93904fb44c5d5af83', 'MIIEowIBAAKCAQEAocsPXEfQyow/uzteyZ4ySbNJ/nbDV9BNB5Z1Fee+DDTZBM+m9CDWOvCvgQtruvSYL38rx0Vap8YqnNfngyHPmsBOu1yC/5MWMM+PUPKqC0gSwZo0ac3d+vnhrdAGRC6Con5ypfUAF//cARM7Z0jOJp2vWCGtAwrZODMG7BN3tzKKor5y71fCQ/860CPNj+UO/nc9hPp8CMZeYR+/BSUbN3gMd2hwU253PPZGKbzYpyODUkBW81moOaYFYUR8ZLDFWwv+O8Qhk9sPXZBvBtdg4qdgbOlR8OGot9dlZp9ppULBP3L8qzoUypn6PJ9BuxZD0OZPUt8rRPgWSC1NBn40JwIDAQABAoIBACwJhx7I9U2MKvSdTkl2Y52nzSYT0UufiLfd/fR+ZT4RTFiA5KJiJa1ZQYvQPfxfAuU6qgHNIMW1rHeQVJVeqJEocsWQgp9pu4qBlDKmOh5j+ab0f/ukax0a0pBqpKn+NtRfsnj1XJcrrwgj4fFU4belB8rnL546X7FqQiITkmwsSL+qcqEcQi8QybNpeFI1J4qp36VOeQQtPx35MoJO6qhLRHUvTMIs335TGclOluEcvBG6h0IbUty7RR0bMCAFyL/GQRfLK6I1bAbwiGZbPBSj8g4uaqbtoXLUYFO8CWQfVWSufp2xg7joX/MZOHolpIrT+9mvPNDWKkHz1mCHePECgYEAzg4h33H8nMbZEVT/OKHOv+zNCGZYFFGiYQQMu9rkK3E9tAlz+KS3Rqkbu3rfc9En9YJSePS5FdV2luDXNLDwOJzqSuEKHUgfej3QtyKZ0k816/GPQHwtnQAjdEGo1zFt8tRSDJCRLyrWCOvfvXkOhIpzfM8yQYZr3oy+57iPsQsCgYEAyQJwoMG2moUOnWFD0G98m/AIm0GBrem9Yy75HZQle6Aj8kxvAB1BV0fIZmWcJlquWvG7giL8PA146beeM1z8J4Otq+txNMFIvUYlKdQsEhgFWbNJaJtco/NSMUPI48s4htVjTwzaEr0gfywMlvPXmkycsJ8DeTMf0gICjf/rctUCgYA1MfLOs6eqda9eSRombadpZuG7tGgyvqxz7U64RJc5G9/5FaOp4iQqisfNC7iaNrnhKMbP0SXjHBukceE9ChzKEMBvmk85JgByV3kFflpFsAtUoSUBWc2QxZNhlP3YqDshQGz6L2w7yAw/e11kldcIopGT/A2WWvSIHL/AJKZ75wKBgQCsbudnQEYBL/tljra566mdq7/2T6+LHWRV3Une0eOFtc6gYGkoU/BlizKihK/7RvcjFROhg5mVeokrFVUlArNcwl0Nih/G/bgEhlQDGfvkWa7dUCu9DOnfyy/LCzceEosfP5a9jWzVfmCON59lLzDZMLeJ3B8CrZSydv7yCYRVMQKBgAs9gvHxzfaPcOIc2UTtpXEDT1rCeu7kshomZehZD9xxCaYluJ6jmHvaTVjzWMaCr9ApHm/WwTDyTwaUg5/eR4myLLW8BIgPv0Xjg0bpqZY0Y6w81xgkpIzvE1tJ4LMV5H2AX1ehqxe76TpRtIBD0lc2RnDb+KpFaDwxDqX7H3PU', 'fede4b279877530a9ab9639ddedbf8a4', 'rsa', '', 1722582288);
INSERT INTO `user_ca` VALUES (2, '2024-08-02 16:39:48.214', '2024-08-02 16:39:48.214', NULL, 0, 'MIIBCgKCAQEA0pbXa80HgeAuUIMsvE+Q0PfZ/N7z7sCFfptZCMBd+bJhu7iM10nkb3emlVlk8+x4hquE3Fx9kaqcM0JYXp+CaCJ/5aoDAOJRitcx7RIHj253OsT3e7tLOtKzS6kM3KhsrVU+sjRdltpuGAn5K7GlZOQfAK/29HfATVlSCZ0ON1fSyBlLBq1vxL3rncFPjPMzLC1sltbD35Kj7NXZe7qMgS1gEDTrUU10KYSY53OynSCWOP6DtmbwE7+2FsYVEbClqNY/weAHe1ndqj38ZFAAOmq40OtuFz17fwUN3mVWdJhWuRwWdkGNByerCW//KBZneYgFPYxQIbW3WZsMEhVCpwIDAQAB', '7ee437699aa216f6d12b80505bc9534e', 'MIIEowIBAAKCAQEA0pbXa80HgeAuUIMsvE+Q0PfZ/N7z7sCFfptZCMBd+bJhu7iM10nkb3emlVlk8+x4hquE3Fx9kaqcM0JYXp+CaCJ/5aoDAOJRitcx7RIHj253OsT3e7tLOtKzS6kM3KhsrVU+sjRdltpuGAn5K7GlZOQfAK/29HfATVlSCZ0ON1fSyBlLBq1vxL3rncFPjPMzLC1sltbD35Kj7NXZe7qMgS1gEDTrUU10KYSY53OynSCWOP6DtmbwE7+2FsYVEbClqNY/weAHe1ndqj38ZFAAOmq40OtuFz17fwUN3mVWdJhWuRwWdkGNByerCW//KBZneYgFPYxQIbW3WZsMEhVCpwIDAQABAoIBAHykt2Cg/B6NGByjZ2kWFvr4mFephuv7m4fyXuJlKbpPMLWxZ/tNSx4Gzdx8CtqtTjE3d+4N+GytdMKKNE/dTYxhTx4aYKi3S9hHws806rCki4GHb4wWsVJhv6m6p1g7gAef+vnzMKnHI36q+IY6IVD0DL2VaTANaQ7fHUV/Zhiumb2cSc0O3DWeYZACqtSu4fUlwN7TCkmdWDIwzGNnoUq7+ciCFfD7RagTJdLP+UmDX5ENSZQ0PhVU3OrgVhLSbTdlg6tyfwFxfXfRN73vnBUtYDXLrgiqs4xktupBXsIbTZTuwGcz5Gd7a0eIuNyyYtqynzqWa7q+BlFYABthFZECgYEA5Y3GGOaUJpbjSXiHwsUjds0V56IjGcEzfaC1P5X0thujbne3Bd0pYH+DTQAIxRZcALNOStobfYnOCWIFBlmYQJlvFkAZ60IlgBj82VZ/fwfqLlyPZ/CM5QszTHqPaaEzCjZxc8ORoZVFWEBqaLTdnvREPzk1LPmJvVi65xvVOk0CgYEA6tm+5IE29nFgrW2fD12BoafqCNbYe7FE1NxjjqnjNuhIU6L/0FdK4OQJd6+tsGrLVcDyjrml1YyoJyuyn9nKLPOCEYR6hzeaqsZI3IXJeCaMEV6sfOp7qDzh2ZPukx0c/sqgdpl5gG3ix8OfkV9IXIRUv431V39hi6KeDykqQsMCgYEAzmI62qRkSpWaX7H4DUsE6YjVyt3hRQSI8MgObM67pwcOVwFXEfSXgOTj/dSsDdZefVq4Z1/VG3dMQO9ZqS3gTyQj1hNnujZMVPEiU50LIADTsT8Sx9ZkNoFta7M9QTDnHKV6NyR7yJTWNdQkV2NVZdYjjdw3TOpFxnECCUri2HkCgYBcruJvO699QvFsgD6tybFHIwVxx5hX+HW5q+B+hP1uxe0FVOawSkL6zgQOf/6ECSGGDkernAieZoxWKSCthT4Mz20djLLuejtH+kj2/rWr/Up58GsSQVfb5Wscew1EcdDnnQAGzjrtPHCrdcCblobwPClWD6grh9HvUZYfVZegEQKBgGGfT+qQWCaAXaLP+md0KQPwLSVLdHJJwUFkeavpwhbxOG+4zOyuvNrKO5bHz4CyqZswtHfxjQyU5lugsVQghTRVXi/fyeFMmXSTTFuLrRKsxG4Lhc6Fru3LvLN0dRtp6XZJG4LCYfUgsdNiAAGarehI5qWRkMyY4xtUyapbkzJU', 'ea7bf7ae5c26b4faacbc873ce8ff5ee8', 'rsa', '', 1722587988);
INSERT INTO `user_ca` VALUES (3, '2024-08-02 16:43:50.731', '2024-08-02 16:43:50.731', NULL, 0, 'MIIBCgKCAQEAsqGKf2Pf5HNXrbmy2lzBF7t1g8c9bGTWzs4V/s93Voe8A98ff5oeS4WSekEwR9JnqIh1rPDRbt2c2jxEwW5Qpxfez0lJPfUkmooSOfkS0CBK+AMR8jSO8nVaeQdKuxxwnSzqtHWXat8EfVFA0cdpTCaEsuuk1kyza0BBJNTYOIVbaWpLEq7a9lGvf5OE4tYC4V2wRBnABWTmuaPIdDP7HANes+NhtYao4aRXaavuYTOpuToQfytB2ZZ2EXR1bbOYXYOZ3tYDsz+chqKzWCtOBoxZAUMFSjQ1Hj0fVv9BxA4oc4XM5YQhU6uMqXo3MFBpMUBptGTUop2jhCf/0dqrKQIDAQAB', '7b08768a3e23917cce991d95961a0aba', 'MIIEpQIBAAKCAQEAsqGKf2Pf5HNXrbmy2lzBF7t1g8c9bGTWzs4V/s93Voe8A98ff5oeS4WSekEwR9JnqIh1rPDRbt2c2jxEwW5Qpxfez0lJPfUkmooSOfkS0CBK+AMR8jSO8nVaeQdKuxxwnSzqtHWXat8EfVFA0cdpTCaEsuuk1kyza0BBJNTYOIVbaWpLEq7a9lGvf5OE4tYC4V2wRBnABWTmuaPIdDP7HANes+NhtYao4aRXaavuYTOpuToQfytB2ZZ2EXR1bbOYXYOZ3tYDsz+chqKzWCtOBoxZAUMFSjQ1Hj0fVv9BxA4oc4XM5YQhU6uMqXo3MFBpMUBptGTUop2jhCf/0dqrKQIDAQABAoIBAQCyCiAoaDs3T18MYBHLYrdyF/1AosytxLP026NAnTesnwzeLv0OCWY4j6E3CNI+Q4/dgmUHQd78SRWUzo0Y3gK1CC2jHMBLaJJbFVKlRYNivUz7dNPmPExdjGnoR0RMykPT5Hjp9Go+spme7eVVMRPaOmn4irV4gmoxc5F6TEpVyarFFZA6N5Mcu1c6Ga8fzbJQ/l4ikLejNii43OJayB+3vpXpZMD+kaNT9ViV4vWjzw+0DTOIu/V1L2tEUgJSUqgsQlHphmA0Zqh4SvSt82JweGa9iblFaTfdrYPQoHeKpnhFrGMS+0SO0Y3cIBHghrO/wOrvLCEKOTRrfqkAMKUhAoGBANJ0gtBX6aI6lhTgb0fbLGUd1SgYRWh8YGALbcOWa58iffNXfZR8hMoFTqZD0NBS6CU9BD1oH4ShqY7jmSCMiVfFdhO+m5ZyksROHyI/wRrDAG9JxpaQs2TAqi6f1u1sMgK1PxLGXBysCVCreC9FyYvp4a3iQbPguks22oVh1n2HAoGBANlJ7uOriJA5Pw2SsNJsOaUrFzlwd2bqEJzx+ZapIXSmeQtDVf1+acynYFWVXaisfEof2v7eDCdgWiLkgClQEtpa32kL2PFFSd/X/rQE+bv7C+Z+1CnwVJXFPQa7l8o4M2x1zOqYsDgI5jCmwnxMIKK/h3N9BD58nqezmLVq7j3PAoGBALoasftPYplFOaaoeX+Pf4jDgtV78RTet9L3CE+nWvBUbhCcU77VBhAn9U7uNV+jkQotFOzAgO6mjux+s9CmtmVZhnsbWTcHhmX1t25v9nGV83J2SRioSRtVjKjCTCuUO/1NDcOqiZYJRAuENjmHex0w65L6u6M3nUYZ8c91n1lNAoGAXB/2gtyVN+CtlQmeonh0ME0GubmvRhjzFjeZhIEMyDJUm6ve520TWkuUknSTkxIeWCcNaIWoTYGSdby/gcSLWbyxgA4f2ptv6NHhEV4Sq9qSfxVCys00b5OCfjpG7tvJIbEIQmaeQAO15OavCpgUNixSN2rHw95+PcftdO1zOQ0CgYEAnk5KX3AGwtoKSzj9t6F8TF0+pWhFBGBLomqDLPnZfo0yBuT71pNTJMfAgJ8enknK9MkdkI0gHSmesb/TlBkfdMSczjbr3i1NAX6i9qvksbDDhfnBxwRqx1iCMHqBXlpQgI2u6xSy4YF9pAWmMlq79Lh2U0u3mueExF0w/WfGkeU=', '785d59f44ab0af3a821226a9c0510730', 'rsa', '', 1722588230);
INSERT INTO `user_ca` VALUES (4, '2024-08-02 16:45:38.756', '2024-08-02 16:45:38.756', NULL, 11, 'MIIBCgKCAQEArjb7qZd0SNOBmBupJeT0fYz/NUwlCoPGKheCcMkEwW3GbOA8bpvs0kL8mQaaicG1EEmcJghOeJ+7VJq56baju4LJUOnqyYHPbDa+pSGvgqTjN2W0CMF948J5+1HUWbTeJ3RCAgC0aeBJ4IGIWA7g21+utghRAJR27LagZZmesN3fMqvX/KXiy4AW6njOrkJVGGdO2Pbi/r0u6y+FYELsuo628o4V0Df/Qd/iWjTtygB4o2Y3iXk6l1RkUO8mjRY8thXIsuaRBZFXBWFYAfadCkKMs9+5vfkjDv04ikQJ1mt1VRqz6pKxDmnh0iZL1fqFW/2/YBj/Jo3oBGIufeJrwQIDAQAB', '90049c330d326572226962cf9ca647e6', 'MIIEpQIBAAKCAQEArjb7qZd0SNOBmBupJeT0fYz/NUwlCoPGKheCcMkEwW3GbOA8bpvs0kL8mQaaicG1EEmcJghOeJ+7VJq56baju4LJUOnqyYHPbDa+pSGvgqTjN2W0CMF948J5+1HUWbTeJ3RCAgC0aeBJ4IGIWA7g21+utghRAJR27LagZZmesN3fMqvX/KXiy4AW6njOrkJVGGdO2Pbi/r0u6y+FYELsuo628o4V0Df/Qd/iWjTtygB4o2Y3iXk6l1RkUO8mjRY8thXIsuaRBZFXBWFYAfadCkKMs9+5vfkjDv04ikQJ1mt1VRqz6pKxDmnh0iZL1fqFW/2/YBj/Jo3oBGIufeJrwQIDAQABAoIBACFouepT+SDeLsQUyeNSfPB20ZpO9+uEScRYrdMjOZ4+hwbqGay8ijXoUril44z3cNmmxCMnBKVtPGeFJJIFTHDxK9owga4TDswZ04WnaBkNteUxw8zKDnJXIT1o++gYH2JBkUcBC5giQwiizprUf5qSRtbx0VbjIctTWn8hbqVMFjI7VHbOWRZUJvDv36jI900zj55wWailbCnMTHXo/E8hRFdU1u+T2jhKTrlFQF+nIx+Mqj86edDBPUExQVAKmIgDRytFpKLRbcyNmCylIvW0IhUj0yWPUGNtfuJbetsgL+Ysj3035ZuKT/KOdfZ5gEe0Oz2d7pmnTGMgR94o7mECgYEA2PuVpvVM3pZw2tEAr0xAHhmmcXX0TXoqi/JBsS46EYMSkFB2aMxaaJcgMpc6plRKmmxuqtxhON7pBm1Wrllwxy7VAyDIndyHPgb6A7FwSVgHGOc/sIeXJrxZCoWeqKSb2nzdYEdmDLoSWj2fPTkisA7HVAyOkz7Qr+9Zy/32aucCgYEAzYqmmEw4Q9OrZl0efGp4MQDQP0f13ycOYiH74NxfgJy6fc0Wkgms1C3Vitx6+pvZ/La6FFZOaXS7QSjhlauAIY3bRLQ9PSG+9qbUFgVLgj46UOa9hlUhf3HO/UxMI0z2U4EhrBW+Q8p8H4RDvyddUUjiRnYQHtYu1mxxl33JhxcCgYEAg6Yph3IfoJGgiXLzumhGIr0nWoHZp0RvGhfzOaA1hTiEg+D95tnOJNVEqM1OvWo5FXaisrIhntLIiai/RetoXMGSZmJ2tij+OWkcOi9nVGi/baGo599smwj98T9fIy00jpzh6C+JwEEJL41h11CaNj3RPLKdssMS9RwA3J8K6pMCgYEAie7d1Z5ws7O8djlPWBbrCZeOKKlV6qb/FAPSUBgTTu/ehsW2u9IaVVlc0I8dSEiGhPzq6vLwM/zOHPAeUp6AogPcfcKmxft7d3KFKTHsdj7kS55M2XUPwcLSMVGeVqIzO96QhjchM8USfuP3wrPb/8x1bOuV7du2Gtczc4JKV4kCgYEAl1MFsc4qwQA6aP1W8D72iUjtGBywIyQ3r3ScGCfmxW18utUkJoNhVcMOTHpRaf/50oWh90S+q+AXP332EnxnC89EHnHoF0sK9cWMRgBDPcZvRrkS4up+YUt0A14iCi1PEHfjCb3OQwrYYMLSOGQdRWNFFyaDPWLTVlfYh/gRvOk=', 'fea15bca0c82996ef108942b6b9bc0f7', 'rsa', '', 1722588338);

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户缴费充值有关记录信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_fee
-- ----------------------------

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
  `user_level` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户级别',
  `user_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户手机号',
  `user_source_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示超级节点 2表示平台注册 3业务应用',
  `user_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '存储用户来源的详细信息',
  `user_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '1表示个人 2表示企业 3表示政府部门',
  `did` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户did',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户的详细信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_info
-- ----------------------------
INSERT INTO `user_info` VALUES (1, '2024-08-02 14:18:15.223', '2024-08-02 14:18:15.223', NULL, 'string', 'string', '0', 'string', '3', '你好小艾:873353f8be6cecb70be611019c4df7a1', '1', '072b287765436146d41d35ae93881d6d');
INSERT INTO `user_info` VALUES (2, '2024-08-02 14:20:02.276', '2024-08-02 14:20:02.276', NULL, '万宝龙', '万总', '0', '11111111', '3', '你好小艾:873353f8be6cecb70be611019c4df7a1', '1', '17ae8c405bd8434043cc61c545c84ca1');
INSERT INTO `user_info` VALUES (3, '2024-08-02 14:20:13.283', '2024-08-02 14:20:13.283', NULL, '万宝龙', '万总1', '0', '11111111', '3', '你好小艾:873353f8be6cecb70be611019c4df7a1', '1', '4b1d8321355f36d6a99186124344179a');
INSERT INTO `user_info` VALUES (4, '2024-08-02 14:20:22.578', '2024-08-02 14:20:22.578', NULL, '万宝龙', '万总2', '0', '11111111', '3', '你好小艾:873353f8be6cecb70be611019c4df7a1', '1', 'd6123b40f5483b1905637ed846f62c93');
INSERT INTO `user_info` VALUES (5, '2024-08-02 14:21:32.637', '2024-08-02 14:21:32.637', NULL, '万宝龙', '万总3', '0', '11111111', '3', '你好小艾:873353f8be6cecb70be611019c4df7a1', '1', 'c0f15a79d896751f703186a8a867a187');
INSERT INTO `user_info` VALUES (6, '2024-08-02 14:21:37.697', '2024-08-02 14:21:37.697', NULL, '万宝龙', '万总4', '0', '11111111', '3', '你好小艾:873353f8be6cecb70be611019c4df7a1', '1', '33f9bb2808f4ca72fddb08d08c84f8d2');
INSERT INTO `user_info` VALUES (7, '2024-08-02 15:03:47.924', '2024-08-02 15:03:47.924', NULL, '万宝龙', '万总5', '0', '11111111', '3', '你好小艾:873353f8be6cecb70be611019c4df7a1', '1', 'cdf5c592dc6d176d13db1f59ade13978');
INSERT INTO `user_info` VALUES (8, '2024-08-02 15:04:48.863', '2024-08-02 15:04:48.863', NULL, '万宝龙', '万总6', '0', '11111111', '3', '你好小艾:873353f8be6cecb70be611019c4df7a1', '1', 'fede4b279877530a9ab9639ddedbf8a4');
INSERT INTO `user_info` VALUES (9, '2024-08-02 16:39:48.210', '2024-08-02 16:39:48.210', NULL, '万宝龙', '万总7', '0', '11111111', '3', '你好小艾:873353f8be6cecb70be611019c4df7a1', '1', 'ea7bf7ae5c26b4faacbc873ce8ff5ee8');
INSERT INTO `user_info` VALUES (10, '2024-08-02 16:43:50.726', '2024-08-02 16:43:50.726', NULL, '万宝龙', '万总8', '0', '11111111', '3', '你好小艾:873353f8be6cecb70be611019c4df7a1', '1', '785d59f44ab0af3a821226a9c0510730');
INSERT INTO `user_info` VALUES (11, '2024-08-02 16:45:38.752', '2024-08-02 16:45:38.752', NULL, '万宝龙', '万总9', '0', '11111111', '3', '你好小艾:873353f8be6cecb70be611019c4df7a1', '1', 'fea15bca0c82996ef108942b6b9bc0f7');

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户登录信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_log
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '不同类别用户补充信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_other
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户数据空间变化记录表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_pdc
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户认证信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_prove
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
