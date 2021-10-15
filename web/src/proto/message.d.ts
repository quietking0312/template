import * as $protobuf from "protobufjs";
/** Namespace PROTO_MESSAGE. */
export namespace PROTO_MESSAGE {

    /** Properties of a Test. */
    interface ITest {

        /** Test data */
        data?: (PROTO_MESSAGE.Test.Iitems[]|null);
    }

    /** Represents a Test. */
    class Test implements ITest {

        /**
         * Constructs a new Test.
         * @param [properties] Properties to set
         */
        constructor(properties?: PROTO_MESSAGE.ITest);

        /** Test data. */
        public data: PROTO_MESSAGE.Test.Iitems[];

        /**
         * Creates a new Test instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Test instance
         */
        public static create(properties?: PROTO_MESSAGE.ITest): PROTO_MESSAGE.Test;

        /**
         * Encodes the specified Test message. Does not implicitly {@link PROTO_MESSAGE.Test.verify|verify} messages.
         * @param message Test message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: PROTO_MESSAGE.ITest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Test message, length delimited. Does not implicitly {@link PROTO_MESSAGE.Test.verify|verify} messages.
         * @param message Test message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: PROTO_MESSAGE.ITest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Test message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Test
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): PROTO_MESSAGE.Test;

        /**
         * Decodes a Test message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Test
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): PROTO_MESSAGE.Test;

        /**
         * Verifies a Test message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Test message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Test
         */
        public static fromObject(object: { [k: string]: any }): PROTO_MESSAGE.Test;

        /**
         * Creates a plain object from a Test message. Also converts values to other types if specified.
         * @param message Test
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: PROTO_MESSAGE.Test, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Test to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    namespace Test {

        /** Properties of an items. */
        interface Iitems {

            /** items time */
            time?: (number|Long|null);
        }

        /** Represents an items. */
        class items implements Iitems {

            /**
             * Constructs a new items.
             * @param [properties] Properties to set
             */
            constructor(properties?: PROTO_MESSAGE.Test.Iitems);

            /** items time. */
            public time: (number|Long);

            /**
             * Creates a new items instance using the specified properties.
             * @param [properties] Properties to set
             * @returns items instance
             */
            public static create(properties?: PROTO_MESSAGE.Test.Iitems): PROTO_MESSAGE.Test.items;

            /**
             * Encodes the specified items message. Does not implicitly {@link PROTO_MESSAGE.Test.items.verify|verify} messages.
             * @param message items message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: PROTO_MESSAGE.Test.Iitems, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified items message, length delimited. Does not implicitly {@link PROTO_MESSAGE.Test.items.verify|verify} messages.
             * @param message items message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: PROTO_MESSAGE.Test.Iitems, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes an items message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns items
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): PROTO_MESSAGE.Test.items;

            /**
             * Decodes an items message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns items
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): PROTO_MESSAGE.Test.items;

            /**
             * Verifies an items message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates an items message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns items
             */
            public static fromObject(object: { [k: string]: any }): PROTO_MESSAGE.Test.items;

            /**
             * Creates a plain object from an items message. Also converts values to other types if specified.
             * @param message items
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: PROTO_MESSAGE.Test.items, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this items to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }
    }
}
