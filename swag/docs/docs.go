// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts/{id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Show an account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Dedicated cloud updated successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.NormalJsonResultWithoutTask"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.DedicatedCloud"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/goodbye": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goodbye"
                ],
                "responses": {
                    "200": {
                        "description": "detach volume successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.NormalJsonResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.DetachVolumeResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "go-career_swag_model.BlockDeviceMapping": {
            "type": "object",
            "properties": {
                "address_unit": {
                    "description": "挂载scsi类型卷需要自定义地址",
                    "type": "integer"
                },
                "attachment_id": {
                    "type": "string"
                },
                "boot": {
                    "$ref": "#/definitions/model.DomainDeviceBoot"
                },
                "boot_index": {
                    "description": "块设备启动顺序",
                    "type": "integer"
                },
                "cache_mode": {
                    "type": "string"
                },
                "connection_info": {
                    "description": "块设备连接信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ConnectionInfo"
                        }
                    ]
                },
                "created_at": {
                    "description": "块设备创建时间",
                    "type": "integer"
                },
                "delete_on_termination": {
                    "description": "在虚机删除时是否删除该块设备",
                    "type": "boolean"
                },
                "deleted": {
                    "description": "块设备是否被标记删除",
                    "type": "boolean"
                },
                "deleted_at": {
                    "description": "块设备删除时间",
                    "type": "integer"
                },
                "destination_type": {
                    "description": "块设备目标类型，如local、volume、blank等",
                    "type": "string"
                },
                "device_name": {
                    "description": "设备挂载的名称",
                    "type": "string"
                },
                "device_type": {
                    "description": "设备类型",
                    "type": "string"
                },
                "disk_bus": {
                    "type": "string"
                },
                "disk_io_tune": {
                    "$ref": "#/definitions/model.DiskIOTune"
                },
                "dpu_disk_queue": {
                    "description": "dpu 单盘队列数",
                    "type": "integer"
                },
                "enable_dpu": {
                    "type": "boolean"
                },
                "encryption": {
                    "$ref": "#/definitions/model.Encryption"
                },
                "guest_format": {
                    "type": "string"
                },
                "id": {
                    "description": "块设备的Id 根据类型来判断具体的id",
                    "type": "string"
                },
                "instance_id": {
                    "description": "块设备所关联的虚机Id",
                    "type": "string"
                },
                "mount_device": {
                    "type": "string"
                },
                "sgio": {
                    "type": "string"
                },
                "shareable": {
                    "description": "如果设备支持多挂载，Shareable为true",
                    "type": "boolean"
                },
                "source_type": {
                    "description": "块设备源类型，如image、volume,snapshot等",
                    "type": "string"
                },
                "updated_at": {
                    "description": "块设备更新时间",
                    "type": "integer"
                },
                "volume_id": {
                    "type": "string"
                },
                "volume_size": {
                    "description": "块设备对应卷的大小，单位Gb",
                    "type": "integer"
                },
                "volume_type": {
                    "type": "string"
                }
            }
        },
        "model.ConnectionInfo": {
            "type": "object",
            "properties": {
                "connector": {
                    "$ref": "#/definitions/model.Connector"
                },
                "data": {},
                "driver_volume_type": {
                    "description": "volume类型",
                    "type": "string"
                },
                "multiattach": {
                    "description": "卷是否支持多连接",
                    "type": "boolean"
                },
                "serial": {
                    "type": "string"
                }
            }
        },
        "model.Connector": {
            "type": "object",
            "properties": {
                "do_local_attach": {
                    "type": "boolean"
                },
                "host": {
                    "type": "string"
                },
                "initiator": {
                    "type": "string"
                },
                "multi_path": {
                    "type": "boolean"
                },
                "os_type": {
                    "type": "string"
                },
                "platform": {
                    "type": "string"
                },
                "system uuid": {
                    "description": "暂时不用，使用nvme类型后端存储时需要配置",
                    "type": "string"
                }
            }
        },
        "model.DiskIOTune": {
            "type": "object",
            "properties": {
                "read_bytes_sec": {
                    "type": "integer"
                },
                "read_iops_sec": {
                    "type": "integer"
                },
                "total_bytes_sec": {
                    "type": "integer"
                },
                "total_iops_sec": {
                    "type": "integer"
                },
                "write_bytes_sec": {
                    "type": "integer"
                },
                "write_iops_sec": {
                    "type": "integer"
                }
            }
        },
        "model.DomainDeviceBoot": {
            "type": "object",
            "properties": {
                "loadparm": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                }
            }
        },
        "model.Encryption": {
            "type": "object",
            "properties": {
                "cipher": {
                    "type": "string"
                },
                "cmkuuid": {
                    "type": "string"
                },
                "control_location": {
                    "type": "string"
                },
                "encryption_ciphertext": {
                    "type": "string"
                },
                "encryption_key_id": {
                    "type": "string"
                },
                "format": {
                    "type": "string"
                },
                "key_size": {
                    "type": "integer"
                },
                "plain_text": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "secret": {
                    "$ref": "#/definitions/model.Secret"
                },
                "secret_uuid": {
                    "type": "string"
                }
            }
        },
        "model.Secret": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "response.DedicatedCloud": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "创建时间",
                    "type": "integer"
                },
                "hosts": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "description": "专属云ID",
                    "type": "string"
                },
                "name": {
                    "description": "专属云名称",
                    "type": "string"
                },
                "project_id": {
                    "description": "租户id",
                    "type": "string"
                },
                "updated_at": {
                    "description": "更新时间",
                    "type": "integer"
                },
                "updated_by": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "response.DetachVolumeResponse": {
            "type": "object",
            "properties": {
                "block_devices_mapping": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/go-career_swag_model.BlockDeviceMapping"
                    }
                }
            }
        },
        "result.NormalJsonResult": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Http状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据内容"
                },
                "events": {},
                "message": {
                    "description": "消息内容"
                },
                "task_id": {
                    "description": "任务id",
                    "type": "string"
                }
            }
        },
        "result.NormalJsonResultWithoutTask": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Http状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据内容"
                },
                "message": {
                    "description": "消息内容"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
