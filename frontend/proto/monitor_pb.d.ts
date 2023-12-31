// @generated by protoc-gen-es v1.2.1
// @generated from file monitor.proto (package xmon, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message xmon.ListNodesRequest
 */
export declare class ListNodesRequest extends Message<ListNodesRequest> {
  constructor(data?: PartialMessage<ListNodesRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "xmon.ListNodesRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListNodesRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListNodesRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListNodesRequest;

  static equals(a: ListNodesRequest | PlainMessage<ListNodesRequest> | undefined, b: ListNodesRequest | PlainMessage<ListNodesRequest> | undefined): boolean;
}

/**
 * @generated from message xmon.ListNodesResponse
 */
export declare class ListNodesResponse extends Message<ListNodesResponse> {
  /**
   * @generated from field: repeated xmon.ListNodesResponse.Node nodes = 1;
   */
  nodes: ListNodesResponse_Node[];

  constructor(data?: PartialMessage<ListNodesResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "xmon.ListNodesResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListNodesResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListNodesResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListNodesResponse;

  static equals(a: ListNodesResponse | PlainMessage<ListNodesResponse> | undefined, b: ListNodesResponse | PlainMessage<ListNodesResponse> | undefined): boolean;
}

/**
 * @generated from message xmon.ListNodesResponse.Client
 */
export declare class ListNodesResponse_Client extends Message<ListNodesResponse_Client> {
  /**
   * @generated from field: string version = 2;
   */
  version: string;

  /**
   * @generated from field: xmon.ListNodesResponse.Client.Status status = 3;
   */
  status: ListNodesResponse_Client_Status;

  /**
   * @generated from field: google.protobuf.Timestamp lastPing = 4;
   */
  lastPing?: Timestamp;

  constructor(data?: PartialMessage<ListNodesResponse_Client>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "xmon.ListNodesResponse.Client";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListNodesResponse_Client;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListNodesResponse_Client;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListNodesResponse_Client;

  static equals(a: ListNodesResponse_Client | PlainMessage<ListNodesResponse_Client> | undefined, b: ListNodesResponse_Client | PlainMessage<ListNodesResponse_Client> | undefined): boolean;
}

/**
 * @generated from enum xmon.ListNodesResponse.Client.Status
 */
export declare enum ListNodesResponse_Client_Status {
  /**
   * @generated from enum value: HEALTHY = 0;
   */
  HEALTHY = 0,

  /**
   * @generated from enum value: UNHEALTHY = 1;
   */
  UNHEALTHY = 1,
}

/**
 * @generated from message xmon.ListNodesResponse.Node
 */
export declare class ListNodesResponse_Node extends Message<ListNodesResponse_Node> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: xmon.ListNodesResponse.Client executionClient = 2;
   */
  executionClient?: ListNodesResponse_Client;

  /**
   * @generated from field: xmon.ListNodesResponse.Client consensusClient = 3;
   */
  consensusClient?: ListNodesResponse_Client;

  constructor(data?: PartialMessage<ListNodesResponse_Node>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "xmon.ListNodesResponse.Node";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListNodesResponse_Node;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListNodesResponse_Node;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListNodesResponse_Node;

  static equals(a: ListNodesResponse_Node | PlainMessage<ListNodesResponse_Node> | undefined, b: ListNodesResponse_Node | PlainMessage<ListNodesResponse_Node> | undefined): boolean;
}

/**
 * @generated from message xmon.ListBlocksRequest
 */
export declare class ListBlocksRequest extends Message<ListBlocksRequest> {
  constructor(data?: PartialMessage<ListBlocksRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "xmon.ListBlocksRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListBlocksRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListBlocksRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListBlocksRequest;

  static equals(a: ListBlocksRequest | PlainMessage<ListBlocksRequest> | undefined, b: ListBlocksRequest | PlainMessage<ListBlocksRequest> | undefined): boolean;
}

/**
 * @generated from message xmon.ListBlocksResponse
 */
export declare class ListBlocksResponse extends Message<ListBlocksResponse> {
  /**
   * @generated from field: repeated xmon.ListBlocksResponse.BlockData blocks = 1;
   */
  blocks: ListBlocksResponse_BlockData[];

  constructor(data?: PartialMessage<ListBlocksResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "xmon.ListBlocksResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListBlocksResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListBlocksResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListBlocksResponse;

  static equals(a: ListBlocksResponse | PlainMessage<ListBlocksResponse> | undefined, b: ListBlocksResponse | PlainMessage<ListBlocksResponse> | undefined): boolean;
}

/**
 * @generated from message xmon.ListBlocksResponse.BlockData
 */
export declare class ListBlocksResponse_BlockData extends Message<ListBlocksResponse_BlockData> {
  /**
   * @generated from field: int64 blockNumber = 1;
   */
  blockNumber: bigint;

  /**
   * @generated from field: map<string, string> hashes = 2;
   */
  hashes: { [key: string]: string };

  /**
   * @generated from field: bool isSynced = 3;
   */
  isSynced: boolean;

  constructor(data?: PartialMessage<ListBlocksResponse_BlockData>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "xmon.ListBlocksResponse.BlockData";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListBlocksResponse_BlockData;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListBlocksResponse_BlockData;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListBlocksResponse_BlockData;

  static equals(a: ListBlocksResponse_BlockData | PlainMessage<ListBlocksResponse_BlockData> | undefined, b: ListBlocksResponse_BlockData | PlainMessage<ListBlocksResponse_BlockData> | undefined): boolean;
}

/**
 * @generated from message xmon.ListSlotsRequest
 */
export declare class ListSlotsRequest extends Message<ListSlotsRequest> {
  constructor(data?: PartialMessage<ListSlotsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "xmon.ListSlotsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSlotsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSlotsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSlotsRequest;

  static equals(a: ListSlotsRequest | PlainMessage<ListSlotsRequest> | undefined, b: ListSlotsRequest | PlainMessage<ListSlotsRequest> | undefined): boolean;
}

/**
 * @generated from message xmon.ListSlotsResponse
 */
export declare class ListSlotsResponse extends Message<ListSlotsResponse> {
  /**
   * node -> []slot
   *
   * @generated from field: map<string, xmon.ListSlotsResponse.SlotList> slots = 1;
   */
  slots: { [key: string]: ListSlotsResponse_SlotList };

  constructor(data?: PartialMessage<ListSlotsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "xmon.ListSlotsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSlotsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSlotsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSlotsResponse;

  static equals(a: ListSlotsResponse | PlainMessage<ListSlotsResponse> | undefined, b: ListSlotsResponse | PlainMessage<ListSlotsResponse> | undefined): boolean;
}

/**
 * @generated from message xmon.ListSlotsResponse.Slot
 */
export declare class ListSlotsResponse_Slot extends Message<ListSlotsResponse_Slot> {
  /**
   * @generated from field: int64 number = 1;
   */
  number: bigint;

  /**
   * @generated from field: string hash = 2;
   */
  hash: string;

  constructor(data?: PartialMessage<ListSlotsResponse_Slot>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "xmon.ListSlotsResponse.Slot";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSlotsResponse_Slot;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSlotsResponse_Slot;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSlotsResponse_Slot;

  static equals(a: ListSlotsResponse_Slot | PlainMessage<ListSlotsResponse_Slot> | undefined, b: ListSlotsResponse_Slot | PlainMessage<ListSlotsResponse_Slot> | undefined): boolean;
}

/**
 * @generated from message xmon.ListSlotsResponse.SlotList
 */
export declare class ListSlotsResponse_SlotList extends Message<ListSlotsResponse_SlotList> {
  /**
   * @generated from field: repeated xmon.ListSlotsResponse.Slot slots = 1;
   */
  slots: ListSlotsResponse_Slot[];

  constructor(data?: PartialMessage<ListSlotsResponse_SlotList>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "xmon.ListSlotsResponse.SlotList";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSlotsResponse_SlotList;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSlotsResponse_SlotList;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSlotsResponse_SlotList;

  static equals(a: ListSlotsResponse_SlotList | PlainMessage<ListSlotsResponse_SlotList> | undefined, b: ListSlotsResponse_SlotList | PlainMessage<ListSlotsResponse_SlotList> | undefined): boolean;
}

