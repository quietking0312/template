/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

import * as $protobuf from "protobufjs/minimal";

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

    PROTO_MESSAGE.Login = (function() {

        /**
         * Properties of a Login.
         * @memberof PROTO_MESSAGE
         * @interface ILogin
         * @property {string|null} [username] Login username
         * @property {string|null} [password] Login password
         */

        /**
         * Constructs a new Login.
         * @memberof PROTO_MESSAGE
         * @classdesc Represents a Login.
         * @implements ILogin
         * @constructor
         * @param {PROTO_MESSAGE.ILogin=} [properties] Properties to set
         */
        function Login(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Login username.
         * @member {string} username
         * @memberof PROTO_MESSAGE.Login
         * @instance
         */
        Login.prototype.username = "";

        /**
         * Login password.
         * @member {string} password
         * @memberof PROTO_MESSAGE.Login
         * @instance
         */
        Login.prototype.password = "";

        /**
         * Creates a new Login instance using the specified properties.
         * @function create
         * @memberof PROTO_MESSAGE.Login
         * @static
         * @param {PROTO_MESSAGE.ILogin=} [properties] Properties to set
         * @returns {PROTO_MESSAGE.Login} Login instance
         */
        Login.create = function create(properties) {
            return new Login(properties);
        };

        /**
         * Encodes the specified Login message. Does not implicitly {@link PROTO_MESSAGE.Login.verify|verify} messages.
         * @function encode
         * @memberof PROTO_MESSAGE.Login
         * @static
         * @param {PROTO_MESSAGE.ILogin} message Login message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Login.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.username != null && Object.hasOwnProperty.call(message, "username"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.username);
            if (message.password != null && Object.hasOwnProperty.call(message, "password"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.password);
            return writer;
        };

        /**
         * Encodes the specified Login message, length delimited. Does not implicitly {@link PROTO_MESSAGE.Login.verify|verify} messages.
         * @function encodeDelimited
         * @memberof PROTO_MESSAGE.Login
         * @static
         * @param {PROTO_MESSAGE.ILogin} message Login message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Login.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Login message from the specified reader or buffer.
         * @function decode
         * @memberof PROTO_MESSAGE.Login
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {PROTO_MESSAGE.Login} Login
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Login.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.PROTO_MESSAGE.Login();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.username = reader.string();
                    break;
                case 2:
                    message.password = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Login message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof PROTO_MESSAGE.Login
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {PROTO_MESSAGE.Login} Login
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Login.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Login message.
         * @function verify
         * @memberof PROTO_MESSAGE.Login
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Login.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.username != null && message.hasOwnProperty("username"))
                if (!$util.isString(message.username))
                    return "username: string expected";
            if (message.password != null && message.hasOwnProperty("password"))
                if (!$util.isString(message.password))
                    return "password: string expected";
            return null;
        };

        /**
         * Creates a Login message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof PROTO_MESSAGE.Login
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {PROTO_MESSAGE.Login} Login
         */
        Login.fromObject = function fromObject(object) {
            if (object instanceof $root.PROTO_MESSAGE.Login)
                return object;
            let message = new $root.PROTO_MESSAGE.Login();
            if (object.username != null)
                message.username = String(object.username);
            if (object.password != null)
                message.password = String(object.password);
            return message;
        };

        /**
         * Creates a plain object from a Login message. Also converts values to other types if specified.
         * @function toObject
         * @memberof PROTO_MESSAGE.Login
         * @static
         * @param {PROTO_MESSAGE.Login} message Login
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Login.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.username = "";
                object.password = "";
            }
            if (message.username != null && message.hasOwnProperty("username"))
                object.username = message.username;
            if (message.password != null && message.hasOwnProperty("password"))
                object.password = message.password;
            return object;
        };

        /**
         * Converts this Login to JSON.
         * @function toJSON
         * @memberof PROTO_MESSAGE.Login
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Login.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Login;
    })();

    return PROTO_MESSAGE;
})();

