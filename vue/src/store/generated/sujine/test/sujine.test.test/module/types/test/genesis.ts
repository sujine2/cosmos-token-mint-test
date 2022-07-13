/* eslint-disable */
import { Params } from "../test/params";
import { SellCoin } from "../test/sell_coin";
import { SymbolList } from "../test/symbol_list";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sujine.test.test";

/** GenesisState defines the test module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  sellCoinList: SellCoin[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  symbolListList: SymbolList[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.sellCoinList) {
      SellCoin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.symbolListList) {
      SymbolList.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.sellCoinList = [];
    message.symbolListList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.sellCoinList.push(SellCoin.decode(reader, reader.uint32()));
          break;
        case 3:
          message.symbolListList.push(
            SymbolList.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.sellCoinList = [];
    message.symbolListList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.sellCoinList !== undefined && object.sellCoinList !== null) {
      for (const e of object.sellCoinList) {
        message.sellCoinList.push(SellCoin.fromJSON(e));
      }
    }
    if (object.symbolListList !== undefined && object.symbolListList !== null) {
      for (const e of object.symbolListList) {
        message.symbolListList.push(SymbolList.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.sellCoinList) {
      obj.sellCoinList = message.sellCoinList.map((e) =>
        e ? SellCoin.toJSON(e) : undefined
      );
    } else {
      obj.sellCoinList = [];
    }
    if (message.symbolListList) {
      obj.symbolListList = message.symbolListList.map((e) =>
        e ? SymbolList.toJSON(e) : undefined
      );
    } else {
      obj.symbolListList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.sellCoinList = [];
    message.symbolListList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.sellCoinList !== undefined && object.sellCoinList !== null) {
      for (const e of object.sellCoinList) {
        message.sellCoinList.push(SellCoin.fromPartial(e));
      }
    }
    if (object.symbolListList !== undefined && object.symbolListList !== null) {
      for (const e of object.symbolListList) {
        message.symbolListList.push(SymbolList.fromPartial(e));
      }
    }
    return message;
  },
};

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
