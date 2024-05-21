//
//API цен и скидок
//
//С помощью этих методов можно устанавливать цены и скидки. Максимум — 10 запросов за 6 секунд суммарно для всех методов. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/prices/v1/task_created.proto (package wb.prices.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { TaskCreatedData } from "./task_created_data_pb.js";

/**
 * @generated from message wb.prices.v1.TaskCreated
 */
export class TaskCreated extends Message<TaskCreated> {
  /**
   * @generated from field: wb.prices.v1.TaskCreatedData data = 1;
   */
  data?: TaskCreatedData;

  /**
   * Флаг ошибки
   *
   * @generated from field: bool error = 2;
   */
  error = false;

  /**
   * Текст ошибки
   *
   * @generated from field: string errorText = 3;
   */
  errorText = "";

  constructor(data?: PartialMessage<TaskCreated>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.prices.v1.TaskCreated";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: TaskCreatedData },
    { no: 2, name: "error", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "errorText", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskCreated {
    return new TaskCreated().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskCreated {
    return new TaskCreated().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskCreated {
    return new TaskCreated().fromJsonString(jsonString, options);
  }

  static equals(a: TaskCreated | PlainMessage<TaskCreated> | undefined, b: TaskCreated | PlainMessage<TaskCreated> | undefined): boolean {
    return proto3.util.equals(TaskCreated, a, b);
  }
}

