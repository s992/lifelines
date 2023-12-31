// @generated by protoc-gen-connect-query v0.6.0
// @generated from file lifelines/v1/lifelines.proto (package lifelines.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateTagRequest, CreateTagResponse, ListTagsRequest, ListTagsResponse, SearchTagsRequest, SearchTagsResponse } from "./lifelines_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { createQueryService, createUnaryHooks } from "@connectrpc/connect-query";

export const typeName = "lifelines.v1.TagService";

/**
 * @generated from service lifelines.v1.TagService
 */
export const TagService = {
  typeName: "lifelines.v1.TagService",
  methods: {
    /**
     * @generated from rpc lifelines.v1.TagService.ListTags
     */
    listTags: {
      name: "ListTags",
      I: ListTagsRequest,
      O: ListTagsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc lifelines.v1.TagService.SearchTags
     */
    searchTags: {
      name: "SearchTags",
      I: SearchTagsRequest,
      O: SearchTagsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc lifelines.v1.TagService.CreateTag
     */
    createTag: {
      name: "CreateTag",
      I: CreateTagRequest,
      O: CreateTagResponse,
      kind: MethodKind.Unary,
    },
  }
};

const $queryService = createQueryService({  service: TagService,});

/**
 * @generated from rpc lifelines.v1.TagService.ListTags
 */
export const listTags = {   ...$queryService.listTags,  ...createUnaryHooks($queryService.listTags)};

/**
 * @generated from rpc lifelines.v1.TagService.SearchTags
 */
export const searchTags = {   ...$queryService.searchTags,  ...createUnaryHooks($queryService.searchTags)};

/**
 * @generated from rpc lifelines.v1.TagService.CreateTag
 */
export const createTag = {   ...$queryService.createTag,  ...createUnaryHooks($queryService.createTag)};
