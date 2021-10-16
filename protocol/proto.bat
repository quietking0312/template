protoc.exe --plugin=protoc-gen-go.exe --go_out=../core/protocol message.proto

:: protoc.exe --js_out=import_style=commonjs,binary:../web/src/proto message.proto
pbjs -t static-module -w commonjs --es6 -o ../web/src/proto/message.js message.proto && pbts -o ../web/src/proto/message.d.ts ../web/src/proto/message.js
:: protoc.exe --plugin=../web/node_modules/.bin/protoc-gen-ts --ts_out ../web/src/proto --ts_opt long_type_string --ts_opt disable_service_client message.proto
