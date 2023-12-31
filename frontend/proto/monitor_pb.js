// @generated by protoc-gen-es v1.2.1
// @generated from file monitor.proto (package xmon, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message xmon.ListNodesRequest
 */
export const ListNodesRequest = proto3.makeMessageType(
  "xmon.ListNodesRequest",
  [],
);

/**
 * @generated from message xmon.ListNodesResponse
 */
export const ListNodesResponse = proto3.makeMessageType(
  "xmon.ListNodesResponse",
  () => [
    { no: 1, name: "nodes", kind: "message", T: ListNodesResponse_Node, repeated: true },
  ],
);

/**
 * @generated from message xmon.ListNodesResponse.Client
 */
export const ListNodesResponse_Client = proto3.makeMessageType(
  "xmon.ListNodesResponse.Client",
  () => [
    { no: 2, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "status", kind: "enum", T: proto3.getEnumType(ListNodesResponse_Client_Status) },
    { no: 4, name: "lastPing", kind: "message", T: Timestamp },
  ],
  {localName: "ListNodesResponse_Client"},
);

/**
 * @generated from enum xmon.ListNodesResponse.Client.Status
 */
export const ListNodesResponse_Client_Status = proto3.makeEnum(
  "xmon.ListNodesResponse.Client.Status",
  [
    {no: 0, name: "HEALTHY"},
    {no: 1, name: "UNHEALTHY"},
  ],
);

/**
 * @generated from message xmon.ListNodesResponse.Node
 */
export const ListNodesResponse_Node = proto3.makeMessageType(
  "xmon.ListNodesResponse.Node",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "executionClient", kind: "message", T: ListNodesResponse_Client },
    { no: 3, name: "consensusClient", kind: "message", T: ListNodesResponse_Client },
  ],
  {localName: "ListNodesResponse_Node"},
);

/**
 * @generated from message xmon.ListBlocksRequest
 */
export const ListBlocksRequest = proto3.makeMessageType(
  "xmon.ListBlocksRequest",
  [],
);

/**
 * @generated from message xmon.ListBlocksResponse
 */
export const ListBlocksResponse = proto3.makeMessageType(
  "xmon.ListBlocksResponse",
  () => [
    { no: 1, name: "blocks", kind: "message", T: ListBlocksResponse_BlockData, repeated: true },
  ],
);

/**
 * @generated from message xmon.ListBlocksResponse.BlockData
 */
export const ListBlocksResponse_BlockData = proto3.makeMessageType(
  "xmon.ListBlocksResponse.BlockData",
  () => [
    { no: 1, name: "blockNumber", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "hashes", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 3, name: "isSynced", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
  {localName: "ListBlocksResponse_BlockData"},
);

/**
 * @generated from message xmon.ListSlotsRequest
 */
export const ListSlotsRequest = proto3.makeMessageType(
  "xmon.ListSlotsRequest",
  [],
);

/**
 * @generated from message xmon.ListSlotsResponse
 */
export const ListSlotsResponse = proto3.makeMessageType(
  "xmon.ListSlotsResponse",
  () => [
    { no: 1, name: "slots", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: ListSlotsResponse_SlotList} },
  ],
);

/**
 * @generated from message xmon.ListSlotsResponse.Slot
 */
export const ListSlotsResponse_Slot = proto3.makeMessageType(
  "xmon.ListSlotsResponse.Slot",
  () => [
    { no: 1, name: "number", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "hash", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
  {localName: "ListSlotsResponse_Slot"},
);

/**
 * @generated from message xmon.ListSlotsResponse.SlotList
 */
export const ListSlotsResponse_SlotList = proto3.makeMessageType(
  "xmon.ListSlotsResponse.SlotList",
  () => [
    { no: 1, name: "slots", kind: "message", T: ListSlotsResponse_Slot, repeated: true },
  ],
  {localName: "ListSlotsResponse_SlotList"},
);

