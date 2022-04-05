package genmysql

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/4 下午10:19
 * @Desc:
 */

type keys struct {
    NonUnique  int    `gorm:"column:Non_unique"`
    KeyName    string `gorm:"column:Key_name"`
    ColumnName string `gorm:"column:Column_name"`
    IndexType  string `gorm:"column:Index_type"`
}

type docKey struct {
    isUnique    bool     // 是否为唯一索引
    KeyName     string   // 键名
    ColumnNames []string // 键名包含的字段
}

// genColumns show full columns
type genColumns struct {
    Field   string  `gorm:"column:Field"`
    Type    string  `gorm:"column:Type"`
    Key     string  `gorm:"column:Key"`
    Desc    string  `gorm:"column:Comment"`
    Extra   string  `gorm:"column:Extra"` // 主键时，是否为自动递增。自增：Extra==auto_increment
    Null    string  `gorm:"column:Null"`
    Default *string `gorm:"column:Default"`
}

const (
    GORM_MODEL_FIELD_ID     = "id"
    GORM_MODEL_FIELD_CREATE = "created_at"
    GORM_MODEL_FIELD_UPDATE = "updated_at"
    GORM_MODEL_FIELD_DELETE = "deleted_at"
)

const (
    KEY_NAME_PRIMARY = "PRIMARY"
)

const FIELD_IS_NULL = "YES"
