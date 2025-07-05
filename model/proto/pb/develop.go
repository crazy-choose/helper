package pb

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ISTCNName 获取枚举值的中文名称
func ISTCNName(status InstStatus) string {
	// 获取枚举类型描述符
	enumType := status.Descriptor()
	// 根据枚举值的数值获取枚举值描述符
	valueDesc := enumType.Values().ByNumber(protoreflect.EnumNumber(status))
	if valueDesc == nil {
		return fmt.Sprintf("未知状态(%d)", status)
	}

	// 获取枚举值描述符的选项
	opts := valueDesc.Options()
	if opts == nil {
		return string(valueDesc.Name()) // 如果没有选项，则返回枚举值的名称
	}

	// 检查并获取扩展选项 `cn_name`
	if opts.ProtoReflect().IsValid() {
		// 使用全局变量 E_CnName 来获取扩展值
		if extValue := proto.GetExtension(opts, E_CnName); extValue != nil {
			if cnName, ok := extValue.(string); ok {
				return cnName
			}
		}
	}

	return string(valueDesc.Name())
}
