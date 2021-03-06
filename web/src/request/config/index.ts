import { ConfigOptions } from "@/request/config/types";

const config: ConfigOptions = {
    base_url: {
        dev: "http://127.0.0.1:9001/api",
        pro: "/api"     // 必须加 / 否则axios 路径拼接会出错
    },

    result_code: 0,

    request_timeout: 5000,

    default_headers: 'application/json'
}

export default config
