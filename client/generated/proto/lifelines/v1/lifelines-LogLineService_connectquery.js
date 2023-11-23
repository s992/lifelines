// @generated by protoc-gen-connect-query v0.6.0
// @generated from file lifelines/v1/lifelines.proto (package lifelines.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateLogLineRequest, CreateLogLineResponse, ListLogLinesRequest, ListLogLinesResponse } from "./lifelines_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { createQueryService, createUnaryHooks } from "@connectrpc/connect-query";

export const typeName = "lifelines.v1.LogLineService";

/**
 * @generated from service lifelines.v1.LogLineService
 */
export const LogLineService = {
  typeName: "lifelines.v1.LogLineService",
  methods: {
    /**
     * @generated from rpc lifelines.v1.LogLineService.ListLogLines
     */
    listLogLines: {
      name: "ListLogLines",
      I: ListLogLinesRequest,
      O: ListLogLinesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc lifelines.v1.LogLineService.CreateLogLine
     */
    createLogLine: {
      name: "CreateLogLine",
      I: CreateLogLineRequest,
      O: CreateLogLineResponse,
      kind: MethodKind.Unary,
    },
  }
};

const $queryService = createQueryService({  service: LogLineService,});

/**
 * @generated from rpc lifelines.v1.LogLineService.ListLogLines
 */
export const listLogLines = {   ...$queryService.listLogLines,  ...createUnaryHooks($queryService.listLogLines)};

/**
 * @generated from rpc lifelines.v1.LogLineService.CreateLogLine
 */
export const createLogLine = {   ...$queryService.createLogLine,  ...createUnaryHooks($queryService.createLogLine)};