/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../test/params";
import { SellCoin } from "../test/sell_coin";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { SymbolList } from "../test/symbol_list";

export const protobufPackage = "sujine.test.test";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetInfoRequest {
  creator: string;
  denom: string;
}

export interface QueryGetInfoResponse {
  denom: string;
  info: string;
}

export interface QueryGetSellCoinRequest {
  index: string;
}

export interface QueryGetSellCoinResponse {
  sellCoin: SellCoin | undefined;
}

export interface QueryAllSellCoinRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSellCoinResponse {
  sellCoin: SellCoin[];
  pagination: PageResponse | undefined;
}

export interface QueryGetSymbolListRequest {
  index: string;
}

export interface QueryGetSymbolListResponse {
  symbolList: SymbolList | undefined;
}

export interface QueryAllSymbolListRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSymbolListResponse {
  symbolList: SymbolList[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetInfoRequest: object = { creator: "", denom: "" };

export const QueryGetInfoRequest = {
  encode(
    message: QueryGetInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.denom !== "") {
      writer.uint32(18).string(message.denom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetInfoRequest } as QueryGetInfoRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetInfoRequest {
    const message = { ...baseQueryGetInfoRequest } as QueryGetInfoRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    return message;
  },

  toJSON(message: QueryGetInfoRequest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.denom !== undefined && (obj.denom = message.denom);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetInfoRequest>): QueryGetInfoRequest {
    const message = { ...baseQueryGetInfoRequest } as QueryGetInfoRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    return message;
  },
};

const baseQueryGetInfoResponse: object = { denom: "", info: "" };

export const QueryGetInfoResponse = {
  encode(
    message: QueryGetInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    if (message.info !== "") {
      writer.uint32(18).string(message.info);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetInfoResponse } as QueryGetInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;
        case 2:
          message.info = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetInfoResponse {
    const message = { ...baseQueryGetInfoResponse } as QueryGetInfoResponse;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = String(object.info);
    } else {
      message.info = "";
    }
    return message;
  },

  toJSON(message: QueryGetInfoResponse): unknown {
    const obj: any = {};
    message.denom !== undefined && (obj.denom = message.denom);
    message.info !== undefined && (obj.info = message.info);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetInfoResponse>): QueryGetInfoResponse {
    const message = { ...baseQueryGetInfoResponse } as QueryGetInfoResponse;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = object.info;
    } else {
      message.info = "";
    }
    return message;
  },
};

const baseQueryGetSellCoinRequest: object = { index: "" };

export const QueryGetSellCoinRequest = {
  encode(
    message: QueryGetSellCoinRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetSellCoinRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSellCoinRequest,
    } as QueryGetSellCoinRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSellCoinRequest {
    const message = {
      ...baseQueryGetSellCoinRequest,
    } as QueryGetSellCoinRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetSellCoinRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSellCoinRequest>
  ): QueryGetSellCoinRequest {
    const message = {
      ...baseQueryGetSellCoinRequest,
    } as QueryGetSellCoinRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetSellCoinResponse: object = {};

export const QueryGetSellCoinResponse = {
  encode(
    message: QueryGetSellCoinResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.sellCoin !== undefined) {
      SellCoin.encode(message.sellCoin, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSellCoinResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSellCoinResponse,
    } as QueryGetSellCoinResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sellCoin = SellCoin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSellCoinResponse {
    const message = {
      ...baseQueryGetSellCoinResponse,
    } as QueryGetSellCoinResponse;
    if (object.sellCoin !== undefined && object.sellCoin !== null) {
      message.sellCoin = SellCoin.fromJSON(object.sellCoin);
    } else {
      message.sellCoin = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSellCoinResponse): unknown {
    const obj: any = {};
    message.sellCoin !== undefined &&
      (obj.sellCoin = message.sellCoin
        ? SellCoin.toJSON(message.sellCoin)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSellCoinResponse>
  ): QueryGetSellCoinResponse {
    const message = {
      ...baseQueryGetSellCoinResponse,
    } as QueryGetSellCoinResponse;
    if (object.sellCoin !== undefined && object.sellCoin !== null) {
      message.sellCoin = SellCoin.fromPartial(object.sellCoin);
    } else {
      message.sellCoin = undefined;
    }
    return message;
  },
};

const baseQueryAllSellCoinRequest: object = {};

export const QueryAllSellCoinRequest = {
  encode(
    message: QueryAllSellCoinRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllSellCoinRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSellCoinRequest,
    } as QueryAllSellCoinRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSellCoinRequest {
    const message = {
      ...baseQueryAllSellCoinRequest,
    } as QueryAllSellCoinRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSellCoinRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSellCoinRequest>
  ): QueryAllSellCoinRequest {
    const message = {
      ...baseQueryAllSellCoinRequest,
    } as QueryAllSellCoinRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllSellCoinResponse: object = {};

export const QueryAllSellCoinResponse = {
  encode(
    message: QueryAllSellCoinResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.sellCoin) {
      SellCoin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSellCoinResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSellCoinResponse,
    } as QueryAllSellCoinResponse;
    message.sellCoin = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sellCoin.push(SellCoin.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSellCoinResponse {
    const message = {
      ...baseQueryAllSellCoinResponse,
    } as QueryAllSellCoinResponse;
    message.sellCoin = [];
    if (object.sellCoin !== undefined && object.sellCoin !== null) {
      for (const e of object.sellCoin) {
        message.sellCoin.push(SellCoin.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSellCoinResponse): unknown {
    const obj: any = {};
    if (message.sellCoin) {
      obj.sellCoin = message.sellCoin.map((e) =>
        e ? SellCoin.toJSON(e) : undefined
      );
    } else {
      obj.sellCoin = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSellCoinResponse>
  ): QueryAllSellCoinResponse {
    const message = {
      ...baseQueryAllSellCoinResponse,
    } as QueryAllSellCoinResponse;
    message.sellCoin = [];
    if (object.sellCoin !== undefined && object.sellCoin !== null) {
      for (const e of object.sellCoin) {
        message.sellCoin.push(SellCoin.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetSymbolListRequest: object = { index: "" };

export const QueryGetSymbolListRequest = {
  encode(
    message: QueryGetSymbolListRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSymbolListRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSymbolListRequest,
    } as QueryGetSymbolListRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSymbolListRequest {
    const message = {
      ...baseQueryGetSymbolListRequest,
    } as QueryGetSymbolListRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetSymbolListRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSymbolListRequest>
  ): QueryGetSymbolListRequest {
    const message = {
      ...baseQueryGetSymbolListRequest,
    } as QueryGetSymbolListRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetSymbolListResponse: object = {};

export const QueryGetSymbolListResponse = {
  encode(
    message: QueryGetSymbolListResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.symbolList !== undefined) {
      SymbolList.encode(message.symbolList, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSymbolListResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSymbolListResponse,
    } as QueryGetSymbolListResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbolList = SymbolList.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSymbolListResponse {
    const message = {
      ...baseQueryGetSymbolListResponse,
    } as QueryGetSymbolListResponse;
    if (object.symbolList !== undefined && object.symbolList !== null) {
      message.symbolList = SymbolList.fromJSON(object.symbolList);
    } else {
      message.symbolList = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSymbolListResponse): unknown {
    const obj: any = {};
    message.symbolList !== undefined &&
      (obj.symbolList = message.symbolList
        ? SymbolList.toJSON(message.symbolList)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSymbolListResponse>
  ): QueryGetSymbolListResponse {
    const message = {
      ...baseQueryGetSymbolListResponse,
    } as QueryGetSymbolListResponse;
    if (object.symbolList !== undefined && object.symbolList !== null) {
      message.symbolList = SymbolList.fromPartial(object.symbolList);
    } else {
      message.symbolList = undefined;
    }
    return message;
  },
};

const baseQueryAllSymbolListRequest: object = {};

export const QueryAllSymbolListRequest = {
  encode(
    message: QueryAllSymbolListRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSymbolListRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSymbolListRequest,
    } as QueryAllSymbolListRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSymbolListRequest {
    const message = {
      ...baseQueryAllSymbolListRequest,
    } as QueryAllSymbolListRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSymbolListRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSymbolListRequest>
  ): QueryAllSymbolListRequest {
    const message = {
      ...baseQueryAllSymbolListRequest,
    } as QueryAllSymbolListRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllSymbolListResponse: object = {};

export const QueryAllSymbolListResponse = {
  encode(
    message: QueryAllSymbolListResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.symbolList) {
      SymbolList.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSymbolListResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSymbolListResponse,
    } as QueryAllSymbolListResponse;
    message.symbolList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbolList.push(SymbolList.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSymbolListResponse {
    const message = {
      ...baseQueryAllSymbolListResponse,
    } as QueryAllSymbolListResponse;
    message.symbolList = [];
    if (object.symbolList !== undefined && object.symbolList !== null) {
      for (const e of object.symbolList) {
        message.symbolList.push(SymbolList.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSymbolListResponse): unknown {
    const obj: any = {};
    if (message.symbolList) {
      obj.symbolList = message.symbolList.map((e) =>
        e ? SymbolList.toJSON(e) : undefined
      );
    } else {
      obj.symbolList = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSymbolListResponse>
  ): QueryAllSymbolListResponse {
    const message = {
      ...baseQueryAllSymbolListResponse,
    } as QueryAllSymbolListResponse;
    message.symbolList = [];
    if (object.symbolList !== undefined && object.symbolList !== null) {
      for (const e of object.symbolList) {
        message.symbolList.push(SymbolList.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of GetInfo items. */
  GetInfo(request: QueryGetInfoRequest): Promise<QueryGetInfoResponse>;
  /** Queries a SellCoin by index. */
  SellCoin(request: QueryGetSellCoinRequest): Promise<QueryGetSellCoinResponse>;
  /** Queries a list of SellCoin items. */
  SellCoinAll(
    request: QueryAllSellCoinRequest
  ): Promise<QueryAllSellCoinResponse>;
  /** Queries a SymbolList by index. */
  SymbolList(
    request: QueryGetSymbolListRequest
  ): Promise<QueryGetSymbolListResponse>;
  /** Queries a list of SymbolList items. */
  SymbolListAll(
    request: QueryAllSymbolListRequest
  ): Promise<QueryAllSymbolListResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("sujine.test.test.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  GetInfo(request: QueryGetInfoRequest): Promise<QueryGetInfoResponse> {
    const data = QueryGetInfoRequest.encode(request).finish();
    const promise = this.rpc.request("sujine.test.test.Query", "GetInfo", data);
    return promise.then((data) =>
      QueryGetInfoResponse.decode(new Reader(data))
    );
  }

  SellCoin(
    request: QueryGetSellCoinRequest
  ): Promise<QueryGetSellCoinResponse> {
    const data = QueryGetSellCoinRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sujine.test.test.Query",
      "SellCoin",
      data
    );
    return promise.then((data) =>
      QueryGetSellCoinResponse.decode(new Reader(data))
    );
  }

  SellCoinAll(
    request: QueryAllSellCoinRequest
  ): Promise<QueryAllSellCoinResponse> {
    const data = QueryAllSellCoinRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sujine.test.test.Query",
      "SellCoinAll",
      data
    );
    return promise.then((data) =>
      QueryAllSellCoinResponse.decode(new Reader(data))
    );
  }

  SymbolList(
    request: QueryGetSymbolListRequest
  ): Promise<QueryGetSymbolListResponse> {
    const data = QueryGetSymbolListRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sujine.test.test.Query",
      "SymbolList",
      data
    );
    return promise.then((data) =>
      QueryGetSymbolListResponse.decode(new Reader(data))
    );
  }

  SymbolListAll(
    request: QueryAllSymbolListRequest
  ): Promise<QueryAllSymbolListResponse> {
    const data = QueryAllSymbolListRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sujine.test.test.Query",
      "SymbolListAll",
      data
    );
    return promise.then((data) =>
      QueryAllSymbolListResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
