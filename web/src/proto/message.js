/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

import * as $protobuf from "protobufjs/minimal"

// Common aliases
const $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const PROTO_MESSAGE = $root.PROTO_MESSAGE = (() => {

    /**
     * Namespace PROTO_MESSAGE.
     * @exports PROTO_MESSAGE
     * @namespace
     */
    const PROTO_MESSAGE = {};

    PROTO_MESSAGE.Test = (function() {

        /**
         * Properties of a Test.
         * @memberof PROTO_MESSAGE
         * @interface ITest
         * @property {Array.<PROTO_MESSAGE.Test.Iitems>|null} [data] Test data
         */

        /**
         * Constructs a new Test.
         * @memberof PROTO_MESSAGE
         * @classdesc Represents a Test.
         * @implements ITest
         * @constructor
         * @param {PROTO_MESSAGE.ITest=} [properties] Properties to set
         */
        function Test(properties) {
            this.data = [];
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Test data.
         * @member {Array.<PROTO_MESSAGE.Test.Iitems>} data
         * @memberof PROTO_MESSAGE.Test
         * @instance
         */
        Test.prototype.data = $util.emptyArray;

        /**
         * Creates a new Test instance using the specified properties.
         * @function create
         * @memberof PROTO_MESSAGE.Test
         * @static
         * @param {PROTO_MESSAGE.ITest=} [properties] Properties to set
         * @returns {PROTO_MESSAGE.Test} Test instance
         */
        Test.create = function create(properties) {
            return new Test(properties);
        };

        /**
         * Encodes the specified Test message. Does not implicitly {@link PROTO_MESSAGE.Test.verify|verify} messages.
         * @function encode
         * @memberof PROTO_MESSAGE.Test
         * @static
         * @param {PROTO_MESSAGE.ITest} message Test message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Test.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.data != null && message.data.length)
                for (let i = 0; i < message.data.length; ++i)
                    $root.PROTO_MESSAGE.Test.items.encode(message.data[i], writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified Test message, length delimited. Does not implicitly {@link PROTO_MESSAGE.Test.verify|verify} messages.
         * @function encodeDelimited
         * @memberof PROTO_MESSAGE.Test
         * @static
         * @param {PROTO_MESSAGE.ITest} message Test message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Test.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Test message from the specified reader or buffer.
         * @function decode
         * @memberof PROTO_MESSAGE.Test
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {PROTO_MESSAGE.Test} Test
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Test.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.PROTO_MESSAGE.Test();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    if (!(message.data && message.data.length))
                        message.data = [];
                    message.data.push($root.PROTO_MESSAGE.Test.items.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Test message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof PROTO_MESSAGE.Test
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {PROTO_MESSAGE.Test} Test
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Test.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Test message.
         * @function verify
         * @memberof PROTO_MESSAGE.Test
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Test.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.data != null && message.hasOwnProperty("data")) {
                if (!Array.isArray(message.data))
                    return "data: array expected";
                for (let i = 0; i < message.data.length; ++i) {
                    let error = $root.PROTO_MESSAGE.Test.items.verify(message.data[i]);
                    if (error)
                        return "data." + error;
                }
            }
            return null;
        };

        /**
         * Creates a Test message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof PROTO_MESSAGE.Test
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {PROTO_MESSAGE.Test} Test
         */
        Test.fromObject = function fromObject(object) {
            if (object instanceof $root.PROTO_MESSAGE.Test)
                return object;
            let message = new $root.PROTO_MESSAGE.Test();
            if (object.data) {
                if (!Array.isArray(object.data))
                    throw TypeError(".PROTO_MESSAGE.Test.data: array expected");
                message.data = [];
                for (let i = 0; i < object.data.length; ++i) {
                    if (typeof object.data[i] !== "object")
                        throw TypeError(".PROTO_MESSAGE.Test.data: object expected");
                    message.data[i] = $root.PROTO_MESSAGE.Test.items.fromObject(object.data[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a Test message. Also converts values to other types if specified.
         * @function toObject
         * @memberof PROTO_MESSAGE.Test
         * @static
         * @param {PROTO_MESSAGE.Test} message Test
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Test.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults)
                object.data = [];
            if (message.data && message.data.length) {
                object.data = [];
                for (let j = 0; j < message.data.length; ++j)
                    object.data[j] = $root.PROTO_MESSAGE.Test.items.toObject(message.data[j], options);
            }
            return object;
        };

        /**
         * Converts this Test to JSON.
         * @function toJSON
         * @memberof PROTO_MESSAGE.Test
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Test.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        Test.items = (function() {

            /**
             * Properties of an items.
             * @memberof PROTO_MESSAGE.Test
             * @interface Iitems
             * @property {number|Long|null} [time] items time
             */

            /**
             * Constructs a new items.
             * @memberof PROTO_MESSAGE.Test
             * @classdesc Represents an items.
             * @implements Iitems
             * @constructor
             * @param {PROTO_MESSAGE.Test.Iitems=} [properties] Properties to set
             */
            function items(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * items time.
             * @member {number|Long} time
             * @memberof PROTO_MESSAGE.Test.items
             * @instance
             */
            items.prototype.time = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

            /**
             * Creates a new items instance using the specified properties.
             * @function create
             * @memberof PROTO_MESSAGE.Test.items
             * @static
             * @param {PROTO_MESSAGE.Test.Iitems=} [properties] Properties to set
             * @returns {PROTO_MESSAGE.Test.items} items instance
             */
            items.create = function create(properties) {
                return new items(properties);
            };

            /**
             * Encodes the specified items message. Does not implicitly {@link PROTO_MESSAGE.Test.items.verify|verify} messages.
             * @function encode
             * @memberof PROTO_MESSAGE.Test.items
             * @static
             * @param {PROTO_MESSAGE.Test.Iitems} message items message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            items.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.time != null && Object.hasOwnProperty.call(message, "time"))
                    writer.uint32(/* id 1, wireType 0 =*/8).int64(message.time);
                return writer;
            };

            /**
             * Encodes the specified items message, length delimited. Does not implicitly {@link PROTO_MESSAGE.Test.items.verify|verify} messages.
             * @function encodeDelimited
             * @memberof PROTO_MESSAGE.Test.items
             * @static
             * @param {PROTO_MESSAGE.Test.Iitems} message items message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            items.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes an items message from the specified reader or buffer.
             * @function decode
             * @memberof PROTO_MESSAGE.Test.items
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {PROTO_MESSAGE.Test.items} items
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            items.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.PROTO_MESSAGE.Test.items();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.time = reader.int64();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes an items message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof PROTO_MESSAGE.Test.items
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {PROTO_MESSAGE.Test.items} items
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            items.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies an items message.
             * @function verify
             * @memberof PROTO_MESSAGE.Test.items
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            items.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.time != null && message.hasOwnProperty("time"))
                    if (!$util.isInteger(message.time) && !(message.time && $util.isInteger(message.time.low) && $util.isInteger(message.time.high)))
                        return "time: integer|Long expected";
                return null;
            };

            /**
             * Creates an items message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof PROTO_MESSAGE.Test.items
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {PROTO_MESSAGE.Test.items} items
             */
            items.fromObject = function fromObject(object) {
                if (object instanceof $root.PROTO_MESSAGE.Test.items)
                    return object;
                let message = new $root.PROTO_MESSAGE.Test.items();
                if (object.time != null)
                    if ($util.Long)
                        (message.time = $util.Long.fromValue(object.time)).unsigned = false;
                    else if (typeof object.time === "string")
                        message.time = parseInt(object.time, 10);
                    else if (typeof object.time === "number")
                        message.time = object.time;
                    else if (typeof object.time === "object")
                        message.time = new $util.LongBits(object.time.low >>> 0, object.time.high >>> 0).toNumber();
                return message;
            };

            /**
             * Creates a plain object from an items message. Also converts values to other types if specified.
             * @function toObject
             * @memberof PROTO_MESSAGE.Test.items
             * @static
             * @param {PROTO_MESSAGE.Test.items} message items
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            items.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults)
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, false);
                        object.time = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.time = options.longs === String ? "0" : 0;
                if (message.time != null && message.hasOwnProperty("time"))
                    if (typeof message.time === "number")
                        object.time = options.longs === String ? String(message.time) : message.time;
                    else
                        object.time = options.longs === String ? $util.Long.prototype.toString.call(message.time) : options.longs === Number ? new $util.LongBits(message.time.low >>> 0, message.time.high >>> 0).toNumber() : message.time;
                return object;
            };

            /**
             * Converts this items to JSON.
             * @function toJSON
             * @memberof PROTO_MESSAGE.Test.items
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            items.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return items;
        })();

        return Test;
    })();

    return PROTO_MESSAGE;
})();


