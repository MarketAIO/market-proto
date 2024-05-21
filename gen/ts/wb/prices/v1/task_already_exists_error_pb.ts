//
//API цен и скидок
//
//С помощью этих методов можно устанавливать цены и скидки. Максимум — 10 запросов за 6 секунд суммарно для всех методов. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/prices/v1/task_already_exists_error.proto (package wb.prices.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { TaskAlreadyExistsErrorData } from "./task_already_exists_error_data_pb.js";

/**
 * @generated from message wb.prices.v1.TaskAlreadyExistsError
 */
export class TaskAlreadyExistsError extends Message<TaskAlreadyExistsError> {
  /**
   * @generated from field: wb.prices.v1.TaskAlreadyExistsErrorData data = 1;
   */
  data?: TaskAlreadyExistsErrorData;

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

  constructor(data?: PartialMessage<TaskAlreadyExistsError>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.prices.v1.TaskAlreadyExistsError";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: TaskAlreadyExistsErrorData },
    { no: 2, name: "error", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "errorText", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskAlreadyExistsError {
    return new TaskAlreadyExistsError().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskAlreadyExistsError {
    return new TaskAlreadyExistsError().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskAlreadyExistsError {
    return new TaskAlreadyExistsError().fromJsonString(jsonString, options);
  }

  static equals(a: TaskAlreadyExistsError | PlainMessage<TaskAlreadyExistsError> | undefined, b: TaskAlreadyExistsError | PlainMessage<TaskAlreadyExistsError> | undefined): boolean {
    return proto3.util.equals(TaskAlreadyExistsError, a, b);
  }
}
